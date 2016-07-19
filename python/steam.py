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
	
	
	
	def ping(self, input):
		"""
		Parameters:
		input: bool
		
		Returns:
		output: bool
    
		"""
		request = {
			'input': input
			
		}
		response = self.connection.call("Ping", request)
		return response['output']
	
	def register_cluster(self, address):
		"""
		Parameters:
		address: string
		
		Returns:
		cluster_id: int64
    
		"""
		request = {
			'address': address
			
		}
		response = self.connection.call("RegisterCluster", request)
		return response['cluster_id']
	
	def unregister_cluster(self, cluster_id):
		"""
		Parameters:
		cluster_id: int64
		
		Returns:
		None
		"""
		request = {
			'cluster_id': cluster_id
			
		}
		response = self.connection.call("UnregisterCluster", request)
		return 
	
	def start_yarn_cluster(self, cluster_name, engine_id, size, memory, username):
		"""
		Parameters:
		cluster_name: string
		engine_id: int64
		size: int
		memory: string
		username: string
		
		Returns:
		cluster_id: int64
    
		"""
		request = {
			'cluster_name': cluster_name
			'engine_id': engine_id
			'size': size
			'memory': memory
			'username': username
			
		}
		response = self.connection.call("StartYarnCluster", request)
		return response['cluster_id']
	
	def stop_yarn_cluster(self, cluster_id):
		"""
		Parameters:
		cluster_id: int64
		
		Returns:
		None
		"""
		request = {
			'cluster_id': cluster_id
			
		}
		response = self.connection.call("StopYarnCluster", request)
		return 
	
	def get_cluster(self, cluster_id):
		"""
		Parameters:
		cluster_id: int64
		
		Returns:
		cluster: Cluster
    
		"""
		request = {
			'cluster_id': cluster_id
			
		}
		response = self.connection.call("GetCluster", request)
		return response['cluster']
	
	def get_yarn_cluster(self, cluster_id):
		"""
		Parameters:
		cluster_id: int64
		
		Returns:
		cluster: YarnCluster
    
		"""
		request = {
			'cluster_id': cluster_id
			
		}
		response = self.connection.call("GetYarnCluster", request)
		return response['cluster']
	
	def get_clusters(self, offset, limit):
		"""
		Parameters:
		offset: int64
		limit: int64
		
		Returns:
		clusters: Cluster
    
		"""
		request = {
			'offset': offset
			'limit': limit
			
		}
		response = self.connection.call("GetClusters", request)
		return response['clusters']
	
	def get_cluster_status(self, cluster_id):
		"""
		Parameters:
		cluster_id: int64
		
		Returns:
		cluster_status: ClusterStatus
    
		"""
		request = {
			'cluster_id': cluster_id
			
		}
		response = self.connection.call("GetClusterStatus", request)
		return response['cluster_status']
	
	def delete_cluster(self, cluster_id):
		"""
		Parameters:
		cluster_id: int64
		
		Returns:
		None
		"""
		request = {
			'cluster_id': cluster_id
			
		}
		response = self.connection.call("DeleteCluster", request)
		return 
	
	def get_job(self, cluster_id, job_name):
		"""
		Parameters:
		cluster_id: int64
		job_name: string
		
		Returns:
		job: Job
    
		"""
		request = {
			'cluster_id': cluster_id
			'job_name': job_name
			
		}
		response = self.connection.call("GetJob", request)
		return response['job']
	
	def get_jobs(self, cluster_id):
		"""
		Parameters:
		cluster_id: int64
		
		Returns:
		jobs: Job
    
		"""
		request = {
			'cluster_id': cluster_id
			
		}
		response = self.connection.call("GetJobs", request)
		return response['jobs']
	
	def create_project(self, name, description):
		"""
		Parameters:
		name: string
		description: string
		
		Returns:
		project_id: int64
    
		"""
		request = {
			'name': name
			'description': description
			
		}
		response = self.connection.call("CreateProject", request)
		return response['project_id']
	
	def get_projects(self, offset, limit):
		"""
		Parameters:
		offset: int64
		limit: int64
		
		Returns:
		projects: Project
    
		"""
		request = {
			'offset': offset
			'limit': limit
			
		}
		response = self.connection.call("GetProjects", request)
		return response['projects']
	
	def get_project(self, project_id):
		"""
		Parameters:
		project_id: int64
		
		Returns:
		project: Project
    
		"""
		request = {
			'project_id': project_id
			
		}
		response = self.connection.call("GetProject", request)
		return response['project']
	
	def delete_project(self, project_id):
		"""
		Parameters:
		project_id: int64
		
		Returns:
		None
		"""
		request = {
			'project_id': project_id
			
		}
		response = self.connection.call("DeleteProject", request)
		return 
	
	def create_datasource(self, project_id, name, description, path):
		"""
		Parameters:
		project_id: int64
		name: string
		description: string
		path: string
		
		Returns:
		datasource_id: int64
    
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
		Parameters:
		project_id: int64
		offset: int64
		limit: int64
		
		Returns:
		datasources: Datasource
    
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
		Parameters:
		datasource_id: int64
		
		Returns:
		datasource: Datasource
    
		"""
		request = {
			'datasource_id': datasource_id
			
		}
		response = self.connection.call("GetDatasource", request)
		return response['datasource']
	
	def update_datasource(self, datasource_id, name, description, path):
		"""
		Parameters:
		datasource_id: int64
		name: string
		description: string
		path: string
		
		Returns:
		None
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
		Parameters:
		datasource_id: int64
		
		Returns:
		None
		"""
		request = {
			'datasource_id': datasource_id
			
		}
		response = self.connection.call("DeleteDatasource", request)
		return 
	
	def create_dataset(self, cluster_id, datasource_id, name, description, response_column_name):
		"""
		Parameters:
		cluster_id: int64
		datasource_id: int64
		name: string
		description: string
		response_column_name: string
		
		Returns:
		dataset_id: int64
    
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
		Parameters:
		datasource_id: int64
		offset: int64
		limit: int64
		
		Returns:
		datasets: Dataset
    
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
		Parameters:
		dataset_id: int64
		
		Returns:
		dataset: Dataset
    
		"""
		request = {
			'dataset_id': dataset_id
			
		}
		response = self.connection.call("GetDataset", request)
		return response['dataset']
	
	def update_dataset(self, dataset_id, name, description, response_column_name):
		"""
		Parameters:
		dataset_id: int64
		name: string
		description: string
		response_column_name: string
		
		Returns:
		None
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
		Parameters:
		dataset_id: int64
		ratio1: int
		ratio2: int
		
		Returns:
		dataset_ids: int64
    
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
		Parameters:
		dataset_id: int64
		
		Returns:
		None
		"""
		request = {
			'dataset_id': dataset_id
			
		}
		response = self.connection.call("DeleteDataset", request)
		return 
	
	def build_model(self, cluster_id, dataset_id, algorithm):
		"""
		Parameters:
		cluster_id: int64
		dataset_id: int64
		algorithm: string
		
		Returns:
		model_id: int64
    
		"""
		request = {
			'cluster_id': cluster_id
			'dataset_id': dataset_id
			'algorithm': algorithm
			
		}
		response = self.connection.call("BuildModel", request)
		return response['model_id']
	
	def build_auto_model(self, cluster_id, dataset, target_name, max_run_time):
		"""
		Parameters:
		cluster_id: int64
		dataset: string
		target_name: string
		max_run_time: int
		
		Returns:
		model: Model
    
		"""
		request = {
			'cluster_id': cluster_id
			'dataset': dataset
			'target_name': target_name
			'max_run_time': max_run_time
			
		}
		response = self.connection.call("BuildAutoModel", request)
		return response['model']
	
	def get_model(self, model_id):
		"""
		Parameters:
		model_id: int64
		
		Returns:
		model: Model
    
		"""
		request = {
			'model_id': model_id
			
		}
		response = self.connection.call("GetModel", request)
		return response['model']
	
	def get_models(self, project_id, offset, limit):
		"""
		Parameters:
		project_id: int64
		offset: int64
		limit: int64
		
		Returns:
		models: Model
    
		"""
		request = {
			'project_id': project_id
			'offset': offset
			'limit': limit
			
		}
		response = self.connection.call("GetModels", request)
		return response['models']
	
	def get_cluster_models(self, cluster_id):
		"""
		Parameters:
		cluster_id: int64
		
		Returns:
		models: Model
    
		"""
		request = {
			'cluster_id': cluster_id
			
		}
		response = self.connection.call("GetClusterModels", request)
		return response['models']
	
	def import_model_from_cluster(self, cluster_id, project_id, model_name):
		"""
		Parameters:
		cluster_id: int64
		project_id: int64
		model_name: string
		
		Returns:
		model: Model
    
		"""
		request = {
			'cluster_id': cluster_id
			'project_id': project_id
			'model_name': model_name
			
		}
		response = self.connection.call("ImportModelFromCluster", request)
		return response['model']
	
	def delete_model(self, model_id):
		"""
		Parameters:
		model_id: int64
		
		Returns:
		None
		"""
		request = {
			'model_id': model_id
			
		}
		response = self.connection.call("DeleteModel", request)
		return 
	
	def start_scoring_service(self, model_id, port):
		"""
		Parameters:
		model_id: int64
		port: int
		
		Returns:
		service: ScoringService
    
		"""
		request = {
			'model_id': model_id
			'port': port
			
		}
		response = self.connection.call("StartScoringService", request)
		return response['service']
	
	def stop_scoring_service(self, service_id):
		"""
		Parameters:
		service_id: int64
		
		Returns:
		None
		"""
		request = {
			'service_id': service_id
			
		}
		response = self.connection.call("StopScoringService", request)
		return 
	
	def get_scoring_service(self, service_id):
		"""
		Parameters:
		service_id: int64
		
		Returns:
		service: ScoringService
    
		"""
		request = {
			'service_id': service_id
			
		}
		response = self.connection.call("GetScoringService", request)
		return response['service']
	
	def get_scoring_services(self, offset, limit):
		"""
		Parameters:
		offset: int64
		limit: int64
		
		Returns:
		services: ScoringService
    
		"""
		request = {
			'offset': offset
			'limit': limit
			
		}
		response = self.connection.call("GetScoringServices", request)
		return response['services']
	
	def get_scoring_services_for_model(self, model_id, offset, limit):
		"""
		Parameters:
		model_id: int64
		offset: int64
		limit: int64
		
		Returns:
		services: ScoringService
    
		"""
		request = {
			'model_id': model_id
			'offset': offset
			'limit': limit
			
		}
		response = self.connection.call("GetScoringServicesForModel", request)
		return response['services']
	
	def delete_scoring_service(self, service_id):
		"""
		Parameters:
		service_id: int64
		
		Returns:
		None
		"""
		request = {
			'service_id': service_id
			
		}
		response = self.connection.call("DeleteScoringService", request)
		return 
	
	def add_engine(self, engine_name, engine_path):
		"""
		Parameters:
		engine_name: string
		engine_path: string
		
		Returns:
		engine_id: int64
    
		"""
		request = {
			'engine_name': engine_name
			'engine_path': engine_path
			
		}
		response = self.connection.call("AddEngine", request)
		return response['engine_id']
	
	def get_engine(self, engine_id):
		"""
		Parameters:
		engine_id: int64
		
		Returns:
		engine: Engine
    
		"""
		request = {
			'engine_id': engine_id
			
		}
		response = self.connection.call("GetEngine", request)
		return response['engine']
	
	def get_engines(self):
		"""
		Parameters:
		
		Returns:
		engines: Engine
    
		"""
		request = {
			
		}
		response = self.connection.call("GetEngines", request)
		return response['engines']
	
	def delete_engine(self, engine_id):
		"""
		Parameters:
		engine_id: int64
		
		Returns:
		None
		"""
		request = {
			'engine_id': engine_id
			
		}
		response = self.connection.call("DeleteEngine", request)
		return 
	
	def get_supported_entity_types(self):
		"""
		Parameters:
		
		Returns:
		entity_types: EntityType
    
		"""
		request = {
			
		}
		response = self.connection.call("GetSupportedEntityTypes", request)
		return response['entity_types']
	
	def get_supported_permissions(self):
		"""
		Parameters:
		
		Returns:
		permissions: Permission
    
		"""
		request = {
			
		}
		response = self.connection.call("GetSupportedPermissions", request)
		return response['permissions']
	
	def get_supported_cluster_types(self):
		"""
		Parameters:
		
		Returns:
		cluster_types: ClusterType
    
		"""
		request = {
			
		}
		response = self.connection.call("GetSupportedClusterTypes", request)
		return response['cluster_types']
	
	def get_permissions_for_role(self, role_id):
		"""
		Parameters:
		role_id: int64
		
		Returns:
		permissions: Permission
    
		"""
		request = {
			'role_id': role_id
			
		}
		response = self.connection.call("GetPermissionsForRole", request)
		return response['permissions']
	
	def get_permissions_for_identity(self, identity_id):
		"""
		Parameters:
		identity_id: int64
		
		Returns:
		permissions: Permission
    
		"""
		request = {
			'identity_id': identity_id
			
		}
		response = self.connection.call("GetPermissionsForIdentity", request)
		return response['permissions']
	
	def create_role(self, name, description):
		"""
		Parameters:
		name: string
		description: string
		
		Returns:
		role_id: int64
    
		"""
		request = {
			'name': name
			'description': description
			
		}
		response = self.connection.call("CreateRole", request)
		return response['role_id']
	
	def get_roles(self, offset, limit):
		"""
		Parameters:
		offset: int64
		limit: int64
		
		Returns:
		roles: Role
    
		"""
		request = {
			'offset': offset
			'limit': limit
			
		}
		response = self.connection.call("GetRoles", request)
		return response['roles']
	
	def get_roles_for_identity(self, identity_id):
		"""
		Parameters:
		identity_id: int64
		
		Returns:
		roles: Role
    
		"""
		request = {
			'identity_id': identity_id
			
		}
		response = self.connection.call("GetRolesForIdentity", request)
		return response['roles']
	
	def get_role(self, role_id):
		"""
		Parameters:
		role_id: int64
		
		Returns:
		role: Role
    
		"""
		request = {
			'role_id': role_id
			
		}
		response = self.connection.call("GetRole", request)
		return response['role']
	
	def get_role_by_name(self, name):
		"""
		Parameters:
		name: string
		
		Returns:
		role: Role
    
		"""
		request = {
			'name': name
			
		}
		response = self.connection.call("GetRoleByName", request)
		return response['role']
	
	def update_role(self, role_id, name, description):
		"""
		Parameters:
		role_id: int64
		name: string
		description: string
		
		Returns:
		None
		"""
		request = {
			'role_id': role_id
			'name': name
			'description': description
			
		}
		response = self.connection.call("UpdateRole", request)
		return 
	
	def link_role_and_permissions(self, role_id, permission_ids):
		"""
		Parameters:
		role_id: int64
		permission_ids: int64
		
		Returns:
		None
		"""
		request = {
			'role_id': role_id
			'permission_ids': permission_ids
			
		}
		response = self.connection.call("LinkRoleAndPermissions", request)
		return 
	
	def delete_role(self, role_id):
		"""
		Parameters:
		role_id: int64
		
		Returns:
		None
		"""
		request = {
			'role_id': role_id
			
		}
		response = self.connection.call("DeleteRole", request)
		return 
	
	def create_workgroup(self, name, description):
		"""
		Parameters:
		name: string
		description: string
		
		Returns:
		workgroup_id: int64
    
		"""
		request = {
			'name': name
			'description': description
			
		}
		response = self.connection.call("CreateWorkgroup", request)
		return response['workgroup_id']
	
	def get_workgroups(self, offset, limit):
		"""
		Parameters:
		offset: int64
		limit: int64
		
		Returns:
		workgroups: Workgroup
    
		"""
		request = {
			'offset': offset
			'limit': limit
			
		}
		response = self.connection.call("GetWorkgroups", request)
		return response['workgroups']
	
	def get_workgroups_for_identity(self, identity_id):
		"""
		Parameters:
		identity_id: int64
		
		Returns:
		workgroups: Workgroup
    
		"""
		request = {
			'identity_id': identity_id
			
		}
		response = self.connection.call("GetWorkgroupsForIdentity", request)
		return response['workgroups']
	
	def get_workgroup(self, workgroup_id):
		"""
		Parameters:
		workgroup_id: int64
		
		Returns:
		workgroup: Workgroup
    
		"""
		request = {
			'workgroup_id': workgroup_id
			
		}
		response = self.connection.call("GetWorkgroup", request)
		return response['workgroup']
	
	def get_workgroup_by_name(self, name):
		"""
		Parameters:
		name: string
		
		Returns:
		workgroup: Workgroup
    
		"""
		request = {
			'name': name
			
		}
		response = self.connection.call("GetWorkgroupByName", request)
		return response['workgroup']
	
	def update_workgroup(self, workgroup_id, name, description):
		"""
		Parameters:
		workgroup_id: int64
		name: string
		description: string
		
		Returns:
		None
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
		Parameters:
		workgroup_id: int64
		
		Returns:
		None
		"""
		request = {
			'workgroup_id': workgroup_id
			
		}
		response = self.connection.call("DeleteWorkgroup", request)
		return 
	
	def create_identity(self, name, password):
		"""
		Parameters:
		name: string
		password: string
		
		Returns:
		identity_id: int64
    
		"""
		request = {
			'name': name
			'password': password
			
		}
		response = self.connection.call("CreateIdentity", request)
		return response['identity_id']
	
	def get_identities(self, offset, limit):
		"""
		Parameters:
		offset: int64
		limit: int64
		
		Returns:
		identities: Identity
    
		"""
		request = {
			'offset': offset
			'limit': limit
			
		}
		response = self.connection.call("GetIdentities", request)
		return response['identities']
	
	def get_identities_for_workgroup(self, workgroup_id):
		"""
		Parameters:
		workgroup_id: int64
		
		Returns:
		identities: Identity
    
		"""
		request = {
			'workgroup_id': workgroup_id
			
		}
		response = self.connection.call("GetIdentitiesForWorkgroup", request)
		return response['identities']
	
	def get_identities_for_role(self, role_id):
		"""
		Parameters:
		role_id: int64
		
		Returns:
		identities: Identity
    
		"""
		request = {
			'role_id': role_id
			
		}
		response = self.connection.call("GetIdentitiesForRole", request)
		return response['identities']
	
	def get_identity(self, identity_id):
		"""
		Parameters:
		identity_id: int64
		
		Returns:
		identity: Identity
    
		"""
		request = {
			'identity_id': identity_id
			
		}
		response = self.connection.call("GetIdentity", request)
		return response['identity']
	
	def get_identity_by_name(self, name):
		"""
		Parameters:
		name: string
		
		Returns:
		identity: Identity
    
		"""
		request = {
			'name': name
			
		}
		response = self.connection.call("GetIdentityByName", request)
		return response['identity']
	
	def link_identity_and_workgroup(self, identity_id, workgroup_id):
		"""
		Parameters:
		identity_id: int64
		workgroup_id: int64
		
		Returns:
		None
		"""
		request = {
			'identity_id': identity_id
			'workgroup_id': workgroup_id
			
		}
		response = self.connection.call("LinkIdentityAndWorkgroup", request)
		return 
	
	def unlink_identity_and_workgroup(self, identity_id, workgroup_id):
		"""
		Parameters:
		identity_id: int64
		workgroup_id: int64
		
		Returns:
		None
		"""
		request = {
			'identity_id': identity_id
			'workgroup_id': workgroup_id
			
		}
		response = self.connection.call("UnlinkIdentityAndWorkgroup", request)
		return 
	
	def link_identity_and_role(self, identity_id, role_id):
		"""
		Parameters:
		identity_id: int64
		role_id: int64
		
		Returns:
		None
		"""
		request = {
			'identity_id': identity_id
			'role_id': role_id
			
		}
		response = self.connection.call("LinkIdentityAndRole", request)
		return 
	
	def unlink_identity_and_role(self, identity_id, role_id):
		"""
		Parameters:
		identity_id: int64
		role_id: int64
		
		Returns:
		None
		"""
		request = {
			'identity_id': identity_id
			'role_id': role_id
			
		}
		response = self.connection.call("UnlinkIdentityAndRole", request)
		return 
	
	def update_identity(self, identity_id, password):
		"""
		Parameters:
		identity_id: int64
		password: string
		
		Returns:
		None
		"""
		request = {
			'identity_id': identity_id
			'password': password
			
		}
		response = self.connection.call("UpdateIdentity", request)
		return 
	
	def deactivate_identity(self, identity_id):
		"""
		Parameters:
		identity_id: int64
		
		Returns:
		None
		"""
		request = {
			'identity_id': identity_id
			
		}
		response = self.connection.call("DeactivateIdentity", request)
		return 
	
	def share_entity(self, kind, workgroup_id, entity_type_id, entity_id):
		"""
		Parameters:
		kind: string
		workgroup_id: int64
		entity_type_id: int64
		entity_id: int64
		
		Returns:
		None
		"""
		request = {
			'kind': kind
			'workgroup_id': workgroup_id
			'entity_type_id': entity_type_id
			'entity_id': entity_id
			
		}
		response = self.connection.call("ShareEntity", request)
		return 
	
	def get_entity_privileges(self, entity_type_id, entity_id):
		"""
		Parameters:
		entity_type_id: int64
		entity_id: int64
		
		Returns:
		privileges: EntityPrivilege
    
		"""
		request = {
			'entity_type_id': entity_type_id
			'entity_id': entity_id
			
		}
		response = self.connection.call("GetEntityPrivileges", request)
		return response['privileges']
	
	def unshare_entity(self, kind, workgroup_id, entity_type_id, entity_id):
		"""
		Parameters:
		kind: string
		workgroup_id: int64
		entity_type_id: int64
		entity_id: int64
		
		Returns:
		None
		"""
		request = {
			'kind': kind
			'workgroup_id': workgroup_id
			'entity_type_id': entity_type_id
			'entity_id': entity_id
			
		}
		response = self.connection.call("UnshareEntity", request)
		return 
	
	def get_entity_history(self, entity_type_id, entity_id, offset, limit):
		"""
		Parameters:
		entity_type_id: int64
		entity_id: int64
		offset: int64
		limit: int64
		
		Returns:
		history: EntityHistory
    
		"""
		request = {
			'entity_type_id': entity_type_id
			'entity_id': entity_id
			'offset': offset
			'limit': limit
			
		}
		response = self.connection.call("GetEntityHistory", request)
		return response['history']
	
	

