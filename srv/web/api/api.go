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

package api

// --- Type Definitions ---

type Config struct {
	KerberosEnabled     bool
	ClusterProxyAddress string
}

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
	Id            int64
	Name          string
	Description   string
	ModelCategory string
	CreatedAt     int64
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
	JSONProperties     string
	CreatedAt          int64
}

type Model struct {
	Id                  int64
	TrainingDatasetId   int64
	ValidationDatasetId int64
	Name                string
	ClusterName         string
	ModelKey            string
	Algorithm           string
	ModelCategory       string
	DatasetName         string
	ResponseColumnName  string
	LogicalName         string
	Location            string
	ModelObjectType     string
	MaxRuntime          int
	JSONMetrics         string
	CreatedAt           int64
	LabelId             int64
	LabelName           string
}

type BinomialModel struct {
	Id                  int64
	TrainingDatasetId   int64
	ValidationDatasetId int64
	Name                string
	ClusterName         string
	ModelKey            string
	Algorithm           string
	ModelCategory       string
	DatasetName         string
	ResponseColumnName  string
	LogicalName         string
	Location            string
	ModelObjectType     string
	MaxRuntime          int
	JSONMetrics         string
	CreatedAt           int64
	LabelId             int64
	LabelName           string
	Mse                 float64
	RSquared            float64
	Logloss             float64
	Auc                 float64
	Gini                float64
}

type MultinomialModel struct {
	Id                  int64
	TrainingDatasetId   int64
	ValidationDatasetId int64
	Name                string
	ClusterName         string
	ModelKey            string
	Algorithm           string
	ModelCategory       string
	DatasetName         string
	ResponseColumnName  string
	LogicalName         string
	Location            string
	ModelObjectType     string
	MaxRuntime          int
	JSONMetrics         string
	CreatedAt           int64
	LabelId             int64
	LabelName           string
	Mse                 float64
	RSquared            float64
	Logloss             float64
}

type RegressionModel struct {
	Id                   int64
	TrainingDatasetId    int64
	ValidationDatasetId  int64
	Name                 string
	ClusterName          string
	ModelKey             string
	Algorithm            string
	ModelCategory        string
	DatasetName          string
	ResponseColumnName   string
	LogicalName          string
	Location             string
	ModelObjectType      string
	MaxRuntime           int
	JSONMetrics          string
	CreatedAt            int64
	LabelId              int64
	LabelName            string
	Mse                  float64
	RSquared             float64
	MeanResidualDeviance float64
}

type Label struct {
	Id          int64
	ProjectId   int64
	ModelId     int64
	Name        string
	Description string
	CreatedAt   int64
}

