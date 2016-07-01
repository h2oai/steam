/// <reference path="xhr.ts" />
// ----------------------------------
// --- Generated with go:generate ---
// ---        DO NOT EDIT         ---
// ----------------------------------

module Proxy {

	// --- Types ---

	export interface Cluster {
		id: number
		name: string
		type_id: number
		detail_id: number
		address: string
		state: string
		created_at: number
	}

	export interface YarnCluster {
		id: number
		engine_id: number
		size: number
		application_id: string
		memory: string
		username: string
	}

	export interface ClusterStatus {
		version: string
		status: string
		max_memory: string
		total_cpu_count: number
		total_allowed_cpu_count: number
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
		name: string
		cluster_name: string
		algorithm: string
		dataset_name: string
		response_column_name: string
		logical_name: string
		location: string
		max_runtime: number
		created_at: number
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

	export interface Engine {
		id: number
		name: string
		location: string
		created_at: number
	}

	export interface EntityType {
		id: number
		name: string
	}

	export interface ClusterType {
		id: number
		name: string
	}

	export interface EntityHistory {
		identity_id: number
		action: string
		description: string
		created_at: number
	}

	export interface Permission {
		id: number
		code: string
		description: string
	}

	export interface Privilege {
		kind: string
		workgroup_id: number
	}

	export interface EntityPrivilege {
		kind: string
		workgroup_id: number
		workgroup_name: string
		workgroup_description: string
	}

	export interface Role {
		id: number
		name: string
		description: string
		created: number
	}

	export interface Identity {
		id: number
		name: string
		is_active: boolean
		last_login: number
		created: number
	}

	export interface Workgroup {
		id: number
		name: string
		description: string
		created: number
	}

	// --- Contract ---

