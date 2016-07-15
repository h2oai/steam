# <a name="CLI Command Reference"></a>CLI Command Reference Appendix

- [`create identity`](#createidentity)
- [`create role`](#createrole)
- [`create workgroup`](#createworkgroup)
- [`deactivate identity`](#deactivateidentity)
- [`delete cluster`](#deletecluster)
- [`delete engine`](#deleteengine)
- [`delete model`](#deletemodel)
- [`delete role`](#deleterole)
- [`delete service`](#deleteservice)
- [`delete workgroup`](#deleteworkgroup)
- [`deploy engine`](#deployengine)
- [`get cluster`](#getcluster)
- [`get clusters`](#getclusters)
- [`get engine`](#getengine)
- [`get engines`](#getengines)
- [`get entities`](#getentities)
- [`get history`](#gethistory)
- [`get identities`](#getidentities)
- [`get identity`](#getidentity)
- [`get model`](#getmodel)
- [`get models`](#getmodels)
- [`get permissions`](#getpermissions)
- [`get role`](#getrole)
- [`get roles`](#getroles)
- [`get service`](#getservice)
- [`get services`](#getservices)
- [`get workgroup`](#getworkgroup)
- [`get workgroups`](#getworkgroups)
- [`import model`](#importmodel)
- [`link identity`](#linkidentity)
- [`link role`](#linkrole)
- [`login`](#login)
- [`register cluster`](#registercluster)
- [`reset`](#reset)
- [`start cluster`](#startcluster)
- [`stop cluster`](#stopcluster)
- [`stop service`](#stopservice)
- [`unlink identity`](#unlinkidentity)
- [`unregister cluster`](#unregistercluster)
- [`update role`](#updaterole)
- [`update workgroup`](#updateworkgroup)

------

<a name="createidentity"></a>
### `create identity`

**Description**

Creates a new user.

**Usage**

	./steam create identity [username] [password]


**Parameters**

- `[username]`: Enter a unique string for the new user name
- `[password]`: Enter a string for the new user's password

**Example**

The following example creates a new user with a username of "minky" and a password of "m1n5kypassword". 
 
	./steam create identity minsky m1n5kypassword
	Created user minsky ID: 2
	
------

### <a name="createrole"></a>`create role`

**Description**

Creates a new role.

**Usage**

	./steam create role [rolename] --desc="[description]"


**Parameters**

- `[rolename]`: Enter a unique string for the new role
- `--desc="[description]"`: Optionally enter a string that describes the new role

**Example**

The following example creates an engineer role. 
 
	./steam create role engineer --desc="a default engineer role"
	Created role engineer ID: 2
		
------

### <a name="createworkgroup"></a>`create workgroup`

**Description**

Creates a new workgroup.

**Usage**

	./steam create workgroup [workgroupname] --desc="[description]"

**Parameters**

- `[workgroupname]`: Enter a unique string for the new workgroup
- `--desc="[description]"`: Optionally enter a string that describes the new workgroup

**Example**

The following example creates a data preparation workgroup. 
 
	./steam create workgroup preparation --desc="data prep group"	Created workgroup preparation ID: 1
		
------

### <a name="deactivateidentity"></a>`deactivate identity`

**Description**

Deactivates an identity based on the specified username.

**Usage**

	./steam deactivate identity [username]

**Parameters**

- `[username]`: Specify the username of the identity that you want to deactivate.

**Example**

The following example deactivates user "minsky". 

	./steam deactivate minsky 

-----

### <a name="deletecluster"></a>`delete cluster`

**Description**

Deletes the specified YARN cluster from the database. Note that this command can only be used with YARN clusters (i.e., those started using [`start cluster`](#start cluster).) This command will not work with local clusters. In addition, this commmand will only work on cluster that have been stopped using [`stop cluster`](#stop cluster).

**Usage**

	./steam delete cluster [id]

**Parameters**

- `[id]`: Specify the ID of the cluster that you want to delete.

**Example**

The following example deletes cluster 2.

	./steam get engines
	NAME			ID	AGE
	automl-hdp2.2.jar	1	2016-07-14 11:48:42 -0700 PDT
	h2o-genmodel.jar	2	2016-07-14 11:49:47 -0700 PDT
	./steam delete engine 1
	Engine deleted: 1

------

### <a name="deleteengine"></a>`delete engine`

**Description**

Deletes the specified engine from the database.

**Usage**

	./steam delete engine [id]

**Parameters**

- `[id]`: Specify the ID of the engine that you want to delete.

**Example**

The following example retrieves a list of engines currently added to the database. It then specifies to delete that automodel-hdp2.2.jar engine.

	./steam get engines
	NAME			ID	AGE
	automl-hdp2.2.jar	1	2016-07-14 11:48:42 -0700 PDT
	h2o-genmodel.jar	2	2016-07-14 11:49:47 -0700 PDT
	./steam delete engine 1
	Engine deleted: 1

------

### <a name="deletemodel"></a>`delete model`

**Description**

Deletes a model from the database based on the model's ID.

**Usage**

	./steam delete model [modelId]

**Parameters**

- `[modelId]`: Specify the ID of the model that you want to delete.

**Example**

The following example deletes model 3 from the database. Note that you can use [`get models`]'(#get models) to retrieve a list of models.

	./steam delete model 3

-----

### <a name="deleterole"></a>`delete role`

**Description**

Deletes a role from the database based on its ID.

**Usage**

	./steam delete role [roleId]

**Parameters**

- `[roleId]`: Specify the ID of the role that you want to delete.

**Example**

The following example deletes role 3 from the database. Note that you can use [`get roles`]'(#get roles) to retrieve a list of roles. In the case below, this role corresponds to the default data science role. 

	./steam delete role 3

-----

### <a name="deleteservice"></a>`delete service`

**Description**

A service represents a successfully deployed model on the Steam Scoring Service. This command deletes a service from the database based on its ID. Note that you must first stop a service before it can be deleted. (See [`stop service`](#stop service).)

**Usage**

	./steam delete service [serviceId]

**Parameters**

- `[serviceId]`: Specify the ID of the service that you want to delete. Note that you can use [`get services`](#get services) to retrieve a list of services. 


**Example**

The following example stops and then deletes service 2. This service will no longer be available on the database.

	./steam stop service 2
	./steam delete service 2

-----

### <a name="deleteworkgroup"></a>`delete workgroup`

**Description**

Deletes a workgroup from the database based on its ID.

**Usage**

	./steam delete workgroup [workgroupId]

**Parameters**

- `[workgroupId]`: Specify the ID of the role that you want to delete.

**Example**

The following example deletes workgroup 3 from the database. Note that you can use [`get workgroups`]'(#get workgroups) to retrieve a list of workgroups.  

	./steam delete workgroup 3

-----

### <a name="deployengine"></a>`deploy engine` 

**Description**

Deploys an H2O engine. After an engine is successfully deployed, it can be specified when starting a cluster. (See [`start cluster`](#start cluster).) 

**Usage**

	./steam deploy engine [path/to/engine]

**Parameters**

- `[path/to/engine]`: Specify the location of the engine that you want to deploy. 

**Example**

The following specifies to deploy the H2O AutoML engine.

	./steam deploy engine ../engines/automl-hdp2.2.jar

-----

### <a name="getcluster"></a>`get cluster`

**Description** 

Retrieves detailed information for a specific cluster based on its ID.

**Usage**

	./steam get cluster [clusterId]

**Parameters**

- `[clusterId]`: Specify the ID of the cluster that you want to retrieve

**Example**

The following example retrieves information for cluster ID 1.

	./steam get cluster 1
					user
	TYPE:			external
	STATE:			healthy
	H2O VERSION:	3.8.2.8
	MEMORY:			3.56 GB
	TOTAL CPUS:		8
	ID:				1
	AGE:			2016-07-15 09:23:16 -0700 PDT

-----

### <a name="getclusters"></a>`get clusters`

**Description** 

Retrieves a list of clusters.

**Usage**

	./steam get clusters

**Parameters**

None

**Example**

The following example retrieves a list of clusters that are running H2O and are registered in Steam. (See [`register cluster`](#register cluster).)

	./steam get clusters
	NAME		ID	ADDRESS			STATE	TYPE		AGE
	user     	1	localhost:54321	started	external	2016-07-01 11:45:58 -0700 PDT

-----

### <a name="getengine"></a>`get engine`

**Description** 

Retrieves information for a specific engine based on its ID.

**Usage**

	./steam get engine [engineId]

**Parameters**

- `[engineId]`: Specify the ID of the engine that you want to retrieve

**Example**

The following example retrieves information about engine 1.

	./steam get engine 1
		h2o-genmodel.jar
	ID:		1
	AGE:	2016-07-15 09:44:10 -0700 PDT

-----

### <a name="getengines"></a>`get engines`

**Description** 

Retrieves a list of deployed engines.

**Usage**

	./steam get engines

**Parameters**

None

**Example**

The following example retrieves a list of engines that have been deployed. (Refer to [`deploy engine`](#deploy engine).)

	./steam get engines
	NAME				ID	AGE
	h2o-genmodel.jar	1	2016-07-01 13:30:50 -0700 PDT
	h2o.jar				2	2016-07-01 13:32:10 -0700 PDT

-----

### <a name="getentities"></a>`get entities`

**Description** 

Retrieves a list of supported Steam entity types.

**Usage**

	./steam get entities

**Parameters**

None

**Example**

The following example retrieves a list of the supported Steam entity types.

	./steam get entities
	NAME		ID
	role		1
	workgroup	2
	identity	3
	engine		4
	cluster		5
	project		6
	model		7
	service		8

-----

### <a name="gethistory"></a>`get history`

**Description** 

Retrieves recent activity information related to a specific user or for a specific cluster.

**Usage**

	./steam get history [identity [identityName] | cluster [clusterId]]

**Parameters**

- `identity [identityName]`: Specifies to retrieve activity information related to a specific user
- `cluster [clusterId]`: Specifies to retrieve a activity information related to a specific cluster

**Example**

The following example retrieves information for user "bob".

	./steam get history identity bob
	USER	ACTION	DESCRITPION						TIME
	1		link	{"id":"2","name":"preparation","type":"workgroup"}	2016-07-15 09:32:55 -0700 PDT
	1		link	{"id":"2","name":"engineer","type":"role"}		2016-07-15 09:32:44 -0700 PDT
	1		create	{"name":"bob"}						2016-07-15 09:32:32 -0700 PDT

-----

### <a name="getidentities"></a>`get identities`

**Description** 

Retrieves a list of users.

**Usage**

	./steam get identities

**Parameters**

None

**Example**

The following example retrieves a list of users that are available on the database.

	./steam get identities
	NAME		ID	LAST LOGIN			AGE
	bob			2	0000-12-31 16:00:00 -0800 PST	2016-07-15 09:32:32 -0700 PDT
	jim			3	0000-12-31 16:00:00 -0800 PST	2016-07-15 09:32:38 -0700 PDT
	superuser	1	0000-12-31 16:00:00 -0800 PST	2016-07-15 09:21:58 -0700 PDT

-----

### <a name="getidentity"></a>`get identity`

**Description** 

Retrieve information about a specific user.

**Usage**

	./steam get identity [identityId]

**Parameters**

- `[identityId]`: Specify the ID of the user you want to retrieve

**Example**

The following example retrieves information about user Jim.

	./steam get identity jim
				jim
	STATUS:		Active
	LAST LOGIN:	0000-12-31 16:00:00 -0800 PST
	ID:		3
	AGE:		2016-07-15 09:32:38 -0700 PDT

	WORKGROUP	DESCRIPTION
	production	production group

	ROLE		DESCRIPTION
	datascience	a default data scientist role

	PERMISSIONS
	Manage models
	View clusters
	Manage projects

-----

### <a name="getmodel"></a>`get model`

**Description** 

Retrieves detailed information for a specific model.

**Usage**

	./steam get model [modelId]

**Parameters**

- `[modelId]`: Specify the ID of the model that you want to retrieve

**Example**

The following example retrieves information for model 2.

	./steam get model 2
	
-----

### <a name="getmodels"></a>`get models`

**Description** 

Retrieves a list of models.

**Usage**

	./steam get models

**Parameters**

None

**Example**

The following example retrieves a list of models that are available on the database.

	./steam get models
	
-----

### <a name="getpermissions"></a>`get permissions`

**Description** 

Retrieves a list of permissions available in Steam along with the corresponding code. These permissions are currently hard coded into Steam. 

**Usage**

	./steam get permissions

**Parameters**

None

**Example**

The following example retrieves a list of Steam permissions.

	./steam get permissions
	ID	DESCRIPTION		CODE
	9	Manage clusters		ManageCluster
	7	Manage engines		ManageEngine
	5	Manage identities	ManageIdentity
	13	Manage models		ManageModel
	11	Manage projects		ManageProject
	1	Manage roles		ManageRole
	15	Manage services		ManageService
	3	Manage workgroups	ManageWorkgroup
	10	View clusters		ViewCluster
	8	View engines		ViewEngine
	6	View identities		ViewIdentity
	14	View models		ViewModel
	12	View projects		ViewProject
	2	View roles		ViewRole
	16	View services		ViewService
	4	View workgroups		ViewWorkgroup

-----

### <a name="getrole"></a>`get role`

**Description** 

Retrieves detailed information for a specific role based on its name.

**Usage**

	./steam get role [roleName]

**Parameters**

- `[roleName]`: Specify the name of the role that you want to retrieve

**Example**

The following example retrieves information about the datascience role.

	./steam get role datascience
				datascience
	DESCRIPTION:	a default data scientist role
	ID:		3
	AGE:	2016-07-15 09:32:10 -0700 PDT

	IDENTITES: 1
	NAME	STATUS	LAST LOGIN
	jim		Active	0000-12-31 16:00:00 -0800 PST

	PERMISSIONS
	Manage models
	Manage projects
	View clusters

-----

### <a name="getroles"></a>`get roles`

**Description** 

Retrieves a list of roles.

**Usage**

	./steam get roles

**Parameters**

None

**Example**

The following example retrieves a list of roles that are available on the database.

	./steam get roles
	NAME		ID	DESCRIPTION			AGE
	Superuser	1	Superuser			2016-07-14 09:25:30 -0700 PDT
	datascience	3	a default data scientist role	2016-07-14 15:39:03 -0700 PDT
	engineer	2	a default engineer role		2016-07-14 15:38:10 -0700 PDT

-----

### <a name="getservice"></a>`get service`

**Description** 

A service represents a successfully deployed model on the Steam Scoring Service. This command retrieves detailed information about a specific service based on its ID.

**Usage**

	./steam get service [serviceId]

**Parameters**

- `[serviceId]`: Specify the ID of the service that you want to retrieve

**Example**

The following example retrieve information about service 2.

	./steam get service 2

-----

### <a name="getservices"></a>`get services`

**Description** 

A service represents a successfully deployed model on the Steam Scoring Service. This command retrieves a list of services available on the database.

**Usage**

	./steam get services

**Parameters**

None

**Example**

The following example retrieves a list of services that are available on the database.

	./steam get services

-----

### <a name="getworkgroup"></a>`get workgroup`

**Description** 

Retrieves information for a specific workgroup based on its name.

**Usage**

	./steam get workgroup [workgroupName]

**Parameters**

- `[workgroupName]`: Specify the name of the workgroup that you want to retrieve

**Example**

The following example retrieves information about the production workgroup

	./steam get workgroup production
					production
	DESCRIPTION:	production group
	ID:		3
	AGE:	2016-07-15 09:32:27 -0700 PDT

	IDENTITIES: 1
	NAME	STATUS	LAST LOGIN
	jim		Active	0000-12-31 16:00:00 -0800 PST
	
-----

### <a name="getworkgroups"></a>`get workgroups`

**Description** 

Retrieves a list of workgroups currently available on the database.

**Usage**

	./steam get workgroups --identity=[identityName]

**Parameters**

- `--identity=[identityName]`: Optionally specify to view all workgroups associated with a specific user name

**Example**

The following example retrieves a list of workgroups that are available on the database.

	./steam get workgroups
	NAME		ID	DESCRIPTION		AGE
	preparation	2	data prep group		2016-07-15 09:32:21 -0700 PDT
	production	3	production group	2016-07-15 09:32:27 -0700 PDT

-----

### <a name="importmodel"></a>`import model`

**Description** 

Imports a model from H2O based on its ID. 

**Usage**

	./steam import model [clusterId] [modelName]

**Parameters**

- `[clusterId`]: Specify the H2O cluster that contains the model you want to import
- `[modelName]`: Specify the name of the that you want to import into steam. 

**Example**

The following example specifies to import the GBM_model_python_1468599779202_1 model from Cluster 1.

	./steam import model 1 GBM_model_python_1468599779202_1

-----

### <a name="linkidentity"></a>`link identity`

**Description** 

Links a user to a specific role or workgroup. 

**Usage**

	./steam link identity [identityName] [role [roleId] | workgroup [workgroupId]]

**Parameters**

- `[identityName]`: Specify the user that will be linked to a role or workgroup.
- `role [roleId]`: Specify the role that the user will be linked to.
- `workgroup [workgroupId]`: Specify the workgroup that the the user will be linked to.

**Example**

The following example links user Jim to datascience role and then to the production workgroup.

	./steam link identity jim role datascience
	./steam link identity jim workgroup production

-----

### <a name="linkrole"></a>`link role`

**Description** 

Links a role to a certain set of permissions. 

**Usage**

	./steam link role [roleId] [permissionId1 permissionId2 ...]

**Parameters**

- `[roleId]`: Specify the role that the user will be linked to.
- `[permissionId]`: Specify a single permission or a list of permissions to assign to this role.  

**Example**

The following example links the datascience role to the ManageProject, ManageModel, and ViewCluster permissions.

	./steam link role datascience ManageProject ManageModel ViewCluster 

-----

### <a name="login"></a>`login`

**Description**

Logs a user in to Steam

**Usage**

	./steam login [address:port] --username=[userName] --password=[password]
	
**Parameters**

- `[address:port]`: Specify the address and port of the Steam server.
- `--username=[userName]`: Specify the username.
- `--password=[password]`: Specify the user's password.

**Example**

The following example logs user Bob into a Steam instance running on localhost:9000.

	./steam login localhost:9000 --username=bob --password=bobSpassword
	Login credentials saved for server localhost:9000

-----

### <a name="registercluster"></a>`register cluster`

**Description**

Registers a cluster that is currently running H2O (typically a local cluster). Once registered, the cluster can be used to perform machine learning tasks through Python, R, and Flow. The cluster will also be visible in the Steam web UI. 

Note that clusters that are started using this command can be stopped from within the web UI or using [`unregister cluster`](#unregister cluster). You will receive an error if you attemt to stop registered clusters using the `stop cluster` command. 

**Usage**

	./steam register cluster [address]

**Parameters**

- `[address]`: Specify the IP address and port of the cluster that you want to register.

**Example**

The following example registers Steam on localhost:54323. Note that this will only be successful if H2O is already running on this cluster. 

	./steam register cluster localhost:54323
	Successfully connected to cluster 2 at address localhost:54323

-----

### <a name="reset"></a>`reset`

**Description**

Resets the current Steam cluster instance. This removes the current authentication from Steam. You will have to re-authenticate in order to continue to use Steam. 

**Usage**

	./steam reset

**Parameters**

None

**Examples**

The following example resets the current Steam instance.

	./steam reset
	Configuration reset successfully. Use 'steam login <server-address>' to re-authenticate to Steam

-----

### <a name="startcluster"></a>`start cluster`

**Description**

After you have deployed engine, you can use this command to start a new cluster through YARN using a specified engine. Note that this command is only valid when starting Steam on a YARN cluster. To start Steam on a local cluster, use [`register cluster`](#register cluster) instead.

**Usage**

	./steam start cluster [id] [engineId] --size=[numNodes] --memory=[string]

**Parameters**

- `[id]`: Enter an ID for this new cluster.
- `[engineId]`: Specify the ID of the engine that this cluster will use. If necessary, use [`get engines`](#get engines) to retrieve a list of all available engines.
- `--size=[numNodes]`: Specify an integer for the number of nodes in this cluster.
- `--memory=[string]`: Enter a string specifying the amount of memory available to Steam in each node (for example, "1024m", "2g", etc.)

**Example**

The following example retrieves a list of engines, then starts a cluster through YARN using one from the list. The cluster is configured with 2 nodes that are 2 gigabytes each. 

	./steam get engines
	NAME				ID	AGE
	h2o-genmodel.jar	1	2016-07-01 13:30:50 -0700 PDT
	h2o.jar			2	2016-07-01 13:32:10 -0700 PDT
	./steam start cluster 9 1 --size=2 --memory=2g
	
-----

### <a name="stopcluster"></a>`stop cluster`

**Description**

Stops a YARN cluster that was started through the CLI or web UI. (See [`start cluster`](#start cluster).) Note that you will receive an error if you attempt to stop a cluster that was started using `register cluster`. 

**Usage**

	./steam stop cluster [id] 

**Parameters**

- `[id]`: Specify the ID of the cluster that you want to stop. If necessary, use [`get clusters`](#get clusters) to retrieve a list of clusters. 

**Example**

The following example stops a cluster that has an ID of 9.

	./steam stop cluster 9

-----

### <a name="stopservice"></a>`stop service`

**Description**

A service represents a successfully deployed model on the Steam Scoring Service. Use this command to stop a service. 

**Usage**

	./steam stop service [serviceId] 

**Parameters**

- `[serviceId]`: Specify the ID of the scoring service that you want to stop. If necessary, use [`get services`](#get clusters) to retrieve a list of running services. 

**Example**

The following example stops a service that has an ID of 2.

	./steam stop service 2

-----

### <a name="unlinkidentity"></a>`unlink identity`

**Description** 

Removes a user's permissions from a specific role or workgroup.

**Usage**

	./steam unlink identity [identityName] [role [roleId] | workgroup [workgroupId]]

**Parameters**

- `[identityName]`: Specify the user that will be unlinked from a role or workgroup
- `role [roleId]`: Specify the role that the user will be unlinked from
- `workgroup [workgroupId]`: Specify the workgroup that the the user will be unlinked from

**Example**

The following example removes user Jim from the datascience role and then from the production workgroup.

	./steam unlink identity jim role datascience
	./steam unlink identity jim workgroup production

-----

### <a name="unregistercluster"></a>`unregister cluster`

**Description**

Stops a cluster that was registered through the CLI or the web UI. (See [`register cluster`](#register cluster).) Note that this does not delete the cluster. Also note that you will receive an error if you attempt to unregister a cluster that was started using `start cluster`. 

**Usage**

	./steam unregister cluster [id] 

**Parameters**

- `[id]`: Specify the ID of the cluster that you want to stop. If necessary, use [`get clusters`](#get clusters) to retrieve a list of clusters. 

**Example**

The following example stops a cluster that has an ID of 9. 

	./steam unregister cluster 2
	Successfully unregisted cluster %d 2

-----

### <a name="updatedrole"></a>`update role`

**Description**

Edits the description and/or name of an existing role. When a role is edited, the edit will automatically propagate to all identities that are associated with this role.

**Usage**

	./steam update role [rolename] --desc="[description]" --name="[newRoleName]

**Parameters**

- `[rolename]`: Enter the role name that you want to edit
- `desc="[description]"`: Optionally enter a string that describes the new role
- `name="[newRoleName]"`: Enter a unique string for the new role name

**Example**

The following example changes the name of the engineer role to be "science engineer".  
 
	./steam update role engineer --desc="A better engineer" --name="science engineer"
	Successfully updated role: engineer
		
------

### <a name="createworkgroup"></a>`create workgroup`

**Description**

Edits the description and/or name of an existing workgroup. When a workgroup is edited, the edit will automatically propagate to all identities that are associated with this workgroup.

**Usage**

	./steam update workgroup [workgroupname] --desc="[description]" --name="[newWorkgroupName]


**Parameters**

- `[workgroup]`: Enter the workgroup name that you want to edit
- `desc="[description]"`: Optionally enter a string that describes the new workgroup
- `name="[newWorkgroupName]"`: Enter a unique string for the new workgroup name

**Example**

The following example changes the name of the production workgroup to be "deploy". 
 
	./steam update workgroup production --desc="A deploy workgroup" --name="deploy"
	Successfully updated workgroup: production

		