type ScoringService struct {
	Id        int64
	ModelId   int64
	Name      string
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

type UserRole struct {
	Kind         string
	IdentityId   int64
	IdentityName string
	RoleId       int64
	RoleName     string
}

type Workgroup struct {
	Id          int64
	Name        string
	Description string
	Created     int64
}

// --- API Facade ---

type Service struct {
	PingServer                    PingServer                    `help:"Ping the Steam server"`
	GetConfig                     GetConfig                     `help:Get Steam start up configurations`
	RegisterCluster               RegisterCluster               `help:"Connect to a cluster"`
	UnregisterCluster             UnregisterCluster             `help:"Disconnect from a cluster"`
	StartClusterOnYarn            StartClusterOnYarn            `help:"Start a cluster using Yarn"`
	StopClusterOnYarn             StopClusterOnYarn             `help:"Stop a cluster using Yarn"`
	GetCluster                    GetCluster                    `help:"Get cluster details"`
	GetClusterOnYarn              GetClusterOnYarn              `help:"Get cluster details (Yarn only)"`
	GetClusters                   GetClusters                   `help:"List clusters"`
	GetClusterStatus              GetClusterStatus              `help:"Get cluster status"`
	DeleteCluster                 DeleteCluster                 `help:"Delete a cluster"`
	GetJob                        GetJob                        `help:"Get job details"`
	GetJobs                       GetJobs                       `help:"List jobs"`
	CreateProject                 CreateProject                 `help:"Create a project"`
	GetProjects                   GetProjects                   `help:"List projects"`
	GetProject                    GetProject                    `help:"Get project details"`
	DeleteProject                 DeleteProject                 `help:"Delete a project"`
	CreateDatasource              CreateDatasource              `help:"Create a datasource"`
	GetDatasources                GetDatasources                `help:"List datasources"`
	GetDatasource                 GetDatasource                 `help:"Get datasource details"`
	UpdateDatasource              UpdateDatasource              `help:"Update a datasource"`
	DeleteDatasource              DeleteDatasource              `help:"Delete a datasource"`
	CreateDataset                 CreateDataset                 `help:"Create a dataset"`
	GetDatasets                   GetDatasets                   `help:"List datasets"`
	GetDataset                    GetDataset                    `help:"Get dataset details"`
	GetDatasetsFromCluster        GetDatasetsFromCluster        `help:"Get a list of datasets on a cluster"`
	UpdateDataset                 UpdateDataset                 `help:"Update a dataset"`
	SplitDataset                  SplitDataset                  `help:"Split a dataset"`
	DeleteDataset                 DeleteDataset                 `help:"Delete a dataset"`
	BuildModel                    BuildModel                    `help:"Build a model"`
	BuildModelAuto                BuildModelAuto                `help:"Build an AutoML model"`
	GetModel                      GetModel                      `help:"Get model details"`
	GetModels                     GetModels                     `help:"List models"`
	GetModelsFromCluster          GetModelsFromCluster          `help:"List models from a cluster"`
	FindModelsCount               FindModelsCount               `help:"Get a count models in a project"`
	GetAllBinomialSortCriteria    GetAllBinomialSortCriteria    `help:"List sort criteria for a binomial models"`
	FindModelsBinomial            FindModelsBinomial            `help:"List binomial models"`
	GetModelBinomial              GetModelBinomial              `help:"View a binomial model"`
	GetAllMultinomialSortCriteria GetAllMultinomialSortCriteria `help:"List sort criteria for a multinomial models"`
	FindModelsMultinomial         FindModelsMultinomial         `help:"List multinomial models"`
	GetModelMultinomial           GetModelMultinomial           `help:"View a binomial model"`
	GetAllRegressionSortCriteria  GetAllRegressionSortCriteria  `help:"List sort criteria for a regression models"`
	FindModelsRegression          FindModelsRegression          `help:"List regression models"`
	GetModelRegression            GetModelRegression            `help:"View a binomial model"`
	ImportModelFromCluster        ImportModelFromCluster        `help:"Import models from a cluster"`
	CheckMojo                     CheckMojo                     `help:"Check if a model category can generate MOJOs"`
	ImportModelPojo               ImportModelPojo               `help:"Import a model's POJO from a cluster"`
	ImportModelMojo               ImportModelMojo               `help:"Import a model's MOJO from a cluster"`
	DeleteModel                   DeleteModel                   `help:"Delete a model"`
	CreateLabel                   CreateLabel                   `help:"Create a label"`
	UpdateLabel                   UpdateLabel                   `help:"Update a label"`
	DeleteLabel                   DeleteLabel                   `help:"Delete a label"`
	LinkLabelWithModel            LinkLabelWithModel            `help:"Label a model"`
	UnlinkLabelFromModel          UnlinkLabelFromModel          `help:"Remove a label from a model"`
	GetLabelsForProject           GetLabelsForProject           `help:"List labels for a project, with corresponding models, if any"`
	StartService                  StartService                  `help:"Start a service"`
	StopService                   StopService                   `help:"Stop a service"`
	GetService                    GetService                    `help:"Get service details"`
	GetServices                   GetServices                   `help:"List all services"`
	GetServicesForProject         GetServicesForProject         `help:"List services for a project"`
	GetServicesForModel           GetServicesForModel           `help:"List services for a model"`
	DeleteService                 DeleteService                 `help:"Delete a service"`
	GetEngine                     GetEngine                     `help:"Get engine details"`
	GetEngines                    GetEngines                    `help:"List engines"`
	DeleteEngine                  DeleteEngine                  `help:"Delete an engine"`
	GetAllEntityTypes             GetAllEntityTypes             `help:"List all entity types"`
	GetAllPermissions             GetAllPermissions             `help:"List all permissions"`
	GetAllClusterTypes            GetAllClusterTypes            `help:"List all cluster types"`
	GetPermissionsForRole         GetPermissionsForRole         `help:"List permissions for a role"`
	GetPermissionsForIdentity     GetPermissionsForIdentity     `help:"List permissions for an identity"`
	CreateRole                    CreateRole                    `help:"Create a role"`
	GetRoles                      GetRoles                      `help:"List roles"`
	GetRolesForIdentity           GetRolesForIdentity           `help:"List roles for an identity"`
	GetRole                       GetRole                       `help:"Get role details"`
	GetRoleByName                 GetRoleByName                 `help:"Get role details by name"`
	UpdateRole                    UpdateRole                    `help:"Update a role"`
	LinkRoleWithPermissions       LinkRoleWithPermissions       `help:"Link a role with permissions"`
	LinkRoleWithPermission        LinkRoleWithPermission        `help:"Link a role with a permission"`
	UnlinkRoleFromPermission      UnlinkRoleFromPermission      `help:"Unlink a role from a permission"`
	DeleteRole                    DeleteRole                    `help:"Delete a role"`
	CreateWorkgroup               CreateWorkgroup               `help:"Create a workgroup"`
	GetWorkgroups                 GetWorkgroups                 `help:"List workgroups"`
	GetWorkgroupsForIdentity      GetWorkgroupsForIdentity      `help:"List workgroups for an identity"`
	GetWorkgroup                  GetWorkgroup                  `help:"Get workgroup details"`
	GetWorkgroupByName            GetWorkgroupByName            `help:"Get workgroup details by name"`
	UpdateWorkgroup               UpdateWorkgroup               `help:"Update a workgroup"`
	DeleteWorkgroup               DeleteWorkgroup               `help:"Delete a workgroup"`
	CreateIdentity                CreateIdentity                `help:"Create an identity"`
	GetIdentities                 GetIdentities                 `help:"List identities"`
	GetIdentitiesForWorkgroup     GetIdentitiesForWorkgroup     `help:"List identities for a workgroup"`
	GetIdentitiesForRole          GetIdentitiesForRole          `help:"List identities for a role"`
	GetIdentitiesForEntity        GetIdentitiesForEntity        `help:"Get a list of identities and roles with access to an entity"`
	GetIdentity                   GetIdentity                   `help:"Get identity details"`
	GetIdentityByName             GetIdentityByName             `help:"Get identity details by name"`
	LinkIdentityWithWorkgroup     LinkIdentityWithWorkgroup     `help:"Link an identity with a workgroup"`
	UnlinkIdentityFromWorkgroup   UnlinkIdentityFromWorkgroup   `help:"Unlink an identity from a workgroup"`
	LinkIdentityWithRole          LinkIdentityWithRole          `help:"Link an identity with a role"`
	UnlinkIdentityFromRole        UnlinkIdentityFromRole        `help:"Unlink an identity from a role"`
	UpdateIdentity                UpdateIdentity                `help:"Update an identity"`
	ActivateIdentity              ActivateIdentity              `help:"Activate an identity"`
	DeactivateIdentity            DeactivateIdentity            `help:"Deactivate an identity"`
	ShareEntity                   ShareEntity                   `help:"Share an entity with a workgroup"`
	GetPrivileges                 GetPrivileges                 `help:"List privileges for an entity"`
	UnshareEntity                 UnshareEntity                 `help:"Unshare an entity"`
	GetHistory                    GetHistory                    `help:"List audit trail records for an entity"`
	CreatePackage                 CreatePackage                 `help:"Create a package for a project"`
	GetPackages                   GetPackages                   `help:"List packages for a project "`
	GetPackageDirectories         GetPackageDirectories         `help:"List directories in a project package"`
	GetPackageFiles               GetPackageFiles               `help:"List files in a project package"`
	DeletePackage                 DeletePackage                 `help:"Delete a project package"`
	DeletePackageDirectory        DeletePackageDirectory        `help:"Delete a directory in a project package"`
	DeletePackageFile             DeletePackageFile             `help:"Delete a file in a project package"`
	SetAttributesForPackage       SetAttributesForPackage       `help:"Set attributes on a project package"`
	GetAttributesForPackage       GetAttributesForPackage       `help:"List attributes for a project package"`
}

// --- API Method Definitions ---

// Note: Define each method as a struct, with fields representing parameters and returns.
//       Place a dummy field "_" to separate inputs and outputs.

type PingServer struct {
	Input  string `help:"Message to send"`
	_      int
	Output string `help:"Echoed message"`
}
type GetConfig struct {
	_      int
	Config Config `help:"An object containing Steam startup configurations"`
}
type RegisterCluster struct {
	Address   string
	_         int
	ClusterId int64
}
type UnregisterCluster struct {
	ClusterId int64
}
type StartClusterOnYarn struct {
	ClusterName string
	EngineId    int64
	Size        int
	Memory      string
	Secure      bool
	Keytab      string
	_           int
	ClusterId   int64
}
type StopClusterOnYarn struct {
	ClusterId int64
	Keytab    string
}
type GetCluster struct {
	ClusterId int64
	_         int
	Cluster   Cluster
}
type GetClusterOnYarn struct {
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
	Name          string
	Description   string
	ModelCategory string
	_             int
	ProjectId     int64
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
type GetDatasetsFromCluster struct {
	ClusterId int64
	_         int
	Dataset   []Dataset
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
type BuildModelAuto struct {
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
type GetModelsFromCluster struct {
	ClusterId int64
	FrameKey  string
	_         int
	Models    []Model
}
type FindModelsCount struct {
	ProjectId int64
	_         int
	Count     int64
}
type GetAllBinomialSortCriteria struct {
	_        int
	Criteria []string
}
type FindModelsBinomial struct {
	ProjectId int64
	NamePart  string
	SortBy    string
	Ascending bool
	Offset    int64
	Limit     int64
	_         int
	Models    []BinomialModel
}
type GetModelBinomial struct {
	ModelId int64
	_       int
	Model   BinomialModel
}
type GetAllMultinomialSortCriteria struct {
	_        int
	Criteria []string
}
type FindModelsMultinomial struct {
	ProjectId int64
	NamePart  string
	SortBy    string
	Ascending bool
	Offset    int64
	Limit     int64
	_         int
	Models    []MultinomialModel
}
type GetModelMultinomial struct {
	ModelId int64
	_       int
	Model   MultinomialModel
}
type GetAllRegressionSortCriteria struct {
	_        int
	Criteria []string
}
type FindModelsRegression struct {
	ProjectId int64
	NamePart  string
	SortBy    string
	Ascending bool
	Offset    int64
	Limit     int64
	_         int
	Models    []RegressionModel
}
type GetModelRegression struct {
	ModelId int64
	_       int
	Model   RegressionModel
}
type ImportModelFromCluster struct {
	ClusterId int64
	ProjectId int64
	ModelKey  string
	ModelName string
	_         int
	ModelId   int64
}
type CheckMojo struct {
	Algo    string
	_       int
	CanMojo bool
}
type ImportModelPojo struct {
	ModelId int64
	_       int
}
type ImportModelMojo struct {
	ModelId int64
	_       int
}
type DeleteModel struct {
	ModelId int64
}
type CreateLabel struct {
	ProjectId   int64
	Name        string
	Description string
	_           int
	LabelId     int64
}
type UpdateLabel struct {
	LabelId     int64
	Name        string
	Description string
}
type DeleteLabel struct {
	LabelId int64
}
type LinkLabelWithModel struct {
	LabelId int64
	ModelId int64
}
type UnlinkLabelFromModel struct {
	LabelId int64
	ModelId int64
}
type GetLabelsForProject struct {
	ProjectId int64
	_         int
	Labels    []Label
}
type StartService struct {
	ModelId     int64
	Name        string
	PackageName string
	_           int
	ServiceId   int64
}
type StopService struct {
	ServiceId int64
}
type GetService struct {
	ServiceId int64
	_         int
	Service   ScoringService
}
type GetServices struct {
	Offset   int64
	Limit    int64
	_        int
	Services []ScoringService
}
type GetServicesForProject struct {
	ProjectId int64
	Offset    int64
	Limit     int64
	_         int
	Services  []ScoringService
}
type GetServicesForModel struct {
	ModelId  int64
	Offset   int64
	Limit    int64
	_        int
	Services []ScoringService
}
type DeleteService struct {
	ServiceId int64
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
type GetAllEntityTypes struct {
	_           int
	EntityTypes []EntityType `help:"A list of Steam entity types."`
}
type GetAllPermissions struct {
	_           int
	Permissions []Permission `help:"A list of Steam permissions."`
}
type GetAllClusterTypes struct {
	_            int
	ClusterTypes []ClusterType
}
type GetPermissionsForRole struct {
	RoleId      int64 `help:"Integer ID of a role in Steam."`
	_           int
	Permissions []Permission `help:"A list of Steam permissions."`
}
type GetPermissionsForIdentity struct {
	IdentityId  int64 `help:"Integer ID of an identity in Steam."`
	_           int
	Permissions []Permission `help:"A list of Steam permissions."`
}
type CreateRole struct {
	Name        string `help:"A string name."`
	Description string `help:"A string description"`
	_           int
	RoleId      int64 `help:"Integer ID of the role in Steam."`
}
type GetRoles struct {
	Offset int64 `help:"An offset to start the search on."`
	Limit  int64 `help:"The maximum returned objects."`
	_      int
	Roles  []Role `help:"A list of Steam roles."`
}
type GetRolesForIdentity struct {
	IdentityId int64 `help:"Integer ID of an identity in Steam."`
	_          int
	Roles      []Role `help:"A list of Steam roles."`
}
type GetRole struct {
	RoleId int64 `help:"Integer ID of a role in Steam."`
	_      int
	Role   Role `help:"A Steam role."`
}
type GetRoleByName struct {
	Name string `help:"A role name."`
	_    int
	Role Role `help:"A Steam role."`
}
type UpdateRole struct {
	RoleId      int64  `help:"Integer ID of a role in Steam."`
	Name        string `help:"A string name."`
	Description string `help:"A string description"`
}
type LinkRoleWithPermissions struct {
	RoleId        int64   `help:"Integer ID of a role in Steam."`
	PermissionIds []int64 `help:"A list of Integer IDs for permissions in Steam."`
}
type LinkRoleWithPermission struct {
	RoleId       int64 `help:"Integer ID of a role in Steam."`
	PermissionId int64 `help:"Integer ID of a permission in Steam."`
}
type UnlinkRoleFromPermission struct {
	RoleId       int64 `help:"Integer ID of a role in Steam."`
	PermissionId int64 `help:"Integer ID of a permission in Steam."`
}
type DeleteRole struct {
	RoleId int64 `help:"Integer ID of a role in Steam."`
}
type CreateWorkgroup struct {
	Name        string `help:"A string name."`
	Description string `help:"A string description"`
	_           int
	WorkgroupId int64 `help:"Integer ID of the workgroup in Steam."`
}
type GetWorkgroups struct {
	Offset     int64 `help:"An offset to start the search on."`
	Limit      int64 `help:"The maximum returned objects."`
	_          int
	Workgroups []Workgroup `help:"A list of workgroups in Steam."`
}
type GetWorkgroupsForIdentity struct {
	IdentityId int64 `help:"Integer ID of an identity in Steam."`
	_          int
	Workgroups []Workgroup `help:"A list of workgroups in Steam."`
}
type GetWorkgroup struct {
	WorkgroupId int64 `help:"Integer ID of a workgroup in Steam."`
	_           int
	Workgroup   Workgroup `help:"A workgroup in Steam."`
}
type GetWorkgroupByName struct {
	Name      string `help:"A string name."`
	_         int
	Workgroup Workgroup `help:"A workgroup in Steam."`
}
type UpdateWorkgroup struct {
	WorkgroupId int64  `help:"Integer ID of a workgrou in Steam."`
	Name        string `help:"A string name."`
	Description string `help:"A string description"`
}
type DeleteWorkgroup struct {
	WorkgroupId int64 `help:"Integer ID of a workgroup in Steam."`
}
type CreateIdentity struct {
	Name       string `help:"A string name."`
	Password   string `help:"A string password"`
	_          int
	IdentityId int64 `help:"Integer ID of the identity in Steam."`
}
type GetIdentities struct {
	Offset     int64 `help:"An offset to start the search on."`
	Limit      int64 `help:"The maximum returned objects."`
	_          int
	Identities []Identity `help:"A list of identities in Steam."`
}
type GetIdentitiesForWorkgroup struct {
	WorkgroupId int64 `help:"Integer ID of a workgroup in Steam."`
	_           int
	Identities  []Identity `help:"A list of identities in Steam."`
}
type GetIdentitiesForRole struct {
	RoleId     int64 `help:"Integer ID of a role in Steam."`
	_          int
	Identities []Identity `help:"A list of identities in Steam."`
}
type GetIdentity struct {
	IdentityId int64 `help:"Integer ID of an identity in Steam."`
	_          int
	Identity   Identity `help:"An identity in Steam."`
}
type GetIdentityByName struct {
	Name     string `help:"An identity name."`
	_        int
	Identity Identity `help:"An identity in Steam."`
}
type GetIdentitiesForEntity struct {
	EntityType int64 `help:"An entity type ID."`
	EntityId   int64 `help:"An entity ID."`
	_          int
	Users      []UserRole `help:"A list of identites and roles"`
}
type LinkIdentityWithWorkgroup struct {
	IdentityId  int64 `help:"Integer ID of an identity in Steam."`
	WorkgroupId int64 `help:"Integer ID of a workgroup in Steam."`
}
type UnlinkIdentityFromWorkgroup struct {
	IdentityId  int64 `help:"Integer ID of an identity in Steam."`
	WorkgroupId int64 `help:"Integer ID of a workgroup in Steam."`
}
type LinkIdentityWithRole struct {
	IdentityId int64 `help:"Integer ID of an identity in Steam."`
	RoleId     int64 `help:"Integer ID of a role in Steam."`
}
type UnlinkIdentityFromRole struct {
	IdentityId int64 `help:"Integer ID of an identity in Steam."`
	RoleId     int64 `help:"Integer ID of a role in Steam."`
}
type UpdateIdentity struct {
	IdentityId int64  `help:"Integer ID of an identity in Steam."`
	Password   string `help:"Password for identity"`
}
type ActivateIdentity struct {
	IdentityId int64 `help:"Integer ID of an identity in Steam."`
}
type DeactivateIdentity struct {
	IdentityId int64 `help:"Integer ID of an identity in Steam."`
}
type ShareEntity struct {
	Kind         string `help:"Type of permission. Can be view, edit, or own."`
	WorkgroupId  int64  `help:"Integer ID of a workgroup in Steam."`
	EntityTypeId int64  `help:"Integer ID for the type of entity."`
	EntityId     int64  `help:"Integer ID for an entity in Steam."`
}
type GetPrivileges struct {
	EntityTypeId int64 `help:"Integer ID for the type of entity."`
	EntityId     int64 `help:"Integer ID for an entity in Steam."`
	_            int
	Privileges   []EntityPrivilege `help:"A list of entity privileges"`
}
type UnshareEntity struct {
	Kind         string `help:"Type of permission. Can be view, edit, or own."`
	WorkgroupId  int64  `help:"Integer ID of a workgroup in Steam."`
	EntityTypeId int64  `help:"Integer ID for the type of entity."`
	EntityId     int64  `help:"Integer ID for an entity in Steam."`
}
type GetHistory struct {
	EntityTypeId int64 `help:"Integer ID for the type of entity."`
	EntityId     int64 `help:"Integer ID for an entity in Steam."`
	Offset       int64 `help:"An offset to start the search on."`
	Limit        int64 `help:"The maximum returned objects."`
	_            int
	History      []EntityHistory `help:"A list of actions performed on the entity."`
}

type CreatePackage struct {
	ProjectId int64
	Name      string
}

type GetPackages struct {
	ProjectId int64
	_         int
	Packages  []string
}

type GetPackageDirectories struct {
	ProjectId    int64
	PackageName  string
	RelativePath string
	_            int
	Directories  []string
}

type GetPackageFiles struct {
	ProjectId    int64
	PackageName  string
	RelativePath string
	_            int
	Files        []string
}

type DeletePackage struct {
	ProjectId int64
	Name      string
}

type DeletePackageDirectory struct {
	ProjectId    int64
	PackageName  string
	RelativePath string
}

type DeletePackageFile struct {
	ProjectId    int64
	PackageName  string
	RelativePath string
}

type SetAttributesForPackage struct {
	ProjectId   int64
	PackageName string
	Attributes  string
}

type GetAttributesForPackage struct {
	ProjectId   int64
	PackageName string
	_           int
	Attributes  string
}
