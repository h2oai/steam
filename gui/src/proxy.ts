// ----------------------------------
// --- Generated with go:generate ---
// ---        DO NOT EDIT         ---
// ----------------------------------

module Proxy {

	// --- Aliases ---

	export type CloudState = string
	export type ScoringServiceState = string
	export type Timestamp = number

	// --- Consts ---

	export var CloudStarted: CloudState = "Started"
	export var CloudHealthy: CloudState = "Healthy"
	export var CloudStopped: CloudState = "Stopped"
	export var CloudUnknown: CloudState = "Unknown"


	export var ScoringServiceStarted: ScoringServiceState = "Started"
	export var ScoringServiceStopped: ScoringServiceState = "Stopped"


	// --- Types ---

	export interface Cloud {
		created_at: Timestamp
		name: string
		engine_name: string
		engine_version: string
		size: number
		memory: string
		total_cores: number
		allowed_cores: number
		state: CloudState
		address: string
		username: string
		application_id: string
	}

	export interface Model {
		name: string
		cloud_name: string
		algo: string
		dataset: string
		target_name: string
		max_runtime: number
		java_model_path: string
		gen_model_path: string
		created_at: Timestamp
	}

	export interface ScoringService {
		model_name: string
		address: string
		port: number
		state: ScoringServiceState
		pid: number
		created_at: Timestamp
	}

	export interface Engine {
		name: string
		path: string
		created_at: Timestamp
	}

	// --- Contract ---

	export interface Service {
		ping: (status: boolean, go: (error: Error, status: boolean) => void) => void
		startCloud: (cloudName: string, engineName: string, size: number, memory: string, username: string, go: (error: Error, cloud: Cloud) => void) => void
		stopCloud: (cloudName: string, go: (error: Error) => void) => void
		getCloud: (cloudName: string, go: (error: Error, cloud: Cloud) => void) => void
		getClouds: (go: (error: Error, clouds: Cloud[]) => void) => void
		getCloudStatus: (cloudName: string, go: (error: Error, cloud: Cloud) => void) => void
		deleteCloud: (cloudName: string, go: (error: Error) => void) => void
		buildModel: (cloudName: string, dataset: string, targetName: string, maxRunTime: number, go: (error: Error, model: Model) => void) => void
		getModel: (modelName: string, go: (error: Error, model: Model) => void) => void
		getModels: (go: (error: Error, models: Model[]) => void) => void
		getCloudModels: (cloudName: string, go: (error: Error, models: Model[]) => void) => void
		getModelFromCloud: (cloudName: string, modelName: string, go: (error: Error, model: Model) => void) => void
		deleteModel: (modelName: string, go: (error: Error) => void) => void
		startScoringService: (modelName: string, port: number, go: (error: Error, service: ScoringService) => void) => void
		stopScoringService: (modelName: string, port: number, go: (error: Error) => void) => void
		getScoringService: (modelName: string, go: (error: Error, service: ScoringService) => void) => void
		getScoringServices: (go: (error: Error, services: ScoringService[]) => void) => void
		deleteScoringService: (modelName: string, port: number, go: (error: Error) => void) => void
		addEngine: (engineName: string, enginePath: string, go: (error: Error) => void) => void
		getEngine: (engineName: string, go: (error: Error, engine: Engine) => void) => void
		getEngines: (go: (error: Error, engines: Engine[]) => void) => void
		deleteEngine: (engineName: string, go: (error: Error) => void) => void
	}

	// --- Messages ---

	interface PingIn {
		status: boolean
	}

	interface PingOut {
		status: boolean
	}

	interface StartCloudIn {
		cloud_name: string
		engine_name: string
		size: number
		memory: string
		username: string
	}

	interface StartCloudOut {
		cloud: Cloud
	}

	interface StopCloudIn {
		cloud_name: string
	}

	interface StopCloudOut {
	}

	interface GetCloudIn {
		cloud_name: string
	}

	interface GetCloudOut {
		cloud: Cloud
	}

	interface GetCloudsIn {
	}

	interface GetCloudsOut {
		clouds: Cloud[]
	}

	interface GetCloudStatusIn {
		cloud_name: string
	}

	interface GetCloudStatusOut {
		cloud: Cloud
	}

	interface DeleteCloudIn {
		cloud_name: string
	}

	interface DeleteCloudOut {
	}

	interface BuildModelIn {
		cloud_name: string
		dataset: string
		target_name: string
		max_run_time: number
	}

	interface BuildModelOut {
		model: Model
	}

	interface GetModelIn {
		model_name: string
	}

	interface GetModelOut {
		model: Model
	}

	interface GetModelsIn {
	}

	interface GetModelsOut {
		models: Model[]
	}

	interface GetCloudModelsIn {
		cloud_name: string
	}

	interface GetCloudModelsOut {
		models: Model[]
	}

	interface GetModelFromCloudIn {
		cloud_name: string
		model_name: string
	}

	interface GetModelFromCloudOut {
		model: Model
	}

	interface DeleteModelIn {
		model_name: string
	}

	interface DeleteModelOut {
	}

	interface StartScoringServiceIn {
		model_name: string
		port: number
	}

	interface StartScoringServiceOut {
		service: ScoringService
	}

	interface StopScoringServiceIn {
		model_name: string
		port: number
	}

	interface StopScoringServiceOut {
	}

	interface GetScoringServiceIn {
		model_name: string
	}

	interface GetScoringServiceOut {
		service: ScoringService
	}

	interface GetScoringServicesIn {
	}

