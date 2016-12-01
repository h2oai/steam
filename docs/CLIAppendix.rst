Appendix B: CLI Command Reference 
=================================

This document describes the CLI commands available in Steam. In addition to this reference, you can view information about each command when you're in the CLI by typing ``./steam help``. 

**Note**: To access the Steam CLI, open a terminal window and log in to Steam as the superuser:

::

    ./steam login <yarn_edge_node>:<port> --username=superuser --password=superuser

----- 

``build model``
~~~~~~~~~~~~~~~

**Description**

Builds a model using either a specified algorithm or through AutoML.

**Usage**

::

    ./steam build model --cluster-id="[cluster]" --dataset-id="[dataset]" --algorithm="[algorithm]"

    ./steam build model --auto --cluster-id="[cluster]" --dataset-id="[dataset]" --target-name="[model_name]" --max-run-time="[seconds]"

**Parameters**

-  ``--cluster-id="[cluster]"``: Specify the ID of the cluster that
   contains the dataset and will contain this model
-  ``--dataset-id="[dataset]"``: Specify the ID of the dataset to use to
   build the model
-  ``--algorithm="[algorithm]"``: Specify the algorithm to use for the
   model. This option cannot be used with ``--auto``. Options include:

   -  ``gbm``: Build a Gradient Boosting Machine model
   -  ``glm``: Build a Generalized Linear model
   -  ``glrm``: Build a Generalized Low Rank model
   -  ``rf``: Build a Random Forest model
   -  ``svm``: Build a Support Vector Machine model
   -  ``dl``: Build a Deep Learning model
   -  ``nb``: Build a Naive Bayes model

-  ``--auto``: Specify to use AutoML to build the model
-  ``--target-name=[model_name]``: Specify a name for the AutoML model
-  ``--max_run_time``: When building an AutoML model, speicfy the
   maximum runtime in seconds to allow for the model to build.

**Example**

The following example builds a gbm model from the airlines dataset. This
dataset was added using `create dataset`_ and has an ID of 1.

::

    ./steam build model --cluster-id="1" --dataset-id="1" --algorithm="gbm"

--------------

``create dataset``
~~~~~~~~~~~~~~~~~~

**Description**

Creates a dataset from an available source file. Once created, the
dataset can be used to build a model.

**Usage**

::

    ./steam create dataset --cluster-id=[cluster] --datasource-id=[source] --name="[datasetname]" --description="[description]" --response-column-name="[column]"

**Parameters**

-  ``--cluster-id=[cluster]``: Specify the ID of the cluster running H2O
   that will contain this dataset
-  ``--datasource-id=[source]``: Specify the ID of the datasource that
   will be used to create this dataset
-  ``--name="[datasetname]"``: Optionally enter a name for this dataset
-  ``--description="[description]"``: Optionally provide a description
   for this dataset
-  ``--response-column-name="[column"]``: Specify the column that will
   be used when making predictions

**Example**

The following example creates a dataset from a source file that was added using `create datasource`_. In this example, Steam will generate a name for the dataset. Note that H2O must be running on the specified cluster.

::

    ./steam create dataset --cluster-id=1 --datasource-id=1 --response-column-name="Origin"
    DatasetId:  1

--------------

``create datasource``
~~~~~~~~~~~~~~~~~~~~~

**Description**

Adds a datasource to the Steam database. Once added, this source file
can be used to create a dataset.

**Usage**

::

    ./steam create datasource --name="[sourcename]" --description="[description]" --path="[path]" --project-id=[id]

**Parameters**

-  ``--name="[datasetname]"``: Optionally enter a name for this dataset
-  ``--description="[description]"``: Optionally provide a description
   for this dataset
-  ``--path="[path]"``: Enter the path for the source file. This path is
   relative to the H2O cluster.
-  ``--project-id=[id]``: Specify the ID of the project that will
   contain this source file

**Example**

The following example creates a project, then adds the allyears2k.csv
file to the Steam database.

::

    ./steam create project --name="Prediction" --description="Prediction project"
    ProjectId:  1
    ./steam create datasource --name="allyears2k.csv" --description="airline data" --path="../../Desktop/allyears2k.csv" --project-id=1
    DatasourceId:   1

--------------

``create identity``
~~~~~~~~~~~~~~~~~~~

**Description**

Creates a new user.

**Usage**

::

    ./steam create identity --name="[username]" --password="[password]"

**Parameters**

-  ``--name="[username]"``: Enter a unique string for the new user name
-  ``--password="[password]``: Enter a string for the new user's
   password

**Example**

The following example creates two users: bob and jim.

::

    ./steam create identity --name="bob" --password="bobSpassword"
    IdentityId: 2
    ./steam create identity --name="jim" --password="j1mSpassword"
    IdentityId: 3

--------------

``create project``
~~~~~~~~~~~~~~~~~~

