/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package web

import (
	"fmt"
	"path"
	"strconv"
	"time"

	"github.com/h2oai/steam/bindings"
	"github.com/h2oai/steam/lib/fs"
	"github.com/h2oai/steam/lib/yarn"
	"github.com/h2oai/steam/master/az"
	"github.com/h2oai/steam/master/data"
	"github.com/h2oai/steam/srv/h2ov3"
	"github.com/h2oai/steam/srv/web"
	"github.com/pkg/errors"
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

// --- ------- ---
// --- ------- ---
// --- Cluster ---
// --- ------- ---
// --- ------- ---

func (s *Service) RegisterCluster(pz az.Principal, address string) (int64, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageCluster); err != nil {
		return 0, errors.Wrap(err, "checking permission")
	}
	// Getting cluster information
	h := h2ov3.NewClient(address)
	cloud, err := h.GetCloudStatus()
	if err != nil {
		return 0, errors.Wrap(err, "communicating with cluster")
	}
	// Check that address is unique
	if _, exists, err := s.ds.ReadCluster(
		data.ByAddress(address),
		data.ByPrivilege(pz),
	); err != nil {
		return 0, errors.Wrap(err, "reading cluster from database")
	} else if exists {
		return 0, errors.Wrapf(err, "a cluster with the address %s is already registered", address)
	}
	// Create cluster
	// TODO: change s.ds.ClusterType to data.ClusterType.External
	clusterId, err := s.ds.CreateCluster(cloud.CloudName, s.ds.ClusterType.External,
		data.WithAddress(address), data.WithState(data.States.Started),
		data.WithPrivilege(pz, data.Owns), data.WithAudit(pz),
	)
	return clusterId, errors.Wrap(err, "creating cluster in database")
}

func (s *Service) UnregisterCluster(pz az.Principal, clusterId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageCluster); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckOwns(s.ds.EntityType.Cluster, clusterId); err != nil {
		return errors.Wrap(err, "checking ownership")
	}
	// Fetch cluster details
	cluster, exists, err := s.ds.ReadCluster(data.ById(clusterId))
	if err != nil {
		return errors.Wrap(err, "reading cluster from database")
	} else if !exists {
		return errors.New("cannot locate cluster in database")
	}
	if cluster.ClusterTypeId != s.ds.ClusterType.External {
		return errors.New("cannot unregister non-external clusters")
	}
	// Delete cluster
	return errors.Wrap(s.ds.DeleteCluster(clusterId, data.WithAudit(pz)), "deleting cluster from database")
}

func (s *Service) StartClusterOnYarn(pz az.Principal, clusterName string, engineId int64, size int, memory, keytab string) (int64, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageCluster); err != nil {
		return 0, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Engine, engineId); err != nil {
		return 0, errors.Wrap(err, "checking view privileges")
	}
	// Check that name is unique to user
	_, exists, err := s.ds.ReadCluster(data.ByName(clusterName), data.ByPrivilege(pz))
	if err != nil {
		return 0, errors.Wrap(err, "reading cluster from database")
	} else if exists {
		return 0, fmt.Errorf("a cluster with name %s already exists for this user", clusterName)
	}
	// Fetch identity details
	identity, exists, err := s.ds.ReadIdentity(data.ById(pz.Id()))
	if err != nil {
		return 0, errors.Wrap(err, "reading identity from database")
	} else if !exists {
		return 0, errors.New("unable to locate identity in database")
	}
	// Fetch engine details
	engine, exists, err := s.ds.ReadEngine(data.ById(engineId))
	if err != nil {
		return 0, errors.Wrap(err, "reading engine from database")
	} else if !exists {
		return 0, errors.New("unable to locate engine in database")
	}
	// FIXME implement keytab generation on the fly
	keytabPath := path.Join(s.workingDir, fs.KTDir, keytab)
	// Start cluster in yarn
	appId, address, out, err := yarn.StartCloud(size, s.kerberosEnabled, memory, clusterName,
		engine.Location, identity.Name, keytabPath)
	if err != nil {
		return 0, errors.Wrap(err, "starting yarn cluster")
	}
	// Create cluster
	clusterId, err := s.ds.CreateCluster(clusterName, s.ds.ClusterType.Yarn,
		data.WithYarnDetail(engineId, int64(size), appId, memory, out),
		data.WithAddress(address), data.WithState(data.States.Started),
		data.WithPrivilege(pz, data.Owns), data.WithAudit(pz),
	)
	return clusterId, errors.Wrap(err, "creating cluster in database")
}

