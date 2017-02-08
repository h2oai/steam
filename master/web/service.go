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
	"crypto/tls"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/user"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/h2oai/steam/bindings"
	"github.com/h2oai/steam/lib/fs"
	"github.com/h2oai/steam/lib/haproxy"
	"github.com/h2oai/steam/lib/kerberos"
	"github.com/h2oai/steam/lib/ldap"
	"github.com/h2oai/steam/lib/svc"
	"github.com/h2oai/steam/lib/yarn"
	"github.com/h2oai/steam/master/auth"
	"github.com/h2oai/steam/master/az"
	"github.com/h2oai/steam/master/data"
	"github.com/h2oai/steam/srv/compiler"
	"github.com/h2oai/steam/srv/h2ov3"
	"github.com/h2oai/steam/srv/web"

	"github.com/pkg/errors"
)

type Service struct {
	version                   string
	workingDir                string
	ds                        *data.Datastore
	tlsConfig                 *tls.Config
	compilationServiceAddress string
	scoringServiceAddress     string
	clusterProxyAddress       string
	scoringServicePortMin     int
	scoringServicePortMax     int
	kerberosEnabled           bool
}

func NewService(
	version, workingDir string,
	ds *data.Datastore,
	tlsConfig *tls.Config,
	compilationServiceAddress, scoringServiceAddress, clusterProxyAddress string,
	scoringServicePortsRange [2]int,
) *Service {
	kerberos, err := ViewGlobalKerberos(ds)
	if err != nil {
		// FIXME: THIS SHOULD BE HANDLED MORE GRACEFULLY
		panic(err)
	}
	return &Service{
		version, workingDir,
		ds,
		tlsConfig,
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
	var authType string
	if auth, exists, err := s.ds.ReadAuthentication(data.ByEnabled); err != nil {
		return nil, errors.Wrap(err, "reading authentication from database")
	} else if exists {
		authType = auth.Key
	} else {
		authType = data.LocalAuth
	}

	permissions, err := s.ds.ReadPermissions(data.ForIdentity(pz.Id()))
	if err != nil {
		return nil, errors.Wrap(err, "reading identity from database")
	}
	return &web.Config{
		Version:             s.version,
		AuthenticationType:  authType,
		KerberosEnabled:     s.kerberosEnabled,
		ClusterProxyAddress: s.clusterProxyAddress,
		Username:            pz.Name(),
		Permissions:         toPermissions(permissions),
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
	h := h2ov3.NewClient(address, "/", "")
	cloud, err := h.GetCloudStatus("")
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

func (s Service) reloadProxyConf(name string) {
	clusters, err := s.ds.ReadClusters(data.ByState(data.States.Started))
	if err != nil {
		log.Println("Failed to read clusters.")
	}
	proxyClusters := make([]data.Cluster, 0)
	for _, cluster := range clusters {
		if cluster.ContextPath.Valid {
			proxyClusters = append(proxyClusters, cluster)
		}
	}

	if err := haproxy.Reload(proxyClusters, s.clusterProxyAddress, fs.GetAssetsPath(s.workingDir, "cert.pem")); err != nil {
		log.Println("Failed to reload proxy configuration.")
	}
}

func (s *Service) StartClusterOnYarn(pz az.Principal, clusterName string, engineId int64, size int, memory string, secure bool, keytab string) (int64, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageCluster); err != nil {
		return 0, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckPermission(s.ds.Permission.ViewEngine); err != nil {
		return 0, errors.Wrap(err, "checking permission")
	}
	// Validate name
	nameRegEx := `^[a-zA-Z][a-zA-Z0-9_\-]*$`
	if !regexp.MustCompile(nameRegEx).MatchString(clusterName) {
		return 0, fmt.Errorf("invalid cluster name: %s. Must match %s", clusterName, nameRegEx)
	}
	// Check that name is unique to user that are still running
	_, exists, err := s.ds.ReadCluster(data.ByName(clusterName), data.ByState(data.States.Started),
		data.ByPrivilege(pz))
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
	// Check SSL file
	if _, err := os.Stat(fs.GetAssetsPath(s.workingDir, "cert.pem")); os.IsNotExist(err) {
		return 0, errors.New("SSL \"" + fs.GetAssetsPath(s.workingDir, "cert.pem") + "\" cert file does not exist")
	}
	// Get UID and GID for impersonation
	uid, gid, err := getUser(pz.Name())
	if err != nil {
		return 0, errors.Wrap(err, "get user")
	}
	// Write keytab if not exists already
	var keytabPath string
	if s.kerberosEnabled {
		kt, err := s.viewKeytab(pz)
		if err != nil {
			return 0, errors.Wrap(err, "viewing keytab")
		}
		keytabPath, err = kerberos.WriteKeytab(kt, s.workingDir, int(uid), int(gid))
		if err != nil {
			return 0, errors.Wrap(err, "writing keytab file")
		}
	}
	// Start cluster in yarn
	appId, address, out, token, contextPath, err := yarn.StartCloud(size, s.kerberosEnabled, memory,
		clusterName, engine.Location, identity.Name, keytabPath, secure, uid, gid)
	if err != nil {
		return 0, errors.Wrap(err, "starting yarn cluster")
	}
	// Create cluster
	clusterId, err := s.ds.CreateCluster(clusterName, s.ds.ClusterType.Yarn,
		data.WithYarnDetail(engineId, int64(size), identity.Name, appId, memory, out, contextPath, token),
		data.WithAddress(address), data.WithState(data.States.Started),
		data.WithPrivilege(pz, data.Owns), data.WithAudit(pz),
	)
	s.reloadProxyConf(identity.Name)
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
	// Fetch identity details
	identity, exists, err := s.ds.ReadIdentity(data.ById(pz.Id()))
	if err != nil {
		return errors.Wrap(err, "reading identity from cluster")
	} else if !exists {
		return errors.New("failed locating identity")
	}
	// FIXME: if cluster status is not started assume it failed
	if cluster.State != data.States.Started {
		err = s.ds.DeleteCluster(clusterId, data.WithAudit(pz))
		s.reloadProxyConf(identity.Name)
		return errors.Wrap(err, "deleting cluster from database")
	}
	// Fetch yarn details
	yarnDetails, exists, err := s.ds.ReadClusterYarnDetail(data.ById(cluster.DetailId.Int64))
	if err != nil {
		return errors.Wrap(err, "reading yarn details from cluster")
	} else if !exists {
		return errors.New("failed locating yarn details")
	}
	// Get UID and GID for impersonation
	uid, gid, err := getUser(pz.Name())
	if err != nil {
		return errors.Wrap(err, "get user")
	}
	var keytabPath string
	if s.kerberosEnabled {
		// Write keytab if not exists already
		kt, err := s.viewKeytab(pz)
		if err != nil {
			return errors.Wrap(err, "viewing keytab")
		}
		keytabPath, err = kerberos.WriteKeytab(kt, s.workingDir, int(uid), int(gid))
		if err != nil {
			return errors.Wrap(err, "writing keytab file")
		}
	}
	// Stop clouds
	if err := yarn.StopCloud(s.kerberosEnabled, cluster.Name, yarnDetails.ApplicationId,
		yarnDetails.OutputDir, identity.Name, keytabPath, uid, gid); err != nil {
		return errors.Wrap(err, "stopping cluster")
	}
	// Delete cluster
	err = s.ds.DeleteCluster(clusterId, data.WithAudit(pz))
	s.reloadProxyConf(identity.Name)
	return errors.Wrap(err, "deleting cluster from database")
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
	h2o := h2ov3.NewClient(cluster.Address.String, cluster.ContextPath.String, cluster.Token.String)
	status, err := h2o.GetCloudStatus(cluster.Token.String)
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
	return errors.Wrap(s.ds.DeleteCluster(clusterId, data.WithAudit(pz)), "deleting cluster from database")
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
	h2o := h2ov3.NewClient(cluster.Address.String, cluster.ContextPath.String, cluster.Token.String)
	rawModel, r, err := h2o.GetModelsFetch(modelKey)
	if err != nil {
		return 0, errors.Wrap(err, "fetching model from H2O")
	}
	m := r.Models[0]
	metricsFunc, err := setMetrics(string(m.Output.ModelCategory), m.Output.TrainingMetrics)
	if err != nil {
		return 0, errors.Wrap(err, "setting model metrics type")
	}
	// Create Model
	modelId, err := s.ds.CreateModel(modelName, modelKey, m.AlgoFullName,
		string(m.Output.ModelCategory), m.ResponseColumnName,
		data.WithProjectId(projectId), data.WithCluster(clusterId, cluster.Name),
		data.WithRawSchema(string(rawModel), "1"), metricsFunc,
		data.WithPrivilege(pz, data.Owns), data.WithAudit(pz),
	)
	return modelId, errors.Wrap(err, "creating model in database")
}

func setMetrics(category string, metrics *bindings.ModelMetrics) (data.QueryOpt, error) {
	switch category {
	case "Binomial":
		return data.WithBinomialModel(metrics.Mse, metrics.R2, metrics.Logloss, metrics.Auc, metrics.Gini), nil
	case "Multinomial":
		return data.WithMultinomialModel(metrics.Mse, metrics.R2, metrics.Logloss), nil
	case "Regression":
		return data.WithRegressionModel(metrics.Mse, metrics.R2, metrics.MeanResidualDeviance), nil
	}
	return nil, fmt.Errorf("unsupported model category: %s", category)
}

func (s *Service) CheckMojo(pz az.Principal, algo string) (bool, error) {
	switch algo {
	case "Gradient Boosting Method", "Distributed Random Forest", "Deep Water":
		return true, nil
	}
	return false, nil
}

func (s *Service) viewModel(pz az.Principal, modelId int64) (data.Model, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageModel); err != nil {
		return data.Model{}, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Model, modelId); err != nil {
		return data.Model{}, errors.Wrap(err, "checking view privileges")
	}
	// Fetch model details
	model, exists, err := s.ds.ReadModel(data.ById(modelId))
	if err != nil {
		return data.Model{}, errors.Wrap(err, "reading model from database")
	} else if !exists {
		return data.Model{}, errors.New("unable to locate model")
	}
	return model, nil
}

func (s *Service) GetModel(pz az.Principal, modelId int64) (*web.Model, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewModel); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Model, modelId); err != nil {
		return nil, errors.Wrap(err, "checking privileges")
	}
	// Fetch model
	model, exists, err := s.ds.ReadLabelModel(data.ById(modelId))
	if err != nil {
		return nil, errors.Wrap(err, "reading model from database")
	} else if !exists {
		return nil, errors.New("unable to locate model")
	}
	return toModel(model), nil
}

// FIXME: should be GetModelsByProject
func (s *Service) GetModels(pz az.Principal, projectId int64, offset, limit uint) ([]*web.Model, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewModel); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetch models
	models, err := s.ds.ReadLabelModels(data.ByPrivilege(pz), data.ByProjectId(projectId),
		data.WithOffset(offset), data.WithLimit(limit),
	)
	return toModels(models), errors.Wrap(err, "reading models from database")
}

