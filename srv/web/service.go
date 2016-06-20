// ----------------------------------
// --- Generated with go:generate ---
// ---        DO NOT EDIT         ---
// ----------------------------------

package web

import(
	"github.com/h2oai/steamY/master/az"
	"net/http"
)

// --- Types ---

type Cluster struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	TypeId int64 `json:"type_id"`
	DetailId int64 `json:"detail_id"`
	Address string `json:"address"`
	State string `json:"state"`
	CreatedAt int64 `json:"created_at"`
}

type YarnCluster struct {
	Id int64 `json:"id"`
	EngineId int64 `json:"engine_id"`
	Size int `json:"size"`
	ApplicationId string `json:"application_id"`
	Memory string `json:"memory"`
	Username string `json:"username"`
}

type ClusterStatus struct {
	Version string `json:"version"`
	Status string `json:"status"`
	MaxMemory string `json:"max_memory"`
	TotalCpuCount int `json:"total_cpu_count"`
	TotalAllowedCpuCount int `json:"total_allowed_cpu_count"`
}

type Job struct {
	Name string `json:"name"`
	ClusterName string `json:"cluster_name"`
	Description string `json:"description"`
	Progress string `json:"progress"`
	StartedAt int64 `json:"started_at"`
	CompletedAt int64 `json:"completed_at"`
}

type Model struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	ClusterName string `json:"cluster_name"`
	Algorithm string `json:"algorithm"`
	DatasetName string `json:"dataset_name"`
	ResponseColumnName string `json:"response_column_name"`
	LogicalName string `json:"logical_name"`
	Location string `json:"location"`
	MaxRuntime int `json:"max_runtime"`
	CreatedAt int64 `json:"created_at"`
}

type ScoringService struct {
	Id int64 `json:"id"`
	ModelId int64 `json:"model_id"`
	Address string `json:"address"`
	Port int `json:"port"`
	ProcessId int `json:"process_id"`
	State string `json:"state"`
	CreatedAt int64 `json:"created_at"`
}

type Engine struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
	CreatedAt int64 `json:"created_at"`
}

// --- Interfaces ---


		type Az interface {
			Identify(r *http.Request) (az.Principal, error)
		}

		type Service interface {
	Ping(pz az.Principal,status bool) (bool, error)
	RegisterCluster(pz az.Principal,address string) (int64, error)
	UnregisterCluster(pz az.Principal,clusterId int64) error
	StartYarnCluster(pz az.Principal,clusterName string, engineId int64, size int, memory string, username string) (int64, error)
	StopYarnCluster(pz az.Principal,clusterId int64) error
	GetCluster(pz az.Principal,clusterId int64) (*Cluster, error)
	GetYarnCluster(pz az.Principal,clusterId int64) (*YarnCluster, error)
	GetClusters(pz az.Principal,offset int64, limit int64) ([]*Cluster, error)
	GetClusterStatus(pz az.Principal,clusterId int64) (*ClusterStatus, error)
	DeleteCluster(pz az.Principal,clusterId int64) error
	GetJob(pz az.Principal,clusterId int64, jobName string) (*Job, error)
	GetJobs(pz az.Principal,clusterId int64) ([]*Job, error)
	BuildModel(pz az.Principal,clusterId int64, dataset string, targetName string, maxRunTime int) (*Model, error)
	GetModel(pz az.Principal,modelId int64) (*Model, error)
	GetModels(pz az.Principal,offset int64, limit int64) ([]*Model, error)
	GetClusterModels(pz az.Principal,clusterId int64) ([]*Model, error)
	GetModelFromCluster(pz az.Principal,clusterId int64, modelName string) (*Model, error)
	DeleteModel(pz az.Principal,modelId int64) error
	StartScoringService(pz az.Principal,modelId int64, port int) (*ScoringService, error)
	StopScoringService(pz az.Principal,serviceId int64) error
	GetScoringService(pz az.Principal,serviceId int64) (*ScoringService, error)
	GetScoringServices(pz az.Principal,offset int64, limit int64) ([]*ScoringService, error)
	DeleteScoringService(pz az.Principal,serviceId int64) error
	AddEngine(pz az.Principal,engineName string, enginePath string) (int64, error)
	GetEngine(pz az.Principal,engineId int64) (*Engine, error)
	GetEngines(pz az.Principal,) ([]*Engine, error)
	DeleteEngine(pz az.Principal,engineId int64) error
}