**Description**

Creates a project in the Steam database. Once created, datasources can
be added to the project, ensuring that allo associated datasets and
models are contained in this single location.

**Usage**

::

    ./steam create project --name="[projectName]" --description="[description]"

**Parameters**

-  ``--name="[projectName]"``: Enter a unique name for the project
-  ``--description="[description]"``: Enter a description for the
   project

**Example**

The following example creates a Prediction project.

::

    ./steam create project --name="Prediction" --description="Prediction project"
    ProjectId:  1

--------------

``create role``
~~~~~~~~~~~~~~~

**Description**

Creates a new role.

**Usage**

::

    ./steam create role --name="[rolename]" --description="[description]"

**Parameters**

-  ``--name="[rolename]"``: Enter a unique string for the new role
-  ``--description="[description]"``: Optionally enter a string that
   describes the new role

**Example**

The following examples create an engineer role and then a datascience
role.

::

    ./steam create role --name="engineer" --description="a default engineer role"
    RoleId: 2
    ./steam create role --name="datascience" --description="a default data science role"
    RoleId: 3

--------------

``create workgroup``
~~~~~~~~~~~~~~~~~~~~

**Description**

Creates a new workgroup.

**Usage**

::

    ./steam create workgroup --name="[workgroupname]" --description="[description]"

**Parameters**

-  ``--name="[workgroupname]"``: Enter a unique string for the new
   workgroup
-  ``--description="[description]"``: Optionally enter a string that
   describes the new workgroup

**Example**

The following example creates a data preparation and a production
workgroup.

::

    ./steam create workgroup --name="preparation" --description="data prep group"   
    WorkgroupId:    1
    ./steam create workgroup --name="production" --description="production group"   
    WorkgroupId:    2
        

--------------

``deactivate identity``
~~~~~~~~~~~~~~~~~~~~~~~

**Description**

Deactivates an identity based on the specified username.

**Usage**

::

    ./steam deactivate identity --identity-id=[identityId]

**Parameters**

-  ``--identity-id=[identityId]``: Specify the identity of the user you
   want to deactivate.

**Example**

The following example deactivates a user whose ID is 3.

::

    ./steam deactivate identity --identity-id=3

--------------

``delete cluster``
~~~~~~~~~~~~~~~~~~

**Description**

Deletes the specified YARN cluster from the database. Note that this command can only be used with YARN clusters (i.e., those started using `start cluster`_.) This command will not work with local clusters. In addition, this commmand will only work on cluster that have been stopped using `stop cluster`_.

**Usage**

::

    ./steam delete cluster --cluster-id=[clusterId]

**Parameters**

-  ``--cluster-id=[clusterId]``: Specify the ID of the cluster that you
   want to delete.

**Example**

The following example retrieves a list of clusters, then stops and
deletes cluster 2.

::

    ./steam get clusters
    Id  Name    TypeId  DetailId    Address State   CreatedAt
    1   user    1       0           localhost:54321 started 1473883790
    2   user    1       0           localhost:54323 started 1474323838
    ./steam stop cluster --cluster-id=2
    ./steam delete cluster --cluster-id=2
    Cluster deleted: 1

--------------

``delete dataset``
~~~~~~~~~~~~~~~~~~

**Decription**

Deletes the specified dataset from the Steam database.

**Note**: You cannot delete a dataset that was used to build an existing
model. You must delete the model(s) first before you can delete the
dataset that was used to build the model.

**Usage**

::

    ./steam delete dataset --dataset-id=[datasetId]

**Parameters**

-  ``--dataset-id=[datasetId]``: Specify the ID of the dataset that that you want to delete. Note that you can use `get datasets`_ to retrieve a list of datasets in  the database.

**Example**

The following example deletes a dataset whose ID is 2.

::

    ./steam delete dataset --dataset-id=2

--------------

``delete datasource``
~~~~~~~~~~~~~~~~~~~~~

**Decription**

Deletes the specified data source file from the Steam database.

**Note**: You cannot delete a datasource that was used to build an
existing dataset. You must delete the dataset(s) first before you can
delete its source file.

**Usage**

::

    ./steam delete datasource --datasource-id=[datasourceId]

**Parameters**

-  ``--datasource-id=[datasourceId]``: Specify the ID of the file that
   that you want to delete. Note that you can use
   `get datasources`_ to retrieve a list of
   datasources in the database.

**Example**

The following example deletes a datasource whose ID is 4.

::

    ./steam delete datasource --datasource-id=4

--------------

``delete engine``
~~~~~~~~~~~~~~~~~

**Description**

Deletes the specified engine from the database.

**Usage**

::

    ./steam delete engine --engine-id=[engineId]

**Parameters**

-  ``--engine-id=[engineId]``: Specify the ID of the engine that you
   want to delete.

**Example**

