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

// ------------------------------
// --- This is generated code ---
// ---      DO NOT EDIT       ---
// ------------------------------

package web

import (
	"encoding/json"
	"github.com/h2oai/steam/master/az"
	"github.com/rs/xid"
	"log"
	"net/http"
)

// --- Types ---

type BinomialModel struct {
	Id                  int64   `json:"id"`
	TrainingDatasetId   int64   `json:"training_dataset_id"`
	ValidationDatasetId int64   `json:"validation_dataset_id"`
	Name                string  `json:"name"`
	ClusterName         string  `json:"cluster_name"`
	ModelKey            string  `json:"model_key"`
	Algorithm           string  `json:"algorithm"`
	ModelCategory       string  `json:"model_category"`
	DatasetName         string  `json:"dataset_name"`
	ResponseColumnName  string  `json:"response_column_name"`
	LogicalName         string  `json:"logical_name"`
	Location            string  `json:"location"`
	ModelObjectType     string  `json:"model_object_type"`
	MaxRuntime          int     `json:"max_runtime"`
	JSONMetrics         string  `json:"json_metrics"`
	CreatedAt           int64   `json:"created_at"`
	LabelId             int64   `json:"label_id"`
	LabelName           string  `json:"label_name"`
	Mse                 float64 `json:"mse"`
	RSquared            float64 `json:"r_squared"`
	Logloss             float64 `json:"logloss"`
	Auc                 float64 `json:"auc"`
	Gini                float64 `json:"gini"`
}

type Cluster struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	TypeId    int64  `json:"type_id"`
	DetailId  int64  `json:"detail_id"`
	Address   string `json:"address"`
	Token     string `json:"token"`
	State     string `json:"state"`
	CreatedAt int64  `json:"created_at"`
}

type ClusterStatus struct {
	Version              string `json:"version"`
	Status               string `json:"status"`
	MaxMemory            string `json:"max_memory"`
	TotalCpuCount        int    `json:"total_cpu_count"`
	TotalAllowedCpuCount int    `json:"total_allowed_cpu_count"`
}

type ClusterType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Config struct {
	KerberosEnabled     bool   `json:"kerberos_enabled"`
	ClusterProxyAddress string `json:"cluster_proxy_address"`
}

type Dataset struct {
	Id                 int64  `json:"id"`
	DatasourceId       int64  `json:"datasource_id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	FrameName          string `json:"frame_name"`
	ResponseColumnName string `json:"response_column_name"`
	JSONProperties     string `json:"json_properties"`
	CreatedAt          int64  `json:"created_at"`
}

type Datasource struct {
	Id            int64  `json:"id"`
	ProjectId     int64  `json:"project_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Kind          string `json:"kind"`
	Configuration string `json:"configuration"`
	CreatedAt     int64  `json:"created_at"`
}

type Engine struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	CreatedAt int64  `json:"created_at"`
}

type EntityHistory struct {
	IdentityId  int64  `json:"identity_id"`
	Action      string `json:"action"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

type EntityPrivilege struct {
	Kind                 string `json:"kind"`
	WorkgroupId          int64  `json:"workgroup_id"`
	WorkgroupName        string `json:"workgroup_name"`
	WorkgroupDescription string `json:"workgroup_description"`
}

type EntityType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Identity struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	LastLogin int64  `json:"last_login"`
	Created   int64  `json:"created"`
}

type Job struct {
	Name        string `json:"name"`
	ClusterName string `json:"cluster_name"`
	Description string `json:"description"`
	Progress    string `json:"progress"`
	StartedAt   int64  `json:"started_at"`
	CompletedAt int64  `json:"completed_at"`
}

type Label struct {
	Id          int64  `json:"id"`
	ProjectId   int64  `json:"project_id"`
	ModelId     int64  `json:"model_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

type Model struct {
	Id                  int64  `json:"id"`
	TrainingDatasetId   int64  `json:"training_dataset_id"`
	ValidationDatasetId int64  `json:"validation_dataset_id"`
	Name                string `json:"name"`
	ClusterName         string `json:"cluster_name"`
	ModelKey            string `json:"model_key"`
	Algorithm           string `json:"algorithm"`
	ModelCategory       string `json:"model_category"`
	DatasetName         string `json:"dataset_name"`
	ResponseColumnName  string `json:"response_column_name"`
	LogicalName         string `json:"logical_name"`
	Location            string `json:"location"`
	ModelObjectType     string `json:"model_object_type"`
	MaxRuntime          int    `json:"max_runtime"`
	JSONMetrics         string `json:"json_metrics"`
	CreatedAt           int64  `json:"created_at"`
	LabelId             int64  `json:"label_id"`
	LabelName           string `json:"label_name"`
}

type MultinomialModel struct {
	Id                  int64   `json:"id"`
	TrainingDatasetId   int64   `json:"training_dataset_id"`
	ValidationDatasetId int64   `json:"validation_dataset_id"`
	Name                string  `json:"name"`
	ClusterName         string  `json:"cluster_name"`
	ModelKey            string  `json:"model_key"`
	Algorithm           string  `json:"algorithm"`
	ModelCategory       string  `json:"model_category"`
	DatasetName         string  `json:"dataset_name"`
	ResponseColumnName  string  `json:"response_column_name"`
	LogicalName         string  `json:"logical_name"`
	Location            string  `json:"location"`
	ModelObjectType     string  `json:"model_object_type"`
	MaxRuntime          int     `json:"max_runtime"`
	JSONMetrics         string  `json:"json_metrics"`
	CreatedAt           int64   `json:"created_at"`
	LabelId             int64   `json:"label_id"`
	LabelName           string  `json:"label_name"`
	Mse                 float64 `json:"mse"`
	RSquared            float64 `json:"r_squared"`
	Logloss             float64 `json:"logloss"`
}

type Permission struct {
	Id          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Project struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ModelCategory string `json:"model_category"`
	CreatedAt     int64  `json:"created_at"`
}

type RegressionModel struct {
	Id                   int64   `json:"id"`
	TrainingDatasetId    int64   `json:"training_dataset_id"`
	ValidationDatasetId  int64   `json:"validation_dataset_id"`
	Name                 string  `json:"name"`
	ClusterName          string  `json:"cluster_name"`
	ModelKey             string  `json:"model_key"`
	Algorithm            string  `json:"algorithm"`
	ModelCategory        string  `json:"model_category"`
	DatasetName          string  `json:"dataset_name"`
	ResponseColumnName   string  `json:"response_column_name"`
	LogicalName          string  `json:"logical_name"`
	Location             string  `json:"location"`
	ModelObjectType      string  `json:"model_object_type"`
	MaxRuntime           int     `json:"max_runtime"`
	JSONMetrics          string  `json:"json_metrics"`
	CreatedAt            int64   `json:"created_at"`
	LabelId              int64   `json:"label_id"`
	LabelName            string  `json:"label_name"`
	Mse                  float64 `json:"mse"`
	RSquared             float64 `json:"r_squared"`
	MeanResidualDeviance float64 `json:"mean_residual_deviance"`
}

type Role struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Created     int64  `json:"created"`
}

