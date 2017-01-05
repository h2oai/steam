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




// --- Types ---
import * as Proxy from './xhr';

export interface BinomialModel {

  id: number

  training_dataset_id: number

  validation_dataset_id: number

  name: string

  cluster_name: string

  model_key: string

  algorithm: string

  model_category: string

  dataset_name: string

  response_column_name: string

  logical_name: string

  location: string

  model_object_type: string

  max_runtime: number

  json_metrics: string

  created_at: number

  label_id: number

  label_name: string

  mse: number

  r_squared: number

  logloss: number

  auc: number

  gini: number

}

export interface Cluster {

  id: number

  name: string

  type_id: number

  detail_id: number

  address: string

  state: string

  created_at: number

}

export interface ClusterStatus {

  version: string

  status: string

  max_memory: string

  total_cpu_count: number

  total_allowed_cpu_count: number

}

export interface ClusterType {

  id: number

  name: string

}

export interface Config {

  kerberos_enabled: boolean

  cluster_proxy_address: string

}

export interface Dataset {

  id: number

  datasource_id: number

  name: string

  description: string

  frame_name: string

  response_column_name: string

  json_properties: string

  created_at: number

}

export interface Datasource {

  id: number

  project_id: number

  name: string

  description: string

  kind: string

  configuration: string

  created_at: number

}

export interface Engine {

  id: number

  name: string

  location: string

  created_at: number

}

export interface EntityHistory {

  identity_id: number

  action: string

  description: string

  created_at: number

}

export interface EntityPrivilege {

  kind: string

  workgroup_id: number

  workgroup_name: string

  workgroup_description: string

}

export interface EntityType {

  id: number

  name: string

}

export interface Identity {

  id: number

  name: string

  is_active: boolean

  last_login: number

  created: number

}

export interface Job {

  name: string

  cluster_name: string

  description: string

  progress: string

  started_at: number

  completed_at: number

}

export interface Label {

  id: number

  project_id: number

  model_id: number

  name: string

  description: string

  created_at: number

}

export interface Model {

  id: number

  training_dataset_id: number

  validation_dataset_id: number

  name: string

  cluster_name: string

  model_key: string

  algorithm: string

  model_category: string

  dataset_name: string

  response_column_name: string

  logical_name: string

  location: string

  model_object_type: string

  max_runtime: number

  json_metrics: string

  created_at: number

  label_id: number

  label_name: string

}

export interface MultinomialModel {

  id: number

  training_dataset_id: number

  validation_dataset_id: number

  name: string

  cluster_name: string

  model_key: string

  algorithm: string

  model_category: string

  dataset_name: string

  response_column_name: string

  logical_name: string

  location: string

  model_object_type: string

  max_runtime: number

  json_metrics: string

  created_at: number

  label_id: number

  label_name: string

  mse: number

  r_squared: number

  logloss: number

}

export interface Permission {

  id: number

  code: string

  description: string

}

export interface Project {

  id: number

  name: string

  description: string

  model_category: string

  created_at: number

}

export interface RegressionModel {

  id: number

  training_dataset_id: number

  validation_dataset_id: number

  name: string

  cluster_name: string

  model_key: string

  algorithm: string

  model_category: string

  dataset_name: string

  response_column_name: string

  logical_name: string

  location: string

  model_object_type: string

  max_runtime: number

  json_metrics: string

  created_at: number

  label_id: number

  label_name: string

  mse: number

  r_squared: number

  mean_residual_deviance: number

}

export interface Role {

  id: number

  name: string

  description: string

  created: number

}

export interface ScoringService {

  id: number

  model_id: number

  name: string

  address: string

  port: number

  process_id: number

  state: string

  created_at: number

}

export interface UserRole {

  kind: string

  identity_id: number

  identity_name: string

  role_id: number

  role_name: string

}

export interface Workgroup {

  id: number

  name: string

  description: string

  created: number

}

export interface YarnCluster {

  id: number

  engine_id: number

  size: number

  application_id: string

  memory: string

  username: string

}


// --- Contract ---


export interface Service {

  // Ping the Steam server
  pingServer: (input: string, go: (error: Error, output: string) => void) => void

  // No description available
  getConfig: (go: (error: Error, config: Config) => void) => void

  // Connect to a cluster
  registerCluster: (address: string, go: (error: Error, clusterId: number) => void) => void

  // Disconnect from a cluster
  unregisterCluster: (clusterId: number, go: (error: Error) => void) => void

  // Start a cluster using Yarn
  startClusterOnYarn: (clusterName: string, engineId: number, size: number, memory: string, secure: boolean, keytab: string, go: (error: Error, clusterId: number) => void) => void

  // Stop a cluster using Yarn
  stopClusterOnYarn: (clusterId: number, keytab: string, go: (error: Error) => void) => void

  // Get cluster details
  getCluster: (clusterId: number, go: (error: Error, cluster: Cluster) => void) => void

  // Get cluster details (Yarn only)
  getClusterOnYarn: (clusterId: number, go: (error: Error, cluster: YarnCluster) => void) => void

  // List clusters
  getClusters: (offset: number, limit: number, go: (error: Error, clusters: Cluster[]) => void) => void

  // Get cluster status
  getClusterStatus: (clusterId: number, go: (error: Error, clusterStatus: ClusterStatus) => void) => void

  // Delete a cluster
  deleteCluster: (clusterId: number, go: (error: Error) => void) => void

  // Get job details
  getJob: (clusterId: number, jobName: string, go: (error: Error, job: Job) => void) => void

  // List jobs
  getJobs: (clusterId: number, go: (error: Error, jobs: Job[]) => void) => void

  // Create a project
  createProject: (name: string, description: string, modelCategory: string, go: (error: Error, projectId: number) => void) => void

  // List projects
  getProjects: (offset: number, limit: number, go: (error: Error, projects: Project[]) => void) => void

  // Get project details
  getProject: (projectId: number, go: (error: Error, project: Project) => void) => void

  // Delete a project
  deleteProject: (projectId: number, go: (error: Error) => void) => void

  // Create a datasource
  createDatasource: (projectId: number, name: string, description: string, path: string, go: (error: Error, datasourceId: number) => void) => void

  // List datasources
  getDatasources: (projectId: number, offset: number, limit: number, go: (error: Error, datasources: Datasource[]) => void) => void

  // Get datasource details
  getDatasource: (datasourceId: number, go: (error: Error, datasource: Datasource) => void) => void

  // Update a datasource
  updateDatasource: (datasourceId: number, name: string, description: string, path: string, go: (error: Error) => void) => void

  // Delete a datasource
  deleteDatasource: (datasourceId: number, go: (error: Error) => void) => void

  // Create a dataset
  createDataset: (clusterId: number, datasourceId: number, name: string, description: string, responseColumnName: string, go: (error: Error, datasetId: number) => void) => void

  // List datasets
  getDatasets: (datasourceId: number, offset: number, limit: number, go: (error: Error, datasets: Dataset[]) => void) => void

  // Get dataset details
  getDataset: (datasetId: number, go: (error: Error, dataset: Dataset) => void) => void

  // Get a list of datasets on a cluster
  getDatasetsFromCluster: (clusterId: number, go: (error: Error, dataset: Dataset[]) => void) => void

  // Update a dataset
  updateDataset: (datasetId: number, name: string, description: string, responseColumnName: string, go: (error: Error) => void) => void

  // Split a dataset
  splitDataset: (datasetId: number, ratio1: number, ratio2: number, go: (error: Error, datasetIds: number[]) => void) => void

  // Delete a dataset
  deleteDataset: (datasetId: number, go: (error: Error) => void) => void

  // Build a model
  buildModel: (clusterId: number, datasetId: number, algorithm: string, go: (error: Error, modelId: number) => void) => void

  // Build an AutoML model
  buildModelAuto: (clusterId: number, dataset: string, targetName: string, maxRunTime: number, go: (error: Error, model: Model) => void) => void

  // Get model details
  getModel: (modelId: number, go: (error: Error, model: Model) => void) => void

  // List models
  getModels: (projectId: number, offset: number, limit: number, go: (error: Error, models: Model[]) => void) => void

  // List models from a cluster
  getModelsFromCluster: (clusterId: number, frameKey: string, go: (error: Error, models: Model[]) => void) => void

  // Get a count models in a project
  findModelsCount: (projectId: number, go: (error: Error, count: number) => void) => void

  // List sort criteria for a binomial models
  getAllBinomialSortCriteria: (go: (error: Error, criteria: string[]) => void) => void

  // List binomial models
  findModelsBinomial: (projectId: number, namePart: string, sortBy: string, ascending: boolean, offset: number, limit: number, go: (error: Error, models: BinomialModel[]) => void) => void

  // View a binomial model
  getModelBinomial: (modelId: number, go: (error: Error, model: BinomialModel) => void) => void

  // List sort criteria for a multinomial models
  getAllMultinomialSortCriteria: (go: (error: Error, criteria: string[]) => void) => void

  // List multinomial models
  findModelsMultinomial: (projectId: number, namePart: string, sortBy: string, ascending: boolean, offset: number, limit: number, go: (error: Error, models: MultinomialModel[]) => void) => void

