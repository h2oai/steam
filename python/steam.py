// ------------------------------
// --- This is generated code ---
// ---      DO NOT EDIT       ---
// ------------------------------


import httplib
import base64
import string
import json
import sys
import logging
from collections import namedtuple

class RPCError(Exception):
	def __init__(self, value):
		self.value = value
	def __str__(self):
		return repr(self.value)

class HTTPConnection:
	def __init__(self, host, port, username, password):
		self.host = host
		self.port = port
		self.username = username
		self.password = password
		self.uid = 0

	def call(self, method, params):
		self.uid = self.uid + 1
		request = {
			'id': self.uid,
			'method': 'web.' + method,
			'params': [params]
		}
		payload = json.dumps(request)

		ws = httplib.HTTP(self.host, self.port)
		ws.putrequest("POST", '/web')

		ws.putheader("Host", self.host)
		ws.putheader("User-Agent", "Steam Python Client")
		ws.putheader("Content-type", "application/json; charset=\"UTF-8\"")
		ws.putheader("Content-length", "%d" % len(payload))
		auth = base64.encodestring('%s:%s' % (self.username, self.password)).replace('\n', '')
		ws.putheader("Authorization", "Basic %s" % auth)
		ws.endheaders()

		ws.send(payload)

		logging.info('%s@%s:%d %s(%s)', self.username, self.host, self.port, method, json.dumps(params))

		code, status, header = ws.getreply()
		reply = ws.getfile().read()

		# print 'code:', code
		# print 'status:', status
		# print 'reply:', reply

		if code != 200:
			logging.exception('%s %s %s', code, status, reply)
			raise RPCError(reply)

		response = json.loads(reply)
		error = response['error']

		if error is None:
			result = response['result']
			logging.info('%s %s %s', code, status, json.dumps(result))
			return result
		else:
			logging.exception('%s %s %s', code, status, error)
			raise RPCError(error)

class View(object):
	def __init__(self, d):
		self.__dict__ = d
	def __str__(self):
		return json.dumps(self.__dict__)