The following example retrieves a list of engines currently added to the
database. It then specifies to delete that h2o-genmodel.jar engine.

::

    ./steam get engines
    Id  Name                Location            CreatedAt
    1   h2o-genmodel.jar    ../Desktop/engines  1473874219
    ./steam delete engine --engine-id=1

--------------

``delete model``
~~~~~~~~~~~~~~~~

**Description**

Deletes a model from the database based on the model's ID.

**Usage**

::

    ./steam delete model --model-id=[modelId]

**Parameters**

-  ``--model-id=[modelId]``: Specify the ID of the model that you want
   to delete.

**Example**

The following example deletes model 3 from the database. Note that you
can use `get models`_ to retrieve a list of models.

::

    ./steam delete model --model-id=3

--------------

``delete project``
~~~~~~~~~~~~~~~~~~

**Description**

Deletes a project from the database based on its ID.

**Note**: You cannot delete a project that includes existing data
(datasources, datasets, or models).

**Usage**

::

    ./steam delete project --project-id=[projectId]

**Parameters**

-  ``--project-id=[projectId]``: Specify the ID of the project that you
   want to delete.

**Example**

The following example deletes project 3 from the database. Note that you
can use `get projects`_ to retrieve a list of
projects.

::

    ./steam delete project --project-id=3

--------------

``delete role``
~~~~~~~~~~~~~~~

**Description**

Deletes a role from the database based on its ID.

**Usage**

::

    ./steam delete role --role-id=[roleId]

**Parameters**

-  ``--role-id=[roleId]``: Specify the ID of the role that you want to
   delete.

**Example**

The following example deletes role 3 from the database. Note that you
can use `get roles`_ to retrieve a list of roles. In
the case below, this role corresponds to the default data science role.

::

    ./steam delete role --role-id=3

--------------

``delete service``
~~~~~~~~~~~~~~~~~~

**Description**

A service represents a successfully deployed model on the Steam Prediction
Service. This command deletes a service from the database based on its
ID. Note that you must first stop a service before it can be deleted.
(See `stop service`_.)

**Usage**

::

    ./steam delete service --service-id=[id]

**Parameters**

-  ``--service-id=[id]``: Specify the ID of the service that you want to
   delete. Note that you can use `get services`_ to
   retrieve a list of services.

**Example**

The following example stops and then deletes service 2. This service
will no longer be available on the database.

::

    ./steam stop service --service-id=2
    ./steam delete service --service-id=2

--------------

``delete workgroup``
~~~~~~~~~~~~~~~~~~~~

**Description**

Deletes a workgroup from the database based on its ID.

**Usage**

::

    ./steam delete workgroup --workgroup-id=[workgroupId]

**Parameters**

-  ``--workgroup-id=[workgroupId]``: Specify the ID of the workgroup
   that you want to delete.

**Example**

The following example deletes workgroup 3 from the database. Note that
you can use `get workgroups`_ to retrieve a list of workgroups.

::

    ./steam delete workgroup --workgroup-id=3

--------------

``get all cluster-types``
~~~~~~~~~~~~~~~~~~~~~~~~~

**Description**

Retrieves a list of cluster types that are available in Steam along with
the corresponding code. 

**Usage**

::

    ./steam get all --cluster-types

**Parameters**

None

**Example**

The following example retrieves a list of the Steam cluster types.

::

    ./steam get all --cluster-types
    Id  Name        
    1   external
    2   yarn

--------------

``get all entity-types``
~~~~~~~~~~~~~~~~~~~~~~~~

**Description**

Retrieves a list of entity types that are available in Steam along with
the corresponding code. 

**Usage**

::

    ./steam get all --entity-types

**Parameters**

None

**Example**

The following example retrieves a list of Steam entity types.

::

    ./steam get all --entity-types
    Id  Name
    1   role        
    2   workgroup   
    3   identity    
    4   engine      
    5   cluster     
    6   project     
    7   datasource  
    8   dataset     
    9   model       
    10  label       
    11  service     

--------------

``get all permissions``
~~~~~~~~~~~~~~~~~~~~~~~

**Description**

Retrieves a list of permissions available in Steam along with the corresponding code. A permission code is used when linking roles to permissions.

**Note**: Permission IDs are randomly generated during installation, and the IDs will vary between Steam installations. 

**Usage**

::

    ./steam get all --permissions

**Parameters**

None

**Example**

The following example retrieves a list of Steam permissions.