  // View a binomial model
  getModelMultinomial: (modelId: number, go: (error: Error, model: MultinomialModel) => void) => void

  // List sort criteria for a regression models
  getAllRegressionSortCriteria: (go: (error: Error, criteria: string[]) => void) => void

  // List regression models
  findModelsRegression: (projectId: number, namePart: string, sortBy: string, ascending: boolean, offset: number, limit: number, go: (error: Error, models: RegressionModel[]) => void) => void

  // View a binomial model
  getModelRegression: (modelId: number, go: (error: Error, model: RegressionModel) => void) => void

  // Import models from a cluster
  importModelFromCluster: (clusterId: number, projectId: number, modelKey: string, modelName: string, go: (error: Error, modelId: number) => void) => void

  // Check if a model category can generate MOJOs
  checkMojo: (algo: string, go: (error: Error, canMojo: boolean) => void) => void

  // Import a model's POJO from a cluster
  importModelPojo: (modelId: number, go: (error: Error) => void) => void

  // Import a model's MOJO from a cluster
  importModelMojo: (modelId: number, go: (error: Error) => void) => void

  // Delete a model
  deleteModel: (modelId: number, go: (error: Error) => void) => void

  // Create a label
  createLabel: (projectId: number, name: string, description: string, go: (error: Error, labelId: number) => void) => void

  // Update a label
  updateLabel: (labelId: number, name: string, description: string, go: (error: Error) => void) => void

  // Delete a label
  deleteLabel: (labelId: number, go: (error: Error) => void) => void

  // Label a model
  linkLabelWithModel: (labelId: number, modelId: number, go: (error: Error) => void) => void

  // Remove a label from a model
  unlinkLabelFromModel: (labelId: number, modelId: number, go: (error: Error) => void) => void

  // List labels for a project, with corresponding models, if any
  getLabelsForProject: (projectId: number, go: (error: Error, labels: Label[]) => void) => void

  // Start a service
  startService: (modelId: number, name: string, packageName: string, go: (error: Error, serviceId: number) => void) => void

  // Stop a service
  stopService: (serviceId: number, go: (error: Error) => void) => void

  // Get service details
  getService: (serviceId: number, go: (error: Error, service: ScoringService) => void) => void

  // List all services
  getServices: (offset: number, limit: number, go: (error: Error, services: ScoringService[]) => void) => void

  // List services for a project
  getServicesForProject: (projectId: number, offset: number, limit: number, go: (error: Error, services: ScoringService[]) => void) => void

  // List services for a model
  getServicesForModel: (modelId: number, offset: number, limit: number, go: (error: Error, services: ScoringService[]) => void) => void

  // Delete a service
  deleteService: (serviceId: number, go: (error: Error) => void) => void

  // Get engine details
  getEngine: (engineId: number, go: (error: Error, engine: Engine) => void) => void

  // List engines
  getEngines: (go: (error: Error, engines: Engine[]) => void) => void

  // Delete an engine
  deleteEngine: (engineId: number, go: (error: Error) => void) => void

  // List all entity types
  getAllEntityTypes: (go: (error: Error, entityTypes: EntityType[]) => void) => void

  // List all permissions
  getAllPermissions: (go: (error: Error, permissions: Permission[]) => void) => void

  // List all cluster types
  getAllClusterTypes: (go: (error: Error, clusterTypes: ClusterType[]) => void) => void

  // List permissions for a role
  getPermissionsForRole: (roleId: number, go: (error: Error, permissions: Permission[]) => void) => void

  // List permissions for an identity
  getPermissionsForIdentity: (identityId: number, go: (error: Error, permissions: Permission[]) => void) => void

  // Create a role
  createRole: (name: string, description: string, go: (error: Error, roleId: number) => void) => void

  // List roles
  getRoles: (offset: number, limit: number, go: (error: Error, roles: Role[]) => void) => void

  // List roles for an identity
  getRolesForIdentity: (identityId: number, go: (error: Error, roles: Role[]) => void) => void

  // Get role details
  getRole: (roleId: number, go: (error: Error, role: Role) => void) => void

  // Get role details by name
  getRoleByName: (name: string, go: (error: Error, role: Role) => void) => void

  // Update a role
  updateRole: (roleId: number, name: string, description: string, go: (error: Error) => void) => void

  // Link a role with permissions
  linkRoleWithPermissions: (roleId: number, permissionIds: number[], go: (error: Error) => void) => void

  // Link a role with a permission
  linkRoleWithPermission: (roleId: number, permissionId: number, go: (error: Error) => void) => void

  // Unlink a role from a permission
  unlinkRoleFromPermission: (roleId: number, permissionId: number, go: (error: Error) => void) => void

  // Delete a role
  deleteRole: (roleId: number, go: (error: Error) => void) => void

  // Create a workgroup
  createWorkgroup: (name: string, description: string, go: (error: Error, workgroupId: number) => void) => void

  // List workgroups
  getWorkgroups: (offset: number, limit: number, go: (error: Error, workgroups: Workgroup[]) => void) => void

  // List workgroups for an identity
  getWorkgroupsForIdentity: (identityId: number, go: (error: Error, workgroups: Workgroup[]) => void) => void

  // Get workgroup details
  getWorkgroup: (workgroupId: number, go: (error: Error, workgroup: Workgroup) => void) => void

  // Get workgroup details by name
  getWorkgroupByName: (name: string, go: (error: Error, workgroup: Workgroup) => void) => void

  // Update a workgroup
  updateWorkgroup: (workgroupId: number, name: string, description: string, go: (error: Error) => void) => void

  // Delete a workgroup
  deleteWorkgroup: (workgroupId: number, go: (error: Error) => void) => void

  // Create an identity
  createIdentity: (name: string, password: string, go: (error: Error, identityId: number) => void) => void

  // List identities
  getIdentities: (offset: number, limit: number, go: (error: Error, identities: Identity[]) => void) => void

  // List identities for a workgroup
  getIdentitiesForWorkgroup: (workgroupId: number, go: (error: Error, identities: Identity[]) => void) => void

  // List identities for a role
  getIdentitiesForRole: (roleId: number, go: (error: Error, identities: Identity[]) => void) => void

  // Get a list of identities and roles with access to an entity
  getIdentitiesForEntity: (entityType: number, entityId: number, go: (error: Error, users: UserRole[]) => void) => void

  // Get identity details
  getIdentity: (identityId: number, go: (error: Error, identity: Identity) => void) => void

  // Get identity details by name
  getIdentityByName: (name: string, go: (error: Error, identity: Identity) => void) => void

  // Link an identity with a workgroup
  linkIdentityWithWorkgroup: (identityId: number, workgroupId: number, go: (error: Error) => void) => void

  // Unlink an identity from a workgroup
  unlinkIdentityFromWorkgroup: (identityId: number, workgroupId: number, go: (error: Error) => void) => void

  // Link an identity with a role
  linkIdentityWithRole: (identityId: number, roleId: number, go: (error: Error) => void) => void

  // Unlink an identity from a role
  unlinkIdentityFromRole: (identityId: number, roleId: number, go: (error: Error) => void) => void

  // Update an identity
  updateIdentity: (identityId: number, password: string, go: (error: Error) => void) => void

  // Activate an identity
  activateIdentity: (identityId: number, go: (error: Error) => void) => void

  // Deactivate an identity
  deactivateIdentity: (identityId: number, go: (error: Error) => void) => void

  // Share an entity with a workgroup
  shareEntity: (kind: string, workgroupId: number, entityTypeId: number, entityId: number, go: (error: Error) => void) => void

  // List privileges for an entity
  getPrivileges: (entityTypeId: number, entityId: number, go: (error: Error, privileges: EntityPrivilege[]) => void) => void

  // Unshare an entity
  unshareEntity: (kind: string, workgroupId: number, entityTypeId: number, entityId: number, go: (error: Error) => void) => void

  // List audit trail records for an entity
  getHistory: (entityTypeId: number, entityId: number, offset: number, limit: number, go: (error: Error, history: EntityHistory[]) => void) => void

  // Create a package for a project
  createPackage: (projectId: number, name: string, go: (error: Error) => void) => void

  // List packages for a project
  getPackages: (projectId: number, go: (error: Error, packages: string[]) => void) => void

  // List directories in a project package
  getPackageDirectories: (projectId: number, packageName: string, relativePath: string, go: (error: Error, directories: string[]) => void) => void

  // List files in a project package
  getPackageFiles: (projectId: number, packageName: string, relativePath: string, go: (error: Error, files: string[]) => void) => void

  // Delete a project package
  deletePackage: (projectId: number, name: string, go: (error: Error) => void) => void

  // Delete a directory in a project package
  deletePackageDirectory: (projectId: number, packageName: string, relativePath: string, go: (error: Error) => void) => void

  // Delete a file in a project package
  deletePackageFile: (projectId: number, packageName: string, relativePath: string, go: (error: Error) => void) => void