class RPCClient:
	def __init__(self, connection):
		self.connection = connection
	
	
	
	def ping_server(self, input):
		"""
		Ping the Steam server

		Parameters:
		input: Message to send (string)

		Returns:
		output: Echoed message (string)
		"""
		request = {
			'input': input
		}
		response = self.connection.call("PingServer", request)
		return response['output']
	
	def register_cluster(self, address):
		"""
		Connect to a cluster

		Parameters:
		address: No description available (string)

		Returns:
		cluster_id: No description available (int64)
		"""
		request = {
			'address': address
		}
		response = self.connection.call("RegisterCluster", request)
		return response['cluster_id']
	
	def unregister_cluster(self, cluster_id):
		"""
		Disconnect from a cluster

		Parameters:
		cluster_id: No description available (int64)

		Returns:None
		"""
		request = {
			'cluster_id': cluster_id
		}
		response = self.connection.call("UnregisterCluster", request)
		return 
	
	def start_cluster_on_yarn(self, cluster_name, engine_id, size, memory, username):
		"""
		Start a cluster using Yarn

		Parameters:
		cluster_name: No description available (string)
		engine_id: No description available (int64)
		size: No description available (int)
		memory: No description available (string)
		username: No description available (string)

		Returns:
		cluster_id: No description available (int64)
		"""
		request = {
			'cluster_name': cluster_name
			'engine_id': engine_id
			'size': size
			'memory': memory
			'username': username
		}
		response = self.connection.call("StartClusterOnYarn", request)
		return response['cluster_id']
	
	def stop_cluster_on_yarn(self, cluster_id):
		"""
		Stop a cluster using Yarn

		Parameters:
		cluster_id: No description available (int64)

		Returns:None
		"""
		request = {
			'cluster_id': cluster_id
		}
		response = self.connection.call("StopClusterOnYarn", request)
		return 
	
	def get_cluster(self, cluster_id):
		"""
		Get cluster details

		Parameters:
		cluster_id: No description available (int64)

		Returns:
		cluster: No description available (Cluster)
		"""
		request = {
			'cluster_id': cluster_id
		}
		response = self.connection.call("GetCluster", request)
		return response['cluster']
	
	def get_cluster_on_yarn(self, cluster_id):
		"""
		Get cluster details (Yarn only)

		Parameters:
		cluster_id: No description available (int64)

		Returns:
		cluster: No description available (YarnCluster)
		"""
		request = {
			'cluster_id': cluster_id
		}
		response = self.connection.call("GetClusterOnYarn", request)
		return response['cluster']
	
	def get_clusters(self, offset, limit):
		"""
		List clusters

		Parameters:
		offset: No description available (int64)
		limit: No description available (int64)

		Returns:
		clusters: No description available (Cluster)
		"""
		request = {
			'offset': offset
			'limit': limit
		}
		response = self.connection.call("GetClusters", request)
		return response['clusters']
	
	def get_cluster_status(self, cluster_id):
		"""
		Get cluster status

		Parameters:
		cluster_id: No description available (int64)

		Returns:
		cluster_status: No description available (ClusterStatus)
		"""
		request = {
			'cluster_id': cluster_id
		}
		response = self.connection.call("GetClusterStatus", request)
		return response['cluster_status']
	
	def delete_cluster(self, cluster_id):
		"""
		Delete a cluster

		Parameters:
		cluster_id: No description available (int64)

		Returns:None
		"""
		request = {
			'cluster_id': cluster_id
		}
		response = self.connection.call("DeleteCluster", request)
		return 
	
	def get_job(self, cluster_id, job_name):
		"""
		Get job details

		Parameters:
		cluster_id: No description available (int64)
		job_name: No description available (string)

		Returns:
		job: No description available (Job)
		"""
		request = {
			'cluster_id': cluster_id
			'job_name': job_name
		}
		response = self.connection.call("GetJob", request)
		return response['job']
	
	def get_jobs(self, cluster_id):
		"""
		List jobs

		Parameters:
		cluster_id: No description available (int64)

		Returns:
		jobs: No description available (Job)
		"""
		request = {
			'cluster_id': cluster_id
		}
		response = self.connection.call("GetJobs", request)
		return response['jobs']
	
	def create_project(self, name, description):
		"""
		Create a project

		Parameters:
		name: No description available (string)
		description: No description available (string)

		Returns:
		project_id: No description available (int64)
		"""
		request = {
			'name': name
			'description': description
		}
		response = self.connection.call("CreateProject", request)
		return response['project_id']
	
	def get_projects(self, offset, limit):
		"""
		List projects

		Parameters:
		offset: No description available (int64)
		limit: No description available (int64)

		Returns:
		projects: No description available (Project)
		"""
		request = {
			'offset': offset
			'limit': limit
		}
		response = self.connection.call("GetProjects", request)
		return response['projects']
	
	def get_project(self, project_id):
		"""
		Get project details

		Parameters:
		project_id: No description available (int64)

		Returns:
		project: No description available (Project)
		"""
		request = {
			'project_id': project_id
		}
		response = self.connection.call("GetProject", request)
		return response['project']
	
	def delete_project(self, project_id):
		"""
		Delete a project

		Parameters:
		project_id: No description available (int64)

		Returns:None
		"""
		request = {
			'project_id': project_id
		}
		response = self.connection.call("DeleteProject", request)
		return 
	
	def create_datasource(self, project_id, name, description, path):
		"""
		Create a datasource

		Parameters:
		project_id: No description available (int64)
		name: No description available (string)
		description: No description available (string)
		path: No description available (string)

		Returns:
		datasource_id: No description available (int64)
		"""
		request = {
			'project_id': project_id
			'name': name
			'description': description
			'path': path
		}
		response = self.connection.call("CreateDatasource", request)
		return response['datasource_id']
	
	def get_datasources(self, project_id, offset, limit):
		"""
		List datasources

		Parameters:
		project_id: No description available (int64)
		offset: No description available (int64)
		limit: No description available (int64)

		Returns:
		datasources: No description available (Datasource)
		"""
		request = {
			'project_id': project_id
			'offset': offset
			'limit': limit
		}
		response = self.connection.call("GetDatasources", request)
		return response['datasources']
	
	def get_datasource(self, datasource_id):
		"""
		Get datasource details

		Parameters:
		datasource_id: No description available (int64)

		Returns:
		datasource: No description available (Datasource)
		"""
		request = {
			'datasource_id': datasource_id
		}
		response = self.connection.call("GetDatasource", request)
		return response['datasource']
	
	def update_datasource(self, datasource_id, name, description, path):
		"""
		Update a datasource

		Parameters:
		datasource_id: No description available (int64)
		name: No description available (string)
		description: No description available (string)
		path: No description available (string)

		Returns:None
		"""
		request = {
			'datasource_id': datasource_id
			'name': name
			'description': description
			'path': path
		}
		response = self.connection.call("UpdateDatasource", request)
		return 
	
	def delete_datasource(self, datasource_id):
		"""
		Delete a datasource

		Parameters:
		datasource_id: No description available (int64)

		Returns:None
		"""
		request = {
			'datasource_id': datasource_id
		}
		response = self.connection.call("DeleteDatasource", request)
		return 
	
	def create_dataset(self, cluster_id, datasource_id, name, description, response_column_name):
		"""
		Create a dataset

		Parameters:
		cluster_id: No description available (int64)
		datasource_id: No description available (int64)
		name: No description available (string)
		description: No description available (string)
		response_column_name: No description available (string)

		Returns:
		dataset_id: No description available (int64)
		"""
		request = {
			'cluster_id': cluster_id
			'datasource_id': datasource_id
			'name': name
			'description': description
			'response_column_name': response_column_name
		}
		response = self.connection.call("CreateDataset", request)
		return response['dataset_id']
	
	def get_datasets(self, datasource_id, offset, limit):
		"""
		List datasets

		Parameters:
		datasource_id: No description available (int64)
		offset: No description available (int64)
		limit: No description available (int64)

		Returns:
		datasets: No description available (Dataset)
		"""
		request = {
			'datasource_id': datasource_id
			'offset': offset
			'limit': limit
		}
		response = self.connection.call("GetDatasets", request)
		return response['datasets']
	
	def get_dataset(self, dataset_id):
		"""
		Get dataset details

		Parameters:
		dataset_id: No description available (int64)

		Returns:
		dataset: No description available (Dataset)
		"""
		request = {
			'dataset_id': dataset_id
		}
		response = self.connection.call("GetDataset", request)
		return response['dataset']
	
	def update_dataset(self, dataset_id, name, description, response_column_name):
		"""
		Update a dataset

		Parameters:
		dataset_id: No description available (int64)
		name: No description available (string)
		description: No description available (string)
		response_column_name: No description available (string)

		Returns:None
		"""
		request = {
			'dataset_id': dataset_id
			'name': name
			'description': description
			'response_column_name': response_column_name
		}
		response = self.connection.call("UpdateDataset", request)
		return 
	
	def split_dataset(self, dataset_id, ratio1, ratio2):
		"""
		Split a dataset

		Parameters:
		dataset_id: No description available (int64)
		ratio1: No description available (int)
		ratio2: No description available (int)

		Returns:
		dataset_ids: No description available (int64)
		"""
		request = {
			'dataset_id': dataset_id
			'ratio1': ratio1
			'ratio2': ratio2
		}
		response = self.connection.call("SplitDataset", request)
		return response['dataset_ids']
	
	def delete_dataset(self, dataset_id):
		"""
		Delete a dataset

		Parameters:
		dataset_id: No description available (int64)

		Returns:None
		"""
		request = {
			'dataset_id': dataset_id
		}
		response = self.connection.call("DeleteDataset", request)
		return 
	
	def build_model(self, cluster_id, dataset_id, algorithm):
		"""
		Build a model

		Parameters:
		cluster_id: No description available (int64)
		dataset_id: No description available (int64)
		algorithm: No description available (string)

		Returns:
		model_id: No description available (int64)
		"""
		request = {
			'cluster_id': cluster_id
			'dataset_id': dataset_id
			'algorithm': algorithm
		}
		response = self.connection.call("BuildModel", request)
		return response['model_id']
	
	def build_model_auto(self, cluster_id, dataset, target_name, max_run_time):
		"""
		Build an AutoML model

		Parameters:
		cluster_id: No description available (int64)
		dataset: No description available (string)
		target_name: No description available (string)
		max_run_time: No description available (int)

		Returns:
		model: No description available (Model)
		"""
		request = {
			'cluster_id': cluster_id
			'dataset': dataset
			'target_name': target_name
			'max_run_time': max_run_time
		}
		response = self.connection.call("BuildModelAuto", request)
		return response['model']
	
	def get_model(self, model_id):
		"""
		Get model details

		Parameters:
		model_id: No description available (int64)

		Returns:
		model: No description available (Model)
		"""
		request = {
			'model_id': model_id
		}
		response = self.connection.call("GetModel", request)
		return response['model']
	
	def get_models(self, project_id, offset, limit):
		"""
		List models

		Parameters:
		project_id: No description available (int64)
		offset: No description available (int64)
		limit: No description available (int64)

		Returns:
		models: No description available (Model)
		"""
		request = {
			'project_id': project_id
			'offset': offset
			'limit': limit
		}
		response = self.connection.call("GetModels", request)
		return response['models']
	
	def get_models_from_cluster(self, cluster_id):
		"""
		List models from a cluster

		Parameters:
		cluster_id: No description available (int64)

		Returns:
		models: No description available (Model)
		"""
		request = {
			'cluster_id': cluster_id
		}
		response = self.connection.call("GetModelsFromCluster", request)
		return response['models']
	
	def import_model_from_cluster(self, cluster_id, project_id, model_key, model_name):
		"""
		Import models from a cluster

		Parameters:
		cluster_id: No description available (int64)
		project_id: No description available (int64)
		model_key: No description available (string)
		model_name: No description available (string)

		Returns:
		model_id: No description available (int64)
		"""
		request = {
			'cluster_id': cluster_id
			'project_id': project_id
			'model_key': model_key
			'model_name': model_name
		}
		response = self.connection.call("ImportModelFromCluster", request)
		return response['model_id']
	
	def rename_model(self, model_id, model_name):
		"""
		Update a model name in the database

		Parameters:
		model_id: No description available (int64)
		model_name: No description available (string)

		Returns:None
		"""
		request = {
			'model_id': model_id
			'model_name': model_name
		}
		response = self.connection.call("RenameModel", request)
		return 
	
	def delete_model(self, model_id):
		"""
		Delete a model

		Parameters:
		model_id: No description available (int64)

		Returns:None
		"""
		request = {
			'model_id': model_id
		}
		response = self.connection.call("DeleteModel", request)
		return 
	
	def start_service(self, model_id, name, port):
		"""
		Start a service

		Parameters:
		model_id: No description available (int64)
		name: No description available (string)
		port: No description available (int)

		Returns:
		service_id: No description available (int64)
		"""
		request = {
			'model_id': model_id
			'name': name
			'port': port
		}
		response = self.connection.call("StartService", request)
		return response['service_id']
	
	def stop_service(self, service_id):
		"""
		Stop a service

		Parameters:
		service_id: No description available (int64)

		Returns:None
		"""
		request = {
			'service_id': service_id
		}
		response = self.connection.call("StopService", request)
		return 
	
	def get_service(self, service_id):
		"""
		Get service details

		Parameters:
		service_id: No description available (int64)

		Returns:
		service: No description available (ScoringService)
		"""
		request = {
			'service_id': service_id
		}
		response = self.connection.call("GetService", request)
		return response['service']
	
	def get_services(self, offset, limit):
		"""
		List services

		Parameters:
		offset: No description available (int64)
		limit: No description available (int64)

		Returns:
		services: No description available (ScoringService)
		"""
		request = {
			'offset': offset
			'limit': limit
		}
		response = self.connection.call("GetServices", request)
		return response['services']
	
	def get_services_for_model(self, model_id, offset, limit):
		"""
		List services for a model

		Parameters:
		model_id: No description available (int64)
		offset: No description available (int64)
		limit: No description available (int64)

		Returns:
		services: No description available (ScoringService)
		"""
		request = {
			'model_id': model_id
			'offset': offset
			'limit': limit
		}
		response = self.connection.call("GetServicesForModel", request)
		return response['services']
	
	def rename_service(self, service_id, service_name):
		"""
		Update a service name in the database

		Parameters:
		service_id: No description available (int64)
		service_name: No description available (string)

		Returns:None
		"""
		request = {
			'service_id': service_id
			'service_name': service_name
		}
		response = self.connection.call("RenameService", request)
		return 
	
	def delete_service(self, service_id):
		"""
		Delete a service

		Parameters:
		service_id: No description available (int64)

		Returns:None
		"""
		request = {
			'service_id': service_id
		}
		response = self.connection.call("DeleteService", request)
		return 
	
	def add_engine(self, engine_name, engine_path):
		"""
		Add an engine

		Parameters:
		engine_name: No description available (string)
		engine_path: No description available (string)

		Returns:
		engine_id: No description available (int64)
		"""
		request = {
			'engine_name': engine_name
			'engine_path': engine_path
		}
		response = self.connection.call("AddEngine", request)
		return response['engine_id']
	
	def get_engine(self, engine_id):
		"""
		Get engine details

		Parameters:
		engine_id: No description available (int64)

		Returns:
		engine: No description available (Engine)
		"""
		request = {
			'engine_id': engine_id
		}
		response = self.connection.call("GetEngine", request)
		return response['engine']
	
	def get_engines(self):
		"""
		List engines

		Parameters:

		Returns:
		engines: No description available (Engine)
		"""
		request = {
		}
		response = self.connection.call("GetEngines", request)
		return response['engines']
	
	def delete_engine(self, engine_id):
		"""
		Delete an engine

		Parameters:
		engine_id: No description available (int64)

		Returns:None
		"""
		request = {
			'engine_id': engine_id
		}
		response = self.connection.call("DeleteEngine", request)
		return 
	
	def get_all_entity_types(self):
		"""
		List all entity types

		Parameters:

		Returns:
		entity_types: No description available (EntityType)
		"""
		request = {
		}
		response = self.connection.call("GetAllEntityTypes", request)
		return response['entity_types']
	
	def get_all_permissions(self):
		"""
		List all permissions

		Parameters:

		Returns:
		permissions: No description available (Permission)
		"""
		request = {
		}
		response = self.connection.call("GetAllPermissions", request)
		return response['permissions']
	
	def get_all_cluster_types(self):
		"""
		List all cluster types

		Parameters:

		Returns:
		cluster_types: No description available (ClusterType)
		"""
		request = {
		}
		response = self.connection.call("GetAllClusterTypes", request)
		return response['cluster_types']
	
	def get_permissions_for_role(self, role_id):
		"""
		List permissions for a role

		Parameters:
		role_id: No description available (int64)

		Returns:
		permissions: No description available (Permission)
		"""
		request = {
			'role_id': role_id
		}
		response = self.connection.call("GetPermissionsForRole", request)
		return response['permissions']
	
	def get_permissions_for_identity(self, identity_id):
		"""
		List permissions for an identity

		Parameters:
		identity_id: No description available (int64)

		Returns:
		permissions: No description available (Permission)
		"""
		request = {
			'identity_id': identity_id
		}
		response = self.connection.call("GetPermissionsForIdentity", request)
		return response['permissions']
	
	def create_role(self, name, description):
		"""
		Create a role

		Parameters:
		name: No description available (string)
		description: No description available (string)

		Returns:
		role_id: No description available (int64)
		"""
		request = {
			'name': name
			'description': description
		}
		response = self.connection.call("CreateRole", request)
		return response['role_id']
	
	def get_roles(self, offset, limit):
		"""
		List roles

		Parameters:
		offset: No description available (int64)
		limit: No description available (int64)

		Returns:
		roles: No description available (Role)
		"""
		request = {
			'offset': offset
			'limit': limit
		}
		response = self.connection.call("GetRoles", request)
		return response['roles']
	
	def get_roles_for_identity(self, identity_id):
		"""
		List roles for an identity

		Parameters:
		identity_id: No description available (int64)

		Returns:
		roles: No description available (Role)
		"""
		request = {
			'identity_id': identity_id
		}
		response = self.connection.call("GetRolesForIdentity", request)
		return response['roles']
	
	def get_role(self, role_id):
		"""
		Get role details

		Parameters:
		role_id: No description available (int64)

		Returns:
		role: No description available (Role)
		"""
		request = {
			'role_id': role_id
		}
		response = self.connection.call("GetRole", request)
		return response['role']
	
	def get_role_by_name(self, name):
		"""
		Get role details by name

		Parameters:
		name: No description available (string)

		Returns:
		role: No description available (Role)
		"""
		request = {
			'name': name
		}
		response = self.connection.call("GetRoleByName", request)
		return response['role']
	
	def update_role(self, role_id, name, description):
		"""
		Update a role

		Parameters:
		role_id: No description available (int64)
		name: No description available (string)
		description: No description available (string)

		Returns:None
		"""
		request = {
			'role_id': role_id
			'name': name
			'description': description
		}
		response = self.connection.call("UpdateRole", request)
		return 
	
	def link_role_with_permissions(self, role_id, permission_ids):
		"""
		Link a role with permissions

		Parameters:
		role_id: No description available (int64)
		permission_ids: No description available (int64)

		Returns:None
		"""
		request = {
			'role_id': role_id
			'permission_ids': permission_ids
		}
		response = self.connection.call("LinkRoleWithPermissions", request)
		return 
	
	def link_role_with_permission(self, role_id, permission_id):
		"""
		Link a role with a permission

		Parameters:
		role_id: No description available (int64)
		permission_id: No description available (int64)

		Returns:None
		"""
		request = {
			'role_id': role_id
			'permission_id': permission_id
		}
		response = self.connection.call("LinkRoleWithPermission", request)
		return 
	
	def unlink_role_from_permission(self, role_id, permission_id):
		"""
		Unlink a role from a permission

		Parameters:
		role_id: No description available (int64)
		permission_id: No description available (int64)

		Returns:None
		"""
		request = {
			'role_id': role_id
			'permission_id': permission_id
		}
		response = self.connection.call("UnlinkRoleFromPermission", request)
		return 
	
	def delete_role(self, role_id):
		"""
		Delete a role

		Parameters:
		role_id: No description available (int64)

		Returns:None
		"""
		request = {
			'role_id': role_id
		}
		response = self.connection.call("DeleteRole", request)
		return 
	
	def create_workgroup(self, name, description):
		"""
		Create a workgroup

		Parameters:
		name: No description available (string)
		description: No description available (string)

		Returns:
		workgroup_id: No description available (int64)
		"""
		request = {
			'name': name
			'description': description
		}
		response = self.connection.call("CreateWorkgroup", request)
		return response['workgroup_id']
	
	def get_workgroups(self, offset, limit):
		"""
		List workgroups

		Parameters:
		offset: No description available (int64)
		limit: No description available (int64)

		Returns:
		workgroups: No description available (Workgroup)
		"""
		request = {
			'offset': offset
			'limit': limit
		}
		response = self.connection.call("GetWorkgroups", request)
		return response['workgroups']
	
	def get_workgroups_for_identity(self, identity_id):
		"""
		List workgroups for an identity

		Parameters:
		identity_id: No description available (int64)

		Returns:
		workgroups: No description available (Workgroup)
		"""
		request = {
			'identity_id': identity_id
		}
		response = self.connection.call("GetWorkgroupsForIdentity", request)
		return response['workgroups']
	
	def get_workgroup(self, workgroup_id):
		"""
		Get workgroup details

		Parameters:
		workgroup_id: No description available (int64)

		Returns:
		workgroup: No description available (Workgroup)
		"""
		request = {
			'workgroup_id': workgroup_id
		}
		response = self.connection.call("GetWorkgroup", request)
		return response['workgroup']
	
	def get_workgroup_by_name(self, name):
		"""
		Get workgroup details by name

		Parameters:
		name: No description available (string)

		Returns:
		workgroup: No description available (Workgroup)
		"""
		request = {
			'name': name
		}
		response = self.connection.call("GetWorkgroupByName", request)
		return response['workgroup']
	
	def update_workgroup(self, workgroup_id, name, description):
		"""
		Update a workgroup

		Parameters:
		workgroup_id: No description available (int64)
		name: No description available (string)
		description: No description available (string)

		Returns:None
		"""
		request = {
			'workgroup_id': workgroup_id
			'name': name
			'description': description
		}
		response = self.connection.call("UpdateWorkgroup", request)
		return 
	
	def delete_workgroup(self, workgroup_id):
		"""
		Delete a workgroup

		Parameters:
		workgroup_id: No description available (int64)

		Returns:None
		"""
		request = {
			'workgroup_id': workgroup_id
		}
		response = self.connection.call("DeleteWorkgroup", request)
		return 
	
	def create_identity(self, name, password):
		"""
		Create an identity

		Parameters:
		name: No description available (string)
		password: No description available (string)

		Returns:
		identity_id: No description available (int64)
		"""
		request = {
			'name': name
			'password': password
		}
		response = self.connection.call("CreateIdentity", request)
		return response['identity_id']
	
	def get_identities(self, offset, limit):
		"""
		List identities

		Parameters:
		offset: No description available (int64)
		limit: No description available (int64)

		Returns:
		identities: No description available (Identity)
		"""
		request = {
			'offset': offset
			'limit': limit
		}
		response = self.connection.call("GetIdentities", request)
		return response['identities']
	
	def get_identities_for_workgroup(self, workgroup_id):
		"""
		List identities for a workgroup

		Parameters:
		workgroup_id: No description available (int64)

		Returns:
		identities: No description available (Identity)
		"""
		request = {
			'workgroup_id': workgroup_id
		}
		response = self.connection.call("GetIdentitiesForWorkgroup", request)
		return response['identities']
	
	def get_identities_for_role(self, role_id):
		"""
		List identities for a role

		Parameters:
		role_id: No description available (int64)

		Returns:
		identities: No description available (Identity)
		"""
		request = {
			'role_id': role_id
		}
		response = self.connection.call("GetIdentitiesForRole", request)
		return response['identities']
	
	def get_identity(self, identity_id):
		"""
		Get identity details

		Parameters:
		identity_id: No description available (int64)

		Returns:
		identity: No description available (Identity)
		"""
		request = {
			'identity_id': identity_id
		}
		response = self.connection.call("GetIdentity", request)
		return response['identity']
	
	def get_identity_by_name(self, name):
		"""
		Get identity details by name

		Parameters:
		name: No description available (string)

		Returns:
		identity: No description available (Identity)
		"""
		request = {
			'name': name
		}
		response = self.connection.call("GetIdentityByName", request)
		return response['identity']
	
	def link_identity_with_workgroup(self, identity_id, workgroup_id):
		"""
		Link an identity with a workgroup

		Parameters:
		identity_id: No description available (int64)
		workgroup_id: No description available (int64)

		Returns:None
		"""
		request = {
			'identity_id': identity_id
			'workgroup_id': workgroup_id
		}
		response = self.connection.call("LinkIdentityWithWorkgroup", request)
		return 
	
	def unlink_identity_from_workgroup(self, identity_id, workgroup_id):
		"""
		Unlink an identity from a workgroup

		Parameters:
		identity_id: No description available (int64)
		workgroup_id: No description available (int64)

		Returns:None
		"""
		request = {
			'identity_id': identity_id
			'workgroup_id': workgroup_id
		}
		response = self.connection.call("UnlinkIdentityFromWorkgroup", request)
		return 
	
	def link_identity_with_role(self, identity_id, role_id):
		"""
		Link an identity with a role

		Parameters:
		identity_id: No description available (int64)
		role_id: No description available (int64)

		Returns:None
		"""
		request = {
			'identity_id': identity_id
			'role_id': role_id
		}
		response = self.connection.call("LinkIdentityWithRole", request)
		return 
	
	def unlink_identity_from_role(self, identity_id, role_id):
		"""
		Unlink an identity from a role

		Parameters:
		identity_id: No description available (int64)
		role_id: No description available (int64)

		Returns:None
		"""
		request = {
			'identity_id': identity_id
			'role_id': role_id
		}
		response = self.connection.call("UnlinkIdentityFromRole", request)
		return 
	
	def update_identity(self, identity_id, password):
		"""
		Update an identity

		Parameters:
		identity_id: No description available (int64)
		password: No description available (string)

		Returns:None
		"""
		request = {
			'identity_id': identity_id
			'password': password
		}
		response = self.connection.call("UpdateIdentity", request)
		return 
	
	def deactivate_identity(self, identity_id):
		"""
		Deactivate an identity

		Parameters:
		identity_id: No description available (int64)

		Returns:None
		"""
		request = {
			'identity_id': identity_id
		}
		response = self.connection.call("DeactivateIdentity", request)
		return 
	
	def share_entity(self, kind, workgroup_id, entity_type_id, entity_id):
		"""
		Share an entity with a workgroup

		Parameters:
		kind: No description available (string)
		workgroup_id: No description available (int64)
		entity_type_id: No description available (int64)
		entity_id: No description available (int64)

		Returns:None
		"""
		request = {
			'kind': kind
			'workgroup_id': workgroup_id
			'entity_type_id': entity_type_id
			'entity_id': entity_id
		}
		response = self.connection.call("ShareEntity", request)
		return 
	
	def get_privileges(self, entity_type_id, entity_id):
		"""
		List privileges for an entity

		Parameters:
		entity_type_id: No description available (int64)
		entity_id: No description available (int64)

		Returns:
		privileges: No description available (EntityPrivilege)
		"""
		request = {
			'entity_type_id': entity_type_id
			'entity_id': entity_id
		}
		response = self.connection.call("GetPrivileges", request)
		return response['privileges']
	
	def unshare_entity(self, kind, workgroup_id, entity_type_id, entity_id):
		"""
		Unshare an entity

		Parameters:
		kind: No description available (string)
		workgroup_id: No description available (int64)
		entity_type_id: No description available (int64)
		entity_id: No description available (int64)

		Returns:None
		"""
		request = {
			'kind': kind
			'workgroup_id': workgroup_id
			'entity_type_id': entity_type_id
			'entity_id': entity_id
		}
		response = self.connection.call("UnshareEntity", request)
		return 
	
	def get_history(self, entity_type_id, entity_id, offset, limit):
		"""
		List audit trail records for an entity

		Parameters:
		entity_type_id: No description available (int64)
		entity_id: No description available (int64)
		offset: No description available (int64)
		limit: No description available (int64)

		Returns:
		history: No description available (EntityHistory)
		"""
		request = {
			'entity_type_id': entity_type_id
			'entity_id': entity_id
			'offset': offset
			'limit': limit
		}
		response = self.connection.call("GetHistory", request)
		return response['history']
	
	