::

    ./steam get all --permissions
    Id  Code                Description     
    9   ManageCluster       Manage clusters
    15  ManageDataset       Manage datasets
    13  ManageDatasource    Manage datasources
    7   ManageEngine        Manage engines
    5   ManageIdentity  Manage identities
    19  ManageLabel     Manage labels   
    17  ManageModel     Manage models   
    11  ManageProject       Manage projects
    1   ManageRole      Manage roles
    21  ManageService       Manage services
    3   ManageWorkgroup Manage workgroups
    10  ViewCluster     View clusters
    16  ViewDataset     View datasets
    14  ViewDatasource  View datasources
    8   ViewEngine      View engines
    6   ViewIdentity        View identities
    20  ViewLabel           View labels
    18  ViewModel           View models
    12  ViewProject     View projects
    2   ViewRole            View roles
    22  ViewService     View services
    4   ViewWorkgroup       View workgroups 

--------------

``get cluster``
~~~~~~~~~~~~~~~

**Description**

Retrieves detailed information for a specific cluster based on its ID.

**Usage**

::

    ./steam get cluster --cluster-id=[clusterId]

**Parameters**

-  ``--cluster-id=[clusterId]``: Specify the ID of the cluster that you
   want to retrieve

**Example**

The following example retrieves information for cluster ID 1.

::

    ./steam get cluster --cluster-id=1
    Attribute       Value
    Id:             1
    Name:           H2O_from_python_techwriter_hh4m3i
    TypeId:     1
    DetailId:       0
    Address:        localhost:54321
    State:          started
    CreatedAt:  1473883790

--------------

``get clusters``
~~~~~~~~~~~~~~~~

**Description**

Retrieves a list of clusters.

**Usage**

::

    ./steam get clusters --limit=[num]

**Parameters**

-  ``--limit=[num]``: Specify the maximum number of clusters that you want to retrieve.

**Example**

The following example retrieves a list of clusters that are running H2O
and are registered in Steam. (See `register cluster`_.)

::

    ./steam get clusters --limit=10
    Id  Name                      TypeId  DetailId  AddressState            CreatedAt 
    1 H2O_from_python_usr_6lvjb7  1       0         localhost:54321 started 1476306145

--------------

``get dataset``
~~~~~~~~~~~~~~~

**Description**

Retrieves information about a specific dataset based on its ID.

**Usage**

::

    ./steam get dataset --dataset-id=[datasetId]

**Parameters**

-  ``--dataset-id=[datasetId]``: Specify the ID of the dataset that you
   want to retrieve.

**Example**

The following example retrieves information about a dataset whose ID is
1. Note that you can use `get datasets`_ to retrieve
a list of all datasets.

::

    ./steam get dataset --dataset-id=1
    Attribute               Value
    Id:                     1
    DatasourceId:           2
    Name:               
    Description:        
    FrameName:          allyears2k.hex
    ResponseColumnName: Origin  
    JSONProperties:     {...<properties>...}
    CreatedAt:          1474321931

--------------

``get datasets``
~~~~~~~~~~~~~~~~

**Description**

Retrieves a list of all datasets available in the database.

**Usage**

::

    ./steam get datasets --limit=[num]

**Parameters**

-  ``--limit=[num]``: Specify the maximum number of datasets that you want to retrieve.

**Example**

The following example retrieves a list of all datasets.

::

    ./steam get datasets --limit=100
    Id  DatasourceId    Name    Description FrameName       ResponseColumnName  JSONProperties          CreatedAt
    1   2                                   prostate.csv    CAPSULE             {...<properties>...}    1473887458
    2   1                                   allyears2k.csv  Origin              {...<properties>...}    1474321931

--------------

``get datasource``
~~~~~~~~~~~~~~~~~~

**Description**

Retrieves information about a specific source file based on its ID.

**Usage**

::

    ./steam get datasource --datasource-id=[datasourceId]

**Parameters**

-  ``--datasource-id=[datasourceId]``: Specify the ID of the datasource
   that you want to retrieve.

**Example**

The following example retrieves information about a datasource whose ID
is 1. Note that you can use `get datasources`_ to
retrieve a list of all datasources.

::

    ./steam get datasource --datasource-id=1
    Attribute           Value
    Id:                 1
    ProjectId:      1
    Name:               allyears2k.csv
    Description:        airline data
    Kind:               CSV 
    Configuration:  {"path":"../Desktop"}
    CreatedAt:      1473879765

--------------

``get datasources``
~~~~~~~~~~~~~~~~~~~

**Description**

Retrieves a list of all datasources available in the database.

**Usage**

::

    ./steam get datasources --limit=[num]

**Parameters**

-  ``--limit=[num]``: Specify the maximum number of datasources that you want to retrieve.

**Example**

The following example retrieves a list of all datasources.

::

    ./steam get datasources --limit=100

    Id  ProjectId   Name            Description     Kind    Configuration           CreatedAt
    1   1           allyears2k.csv  airline data    CSV     {"path":"../Desktop"}   1473879765
    2   1           prostate.csv    prostate data   CSV     {"path":"../Desktop"}   1473880195

--------------

``get engine``
~~~~~~~~~~~~~~

**Description**