  // Set attributes on a project package
  setAttributesForPackage: (projectId: number, packageName: string, attributes: string, go: (error: Error) => void) => void

  // List attributes for a project package
  getAttributesForPackage: (projectId: number, packageName: string, go: (error: Error, attributes: string) => void) => void

}

// --- Messages ---

interface PingServerIn {

  input: string

}

interface PingServerOut {

  output: string

}

interface GetConfigIn {

}

interface GetConfigOut {

  config: Config

}

interface RegisterClusterIn {

  address: string

}

interface RegisterClusterOut {

  cluster_id: number

}

interface UnregisterClusterIn {

  cluster_id: number

}

interface UnregisterClusterOut {

}

interface StartClusterOnYarnIn {

  cluster_name: string

  engine_id: number

  size: number

  memory: string

  secure: boolean

  keytab: string

}

interface StartClusterOnYarnOut {

  cluster_id: number

}

interface StopClusterOnYarnIn {

  cluster_id: number

  keytab: string

}

interface StopClusterOnYarnOut {

}

interface GetClusterIn {

  cluster_id: number

}

interface GetClusterOut {

  cluster: Cluster

}

interface GetClusterOnYarnIn {

  cluster_id: number

}

interface GetClusterOnYarnOut {

  cluster: YarnCluster

}

interface GetClustersIn {

  offset: number

  limit: number

}

interface GetClustersOut {

  clusters: Cluster[]

}

interface GetClusterStatusIn {

  cluster_id: number

}

interface GetClusterStatusOut {

  cluster_status: ClusterStatus

}

interface DeleteClusterIn {

  cluster_id: number

}

interface DeleteClusterOut {

}

interface GetJobIn {

  cluster_id: number

  job_name: string

}

interface GetJobOut {

  job: Job

}

interface GetJobsIn {

  cluster_id: number

}

interface GetJobsOut {

  jobs: Job[]

}

interface CreateProjectIn {

  name: string

  description: string

  model_category: string

}

interface CreateProjectOut {

  project_id: number

}

interface GetProjectsIn {

  offset: number

  limit: number

}

interface GetProjectsOut {

  projects: Project[]

}

interface GetProjectIn {

  project_id: number

}

interface GetProjectOut {

  project: Project

}

interface DeleteProjectIn {

  project_id: number

}

interface DeleteProjectOut {

}

interface CreateDatasourceIn {

  project_id: number

  name: string

  description: string

  path: string

}

interface CreateDatasourceOut {

  datasource_id: number

}

interface GetDatasourcesIn {

  project_id: number

  offset: number

  limit: number

}

interface GetDatasourcesOut {

  datasources: Datasource[]

}

interface GetDatasourceIn {

  datasource_id: number

}

interface GetDatasourceOut {

  datasource: Datasource

}

interface UpdateDatasourceIn {

  datasource_id: number

  name: string

  description: string

  path: string

}

interface UpdateDatasourceOut {

}

interface DeleteDatasourceIn {

  datasource_id: number

}

interface DeleteDatasourceOut {

}

interface CreateDatasetIn {

  cluster_id: number

  datasource_id: number

  name: string

  description: string

  response_column_name: string

}

interface CreateDatasetOut {

  dataset_id: number

}

interface GetDatasetsIn {

  datasource_id: number

  offset: number

  limit: number

}

interface GetDatasetsOut {

  datasets: Dataset[]

}

interface GetDatasetIn {

  dataset_id: number

}

interface GetDatasetOut {

  dataset: Dataset

}

interface GetDatasetsFromClusterIn {

  cluster_id: number

}

interface GetDatasetsFromClusterOut {

  dataset: Dataset[]

}

interface UpdateDatasetIn {

  dataset_id: number

  name: string

  description: string

  response_column_name: string

}

interface UpdateDatasetOut {

}

interface SplitDatasetIn {

  dataset_id: number

  ratio1: number

  ratio2: number

}

interface SplitDatasetOut {

  dataset_ids: number[]

}

interface DeleteDatasetIn {

  dataset_id: number

}

interface DeleteDatasetOut {

}

interface BuildModelIn {

  cluster_id: number

  dataset_id: number

  algorithm: string

}

interface BuildModelOut {

  model_id: number

}

interface BuildModelAutoIn {

  cluster_id: number

  dataset: string

  target_name: string

  max_run_time: number

}

interface BuildModelAutoOut {

  model: Model

}

interface GetModelIn {

  model_id: number

}

interface GetModelOut {

  model: Model

}

interface GetModelsIn {

  project_id: number

  offset: number

  limit: number

}

interface GetModelsOut {

  models: Model[]

}

interface GetModelsFromClusterIn {

  cluster_id: number

  frame_key: string

}

interface GetModelsFromClusterOut {

  models: Model[]

}

interface FindModelsCountIn {

  project_id: number

}

interface FindModelsCountOut {

  count: number

}

interface GetAllBinomialSortCriteriaIn {

}

interface GetAllBinomialSortCriteriaOut {

  criteria: string[]

}

interface FindModelsBinomialIn {

  project_id: number

  name_part: string

  sort_by: string

  ascending: boolean

  offset: number

  limit: number

}

interface FindModelsBinomialOut {

  models: BinomialModel[]

}

interface GetModelBinomialIn {

  model_id: number

}

interface GetModelBinomialOut {

  model: BinomialModel

}

interface GetAllMultinomialSortCriteriaIn {

}

interface GetAllMultinomialSortCriteriaOut {

  criteria: string[]

}

interface FindModelsMultinomialIn {

  project_id: number

  name_part: string

  sort_by: string

  ascending: boolean

  offset: number

  limit: number

}

interface FindModelsMultinomialOut {

  models: MultinomialModel[]

}

interface GetModelMultinomialIn {

  model_id: number

}

interface GetModelMultinomialOut {

  model: MultinomialModel

}

interface GetAllRegressionSortCriteriaIn {

}

interface GetAllRegressionSortCriteriaOut {

  criteria: string[]

}

interface FindModelsRegressionIn {

  project_id: number

  name_part: string

  sort_by: string

  ascending: boolean

  offset: number

  limit: number

}

interface FindModelsRegressionOut {

  models: RegressionModel[]

}

interface GetModelRegressionIn {

  model_id: number

}

interface GetModelRegressionOut {

  model: RegressionModel

}

interface ImportModelFromClusterIn {

  cluster_id: number

  project_id: number

  model_key: string

  model_name: string

}

interface ImportModelFromClusterOut {

  model_id: number

}

interface CheckMojoIn {

  algo: string

}

interface CheckMojoOut {

  can_mojo: boolean

}

interface ImportModelPojoIn {

  model_id: number

}

interface ImportModelPojoOut {

}

interface ImportModelMojoIn {

  model_id: number

}

interface ImportModelMojoOut {

}

interface DeleteModelIn {

  model_id: number

}

interface DeleteModelOut {

}

interface CreateLabelIn {

  project_id: number

  name: string

  description: string

}

interface CreateLabelOut {

  label_id: number

}

interface UpdateLabelIn {

  label_id: number

  name: string

  description: string

}

interface UpdateLabelOut {

}

interface DeleteLabelIn {

  label_id: number

}

interface DeleteLabelOut {

}

interface LinkLabelWithModelIn {

  label_id: number

  model_id: number

}

interface LinkLabelWithModelOut {

}

interface UnlinkLabelFromModelIn {

  label_id: number

  model_id: number

}

interface UnlinkLabelFromModelOut {

}

interface GetLabelsForProjectIn {

  project_id: number

}

interface GetLabelsForProjectOut {

  labels: Label[]

}

interface StartServiceIn {

  model_id: number

  name: string

  package_name: string

}

interface StartServiceOut {

  service_id: number

}

interface StopServiceIn {

  service_id: number

}

interface StopServiceOut {

}

interface GetServiceIn {

  service_id: number

}

interface GetServiceOut {

  service: ScoringService

}

interface GetServicesIn {

  offset: number

  limit: number

}

interface GetServicesOut {

  services: ScoringService[]

}

interface GetServicesForProjectIn {

  project_id: number

  offset: number

  limit: number

}

interface GetServicesForProjectOut {

  services: ScoringService[]

}

interface GetServicesForModelIn {

  model_id: number

  offset: number

  limit: number

}

interface GetServicesForModelOut {

  services: ScoringService[]

}

interface DeleteServiceIn {

  service_id: number

}

interface DeleteServiceOut {

}

interface GetEngineIn {

  engine_id: number

}

interface GetEngineOut {

  engine: Engine

}

interface GetEnginesIn {

}

interface GetEnginesOut {

  engines: Engine[]

}

interface DeleteEngineIn {

  engine_id: number

}

interface DeleteEngineOut {

}

interface GetAllEntityTypesIn {

}

interface GetAllEntityTypesOut {

  entity_types: EntityType[]

}

interface GetAllPermissionsIn {

}

interface GetAllPermissionsOut {

  permissions: Permission[]

}

interface GetAllClusterTypesIn {

}

interface GetAllClusterTypesOut {

  cluster_types: ClusterType[]

}

interface GetPermissionsForRoleIn {

  role_id: number

}

