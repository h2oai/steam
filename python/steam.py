# ------------------------------
# --- This is generated code ---
# ---      DO NOT EDIT       ---
# ------------------------------


import httplib
import base64
import string
import json
import sys
import logging
import ssl
from collections import namedtuple

class RPCError(Exception):
	def __init__(self, value):
		self.value = value
	def __str__(self):
		return repr(self.value)

class HTTPConnection:
	def __init__(self, host, port, username, password, verify_ssl=True):
		self.host = host
		self.port = port
		self.username = username
		self.password = password
		self.verify_ssl = verify_ssl
		self.uid = 0

	def call(self, method, params):
		self.uid = self.uid + 1
		request = {
			'id': self.uid,
			'method': 'web.' + method,
			'params': [params]
		}
		payload = json.dumps(request)
		ssl._https_verify_certificates(self.verify_ssl)
		ws = httplib.HTTPS(self.host, self.port)
		
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

steamconn = None

def connect(host, port, username, passw, verify_ssl=True):
	global steamconn
	steamconn = HTTPConnection(host, port, username, passw, verify_ssl)

def get_cluster_connection(self, cluster):
	proxconf = self.get_config()
	cport = proxconf['cluster_proxy_address'].split(':', 1)[1]
	conf = {'https':True, 'verify_ssl_certificates':self.connection.verify_ssl, \
		'port': int(cport), \
		'context_path':'%s_%s' % (proxconf['username'], cluster['name']), \
		'cookies':["%s=%s" % (cluster['name'], cluster['token'])], \
		'ip':self.connection.host}
	return conf
	
	


def ping_server(input):
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
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("PingServer", request)
	return response['output']