Retrieves information for a specific engine based on its ID.

**Usage**

::

    ./steam get engine --engine-id=[engineId]

**Parameters**

-  ``--engine-id=[engineId]``: Specify the ID of the engine that you
   want to retrieve

**Example**

The following example retrieves information about engine 1.

::

    ./steam get engine --engine-id=1
    Attribute       Value
    ID:             1
    Name:           h2o-genmodel.jar            
    Location:       ../Desktop/engines
    CreatedAt:  1473874219

--------------

``get engines``
~~~~~~~~~~~~~~~

**Description**

Retrieves a list of deployed engines.

**Usage**

::

    ./steam get engines

**Parameters**

None

**Example**

The following example retrieves a list of engines that have been
added. (Refer to `upload engine`_.)

::

    ./steam get engines
    Id  Name                Location            CreatedAt
    1   h2o-genmodel.jar    ../Desktop/engines  1473874219

--------------

``get identities``
~~~~~~~~~~~~~~~~~~

**Description**

Retrieves a list of users.

**Usage**

::

    ./steam get identities --limit=[num]

**Parameters**

-  ``--limit=[num]``: Specify the maximum number of identities that you want to retrieve.

**Example**

The following example retrieves a list of users that are available on
the database.

::

    ./steam get identities --limit=100
    Id    NAME        IsActive  LastLogin     Created          
     2    bob         true      -62135596804  1473883790
     3    jim         false     -62135596746  1474323838
     1    superuser   true      -62135596800  1476306094

--------------

``get identity``
~~~~~~~~~~~~~~~~

**Description**

Retrieve information about a specific user.

**Usage**

::

    ./steam get identity --identity-id=[identityId]
    ./steam get identity --by-name --name="[username]"

**Parameters**

-  ``[identityId]``: Specify the ID of the user you want to retrieve

**Example**

The following example retrieves information about a user whose ID is 2.

::

    ./steam get identity 2
    Attribute       Value       
    Id:             2       
    Name:           bob     
    IsActive:       true        
    LastLogin:      -62135596800    
    Created:        1474305548

--------------

``get model``
~~~~~~~~~~~~~

**Description**

Retrieves detailed information for a specific model.

**Usage**

::

    ./steam get model --model-id=[modelId]

**Parameters**

-  ``--model-id=[modelId]``: Specify the ID of the model that you want
   to retrieve

**Example**

The following example retrieves information for model 2.

::

    ./steam get model --model-id2

--------------

``get models``
~~~~~~~~~~~~~~

**Description**

Retrieves a list of models.

**Usage**

::

    ./steam get models --limit=[num]

**Parameters**

-  ``--limit=[num]``: Specify the maximum number of models that you want to retrieve.

**Example**

The following example retrieves a list of models that are available on
the database.

::

    ./steam get models --limit=100

--------------

``get permissions``
~~~~~~~~~~~~~~~~~~~

**Description**

Retrieves permission information for an identity or role.

**Usage**

::

    ./steam get permissions --for-role --role-id=[roleId] 
    ./steam get permissions --for-identity --identity-id=[identityId] 

**Parameters**

-  ``--role-id=[roleId]``: When retrieving permissions for a role,
   specify the ID of the role that you want to view
-  ``--identity-id=[identityId]``: When retrieving permissions for an
   identity, specify the ID that you want to view

**Examples**

The following example retrieves the permissions assigned to a role whose
ID is 2.

::

    Id  Code            Description     
    18  ViewModel       View models     
    12  ViewProject     View projects       
    4   ViewWorkgroup    View workgroups    

--------------

``get project``
~~~~~~~~~~~~~~~

**Description**

Retrieves detailed information for a specific project based on its ID.

**Usage**

::

    ./steam get project --project-id=[id]

**Parameters**

-  ``--project-id=[id]``: Specify the ID of the project that you want to
   retrieve

**Examples**

The following example retrieves information about a project whose ID is
1. Note that you can use `get projects`_ to retrieve
a list of all projects and IDs.

::

    ./steam get project --project-id=1
    Attribute       Value               
    Id:             1               
    Name:           Prediction          
    Description:    Prediction project  
    ModelCategory:                  
    CreatedAt:      1473878624  

--------------

``get projects``
~~~~~~~~~~~~~~~~

**Description**

Retrieves a list of all projects in the Steam database.

**Usage**

::

    ./steam get projects --limit=[num]

**Parameters**

-  ``--limit=[num]``: Specify the maximum number of projects that you want to retrieve.

**Example**

The following example retrieves a list of projects that are available on
the database.

::

    ./steam get projects --limit=10

    Id  Name        Description             ModelCategory   CreatedAt
    1   Prediction  Prediction project      Classification  1473878624
    2   Churn       Customer churn project  Regression      1473879033

--------------

``get role``
~~~~~~~~~~~~