interface GetPermissionsForRoleOut {

  permissions: Permission[]

}

interface GetPermissionsForIdentityIn {

  identity_id: number

}

interface GetPermissionsForIdentityOut {

  permissions: Permission[]

}

interface CreateRoleIn {

  name: string

  description: string

}

interface CreateRoleOut {

  role_id: number

}

interface GetRolesIn {

  offset: number

  limit: number

}

interface GetRolesOut {

  roles: Role[]

}

interface GetRolesForIdentityIn {

  identity_id: number

}

interface GetRolesForIdentityOut {

  roles: Role[]

}

interface GetRoleIn {

  role_id: number

}

interface GetRoleOut {

  role: Role

}

interface GetRoleByNameIn {

  name: string

}

interface GetRoleByNameOut {

  role: Role

}

interface UpdateRoleIn {

  role_id: number

  name: string

  description: string

}

interface UpdateRoleOut {

}

interface LinkRoleWithPermissionsIn {

  role_id: number

  permission_ids: number[]

}

interface LinkRoleWithPermissionsOut {

}

interface LinkRoleWithPermissionIn {

  role_id: number

  permission_id: number

}

interface LinkRoleWithPermissionOut {

}

interface UnlinkRoleFromPermissionIn {

  role_id: number

  permission_id: number

}

interface UnlinkRoleFromPermissionOut {

}

interface DeleteRoleIn {

  role_id: number

}

interface DeleteRoleOut {

}

interface CreateWorkgroupIn {

  name: string

  description: string

}

interface CreateWorkgroupOut {

  workgroup_id: number

}

interface GetWorkgroupsIn {

  offset: number

  limit: number

}

interface GetWorkgroupsOut {

  workgroups: Workgroup[]

}

interface GetWorkgroupsForIdentityIn {

  identity_id: number

}

interface GetWorkgroupsForIdentityOut {

  workgroups: Workgroup[]

}

interface GetWorkgroupIn {

  workgroup_id: number

}

interface GetWorkgroupOut {

  workgroup: Workgroup

}

interface GetWorkgroupByNameIn {

  name: string

}

interface GetWorkgroupByNameOut {

  workgroup: Workgroup

}

interface UpdateWorkgroupIn {

  workgroup_id: number

  name: string

  description: string

}

interface UpdateWorkgroupOut {

}

interface DeleteWorkgroupIn {

  workgroup_id: number

}

interface DeleteWorkgroupOut {

}

interface CreateIdentityIn {

  name: string

  password: string

}

interface CreateIdentityOut {

  identity_id: number

}

interface GetIdentitiesIn {

  offset: number

  limit: number

}

interface GetIdentitiesOut {

  identities: Identity[]

}

interface GetIdentitiesForWorkgroupIn {

  workgroup_id: number

}

interface GetIdentitiesForWorkgroupOut {

  identities: Identity[]

}

interface GetIdentitiesForRoleIn {

  role_id: number

}

interface GetIdentitiesForRoleOut {

  identities: Identity[]

}

interface GetIdentitiesForEntityIn {

  entity_type: number

  entity_id: number

}

interface GetIdentitiesForEntityOut {

  users: UserRole[]

}

interface GetIdentityIn {

  identity_id: number

}

interface GetIdentityOut {

  identity: Identity

}

interface GetIdentityByNameIn {

  name: string

}

interface GetIdentityByNameOut {

  identity: Identity

}

interface LinkIdentityWithWorkgroupIn {

  identity_id: number

  workgroup_id: number

}

interface LinkIdentityWithWorkgroupOut {

}

interface UnlinkIdentityFromWorkgroupIn {

  identity_id: number

  workgroup_id: number

}

interface UnlinkIdentityFromWorkgroupOut {

}

interface LinkIdentityWithRoleIn {

  identity_id: number

  role_id: number

}

interface LinkIdentityWithRoleOut {

}

interface UnlinkIdentityFromRoleIn {

  identity_id: number

  role_id: number

}

interface UnlinkIdentityFromRoleOut {

}

interface UpdateIdentityIn {

  identity_id: number

  password: string

}

interface UpdateIdentityOut {

}

interface ActivateIdentityIn {

  identity_id: number

}

interface ActivateIdentityOut {

}

interface DeactivateIdentityIn {

  identity_id: number

}

interface DeactivateIdentityOut {

}

interface ShareEntityIn {

  kind: string

  workgroup_id: number

  entity_type_id: number

  entity_id: number

}

interface ShareEntityOut {

}

interface GetPrivilegesIn {

  entity_type_id: number

  entity_id: number

}

interface GetPrivilegesOut {

  privileges: EntityPrivilege[]

}

interface UnshareEntityIn {

  kind: string

  workgroup_id: number

  entity_type_id: number

  entity_id: number

}

interface UnshareEntityOut {

}

interface GetHistoryIn {

  entity_type_id: number

  entity_id: number

  offset: number

  limit: number

}

interface GetHistoryOut {

  history: EntityHistory[]

}

interface CreatePackageIn {

  project_id: number

  name: string

}

interface CreatePackageOut {

}

interface GetPackagesIn {

  project_id: number

}

interface GetPackagesOut {

  packages: string[]

}

interface GetPackageDirectoriesIn {

  project_id: number

  package_name: string

  relative_path: string

}

interface GetPackageDirectoriesOut {

  directories: string[]

}

interface GetPackageFilesIn {

  project_id: number

  package_name: string

  relative_path: string

}

interface GetPackageFilesOut {

  files: string[]

}

interface DeletePackageIn {

  project_id: number

  name: string

}

interface DeletePackageOut {

}

interface DeletePackageDirectoryIn {

  project_id: number

  package_name: string

  relative_path: string

}

interface DeletePackageDirectoryOut {

}

interface DeletePackageFileIn {

  project_id: number

  package_name: string

  relative_path: string

}

interface DeletePackageFileOut {

}

interface SetAttributesForPackageIn {

  project_id: number

  package_name: string

  attributes: string

}

interface SetAttributesForPackageOut {

}

interface GetAttributesForPackageIn {

  project_id: number

  package_name: string

}

interface GetAttributesForPackageOut {

  attributes: string

}



// --- Client Stub ---


export function pingServer(input: string, go: (error: Error, output: string) => void): void {
  const req: PingServerIn = { input: input };
  Proxy.Call("PingServer", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: PingServerOut = <PingServerOut> data;
      return go(null, d.output);
    }
  });
}

export function getConfig(go: (error: Error, config: Config) => void): void {
  const req: GetConfigIn = {  };
  Proxy.Call("GetConfig", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetConfigOut = <GetConfigOut> data;
      return go(null, d.config);
    }
  });
}

export function registerCluster(address: string, go: (error: Error, clusterId: number) => void): void {
  const req: RegisterClusterIn = { address: address };
  Proxy.Call("RegisterCluster", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: RegisterClusterOut = <RegisterClusterOut> data;
      return go(null, d.cluster_id);
    }
  });
}

export function unregisterCluster(clusterId: number, go: (error: Error) => void): void {
  const req: UnregisterClusterIn = { cluster_id: clusterId };
  Proxy.Call("UnregisterCluster", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: UnregisterClusterOut = <UnregisterClusterOut> data;
      return go(null);
    }
  });
}

export function startClusterOnYarn(clusterName: string, engineId: number, size: number, memory: string, secure: boolean, keytab: string, go: (error: Error, clusterId: number) => void): void {
  const req: StartClusterOnYarnIn = { cluster_name: clusterName, engine_id: engineId, size: size, memory: memory, secure: secure, keytab: keytab };
  Proxy.Call("StartClusterOnYarn", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: StartClusterOnYarnOut = <StartClusterOnYarnOut> data;
      return go(null, d.cluster_id);
    }
  });
}

export function stopClusterOnYarn(clusterId: number, keytab: string, go: (error: Error) => void): void {
  const req: StopClusterOnYarnIn = { cluster_id: clusterId, keytab: keytab };
  Proxy.Call("StopClusterOnYarn", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: StopClusterOnYarnOut = <StopClusterOnYarnOut> data;
      return go(null);
    }
  });
}

export function getCluster(clusterId: number, go: (error: Error, cluster: Cluster) => void): void {
  const req: GetClusterIn = { cluster_id: clusterId };
  Proxy.Call("GetCluster", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetClusterOut = <GetClusterOut> data;
      return go(null, d.cluster);
    }
  });
}

export function getClusterOnYarn(clusterId: number, go: (error: Error, cluster: YarnCluster) => void): void {
  const req: GetClusterOnYarnIn = { cluster_id: clusterId };
  Proxy.Call("GetClusterOnYarn", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetClusterOnYarnOut = <GetClusterOnYarnOut> data;
      return go(null, d.cluster);
    }
  });
}

export function getClusters(offset: number, limit: number, go: (error: Error, clusters: Cluster[]) => void): void {
  const req: GetClustersIn = { offset: offset, limit: limit };
  Proxy.Call("GetClusters", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetClustersOut = <GetClustersOut> data;
      return go(null, d.clusters);
    }
  });
}

