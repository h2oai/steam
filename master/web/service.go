package web

import (
	"fmt"
	"github.com/h2oai/steamY/bindings"
	"github.com/h2oai/steamY/lib/fs"
	"github.com/h2oai/steamY/lib/svc"
	"github.com/h2oai/steamY/lib/yarn"
	"github.com/h2oai/steamY/master/auth"
	"github.com/h2oai/steamY/master/az"
	"github.com/h2oai/steamY/master/data"
	"github.com/h2oai/steamY/srv/compiler" // FIXME rename comp to compiler
	"github.com/h2oai/steamY/srv/h2ov3"
	"github.com/h2oai/steamY/srv/web"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type Service struct {
	workingDir                string
	ds                        *data.Datastore
	compilationServiceAddress string
	scoringServiceAddress     string
	kerberosEnabled           bool
	username                  string
	keytab                    string
}

func NewService(workingDir string, ds *data.Datastore, compilationServiceAddress, scoringServiceAddress string, kerberos bool, username, keytab string) *Service {
	return &Service{
		workingDir,
		ds,
		compilationServiceAddress,
		scoringServiceAddress,
		kerberos,
		username,
		keytab,
	}
}

func toTimestamp(t time.Time) int64 {
	return t.UTC().Unix()
}

func now() int64 {
	return toTimestamp(time.Now())
}

func (s *Service) Ping(pz az.Principal, status bool) (bool, error) {
	return status, nil
}

func (s *Service) RegisterCluster(pz az.Principal, address string) (int64, error) {

	if err := pz.CheckPermission(s.ds.Permissions.ManageCluster); err != nil {
		return 0, err
	}

	h := h2ov3.NewClient(address)
	cloud, err := h.GetCloud()
	if err != nil {
		return 0, fmt.Errorf("Could not communicate with cloud %s.", address)
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
		return 0, fmt.Errorf("Failed storing cluster:", err)
	}
	// s.clusterProxy.NewProxy(c.ID, c.Address)
	// if err := s.pollCloud(toCloud(c)); err != nil {
	// 	return nil, err
	// }

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

func (s *Service) StartYarnCluster(pz az.Principal, clusterName string, engineId int64, size int, memory, username string) (int64, error) {

	if err := pz.CheckPermission(s.ds.Permissions.ManageCluster); err != nil {
		return 0, err
	}

	// Cluster should have a unique name
	_, ok, err := s.ds.ReadClusterByName(pz, clusterName)
	if err != nil {
		return 0, err
	}
	if ok {
		return 0, fmt.Errorf("Failed starting cluster: a cluster with the name %s already exists.", clusterName)
	}

	engine, err := s.ds.ReadEngine(pz, engineId)
	if err != nil {
		return 0, fmt.Errorf("Failed starting cluster: cannot locate the specified engine %d: %s", engineId, err)
	}

	applicationId, address, out, err := yarn.StartCloud(size, s.kerberosEnabled, memory, clusterName, engine.Location, s.username, s.keytab) // FIXME: THIS IS USING ADMIN TO START ALL CLOUDS
	if err != nil {
		log.Println(pz, err)
		return 0, err
	}

	yarnCluster := data.YarnCluster{
		0,
		engineId,
		int64(size),
		applicationId,
		memory,
		username,
		out,
	}

	clusterId, err := s.ds.CreateYarnCluster(pz, clusterName, address, data.StartedState, yarnCluster)

	if err != nil {
		return 0, err
	}

	// // Create an instance of this cloud in cloudActivity map
	// if err := s.pollCloud(toCluster(c)); err != nil {
	// 	return nil, err
	// }

	// Create a reverse proxy for cluster.
	// FIXME
	// s.clusterProxy.NewProxy(c.ID, c.Address)

	return clusterId, nil
}

func (s *Service) StopYarnCluster(pz az.Principal, clusterId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageCluster); err != nil {
		return err
	}

	// Cluster should exist
	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return err
	}

	if cluster.TypeId != s.ds.ClusterTypes.Yarn {
		return fmt.Errorf("Cluster %d was not started through YARN", clusterId)
	}

	// Bail out if already stopped
	if cluster.State == data.StoppedState {
		return fmt.Errorf("Cluster %d is already stopped", clusterId)
	}

	yarnCluster, err := s.ds.ReadYarnCluster(pz, clusterId)
	if err != nil {
		return err
	}

	if err := yarn.StopCloud(s.kerberosEnabled, cluster.Name, yarnCluster.ApplicationId, yarnCluster.OutputDir, s.username, s.keytab); err != nil { //FIXME: this is using adming kerberos credentials
		log.Println(err)
		return err
	}

	return s.ds.UpdateClusterState(pz, clusterId, data.StoppedState)
}