func (s *Service) GetModelsFromCluster(pz az.Principal, clusterId int64, frameKey string) ([]*web.Model, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewCluster); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.Permission.ViewCluster, clusterId); err != nil {
		return nil, errors.Wrap(err, "checking view privileges")
	}
	// Fetch cluster details
	cluster, exists, err := s.ds.ReadCluster(data.ById(clusterId))
	if err != nil {
		return nil, errors.Wrap(err, "reading cluster from database")
	} else if !exists {
		return nil, errors.New("unable to locate cluster")
	}
	// H2O connection for models
	h2o := h2ov3.NewClient(cluster.Address.String, cluster.ContextPath.String, cluster.Token.String)
	_, frame, err := h2o.GetFramesFetch(frameKey, true)
	return h2oToModels(frame.CompatibleModels), errors.Wrap(err, "fetching H2O frame")
}

func (s *Service) FindModelsCount(pz az.Principal, projectId int64) (int64, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewProject); err != nil {
		return 0, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Project, projectId); err != nil {
		return 0, errors.Wrap(err, "checking view privileges")
	}
	// Fetch models count
	count, err := s.ds.Count("model", data.ByProjectId(projectId))
	return count, errors.Wrap(err, "reading models from database")
}

// TODO: hardcoded; should be determined by h2o metrics
func (s *Service) GetAllBinomialSortCriteria(pz az.Principal) ([]string, error) {
	return []string{"mse", "r_squared", "logloss", "auc", "gini"}, nil
}

func (s *Service) FindModelsBinomial(pz az.Principal, projectId int64, namePart, sortBy string, ascending bool, offset, limit uint) ([]*web.BinomialModel, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewModel); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetch project details
	if _, err := s.viewProject(pz, projectId); err != nil {
		return nil, errors.Wrap(err, "checking view privilege")
	}
	// Fetch Binomial Models
	models, err := s.ds.ReadBinomialModels(data.ByPrivilege(pz), data.ByProjectId(projectId),
		data.WithFilterByName(namePart), data.WithOrderBy(sortBy, ascending),
		data.WithOffset(offset), data.WithLimit(limit),
	)
	return toBinomialModels(models), errors.Wrap(err, "reading binomial models from database")
}

func (s *Service) GetModelBinomial(pz az.Principal, modelId int64) (*web.BinomialModel, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewModel); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Model, modelId); err != nil {
		return nil, errors.Wrap(err, "checking view privileges")
	}
	// Fetch Binomial Model
	model, exists, err := s.ds.ReadBinomialModel(data.ById(modelId))
	if err != nil {
		return nil, errors.Wrap(err, "reading binomial model from database")
	} else if !exists {
		return nil, errors.New("unable to locate binomial model")
	}
	return toBinomialModel(model), nil
}

// TODO: hardcoded; should be determined by h2o metrics
func (s *Service) GetAllMultinomialSortCriteria(pz az.Principal) ([]string, error) {
	return []string{"mse", "r_squared", "logloss"}, nil
}

func (s *Service) FindModelsMultinomial(pz az.Principal, projectId int64, namePart, sortBy string, ascending bool, offset, limit uint) ([]*web.MultinomialModel, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewModel); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetch project details
	if _, err := s.viewProject(pz, projectId); err != nil {
		return nil, errors.Wrap(err, "checking view privilege")
	}
	// Fetch Multinomial Models
	models, err := s.ds.ReadMultinomialModels(data.ByPrivilege(pz), data.ByProjectId(projectId),
		data.WithFilterByName(namePart), data.WithOrderBy(sortBy, ascending),
		data.WithOffset(offset), data.WithLimit(limit),
	)
	return toMultinomialModels(models), errors.Wrap(err, "reading multinomial models from database")
}

func (s *Service) GetModelMultinomial(pz az.Principal, modelId int64) (*web.MultinomialModel, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewModel); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Model, modelId); err != nil {
		return nil, errors.Wrap(err, "checking view privileges")
	}
	// Fetch Multinomial Model
	model, exists, err := s.ds.ReadMultinomialModel(data.ById(modelId))
	if err != nil {
		return nil, errors.Wrap(err, "reading multinomial model from database")
	} else if !exists {
		return nil, errors.New("unable to locate multinomial model")
	}
	return toMultinomialModel(model), nil
}

func (s *Service) FindModelsRegression(pz az.Principal, projectId int64, namePart, sortBy string, ascending bool, offset, limit uint) ([]*web.RegressionModel, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewModel); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetch project details
	if _, err := s.viewProject(pz, projectId); err != nil {
		return nil, errors.Wrap(err, "checking view privilege")
	}
	// Fetch Regression Models
	models, err := s.ds.ReadRegressionModels(data.ByPrivilege(pz), data.ByProjectId(projectId),
		data.WithFilterByName(namePart), data.WithOrderBy(sortBy, ascending),
		data.WithOffset(offset), data.WithLimit(limit),
	)
	return toRegressionModels(models), errors.Wrap(err, "reading regression models from database")
}

// TODO: hardcoded; should be determined by h2o metrics
func (s *Service) GetAllRegressionSortCriteria(pz az.Principal) ([]string, error) {
	return []string{"mse", "r_squared", "mean_residual_deviance"}, nil
}

func (s *Service) GetModelRegression(pz az.Principal, modelId int64) (*web.RegressionModel, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewModel); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Model, modelId); err != nil {
		return nil, errors.Wrap(err, "checking view privileges")
	}
	// Fetch Regression Model
	model, exists, err := s.ds.ReadRegressionModel(data.ById(modelId))
	if err != nil {
		return nil, errors.Wrap(err, "reading regression model from database")
	} else if !exists {
		return nil, errors.New("unable to locate regression model")
	}
	return toRegressionModel(model), nil
}

func (s *Service) ImportModelPojo(pz az.Principal, modelId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageModel); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Model, modelId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	// Fetch model details
	model, exists, err := s.ds.ReadModel(data.ById(modelId))
	if err != nil {
		return errors.Wrap(err, "reading model from database")
	} else if !exists {
		return errors.New("unable to locate model")
	}
	// Fetch cluster details
	cluster, err := s.viewCluster(pz, model.ClusterId)
	if err != nil {
		return err
	}
	// Get Pojo from H2O
	h2o := h2ov3.NewClient(cluster.Address.String, cluster.ContextPath.String, cluster.Token.String)
	modelPath := fs.GetModelPath(s.workingDir, modelId)
	javaModelPath, err := h2o.ExportJavaModel(model.ModelKey, modelPath)
	if err != nil {
		return errors.Wrap(err, "exporting java model from H2O")
	}
	if _, err := h2o.ExportGenModel(modelPath); err != nil {
		return errors.Wrap(err, "exporting java dependency from H2O")
	}
	opts := []data.QueryOpt{data.WithModelObjectType("pojo"), data.WithAudit(pz)}
	if !model.LogicalName.Valid {
		opts = append(opts, data.WithLocation(modelId, fs.GetBasenameWithoutExt(javaModelPath)))
	}
	return errors.Wrap(s.ds.UpdateModel(modelId, opts...), "updating model in database")
}

func (s *Service) ImportModelMojo(pz az.Principal, modelId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageModel); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Model, modelId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	// Fetch model details
	model, exists, err := s.ds.ReadModel(data.ById(modelId))
	if err != nil {
		return errors.Wrap(err, "reading model from database")
	} else if !exists {
		return errors.New("unable to locate model")
	}
	// Verify model CanMojo
	if ok, _ := s.CheckMojo(pz, model.Algorithm); !ok {
		return fmt.Errorf("unable to import mojo from model of type '%s'", model.Algorithm)
	}
	// Fetch cluster details
	cluster, err := s.viewCluster(pz, model.ClusterId)
	if err != nil {
		return err
	}
	// Get Pojo from H2O
	h2o := h2ov3.NewClient(cluster.Address.String, cluster.ContextPath.String, cluster.Token.String)
	modelPath := fs.GetModelPath(s.workingDir, modelId)
	javaModelPath, err := h2o.ExportMOJO(model.ModelKey, modelPath)
	if err != nil {
		return errors.Wrap(err, "exporting MOJO from H2O")
	}
	if _, err := h2o.ExportGenModel(modelPath); err != nil {
		return errors.Wrap(err, "exporting java dependency from H2O")
	}
	// External checks for DeepWater
	if model.Algorithm == "Deep Water" {
		if _, err := h2o.ExportDeepWaterAll(modelPath); err != nil {
			return errors.Wrap(err, "exporting Deep Water dependency")
		}
	}
	opts := []data.QueryOpt{data.WithModelObjectType("mojo"), data.WithAudit(pz)}
	if !model.LogicalName.Valid {
		opts = append(opts, data.WithLocation(modelId, fs.GetBasenameWithoutExt(javaModelPath)))
	}
	return errors.Wrap(s.ds.UpdateModel(modelId, opts...), "updating model in database")
}

func (s *Service) DeleteModel(pz az.Principal, modelId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageModel); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckOwns(s.ds.EntityType.Model, modelId); err != nil {
		return errors.Wrap(err, "checking ownership")
	}
	// FIXME delete assets from disk
	// Fetch model details
	_, exists, err := s.ds.ReadModel(data.ById(modelId))
	if err != nil {
		return errors.Wrap(err, "reading model from database")
	} else if !exists {
		return errors.New("unable to locate model")
	}

	services, err := s.ds.ReadServices(data.ByModelId(modelId))
	if err != nil {
		return errors.Wrap(err, "reading services from database")
	}
	for _, service := range services {
		switch service.State {
		case data.States.Stopped: //FIXME: allow for other states other that started/stopped
			return errors.New("unable to delete a model with at least one prediction service")
		}
	}

	return errors.Wrap(s.ds.DeleteModel(modelId, data.WithAudit(pz)), "deleting model from database")
}

// --- ----- ---
// --- ----- ---
// --- Label ---
// --- ----- ---
// --- ----- ---

func (s *Service) CreateLabel(pz az.Principal, projectId int64, name, description string) (int64, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageLabel); err != nil {
		return 0, errors.Wrap(err, "checking permission")
	}
	// Pre-add checks
	name = strings.TrimSpace(name)
	if name == "" {
		return 0, errors.New("label name cannot be empty")
	}
	if _, exists, err := s.ds.ReadLabel(data.ByProjectId(projectId), data.ByName(name)); err != nil {
		return 0, errors.Wrap(err, "reading label from project")
	} else if exists {
		return 0, fmt.Errorf("a label named %q is alread associated with this project", name)
	}
	// Create label
	labelId, err := s.ds.CreateLabel(projectId, name, description,
		data.WithAudit(pz), data.WithPrivilege(pz, data.Owns),
	)
	return labelId, errors.Wrap(err, "creating label in database")
}

func (s *Service) UpdateLabel(pz az.Principal, labelId int64, name, description string) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageLabel); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Label, labelId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	// Fetch label details
	label, exists, err := s.ds.ReadLabel(data.ById(labelId))
	if err != nil {
		return errors.Wrap(err, "reading label from project")
	} else if !exists {
		return errors.New("unable to locate label")
	}
	// Pre-updatechecks
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("label name cannot be empty")
	}
	if _, exists, err := s.ds.ReadLabel(data.ByProjectId(label.ProjectId.Int64), data.ByName(name)); err != nil {
		return errors.Wrap(err, "reading label from project")
	} else if exists {
		return fmt.Errorf("a label named %q is alread associated with this project", name)
	}
	// Update label
	err = s.ds.UpdateLabel(labelId, data.WithName(name), data.WithDescription(description),
		data.WithAudit(pz),
	)
	return errors.Wrap(err, "updating label in database")
}

