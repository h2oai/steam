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
	Properties         string
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
	MaxRuntime          int
	Metrics             string
	CreatedAt           int64
	LabelId             int64
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
	MaxRuntime          int
	Metrics             string
	CreatedAt           int64
	LabelId             int64
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
	MaxRuntime          int
	Metrics             string
	CreatedAt           int64
	LabelId             int64
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
	MaxRuntime           int
	Metrics              string
	CreatedAt            int64
	LabelId              int64
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
	PingServer                    PingServer                    `help:"Ping the Steam server"`
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
	AddEngine                     AddEngine                     `help:"Add an engine"`
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
	GetIdentity                   GetIdentity                   `help:"Get identity details"`
	GetIdentityByName             GetIdentityByName             `help:"Get identity details by name"`
	LinkIdentityWithWorkgroup     LinkIdentityWithWorkgroup     `help:"Link an identity with a workgroup"`
	UnlinkIdentityFromWorkgroup   UnlinkIdentityFromWorkgroup   `help:"Unlink an identity from a workgroup"`
	LinkIdentityWithRole          LinkIdentityWithRole          `help:"Link an identity with a role"`
	UnlinkIdentityFromRole        UnlinkIdentityFromRole        `help:"Unlink an identity from a role"`
	UpdateIdentity                UpdateIdentity                `help:"Update an identity"`
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
	Username    string
	_           int
	ClusterId   int64
}
type StopClusterOnYarn struct {
	ClusterId int64
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
type GetAllEntityTypes struct {
	_           int
	EntityTypes []EntityType
}
type GetAllPermissions struct {
	_           int
	Permissions []Permission
}
type GetAllClusterTypes struct {
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
type LinkRoleWithPermissions struct {
	RoleId        int64
	PermissionIds []int64
}
type LinkRoleWithPermission struct {
	RoleId       int64
	PermissionId int64
}
type UnlinkRoleFromPermission struct {
	RoleId       int64
	PermissionId int64
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
type LinkIdentityWithWorkgroup struct {
	IdentityId  int64
	WorkgroupId int64
}
type UnlinkIdentityFromWorkgroup struct {
	IdentityId  int64
	WorkgroupId int64
}
type LinkIdentityWithRole struct {
	IdentityId int64
	RoleId     int64
}
type UnlinkIdentityFromRole struct {
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
type GetPrivileges struct {
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
type GetHistory struct {
	EntityTypeId int64
	EntityId     int64
	Offset       int64
	Limit        int64
	_            int
	History      []EntityHistory
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
