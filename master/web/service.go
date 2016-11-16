/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package web

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/h2oai/steam/bindings"
	"github.com/h2oai/steam/lib/fs"
	"github.com/h2oai/steam/lib/svc"
	"github.com/h2oai/steam/lib/yarn"
	"github.com/h2oai/steam/master/auth"
	"github.com/h2oai/steam/master/az"
	"github.com/h2oai/steam/master/data"
	"github.com/h2oai/steam/srv/compiler"
	"github.com/h2oai/steam/srv/h2ov3"
	"github.com/h2oai/steam/srv/web"
	"github.com/pkg/errors"
	"github.com/h2oai/steam/lib/haproxy"
)

type Service struct {
	workingDir                string
	ds                        *data.Datastore
	compilationServiceAddress string
	scoringServiceAddress     string
	clusterProxyAddress       string
	scoringServicePortMin     int
	scoringServicePortMax     int
	kerberosEnabled           bool
}

func NewService(
	workingDir string,
	ds *data.Datastore,
	compilationServiceAddress, scoringServiceAddress, clusterProxyAddress string,
	scoringServicePortsRange [2]int,
	kerberos bool,
) *Service {
	return &Service{
		workingDir,
		ds,
		compilationServiceAddress, scoringServiceAddress, clusterProxyAddress,
		scoringServicePortsRange[0], scoringServicePortsRange[1],
		kerberos,
	}
}

func toTimestamp(t time.Time) int64 {
	return t.UTC().Unix()
}

func now() int64 {
	return toTimestamp(time.Now())
}

func (s *Service) PingServer(pz az.Principal, status string) (string, error) {
	return status, nil
}

func (s *Service) GetConfig(pz az.Principal) (*web.Config, error) {
	return &web.Config{
		KerberosEnabled:     s.kerberosEnabled,
		ClusterProxyAddress: s.clusterProxyAddress,
	}, nil
}

func (s *Service) RegisterCluster(pz az.Principal, address string) (int64, error) {

	if err := pz.CheckPermission(s.ds.Permissions.ManageCluster); err != nil {
		return 0, err
	}

	h := h2ov3.NewClient(address)
	cloud, err := h.GetCloudStatus()
	if err != nil {
		return 0, fmt.Errorf("Could not communicate with an h2o cluster at %s", address)
	}

	_, ok, err := s.ds.ReadClusterByAddress(pz, address)
	if err != nil {
		return 0, err
	}

	if ok {
		return 0, fmt.Errorf("A cluster with the address %s is already registered", address)
	}

	clusterId, err := s.ds.CreateExternalCluster(pz, cloud.CloudName, address, data.StartedState)
	if err != nil {
		return 0, err
	}

	return clusterId, nil
}

func (s *Service) UnregisterCluster(pz az.Principal, clusterId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageCluster); err != nil {
		return err
	}

	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return err
	}

	if cluster.TypeId != s.ds.ClusterTypes.External {
		return fmt.Errorf("Cannot unregister internal clusters.")
	}

	if err := s.ds.DeleteCluster(pz, clusterId); err != nil {
		return err
	}

	return nil
}

func (s *Service) StartClusterOnYarn(pz az.Principal, clusterName string, engineId int64, size int, memory string, secure bool, keytab string) (int64, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ManageCluster); err != nil {
		return 0, err
	}

	// Cluster should have a unique name
	_, ok, err := s.ds.ReadClusterByName(pz, clusterName)
	if err != nil {
		return 0, err
	}
	if ok {
		return 0, fmt.Errorf("A cluster with the name %s already exists.", clusterName)
	}

	identity, err := s.ds.ReadIdentity(pz, pz.Id())
	if err != nil {
		return 0, err
	}

	engine, err := s.ds.ReadEngine(pz, engineId)
	if err != nil {
		return 0, err
	}

	// FIXME check if file exists
	keytabPath := path.Join(s.workingDir, fs.KTDir, keytab)

	appId, address, out, token, err := yarn.StartCloud(size, s.kerberosEnabled, memory,
		clusterName, engine.Location, identity.Name, keytabPath, secure)
	if err != nil {
		return 0, err
	}

	yarnCluster := data.YarnCluster{
		0,
		engineId,
		int64(size),
		appId,
		memory,
		identity.Name,
		out,
	}

	clusterId, err := s.ds.CreateYarnCluster(pz, clusterName, address, token, data.StartedState, yarnCluster)
	if err != nil {
		log.Println("Failed to create yarn cluster.")
		return 0, err
	}

	if err := reloadProxyConf(s, pz, identity.Name); err != nil {
		return 0, err
	}

	return clusterId, nil
}

func (s *Service) StopClusterOnYarn(pz az.Principal, clusterId int64, keytab string) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageCluster); err != nil {
		return err
	}

	// Cluster should exist
	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return errors.Wrap(err, "failed reading cluster")
	}

	if cluster.TypeId != s.ds.ClusterTypes.Yarn {
		return fmt.Errorf("Cluster %d was not started through YARN", clusterId)
	}

	// Bail out if already stopped
	if cluster.State == data.StoppedState {
		return fmt.Errorf("Cluster %d is already stopped", clusterId)
	}
	// Get cluster information
	yarnCluster, err := s.ds.ReadYarnCluster(pz, clusterId)
	if err != nil {
		return errors.Wrap(err, "failed reading yarn cluster")
	}
	//	Get username
	identity, err := s.ds.ReadIdentity(pz, pz.Id())
	if err != nil {
		return errors.Wrap(err, "failed reading identity")
	}

	// FIXME check if file exists
	keytabPath := path.Join(s.workingDir, fs.KTDir, keytab)
	if err := yarn.StopCloud(s.kerberosEnabled, cluster.Name, yarnCluster.ApplicationId,
		yarnCluster.OutputDir, identity.Name, keytabPath); err != nil {
		return err
	}

	res := s.ds.DeleteCluster(pz, clusterId)

	if err := reloadProxyConf(s, pz, identity.Name); err != nil {
		return err
	}

	return res
}

func reloadProxyConf(s *Service, pz az.Principal, name string) error {
	clusters, err := s.ds.ReadAllClusters(pz)
	if err != nil {
		log.Println("Failed to read clusters.")
		return err
	}

	uid, gid, err := yarn.GetUser(name)
	if err := haproxy.Reload(clusters, uid, gid); err != nil {
		log.Println("Failed to reload proxy configuration.")
	}
	return nil
}

func (s *Service) GetCluster(pz az.Principal, clusterId int64) (*web.Cluster, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewCluster); err != nil {
		return nil, err
	}

	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return nil, err
	}
	return toCluster(cluster), nil
}

func (s *Service) GetClusterOnYarn(pz az.Principal, clusterId int64) (*web.YarnCluster, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewCluster); err != nil {
		return nil, err
	}
	cluster, err := s.ds.ReadYarnCluster(pz, clusterId)
	if err != nil {
		return nil, err
	}
	return toYarnCluster(cluster), nil
}

// func (s *Service) getCloud(pz az.Principal, cloudId int64) (*data.Cluster, error) {
// 	c, err := s.ds.ReadCluster(pz, cloudId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if c == nil {
// 		return nil, fmt.Errorf("Cloud %d does not exist.", cloudId)
// 	}
// 	return c, nil
// }

func (s *Service) GetClusters(pz az.Principal, offset, limit int64) ([]*web.Cluster, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewCluster); err != nil {
		return nil, err
	}
	clusters, err := s.ds.ReadClusters(pz, offset, limit)
	if err != nil {
		return nil, err
	}

	cs := make([]*web.Cluster, len(clusters))
	for i, cluster := range clusters {
		cs[i] = toCluster(cluster)
	}
	return cs, nil
}

// Returns the Cloud status from H2O
// This method should only be called if the cluster reports a non-Stopped status
// If the cloud was shut down from the outside of steam, will report Unknown
// / status for cloud
//
// TODO: Maybe this should only report if non-Stopped,non-Unknown status
//       In the case of Unknown, should only check if forced?
func (s *Service) GetClusterStatus(pz az.Principal, cloudId int64) (*web.ClusterStatus, error) { // Only called if cloud status != found
	if err := pz.CheckPermission(s.ds.Permissions.ViewCluster); err != nil {
		return nil, err
	}

	cluster, err := s.ds.ReadCluster(pz, cloudId)
	if err != nil {
		return nil, err
	}

	h2o := h2ov3.NewClient(cluster.Address)

	stat, err := h2o.GetCloudStatus()
	if err != nil {
		return &web.ClusterStatus{Status: "Unknown"}, nil
	}

	var (
		tot, all int32
		mem      int64
	)
	for _, n := range stat.Nodes {
		mem += n.MaxMem
		tot += n.NumCpus
		all += n.CpusAllowed
	}

	// FIXME: this needs a better impl
	var health string
	if stat.CloudHealthy {
		health = "healthy"
	} else {
		health = "unknown"
	}

	return &web.ClusterStatus{
		stat.Version,
		health,
		toSizeBytes(mem),
		int(tot),
		int(all),
	}, nil
}