func (s *Service) DeleteLabel(pz az.Principal, labelId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageLabel); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckOwns(s.ds.EntityType.Label, labelId); err != nil {
		return errors.Wrap(err, "checking ownership")
	}
	// Fetch label details
	if _, exists, err := s.ds.ReadLabel(data.ById(labelId)); err != nil {
		return errors.Wrap(err, "reading label from database")
	} else if !exists {
		return errors.New("unable to locate label")
	}
	return errors.Wrap(s.ds.DeleteLabel(labelId, data.WithAudit(pz)), "deleting label from database")
}

func (s *Service) LinkLabelWithModel(pz az.Principal, labelId, modelId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageLabel); err != nil {
		return errors.Wrap(err, "checking label permission")
	}
	if err := pz.CheckPermission(s.ds.Permission.ManageModel); err != nil {
		return errors.Wrap(err, "checking model permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Model, modelId); err != nil {
		return errors.Wrap(err, "checking model edit privileges")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Label, labelId); err != nil {
		return errors.Wrap(err, "checking label edit privileges")
	}
	// Check pre-link details details
	if _, err := s.viewModel(pz, modelId); err != nil {
		return err
	}
	if label, exists, err := s.ds.ReadLabel(data.ByModelId(modelId)); err != nil {
		return errors.Wrap(err, "reading label from database")
	} else if exists {
		if err := s.ds.UpdateLabel(label.Id.Int64, data.WithNil("model_id"), data.WithUnlinkAudit(pz)); err != nil {
			return errors.Wrap(err, "unlinking label and model in database")
		}
	}
	err := s.ds.UpdateLabel(labelId, data.WithModelId(modelId),
		data.WithLinkAudit(pz),
	)
	return errors.Wrap(err, "linking label and mode in database")
}

func (s *Service) UnlinkLabelFromModel(pz az.Principal, labelId, modelId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageLabel); err != nil {
		return errors.Wrap(err, "checking label permission")
	}
	if err := pz.CheckPermission(s.ds.Permission.ManageModel); err != nil {
		return errors.Wrap(err, "checking model permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Model, modelId); err != nil {
		return errors.Wrap(err, "checking model edit privileges")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Label, labelId); err != nil {
		return errors.Wrap(err, "checking label edit privileges")
	}
	// Fetch label details
	if _, exists, err := s.ds.ReadLabel(data.ByModelId(modelId)); err != nil {
		return errors.Wrap(err, "reading label from database")
	} else if !exists {
		return errors.New("unable to locate label")
	}
	if _, err := s.viewModel(pz, modelId); err != nil {
		return err
	}
	err := s.ds.UpdateLabel(labelId, data.WithModelId(modelId), data.WithUnlinkAudit(pz))
	return errors.Wrap(err, "unlinking label and model in database")
}

func (s *Service) GetLabelsForProject(pz az.Principal, projectId int64) ([]*web.Label, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewLabel); err != nil {
		return nil, errors.Wrap(err, "checking label permission")
	}
	if err := pz.CheckPermission(s.ds.Permission.ViewProject); err != nil {
		return nil, errors.Wrap(err, "checking project permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Project, projectId); err != nil {
		return nil, errors.Wrap(err, "checking label view privileges")
	}
	// Fetch label details
	labels, err := s.ds.ReadLabels(data.ByProjectId(projectId))
	return toLabels(labels), errors.Wrap(err, "reading labels in database")
}

// --- ------- ---
// --- ------- ---
// --- Service ---
// --- ------- ---
// --- ------- ---

// Service helper funcs
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
	return 0, fmt.Errorf("no open port found within range %d:%d", s.scoringServicePortMin, s.scoringServicePortMax)
}

func (s *Service) StartService(pz az.Principal, modelId int64, name, packageName string) (int64, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageService); err != nil {
		return 0, errors.Wrap(err, "checking service permission")
	}
	// Fetch model details
	model, err := s.viewModel(pz, modelId)
	if err != nil {
		return 0, err
	}
	// Start service
	artifact := compiler.ArtifactWar
	if len(packageName) > 0 {
		artifact = compiler.ArtifactPythonWar
	}
	warFilePath, err := compiler.CompileModel(s.compilationServiceAddress, s.workingDir,
		model.ProjectId, model.Id, model.LogicalName.String, model.ModelObjectType.String,
		model.Algorithm, artifact, packageName,
	)
	if err != nil {
		return 0, errors.Wrap(err, "compiling model")
	}
	port, err := s.assignPort()
	if err != nil {
		return 0, errors.Wrap(err, "assigning port")
	}
	pid, err := svc.Start(warFilePath, fs.GetAssetsPath(s.workingDir, "jetty-runner.jar"),
		port, name, pz.Name(),
	)
	if err != nil {
		return 0, errors.Wrap(err, "starting service")
	}
	// Create service
	serviceId, err := s.ds.CreateService(model.ProjectId, model.Id, name,
		data.WithHost(s.scoringServiceAddress), data.WithPort(port),
		data.WithProcessId(pid), data.WithState(data.States.Started),
		data.WithPrivilege(pz, data.Owns), data.WithAudit(pz),
	)
	return serviceId, errors.Wrap(err, "creating service in database")
}

func (s *Service) StopService(pz az.Principal, serviceId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageService); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckOwns(s.ds.EntityType.Service, serviceId); err != nil {
		return errors.Wrap(err, "checking ownership")
	}
	// Fetch cluster detail
	service, exists, err := s.ds.ReadService(data.ById(serviceId))
	if err != nil {
		return errors.Wrap(err, "reading service from database")
	} else if !exists {
		return errors.New("unable to locate service")
	}
	if service.State != data.States.Stopped {
		if err := svc.Stop(int(service.ProcessId.Int64)); err != nil {
			return errors.Wrap(err, "stopping service")
		}
	}
	// Delete cluster
	err = s.ds.DeleteService(serviceId, data.WithAudit(pz))
	return errors.Wrap(err, "deleting service from database")
}

func (s *Service) viewService(pz az.Principal, serviceId int64) (data.Service, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewService); err != nil {
		return data.Service{}, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Service, serviceId); err != nil {
		return data.Service{}, errors.Wrap(err, "checking view privileges")
	}
	// Fetch service details
	service, exists, err := s.ds.ReadService(data.ById(serviceId))
	if err != err {
		return data.Service{}, errors.Wrap(err, "reading service from database")
	} else if !exists {
		return data.Service{}, errors.New("unable to locate service")
	}
	return service, nil
}

func (s *Service) GetService(pz az.Principal, serviceId int64) (*web.ScoringService, error) {
	// Fetch cluster
	service, err := s.viewService(pz, serviceId)
	return toService(service), err
}

func (s *Service) GetServices(pz az.Principal, offset, limit uint) ([]*web.ScoringService, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewService); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetch services
	services, err := s.ds.ReadServices(data.ByPrivilege(pz),
		data.WithOffset(offset), data.WithLimit(limit),
	)
	return toServices(services), errors.Wrap(err, "reading services from database")
}

func (s *Service) GetServicesForProject(pz az.Principal, projectId int64, offset, limit uint) ([]*web.ScoringService, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewService); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetch project and service details
	if _, err := s.viewProject(pz, projectId); err != nil {
		return nil, err
	}
	services, err := s.ds.ReadServices(data.ByProjectId(projectId),
		data.WithOffset(offset), data.WithLimit(limit),
		data.ByPrivilege(pz),
	)
	return toServices(services), errors.Wrap(err, "reading services from database")
}

func (s *Service) GetServicesForModel(pz az.Principal, modelId int64, offset, limit uint) ([]*web.ScoringService, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewService); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetch model and service details
	if _, err := s.viewModel(pz, modelId); err != nil {
		return nil, err
	}
	services, err := s.ds.ReadServices(data.ByModelId(modelId),
		data.WithOffset(offset), data.WithLimit(limit),
		data.ByPrivilege(pz),
	)
	return toServices(services), errors.Wrap(err, "reading services from database")
}

func (s *Service) DeleteService(pz az.Principal, serviceId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageService); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckOwns(s.ds.EntityType.Service, serviceId); err != nil {
		return errors.Wrap(err, "checking ownership")
	}
	// Fetch service
	service, exists, err := s.ds.ReadService(data.ById(serviceId))
	if err != nil {
		return errors.Wrap(err, "reading service from database")
	} else if !exists {
		return errors.New("unable to locate service")
	}
	if service.State == data.States.Started {
		return errors.New("unable to delete a running service")
	}
	err = s.ds.DeleteService(serviceId, data.WithAudit(pz))
	return errors.Wrap(err, "deleting service from database")
}

// --- ------ ---
// --- ------ ---
// --- Engine ---
// --- ------ ---
// --- ------ ---

// viewEngine is a small wrapper around the perm/fetch db calls that returns an
// engine in view only scenarios
func (s *Service) viewEngine(pz az.Principal, engineId int64) (data.Engine, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewEngine); err != nil {
		return data.Engine{}, errors.Wrap(err, "reading permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Engine, engineId); err != nil {
		return data.Engine{}, errors.Wrap(err, "checking view privileges")
	}
	// Fetch engine
	engine, exists, err := s.ds.ReadEngine(data.ById(engineId))
	if err != nil {
		return data.Engine{}, errors.Wrap(err, "reading engine from database")
	} else if !exists {
		return data.Engine{}, errors.New("unable to locate engine")
	}
	return engine, nil
}

func (s *Service) GetEngine(pz az.Principal, engineId int64) (*web.Engine, error) {
	// Fetch engine using view wrapper
	engine, err := s.viewEngine(pz, engineId)
	return toEngine(engine), err
}

func (s *Service) GetEngines(pz az.Principal) ([]*web.Engine, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewEngine); err != nil {
		return nil, errors.Wrap(err, "reading permission")
	}
	// Fetch clusters with privilege
	engines, err := s.ds.ReadEngines()
	return toEngines(engines), errors.Wrap(err, "reading engines from database")
}

func (s *Service) DeleteEngine(pz az.Principal, engineId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageEngine); err != nil {
		return errors.Wrap(err, "reading permission")
	}
	if err := pz.CheckOwns(s.ds.EntityType.Engine, engineId); err != nil {
		return errors.Wrap(err, "checking ownership")
	}
	// Fetch engine details
	_, exists, err := s.ds.ReadEngine(data.ById(engineId))
	if err != nil {
		return errors.Wrap(err, "reading engine fromd database")
	} else if !exists {
		return errors.New("unable to locate engine")
	}
	// Delete engine
	err = s.ds.DeleteEngine(engineId, data.WithAudit(pz))
	return errors.Wrap(err, "deleting engine from database")
}

// --- ---- ---
// --- ---- ---
// --- Role ---
// --- ---- ---
// --- ---- ---