// --- Messages ---

type PingIn struct {
	Status bool `json:"status"`
}

type PingOut struct {
	Status bool `json:"status"`
}

type RegisterClusterIn struct {
	Address string `json:"address"`
}

type RegisterClusterOut struct {
	ClusterId int64 `json:"cluster_id"`
}

type UnregisterClusterIn struct {
	ClusterId int64 `json:"cluster_id"`
}

type UnregisterClusterOut struct {
}

type StartYarnClusterIn struct {
	ClusterName string `json:"cluster_name"`
	EngineId int64 `json:"engine_id"`
	Size int `json:"size"`
	Memory string `json:"memory"`
	Username string `json:"username"`
}

type StartYarnClusterOut struct {
	ClusterId int64 `json:"cluster_id"`
}

type StopYarnClusterIn struct {
	ClusterId int64 `json:"cluster_id"`
}

type StopYarnClusterOut struct {
}

type GetClusterIn struct {
	ClusterId int64 `json:"cluster_id"`
}

type GetClusterOut struct {
	Cluster *Cluster `json:"cluster"`
}

type GetYarnClusterIn struct {
	ClusterId int64 `json:"cluster_id"`
}

type GetYarnClusterOut struct {
	Cluster *YarnCluster `json:"cluster"`
}

type GetClustersIn struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
}

type GetClustersOut struct {
	Clusters []*Cluster `json:"clusters"`
}

type GetClusterStatusIn struct {
	ClusterId int64 `json:"cluster_id"`
}

type GetClusterStatusOut struct {
	ClusterStatus *ClusterStatus `json:"cluster_status"`
}

type DeleteClusterIn struct {
	ClusterId int64 `json:"cluster_id"`
}

type DeleteClusterOut struct {
}

type GetJobIn struct {
	ClusterId int64 `json:"cluster_id"`
	JobName string `json:"job_name"`
}

type GetJobOut struct {
	Job *Job `json:"job"`
}

type GetJobsIn struct {
	ClusterId int64 `json:"cluster_id"`
}

type GetJobsOut struct {
	Jobs []*Job `json:"jobs"`
}

type BuildModelIn struct {
	ClusterId int64 `json:"cluster_id"`
	Dataset string `json:"dataset"`
	TargetName string `json:"target_name"`
	MaxRunTime int `json:"max_run_time"`
}

type BuildModelOut struct {
	Model *Model `json:"model"`
}

type GetModelIn struct {
	ModelId int64 `json:"model_id"`
}

type GetModelOut struct {
	Model *Model `json:"model"`
}

type GetModelsIn struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
}

type GetModelsOut struct {
	Models []*Model `json:"models"`
}

type GetClusterModelsIn struct {
	ClusterId int64 `json:"cluster_id"`
}

type GetClusterModelsOut struct {
	Models []*Model `json:"models"`
}

type GetModelFromClusterIn struct {
	ClusterId int64 `json:"cluster_id"`
	ModelName string `json:"model_name"`
}

type GetModelFromClusterOut struct {
	Model *Model `json:"model"`
}

type DeleteModelIn struct {
	ModelId int64 `json:"model_id"`
}

type DeleteModelOut struct {
}

type StartScoringServiceIn struct {
	ModelId int64 `json:"model_id"`
	Port int `json:"port"`
}

type StartScoringServiceOut struct {
	Service *ScoringService `json:"service"`
}

type StopScoringServiceIn struct {
	ServiceId int64 `json:"service_id"`
}

type StopScoringServiceOut struct {
}

type GetScoringServiceIn struct {
	ServiceId int64 `json:"service_id"`
}

type GetScoringServiceOut struct {
	Service *ScoringService `json:"service"`
}

type GetScoringServicesIn struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
}

type GetScoringServicesOut struct {
	Services []*ScoringService `json:"services"`
}

type DeleteScoringServiceIn struct {
	ServiceId int64 `json:"service_id"`
}

type DeleteScoringServiceOut struct {
}

type AddEngineIn struct {
	EngineName string `json:"engine_name"`
	EnginePath string `json:"engine_path"`
}

type AddEngineOut struct {
	EngineId int64 `json:"engine_id"`
}