	export interface Service {
		ping: (status: boolean, go: (error: Error, status: boolean) => void) => void
		registerCluster: (address: string, go: (error: Error, clusterId: number) => void) => void
		unregisterCluster: (clusterId: number, go: (error: Error) => void) => void
		startYarnCluster: (clusterName: string, engineId: number, size: number, memory: string, username: string, go: (error: Error, clusterId: number) => void) => void
		stopYarnCluster: (clusterId: number, go: (error: Error) => void) => void
		getCluster: (clusterId: number, go: (error: Error, cluster: Cluster) => void) => void
		getYarnCluster: (clusterId: number, go: (error: Error, cluster: YarnCluster) => void) => void
		getClusters: (offset: number, limit: number, go: (error: Error, clusters: Cluster[]) => void) => void
		getClusterStatus: (clusterId: number, go: (error: Error, clusterStatus: ClusterStatus) => void) => void
		deleteCluster: (clusterId: number, go: (error: Error) => void) => void
		getJob: (clusterId: number, jobName: string, go: (error: Error, job: Job) => void) => void
		getJobs: (clusterId: number, go: (error: Error, jobs: Job[]) => void) => void
		buildModel: (clusterId: number, dataset: string, targetName: string, maxRunTime: number, go: (error: Error, model: Model) => void) => void
		getModel: (modelId: number, go: (error: Error, model: Model) => void) => void
		getModels: (offset: number, limit: number, go: (error: Error, models: Model[]) => void) => void
		getClusterModels: (clusterId: number, go: (error: Error, models: Model[]) => void) => void
		importModelFromCluster: (clusterId: number, modelName: string, go: (error: Error, model: Model) => void) => void
		deleteModel: (modelId: number, go: (error: Error) => void) => void
		startScoringService: (modelId: number, port: number, go: (error: Error, service: ScoringService) => void) => void
		stopScoringService: (serviceId: number, go: (error: Error) => void) => void
		getScoringService: (serviceId: number, go: (error: Error, service: ScoringService) => void) => void
		getScoringServices: (offset: number, limit: number, go: (error: Error, services: ScoringService[]) => void) => void
		getScoringServicesForModel: (modelId: number, offset: number, limit: number, go: (error: Error, services: ScoringService[]) => void) => void
		deleteScoringService: (serviceId: number, go: (error: Error) => void) => void
		addEngine: (engineName: string, enginePath: string, go: (error: Error, engineId: number) => void) => void
		getEngine: (engineId: number, go: (error: Error, engine: Engine) => void) => void
		getEngines: (go: (error: Error, engines: Engine[]) => void) => void
		deleteEngine: (engineId: number, go: (error: Error) => void) => void
		getSupportedEntityTypes: (go: (error: Error, entityTypes: EntityType[]) => void) => void
		getSupportedPermissions: (go: (error: Error, permissions: Permission[]) => void) => void
		getSupportedClusterTypes: (go: (error: Error, clusterTypes: ClusterType[]) => void) => void
		getPermissionsForRole: (roleId: number, go: (error: Error, permissions: Permission[]) => void) => void
		getPermissionsForIdentity: (identityId: number, go: (error: Error, permissions: Permission[]) => void) => void
		createRole: (name: string, description: string, go: (error: Error, roleId: number) => void) => void
		getRoles: (offset: number, limit: number, go: (error: Error, roles: Role[]) => void) => void
		getRolesForIdentity: (identityId: number, go: (error: Error, roles: Role[]) => void) => void
		getRole: (roleId: number, go: (error: Error, role: Role) => void) => void
		getRoleByName: (name: string, go: (error: Error, role: Role) => void) => void
		updateRole: (roleId: number, name: string, description: string, go: (error: Error) => void) => void
		linkRoleAndPermissions: (roleId: number, permissionIds: number[], go: (error: Error) => void) => void
		deleteRole: (roleId: number, go: (error: Error) => void) => void
		createWorkgroup: (name: string, description: string, go: (error: Error, workgroupId: number) => void) => void
		getWorkgroups: (offset: number, limit: number, go: (error: Error, workgroups: Workgroup[]) => void) => void
		getWorkgroupsForIdentity: (identityId: number, go: (error: Error, workgroups: Workgroup[]) => void) => void
		getWorkgroup: (workgroupId: number, go: (error: Error, workgroup: Workgroup) => void) => void
		getWorkgroupByName: (name: string, go: (error: Error, workgroup: Workgroup) => void) => void
		updateWorkgroup: (workgroupId: number, name: string, description: string, go: (error: Error) => void) => void
		deleteWorkgroup: (workgroupId: number, go: (error: Error) => void) => void
		createIdentity: (name: string, password: string, go: (error: Error, identityId: number) => void) => void
		getIdentities: (offset: number, limit: number, go: (error: Error, identities: Identity[]) => void) => void
		getIdentitiesForWorkgroup: (workgroupId: number, go: (error: Error, identities: Identity[]) => void) => void
		getIdentititesForRole: (roleId: number, go: (error: Error, identities: Identity[]) => void) => void
		getIdentity: (identityId: number, go: (error: Error, identity: Identity) => void) => void
		getIdentityByName: (name: string, go: (error: Error, identity: Identity) => void) => void
		linkIdentityAndWorkgroup: (identityId: number, workgroupId: number, go: (error: Error) => void) => void
		unlinkIdentityAndWorkgroup: (identityId: number, workgroupId: number, go: (error: Error) => void) => void
		linkIdentityAndRole: (identityId: number, roleId: number, go: (error: Error) => void) => void
		unlinkIdentityAndRole: (identityId: number, roleId: number, go: (error: Error) => void) => void
		deactivateIdentity: (identityId: number, go: (error: Error) => void) => void
		shareEntity: (kind: string, workgroupId: number, entityTypeId: number, entityId: number, go: (error: Error) => void) => void
		getEntityPrivileges: (entityTypeId: number, entityId: number, go: (error: Error, privileges: EntityPrivilege[]) => void) => void
		unshareEntity: (kind: string, workgroupId: number, entityTypeId: number, entityId: number, go: (error: Error) => void) => void
		getEntityHistory: (entityTypeId: number, entityId: number, offset: number, limit: number, go: (error: Error, history: EntityHistory[]) => void) => void
	}

	// --- Messages ---

	interface PingIn {
		status: boolean
	}

	interface PingOut {
		status: boolean
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

	interface StartYarnClusterIn {
		cluster_name: string
		engine_id: number
		size: number
		memory: string
		username: string
	}

	interface StartYarnClusterOut {
		cluster_id: number
	}

	interface StopYarnClusterIn {
		cluster_id: number
	}

	interface StopYarnClusterOut {
	}

	interface GetClusterIn {
		cluster_id: number
	}

	interface GetClusterOut {
		cluster: Cluster
	}

	interface GetYarnClusterIn {
		cluster_id: number
	}

