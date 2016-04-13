// ----------------------------------
// --- Generated with go:generate ---
// ---        DO NOT EDIT         ---
// ----------------------------------

module Proxy {

	// --- Aliases ---

	export type CloudState = string
	export type Timestamp = number

	// --- Types ---

	export interface Cloud {
		name: string
		pack: string
		state: CloudState
	}

	export interface CloudModelSynopsis {
		algorithm: string
		algorithm_full_name: string
		frame_name: string
		model_name: string
		response_column_name: string
		modified_at: Timestamp
	}

	// --- Contract ---

	export interface Service {
		ping: (status: boolean, go: (error: Error, status: boolean) => void) => void
		startCloud: (size: number, kerberos: boolean, name: string, username: string, keytab: string, go: (error: Error, apID: string) => void) => void
		stopCloud: (kerberos: boolean, name: string, id: string, username: string, keytab: string, go: (error: Error) => void) => void
		getCloud: (address: string, go: (error: Error, cloud: Cloud) => void) => void
		buildAutoML: (address: string, dataset: string, targetName: string, maxTime: number, go: (error: Error, modelID: string) => void) => void
		deployPojo: (address: string, javaModel: string, jar: string, go: (error: Error) => void) => void
		shutdown: (address: string, go: (error: Error) => void) => void
	}

	// --- Messages ---

	interface PingIn {
		status: boolean
	}

	interface PingOut {
		status: boolean
	}

	interface StartCloudIn {
		size: number
		kerberos: boolean
		name: string
		username: string
		keytab: string
	}

	interface StartCloudOut {
		ap_id: string
	}

	interface StopCloudIn {
		kerberos: boolean
		name: string
		id: string
		username: string
		keytab: string
	}

	interface StopCloudOut {
	}

	interface GetCloudIn {
		address: string
	}

	interface GetCloudOut {
		cloud: Cloud
	}

	interface BuildAutoMLIn {
		address: string
		dataset: string
		target_name: string
		max_time: number
	}

	interface BuildAutoMLOut {
		model_id: string
	}

	interface DeployPojoIn {
		address: string
		java_model: string
		jar: string
	}

	interface DeployPojoOut {
	}

	interface ShutdownIn {
		address: string
	}

	interface ShutdownOut {
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
	export function startCloud(size: number, kerberos: boolean, name: string, username: string, keytab: string, go: (error: Error, apID: string) => void): void {
		var req: StartCloudIn = {
			size: size,
			kerberos: kerberos,
			name: name,
			username: username,
			keytab: keytab
		}
		Proxy.Call("StartCloud", req, function(error, data) {
			return error ? go(error, null) : go(null, (<StartCloudOut>data).apID)
		})

	}
	export function stopCloud(kerberos: boolean, name: string, id: string, username: string, keytab: string, go: (error: Error) => void): void {
		var req: StopCloudIn = {
			kerberos: kerberos,
			name: name,
			id: id,
			username: username,
			keytab: keytab
		}
		Proxy.Call("StopCloud", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function getCloud(address: string, go: (error: Error, cloud: Cloud) => void): void {
		var req: GetCloudIn = {
			address: address
		}
		Proxy.Call("GetCloud", req, function(error, data) {
			return error ? go(error, null) : go(null, (<GetCloudOut>data).cloud)
		})

	}
	export function buildAutoML(address: string, dataset: string, targetName: string, maxTime: number, go: (error: Error, modelID: string) => void): void {
		var req: BuildAutoMLIn = {
			address: address,
			dataset: dataset,
			target_name: targetName,
			max_time: maxTime
		}
		Proxy.Call("BuildAutoML", req, function(error, data) {
			return error ? go(error, null) : go(null, (<BuildAutoMLOut>data).modelID)
		})

	}
	export function deployPojo(address: string, javaModel: string, jar: string, go: (error: Error) => void): void {
		var req: DeployPojoIn = {
			address: address,
			java_model: javaModel,
			jar: jar
		}
		Proxy.Call("DeployPojo", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
	export function shutdown(address: string, go: (error: Error) => void): void {
		var req: ShutdownIn = {
			address: address
		}
		Proxy.Call("Shutdown", req, function(error, data) {
			return error ? go(error) : go(null)
		})

	}
}