def get_config():
	"""
	Get Steam start up configurations

	Parameters:

	Returns:
	config: An object containing Steam startup configurations (Config)
	"""
	request = {
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetConfig", request)
	return response['config']

def check_admin():
	"""
	Check if an identity has admin privileges

	Parameters:

	Returns:
	is_admin: No description available (bool)
	"""
	request = {
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("CheckAdmin", request)
	return response['is_admin']

def set_local_config():
	"""
	Set security configuration to local

	Parameters:

	Returns:None
	"""
	request = {
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("SetLocalConfig", request)
	return 

def set_ldap_config(config):
	"""
	Set LDAP security configuration

	Parameters:
	config: No description available (LdapConfig)

	Returns:None
	"""
	request = {
		'config': config
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("SetLdapConfig", request)
	return 

def get_ldap_config():
	"""
	Get LDAP security configurations

	Parameters:

	Returns:
	config: No description available (LdapConfig)
	exists: No description available (bool)
	"""
	request = {
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetLdapConfig", request)
	return response['config'], response['exists']

def test_ldap_config(config):
	"""
	Test LDAP security configurations

	Parameters:
	config: No description available (LdapConfig)

	Returns:None
	"""
	request = {
		'config': config
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("TestLdapConfig", request)
	return 

def register_cluster(address):
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
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("RegisterCluster", request)
	return response['cluster_id']

def unregister_cluster(cluster_id):
	"""
	Disconnect from a cluster

	Parameters:
	cluster_id: No description available (int64)

	Returns:None
	"""
	request = {
		'cluster_id': cluster_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("UnregisterCluster", request)
	return 

def start_cluster_on_yarn(cluster_name,engine_id,size,memory,secure,keytab):
	"""
	Start a cluster using Yarn

	Parameters:
	cluster_name: No description available (string)
	engine_id: No description available (int64)
	size: No description available (int)
	memory: No description available (string)
	secure: No description available (bool)
	keytab: No description available (string)

	Returns:
	cluster_id: No description available (int64)
	"""
	request = {
		'cluster_name': cluster_name,
		'engine_id': engine_id,
		'size': size,
		'memory': memory,
		'secure': secure,
		'keytab': keytab
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("StartClusterOnYarn", request)
	return response['cluster_id']

def stop_cluster_on_yarn(cluster_id,keytab):
	"""
	Stop a cluster using Yarn

	Parameters:
	cluster_id: No description available (int64)
	keytab: No description available (string)

	Returns:None
	"""
	request = {
		'cluster_id': cluster_id,
		'keytab': keytab
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("StopClusterOnYarn", request)
	return 

def get_cluster(cluster_id):
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
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetCluster", request)
	return response['cluster']

def get_cluster_on_yarn(cluster_id):
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
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetClusterOnYarn", request)
	return response['cluster']

def get_clusters(offset,limit):
	"""
	List clusters

	Parameters:
	offset: No description available (uint)
	limit: No description available (uint)

	Returns:
	clusters: No description available (Cluster)
	"""
	request = {
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetClusters", request)
	return response['clusters']

def get_cluster_status(cluster_id):
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
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetClusterStatus", request)
	return response['cluster_status']

def delete_cluster(cluster_id):
	"""
	Delete a cluster

	Parameters:
	cluster_id: No description available (int64)

	Returns:None
	"""
	request = {
		'cluster_id': cluster_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("DeleteCluster", request)
	return 

def get_job(cluster_id,job_name):
	"""
	Get job details

	Parameters:
	cluster_id: No description available (int64)
	job_name: No description available (string)

	Returns:
	job: No description available (Job)
	"""
	request = {
		'cluster_id': cluster_id,
		'job_name': job_name
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetJob", request)
	return response['job']

def get_jobs(cluster_id):
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
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetJobs", request)
	return response['jobs']

def create_project(name,description,model_category):
	"""
	Create a project

	Parameters:
	name: No description available (string)
	description: No description available (string)
	model_category: No description available (string)

	Returns:
	project_id: No description available (int64)
	"""
	request = {
		'name': name,
		'description': description,
		'model_category': model_category
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("CreateProject", request)
	return response['project_id']

def get_projects(offset,limit):
	"""
	List projects

	Parameters:
	offset: No description available (uint)
	limit: No description available (uint)

	Returns:
	projects: No description available (Project)
	"""
	request = {
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetProjects", request)
	return response['projects']

def get_project(project_id):
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
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetProject", request)
	return response['project']

def delete_project(project_id):
	"""
	Delete a project

	Parameters:
	project_id: No description available (int64)

	Returns:None
	"""
	request = {
		'project_id': project_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("DeleteProject", request)
	return 

def create_datasource(project_id,name,description,path):
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
		'project_id': project_id,
		'name': name,
		'description': description,
		'path': path
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("CreateDatasource", request)
	return response['datasource_id']

def get_datasources(project_id,offset,limit):
	"""
	List datasources

	Parameters:
	project_id: No description available (int64)
	offset: No description available (uint)
	limit: No description available (uint)

	Returns:
	datasources: No description available (Datasource)
	"""
	request = {
		'project_id': project_id,
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetDatasources", request)
	return response['datasources']

def get_datasource(datasource_id):
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
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetDatasource", request)
	return response['datasource']

def update_datasource(datasource_id,name,description,path):
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
		'datasource_id': datasource_id,
		'name': name,
		'description': description,
		'path': path
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("UpdateDatasource", request)
	return 

def delete_datasource(datasource_id):
	"""
	Delete a datasource

	Parameters:
	datasource_id: No description available (int64)

	Returns:None
	"""
	request = {
		'datasource_id': datasource_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("DeleteDatasource", request)
	return 

def create_dataset(cluster_id,datasource_id,name,description,response_column_name):
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
		'cluster_id': cluster_id,
		'datasource_id': datasource_id,
		'name': name,
		'description': description,
		'response_column_name': response_column_name
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("CreateDataset", request)
	return response['dataset_id']

def get_datasets(datasource_id,offset,limit):
	"""
	List datasets

	Parameters:
	datasource_id: No description available (int64)
	offset: No description available (uint)
	limit: No description available (uint)

	Returns:
	datasets: No description available (Dataset)
	"""
	request = {
		'datasource_id': datasource_id,
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetDatasets", request)
	return response['datasets']

def get_dataset(dataset_id):
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
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetDataset", request)
	return response['dataset']

def get_datasets_from_cluster(cluster_id):
	"""
	Get a list of datasets on a cluster

	Parameters:
	cluster_id: No description available (int64)

	Returns:
	dataset: No description available (Dataset)
	"""
	request = {
		'cluster_id': cluster_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetDatasetsFromCluster", request)
	return response['dataset']

def update_dataset(dataset_id,name,description,response_column_name):
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
		'dataset_id': dataset_id,
		'name': name,
		'description': description,
		'response_column_name': response_column_name
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("UpdateDataset", request)
	return 

def split_dataset(dataset_id,ratio1,ratio2):
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
		'dataset_id': dataset_id,
		'ratio1': ratio1,
		'ratio2': ratio2
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("SplitDataset", request)
	return response['dataset_ids']

def delete_dataset(dataset_id):
	"""
	Delete a dataset

	Parameters:
	dataset_id: No description available (int64)

	Returns:None
	"""
	request = {
		'dataset_id': dataset_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("DeleteDataset", request)
	return 

def build_model(cluster_id,dataset_id,algorithm):
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
		'cluster_id': cluster_id,
		'dataset_id': dataset_id,
		'algorithm': algorithm
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("BuildModel", request)
	return response['model_id']

def build_model_auto(cluster_id,dataset,target_name,max_run_time):
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
		'cluster_id': cluster_id,
		'dataset': dataset,
		'target_name': target_name,
		'max_run_time': max_run_time
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("BuildModelAuto", request)
	return response['model']

def get_model(model_id):
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
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetModel", request)
	return response['model']

def get_models(project_id,offset,limit):
	"""
	List models

	Parameters:
	project_id: No description available (int64)
	offset: No description available (uint)
	limit: No description available (uint)

	Returns:
	models: No description available (Model)
	"""
	request = {
		'project_id': project_id,
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetModels", request)
	return response['models']

def get_models_from_cluster(cluster_id,frame_key):
	"""
	List models from a cluster

	Parameters:
	cluster_id: No description available (int64)
	frame_key: No description available (string)

	Returns:
	models: No description available (Model)
	"""
	request = {
		'cluster_id': cluster_id,
		'frame_key': frame_key
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetModelsFromCluster", request)
	return response['models']

def find_models_count(project_id):
	"""
	Get a count models in a project

	Parameters:
	project_id: No description available (int64)

	Returns:
	count: No description available (int64)
	"""
	request = {
		'project_id': project_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("FindModelsCount", request)
	return response['count']

def get_all_binomial_sort_criteria():
	"""
	List sort criteria for a binomial models

	Parameters:

	Returns:
	criteria: No description available (string)
	"""
	request = {
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetAllBinomialSortCriteria", request)
	return response['criteria']

def find_models_binomial(project_id,name_part,sort_by,ascending,offset,limit):
	"""
	List binomial models

	Parameters:
	project_id: No description available (int64)
	name_part: No description available (string)
	sort_by: No description available (string)
	ascending: No description available (bool)
	offset: No description available (uint)
	limit: No description available (uint)

	Returns:
	models: No description available (BinomialModel)
	"""
	request = {
		'project_id': project_id,
		'name_part': name_part,
		'sort_by': sort_by,
		'ascending': ascending,
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("FindModelsBinomial", request)
	return response['models']

def get_model_binomial(model_id):
	"""
	View a binomial model

	Parameters:
	model_id: No description available (int64)

	Returns:
	model: No description available (BinomialModel)
	"""
	request = {
		'model_id': model_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetModelBinomial", request)
	return response['model']

def get_all_multinomial_sort_criteria():
	"""
	List sort criteria for a multinomial models

	Parameters:

	Returns:
	criteria: No description available (string)
	"""
	request = {
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetAllMultinomialSortCriteria", request)
	return response['criteria']

def find_models_multinomial(project_id,name_part,sort_by,ascending,offset,limit):
	"""
	List multinomial models

	Parameters:
	project_id: No description available (int64)
	name_part: No description available (string)
	sort_by: No description available (string)
	ascending: No description available (bool)
	offset: No description available (uint)
	limit: No description available (uint)

	Returns:
	models: No description available (MultinomialModel)
	"""
	request = {
		'project_id': project_id,
		'name_part': name_part,
		'sort_by': sort_by,
		'ascending': ascending,
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("FindModelsMultinomial", request)
	return response['models']

def get_model_multinomial(model_id):
	"""
	View a binomial model

	Parameters:
	model_id: No description available (int64)

	Returns:
	model: No description available (MultinomialModel)
	"""
	request = {
		'model_id': model_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetModelMultinomial", request)
	return response['model']

def get_all_regression_sort_criteria():
	"""
	List sort criteria for a regression models

	Parameters:

	Returns:
	criteria: No description available (string)
	"""
	request = {
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetAllRegressionSortCriteria", request)
	return response['criteria']

def find_models_regression(project_id,name_part,sort_by,ascending,offset,limit):
	"""
	List regression models

	Parameters:
	project_id: No description available (int64)
	name_part: No description available (string)
	sort_by: No description available (string)
	ascending: No description available (bool)
	offset: No description available (uint)
	limit: No description available (uint)

	Returns:
	models: No description available (RegressionModel)
	"""
	request = {
		'project_id': project_id,
		'name_part': name_part,
		'sort_by': sort_by,
		'ascending': ascending,
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("FindModelsRegression", request)
	return response['models']

def get_model_regression(model_id):
	"""
	View a binomial model

	Parameters:
	model_id: No description available (int64)

	Returns:
	model: No description available (RegressionModel)
	"""
	request = {
		'model_id': model_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetModelRegression", request)
	return response['model']

def import_model_from_cluster(cluster_id,project_id,model_key,model_name):
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
		'cluster_id': cluster_id,
		'project_id': project_id,
		'model_key': model_key,
		'model_name': model_name
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("ImportModelFromCluster", request)
	return response['model_id']

def check_mojo(algo):
	"""
	Check if a model category can generate MOJOs

	Parameters:
	algo: No description available (string)

	Returns:
	can_mojo: No description available (bool)
	"""
	request = {
		'algo': algo
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("CheckMojo", request)
	return response['can_mojo']

def import_model_pojo(model_id):
	"""
	Import a model's POJO from a cluster

	Parameters:
	model_id: No description available (int64)

	Returns:None
	"""
	request = {
		'model_id': model_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("ImportModelPojo", request)
	return 

def import_model_mojo(model_id):
	"""
	Import a model's MOJO from a cluster

	Parameters:
	model_id: No description available (int64)

	Returns:None
	"""
	request = {
		'model_id': model_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("ImportModelMojo", request)
	return 

def delete_model(model_id):
	"""
	Delete a model

	Parameters:
	model_id: No description available (int64)

	Returns:None
	"""
	request = {
		'model_id': model_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("DeleteModel", request)
	return 

def create_label(project_id,name,description):
	"""
	Create a label

	Parameters:
	project_id: No description available (int64)
	name: No description available (string)
	description: No description available (string)

	Returns:
	label_id: No description available (int64)
	"""
	request = {
		'project_id': project_id,
		'name': name,
		'description': description
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("CreateLabel", request)
	return response['label_id']

def update_label(label_id,name,description):
	"""
	Update a label

	Parameters:
	label_id: No description available (int64)
	name: No description available (string)
	description: No description available (string)

	Returns:None
	"""
	request = {
		'label_id': label_id,
		'name': name,
		'description': description
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("UpdateLabel", request)
	return 

def delete_label(label_id):
	"""
	Delete a label

	Parameters:
	label_id: No description available (int64)

	Returns:None
	"""
	request = {
		'label_id': label_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("DeleteLabel", request)
	return 

def link_label_with_model(label_id,model_id):
	"""
	Label a model

	Parameters:
	label_id: No description available (int64)
	model_id: No description available (int64)

	Returns:None
	"""
	request = {
		'label_id': label_id,
		'model_id': model_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("LinkLabelWithModel", request)
	return 

def unlink_label_from_model(label_id,model_id):
	"""
	Remove a label from a model

	Parameters:
	label_id: No description available (int64)
	model_id: No description available (int64)

	Returns:None
	"""
	request = {
		'label_id': label_id,
		'model_id': model_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("UnlinkLabelFromModel", request)
	return 

def get_labels_for_project(project_id):
	"""
	List labels for a project, with corresponding models, if any

	Parameters:
	project_id: No description available (int64)

	Returns:
	labels: No description available (Label)
	"""
	request = {
		'project_id': project_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetLabelsForProject", request)
	return response['labels']

def start_service(model_id,name,package_name):
	"""
	Start a service

	Parameters:
	model_id: No description available (int64)
	name: No description available (string)
	package_name: No description available (string)

	Returns:
	service_id: No description available (int64)
	"""
	request = {
		'model_id': model_id,
		'name': name,
		'package_name': package_name
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("StartService", request)
	return response['service_id']

def stop_service(service_id):
	"""
	Stop a service

	Parameters:
	service_id: No description available (int64)

	Returns:None
	"""
	request = {
		'service_id': service_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("StopService", request)
	return 

def get_service(service_id):
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
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetService", request)
	return response['service']

def get_services(offset,limit):
	"""
	List all services

	Parameters:
	offset: No description available (uint)
	limit: No description available (uint)

	Returns:
	services: No description available (ScoringService)
	"""
	request = {
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetServices", request)
	return response['services']

def get_services_for_project(project_id,offset,limit):
	"""
	List services for a project

	Parameters:
	project_id: No description available (int64)
	offset: No description available (uint)
	limit: No description available (uint)

	Returns:
	services: No description available (ScoringService)
	"""
	request = {
		'project_id': project_id,
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetServicesForProject", request)
	return response['services']

def get_services_for_model(model_id,offset,limit):
	"""
	List services for a model

	Parameters:
	model_id: No description available (int64)
	offset: No description available (uint)
	limit: No description available (uint)

	Returns:
	services: No description available (ScoringService)
	"""
	request = {
		'model_id': model_id,
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetServicesForModel", request)
	return response['services']

def delete_service(service_id):
	"""
	Delete a service

	Parameters:
	service_id: No description available (int64)

	Returns:None
	"""
	request = {
		'service_id': service_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("DeleteService", request)
	return 

def get_engine(engine_id):
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
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetEngine", request)
	return response['engine']

def get_engines():
	"""
	List engines

	Parameters:

	Returns:
	engines: No description available (Engine)
	"""
	request = {
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetEngines", request)
	return response['engines']

def delete_engine(engine_id):
	"""
	Delete an engine

	Parameters:
	engine_id: No description available (int64)

	Returns:None
	"""
	request = {
		'engine_id': engine_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("DeleteEngine", request)
	return 

def get_all_entity_types():
	"""
	List all entity types

	Parameters:

	Returns:
	entity_types: A list of Steam entity types. (EntityType)
	"""
	request = {
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetAllEntityTypes", request)
	return response['entity_types']

def get_all_permissions():
	"""
	List all permissions

	Parameters:

	Returns:
	permissions: A list of Steam permissions. (Permission)
	"""
	request = {
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetAllPermissions", request)
	return response['permissions']

def get_all_cluster_types():
	"""
	List all cluster types

	Parameters:

	Returns:
	cluster_types: No description available (ClusterType)
	"""
	request = {
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetAllClusterTypes", request)
	return response['cluster_types']

def get_permissions_for_role(role_id):
	"""
	List permissions for a role

	Parameters:
	role_id: Integer ID of a role in Steam. (int64)

	Returns:
	permissions: A list of Steam permissions. (Permission)
	"""
	request = {
		'role_id': role_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetPermissionsForRole", request)
	return response['permissions']

def get_permissions_for_identity(identity_id):
	"""
	List permissions for an identity

	Parameters:
	identity_id: Integer ID of an identity in Steam. (int64)

	Returns:
	permissions: A list of Steam permissions. (Permission)
	"""
	request = {
		'identity_id': identity_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetPermissionsForIdentity", request)
	return response['permissions']

def create_role(name,description):
	"""
	Create a role

	Parameters:
	name: A string name. (string)
	description: A string description (string)

	Returns:
	role_id: Integer ID of the role in Steam. (int64)
	"""
	request = {
		'name': name,
		'description': description
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("CreateRole", request)
	return response['role_id']

def get_roles(offset,limit):
	"""
	List roles

	Parameters:
	offset: An offset uint start the search on. (uint)
	limit: The maximum uint objects. (uint)

	Returns:
	roles: A list of Steam roles. (Role)
	"""
	request = {
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetRoles", request)
	return response['roles']

def get_roles_for_identity(identity_id):
	"""
	List roles for an identity

	Parameters:
	identity_id: Integer ID of an identity in Steam. (int64)

	Returns:
	roles: A list of Steam roles. (Role)
	"""
	request = {
		'identity_id': identity_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetRolesForIdentity", request)
	return response['roles']

def get_role(role_id):
	"""
	Get role details

	Parameters:
	role_id: Integer ID of a role in Steam. (int64)

	Returns:
	role: A Steam role. (Role)
	"""
	request = {
		'role_id': role_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetRole", request)
	return response['role']

def get_role_by_name(name):
	"""
	Get role details by name

	Parameters:
	name: A role name. (string)

	Returns:
	role: A Steam role. (Role)
	"""
	request = {
		'name': name
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetRoleByName", request)
	return response['role']

def update_role(role_id,name,description):
	"""
	Update a role

	Parameters:
	role_id: Integer ID of a role in Steam. (int64)
	name: A string name. (string)
	description: A string description (string)

	Returns:None
	"""
	request = {
		'role_id': role_id,
		'name': name,
		'description': description
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("UpdateRole", request)
	return 

def link_role_with_permissions(role_id,permission_ids):
	"""
	Link a role with permissions

	Parameters:
	role_id: Integer ID of a role in Steam. (int64)
	permission_ids: A list of Integer IDs for permissions in Steam. (int64)

	Returns:None
	"""
	request = {
		'role_id': role_id,
		'permission_ids': permission_ids
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("LinkRoleWithPermissions", request)
	return 

def link_role_with_permission(role_id,permission_id):
	"""
	Link a role with a permission

	Parameters:
	role_id: Integer ID of a role in Steam. (int64)
	permission_id: Integer ID of a permission in Steam. (int64)

	Returns:None
	"""
	request = {
		'role_id': role_id,
		'permission_id': permission_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("LinkRoleWithPermission", request)
	return 

def unlink_role_from_permission(role_id,permission_id):
	"""
	Unlink a role from a permission

	Parameters:
	role_id: Integer ID of a role in Steam. (int64)
	permission_id: Integer ID of a permission in Steam. (int64)

	Returns:None
	"""
	request = {
		'role_id': role_id,
		'permission_id': permission_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("UnlinkRoleFromPermission", request)
	return 

def delete_role(role_id):
	"""
	Delete a role

	Parameters:
	role_id: Integer ID of a role in Steam. (int64)

	Returns:None
	"""
	request = {
		'role_id': role_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("DeleteRole", request)
	return 

def create_workgroup(name,description):
	"""
	Create a workgroup

	Parameters:
	name: A string name. (string)
	description: A string description (string)

	Returns:
	workgroup_id: Integer ID of the workgroup in Steam. (int64)
	"""
	request = {
		'name': name,
		'description': description
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("CreateWorkgroup", request)
	return response['workgroup_id']

def get_workgroups(offset,limit):
	"""
	List workgroups

	Parameters:
	offset: An offset uint start the search on. (uint)
	limit: The maximum uint objects. (uint)

	Returns:
	workgroups: A list of workgroups in Steam. (Workgroup)
	"""
	request = {
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetWorkgroups", request)
	return response['workgroups']

def get_workgroups_for_identity(identity_id):
	"""
	List workgroups for an identity

	Parameters:
	identity_id: Integer ID of an identity in Steam. (int64)

	Returns:
	workgroups: A list of workgroups in Steam. (Workgroup)
	"""
	request = {
		'identity_id': identity_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetWorkgroupsForIdentity", request)
	return response['workgroups']

def get_workgroup(workgroup_id):
	"""
	Get workgroup details

	Parameters:
	workgroup_id: Integer ID of a workgroup in Steam. (int64)

	Returns:
	workgroup: A workgroup in Steam. (Workgroup)
	"""
	request = {
		'workgroup_id': workgroup_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetWorkgroup", request)
	return response['workgroup']

def get_workgroup_by_name(name):
	"""
	Get workgroup details by name

	Parameters:
	name: A string name. (string)

	Returns:
	workgroup: A workgroup in Steam. (Workgroup)
	"""
	request = {
		'name': name
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetWorkgroupByName", request)
	return response['workgroup']

def update_workgroup(workgroup_id,name,description):
	"""
	Update a workgroup

	Parameters:
	workgroup_id: Integer ID of a workgrou in Steam. (int64)
	name: A string name. (string)
	description: A string description (string)

	Returns:None
	"""
	request = {
		'workgroup_id': workgroup_id,
		'name': name,
		'description': description
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("UpdateWorkgroup", request)
	return 

def delete_workgroup(workgroup_id):
	"""
	Delete a workgroup

	Parameters:
	workgroup_id: Integer ID of a workgroup in Steam. (int64)

	Returns:None
	"""
	request = {
		'workgroup_id': workgroup_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("DeleteWorkgroup", request)
	return 

def create_identity(name,password):
	"""
	Create an identity

	Parameters:
	name: A string name. (string)
	password: A string password (string)

	Returns:
	identity_id: Integer ID of the identity in Steam. (int64)
	"""
	request = {
		'name': name,
		'password': password
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("CreateIdentity", request)
	return response['identity_id']

def get_identities(offset,limit):
	"""
	List identities

	Parameters:
	offset: An offset uint start the search on. (uint)
	limit: The maximum uint objects. (uint)

	Returns:
	identities: A list of identities in Steam. (Identity)
	"""
	request = {
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetIdentities", request)
	return response['identities']

def get_identities_for_workgroup(workgroup_id):
	"""
	List identities for a workgroup

	Parameters:
	workgroup_id: Integer ID of a workgroup in Steam. (int64)

	Returns:
	identities: A list of identities in Steam. (Identity)
	"""
	request = {
		'workgroup_id': workgroup_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetIdentitiesForWorkgroup", request)
	return response['identities']

def get_identities_for_role(role_id):
	"""
	List identities for a role

	Parameters:
	role_id: Integer ID of a role in Steam. (int64)

	Returns:
	identities: A list of identities in Steam. (Identity)
	"""
	request = {
		'role_id': role_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetIdentitiesForRole", request)
	return response['identities']

def get_identities_for_entity(entity_type,entity_id):
	"""
	Get a list of identities and roles with access to an entity

	Parameters:
	entity_type: An entity type ID. (int64)
	entity_id: An entity ID. (int64)

	Returns:
	users: A list of identites and roles (UserRole)
	"""
	request = {
		'entity_type': entity_type,
		'entity_id': entity_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetIdentitiesForEntity", request)
	return response['users']

def get_identity(identity_id):
	"""
	Get identity details

	Parameters:
	identity_id: Integer ID of an identity in Steam. (int64)

	Returns:
	identity: An identity in Steam. (Identity)
	"""
	request = {
		'identity_id': identity_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetIdentity", request)
	return response['identity']

def get_identity_by_name(name):
	"""
	Get identity details by name

	Parameters:
	name: An identity name. (string)

	Returns:
	identity: An identity in Steam. (Identity)
	"""
	request = {
		'name': name
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetIdentityByName", request)
	return response['identity']

def link_identity_with_workgroup(identity_id,workgroup_id):
	"""
	Link an identity with a workgroup

	Parameters:
	identity_id: Integer ID of an identity in Steam. (int64)
	workgroup_id: Integer ID of a workgroup in Steam. (int64)

	Returns:None
	"""
	request = {
		'identity_id': identity_id,
		'workgroup_id': workgroup_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("LinkIdentityWithWorkgroup", request)
	return 

def unlink_identity_from_workgroup(identity_id,workgroup_id):
	"""
	Unlink an identity from a workgroup

	Parameters:
	identity_id: Integer ID of an identity in Steam. (int64)
	workgroup_id: Integer ID of a workgroup in Steam. (int64)

	Returns:None
	"""
	request = {
		'identity_id': identity_id,
		'workgroup_id': workgroup_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("UnlinkIdentityFromWorkgroup", request)
	return 

def link_identity_with_role(identity_id,role_id):
	"""
	Link an identity with a role

	Parameters:
	identity_id: Integer ID of an identity in Steam. (int64)
	role_id: Integer ID of a role in Steam. (int64)

	Returns:None
	"""
	request = {
		'identity_id': identity_id,
		'role_id': role_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("LinkIdentityWithRole", request)
	return 

def unlink_identity_from_role(identity_id,role_id):
	"""
	Unlink an identity from a role

	Parameters:
	identity_id: Integer ID of an identity in Steam. (int64)
	role_id: Integer ID of a role in Steam. (int64)

	Returns:None
	"""
	request = {
		'identity_id': identity_id,
		'role_id': role_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("UnlinkIdentityFromRole", request)
	return 

def update_identity(identity_id,password):
	"""
	Update an identity

	Parameters:
	identity_id: Integer ID of an identity in Steam. (int64)
	password: Password for identity (string)

	Returns:None
	"""
	request = {
		'identity_id': identity_id,
		'password': password
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("UpdateIdentity", request)
	return 

def activate_identity(identity_id):
	"""
	Activate an identity

	Parameters:
	identity_id: Integer ID of an identity in Steam. (int64)

	Returns:None
	"""
	request = {
		'identity_id': identity_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("ActivateIdentity", request)
	return 

def deactivate_identity(identity_id):
	"""
	Deactivate an identity

	Parameters:
	identity_id: Integer ID of an identity in Steam. (int64)

	Returns:None
	"""
	request = {
		'identity_id': identity_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("DeactivateIdentity", request)
	return 

def share_entity(kind,workgroup_id,entity_type_id,entity_id):
	"""
	Share an entity with a workgroup

	Parameters:
	kind: Type of permission. Can be view, edit, or own. (string)
	workgroup_id: Integer ID of a workgroup in Steam. (int64)
	entity_type_id: Integer ID for the type of entity. (int64)
	entity_id: Integer ID for an entity in Steam. (int64)

	Returns:None
	"""
	request = {
		'kind': kind,
		'workgroup_id': workgroup_id,
		'entity_type_id': entity_type_id,
		'entity_id': entity_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("ShareEntity", request)
	return 

def get_privileges(entity_type_id,entity_id):
	"""
	List privileges for an entity

	Parameters:
	entity_type_id: Integer ID for the type of entity. (int64)
	entity_id: Integer ID for an entity in Steam. (int64)

	Returns:
	privileges: A list of entity privileges (EntityPrivilege)
	"""
	request = {
		'entity_type_id': entity_type_id,
		'entity_id': entity_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetPrivileges", request)
	return response['privileges']

def unshare_entity(kind,workgroup_id,entity_type_id,entity_id):
	"""
	Unshare an entity

	Parameters:
	kind: Type of permission. Can be view, edit, or own. (string)
	workgroup_id: Integer ID of a workgroup in Steam. (int64)
	entity_type_id: Integer ID for the type of entity. (int64)
	entity_id: Integer ID for an entity in Steam. (int64)

	Returns:None
	"""
	request = {
		'kind': kind,
		'workgroup_id': workgroup_id,
		'entity_type_id': entity_type_id,
		'entity_id': entity_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("UnshareEntity", request)
	return 

def get_history(entity_type_id,entity_id,offset,limit):
	"""
	List audit trail records for an entity

	Parameters:
	entity_type_id: Integer ID for the type of entity. (int64)
	entity_id: Integer ID for an entity in Steam. (int64)
	offset: An offset uint start the search on. (uint)
	limit: The maximum uint objects. (uint)

	Returns:
	history: A list of actions performed on the entity. (EntityHistory)
	"""
	request = {
		'entity_type_id': entity_type_id,
		'entity_id': entity_id,
		'offset': offset,
		'limit': limit
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetHistory", request)
	return response['history']

def create_package(project_id,name):
	"""
	Create a package for a project

	Parameters:
	project_id: No description available (int64)
	name: No description available (string)

	Returns:None
	"""
	request = {
		'project_id': project_id,
		'name': name
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("CreatePackage", request)
	return 

def get_packages(project_id):
	"""
	List packages for a project 

	Parameters:
	project_id: No description available (int64)

	Returns:
	packages: No description available (string)
	"""
	request = {
		'project_id': project_id
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetPackages", request)
	return response['packages']

def get_package_directories(project_id,package_name,relative_path):
	"""
	List directories in a project package

	Parameters:
	project_id: No description available (int64)
	package_name: No description available (string)
	relative_path: No description available (string)

	Returns:
	directories: No description available (string)
	"""
	request = {
		'project_id': project_id,
		'package_name': package_name,
		'relative_path': relative_path
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetPackageDirectories", request)
	return response['directories']

def get_package_files(project_id,package_name,relative_path):
	"""
	List files in a project package

	Parameters:
	project_id: No description available (int64)
	package_name: No description available (string)
	relative_path: No description available (string)

	Returns:
	files: No description available (string)
	"""
	request = {
		'project_id': project_id,
		'package_name': package_name,
		'relative_path': relative_path
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetPackageFiles", request)
	return response['files']

def delete_package(project_id,name):
	"""
	Delete a project package

	Parameters:
	project_id: No description available (int64)
	name: No description available (string)

	Returns:None
	"""
	request = {
		'project_id': project_id,
		'name': name
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("DeletePackage", request)
	return 

def delete_package_directory(project_id,package_name,relative_path):
	"""
	Delete a directory in a project package

	Parameters:
	project_id: No description available (int64)
	package_name: No description available (string)
	relative_path: No description available (string)

	Returns:None
	"""
	request = {
		'project_id': project_id,
		'package_name': package_name,
		'relative_path': relative_path
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("DeletePackageDirectory", request)
	return 

def delete_package_file(project_id,package_name,relative_path):
	"""
	Delete a file in a project package

	Parameters:
	project_id: No description available (int64)
	package_name: No description available (string)
	relative_path: No description available (string)

	Returns:None
	"""
	request = {
		'project_id': project_id,
		'package_name': package_name,
		'relative_path': relative_path
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("DeletePackageFile", request)
	return 

def set_attributes_for_package(project_id,package_name,attributes):
	"""
	Set attributes on a project package

	Parameters:
	project_id: No description available (int64)
	package_name: No description available (string)
	attributes: No description available (string)

	Returns:None
	"""
	request = {
		'project_id': project_id,
		'package_name': package_name,
		'attributes': attributes
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("SetAttributesForPackage", request)
	return 

def get_attributes_for_package(project_id,package_name):
	"""
	List attributes for a project package

	Parameters:
	project_id: No description available (int64)
	package_name: No description available (string)

	Returns:
	attributes: No description available (string)
	"""
	request = {
		'project_id': project_id,
		'package_name': package_name
	}
	global steamconn
	if steamconn is None:
		raise Exception('Not connected to Steam. Did you run `steam.connect()`?') 
	response = steamconn.call("GetAttributesForPackage", request)
	return response['attributes']



