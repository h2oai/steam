
# ----------------------------------
# --- Generated with go:generate ---
# ---        DO NOT EDIT         ---
# ----------------------------------

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
	

	def ping(self, status):
		"""
		Returns status (bool)
		"""
		request = {
			'status': status
		}
		response = self.connection.call("Ping", request)
		return response['status']

	def register_cluster(self, address):
		"""
		Returns clusterId (int64)
		"""
		request = {
			'address': address
		}
		response = self.connection.call("RegisterCluster", request)
		return response['cluster_id']

	def unregister_cluster(self, cluster_id):
		"""
		Returns None
		"""
		request = {
			'cluster_id': cluster_id
		}
		self.connection.call("UnregisterCluster", request)


	def start_yarn_cluster(self, cluster_name, engine_id, size, memory, username):
		"""
		Returns clusterId (int64)
		"""
		request = {
			'cluster_name': cluster_name,
			'engine_id': engine_id,
			'size': size,
			'memory': memory,
			'username': username
		}
		response = self.connection.call("StartYarnCluster", request)
		return response['cluster_id']

	def stop_yarn_cluster(self, cluster_id):
		"""
		Returns None
		"""
		request = {
			'cluster_id': cluster_id
		}
		self.connection.call("StopYarnCluster", request)


	def get_cluster(self, cluster_id):
		"""
		Returns cluster (*Cluster)
		"""
		request = {
			'cluster_id': cluster_id
		}
		response = self.connection.call("GetCluster", request)
		return View(response['cluster'])

	def get_yarn_cluster(self, cluster_id):
		"""
		Returns cluster (*YarnCluster)
		"""
		request = {
			'cluster_id': cluster_id
		}
		response = self.connection.call("GetYarnCluster", request)
		return View(response['cluster'])

	def get_clusters(self, offset, limit):
		"""
		Returns clusters ([]*Cluster)
		"""
		request = {
			'offset': offset,
			'limit': limit
		}
		response = self.connection.call("GetClusters", request)
		return [View(o) for o in response['clusters']]

	def get_cluster_status(self, cluster_id):
		"""
		Returns clusterStatus (*ClusterStatus)
		"""
		request = {
			'cluster_id': cluster_id
		}
		response = self.connection.call("GetClusterStatus", request)
		return View(response['cluster_status'])

	def delete_cluster(self, cluster_id):
		"""
		Returns None
		"""
		request = {
			'cluster_id': cluster_id
		}
		self.connection.call("DeleteCluster", request)


	def get_job(self, cluster_id, job_name):
		"""
		Returns job (*Job)
		"""
		request = {
			'cluster_id': cluster_id,
			'job_name': job_name
		}
		response = self.connection.call("GetJob", request)
		return View(response['job'])

	def get_jobs(self, cluster_id):
		"""
		Returns jobs ([]*Job)
		"""
		request = {
			'cluster_id': cluster_id
		}
		response = self.connection.call("GetJobs", request)
		return [View(o) for o in response['jobs']]

	def create_project(self, name, description):
		"""
		Returns projectId (int64)
		"""
		request = {
			'name': name,
			'description': description
		}
		response = self.connection.call("CreateProject", request)
		return response['project_id']

	def get_projects(self, offset, limit):
		"""
		Returns projects ([]*Project)
		"""
		request = {
			'offset': offset,
			'limit': limit
		}
		response = self.connection.call("GetProjects", request)
		return [View(o) for o in response['projects']]

	def get_project(self, project_id):
		"""
		Returns project (*Project)
		"""
		request = {
			'project_id': project_id
		}
		response = self.connection.call("GetProject", request)
		return View(response['project'])

	def delete_project(self, project_id):
		"""
		Returns None
		"""
		request = {
			'project_id': project_id
		}
		self.connection.call("DeleteProject", request)


	def create_datasource(self, project_id, name, description, path):
		"""
		Returns datasourceId (int64)
		"""
		request = {
			'project_id': project_id,
			'name': name,
			'description': description,
			'path': path
		}
		response = self.connection.call("CreateDatasource", request)
		return response['datasource_id']

	def get_datasources(self, project_id, offset, limit):
		"""
		Returns datasources ([]*Datasource)
		"""
		request = {
			'project_id': project_id,
			'offset': offset,
			'limit': limit
		}
		response = self.connection.call("GetDatasources", request)
		return [View(o) for o in response['datasources']]

	def get_datasource(self, datasource_id):
		"""
		Returns datasource (*Datasources)
		"""
		request = {
			'datasource_id': datasource_id
		}
		response = self.connection.call("GetDatasource", request)
		return View(response['datasource'])

	def update_datasource(self, name, description, path):
		"""
		Returns None
		"""
		request = {
			'name': name,
			'description': description,
			'path': path
		}
		self.connection.call("UpdateDatasource", request)


	def delete_datasource(self, datasource_id):
		"""
		Returns None
		"""
		request = {
			'datasource_id': datasource_id
		}
		self.connection.call("DeleteDatasource", request)


	def create_dataset(self, cluster_id, datasource_id, name, description, response_column_name):
		"""
		Returns datasetId (int64)
		"""
		request = {
			'cluster_id': cluster_id,
			'datasource_id': datasource_id,
			'name': name,
			'description': description,
			'response_column_name': response_column_name
		}
		response = self.connection.call("CreateDataset", request)
		return response['dataset_id']

	def get_datasets(self, datasource_id, offset, limit):
		"""
		Returns datasets ([]*Dataset)
		"""
		request = {
			'datasource_id': datasource_id,
			'offset': offset,
			'limit': limit
		}
		response = self.connection.call("GetDatasets", request)
		return [View(o) for o in response['datasets']]

	def get_dataset(self, dataset_id):
		"""
		Returns dataset (*Dataset)
		"""
		request = {
			'dataset_id': dataset_id
		}
		response = self.connection.call("GetDataset", request)
		return View(response['dataset'])

	def update_dataset(self, dataset_id, name, description, response_column_name):
		"""
		Returns None
		"""
		request = {
			'dataset_id': dataset_id,
			'name': name,
			'description': description,
			'response_column_name': response_column_name
		}
		self.connection.call("UpdateDataset", request)


	def split_dataset(self, dataset_id, ratio1, ratio2):
		"""
		Returns datasetIds ([]int64)
		"""
		request = {
			'dataset_id': dataset_id,
			'ratio1': ratio1,
			'ratio2': ratio2
		}
		response = self.connection.call("SplitDataset", request)
		return response['dataset_ids']

	def delete_dataset(self, dataset_id):
		"""
		Returns None
		"""
		request = {
			'dataset_id': dataset_id
		}
		self.connection.call("DeleteDataset", request)


	def build_model(self, cluster_id, dataset_id, algorithm):
		"""
		Returns modelId (int64)
		"""
		request = {
			'cluster_id': cluster_id,
			'dataset_id': dataset_id,
			'algorithm': algorithm
		}
		response = self.connection.call("BuildModel", request)
		return response['model_id']

	def build_auto_model(self, cluster_id, dataset, target_name, max_run_time):
		"""
		Returns model (*Model)
		"""
		request = {
			'cluster_id': cluster_id,
			'dataset': dataset,
			'target_name': target_name,
			'max_run_time': max_run_time
		}
		response = self.connection.call("BuildAutoModel", request)
		return View(response['model'])

	def get_model(self, model_id):
		"""
		Returns model (*Model)
		"""
		request = {
			'model_id': model_id
		}
		response = self.connection.call("GetModel", request)
		return View(response['model'])

	def get_models(self, project_id, offset, limit):
		"""
		Returns models ([]*Model)
		"""
		request = {
			'project_id': project_id,
			'offset': offset,
			'limit': limit
		}
		response = self.connection.call("GetModels", request)
		return [View(o) for o in response['models']]

	def get_cluster_models(self, cluster_id):
		"""
		Returns models ([]*Model)
		"""
		request = {
			'cluster_id': cluster_id
		}
		response = self.connection.call("GetClusterModels", request)
		return [View(o) for o in response['models']]

	def import_model_from_cluster(self, cluster_id, project_id, model_name):
		"""
		Returns model (*Model)
		"""
		request = {
			'cluster_id': cluster_id,
			'project_id': project_id,
			'model_name': model_name
		}
		response = self.connection.call("ImportModelFromCluster", request)
		return View(response['model'])

	def delete_model(self, model_id):
		"""
		Returns None
		"""
		request = {
			'model_id': model_id
		}
		self.connection.call("DeleteModel", request)


	def start_scoring_service(self, model_id, port):
		"""
		Returns service (*ScoringService)
		"""
		request = {
			'model_id': model_id,
			'port': port
		}
		response = self.connection.call("StartScoringService", request)
		return View(response['service'])

	def stop_scoring_service(self, service_id):
		"""
		Returns None
		"""
		request = {
			'service_id': service_id
		}
		self.connection.call("StopScoringService", request)


	def get_scoring_service(self, service_id):
		"""
		Returns service (*ScoringService)
		"""
		request = {
			'service_id': service_id
		}
		response = self.connection.call("GetScoringService", request)
		return View(response['service'])

	def get_scoring_services(self, offset, limit):
		"""
		Returns services ([]*ScoringService)
		"""
		request = {
			'offset': offset,
			'limit': limit
		}
		response = self.connection.call("GetScoringServices", request)
		return [View(o) for o in response['services']]

	def get_scoring_services_for_model(self, model_id, offset, limit):
		"""
		Returns services ([]*ScoringService)
		"""
		request = {
			'model_id': model_id,
			'offset': offset,
			'limit': limit
		}
		response = self.connection.call("GetScoringServicesForModel", request)
		return [View(o) for o in response['services']]

	def delete_scoring_service(self, service_id):
		"""
		Returns None
		"""
		request = {
			'service_id': service_id
		}
		self.connection.call("DeleteScoringService", request)


	def add_engine(self, engine_name, engine_path):
		"""
		Returns engineId (int64)
		"""
		request = {
			'engine_name': engine_name,
			'engine_path': engine_path
		}
		response = self.connection.call("AddEngine", request)
		return response['engine_id']

	def get_engine(self, engine_id):
		"""
		Returns engine (*Engine)
		"""
		request = {
			'engine_id': engine_id
		}
		response = self.connection.call("GetEngine", request)
		return View(response['engine'])

	def get_engines(self):
		"""
		Returns engines ([]*Engine)
		"""
		request = {
		}
		response = self.connection.call("GetEngines", request)
		return [View(o) for o in response['engines']]

	def delete_engine(self, engine_id):
		"""
		Returns None
		"""
		request = {
			'engine_id': engine_id
		}
		self.connection.call("DeleteEngine", request)


	def get_supported_entity_types(self):
		"""
		Returns entityTypes ([]*EntityType)
		"""
		request = {
		}
		response = self.connection.call("GetSupportedEntityTypes", request)
		return [View(o) for o in response['entity_types']]

	def get_supported_permissions(self):
		"""
		Returns permissions ([]*Permission)
		"""
		request = {
		}
		response = self.connection.call("GetSupportedPermissions", request)
		return [View(o) for o in response['permissions']]

	def get_supported_cluster_types(self):
		"""
		Returns clusterTypes ([]*ClusterType)
		"""
		request = {
		}
		response = self.connection.call("GetSupportedClusterTypes", request)
		return [View(o) for o in response['cluster_types']]

	def get_permissions_for_role(self, role_id):
		"""
		Returns permissions ([]*Permission)
		"""
		request = {
			'role_id': role_id
		}
		response = self.connection.call("GetPermissionsForRole", request)
		return [View(o) for o in response['permissions']]

	def get_permissions_for_identity(self, identity_id):
		"""
		Returns permissions ([]*Permission)
		"""
		request = {
			'identity_id': identity_id
		}
		response = self.connection.call("GetPermissionsForIdentity", request)
		return [View(o) for o in response['permissions']]

	def create_role(self, name, description):
		"""
		Returns roleId (int64)
		"""
		request = {
			'name': name,
			'description': description
		}
		response = self.connection.call("CreateRole", request)
		return response['role_id']

	def get_roles(self, offset, limit):
		"""
		Returns roles ([]*Role)
		"""
		request = {
			'offset': offset,
			'limit': limit
		}
		response = self.connection.call("GetRoles", request)
		return [View(o) for o in response['roles']]

	def get_roles_for_identity(self, identity_id):
		"""
		Returns roles ([]*Role)
		"""
		request = {
			'identity_id': identity_id
		}
		response = self.connection.call("GetRolesForIdentity", request)
		return [View(o) for o in response['roles']]

	def get_role(self, role_id):
		"""
		Returns role (*Role)
		"""
		request = {
			'role_id': role_id
		}
		response = self.connection.call("GetRole", request)
		return View(response['role'])

	def get_role_by_name(self, name):
		"""
		Returns role (*Role)
		"""
		request = {
			'name': name
		}
		response = self.connection.call("GetRoleByName", request)
		return View(response['role'])

	def update_role(self, role_id, name, description):
		"""
		Returns None
		"""
		request = {
			'role_id': role_id,
			'name': name,
			'description': description
		}
		self.connection.call("UpdateRole", request)


	def link_role_and_permissions(self, role_id, permission_ids):
		"""
		Returns None
		"""
		request = {
			'role_id': role_id,
			'permission_ids': permission_ids
		}
		self.connection.call("LinkRoleAndPermissions", request)


	def delete_role(self, role_id):
		"""
		Returns None
		"""
		request = {
			'role_id': role_id
		}
		self.connection.call("DeleteRole", request)


	def create_workgroup(self, name, description):
		"""
		Returns workgroupId (int64)
		"""
		request = {
			'name': name,
			'description': description
		}
		response = self.connection.call("CreateWorkgroup", request)
		return response['workgroup_id']

	def get_workgroups(self, offset, limit):
		"""
		Returns workgroups ([]*Workgroup)
		"""
		request = {
			'offset': offset,
			'limit': limit
		}
		response = self.connection.call("GetWorkgroups", request)
		return [View(o) for o in response['workgroups']]

	def get_workgroups_for_identity(self, identity_id):
		"""
		Returns workgroups ([]*Workgroup)
		"""
		request = {
			'identity_id': identity_id
		}
		response = self.connection.call("GetWorkgroupsForIdentity", request)
		return [View(o) for o in response['workgroups']]

	def get_workgroup(self, workgroup_id):
		"""
		Returns workgroup (*Workgroup)
		"""
		request = {
			'workgroup_id': workgroup_id
		}
		response = self.connection.call("GetWorkgroup", request)
		return View(response['workgroup'])

	def get_workgroup_by_name(self, name):
		"""
		Returns workgroup (*Workgroup)
		"""
		request = {
			'name': name
		}
		response = self.connection.call("GetWorkgroupByName", request)
		return View(response['workgroup'])

	def update_workgroup(self, workgroup_id, name, description):
		"""
		Returns None
		"""
		request = {
			'workgroup_id': workgroup_id,
			'name': name,
			'description': description
		}
		self.connection.call("UpdateWorkgroup", request)


	def delete_workgroup(self, workgroup_id):
		"""
		Returns None
		"""
		request = {
			'workgroup_id': workgroup_id
		}
		self.connection.call("DeleteWorkgroup", request)


	def create_identity(self, name, password):
		"""
		Returns identityId (int64)
		"""
		request = {
			'name': name,
			'password': password
		}
		response = self.connection.call("CreateIdentity", request)
		return response['identity_id']

	def get_identities(self, offset, limit):
		"""
		Returns identities ([]*Identity)
		"""
		request = {
			'offset': offset,
			'limit': limit
		}
		response = self.connection.call("GetIdentities", request)
		return [View(o) for o in response['identities']]

	def get_identities_for_workgroup(self, workgroup_id):
		"""
		Returns identities ([]*Identity)
		"""
		request = {
			'workgroup_id': workgroup_id
		}
		response = self.connection.call("GetIdentitiesForWorkgroup", request)
		return [View(o) for o in response['identities']]

	def get_identities_for_role(self, role_id):
		"""
		Returns identities ([]*Identity)
		"""
		request = {
			'role_id': role_id
		}
		response = self.connection.call("GetIdentitiesForRole", request)
		return [View(o) for o in response['identities']]

	def get_identity(self, identity_id):
		"""
		Returns identity (*Identity)
		"""
		request = {
			'identity_id': identity_id
		}
		response = self.connection.call("GetIdentity", request)
		return View(response['identity'])

	def get_identity_by_name(self, name):
		"""
		Returns identity (*Identity)
		"""
		request = {
			'name': name
		}
		response = self.connection.call("GetIdentityByName", request)
		return View(response['identity'])

	def link_identity_and_workgroup(self, identity_id, workgroup_id):
		"""
		Returns None
		"""
		request = {
			'identity_id': identity_id,
			'workgroup_id': workgroup_id
		}
		self.connection.call("LinkIdentityAndWorkgroup", request)


	def unlink_identity_and_workgroup(self, identity_id, workgroup_id):
		"""
		Returns None
		"""
		request = {
			'identity_id': identity_id,
			'workgroup_id': workgroup_id
		}
		self.connection.call("UnlinkIdentityAndWorkgroup", request)


	def link_identity_and_role(self, identity_id, role_id):
		"""
		Returns None
		"""
		request = {
			'identity_id': identity_id,
			'role_id': role_id
		}
		self.connection.call("LinkIdentityAndRole", request)


	def unlink_identity_and_role(self, identity_id, role_id):
		"""
		Returns None
		"""
		request = {
			'identity_id': identity_id,
			'role_id': role_id
		}
		self.connection.call("UnlinkIdentityAndRole", request)


	def update_identity(self, identity_id, password):
		"""
		Returns None
		"""
		request = {
			'identity_id': identity_id,
			'password': password
		}
		self.connection.call("UpdateIdentity", request)


	def deactivate_identity(self, identity_id):
		"""
		Returns None
		"""
		request = {
			'identity_id': identity_id
		}
		self.connection.call("DeactivateIdentity", request)


	def share_entity(self, kind, workgroup_id, entity_type_id, entity_id):
		"""
		Returns None
		"""
		request = {
			'kind': kind,
			'workgroup_id': workgroup_id,
			'entity_type_id': entity_type_id,
			'entity_id': entity_id
		}
		self.connection.call("ShareEntity", request)


	def get_entity_privileges(self, entity_type_id, entity_id):
		"""
		Returns privileges ([]*EntityPrivilege)
		"""
		request = {
			'entity_type_id': entity_type_id,
			'entity_id': entity_id
		}
		response = self.connection.call("GetEntityPrivileges", request)
		return [View(o) for o in response['privileges']]

	def unshare_entity(self, kind, workgroup_id, entity_type_id, entity_id):
		"""
		Returns None
		"""
		request = {
			'kind': kind,
			'workgroup_id': workgroup_id,
			'entity_type_id': entity_type_id,
			'entity_id': entity_id
		}
		self.connection.call("UnshareEntity", request)


	def get_entity_history(self, entity_type_id, entity_id, offset, limit):
		"""
		Returns history ([]*EntityHistory)
		"""
		request = {
			'entity_type_id': entity_type_id,
			'entity_id': entity_id,
			'offset': offset,
			'limit': limit
		}
		response = self.connection.call("GetEntityHistory", request)
		return [View(o) for o in response['history']]