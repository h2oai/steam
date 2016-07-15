# Steam Standard Standalone Demo


This demo describes how to use Steam without the need for a local running instance of YARN. This demo will walk through the following procedures:

- Installing and starting Steam, the Compilation Service, and H2O
- Adding Roles, Workgroups, and Users to the database
- Building a simple model in Python (Optional for users who don't have an existing demo)
- Deploying the model using Steam

During this demo, three terminal windows will remain open for the Steam, Scoring, and H2O services. A fourth terminal window will be used to run H2O commands in the Python or R example. 

Finally, these steps were created using H2O version 3.8.2.8, and that version resides in a Downloads folder. Wherever used, this version number and path should be adjusted to match your version and path.

## Steam Overview

Steam is an “instant on” platform that streamlines the entire process of building and deploying applications. It is the industry’s first data science hub that lets data scientists and developers collaboratively build, deploy, and refine predictive applications across large scale data sets. Data scientists can publish Python and R code as REST APIs and easily integrate with production applications.

## Requirements

- Web browser with an Internet connection
- s3cmd
	- available from <a href="http://s3tools.org/s3cmd" target="_blank">s3tools.org</a>
	- Requires AWS credentials. Contact your internal support for access.
	- Requires <a href="http://brew.sh/" target="_blank">Homebrew</a> to install to run the s3 commands:
	
	```
		brew install s3cmd
		s3cmd --configure
	```
- JDK 1.7 or greater
- PostgreSQL
	- available from <a href="https://www.postgresql.org/" target="_blank">PostgreSQL.org</a>
- Steam binary for Linux or OS X 
	- `s3cmd get s3://steam-release/steamY-master-linux-amd64.tar.gz`
	- `s3cmd get s3://steam-release/steamY-master-darwin-amd64.tar.gz`
- H2O jar file
	- available from the <a href="http://www.h2o.ai/download/h2o/choose" target="_blank">H2O Download</a> page

###Optional

The following are required if you use a Python or R demo.

**Python**

- A dataset that will be used to generate a model. This demo uses the well-known iris.csv dataset with headers (available online), and the dataset is saved onto the desktop. 
- Python 2.7

**R**

- A dataset that will be used to generate a model. 
- Comprehensive R Archive Network (R). Available from <a href="https://cran.r-project.org/mirrors.html", target="_blank">https://cran.r-project.org/mirrors.html</a>. 

## User Management Overview

Before using Steam, it is important to understand User Management within your YARN environment. In Steam, User Management is supported in a postresql database. The User Management functions in Steam determine the level of access that users have for Steam features. The Steam database supports setup via CLI commands. Refer to the [CLI Command Reference Appendix](#CLI Command Reference) at the end of this document for a list of all available CLI commands. 

For more information on Steam User Management, refer to the following sections. 

- [Terms](#terms)
- [Privileges/Access Control](#privileges)
- [Authorization](#authorization)
- [User Management Workflow](#user management workflow)
- [CLI Command Reference](#CLI Command Reference)

<a name="terms"></a>
### Terms

The following lists common terms used when describing Steam User Management.  

- **Entities** represent *objects* in Steam. Examples of entities include Roles, Workgroups, Identities, Clusters, Projects, Models, and Services (engines). 

- **Identities** represent *users* in Steam. Users sign in using an Identity, and then perform operations in Steam.

- **Permissions** determine what operations you can perform. Examples of permissions include *Manage Clusters*, *View Clusters*, *Manage Models*, *View Models*, and so on.

- **Privileges** determine the entities that you can perform operations on (i.e., data / access control).


<a name="privileges"></a>
### Privileges/Access Control

Privileges are uniquely identified by the entity in question and the kind of privilege you have on the entity.

The following privileges are available on an entity:

- **Own** privileges allow you to share, view, edit, and delete entities.

- **Edit** privileges allow you to view and edit entities, but not share or delete them.

- **View** privileges allow you to view entities, but not share, edit, or delete them.

When you create an entity, you immediately *Own* it. You can then share this entity with others and award them either *Edit* or *View* privileges. Entities are allowed to have more than one owner, so you can also add additional owners to entities. 

The following table lists the kind of privileges you need in order to perform specific operations on entities:


        Entity               Own  Edit View
        -----------------------------------
        Role
          Read               x    x    x
          Update             x    x
          Assign Permission  x    x
          Delete             x
          Share              x
          
        Workgroup
          Read               x    x    x
          Update             x    x
          Delete             x
          Share              x
        
        Identity
          Read               x    x    x
          Assign Role        x    x
          Assign Workgroup   x    x
          Update             x    x
          Delete             x
          Share              x
        
        Cluster
          Read               x    x    x
          Start/Stop         x
        
        Project
          Read               x    x    x
          Assign Model       x    x
          Update             x    x
          Delete             x
          Share              x
        
        Engine, Model
          Read               x    x    x
          Update             x    x
          Delete             x
          Share              x

        

<a name="authorization"></a>
### Authorization

Permissions and privileges are set up using Roles and Workgroups, respectively.

- Identities cannot be linked directly to permissions. For that, you'll need Roles.

- Identities cannot be linked directly to privileges on entities. For that, you'll need Workgroups, i.e. when you share entities with others, you would be sharing those entities with workgroups, not individuals.

#### Roles
A **Role** is a named set of permissions. Roles allow you define a cohesive set of permissions into operational roles and then have multiple identities *play* those roles, regardless of access control.
For example:

- a *Data Scientist* role can be composed of the permissions *View Clusters*, *Manage Models*, *View Models*.
- an *Operations* role can be composed of the permissions *View Models*, *View Services*, *Manage Services*,
- a *Manager* role can be composed of the permissions *Manage Roles*, *View Roles*, *Manage Workgroups*, *View Workgroups*

#### Workgroups
A **Workgroup** is a named set of identities. Workgroups allow you to form collections of identities for access control purposes. For example, a *Demand Forecasting* workgroup can be composed of all the users working on demand forecasting, regardless of their role. This workgroup can be then used to control access to all the clusters, projects, models and services that are used for demand forecasting. 

#### Next Steps

Now that you understand User Management, you can begin building and then setting up the H2O Scoring Service and Steam.


## Starting Steam

This section describes how to set up and start Steam and start the Steam CLI for user management. Five terminal windows will be open the first time you run this setup; four terminal windows will be open for subsequent logins.

1. Open a terminal window and start postgresql. This should be started from the folder where posgresql was installed.

		postgres -D /usr/local/var/postgres

2. Open a second terminal window and download the steamY binary. Note that the command below downloads the OS X binary. Replace `darwin` with `linux` in the steps that follow to build on Linux.

	`user$ s3cmd get s3://steam-release/steamY-master-darwin-amd64.tar.gz`

3. Untar the steamY binary.

	`user$ tar xvf steamY-master-darwin-amd64.tar.gz`

4. Open a third terminal window to create a new user for the Steam database and then create the database. The commands below only need to be performed once. The example below creates a steam **superuser** with a password ``superuser`` before creating the Steam database. Be sure to provide a secure password, and be sure to remember the password that you enter. This will be required each time you log in to Steam. 

		createuser -P steam 
		Enter password for new role: superuser
		Enter it again: superuser
		# Change directories to the Steam /var/master/scripts folder.
		cd steam-master-darwin-amd64/var/master/scripts
		./create-database.sh

5. Change directories to your Steam directory, and start the Jetty server.

		user$ cd steam-master-darwin-amd64
		user$ java -jar var/master/assets/jetty-runner.jar var/master/assets/ROOT.war
		
	>***Note***: The Jetty server defaults to port 8080. You can optionally provide a `--port` value for **jetty-runner.jar**.
		
6. Open a fourth terminal window. From within the **steam-maseter-darwin-amd64** folder, start the Steam compilation and scoring service using the password that you provided in Step 2. This starts Steam on localhost:9000.

		./steam serve master --superuser-name=superuser --superuser-password=superuser

	>***Note***: This starts the Steam web service on `localhost:9000`, the compilation service on `localhost:8080` (same as the Jetty server), and the scoring service on `localhost`. You can change these using `--compilation-service-address=<ip_address:port>` and `--scoring-service-address=<ip_address>`. Use `./steam help serve master` or `./steam serve master -h` to view additional options.

7. <a name="step7"></a>Open a fifth terminal window. From within the Steam folder, log in to the maching running Steam (localhost:9000). Use the password that you provided in Step 2.

		./steam login localhost:9000 --username=superuser --password=st3amUser


8. Run the following to verify that the CLI is working correctly.

		./steam help
		
At this point, you can open a browser and navigate to localhost:9000. Login credentials will automatically be supplied. 

The next section describes how to add additional users to the Steam database.

<a name="user management workflow"></a>
### User Management Workflow

The steps below provide a common workflow to follow when creating users. This workflow is followed in the example that follows.

1. Define roles based on operational needs.
2. Define workgroups based on data / access control needs.
3. Then add a new user:

 -	Create the user's identity.
 - Associate the user with one or more roles.
 - Optionally, associate the user with one or more workgroups. 

#### Example

The following example creates sample roles, workgroups, and users using the CLI. Refer to the [CLI Command Reference](#CLI Command Reference) section for information about all of the commands available in the CLI. These commands are run from the terminal window used to log in to Steam ([Step 7](#step7) above).

		# Create engineer role and link that role to permissions
		./steam create role engineer --desc="a default engineer role"
		./steam link role engineer ViewModel ViewProject ViewWorkgroup
		
		# Create data scientist role and link that role to permissions
		./steam create role datascience --desc="a default data scientist role"
		./steam link role datascience ManageProject ManageModel ViewCluster 
		
		# Create preparation and production workgroups
		./steam create workgroup preparation --desc="data prep group"
		./steam create workgroup production --desc="production group"
		
		# Create two users - Bob and Jim
		./steam create identity bob bobSpassword
		./steam create identity jim j1mSpassword
		
		# Link Bob to engineer role; link Jim to datascience role
		./steam link identity bob role engineer
		./steam link identity jim role datascience
		
		# Link Bob to preparation workgroup; link Jim to production workgroup
		./steam link identity bob workgroup preparation
		./steam link identity jim workgroup production


### Stopping the Steam Database

Use Ctrl+C in each of the Steam, Compilation Service, and postgres terminal windows to stop the services end your session. 


## Using H2O with Steam

Now that Steam is running, and users are set up, this section describes how to use H2O with Steam.

1. Open another terminal window. Navigate to the folder with your H2O jar file and start H2O. This will create a one-node cluster on your local machine on port 54321.

		user$ cd ~/Downloads/h2o-3.8.2.8
		user$ java -jar h2o.jar 
		
2. Point your browser to the Steam URL, for example, http://localhost:9000/.
 
3. In the left pane, select the **Clusters** tab (selected by default), then click the **Connect To Cluster** button to setup Steam with H2O. Specify the IP address and port of the cluster currently running H2O (for example, localhost:54321), then click **Register Cluster**. 

	![](images/register_cluster.png)

You are now ready to build a model on this cluster in Python. 

>***Note***: After you connect to a cluster, click on **Cluster Details**, select your cluster, then click the **Address** link on that page to launch H2O Flow. 

## Building a Model in Python (Optional)

>**Notes**: This section can be skipped if you already have demo steps that you use in R, Python, or Flow. If you use another demo, be sure that you initialize H2O on your local cluster so that the data will be available in Steam. 
 
 >	Additional demos for Python are available <a href="https://github.com/h2oai/h2o-3/tree/master/h2o-py/demos" target="_blank">here</a>.
 
 > Demos for R are available <a href="https://github.com/h2oai/h2o-3/tree/master/h2o-r/demos" target="_blank">here</a>. 
 
 > A demo of Flow can be viewed <a href="https://www.youtube.com/watch?feature=player_embedded&v=wzeuFfbW7WE" target="_blank">here</a>. 

The steps below show how to build model using the Iris dataset and the GBM algorithm. The steps will be run using H2O in Python. Once created, the model can be deployed in Steam. 

1. Open a terminal window. Change directories to the H2O folder, and start Python. Import the modules that will be used for this demo. 

		$ cd ~/Downloads/h2o-3.8.2.8
		$ python
		>>> import h2o
		>>> from h2o.estimators.gbm import H2OGradientBoostingEstimator

2. Initialize H2O using localhost and port 54321. (Note that if started Steam on a different machine, then replace `localhost` with the IP address of that machine.)

		>>> h2o.init(ip="localhost", port=54321)
		------------------------------  -------------------------------------
		H2O cluster uptime:             2 minutes 37 seconds 168 milliseconds
		H2O cluster version:            3.8.2.8
		H2O cluster name:               user
		H2O cluster total nodes:        1
		H2O cluster total free memory:  3.35 GB
		H2O cluster total cores:        8
		H2O cluster allowed cores:      8
		H2O cluster healthy:            True
		H2O Connection ip:              127.0.0.1
		H2O Connection port:            54321
		H2O Connection proxy:
		Python Version:                 2.7.9
		------------------------------  -------------------------------------

3. Upload the Iris dataset. Note that in this example, Python is running from the Downloads folder, and the Iris dataset is on the Desktop:

		>>> df=h2o.upload_file("../../Desktop/iris.csv")

4. Specify the configuration options to use when building a GBM model.

		>>> gbm_regressor = H2OGradientBoostingEstimator(distribution="gaussian", ntrees=10, max_depth=3, min_rows=2, learn_rate="0.2")

5. Train the model using the Iris dataset (`df` object) and the GBM configuration options. 

		>>> gbm_regressor.train(x=range(1, df.ncol), y=0, training_frame=df)

6. Optionally view the model details.

		>>> gbm_regressor

Once created, the model will be visible in the Steam UI. 

## Deploying a Model in Steam

1. In the Steam UI, Select **Cluster** > **Models**. Select a model from your demo, and then click **Import Model to Steam**. This pulls the model into Steam. Once imported, the model can then be deployed to the scoring service.

  ![](images/import_model.png)

2. Select the **Models** tab in the left pane. You should see the model that you just imported. Select this model, and then click **Deploy This Model** to create scoring services for the model.
 

	![](images/deploy_model.png)

3. Specify the port number for the scoring service (defaults to 8000), then click **Deploy**.

4. Select the **Services** tab in the left pane.

5. Select a service (in this case, the model you just deployed), and click the link in the **Endpoint** field to reach the scoring service.

	![](images/select_service.png)

6. Make predictions using one of the following methods:
    
    - Specify input values based on column data from the original dataset
    
     OR
    
    - Enter a query string using the format `field1=value1&field2=value2` (for example, `sepal_width=3&petal_len=5`)

 Use the **Clear** button to clear all entries and begin a new prediction.
     
 You can view additional statistics about the scoring service by clicking the **More Stats** button.
 
## What's Next?
 
Now that you have completed your first demo, you are ready to begin creating models using your own data. Additional users can then be give access to this data based on the user's role and workgroup.  

---------

<a name="CLI Command Reference"></a>
# CLI Command Reference Appendix

- [`create identity`](#create identity)
- [`create role`](#create role)
- [`create workgroup`](#create workgroup)
- [`deactivate identity`](#deactivate identity)
- [`delete cluster`](#delete cluster)
- [`delete engine`](#delete engine)
- [`delete model`](#delete model)
- [`delete role`](#delete role)
- [`delete service`](#delete service)
- [`delete workgroup`](#delete workgroup)
- [`deploy engine`](#deploy engine)
- [`get clusters`](#get clusters)
- [`get engines`](#get engines)
- [`get roles`](#get roles)
- [`register cluster`](#register cluster)
- [`start cluster`](#start cluster)
- [`stop cluster`](#stop cluster)
- [`unregister cluster`](#unregister cluster)


------

#### <a name="create identity"></a>`create identity`

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

#### <a name="create role"></a>`create role`

**Description**

Creates a new role.

**Usage**

	./steam create role [rolename] desc="[description]"


**Parameters**

- `[rolename]`: Enter a unique string for the new role
- `desc="[description]"`: Optionally enter a string that describes the new role

**Example**

The following example creates an engineer role. 
 
	./steam create role engineer --desc="a default engineer role"
	Created role engineer ID: 2
		
------

#### <a name="create workgroup"></a>`create workgroup`

**Description**

Creates a new workgroup.

**Usage**

	./steam create workgroup [workgroupname] desc="[description]"


**Parameters**

- `[workgroupname]`: Enter a unique string for the new workgroup
- `desc="[description]"`: Optionally enter a string that describes the new workgroup

**Example**

The following example creates a data preparation workgroup. 
 
	./steam create workgroup preparation --desc="data prep group"	
	Created workgrou preparation ID: 1
		
------

#### <a name="deactivate identity"></a>`deactivate identity`

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

#### <a name="delete cluster"></a>`delete cluster`

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

#### <a name="delete engine"></a>`delete engine`

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

-----
#### <a name="delete model"></a>`delete model`

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

#### <a name="delete role"></a>`delete role`

**Description**

Deletes a role from the database based on its ID.

**Usage**

	./steam delete role [roleId]

**Parameters**

- `[roleId]`: Specify the ID of the role that you want to delete.

**Example**

The following example deletes role 3 from the database. Note that you can use [`get roles`]'(#get roles) to retrieve a list of roles. In the case below, this role corresponds to the default data science role. 

	./steam delete roles 3

-----

#### <a name="delete service"></a>`delete service`

**Description**

Deletes the specified service from the database.

**Usage**


**Parameters**


**Example**

-----

#### <a name="delete workgroup"></a>`delete workgroup`

**Description**

Deletes the specified workgroup from the database.

**Usage**


**Parameters**


**Example**

-----

#### <a name="deploy engine"></a>`deploy engine` 

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

#### <a name="get clusters"></a>`get clusters`

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

#### <a name="get engines"></a>`get engines`

**Description** 

Retrieves a list of deployed engines.

**Usage**

	./steam get engines

**Parameters**

None

**Example**

The following example retrieves a list of engines that have been deployed. (Refer to [`deploy engine`](#deploy engine).)

	./steam get engines
	NAME			ID	AGE
	h2o-genmodel.jar	1	2016-07-01 13:30:50 -0700 PDT
	h2o.jar			2	2016-07-01 13:32:10 -0700 PDT

-----

#### <a name="get roles"></a>`get roles`

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


#### <a name="register cluster"></a>`register cluster`

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

#### <a name="start cluster"></a>`start cluster`

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

#### <a name="stop cluster"></a>`stop cluster`

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

#### <a name="unregister cluster"></a>`unregister cluster`

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