func (s *Service) DeleteCluster(pz az.Principal, clusterId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageCluster); err != nil {
		return err
	}

	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return err
	}

	if cluster.State != data.StoppedState {
		return fmt.Errorf("Cannot delete a running cluster")
	}

	return s.ds.DeleteCluster(pz, clusterId)
}

type Jobs []*web.Job

func (k Jobs) Len() int {
	return len(k)
}

func (k Jobs) Less(i, j int) bool {
	switch {
	case k[i].Progress == "DONE" && k[j].Progress == "DONE":
		return k[i].CompletedAt < k[j].CompletedAt
	case k[i].Progress == "DONE":
		return true
	case k[j].Progress == "DONE":
		return false
	default:
		return k[i].CompletedAt < k[j].CompletedAt
	}
}

func (k Jobs) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

// FIXME where is this API used?
func (s *Service) GetJob(pz az.Principal, clusterId int64, jobName string) (*web.Job, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewCluster); err != nil {
		return nil, err
	}

	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return nil, err
	}

	h := h2ov3.NewClient(cluster.Address)

	j, err := h.GetJobsFetch(jobName)
	if err != nil {
		return nil, err
	}
	job := j.Jobs[0]

	return toJob(job), nil
}

func (s *Service) GetJobs(pz az.Principal, clusterId int64) ([]*web.Job, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewCluster); err != nil {
		return nil, err
	}

	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return nil, err
	}

	h := h2ov3.NewClient(cluster.Address)

	j, err := h.GetJobsList()
	if err != nil {
		return nil, err
	}

	jobs := make([]*web.Job, len(j.Jobs))
	for i, job := range j.Jobs {
		jobs[i] = toJob(job)
	}

	sort.Sort(sort.Reverse(Jobs(jobs)))

	return jobs, nil
}

// --- Project ---

func (s *Service) CreateProject(pz az.Principal, name, description, modelCategory string) (int64, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ManageProject); err != nil {
		return 0, err
	}

	projectId, err := s.ds.CreateProject(pz, name, description, modelCategory)
	if err != nil {
		return 0, err
	}

	return projectId, nil
}

func (s *Service) GetProjects(pz az.Principal, offset, limit int64) ([]*web.Project, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewProject); err != nil {
		return nil, err
	}

	projects, err := s.ds.ReadProjects(pz, offset, limit)
	if err != nil {
		return nil, err
	}

	return toProjects(projects), nil
}

func (s *Service) GetProject(pz az.Principal, projectId int64) (*web.Project, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewProject); err != nil {
		return nil, err
	}

	project, err := s.ds.ReadProject(pz, projectId)
	if err != nil {
		return nil, err
	}

	return toProject(project), nil
}

func (s *Service) deleteImplicitDatasouce(pz az.Principal, d data.Datasource) error {
	// Verify Datasource is Implicit
	if d.Kind != "Implicit" {
		return fmt.Errorf("non-implicit datasource: %s", d.Name)
	}

	// Read and delete all Datasets associated with Implicit Datasource
	dts, err := s.ds.ReadDatasets(pz, d.Id, 0, 10000)
	if err != nil {
		return errors.Wrap(err, "failed to read datasets")
	}
	for _, dt := range dts {
		if err := s.ds.DeleteDataset(pz, dt.Id); err != nil {
			return errors.Wrapf(err, "failed to delete dataset %s", dt.Name)
		}
	}

	// Delete Datasource
	return s.ds.DeleteDatasource(pz, d.Id)
}

func (s *Service) DeleteProject(pz az.Principal, projectId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageProject); err != nil {
		return err
	}

	// Verify no Models in Project
	if _, ok, err := s.ds.ReadModelsForProject(pz, projectId, 0, 1); err != nil {
		return err
	} else if ok {
		return fmt.Errorf("This project still contains at least one model.")
	}

	// Verify no Datasources in Project
	if _, ok, err := s.ds.ReadDatasourceByProject(pz, projectId); err != nil {
		return err
	} else if ok {
		// If Datasources present, remove all Implicit Datasources/Datasets
		ds, err := s.ds.ReadDatasources(pz, projectId, 0, 10000)
		if err != nil {
			return errors.Wrap(err, "failed to read datasources for project")
		}
		for _, d := range ds {
			if d.Kind != "Implicit" {
				return fmt.Errorf("project contains at least one datasource")
			}
			if err := s.deleteImplicitDatasouce(pz, d); err != nil {
				return errors.Wrapf(err, "failed to delete implicit datasource %s", d.Name)
			}
		}
	}

	return s.ds.DeleteProject(pz, projectId)
}

// --- Datasource ---

func (s *Service) CreateDatasource(pz az.Principal, projectId int64, name, description, filePath string) (int64, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ManageDatasource); err != nil {
		return 0, err
	}

	mapPath := map[string]string{"path": filePath}
	jsonPath, err := json.Marshal(mapPath)
	if err != nil {
		return 0, err
	}

	datasource := data.Datasource{
		0,
		projectId,
		name,
		description,
		"CSV", // FIXME: this is hardcoded
		string(jsonPath),
		time.Now(),
	}

	datasrcId, err := s.ds.CreateDatasource(pz, datasource)
	if err != nil {
		return 0, err
	}

	return datasrcId, nil
}

func (s *Service) GetDatasources(pz az.Principal, projectId, offset, limit int64) ([]*web.Datasource, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewDatasource); err != nil {
		return nil, err
	}

	datasources, err := s.ds.ReadDatasources(pz, projectId, offset, limit)
	if err != nil {
		return nil, err
	}

	return toDatasources(datasources), nil
}

func (s *Service) GetDatasource(pz az.Principal, datasourceId int64) (*web.Datasource, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewDatasource); err != nil {
		return nil, err
	}

	datasource, err := s.ds.ReadDatasource(pz, datasourceId)
	if err != nil {
		return nil, err
	}

	return toDatasource(datasource), nil
}

func (s *Service) UpdateDatasource(pz az.Principal, datasourceId int64, name, description, filePath string) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageDatasource); err != nil {
		return err
	}

	mapPath := map[string]string{"path": filePath}
	jsonPath, err := json.Marshal(mapPath)
	if err != nil {
		return err
	}

	datasource := data.Datasource{
		0,
		0,
		name,
		description,
		"CSV", // FIXME this is hardcoded
		string(jsonPath),
		time.Now(),
	}

	if err := s.ds.UpdateDatasource(pz, datasourceId, datasource); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteDatasource(pz az.Principal, datasourceId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageDatasource); err != nil {
		return err
	}

	_, ok, err := s.ds.ReadDatasetByDatasource(pz, datasourceId)
	if err != nil {
		return err
	}

	if ok {
		return fmt.Errorf("A dataset is still using this datasource.")
	}

	if err := s.ds.DeleteDatasource(pz, datasourceId); err != nil {
		return err
	}

	return nil
}

// --- Dataset ---

func (s *Service) importDataset(name, configuration, address string) ([]byte, string, error) {
	h2o := h2ov3.NewClient(address)

	// Translate json to string path
	rawJson := make(map[string]string)
	if err := json.Unmarshal([]byte(configuration), &rawJson); err != nil {
		return nil, "", err
	}
	filePath, ok := rawJson["path"]
	if !ok {
		return nil, "", fmt.Errorf("Cannot locate path: Empty datasource configuration")
	}

	importBody, err := h2o.PostImportFilesImportfiles(filePath)
	if err != nil {
		return nil, "", err
	}

	parseSetupBody, err := h2o.PostParseSetupGuesssetup(importBody.DestinationFrames)
	if err != nil {
		return nil, "", err
	}

	parseParms := bindings.NewParseV3()
	parseParms.FromParseSetup(*parseSetupBody)
	parseParms.Blocking = true
	parseBody, err := h2o.PostParseParse(parseParms)
	if err != nil {
		log.Fatalln(err)
	}

	job, err := h2o.JobPoll(parseBody.Job.Key.Name)
	if err != nil {
		return nil, "", err
	}
	rawFrame, _, err := h2o.GetFramesFetch(job.Dest.Name, false)
	if err != nil {
		return nil, "", err
	}

	return rawFrame, parseParms.DestinationFrame.Name, err
}

func (s *Service) CreateDataset(pz az.Principal, clusterId int64, datasourceId int64, name, description string, responseColumnName string) (int64, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ManageDataset); err != nil {
		return 0, err
	}

	datasource, err := s.ds.ReadDatasource(pz, datasourceId)
	if err != nil {
		return 0, err
	}
	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return 0, err
	}

	properties, frameName, err := s.importDataset(name, datasource.Configuration, cluster.Address)
	if err != nil {
		return 0, err
	}

	dataset := data.Dataset{
		0,
		datasourceId,
		name,
		description,
		frameName,
		responseColumnName,
		string(properties),
		"1",
		time.Now(),
	}

	datasetId, err := s.ds.CreateDataset(pz, dataset)
	if err != nil {
		return 0, err
	}

	return datasetId, nil
}

func (s *Service) GetDatasets(pz az.Principal, datasourceId int64, offset, limit int64) ([]*web.Dataset, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewDataset); err != nil {
		return nil, err
	}

	datasets, err := s.ds.ReadDatasets(pz, datasourceId, offset, limit)
	if err != nil {
		return nil, err
	}

	return toDatasets(datasets), nil
}

