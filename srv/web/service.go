// ----------------------------------
// --- Generated with go:generate ---
// ---        DO NOT EDIT         ---
// ----------------------------------

package web

import (
	"github.com/h2oai/steamY/master/az"
	"net/http"
)

// --- Types ---

type Cluster struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	TypeId    int64  `json:"type_id"`
	DetailId  int64  `json:"detail_id"`
	Address   string `json:"address"`
	State     string `json:"state"`
	CreatedAt int64  `json:"created_at"`
}

type YarnCluster struct {
	Id            int64  `json:"id"`
	EngineId      int64  `json:"engine_id"`
	Size          int    `json:"size"`
	ApplicationId string `json:"application_id"`
	Memory        string `json:"memory"`
	Username      string `json:"username"`
}

type ClusterStatus struct {
	Version              string `json:"version"`
	Status               string `json:"status"`
	MaxMemory            string `json:"max_memory"`
	TotalCpuCount        int    `json:"total_cpu_count"`
	TotalAllowedCpuCount int    `json:"total_allowed_cpu_count"`
}

type Job struct {
	Name        string `json:"name"`
	ClusterName string `json:"cluster_name"`
	Description string `json:"description"`
	Progress    string `json:"progress"`
	StartedAt   int64  `json:"started_at"`
	CompletedAt int64  `json:"completed_at"`
}

type Model struct {
	Id                 int64  `json:"id"`
	Name               string `json:"name"`
	ClusterName        string `json:"cluster_name"`
	Algorithm          string `json:"algorithm"`
	DatasetName        string `json:"dataset_name"`
	ResponseColumnName string `json:"response_column_name"`
	LogicalName        string `json:"logical_name"`
	Location           string `json:"location"`
	MaxRuntime         int    `json:"max_runtime"`
	CreatedAt          int64  `json:"created_at"`
}

type ScoringService struct {
	Id        int64  `json:"id"`
	ModelId   int64  `json:"model_id"`
	Address   string `json:"address"`
	Port      int    `json:"port"`
	ProcessId int    `json:"process_id"`
	State     string `json:"state"`
	CreatedAt int64  `json:"created_at"`
}

type Engine struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	CreatedAt int64  `json:"created_at"`
}