func (s *Service) StopClusterOnYarn(pz az.Principal, clusterId int64, keytab string) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageCluster); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckOwns(s.ds.EntityType.Cluster, clusterId); err != nil {
		return errors.Wrap(err, "checking ownership")
	}
	// Fetch cluster details
	cluster, exists, err := s.ds.ReadCluster(data.ById(clusterId))
	if err != nil {
		return errors.Wrap(err, "reading cluster from database")
	} else if !exists {
		return errors.New("failed locating cluster")
	}
	if cluster.ClusterTypeId != s.ds.ClusterType.Yarn {
		return errors.New("cluster was not started through YARN")
	}
	// Fetch yarn details
	yarnDetails, exists, err := s.ds.ReadClusterYarnDetail(data.ById(cluster.DetailId.Int64))
	if err != nil {
		return errors.Wrap(err, "reading yarn details from cluster")
	} else if !exists {
		return errors.New("failed locating yarn details")
	}
	// Fetch identity details
	identity, exists, err := s.ds.ReadIdentity(data.ById(pz.Id()))
	if err != nil {
		return errors.Wrap(err, "reading identity from cluster")
	} else if !exists {
		return errors.New("failed locating identity")
	}
	// FIXME implement keytab generation on the fly
	keytabPath := path.Join(s.workingDir, fs.KTDir, keytab)
	// Stop clouds
	if err := yarn.StopCloud(s.kerberosEnabled, cluster.Name, yarnDetails.ApplicationId,
		yarnDetails.OutputDir, identity.Name, keytabPath); err != nil {
		return errors.Wrap(err, "stopping cluster")
	}
	// Delete cluster
	return errors.Wrap(s.ds.DeleteCluster(clusterId, data.WithAudit(pz)), "deleting cluster from database")
}

func (s *Service) viewCluster(pz az.Principal, clusterId int64) (data.Cluster, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewCluster); err != nil {
		return data.Cluster{}, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Cluster, clusterId); err != nil {
		return data.Cluster{}, errors.Wrap(err, "checking view privileges")
	}
	// Fetch cluster details
	cluster, exists, err := s.ds.ReadCluster(data.ById(clusterId))
	if err != nil {
		return data.Cluster{}, errors.Wrap(err, "reading cluster from database")
	} else if !exists {
		return data.Cluster{}, errors.New("unable to locate cluster")
	}
	return cluster, nil
}

func (s *Service) GetCluster(pz az.Principal, clusterId int64) (*web.Cluster, error) {
	// Fetch cluster
	cluster, err := s.viewCluster(pz, clusterId)
	return toCluster(cluster), err
}

func (s *Service) GetClusterOnYarn(pz az.Principal, clusterId int64) (*web.YarnCluster, error) {
	// Fetch cluster
	cluster, err := s.viewCluster(pz, clusterId)
	if err != nil {
		return nil, err
	}
	// Fetch yarn details
	yarnDetails, exists, err := s.ds.ReadClusterYarnDetail(data.ById(cluster.DetailId.Int64))
	if err != nil {
		return nil, errors.Wrap(err, "reading yarn details from cluster")
	} else if !exists {
		return nil, errors.New("unable to locate yarn details")
	}

	return toYarnCluster(cluster, yarnDetails), nil
}