func (s *Service) GetDataset(pz az.Principal, datasetId int64) (*web.Dataset, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewDataset); err != nil {
		return nil, err
	}

	dataset, err := s.ds.ReadDataset(pz, datasetId)
	if err != nil {
		return nil, err
	}

	return toDataset(dataset), nil
}

// -- H2O to STEAM conversions --

func framesToDatasets(frames *bindings.FramesV3) []data.Dataset {
	array := make([]data.Dataset, len(frames.Frames))
	for i, frame := range frames.Frames {
		array[i] = frameToDataset(frame)
	}
	return array
}

func frameToDataset(frame *bindings.FrameBase) data.Dataset {
	return data.Dataset{
		0,
		0,
		frame.FrameId.Name,
		"",
		frame.FrameId.Name,
		"",
		"",
		"",
		time.Now(),
	}
}

func (s *Service) GetDatasetsFromCluster(pz az.Principal, clusterId int64) ([]*web.Dataset, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewCluster); err != nil {
		return nil, err
	}

	// Get cluster information
	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return nil, err
	}

	// Start h2o communication
	h2o := h2ov3.NewClient(cluster.Address)
	frames, err := h2o.GetFramesList()
	if err != nil {
		return nil, err
	}

	return toDatasets(framesToDatasets(frames)), nil
}

func (s *Service) UpdateDataset(pz az.Principal, datasetId int64, name, description, responseColumnName string) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageDataset); err != nil {
		return err
	}

	dataset := data.Dataset{
		0,
		0,
		name,
		description,
		"",
		responseColumnName,
		"",
		"1",
		time.Now(),
	}

	if err := s.ds.UpdateDataset(pz, datasetId, dataset); err != nil {
		return err
	}

	return nil
}

func (s *Service) SplitDataset(pz az.Principal, datasetId int64, ratio1 int, ratio2 int) ([]int64, error) {
	return nil, nil // XXX
}

func (s *Service) DeleteDataset(pz az.Principal, datasetId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageDataset); err != nil {
		return err
	}

	_, ok, err := s.ds.ReadModelByDataset(pz, datasetId)
	if err != nil {
		return err
	}

	if ok {
		return fmt.Errorf("A model is still using this dataset.")
	}

	if err := s.ds.DeleteDataset(pz, datasetId); err != nil {
		return err
	}

	return nil
}

// --- Model ---

func (s *Service) exportModel(h2o *h2ov3.H2O, modelName string, modelId int64) (string, string, error) {
	modelPath := fs.GetModelPath(s.workingDir, modelId)
	javaModelPath, err := h2o.ExportJavaModel(modelName, modelPath)
	if err != nil {
		return "", "", err
	}
	logicalName := fs.GetBasenameWithoutExt(javaModelPath)

	if _, err := h2o.ExportGenModel(modelPath); err != nil {
		return "", "", err
	}

	return strconv.FormatInt(modelId, 10), logicalName, err
}

func (s *Service) BuildModel(pz az.Principal, clusterId int64, datasetId int64, algorithm string) (int64, error) {
	return 0, nil // XXX Build default model, save to DB, return model id
}

func (s *Service) BuildModelAuto(pz az.Principal, clusterId int64, dataset, targetName string, maxRunTime int) (*web.Model, error) {

	return nil, fmt.Errorf("AutoML is currently not supported")

	if err := pz.CheckPermission(s.ds.Permissions.ManageModel); err != nil {
		return nil, err
	}
	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return nil, err
	}
	if cluster.State == data.StoppedState {
		return nil, fmt.Errorf("Cluster is not running")
	}

	h2o := h2ov3.NewClient(cluster.Address)

	modelKey, err := h2o.AutoML(dataset, targetName, maxRunTime) // TODO: can be a goroutine
	if err != nil {
		return nil, err
	}

	modelId, err := s.ds.CreateModel(pz, data.Model{
		0,
		0,               //FIXME -- should be a valid project ID to prevent a FK violation.
		0,               // FIXME -- should be a valid dataset ID to prevent a FK violation.
		sql.NullInt64{}, // FIXME -- should be a valid dataset ID to prevent a FK violation.
		modelKey,        // TODO this should be a modelName
		cluster.Id,
		cluster.Name,
		modelKey,
		"AutoML",
		"", // ModelCategory
		dataset,
		targetName,
		sql.NullString{},
		"",
		sql.NullString{},
		int64(maxRunTime),
		"",  // TODO Sebastian: put raw metrics json here (do not unmarshal/marshal json from h2o)
		"1", // MUST be "1"; will change when H2O's API version is bumped.
		time.Now(),
		sql.NullInt64{0, false},
		sql.NullString{"", false},
	})
	if err != nil {
		return nil, err
	}

	location, logicalName, err := s.exportModel(h2o, modelKey, modelId)
	if err != nil {
		return nil, err
	}

	if err := s.ds.UpdateModelLocation(pz, modelId, location, logicalName); err != nil {
		return nil, err
	}

	model, err := s.ds.ReadModel(pz, modelId)
	if err != nil {
		return nil, err
	}

	return toModel(model), nil
}

func (s *Service) GetModel(pz az.Principal, modelId int64) (*web.Model, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewModel); err != nil {
		return nil, err
	}
	model, err := s.ds.ReadModel(pz, modelId)
	if err != nil {
		return nil, err
	}
	return toModel(model), nil
}

func (s *Service) GetModels(pz az.Principal, projectId int64, offset, limit int64) ([]*web.Model, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewModel); err != nil {
		return nil, err
	}

	// Verifty project exists
	if _, err := s.ds.ReadProject(pz, projectId); err != nil {
		return nil, err
	}

	ms, _, err := s.ds.ReadModelsForProject(pz, projectId, offset, limit)
	if err != nil {
		return nil, err
	}

	models := make([]*web.Model, len(ms))
	for i, m := range ms {
		models[i] = toModel(m)
	}

	return models, nil
}

// Use this instead of model.DataFrame.Name because model.DataFrame can be nil
func dataFrameName(m *bindings.ModelSchema) string {
	if m.DataFrame != nil {
		return m.DataFrame.Name
	}

	return ""
}

func h2oToModel(model *bindings.ModelSchema) data.Model {
	return data.Model{
		Name:               model.ModelId.Name,
		ModelKey:           model.ModelId.Name,
		Algorithm:          model.AlgoFullName,
		ModelCategory:      string(model.Output.ModelCategory),
		DatasetName:        dataFrameName(model),
		ResponseColumnName: model.ResponseColumnName,
		Created:            time.Now(),
	}
}

func h2oToModels(models []*bindings.ModelSchema) []data.Model {
	array := make([]data.Model, len(models))
	for i, model := range models {
		array[i] = h2oToModel(model)
	}
	return array
}

func (s *Service) GetModelsFromCluster(pz az.Principal, clusterId int64, frameKey string) ([]*web.Model, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewCluster); err != nil {
		return nil, err
	}

	// Get cluster information
	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return nil, err
	}

	// Connect to h2o
	h2o := h2ov3.NewClient(cluster.Address)
	_, frame, err := h2o.GetFramesFetch(frameKey, true)
	if err != nil {
		return nil, err
	}

	models := h2oToModels(frame.CompatibleModels)

	ms := make([]*web.Model, len(models))
	for i, m := range models {
		ms[i] = toModel(m)
	}

	return ms, nil
}

func (s *Service) FindModelsCount(pz az.Principal, projectId int64) (int64, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewProject); err != nil {
		return 0, err
	}

	if _, err := s.ds.ReadProject(pz, projectId); err != nil {
		return 0, err
	}

	return s.ds.CountModelsForProject(pz, projectId)
}

// TODO: hardcoded; should be determined by h2o metrics
func (s *Service) GetAllBinomialSortCriteria(pz az.Principal) ([]string, error) {
	return []string{"mse", "r_squared", "logloss", "auc", "gini"}, nil
}

func (s *Service) FindModelsBinomial(pz az.Principal, projectId int64, namePart, sortBy string, ascending bool, offset, limit int64) ([]*web.BinomialModel, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewProject); err != nil {
		return nil, err
	}

	// Verify project exists
	if _, err := s.ds.ReadProject(pz, projectId); err != nil {
		return nil, errors.Wrap(err, "failed to read project from database")
	}

	models, err := s.ds.ReadBinomialModels(pz, projectId, namePart, sortBy, ascending, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read binomial models from database")
	}

	return toBinomialModels(models), nil
}

func (s *Service) GetModelBinomial(pz az.Principal, modelId int64) (*web.BinomialModel, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewModel); err != nil {
		return nil, err
	}

	model, err := s.ds.ReadBinomialModel(pz, modelId)
	if err != nil {
		return nil, err
	}

	return toBinomialModel(model), nil
}

// TODO: hardcoded; should be determined by h2o metrics
func (s *Service) GetAllMultinomialSortCriteria(pz az.Principal) ([]string, error) {
	return []string{"mse", "r_squared", "logloss"}, nil
}