type EntityType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type EntityHistory struct {
	IdentityId  int64  `json:"identity_id"`
	Action      string `json:"action"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

type Permission struct {
	Id          int64  `json:"id"`
	Code        int64  `json:"code"`
	Description string `json:"description"`
}

type Privilege struct {
	Kind        string `json:"kind"`
	WorkgroupId int64  `json:"workgroup_id"`
}

type EntityPrivilege struct {
	Kind                 string `json:"kind"`
	WorkgroupId          int64  `json:"workgroup_id"`
	WorkgroupName        string `json:"workgroup_name"`
	WorkgroupDescription string `json:"workgroup_description"`
}

type Role struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Created     int64  `json:"created"`
}

type Identity struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	LastLogin int64  `json:"last_login"`
	Created   int64  `json:"created"`
}

type Workgroup struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Created     int64  `json:"created"`
}

// --- Interfaces ---

type Az interface {
	Identify(r *http.Request) (az.Principal, error)
}

type Service interface {
	Ping(pz az.Principal, status bool) (bool, error)
	RegisterCluster(pz az.Principal, address string) (int64, error)
	UnregisterCluster(pz az.Principal, clusterId int64) error
	StartYarnCluster(pz az.Principal, clusterName string, engineId int64, size int, memory string, username string) (int64, error)
	StopYarnCluster(pz az.Principal, clusterId int64) error
	GetCluster(pz az.Principal, clusterId int64) (*Cluster, error)
	GetYarnCluster(pz az.Principal, clusterId int64) (*YarnCluster, error)
	GetClusters(pz az.Principal, offset int64, limit int64) ([]*Cluster, error)
	GetClusterStatus(pz az.Principal, clusterId int64) (*ClusterStatus, error)
	DeleteCluster(pz az.Principal, clusterId int64) error
	GetJob(pz az.Principal, clusterId int64, jobName string) (*Job, error)
	GetJobs(pz az.Principal, clusterId int64) ([]*Job, error)
	BuildModel(pz az.Principal, clusterId int64, dataset string, targetName string, maxRunTime int) (*Model, error)
	GetModel(pz az.Principal, modelId int64) (*Model, error)
	GetModels(pz az.Principal, offset int64, limit int64) ([]*Model, error)
	GetClusterModels(pz az.Principal, clusterId int64) ([]*Model, error)
	ImportModelFromCluster(pz az.Principal, clusterId int64, modelName string) (*Model, error)
	DeleteModel(pz az.Principal, modelId int64) error
	StartScoringService(pz az.Principal, modelId int64, port int) (*ScoringService, error)
	StopScoringService(pz az.Principal, serviceId int64) error
	GetScoringService(pz az.Principal, serviceId int64) (*ScoringService, error)
	GetScoringServices(pz az.Principal, offset int64, limit int64) ([]*ScoringService, error)
	DeleteScoringService(pz az.Principal, serviceId int64) error
	AddEngine(pz az.Principal, engineName string, enginePath string) (int64, error)
	GetEngine(pz az.Principal, engineId int64) (*Engine, error)
	GetEngines(pz az.Principal) ([]*Engine, error)
	DeleteEngine(pz az.Principal, engineId int64) error
	GetSupportedEntityTypes(pz az.Principal) ([]*EntityType, error)
	GetSupportedPermissions(pz az.Principal) ([]*Permission, error)
	GetPermissionsForRole(pz az.Principal, roleId int64) ([]*Permission, error)
	GetPermissionsForIdentity(pz az.Principal, identityId int64) ([]*Permission, error)
	CreateRole(pz az.Principal, name string, description string) (int64, error)
	GetRoles(pz az.Principal, offset int64, limit int64) ([]*Role, error)
	GetRolesForIdentity(pz az.Principal, identityId int64) ([]*Role, error)
	GetRole(pz az.Principal, roleId int64) (*Role, error)
	UpdateRole(pz az.Principal, roleId int64, name string, description string) error
	LinkRoleAndPermissions(pz az.Principal, roleId int64, permissionIds []int64) error
	DeleteRole(pz az.Principal, roleId int64) error
	CreateWorkgroup(pz az.Principal, name string, description string) (int64, error)
	GetWorkgroups(pz az.Principal, offset int64, limit int64) ([]*Workgroup, error)
	GetWorkgroupsForIdentity(pz az.Principal, identityId int64) ([]*Workgroup, error)
	GetWorkgroup(pz az.Principal, workgroupId int64) (*Workgroup, error)
	UpdateWorkgroup(pz az.Principal, workgroupId int64, name string, description string) error
	DeleteWorkgroup(pz az.Principal, workgroupId int64) error
	CreateIdentity(pz az.Principal, name string, password string) (int64, error)
	GetIdentities(pz az.Principal, offset int64, limit int64) ([]*Identity, error)
	GetIdentitiesForWorkgroup(pz az.Principal, workgroupId int64) ([]*Identity, error)
	GetIdentititesForRole(pz az.Principal, roleId int64) ([]*Identity, error)
	GetIdentity(pz az.Principal, identityId int64) (*Identity, error)
	LinkIdentityAndWorkgroup(pz az.Principal, identityId int64, workgroupId int64) error
	UnlinkIdentityAndWorkgroup(pz az.Principal, identityId int64, workgroupId int64) error
	LinkIdentityAndRole(pz az.Principal, identityId int64, roleId int64) error
	UnlinkIdentityAndRole(pz az.Principal, identityId int64, roleId int64) error
	DeactivateIdentity(pz az.Principal, identityId int64) error
	ShareEntity(pz az.Principal, kind string, workgroupId int64, entityTypeId int64, entityId int64) error
	GetEntityPrivileges(pz az.Principal, entityTypeId int64, entityId int64) ([]*EntityPrivilege, error)
	UnshareEntity(pz az.Principal, kind string, workgroupId int64, entityTypeId int64, entityId int64) error
	GetEntityHistory(pz az.Principal, entityTypeId int64, entityId int64, offset int64, limit int64) ([]*EntityHistory, error)
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
	EngineId    int64  `json:"engine_id"`
	Size        int    `json:"size"`
	Memory      string `json:"memory"`
	Username    string `json:"username"`
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
	Limit  int64 `json:"limit"`
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
	ClusterId int64  `json:"cluster_id"`
	JobName   string `json:"job_name"`
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
	ClusterId  int64  `json:"cluster_id"`
	Dataset    string `json:"dataset"`
	TargetName string `json:"target_name"`
	MaxRunTime int    `json:"max_run_time"`
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
	Limit  int64 `json:"limit"`
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

type ImportModelFromClusterIn struct {
	ClusterId int64  `json:"cluster_id"`
	ModelName string `json:"model_name"`
}

type ImportModelFromClusterOut struct {
	Model *Model `json:"model"`
}

type DeleteModelIn struct {
	ModelId int64 `json:"model_id"`
}

type DeleteModelOut struct {
}

type StartScoringServiceIn struct {
	ModelId int64 `json:"model_id"`
	Port    int   `json:"port"`
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
	Limit  int64 `json:"limit"`
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

type GetSupportedEntityTypesIn struct {
}

type GetSupportedEntityTypesOut struct {
	EntityTypes []*EntityType `json:"entity_types"`
}

type GetSupportedPermissionsIn struct {
}

type GetSupportedPermissionsOut struct {
	Permissions []*Permission `json:"permissions"`
}

type GetPermissionsForRoleIn struct {
	RoleId int64 `json:"role_id"`
}

type GetPermissionsForRoleOut struct {
	Permissions []*Permission `json:"permissions"`
}

type GetPermissionsForIdentityIn struct {
	IdentityId int64 `json:"identity_id"`
}

type GetPermissionsForIdentityOut struct {
	Permissions []*Permission `json:"permissions"`
}

type CreateRoleIn struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateRoleOut struct {
	RoleId int64 `json:"role_id"`
}

type GetRolesIn struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetRolesOut struct {
	Roles []*Role `json:"roles"`
}

type GetRolesForIdentityIn struct {
	IdentityId int64 `json:"identity_id"`
}

type GetRolesForIdentityOut struct {
	Roles []*Role `json:"roles"`
}

type GetRoleIn struct {
	RoleId int64 `json:"role_id"`
}

type GetRoleOut struct {
	Role *Role `json:"role"`
}

type UpdateRoleIn struct {
	RoleId      int64  `json:"role_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateRoleOut struct {
}

