// ------------------------------
// --- This is generated code ---
// ---      DO NOT EDIT       ---
// ------------------------------



package web

import (
	"github.com/h2oai/steamY/master/az"
	"log"
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

type ClusterStatus struct {

  Version string `json:"version"`

  Status string `json:"status"`

  MaxMemory string `json:"max_memory"`

  TotalCpuCount int `json:"total_cpu_count"`

  TotalAllowedCpuCount int `json:"total_allowed_cpu_count"`

}

type ClusterType struct {

  Id int64 `json:"id"`

  Name string `json:"name"`

}

type Dataset struct {

  Id int64 `json:"id"`

  DatasourceId int64 `json:"datasource_id"`

  Name string `json:"name"`

  Description string `json:"description"`

  FrameName string `json:"frame_name"`

  ResponseColumnName string `json:"response_column_name"`

  Properties string `json:"properties"`

  CreatedAt int64 `json:"created_at"`

}

type Datasource struct {

  Id int64 `json:"id"`

  ProjectId int64 `json:"project_id"`

  Name string `json:"name"`

  Description string `json:"description"`

  Kind string `json:"kind"`

  Configuration string `json:"configuration"`

  CreatedAt int64 `json:"created_at"`

}

type Engine struct {

  Id int64 `json:"id"`

  Name string `json:"name"`

  Location string `json:"location"`

  CreatedAt int64 `json:"created_at"`

}

type EntityHistory struct {

  IdentityId int64 `json:"identity_id"`

  Action string `json:"action"`

  Description string `json:"description"`

  CreatedAt int64 `json:"created_at"`

}

type EntityPrivilege struct {

  Kind string `json:"kind"`

  WorkgroupId int64 `json:"workgroup_id"`

  WorkgroupName string `json:"workgroup_name"`

  WorkgroupDescription string `json:"workgroup_description"`

}

type EntityType struct {

  Id int64 `json:"id"`

  Name string `json:"name"`

}

type Identity struct {

  Id int64 `json:"id"`

  Name string `json:"name"`

  IsActive bool `json:"is_active"`

  LastLogin int64 `json:"last_login"`

  Created int64 `json:"created"`

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

  TrainingDatasetId int64 `json:"training_dataset_id"`

  ValidationDatasetId int64 `json:"validation_dataset_id"`

  Name string `json:"name"`

  ClusterName string `json:"cluster_name"`

  Algorithm string `json:"algorithm"`

  DatasetName string `json:"dataset_name"`

  ResponseColumnName string `json:"response_column_name"`

  LogicalName string `json:"logical_name"`

  Location string `json:"location"`

  MaxRuntime int `json:"max_runtime"`

  Metrics string `json:"metrics"`

  CreatedAt int64 `json:"created_at"`

}

type Permission struct {

  Id int64 `json:"id"`

  Code string `json:"code"`

  Description string `json:"description"`

}

type Project struct {

  Id int64 `json:"id"`

  Name string `json:"name"`

  Description string `json:"description"`

  CreatedAt int64 `json:"created_at"`

}

type Role struct {

  Id int64 `json:"id"`

  Name string `json:"name"`

  Description string `json:"description"`

  Created int64 `json:"created"`

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

type Workgroup struct {

  Id int64 `json:"id"`

  Name string `json:"name"`

  Description string `json:"description"`

  Created int64 `json:"created"`

}

type YarnCluster struct {

  Id int64 `json:"id"`

  EngineId int64 `json:"engine_id"`

  Size int `json:"size"`

  ApplicationId string `json:"application_id"`

  Memory string `json:"memory"`

  Username string `json:"username"`

}


// --- Interface ---

type Az interface {
	Identify(r *http.Request) (az.Principal, error)
}


type Service interface {
  
  Ping(pz az.Principal, input bool) (bool, error)
  
  RegisterCluster(pz az.Principal, address string) (int64, error)
  
  UnregisterCluster(pz az.Principal, clusterId int64) (error)
  
  StartYarnCluster(pz az.Principal, clusterName string, engineId int64, size int, memory string, username string) (int64, error)
  
  StopYarnCluster(pz az.Principal, clusterId int64) (error)
  
  GetCluster(pz az.Principal, clusterId int64) (*Cluster, error)
  
  GetYarnCluster(pz az.Principal, clusterId int64) (*YarnCluster, error)
  
  GetClusters(pz az.Principal, offset int64, limit int64) ([]*Cluster, error)
  
  GetClusterStatus(pz az.Principal, clusterId int64) (*ClusterStatus, error)
  
  DeleteCluster(pz az.Principal, clusterId int64) (error)
  
  GetJob(pz az.Principal, clusterId int64, jobName string) (*Job, error)
  
  GetJobs(pz az.Principal, clusterId int64) ([]*Job, error)
  
  CreateProject(pz az.Principal, name string, description string) (int64, error)
  
  GetProjects(pz az.Principal, offset int64, limit int64) ([]*Project, error)
  
  GetProject(pz az.Principal, projectId int64) (*Project, error)
  
  DeleteProject(pz az.Principal, projectId int64) (error)
  
  CreateDatasource(pz az.Principal, projectId int64, name string, description string, path string) (int64, error)
  
  GetDatasources(pz az.Principal, projectId int64, offset int64, limit int64) ([]*Datasource, error)
  
  GetDatasource(pz az.Principal, datasourceId int64) (*Datasource, error)
  
  UpdateDatasource(pz az.Principal, datasourceId int64, name string, description string, path string) (error)
  
  DeleteDatasource(pz az.Principal, datasourceId int64) (error)
  
  CreateDataset(pz az.Principal, clusterId int64, datasourceId int64, name string, description string, responseColumnName string) (int64, error)
  
  GetDatasets(pz az.Principal, datasourceId int64, offset int64, limit int64) ([]*Dataset, error)
  
  GetDataset(pz az.Principal, datasetId int64) (*Dataset, error)
  
  UpdateDataset(pz az.Principal, datasetId int64, name string, description string, responseColumnName string) (error)
  
  SplitDataset(pz az.Principal, datasetId int64, ratio1 int, ratio2 int) ([]int64, error)
  
  DeleteDataset(pz az.Principal, datasetId int64) (error)
  
  BuildModel(pz az.Principal, clusterId int64, datasetId int64, algorithm string) (int64, error)
  
  BuildAutoModel(pz az.Principal, clusterId int64, dataset string, targetName string, maxRunTime int) (*Model, error)
  
  GetModel(pz az.Principal, modelId int64) (*Model, error)
  
  GetModels(pz az.Principal, projectId int64, offset int64, limit int64) ([]*Model, error)
  
  GetClusterModels(pz az.Principal, clusterId int64) ([]*Model, error)
  
  ImportModelFromCluster(pz az.Principal, clusterId int64, projectId int64, modelName string) (*Model, error)
  
  DeleteModel(pz az.Principal, modelId int64) (error)
  
  StartScoringService(pz az.Principal, modelId int64, port int) (*ScoringService, error)
  
  StopScoringService(pz az.Principal, serviceId int64) (error)
  
  GetScoringService(pz az.Principal, serviceId int64) (*ScoringService, error)
  
  GetScoringServices(pz az.Principal, offset int64, limit int64) ([]*ScoringService, error)
  
  GetScoringServicesForModel(pz az.Principal, modelId int64, offset int64, limit int64) ([]*ScoringService, error)
  
  DeleteScoringService(pz az.Principal, serviceId int64) (error)
  
  AddEngine(pz az.Principal, engineName string, enginePath string) (int64, error)
  
  GetEngine(pz az.Principal, engineId int64) (*Engine, error)
  
  GetEngines(pz az.Principal) ([]*Engine, error)
  
  DeleteEngine(pz az.Principal, engineId int64) (error)
  
  GetSupportedEntityTypes(pz az.Principal) ([]*EntityType, error)
  
  GetSupportedPermissions(pz az.Principal) ([]*Permission, error)
  
  GetSupportedClusterTypes(pz az.Principal) ([]*ClusterType, error)
  
  GetPermissionsForRole(pz az.Principal, roleId int64) ([]*Permission, error)
  
  GetPermissionsForIdentity(pz az.Principal, identityId int64) ([]*Permission, error)
  
  CreateRole(pz az.Principal, name string, description string) (int64, error)
  
  GetRoles(pz az.Principal, offset int64, limit int64) ([]*Role, error)
  
  GetRolesForIdentity(pz az.Principal, identityId int64) ([]*Role, error)
  
  GetRole(pz az.Principal, roleId int64) (*Role, error)
  
  GetRoleByName(pz az.Principal, name string) (*Role, error)
  
  UpdateRole(pz az.Principal, roleId int64, name string, description string) (error)
  
  LinkRoleAndPermissions(pz az.Principal, roleId int64, permissionIds []int64) (error)
  
  DeleteRole(pz az.Principal, roleId int64) (error)
  
  CreateWorkgroup(pz az.Principal, name string, description string) (int64, error)
  
  GetWorkgroups(pz az.Principal, offset int64, limit int64) ([]*Workgroup, error)
  
  GetWorkgroupsForIdentity(pz az.Principal, identityId int64) ([]*Workgroup, error)
  
  GetWorkgroup(pz az.Principal, workgroupId int64) (*Workgroup, error)
  
  GetWorkgroupByName(pz az.Principal, name string) (*Workgroup, error)
  
  UpdateWorkgroup(pz az.Principal, workgroupId int64, name string, description string) (error)
  
  DeleteWorkgroup(pz az.Principal, workgroupId int64) (error)
  
  CreateIdentity(pz az.Principal, name string, password string) (int64, error)
  
  GetIdentities(pz az.Principal, offset int64, limit int64) ([]*Identity, error)
  
  GetIdentitiesForWorkgroup(pz az.Principal, workgroupId int64) ([]*Identity, error)
  
  GetIdentitiesForRole(pz az.Principal, roleId int64) ([]*Identity, error)
  
  GetIdentity(pz az.Principal, identityId int64) (*Identity, error)
  
  GetIdentityByName(pz az.Principal, name string) (*Identity, error)
  
  LinkIdentityAndWorkgroup(pz az.Principal, identityId int64, workgroupId int64) (error)
  
  UnlinkIdentityAndWorkgroup(pz az.Principal, identityId int64, workgroupId int64) (error)
  
  LinkIdentityAndRole(pz az.Principal, identityId int64, roleId int64) (error)
  
  UnlinkIdentityAndRole(pz az.Principal, identityId int64, roleId int64) (error)
  
  UpdateIdentity(pz az.Principal, identityId int64, password string) (error)
  
  DeactivateIdentity(pz az.Principal, identityId int64) (error)
  
  ShareEntity(pz az.Principal, kind string, workgroupId int64, entityTypeId int64, entityId int64) (error)
  
  GetEntityPrivileges(pz az.Principal, entityTypeId int64, entityId int64) ([]*EntityPrivilege, error)
  
  UnshareEntity(pz az.Principal, kind string, workgroupId int64, entityTypeId int64, entityId int64) (error)
  
  GetEntityHistory(pz az.Principal, entityTypeId int64, entityId int64, offset int64, limit int64) ([]*EntityHistory, error)
  
}



// --- Messages ---



type PingIn struct {
  
  Input bool `json:"input"`
  
}
type PingOut struct {
  
  Output bool `json:"output"`
  
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

type CreateProjectIn struct {
  
  Name string `json:"name"`
  
  Description string `json:"description"`
  
}
type CreateProjectOut struct {
  
  ProjectId int64 `json:"project_id"`
  
}

type GetProjectsIn struct {
  
  Offset int64 `json:"offset"`
  
  Limit int64 `json:"limit"`
  
}
type GetProjectsOut struct {
  
  Projects []*Project `json:"projects"`
  
}

type GetProjectIn struct {
  
  ProjectId int64 `json:"project_id"`
  
}
type GetProjectOut struct {
  
  Project *Project `json:"project"`
  
}

type DeleteProjectIn struct {
  
  ProjectId int64 `json:"project_id"`
  
}
type DeleteProjectOut struct {
  
}

type CreateDatasourceIn struct {
  
  ProjectId int64 `json:"project_id"`
  
  Name string `json:"name"`
  
  Description string `json:"description"`
  
  Path string `json:"path"`
  
}
type CreateDatasourceOut struct {
  
  DatasourceId int64 `json:"datasource_id"`
  
}

type GetDatasourcesIn struct {
  
  ProjectId int64 `json:"project_id"`
  
  Offset int64 `json:"offset"`
  
  Limit int64 `json:"limit"`
  
}
type GetDatasourcesOut struct {
  
  Datasources []*Datasource `json:"datasources"`
  
}

type GetDatasourceIn struct {
  
  DatasourceId int64 `json:"datasource_id"`
  
}
type GetDatasourceOut struct {
  
  Datasource *Datasource `json:"datasource"`
  
}

type UpdateDatasourceIn struct {
  
  DatasourceId int64 `json:"datasource_id"`
  
  Name string `json:"name"`
  
  Description string `json:"description"`
  
  Path string `json:"path"`
  
}
type UpdateDatasourceOut struct {
  
}

type DeleteDatasourceIn struct {
  
  DatasourceId int64 `json:"datasource_id"`
  
}
type DeleteDatasourceOut struct {
  
}

type CreateDatasetIn struct {
  
  ClusterId int64 `json:"cluster_id"`
  
  DatasourceId int64 `json:"datasource_id"`
  
  Name string `json:"name"`
  
  Description string `json:"description"`
  
  ResponseColumnName string `json:"response_column_name"`
  
}
type CreateDatasetOut struct {
  
  DatasetId int64 `json:"dataset_id"`
  
}

type GetDatasetsIn struct {
  
  DatasourceId int64 `json:"datasource_id"`
  
  Offset int64 `json:"offset"`
  
  Limit int64 `json:"limit"`
  
}
type GetDatasetsOut struct {
  
  Datasets []*Dataset `json:"datasets"`
  
}

type GetDatasetIn struct {
  
  DatasetId int64 `json:"dataset_id"`
  
}
type GetDatasetOut struct {
  
  Dataset *Dataset `json:"dataset"`
  
}

type UpdateDatasetIn struct {
  
  DatasetId int64 `json:"dataset_id"`
  
  Name string `json:"name"`
  
  Description string `json:"description"`
  
  ResponseColumnName string `json:"response_column_name"`
  
}
type UpdateDatasetOut struct {
  
}

type SplitDatasetIn struct {
  
  DatasetId int64 `json:"dataset_id"`
  
  Ratio1 int `json:"ratio1"`
  
  Ratio2 int `json:"ratio2"`
  
}
type SplitDatasetOut struct {
  
  DatasetIds []int64 `json:"dataset_ids"`
  
}

type DeleteDatasetIn struct {
  
  DatasetId int64 `json:"dataset_id"`
  
}
type DeleteDatasetOut struct {
  
}

type BuildModelIn struct {
  
  ClusterId int64 `json:"cluster_id"`
  
  DatasetId int64 `json:"dataset_id"`
  
  Algorithm string `json:"algorithm"`
  
}
type BuildModelOut struct {
  
  ModelId int64 `json:"model_id"`
  
}

type BuildAutoModelIn struct {
  
  ClusterId int64 `json:"cluster_id"`
  
  Dataset string `json:"dataset"`
  
  TargetName string `json:"target_name"`
  
  MaxRunTime int `json:"max_run_time"`
  
}
type BuildAutoModelOut struct {
  
  Model *Model `json:"model"`
  
}

type GetModelIn struct {
  
  ModelId int64 `json:"model_id"`
  
}
type GetModelOut struct {
  
  Model *Model `json:"model"`
  
}

type GetModelsIn struct {
  
  ProjectId int64 `json:"project_id"`
  
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

type ImportModelFromClusterIn struct {
  
  ClusterId int64 `json:"cluster_id"`
  
  ProjectId int64 `json:"project_id"`
  
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

type GetScoringServicesForModelIn struct {
  
  ModelId int64 `json:"model_id"`
  
  Offset int64 `json:"offset"`
  
  Limit int64 `json:"limit"`
  
}
type GetScoringServicesForModelOut struct {
  
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

type GetSupportedClusterTypesIn struct {
  
}
type GetSupportedClusterTypesOut struct {
  
  ClusterTypes []*ClusterType `json:"cluster_types"`
  
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
  
  Name string `json:"name"`
  
  Description string `json:"description"`
  
}
type CreateRoleOut struct {
  
  RoleId int64 `json:"role_id"`
  
}

type GetRolesIn struct {
  
  Offset int64 `json:"offset"`
  
  Limit int64 `json:"limit"`
  
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

type GetRoleByNameIn struct {
  
  Name string `json:"name"`
  
}
type GetRoleByNameOut struct {
  
  Role *Role `json:"role"`
  
}

type UpdateRoleIn struct {
  
  RoleId int64 `json:"role_id"`
  
  Name string `json:"name"`
  
  Description string `json:"description"`
  
}
type UpdateRoleOut struct {
  
}

type LinkRoleAndPermissionsIn struct {
  
  RoleId int64 `json:"role_id"`
  
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
  
  Name string `json:"name"`
  
  Description string `json:"description"`
  
}
type CreateWorkgroupOut struct {
  
  WorkgroupId int64 `json:"workgroup_id"`
  
}

type GetWorkgroupsIn struct {
  
  Offset int64 `json:"offset"`
  
  Limit int64 `json:"limit"`
  
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

type GetWorkgroupByNameIn struct {
  
  Name string `json:"name"`
  
}
type GetWorkgroupByNameOut struct {
  
  Workgroup *Workgroup `json:"workgroup"`
  
}

type UpdateWorkgroupIn struct {
  
  WorkgroupId int64 `json:"workgroup_id"`
  
  Name string `json:"name"`
  
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
  
  Name string `json:"name"`
  
  Password string `json:"password"`
  
}
type CreateIdentityOut struct {
  
  IdentityId int64 `json:"identity_id"`
  
}

type GetIdentitiesIn struct {
  
  Offset int64 `json:"offset"`
  
  Limit int64 `json:"limit"`
  
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

type GetIdentitiesForRoleIn struct {
  
  RoleId int64 `json:"role_id"`
  
}
type GetIdentitiesForRoleOut struct {
  
  Identities []*Identity `json:"identities"`
  
}

type GetIdentityIn struct {
  
  IdentityId int64 `json:"identity_id"`
  
}
type GetIdentityOut struct {
  
  Identity *Identity `json:"identity"`
  
}

type GetIdentityByNameIn struct {
  
  Name string `json:"name"`
  
}
type GetIdentityByNameOut struct {
  
  Identity *Identity `json:"identity"`
  
}

type LinkIdentityAndWorkgroupIn struct {
  
  IdentityId int64 `json:"identity_id"`
  
  WorkgroupId int64 `json:"workgroup_id"`
  
}
type LinkIdentityAndWorkgroupOut struct {
  
}

type UnlinkIdentityAndWorkgroupIn struct {
  
  IdentityId int64 `json:"identity_id"`
  
  WorkgroupId int64 `json:"workgroup_id"`
  
}
type UnlinkIdentityAndWorkgroupOut struct {
  
}

type LinkIdentityAndRoleIn struct {
  
  IdentityId int64 `json:"identity_id"`
  
  RoleId int64 `json:"role_id"`
  
}
type LinkIdentityAndRoleOut struct {
  
}

type UnlinkIdentityAndRoleIn struct {
  
  IdentityId int64 `json:"identity_id"`
  
  RoleId int64 `json:"role_id"`
  
}
type UnlinkIdentityAndRoleOut struct {
  
}

type UpdateIdentityIn struct {
  
  IdentityId int64 `json:"identity_id"`
  
  Password string `json:"password"`
  
}
type UpdateIdentityOut struct {
  
}

type DeactivateIdentityIn struct {
  
  IdentityId int64 `json:"identity_id"`
  
}
type DeactivateIdentityOut struct {
  
}

type ShareEntityIn struct {
  
  Kind string `json:"kind"`
  
  WorkgroupId int64 `json:"workgroup_id"`
  
  EntityTypeId int64 `json:"entity_type_id"`
  
  EntityId int64 `json:"entity_id"`
  
}
type ShareEntityOut struct {
  
}

type GetEntityPrivilegesIn struct {
  
  EntityTypeId int64 `json:"entity_type_id"`
  
  EntityId int64 `json:"entity_id"`
  
}
type GetEntityPrivilegesOut struct {
  
  Privileges []*EntityPrivilege `json:"privileges"`
  
}

type UnshareEntityIn struct {
  
  Kind string `json:"kind"`
  
  WorkgroupId int64 `json:"workgroup_id"`
  
  EntityTypeId int64 `json:"entity_type_id"`
  
  EntityId int64 `json:"entity_id"`
  
}
type UnshareEntityOut struct {
  
}

type GetEntityHistoryIn struct {
  
  EntityTypeId int64 `json:"entity_type_id"`
  
  EntityId int64 `json:"entity_id"`
  
  Offset int64 `json:"offset"`
  
  Limit int64 `json:"limit"`
  
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



func (this *Remote) Ping(input bool) (bool, error) {
  in := PingIn{ input  }
  var out PingOut
  err := this.Proc.Call("Ping", &in, &out)
  if err != nil {
    return false, err
  }
  return out.Output, nil
}

func (this *Remote) RegisterCluster(address string) (int64, error) {
  in := RegisterClusterIn{ address  }
  var out RegisterClusterOut
  err := this.Proc.Call("RegisterCluster", &in, &out)
  if err != nil {
    return 0, err
  }
  return out.ClusterId, nil
}

func (this *Remote) UnregisterCluster(clusterId int64) (error) {
  in := UnregisterClusterIn{ clusterId  }
  var out UnregisterClusterOut
  err := this.Proc.Call("UnregisterCluster", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) StartYarnCluster(clusterName string, engineId int64, size int, memory string, username string) (int64, error) {
  in := StartYarnClusterIn{ clusterName , engineId , size , memory , username  }
  var out StartYarnClusterOut
  err := this.Proc.Call("StartYarnCluster", &in, &out)
  if err != nil {
    return 0, err
  }
  return out.ClusterId, nil
}

func (this *Remote) StopYarnCluster(clusterId int64) (error) {
  in := StopYarnClusterIn{ clusterId  }
  var out StopYarnClusterOut
  err := this.Proc.Call("StopYarnCluster", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) GetCluster(clusterId int64) (*Cluster, error) {
  in := GetClusterIn{ clusterId  }
  var out GetClusterOut
  err := this.Proc.Call("GetCluster", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Cluster, nil
}

func (this *Remote) GetYarnCluster(clusterId int64) (*YarnCluster, error) {
  in := GetYarnClusterIn{ clusterId  }
  var out GetYarnClusterOut
  err := this.Proc.Call("GetYarnCluster", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Cluster, nil
}

func (this *Remote) GetClusters(offset int64, limit int64) ([]*Cluster, error) {
  in := GetClustersIn{ offset , limit  }
  var out GetClustersOut
  err := this.Proc.Call("GetClusters", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Clusters, nil
}

func (this *Remote) GetClusterStatus(clusterId int64) (*ClusterStatus, error) {
  in := GetClusterStatusIn{ clusterId  }
  var out GetClusterStatusOut
  err := this.Proc.Call("GetClusterStatus", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.ClusterStatus, nil
}

func (this *Remote) DeleteCluster(clusterId int64) (error) {
  in := DeleteClusterIn{ clusterId  }
  var out DeleteClusterOut
  err := this.Proc.Call("DeleteCluster", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) GetJob(clusterId int64, jobName string) (*Job, error) {
  in := GetJobIn{ clusterId , jobName  }
  var out GetJobOut
  err := this.Proc.Call("GetJob", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Job, nil
}

func (this *Remote) GetJobs(clusterId int64) ([]*Job, error) {
  in := GetJobsIn{ clusterId  }
  var out GetJobsOut
  err := this.Proc.Call("GetJobs", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Jobs, nil
}

func (this *Remote) CreateProject(name string, description string) (int64, error) {
  in := CreateProjectIn{ name , description  }
  var out CreateProjectOut
  err := this.Proc.Call("CreateProject", &in, &out)
  if err != nil {
    return 0, err
  }
  return out.ProjectId, nil
}

func (this *Remote) GetProjects(offset int64, limit int64) ([]*Project, error) {
  in := GetProjectsIn{ offset , limit  }
  var out GetProjectsOut
  err := this.Proc.Call("GetProjects", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Projects, nil
}

func (this *Remote) GetProject(projectId int64) (*Project, error) {
  in := GetProjectIn{ projectId  }
  var out GetProjectOut
  err := this.Proc.Call("GetProject", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Project, nil
}

func (this *Remote) DeleteProject(projectId int64) (error) {
  in := DeleteProjectIn{ projectId  }
  var out DeleteProjectOut
  err := this.Proc.Call("DeleteProject", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) CreateDatasource(projectId int64, name string, description string, path string) (int64, error) {
  in := CreateDatasourceIn{ projectId , name , description , path  }
  var out CreateDatasourceOut
  err := this.Proc.Call("CreateDatasource", &in, &out)
  if err != nil {
    return 0, err
  }
  return out.DatasourceId, nil
}

func (this *Remote) GetDatasources(projectId int64, offset int64, limit int64) ([]*Datasource, error) {
  in := GetDatasourcesIn{ projectId , offset , limit  }
  var out GetDatasourcesOut
  err := this.Proc.Call("GetDatasources", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Datasources, nil
}

func (this *Remote) GetDatasource(datasourceId int64) (*Datasource, error) {
  in := GetDatasourceIn{ datasourceId  }
  var out GetDatasourceOut
  err := this.Proc.Call("GetDatasource", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Datasource, nil
}

func (this *Remote) UpdateDatasource(datasourceId int64, name string, description string, path string) (error) {
  in := UpdateDatasourceIn{ datasourceId , name , description , path  }
  var out UpdateDatasourceOut
  err := this.Proc.Call("UpdateDatasource", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) DeleteDatasource(datasourceId int64) (error) {
  in := DeleteDatasourceIn{ datasourceId  }
  var out DeleteDatasourceOut
  err := this.Proc.Call("DeleteDatasource", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) CreateDataset(clusterId int64, datasourceId int64, name string, description string, responseColumnName string) (int64, error) {
  in := CreateDatasetIn{ clusterId , datasourceId , name , description , responseColumnName  }
  var out CreateDatasetOut
  err := this.Proc.Call("CreateDataset", &in, &out)
  if err != nil {
    return 0, err
  }
  return out.DatasetId, nil
}

func (this *Remote) GetDatasets(datasourceId int64, offset int64, limit int64) ([]*Dataset, error) {
  in := GetDatasetsIn{ datasourceId , offset , limit  }
  var out GetDatasetsOut
  err := this.Proc.Call("GetDatasets", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Datasets, nil
}

func (this *Remote) GetDataset(datasetId int64) (*Dataset, error) {
  in := GetDatasetIn{ datasetId  }
  var out GetDatasetOut
  err := this.Proc.Call("GetDataset", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Dataset, nil
}

func (this *Remote) UpdateDataset(datasetId int64, name string, description string, responseColumnName string) (error) {
  in := UpdateDatasetIn{ datasetId , name , description , responseColumnName  }
  var out UpdateDatasetOut
  err := this.Proc.Call("UpdateDataset", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) SplitDataset(datasetId int64, ratio1 int, ratio2 int) ([]int64, error) {
  in := SplitDatasetIn{ datasetId , ratio1 , ratio2  }
  var out SplitDatasetOut
  err := this.Proc.Call("SplitDataset", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.DatasetIds, nil
}

func (this *Remote) DeleteDataset(datasetId int64) (error) {
  in := DeleteDatasetIn{ datasetId  }
  var out DeleteDatasetOut
  err := this.Proc.Call("DeleteDataset", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) BuildModel(clusterId int64, datasetId int64, algorithm string) (int64, error) {
  in := BuildModelIn{ clusterId , datasetId , algorithm  }
  var out BuildModelOut
  err := this.Proc.Call("BuildModel", &in, &out)
  if err != nil {
    return 0, err
  }
  return out.ModelId, nil
}

func (this *Remote) BuildAutoModel(clusterId int64, dataset string, targetName string, maxRunTime int) (*Model, error) {
  in := BuildAutoModelIn{ clusterId , dataset , targetName , maxRunTime  }
  var out BuildAutoModelOut
  err := this.Proc.Call("BuildAutoModel", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Model, nil
}

func (this *Remote) GetModel(modelId int64) (*Model, error) {
  in := GetModelIn{ modelId  }
  var out GetModelOut
  err := this.Proc.Call("GetModel", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Model, nil
}

func (this *Remote) GetModels(projectId int64, offset int64, limit int64) ([]*Model, error) {
  in := GetModelsIn{ projectId , offset , limit  }
  var out GetModelsOut
  err := this.Proc.Call("GetModels", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Models, nil
}

func (this *Remote) GetClusterModels(clusterId int64) ([]*Model, error) {
  in := GetClusterModelsIn{ clusterId  }
  var out GetClusterModelsOut
  err := this.Proc.Call("GetClusterModels", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Models, nil
}

func (this *Remote) ImportModelFromCluster(clusterId int64, projectId int64, modelName string) (*Model, error) {
  in := ImportModelFromClusterIn{ clusterId , projectId , modelName  }
  var out ImportModelFromClusterOut
  err := this.Proc.Call("ImportModelFromCluster", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Model, nil
}

func (this *Remote) DeleteModel(modelId int64) (error) {
  in := DeleteModelIn{ modelId  }
  var out DeleteModelOut
  err := this.Proc.Call("DeleteModel", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) StartScoringService(modelId int64, port int) (*ScoringService, error) {
  in := StartScoringServiceIn{ modelId , port  }
  var out StartScoringServiceOut
  err := this.Proc.Call("StartScoringService", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Service, nil
}

func (this *Remote) StopScoringService(serviceId int64) (error) {
  in := StopScoringServiceIn{ serviceId  }
  var out StopScoringServiceOut
  err := this.Proc.Call("StopScoringService", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) GetScoringService(serviceId int64) (*ScoringService, error) {
  in := GetScoringServiceIn{ serviceId  }
  var out GetScoringServiceOut
  err := this.Proc.Call("GetScoringService", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Service, nil
}

func (this *Remote) GetScoringServices(offset int64, limit int64) ([]*ScoringService, error) {
  in := GetScoringServicesIn{ offset , limit  }
  var out GetScoringServicesOut
  err := this.Proc.Call("GetScoringServices", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Services, nil
}

func (this *Remote) GetScoringServicesForModel(modelId int64, offset int64, limit int64) ([]*ScoringService, error) {
  in := GetScoringServicesForModelIn{ modelId , offset , limit  }
  var out GetScoringServicesForModelOut
  err := this.Proc.Call("GetScoringServicesForModel", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Services, nil
}

func (this *Remote) DeleteScoringService(serviceId int64) (error) {
  in := DeleteScoringServiceIn{ serviceId  }
  var out DeleteScoringServiceOut
  err := this.Proc.Call("DeleteScoringService", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) AddEngine(engineName string, enginePath string) (int64, error) {
  in := AddEngineIn{ engineName , enginePath  }
  var out AddEngineOut
  err := this.Proc.Call("AddEngine", &in, &out)
  if err != nil {
    return 0, err
  }
  return out.EngineId, nil
}

func (this *Remote) GetEngine(engineId int64) (*Engine, error) {
  in := GetEngineIn{ engineId  }
  var out GetEngineOut
  err := this.Proc.Call("GetEngine", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Engine, nil
}

func (this *Remote) GetEngines() ([]*Engine, error) {
  in := GetEnginesIn{  }
  var out GetEnginesOut
  err := this.Proc.Call("GetEngines", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Engines, nil
}

func (this *Remote) DeleteEngine(engineId int64) (error) {
  in := DeleteEngineIn{ engineId  }
  var out DeleteEngineOut
  err := this.Proc.Call("DeleteEngine", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) GetSupportedEntityTypes() ([]*EntityType, error) {
  in := GetSupportedEntityTypesIn{  }
  var out GetSupportedEntityTypesOut
  err := this.Proc.Call("GetSupportedEntityTypes", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.EntityTypes, nil
}

func (this *Remote) GetSupportedPermissions() ([]*Permission, error) {
  in := GetSupportedPermissionsIn{  }
  var out GetSupportedPermissionsOut
  err := this.Proc.Call("GetSupportedPermissions", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Permissions, nil
}

func (this *Remote) GetSupportedClusterTypes() ([]*ClusterType, error) {
  in := GetSupportedClusterTypesIn{  }
  var out GetSupportedClusterTypesOut
  err := this.Proc.Call("GetSupportedClusterTypes", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.ClusterTypes, nil
}

func (this *Remote) GetPermissionsForRole(roleId int64) ([]*Permission, error) {
  in := GetPermissionsForRoleIn{ roleId  }
  var out GetPermissionsForRoleOut
  err := this.Proc.Call("GetPermissionsForRole", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Permissions, nil
}

func (this *Remote) GetPermissionsForIdentity(identityId int64) ([]*Permission, error) {
  in := GetPermissionsForIdentityIn{ identityId  }
  var out GetPermissionsForIdentityOut
  err := this.Proc.Call("GetPermissionsForIdentity", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Permissions, nil
}

func (this *Remote) CreateRole(name string, description string) (int64, error) {
  in := CreateRoleIn{ name , description  }
  var out CreateRoleOut
  err := this.Proc.Call("CreateRole", &in, &out)
  if err != nil {
    return 0, err
  }
  return out.RoleId, nil
}

func (this *Remote) GetRoles(offset int64, limit int64) ([]*Role, error) {
  in := GetRolesIn{ offset , limit  }
  var out GetRolesOut
  err := this.Proc.Call("GetRoles", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Roles, nil
}

func (this *Remote) GetRolesForIdentity(identityId int64) ([]*Role, error) {
  in := GetRolesForIdentityIn{ identityId  }
  var out GetRolesForIdentityOut
  err := this.Proc.Call("GetRolesForIdentity", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Roles, nil
}

func (this *Remote) GetRole(roleId int64) (*Role, error) {
  in := GetRoleIn{ roleId  }
  var out GetRoleOut
  err := this.Proc.Call("GetRole", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Role, nil
}

func (this *Remote) GetRoleByName(name string) (*Role, error) {
  in := GetRoleByNameIn{ name  }
  var out GetRoleByNameOut
  err := this.Proc.Call("GetRoleByName", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Role, nil
}

func (this *Remote) UpdateRole(roleId int64, name string, description string) (error) {
  in := UpdateRoleIn{ roleId , name , description  }
  var out UpdateRoleOut
  err := this.Proc.Call("UpdateRole", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) LinkRoleAndPermissions(roleId int64, permissionIds []int64) (error) {
  in := LinkRoleAndPermissionsIn{ roleId , permissionIds  }
  var out LinkRoleAndPermissionsOut
  err := this.Proc.Call("LinkRoleAndPermissions", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) DeleteRole(roleId int64) (error) {
  in := DeleteRoleIn{ roleId  }
  var out DeleteRoleOut
  err := this.Proc.Call("DeleteRole", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) CreateWorkgroup(name string, description string) (int64, error) {
  in := CreateWorkgroupIn{ name , description  }
  var out CreateWorkgroupOut
  err := this.Proc.Call("CreateWorkgroup", &in, &out)
  if err != nil {
    return 0, err
  }
  return out.WorkgroupId, nil
}

func (this *Remote) GetWorkgroups(offset int64, limit int64) ([]*Workgroup, error) {
  in := GetWorkgroupsIn{ offset , limit  }
  var out GetWorkgroupsOut
  err := this.Proc.Call("GetWorkgroups", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Workgroups, nil
}

func (this *Remote) GetWorkgroupsForIdentity(identityId int64) ([]*Workgroup, error) {
  in := GetWorkgroupsForIdentityIn{ identityId  }
  var out GetWorkgroupsForIdentityOut
  err := this.Proc.Call("GetWorkgroupsForIdentity", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Workgroups, nil
}

func (this *Remote) GetWorkgroup(workgroupId int64) (*Workgroup, error) {
  in := GetWorkgroupIn{ workgroupId  }
  var out GetWorkgroupOut
  err := this.Proc.Call("GetWorkgroup", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Workgroup, nil
}

func (this *Remote) GetWorkgroupByName(name string) (*Workgroup, error) {
  in := GetWorkgroupByNameIn{ name  }
  var out GetWorkgroupByNameOut
  err := this.Proc.Call("GetWorkgroupByName", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Workgroup, nil
}

func (this *Remote) UpdateWorkgroup(workgroupId int64, name string, description string) (error) {
  in := UpdateWorkgroupIn{ workgroupId , name , description  }
  var out UpdateWorkgroupOut
  err := this.Proc.Call("UpdateWorkgroup", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) DeleteWorkgroup(workgroupId int64) (error) {
  in := DeleteWorkgroupIn{ workgroupId  }
  var out DeleteWorkgroupOut
  err := this.Proc.Call("DeleteWorkgroup", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) CreateIdentity(name string, password string) (int64, error) {
  in := CreateIdentityIn{ name , password  }
  var out CreateIdentityOut
  err := this.Proc.Call("CreateIdentity", &in, &out)
  if err != nil {
    return 0, err
  }
  return out.IdentityId, nil
}

func (this *Remote) GetIdentities(offset int64, limit int64) ([]*Identity, error) {
  in := GetIdentitiesIn{ offset , limit  }
  var out GetIdentitiesOut
  err := this.Proc.Call("GetIdentities", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Identities, nil
}

func (this *Remote) GetIdentitiesForWorkgroup(workgroupId int64) ([]*Identity, error) {
  in := GetIdentitiesForWorkgroupIn{ workgroupId  }
  var out GetIdentitiesForWorkgroupOut
  err := this.Proc.Call("GetIdentitiesForWorkgroup", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Identities, nil
}

func (this *Remote) GetIdentitiesForRole(roleId int64) ([]*Identity, error) {
  in := GetIdentitiesForRoleIn{ roleId  }
  var out GetIdentitiesForRoleOut
  err := this.Proc.Call("GetIdentitiesForRole", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Identities, nil
}

func (this *Remote) GetIdentity(identityId int64) (*Identity, error) {
  in := GetIdentityIn{ identityId  }
  var out GetIdentityOut
  err := this.Proc.Call("GetIdentity", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Identity, nil
}

func (this *Remote) GetIdentityByName(name string) (*Identity, error) {
  in := GetIdentityByNameIn{ name  }
  var out GetIdentityByNameOut
  err := this.Proc.Call("GetIdentityByName", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Identity, nil
}

func (this *Remote) LinkIdentityAndWorkgroup(identityId int64, workgroupId int64) (error) {
  in := LinkIdentityAndWorkgroupIn{ identityId , workgroupId  }
  var out LinkIdentityAndWorkgroupOut
  err := this.Proc.Call("LinkIdentityAndWorkgroup", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) UnlinkIdentityAndWorkgroup(identityId int64, workgroupId int64) (error) {
  in := UnlinkIdentityAndWorkgroupIn{ identityId , workgroupId  }
  var out UnlinkIdentityAndWorkgroupOut
  err := this.Proc.Call("UnlinkIdentityAndWorkgroup", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) LinkIdentityAndRole(identityId int64, roleId int64) (error) {
  in := LinkIdentityAndRoleIn{ identityId , roleId  }
  var out LinkIdentityAndRoleOut
  err := this.Proc.Call("LinkIdentityAndRole", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) UnlinkIdentityAndRole(identityId int64, roleId int64) (error) {
  in := UnlinkIdentityAndRoleIn{ identityId , roleId  }
  var out UnlinkIdentityAndRoleOut
  err := this.Proc.Call("UnlinkIdentityAndRole", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) UpdateIdentity(identityId int64, password string) (error) {
  in := UpdateIdentityIn{ identityId , password  }
  var out UpdateIdentityOut
  err := this.Proc.Call("UpdateIdentity", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) DeactivateIdentity(identityId int64) (error) {
  in := DeactivateIdentityIn{ identityId  }
  var out DeactivateIdentityOut
  err := this.Proc.Call("DeactivateIdentity", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) ShareEntity(kind string, workgroupId int64, entityTypeId int64, entityId int64) (error) {
  in := ShareEntityIn{ kind , workgroupId , entityTypeId , entityId  }
  var out ShareEntityOut
  err := this.Proc.Call("ShareEntity", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) GetEntityPrivileges(entityTypeId int64, entityId int64) ([]*EntityPrivilege, error) {
  in := GetEntityPrivilegesIn{ entityTypeId , entityId  }
  var out GetEntityPrivilegesOut
  err := this.Proc.Call("GetEntityPrivileges", &in, &out)
  if err != nil {
    return nil, err
  }
  return out.Privileges, nil
}

func (this *Remote) UnshareEntity(kind string, workgroupId int64, entityTypeId int64, entityId int64) (error) {
  in := UnshareEntityIn{ kind , workgroupId , entityTypeId , entityId  }
  var out UnshareEntityOut
  err := this.Proc.Call("UnshareEntity", &in, &out)
  if err != nil {
    return err
  }
  return nil
}

func (this *Remote) GetEntityHistory(entityTypeId int64, entityId int64, offset int64, limit int64) ([]*EntityHistory, error) {
  in := GetEntityHistoryIn{ entityTypeId , entityId , offset , limit  }
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
	log.Println(pz, "Ping")

	val0, err := this.Service.Ping(pz, in.Input)
	if err != nil {
		log.Printf("%s Failed Ping: %v", pz, err)
		return err
	}
  
	out.Output = val0 
  
	return nil
}

func (this *Impl) RegisterCluster(r *http.Request, in *RegisterClusterIn, out *RegisterClusterOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "RegisterCluster")

	val0, err := this.Service.RegisterCluster(pz, in.Address)
	if err != nil {
		log.Printf("%s Failed RegisterCluster: %v", pz, err)
		return err
	}
  
	out.ClusterId = val0 
  
	return nil
}

func (this *Impl) UnregisterCluster(r *http.Request, in *UnregisterClusterIn, out *UnregisterClusterOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "UnregisterCluster")

	err := this.Service.UnregisterCluster(pz, in.ClusterId)
	if err != nil {
		log.Printf("%s Failed UnregisterCluster: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) StartYarnCluster(r *http.Request, in *StartYarnClusterIn, out *StartYarnClusterOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "StartYarnCluster")

	val0, err := this.Service.StartYarnCluster(pz, in.ClusterName, in.EngineId, in.Size, in.Memory, in.Username)
	if err != nil {
		log.Printf("%s Failed StartYarnCluster: %v", pz, err)
		return err
	}
  
	out.ClusterId = val0 
  
	return nil
}

func (this *Impl) StopYarnCluster(r *http.Request, in *StopYarnClusterIn, out *StopYarnClusterOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "StopYarnCluster")

	err := this.Service.StopYarnCluster(pz, in.ClusterId)
	if err != nil {
		log.Printf("%s Failed StopYarnCluster: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) GetCluster(r *http.Request, in *GetClusterIn, out *GetClusterOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetCluster")

	val0, err := this.Service.GetCluster(pz, in.ClusterId)
	if err != nil {
		log.Printf("%s Failed GetCluster: %v", pz, err)
		return err
	}
  
	out.Cluster = val0 
  
	return nil
}

func (this *Impl) GetYarnCluster(r *http.Request, in *GetYarnClusterIn, out *GetYarnClusterOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetYarnCluster")

	val0, err := this.Service.GetYarnCluster(pz, in.ClusterId)
	if err != nil {
		log.Printf("%s Failed GetYarnCluster: %v", pz, err)
		return err
	}
  
	out.Cluster = val0 
  
	return nil
}

func (this *Impl) GetClusters(r *http.Request, in *GetClustersIn, out *GetClustersOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetClusters")

	val0, err := this.Service.GetClusters(pz, in.Offset, in.Limit)
	if err != nil {
		log.Printf("%s Failed GetClusters: %v", pz, err)
		return err
	}
  
	out.Clusters = val0 
  
	return nil
}

func (this *Impl) GetClusterStatus(r *http.Request, in *GetClusterStatusIn, out *GetClusterStatusOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetClusterStatus")

	val0, err := this.Service.GetClusterStatus(pz, in.ClusterId)
	if err != nil {
		log.Printf("%s Failed GetClusterStatus: %v", pz, err)
		return err
	}
  
	out.ClusterStatus = val0 
  
	return nil
}

func (this *Impl) DeleteCluster(r *http.Request, in *DeleteClusterIn, out *DeleteClusterOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "DeleteCluster")

	err := this.Service.DeleteCluster(pz, in.ClusterId)
	if err != nil {
		log.Printf("%s Failed DeleteCluster: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) GetJob(r *http.Request, in *GetJobIn, out *GetJobOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetJob")

	val0, err := this.Service.GetJob(pz, in.ClusterId, in.JobName)
	if err != nil {
		log.Printf("%s Failed GetJob: %v", pz, err)
		return err
	}
  
	out.Job = val0 
  
	return nil
}

func (this *Impl) GetJobs(r *http.Request, in *GetJobsIn, out *GetJobsOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetJobs")

	val0, err := this.Service.GetJobs(pz, in.ClusterId)
	if err != nil {
		log.Printf("%s Failed GetJobs: %v", pz, err)
		return err
	}
  
	out.Jobs = val0 
  
	return nil
}

func (this *Impl) CreateProject(r *http.Request, in *CreateProjectIn, out *CreateProjectOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "CreateProject")

	val0, err := this.Service.CreateProject(pz, in.Name, in.Description)
	if err != nil {
		log.Printf("%s Failed CreateProject: %v", pz, err)
		return err
	}
  
	out.ProjectId = val0 
  
	return nil
}

func (this *Impl) GetProjects(r *http.Request, in *GetProjectsIn, out *GetProjectsOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetProjects")

	val0, err := this.Service.GetProjects(pz, in.Offset, in.Limit)
	if err != nil {
		log.Printf("%s Failed GetProjects: %v", pz, err)
		return err
	}
  
	out.Projects = val0 
  
	return nil
}

func (this *Impl) GetProject(r *http.Request, in *GetProjectIn, out *GetProjectOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetProject")

	val0, err := this.Service.GetProject(pz, in.ProjectId)
	if err != nil {
		log.Printf("%s Failed GetProject: %v", pz, err)
		return err
	}
  
	out.Project = val0 
  
	return nil
}

func (this *Impl) DeleteProject(r *http.Request, in *DeleteProjectIn, out *DeleteProjectOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "DeleteProject")

	err := this.Service.DeleteProject(pz, in.ProjectId)
	if err != nil {
		log.Printf("%s Failed DeleteProject: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) CreateDatasource(r *http.Request, in *CreateDatasourceIn, out *CreateDatasourceOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "CreateDatasource")

	val0, err := this.Service.CreateDatasource(pz, in.ProjectId, in.Name, in.Description, in.Path)
	if err != nil {
		log.Printf("%s Failed CreateDatasource: %v", pz, err)
		return err
	}
  
	out.DatasourceId = val0 
  
	return nil
}

func (this *Impl) GetDatasources(r *http.Request, in *GetDatasourcesIn, out *GetDatasourcesOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetDatasources")

	val0, err := this.Service.GetDatasources(pz, in.ProjectId, in.Offset, in.Limit)
	if err != nil {
		log.Printf("%s Failed GetDatasources: %v", pz, err)
		return err
	}
  
	out.Datasources = val0 
  
	return nil
}

func (this *Impl) GetDatasource(r *http.Request, in *GetDatasourceIn, out *GetDatasourceOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetDatasource")

	val0, err := this.Service.GetDatasource(pz, in.DatasourceId)
	if err != nil {
		log.Printf("%s Failed GetDatasource: %v", pz, err)
		return err
	}
  
	out.Datasource = val0 
  
	return nil
}

func (this *Impl) UpdateDatasource(r *http.Request, in *UpdateDatasourceIn, out *UpdateDatasourceOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "UpdateDatasource")

	err := this.Service.UpdateDatasource(pz, in.DatasourceId, in.Name, in.Description, in.Path)
	if err != nil {
		log.Printf("%s Failed UpdateDatasource: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) DeleteDatasource(r *http.Request, in *DeleteDatasourceIn, out *DeleteDatasourceOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "DeleteDatasource")

	err := this.Service.DeleteDatasource(pz, in.DatasourceId)
	if err != nil {
		log.Printf("%s Failed DeleteDatasource: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) CreateDataset(r *http.Request, in *CreateDatasetIn, out *CreateDatasetOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "CreateDataset")

	val0, err := this.Service.CreateDataset(pz, in.ClusterId, in.DatasourceId, in.Name, in.Description, in.ResponseColumnName)
	if err != nil {
		log.Printf("%s Failed CreateDataset: %v", pz, err)
		return err
	}
  
	out.DatasetId = val0 
  
	return nil
}

func (this *Impl) GetDatasets(r *http.Request, in *GetDatasetsIn, out *GetDatasetsOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetDatasets")

	val0, err := this.Service.GetDatasets(pz, in.DatasourceId, in.Offset, in.Limit)
	if err != nil {
		log.Printf("%s Failed GetDatasets: %v", pz, err)
		return err
	}
  
	out.Datasets = val0 
  
	return nil
}

func (this *Impl) GetDataset(r *http.Request, in *GetDatasetIn, out *GetDatasetOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetDataset")

	val0, err := this.Service.GetDataset(pz, in.DatasetId)
	if err != nil {
		log.Printf("%s Failed GetDataset: %v", pz, err)
		return err
	}
  
	out.Dataset = val0 
  
	return nil
}

func (this *Impl) UpdateDataset(r *http.Request, in *UpdateDatasetIn, out *UpdateDatasetOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "UpdateDataset")

	err := this.Service.UpdateDataset(pz, in.DatasetId, in.Name, in.Description, in.ResponseColumnName)
	if err != nil {
		log.Printf("%s Failed UpdateDataset: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) SplitDataset(r *http.Request, in *SplitDatasetIn, out *SplitDatasetOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "SplitDataset")

	val0, err := this.Service.SplitDataset(pz, in.DatasetId, in.Ratio1, in.Ratio2)
	if err != nil {
		log.Printf("%s Failed SplitDataset: %v", pz, err)
		return err
	}
  
	out.DatasetIds = val0 
  
	return nil
}

func (this *Impl) DeleteDataset(r *http.Request, in *DeleteDatasetIn, out *DeleteDatasetOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "DeleteDataset")

	err := this.Service.DeleteDataset(pz, in.DatasetId)
	if err != nil {
		log.Printf("%s Failed DeleteDataset: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) BuildModel(r *http.Request, in *BuildModelIn, out *BuildModelOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "BuildModel")

	val0, err := this.Service.BuildModel(pz, in.ClusterId, in.DatasetId, in.Algorithm)
	if err != nil {
		log.Printf("%s Failed BuildModel: %v", pz, err)
		return err
	}
  
	out.ModelId = val0 
  
	return nil
}

func (this *Impl) BuildAutoModel(r *http.Request, in *BuildAutoModelIn, out *BuildAutoModelOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "BuildAutoModel")

	val0, err := this.Service.BuildAutoModel(pz, in.ClusterId, in.Dataset, in.TargetName, in.MaxRunTime)
	if err != nil {
		log.Printf("%s Failed BuildAutoModel: %v", pz, err)
		return err
	}
  
	out.Model = val0 
  
	return nil
}

func (this *Impl) GetModel(r *http.Request, in *GetModelIn, out *GetModelOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetModel")

	val0, err := this.Service.GetModel(pz, in.ModelId)
	if err != nil {
		log.Printf("%s Failed GetModel: %v", pz, err)
		return err
	}
  
	out.Model = val0 
  
	return nil
}

func (this *Impl) GetModels(r *http.Request, in *GetModelsIn, out *GetModelsOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetModels")

	val0, err := this.Service.GetModels(pz, in.ProjectId, in.Offset, in.Limit)
	if err != nil {
		log.Printf("%s Failed GetModels: %v", pz, err)
		return err
	}
  
	out.Models = val0 
  
	return nil
}

func (this *Impl) GetClusterModels(r *http.Request, in *GetClusterModelsIn, out *GetClusterModelsOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetClusterModels")

	val0, err := this.Service.GetClusterModels(pz, in.ClusterId)
	if err != nil {
		log.Printf("%s Failed GetClusterModels: %v", pz, err)
		return err
	}
  
	out.Models = val0 
  
	return nil
}

func (this *Impl) ImportModelFromCluster(r *http.Request, in *ImportModelFromClusterIn, out *ImportModelFromClusterOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "ImportModelFromCluster")

	val0, err := this.Service.ImportModelFromCluster(pz, in.ClusterId, in.ProjectId, in.ModelName)
	if err != nil {
		log.Printf("%s Failed ImportModelFromCluster: %v", pz, err)
		return err
	}
  
	out.Model = val0 
  
	return nil
}

func (this *Impl) DeleteModel(r *http.Request, in *DeleteModelIn, out *DeleteModelOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "DeleteModel")

	err := this.Service.DeleteModel(pz, in.ModelId)
	if err != nil {
		log.Printf("%s Failed DeleteModel: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) StartScoringService(r *http.Request, in *StartScoringServiceIn, out *StartScoringServiceOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "StartScoringService")

	val0, err := this.Service.StartScoringService(pz, in.ModelId, in.Port)
	if err != nil {
		log.Printf("%s Failed StartScoringService: %v", pz, err)
		return err
	}
  
	out.Service = val0 
  
	return nil
}

func (this *Impl) StopScoringService(r *http.Request, in *StopScoringServiceIn, out *StopScoringServiceOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "StopScoringService")

	err := this.Service.StopScoringService(pz, in.ServiceId)
	if err != nil {
		log.Printf("%s Failed StopScoringService: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) GetScoringService(r *http.Request, in *GetScoringServiceIn, out *GetScoringServiceOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetScoringService")

	val0, err := this.Service.GetScoringService(pz, in.ServiceId)
	if err != nil {
		log.Printf("%s Failed GetScoringService: %v", pz, err)
		return err
	}
  
	out.Service = val0 
  
	return nil
}

func (this *Impl) GetScoringServices(r *http.Request, in *GetScoringServicesIn, out *GetScoringServicesOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetScoringServices")

	val0, err := this.Service.GetScoringServices(pz, in.Offset, in.Limit)
	if err != nil {
		log.Printf("%s Failed GetScoringServices: %v", pz, err)
		return err
	}
  
	out.Services = val0 
  
	return nil
}

func (this *Impl) GetScoringServicesForModel(r *http.Request, in *GetScoringServicesForModelIn, out *GetScoringServicesForModelOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetScoringServicesForModel")

	val0, err := this.Service.GetScoringServicesForModel(pz, in.ModelId, in.Offset, in.Limit)
	if err != nil {
		log.Printf("%s Failed GetScoringServicesForModel: %v", pz, err)
		return err
	}
  
	out.Services = val0 
  
	return nil
}

func (this *Impl) DeleteScoringService(r *http.Request, in *DeleteScoringServiceIn, out *DeleteScoringServiceOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "DeleteScoringService")

	err := this.Service.DeleteScoringService(pz, in.ServiceId)
	if err != nil {
		log.Printf("%s Failed DeleteScoringService: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) AddEngine(r *http.Request, in *AddEngineIn, out *AddEngineOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "AddEngine")

	val0, err := this.Service.AddEngine(pz, in.EngineName, in.EnginePath)
	if err != nil {
		log.Printf("%s Failed AddEngine: %v", pz, err)
		return err
	}
  
	out.EngineId = val0 
  
	return nil
}

func (this *Impl) GetEngine(r *http.Request, in *GetEngineIn, out *GetEngineOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetEngine")

	val0, err := this.Service.GetEngine(pz, in.EngineId)
	if err != nil {
		log.Printf("%s Failed GetEngine: %v", pz, err)
		return err
	}
  
	out.Engine = val0 
  
	return nil
}

func (this *Impl) GetEngines(r *http.Request, in *GetEnginesIn, out *GetEnginesOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetEngines")

	val0, err := this.Service.GetEngines(pz)
	if err != nil {
		log.Printf("%s Failed GetEngines: %v", pz, err)
		return err
	}
  
	out.Engines = val0 
  
	return nil
}

func (this *Impl) DeleteEngine(r *http.Request, in *DeleteEngineIn, out *DeleteEngineOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "DeleteEngine")

	err := this.Service.DeleteEngine(pz, in.EngineId)
	if err != nil {
		log.Printf("%s Failed DeleteEngine: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) GetSupportedEntityTypes(r *http.Request, in *GetSupportedEntityTypesIn, out *GetSupportedEntityTypesOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetSupportedEntityTypes")

	val0, err := this.Service.GetSupportedEntityTypes(pz)
	if err != nil {
		log.Printf("%s Failed GetSupportedEntityTypes: %v", pz, err)
		return err
	}
  
	out.EntityTypes = val0 
  
	return nil
}

func (this *Impl) GetSupportedPermissions(r *http.Request, in *GetSupportedPermissionsIn, out *GetSupportedPermissionsOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetSupportedPermissions")

	val0, err := this.Service.GetSupportedPermissions(pz)
	if err != nil {
		log.Printf("%s Failed GetSupportedPermissions: %v", pz, err)
		return err
	}
  
	out.Permissions = val0 
  
	return nil
}

func (this *Impl) GetSupportedClusterTypes(r *http.Request, in *GetSupportedClusterTypesIn, out *GetSupportedClusterTypesOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetSupportedClusterTypes")

	val0, err := this.Service.GetSupportedClusterTypes(pz)
	if err != nil {
		log.Printf("%s Failed GetSupportedClusterTypes: %v", pz, err)
		return err
	}
  
	out.ClusterTypes = val0 
  
	return nil
}

func (this *Impl) GetPermissionsForRole(r *http.Request, in *GetPermissionsForRoleIn, out *GetPermissionsForRoleOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetPermissionsForRole")

	val0, err := this.Service.GetPermissionsForRole(pz, in.RoleId)
	if err != nil {
		log.Printf("%s Failed GetPermissionsForRole: %v", pz, err)
		return err
	}
  
	out.Permissions = val0 
  
	return nil
}

func (this *Impl) GetPermissionsForIdentity(r *http.Request, in *GetPermissionsForIdentityIn, out *GetPermissionsForIdentityOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetPermissionsForIdentity")

	val0, err := this.Service.GetPermissionsForIdentity(pz, in.IdentityId)
	if err != nil {
		log.Printf("%s Failed GetPermissionsForIdentity: %v", pz, err)
		return err
	}
  
	out.Permissions = val0 
  
	return nil
}

func (this *Impl) CreateRole(r *http.Request, in *CreateRoleIn, out *CreateRoleOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "CreateRole")

	val0, err := this.Service.CreateRole(pz, in.Name, in.Description)
	if err != nil {
		log.Printf("%s Failed CreateRole: %v", pz, err)
		return err
	}
  
	out.RoleId = val0 
  
	return nil
}

func (this *Impl) GetRoles(r *http.Request, in *GetRolesIn, out *GetRolesOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetRoles")

	val0, err := this.Service.GetRoles(pz, in.Offset, in.Limit)
	if err != nil {
		log.Printf("%s Failed GetRoles: %v", pz, err)
		return err
	}
  
	out.Roles = val0 
  
	return nil
}

func (this *Impl) GetRolesForIdentity(r *http.Request, in *GetRolesForIdentityIn, out *GetRolesForIdentityOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetRolesForIdentity")

	val0, err := this.Service.GetRolesForIdentity(pz, in.IdentityId)
	if err != nil {
		log.Printf("%s Failed GetRolesForIdentity: %v", pz, err)
		return err
	}
  
	out.Roles = val0 
  
	return nil
}

func (this *Impl) GetRole(r *http.Request, in *GetRoleIn, out *GetRoleOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetRole")

	val0, err := this.Service.GetRole(pz, in.RoleId)
	if err != nil {
		log.Printf("%s Failed GetRole: %v", pz, err)
		return err
	}
  
	out.Role = val0 
  
	return nil
}

func (this *Impl) GetRoleByName(r *http.Request, in *GetRoleByNameIn, out *GetRoleByNameOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetRoleByName")

	val0, err := this.Service.GetRoleByName(pz, in.Name)
	if err != nil {
		log.Printf("%s Failed GetRoleByName: %v", pz, err)
		return err
	}
  
	out.Role = val0 
  
	return nil
}

func (this *Impl) UpdateRole(r *http.Request, in *UpdateRoleIn, out *UpdateRoleOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "UpdateRole")

	err := this.Service.UpdateRole(pz, in.RoleId, in.Name, in.Description)
	if err != nil {
		log.Printf("%s Failed UpdateRole: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) LinkRoleAndPermissions(r *http.Request, in *LinkRoleAndPermissionsIn, out *LinkRoleAndPermissionsOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "LinkRoleAndPermissions")

	err := this.Service.LinkRoleAndPermissions(pz, in.RoleId, in.PermissionIds)
	if err != nil {
		log.Printf("%s Failed LinkRoleAndPermissions: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) DeleteRole(r *http.Request, in *DeleteRoleIn, out *DeleteRoleOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "DeleteRole")

	err := this.Service.DeleteRole(pz, in.RoleId)
	if err != nil {
		log.Printf("%s Failed DeleteRole: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) CreateWorkgroup(r *http.Request, in *CreateWorkgroupIn, out *CreateWorkgroupOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "CreateWorkgroup")

	val0, err := this.Service.CreateWorkgroup(pz, in.Name, in.Description)
	if err != nil {
		log.Printf("%s Failed CreateWorkgroup: %v", pz, err)
		return err
	}
  
	out.WorkgroupId = val0 
  
	return nil
}

func (this *Impl) GetWorkgroups(r *http.Request, in *GetWorkgroupsIn, out *GetWorkgroupsOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetWorkgroups")

	val0, err := this.Service.GetWorkgroups(pz, in.Offset, in.Limit)
	if err != nil {
		log.Printf("%s Failed GetWorkgroups: %v", pz, err)
		return err
	}
  
	out.Workgroups = val0 
  
	return nil
}

func (this *Impl) GetWorkgroupsForIdentity(r *http.Request, in *GetWorkgroupsForIdentityIn, out *GetWorkgroupsForIdentityOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetWorkgroupsForIdentity")

	val0, err := this.Service.GetWorkgroupsForIdentity(pz, in.IdentityId)
	if err != nil {
		log.Printf("%s Failed GetWorkgroupsForIdentity: %v", pz, err)
		return err
	}
  
	out.Workgroups = val0 
  
	return nil
}

func (this *Impl) GetWorkgroup(r *http.Request, in *GetWorkgroupIn, out *GetWorkgroupOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetWorkgroup")

	val0, err := this.Service.GetWorkgroup(pz, in.WorkgroupId)
	if err != nil {
		log.Printf("%s Failed GetWorkgroup: %v", pz, err)
		return err
	}
  
	out.Workgroup = val0 
  
	return nil
}

func (this *Impl) GetWorkgroupByName(r *http.Request, in *GetWorkgroupByNameIn, out *GetWorkgroupByNameOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetWorkgroupByName")

	val0, err := this.Service.GetWorkgroupByName(pz, in.Name)
	if err != nil {
		log.Printf("%s Failed GetWorkgroupByName: %v", pz, err)
		return err
	}
  
	out.Workgroup = val0 
  
	return nil
}

func (this *Impl) UpdateWorkgroup(r *http.Request, in *UpdateWorkgroupIn, out *UpdateWorkgroupOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "UpdateWorkgroup")

	err := this.Service.UpdateWorkgroup(pz, in.WorkgroupId, in.Name, in.Description)
	if err != nil {
		log.Printf("%s Failed UpdateWorkgroup: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) DeleteWorkgroup(r *http.Request, in *DeleteWorkgroupIn, out *DeleteWorkgroupOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "DeleteWorkgroup")

	err := this.Service.DeleteWorkgroup(pz, in.WorkgroupId)
	if err != nil {
		log.Printf("%s Failed DeleteWorkgroup: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) CreateIdentity(r *http.Request, in *CreateIdentityIn, out *CreateIdentityOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "CreateIdentity")

	val0, err := this.Service.CreateIdentity(pz, in.Name, in.Password)
	if err != nil {
		log.Printf("%s Failed CreateIdentity: %v", pz, err)
		return err
	}
  
	out.IdentityId = val0 
  
	return nil
}

func (this *Impl) GetIdentities(r *http.Request, in *GetIdentitiesIn, out *GetIdentitiesOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetIdentities")

	val0, err := this.Service.GetIdentities(pz, in.Offset, in.Limit)
	if err != nil {
		log.Printf("%s Failed GetIdentities: %v", pz, err)
		return err
	}
  
	out.Identities = val0 
  
	return nil
}

func (this *Impl) GetIdentitiesForWorkgroup(r *http.Request, in *GetIdentitiesForWorkgroupIn, out *GetIdentitiesForWorkgroupOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetIdentitiesForWorkgroup")

	val0, err := this.Service.GetIdentitiesForWorkgroup(pz, in.WorkgroupId)
	if err != nil {
		log.Printf("%s Failed GetIdentitiesForWorkgroup: %v", pz, err)
		return err
	}
  
	out.Identities = val0 
  
	return nil
}

func (this *Impl) GetIdentitiesForRole(r *http.Request, in *GetIdentitiesForRoleIn, out *GetIdentitiesForRoleOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetIdentitiesForRole")

	val0, err := this.Service.GetIdentitiesForRole(pz, in.RoleId)
	if err != nil {
		log.Printf("%s Failed GetIdentitiesForRole: %v", pz, err)
		return err
	}
  
	out.Identities = val0 
  
	return nil
}

func (this *Impl) GetIdentity(r *http.Request, in *GetIdentityIn, out *GetIdentityOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetIdentity")

	val0, err := this.Service.GetIdentity(pz, in.IdentityId)
	if err != nil {
		log.Printf("%s Failed GetIdentity: %v", pz, err)
		return err
	}
  
	out.Identity = val0 
  
	return nil
}

func (this *Impl) GetIdentityByName(r *http.Request, in *GetIdentityByNameIn, out *GetIdentityByNameOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetIdentityByName")

	val0, err := this.Service.GetIdentityByName(pz, in.Name)
	if err != nil {
		log.Printf("%s Failed GetIdentityByName: %v", pz, err)
		return err
	}
  
	out.Identity = val0 
  
	return nil
}

func (this *Impl) LinkIdentityAndWorkgroup(r *http.Request, in *LinkIdentityAndWorkgroupIn, out *LinkIdentityAndWorkgroupOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "LinkIdentityAndWorkgroup")

	err := this.Service.LinkIdentityAndWorkgroup(pz, in.IdentityId, in.WorkgroupId)
	if err != nil {
		log.Printf("%s Failed LinkIdentityAndWorkgroup: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) UnlinkIdentityAndWorkgroup(r *http.Request, in *UnlinkIdentityAndWorkgroupIn, out *UnlinkIdentityAndWorkgroupOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "UnlinkIdentityAndWorkgroup")

	err := this.Service.UnlinkIdentityAndWorkgroup(pz, in.IdentityId, in.WorkgroupId)
	if err != nil {
		log.Printf("%s Failed UnlinkIdentityAndWorkgroup: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) LinkIdentityAndRole(r *http.Request, in *LinkIdentityAndRoleIn, out *LinkIdentityAndRoleOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "LinkIdentityAndRole")

	err := this.Service.LinkIdentityAndRole(pz, in.IdentityId, in.RoleId)
	if err != nil {
		log.Printf("%s Failed LinkIdentityAndRole: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) UnlinkIdentityAndRole(r *http.Request, in *UnlinkIdentityAndRoleIn, out *UnlinkIdentityAndRoleOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "UnlinkIdentityAndRole")

	err := this.Service.UnlinkIdentityAndRole(pz, in.IdentityId, in.RoleId)
	if err != nil {
		log.Printf("%s Failed UnlinkIdentityAndRole: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) UpdateIdentity(r *http.Request, in *UpdateIdentityIn, out *UpdateIdentityOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "UpdateIdentity")

	err := this.Service.UpdateIdentity(pz, in.IdentityId, in.Password)
	if err != nil {
		log.Printf("%s Failed UpdateIdentity: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) DeactivateIdentity(r *http.Request, in *DeactivateIdentityIn, out *DeactivateIdentityOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "DeactivateIdentity")

	err := this.Service.DeactivateIdentity(pz, in.IdentityId)
	if err != nil {
		log.Printf("%s Failed DeactivateIdentity: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) ShareEntity(r *http.Request, in *ShareEntityIn, out *ShareEntityOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "ShareEntity")

	err := this.Service.ShareEntity(pz, in.Kind, in.WorkgroupId, in.EntityTypeId, in.EntityId)
	if err != nil {
		log.Printf("%s Failed ShareEntity: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) GetEntityPrivileges(r *http.Request, in *GetEntityPrivilegesIn, out *GetEntityPrivilegesOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetEntityPrivileges")

	val0, err := this.Service.GetEntityPrivileges(pz, in.EntityTypeId, in.EntityId)
	if err != nil {
		log.Printf("%s Failed GetEntityPrivileges: %v", pz, err)
		return err
	}
  
	out.Privileges = val0 
  
	return nil
}

func (this *Impl) UnshareEntity(r *http.Request, in *UnshareEntityIn, out *UnshareEntityOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "UnshareEntity")

	err := this.Service.UnshareEntity(pz, in.Kind, in.WorkgroupId, in.EntityTypeId, in.EntityId)
	if err != nil {
		log.Printf("%s Failed UnshareEntity: %v", pz, err)
		return err
	}
  
	return nil
}

func (this *Impl) GetEntityHistory(r *http.Request, in *GetEntityHistoryIn, out *GetEntityHistoryOut) error {
	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}
	log.Println(pz, "GetEntityHistory")

	val0, err := this.Service.GetEntityHistory(pz, in.EntityTypeId, in.EntityId, in.Offset, in.Limit)
	if err != nil {
		log.Printf("%s Failed GetEntityHistory: %v", pz, err)
		return err
	}
  
	out.History = val0 
  
	return nil
}