func (s *Service) FindModelsMultinomial(pz az.Principal, projectId int64, namePart, sortBy string, ascending bool, offset, limit int64) ([]*web.MultinomialModel, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewProject); err != nil {
		return nil, err
	}

	// Verify project exists
	if _, err := s.ds.ReadProject(pz, projectId); err != nil {
		return nil, err
	}

	models, err := s.ds.ReadMultinomialModels(pz, projectId, namePart, sortBy, ascending, offset, limit)
	if err != nil {
		return nil, err
	}

	return toMultinomialModels(models), nil
}

func (s *Service) GetModelMultinomial(pz az.Principal, modelId int64) (*web.MultinomialModel, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewModel); err != nil {
		return nil, err
	}

	model, err := s.ds.ReadMultinomialModel(pz, modelId)
	if err != nil {
		return nil, err
	}

	return toMultinomialModel(model), nil
}

// TODO: hardcoded; should be determined by h2o metrics
func (s *Service) GetAllRegressionSortCriteria(pz az.Principal) ([]string, error) {
	return []string{"mse", "r_squared", "mean_residual_deviance"}, nil
}

func (s *Service) FindModelsRegression(pz az.Principal, projectId int64, namePart, sortBy string, ascending bool, offset, limit int64) ([]*web.RegressionModel, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewProject); err != nil {
		return nil, err
	}

	// Verify project exists
	if _, err := s.ds.ReadProject(pz, projectId); err != nil {
		return nil, err
	}

	models, err := s.ds.ReadRegressionModels(pz, projectId, namePart, sortBy, ascending, offset, limit)
	if err != nil {
		return nil, err
	}

	return toRegressionModels(models), nil
}

func (s *Service) GetModelRegression(pz az.Principal, modelId int64) (*web.RegressionModel, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewModel); err != nil {
		return nil, err
	}

	model, err := s.ds.ReadRegressionModel(pz, modelId)
	if err != nil {
		return nil, err
	}

	return toRegressionModel(model), nil
}

func (s *Service) ImportModelFromCluster(pz az.Principal, clusterId, projectId int64, modelKey, modelName string) (int64, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ManageModel); err != nil {
		return 0, err
	}

	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return 0, err
	}

	// Default modelName to modelKey
	if modelName == "" {
		modelName = modelKey
	}

	// get model from the cloud
	h2o := h2ov3.NewClient(cluster.Address)
	rawModel, r, err := h2o.GetModelsFetch(modelKey)
	if err != nil {
		return 0, err
	}

	m := r.Models[0]

	// fetch raw frame json from H2O
	rawFrame, _, err := h2o.GetFramesFetch(dataFrameName(m), false)
	if err != nil {
		return 0, err
	}

	datasourceId, err := s.ds.CreateDatasource(pz, data.Datasource{
		0,
		projectId,
		modelName + " Datasource",
		"Datasource for model " + modelName,
		"Implicit",
		"",
		time.Now(),
	})
	if err != nil {
		return 0, err
	}
	trainingDatasetId, err := s.ds.CreateDataset(pz, data.Dataset{
		0,
		datasourceId,
		modelName + " Dataset",
		"Dataset for model " + modelName,
		m.DataFrame.Name,
		m.ResponseColumnName,
		string(rawFrame),
		"1", // MUST be "1"; will change when H2O's API version is bumped.
		time.Now(),
	})
	if err != nil {
		return 0, err
	}

	// TODO: create a function to make this statically typed
	model := data.Model{
		ProjectId:          projectId,
		TrainingDatasetId:  trainingDatasetId,
		Name:               modelName,
		ClusterName:        cluster.Name,
		ClusterId:          cluster.Id,
		ModelKey:           modelKey,
		Algorithm:          m.AlgoFullName,
		ModelCategory:      string(m.Output.ModelCategory),
		DatasetName:        dataFrameName(m),
		ResponseColumnName: m.ResponseColumnName,
		Metrics:            string(rawModel),
		MetricsVersion:     "1",
		Created:            time.Now(),
	}

	modelId, err := s.ds.CreateModel(pz, model)
	if err != nil {
		return 0, err
	}

	if err := s.createMetricsTable(pz, modelId, m.Output.TrainingMetrics, string(m.Output.ModelCategory)); err != nil {
		return 0, err
	}

	return modelId, nil
}

func (s *Service) createMetricsTable(pz az.Principal, modelId int64, metrics *bindings.ModelMetrics, category string) error {
	switch category {
	case "Binomial":
		if err := s.ds.CreateBinomialModel(
			pz,
			modelId,
			metrics.Mse,
			metrics.R2,
			metrics.Logloss,
			metrics.Auc,
			metrics.Gini,
		); err != nil {
			return err
		}
	case "Multinomial":
		if err := s.ds.CreateMultinomialModel(
			pz,
			modelId,
			metrics.Mse,
			metrics.R2,
			metrics.Logloss,
		); err != nil {
			return err
		}
	case "Regression":
		if err := s.ds.CreateRegressionModel(
			pz,
			modelId,
			metrics.Mse,
			metrics.R2,
			metrics.MeanResidualDeviance); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Model category %s not supported", category)
	}

	return nil
}

func (s *Service) CheckMojo(pz az.Principal, algo string) (bool, error) {
	switch algo {
	case "Gradient Boosting Method", "Distributed Random Forest", "Deep Water":
		return true, nil
	}
	return false, nil
}

func (s *Service) ImportModelPojo(pz az.Principal, modelId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageModel); err != nil {
		return errors.Wrap(err, "failed checking permissions")
	}

	m, err := s.ds.ReadModel(pz, modelId)
	if err != nil {
		return errors.Wrap(err, "failed reading model from database")
	}

	c, err := s.ds.ReadCluster(pz, m.ClusterId)
	if err != nil {
		return errors.Wrap(err, "failed reading cluster from database")
	}
	h2o := h2ov3.NewClient(c.Address)

	modelPath := fs.GetModelPath(s.workingDir, modelId)
	javaModelPath, err := h2o.ExportJavaModel(m.ModelKey, modelPath)
	if err != nil {
		return errors.Wrap(err, "failed exporting java model from h2o")
	}
	if _, err := h2o.ExportGenModel(modelPath); err != nil {
		return errors.Wrap(err, "failed to export java dependency")
	}

	if !m.LogicalName.Valid {
		s.ds.UpdateModelLocation(pz, modelId, strconv.FormatInt(modelId, 10), fs.GetBasenameWithoutExt(javaModelPath))
	}

	return s.ds.UpdateModelObjectType(pz, modelId, "pojo")
}

func (s *Service) ImportModelMojo(pz az.Principal, modelId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageModel); err != nil {
		return errors.Wrap(err, "failed checking permissions")
	}

	m, err := s.ds.ReadModel(pz, modelId)
	if err != nil {
		return errors.Wrap(err, "failed reading model from database")
	}

	// Doesn't return errors
	if ok, _ := s.CheckMojo(pz, m.Algorithm); !ok {
		return fmt.Errorf("model of type %s does not have MOJO support", m.Algorithm)
	}

	c, err := s.ds.ReadCluster(pz, m.ClusterId)
	if err != nil {
		return errors.Wrap(err, "failed reading cluster from database")
	}
	h2o := h2ov3.NewClient(c.Address)

	modelPath := fs.GetModelPath(s.workingDir, modelId)
	mojoPath, err := h2o.ExportMOJO(m.ModelKey, modelPath)
	if err != nil {
		return errors.Wrap(err, "failed exporting MOJO from h2o")
	}
	if _, err := h2o.ExportGenModel(modelPath); err != nil {
		return errors.Wrap(err, "failed to export java dependency")
	}

	// MOJO only check
	if m.Algorithm == "Deep Water" {
		if _, err := h2o.ExportDeepWaterAll(modelPath); err != nil {
			return errors.Wrap(err, "exporting Deep Water dependency")
		}
	}

	if !m.LogicalName.Valid {
		s.ds.UpdateModelLocation(pz, modelId, strconv.FormatInt(modelId, 10), fs.GetBasenameWithoutExt(mojoPath))
	}

	return s.ds.UpdateModelObjectType(pz, modelId, "mojo")
}

func (s *Service) DeleteModel(pz az.Principal, modelId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageModel); err != nil {
		return err
	}

	// FIXME delete assets from disk

	_, err := s.ds.ReadModel(pz, modelId)
	if err != nil {
		return err
	}

	services, err := s.ds.ReadServicesForModelId(pz, modelId)
	if err != nil {
		return err
	}

	if len(services) > 0 {
		for _, service := range services {
			if service.State != data.StoppedState {
				return fmt.Errorf("A scoring service for this model is deployed and running at %s:%d", service.Address, service.Port)
			}
		}
	}

	return s.ds.DeleteModel(pz, modelId)
}

func (s *Service) CreateLabel(pz az.Principal, projectId int64, name, description string) (int64, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ManageLabel); err != nil {
		return 0, err
	}

	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return 0, fmt.Errorf("Label name cannot be empty")
	}

	if err := s.checkForDuplicateLabel(pz, projectId, name); err != nil {
		return 0, err
	}

	id, err := s.ds.CreateLabel(pz, projectId, name, description)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Service) UpdateLabel(pz az.Principal, labelId int64, name, description string) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageLabel); err != nil {
		return err
	}

	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return fmt.Errorf("Label name cannot be empty")
	}

	label, err := s.ds.ReadLabel(pz, labelId)
	if err != nil {
		return err
	}

	if err := s.checkForDuplicateLabel(pz, label.ProjectId, name); err != nil {
		return err
	}

	return s.ds.UpdateLabel(pz, labelId, name, description)
}