// FIXME - why is this required?

// func (s *Service) Shutdown(address string) error {
// 	h := h2ov3.NewClient(address)

// 	if err := h.Shutdown(); err != nil {
// 		return err
// 	}

// 	return nil
// }

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

func (s *Service) GetYarnCluster(pz az.Principal, clusterId int64) (*web.YarnCluster, error) {
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

	cloud, err := h2o.GetCloud()

	var (
		tot, all int32
		mem      int64
	)
	for _, n := range cloud.Nodes {
		mem += n.MaxMem
		tot += n.NumCpus
		all += n.CpusAllowed
	}

	// FIXME: this needs a better impl
	var health string
	if cloud.CloudHealthy {
		health = "healthy"
	} else {
		health = "unknown"
	}

	return &web.ClusterStatus{
		cloud.Version,
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

	j, err := h.GetJob(jobName)
	if err != nil {
		return nil, err //FIXME format error
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

	j, err := h.GetJobs()
	if err != nil {
		return nil, err //FIXME format error
	}

	jobs := make([]*web.Job, len(j.Jobs))
	for i, job := range j.Jobs {
		jobs[i] = toJob(job)
	}

	sort.Sort(sort.Reverse(Jobs(jobs)))

	return jobs, nil
}

func (s *Service) exportModel(h2o *h2ov3.H2O, modelName string) (string, string, error) {

	// FIXME: allow overwriting of existing java-model/genmodel/metrics, if any.
	// FIXME: purge war file if overwriting, so that a fresh war file can be built the next time around.

	var location, logicalName string
	location = fs.GetModelPath(s.workingDir, modelName)
	javaModelPath, err := h2o.ExportJavaModel(modelName, location)
	if err != nil {
		return location, logicalName, err
	}
	logicalName = fs.GetBasenameWithoutExt(javaModelPath)

	if _, err := h2o.ExportGenModel(location); err != nil {
		return location, logicalName, err
	}

	return location, logicalName, err
}

func (s *Service) BuildModel(pz az.Principal, clusterId int64, dataset, targetName string, maxRunTime int) (*web.Model, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ManageModel); err != nil {
		return nil, err
	}
	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return nil, err
	}
	if cluster.State == data.StoppedState {
		return nil, fmt.Errorf("Failed building model: cluster is not running")
	}

	h2o := h2ov3.NewClient(cluster.Address)

	modelName, err := h2o.AutoML(dataset, targetName, maxRunTime) // TODO: can be a goroutine
	if err != nil {
		return nil, err
	}

	location, logicalName, err := s.exportModel(h2o, modelName)
	if err != nil {
		return nil, err
	}

	modelId, err := s.ds.CreateModel(pz, data.Model{
		0,
		modelName,
		cluster.Name,
		"AutoML",
		dataset,
		targetName,
		logicalName,
		location,
		int64(maxRunTime),
		time.Now(),
	})
	if err != nil {
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

func (s *Service) GetModels(pz az.Principal, offset, limit int64) ([]*web.Model, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewModel); err != nil {
		return nil, err
	}
	ms, err := s.ds.ReadModels(pz, offset, limit)
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
func dataFrameName(m *bindings.ModelSchemaBase) string {
	if m.DataFrame != nil {
		return m.DataFrame.Name
	}

	return ""
}

func (s *Service) GetClusterModels(pz az.Principal, clusterId int64) ([]*web.Model, error) {
	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return nil, err
	}

	h := h2ov3.NewClient(cluster.Address)
	ms, err := h.GetModels()
	if err != nil {
		return nil, fmt.Errorf("Failed fetching models from cluster: %s", err)
	}

	models := make([]*web.Model, len(ms.Models))
	for i, m := range ms.Models {
		models[i] = &web.Model{
			0,
			m.ModelId.Name,
			cluster.Name,
			m.AlgoFullName,
			dataFrameName(m),
			m.ResponseColumnName,
			"",
			"",
			0,
			m.Timestamp,
		}
	}

	return models, nil
}

func (s *Service) ImportModelFromCluster(pz az.Principal, clusterId int64, modelName string) (*web.Model, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewModel); err != nil {
		return nil, err
	}

	cluster, err := s.ds.ReadCluster(pz, clusterId)
	if err != nil {
		return nil, err
	}

	log.Printf("Started: Searching for model %s in cluster %s...", modelName, cluster.Name)
	// get model from the cloud
	h2o := h2ov3.NewClient(cluster.Address)
	r, err := h2o.GetModel(modelName)
	if err != nil {
		return nil, err
	}

	location, logicalName, err := s.exportModel(h2o, modelName)
	if err != nil {
		return nil, err
	}

	m := r.Models[0]

	modelId, err := s.ds.CreateModel(pz, data.Model{
		0,
		modelName,
		cluster.Name,
		m.AlgoFullName,
		dataFrameName(m),
		m.ResponseColumnName,
		logicalName,
		location,
		0,
		time.Now(),
	})
	if err != nil {
		return nil, err
	}

	model, err := s.ds.ReadModel(pz, modelId)
	if err != nil {
		return nil, err
	}

	mod := toModel(model)
	mod.Algorithm = m.AlgoFullName

	return mod, nil
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
				return fmt.Errorf("Failed deleting model: a scoring service for this model is deployed and running at %s:%d", service.Address, service.Port)
			}
		}
	}

	return s.ds.DeleteModel(pz, modelId)
}