**Description**

Retrieves detailed information for a specific role based on its name.

**Usage**

::

    ./steam get role --role-id=[id]

**Parameters**

-  ``--role-id=[id]``: Specify the ID of the role that you want to
   retrieve

**Example**

The following example retrieves information about the datascience role.

::

    ./steam get role --role-id=2
    Attribute       Value
    Id:             2
    Name:           datascience
    Description:    a default data science role
    Created:        1473874053

--------------

``get roles``
~~~~~~~~~~~~~

**Description**

Retrieves a list of roles.

**Usage**

::

    ./steam get roles --limit=[num]

**Parameters**

-  ``--limit=[num]``: Specify the maximum number of identities that you want to retrieve.

**Example**

The following example retrieves a list of roles that are available on
the database.

::

    ./steam get roles --limit=10
    Id    Name        Description                 Created
    1     Superuser   Superuser                   1473874053
    2     datascience a default data science role 1473893347  

--------------

``get service``
~~~~~~~~~~~~~~~

**Description**

A service represents a successfully deployed model on the Steam Prediction
Service. This command retrieves detailed information about a specific
service based on its ID.

**Usage**

::

    ./steam get service [serviceId]

**Parameters**

-  ``[serviceId]``: Specify the ID of the service that you want to
   retrieve

**Example**

The following example retrieve information about service 2.

::

    ./steam get service 2

--------------

``get services``
~~~~~~~~~~~~~~~~

**Description**

A service represents a successfully deployed model on the Steam Prediction
Service. This command retrieves a list of services available on the
database.

**Usage**

::

    ./steam get services --limit=[num]

**Parameters**

-  ``--limit=[num]``: Specify the maximum number of services that you want to retrieve.

**Example**

The following example retrieves a list of services that are available on
the database.

::

    ./steam get services --limit=10
    Id  ModelId Name      Address     Port  ProcessId State   CreatedAt 
    1   1       IrisModel 172.16.2.89 50336 26200     started 1476306364

--------------

``get workgroup``
~~~~~~~~~~~~~~~~~

**Description**

Retrieves information for a specific workgroup based on its name.

**Usage**

::

    ./steam get workgroup [workgroupName]

**Parameters**

-  ``[workgroupName]``: Specify the name of the workgroup that you want
   to retrieve

**Example**

The following example retrieves information about the production
workgroup

::

    ./steam get workgroup production
                    production
    DESCRIPTION:    production group
    ID:     3
    AGE:    2016-07-15 09:32:27 -0700 PDT

    IDENTITIES: 1
    NAME    STATUS  LAST LOGIN
    jim     Active  0000-12-31 16:00:00 -0800 PST

--------------

``get workgroups``
~~~~~~~~~~~~~~~~~~

**Description**

Retrieves a list of workgroups currently available on the database.

**Usage**

::

    ./steam get workgroups --identity=[identityName] --limit=[num]

**Parameters**

-  ``--identity=[identityName]``: Optionally specify to view all
   workgroups associated with a specific user name
-  ``--limit=[num]``: Specify the maximum number of workgroups that you want to retrieve

**Example**

The following example retrieves a list of workgroups that are available
on the database.

::

    ./steam get workgroups --limit=1
    Id    Name        Description         Created
    2     preparation data prep group     1473874219
    3     production  production group    1473879765

--------------

``import model``
~~~~~~~~~~~~~~~~

**Description**

Imports a model from H2O based on its ID.

**Usage**

::

    ./steam import model [clusterId] [modelName]

**Parameters**

-  ``[clusterId``]: Specify the H2O cluster that contains the model you
   want to import
-  ``[modelName]``: Specify the name of the that you want to import into
   steam.

**Example**

The following example specifies to import the
GBM_model_python_1468599779202_1 model from Cluster 1.

::

    ./steam import model 1 GBM_model_python_1468599779202_1

--------------

``link identity``
~~~~~~~~~~~~~~~~~

**Description**

Links a user to a specific role or workgroup.

**Usage**

::

    ./steam link identity --with-role --identity-id=[identityId] --role-id=[roleId]
    ./steam link identity --with-workgroup --identity-id=[identityId] --workgroup-id=[workgroupId]

**Parameters**

-  Link identity to a specific role:

   -  ``--with-role``: Enable this flag to associate an identity with a
      role
   -  ``--identity-id=[identityId]``: Specify the ID of user that will
      be linked to a role
   -  ``--role-id=[roleId]``: Specify the ID of the role that the user
      will be linked to

-  Link identity to a specific workgroup:

   -  ``--with-workgroup``: Enable this flag to associate an identity
      with a workgroup
   -  ``--identity-id=[identityId]``: Specify the ID of user that will
      be linked to a workgroup
   -  ``--workgroup-id=[workgroupId]``: Specify the ID of the workgroup
      that the the user will be linked to