type LinkRoleAndPermissionsIn struct {
	RoleId        int64   `json:"role_id"`
	PermissionIds []int64 `json:"permission_ids"`
}

type LinkRoleAndPermissionsOut struct {
}

type DeleteRoleIn struct {
	RoleId int64 `json:"role_id"`
}

type DeleteRoleOut struct {
}

type CreateWorkgroupIn struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateWorkgroupOut struct {
	WorkgroupId int64 `json:"workgroup_id"`
}

type GetWorkgroupsIn struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetWorkgroupsOut struct {
	Workgroups []*Workgroup `json:"workgroups"`
}

type GetWorkgroupsForIdentityIn struct {
	IdentityId int64 `json:"identity_id"`
}

type GetWorkgroupsForIdentityOut struct {
	Workgroups []*Workgroup `json:"workgroups"`
}

type GetWorkgroupIn struct {
	WorkgroupId int64 `json:"workgroup_id"`
}

type GetWorkgroupOut struct {
	Workgroup *Workgroup `json:"workgroup"`
}

type UpdateWorkgroupIn struct {
	WorkgroupId int64  `json:"workgroup_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateWorkgroupOut struct {
}

type DeleteWorkgroupIn struct {
	WorkgroupId int64 `json:"workgroup_id"`
}

type DeleteWorkgroupOut struct {
}

type CreateIdentityIn struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateIdentityOut struct {
	IdentityId int64 `json:"identity_id"`
}

type GetIdentitiesIn struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetIdentitiesOut struct {
	Identities []*Identity `json:"identities"`
}

type GetIdentitiesForWorkgroupIn struct {
	WorkgroupId int64 `json:"workgroup_id"`
}

type GetIdentitiesForWorkgroupOut struct {
	Identities []*Identity `json:"identities"`
}

type GetIdentititesForRoleIn struct {
	RoleId int64 `json:"role_id"`
}

type GetIdentititesForRoleOut struct {
	Identities []*Identity `json:"identities"`
}

type GetIdentityIn struct {
	IdentityId int64 `json:"identity_id"`
}

type GetIdentityOut struct {
	Identity *Identity `json:"identity"`
}

type LinkIdentityAndWorkgroupIn struct {
	IdentityId  int64 `json:"identity_id"`
	WorkgroupId int64 `json:"workgroup_id"`
}

type LinkIdentityAndWorkgroupOut struct {
}

type UnlinkIdentityAndWorkgroupIn struct {
	IdentityId  int64 `json:"identity_id"`
	WorkgroupId int64 `json:"workgroup_id"`
}

type UnlinkIdentityAndWorkgroupOut struct {
}

type LinkIdentityAndRoleIn struct {
	IdentityId int64 `json:"identity_id"`
	RoleId     int64 `json:"role_id"`
}

type LinkIdentityAndRoleOut struct {
}

type UnlinkIdentityAndRoleIn struct {
	IdentityId int64 `json:"identity_id"`
	RoleId     int64 `json:"role_id"`
}

type UnlinkIdentityAndRoleOut struct {
}

type DeactivateIdentityIn struct {
	IdentityId int64 `json:"identity_id"`
}

type DeactivateIdentityOut struct {
}

type ShareEntityIn struct {
	Kind         string `json:"kind"`
	WorkgroupId  int64  `json:"workgroup_id"`
	EntityTypeId int64  `json:"entity_type_id"`
	EntityId     int64  `json:"entity_id"`
}

type ShareEntityOut struct {
}

type GetEntityPrivilegesIn struct {
	EntityTypeId int64 `json:"entity_type_id"`
	EntityId     int64 `json:"entity_id"`
}

type GetEntityPrivilegesOut struct {
	Privileges []*EntityPrivilege `json:"privileges"`
}

type UnshareEntityIn struct {
	Kind         string `json:"kind"`
	WorkgroupId  int64  `json:"workgroup_id"`
	EntityTypeId int64  `json:"entity_type_id"`
	EntityId     int64  `json:"entity_id"`
}

type UnshareEntityOut struct {
}

type GetEntityHistoryIn struct {
	EntityTypeId int64 `json:"entity_type_id"`
	EntityId     int64 `json:"entity_id"`
	Offset       int64 `json:"offset"`
	Limit        int64 `json:"limit"`
}

type GetEntityHistoryOut struct {
	History []*EntityHistory `json:"history"`
}

// --- Client Stub ---

type Remote struct {
	Proc Proc
}

type Proc interface {
	Call(name string, in, out interface{}) error
}

func (this *Remote) Ping(status bool) (bool, error) {
	in := PingIn{status}
	var out PingOut
	err := this.Proc.Call("Ping", &in, &out)
	if err != nil {
		return false, err
	}
	return out.Status, nil
}

func (this *Remote) RegisterCluster(address string) (int64, error) {
	in := RegisterClusterIn{address}
	var out RegisterClusterOut
	err := this.Proc.Call("RegisterCluster", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.ClusterId, nil
}

func (this *Remote) UnregisterCluster(clusterId int64) error {
	in := UnregisterClusterIn{clusterId}
	var out UnregisterClusterOut
	err := this.Proc.Call("UnregisterCluster", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) StartYarnCluster(clusterName string, engineId int64, size int, memory string, username string) (int64, error) {
	in := StartYarnClusterIn{clusterName, engineId, size, memory, username}
	var out StartYarnClusterOut
	err := this.Proc.Call("StartYarnCluster", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.ClusterId, nil
}

func (this *Remote) StopYarnCluster(clusterId int64) error {
	in := StopYarnClusterIn{clusterId}
	var out StopYarnClusterOut
	err := this.Proc.Call("StopYarnCluster", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetCluster(clusterId int64) (*Cluster, error) {
	in := GetClusterIn{clusterId}
	var out GetClusterOut
	err := this.Proc.Call("GetCluster", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Cluster, nil
}

func (this *Remote) GetYarnCluster(clusterId int64) (*YarnCluster, error) {
	in := GetYarnClusterIn{clusterId}
	var out GetYarnClusterOut
	err := this.Proc.Call("GetYarnCluster", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Cluster, nil
}

func (this *Remote) GetClusters(offset int64, limit int64) ([]*Cluster, error) {
	in := GetClustersIn{offset, limit}
	var out GetClustersOut
	err := this.Proc.Call("GetClusters", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Clusters, nil
}

func (this *Remote) GetClusterStatus(clusterId int64) (*ClusterStatus, error) {
	in := GetClusterStatusIn{clusterId}
	var out GetClusterStatusOut
	err := this.Proc.Call("GetClusterStatus", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.ClusterStatus, nil
}

func (this *Remote) DeleteCluster(clusterId int64) error {
	in := DeleteClusterIn{clusterId}
	var out DeleteClusterOut
	err := this.Proc.Call("DeleteCluster", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetJob(clusterId int64, jobName string) (*Job, error) {
	in := GetJobIn{clusterId, jobName}
	var out GetJobOut
	err := this.Proc.Call("GetJob", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Job, nil
}

func (this *Remote) GetJobs(clusterId int64) ([]*Job, error) {
	in := GetJobsIn{clusterId}
	var out GetJobsOut
	err := this.Proc.Call("GetJobs", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Jobs, nil
}

func (this *Remote) BuildModel(clusterId int64, dataset string, targetName string, maxRunTime int) (*Model, error) {
	in := BuildModelIn{clusterId, dataset, targetName, maxRunTime}
	var out BuildModelOut
	err := this.Proc.Call("BuildModel", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Model, nil
}

func (this *Remote) GetModel(modelId int64) (*Model, error) {
	in := GetModelIn{modelId}
	var out GetModelOut
	err := this.Proc.Call("GetModel", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Model, nil
}

func (this *Remote) GetModels(offset int64, limit int64) ([]*Model, error) {
	in := GetModelsIn{offset, limit}
	var out GetModelsOut
	err := this.Proc.Call("GetModels", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Models, nil
}

func (this *Remote) GetClusterModels(clusterId int64) ([]*Model, error) {
	in := GetClusterModelsIn{clusterId}
	var out GetClusterModelsOut
	err := this.Proc.Call("GetClusterModels", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Models, nil
}

func (this *Remote) ImportModelFromCluster(clusterId int64, modelName string) (*Model, error) {
	in := ImportModelFromClusterIn{clusterId, modelName}
	var out ImportModelFromClusterOut
	err := this.Proc.Call("ImportModelFromCluster", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Model, nil
}

func (this *Remote) DeleteModel(modelId int64) error {
	in := DeleteModelIn{modelId}
	var out DeleteModelOut
	err := this.Proc.Call("DeleteModel", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) StartScoringService(modelId int64, port int) (*ScoringService, error) {
	in := StartScoringServiceIn{modelId, port}
	var out StartScoringServiceOut
	err := this.Proc.Call("StartScoringService", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Service, nil
}

func (this *Remote) StopScoringService(serviceId int64) error {
	in := StopScoringServiceIn{serviceId}
	var out StopScoringServiceOut
	err := this.Proc.Call("StopScoringService", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetScoringService(serviceId int64) (*ScoringService, error) {
	in := GetScoringServiceIn{serviceId}
	var out GetScoringServiceOut
	err := this.Proc.Call("GetScoringService", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Service, nil
}

func (this *Remote) GetScoringServices(offset int64, limit int64) ([]*ScoringService, error) {
	in := GetScoringServicesIn{offset, limit}
	var out GetScoringServicesOut
	err := this.Proc.Call("GetScoringServices", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Services, nil
}

func (this *Remote) DeleteScoringService(serviceId int64) error {
	in := DeleteScoringServiceIn{serviceId}
	var out DeleteScoringServiceOut
	err := this.Proc.Call("DeleteScoringService", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) AddEngine(engineName string, enginePath string) (int64, error) {
	in := AddEngineIn{engineName, enginePath}
	var out AddEngineOut
	err := this.Proc.Call("AddEngine", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.EngineId, nil
}

func (this *Remote) GetEngine(engineId int64) (*Engine, error) {
	in := GetEngineIn{engineId}
	var out GetEngineOut
	err := this.Proc.Call("GetEngine", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Engine, nil
}

func (this *Remote) GetEngines() ([]*Engine, error) {
	in := GetEnginesIn{}
	var out GetEnginesOut
	err := this.Proc.Call("GetEngines", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Engines, nil
}

func (this *Remote) DeleteEngine(engineId int64) error {
	in := DeleteEngineIn{engineId}
	var out DeleteEngineOut
	err := this.Proc.Call("DeleteEngine", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetSupportedEntityTypes() ([]*EntityType, error) {
	in := GetSupportedEntityTypesIn{}
	var out GetSupportedEntityTypesOut
	err := this.Proc.Call("GetSupportedEntityTypes", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.EntityTypes, nil
}

func (this *Remote) GetSupportedPermissions() ([]*Permission, error) {
	in := GetSupportedPermissionsIn{}
	var out GetSupportedPermissionsOut
	err := this.Proc.Call("GetSupportedPermissions", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Permissions, nil
}

func (this *Remote) GetPermissionsForRole(roleId int64) ([]*Permission, error) {
	in := GetPermissionsForRoleIn{roleId}
	var out GetPermissionsForRoleOut
	err := this.Proc.Call("GetPermissionsForRole", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Permissions, nil
}

func (this *Remote) GetPermissionsForIdentity(identityId int64) ([]*Permission, error) {
	in := GetPermissionsForIdentityIn{identityId}
	var out GetPermissionsForIdentityOut
	err := this.Proc.Call("GetPermissionsForIdentity", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Permissions, nil
}

func (this *Remote) CreateRole(name string, description string) (int64, error) {
	in := CreateRoleIn{name, description}
	var out CreateRoleOut
	err := this.Proc.Call("CreateRole", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.RoleId, nil
}

func (this *Remote) GetRoles(offset int64, limit int64) ([]*Role, error) {
	in := GetRolesIn{offset, limit}
	var out GetRolesOut
	err := this.Proc.Call("GetRoles", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Roles, nil
}

func (this *Remote) GetRolesForIdentity(identityId int64) ([]*Role, error) {
	in := GetRolesForIdentityIn{identityId}
	var out GetRolesForIdentityOut
	err := this.Proc.Call("GetRolesForIdentity", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Roles, nil
}

func (this *Remote) GetRole(roleId int64) (*Role, error) {
	in := GetRoleIn{roleId}
	var out GetRoleOut
	err := this.Proc.Call("GetRole", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Role, nil
}

func (this *Remote) UpdateRole(roleId int64, name string, description string) error {
	in := UpdateRoleIn{roleId, name, description}
	var out UpdateRoleOut
	err := this.Proc.Call("UpdateRole", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) LinkRoleAndPermissions(roleId int64, permissionIds []int64) error {
	in := LinkRoleAndPermissionsIn{roleId, permissionIds}
	var out LinkRoleAndPermissionsOut
	err := this.Proc.Call("LinkRoleAndPermissions", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) DeleteRole(roleId int64) error {
	in := DeleteRoleIn{roleId}
	var out DeleteRoleOut
	err := this.Proc.Call("DeleteRole", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) CreateWorkgroup(name string, description string) (int64, error) {
	in := CreateWorkgroupIn{name, description}
	var out CreateWorkgroupOut
	err := this.Proc.Call("CreateWorkgroup", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.WorkgroupId, nil
}

func (this *Remote) GetWorkgroups(offset int64, limit int64) ([]*Workgroup, error) {
	in := GetWorkgroupsIn{offset, limit}
	var out GetWorkgroupsOut
	err := this.Proc.Call("GetWorkgroups", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Workgroups, nil
}

func (this *Remote) GetWorkgroupsForIdentity(identityId int64) ([]*Workgroup, error) {
	in := GetWorkgroupsForIdentityIn{identityId}
	var out GetWorkgroupsForIdentityOut
	err := this.Proc.Call("GetWorkgroupsForIdentity", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Workgroups, nil
}

func (this *Remote) GetWorkgroup(workgroupId int64) (*Workgroup, error) {
	in := GetWorkgroupIn{workgroupId}
	var out GetWorkgroupOut
	err := this.Proc.Call("GetWorkgroup", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Workgroup, nil
}

func (this *Remote) UpdateWorkgroup(workgroupId int64, name string, description string) error {
	in := UpdateWorkgroupIn{workgroupId, name, description}
	var out UpdateWorkgroupOut
	err := this.Proc.Call("UpdateWorkgroup", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) DeleteWorkgroup(workgroupId int64) error {
	in := DeleteWorkgroupIn{workgroupId}
	var out DeleteWorkgroupOut
	err := this.Proc.Call("DeleteWorkgroup", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) CreateIdentity(name string, password string) (int64, error) {
	in := CreateIdentityIn{name, password}
	var out CreateIdentityOut
	err := this.Proc.Call("CreateIdentity", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.IdentityId, nil
}

func (this *Remote) GetIdentities(offset int64, limit int64) ([]*Identity, error) {
	in := GetIdentitiesIn{offset, limit}
	var out GetIdentitiesOut
	err := this.Proc.Call("GetIdentities", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Identities, nil
}

func (this *Remote) GetIdentitiesForWorkgroup(workgroupId int64) ([]*Identity, error) {
	in := GetIdentitiesForWorkgroupIn{workgroupId}
	var out GetIdentitiesForWorkgroupOut
	err := this.Proc.Call("GetIdentitiesForWorkgroup", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Identities, nil
}

func (this *Remote) GetIdentititesForRole(roleId int64) ([]*Identity, error) {
	in := GetIdentititesForRoleIn{roleId}
	var out GetIdentititesForRoleOut
	err := this.Proc.Call("GetIdentititesForRole", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Identities, nil
}

func (this *Remote) GetIdentity(identityId int64) (*Identity, error) {
	in := GetIdentityIn{identityId}
	var out GetIdentityOut
	err := this.Proc.Call("GetIdentity", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Identity, nil
}

func (this *Remote) LinkIdentityAndWorkgroup(identityId int64, workgroupId int64) error {
	in := LinkIdentityAndWorkgroupIn{identityId, workgroupId}
	var out LinkIdentityAndWorkgroupOut
	err := this.Proc.Call("LinkIdentityAndWorkgroup", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) UnlinkIdentityAndWorkgroup(identityId int64, workgroupId int64) error {
	in := UnlinkIdentityAndWorkgroupIn{identityId, workgroupId}
	var out UnlinkIdentityAndWorkgroupOut
	err := this.Proc.Call("UnlinkIdentityAndWorkgroup", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) LinkIdentityAndRole(identityId int64, roleId int64) error {
	in := LinkIdentityAndRoleIn{identityId, roleId}
	var out LinkIdentityAndRoleOut
	err := this.Proc.Call("LinkIdentityAndRole", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) UnlinkIdentityAndRole(identityId int64, roleId int64) error {
	in := UnlinkIdentityAndRoleIn{identityId, roleId}
	var out UnlinkIdentityAndRoleOut
	err := this.Proc.Call("UnlinkIdentityAndRole", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) DeactivateIdentity(identityId int64) error {
	in := DeactivateIdentityIn{identityId}
	var out DeactivateIdentityOut
	err := this.Proc.Call("DeactivateIdentity", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) ShareEntity(kind string, workgroupId int64, entityTypeId int64, entityId int64) error {
	in := ShareEntityIn{kind, workgroupId, entityTypeId, entityId}
	var out ShareEntityOut
	err := this.Proc.Call("ShareEntity", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetEntityPrivileges(entityTypeId int64, entityId int64) ([]*EntityPrivilege, error) {
	in := GetEntityPrivilegesIn{entityTypeId, entityId}
	var out GetEntityPrivilegesOut
	err := this.Proc.Call("GetEntityPrivileges", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Privileges, nil
}

func (this *Remote) UnshareEntity(kind string, workgroupId int64, entityTypeId int64, entityId int64) error {
	in := UnshareEntityIn{kind, workgroupId, entityTypeId, entityId}
	var out UnshareEntityOut
	err := this.Proc.Call("UnshareEntity", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetEntityHistory(entityTypeId int64, entityId int64, offset int64, limit int64) ([]*EntityHistory, error) {
	in := GetEntityHistoryIn{entityTypeId, entityId, offset, limit}
	var out GetEntityHistoryOut
	err := this.Proc.Call("GetEntityHistory", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.History, nil
}

// --- Server Stub ---

type Impl struct {
	Service Service
	Az      az.Az
}

func (this *Impl) Ping(r *http.Request, in *PingIn, out *PingOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.Ping(pz, in.Status)
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
	it, err := this.Service.RegisterCluster(pz, in.Address)
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
	err := this.Service.UnregisterCluster(pz, in.ClusterId)
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
	it, err := this.Service.StartYarnCluster(pz, in.ClusterName, in.EngineId, in.Size, in.Memory, in.Username)
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
	err := this.Service.StopYarnCluster(pz, in.ClusterId)
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
	it, err := this.Service.GetCluster(pz, in.ClusterId)
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
	it, err := this.Service.GetYarnCluster(pz, in.ClusterId)
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
	it, err := this.Service.GetClusters(pz, in.Offset, in.Limit)
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
	it, err := this.Service.GetClusterStatus(pz, in.ClusterId)
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
	err := this.Service.DeleteCluster(pz, in.ClusterId)
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
	it, err := this.Service.GetJob(pz, in.ClusterId, in.JobName)
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
	it, err := this.Service.GetJobs(pz, in.ClusterId)
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
	it, err := this.Service.BuildModel(pz, in.ClusterId, in.Dataset, in.TargetName, in.MaxRunTime)
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
	it, err := this.Service.GetModel(pz, in.ModelId)
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
	it, err := this.Service.GetModels(pz, in.Offset, in.Limit)
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
	it, err := this.Service.GetClusterModels(pz, in.ClusterId)
	if err != nil {
		return err
	}
	out.Models = it
	return nil
}

func (this *Impl) ImportModelFromCluster(r *http.Request, in *ImportModelFromClusterIn, out *ImportModelFromClusterOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.ImportModelFromCluster(pz, in.ClusterId, in.ModelName)
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
	err := this.Service.DeleteModel(pz, in.ModelId)
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
	it, err := this.Service.StartScoringService(pz, in.ModelId, in.Port)
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
	err := this.Service.StopScoringService(pz, in.ServiceId)
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
	it, err := this.Service.GetScoringService(pz, in.ServiceId)
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
	it, err := this.Service.GetScoringServices(pz, in.Offset, in.Limit)
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
	err := this.Service.DeleteScoringService(pz, in.ServiceId)
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
	it, err := this.Service.AddEngine(pz, in.EngineName, in.EnginePath)
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
	it, err := this.Service.GetEngine(pz, in.EngineId)
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
	it, err := this.Service.GetEngines(pz)
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
	err := this.Service.DeleteEngine(pz, in.EngineId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) GetSupportedEntityTypes(r *http.Request, in *GetSupportedEntityTypesIn, out *GetSupportedEntityTypesOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetSupportedEntityTypes(pz)
	if err != nil {
		return err
	}
	out.EntityTypes = it
	return nil
}

func (this *Impl) GetSupportedPermissions(r *http.Request, in *GetSupportedPermissionsIn, out *GetSupportedPermissionsOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetSupportedPermissions(pz)
	if err != nil {
		return err
	}
	out.Permissions = it
	return nil
}

func (this *Impl) GetPermissionsForRole(r *http.Request, in *GetPermissionsForRoleIn, out *GetPermissionsForRoleOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetPermissionsForRole(pz, in.RoleId)
	if err != nil {
		return err
	}
	out.Permissions = it
	return nil
}

func (this *Impl) GetPermissionsForIdentity(r *http.Request, in *GetPermissionsForIdentityIn, out *GetPermissionsForIdentityOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetPermissionsForIdentity(pz, in.IdentityId)
	if err != nil {
		return err
	}
	out.Permissions = it
	return nil
}

func (this *Impl) CreateRole(r *http.Request, in *CreateRoleIn, out *CreateRoleOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.CreateRole(pz, in.Name, in.Description)
	if err != nil {
		return err
	}
	out.RoleId = it
	return nil
}

func (this *Impl) GetRoles(r *http.Request, in *GetRolesIn, out *GetRolesOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetRoles(pz, in.Offset, in.Limit)
	if err != nil {
		return err
	}
	out.Roles = it
	return nil
}

func (this *Impl) GetRolesForIdentity(r *http.Request, in *GetRolesForIdentityIn, out *GetRolesForIdentityOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetRolesForIdentity(pz, in.IdentityId)
	if err != nil {
		return err
	}
	out.Roles = it
	return nil
}

func (this *Impl) GetRole(r *http.Request, in *GetRoleIn, out *GetRoleOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetRole(pz, in.RoleId)
	if err != nil {
		return err
	}
	out.Role = it
	return nil
}

func (this *Impl) UpdateRole(r *http.Request, in *UpdateRoleIn, out *UpdateRoleOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.UpdateRole(pz, in.RoleId, in.Name, in.Description)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) LinkRoleAndPermissions(r *http.Request, in *LinkRoleAndPermissionsIn, out *LinkRoleAndPermissionsOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.LinkRoleAndPermissions(pz, in.RoleId, in.PermissionIds)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) DeleteRole(r *http.Request, in *DeleteRoleIn, out *DeleteRoleOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.DeleteRole(pz, in.RoleId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) CreateWorkgroup(r *http.Request, in *CreateWorkgroupIn, out *CreateWorkgroupOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.CreateWorkgroup(pz, in.Name, in.Description)
	if err != nil {
		return err
	}
	out.WorkgroupId = it
	return nil
}

func (this *Impl) GetWorkgroups(r *http.Request, in *GetWorkgroupsIn, out *GetWorkgroupsOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetWorkgroups(pz, in.Offset, in.Limit)
	if err != nil {
		return err
	}
	out.Workgroups = it
	return nil
}

func (this *Impl) GetWorkgroupsForIdentity(r *http.Request, in *GetWorkgroupsForIdentityIn, out *GetWorkgroupsForIdentityOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetWorkgroupsForIdentity(pz, in.IdentityId)
	if err != nil {
		return err
	}
	out.Workgroups = it
	return nil
}

func (this *Impl) GetWorkgroup(r *http.Request, in *GetWorkgroupIn, out *GetWorkgroupOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetWorkgroup(pz, in.WorkgroupId)
	if err != nil {
		return err
	}
	out.Workgroup = it
	return nil
}

func (this *Impl) UpdateWorkgroup(r *http.Request, in *UpdateWorkgroupIn, out *UpdateWorkgroupOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.UpdateWorkgroup(pz, in.WorkgroupId, in.Name, in.Description)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) DeleteWorkgroup(r *http.Request, in *DeleteWorkgroupIn, out *DeleteWorkgroupOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.DeleteWorkgroup(pz, in.WorkgroupId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) CreateIdentity(r *http.Request, in *CreateIdentityIn, out *CreateIdentityOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.CreateIdentity(pz, in.Name, in.Password)
	if err != nil {
		return err
	}
	out.IdentityId = it
	return nil
}

func (this *Impl) GetIdentities(r *http.Request, in *GetIdentitiesIn, out *GetIdentitiesOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetIdentities(pz, in.Offset, in.Limit)
	if err != nil {
		return err
	}
	out.Identities = it
	return nil
}

func (this *Impl) GetIdentitiesForWorkgroup(r *http.Request, in *GetIdentitiesForWorkgroupIn, out *GetIdentitiesForWorkgroupOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetIdentitiesForWorkgroup(pz, in.WorkgroupId)
	if err != nil {
		return err
	}
	out.Identities = it
	return nil
}

func (this *Impl) GetIdentititesForRole(r *http.Request, in *GetIdentititesForRoleIn, out *GetIdentititesForRoleOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetIdentititesForRole(pz, in.RoleId)
	if err != nil {
		return err
	}
	out.Identities = it
	return nil
}

func (this *Impl) GetIdentity(r *http.Request, in *GetIdentityIn, out *GetIdentityOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetIdentity(pz, in.IdentityId)
	if err != nil {
		return err
	}
	out.Identity = it
	return nil
}

func (this *Impl) LinkIdentityAndWorkgroup(r *http.Request, in *LinkIdentityAndWorkgroupIn, out *LinkIdentityAndWorkgroupOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.LinkIdentityAndWorkgroup(pz, in.IdentityId, in.WorkgroupId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) UnlinkIdentityAndWorkgroup(r *http.Request, in *UnlinkIdentityAndWorkgroupIn, out *UnlinkIdentityAndWorkgroupOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.UnlinkIdentityAndWorkgroup(pz, in.IdentityId, in.WorkgroupId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) LinkIdentityAndRole(r *http.Request, in *LinkIdentityAndRoleIn, out *LinkIdentityAndRoleOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.LinkIdentityAndRole(pz, in.IdentityId, in.RoleId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) UnlinkIdentityAndRole(r *http.Request, in *UnlinkIdentityAndRoleIn, out *UnlinkIdentityAndRoleOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.UnlinkIdentityAndRole(pz, in.IdentityId, in.RoleId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) DeactivateIdentity(r *http.Request, in *DeactivateIdentityIn, out *DeactivateIdentityOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.DeactivateIdentity(pz, in.IdentityId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) ShareEntity(r *http.Request, in *ShareEntityIn, out *ShareEntityOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.ShareEntity(pz, in.Kind, in.WorkgroupId, in.EntityTypeId, in.EntityId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) GetEntityPrivileges(r *http.Request, in *GetEntityPrivilegesIn, out *GetEntityPrivilegesOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetEntityPrivileges(pz, in.EntityTypeId, in.EntityId)
	if err != nil {
		return err
	}
	out.Privileges = it
	return nil
}

func (this *Impl) UnshareEntity(r *http.Request, in *UnshareEntityIn, out *UnshareEntityOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	err := this.Service.UnshareEntity(pz, in.Kind, in.WorkgroupId, in.EntityTypeId, in.EntityId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) GetEntityHistory(r *http.Request, in *GetEntityHistoryIn, out *GetEntityHistoryOut) error {

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	it, err := this.Service.GetEntityHistory(pz, in.EntityTypeId, in.EntityId, in.Offset, in.Limit)
	if err != nil {
		return err
	}
	out.History = it
	return nil
}