func (s *Service) StartScoringService(pz az.Principal, modelId int64, port int) (*web.ScoringService, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ManageService); err != nil {
		return nil, err
	}

	// FIXME: change sequence to:
	// 1. insert a record into the Service table with the state "starting"
	// 2. attempt to compile and start the service
	// 3. update the Service record state to "started" if successful, or "failed" if not.

	model, err := s.ds.ReadModel(pz, modelId)
	if err != nil {
		return nil, err
	}

	compilationService := compiler.NewServer(s.compilationServiceAddress)
	if err := compilationService.Ping(); err != nil {
		return nil, fmt.Errorf("Failed connecting to compilation service at %s", s.compilationServiceAddress)
	}

	// do not recompile if war file is already available
	warFilePath := fs.GetWarFilePath(s.workingDir, model.Name, model.LogicalName)
	if _, err := os.Stat(warFilePath); os.IsNotExist(err) {
		warFilePath, err = compilationService.CompilePojo(
			fs.GetJavaModelPath(s.workingDir, model.Name, model.LogicalName),
			fs.GetGenModelPath(s.workingDir, model.Name),
			"makewar",
		)
		if err != nil {
			return nil, err //FIXME format error
		}
	}

	pid, err := svc.Start(
		warFilePath,
		fs.GetAssetsPath(s.workingDir, "jetty-runner.jar"),
		s.scoringServiceAddress,
		port,
	)
	if err != nil {
		return nil, err
	}

	address, err := fs.GetExternalHost() // FIXME there is no need to re-scan this every time. Can be a property on *Service at init time.
	if err != nil {
		return nil, err
	}

	log.Printf("Scoring service started at %s:%d\n", address, port)

	service := data.Service{
		0,
		model.Id,
		address,
		int64(port), // FIXME change to int
		int64(pid),  // FIXME change to int
		data.StartedState,
		time.Now(),
	}

	serviceId, err := s.ds.CreateService(pz, service)
	if err != nil {
		return nil, err
	}

	service, err = s.ds.ReadService(pz, serviceId)
	if err != nil {
		return nil, err
	}

	// s.scoreActivity.Lock()
	// s.scoreActivity.latest[modelName] = ss.CreatedAt
	// s.scoreActivity.Unlock()

	return toScoringService(service), nil
}

func (s *Service) StopScoringService(pz az.Principal, serviceId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageService); err != nil {
		return err
	}
	service, err := s.ds.ReadService(pz, serviceId)
	if err != nil {
		return err
	}

	if service.State == data.StoppedState {
		return fmt.Errorf("Scoring service is already stopped")
	}

	if err := svc.Stop(int(service.ProcessId)); err != nil {
		return err
	}

	if err := s.ds.UpdateServiceState(pz, serviceId, data.StoppedState); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetScoringService(pz az.Principal, serviceId int64) (*web.ScoringService, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewService); err != nil {
		return nil, err
	}

	service, err := s.ds.ReadService(pz, serviceId)
	if err != nil {
		return nil, err
	}
	return toScoringService(service), nil
}