type ScoringService struct {
	Id        int64  `json:"id"`
	ModelId   int64  `json:"model_id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Port      int    `json:"port"`
	ProcessId int    `json:"process_id"`
	State     string `json:"state"`
	CreatedAt int64  `json:"created_at"`
}

type UserRole struct {
	Kind         string `json:"kind"`
	IdentityId   int64  `json:"identity_id"`
	IdentityName string `json:"identity_name"`
	RoleId       int64  `json:"role_id"`
	RoleName     string `json:"role_name"`
}

type Workgroup struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Created     int64  `json:"created"`
}

type YarnCluster struct {
	Id            int64  `json:"id"`
	EngineId      int64  `json:"engine_id"`
	Size          int    `json:"size"`
	ApplicationId string `json:"application_id"`
	Memory        string `json:"memory"`
	Username      string `json:"username"`
}

// --- Interface ---

type Az interface {
	Identify(r *http.Request) (az.Principal, error)
}
type Service interface {
	PingServer(pz az.Principal, input string) (string, error)
	GetConfig(pz az.Principal) (*Config, error)
	RegisterCluster(pz az.Principal, address string) (int64, error)
	UnregisterCluster(pz az.Principal, clusterId int64) error
	StartClusterOnYarn(pz az.Principal, clusterName string, engineId int64, size int, memory string, secure bool, keytab string) (int64, error)
	StopClusterOnYarn(pz az.Principal, clusterId int64, keytab string) error
	GetCluster(pz az.Principal, clusterId int64) (*Cluster, error)
	GetClusterOnYarn(pz az.Principal, clusterId int64) (*YarnCluster, error)
	GetClusters(pz az.Principal, offset int64, limit int64) ([]*Cluster, error)
	GetClusterStatus(pz az.Principal, clusterId int64) (*ClusterStatus, error)
	DeleteCluster(pz az.Principal, clusterId int64) error
	GetJob(pz az.Principal, clusterId int64, jobName string) (*Job, error)
	GetJobs(pz az.Principal, clusterId int64) ([]*Job, error)
	CreateProject(pz az.Principal, name string, description string, modelCategory string) (int64, error)
	GetProjects(pz az.Principal, offset int64, limit int64) ([]*Project, error)
	GetProject(pz az.Principal, projectId int64) (*Project, error)
	DeleteProject(pz az.Principal, projectId int64) error
	CreateDatasource(pz az.Principal, projectId int64, name string, description string, path string) (int64, error)
	GetDatasources(pz az.Principal, projectId int64, offset int64, limit int64) ([]*Datasource, error)
	GetDatasource(pz az.Principal, datasourceId int64) (*Datasource, error)
	UpdateDatasource(pz az.Principal, datasourceId int64, name string, description string, path string) error
	DeleteDatasource(pz az.Principal, datasourceId int64) error
	CreateDataset(pz az.Principal, clusterId int64, datasourceId int64, name string, description string, responseColumnName string) (int64, error)
	GetDatasets(pz az.Principal, datasourceId int64, offset int64, limit int64) ([]*Dataset, error)
	GetDataset(pz az.Principal, datasetId int64) (*Dataset, error)
	GetDatasetsFromCluster(pz az.Principal, clusterId int64) ([]*Dataset, error)
	UpdateDataset(pz az.Principal, datasetId int64, name string, description string, responseColumnName string) error
	SplitDataset(pz az.Principal, datasetId int64, ratio1 int, ratio2 int) ([]int64, error)
	DeleteDataset(pz az.Principal, datasetId int64) error
	BuildModel(pz az.Principal, clusterId int64, datasetId int64, algorithm string) (int64, error)
	BuildModelAuto(pz az.Principal, clusterId int64, dataset string, targetName string, maxRunTime int) (*Model, error)
	GetModel(pz az.Principal, modelId int64) (*Model, error)
	GetModels(pz az.Principal, projectId int64, offset int64, limit int64) ([]*Model, error)
	GetModelsFromCluster(pz az.Principal, clusterId int64, frameKey string) ([]*Model, error)
	FindModelsCount(pz az.Principal, projectId int64) (int64, error)
	GetAllBinomialSortCriteria(pz az.Principal) ([]string, error)
	FindModelsBinomial(pz az.Principal, projectId int64, namePart string, sortBy string, ascending bool, offset int64, limit int64) ([]*BinomialModel, error)
	GetModelBinomial(pz az.Principal, modelId int64) (*BinomialModel, error)
	GetAllMultinomialSortCriteria(pz az.Principal) ([]string, error)
	FindModelsMultinomial(pz az.Principal, projectId int64, namePart string, sortBy string, ascending bool, offset int64, limit int64) ([]*MultinomialModel, error)
	GetModelMultinomial(pz az.Principal, modelId int64) (*MultinomialModel, error)
	GetAllRegressionSortCriteria(pz az.Principal) ([]string, error)
	FindModelsRegression(pz az.Principal, projectId int64, namePart string, sortBy string, ascending bool, offset int64, limit int64) ([]*RegressionModel, error)
	GetModelRegression(pz az.Principal, modelId int64) (*RegressionModel, error)
	ImportModelFromCluster(pz az.Principal, clusterId int64, projectId int64, modelKey string, modelName string) (int64, error)
	CheckMojo(pz az.Principal, algo string) (bool, error)
	ImportModelPojo(pz az.Principal, modelId int64) error
	ImportModelMojo(pz az.Principal, modelId int64) error
	DeleteModel(pz az.Principal, modelId int64) error
	CreateLabel(pz az.Principal, projectId int64, name string, description string) (int64, error)
	UpdateLabel(pz az.Principal, labelId int64, name string, description string) error
	DeleteLabel(pz az.Principal, labelId int64) error
	LinkLabelWithModel(pz az.Principal, labelId int64, modelId int64) error
	UnlinkLabelFromModel(pz az.Principal, labelId int64, modelId int64) error
	GetLabelsForProject(pz az.Principal, projectId int64) ([]*Label, error)
	StartService(pz az.Principal, modelId int64, name string, packageName string) (int64, error)
	StopService(pz az.Principal, serviceId int64) error
	GetService(pz az.Principal, serviceId int64) (*ScoringService, error)
	GetServices(pz az.Principal, offset int64, limit int64) ([]*ScoringService, error)
	GetServicesForProject(pz az.Principal, projectId int64, offset int64, limit int64) ([]*ScoringService, error)
	GetServicesForModel(pz az.Principal, modelId int64, offset int64, limit int64) ([]*ScoringService, error)
	DeleteService(pz az.Principal, serviceId int64) error
	GetEngine(pz az.Principal, engineId int64) (*Engine, error)
	GetEngines(pz az.Principal) ([]*Engine, error)
	DeleteEngine(pz az.Principal, engineId int64) error
	GetAllEntityTypes(pz az.Principal) ([]*EntityType, error)
	GetAllPermissions(pz az.Principal) ([]*Permission, error)
	GetAllClusterTypes(pz az.Principal) ([]*ClusterType, error)
	GetPermissionsForRole(pz az.Principal, roleId int64) ([]*Permission, error)
	GetPermissionsForIdentity(pz az.Principal, identityId int64) ([]*Permission, error)
	CreateRole(pz az.Principal, name string, description string) (int64, error)
	GetRoles(pz az.Principal, offset int64, limit int64) ([]*Role, error)
	GetRolesForIdentity(pz az.Principal, identityId int64) ([]*Role, error)
	GetRole(pz az.Principal, roleId int64) (*Role, error)
	GetRoleByName(pz az.Principal, name string) (*Role, error)
	UpdateRole(pz az.Principal, roleId int64, name string, description string) error
	LinkRoleWithPermissions(pz az.Principal, roleId int64, permissionIds []int64) error
	LinkRoleWithPermission(pz az.Principal, roleId int64, permissionId int64) error
	UnlinkRoleFromPermission(pz az.Principal, roleId int64, permissionId int64) error
	DeleteRole(pz az.Principal, roleId int64) error
	CreateWorkgroup(pz az.Principal, name string, description string) (int64, error)
	GetWorkgroups(pz az.Principal, offset int64, limit int64) ([]*Workgroup, error)
	GetWorkgroupsForIdentity(pz az.Principal, identityId int64) ([]*Workgroup, error)
	GetWorkgroup(pz az.Principal, workgroupId int64) (*Workgroup, error)
	GetWorkgroupByName(pz az.Principal, name string) (*Workgroup, error)
	UpdateWorkgroup(pz az.Principal, workgroupId int64, name string, description string) error
	DeleteWorkgroup(pz az.Principal, workgroupId int64) error
	CreateIdentity(pz az.Principal, name string, password string) (int64, error)
	GetIdentities(pz az.Principal, offset int64, limit int64) ([]*Identity, error)
	GetIdentitiesForWorkgroup(pz az.Principal, workgroupId int64) ([]*Identity, error)
	GetIdentitiesForRole(pz az.Principal, roleId int64) ([]*Identity, error)
	GetIdentitiesForEntity(pz az.Principal, entityType int64, entityId int64) ([]*UserRole, error)
	GetIdentity(pz az.Principal, identityId int64) (*Identity, error)
	GetIdentityByName(pz az.Principal, name string) (*Identity, error)
	LinkIdentityWithWorkgroup(pz az.Principal, identityId int64, workgroupId int64) error
	UnlinkIdentityFromWorkgroup(pz az.Principal, identityId int64, workgroupId int64) error
	LinkIdentityWithRole(pz az.Principal, identityId int64, roleId int64) error
	UnlinkIdentityFromRole(pz az.Principal, identityId int64, roleId int64) error
	UpdateIdentity(pz az.Principal, identityId int64, password string) error
	ActivateIdentity(pz az.Principal, identityId int64) error
	DeactivateIdentity(pz az.Principal, identityId int64) error
	ShareEntity(pz az.Principal, kind string, workgroupId int64, entityTypeId int64, entityId int64) error
	GetPrivileges(pz az.Principal, entityTypeId int64, entityId int64) ([]*EntityPrivilege, error)
	UnshareEntity(pz az.Principal, kind string, workgroupId int64, entityTypeId int64, entityId int64) error
	GetHistory(pz az.Principal, entityTypeId int64, entityId int64, offset int64, limit int64) ([]*EntityHistory, error)
	CreatePackage(pz az.Principal, projectId int64, name string) error
	GetPackages(pz az.Principal, projectId int64) ([]string, error)
	GetPackageDirectories(pz az.Principal, projectId int64, packageName string, relativePath string) ([]string, error)
	GetPackageFiles(pz az.Principal, projectId int64, packageName string, relativePath string) ([]string, error)
	DeletePackage(pz az.Principal, projectId int64, name string) error
	DeletePackageDirectory(pz az.Principal, projectId int64, packageName string, relativePath string) error
	DeletePackageFile(pz az.Principal, projectId int64, packageName string, relativePath string) error
	SetAttributesForPackage(pz az.Principal, projectId int64, packageName string, attributes string) error
	GetAttributesForPackage(pz az.Principal, projectId int64, packageName string) (string, error)
}

// --- Messages ---
type PingServerIn struct {
	Input string `json:"input"`
}

type PingServerOut struct {
	Output string `json:"output"`
}

type GetConfigIn struct {
}

type GetConfigOut struct {
	Config *Config `json:"config"`
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

type StartClusterOnYarnIn struct {
	ClusterName string `json:"cluster_name"`
	EngineId    int64  `json:"engine_id"`
	Size        int    `json:"size"`
	Memory      string `json:"memory"`
	Secure      bool   `json:"secure"`
	Keytab      string `json:"keytab"`
}

type StartClusterOnYarnOut struct {
	ClusterId int64 `json:"cluster_id"`
}

type StopClusterOnYarnIn struct {
	ClusterId int64  `json:"cluster_id"`
	Keytab    string `json:"keytab"`
}

type StopClusterOnYarnOut struct {
}

type GetClusterIn struct {
	ClusterId int64 `json:"cluster_id"`
}

type GetClusterOut struct {
	Cluster *Cluster `json:"cluster"`
}

type GetClusterOnYarnIn struct {
	ClusterId int64 `json:"cluster_id"`
}

type GetClusterOnYarnOut struct {
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

type CreateProjectIn struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	ModelCategory string `json:"model_category"`
}

type CreateProjectOut struct {
	ProjectId int64 `json:"project_id"`
}

type GetProjectsIn struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
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
	ProjectId   int64  `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
}

type CreateDatasourceOut struct {
	DatasourceId int64 `json:"datasource_id"`
}

type GetDatasourcesIn struct {
	ProjectId int64 `json:"project_id"`
	Offset    int64 `json:"offset"`
	Limit     int64 `json:"limit"`
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
	DatasourceId int64  `json:"datasource_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Path         string `json:"path"`
}

type UpdateDatasourceOut struct {
}

type DeleteDatasourceIn struct {
	DatasourceId int64 `json:"datasource_id"`
}

type DeleteDatasourceOut struct {
}

type CreateDatasetIn struct {
	ClusterId          int64  `json:"cluster_id"`
	DatasourceId       int64  `json:"datasource_id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	ResponseColumnName string `json:"response_column_name"`
}

type CreateDatasetOut struct {
	DatasetId int64 `json:"dataset_id"`
}

