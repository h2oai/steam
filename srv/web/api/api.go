package api

// --- Type Definitions ---

type Cluster struct {
	Id        int64
	Name      string
	TypeId    int64
	DetailId  int64
	Address   string
	State     string
	CreatedAt int64
}

type YarnCluster struct {
	Id            int64
	EngineId      int64
	Size          int
	ApplicationId string
	Memory        string
	Username      string
}

type ClusterStatus struct {
	Version              string
	Status               string
	MaxMemory            string
	TotalCpuCount        int
	TotalAllowedCpuCount int
}

type Job struct {
	Name        string
	ClusterName string
	Description string
	Progress    string
	StartedAt   int64
	CompletedAt int64
}

type Project struct {
	Id          int64
	Name        string
	Description string
	CreatedAt   int64
}

type Datasource struct {
	Id            int64
	ProjectId     int64
	Name          string
	Description   string
	Kind          string
	Configuration string
	CreatedAt     int64
}

type Dataset struct {
	Id                 int64
	DatasourceId       int64
	Name               string
	Description        string
	FrameName          string
	ResponseColumnName string
	Properties         string
	CreatedAt          int64
}

type Model struct {
	Id                  int64
	TrainingDatasetId   int64
	ValidationDatasetId int64
	Name                string
	ClusterName         string
	Algorithm           string
	DatasetName         string
	ResponseColumnName  string
	LogicalName         string
	Location            string
	MaxRuntime          int
	Metrics             string
	CreatedAt           int64
}

type ScoringService struct {
	Id        int64
	ModelId   int64
	Address   string
	Port      int
	ProcessId int
	State     string
	CreatedAt int64
}

type Engine struct {
	Id        int64
	Name      string
	Location  string
	CreatedAt int64
}

type EntityType struct {
	Id   int64
	Name string
}

type ClusterType struct {
	Id   int64
	Name string
}

type EntityHistory struct {
	IdentityId  int64
	Action      string
	Description string
	CreatedAt   int64
}

type Permission struct {
	Id          int64
	Code        string
	Description string
}

type Privilege struct {
	Kind        string
	WorkgroupId int64
}

type EntityPrivilege struct {
	Kind                 string
	WorkgroupId          int64
	WorkgroupName        string
	WorkgroupDescription string
}

type Role struct {
	Id          int64
	Name        string
	Description string
	Created     int64
}

type Identity struct {
	Id        int64
	Name      string
	IsActive  bool
	LastLogin int64
	Created   int64
}

type Workgroup struct {
	Id          int64
	Name        string
	Description string
	Created     int64
}

// --- API Facade ---

type Service struct {
	Ping                       Ping
	RegisterCluster            RegisterCluster
	UnregisterCluster          UnregisterCluster
	StartYarnCluster           StartYarnCluster
	StopYarnCluster            StopYarnCluster
	GetCluster                 GetCluster
	GetYarnCluster             GetYarnCluster
	GetClusters                GetClusters
	GetClusterStatus           GetClusterStatus
	DeleteCluster              DeleteCluster
	GetJob                     GetJob
	GetJobs                    GetJobs
	CreateProject              CreateProject
	GetProjects                GetProjects
	GetProject                 GetProject
	DeleteProject              DeleteProject
	CreateDatasource           CreateDatasource
	GetDatasources             GetDatasources
	GetDatasource              GetDatasource
	UpdateDatasource           UpdateDatasource
	DeleteDatasource           DeleteDatasource
	CreateDataset              CreateDataset
	GetDatasets                GetDatasets
	GetDataset                 GetDataset
	UpdateDataset              UpdateDataset
	SplitDataset               SplitDataset
	DeleteDataset              DeleteDataset
	BuildModel                 BuildModel
	BuildAutoModel             BuildAutoModel
	GetModel                   GetModel
	GetModels                  GetModels
	GetClusterModels           GetClusterModels
	ImportModelFromCluster     ImportModelFromCluster
	DeleteModel                DeleteModel
	StartScoringService        StartScoringService
	StopScoringService         StopScoringService
	GetScoringService          GetScoringService
	GetScoringServices         GetScoringServices
	GetScoringServicesForModel GetScoringServicesForModel
	DeleteScoringService       DeleteScoringService
	AddEngine                  AddEngine
	GetEngine                  GetEngine
	GetEngines                 GetEngines
	DeleteEngine               DeleteEngine
	GetSupportedEntityTypes    GetSupportedEntityTypes
	GetSupportedPermissions    GetSupportedPermissions
	GetSupportedClusterTypes   GetSupportedClusterTypes
	GetPermissionsForRole      GetPermissionsForRole
	GetPermissionsForIdentity  GetPermissionsForIdentity
	CreateRole                 CreateRole
	GetRoles                   GetRoles
	GetRolesForIdentity        GetRolesForIdentity
	GetRole                    GetRole
	GetRoleByName              GetRoleByName
	UpdateRole                 UpdateRole
	LinkRoleAndPermissions     LinkRoleAndPermissions
	DeleteRole                 DeleteRole
	CreateWorkgroup            CreateWorkgroup
	GetWorkgroups              GetWorkgroups
	GetWorkgroupsForIdentity   GetWorkgroupsForIdentity
	GetWorkgroup               GetWorkgroup
	GetWorkgroupByName         GetWorkgroupByName
	UpdateWorkgroup            UpdateWorkgroup
	DeleteWorkgroup            DeleteWorkgroup
	CreateIdentity             CreateIdentity
	GetIdentities              GetIdentities
	GetIdentitiesForWorkgroup  GetIdentitiesForWorkgroup
	GetIdentitiesForRole       GetIdentitiesForRole
	GetIdentity                GetIdentity
	GetIdentityByName          GetIdentityByName
	LinkIdentityAndWorkgroup   LinkIdentityAndWorkgroup
	UnlinkIdentityAndWorkgroup UnlinkIdentityAndWorkgroup
	LinkIdentityAndRole        LinkIdentityAndRole
	UnlinkIdentityAndRole      UnlinkIdentityAndRole
	UpdateIdentity             UpdateIdentity
	DeactivateIdentity         DeactivateIdentity
	ShareEntity                ShareEntity
	GetEntityPrivileges        GetEntityPrivileges
	UnshareEntity              UnshareEntity
	GetEntityHistory           GetEntityHistory
}