func (s *Service) CreateRole(pz az.Principal, name string, description string) (int64, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageRole); err != nil {
		return 0, errors.Wrap(err, "checking permission")
	}
	// Pre-add checks
	_, exists, err := s.ds.ReadRole(data.ByName(name))
	if err != nil {
		return 0, errors.Wrap(err, "reading roles from database")
	} else if exists {
		return 0, fmt.Errorf("a role with the name '%s' already exists.", name)
	}
	// Create Role
	id, err := s.ds.CreateRole(name, data.WithDescription(description),
		data.WithPrivilege(pz, data.Owns), data.WithAudit(pz),
	)
	return id, errors.Wrap(err, "creating role in database")
}

func (s *Service) GetRoles(pz az.Principal, offset, limit uint) ([]*web.Role, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewRole); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetch roles
	roles, err := s.ds.ReadRoles(data.ByPrivilege(pz),
		data.WithOffset(offset), data.WithLimit(limit),
	)
	return toRoles(roles), errors.Wrap(err, "reading roles from database")
}

func (s *Service) GetRolesForIdentity(pz az.Principal, identityId int64) ([]*web.Role, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewRole); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	if _, err := s.viewIdentity(pz, identityId); err != nil {
		return nil, err
	}
	roles, err := s.ds.ReadRoles(data.ForIdentity(identityId), data.ByPrivilege(pz))
	return toRoles(roles), errors.Wrap(err, "reading roles from database")
}

func (s *Service) viewRole(pz az.Principal, roleId int64) (data.Role, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewRole); err != nil {
		return data.Role{}, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Role, roleId); err != nil {
		return data.Role{}, errors.Wrap(err, "checking view privileges")
	}
	// Fetch role
	role, exists, err := s.ds.ReadRole(data.ById(roleId))
	if err != nil {
		return data.Role{}, errors.Wrap(err, "reading roles from database")
	} else if !exists {
		return data.Role{}, errors.New("unable to locate role")
	}
	return role, nil
}

func (s *Service) GetRole(pz az.Principal, roleId int64) (*web.Role, error) {
	// Fetch using interal wrapper
	role, err := s.viewRole(pz, roleId)
	return toRole(role), err
}

func (s *Service) GetRoleByName(pz az.Principal, name string) (*web.Role, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewRole); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetch role
	role, exists, err := s.ds.ReadRole(data.ByName(name))
	if err != nil {
		return nil, errors.Wrap(err, "reading role from database")
	} else if !exists {
		return nil, errors.New("unable to locate role")
	}
	// Checking privileges
	if err := pz.CheckView(s.ds.EntityType.Role, role.Id); err != nil {
		return nil, errors.Wrap(err, "checking view privileges")
	}
	return toRole(role), nil
}

func (s *Service) UpdateRole(pz az.Principal, roleId int64, name string, description string) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageRole); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Role, roleId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	// Pre-add checks
	if role, exists, err := s.ds.ReadRole(data.ByName(name)); err != nil {
		return errors.Wrap(err, "reading role from database")
	} else if exists && role.Id != roleId {
		return fmt.Errorf("another role with the name '%s' already exists", name)
	}
	if _, exists, err := s.ds.ReadRole(data.ById(roleId)); err != nil {
		return errors.Wrap(err, "reading role from database")
	} else if !exists {
		return errors.New("unable to locate role")
	}
	// Update role
	err := s.ds.UpdateRole(roleId, data.WithName(name), data.WithDescription(description),
		data.WithAudit(pz),
	)
	return errors.Wrap(err, "updating role in database")
}

func (s *Service) LinkRoleWithPermissions(pz az.Principal, roleId int64, permissionIds []int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageRole); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Role, roleId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	if _, exists, err := s.ds.ReadRole(data.ById(roleId)); err != nil {
		return errors.Wrap(err, "reading role from database")
	} else if !exists {
		return errors.New("unable to locate role")
	}
	// Linking role permissions
	err := s.ds.UpdateRole(roleId, data.LinkPermissions(true, permissionIds...), data.WithLinkAudit(pz))
	return errors.Wrap(err, "updating role in database")
}

func (s *Service) LinkRoleWithPermission(pz az.Principal, roleId int64, permissionId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageRole); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Role, roleId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	if _, exists, err := s.ds.ReadRole(data.ById(roleId)); err != nil {
		return errors.Wrap(err, "reading role from database")
	} else if !exists {
		return errors.New("unable to locate role")
	}
	// Linking role permission
	err := s.ds.UpdateRole(roleId, data.LinkPermissions(false, permissionId), data.WithLinkAudit(pz))
	return errors.Wrap(err, "updating role in database")
}

func (s *Service) UnlinkRoleFromPermission(pz az.Principal, roleId int64, permissionId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageRole); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Role, roleId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	if _, exists, err := s.ds.ReadRole(data.ById(roleId)); err != nil {
		return errors.Wrap(err, "reading role from database")
	} else if !exists {
		return errors.New("unable to locate role")
	}
	// Uninking role permission
	err := s.ds.UpdateRole(roleId, data.UnlinkPermissions(permissionId), data.WithUnlinkAudit(pz))
	return errors.Wrap(err, "updating role in database")
}

func (s *Service) DeleteRole(pz az.Principal, roleId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageRole); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckOwns(s.ds.EntityType.Role, roleId); err != nil {
		return errors.Wrap(err, "checking ownership")
	}
	// Fetch role details
	if _, exists, err := s.ds.ReadRole(data.ById(roleId)); err != nil {
		return errors.Wrap(err, "reading role from database")
	} else if !exists {
		return errors.New("unable to locate role")
	}
	// Delete role
	err := s.ds.DeleteRole(roleId, data.WithAudit(pz))
	return errors.Wrap(err, "deleting role from database")
}

// --- --------- ---
// --- --------- ---
// --- Workgroup ---
// --- --------- ---
// --- --------- ---

func (s *Service) CreateWorkgroup(pz az.Principal, name string, description string) (int64, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageWorkgroup); err != nil {
		return 0, errors.Wrap(err, "checking permission")
	}
	// Pre-add checks
	if _, exists, err := s.ds.ReadWorkgroup(data.ByName(name)); err != nil {
		return 0, errors.Wrap(err, "reading workgroup from database")
	} else if exists {
		return 0, fmt.Errorf("another workgroup with the name '%s' alredy exists", name)
	}
	// Create workgroup
	id, err := s.ds.CreateWorkgroup("workgroup", name, data.WithDescription(description),
		data.WithPrivilege(pz, data.Owns), data.WithAudit(pz),
	)
	return id, errors.Wrap(err, "creating workgroup in database")
}

func (s *Service) GetWorkgroups(pz az.Principal, offset, limit uint) ([]*web.Workgroup, error) {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewWorkgroup); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetching workgroup details
	workgroups, err := s.ds.ReadWorkgroups(data.ByPrivilege(pz),
		data.WithOffset(offset), data.WithLimit(limit),
	)
	return toWorkgroups(workgroups), errors.Wrap(err, "reading workgroups from database")
}

func (s *Service) GetWorkgroupsForIdentity(pz az.Principal, identityId int64) ([]*web.Workgroup, error) {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewWorkgroup); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetching identity/workgroup details (using internal wrapper)
	if _, err := s.viewIdentity(pz, identityId); err != nil {
		return nil, err
	}
	workgroups, err := s.ds.ReadWorkgroups(data.ByPrivilege(pz), data.ForIdentity(identityId))
	return toWorkgroups(workgroups), errors.Wrap(err, "reading workgroups from database")
}

func (s *Service) viewWorkgroup(pz az.Principal, workgroupId int64) (data.Workgroup, error) {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewWorkgroup); err != nil {
		return data.Workgroup{}, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Workgroup, workgroupId); err != nil {
		return data.Workgroup{}, errors.Wrap(err, "checking view privileges")
	}
	// Fetching workgroup details
	workgroup, exists, err := s.ds.ReadWorkgroup(data.ById(workgroupId))
	if err != nil {
		return data.Workgroup{}, errors.Wrap(err, "reading workgroup from datbase")
	} else if !exists {
		return data.Workgroup{}, errors.New("unable to locate workgroup")
	}
	return workgroup, nil
}

func (s *Service) GetWorkgroup(pz az.Principal, workgroupId int64) (*web.Workgroup, error) {
	// Fetching using wrapper
	workgroup, err := s.viewWorkgroup(pz, workgroupId)
	return toWorkgroup(workgroup), errors.Wrap(err, "reading workgroup from database")
}

func (s *Service) GetWorkgroupByName(pz az.Principal, name string) (*web.Workgroup, error) {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewWorkgroup); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetching workgroup details
	workgroup, exists, err := s.ds.ReadWorkgroup(data.ByName(name))
	if err != nil {
		return nil, errors.Wrap(err, "reading workgroup from databse")
	} else if !exists {
		return nil, errors.New("unable to locate workgroup")
	}
	// checking privileges
	err = pz.CheckView(s.ds.EntityType.Workgroup, workgroup.Id)
	return toWorkgroup(workgroup), errors.Wrap(err, "checking view privileges")
}

func (s *Service) UpdateWorkgroup(pz az.Principal, workgroupId int64, name string, description string) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageWorkgroup); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Workgroup, workgroupId); err != nil {
		return errors.Wrap(err, "checking edit priviliges")
	}
	// Pre update checks
	if workgroup, exists, err := s.ds.ReadWorkgroup(data.ByName(name)); err != nil {
		return errors.Wrap(err, "reading workgroup from databse")
	} else if exists && workgroup.Id != workgroupId {
		return fmt.Errorf("another workgroup with the name '%s' already exists", name)
	}
	_, exists, err := s.ds.ReadWorkgroup(data.ById(workgroupId))
	if err != nil {
		return errors.Wrap(err, "reading workgroup from databse")
	} else if !exists {
		return errors.New("unable to locate workgroup")
	}
	// Update workgroup
	err = s.ds.UpdateWorkgroup(workgroupId, data.WithName(name), data.WithDescription(description),
		data.WithAudit(pz),
	)
	return errors.Wrap(err, "updating workgroup in database")
}

func (s *Service) DeleteWorkgroup(pz az.Principal, workgroupId int64) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageWorkgroup); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckOwns(s.ds.EntityType.Workgroup, workgroupId); err != nil {
		return errors.Wrap(err, "checking ownership")
	}
	// Pre-delete checks
	if _, exists, err := s.ds.ReadWorkgroup(data.ById(workgroupId)); err != nil {
		return errors.Wrap(err, "reading workgroup from databse")
	} else if !exists {
		return errors.New("unable to locate workgroup")
	}
	// Deleting workgroup
	err := s.ds.DeleteWorkgroup(workgroupId, data.WithAudit(pz))
	return errors.Wrap(err, "deleting workgroup from database")
}

// --- -------- ---
// --- -------- ---
// --- Identity ---
// --- -------- ---
// --- -------- ---

func (s *Service) CreateIdentity(pz az.Principal, name string, password string) (int64, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageIdentity); err != nil {
		return 0, errors.Wrap(err, "checking permission")
	}
	// Prep Identity
	hash, err := auth.HashPassword(password)
	if err != nil {
		return 0, errors.Wrap(err, "hashing password")
	}
	// Create identity
	id, err := s.ds.CreateIdentity(name, data.WithPassword(hash), data.WithDefaultIdentityWorkgroup,
		data.WithPrivilege(pz, data.Owns), data.WithAudit(pz), data.WithSelfView,
	)
	return id, errors.Wrap(err, "creating identity in database")
}