export function getClusterStatus(clusterId: number, go: (error: Error, clusterStatus: ClusterStatus) => void): void {
  const req: GetClusterStatusIn = { cluster_id: clusterId };
  Proxy.Call("GetClusterStatus", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetClusterStatusOut = <GetClusterStatusOut> data;
      return go(null, d.cluster_status);
    }
  });
}

export function deleteCluster(clusterId: number, go: (error: Error) => void): void {
  const req: DeleteClusterIn = { cluster_id: clusterId };
  Proxy.Call("DeleteCluster", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: DeleteClusterOut = <DeleteClusterOut> data;
      return go(null);
    }
  });
}

export function getJob(clusterId: number, jobName: string, go: (error: Error, job: Job) => void): void {
  const req: GetJobIn = { cluster_id: clusterId, job_name: jobName };
  Proxy.Call("GetJob", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetJobOut = <GetJobOut> data;
      return go(null, d.job);
    }
  });
}

export function getJobs(clusterId: number, go: (error: Error, jobs: Job[]) => void): void {
  const req: GetJobsIn = { cluster_id: clusterId };
  Proxy.Call("GetJobs", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetJobsOut = <GetJobsOut> data;
      return go(null, d.jobs);
    }
  });
}

export function createProject(name: string, description: string, modelCategory: string, go: (error: Error, projectId: number) => void): void {
  const req: CreateProjectIn = { name: name, description: description, model_category: modelCategory };
  Proxy.Call("CreateProject", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: CreateProjectOut = <CreateProjectOut> data;
      return go(null, d.project_id);
    }
  });
}

export function getProjects(offset: number, limit: number, go: (error: Error, projects: Project[]) => void): void {
  const req: GetProjectsIn = { offset: offset, limit: limit };
  Proxy.Call("GetProjects", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetProjectsOut = <GetProjectsOut> data;
      return go(null, d.projects);
    }
  });
}

export function getProject(projectId: number, go: (error: Error, project: Project) => void): void {
  const req: GetProjectIn = { project_id: projectId };
  Proxy.Call("GetProject", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetProjectOut = <GetProjectOut> data;
      return go(null, d.project);
    }
  });
}

export function deleteProject(projectId: number, go: (error: Error) => void): void {
  const req: DeleteProjectIn = { project_id: projectId };
  Proxy.Call("DeleteProject", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: DeleteProjectOut = <DeleteProjectOut> data;
      return go(null);
    }
  });
}

export function createDatasource(projectId: number, name: string, description: string, path: string, go: (error: Error, datasourceId: number) => void): void {
  const req: CreateDatasourceIn = { project_id: projectId, name: name, description: description, path: path };
  Proxy.Call("CreateDatasource", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: CreateDatasourceOut = <CreateDatasourceOut> data;
      return go(null, d.datasource_id);
    }
  });
}

export function getDatasources(projectId: number, offset: number, limit: number, go: (error: Error, datasources: Datasource[]) => void): void {
  const req: GetDatasourcesIn = { project_id: projectId, offset: offset, limit: limit };
  Proxy.Call("GetDatasources", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetDatasourcesOut = <GetDatasourcesOut> data;
      return go(null, d.datasources);
    }
  });
}

export function getDatasource(datasourceId: number, go: (error: Error, datasource: Datasource) => void): void {
  const req: GetDatasourceIn = { datasource_id: datasourceId };
  Proxy.Call("GetDatasource", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetDatasourceOut = <GetDatasourceOut> data;
      return go(null, d.datasource);
    }
  });
}

export function updateDatasource(datasourceId: number, name: string, description: string, path: string, go: (error: Error) => void): void {
  const req: UpdateDatasourceIn = { datasource_id: datasourceId, name: name, description: description, path: path };
  Proxy.Call("UpdateDatasource", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: UpdateDatasourceOut = <UpdateDatasourceOut> data;
      return go(null);
    }
  });
}

export function deleteDatasource(datasourceId: number, go: (error: Error) => void): void {
  const req: DeleteDatasourceIn = { datasource_id: datasourceId };
  Proxy.Call("DeleteDatasource", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: DeleteDatasourceOut = <DeleteDatasourceOut> data;
      return go(null);
    }
  });
}

export function createDataset(clusterId: number, datasourceId: number, name: string, description: string, responseColumnName: string, go: (error: Error, datasetId: number) => void): void {
  const req: CreateDatasetIn = { cluster_id: clusterId, datasource_id: datasourceId, name: name, description: description, response_column_name: responseColumnName };
  Proxy.Call("CreateDataset", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: CreateDatasetOut = <CreateDatasetOut> data;
      return go(null, d.dataset_id);
    }
  });
}

export function getDatasets(datasourceId: number, offset: number, limit: number, go: (error: Error, datasets: Dataset[]) => void): void {
  const req: GetDatasetsIn = { datasource_id: datasourceId, offset: offset, limit: limit };
  Proxy.Call("GetDatasets", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetDatasetsOut = <GetDatasetsOut> data;
      return go(null, d.datasets);
    }
  });
}

export function getDataset(datasetId: number, go: (error: Error, dataset: Dataset) => void): void {
  const req: GetDatasetIn = { dataset_id: datasetId };
  Proxy.Call("GetDataset", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetDatasetOut = <GetDatasetOut> data;
      return go(null, d.dataset);
    }
  });
}

export function getDatasetsFromCluster(clusterId: number, go: (error: Error, dataset: Dataset[]) => void): void {
  const req: GetDatasetsFromClusterIn = { cluster_id: clusterId };
  Proxy.Call("GetDatasetsFromCluster", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetDatasetsFromClusterOut = <GetDatasetsFromClusterOut> data;
      return go(null, d.dataset);
    }
  });
}

export function updateDataset(datasetId: number, name: string, description: string, responseColumnName: string, go: (error: Error) => void): void {
  const req: UpdateDatasetIn = { dataset_id: datasetId, name: name, description: description, response_column_name: responseColumnName };
  Proxy.Call("UpdateDataset", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: UpdateDatasetOut = <UpdateDatasetOut> data;
      return go(null);
    }
  });
}

export function splitDataset(datasetId: number, ratio1: number, ratio2: number, go: (error: Error, datasetIds: number[]) => void): void {
  const req: SplitDatasetIn = { dataset_id: datasetId, ratio1: ratio1, ratio2: ratio2 };
  Proxy.Call("SplitDataset", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: SplitDatasetOut = <SplitDatasetOut> data;
      return go(null, d.dataset_ids);
    }
  });
}

export function deleteDataset(datasetId: number, go: (error: Error) => void): void {
  const req: DeleteDatasetIn = { dataset_id: datasetId };
  Proxy.Call("DeleteDataset", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: DeleteDatasetOut = <DeleteDatasetOut> data;
      return go(null);
    }
  });
}

export function buildModel(clusterId: number, datasetId: number, algorithm: string, go: (error: Error, modelId: number) => void): void {
  const req: BuildModelIn = { cluster_id: clusterId, dataset_id: datasetId, algorithm: algorithm };
  Proxy.Call("BuildModel", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: BuildModelOut = <BuildModelOut> data;
      return go(null, d.model_id);
    }
  });
}

export function buildModelAuto(clusterId: number, dataset: string, targetName: string, maxRunTime: number, go: (error: Error, model: Model) => void): void {
  const req: BuildModelAutoIn = { cluster_id: clusterId, dataset: dataset, target_name: targetName, max_run_time: maxRunTime };
  Proxy.Call("BuildModelAuto", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: BuildModelAutoOut = <BuildModelAutoOut> data;
      return go(null, d.model);
    }
  });
}

export function getModel(modelId: number, go: (error: Error, model: Model) => void): void {
  const req: GetModelIn = { model_id: modelId };
  Proxy.Call("GetModel", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetModelOut = <GetModelOut> data;
      return go(null, d.model);
    }
  });
}

export function getModels(projectId: number, offset: number, limit: number, go: (error: Error, models: Model[]) => void): void {
  const req: GetModelsIn = { project_id: projectId, offset: offset, limit: limit };
  Proxy.Call("GetModels", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetModelsOut = <GetModelsOut> data;
      return go(null, d.models);
    }
  });
}

export function getModelsFromCluster(clusterId: number, frameKey: string, go: (error: Error, models: Model[]) => void): void {
  const req: GetModelsFromClusterIn = { cluster_id: clusterId, frame_key: frameKey };
  Proxy.Call("GetModelsFromCluster", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetModelsFromClusterOut = <GetModelsFromClusterOut> data;
      return go(null, d.models);
    }
  });
}

export function findModelsCount(projectId: number, go: (error: Error, count: number) => void): void {
  const req: FindModelsCountIn = { project_id: projectId };
  Proxy.Call("FindModelsCount", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: FindModelsCountOut = <FindModelsCountOut> data;
      return go(null, d.count);
    }
  });
}

export function getAllBinomialSortCriteria(go: (error: Error, criteria: string[]) => void): void {
  const req: GetAllBinomialSortCriteriaIn = {  };
  Proxy.Call("GetAllBinomialSortCriteria", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetAllBinomialSortCriteriaOut = <GetAllBinomialSortCriteriaOut> data;
      return go(null, d.criteria);
    }
  });
}