func (s *Service) checkForDuplicateLabel(pz az.Principal, projectId int64, name string) error {
	labels, err := s.ds.ReadLabelsForProject(pz, projectId)
	if err != nil {
		return err
	}

	for _, label := range labels {
		if label.Name == name {
			return fmt.Errorf("A label named %s is already associated with this project", name)
		}
	}
	return nil
}

func (s *Service) DeleteLabel(pz az.Principal, labelId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageLabel); err != nil {
		return err
	}

	return s.ds.DeleteLabel(pz, labelId)
}

func (s *Service) LinkLabelWithModel(pz az.Principal, labelId, modelId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageLabel); err != nil {
		return err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ManageModel); err != nil {
		return err
	}

	oldLabel, ok, err := s.ds.ReadLabelByModel(pz, modelId)
	if err != nil {
		return err
	}

	if ok {
		if err := s.ds.UnlinkLabelFromModel(pz, oldLabel.Id, modelId); err != nil {
			return err
		}
	}

	return s.ds.LinkLabelWithModel(pz, labelId, modelId)
}

func (s *Service) UnlinkLabelFromModel(pz az.Principal, labelId, modelId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageLabel); err != nil {
		return err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ManageModel); err != nil {
		return err
	}

	return s.ds.UnlinkLabelFromModel(pz, labelId, modelId)
}

func (s *Service) GetLabelsForProject(pz az.Principal, projectId int64) ([]*web.Label, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewLabel); err != nil {
		return nil, err
	}

	labels, err := s.ds.ReadLabelsForProject(pz, projectId)
	if err != nil {
		return nil, err
	}

	return toLabels(labels), nil
}

func isPortOpen(port int) bool {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func (s *Service) assignPort() (int, error) {
	randPort := rand.New(rand.NewSource(time.Now().UnixNano()))
	portRange := s.scoringServicePortMax - (s.scoringServicePortMin + 1)

	for i := s.scoringServicePortMin; i < s.scoringServicePortMax; i++ {
		port := randPort.Intn(portRange) + s.scoringServicePortMin + 1
		if isPortOpen(port) {
			return port, nil
		}
	}
	return 0, fmt.Errorf("No open port found within range %d:%d", s.scoringServicePortMin, s.scoringServicePortMax)
}

func (s *Service) StartService(pz az.Principal, modelId int64, name, packageName string) (int64, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ManageService); err != nil {
		return 0, err
	}

	model, err := s.ds.ReadModel(pz, modelId)
	if err != nil {
		return 0, err
	}

	artifact := compiler.ArtifactWar
	if len(packageName) > 0 {
		artifact = compiler.ArtifactPythonWar
	}

	warFilePath, err := compiler.CompileModel(
		s.compilationServiceAddress,
		s.workingDir,
		model.ProjectId,
		model.Id,
		model.LogicalName.String,
		model.ModelObjectType.String,
		model.Algorithm,
		artifact,
		packageName,
	)
	if err != nil {
		return 0, err
	}

	// Assign a port from allowed range
	port, err := s.assignPort()
	if err != nil {
		return 0, err
	}
	pid, err := svc.Start(
		warFilePath,
		fs.GetAssetsPath(s.workingDir, "jetty-runner.jar"),
		port,
		name,
		pz.Name(),
	)
	if err != nil {
		return 0, err
	}

	log.Printf("Scoring service started at %s:%d\n", s.scoringServiceAddress, port)

	service := data.Service{
		0,
		model.ProjectId,
		model.Id,
		name,
		s.scoringServiceAddress,
		int64(port), // FIXME change to int
		int64(pid),  // FIXME change to int
		data.StartedState,
		time.Now(),
	}

	serviceId, err := s.ds.CreateService(pz, service)
	if err != nil {
		return 0, err
	}

	return serviceId, nil
}

func (s *Service) StopService(pz az.Principal, serviceId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageService); err != nil {
		return err
	}
	service, err := s.ds.ReadService(pz, serviceId)
	if err != nil {
		return err
	}

	if service.State == data.StoppedState {
		return fmt.Errorf("Scoring service on model %s at port %d is already stopped", service.ModelId, service.Port)
	}

	if err := svc.Stop(int(service.ProcessId)); err != nil {
		return err
	}

	if err := s.ds.DeleteService(pz, serviceId); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetService(pz az.Principal, serviceId int64) (*web.ScoringService, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewService); err != nil {
		return nil, err
	}

	service, err := s.ds.ReadService(pz, serviceId)
	if err != nil {
		return nil, err
	}
	return toScoringService(service), nil
}

func (s *Service) GetServices(pz az.Principal, offset, limit int64) ([]*web.ScoringService, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewService); err != nil {
		return nil, err
	}

	services, err := s.ds.ReadServices(pz, offset, limit)
	if err != nil {
		return nil, err
	}
	ss := make([]*web.ScoringService, len(services))
	for i, service := range services {
		ss[i] = toScoringService(service)
	}

	return ss, nil
}

func (s *Service) GetServicesForProject(pz az.Principal, projectId, offset, limit int64) ([]*web.ScoringService, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewService); err != nil {
		return nil, err
	}

	services, err := s.ds.ReadServicesForProjectId(pz, projectId, offset, limit)
	if err != nil {
		return nil, err
	}
	ss := make([]*web.ScoringService, len(services))
	for i, service := range services {
		ss[i] = toScoringService(service)
	}

	return ss, nil
}

func (s *Service) GetServicesForModel(pz az.Principal, modelId, offset, limit int64) ([]*web.ScoringService, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewService); err != nil {
		return nil, err
	}

	services, err := s.ds.ReadServicesForModelId(pz, modelId)
	if err != nil {
		return nil, err
	}

	ss := make([]*web.ScoringService, len(services))
	for i, service := range services {
		ss[i] = toScoringService(service)
	}

	return ss, nil
}

func (s *Service) DeleteService(pz az.Principal, serviceId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageService); err != nil {
		return err
	}

	service, err := s.ds.ReadService(pz, serviceId)
	if err != nil {
		return err
	}

	if service.State != data.StoppedState || service.State != data.FailedState {
		return fmt.Errorf("Cannot delete service when in %s state", service.State)
	}

	if err := s.ds.DeleteService(pz, serviceId); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetEngine(pz az.Principal, engineId int64) (*web.Engine, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewEngine); err != nil {
		return nil, err
	}
	engine, err := s.ds.ReadEngine(pz, engineId)
	if err != nil {
		return nil, err
	}
	return toEngine(engine), nil
}

func (s *Service) GetEngines(pz az.Principal) ([]*web.Engine, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewEngine); err != nil {
		return nil, err
	}

	es, err := s.ds.ReadEngines(pz)
	if err != nil {
		return nil, err
	}

	engines := make([]*web.Engine, len(es))
	for i, e := range es {
		engines[i] = toEngine(e)
	}

	return engines, nil
}

func (s *Service) DeleteEngine(pz az.Principal, engineId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageEngine); err != nil {
		return err
	}

	// FIXME delete assets from disk

	_, err := s.ds.ReadEngine(pz, engineId)
	if err != nil {
		return err
	}

	return s.ds.DeleteEngine(pz, engineId)
}

func (s *Service) GetAllClusterTypes(pz az.Principal) ([]*web.ClusterType, error) {

	// No permission checks required

	return toClusterTypes(s.ds.ReadClusterTypes(pz)), nil
}

func (s *Service) GetAllEntityTypes(pz az.Principal) ([]*web.EntityType, error) {

	// No permission checks required

	return toEntityTypes(s.ds.ReadEntityTypes(pz)), nil
}

func (s *Service) GetAllPermissions(pz az.Principal) ([]*web.Permission, error) {

	// No permission checks required

	permissions, err := s.ds.ReadAllPermissions(pz)
	if err != nil {
		return nil, err
	}
	return toPermissions(permissions), nil
}

func (s *Service) GetPermissionsForRole(pz az.Principal, roleId int64) ([]*web.Permission, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewRole); err != nil {
		return nil, err
	}

	permissions, err := s.ds.ReadPermissionsForRole(pz, roleId)
	if err != nil {
		return nil, err
	}
	return toPermissions(permissions), nil
}

func (s *Service) GetPermissionsForIdentity(pz az.Principal, identityId int64) ([]*web.Permission, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewIdentity); err != nil {
		return nil, err
	}

	permissions, err := s.ds.ReadPermissionsForIdentity(pz, identityId)
	if err != nil {
		return nil, err
	}
	return toPermissions(permissions), nil
}

func (s *Service) CreateRole(pz az.Principal, name string, description string) (int64, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ManageRole); err != nil {
		return 0, err
	}
	return s.ds.CreateRole(pz, name, description)
}

func (s *Service) GetRoles(pz az.Principal, offset, limit int64) ([]*web.Role, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewRole); err != nil {
		return nil, err
	}

	roles, err := s.ds.ReadRoles(pz, offset, limit)
	if err != nil {
		return nil, err
	}
	return toRoles(roles), nil
}

