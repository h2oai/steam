import backend

class SteamClient(object):
	def __init__(self, httpconn):
		self.__conn = backend.SteamConnection(httpconn)
	
	def start_cluster(self, name=None, num_nodes=0, mem_per_node=None, h2o_version=None):
			
		engine = self.__conn.get_engine_by_version(h2o_version)
	
		proxconf = self.__conn.get_config()
		cluster = self.__conn.get_cluster(clid)
		cport = proxconf['cluster_proxy_address'].split(':', 1)[1]
		conf = {'https':True, 'verify_ssl_certificates':self.__conn.verify_ssl, \
			'port': int(cport), \
			'context_path':'%s_%s' % (proxconf['username'], cluster['name']), \
			'cookies':["%s=%s" % (cluster['name'], cluster['token'])], \
			'ip':self.__conn.host}
		return conf	
	
	def stop_cluster(self, config):
		name = config['context_path'].split('_')[1]
		clid = None
		clusts = self.__conn.get_clusters(0, 1000)
		for c in clusts:
			if c['name'] == name:
				clid = c['id']
				break
		
		if clid is None:
			raise LookupError("Failed to locate referenced cluster")

		self.__conn.stop_cluster_on_yarn(clid, None)


	def upload_keytab(self, path):
		self.__conn.upload(target="type=keytab&principal=user", path=path)

	

#def create_login_file(path, username, password, login_file_pass):
	

def login(ip, port=9000, username=None, password=None, login_file=None, login_file_pass=None, verify_ssl=True):
	if password is not None and username is not None:
		steamconn = SteamClient(backend.HTTPSConnection(ip, port, username, password, verify_ssl))
	
	return steamconn



#	def get_cluster_connection(self, clid):
#		proxconf = self.get_config()
#		cluster = self.get_cluster(clid)
#		cport = proxconf['cluster_proxy_address'].split(':', 1)[1]
#		conf = {'https':True, 'verify_ssl_certificates':self.connection.verify_ssl, \
#			'port': int(cport), \
#			'context_path':'%s_%s' % (proxconf['username'], cluster['name']), \
#			'cookies':["%s=%s" % (cluster['name'], cluster['token'])], \
#			'ip':self.connection.host}
#		return conf	
	