type GetEngineIn struct {
	EngineId int64 `json:"engine_id"`
}

type GetEngineOut struct {
	Engine *Engine `json:"engine"`
}

type GetEnginesIn struct {
}

type GetEnginesOut struct {
	Engines []*Engine `json:"engines"`
}

type DeleteEngineIn struct {
	EngineId int64 `json:"engine_id"`
}

type DeleteEngineOut struct {
}

// --- Client Stub ---

type Remote struct {
	Proc Proc
}

type Proc interface {
	Call(name string, in, out interface{}) error
}

func (this *Remote) Ping(status bool) (bool,  error) {
	in := PingIn{status}
	var out PingOut
	err := this.Proc.Call("Ping", &in, &out)
	if err != nil {
		return false, err
	}
	return out.Status, nil
}

func (this *Remote) RegisterCluster(address string) (int64,  error) {
	in := RegisterClusterIn{address}
	var out RegisterClusterOut
	err := this.Proc.Call("RegisterCluster", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.ClusterId, nil
}

func (this *Remote) UnregisterCluster(clusterId int64) ( error) {
	in := UnregisterClusterIn{clusterId}
	var out UnregisterClusterOut
	err := this.Proc.Call("UnregisterCluster", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) StartYarnCluster(clusterName string, engineId int64, size int, memory string, username string) (int64,  error) {
	in := StartYarnClusterIn{clusterName, engineId, size, memory, username}
	var out StartYarnClusterOut
	err := this.Proc.Call("StartYarnCluster", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.ClusterId, nil
}

func (this *Remote) StopYarnCluster(clusterId int64) ( error) {
	in := StopYarnClusterIn{clusterId}
	var out StopYarnClusterOut
	err := this.Proc.Call("StopYarnCluster", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetCluster(clusterId int64) (*Cluster,  error) {
	in := GetClusterIn{clusterId}
	var out GetClusterOut
	err := this.Proc.Call("GetCluster", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Cluster, nil
}

func (this *Remote) GetYarnCluster(clusterId int64) (*YarnCluster,  error) {
	in := GetYarnClusterIn{clusterId}
	var out GetYarnClusterOut
	err := this.Proc.Call("GetYarnCluster", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Cluster, nil
}

func (this *Remote) GetClusters(offset int64, limit int64) ([]*Cluster,  error) {
	in := GetClustersIn{offset, limit}
	var out GetClustersOut
	err := this.Proc.Call("GetClusters", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Clusters, nil
}

func (this *Remote) GetClusterStatus(clusterId int64) (*ClusterStatus,  error) {
	in := GetClusterStatusIn{clusterId}
	var out GetClusterStatusOut
	err := this.Proc.Call("GetClusterStatus", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.ClusterStatus, nil
}

func (this *Remote) DeleteCluster(clusterId int64) ( error) {
	in := DeleteClusterIn{clusterId}
	var out DeleteClusterOut
	err := this.Proc.Call("DeleteCluster", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetJob(clusterId int64, jobName string) (*Job,  error) {
	in := GetJobIn{clusterId, jobName}
	var out GetJobOut
	err := this.Proc.Call("GetJob", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Job, nil
}

func (this *Remote) GetJobs(clusterId int64) ([]*Job,  error) {
	in := GetJobsIn{clusterId}
	var out GetJobsOut
	err := this.Proc.Call("GetJobs", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Jobs, nil
}

func (this *Remote) BuildModel(clusterId int64, dataset string, targetName string, maxRunTime int) (*Model,  error) {
	in := BuildModelIn{clusterId, dataset, targetName, maxRunTime}
	var out BuildModelOut
	err := this.Proc.Call("BuildModel", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Model, nil
}

func (this *Remote) GetModel(modelId int64) (*Model,  error) {
	in := GetModelIn{modelId}
	var out GetModelOut
	err := this.Proc.Call("GetModel", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Model, nil
}

func (this *Remote) GetModels(offset int64, limit int64) ([]*Model,  error) {
	in := GetModelsIn{offset, limit}
	var out GetModelsOut
	err := this.Proc.Call("GetModels", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Models, nil
}

func (this *Remote) GetClusterModels(clusterId int64) ([]*Model,  error) {
	in := GetClusterModelsIn{clusterId}
	var out GetClusterModelsOut
	err := this.Proc.Call("GetClusterModels", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Models, nil
}

func (this *Remote) GetModelFromCluster(clusterId int64, modelName string) (*Model,  error) {
	in := GetModelFromClusterIn{clusterId, modelName}
	var out GetModelFromClusterOut
	err := this.Proc.Call("GetModelFromCluster", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Model, nil
}

func (this *Remote) DeleteModel(modelId int64) ( error) {
	in := DeleteModelIn{modelId}
	var out DeleteModelOut
	err := this.Proc.Call("DeleteModel", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) StartScoringService(modelId int64, port int) (*ScoringService,  error) {
	in := StartScoringServiceIn{modelId, port}
	var out StartScoringServiceOut
	err := this.Proc.Call("StartScoringService", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Service, nil
}

func (this *Remote) StopScoringService(serviceId int64) ( error) {
	in := StopScoringServiceIn{serviceId}
	var out StopScoringServiceOut
	err := this.Proc.Call("StopScoringService", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetScoringService(serviceId int64) (*ScoringService,  error) {
	in := GetScoringServiceIn{serviceId}
	var out GetScoringServiceOut
	err := this.Proc.Call("GetScoringService", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Service, nil
}

func (this *Remote) GetScoringServices(offset int64, limit int64) ([]*ScoringService,  error) {
	in := GetScoringServicesIn{offset, limit}
	var out GetScoringServicesOut
	err := this.Proc.Call("GetScoringServices", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Services, nil
}

func (this *Remote) DeleteScoringService(serviceId int64) ( error) {
	in := DeleteScoringServiceIn{serviceId}
	var out DeleteScoringServiceOut
	err := this.Proc.Call("DeleteScoringService", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) AddEngine(engineName string, enginePath string) (int64,  error) {
	in := AddEngineIn{engineName, enginePath}
	var out AddEngineOut
	err := this.Proc.Call("AddEngine", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.EngineId, nil
}

func (this *Remote) GetEngine(engineId int64) (*Engine,  error) {
	in := GetEngineIn{engineId}
	var out GetEngineOut
	err := this.Proc.Call("GetEngine", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Engine, nil
}

func (this *Remote) GetEngines() ([]*Engine,  error) {
	in := GetEnginesIn{}
	var out GetEnginesOut
	err := this.Proc.Call("GetEngines", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Engines, nil
}

func (this *Remote) DeleteEngine(engineId int64) ( error) {
	in := DeleteEngineIn{engineId}
	var out DeleteEngineOut
	err := this.Proc.Call("DeleteEngine", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

// --- Server Stub ---

type Impl struct {
	Service Service
	Az az.Az
}

func (this *Impl) Ping(r *http.Request, in *PingIn, out *PingOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.Ping(pz,in.Status)
	if err != nil {
		return err
	}
	out.Status = it
	return nil
}

func (this *Impl) RegisterCluster(r *http.Request, in *RegisterClusterIn, out *RegisterClusterOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.RegisterCluster(pz,in.Address)
	if err != nil {
		return err
	}
	out.ClusterId = it
	return nil
}

func (this *Impl) UnregisterCluster(r *http.Request, in *UnregisterClusterIn, out *UnregisterClusterOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.UnregisterCluster(pz,in.ClusterId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) StartYarnCluster(r *http.Request, in *StartYarnClusterIn, out *StartYarnClusterOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.StartYarnCluster(pz,in.ClusterName, in.EngineId, in.Size, in.Memory, in.Username)
	if err != nil {
		return err
	}
	out.ClusterId = it
	return nil
}

func (this *Impl) StopYarnCluster(r *http.Request, in *StopYarnClusterIn, out *StopYarnClusterOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.StopYarnCluster(pz,in.ClusterId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) GetCluster(r *http.Request, in *GetClusterIn, out *GetClusterOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetCluster(pz,in.ClusterId)
	if err != nil {
		return err
	}
	out.Cluster = it
	return nil
}

func (this *Impl) GetYarnCluster(r *http.Request, in *GetYarnClusterIn, out *GetYarnClusterOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetYarnCluster(pz,in.ClusterId)
	if err != nil {
		return err
	}
	out.Cluster = it
	return nil
}

func (this *Impl) GetClusters(r *http.Request, in *GetClustersIn, out *GetClustersOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetClusters(pz,in.Offset, in.Limit)
	if err != nil {
		return err
	}
	out.Clusters = it
	return nil
}

func (this *Impl) GetClusterStatus(r *http.Request, in *GetClusterStatusIn, out *GetClusterStatusOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetClusterStatus(pz,in.ClusterId)
	if err != nil {
		return err
	}
	out.ClusterStatus = it
	return nil
}

func (this *Impl) DeleteCluster(r *http.Request, in *DeleteClusterIn, out *DeleteClusterOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.DeleteCluster(pz,in.ClusterId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) GetJob(r *http.Request, in *GetJobIn, out *GetJobOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetJob(pz,in.ClusterId, in.JobName)
	if err != nil {
		return err
	}
	out.Job = it
	return nil
}

func (this *Impl) GetJobs(r *http.Request, in *GetJobsIn, out *GetJobsOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetJobs(pz,in.ClusterId)
	if err != nil {
		return err
	}
	out.Jobs = it
	return nil
}

func (this *Impl) BuildModel(r *http.Request, in *BuildModelIn, out *BuildModelOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.BuildModel(pz,in.ClusterId, in.Dataset, in.TargetName, in.MaxRunTime)
	if err != nil {
		return err
	}
	out.Model = it
	return nil
}

func (this *Impl) GetModel(r *http.Request, in *GetModelIn, out *GetModelOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetModel(pz,in.ModelId)
	if err != nil {
		return err
	}
	out.Model = it
	return nil
}

func (this *Impl) GetModels(r *http.Request, in *GetModelsIn, out *GetModelsOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetModels(pz,in.Offset, in.Limit)
	if err != nil {
		return err
	}
	out.Models = it
	return nil
}

func (this *Impl) GetClusterModels(r *http.Request, in *GetClusterModelsIn, out *GetClusterModelsOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetClusterModels(pz,in.ClusterId)
	if err != nil {
		return err
	}
	out.Models = it
	return nil
}

func (this *Impl) GetModelFromCluster(r *http.Request, in *GetModelFromClusterIn, out *GetModelFromClusterOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetModelFromCluster(pz,in.ClusterId, in.ModelName)
	if err != nil {
		return err
	}
	out.Model = it
	return nil
}

func (this *Impl) DeleteModel(r *http.Request, in *DeleteModelIn, out *DeleteModelOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.DeleteModel(pz,in.ModelId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) StartScoringService(r *http.Request, in *StartScoringServiceIn, out *StartScoringServiceOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.StartScoringService(pz,in.ModelId, in.Port)
	if err != nil {
		return err
	}
	out.Service = it
	return nil
}

func (this *Impl) StopScoringService(r *http.Request, in *StopScoringServiceIn, out *StopScoringServiceOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.StopScoringService(pz,in.ServiceId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) GetScoringService(r *http.Request, in *GetScoringServiceIn, out *GetScoringServiceOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetScoringService(pz,in.ServiceId)
	if err != nil {
		return err
	}
	out.Service = it
	return nil
}

func (this *Impl) GetScoringServices(r *http.Request, in *GetScoringServicesIn, out *GetScoringServicesOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetScoringServices(pz,in.Offset, in.Limit)
	if err != nil {
		return err
	}
	out.Services = it
	return nil
}

func (this *Impl) DeleteScoringService(r *http.Request, in *DeleteScoringServiceIn, out *DeleteScoringServiceOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.DeleteScoringService(pz,in.ServiceId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) AddEngine(r *http.Request, in *AddEngineIn, out *AddEngineOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.AddEngine(pz,in.EngineName, in.EnginePath)
	if err != nil {
		return err
	}
	out.EngineId = it
	return nil
}

func (this *Impl) GetEngine(r *http.Request, in *GetEngineIn, out *GetEngineOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetEngine(pz,in.EngineId)
	if err != nil {
		return err
	}
	out.Engine = it
	return nil
}

func (this *Impl) GetEngines(r *http.Request, in *GetEnginesIn, out *GetEnginesOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetEngines(pz,)
	if err != nil {
		return err
	}
	out.Engines = it
	return nil
}

func (this *Impl) DeleteEngine(r *http.Request, in *DeleteEngineIn, out *DeleteEngineOut) error {
	
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.DeleteEngine(pz,in.EngineId)
	if err != nil {
		return err
	}
	return nil
}