func (s *Service) GetRolesForIdentity(pz az.Principal, identityId int64) ([]*web.Role, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewIdentity); err != nil {
		return nil, err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewRole); err != nil {
		return nil, err
	}

	roles, err := s.ds.ReadRolesForIdentity(pz, identityId)
	if err != nil {
		return nil, err
	}
	return toRoles(roles), nil
}

func (s *Service) GetRole(pz az.Principal, roleId int64) (*web.Role, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewRole); err != nil {
		return nil, err
	}

	role, err := s.ds.ReadRole(pz, roleId)
	if err != nil {
		return nil, err
	}
	return toRole(role), nil
}

func (s *Service) GetRoleByName(pz az.Principal, name string) (*web.Role, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewRole); err != nil {
		return nil, err
	}

	role, err := s.ds.ReadRoleByName(pz, name)
	if err != nil {
		return nil, err
	}
	return toRole(role), nil
}

func (s *Service) UpdateRole(pz az.Principal, roleId int64, name string, description string) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageRole); err != nil {
		return err
	}

	return s.ds.UpdateRole(pz, roleId, name, description)
}

func (s *Service) LinkRoleWithPermissions(pz az.Principal, roleId int64, permissionIds []int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageRole); err != nil {
		return err
	}

	return s.ds.LinkRoleAndPermissions(pz, roleId, permissionIds)
}

func (s *Service) LinkRoleWithPermission(pz az.Principal, roleId int64, permissionId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageRole); err != nil {
		return err
	}

	return s.ds.LinkRoleWithPermission(pz, roleId, permissionId)
}

func (s *Service) UnlinkRoleFromPermission(pz az.Principal, roleId int64, permissionId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageRole); err != nil {
		return err
	}

	return s.ds.UnlinkRoleFromPermission(pz, roleId, permissionId)
}

func (s *Service) DeleteRole(pz az.Principal, roleId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageRole); err != nil {
		return err
	}

	return s.ds.DeleteRole(pz, roleId)
}

func (s *Service) CreateWorkgroup(pz az.Principal, name string, description string) (int64, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ManageWorkgroup); err != nil {
		return 0, err
	}

	return s.ds.CreateWorkgroup(pz, name, description)
}

func (s *Service) GetWorkgroups(pz az.Principal, offset, limit int64) ([]*web.Workgroup, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewWorkgroup); err != nil {
		return nil, err
	}

	workgroups, err := s.ds.ReadWorkgroups(pz, offset, limit)
	if err != nil {
		return nil, err
	}
	return toWorkgroups(workgroups), nil
}

func (s *Service) GetWorkgroupsForIdentity(pz az.Principal, identityId int64) ([]*web.Workgroup, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewIdentity); err != nil {
		return nil, err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewWorkgroup); err != nil {
		return nil, err
	}

	workgroups, err := s.ds.ReadWorkgroupsForIdentity(pz, identityId)
	if err != nil {
		return nil, err
	}
	return toWorkgroups(workgroups), nil
}

func (s *Service) GetWorkgroup(pz az.Principal, workgroupId int64) (*web.Workgroup, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewWorkgroup); err != nil {
		return nil, err
	}

	workgroup, err := s.ds.ReadWorkgroup(pz, workgroupId)
	if err != nil {
		return nil, err
	}
	return toWorkgroup(workgroup), nil
}

func (s *Service) GetWorkgroupByName(pz az.Principal, name string) (*web.Workgroup, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewWorkgroup); err != nil {
		return nil, err
	}

	workgroup, err := s.ds.ReadWorkgroupByName(pz, name)
	if err != nil {
		return nil, err
	}
	return toWorkgroup(workgroup), nil
}

func (s *Service) UpdateWorkgroup(pz az.Principal, workgroupId int64, name string, description string) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageWorkgroup); err != nil {
		return err
	}

	return s.ds.UpdateWorkgroup(pz, workgroupId, name, description)
}

func (s *Service) DeleteWorkgroup(pz az.Principal, workgroupId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageWorkgroup); err != nil {
		return err
	}

	return s.ds.DeleteWorkgroup(pz, workgroupId)
}

func (s *Service) CreateIdentity(pz az.Principal, name string, password string) (int64, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ManageIdentity); err != nil {
		return 0, err
	}

	hash, err := auth.HashPassword(password)
	if err != nil {
		return 0, err
	}
	id, _, err := s.ds.CreateIdentity(pz, name, hash)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Service) GetIdentities(pz az.Principal, offset, limit int64) ([]*web.Identity, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewIdentity); err != nil {
		return nil, err
	}

	identities, err := s.ds.ReadIdentities(pz, offset, limit)
	if err != nil {
		return nil, err
	}
	return toIdentities(identities), nil
}

func (s *Service) GetIdentitiesForWorkgroup(pz az.Principal, workgroupId int64) ([]*web.Identity, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewIdentity); err != nil {
		return nil, err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewWorkgroup); err != nil {
		return nil, err
	}

	identities, err := s.ds.ReadIdentitiesForWorkgroup(pz, workgroupId)
	if err != nil {
		return nil, err
	}
	return toIdentities(identities), nil
}

func (s *Service) GetIdentitiesForRole(pz az.Principal, roleId int64) ([]*web.Identity, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewIdentity); err != nil {
		return nil, err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewRole); err != nil {
		return nil, err
	}

	identities, err := s.ds.ReadIdentitiesForRole(pz, roleId)
	if err != nil {
		return nil, err
	}
	return toIdentities(identities), nil
}

func (s *Service) GetIdentitiesForEntity(pz az.Principal, entityTypeId, entityId int64) ([]*web.UserRole, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewIdentity); err != nil {
		return nil, err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewRole); err != nil {
		return nil, err
	}

	userRoles, err := s.ds.ReadUsersForEntity(pz, entityTypeId, entityId)
	if err != nil {
		return nil, err
	}

	return toUserRoles(userRoles), nil
}

func (s *Service) GetIdentity(pz az.Principal, identityId int64) (*web.Identity, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewIdentity); err != nil {
		return nil, err
	}

	identity, err := s.ds.ReadIdentity(pz, identityId)
	if err != nil {
		return nil, err
	}
	return toIdentity(identity), err
}

func (s *Service) GetIdentityByName(pz az.Principal, name string) (*web.Identity, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewIdentity); err != nil {
		return nil, err
	}

	identity, err := s.ds.ReadIdentityByName(pz, name)
	if err != nil {
		return nil, err
	}
	return toIdentity(identity), err
}

func (s *Service) LinkIdentityWithWorkgroup(pz az.Principal, identityId int64, workgroupId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageIdentity); err != nil {
		return err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewWorkgroup); err != nil {
		return err
	}

	return s.ds.LinkIdentityAndWorkgroup(pz, identityId, workgroupId)
}

func (s *Service) UnlinkIdentityFromWorkgroup(pz az.Principal, identityId int64, workgroupId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageIdentity); err != nil {
		return err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewWorkgroup); err != nil {
		return err
	}

	return s.ds.UnlinkIdentityAndWorkgroup(pz, identityId, workgroupId)
}

func (s *Service) LinkIdentityWithRole(pz az.Principal, identityId int64, roleId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageIdentity); err != nil {
		return err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewRole); err != nil {
		return err
	}

	return s.ds.LinkIdentityAndRole(pz, identityId, roleId)
}

func (s *Service) UnlinkIdentityFromRole(pz az.Principal, identityId int64, roleId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageIdentity); err != nil {
		return err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewRole); err != nil {
		return err
	}

	return s.ds.UnlinkIdentityAndRole(pz, identityId, roleId)
}

func (s *Service) UpdateIdentity(pz az.Principal, identityId int64, password string) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageIdentity); err != nil {
		return err
	}

	hash, err := auth.HashPassword(password)
	if err != nil {
		return fmt.Errorf("Failed hashing password: %s", err)
	}

	return s.ds.UpdateIdentity(pz, identityId, hash)
}

func (s *Service) ActivateIdentity(pz az.Principal, identityId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageIdentity); err != nil {
		return err
	}

	return s.ds.ActivateIdentity(pz, identityId)
}

func (s *Service) DeactivateIdentity(pz az.Principal, identityId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageIdentity); err != nil {
		return err
	}

	return s.ds.DeactivateIdentity(pz, identityId)
}

func (s *Service) ShareEntity(pz az.Principal, kind string, workgroupId, entityTypeId, entityId int64) error {
	if err := pz.CheckPermission(s.ds.ManagePermissions[entityTypeId]); err != nil {
		return err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewWorkgroup); err != nil {
		return err
	}

	return s.ds.CreatePrivilege(pz, data.Privilege{
		kind,
		workgroupId,
		entityTypeId,
		entityId,
	})
}

func (s *Service) GetPrivileges(pz az.Principal, entityTypeId, entityId int64) ([]*web.EntityPrivilege, error) {
	if err := pz.CheckPermission(s.ds.ViewPermissions[entityTypeId]); err != nil {
		return nil, err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewWorkgroup); err != nil {
		return nil, err
	}

	privileges, err := s.ds.ReadEntityPrivileges(pz, entityTypeId, entityId)
	if err != nil {
		return nil, err
	}
	return toEntityPrivileges(privileges), nil
}