	interface GetYarnClusterOut {
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

	interface BuildModelIn {
		cluster_id: number
		dataset: string
		target_name: string
		max_run_time: number
	}

	interface BuildModelOut {
		model: Model
	}

	interface GetModelIn {
		model_id: number
	}

	interface GetModelOut {
		model: Model
	}

	interface GetModelsIn {
		offset: number
		limit: number
	}

	interface GetModelsOut {
		models: Model[]
	}

	interface GetClusterModelsIn {
		cluster_id: number
	}

	interface GetClusterModelsOut {
		models: Model[]
	}

	interface ImportModelFromClusterIn {
		cluster_id: number
		model_name: string
	}

	interface ImportModelFromClusterOut {
		model: Model
	}

	interface DeleteModelIn {
		model_id: number
	}

	interface DeleteModelOut {
	}

	interface StartScoringServiceIn {
		model_id: number
		port: number
	}

	interface StartScoringServiceOut {
		service: ScoringService
	}

	interface StopScoringServiceIn {
		service_id: number
	}

	interface StopScoringServiceOut {
	}

	interface GetScoringServiceIn {
		service_id: number
	}

	interface GetScoringServiceOut {
		service: ScoringService
	}

	interface GetScoringServicesIn {
		offset: number
		limit: number
	}

	interface GetScoringServicesOut {
		services: ScoringService[]
	}

	interface GetScoringServicesForModelIn {
		model_id: number
		offset: number
		limit: number
	}

	interface GetScoringServicesForModelOut {
		services: ScoringService[]
	}

	interface DeleteScoringServiceIn {
		service_id: number
	}

	interface DeleteScoringServiceOut {
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

	interface GetSupportedEntityTypesIn {
	}

	interface GetSupportedEntityTypesOut {
		entity_types: EntityType[]
	}

	interface GetSupportedPermissionsIn {
	}

	interface GetSupportedPermissionsOut {
		permissions: Permission[]
	}

	interface GetSupportedClusterTypesIn {
	}

	interface GetSupportedClusterTypesOut {
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

	interface LinkRoleAndPermissionsIn {
		role_id: number
		permission_ids: number[]
	}

	interface LinkRoleAndPermissionsOut {
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

	interface GetIdentititesForRoleIn {
		role_id: number
	}

	interface GetIdentititesForRoleOut {
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

	interface LinkIdentityAndWorkgroupIn {
		identity_id: number
		workgroup_id: number
	}

	interface LinkIdentityAndWorkgroupOut {
	}

	interface UnlinkIdentityAndWorkgroupIn {
		identity_id: number
		workgroup_id: number
	}

	interface UnlinkIdentityAndWorkgroupOut {
	}

	interface LinkIdentityAndRoleIn {
		identity_id: number
		role_id: number
	}

	interface LinkIdentityAndRoleOut {
	}

	interface UnlinkIdentityAndRoleIn {
		identity_id: number
		role_id: number
	}

	interface UnlinkIdentityAndRoleOut {
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

	interface GetEntityPrivilegesIn {
		entity_type_id: number
		entity_id: number
	}

	interface GetEntityPrivilegesOut {
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

	interface GetEntityHistoryIn {
		entity_type_id: number
		entity_id: number
		offset: number
		limit: number
	}

	interface GetEntityHistoryOut {
		history: EntityHistory[]
	}

	// --- Client Stub ---