export function findModelsBinomial(projectId: number, namePart: string, sortBy: string, ascending: boolean, offset: number, limit: number, go: (error: Error, models: BinomialModel[]) => void): void {
  const req: FindModelsBinomialIn = { project_id: projectId, name_part: namePart, sort_by: sortBy, ascending: ascending, offset: offset, limit: limit };
  Proxy.Call("FindModelsBinomial", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: FindModelsBinomialOut = <FindModelsBinomialOut> data;
      return go(null, d.models);
    }
  });
}

export function getModelBinomial(modelId: number, go: (error: Error, model: BinomialModel) => void): void {
  const req: GetModelBinomialIn = { model_id: modelId };
  Proxy.Call("GetModelBinomial", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetModelBinomialOut = <GetModelBinomialOut> data;
      return go(null, d.model);
    }
  });
}

export function getAllMultinomialSortCriteria(go: (error: Error, criteria: string[]) => void): void {
  const req: GetAllMultinomialSortCriteriaIn = {  };
  Proxy.Call("GetAllMultinomialSortCriteria", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetAllMultinomialSortCriteriaOut = <GetAllMultinomialSortCriteriaOut> data;
      return go(null, d.criteria);
    }
  });
}

export function findModelsMultinomial(projectId: number, namePart: string, sortBy: string, ascending: boolean, offset: number, limit: number, go: (error: Error, models: MultinomialModel[]) => void): void {
  const req: FindModelsMultinomialIn = { project_id: projectId, name_part: namePart, sort_by: sortBy, ascending: ascending, offset: offset, limit: limit };
  Proxy.Call("FindModelsMultinomial", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: FindModelsMultinomialOut = <FindModelsMultinomialOut> data;
      return go(null, d.models);
    }
  });
}

export function getModelMultinomial(modelId: number, go: (error: Error, model: MultinomialModel) => void): void {
  const req: GetModelMultinomialIn = { model_id: modelId };
  Proxy.Call("GetModelMultinomial", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetModelMultinomialOut = <GetModelMultinomialOut> data;
      return go(null, d.model);
    }
  });
}

export function getAllRegressionSortCriteria(go: (error: Error, criteria: string[]) => void): void {
  const req: GetAllRegressionSortCriteriaIn = {  };
  Proxy.Call("GetAllRegressionSortCriteria", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetAllRegressionSortCriteriaOut = <GetAllRegressionSortCriteriaOut> data;
      return go(null, d.criteria);
    }
  });
}

export function findModelsRegression(projectId: number, namePart: string, sortBy: string, ascending: boolean, offset: number, limit: number, go: (error: Error, models: RegressionModel[]) => void): void {
  const req: FindModelsRegressionIn = { project_id: projectId, name_part: namePart, sort_by: sortBy, ascending: ascending, offset: offset, limit: limit };
  Proxy.Call("FindModelsRegression", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: FindModelsRegressionOut = <FindModelsRegressionOut> data;
      return go(null, d.models);
    }
  });
}

export function getModelRegression(modelId: number, go: (error: Error, model: RegressionModel) => void): void {
  const req: GetModelRegressionIn = { model_id: modelId };
  Proxy.Call("GetModelRegression", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetModelRegressionOut = <GetModelRegressionOut> data;
      return go(null, d.model);
    }
  });
}

export function importModelFromCluster(clusterId: number, projectId: number, modelKey: string, modelName: string, go: (error: Error, modelId: number) => void): void {
  const req: ImportModelFromClusterIn = { cluster_id: clusterId, project_id: projectId, model_key: modelKey, model_name: modelName };
  Proxy.Call("ImportModelFromCluster", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: ImportModelFromClusterOut = <ImportModelFromClusterOut> data;
      return go(null, d.model_id);
    }
  });
}

export function checkMojo(algo: string, go: (error: Error, canMojo: boolean) => void): void {
  const req: CheckMojoIn = { algo: algo };
  Proxy.Call("CheckMojo", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: CheckMojoOut = <CheckMojoOut> data;
      return go(null, d.can_mojo);
    }
  });
}

export function importModelPojo(modelId: number, go: (error: Error) => void): void {
  const req: ImportModelPojoIn = { model_id: modelId };
  Proxy.Call("ImportModelPojo", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: ImportModelPojoOut = <ImportModelPojoOut> data;
      return go(null);
    }
  });
}

export function importModelMojo(modelId: number, go: (error: Error) => void): void {
  const req: ImportModelMojoIn = { model_id: modelId };
  Proxy.Call("ImportModelMojo", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: ImportModelMojoOut = <ImportModelMojoOut> data;
      return go(null);
    }
  });
}

export function deleteModel(modelId: number, go: (error: Error) => void): void {
  const req: DeleteModelIn = { model_id: modelId };
  Proxy.Call("DeleteModel", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: DeleteModelOut = <DeleteModelOut> data;
      return go(null);
    }
  });
}

export function createLabel(projectId: number, name: string, description: string, go: (error: Error, labelId: number) => void): void {
  const req: CreateLabelIn = { project_id: projectId, name: name, description: description };
  Proxy.Call("CreateLabel", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: CreateLabelOut = <CreateLabelOut> data;
      return go(null, d.label_id);
    }
  });
}

export function updateLabel(labelId: number, name: string, description: string, go: (error: Error) => void): void {
  const req: UpdateLabelIn = { label_id: labelId, name: name, description: description };
  Proxy.Call("UpdateLabel", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: UpdateLabelOut = <UpdateLabelOut> data;
      return go(null);
    }
  });
}

export function deleteLabel(labelId: number, go: (error: Error) => void): void {
  const req: DeleteLabelIn = { label_id: labelId };
  Proxy.Call("DeleteLabel", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: DeleteLabelOut = <DeleteLabelOut> data;
      return go(null);
    }
  });
}

export function linkLabelWithModel(labelId: number, modelId: number, go: (error: Error) => void): void {
  const req: LinkLabelWithModelIn = { label_id: labelId, model_id: modelId };
  Proxy.Call("LinkLabelWithModel", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: LinkLabelWithModelOut = <LinkLabelWithModelOut> data;
      return go(null);
    }
  });
}

export function unlinkLabelFromModel(labelId: number, modelId: number, go: (error: Error) => void): void {
  const req: UnlinkLabelFromModelIn = { label_id: labelId, model_id: modelId };
  Proxy.Call("UnlinkLabelFromModel", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: UnlinkLabelFromModelOut = <UnlinkLabelFromModelOut> data;
      return go(null);
    }
  });
}

export function getLabelsForProject(projectId: number, go: (error: Error, labels: Label[]) => void): void {
  const req: GetLabelsForProjectIn = { project_id: projectId };
  Proxy.Call("GetLabelsForProject", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetLabelsForProjectOut = <GetLabelsForProjectOut> data;
      return go(null, d.labels);
    }
  });
}

export function startService(modelId: number, name: string, packageName: string, go: (error: Error, serviceId: number) => void): void {
  const req: StartServiceIn = { model_id: modelId, name: name, package_name: packageName };
  Proxy.Call("StartService", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: StartServiceOut = <StartServiceOut> data;
      return go(null, d.service_id);
    }
  });
}

export function stopService(serviceId: number, go: (error: Error) => void): void {
  const req: StopServiceIn = { service_id: serviceId };
  Proxy.Call("StopService", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: StopServiceOut = <StopServiceOut> data;
      return go(null);
    }
  });
}

export function getService(serviceId: number, go: (error: Error, service: ScoringService) => void): void {
  const req: GetServiceIn = { service_id: serviceId };
  Proxy.Call("GetService", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetServiceOut = <GetServiceOut> data;
      return go(null, d.service);
    }
  });
}

export function getServices(offset: number, limit: number, go: (error: Error, services: ScoringService[]) => void): void {
  const req: GetServicesIn = { offset: offset, limit: limit };
  Proxy.Call("GetServices", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetServicesOut = <GetServicesOut> data;
      return go(null, d.services);
    }
  });
}

export function getServicesForProject(projectId: number, offset: number, limit: number, go: (error: Error, services: ScoringService[]) => void): void {
  const req: GetServicesForProjectIn = { project_id: projectId, offset: offset, limit: limit };
  Proxy.Call("GetServicesForProject", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetServicesForProjectOut = <GetServicesForProjectOut> data;
      return go(null, d.services);
    }
  });
}

export function getServicesForModel(modelId: number, offset: number, limit: number, go: (error: Error, services: ScoringService[]) => void): void {
  const req: GetServicesForModelIn = { model_id: modelId, offset: offset, limit: limit };
  Proxy.Call("GetServicesForModel", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetServicesForModelOut = <GetServicesForModelOut> data;
      return go(null, d.services);
    }
  });
}

export function deleteService(serviceId: number, go: (error: Error) => void): void {
  const req: DeleteServiceIn = { service_id: serviceId };
  Proxy.Call("DeleteService", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: DeleteServiceOut = <DeleteServiceOut> data;
      return go(null);
    }
  });
}