func (s *Service) UnshareEntity(pz az.Principal, kind string, workgroupId, entityTypeId, entityId int64) error {
	if err := pz.CheckPermission(s.ds.ManagePermissions[entityTypeId]); err != nil {
		return err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewWorkgroup); err != nil {
		return err
	}

	return s.ds.DeletePrivilege(pz, data.Privilege{
		kind,
		workgroupId,
		entityTypeId,
		entityId,
	})
}

func (s *Service) GetHistory(pz az.Principal, entityTypeId, entityId, offset, limit int64) ([]*web.EntityHistory, error) {
	if err := pz.CheckPermission(s.ds.ViewPermissions[entityTypeId]); err != nil {
		return nil, err
	}

	history, err := s.ds.ReadHistoryForEntity(pz, entityTypeId, entityId, offset, limit)
	if err != nil {
		return nil, err
	}
	return toEntityHistory(history), nil
}

func (s *Service) CreatePackage(pz az.Principal, projectId int64, name string) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageProject); err != nil {
		return err
	}

	if err := pz.CheckEdit(s.ds.EntityTypes.Project, projectId); err != nil {
		return err
	}

	if err := fs.ValidateName(name); err != nil {
		return fmt.Errorf("Invalid package name: %s", err)
	}

	packagePath := fs.GetPackagePath(s.workingDir, projectId, name)
	if fs.DirExists(packagePath) {
		return fmt.Errorf("Failed creating package directory: %s already exists", name)
	}

	if err := fs.Mkdir(packagePath); err != nil {
		return fmt.Errorf("Failed creating package directory: %s", err)
	}

	return nil
}

func (s *Service) GetPackages(pz az.Principal, projectId int64) ([]string, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewProject); err != nil {
		return nil, err
	}

	if err := pz.CheckView(s.ds.EntityTypes.Project, projectId); err != nil {
		return nil, err
	}

	projectPath := fs.GetProjectPath(s.workingDir, projectId)
	if !fs.DirExists(projectPath) {
		return []string{}, nil
	}

	dirs, err := fs.ListDirs(projectPath)
	if err != nil {
		return nil, fmt.Errorf("Failed listing packages: %s", err)
	}

	return dirs, nil
}

func (s *Service) GetPackageDirectories(pz az.Principal, projectId int64, packageName string, relativePath string) ([]string, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewProject); err != nil {
		return nil, err
	}

	if err := pz.CheckView(s.ds.EntityTypes.Project, projectId); err != nil {
		return nil, err
	}

	packagePath := fs.GetPackagePath(s.workingDir, projectId, packageName)
	if !fs.DirExists(packagePath) {
		return nil, fmt.Errorf("Package %s does not exist")
	}

	packageDirPath, err := fs.GetPackageRelativePath(s.workingDir, projectId, packageName, relativePath)
	if err != nil {
		return nil, err
	}

	if !fs.DirExists(packageDirPath) {
		return []string{}, nil
	}

	dirs, err := fs.ListDirs(packageDirPath)
	if err != nil {
		return nil, fmt.Errorf("Failed listing package directories: %s", err)
	}

	return dirs, nil
}

func (s *Service) GetPackageFiles(pz az.Principal, projectId int64, packageName string, relativePath string) ([]string, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewProject); err != nil {
		return nil, err
	}

	if err := pz.CheckView(s.ds.EntityTypes.Project, projectId); err != nil {
		return nil, err
	}

	packagePath := fs.GetPackagePath(s.workingDir, projectId, packageName)
	if !fs.DirExists(packagePath) {
		return nil, fmt.Errorf("Package %s does not exist")
	}

	packageDirPath, err := fs.GetPackageRelativePath(s.workingDir, projectId, packageName, relativePath)
	if err != nil {
		return nil, err
	}

	if !fs.DirExists(packageDirPath) {
		return []string{}, nil
	}

	files, err := fs.ListFiles(packageDirPath)
	if err != nil {
		return nil, fmt.Errorf("Failed listing package files: %s", err)
	}

	return files, nil
}

func (s *Service) DeletePackage(pz az.Principal, projectId int64, name string) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageProject); err != nil {
		return err
	}

	if err := pz.CheckEdit(s.ds.EntityTypes.Project, projectId); err != nil {
		return err
	}

	packagePath := fs.GetPackagePath(s.workingDir, projectId, name)
	if !fs.DirExists(packagePath) {
		return fmt.Errorf("Package %s does not exist")
	}

	if err := fs.Rmdir(packagePath); err != nil {
		return fmt.Errorf("Failed deleting package: %s", err)
	}

	return nil
}

func (s *Service) DeletePackageDirectory(pz az.Principal, projectId int64, packageName string, relativePath string) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageProject); err != nil {
		return err
	}

	if err := pz.CheckEdit(s.ds.EntityTypes.Project, projectId); err != nil {
		return err
	}

	packagePath := fs.GetPackagePath(s.workingDir, projectId, packageName)
	if !fs.DirExists(packagePath) {
		return fmt.Errorf("Package %s does not exist")
	}

	dirPath, err := fs.GetPackageRelativePath(s.workingDir, projectId, packageName, relativePath)
	if err != nil {
		return err
	}

	if !fs.DirExists(dirPath) {
		return fmt.Errorf("Invalid path %s in package %s", relativePath, packageName)
	}

	if err := fs.Rmdir(dirPath); err != nil {
		return fmt.Errorf("Failed deleting directory: %s", err)
	}

	return nil
}

func (s *Service) DeletePackageFile(pz az.Principal, projectId int64, packageName string, relativePath string) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageProject); err != nil {
		return err
	}

	if err := pz.CheckEdit(s.ds.EntityTypes.Project, projectId); err != nil {
		return err
	}

	packagePath := fs.GetPackagePath(s.workingDir, projectId, packageName)
	if !fs.DirExists(packagePath) {
		return fmt.Errorf("Package %s does not exist")
	}

	filePath, err := fs.GetPackageRelativePath(s.workingDir, projectId, packageName, relativePath)
	if err != nil {
		return err
	}

	if !fs.FileExists(filePath) {
		return fmt.Errorf("Invalid path %s in package %s", relativePath, packageName)
	}

	if err := fs.Rm(filePath); err != nil {
		return fmt.Errorf("Failed deleting file: %s", err)
	}

	return nil
}

func (s *Service) SetAttributesForPackage(pz az.Principal, projectId int64, packageName string, attributes string) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageProject); err != nil {
		return err
	}

	if err := pz.CheckEdit(s.ds.EntityTypes.Project, projectId); err != nil {
		return err
	}

	if err := fs.SetPackageAttributes(s.workingDir, projectId, packageName, []byte(attributes)); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAttributesForPackage(pz az.Principal, projectId int64, packageName string) (string, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewProject); err != nil {
		return "", err
	}

	if err := pz.CheckView(s.ds.EntityTypes.Project, projectId); err != nil {
		return "", err
	}

	b, err := fs.GetPackageAttributes(s.workingDir, projectId, packageName)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Helper function to convert from int to bytes
func toSizeBytes(i int64) string {
	f := float64(i)

	s := 0
	for f > 1024 {
		f /= 1024
		s++
	}
	b := strconv.FormatFloat(f, 'f', 2, 64)

	switch s {
	case 0:
		return b + " B"
	case 1:
		return b + " KB"
	case 2:
		return b + " MB"
	case 3:
		return b + " GB"
	case 4:
		return b + " TB"
	case 5:
		return b + " PB"
	}

	return ""
}

//
// Routines to convert DB structs into API structs
//

func toCluster(c data.Cluster) *web.Cluster {
	return &web.Cluster{
		c.Id, // Name
		c.Name,
		c.TypeId,
		c.DetailId,
		c.Address,
		c.Token,
		c.State,
		toTimestamp(c.Created),
	}
}

func toYarnCluster(c data.YarnCluster) *web.YarnCluster {
	return &web.YarnCluster{
		c.Id,
		c.EngineId,
		int(c.Size), // FIXME change db field to int
		c.ApplicationId,
		c.Memory,
		c.Username,
	}
}

func fromNullInt64(maybeId sql.NullInt64) int64 {
	if maybeId.Valid {
		return maybeId.Int64
	}
	return -1
}

func fromNullString(maybeId sql.NullString) string {
	if maybeId.Valid {
		return maybeId.String
	}
	return ""
}

// TODO: Fix these NULLSQL values
func toModel(m data.Model) *web.Model {
	return &web.Model{
		m.Id,
		m.TrainingDatasetId,
		m.ValidationDatasetId.Int64,
		m.Name,
		m.ClusterName,
		m.ModelKey,
		m.Algorithm,
		m.ModelCategory,
		m.DatasetName,
		m.ResponseColumnName,
		m.LogicalName.String,
		m.Location,
		m.ModelObjectType.String,
		int(m.MaxRunTime), // FIXME change db field to int
		m.Metrics,
		toTimestamp(m.Created),
		fromNullInt64(m.LabelId),
		fromNullString(m.LabelName),
	}
}

