/// <reference path="xhr.ts" />
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
		startCloud: (name: string, size: number, useKerberos: boolean, username: string, keytab: string, go: (error: Error, applicationID: string) => void) => void
		stopCloud: (name: string, useKerberos: boolean, applicationID: string, username: string, keytab: string, go: (error: Error) => void) => void
		getCloud: (address: string, go: (error: Error, cloud: Cloud) => void) => void
		buildAutoML: (address: string, dataset: string, targetName: string, maxRunTime: number, go: (error: Error, modelID: string) => void) => void
		deployPojo: (address: string, javaModelPath: string, genModelPath: string, go: (error: Error) => void) => void
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
		name: string
		size: number
		use_kerberos: boolean
		username: string
		keytab: string
	}

	interface StartCloudOut {
		application_id: string
	}

	interface StopCloudIn {
		name: string
		use_kerberos: boolean
		application_id: string
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
		max_run_time: number
	}

	interface BuildAutoMLOut {
		model_id: string
	}

	interface DeployPojoIn {
		address: string
		java_model_path: string
		gen_model_path: string
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
	export function startCloud(name: string, size: number, useKerberos: boolean, username: string, keytab: string, go: (error: Error, applicationID: string) => void): void {
		var req: StartCloudIn = {
			name: name,
			size: size,
			use_kerberos: useKerberos,
			username: username,
			keytab: keytab
		}
		Proxy.Call("StartCloud", req, function(error, data) {
			return error ? go(error, null) : go(null, (<StartCloudOut>data).application_id)
		})

	}
	export function stopCloud(name: string, useKerberos: boolean, applicationID: string, username: string, keytab: string, go: (error: Error) => void): void {
		var req: StopCloudIn = {
			name: name,
			use_kerberos: useKerberos,
			application_id: applicationID,
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
	export function buildAutoML(address: string, dataset: string, targetName: string, maxRunTime: number, go: (error: Error, modelID: string) => void): void {
		var req: BuildAutoMLIn = {
			address: address,
			dataset: dataset,
			target_name: targetName,
			max_run_time: maxRunTime
		}
		Proxy.Call("BuildAutoML", req, function(error, data) {
			return error ? go(error, null) : go(null, (<BuildAutoMLOut>data).model_id)
		})

	}
	export function deployPojo(address: string, javaModelPath: string, genModelPath: string, go: (error: Error) => void): void {
		var req: DeployPojoIn = {
			address: address,
			java_model_path: javaModelPath,
			gen_model_path: genModelPath
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