func (s *Service) GetScoringServices(pz az.Principal, offset, limit int64) ([]*web.ScoringService, error) {
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

func (s *Service) GetScoringServicesForModel(pz az.Principal, modelId, offset, limit int64) ([]*web.ScoringService, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ViewService); err != nil {
		return nil, err //FIXME format error
	}

	services, err := s.ds.ReadServicesForModelId(pz, modelId)
	if err != nil {
		return nil, err //FIXME format error
	}

	ss := make([]*web.ScoringService, len(services))
	for i, service := range services {
		ss[i] = toScoringService(service)
	}

	return ss, nil
}

func (s *Service) DeleteScoringService(pz az.Principal, serviceId int64) error {
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

// FIXME this should not be here - not an client-facing API
func (s *Service) AddEngine(pz az.Principal, engineName, enginePath string) (int64, error) {
	if err := pz.CheckPermission(s.ds.Permissions.ManageEngine); err != nil {
		return 0, err
	}

	return s.ds.CreateEngine(pz, engineName, enginePath)
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

func (s *Service) GetSupportedClusterTypes(pz az.Principal) ([]*web.ClusterType, error) {

	// No permission checks required

	return toClusterTypes(s.ds.ReadClusterTypes(pz)), nil
}

func (s *Service) GetSupportedEntityTypes(pz az.Principal) ([]*web.EntityType, error) {

	// No permission checks required

	return toEntityTypes(s.ds.ReadEntityTypes(pz)), nil
}

func (s *Service) GetSupportedPermissions(pz az.Principal) ([]*web.Permission, error) {

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

func (s *Service) LinkRoleAndPermissions(pz az.Principal, roleId int64, permissionIds []int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageRole); err != nil {
		return err
	}

	return s.ds.LinkRoleAndPermissions(pz, roleId, permissionIds)
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
		return 0, fmt.Errorf("Failed hashing password: %s", err)
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

func (s *Service) LinkIdentityAndWorkgroup(pz az.Principal, identityId int64, workgroupId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageIdentity); err != nil {
		return err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewWorkgroup); err != nil {
		return err
	}

	return s.ds.LinkIdentityAndWorkgroup(pz, identityId, workgroupId)
}

func (s *Service) UnlinkIdentityAndWorkgroup(pz az.Principal, identityId int64, workgroupId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageIdentity); err != nil {
		return err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewWorkgroup); err != nil {
		return err
	}

	return s.ds.UnlinkIdentityAndWorkgroup(pz, identityId, workgroupId)
}

func (s *Service) LinkIdentityAndRole(pz az.Principal, identityId int64, roleId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageIdentity); err != nil {
		return err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewRole); err != nil {
		return err
	}

	return s.ds.LinkIdentityAndRole(pz, identityId, roleId)
}

func (s *Service) UnlinkIdentityAndRole(pz az.Principal, identityId int64, roleId int64) error {
	if err := pz.CheckPermission(s.ds.Permissions.ManageIdentity); err != nil {
		return err
	}
	if err := pz.CheckPermission(s.ds.Permissions.ViewRole); err != nil {
		return err
	}

	return s.ds.UnlinkIdentityAndRole(pz, identityId, roleId)
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

func (s *Service) GetEntityPrivileges(pz az.Principal, entityTypeId, entityId int64) ([]*web.EntityPrivilege, error) {
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

func (s *Service) GetEntityHistory(pz az.Principal, entityTypeId, entityId, offset, limit int64) ([]*web.EntityHistory, error) {
	if err := pz.CheckPermission(s.ds.ViewPermissions[entityTypeId]); err != nil {
		return nil, err
	}

	history, err := s.ds.ReadHistoryForEntity(pz, entityTypeId, entityId, offset, limit)
	if err != nil {
		return nil, err
	}
	return toEntityHistory(history), nil
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

func toModel(m data.Model) *web.Model {
	return &web.Model{
		m.Id,
		m.Name,
		m.ClusterName,
		m.Algorithm,
		m.DatasetName,
		m.ResponseColumnName,
		m.LogicalName,
		m.Location,
		int(m.MaxRunTime), // FIXME change db field to int
		toTimestamp(m.Created),
	}
}

func toScoringService(s data.Service) *web.ScoringService {
	return &web.ScoringService{
		s.Id,
		s.ModelId,
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
