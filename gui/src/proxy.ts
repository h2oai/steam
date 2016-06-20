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
		getModelFromCluster: (clusterId: number, modelName: string, go: (error: Error, model: Model) => void) => void
		deleteModel: (modelId: number, go: (error: Error) => void) => void
		startScoringService: (modelId: number, port: number, go: (error: Error, service: ScoringService) => void) => void
		stopScoringService: (serviceId: number, go: (error: Error) => void) => void
		getScoringService: (serviceId: number, go: (error: Error, service: ScoringService) => void) => void
		getScoringServices: (offset: number, limit: number, go: (error: Error, services: ScoringService[]) => void) => void
		deleteScoringService: (serviceId: number, go: (error: Error) => void) => void
		addEngine: (engineName: string, enginePath: string, go: (error: Error, engineId: number) => void) => void
		getEngine: (engineId: number, go: (error: Error, engine: Engine) => void) => void
		getEngines: (go: (error: Error, engines: Engine[]) => void) => void
		deleteEngine: (engineId: number, go: (error: Error) => void) => void
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

	interface GetModelFromClusterIn {
		cluster_id: number
		model_name: string
	}

	interface GetModelFromClusterOut {
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
	export function getModelFromCluster(clusterId: number, modelName: string, go: (error: Error, model: Model) => void): void {
		var req: GetModelFromClusterIn = {
			cluster_id: clusterId,
			model_name: modelName
		}
		Proxy.Call("GetModelFromCluster", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetModelFromClusterOut>data).model)
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
}