func (s *Service) GetIdentities(pz az.Principal, offset, limit uint) ([]*web.Identity, error) {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewIdentity); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetch identity details
	identities, err := s.ds.ReadIdentities(data.ByPrivilege(pz),
		data.WithOffset(offset), data.WithLimit(limit),
	)
	return toIdentities(identities), errors.Wrap(err, "reading identities from database")
}

func (s *Service) GetIdentitiesForWorkgroup(pz az.Principal, workgroupId int64) ([]*web.Identity, error) {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewIdentity); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Validate workgroup using view wrapper
	if _, err := s.viewWorkgroup(pz, workgroupId); err != nil {
		return nil, err
	}
	// Fetch identity details
	identities, err := s.ds.ReadIdentities(data.ByPrivilege(pz), data.ForWorkgroup(workgroupId))
	return toIdentities(identities), errors.Wrap(err, "reading identites from database")
}

func (s *Service) GetIdentitiesForRole(pz az.Principal, roleId int64) ([]*web.Identity, error) {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewIdentity); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Validate role using view wrapper
	if _, err := s.viewRole(pz, roleId); err != nil {
		return nil, err
	}
	// Fetch identity details
	identities, err := s.ds.ReadIdentities(data.ByPrivilege(pz), data.ForRole(roleId))
	return toIdentities(identities), errors.Wrap(err, "reading identites from database")
}

// // TODO FIX THIS
func (s *Service) GetIdentitiesForEntity(pz az.Principal, entityTypeId, entityId int64) ([]*web.UserRole, error) {
	return nil, nil
}

// 	// Check permission/privileges
// 	if err := pz.CheckPermission(s.ds.Permission.ViewIdentity); err != nil {
// 		return errors.Wrap(err, "checking permission")
// 	}
// 	if err := pz.CheckView(entityTypeId, entityId); err != nil {
// 		return errors.Wrap(err, "checking view privileges")
// 	}
// 	// TODO: validate that entity exists
// 	identities, err := s.ds.ReadIdentities(data.ByPrivilege(pz), data.ForEntity(entityTypeId, entity))
// }

func (s *Service) viewIdentity(pz az.Principal, identityId int64) (data.Identity, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewIdentity); err != nil {
		return data.Identity{}, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Identity, identityId); err != nil {
		return data.Identity{}, errors.Wrap(err, "checking view privileges")
	}
	// Fetch identity details
	identity, exists, err := s.ds.ReadIdentity(data.ById(identityId))
	if err != nil {
		return data.Identity{}, errors.Wrap(err, "reading identity from cluster")
	} else if !exists {
		return data.Identity{}, errors.New("unable to locate identity")
	}
	return identity, nil
}

func (s *Service) GetIdentity(pz az.Principal, identityId int64) (*web.Identity, error) {
	// Fetch identity using view wrapper
	identity, err := s.viewIdentity(pz, identityId)
	return toIdentity(identity), err
}

func (s *Service) GetIdentityByName(pz az.Principal, name string) (*web.Identity, error) {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewIdentity); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetch identity details
	identity, exists, err := s.ds.ReadIdentity(data.ByName(name))
	if err != nil {
		return nil, errors.Wrap(err, "reading identity from datbase")
	} else if !exists {
		return nil, errors.New("unable to locate identity")
	}
	// Check privilege
	err = pz.CheckView(s.ds.EntityType.Identity, identity.Id)
	return toIdentity(identity), errors.Wrap(err, "checking view privileges")
}

func (s *Service) LinkIdentityWithWorkgroup(pz az.Principal, identityId int64, workgroupId int64) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageIdentity); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckPermission(s.ds.Permission.ManageWorkgroup); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Identity, identityId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Workgroup, workgroupId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	// Validate entities
	if _, exists, err := s.ds.ReadIdentity(data.ById(identityId)); err != nil {
		return errors.Wrap(err, "reading identity from datbase")
	} else if !exists {
		return errors.New("unable to locate identity")
	}
	if _, exists, err := s.ds.ReadWorkgroup(data.ById(workgroupId)); err != nil {
		return errors.Wrap(err, "reading workgroup from datbase")
	} else if !exists {
		return errors.New("unable to locate workgroup")
	}
	// Link workgroup to identity
	err := s.ds.UpdateIdentity(identityId, data.LinkWorkgroup(workgroupId),
		data.WithLinkAudit(pz),
	)
	return errors.Wrap(err, "updating identity in database")
}

func (s *Service) UnlinkIdentityFromWorkgroup(pz az.Principal, identityId int64, workgroupId int64) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageIdentity); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckPermission(s.ds.Permission.ManageWorkgroup); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Identity, identityId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Workgroup, workgroupId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	// Validate entities
	if _, exists, err := s.ds.ReadIdentity(data.ById(identityId)); err != nil {
		return errors.Wrap(err, "reading identity from datbase")
	} else if !exists {
		return errors.New("unable to locate identity")
	}
	if _, exists, err := s.ds.ReadWorkgroup(data.ById(workgroupId)); err != nil {
		return errors.Wrap(err, "reading workgroup from datbase")
	} else if !exists {
		return errors.New("unable to locate workgroup")
	}
	// Unlink workgroup to identity
	err := s.ds.UpdateIdentity(identityId, data.UnlinkWorkgroup(workgroupId),
		data.WithUnlinkAudit(pz),
	)
	return errors.Wrap(err, "updating identity in database")
}

func (s *Service) LinkIdentityWithRole(pz az.Principal, identityId int64, roleId int64) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageIdentity); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckPermission(s.ds.Permission.ManageRole); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Identity, identityId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Role, roleId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	// Validate entities
	if _, exists, err := s.ds.ReadIdentity(data.ById(identityId)); err != nil {
		return errors.Wrap(err, "reading identity from datbase")
	} else if !exists {
		return errors.New("unable to locate identity")
	}
	if _, exists, err := s.ds.ReadRole(data.ById(roleId)); err != nil {
		return errors.Wrap(err, "reading role from datbase")
	} else if !exists {
		return errors.New("unable to locate role")
	}
	// Link role to identity
	err := s.ds.UpdateIdentity(identityId, data.LinkRole(roleId, false),
		data.WithLinkAudit(pz),
	)
	return errors.Wrap(err, "updating identity in database")
}

func (s *Service) UnlinkIdentityFromRole(pz az.Principal, identityId int64, roleId int64) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageIdentity); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckPermission(s.ds.Permission.ManageRole); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Identity, identityId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Role, roleId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	// Validate entities
	if _, exists, err := s.ds.ReadIdentity(data.ById(identityId)); err != nil {
		return errors.Wrap(err, "reading identity from datbase")
	} else if !exists {
		return errors.New("unable to locate identity")
	}
	if _, exists, err := s.ds.ReadRole(data.ById(roleId)); err != nil {
		return errors.Wrap(err, "reading role from datbase")
	} else if !exists {
		return errors.New("unable to locate role")
	}
	// Unlink role to identity
	err := s.ds.UpdateIdentity(identityId, data.UnlinkRole(roleId),
		data.WithUnlinkAudit(pz),
	)
	return errors.Wrap(err, "updating identity in database")
}

func (s *Service) UpdateIdentity(pz az.Principal, identityId int64, password string) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageIdentity); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Identity, identityId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	// Validate entities
	if _, exists, err := s.ds.ReadIdentity(data.ById(identityId)); err != nil {
		return errors.Wrap(err, "reading identity from datbase")
	} else if !exists {
		return errors.New("unable to locate identity")
	}
	// Hash password
	hash, err := auth.HashPassword(password)
	if err != nil {
		return errors.Wrap(err, "hashing password")
	}
	// Update identity
	err = s.ds.UpdateIdentity(identityId, data.WithPassword(hash))
	return errors.Wrap(err, "updating identity in database")
}

func (s *Service) ActivateIdentity(pz az.Principal, identityId int64) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageIdentity); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Identity, identityId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	// Update identity
	err := s.ds.UpdateIdentity(identityId, data.WithActivity(true))
	return errors.Wrap(err, "updating identity in database")
}

func (s *Service) DeactivateIdentity(pz az.Principal, identityId int64) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageIdentity); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Identity, identityId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	// Update identity
	err := s.ds.UpdateIdentity(identityId, data.WithActivity(false))
	return errors.Wrap(err, "updating identity in database")
}

// --- ------- ---
// --- ------- ---
// --- Package ---
// --- ------- ---
// --- ------- ---

func (s *Service) CreatePackage(pz az.Principal, projectId int64, name string) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageProject); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Project, projectId); err != nil {
		return errors.Wrap(err, "checking edit privileges")
	}
	// Pre-add Checks
	if err := fs.ValidateName(name); err != nil {
		return errors.Wrap(err, "validating package name")
	}
	packagePath := fs.GetPackagePath(s.workingDir, projectId, name)
	if fs.DirExists(packagePath) {
		return fmt.Errorf("failed creating package directory: %s already exists", name)
	}
	// Create package directory
	return errors.Wrap(fs.Mkdir(packagePath), "creating package directory")
}

func (s *Service) GetPackages(pz az.Principal, projectId int64) ([]string, error) {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewProject); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Project, projectId); err != nil {
		return nil, errors.Wrap(err, "checking view privileges")
	}
	// Fetch project path
	projectPath := fs.GetProjectPath(s.workingDir, projectId)
	if !fs.DirExists(projectPath) {
		return []string{}, nil
	}
	dirs, err := fs.ListDirs(projectPath)
	return dirs, errors.Wrap(err, "listing package directories")
}

func (s *Service) GetPackageDirectories(pz az.Principal, projectId int64, packageName string, relativePath string) ([]string, error) {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewProject); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Project, projectId); err != nil {
		return nil, errors.Wrap(err, "checking view privileges")
	}
	// Fetching pack directories
	packagePath := fs.GetPackagePath(s.workingDir, projectId, packageName)
	if !fs.DirExists(packagePath) {
		return nil, fmt.Errorf("package %s does not exist")
	}
	packageDirPath, err := fs.GetPackageRelativePath(s.workingDir, projectId, packageName, relativePath)
	if err != nil {
		return nil, errors.Wrap(err, "geting relative package path")
	}
	if !fs.DirExists(packageDirPath) {
		return []string{}, nil
	}
	// Fetching directories
	dirs, err := fs.ListDirs(packageDirPath)
	return dirs, errors.Wrap(err, "listing package directories")
}

func (s *Service) GetPackageFiles(pz az.Principal, projectId int64, packageName string, relativePath string) ([]string, error) {
	if err := pz.CheckPermission(s.ds.Permission.ViewProject); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Project, projectId); err != nil {
		return nil, errors.Wrap(err, "checking view privileges")
	}
	// Fetching package paths
	packagePath := fs.GetPackagePath(s.workingDir, projectId, packageName)
	if !fs.DirExists(packagePath) {
		return nil, fmt.Errorf("package %s does not exist", packageName)
	}
	packageDirPath, err := fs.GetPackageRelativePath(s.workingDir, projectId, packageName, relativePath)
	if err != nil {
		return nil, errors.Wrap(err, "getting relative package path")
	}
	if !fs.DirExists(packageDirPath) {
		return []string{}, nil
	}
	// Fetching files
	files, err := fs.ListFiles(packageDirPath)
	return files, errors.Wrap(err, "listing package files")
}