export function getEngine(engineId: number, go: (error: Error, engine: Engine) => void): void {
  const req: GetEngineIn = { engine_id: engineId };
  Proxy.Call("GetEngine", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetEngineOut = <GetEngineOut> data;
      return go(null, d.engine);
    }
  });
}

export function getEngines(go: (error: Error, engines: Engine[]) => void): void {
  const req: GetEnginesIn = {  };
  Proxy.Call("GetEngines", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetEnginesOut = <GetEnginesOut> data;
      return go(null, d.engines);
    }
  });
}

export function deleteEngine(engineId: number, go: (error: Error) => void): void {
  const req: DeleteEngineIn = { engine_id: engineId };
  Proxy.Call("DeleteEngine", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: DeleteEngineOut = <DeleteEngineOut> data;
      return go(null);
    }
  });
}

export function getAllEntityTypes(go: (error: Error, entityTypes: EntityType[]) => void): void {
  const req: GetAllEntityTypesIn = {  };
  Proxy.Call("GetAllEntityTypes", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetAllEntityTypesOut = <GetAllEntityTypesOut> data;
      return go(null, d.entity_types);
    }
  });
}

export function getAllPermissions(go: (error: Error, permissions: Permission[]) => void): void {
  const req: GetAllPermissionsIn = {  };
  Proxy.Call("GetAllPermissions", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetAllPermissionsOut = <GetAllPermissionsOut> data;
      return go(null, d.permissions);
    }
  });
}

export function getAllClusterTypes(go: (error: Error, clusterTypes: ClusterType[]) => void): void {
  const req: GetAllClusterTypesIn = {  };
  Proxy.Call("GetAllClusterTypes", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetAllClusterTypesOut = <GetAllClusterTypesOut> data;
      return go(null, d.cluster_types);
    }
  });
}

export function getPermissionsForRole(roleId: number, go: (error: Error, permissions: Permission[]) => void): void {
  const req: GetPermissionsForRoleIn = { role_id: roleId };
  Proxy.Call("GetPermissionsForRole", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetPermissionsForRoleOut = <GetPermissionsForRoleOut> data;
      return go(null, d.permissions);
    }
  });
}

export function getPermissionsForIdentity(identityId: number, go: (error: Error, permissions: Permission[]) => void): void {
  const req: GetPermissionsForIdentityIn = { identity_id: identityId };
  Proxy.Call("GetPermissionsForIdentity", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetPermissionsForIdentityOut = <GetPermissionsForIdentityOut> data;
      return go(null, d.permissions);
    }
  });
}

export function createRole(name: string, description: string, go: (error: Error, roleId: number) => void): void {
  const req: CreateRoleIn = { name: name, description: description };
  Proxy.Call("CreateRole", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: CreateRoleOut = <CreateRoleOut> data;
      return go(null, d.role_id);
    }
  });
}

export function getRoles(offset: number, limit: number, go: (error: Error, roles: Role[]) => void): void {
  const req: GetRolesIn = { offset: offset, limit: limit };
  Proxy.Call("GetRoles", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetRolesOut = <GetRolesOut> data;
      return go(null, d.roles);
    }
  });
}

export function getRolesForIdentity(identityId: number, go: (error: Error, roles: Role[]) => void): void {
  const req: GetRolesForIdentityIn = { identity_id: identityId };
  Proxy.Call("GetRolesForIdentity", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetRolesForIdentityOut = <GetRolesForIdentityOut> data;
      return go(null, d.roles);
    }
  });
}

export function getRole(roleId: number, go: (error: Error, role: Role) => void): void {
  const req: GetRoleIn = { role_id: roleId };
  Proxy.Call("GetRole", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetRoleOut = <GetRoleOut> data;
      return go(null, d.role);
    }
  });
}

export function getRoleByName(name: string, go: (error: Error, role: Role) => void): void {
  const req: GetRoleByNameIn = { name: name };
  Proxy.Call("GetRoleByName", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetRoleByNameOut = <GetRoleByNameOut> data;
      return go(null, d.role);
    }
  });
}

export function updateRole(roleId: number, name: string, description: string, go: (error: Error) => void): void {
  const req: UpdateRoleIn = { role_id: roleId, name: name, description: description };
  Proxy.Call("UpdateRole", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: UpdateRoleOut = <UpdateRoleOut> data;
      return go(null);
    }
  });
}

export function linkRoleWithPermissions(roleId: number, permissionIds: number[], go: (error: Error) => void): void {
  const req: LinkRoleWithPermissionsIn = { role_id: roleId, permission_ids: permissionIds };
  Proxy.Call("LinkRoleWithPermissions", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: LinkRoleWithPermissionsOut = <LinkRoleWithPermissionsOut> data;
      return go(null);
    }
  });
}

export function linkRoleWithPermission(roleId: number, permissionId: number, go: (error: Error) => void): void {
  const req: LinkRoleWithPermissionIn = { role_id: roleId, permission_id: permissionId };
  Proxy.Call("LinkRoleWithPermission", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: LinkRoleWithPermissionOut = <LinkRoleWithPermissionOut> data;
      return go(null);
    }
  });
}

export function unlinkRoleFromPermission(roleId: number, permissionId: number, go: (error: Error) => void): void {
  const req: UnlinkRoleFromPermissionIn = { role_id: roleId, permission_id: permissionId };
  Proxy.Call("UnlinkRoleFromPermission", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: UnlinkRoleFromPermissionOut = <UnlinkRoleFromPermissionOut> data;
      return go(null);
    }
  });
}

export function deleteRole(roleId: number, go: (error: Error) => void): void {
  const req: DeleteRoleIn = { role_id: roleId };
  Proxy.Call("DeleteRole", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: DeleteRoleOut = <DeleteRoleOut> data;
      return go(null);
    }
  });
}

export function createWorkgroup(name: string, description: string, go: (error: Error, workgroupId: number) => void): void {
  const req: CreateWorkgroupIn = { name: name, description: description };
  Proxy.Call("CreateWorkgroup", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: CreateWorkgroupOut = <CreateWorkgroupOut> data;
      return go(null, d.workgroup_id);
    }
  });
}

export function getWorkgroups(offset: number, limit: number, go: (error: Error, workgroups: Workgroup[]) => void): void {
  const req: GetWorkgroupsIn = { offset: offset, limit: limit };
  Proxy.Call("GetWorkgroups", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetWorkgroupsOut = <GetWorkgroupsOut> data;
      return go(null, d.workgroups);
    }
  });
}

export function getWorkgroupsForIdentity(identityId: number, go: (error: Error, workgroups: Workgroup[]) => void): void {
  const req: GetWorkgroupsForIdentityIn = { identity_id: identityId };
  Proxy.Call("GetWorkgroupsForIdentity", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetWorkgroupsForIdentityOut = <GetWorkgroupsForIdentityOut> data;
      return go(null, d.workgroups);
    }
  });
}

export function getWorkgroup(workgroupId: number, go: (error: Error, workgroup: Workgroup) => void): void {
  const req: GetWorkgroupIn = { workgroup_id: workgroupId };
  Proxy.Call("GetWorkgroup", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetWorkgroupOut = <GetWorkgroupOut> data;
      return go(null, d.workgroup);
    }
  });
}

export function getWorkgroupByName(name: string, go: (error: Error, workgroup: Workgroup) => void): void {
  const req: GetWorkgroupByNameIn = { name: name };
  Proxy.Call("GetWorkgroupByName", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetWorkgroupByNameOut = <GetWorkgroupByNameOut> data;
      return go(null, d.workgroup);
    }
  });
}

export function updateWorkgroup(workgroupId: number, name: string, description: string, go: (error: Error) => void): void {
  const req: UpdateWorkgroupIn = { workgroup_id: workgroupId, name: name, description: description };
  Proxy.Call("UpdateWorkgroup", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: UpdateWorkgroupOut = <UpdateWorkgroupOut> data;
      return go(null);
    }
  });
}

export function deleteWorkgroup(workgroupId: number, go: (error: Error) => void): void {
  const req: DeleteWorkgroupIn = { workgroup_id: workgroupId };
  Proxy.Call("DeleteWorkgroup", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: DeleteWorkgroupOut = <DeleteWorkgroupOut> data;
      return go(null);
    }
  });
}

export function createIdentity(name: string, password: string, go: (error: Error, identityId: number) => void): void {
  const req: CreateIdentityIn = { name: name, password: password };
  Proxy.Call("CreateIdentity", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: CreateIdentityOut = <CreateIdentityOut> data;
      return go(null, d.identity_id);
    }
  });
}

export function getIdentities(offset: number, limit: number, go: (error: Error, identities: Identity[]) => void): void {
  const req: GetIdentitiesIn = { offset: offset, limit: limit };
  Proxy.Call("GetIdentities", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetIdentitiesOut = <GetIdentitiesOut> data;
      return go(null, d.identities);
    }
  });
}