	export function ping(status: boolean, go: (error: Error, status: boolean) => void): void {
		var req: PingIn = {
			status: status
		}
		Proxy.Call("Ping", req, function(error, data) {
			return error ? go(error, null) : go(null, (<PingOut>data).status)
		})

	}
	export function registerCluster(address: string, go: (error: Error, clusterId: number) => void): void {
		var req: RegisterClusterIn = {
			address: address
		}
		Proxy.Call("RegisterCluster", req, function(error, data) {
			return error ? go(error, null) : go(null, (<RegisterClusterOut>data).cluster_id)
		})

	}
	export function unregisterCluster(clusterId: number, go: (error: Error) => void): void {
		var req: UnregisterClusterIn = {
			cluster_id: clusterId
		}
		Proxy.Call("UnregisterCluster", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function startYarnCluster(clusterName: string, engineId: number, size: number, memory: string, username: string, go: (error: Error, clusterId: number) => void): void {
		var req: StartYarnClusterIn = {
			cluster_name: clusterName,
			engine_id: engineId,
			size: size,
			memory: memory,
			username: username
		}
		Proxy.Call("StartYarnCluster", req, function(error, data) {
			return error ? go(error, null) : go(null, (<StartYarnClusterOut>data).cluster_id)
		})

	}
	export function stopYarnCluster(clusterId: number, go: (error: Error) => void): void {
		var req: StopYarnClusterIn = {
			cluster_id: clusterId
		}
		Proxy.Call("StopYarnCluster", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function getCluster(clusterId: number, go: (error: Error, cluster: Cluster) => void): void {
		var req: GetClusterIn = {
			cluster_id: clusterId
		}
		Proxy.Call("GetCluster", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetClusterOut>data).cluster)
		})

	}
	export function getYarnCluster(clusterId: number, go: (error: Error, cluster: YarnCluster) => void): void {
		var req: GetYarnClusterIn = {
			cluster_id: clusterId
		}
		Proxy.Call("GetYarnCluster", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetYarnClusterOut>data).cluster)
		})

	}
	export function getClusters(offset: number, limit: number, go: (error: Error, clusters: Cluster[]) => void): void {
		var req: GetClustersIn = {
			offset: offset,
			limit: limit
		}
		Proxy.Call("GetClusters", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetClustersOut>data).clusters)
		})

	}
	export function getClusterStatus(clusterId: number, go: (error: Error, clusterStatus: ClusterStatus) => void): void {
		var req: GetClusterStatusIn = {
			cluster_id: clusterId
		}
		Proxy.Call("GetClusterStatus", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetClusterStatusOut>data).cluster_status)
		})

	}
	export function deleteCluster(clusterId: number, go: (error: Error) => void): void {
		var req: DeleteClusterIn = {
			cluster_id: clusterId
		}
		Proxy.Call("DeleteCluster", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function getJob(clusterId: number, jobName: string, go: (error: Error, job: Job) => void): void {
		var req: GetJobIn = {
			cluster_id: clusterId,
			job_name: jobName
		}
		Proxy.Call("GetJob", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetJobOut>data).job)
		})

	}
	export function getJobs(clusterId: number, go: (error: Error, jobs: Job[]) => void): void {
		var req: GetJobsIn = {
			cluster_id: clusterId
		}
		Proxy.Call("GetJobs", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetJobsOut>data).jobs)
		})

	}
	export function buildModel(clusterId: number, dataset: string, targetName: string, maxRunTime: number, go: (error: Error, model: Model) => void): void {
		var req: BuildModelIn = {
			cluster_id: clusterId,
			dataset: dataset,
			target_name: targetName,
			max_run_time: maxRunTime
		}
		Proxy.Call("BuildModel", req, function(error, data) {
			return error ? go(error, null) : go(null, (<BuildModelOut>data).model)
		})

	}
	export function getModel(modelId: number, go: (error: Error, model: Model) => void): void {
		var req: GetModelIn = {
			model_id: modelId
		}
		Proxy.Call("GetModel", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetModelOut>data).model)
		})

	}
	export function getModels(offset: number, limit: number, go: (error: Error, models: Model[]) => void): void {
		var req: GetModelsIn = {
			offset: offset,
			limit: limit
		}
		Proxy.Call("GetModels", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetModelsOut>data).models)
		})

	}
	export function getClusterModels(clusterId: number, go: (error: Error, models: Model[]) => void): void {
		var req: GetClusterModelsIn = {
			cluster_id: clusterId
		}
		Proxy.Call("GetClusterModels", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetClusterModelsOut>data).models)
		})

	}
	export function importModelFromCluster(clusterId: number, modelName: string, go: (error: Error, model: Model) => void): void {
		var req: ImportModelFromClusterIn = {
			cluster_id: clusterId,
			model_name: modelName
		}
		Proxy.Call("ImportModelFromCluster", req, function(error, data) {
			return error ? go(error, null) : go(null, (<ImportModelFromClusterOut>data).model)
		})

	}
	export function deleteModel(modelId: number, go: (error: Error) => void): void {
		var req: DeleteModelIn = {
			model_id: modelId
		}
		Proxy.Call("DeleteModel", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function startScoringService(modelId: number, port: number, go: (error: Error, service: ScoringService) => void): void {
		var req: StartScoringServiceIn = {
			model_id: modelId,
			port: port
		}
		Proxy.Call("StartScoringService", req, function(error, data) {
			return error ? go(error, null) : go(null, (<StartScoringServiceOut>data).service)
		})

	}
	export function stopScoringService(serviceId: number, go: (error: Error) => void): void {
		var req: StopScoringServiceIn = {
			service_id: serviceId
		}
		Proxy.Call("StopScoringService", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function getScoringService(serviceId: number, go: (error: Error, service: ScoringService) => void): void {
		var req: GetScoringServiceIn = {
			service_id: serviceId
		}
		Proxy.Call("GetScoringService", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetScoringServiceOut>data).service)
		})

	}
	export function getScoringServices(offset: number, limit: number, go: (error: Error, services: ScoringService[]) => void): void {
		var req: GetScoringServicesIn = {
			offset: offset,
			limit: limit
		}
		Proxy.Call("GetScoringServices", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetScoringServicesOut>data).services)
		})

	}
	export function getScoringServicesForModel(modelId: number, offset: number, limit: number, go: (error: Error, services: ScoringService[]) => void): void {
		var req: GetScoringServicesForModelIn = {
			model_id: modelId,
			offset: offset,
			limit: limit
		}
		Proxy.Call("GetScoringServicesForModel", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetScoringServicesForModelOut>data).services)
		})

	}
	export function deleteScoringService(serviceId: number, go: (error: Error) => void): void {
		var req: DeleteScoringServiceIn = {
			service_id: serviceId
		}
		Proxy.Call("DeleteScoringService", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function addEngine(engineName: string, enginePath: string, go: (error: Error, engineId: number) => void): void {
		var req: AddEngineIn = {
			engine_name: engineName,
			engine_path: enginePath
		}
		Proxy.Call("AddEngine", req, function(error, data) {
			return error ? go(error, null) : go(null, (<AddEngineOut>data).engine_id)
		})

	}
	export function getEngine(engineId: number, go: (error: Error, engine: Engine) => void): void {
		var req: GetEngineIn = {
			engine_id: engineId
		}
		Proxy.Call("GetEngine", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetEngineOut>data).engine)
		})

	}
	export function getEngines(go: (error: Error, engines: Engine[]) => void): void {
		var req: GetEnginesIn = {
		}
		Proxy.Call("GetEngines", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetEnginesOut>data).engines)
		})

	}
	export function deleteEngine(engineId: number, go: (error: Error) => void): void {
		var req: DeleteEngineIn = {
			engine_id: engineId
		}
		Proxy.Call("DeleteEngine", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function getSupportedEntityTypes(go: (error: Error, entityTypes: EntityType[]) => void): void {
		var req: GetSupportedEntityTypesIn = {
		}
		Proxy.Call("GetSupportedEntityTypes", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetSupportedEntityTypesOut>data).entity_types)
		})

	}
	export function getSupportedPermissions(go: (error: Error, permissions: Permission[]) => void): void {
		var req: GetSupportedPermissionsIn = {
		}
		Proxy.Call("GetSupportedPermissions", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetSupportedPermissionsOut>data).permissions)
		})

	}
	export function getSupportedClusterTypes(go: (error: Error, clusterTypes: ClusterType[]) => void): void {
		var req: GetSupportedClusterTypesIn = {
		}
		Proxy.Call("GetSupportedClusterTypes", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetSupportedClusterTypesOut>data).cluster_types)
		})

	}
	export function getPermissionsForRole(roleId: number, go: (error: Error, permissions: Permission[]) => void): void {
		var req: GetPermissionsForRoleIn = {
			role_id: roleId
		}
		Proxy.Call("GetPermissionsForRole", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetPermissionsForRoleOut>data).permissions)
		})

	}
	export function getPermissionsForIdentity(identityId: number, go: (error: Error, permissions: Permission[]) => void): void {
		var req: GetPermissionsForIdentityIn = {
			identity_id: identityId
		}
		Proxy.Call("GetPermissionsForIdentity", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetPermissionsForIdentityOut>data).permissions)
		})

	}
	export function createRole(name: string, description: string, go: (error: Error, roleId: number) => void): void {
		var req: CreateRoleIn = {
			name: name,
			description: description
		}
		Proxy.Call("CreateRole", req, function(error, data) {
			return error ? go(error, null) : go(null, (<CreateRoleOut>data).role_id)
		})

	}
	export function getRoles(offset: number, limit: number, go: (error: Error, roles: Role[]) => void): void {
		var req: GetRolesIn = {
			offset: offset,
			limit: limit
		}
		Proxy.Call("GetRoles", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetRolesOut>data).roles)
		})

	}
	export function getRolesForIdentity(identityId: number, go: (error: Error, roles: Role[]) => void): void {
		var req: GetRolesForIdentityIn = {
			identity_id: identityId
		}
		Proxy.Call("GetRolesForIdentity", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetRolesForIdentityOut>data).roles)
		})

	}
	export function getRole(roleId: number, go: (error: Error, role: Role) => void): void {
		var req: GetRoleIn = {
			role_id: roleId
		}
		Proxy.Call("GetRole", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetRoleOut>data).role)
		})

	}
	export function getRoleByName(name: string, go: (error: Error, role: Role) => void): void {
		var req: GetRoleByNameIn = {
			name: name
		}
		Proxy.Call("GetRoleByName", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetRoleByNameOut>data).role)
		})

	}
	export function updateRole(roleId: number, name: string, description: string, go: (error: Error) => void): void {
		var req: UpdateRoleIn = {
			role_id: roleId,
			name: name,
			description: description
		}
		Proxy.Call("UpdateRole", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function linkRoleAndPermissions(roleId: number, permissionIds: number[], go: (error: Error) => void): void {
		var req: LinkRoleAndPermissionsIn = {
			role_id: roleId,
			permission_ids: permissionIds
		}
		Proxy.Call("LinkRoleAndPermissions", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function deleteRole(roleId: number, go: (error: Error) => void): void {
		var req: DeleteRoleIn = {
			role_id: roleId
		}
		Proxy.Call("DeleteRole", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function createWorkgroup(name: string, description: string, go: (error: Error, workgroupId: number) => void): void {
		var req: CreateWorkgroupIn = {
			name: name,
			description: description
		}
		Proxy.Call("CreateWorkgroup", req, function(error, data) {
			return error ? go(error, null) : go(null, (<CreateWorkgroupOut>data).workgroup_id)
		})

	}
	export function getWorkgroups(offset: number, limit: number, go: (error: Error, workgroups: Workgroup[]) => void): void {
		var req: GetWorkgroupsIn = {
			offset: offset,
			limit: limit
		}
		Proxy.Call("GetWorkgroups", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetWorkgroupsOut>data).workgroups)
		})

	}
	export function getWorkgroupsForIdentity(identityId: number, go: (error: Error, workgroups: Workgroup[]) => void): void {
		var req: GetWorkgroupsForIdentityIn = {
			identity_id: identityId
		}
		Proxy.Call("GetWorkgroupsForIdentity", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetWorkgroupsForIdentityOut>data).workgroups)
		})

	}
	export function getWorkgroup(workgroupId: number, go: (error: Error, workgroup: Workgroup) => void): void {
		var req: GetWorkgroupIn = {
			workgroup_id: workgroupId
		}
		Proxy.Call("GetWorkgroup", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetWorkgroupOut>data).workgroup)
		})

	}
	export function getWorkgroupByName(name: string, go: (error: Error, workgroup: Workgroup) => void): void {
		var req: GetWorkgroupByNameIn = {
			name: name
		}
		Proxy.Call("GetWorkgroupByName", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetWorkgroupByNameOut>data).workgroup)
		})

	}
	export function updateWorkgroup(workgroupId: number, name: string, description: string, go: (error: Error) => void): void {
		var req: UpdateWorkgroupIn = {
			workgroup_id: workgroupId,
			name: name,
			description: description
		}
		Proxy.Call("UpdateWorkgroup", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function deleteWorkgroup(workgroupId: number, go: (error: Error) => void): void {
		var req: DeleteWorkgroupIn = {
			workgroup_id: workgroupId
		}
		Proxy.Call("DeleteWorkgroup", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function createIdentity(name: string, password: string, go: (error: Error, identityId: number) => void): void {
		var req: CreateIdentityIn = {
			name: name,
			password: password
		}
		Proxy.Call("CreateIdentity", req, function(error, data) {
			return error ? go(error, null) : go(null, (<CreateIdentityOut>data).identity_id)
		})

	}
	export function getIdentities(offset: number, limit: number, go: (error: Error, identities: Identity[]) => void): void {
		var req: GetIdentitiesIn = {
			offset: offset,
			limit: limit
		}
		Proxy.Call("GetIdentities", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetIdentitiesOut>data).identities)
		})

	}
	export function getIdentitiesForWorkgroup(workgroupId: number, go: (error: Error, identities: Identity[]) => void): void {
		var req: GetIdentitiesForWorkgroupIn = {
			workgroup_id: workgroupId
		}
		Proxy.Call("GetIdentitiesForWorkgroup", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetIdentitiesForWorkgroupOut>data).identities)
		})

	}
	export function getIdentititesForRole(roleId: number, go: (error: Error, identities: Identity[]) => void): void {
		var req: GetIdentititesForRoleIn = {
			role_id: roleId
		}
		Proxy.Call("GetIdentititesForRole", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetIdentititesForRoleOut>data).identities)
		})

	}
	export function getIdentity(identityId: number, go: (error: Error, identity: Identity) => void): void {
		var req: GetIdentityIn = {
			identity_id: identityId
		}
		Proxy.Call("GetIdentity", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetIdentityOut>data).identity)
		})

	}
	export function getIdentityByName(name: string, go: (error: Error, identity: Identity) => void): void {
		var req: GetIdentityByNameIn = {
			name: name
		}
		Proxy.Call("GetIdentityByName", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetIdentityByNameOut>data).identity)
		})

	}
	export function linkIdentityAndWorkgroup(identityId: number, workgroupId: number, go: (error: Error) => void): void {
		var req: LinkIdentityAndWorkgroupIn = {
			identity_id: identityId,
			workgroup_id: workgroupId
		}
		Proxy.Call("LinkIdentityAndWorkgroup", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function unlinkIdentityAndWorkgroup(identityId: number, workgroupId: number, go: (error: Error) => void): void {
		var req: UnlinkIdentityAndWorkgroupIn = {
			identity_id: identityId,
			workgroup_id: workgroupId
		}
		Proxy.Call("UnlinkIdentityAndWorkgroup", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function linkIdentityAndRole(identityId: number, roleId: number, go: (error: Error) => void): void {
		var req: LinkIdentityAndRoleIn = {
			identity_id: identityId,
			role_id: roleId
		}
		Proxy.Call("LinkIdentityAndRole", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function unlinkIdentityAndRole(identityId: number, roleId: number, go: (error: Error) => void): void {
		var req: UnlinkIdentityAndRoleIn = {
			identity_id: identityId,
			role_id: roleId
		}
		Proxy.Call("UnlinkIdentityAndRole", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function deactivateIdentity(identityId: number, go: (error: Error) => void): void {
		var req: DeactivateIdentityIn = {
			identity_id: identityId
		}
		Proxy.Call("DeactivateIdentity", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function shareEntity(kind: string, workgroupId: number, entityTypeId: number, entityId: number, go: (error: Error) => void): void {
		var req: ShareEntityIn = {
			kind: kind,
			workgroup_id: workgroupId,
			entity_type_id: entityTypeId,
			entity_id: entityId
		}
		Proxy.Call("ShareEntity", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function getEntityPrivileges(entityTypeId: number, entityId: number, go: (error: Error, privileges: EntityPrivilege[]) => void): void {
		var req: GetEntityPrivilegesIn = {
			entity_type_id: entityTypeId,
			entity_id: entityId
		}
		Proxy.Call("GetEntityPrivileges", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetEntityPrivilegesOut>data).privileges)
		})

	}
	export function unshareEntity(kind: string, workgroupId: number, entityTypeId: number, entityId: number, go: (error: Error) => void): void {
		var req: UnshareEntityIn = {
			kind: kind,
			workgroup_id: workgroupId,
			entity_type_id: entityTypeId,
			entity_id: entityId
		}
		Proxy.Call("UnshareEntity", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function getEntityHistory(entityTypeId: number, entityId: number, offset: number, limit: number, go: (error: Error, history: EntityHistory[]) => void): void {
		var req: GetEntityHistoryIn = {
			entity_type_id: entityTypeId,
			entity_id: entityId,
			offset: offset,
			limit: limit
		}
		Proxy.Call("GetEntityHistory", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetEntityHistoryOut>data).history)
		})

	}
}