func (s *Service) DeletePackage(pz az.Principal, projectId int64, name string) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageProject); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Project, projectId); err != nil {
		return errors.Wrap(err, "check edit privileges")
	}
	// Fetching package path
	packagePath := fs.GetPackagePath(s.workingDir, projectId, name)
	if !fs.DirExists(packagePath) {
		return fmt.Errorf("package %s does not exist", name)
	}
	// Delete Package
	return errors.Wrap(fs.Rmdir(packagePath), "deleting package")
}

func (s *Service) DeletePackageDirectory(pz az.Principal, projectId int64, packageName string, relativePath string) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageProject); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Project, projectId); err != nil {
		return errors.Wrap(err, "check edit privileges")
	}
	// Fetching package paths
	packagePath := fs.GetPackagePath(s.workingDir, projectId, packageName)
	if !fs.DirExists(packagePath) {
		return fmt.Errorf("package %s does not exist", packageName)
	}
	dirPath, err := fs.GetPackageRelativePath(s.workingDir, projectId, packageName, relativePath)
	if err != nil {
		return errors.Wrap(err, "fetching relative package path")
	}
	if !fs.DirExists(dirPath) {
		return fmt.Errorf("invalid path %s in package %s", relativePath, packageName)
	}
	// Deleting directory
	return errors.Wrap(fs.Rmdir(dirPath), "deleting director")
}

func (s *Service) DeletePackageFile(pz az.Principal, projectId int64, packageName string, relativePath string) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageProject); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Project, projectId); err != nil {
		return errors.Wrap(err, "check edit privileges")
	}
	// Fetching package paths
	packagePath := fs.GetPackagePath(s.workingDir, projectId, packageName)
	if !fs.DirExists(packagePath) {
		return fmt.Errorf("package %s does not exist", packageName)
	}
	filePath, err := fs.GetPackageRelativePath(s.workingDir, projectId, packageName, relativePath)
	if err != nil {
		return errors.Wrap(err, "fetching relative package path")
	}
	if !fs.FileExists(filePath) {
		return fmt.Errorf("invalid path %s in package %s", relativePath, packageName)
	}
	// Deleting file
	return errors.Wrap(fs.Rm(filePath), "deleting file")
}

func (s *Service) SetAttributesForPackage(pz az.Principal, projectId int64, packageName string, attributes string) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageProject); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckEdit(s.ds.EntityType.Project, projectId); err != nil {
		return errors.Wrap(err, "check edit privileges")
	}
	// Setting attributes
	err := fs.SetPackageAttributes(s.workingDir, projectId, packageName, []byte(attributes))
	return errors.Wrap(err, "setting package attributes")
}

func (s *Service) GetAttributesForPackage(pz az.Principal, projectId int64, packageName string) (string, error) {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewProject); err != nil {
		return "", errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Project, projectId); err != nil {
		return "", errors.Wrap(err, "checking view privileges")
	}
	// Fetching package attributes
	b, err := fs.GetPackageAttributes(s.workingDir, projectId, packageName)
	return string(b), errors.Wrap(err, "getting package attributes")
}

// --- ---- ---
// --- ---- ---
// --- Misc ---
// --- ---- ---
// --- ---- ---

func h2oToDataset(f *bindings.FrameBase) *web.Dataset {
	return &web.Dataset{
		Name:      f.FrameId.Name,
		FrameName: f.FrameId.Name,
		CreatedAt: toTimestamp(time.Now()),
	}
}

func h2oToDatasets(fs *bindings.FramesV3) []*web.Dataset {
	ar := make([]*web.Dataset, len(fs.Frames))
	for i, f := range fs.Frames {
		ar[i] = h2oToDataset(f)
	}
	return ar
}

func (s *Service) GetDatasetsFromCluster(pz az.Principal, clusterId int64) ([]*web.Dataset, error) {
	// View cluster using wrapper
	cluster, err := s.viewCluster(pz, clusterId)
	if err != nil {
		return nil, err
	}
	// Start h2o communication
	h2o := h2ov3.NewClient(cluster.Address.String, cluster.ContextPath.String, cluster.Token.String)
	frames, err := h2o.GetFramesList()
	return h2oToDatasets(frames), errors.Wrap(err, "fetching frames from H2O")
}

func (s *Service) GetAllEntityTypes(pz az.Principal) ([]*web.EntityType, error) {
	order := make([]int, 0, len(s.ds.EntityTypeMap))
	for id, _ := range s.ds.EntityTypeMap {
		order = append(order, int(id))
	}
	sort.Ints(order)

	ar := make([]*web.EntityType, 0, len(s.ds.EntityTypeMap))
	for _, i := range order {
		name := s.ds.EntityTypeMap[int64(i)]
		et := web.EntityType{Id: int64(i), Name: name}
		ar = append(ar, &et)
	}
	return ar, nil
}

func (s *Service) GetAllPermissions(pz az.Principal) ([]*web.Permission, error) {
	order := make([]int, 0, len(s.ds.PermissionMap))
	for id, _ := range s.ds.PermissionMap {
		order = append(order, int(id))
	}
	sort.Ints(order)

	ar := make([]*web.Permission, 0, len(s.ds.PermissionMap))
	for _, i := range order {
		pm := s.ds.PermissionMap[int64(i)]
		p := web.Permission{Id: int64(i), Code: pm.Code, Description: pm.Desc}
		ar = append(ar, &p)
	}
	return ar, nil
}

// func (s *Service) GetAllClusterTypes(pz az.Principal) ([]*web.ClusterType, error) {

// 	// No permission checks required

// 	return toClusterTypes(s.ds.ClusterType)
// }

func (s *Service) GetPermissionsForRole(pz az.Principal, roleId int64) ([]*web.Permission, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewRole); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Role, roleId); err != nil {
		return nil, errors.Wrap(err, "checking view privileges")
	}
	// Fetch permissions
	permissions, err := s.ds.ReadPermissions(data.ForRole(roleId))
	return toPermissions(permissions), errors.Wrap(err, "reading permissions from database")
}

func (s *Service) GetPermissionsForIdentity(pz az.Principal, identityId int64) ([]*web.Permission, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewIdentity); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	// Fetch identity via wrapper
	if _, err := s.viewIdentity(pz, identityId); err != nil {
		return nil, err
	}
	// Fetch permissions
	permissions, err := s.ds.ReadPermissions(data.ForIdentity(identityId))
	return toPermissions(permissions), errors.Wrap(err, "reading permissions from database")
}

func (s *Service) ShareEntity(pz az.Principal, kind string, workgroupId, entityTypeId, entityId int64) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.ManagePermission[entityTypeId]); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckPermission(s.ds.Permission.ViewWorkgroup); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckOwns(entityTypeId, entityId); err != nil {
		return errors.Wrap(err, "checking ownership")
	}
	if err := pz.CheckView(s.ds.EntityType.Workgroup, workgroupId); err != nil {
		return errors.Wrap(err, "checking view privileges")
	}
	entityType, ok := s.ds.EntityTypeMap[entityTypeId]
	if !ok {
		return errors.New(fmt.Sprintf("invalid entity type id: %d", entityTypeId))
	}
	// Create privilege/sharing
	_, err := s.ds.CreatePrivilege(kind, pz.Id(), workgroupId, entityType, entityId,
		data.WithAudit(pz),
	)
	return errors.Wrap(err, "creating privilege in database")
}

func (s *Service) GetPrivileges(pz az.Principal, entityTypeId, entityId int64) ([]*web.EntityPrivilege, error) {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.ViewPermission[entityTypeId]); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckPermission(s.ds.Permission.ViewWorkgroup); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(entityTypeId, entityId); err != nil {
		return nil, errors.Wrap(err, "checking view privileges")
	}
	// Read privileges
	entityType, ok := s.ds.EntityTypeMap[entityTypeId]
	if !ok {
		return nil, fmt.Errorf("unable to locate entity type id: %d", entityTypeId)
	}
	entityPrivileges, err := s.ds.ReadEntityPrivileges(
		data.ByEntityType(entityType), data.ByEntityId(entityId),
	)
	return toEntityPrivileges(entityPrivileges), errors.Wrap(err, "reading entity privileges from database")
}

func (s *Service) UnshareEntity(pz az.Principal, kind string, workgroupId, entityTypeId, entityId int64) error {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.ManagePermission[entityTypeId]); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckPermission(s.ds.Permission.ViewWorkgroup); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(entityTypeId, entityId); err != nil {
		return errors.Wrap(err, "checking view privileges")
	}
	if err := pz.CheckView(s.ds.EntityType.Workgroup, workgroupId); err != nil {
		return errors.Wrap(err, "checking view privileges")
	}
	entityType, ok := s.ds.EntityTypeMap[entityTypeId]
	if !ok {
		return fmt.Errorf("unable to locate entity type id: %d", entityTypeId)
	}
	// Delete privileges
	err := s.ds.DeletePrivilege(data.ByPrivilegeType(kind), data.ByWorkgroupId(workgroupId),
		data.ByEntityType(entityType), data.ByEntityId(entityId),
		data.WithUnshareAudit(pz),
	)
	return errors.Wrap(err, "deleting privilege from database")
}

func (s *Service) GetHistory(pz az.Principal, entityTypeId, entityId int64, offset, limit uint) ([]*web.EntityHistory, error) {
	// Check permission/privileges
	if err := pz.CheckPermission(s.ds.ViewPermission[entityTypeId]); err != nil {
		return nil, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(entityTypeId, entityId); err != nil {
		return nil, errors.Wrap(err, "checking view privileges")
	}
	entityType, ok := s.ds.EntityTypeMap[entityTypeId]
	if !ok {
		return nil, fmt.Errorf("unable to locate entity type id: %d", entityTypeId)
	}
	// Fetch history details
	history, err := s.ds.ReadHistories(
		data.ByEntityType(entityType), data.ByEntityId(entityId),
		data.WithOffset(offset), data.WithLimit(limit),
	)
	return toEntityHistories(history), errors.Wrap(err, "reading history from database")
}

type ldapSerialized struct {
	Address                string
	Bind                   string
	UserBaseDn             string
	UserBaseFilter         string
	UserNameAttribute      string
	GroupBaseDN            string
	GroupNameAttribute     string
	GroupNames             []string
	StaticMemberAttribute  string
	SearchRequestSizeLimit int
	SearchRequestTimeLimit int

	ForceBind bool
	Ldaps     bool
}

func configToSerialized(config *web.LdapConfig) ldapSerialized {
	bind := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.BindDn, config.BindPassword)))
	groups := strings.Split(config.GroupNames, ",")

	return ldapSerialized{
		fmt.Sprintf("%s:%d", config.Host, config.Port),
		bind,
		config.UserBaseDn,
		config.UserBaseFilter,
		config.UserNameAttribute,
		config.GroupBaseDn,
		config.GroupNameAttribute,
		groups,
		config.StaticMemberAttribute,
		config.SearchRequestSizeLimit,
		config.SearchRequestTimeLimit,

		config.ForceBind,
		config.Ldaps,
	}
}