**Example**

The following example links user Jim to datascience role and then to the
production workgroup.

::

    ./steam link identity --with-role --identity-id=3 --role-id=3
    ./steam link identity --with-workgroup --identity-id=3 --workgroup-id=3

--------------

``link role``
~~~~~~~~~~~~~

**Description**

Links a role to a certain set of permissions.

**Usage**

::

    ./steam link role --with-permission --role-id=[roleId] --permission-id=[permissionId]

**Parameters**

-  ``--with-permission``: Enable this flag when setting permissions
-  ``role-id=[roleId]``: Specify the role that the user will be linked
   to.
-  ``permission-id=[permissionId]``: Specify a single permission to
   assign to this role.

**Example**

The following example links the datascience role to the ManageProject, ManageModel, and ViewCluster permissions. Note that you can use `get all permissions`_ to view a list of permission IDs.

::

        ./steam link role --with-permission --role-id=3 --permission-id=11
        ./steam link role --with-permission --role-id=3 --permission-id=17
        ./steam link role --with-permission --role-id=3 --permission-id=10

--------------

``login``
~~~~~~~~~

**Description**

Logs a user in to Steam

**Usage**

::

    ./steam login [address:port] --username=[userName] --password=[password]

**Parameters**

-  ``[address:port]``: Specify the address and port of the Steam server.
-  ``--username=[userName]``: Specify the username.
-  ``--password=[password]``: Specify the user's password.

**Example**

The following example logs user Bob into a Steam instance running on
localhost:9000.

::

    ./steam login localhost:9000 --username=bob --password=bobSpassword
    Login credentials saved for server localhost:9000

--------------

``register cluster``
~~~~~~~~~~~~~~~~~~~~

**Description**

Registers a cluster that is currently running H2O (typically a local
cluster). Once registered, the cluster can be used to perform machine
learning tasks through Python, R, and Flow. The cluster will also be
visible in the Steam web UI.

Note that clusters that are started using this command can be stopped
from within the web UI or using `unregister cluster`_. You will receive an
error if you attemt to stop registered clusters using the
``stop cluster`` command.

**Usage**

::

    ./steam register cluster --address="[address]"

**Parameters**

-  ``--address="[address]"``: Specify the IP address and port of the
   cluster that you want to register.

**Example**

The following example registers Steam on localhost:54323. Note that this
will only be successful if H2O is already running on this cluster.

::

    ./steam register cluster --address="localhost:54323"
    ClusterId:  2

--------------

``reset``
~~~~~~~~~

**Description**

Resets the current Steam cluster instance. This removes the current
authentication from Steam. You will have to re-authenticate in order to
continue to use Steam.

**Usage**

::

    ./steam reset

**Parameters**

None

**Examples**

The following example resets the current Steam instance.

::

    ./steam reset
    Configuration reset successfully. Use 'steam login <server-address>' to re-authenticate to Steam

--------------

``start cluster``
~~~~~~~~~~~~~~~~~

**Description**

After you have deployed engine, you can use this command to start a new
cluster through YARN using a specified engine. Note that this command is
only valid when starting Steam on a YARN cluster. To start Steam on a
local cluster, use `register cluster`_ instead.

**Usage**

::

    ./steam start cluster [id] [engineId] --size=[numNodes] --memory=[string]

**Parameters**

-  ``[id]``: Enter an ID for this new cluster.
-  ``[engineId]``: Specify the ID of the engine that this cluster will
   use. If necessary, use `get engines`_ to retrieve a list of all available engines.
-  ``--size=[numNodes]``: Specify an integer for the number of nodes in
   this cluster.
-  ``--memory=[string]``: Enter a string specifying the amount of memory
   available to Steam in each node (for example, "1024m", "2g", etc.)

**Example**

The following example retrieves a list of engines, then starts a cluster
through YARN using an engine from the list. The cluster is configured
with 2 nodes that are 2 gigabytes each.

::

    ./steam get engines
    NAME                ID  AGE
    h2o-genmodel.jar    1   2016-07-01 13:30:50 -0700 PDT
    h2o.jar         2   2016-07-01 13:32:10 -0700 PDT
    ./steam start cluster 9 1 --size=2 --memory=2g

--------------

``stop cluster``
~~~~~~~~~~~~~~~~

**Description**

Stops a YARN cluster that was started through the CLI or web UI. (See `start cluster`_.) Note that you will receive an error if you attempt to stop a cluster that was started using `register cluster`_.

**Usage**

::

    ./steam stop cluster [id] 

**Parameters**

-  ``[id]``: Specify the ID of the cluster that you want to stop. If
   necessary, use `get clusters`_ to retrieve a list of clusters.

**Example**

The following example stops a cluster that has an ID of 9.

::

    ./steam stop cluster 9

--------------

