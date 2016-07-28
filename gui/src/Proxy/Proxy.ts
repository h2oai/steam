// ------------------------------
// --- This is generated code ---
// ---      DO NOT EDIT       ---
// ------------------------------




// --- Types ---
import * as Proxy from './xhr';

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

export interface Dataset {
  
  id: number
  
  datasource_id: number
  
  name: string
  
  description: string
  
  frame_name: string
  
  response_column_name: string
  
  properties: string
  
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

export interface Model {
  
  id: number
  
  training_dataset_id: number
  
  validation_dataset_id: number
  
  name: string
  
  cluster_name: string
  
  model_key: string
  
  algorithm: string
  
  dataset_name: string
  
  response_column_name: string
  
  logical_name: string
  
  location: string
  
  max_runtime: number
  
  metrics: string
  
  created_at: number
  
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
  
  created_at: number
  
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
  
  address: string
  
  port: number
  
  process_id: number
  
  state: string
  
  created_at: number
  
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
  
  // Connect to a cluster
  registerCluster: (address: string, go: (error: Error, clusterId: number) => void) => void
  
  // Disconnect from a cluster
  unregisterCluster: (clusterId: number, go: (error: Error) => void) => void
  
  // Start a cluster using Yarn
  startClusterOnYarn: (clusterName: string, engineId: number, size: number, memory: string, username: string, go: (error: Error, clusterId: number) => void) => void
  
  // Stop a cluster using Yarn
  stopClusterOnYarn: (clusterId: number, go: (error: Error) => void) => void
  
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
  createProject: (name: string, description: string, go: (error: Error, projectId: number) => void) => void
  
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
  
  // No description available
  filterModelsByName: (projectId: number, namePart: string, offset: number, limit: number, go: (error: Error, models: Model[]) => void) => void
  
  // List models from a cluster
  getModelsFromCluster: (clusterId: number, go: (error: Error, models: Model[]) => void) => void
  
  // Import models from a cluster
  importModelFromCluster: (clusterId: number, projectId: number, modelKey: string, modelName: string, go: (error: Error, modelId: number) => void) => void
  
  // Delete a model
  deleteModel: (modelId: number, go: (error: Error) => void) => void
  
  // Start a service
  startService: (modelId: number, port: number, go: (error: Error, service: ScoringService) => void) => void
  
  // Stop a service
  stopService: (serviceId: number, go: (error: Error) => void) => void
  
  // Get service details
  getService: (serviceId: number, go: (error: Error, service: ScoringService) => void) => void
  
  // List services
  getServices: (offset: number, limit: number, go: (error: Error, services: ScoringService[]) => void) => void
  
  // List services for a model
  getServicesForModel: (modelId: number, offset: number, limit: number, go: (error: Error, services: ScoringService[]) => void) => void
  
  // Delete a service
  deleteService: (serviceId: number, go: (error: Error) => void) => void
  
  // Add an engine
  addEngine: (engineName: string, enginePath: string, go: (error: Error, engineId: number) => void) => void
  
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
  
}

// --- Messages ---

interface PingServerIn {
  
  input: string
  
}

interface PingServerOut {
  
  output: string
  
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
  
  username: string
  
}

interface StartClusterOnYarnOut {
  
  cluster_id: number
  
}

interface StopClusterOnYarnIn {
  
  cluster_id: number
  
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

interface FilterModelsByNameIn {
  
  project_id: number
  
  name_part: string
  
  offset: number
  
  limit: number
  
}

interface FilterModelsByNameOut {
  
  models: Model[]
  
}

interface GetModelsFromClusterIn {
  
  cluster_id: number
  
}

interface GetModelsFromClusterOut {
  
  models: Model[]
  
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

interface DeleteModelIn {
  
  model_id: number
  
}

interface DeleteModelOut {
  
}

interface StartServiceIn {
  
  model_id: number
  
  port: number
  
}

interface StartServiceOut {
  
  service: ScoringService
  
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

interface AddEngineIn {
  
  engine_name: string
  
  engine_path: string
  
}

interface AddEngineOut {
  
  engine_id: number
  
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

export function startClusterOnYarn(clusterName: string, engineId: number, size: number, memory: string, username: string, go: (error: Error, clusterId: number) => void): void {
  const req: StartClusterOnYarnIn = { cluster_name: clusterName, engine_id: engineId, size: size, memory: memory, username: username };
  Proxy.Call("StartClusterOnYarn", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: StartClusterOnYarnOut = <StartClusterOnYarnOut> data;
      return go(null, d.cluster_id);
    }
  });
}

export function stopClusterOnYarn(clusterId: number, go: (error: Error) => void): void {
  const req: StopClusterOnYarnIn = { cluster_id: clusterId };
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

export function createProject(name: string, description: string, go: (error: Error, projectId: number) => void): void {
  const req: CreateProjectIn = { name: name, description: description };
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

export function filterModelsByName(projectId: number, namePart: string, offset: number, limit: number, go: (error: Error, models: Model[]) => void): void {
  const req: FilterModelsByNameIn = { project_id: projectId, name_part: namePart, offset: offset, limit: limit };
  Proxy.Call("FilterModelsByName", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: FilterModelsByNameOut = <FilterModelsByNameOut> data;
      return go(null, d.models);
    }
  });
}

export function getModelsFromCluster(clusterId: number, go: (error: Error, models: Model[]) => void): void {
  const req: GetModelsFromClusterIn = { cluster_id: clusterId };
  Proxy.Call("GetModelsFromCluster", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: GetModelsFromClusterOut = <GetModelsFromClusterOut> data;
      return go(null, d.models);
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

export function startService(modelId: number, port: number, go: (error: Error, service: ScoringService) => void): void {
  const req: StartServiceIn = { model_id: modelId, port: port };
  Proxy.Call("StartService", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: StartServiceOut = <StartServiceOut> data;
      return go(null, d.service);
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

export function addEngine(engineName: string, enginePath: string, go: (error: Error, engineId: number) => void): void {
  const req: AddEngineIn = { engine_name: engineName, engine_path: enginePath };
  Proxy.Call("AddEngine", req, function(error, data) {
    if (error) {
      return go(error, null);
    } else {
      const d: AddEngineOut = <AddEngineOut> data;
      return go(null, d.engine_id);
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