func serializedToConfig(config ldapSerialized) (*web.LdapConfig, error) {
	address := strings.Split(config.Address, ":")
	host := address[0]
	port, err := strconv.Atoi(address[1])
	if err != nil {
		return nil, errors.Wrap(err, "convering port to integer")
	}
	b, err := base64.StdEncoding.DecodeString(config.Bind)
	if err != nil {
		return nil, errors.Wrap(err, "decoding bind") //FIXME format error
	}
	bind := strings.Split(string(b), ":")

	return &web.LdapConfig{
		Host:                   host,
		Port:                   port,
		Ldaps:                  config.Ldaps,
		BindDn:                 bind[0],
		UserBaseDn:             config.UserBaseDn,
		UserBaseFilter:         config.UserBaseFilter,
		UserNameAttribute:      config.UserNameAttribute,
		GroupBaseDn:            config.GroupBaseDN,
		GroupNameAttribute:     config.GroupNameAttribute,
		GroupNames:             strings.Join(config.GroupNames, ","),
		StaticMemberAttribute:  config.StaticMemberAttribute,
		SearchRequestSizeLimit: config.SearchRequestSizeLimit,
		SearchRequestTimeLimit: config.SearchRequestTimeLimit,
		ForceBind:              config.ForceBind,
	}, nil
}

func (s *Service) SetLocalConfig(pz az.Principal) error {
	if authentication, exists, err := s.ds.ReadAuthentication(data.ByEnabled); err != nil {
		return errors.Wrap(err, "reading authentication setting from database")
	} else if exists {
		err := s.ds.UpdateAuthentication(authentication.Id, data.WithEnable(false))
		return errors.Wrap(err, "updaing authentication setting from database")
	}
	return nil
}

func (s *Service) SetLdapConfig(pz az.Principal, config *web.LdapConfig) error {
	if !pz.IsAdmin() {
		return errors.New("only admins can edit LDAP settings")
	}
	if strings.TrimSpace(config.BindPassword) == "" {
		return errors.New("bind password cannot be blank")
	}

	raw := configToSerialized(config)
	meta, err := json.Marshal(raw)
	if err != nil {
		return errors.Wrap(err, "serializing config metadata")
	}

	if authentication, exists, err := s.ds.ReadAuthentication(data.ByKey(data.LDAPAuth)); err != nil {
		return errors.Wrap(err, "reading authentication config from database")
	} else if exists {
		err := s.ds.UpdateAuthentication(authentication.Id, data.WithValue(string(meta)), data.WithEnable(true))
		return errors.Wrap(err, "updating authentication config in database")
	}

	_, err = s.ds.CreateAuthentication(data.LDAPAuth, string(meta), data.WithEnable(true))
	return errors.Wrap(err, "creating authentication config in database")
}

func (s *Service) GetLdapConfig(pz az.Principal) (*web.LdapConfig, bool, error) {
	if !pz.IsAdmin() {
		return nil, false, errors.New("only admins can view LDAP settings")
	}

	authentication, exists, err := s.ds.ReadAuthentication(data.ByKey(data.LDAPAuth))
	if err != nil {
		return nil, false, errors.Wrap(err, "reading authentication config from database")
	} else if !exists {
		return nil, false, nil
	}
	// FIXME: HACK to remove escaped quotes from sqlite???
	authentication.Value = strings.Replace(authentication.Value, "\\", "", -1)

	var deserial ldapSerialized
	if err := json.Unmarshal(json.RawMessage(authentication.Value), &deserial); err != nil {
		return nil, false, errors.Wrap(err, "deserializing config metadata")
	}

	config, err := serializedToConfig(deserial)
	return config, exists, err
}

func (s *Service) TestLdapConfig(pz az.Principal, config *web.LdapConfig) (int, []*web.LdapGroup, error) {
	ct, gs, err := ldap.FromConfig(config, s.tlsConfig).TestConfig()
	return ct, toLdapGroup(gs), err
}

func (s *Service) CheckAdmin(pz az.Principal) (bool, error) { return pz.IsAdmin(), nil }

func (s *Service) GetSteamKeytab(pz az.Principal) (*web.Keytab, bool, error) {
	// Only admin can see/modify
	if !pz.IsAdmin() {
		return nil, false, errors.New("user is not an admin")
	}
	// Read keytab
	keytab, exists, err := s.ds.ReadKeytab(data.ByPrincipalSteam)
	if err != nil {
		return nil, false, errors.Wrap(err, "reading keytab from database")
	}
	return toKeytab(keytab), exists, nil
}

func (s *Service) viewKeytab(pz az.Principal) (data.Keytab, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewKeytab); err != nil {
		return data.Keytab{}, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Identity, pz.Id()); err != nil {
		return data.Keytab{}, errors.Wrap(err, "checking view privileges")
	}
	// Read Keytab
	keytab, exists, err := s.ds.ReadKeytab(data.ByIdentityId(pz.Id()))
	if err != nil {
		return data.Keytab{}, errors.Wrap(err, "reading keytab from database")
	} else if !exists {
		return data.Keytab{}, errors.New("unable to locate keytab for user")
	}
	if err := pz.CheckView(s.ds.EntityType.Keytab, keytab.Id); err != nil {
		return data.Keytab{}, errors.Wrap(err, "checking view privileges")
	}
	return keytab, nil
}

func (s *Service) GetUserKeytab(pz az.Principal) (*web.Keytab, bool, error) {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewKeytab); err != nil {
		return nil, false, errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Identity, pz.Id()); err != nil {
		return nil, false, errors.Wrap(err, "checking view privileges")
	}
	// Read Keytab
	keytab, exists, err := s.ds.ReadKeytab(data.ByIdentityId(pz.Id()))
	if err != nil {
		return nil, false, errors.Wrap(err, "reading keytab from database")
	} else if exists {
		if err := pz.CheckView(s.ds.EntityType.Keytab, keytab.Id); err != nil {
			return nil, false, errors.Wrap(err, "checking view privileges")
		}
	}
	return toKeytab(keytab), exists, nil
}

func (s *Service) TestKeytab(pz az.Principal, keytabId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ViewKeytab); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Identity, pz.Id()); err != nil {
		return errors.Wrap(err, "checking owns privileges")
	}
	// Read Keytab
	keytab, exists, err := s.ds.ReadKeytab(data.ById(keytabId))
	if err != nil {
		return errors.Wrap(err, "reading keytab from database")
	} else if !exists {
		return errors.New("cannot locate keytab")
	}
	var name string
	// Final privileges must be handled differently if a steam/user keytab
	if keytab.Principal.Valid { // Is the steam keytab
		if !pz.IsAdmin() {
			return errors.New("user is not a valid admin")
		}
		name = keytab.Principal.String
	} else {
		if err := pz.CheckOwns(s.ds.EntityType.Keytab, keytabId); err != nil {
			return errors.Wrap(err, "checking view privileges")
		}
		name = pz.Name()
	}
	// Get user uid and gid for impersonation of execs
	uid, gid, err := getUser(name)
	if err != nil {
		return errors.Wrap(err, "get user")
	}
	// Generate keytab on disk, if not available
	kpath, err := kerberos.WriteKeytab(keytab, s.workingDir, int(uid), int(gid))
	if err != nil {
		return errors.Wrap(err, "write keytab")
	}
	// Test by performing a kinit
	return kerberos.Kinit(kpath, name, uid, gid)
}

func (s *Service) DeleteKeytab(pz az.Principal, keytabId int64) error {
	// Check permissions/privileges
	if err := pz.CheckPermission(s.ds.Permission.ManageKeytab); err != nil {
		return errors.Wrap(err, "checking permission")
	}
	if err := pz.CheckView(s.ds.EntityType.Identity, pz.Id()); err != nil {
		return errors.Wrap(err, "checking owns privileges")
	}
	// Read Keytab
	keytab, exists, err := s.ds.ReadKeytab(data.ById(keytabId))
	if err != nil {
		return errors.Wrap(err, "reading keytab from database")
	} else if !exists {
		return errors.New("cannot locate keytab")
	}
	// Final privileges must be handled differently if a steam/user keytab
	if keytab.Principal.Valid { // Is the steam keytab
		if !pz.IsAdmin() {
			return errors.New("user is not a valid admin")
		}
	} else if err := pz.CheckOwns(s.ds.EntityType.Keytab, keytabId); err != nil {
		return errors.Wrap(err, "checking view privileges")
	}
	// Delete local Keytab
	if err := kerberos.DeleteKeytab(keytab.Filename, s.workingDir); err != nil {
		return errors.Wrap(err, "deleting local keytab")
	}
	// Delete Keytab
	return s.ds.DeleteKeytab(keytabId)
}

func ViewGlobalKerberos(ds *data.Datastore) (bool, error) {
	meta, exists, err := ds.ReadMeta(data.ByKey("kerberos"))
	if err != nil {
		return false, errors.Wrap(err, "reading meta table")
	} else if !exists {
		return false, nil
	}
	return strconv.ParseBool(meta.Value)
}

func (s *Service) SetGlobalKerberos(pz az.Principal, enabled bool) error {
	// Check permission (only admin)
	if !pz.IsAdmin() {
		return errors.New("user is not an admin")
	}

	meta, exists, err := s.ds.ReadMeta(data.ByKey("kerberos"))
	if err != nil {
		return errors.Wrap(err, "reading meta table")
	} else if exists {
		if err := s.ds.UpdateMeta(meta.Id, data.WithValue(strconv.FormatBool(enabled))); err != nil {
			errors.Wrap(err, "updating kerberos value")
		}
	} else {
		if _, err := s.ds.CreateMeta("kerberos", data.WithValue(strconv.FormatBool(enabled))); err != nil {
			errors.Wrap(err, "creating kerberos value")
		}
	}
	s.kerberosEnabled = enabled
	return nil
}

// --- ---------- ---
// --- ---------- ---
// --- Deprecated ---
// --- ---------- ---
// --- ---------- ---

func (s *Service) BuildModel(pz az.Principal, clusterId int64, datasetId int64, algorithm string) (int64, error) {
	return 0, errors.New("This API call has been deprecated")
}

func (s *Service) BuildModelAuto(pz az.Principal, clusterId int64, dataset, targetName string, maxRunTime int) (*web.Model, error) {
	return nil, errors.New("This API call has been deprecated")
}

func (s *Service) CreateDataset(pz az.Principal, clusterId int64, datasourceId int64, name, description string, responseColumnName string) (int64, error) {
	return 0, errors.New("This API call has been deprecated")
}

func (s *Service) GetDatasets(pz az.Principal, datasourceId int64, offset, limit uint) ([]*web.Dataset, error) {
	return nil, errors.New("This API call has been deprecated")
}

func (s *Service) GetDataset(pz az.Principal, datasetId int64) (*web.Dataset, error) {
	return nil, errors.New("This API call has been deprecated")
}

func (s *Service) UpdateDataset(pz az.Principal, datasetId int64, name, description, responseColumnName string) error {
	return errors.New("This API call has been deprecated")
}

func (s *Service) SplitDataset(pz az.Principal, datasetId int64, ratio1 int, ratio2 int) ([]int64, error) {
	return nil, errors.New("This API call has been deprecated")
}

func (s *Service) DeleteDataset(pz az.Principal, datasetId int64) error {
	return errors.New("This API call has been deprecated")
}