// XXX: will this change the API?
// func (s *Service) GetClusters(pz az.Principal, offset, limit int64) ([]*web.Cluster, error) {
func (s *Service) GetClusters(pz az.Principal, offset, limit uint) ([]*web.Cluster, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewCluster); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetch clusters
	clusters, err := s.ds.ReadClusters(data.ByPrivilege(pz), data.WithOffset(offset), data.WithLimit(limit))
	return toClusters(clusters), errors.Wrap(err, "reading clusters from database")
}

// XXX will this break the API?
// func (s *Service) GetClusterStatus(pz az.Principal, cloudId int64) (*web.ClusterStatus, error) { // Only called if cloud status != found
func (s *Service) GetClusterStatus(pz az.Principal, clusterId int64) (*web.ClusterStatus, error) { // Only called if cloud status != found
	// Fetch cluster
	cluster, err := s.viewCluster(pz, clusterId)
	if err != nil {
		return nil, err
	}
	// Fetch from h2o
	h2o := h2ov3.NewClient(cluster.Address.String)
	status, err := h2o.GetCloudStatus()
	if err != nil { // Do not bail out, report Unknown status instead
		return &web.ClusterStatus{Status: "Unknown"}, nil
	}
	// Fetch stats
	var totCPUs, allCPUs int32
	var totMemory int64
	for _, node := range status.Nodes {
		totMemory += node.MaxMem
		totCPUs += node.NumCpus
		allCPUs += node.CpusAllowed
	}

	return &web.ClusterStatus{
		status.Version, "Healthy",
		toSizeBytes(totMemory), int(totCPUs), int(allCPUs),
	}, nil
}

func (s *Service) DeleteCluster(pz az.Principal, clusterId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageCluster); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckOwns(s.ds.EntityType.Cluster, clusterId); err != nil {
		return errors.Wrap(err, "checking ownership")
	}
	// Fetch cluster details
	cluster, exists, err := s.ds.ReadCluster(data.ById(clusterId))
	if err != nil {
		return errors.Wrap(err, "reading cluster from database")
	} else if !exists {
		return errors.New("unable to locate cluster")
	}
	if cluster.State != data.States.Stopped {
		return errors.New("cannot delete a running cluster")
	}
	// Delete clsuter
	return errors.Wrap(s.ds.DeleteCluster(clusterId, data.WithAudit(pz)), "deleting cluster")
}

// --- ------- ---
// --- ------- ---
// --- Project ---
// --- ------- ---
// --- ------- ---

func (s *Service) CreateProject(pz az.Principal, name, description, modelCategory string) (int64, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageProject); err != nil {
		return 0, errors.Wrap(err, "checking permission")
	}
	// Create project
	projectId, err := s.ds.CreateProject(name, description, modelCategory,
		data.WithPrivilege(pz, data.Owns), data.WithAudit(pz),
	)
	return projectId, errors.Wrap(err, "creating project in database")
}

func (s *Service) GetProjects(pz az.Principal, offset, limit uint) ([]*web.Project, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewProject); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetch projects
	projects, err := s.ds.ReadProjects(data.ByPrivilege(pz), data.WithOffset(offset), data.WithLimit(limit))
	return toProjects(projects), errors.Wrap(err, "reading projects from database")
}

func (s *Service) viewProject(pz az.Principal, projectId int64) (data.Project, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewProject); err != nil {
		return data.Project{}, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Project, projectId); err != nil {
		return data.Project{}, errors.Wrap(err, "checking view privileges")
	}
	// Fetch project details
	project, exists, err := s.ds.ReadProject(data.ById(projectId))
	if err != nil {
		return data.Project{}, errors.Wrap(err, "reading project from database")
	} else if !exists {
		return data.Project{}, errors.New("unable to locate project")
	}
	return project, nil
}