// TODO: Fix these NULLSQL values
func toBinomialModel(model data.BinomialModel) *web.BinomialModel {
	return &web.BinomialModel{
		model.Id,
		model.TrainingDatasetId,
		model.ValidationDatasetId.Int64,
		model.Name,
		model.ClusterName,
		model.ModelKey,
		model.Algorithm,
		model.ModelCategory,
		model.DatasetName,
		model.ResponseColumnName,
		model.LogicalName.String,
		model.Location,
		model.ModelObjectType.String,
		int(model.MaxRunTime), // FIXME change db field to int
		model.Metrics,
		toTimestamp(model.Created),
		fromNullInt64(model.LabelId),
		fromNullString(model.LabelName),
		model.Mse,
		model.RSquared,
		model.Logloss,
		model.Auc,
		model.Gini,
	}
}

func toBinomialModels(models []data.BinomialModel) []*web.BinomialModel {
	array := make([]*web.BinomialModel, len(models))
	for i, model := range models {
		array[i] = toBinomialModel(model)
	}
	return array
}

// TODO: Fix these NULLSQL values
func toMultinomialModel(model data.MultinomialModel) *web.MultinomialModel {
	return &web.MultinomialModel{
		model.Id,
		model.TrainingDatasetId,
		model.ValidationDatasetId.Int64,
		model.Name,
		model.ClusterName,
		model.ModelKey,
		model.Algorithm,
		model.ModelCategory,
		model.DatasetName,
		model.ResponseColumnName,
		model.LogicalName.String,
		model.Location,
		model.ModelObjectType.String,
		int(model.MaxRunTime), // FIXME change db field to int
		model.Metrics,
		toTimestamp(model.Created),
		fromNullInt64(model.LabelId),
		fromNullString(model.LabelName),
		model.Mse,
		model.RSquared,
		model.Logloss,
	}
}

func toMultinomialModels(models []data.MultinomialModel) []*web.MultinomialModel {
	array := make([]*web.MultinomialModel, len(models))
	for i, model := range models {
		array[i] = toMultinomialModel(model)
	}
	return array
}

func toRegressionModel(model data.RegressionModel) *web.RegressionModel {
	return &web.RegressionModel{
		model.Id,
		model.TrainingDatasetId,
		model.ValidationDatasetId.Int64,
		model.Name,
		model.ClusterName,
		model.ModelKey,
		model.Algorithm,
		model.ModelCategory,
		model.DatasetName,
		model.ResponseColumnName,
		model.LogicalName.String,
		model.Location,
		model.ModelObjectType.String,
		int(model.MaxRunTime), // FIXME change db field to int
		model.Metrics,
		toTimestamp(model.Created),
		fromNullInt64(model.LabelId),
		fromNullString(model.LabelName),
		model.Mse,
		model.RSquared,
		model.MeanResidualDeviance,
	}
}

func toRegressionModels(models []data.RegressionModel) []*web.RegressionModel {
	array := make([]*web.RegressionModel, len(models))
	for i, model := range models {
		array[i] = toRegressionModel(model)
	}
	return array
}

func toScoringService(s data.Service) *web.ScoringService {
	return &web.ScoringService{
		s.Id,
		s.ModelId,
		s.Name,
		s.Address,
		int(s.Port),      // FIXME change db field to int
		int(s.ProcessId), // FIXME change db field to int
		s.State,
		toTimestamp(s.Created),
	}
}

func toEngine(e data.Engine) *web.Engine {
	return &web.Engine{
		e.Id,
		e.Name,
		e.Location,
		toTimestamp(e.Created),
	}
}

func toClusterTypes(entityTypes []data.ClusterType) []*web.ClusterType {
	array := make([]*web.ClusterType, len(entityTypes))
	for i, ct := range entityTypes {
		array[i] = &web.ClusterType{
			ct.Id,
			ct.Name,
		}
	}
	return array
}

func toEntityTypes(entityTypes []data.EntityType) []*web.EntityType {
	array := make([]*web.EntityType, len(entityTypes))
	for i, et := range entityTypes {
		array[i] = &web.EntityType{
			et.Id,
			et.Name,
		}
	}
	return array
}

func toPermissions(permissions []data.Permission) []*web.Permission {
	array := make([]*web.Permission, len(permissions))
	for i, p := range permissions {
		array[i] = &web.Permission{
			p.Id,
			p.Code,
			p.Description,
		}
	}
	return array
}

func toRole(r data.Role) *web.Role {
	return &web.Role{
		r.Id,
		r.Name,
		r.Description,
		toTimestamp(r.Created),
	}
}

func toRoles(roles []data.Role) []*web.Role {
	array := make([]*web.Role, len(roles))
	for i, r := range roles {
		array[i] = toRole(r)
	}
	return array
}

func toWorkgroup(w data.Workgroup) *web.Workgroup {
	return &web.Workgroup{
		w.Id,
		w.Name,
		w.Description,
		toTimestamp(w.Created),
	}
}

func toWorkgroups(workgroups []data.Workgroup) []*web.Workgroup {
	array := make([]*web.Workgroup, len(workgroups))
	for i, r := range workgroups {
		array[i] = toWorkgroup(r)
	}
	return array
}

func toIdentity(u data.Identity) *web.Identity {
	var lastLogin time.Time
	if u.LastLogin.Valid {
		lastLogin = u.LastLogin.Time
	}
	return &web.Identity{
		u.Id,
		u.Name,
		u.IsActive,
		toTimestamp(lastLogin),
		toTimestamp(u.Created),
	}
}

func toIdentities(workgroups []data.Identity) []*web.Identity {
	array := make([]*web.Identity, len(workgroups))
	for i, r := range workgroups {
		array[i] = toIdentity(r)
	}
	return array
}

func toUserRole(u data.IdentityAndRole) *web.UserRole {
	return &web.UserRole{
		u.Kind,
		u.IdentityId,
		u.IdentityName,
		u.RoleId,
		u.RoleName,
	}
}

func toUserRoles(users []data.IdentityAndRole) []*web.UserRole {
	array := make([]*web.UserRole, len(users))
	for i, u := range users {
		array[i] = toUserRole(u)
	}
	return array
}

func toEntityPrivileges(entityPrivileges []data.EntityPrivilege) []*web.EntityPrivilege {
	array := make([]*web.EntityPrivilege, len(entityPrivileges))
	for i, ep := range entityPrivileges {
		array[i] = &web.EntityPrivilege{
			ep.Type,
			ep.WorkgroupId,
			ep.WorkgroupName,
			ep.WorkgroupDescription,
		}
	}
	return array
}

func toEntityHistory(entityHistory []data.EntityHistory) []*web.EntityHistory {
	array := make([]*web.EntityHistory, len(entityHistory))
	for i, h := range entityHistory {
		array[i] = &web.EntityHistory{
			h.IdentityId,
			h.Action,
			h.Description,
			toTimestamp(h.Created),
		}
	}
	return array
}

func toProject(project data.Project) *web.Project {
	return &web.Project{
		project.Id,
		project.Name,
		project.Description,
		project.ModelCategory,
		toTimestamp(project.Created),
	}
}

func toProjects(projects []data.Project) []*web.Project {
	array := make([]*web.Project, len(projects))
	for i, project := range projects {
		array[i] = toProject(project)
	}
	return array
}

func toLabels(labels []data.Label) []*web.Label {
	array := make([]*web.Label, len(labels))
	for i, label := range labels {
		var modelId int64 = -1
		if label.ModelId.Valid {
			modelId = label.ModelId.Int64
		}
		array[i] = &web.Label{
			label.Id,
			label.ProjectId,
			modelId,
			label.Name,
			label.Description,
			toTimestamp(label.Created),
		}
	}
	return array
}

func toDatasource(datasource data.Datasource) *web.Datasource {
	return &web.Datasource{
		datasource.Id,
		datasource.ProjectId,
		datasource.Name,
		datasource.Description,
		datasource.Kind,
		datasource.Configuration,
		toTimestamp(datasource.Created),
	}
}

func toDatasources(datasources []data.Datasource) []*web.Datasource {
	array := make([]*web.Datasource, len(datasources))
	for i, datasource := range datasources {
		array[i] = toDatasource(datasource)
	}
	return array
}

func toDataset(dataset data.Dataset) *web.Dataset {
	return &web.Dataset{
		dataset.Id,
		dataset.DatasourceId,
		dataset.Name,
		dataset.Description,
		dataset.FrameName,
		dataset.ResponseColumnName,
		dataset.Properties,
		toTimestamp(dataset.Created),
	}
}

func toDatasets(datasets []data.Dataset) []*web.Dataset {
	array := make([]*web.Dataset, len(datasets))
	for i, dataset := range datasets {
		array[i] = toDataset(dataset)
	}
	return array
}

//
// Routines to convert H2O structs into API structs
//

func toJob(j *bindings.JobV3) *web.Job {
	var end int64
	if j.Status == "DONE" {
		end = j.StartTime + j.Msec
	}

	return &web.Job{
		j.Key.Name,
		"",
		j.Description,
		j.Status,
		j.StartTime,
		end,
	}
}