func (s *Service) CreateDatasource(pz az.Principal, projectId int64, name, description, filePath string) (int64, error) {
	return 0, errors.New("This API call has been deprecated")
}

func (s *Service) GetDatasources(pz az.Principal, projectId int64, offset, limit uint) ([]*web.Datasource, error) {
	return nil, errors.New("This API call has been deprecated")
}

func (s *Service) GetDatasource(pz az.Principal, datasourceId int64) (*web.Datasource, error) {
	return nil, errors.New("This API call has been deprecated")
}

func (s *Service) UpdateDatasource(pz az.Principal, datasourceId int64, name, description, filePath string) error {
	return errors.New("This API call has been deprecated")
}

func (s *Service) DeleteDatasource(pz az.Principal, datasourceId int64) error {
	return errors.New("This API call has been deprecated")
}

func (s *Service) GetAllClusterTypes(pz az.Principal) ([]*web.ClusterType, error) {
	return nil, errors.New("This API call has been deprecated")
}

func (s *Service) GetJob(pz az.Principal, clusterId int64, jobName string) (*web.Job, error) {
	return nil, errors.New("This API call has been deprecated")
}

func (s *Service) GetJobs(pz az.Principal, clusterId int64) ([]*web.Job, error) {
	return nil, errors.New("This API call has been deprecated")
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

// Helper func to get uid/gid
func getUser(username string) (uint32, uint32, error) {
	u, err := user.Lookup(username)
	if err != nil {
		return 0, 0, errors.Wrap(err, "user lookup")
	}

	uid, err := strconv.ParseUint(u.Uid, 10, 64)
	if err != nil {
		return 0, 0, errors.Wrap(err, "parsing uid")
	}
	gid, err := strconv.ParseUint(u.Gid, 10, 64)
	if err != nil {
		return 0, 0, errors.Wrap(err, "parsing gid")
	}
	return uint32(uid), uint32(gid), nil
}

// //
// // Routines to convert DB structs into API structs
// //

func toCluster(c data.Cluster) *web.Cluster {
	return &web.Cluster{
		c.Id,
		c.Name,
		c.ContextPath.String,
		c.ClusterTypeId,
		c.DetailId.Int64,
		c.Address.String,
		c.Token.String,
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
		c.Username.String,
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
	ar := make([]*web.Project, len(ps))
	for i, p := range ps {
		ar[i] = toProject(p)
	}
	return ar
}

func nullToInt64(nullable sql.NullInt64) int64 {
	if nullable.Valid {
		return nullable.Int64
	}
	return -1
}

func fromFrameKey(key *bindings.FrameKeyV3) string {
	if key != nil {
		return key.Name
	}
	return ""
}

func toModel(m data.LabelModel) *web.Model {
	return &web.Model{
		m.Id,
		0,
		0,
		m.Name,
		m.ClusterName,
		m.ModelKey,
		m.Algorithm,
		m.ModelCategory,
		m.DatasetName.String,
		m.ResponseColumn,
		m.LogicalName.String,
		m.Location.String,
		m.ModelObjectType.String,
		int(m.MaxRunTime.Int64),
		m.Schema.String,
		toTimestamp(m.Created),
		nullToInt64(m.Label.Id),
		m.Label.Name.String,
	}
}

func toModels(ms []data.LabelModel) []*web.Model {
	ar := make([]*web.Model, len(ms))
	for i, m := range ms {
		ar[i] = toModel(m)
	}
	return ar
}

func h2oToModel(m *bindings.ModelSchema) *web.Model {
	return &web.Model{
		Name:               m.ModelId.Name,
		ModelKey:           m.ModelId.Name,
		Algorithm:          m.AlgoFullName,
		ModelCategory:      string(m.Output.ModelCategory),
		DatasetName:        fromFrameKey(m.DataFrame),
		ResponseColumnName: m.ResponseColumnName,
		CreatedAt:          toTimestamp(time.Now()),
	}
}

func h2oToModels(ms []*bindings.ModelSchema) []*web.Model {
	ar := make([]*web.Model, len(ms))
	for i, m := range ms {
		ar[i] = h2oToModel(m)
	}
	return ar
}

func toBinomialModel(m data.BinomialModel) *web.BinomialModel {
	return &web.BinomialModel{
		m.Id,
		0,
		0,
		m.Name,
		m.ClusterName,
		m.ModelKey,
		m.Algorithm,
		m.ModelCategory,
		m.DatasetName.String,
		m.ResponseColumn,
		m.LogicalName.String,
		m.Location.String,
		m.ModelObjectType.String,
		int(m.MaxRunTime.Int64),
		m.Schema.String,
		toTimestamp(m.Created),
		nullToInt64(m.Label.Id),
		m.Label.Name.String,
		m.Binomial.Mse,
		m.Binomial.RSquared,
		m.Binomial.Logloss,
		m.Binomial.Auc,
		m.Binomial.Gini,
	}
}

func toBinomialModels(ms []data.BinomialModel) []*web.BinomialModel {
	ar := make([]*web.BinomialModel, len(ms))
	for i, m := range ms {
		ar[i] = toBinomialModel(m)
	}
	return ar
}

func toMultinomialModel(m data.MultinomialModel) *web.MultinomialModel {
	return &web.MultinomialModel{
		m.Id,
		0,
		0,
		m.Name,
		m.ClusterName,
		m.ModelKey,
		m.Algorithm,
		m.ModelCategory,
		m.DatasetName.String,
		m.ResponseColumn,
		m.LogicalName.String,
		m.Location.String,
		m.ModelObjectType.String,
		int(m.MaxRunTime.Int64),
		m.Schema.String,
		toTimestamp(m.Created),
		nullToInt64(m.Label.Id),
		m.Label.Name.String,
		m.Multinomial.Mse,
		m.Multinomial.RSquared,
		m.Multinomial.Logloss,
	}
}

func toMultinomialModels(ms []data.MultinomialModel) []*web.MultinomialModel {
	ar := make([]*web.MultinomialModel, len(ms))
	for i, m := range ms {
		ar[i] = toMultinomialModel(m)
	}
	return ar
}

func toRegressionModel(m data.RegressionModel) *web.RegressionModel {
	return &web.RegressionModel{
		m.Id,
		0,
		0,
		m.Name,
		m.ClusterName,
		m.ModelKey,
		m.Algorithm,
		m.ModelCategory,
		m.DatasetName.String,
		m.ResponseColumn,
		m.LogicalName.String,
		m.Location.String,
		m.ModelObjectType.String,
		int(m.MaxRunTime.Int64),
		m.Schema.String,
		toTimestamp(m.Created),
		nullToInt64(m.Label.Id),
		m.Label.Name.String,
		m.Regression.Mse,
		m.Regression.RSquared,
		m.Regression.MeanResidualDeviance,
	}
}

func toRegressionModels(ms []data.RegressionModel) []*web.RegressionModel {
	ar := make([]*web.RegressionModel, len(ms))
	for i, m := range ms {
		ar[i] = toRegressionModel(m)
	}
	return ar
}

func toLabel(l data.Label) *web.Label {
	return &web.Label{
		l.Id.Int64,
		l.ProjectId.Int64,
		nullToInt64(l.ModelId),
		l.Name.String,
		l.Description.String,
		toTimestamp(l.Created.Time),
	}
}

func toLabels(ls []data.Label) []*web.Label {
	ar := make([]*web.Label, len(ls))
	for i, l := range ls {
		ar[i] = toLabel(l)
	}
	return ar
}

func toService(s data.Service) *web.ScoringService {
	return &web.ScoringService{
		s.Id,
		s.ModelId,
		s.Name,
		s.Host.String,
		int(s.Port.Int64),
		int(s.ProcessId.Int64),
		s.State,
		toTimestamp(s.Created),
	}
}

func toServices(ss []data.Service) []*web.ScoringService {
	ar := make([]*web.ScoringService, len(ss))
	for i, s := range ss {
		ar[i] = toService(s)
	}
	return ar
}

func toEngine(e data.Engine) *web.Engine {
	return &web.Engine{
		e.Id,
		e.Name,
		e.Location,
		toTimestamp(e.Created),
	}
}

func toEngines(es []data.Engine) []*web.Engine {
	ar := make([]*web.Engine, len(es))
	for i, e := range es {
		ar[i] = toEngine(e)
	}
	return ar
}

func toPermission(p data.Permission) *web.Permission {
	return &web.Permission{
		p.Id,
		p.Code,
		p.Description,
	}
}

func toPermissions(ps []data.Permission) []*web.Permission {
	ar := make([]*web.Permission, len(ps))
	for i, p := range ps {
		ar[i] = toPermission(p)
	}
	return ar
}

func toRole(r data.Role) *web.Role {
	return &web.Role{
		r.Id,
		r.Name,
		r.Description.String,
		toTimestamp(r.Created),
	}
}

func toRoles(rs []data.Role) []*web.Role {
	ar := make([]*web.Role, len(rs))
	for i, r := range rs {
		ar[i] = toRole(r)
	}
	return ar
}

func toWorkgroup(w data.Workgroup) *web.Workgroup {
	return &web.Workgroup{
		w.Id,
		w.Name,
		w.Description.String,
		toTimestamp(w.Created),
	}
}

func toWorkgroups(ws []data.Workgroup) []*web.Workgroup {
	ar := make([]*web.Workgroup, len(ws))
	for i, w := range ws {
		ar[i] = toWorkgroup(w)
	}
	return ar
}

func toIdentity(i data.Identity) *web.Identity {
	return &web.Identity{
		i.Id,
		i.Name,
		i.IsActive,
		toTimestamp(i.LastLogin.Time),
		toTimestamp(i.Created),
	}
}

func toIdentities(is []data.Identity) []*web.Identity {
	ar := make([]*web.Identity, len(is))
	for i, iy := range is {
		ar[i] = toIdentity(iy)
	}
	return ar
}

func toEntityPrivilege(e data.EntityPrivilege) *web.EntityPrivilege {
	return &web.EntityPrivilege{
		e.Type,
		e.WorkgroupId,
		e.Workgroup.Name,
		e.Workgroup.Description.String,
	}
}

func toEntityPrivileges(es []data.EntityPrivilege) []*web.EntityPrivilege {
	ar := make([]*web.EntityPrivilege, len(es))
	for i, e := range es {
		ar[i] = toEntityPrivilege(e)
	}
	return ar
}

func toEntityHistory(h data.History) *web.EntityHistory {
	return &web.EntityHistory{
		h.IdentityId,
		h.Action,
		h.Description.String,
		toTimestamp(h.Created),
	}
}

func toEntityHistories(es []data.History) []*web.EntityHistory {
	ar := make([]*web.EntityHistory, len(es))
	for i, e := range es {
		ar[i] = toEntityHistory(e)
	}
	return ar
}

func toLdapGroup(gs map[string]int) []*web.LdapGroup {
	ar := make([]*web.LdapGroup, 0, len(gs))
	for g, u := range gs {
		ar = append(ar, &web.LdapGroup{Name: g, Users: u})
	}
	return ar
}

func toKeytab(k data.Keytab) *web.Keytab {
	return &web.Keytab{
		Id:        int(k.Id),
		Principal: k.Principal.String,
		Name:      k.Filename,
	}
}