	interface GetScoringServicesOut {
		services: ScoringService[]
	}

	interface DeleteScoringServiceIn {
		model_name: string
		port: number
	}

	interface DeleteScoringServiceOut {
	}

	interface AddEngineIn {
		engine_name: string
		engine_path: string
	}

	interface AddEngineOut {
	}

	interface GetEngineIn {
		engine_name: string
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
		engine_name: string
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
	export function startCloud(cloudName: string, engineName: string, size: number, memory: string, username: string, go: (error: Error, cloud: Cloud) => void): void {
		var req: StartCloudIn = {
			cloud_name: cloudName,
			engine_name: engineName,
			size: size,
			memory: memory,
			username: username
		}
		Proxy.Call("StartCloud", req, function(error, data) {
			return error ? go(error, null) : go(null, (<StartCloudOut>data).cloud)
		})

	}
	export function stopCloud(cloudName: string, go: (error: Error) => void): void {
		var req: StopCloudIn = {
			cloud_name: cloudName
		}
		Proxy.Call("StopCloud", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function getCloud(cloudName: string, go: (error: Error, cloud: Cloud) => void): void {
		var req: GetCloudIn = {
			cloud_name: cloudName
		}
		Proxy.Call("GetCloud", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetCloudOut>data).cloud)
		})

	}
	export function getClouds(go: (error: Error, clouds: Cloud[]) => void): void {
		var req: GetCloudsIn = {
		}
		Proxy.Call("GetClouds", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetCloudsOut>data).clouds)
		})

	}
	export function getCloudStatus(cloudName: string, go: (error: Error, cloud: Cloud) => void): void {
		var req: GetCloudStatusIn = {
			cloud_name: cloudName
		}
		Proxy.Call("GetCloudStatus", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetCloudStatusOut>data).cloud)
		})

	}
	export function deleteCloud(cloudName: string, go: (error: Error) => void): void {
		var req: DeleteCloudIn = {
			cloud_name: cloudName
		}
		Proxy.Call("DeleteCloud", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function buildModel(cloudName: string, dataset: string, targetName: string, maxRunTime: number, go: (error: Error, model: Model) => void): void {
		var req: BuildModelIn = {
			cloud_name: cloudName,
			dataset: dataset,
			target_name: targetName,
			max_run_time: maxRunTime
		}
		Proxy.Call("BuildModel", req, function(error, data) {
			return error ? go(error, null) : go(null, (<BuildModelOut>data).model)
		})

	}
	export function getModel(modelName: string, go: (error: Error, model: Model) => void): void {
		var req: GetModelIn = {
			model_name: modelName
		}
		Proxy.Call("GetModel", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetModelOut>data).model)
		})

	}
	export function getModels(go: (error: Error, models: Model[]) => void): void {
		var req: GetModelsIn = {
		}
		Proxy.Call("GetModels", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetModelsOut>data).models)
		})

	}
	export function getCloudModels(cloudName: string, go: (error: Error, models: Model[]) => void): void {
		var req: GetCloudModelsIn = {
			cloud_name: cloudName
		}
		Proxy.Call("GetCloudModels", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetCloudModelsOut>data).models)
		})

	}
	export function getModelFromCloud(cloudName: string, modelName: string, go: (error: Error, model: Model) => void): void {
		var req: GetModelFromCloudIn = {
			cloud_name: cloudName,
			model_name: modelName
		}
		Proxy.Call("GetModelFromCloud", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetModelFromCloudOut>data).model)
		})

	}
	export function deleteModel(modelName: string, go: (error: Error) => void): void {
		var req: DeleteModelIn = {
			model_name: modelName
		}
		Proxy.Call("DeleteModel", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function startScoringService(modelName: string, port: number, go: (error: Error, service: ScoringService) => void): void {
		var req: StartScoringServiceIn = {
			model_name: modelName,
			port: port
		}
		Proxy.Call("StartScoringService", req, function(error, data) {
			return error ? go(error, null) : go(null, (<StartScoringServiceOut>data).service)
		})

	}
	export function stopScoringService(modelName: string, port: number, go: (error: Error) => void): void {
		var req: StopScoringServiceIn = {
			model_name: modelName,
			port: port
		}
		Proxy.Call("StopScoringService", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function getScoringService(modelName: string, go: (error: Error, service: ScoringService) => void): void {
		var req: GetScoringServiceIn = {
			model_name: modelName
		}
		Proxy.Call("GetScoringService", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetScoringServiceOut>data).service)
		})

	}
	export function getScoringServices(go: (error: Error, services: ScoringService[]) => void): void {
		var req: GetScoringServicesIn = {
		}
		Proxy.Call("GetScoringServices", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetScoringServicesOut>data).services)
		})

	}
	export function deleteScoringService(modelName: string, port: number, go: (error: Error) => void): void {
		var req: DeleteScoringServiceIn = {
			model_name: modelName,
			port: port
		}
		Proxy.Call("DeleteScoringService", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function addEngine(engineName: string, enginePath: string, go: (error: Error) => void): void {
		var req: AddEngineIn = {
			engine_name: engineName,
			engine_path: enginePath
		}
		Proxy.Call("AddEngine", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function getEngine(engineName: string, go: (error: Error, engine: Engine) => void): void {
		var req: GetEngineIn = {
			engine_name: engineName
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
	export function deleteEngine(engineName: string, go: (error: Error) => void): void {
		var req: DeleteEngineIn = {
			engine_name: engineName
		}
		Proxy.Call("DeleteEngine", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
}