// --- API Method Definitions ---

// Note: Define each method as a struct, with fields representing parameters and returns.
//       Place a dummy field "_" to separate inputs and outputs.

type Ping struct {
	Input  bool
	_      int
	Output bool
}
type RegisterCluster struct {
	Address   string
	_         int
	ClusterId int64
}
type UnregisterCluster struct {
	ClusterId int64
}
type StartYarnCluster struct {
	ClusterName string
	EngineId    int64
	Size        int
	Memory      string
	Username    string
	_           int
	ClusterId   int64
}
type StopYarnCluster struct {
	ClusterId int64
}
type GetCluster struct {
	ClusterId int64
	_         int
	Cluster   Cluster
}
type GetYarnCluster struct {
	ClusterId int64
	_         int
	Cluster   YarnCluster
}
type GetClusters struct {
	Offset   int64
	Limit    int64
	_        int
	Clusters []Cluster
}
type GetClusterStatus struct {
	ClusterId     int64
	_             int
	ClusterStatus ClusterStatus
}
type DeleteCluster struct {
	ClusterId int64
}
type GetJob struct {
	ClusterId int64
	JobName   string
	_         int
	Job       Job
}
type GetJobs struct {
	ClusterId int64
	_         int
	Jobs      []Job
}
type CreateProject struct {
	Name        string
	Description string
	_           int
	ProjectId   int64
}
type GetProjects struct {
	Offset   int64
	Limit    int64
	_        int
	Projects []Project
}
type GetProject struct {
	ProjectId int64
	_         int
	Project   Project
}
type DeleteProject struct {
	ProjectId int64
}
type CreateDatasource struct {
	ProjectId    int64
	Name         string
	Description  string
	Path         string
	_            int
	DatasourceId int64
}
type GetDatasources struct {
	ProjectId   int64
	Offset      int64
	Limit       int64
	_           int
	Datasources []Datasource
}
type GetDatasource struct {
	DatasourceId int64
	_            int
	Datasource   Datasource
}
type UpdateDatasource struct {
	DatasourceId int64
	Name         string
	Description  string
	Path         string
}
type DeleteDatasource struct {
	DatasourceId int64
}
type CreateDataset struct {
	ClusterId          int64
	DatasourceId       int64
	Name               string
	Description        string
	ResponseColumnName string
	_                  int
	DatasetId          int64
}
type GetDatasets struct {
	DatasourceId int64
	Offset       int64
	Limit        int64
	_            int
	Datasets     []Dataset
}
type GetDataset struct {
	DatasetId int64
	_         int
	Dataset   Dataset
}
type UpdateDataset struct {
	DatasetId          int64
	Name               string
	Description        string
	ResponseColumnName string
}
type SplitDataset struct {
	DatasetId  int64
	Ratio1     int
	Ratio2     int
	_          int
	DatasetIds []int64
}
type DeleteDataset struct {
	DatasetId int64
}
type BuildModel struct {
	ClusterId int64
	DatasetId int64
	Algorithm string
	_         int
	ModelId   int64
}
type BuildAutoModel struct {
	ClusterId  int64
	Dataset    string
	TargetName string
	MaxRunTime int
	_          int
	Model      Model
}
type GetModel struct {
	ModelId int64
	_       int
	Model   Model
}
type GetModels struct {
	ProjectId int64
	Offset    int64
	Limit     int64
	_         int
	Models    []Model
}
type GetClusterModels struct {
	ClusterId int64
	_         int
	Models    []Model
}
type ImportModelFromCluster struct {
	ClusterId int64
	ProjectId int64
	ModelName string
	_         int
	Model     Model
}
type DeleteModel struct {
	ModelId int64
}
type StartScoringService struct {
	ModelId int64
	Port    int
	_       int
	Service ScoringService
}
type StopScoringService struct {
	ServiceId int64
}
type GetScoringService struct {
	ServiceId int64
	_         int
	Service   ScoringService
}
type GetScoringServices struct {
	Offset   int64
	Limit    int64
	_        int
	Services []ScoringService
}
type GetScoringServicesForModel struct {
	ModelId  int64
	Offset   int64
	Limit    int64
	_        int
	Services []ScoringService
}
type DeleteScoringService struct {
	ServiceId int64
}
type AddEngine struct {
	EngineName string
	EnginePath string
	_          int
	EngineId   int64
}
type GetEngine struct {
	EngineId int64
	_        int
	Engine   Engine
}
type GetEngines struct {
	_       int
	Engines []Engine
}
type DeleteEngine struct {
	EngineId int64
}
type GetSupportedEntityTypes struct {
	_           int
	EntityTypes []EntityType
}
type GetSupportedPermissions struct {
	_           int
	Permissions []Permission
}
type GetSupportedClusterTypes struct {
	_            int
	ClusterTypes []ClusterType
}
type GetPermissionsForRole struct {
	RoleId      int64
	_           int
	Permissions []Permission
}
type GetPermissionsForIdentity struct {
	IdentityId  int64
	_           int
	Permissions []Permission
}
type CreateRole struct {
	Name        string
	Description string
	_           int
	RoleId      int64
}
type GetRoles struct {
	Offset int64
	Limit  int64
	_      int
	Roles  []Role
}
type GetRolesForIdentity struct {
	IdentityId int64
	_          int
	Roles      []Role
}
type GetRole struct {
	RoleId int64
	_      int
	Role   Role
}
type GetRoleByName struct {
	Name string
	_    int
	Role Role
}
type UpdateRole struct {
	RoleId      int64
	Name        string
	Description string
}
type LinkRoleAndPermissions struct {
	RoleId        int64
	PermissionIds []int64
}
type DeleteRole struct {
	RoleId int64
}
type CreateWorkgroup struct {
	Name        string
	Description string
	_           int
	WorkgroupId int64
}
type GetWorkgroups struct {
	Offset     int64
	Limit      int64
	_          int
	Workgroups []Workgroup
}
type GetWorkgroupsForIdentity struct {
	IdentityId int64
	_          int
	Workgroups []Workgroup
}
type GetWorkgroup struct {
	WorkgroupId int64
	_           int
	Workgroup   Workgroup
}
type GetWorkgroupByName struct {
	Name      string
	_         int
	Workgroup Workgroup
}
type UpdateWorkgroup struct {
	WorkgroupId int64
	Name        string
	Description string
}
type DeleteWorkgroup struct {
	WorkgroupId int64
}
type CreateIdentity struct {
	Name       string
	Password   string
	_          int
	IdentityId int64
}
type GetIdentities struct {
	Offset     int64
	Limit      int64
	_          int
	Identities []Identity
}
type GetIdentitiesForWorkgroup struct {
	WorkgroupId int64
	_           int
	Identities  []Identity
}
type GetIdentitiesForRole struct {
	RoleId     int64
	_          int
	Identities []Identity
}
type GetIdentity struct {
	IdentityId int64
	_          int
	Identity   Identity
}
type GetIdentityByName struct {
	Name     string
	_        int
	Identity Identity
}
type LinkIdentityAndWorkgroup struct {
	IdentityId  int64
	WorkgroupId int64
}
type UnlinkIdentityAndWorkgroup struct {
	IdentityId  int64
	WorkgroupId int64
}
type LinkIdentityAndRole struct {
	IdentityId int64
	RoleId     int64
}
type UnlinkIdentityAndRole struct {
	IdentityId int64
	RoleId     int64
}
type UpdateIdentity struct {
	IdentityId int64
	Password   string
}
type DeactivateIdentity struct {
	IdentityId int64
}
type ShareEntity struct {
	Kind         string
	WorkgroupId  int64
	EntityTypeId int64
	EntityId     int64
}
type GetEntityPrivileges struct {
	EntityTypeId int64
	EntityId     int64
	_            int
	Privileges   []EntityPrivilege
}
type UnshareEntity struct {
	Kind         string
	WorkgroupId  int64
	EntityTypeId int64
	EntityId     int64
}
type GetEntityHistory struct {
	EntityTypeId int64
	EntityId     int64
	Offset       int64
	Limit        int64
	_            int
	History      []EntityHistory
}