type GetDatasetsIn struct {
	DatasourceId int64 `json:"datasource_id"`
	Offset       int64 `json:"offset"`
	Limit        int64 `json:"limit"`
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

type GetDatasetsFromClusterIn struct {
	ClusterId int64 `json:"cluster_id"`
}

type GetDatasetsFromClusterOut struct {
	Dataset []*Dataset `json:"dataset"`
}

type UpdateDatasetIn struct {
	DatasetId          int64  `json:"dataset_id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	ResponseColumnName string `json:"response_column_name"`
}

type UpdateDatasetOut struct {
}

type SplitDatasetIn struct {
	DatasetId int64 `json:"dataset_id"`
	Ratio1    int   `json:"ratio1"`
	Ratio2    int   `json:"ratio2"`
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
	ClusterId int64  `json:"cluster_id"`
	DatasetId int64  `json:"dataset_id"`
	Algorithm string `json:"algorithm"`
}

type BuildModelOut struct {
	ModelId int64 `json:"model_id"`
}

type BuildModelAutoIn struct {
	ClusterId  int64  `json:"cluster_id"`
	Dataset    string `json:"dataset"`
	TargetName string `json:"target_name"`
	MaxRunTime int    `json:"max_run_time"`
}

type BuildModelAutoOut struct {
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
	Offset    int64 `json:"offset"`
	Limit     int64 `json:"limit"`
}

type GetModelsOut struct {
	Models []*Model `json:"models"`
}

type GetModelsFromClusterIn struct {
	ClusterId int64  `json:"cluster_id"`
	FrameKey  string `json:"frame_key"`
}

type GetModelsFromClusterOut struct {
	Models []*Model `json:"models"`
}

type FindModelsCountIn struct {
	ProjectId int64 `json:"project_id"`
}

type FindModelsCountOut struct {
	Count int64 `json:"count"`
}

type GetAllBinomialSortCriteriaIn struct {
}

type GetAllBinomialSortCriteriaOut struct {
	Criteria []string `json:"criteria"`
}

type FindModelsBinomialIn struct {
	ProjectId int64  `json:"project_id"`
	NamePart  string `json:"name_part"`
	SortBy    string `json:"sort_by"`
	Ascending bool   `json:"ascending"`
	Offset    int64  `json:"offset"`
	Limit     int64  `json:"limit"`
}

type FindModelsBinomialOut struct {
	Models []*BinomialModel `json:"models"`
}

type GetModelBinomialIn struct {
	ModelId int64 `json:"model_id"`
}

type GetModelBinomialOut struct {
	Model *BinomialModel `json:"model"`
}

type GetAllMultinomialSortCriteriaIn struct {
}

type GetAllMultinomialSortCriteriaOut struct {
	Criteria []string `json:"criteria"`
}

type FindModelsMultinomialIn struct {
	ProjectId int64  `json:"project_id"`
	NamePart  string `json:"name_part"`
	SortBy    string `json:"sort_by"`
	Ascending bool   `json:"ascending"`
	Offset    int64  `json:"offset"`
	Limit     int64  `json:"limit"`
}

type FindModelsMultinomialOut struct {
	Models []*MultinomialModel `json:"models"`
}

type GetModelMultinomialIn struct {
	ModelId int64 `json:"model_id"`
}

type GetModelMultinomialOut struct {
	Model *MultinomialModel `json:"model"`
}

type GetAllRegressionSortCriteriaIn struct {
}

type GetAllRegressionSortCriteriaOut struct {
	Criteria []string `json:"criteria"`
}

type FindModelsRegressionIn struct {
	ProjectId int64  `json:"project_id"`
	NamePart  string `json:"name_part"`
	SortBy    string `json:"sort_by"`
	Ascending bool   `json:"ascending"`
	Offset    int64  `json:"offset"`
	Limit     int64  `json:"limit"`
}

type FindModelsRegressionOut struct {
	Models []*RegressionModel `json:"models"`
}

type GetModelRegressionIn struct {
	ModelId int64 `json:"model_id"`
}

type GetModelRegressionOut struct {
	Model *RegressionModel `json:"model"`
}

type ImportModelFromClusterIn struct {
	ClusterId int64  `json:"cluster_id"`
	ProjectId int64  `json:"project_id"`
	ModelKey  string `json:"model_key"`
	ModelName string `json:"model_name"`
}

type ImportModelFromClusterOut struct {
	ModelId int64 `json:"model_id"`
}

type CheckMojoIn struct {
	Algo string `json:"algo"`
}

type CheckMojoOut struct {
	CanMojo bool `json:"can_mojo"`
}

type ImportModelPojoIn struct {
	ModelId int64 `json:"model_id"`
}

type ImportModelPojoOut struct {
}

type ImportModelMojoIn struct {
	ModelId int64 `json:"model_id"`
}

type ImportModelMojoOut struct {
}

type DeleteModelIn struct {
	ModelId int64 `json:"model_id"`
}

type DeleteModelOut struct {
}

type CreateLabelIn struct {
	ProjectId   int64  `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateLabelOut struct {
	LabelId int64 `json:"label_id"`
}

type UpdateLabelIn struct {
	LabelId     int64  `json:"label_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateLabelOut struct {
}

type DeleteLabelIn struct {
	LabelId int64 `json:"label_id"`
}

type DeleteLabelOut struct {
}

type LinkLabelWithModelIn struct {
	LabelId int64 `json:"label_id"`
	ModelId int64 `json:"model_id"`
}

type LinkLabelWithModelOut struct {
}

type UnlinkLabelFromModelIn struct {
	LabelId int64 `json:"label_id"`
	ModelId int64 `json:"model_id"`
}

type UnlinkLabelFromModelOut struct {
}

type GetLabelsForProjectIn struct {
	ProjectId int64 `json:"project_id"`
}

type GetLabelsForProjectOut struct {
	Labels []*Label `json:"labels"`
}

type StartServiceIn struct {
	ModelId     int64  `json:"model_id"`
	Name        string `json:"name"`
	PackageName string `json:"package_name"`
}

type StartServiceOut struct {
	ServiceId int64 `json:"service_id"`
}

type StopServiceIn struct {
	ServiceId int64 `json:"service_id"`
}

type StopServiceOut struct {
}

type GetServiceIn struct {
	ServiceId int64 `json:"service_id"`
}

type GetServiceOut struct {
	Service *ScoringService `json:"service"`
}

type GetServicesIn struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetServicesOut struct {
	Services []*ScoringService `json:"services"`
}

type GetServicesForProjectIn struct {
	ProjectId int64 `json:"project_id"`
	Offset    int64 `json:"offset"`
	Limit     int64 `json:"limit"`
}

type GetServicesForProjectOut struct {
	Services []*ScoringService `json:"services"`
}

type GetServicesForModelIn struct {
	ModelId int64 `json:"model_id"`
	Offset  int64 `json:"offset"`
	Limit   int64 `json:"limit"`
}

type GetServicesForModelOut struct {
	Services []*ScoringService `json:"services"`
}

type DeleteServiceIn struct {
	ServiceId int64 `json:"service_id"`
}

type DeleteServiceOut struct {
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

type GetAllEntityTypesIn struct {
}

type GetAllEntityTypesOut struct {
	EntityTypes []*EntityType `json:"entity_types"`
}

type GetAllPermissionsIn struct {
}

type GetAllPermissionsOut struct {
	Permissions []*Permission `json:"permissions"`
}

type GetAllClusterTypesIn struct {
}

type GetAllClusterTypesOut struct {
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

type GetRoleByNameIn struct {
	Name string `json:"name"`
}

type GetRoleByNameOut struct {
	Role *Role `json:"role"`
}

type UpdateRoleIn struct {
	RoleId      int64  `json:"role_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateRoleOut struct {
}

type LinkRoleWithPermissionsIn struct {
	RoleId        int64   `json:"role_id"`
	PermissionIds []int64 `json:"permission_ids"`
}

type LinkRoleWithPermissionsOut struct {
}

type LinkRoleWithPermissionIn struct {
	RoleId       int64 `json:"role_id"`
	PermissionId int64 `json:"permission_id"`
}

type LinkRoleWithPermissionOut struct {
}

type UnlinkRoleFromPermissionIn struct {
	RoleId       int64 `json:"role_id"`
	PermissionId int64 `json:"permission_id"`
}

type UnlinkRoleFromPermissionOut struct {
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

type GetWorkgroupByNameIn struct {
	Name string `json:"name"`
}

type GetWorkgroupByNameOut struct {
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

type GetIdentitiesForRoleIn struct {
	RoleId int64 `json:"role_id"`
}

type GetIdentitiesForRoleOut struct {
	Identities []*Identity `json:"identities"`
}

type GetIdentitiesForEntityIn struct {
	EntityType int64 `json:"entity_type"`
	EntityId   int64 `json:"entity_id"`
}

type GetIdentitiesForEntityOut struct {
	Users []*UserRole `json:"users"`
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

type LinkIdentityWithWorkgroupIn struct {
	IdentityId  int64 `json:"identity_id"`
	WorkgroupId int64 `json:"workgroup_id"`
}

type LinkIdentityWithWorkgroupOut struct {
}

type UnlinkIdentityFromWorkgroupIn struct {
	IdentityId  int64 `json:"identity_id"`
	WorkgroupId int64 `json:"workgroup_id"`
}

type UnlinkIdentityFromWorkgroupOut struct {
}

type LinkIdentityWithRoleIn struct {
	IdentityId int64 `json:"identity_id"`
	RoleId     int64 `json:"role_id"`
}

type LinkIdentityWithRoleOut struct {
}

type UnlinkIdentityFromRoleIn struct {
	IdentityId int64 `json:"identity_id"`
	RoleId     int64 `json:"role_id"`
}

type UnlinkIdentityFromRoleOut struct {
}

type UpdateIdentityIn struct {
	IdentityId int64  `json:"identity_id"`
	Password   string `json:"password"`
}

type UpdateIdentityOut struct {
}

type ActivateIdentityIn struct {
	IdentityId int64 `json:"identity_id"`
}

type ActivateIdentityOut struct {
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

type GetPrivilegesIn struct {
	EntityTypeId int64 `json:"entity_type_id"`
	EntityId     int64 `json:"entity_id"`
}

type GetPrivilegesOut struct {
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

type GetHistoryIn struct {
	EntityTypeId int64 `json:"entity_type_id"`
	EntityId     int64 `json:"entity_id"`
	Offset       int64 `json:"offset"`
	Limit        int64 `json:"limit"`
}

type GetHistoryOut struct {
	History []*EntityHistory `json:"history"`
}

type CreatePackageIn struct {
	ProjectId int64  `json:"project_id"`
	Name      string `json:"name"`
}

type CreatePackageOut struct {
}

type GetPackagesIn struct {
	ProjectId int64 `json:"project_id"`
}

type GetPackagesOut struct {
	Packages []string `json:"packages"`
}

type GetPackageDirectoriesIn struct {
	ProjectId    int64  `json:"project_id"`
	PackageName  string `json:"package_name"`
	RelativePath string `json:"relative_path"`
}

type GetPackageDirectoriesOut struct {
	Directories []string `json:"directories"`
}

type GetPackageFilesIn struct {
	ProjectId    int64  `json:"project_id"`
	PackageName  string `json:"package_name"`
	RelativePath string `json:"relative_path"`
}

type GetPackageFilesOut struct {
	Files []string `json:"files"`
}

type DeletePackageIn struct {
	ProjectId int64  `json:"project_id"`
	Name      string `json:"name"`
}

type DeletePackageOut struct {
}

type DeletePackageDirectoryIn struct {
	ProjectId    int64  `json:"project_id"`
	PackageName  string `json:"package_name"`
	RelativePath string `json:"relative_path"`
}

type DeletePackageDirectoryOut struct {
}

type DeletePackageFileIn struct {
	ProjectId    int64  `json:"project_id"`
	PackageName  string `json:"package_name"`
	RelativePath string `json:"relative_path"`
}

type DeletePackageFileOut struct {
}

type SetAttributesForPackageIn struct {
	ProjectId   int64  `json:"project_id"`
	PackageName string `json:"package_name"`
	Attributes  string `json:"attributes"`
}

type SetAttributesForPackageOut struct {
}

type GetAttributesForPackageIn struct {
	ProjectId   int64  `json:"project_id"`
	PackageName string `json:"package_name"`
}

type GetAttributesForPackageOut struct {
	Attributes string `json:"attributes"`
}

// --- Client Stub ---

type Remote struct {
	Proc Proc
}

type Proc interface {
	Call(name string, in, out interface{}) error
}

func (this *Remote) PingServer(input string) (string, error) {
	in := PingServerIn{input}
	var out PingServerOut
	err := this.Proc.Call("PingServer", &in, &out)
	if err != nil {
		return "", err
	}
	return out.Output, nil
}

func (this *Remote) GetConfig() (*Config, error) {
	in := GetConfigIn{}
	var out GetConfigOut
	err := this.Proc.Call("GetConfig", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Config, nil
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

func (this *Remote) StartClusterOnYarn(clusterName string, engineId int64, size int, memory string, secure bool, keytab string) (int64, error) {
	in := StartClusterOnYarnIn{clusterName, engineId, size, memory, secure, keytab}
	var out StartClusterOnYarnOut
	err := this.Proc.Call("StartClusterOnYarn", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.ClusterId, nil
}

func (this *Remote) StopClusterOnYarn(clusterId int64, keytab string) error {
	in := StopClusterOnYarnIn{clusterId, keytab}
	var out StopClusterOnYarnOut
	err := this.Proc.Call("StopClusterOnYarn", &in, &out)
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

func (this *Remote) GetClusterOnYarn(clusterId int64) (*YarnCluster, error) {
	in := GetClusterOnYarnIn{clusterId}
	var out GetClusterOnYarnOut
	err := this.Proc.Call("GetClusterOnYarn", &in, &out)
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

func (this *Remote) CreateProject(name string, description string, modelCategory string) (int64, error) {
	in := CreateProjectIn{name, description, modelCategory}
	var out CreateProjectOut
	err := this.Proc.Call("CreateProject", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.ProjectId, nil
}

func (this *Remote) GetProjects(offset int64, limit int64) ([]*Project, error) {
	in := GetProjectsIn{offset, limit}
	var out GetProjectsOut
	err := this.Proc.Call("GetProjects", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Projects, nil
}

func (this *Remote) GetProject(projectId int64) (*Project, error) {
	in := GetProjectIn{projectId}
	var out GetProjectOut
	err := this.Proc.Call("GetProject", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Project, nil
}

func (this *Remote) DeleteProject(projectId int64) error {
	in := DeleteProjectIn{projectId}
	var out DeleteProjectOut
	err := this.Proc.Call("DeleteProject", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) CreateDatasource(projectId int64, name string, description string, path string) (int64, error) {
	in := CreateDatasourceIn{projectId, name, description, path}
	var out CreateDatasourceOut
	err := this.Proc.Call("CreateDatasource", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.DatasourceId, nil
}

func (this *Remote) GetDatasources(projectId int64, offset int64, limit int64) ([]*Datasource, error) {
	in := GetDatasourcesIn{projectId, offset, limit}
	var out GetDatasourcesOut
	err := this.Proc.Call("GetDatasources", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Datasources, nil
}

func (this *Remote) GetDatasource(datasourceId int64) (*Datasource, error) {
	in := GetDatasourceIn{datasourceId}
	var out GetDatasourceOut
	err := this.Proc.Call("GetDatasource", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Datasource, nil
}

func (this *Remote) UpdateDatasource(datasourceId int64, name string, description string, path string) error {
	in := UpdateDatasourceIn{datasourceId, name, description, path}
	var out UpdateDatasourceOut
	err := this.Proc.Call("UpdateDatasource", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) DeleteDatasource(datasourceId int64) error {
	in := DeleteDatasourceIn{datasourceId}
	var out DeleteDatasourceOut
	err := this.Proc.Call("DeleteDatasource", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) CreateDataset(clusterId int64, datasourceId int64, name string, description string, responseColumnName string) (int64, error) {
	in := CreateDatasetIn{clusterId, datasourceId, name, description, responseColumnName}
	var out CreateDatasetOut
	err := this.Proc.Call("CreateDataset", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.DatasetId, nil
}

func (this *Remote) GetDatasets(datasourceId int64, offset int64, limit int64) ([]*Dataset, error) {
	in := GetDatasetsIn{datasourceId, offset, limit}
	var out GetDatasetsOut
	err := this.Proc.Call("GetDatasets", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Datasets, nil
}

func (this *Remote) GetDataset(datasetId int64) (*Dataset, error) {
	in := GetDatasetIn{datasetId}
	var out GetDatasetOut
	err := this.Proc.Call("GetDataset", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Dataset, nil
}

func (this *Remote) GetDatasetsFromCluster(clusterId int64) ([]*Dataset, error) {
	in := GetDatasetsFromClusterIn{clusterId}
	var out GetDatasetsFromClusterOut
	err := this.Proc.Call("GetDatasetsFromCluster", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Dataset, nil
}

func (this *Remote) UpdateDataset(datasetId int64, name string, description string, responseColumnName string) error {
	in := UpdateDatasetIn{datasetId, name, description, responseColumnName}
	var out UpdateDatasetOut
	err := this.Proc.Call("UpdateDataset", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) SplitDataset(datasetId int64, ratio1 int, ratio2 int) ([]int64, error) {
	in := SplitDatasetIn{datasetId, ratio1, ratio2}
	var out SplitDatasetOut
	err := this.Proc.Call("SplitDataset", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.DatasetIds, nil
}

func (this *Remote) DeleteDataset(datasetId int64) error {
	in := DeleteDatasetIn{datasetId}
	var out DeleteDatasetOut
	err := this.Proc.Call("DeleteDataset", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) BuildModel(clusterId int64, datasetId int64, algorithm string) (int64, error) {
	in := BuildModelIn{clusterId, datasetId, algorithm}
	var out BuildModelOut
	err := this.Proc.Call("BuildModel", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.ModelId, nil
}

func (this *Remote) BuildModelAuto(clusterId int64, dataset string, targetName string, maxRunTime int) (*Model, error) {
	in := BuildModelAutoIn{clusterId, dataset, targetName, maxRunTime}
	var out BuildModelAutoOut
	err := this.Proc.Call("BuildModelAuto", &in, &out)
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

func (this *Remote) GetModels(projectId int64, offset int64, limit int64) ([]*Model, error) {
	in := GetModelsIn{projectId, offset, limit}
	var out GetModelsOut
	err := this.Proc.Call("GetModels", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Models, nil
}

func (this *Remote) GetModelsFromCluster(clusterId int64, frameKey string) ([]*Model, error) {
	in := GetModelsFromClusterIn{clusterId, frameKey}
	var out GetModelsFromClusterOut
	err := this.Proc.Call("GetModelsFromCluster", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Models, nil
}

func (this *Remote) FindModelsCount(projectId int64) (int64, error) {
	in := FindModelsCountIn{projectId}
	var out FindModelsCountOut
	err := this.Proc.Call("FindModelsCount", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.Count, nil
}

func (this *Remote) GetAllBinomialSortCriteria() ([]string, error) {
	in := GetAllBinomialSortCriteriaIn{}
	var out GetAllBinomialSortCriteriaOut
	err := this.Proc.Call("GetAllBinomialSortCriteria", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Criteria, nil
}

func (this *Remote) FindModelsBinomial(projectId int64, namePart string, sortBy string, ascending bool, offset int64, limit int64) ([]*BinomialModel, error) {
	in := FindModelsBinomialIn{projectId, namePart, sortBy, ascending, offset, limit}
	var out FindModelsBinomialOut
	err := this.Proc.Call("FindModelsBinomial", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Models, nil
}

func (this *Remote) GetModelBinomial(modelId int64) (*BinomialModel, error) {
	in := GetModelBinomialIn{modelId}
	var out GetModelBinomialOut
	err := this.Proc.Call("GetModelBinomial", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Model, nil
}

func (this *Remote) GetAllMultinomialSortCriteria() ([]string, error) {
	in := GetAllMultinomialSortCriteriaIn{}
	var out GetAllMultinomialSortCriteriaOut
	err := this.Proc.Call("GetAllMultinomialSortCriteria", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Criteria, nil
}

func (this *Remote) FindModelsMultinomial(projectId int64, namePart string, sortBy string, ascending bool, offset int64, limit int64) ([]*MultinomialModel, error) {
	in := FindModelsMultinomialIn{projectId, namePart, sortBy, ascending, offset, limit}
	var out FindModelsMultinomialOut
	err := this.Proc.Call("FindModelsMultinomial", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Models, nil
}

func (this *Remote) GetModelMultinomial(modelId int64) (*MultinomialModel, error) {
	in := GetModelMultinomialIn{modelId}
	var out GetModelMultinomialOut
	err := this.Proc.Call("GetModelMultinomial", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Model, nil
}

func (this *Remote) GetAllRegressionSortCriteria() ([]string, error) {
	in := GetAllRegressionSortCriteriaIn{}
	var out GetAllRegressionSortCriteriaOut
	err := this.Proc.Call("GetAllRegressionSortCriteria", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Criteria, nil
}

func (this *Remote) FindModelsRegression(projectId int64, namePart string, sortBy string, ascending bool, offset int64, limit int64) ([]*RegressionModel, error) {
	in := FindModelsRegressionIn{projectId, namePart, sortBy, ascending, offset, limit}
	var out FindModelsRegressionOut
	err := this.Proc.Call("FindModelsRegression", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Models, nil
}

func (this *Remote) GetModelRegression(modelId int64) (*RegressionModel, error) {
	in := GetModelRegressionIn{modelId}
	var out GetModelRegressionOut
	err := this.Proc.Call("GetModelRegression", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Model, nil
}

func (this *Remote) ImportModelFromCluster(clusterId int64, projectId int64, modelKey string, modelName string) (int64, error) {
	in := ImportModelFromClusterIn{clusterId, projectId, modelKey, modelName}
	var out ImportModelFromClusterOut
	err := this.Proc.Call("ImportModelFromCluster", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.ModelId, nil
}

func (this *Remote) CheckMojo(algo string) (bool, error) {
	in := CheckMojoIn{algo}
	var out CheckMojoOut
	err := this.Proc.Call("CheckMojo", &in, &out)
	if err != nil {
		return false, err
	}
	return out.CanMojo, nil
}

func (this *Remote) ImportModelPojo(modelId int64) error {
	in := ImportModelPojoIn{modelId}
	var out ImportModelPojoOut
	err := this.Proc.Call("ImportModelPojo", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) ImportModelMojo(modelId int64) error {
	in := ImportModelMojoIn{modelId}
	var out ImportModelMojoOut
	err := this.Proc.Call("ImportModelMojo", &in, &out)
	if err != nil {
		return err
	}
	return nil
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

func (this *Remote) CreateLabel(projectId int64, name string, description string) (int64, error) {
	in := CreateLabelIn{projectId, name, description}
	var out CreateLabelOut
	err := this.Proc.Call("CreateLabel", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.LabelId, nil
}

func (this *Remote) UpdateLabel(labelId int64, name string, description string) error {
	in := UpdateLabelIn{labelId, name, description}
	var out UpdateLabelOut
	err := this.Proc.Call("UpdateLabel", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) DeleteLabel(labelId int64) error {
	in := DeleteLabelIn{labelId}
	var out DeleteLabelOut
	err := this.Proc.Call("DeleteLabel", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) LinkLabelWithModel(labelId int64, modelId int64) error {
	in := LinkLabelWithModelIn{labelId, modelId}
	var out LinkLabelWithModelOut
	err := this.Proc.Call("LinkLabelWithModel", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) UnlinkLabelFromModel(labelId int64, modelId int64) error {
	in := UnlinkLabelFromModelIn{labelId, modelId}
	var out UnlinkLabelFromModelOut
	err := this.Proc.Call("UnlinkLabelFromModel", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetLabelsForProject(projectId int64) ([]*Label, error) {
	in := GetLabelsForProjectIn{projectId}
	var out GetLabelsForProjectOut
	err := this.Proc.Call("GetLabelsForProject", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Labels, nil
}

func (this *Remote) StartService(modelId int64, name string, packageName string) (int64, error) {
	in := StartServiceIn{modelId, name, packageName}
	var out StartServiceOut
	err := this.Proc.Call("StartService", &in, &out)
	if err != nil {
		return 0, err
	}
	return out.ServiceId, nil
}

func (this *Remote) StopService(serviceId int64) error {
	in := StopServiceIn{serviceId}
	var out StopServiceOut
	err := this.Proc.Call("StopService", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetService(serviceId int64) (*ScoringService, error) {
	in := GetServiceIn{serviceId}
	var out GetServiceOut
	err := this.Proc.Call("GetService", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Service, nil
}

func (this *Remote) GetServices(offset int64, limit int64) ([]*ScoringService, error) {
	in := GetServicesIn{offset, limit}
	var out GetServicesOut
	err := this.Proc.Call("GetServices", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Services, nil
}

func (this *Remote) GetServicesForProject(projectId int64, offset int64, limit int64) ([]*ScoringService, error) {
	in := GetServicesForProjectIn{projectId, offset, limit}
	var out GetServicesForProjectOut
	err := this.Proc.Call("GetServicesForProject", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Services, nil
}

func (this *Remote) GetServicesForModel(modelId int64, offset int64, limit int64) ([]*ScoringService, error) {
	in := GetServicesForModelIn{modelId, offset, limit}
	var out GetServicesForModelOut
	err := this.Proc.Call("GetServicesForModel", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Services, nil
}

func (this *Remote) DeleteService(serviceId int64) error {
	in := DeleteServiceIn{serviceId}
	var out DeleteServiceOut
	err := this.Proc.Call("DeleteService", &in, &out)
	if err != nil {
		return err
	}
	return nil
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

func (this *Remote) GetAllEntityTypes() ([]*EntityType, error) {
	in := GetAllEntityTypesIn{}
	var out GetAllEntityTypesOut
	err := this.Proc.Call("GetAllEntityTypes", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.EntityTypes, nil
}

func (this *Remote) GetAllPermissions() ([]*Permission, error) {
	in := GetAllPermissionsIn{}
	var out GetAllPermissionsOut
	err := this.Proc.Call("GetAllPermissions", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Permissions, nil
}

func (this *Remote) GetAllClusterTypes() ([]*ClusterType, error) {
	in := GetAllClusterTypesIn{}
	var out GetAllClusterTypesOut
	err := this.Proc.Call("GetAllClusterTypes", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.ClusterTypes, nil
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

func (this *Remote) GetRoleByName(name string) (*Role, error) {
	in := GetRoleByNameIn{name}
	var out GetRoleByNameOut
	err := this.Proc.Call("GetRoleByName", &in, &out)
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

func (this *Remote) LinkRoleWithPermissions(roleId int64, permissionIds []int64) error {
	in := LinkRoleWithPermissionsIn{roleId, permissionIds}
	var out LinkRoleWithPermissionsOut
	err := this.Proc.Call("LinkRoleWithPermissions", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) LinkRoleWithPermission(roleId int64, permissionId int64) error {
	in := LinkRoleWithPermissionIn{roleId, permissionId}
	var out LinkRoleWithPermissionOut
	err := this.Proc.Call("LinkRoleWithPermission", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) UnlinkRoleFromPermission(roleId int64, permissionId int64) error {
	in := UnlinkRoleFromPermissionIn{roleId, permissionId}
	var out UnlinkRoleFromPermissionOut
	err := this.Proc.Call("UnlinkRoleFromPermission", &in, &out)
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

func (this *Remote) GetWorkgroupByName(name string) (*Workgroup, error) {
	in := GetWorkgroupByNameIn{name}
	var out GetWorkgroupByNameOut
	err := this.Proc.Call("GetWorkgroupByName", &in, &out)
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

func (this *Remote) GetIdentitiesForRole(roleId int64) ([]*Identity, error) {
	in := GetIdentitiesForRoleIn{roleId}
	var out GetIdentitiesForRoleOut
	err := this.Proc.Call("GetIdentitiesForRole", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Identities, nil
}

func (this *Remote) GetIdentitiesForEntity(entityType int64, entityId int64) ([]*UserRole, error) {
	in := GetIdentitiesForEntityIn{entityType, entityId}
	var out GetIdentitiesForEntityOut
	err := this.Proc.Call("GetIdentitiesForEntity", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Users, nil
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

func (this *Remote) GetIdentityByName(name string) (*Identity, error) {
	in := GetIdentityByNameIn{name}
	var out GetIdentityByNameOut
	err := this.Proc.Call("GetIdentityByName", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Identity, nil
}

func (this *Remote) LinkIdentityWithWorkgroup(identityId int64, workgroupId int64) error {
	in := LinkIdentityWithWorkgroupIn{identityId, workgroupId}
	var out LinkIdentityWithWorkgroupOut
	err := this.Proc.Call("LinkIdentityWithWorkgroup", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) UnlinkIdentityFromWorkgroup(identityId int64, workgroupId int64) error {
	in := UnlinkIdentityFromWorkgroupIn{identityId, workgroupId}
	var out UnlinkIdentityFromWorkgroupOut
	err := this.Proc.Call("UnlinkIdentityFromWorkgroup", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) LinkIdentityWithRole(identityId int64, roleId int64) error {
	in := LinkIdentityWithRoleIn{identityId, roleId}
	var out LinkIdentityWithRoleOut
	err := this.Proc.Call("LinkIdentityWithRole", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) UnlinkIdentityFromRole(identityId int64, roleId int64) error {
	in := UnlinkIdentityFromRoleIn{identityId, roleId}
	var out UnlinkIdentityFromRoleOut
	err := this.Proc.Call("UnlinkIdentityFromRole", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) UpdateIdentity(identityId int64, password string) error {
	in := UpdateIdentityIn{identityId, password}
	var out UpdateIdentityOut
	err := this.Proc.Call("UpdateIdentity", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) ActivateIdentity(identityId int64) error {
	in := ActivateIdentityIn{identityId}
	var out ActivateIdentityOut
	err := this.Proc.Call("ActivateIdentity", &in, &out)
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

func (this *Remote) GetPrivileges(entityTypeId int64, entityId int64) ([]*EntityPrivilege, error) {
	in := GetPrivilegesIn{entityTypeId, entityId}
	var out GetPrivilegesOut
	err := this.Proc.Call("GetPrivileges", &in, &out)
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

func (this *Remote) GetHistory(entityTypeId int64, entityId int64, offset int64, limit int64) ([]*EntityHistory, error) {
	in := GetHistoryIn{entityTypeId, entityId, offset, limit}
	var out GetHistoryOut
	err := this.Proc.Call("GetHistory", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.History, nil
}

func (this *Remote) CreatePackage(projectId int64, name string) error {
	in := CreatePackageIn{projectId, name}
	var out CreatePackageOut
	err := this.Proc.Call("CreatePackage", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetPackages(projectId int64) ([]string, error) {
	in := GetPackagesIn{projectId}
	var out GetPackagesOut
	err := this.Proc.Call("GetPackages", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Packages, nil
}

func (this *Remote) GetPackageDirectories(projectId int64, packageName string, relativePath string) ([]string, error) {
	in := GetPackageDirectoriesIn{projectId, packageName, relativePath}
	var out GetPackageDirectoriesOut
	err := this.Proc.Call("GetPackageDirectories", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Directories, nil
}

func (this *Remote) GetPackageFiles(projectId int64, packageName string, relativePath string) ([]string, error) {
	in := GetPackageFilesIn{projectId, packageName, relativePath}
	var out GetPackageFilesOut
	err := this.Proc.Call("GetPackageFiles", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Files, nil
}

func (this *Remote) DeletePackage(projectId int64, name string) error {
	in := DeletePackageIn{projectId, name}
	var out DeletePackageOut
	err := this.Proc.Call("DeletePackage", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) DeletePackageDirectory(projectId int64, packageName string, relativePath string) error {
	in := DeletePackageDirectoryIn{projectId, packageName, relativePath}
	var out DeletePackageDirectoryOut
	err := this.Proc.Call("DeletePackageDirectory", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) DeletePackageFile(projectId int64, packageName string, relativePath string) error {
	in := DeletePackageFileIn{projectId, packageName, relativePath}
	var out DeletePackageFileOut
	err := this.Proc.Call("DeletePackageFile", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) SetAttributesForPackage(projectId int64, packageName string, attributes string) error {
	in := SetAttributesForPackageIn{projectId, packageName, attributes}
	var out SetAttributesForPackageOut
	err := this.Proc.Call("SetAttributesForPackage", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetAttributesForPackage(projectId int64, packageName string) (string, error) {
	in := GetAttributesForPackageIn{projectId, packageName}
	var out GetAttributesForPackageOut
	err := this.Proc.Call("GetAttributesForPackage", &in, &out)
	if err != nil {
		return "", err
	}
	return out.Attributes, nil
}

// --- Server Stub ---

type Impl struct {
	Service Service
	Az      az.Az
}

func (this *Impl) PingServer(r *http.Request, in *PingServerIn, out *PingServerOut) error {
	const name = "PingServer"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.PingServer(pz, in.Input)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Output = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetConfig(r *http.Request, in *GetConfigIn, out *GetConfigOut) error {
	const name = "GetConfig"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetConfig(pz)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Config = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) RegisterCluster(r *http.Request, in *RegisterClusterIn, out *RegisterClusterOut) error {
	const name = "RegisterCluster"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.RegisterCluster(pz, in.Address)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.ClusterId = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) UnregisterCluster(r *http.Request, in *UnregisterClusterIn, out *UnregisterClusterOut) error {
	const name = "UnregisterCluster"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.UnregisterCluster(pz, in.ClusterId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) StartClusterOnYarn(r *http.Request, in *StartClusterOnYarnIn, out *StartClusterOnYarnOut) error {
	const name = "StartClusterOnYarn"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.StartClusterOnYarn(pz, in.ClusterName, in.EngineId, in.Size, in.Memory, in.Secure, in.Keytab)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.ClusterId = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) StopClusterOnYarn(r *http.Request, in *StopClusterOnYarnIn, out *StopClusterOnYarnOut) error {
	const name = "StopClusterOnYarn"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.StopClusterOnYarn(pz, in.ClusterId, in.Keytab)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetCluster(r *http.Request, in *GetClusterIn, out *GetClusterOut) error {
	const name = "GetCluster"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetCluster(pz, in.ClusterId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Cluster = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetClusterOnYarn(r *http.Request, in *GetClusterOnYarnIn, out *GetClusterOnYarnOut) error {
	const name = "GetClusterOnYarn"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetClusterOnYarn(pz, in.ClusterId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Cluster = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetClusters(r *http.Request, in *GetClustersIn, out *GetClustersOut) error {
	const name = "GetClusters"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetClusters(pz, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Clusters = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetClusterStatus(r *http.Request, in *GetClusterStatusIn, out *GetClusterStatusOut) error {
	const name = "GetClusterStatus"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetClusterStatus(pz, in.ClusterId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.ClusterStatus = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) DeleteCluster(r *http.Request, in *DeleteClusterIn, out *DeleteClusterOut) error {
	const name = "DeleteCluster"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.DeleteCluster(pz, in.ClusterId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetJob(r *http.Request, in *GetJobIn, out *GetJobOut) error {
	const name = "GetJob"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetJob(pz, in.ClusterId, in.JobName)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Job = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetJobs(r *http.Request, in *GetJobsIn, out *GetJobsOut) error {
	const name = "GetJobs"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetJobs(pz, in.ClusterId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Jobs = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) CreateProject(r *http.Request, in *CreateProjectIn, out *CreateProjectOut) error {
	const name = "CreateProject"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.CreateProject(pz, in.Name, in.Description, in.ModelCategory)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.ProjectId = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetProjects(r *http.Request, in *GetProjectsIn, out *GetProjectsOut) error {
	const name = "GetProjects"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetProjects(pz, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Projects = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetProject(r *http.Request, in *GetProjectIn, out *GetProjectOut) error {
	const name = "GetProject"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetProject(pz, in.ProjectId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Project = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) DeleteProject(r *http.Request, in *DeleteProjectIn, out *DeleteProjectOut) error {
	const name = "DeleteProject"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.DeleteProject(pz, in.ProjectId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) CreateDatasource(r *http.Request, in *CreateDatasourceIn, out *CreateDatasourceOut) error {
	const name = "CreateDatasource"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.CreateDatasource(pz, in.ProjectId, in.Name, in.Description, in.Path)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.DatasourceId = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetDatasources(r *http.Request, in *GetDatasourcesIn, out *GetDatasourcesOut) error {
	const name = "GetDatasources"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetDatasources(pz, in.ProjectId, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Datasources = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetDatasource(r *http.Request, in *GetDatasourceIn, out *GetDatasourceOut) error {
	const name = "GetDatasource"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetDatasource(pz, in.DatasourceId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Datasource = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) UpdateDatasource(r *http.Request, in *UpdateDatasourceIn, out *UpdateDatasourceOut) error {
	const name = "UpdateDatasource"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.UpdateDatasource(pz, in.DatasourceId, in.Name, in.Description, in.Path)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) DeleteDatasource(r *http.Request, in *DeleteDatasourceIn, out *DeleteDatasourceOut) error {
	const name = "DeleteDatasource"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.DeleteDatasource(pz, in.DatasourceId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) CreateDataset(r *http.Request, in *CreateDatasetIn, out *CreateDatasetOut) error {
	const name = "CreateDataset"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.CreateDataset(pz, in.ClusterId, in.DatasourceId, in.Name, in.Description, in.ResponseColumnName)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.DatasetId = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetDatasets(r *http.Request, in *GetDatasetsIn, out *GetDatasetsOut) error {
	const name = "GetDatasets"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetDatasets(pz, in.DatasourceId, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Datasets = val0

	aux := make([]Dataset, len(out.Datasets))
	for i, val := range out.Datasets {
		aux[i] = *val
		aux[i].JSONProperties = "JSON DATA OMITTED..."
	}

	res, merr := json.Marshal(aux)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetDataset(r *http.Request, in *GetDatasetIn, out *GetDatasetOut) error {
	const name = "GetDataset"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetDataset(pz, in.DatasetId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Dataset = val0

	aux := *out.Dataset
	aux.JSONProperties = "JSON DATA OMITTED..."

	res, merr := json.Marshal(aux)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetDatasetsFromCluster(r *http.Request, in *GetDatasetsFromClusterIn, out *GetDatasetsFromClusterOut) error {
	const name = "GetDatasetsFromCluster"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetDatasetsFromCluster(pz, in.ClusterId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Dataset = val0

	aux := make([]Dataset, len(out.Dataset))
	for i, val := range out.Dataset {
		aux[i] = *val
		aux[i].JSONProperties = "JSON DATA OMITTED..."
	}

	res, merr := json.Marshal(aux)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) UpdateDataset(r *http.Request, in *UpdateDatasetIn, out *UpdateDatasetOut) error {
	const name = "UpdateDataset"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.UpdateDataset(pz, in.DatasetId, in.Name, in.Description, in.ResponseColumnName)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) SplitDataset(r *http.Request, in *SplitDatasetIn, out *SplitDatasetOut) error {
	const name = "SplitDataset"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.SplitDataset(pz, in.DatasetId, in.Ratio1, in.Ratio2)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.DatasetIds = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) DeleteDataset(r *http.Request, in *DeleteDatasetIn, out *DeleteDatasetOut) error {
	const name = "DeleteDataset"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.DeleteDataset(pz, in.DatasetId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) BuildModel(r *http.Request, in *BuildModelIn, out *BuildModelOut) error {
	const name = "BuildModel"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.BuildModel(pz, in.ClusterId, in.DatasetId, in.Algorithm)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.ModelId = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) BuildModelAuto(r *http.Request, in *BuildModelAutoIn, out *BuildModelAutoOut) error {
	const name = "BuildModelAuto"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.BuildModelAuto(pz, in.ClusterId, in.Dataset, in.TargetName, in.MaxRunTime)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Model = val0

	aux := *out.Model
	aux.JSONMetrics = "JSON DATA OMITTED..."

	res, merr := json.Marshal(aux)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetModel(r *http.Request, in *GetModelIn, out *GetModelOut) error {
	const name = "GetModel"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetModel(pz, in.ModelId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Model = val0

	aux := *out.Model
	aux.JSONMetrics = "JSON DATA OMITTED..."

	res, merr := json.Marshal(aux)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetModels(r *http.Request, in *GetModelsIn, out *GetModelsOut) error {
	const name = "GetModels"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetModels(pz, in.ProjectId, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Models = val0

	aux := make([]Model, len(out.Models))
	for i, val := range out.Models {
		aux[i] = *val
		aux[i].JSONMetrics = "JSON DATA OMITTED..."
	}

	res, merr := json.Marshal(aux)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetModelsFromCluster(r *http.Request, in *GetModelsFromClusterIn, out *GetModelsFromClusterOut) error {
	const name = "GetModelsFromCluster"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetModelsFromCluster(pz, in.ClusterId, in.FrameKey)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Models = val0

	aux := make([]Model, len(out.Models))
	for i, val := range out.Models {
		aux[i] = *val
		aux[i].JSONMetrics = "JSON DATA OMITTED..."
	}

	res, merr := json.Marshal(aux)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) FindModelsCount(r *http.Request, in *FindModelsCountIn, out *FindModelsCountOut) error {
	const name = "FindModelsCount"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.FindModelsCount(pz, in.ProjectId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Count = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetAllBinomialSortCriteria(r *http.Request, in *GetAllBinomialSortCriteriaIn, out *GetAllBinomialSortCriteriaOut) error {
	const name = "GetAllBinomialSortCriteria"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetAllBinomialSortCriteria(pz)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Criteria = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) FindModelsBinomial(r *http.Request, in *FindModelsBinomialIn, out *FindModelsBinomialOut) error {
	const name = "FindModelsBinomial"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.FindModelsBinomial(pz, in.ProjectId, in.NamePart, in.SortBy, in.Ascending, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Models = val0

	aux := make([]BinomialModel, len(out.Models))
	for i, val := range out.Models {
		aux[i] = *val
		aux[i].JSONMetrics = "JSON DATA OMITTED..."
	}

	res, merr := json.Marshal(aux)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetModelBinomial(r *http.Request, in *GetModelBinomialIn, out *GetModelBinomialOut) error {
	const name = "GetModelBinomial"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetModelBinomial(pz, in.ModelId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Model = val0

	aux := *out.Model
	aux.JSONMetrics = "JSON DATA OMITTED..."

	res, merr := json.Marshal(aux)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetAllMultinomialSortCriteria(r *http.Request, in *GetAllMultinomialSortCriteriaIn, out *GetAllMultinomialSortCriteriaOut) error {
	const name = "GetAllMultinomialSortCriteria"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetAllMultinomialSortCriteria(pz)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Criteria = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) FindModelsMultinomial(r *http.Request, in *FindModelsMultinomialIn, out *FindModelsMultinomialOut) error {
	const name = "FindModelsMultinomial"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.FindModelsMultinomial(pz, in.ProjectId, in.NamePart, in.SortBy, in.Ascending, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Models = val0

	aux := make([]MultinomialModel, len(out.Models))
	for i, val := range out.Models {
		aux[i] = *val
		aux[i].JSONMetrics = "JSON DATA OMITTED..."
	}

	res, merr := json.Marshal(aux)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetModelMultinomial(r *http.Request, in *GetModelMultinomialIn, out *GetModelMultinomialOut) error {
	const name = "GetModelMultinomial"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetModelMultinomial(pz, in.ModelId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Model = val0

	aux := *out.Model
	aux.JSONMetrics = "JSON DATA OMITTED..."

	res, merr := json.Marshal(aux)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetAllRegressionSortCriteria(r *http.Request, in *GetAllRegressionSortCriteriaIn, out *GetAllRegressionSortCriteriaOut) error {
	const name = "GetAllRegressionSortCriteria"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetAllRegressionSortCriteria(pz)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Criteria = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) FindModelsRegression(r *http.Request, in *FindModelsRegressionIn, out *FindModelsRegressionOut) error {
	const name = "FindModelsRegression"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.FindModelsRegression(pz, in.ProjectId, in.NamePart, in.SortBy, in.Ascending, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Models = val0

	aux := make([]RegressionModel, len(out.Models))
	for i, val := range out.Models {
		aux[i] = *val
		aux[i].JSONMetrics = "JSON DATA OMITTED..."
	}

	res, merr := json.Marshal(aux)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetModelRegression(r *http.Request, in *GetModelRegressionIn, out *GetModelRegressionOut) error {
	const name = "GetModelRegression"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetModelRegression(pz, in.ModelId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Model = val0

	aux := *out.Model
	aux.JSONMetrics = "JSON DATA OMITTED..."

	res, merr := json.Marshal(aux)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) ImportModelFromCluster(r *http.Request, in *ImportModelFromClusterIn, out *ImportModelFromClusterOut) error {
	const name = "ImportModelFromCluster"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.ImportModelFromCluster(pz, in.ClusterId, in.ProjectId, in.ModelKey, in.ModelName)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.ModelId = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) CheckMojo(r *http.Request, in *CheckMojoIn, out *CheckMojoOut) error {
	const name = "CheckMojo"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.CheckMojo(pz, in.Algo)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.CanMojo = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) ImportModelPojo(r *http.Request, in *ImportModelPojoIn, out *ImportModelPojoOut) error {
	const name = "ImportModelPojo"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.ImportModelPojo(pz, in.ModelId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) ImportModelMojo(r *http.Request, in *ImportModelMojoIn, out *ImportModelMojoOut) error {
	const name = "ImportModelMojo"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.ImportModelMojo(pz, in.ModelId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) DeleteModel(r *http.Request, in *DeleteModelIn, out *DeleteModelOut) error {
	const name = "DeleteModel"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.DeleteModel(pz, in.ModelId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) CreateLabel(r *http.Request, in *CreateLabelIn, out *CreateLabelOut) error {
	const name = "CreateLabel"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.CreateLabel(pz, in.ProjectId, in.Name, in.Description)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.LabelId = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) UpdateLabel(r *http.Request, in *UpdateLabelIn, out *UpdateLabelOut) error {
	const name = "UpdateLabel"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.UpdateLabel(pz, in.LabelId, in.Name, in.Description)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) DeleteLabel(r *http.Request, in *DeleteLabelIn, out *DeleteLabelOut) error {
	const name = "DeleteLabel"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.DeleteLabel(pz, in.LabelId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) LinkLabelWithModel(r *http.Request, in *LinkLabelWithModelIn, out *LinkLabelWithModelOut) error {
	const name = "LinkLabelWithModel"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.LinkLabelWithModel(pz, in.LabelId, in.ModelId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) UnlinkLabelFromModel(r *http.Request, in *UnlinkLabelFromModelIn, out *UnlinkLabelFromModelOut) error {
	const name = "UnlinkLabelFromModel"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.UnlinkLabelFromModel(pz, in.LabelId, in.ModelId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetLabelsForProject(r *http.Request, in *GetLabelsForProjectIn, out *GetLabelsForProjectOut) error {
	const name = "GetLabelsForProject"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetLabelsForProject(pz, in.ProjectId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Labels = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) StartService(r *http.Request, in *StartServiceIn, out *StartServiceOut) error {
	const name = "StartService"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.StartService(pz, in.ModelId, in.Name, in.PackageName)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.ServiceId = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) StopService(r *http.Request, in *StopServiceIn, out *StopServiceOut) error {
	const name = "StopService"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.StopService(pz, in.ServiceId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetService(r *http.Request, in *GetServiceIn, out *GetServiceOut) error {
	const name = "GetService"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetService(pz, in.ServiceId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Service = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetServices(r *http.Request, in *GetServicesIn, out *GetServicesOut) error {
	const name = "GetServices"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetServices(pz, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Services = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetServicesForProject(r *http.Request, in *GetServicesForProjectIn, out *GetServicesForProjectOut) error {
	const name = "GetServicesForProject"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetServicesForProject(pz, in.ProjectId, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Services = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetServicesForModel(r *http.Request, in *GetServicesForModelIn, out *GetServicesForModelOut) error {
	const name = "GetServicesForModel"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetServicesForModel(pz, in.ModelId, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Services = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) DeleteService(r *http.Request, in *DeleteServiceIn, out *DeleteServiceOut) error {
	const name = "DeleteService"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.DeleteService(pz, in.ServiceId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetEngine(r *http.Request, in *GetEngineIn, out *GetEngineOut) error {
	const name = "GetEngine"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetEngine(pz, in.EngineId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Engine = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetEngines(r *http.Request, in *GetEnginesIn, out *GetEnginesOut) error {
	const name = "GetEngines"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetEngines(pz)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Engines = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) DeleteEngine(r *http.Request, in *DeleteEngineIn, out *DeleteEngineOut) error {
	const name = "DeleteEngine"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.DeleteEngine(pz, in.EngineId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetAllEntityTypes(r *http.Request, in *GetAllEntityTypesIn, out *GetAllEntityTypesOut) error {
	const name = "GetAllEntityTypes"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetAllEntityTypes(pz)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.EntityTypes = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetAllPermissions(r *http.Request, in *GetAllPermissionsIn, out *GetAllPermissionsOut) error {
	const name = "GetAllPermissions"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetAllPermissions(pz)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Permissions = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetAllClusterTypes(r *http.Request, in *GetAllClusterTypesIn, out *GetAllClusterTypesOut) error {
	const name = "GetAllClusterTypes"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetAllClusterTypes(pz)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.ClusterTypes = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetPermissionsForRole(r *http.Request, in *GetPermissionsForRoleIn, out *GetPermissionsForRoleOut) error {
	const name = "GetPermissionsForRole"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetPermissionsForRole(pz, in.RoleId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Permissions = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetPermissionsForIdentity(r *http.Request, in *GetPermissionsForIdentityIn, out *GetPermissionsForIdentityOut) error {
	const name = "GetPermissionsForIdentity"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetPermissionsForIdentity(pz, in.IdentityId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Permissions = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) CreateRole(r *http.Request, in *CreateRoleIn, out *CreateRoleOut) error {
	const name = "CreateRole"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.CreateRole(pz, in.Name, in.Description)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.RoleId = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetRoles(r *http.Request, in *GetRolesIn, out *GetRolesOut) error {
	const name = "GetRoles"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetRoles(pz, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Roles = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetRolesForIdentity(r *http.Request, in *GetRolesForIdentityIn, out *GetRolesForIdentityOut) error {
	const name = "GetRolesForIdentity"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetRolesForIdentity(pz, in.IdentityId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Roles = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetRole(r *http.Request, in *GetRoleIn, out *GetRoleOut) error {
	const name = "GetRole"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetRole(pz, in.RoleId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Role = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetRoleByName(r *http.Request, in *GetRoleByNameIn, out *GetRoleByNameOut) error {
	const name = "GetRoleByName"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetRoleByName(pz, in.Name)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Role = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) UpdateRole(r *http.Request, in *UpdateRoleIn, out *UpdateRoleOut) error {
	const name = "UpdateRole"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.UpdateRole(pz, in.RoleId, in.Name, in.Description)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) LinkRoleWithPermissions(r *http.Request, in *LinkRoleWithPermissionsIn, out *LinkRoleWithPermissionsOut) error {
	const name = "LinkRoleWithPermissions"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.LinkRoleWithPermissions(pz, in.RoleId, in.PermissionIds)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) LinkRoleWithPermission(r *http.Request, in *LinkRoleWithPermissionIn, out *LinkRoleWithPermissionOut) error {
	const name = "LinkRoleWithPermission"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.LinkRoleWithPermission(pz, in.RoleId, in.PermissionId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) UnlinkRoleFromPermission(r *http.Request, in *UnlinkRoleFromPermissionIn, out *UnlinkRoleFromPermissionOut) error {
	const name = "UnlinkRoleFromPermission"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.UnlinkRoleFromPermission(pz, in.RoleId, in.PermissionId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) DeleteRole(r *http.Request, in *DeleteRoleIn, out *DeleteRoleOut) error {
	const name = "DeleteRole"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.DeleteRole(pz, in.RoleId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) CreateWorkgroup(r *http.Request, in *CreateWorkgroupIn, out *CreateWorkgroupOut) error {
	const name = "CreateWorkgroup"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.CreateWorkgroup(pz, in.Name, in.Description)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.WorkgroupId = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetWorkgroups(r *http.Request, in *GetWorkgroupsIn, out *GetWorkgroupsOut) error {
	const name = "GetWorkgroups"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetWorkgroups(pz, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Workgroups = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetWorkgroupsForIdentity(r *http.Request, in *GetWorkgroupsForIdentityIn, out *GetWorkgroupsForIdentityOut) error {
	const name = "GetWorkgroupsForIdentity"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetWorkgroupsForIdentity(pz, in.IdentityId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Workgroups = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetWorkgroup(r *http.Request, in *GetWorkgroupIn, out *GetWorkgroupOut) error {
	const name = "GetWorkgroup"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetWorkgroup(pz, in.WorkgroupId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Workgroup = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetWorkgroupByName(r *http.Request, in *GetWorkgroupByNameIn, out *GetWorkgroupByNameOut) error {
	const name = "GetWorkgroupByName"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetWorkgroupByName(pz, in.Name)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Workgroup = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) UpdateWorkgroup(r *http.Request, in *UpdateWorkgroupIn, out *UpdateWorkgroupOut) error {
	const name = "UpdateWorkgroup"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.UpdateWorkgroup(pz, in.WorkgroupId, in.Name, in.Description)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) DeleteWorkgroup(r *http.Request, in *DeleteWorkgroupIn, out *DeleteWorkgroupOut) error {
	const name = "DeleteWorkgroup"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.DeleteWorkgroup(pz, in.WorkgroupId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) CreateIdentity(r *http.Request, in *CreateIdentityIn, out *CreateIdentityOut) error {
	const name = "CreateIdentity"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.CreateIdentity(pz, in.Name, in.Password)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.IdentityId = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetIdentities(r *http.Request, in *GetIdentitiesIn, out *GetIdentitiesOut) error {
	const name = "GetIdentities"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetIdentities(pz, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Identities = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetIdentitiesForWorkgroup(r *http.Request, in *GetIdentitiesForWorkgroupIn, out *GetIdentitiesForWorkgroupOut) error {
	const name = "GetIdentitiesForWorkgroup"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetIdentitiesForWorkgroup(pz, in.WorkgroupId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Identities = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetIdentitiesForRole(r *http.Request, in *GetIdentitiesForRoleIn, out *GetIdentitiesForRoleOut) error {
	const name = "GetIdentitiesForRole"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetIdentitiesForRole(pz, in.RoleId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Identities = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetIdentitiesForEntity(r *http.Request, in *GetIdentitiesForEntityIn, out *GetIdentitiesForEntityOut) error {
	const name = "GetIdentitiesForEntity"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetIdentitiesForEntity(pz, in.EntityType, in.EntityId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Users = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetIdentity(r *http.Request, in *GetIdentityIn, out *GetIdentityOut) error {
	const name = "GetIdentity"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetIdentity(pz, in.IdentityId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Identity = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetIdentityByName(r *http.Request, in *GetIdentityByNameIn, out *GetIdentityByNameOut) error {
	const name = "GetIdentityByName"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetIdentityByName(pz, in.Name)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Identity = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) LinkIdentityWithWorkgroup(r *http.Request, in *LinkIdentityWithWorkgroupIn, out *LinkIdentityWithWorkgroupOut) error {
	const name = "LinkIdentityWithWorkgroup"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.LinkIdentityWithWorkgroup(pz, in.IdentityId, in.WorkgroupId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) UnlinkIdentityFromWorkgroup(r *http.Request, in *UnlinkIdentityFromWorkgroupIn, out *UnlinkIdentityFromWorkgroupOut) error {
	const name = "UnlinkIdentityFromWorkgroup"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.UnlinkIdentityFromWorkgroup(pz, in.IdentityId, in.WorkgroupId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) LinkIdentityWithRole(r *http.Request, in *LinkIdentityWithRoleIn, out *LinkIdentityWithRoleOut) error {
	const name = "LinkIdentityWithRole"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.LinkIdentityWithRole(pz, in.IdentityId, in.RoleId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) UnlinkIdentityFromRole(r *http.Request, in *UnlinkIdentityFromRoleIn, out *UnlinkIdentityFromRoleOut) error {
	const name = "UnlinkIdentityFromRole"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.UnlinkIdentityFromRole(pz, in.IdentityId, in.RoleId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) UpdateIdentity(r *http.Request, in *UpdateIdentityIn, out *UpdateIdentityOut) error {
	const name = "UpdateIdentity"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.UpdateIdentity(pz, in.IdentityId, in.Password)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) ActivateIdentity(r *http.Request, in *ActivateIdentityIn, out *ActivateIdentityOut) error {
	const name = "ActivateIdentity"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.ActivateIdentity(pz, in.IdentityId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) DeactivateIdentity(r *http.Request, in *DeactivateIdentityIn, out *DeactivateIdentityOut) error {
	const name = "DeactivateIdentity"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.DeactivateIdentity(pz, in.IdentityId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) ShareEntity(r *http.Request, in *ShareEntityIn, out *ShareEntityOut) error {
	const name = "ShareEntity"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.ShareEntity(pz, in.Kind, in.WorkgroupId, in.EntityTypeId, in.EntityId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetPrivileges(r *http.Request, in *GetPrivilegesIn, out *GetPrivilegesOut) error {
	const name = "GetPrivileges"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetPrivileges(pz, in.EntityTypeId, in.EntityId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Privileges = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) UnshareEntity(r *http.Request, in *UnshareEntityIn, out *UnshareEntityOut) error {
	const name = "UnshareEntity"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.UnshareEntity(pz, in.Kind, in.WorkgroupId, in.EntityTypeId, in.EntityId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetHistory(r *http.Request, in *GetHistoryIn, out *GetHistoryOut) error {
	const name = "GetHistory"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetHistory(pz, in.EntityTypeId, in.EntityId, in.Offset, in.Limit)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.History = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) CreatePackage(r *http.Request, in *CreatePackageIn, out *CreatePackageOut) error {
	const name = "CreatePackage"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.CreatePackage(pz, in.ProjectId, in.Name)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetPackages(r *http.Request, in *GetPackagesIn, out *GetPackagesOut) error {
	const name = "GetPackages"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetPackages(pz, in.ProjectId)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Packages = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetPackageDirectories(r *http.Request, in *GetPackageDirectoriesIn, out *GetPackageDirectoriesOut) error {
	const name = "GetPackageDirectories"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetPackageDirectories(pz, in.ProjectId, in.PackageName, in.RelativePath)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Directories = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetPackageFiles(r *http.Request, in *GetPackageFilesIn, out *GetPackageFilesOut) error {
	const name = "GetPackageFiles"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetPackageFiles(pz, in.ProjectId, in.PackageName, in.RelativePath)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Files = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) DeletePackage(r *http.Request, in *DeletePackageIn, out *DeletePackageOut) error {
	const name = "DeletePackage"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.DeletePackage(pz, in.ProjectId, in.Name)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) DeletePackageDirectory(r *http.Request, in *DeletePackageDirectoryIn, out *DeletePackageDirectoryOut) error {
	const name = "DeletePackageDirectory"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.DeletePackageDirectory(pz, in.ProjectId, in.PackageName, in.RelativePath)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) DeletePackageFile(r *http.Request, in *DeletePackageFileIn, out *DeletePackageFileOut) error {
	const name = "DeletePackageFile"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.DeletePackageFile(pz, in.ProjectId, in.PackageName, in.RelativePath)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) SetAttributesForPackage(r *http.Request, in *SetAttributesForPackageIn, out *SetAttributesForPackageOut) error {
	const name = "SetAttributesForPackage"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	err := this.Service.SetAttributesForPackage(pz, in.ProjectId, in.PackageName, in.Attributes)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}

func (this *Impl) GetAttributesForPackage(r *http.Request, in *GetAttributesForPackageIn, out *GetAttributesForPackageOut) error {
	const name = "GetAttributesForPackage"

	guid := xid.New().String()

	pz, azerr := this.Az.Identify(r)
	if azerr != nil {
		return azerr
	}

	req, merr := json.Marshal(in)
	if merr != nil {
		log.Println(guid, "REQ", pz, name, merr)
	} else {
		log.Println(guid, "REQ", pz, name, string(req))
	}

	val0, err := this.Service.GetAttributesForPackage(pz, in.ProjectId, in.PackageName)
	if err != nil {
		log.Println(guid, "ERR", pz, name, err)
		return err
	}

	out.Attributes = val0

	res, merr := json.Marshal(out)
	if merr != nil {
		log.Println(guid, "RES", pz, name, merr)
	} else {
		log.Println(guid, "RES", pz, name, string(res))
	}

	return nil
}