``stop service``
~~~~~~~~~~~~~~~~

**Description**

A service represents a successfully deployed model on the Steam Prediction
Service. Use this command to stop a service.

**Usage**

::

    ./steam stop service [serviceId] 

**Parameters**

-  ``[serviceId]``: Specify the ID of the scoring service that you want
   to stop. If necessary, use `get services`_ to
   retrieve a list of running services.

**Example**

The following example stops a service that has an ID of 2.

::

    ./steam stop service 2

--------------

``unlink identity``
~~~~~~~~~~~~~~~~~~~

**Description**

Removes a user's permissions from a specific role or workgroup.

**Usage**

::

    ./steam unlink identity [identityName] [role [roleId] | workgroup [workgroupId]]

**Parameters**

-  ``[identityName]``: Specify the user that will be unlinked from a
   role or workgroup
-  ``role [roleId]``: Specify the role that the user will be unlinked
   from
-  ``workgroup [workgroupId]``: Specify the workgroup that the the user
   will be unlinked from

**Example**

The following example removes user Jim from the datascience role and
then from the production workgroup.

::

    ./steam unlink identity jim role datascience
    ./steam unlink identity jim workgroup production

--------------

``unregister cluster``
~~~~~~~~~~~~~~~~~~~~~~

**Description**

Stops a cluster that was registered through the CLI or the web UI. (See `register cluster`_.) Note that this does not delete the cluster. Also note that you will receive an error if you attempt to unregister a cluster that was started using `start cluster`_.

**Usage**

::

    ./steam unregister cluster [id] 

**Parameters**

-  ``[id]``: Specify the ID of the cluster that you want to stop. If
   necessary, use `get clusters`_ to retrieve a list of clusters.

**Example**

The following example stops a cluster that has an ID of 9.

::

    ./steam unregister cluster 2
    Successfully unregisted cluster %d 2

--------------

``update role``
~~~~~~~~~~~~~~~

**Description**

Edits the description and/or name of an existing role. When a role is
edited, the edit will automatically propagate to all identities that are
associated with this role.

**Usage**

::

    ./steam update role [rolename] --desc="[description]" --name="[newRoleName]

**Parameters**

-  ``[rolename]``: Enter the role name that you want to edit
-  ``desc="[description]"``: Optionally enter a string that describes
   the new role
-  ``name="[newRoleName]"``: Enter a unique string for the new role name

**Example**

The following example changes the name of the engineer role to be
"science engineer".

::

    ./steam update role engineer --desc="A better engineer" --name="science engineer"
    Successfully updated role: engineer
        

--------------

``update workgroup``
~~~~~~~~~~~~~~~~~~~~

**Description**

Edits the description and/or name of an existing workgroup. When a
workgroup is edited, the edit will automatically propagate to all
identities that are associated with this workgroup.

**Usage**

::

    ./steam update workgroup [workgroupname] --desc="[description]" --name="[newWorkgroupName]

**Parameters**

-  ``[workgroup]``: Enter the workgroup name that you want to edit
-  ``desc="[description]"``: Optionally enter a string that describes
   the new workgroup
-  ``name="[newWorkgroupName]"``: Enter a unique string for the new
   workgroup name

**Example**

The following example changes the name of the production workgroup to be
"deploy".

::

    ./steam update workgroup production --desc="A deploy workgroup" --name="deploy"
    Successfully updated workgroup: production

--------------

``upload engine``
~~~~~~~~~~~~~~~~~

**Description**

Adds a new engine to the Steam database. After an engine is successfully added, it can be specified when starting a cluster. (See `start cluster`_.)

**Usage**

::

    ./steam upload engine --file-path="[path]"

**Parameters**

-  ``--file-path="[path]"``: Enter the path for the engine that you want to upload

**Example**

The following example adds **h2o-genmodel.jar** to the list of available
engines.

::

    ./steam upload engine --file-path="../Desktop/engines/h2o.genmodel.jar"

--------------

``upload file``
~~~~~~~~~~~~~~~

Adds a new preprocessing file to the Steam database. 

**Usage**

::

    ./steam upload file --file-path="[path]" --project-id=[id] --package-name="[target_name]" --relative-path="[path_to_copy_to]"

**Parameters**

-  ``--file-path="[path]"``: Enter the path for the preprocessing file that you want to upload
-  ``--project-id=[id]``: Preprocessing files must be associated with a project. Enter the ID of the project that will have access to this file.
-  ``--package-name="[target_name]"``: Specify the name for this package
-  ``--relative-path="[path_to_copy_to]"``: Specify the relative path to copy this file to

**Example**

The following example adds a preprocessing file to a project whose ID is 5. The file will be copied to the Steam assets folder.

::

  /steam upload file --file-path="../preprocess/score.py" --package-name="score.py" --project-id=5 --relative-path="var/master/assets"