func (s *Service) GetProject(pz az.Principal, projectId int64) (*web.Project, error) {
	// Fetch project
	project, err := s.viewProject(pz, projectId)
	return toProject(project), err
}

func (s *Service) DeleteProject(pz az.Principal, projectId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageProject); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckOwns(s.ds.EntityType.Project, projectId); err != nil {
		return errors.Wrap(err, "checking ownership")
	}
	// Checks before deletion
	if _, exists, err := s.ds.ReadProject(data.ById(projectId)); err != nil {
		return errors.Wrap(err, "reading project from database")
	} else if !exists {
		return errors.New("unable to locate project")
	}
	if _, exists, err := s.ds.ReadModel(data.ByProjectId(projectId)); err != nil {
		return errors.Wrap(err, "reading model from database")
	} else if exists {
		return errors.New("unable to delete a project with models")
	}

	return s.ds.DeleteProject(projectId, data.WithAudit(pz))
}

// --- ----- ---
// --- ----- ---
// --- Model ---
// --- ----- ---
// --- ----- ---

func (s *Service) ImportModelFromCluster(pz az.Principal, clusterId, projectId int64, modelKey, modelName string) (int64, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageModel); err != nil {
		return 0, errors.Wrap(err, "checking permission")
	}
	// Fetch cluster and project
	cluster, err := s.viewCluster(pz, clusterId)
	if err != nil {
		return 0, err
	}
	if _, err := s.viewProject(pz, projectId); err != nil {
		return 0, err
	}
	// Get model from cloud
	h2o := h2ov3.NewClient(cluster.Address.String)
	rawModel, r, err := h2o.GetModelsFetch(modelKey)
	if err != nil {
		return 0, errors.Wrap(err, "fetching model from H2O")
	}
	m := r.Models[0]
	metricsFunc, err := createMetrics(string(m.Output.ModelCategory), m.Output.TrainingMetrics)
	// Create Model
	modelId, err := s.ds.CreateModel(modelName, modelKey, m.AlgoFullName,
		string(m.Output.ModelCategory), m.ResponseColumnName,
		data.WithProjectId(projectId), data.WithClusterId(clusterId),
		data.WithRawSchema(string(rawModel), "1"), metricsFunc,
		data.WithPrivilege(pz, data.Owns), data.WithAudit(pz),
	)
	return modelId, errors.Wrap(err, "creating model in database")
}

func createMetrics(category string, metrics *bindings.ModelMetrics) (data.QueryOpt, error) {
	switch category {
	case "Binomial":
		return data.WithBinomialModel()
	case "Multinomial":
		return data.WithMultinomialModel()
	case "Regression":
		return data.WithRegressionModel()
	}
	return errors.New("unsupported model category:", category)
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

// //
// // Routines to convert DB structs into API structs
// //

func toCluster(c data.Cluster) *web.Cluster {
	return &web.Cluster{
		c.Id,
		c.Name,
		c.ClusterTypeId,
		c.DetailId.Int64,
		c.Address.String,
		c.State,
		toTimestamp(c.Created),
	}
}

func toClusters(cs []data.Cluster) []*web.Cluster {
	ar := make([]*web.Cluster, len(cs))
	for i, c := range cs {
		ar[i] = toCluster(c)
	}
	return ar
}

func toYarnCluster(c data.Cluster, y data.ClusterYarnDetail) *web.YarnCluster {
	return &web.YarnCluster{
		c.Id,
		y.EngineId,
		int(y.Size),
		y.ApplicationId,
		y.Memory,
		y.Username,
	}
}

func toProject(p data.Project) *web.Project {
	return &web.Project{
		p.Id,
		p.Name,
		p.Description,
		p.ModelCategory,
		toTimestamp(p.Created),
	}
}

func toProjects(ps []data.Project) []*web.Project {
	array := make([]*web.Project, len(ps))
	for i, p := range ps {
		array[i] = toProject(p)
	}
	return array
}