export function getIdentitiesForWorkgroup(workgroupId: number, go: (error: Error, identities: Identity[]) => void): void {
  const req: GetIdentitiesForWorkgroupIn = { workgroup_id: workgroupId };
  Proxy.Call("GetIdentitiesForWorkgroup", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetIdentitiesForWorkgroupOut = <GetIdentitiesForWorkgroupOut> data;
      return go(null, d.identities);
    }
  });
}

export function getIdentitiesForRole(roleId: number, go: (error: Error, identities: Identity[]) => void): void {
  const req: GetIdentitiesForRoleIn = { role_id: roleId };
  Proxy.Call("GetIdentitiesForRole", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetIdentitiesForRoleOut = <GetIdentitiesForRoleOut> data;
      return go(null, d.identities);
    }
  });
}

export function getIdentitiesForEntity(entityType: number, entityId: number, go: (error: Error, users: UserRole[]) => void): void {
  const req: GetIdentitiesForEntityIn = { entity_type: entityType, entity_id: entityId };
  Proxy.Call("GetIdentitiesForEntity", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetIdentitiesForEntityOut = <GetIdentitiesForEntityOut> data;
      return go(null, d.users);
    }
  });
}

export function getIdentity(identityId: number, go: (error: Error, identity: Identity) => void): void {
  const req: GetIdentityIn = { identity_id: identityId };
  Proxy.Call("GetIdentity", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetIdentityOut = <GetIdentityOut> data;
      return go(null, d.identity);
    }
  });
}

export function getIdentityByName(name: string, go: (error: Error, identity: Identity) => void): void {
  const req: GetIdentityByNameIn = { name: name };
  Proxy.Call("GetIdentityByName", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetIdentityByNameOut = <GetIdentityByNameOut> data;
      return go(null, d.identity);
    }
  });
}

export function linkIdentityWithWorkgroup(identityId: number, workgroupId: number, go: (error: Error) => void): void {
  const req: LinkIdentityWithWorkgroupIn = { identity_id: identityId, workgroup_id: workgroupId };
  Proxy.Call("LinkIdentityWithWorkgroup", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: LinkIdentityWithWorkgroupOut = <LinkIdentityWithWorkgroupOut> data;
      return go(null);
    }
  });
}

export function unlinkIdentityFromWorkgroup(identityId: number, workgroupId: number, go: (error: Error) => void): void {
  const req: UnlinkIdentityFromWorkgroupIn = { identity_id: identityId, workgroup_id: workgroupId };
  Proxy.Call("UnlinkIdentityFromWorkgroup", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: UnlinkIdentityFromWorkgroupOut = <UnlinkIdentityFromWorkgroupOut> data;
      return go(null);
    }
  });
}

export function linkIdentityWithRole(identityId: number, roleId: number, go: (error: Error) => void): void {
  const req: LinkIdentityWithRoleIn = { identity_id: identityId, role_id: roleId };
  Proxy.Call("LinkIdentityWithRole", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: LinkIdentityWithRoleOut = <LinkIdentityWithRoleOut> data;
      return go(null);
    }
  });
}

export function unlinkIdentityFromRole(identityId: number, roleId: number, go: (error: Error) => void): void {
  const req: UnlinkIdentityFromRoleIn = { identity_id: identityId, role_id: roleId };
  Proxy.Call("UnlinkIdentityFromRole", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: UnlinkIdentityFromRoleOut = <UnlinkIdentityFromRoleOut> data;
      return go(null);
    }
  });
}

export function updateIdentity(identityId: number, password: string, go: (error: Error) => void): void {
  const req: UpdateIdentityIn = { identity_id: identityId, password: password };
  Proxy.Call("UpdateIdentity", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: UpdateIdentityOut = <UpdateIdentityOut> data;
      return go(null);
    }
  });
}

export function activateIdentity(identityId: number, go: (error: Error) => void): void {
  const req: ActivateIdentityIn = { identity_id: identityId };
  Proxy.Call("ActivateIdentity", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: ActivateIdentityOut = <ActivateIdentityOut> data;
      return go(null);
    }
  });
}

export function deactivateIdentity(identityId: number, go: (error: Error) => void): void {
  const req: DeactivateIdentityIn = { identity_id: identityId };
  Proxy.Call("DeactivateIdentity", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: DeactivateIdentityOut = <DeactivateIdentityOut> data;
      return go(null);
    }
  });
}

export function shareEntity(kind: string, workgroupId: number, entityTypeId: number, entityId: number, go: (error: Error) => void): void {
  const req: ShareEntityIn = { kind: kind, workgroup_id: workgroupId, entity_type_id: entityTypeId, entity_id: entityId };
  Proxy.Call("ShareEntity", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: ShareEntityOut = <ShareEntityOut> data;
      return go(null);
    }
  });
}

export function getPrivileges(entityTypeId: number, entityId: number, go: (error: Error, privileges: EntityPrivilege[]) => void): void {
  const req: GetPrivilegesIn = { entity_type_id: entityTypeId, entity_id: entityId };
  Proxy.Call("GetPrivileges", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetPrivilegesOut = <GetPrivilegesOut> data;
      return go(null, d.privileges);
    }
  });
}

export function unshareEntity(kind: string, workgroupId: number, entityTypeId: number, entityId: number, go: (error: Error) => void): void {
  const req: UnshareEntityIn = { kind: kind, workgroup_id: workgroupId, entity_type_id: entityTypeId, entity_id: entityId };
  Proxy.Call("UnshareEntity", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: UnshareEntityOut = <UnshareEntityOut> data;
      return go(null);
    }
  });
}

export function getHistory(entityTypeId: number, entityId: number, offset: number, limit: number, go: (error: Error, history: EntityHistory[]) => void): void {
  const req: GetHistoryIn = { entity_type_id: entityTypeId, entity_id: entityId, offset: offset, limit: limit };
  Proxy.Call("GetHistory", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetHistoryOut = <GetHistoryOut> data;
      return go(null, d.history);
    }
  });
}

export function createPackage(projectId: number, name: string, go: (error: Error) => void): void {
  const req: CreatePackageIn = { project_id: projectId, name: name };
  Proxy.Call("CreatePackage", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: CreatePackageOut = <CreatePackageOut> data;
      return go(null);
    }
  });
}

export function getPackages(projectId: number, go: (error: Error, packages: string[]) => void): void {
  const req: GetPackagesIn = { project_id: projectId };
  Proxy.Call("GetPackages", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetPackagesOut = <GetPackagesOut> data;
      return go(null, d.packages);
    }
  });
}

export function getPackageDirectories(projectId: number, packageName: string, relativePath: string, go: (error: Error, directories: string[]) => void): void {
  const req: GetPackageDirectoriesIn = { project_id: projectId, package_name: packageName, relative_path: relativePath };
  Proxy.Call("GetPackageDirectories", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetPackageDirectoriesOut = <GetPackageDirectoriesOut> data;
      return go(null, d.directories);
    }
  });
}

export function getPackageFiles(projectId: number, packageName: string, relativePath: string, go: (error: Error, files: string[]) => void): void {
  const req: GetPackageFilesIn = { project_id: projectId, package_name: packageName, relative_path: relativePath };
  Proxy.Call("GetPackageFiles", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetPackageFilesOut = <GetPackageFilesOut> data;
      return go(null, d.files);
    }
  });
}

export function deletePackage(projectId: number, name: string, go: (error: Error) => void): void {
  const req: DeletePackageIn = { project_id: projectId, name: name };
  Proxy.Call("DeletePackage", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: DeletePackageOut = <DeletePackageOut> data;
      return go(null);
    }
  });
}

export function deletePackageDirectory(projectId: number, packageName: string, relativePath: string, go: (error: Error) => void): void {
  const req: DeletePackageDirectoryIn = { project_id: projectId, package_name: packageName, relative_path: relativePath };
  Proxy.Call("DeletePackageDirectory", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: DeletePackageDirectoryOut = <DeletePackageDirectoryOut> data;
      return go(null);
    }
  });
}

export function deletePackageFile(projectId: number, packageName: string, relativePath: string, go: (error: Error) => void): void {
  const req: DeletePackageFileIn = { project_id: projectId, package_name: packageName, relative_path: relativePath };
  Proxy.Call("DeletePackageFile", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: DeletePackageFileOut = <DeletePackageFileOut> data;
      return go(null);
    }
  });
}

export function setAttributesForPackage(projectId: number, packageName: string, attributes: string, go: (error: Error) => void): void {
  const req: SetAttributesForPackageIn = { project_id: projectId, package_name: packageName, attributes: attributes };
  Proxy.Call("SetAttributesForPackage", req, function(error, data) {
    if (error) {
      return go(error);
    } else {
      const d: SetAttributesForPackageOut = <SetAttributesForPackageOut> data;
      return go(null);
    }
  });
}

export function getAttributesForPackage(projectId: number, packageName: string, go: (error: Error, attributes: string) => void): void {
  const req: GetAttributesForPackageIn = { project_id: projectId, package_name: packageName };
  Proxy.Call("GetAttributesForPackage", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetAttributesForPackageOut = <GetAttributesForPackageOut> data;
      return go(null, d.attributes);
    }
  });
}



