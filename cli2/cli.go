// ------------------------------
// --- This is generated code ---
// ---      DO NOT EDIT       ---
// ------------------------------
package cli2

import (
  "github.com/spf13/cobra"
  "log"
  "fmt"
)



func registerGeneratedCommands(c *context, cmd *cobra.Command) {
  cmd.AddCommand(
    add(c),
    build(c),
    create(c),
    deactivate(c),
    delete_(c),
    get(c),
    import_(c),
    link(c),
    ping(c),
    register(c),
    share(c),
    split(c),
    start(c),
    stop(c),
    unlink(c),
    unregister(c),
    unshare(c),
    update(c),
  )
}



var addHelp = `
add [?]
Add entities
Commands:

    $ steam add engine ...
`
func add(c *context) *cobra.Command {
  cmd := newCmd(c, addHelp, nil)

  cmd.AddCommand(addEngine(c))
  return cmd
}


var addEngineHelp = `
engine [?]
Add Engine
Examples:

    Add an engine
    $ steam add engine \
        --engine-name=? \
        --engine-path=?

`
func addEngine(c *context) *cobra.Command {
  
  
  var engineName string
  var enginePath string
  cmd := newCmd(c, addEngineHelp, func(c *context, args []string) {
    
    
  engineId, err := c.remote.AddEngine(
  
    engineName,
  
    enginePath,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", engineId)
  
  return

    
  })
  
  
  cmd.Flags().StringVar(&engineName, "engine-name", engineName, "No description available")
  cmd.Flags().StringVar(&enginePath, "engine-path", enginePath, "No description available")
  return cmd
}




var buildHelp = `
build [?]
Build entities
Commands:

    $ steam build model ...
`
func build(c *context) *cobra.Command {
  cmd := newCmd(c, buildHelp, nil)

  cmd.AddCommand(buildModel(c))
  return cmd
}


var buildModelHelp = `
model [?]
Build Model
Examples:

    Build a model
    $ steam build model \
        --cluster-id=? \
        --dataset-id=? \
        --algorithm=?

    Build an AutoML model
    $ steam build model --auto \
        --cluster-id=? \
        --dataset=? \
        --target-name=? \
        --max-run-time=?

`
func buildModel(c *context) *cobra.Command {
  var auto bool
  
  
  var clusterId int64
  var datasetId int64
  var algorithm string
  var dataset string
  var targetName string
  var maxRunTime int
  cmd := newCmd(c, buildModelHelp, func(c *context, args []string) {
    
    
    
    
    
    if auto {
      
  model, err := c.remote.BuildModelAuto(
  
    clusterId,
  
    dataset,
  
    targetName,
  
    maxRunTime,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", model)
  
  return

    }
    
    
    
    
    if true {
      
  modelId, err := c.remote.BuildModel(
  
    clusterId,
  
    datasetId,
  
    algorithm,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", modelId)
  
  return

    }
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&auto, "auto", auto, "Build an AutoML model")
  
  
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  cmd.Flags().Int64Var(&datasetId, "dataset-id", datasetId, "No description available")
  cmd.Flags().StringVar(&algorithm, "algorithm", algorithm, "No description available")
  cmd.Flags().StringVar(&dataset, "dataset", dataset, "No description available")
  cmd.Flags().StringVar(&targetName, "target-name", targetName, "No description available")
  cmd.Flags().IntVar(&maxRunTime, "max-run-time", maxRunTime, "No description available")
  return cmd
}




var createHelp = `
create [?]
Create entities
Commands:

    $ steam create dataset ...
    $ steam create datasource ...
    $ steam create identity ...
    $ steam create project ...
    $ steam create role ...
    $ steam create workgroup ...
`
func create(c *context) *cobra.Command {
  cmd := newCmd(c, createHelp, nil)

  cmd.AddCommand(createDataset(c))
  cmd.AddCommand(createDatasource(c))
  cmd.AddCommand(createIdentity(c))
  cmd.AddCommand(createProject(c))
  cmd.AddCommand(createRole(c))
  cmd.AddCommand(createWorkgroup(c))
  return cmd
}


var createDatasetHelp = `
dataset [?]
Create Dataset
Examples:

    Create a dataset
    $ steam create dataset \
        --cluster-id=? \
        --datasource-id=? \
        --name=? \
        --description=? \
        --response-column-name=?

`
func createDataset(c *context) *cobra.Command {
  
  
  var clusterId int64
  var datasourceId int64
  var name string
  var description string
  var responseColumnName string
  cmd := newCmd(c, createDatasetHelp, func(c *context, args []string) {
    
    
  datasetId, err := c.remote.CreateDataset(
  
    clusterId,
  
    datasourceId,
  
    name,
  
    description,
  
    responseColumnName,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", datasetId)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  cmd.Flags().Int64Var(&datasourceId, "datasource-id", datasourceId, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().StringVar(&responseColumnName, "response-column-name", responseColumnName, "No description available")
  return cmd
}

var createDatasourceHelp = `
datasource [?]
Create Datasource
Examples:

    Create a datasource
    $ steam create datasource \
        --project-id=? \
        --name=? \
        --description=? \
        --path=?

`
func createDatasource(c *context) *cobra.Command {
  
  
  var projectId int64
  var name string
  var description string
  var path string
  cmd := newCmd(c, createDatasourceHelp, func(c *context, args []string) {
    
    
  datasourceId, err := c.remote.CreateDatasource(
  
    projectId,
  
    name,
  
    description,
  
    path,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", datasourceId)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&projectId, "project-id", projectId, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().StringVar(&path, "path", path, "No description available")
  return cmd
}

var createIdentityHelp = `
identity [?]
Create Identity
Examples:

    Create an identity
    $ steam create identity \
        --name=? \
        --password=?

`
func createIdentity(c *context) *cobra.Command {
  
  
  var password string
  var name string
  cmd := newCmd(c, createIdentityHelp, func(c *context, args []string) {
    
    
  identityId, err := c.remote.CreateIdentity(
  
    name,
  
    password,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", identityId)
  
  return

    
  })
  
  
  cmd.Flags().StringVar(&password, "password", password, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  return cmd
}

var createProjectHelp = `
project [?]
Create Project
Examples:

    Create a project
    $ steam create project \
        --name=? \
        --description=?

`
func createProject(c *context) *cobra.Command {
  
  
  var name string
  var description string
  cmd := newCmd(c, createProjectHelp, func(c *context, args []string) {
    
    
  projectId, err := c.remote.CreateProject(
  
    name,
  
    description,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", projectId)
  
  return

    
  })
  
  
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  return cmd
}

var createRoleHelp = `
role [?]
Create Role
Examples:

    Create a role
    $ steam create role \
        --name=? \
        --description=?

`
func createRole(c *context) *cobra.Command {
  
  
  var name string
  var description string
  cmd := newCmd(c, createRoleHelp, func(c *context, args []string) {
    
    
  roleId, err := c.remote.CreateRole(
  
    name,
  
    description,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", roleId)
  
  return

    
  })
  
  
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  return cmd
}

var createWorkgroupHelp = `
workgroup [?]
Create Workgroup
Examples:

    Create a workgroup
    $ steam create workgroup \
        --name=? \
        --description=?

`
func createWorkgroup(c *context) *cobra.Command {
  
  
  var name string
  var description string
  cmd := newCmd(c, createWorkgroupHelp, func(c *context, args []string) {
    
    
  workgroupId, err := c.remote.CreateWorkgroup(
  
    name,
  
    description,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", workgroupId)
  
  return

    
  })
  
  
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  return cmd
}




var deactivateHelp = `
deactivate [?]
Deactivate entities
Commands:

    $ steam deactivate identity ...
`
func deactivate(c *context) *cobra.Command {
  cmd := newCmd(c, deactivateHelp, nil)

  cmd.AddCommand(deactivateIdentity(c))
  return cmd
}


var deactivateIdentityHelp = `
identity [?]
Deactivate Identity
Examples:

    Deactivate an identity
    $ steam deactivate identity \
        --identity-id=?

`
func deactivateIdentity(c *context) *cobra.Command {
  
  
  var identityId int64
  cmd := newCmd(c, deactivateIdentityHelp, func(c *context, args []string) {
    
    
  err := c.remote.DeactivateIdentity(
  
    identityId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&identityId, "identity-id", identityId, "No description available")
  return cmd
}




var deleteHelp = `
delete [?]
Delete entities
Commands:

    $ steam delete cluster ...
    $ steam delete dataset ...
    $ steam delete datasource ...
    $ steam delete engine ...
    $ steam delete model ...
    $ steam delete project ...
    $ steam delete role ...
    $ steam delete service ...
    $ steam delete workgroup ...
`
func delete_(c *context) *cobra.Command {
  cmd := newCmd(c, deleteHelp, nil)

  cmd.AddCommand(deleteCluster(c))
  cmd.AddCommand(deleteDataset(c))
  cmd.AddCommand(deleteDatasource(c))
  cmd.AddCommand(deleteEngine(c))
  cmd.AddCommand(deleteModel(c))
  cmd.AddCommand(deleteProject(c))
  cmd.AddCommand(deleteRole(c))
  cmd.AddCommand(deleteService(c))
  cmd.AddCommand(deleteWorkgroup(c))
  return cmd
}


var deleteClusterHelp = `
cluster [?]
Delete Cluster
Examples:

    Delete a cluster
    $ steam delete cluster \
        --cluster-id=?

`
func deleteCluster(c *context) *cobra.Command {
  
  
  var clusterId int64
  cmd := newCmd(c, deleteClusterHelp, func(c *context, args []string) {
    
    
  err := c.remote.DeleteCluster(
  
    clusterId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  return cmd
}

var deleteDatasetHelp = `
dataset [?]
Delete Dataset
Examples:

    Delete a dataset
    $ steam delete dataset \
        --dataset-id=?

`
func deleteDataset(c *context) *cobra.Command {
  
  
  var datasetId int64
  cmd := newCmd(c, deleteDatasetHelp, func(c *context, args []string) {
    
    
  err := c.remote.DeleteDataset(
  
    datasetId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&datasetId, "dataset-id", datasetId, "No description available")
  return cmd
}

var deleteDatasourceHelp = `
datasource [?]
Delete Datasource
Examples:

    Delete a datasource
    $ steam delete datasource \
        --datasource-id=?

`
func deleteDatasource(c *context) *cobra.Command {
  
  
  var datasourceId int64
  cmd := newCmd(c, deleteDatasourceHelp, func(c *context, args []string) {
    
    
  err := c.remote.DeleteDatasource(
  
    datasourceId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&datasourceId, "datasource-id", datasourceId, "No description available")
  return cmd
}

var deleteEngineHelp = `
engine [?]
Delete Engine
Examples:

    Delete an engine
    $ steam delete engine \
        --engine-id=?

`
func deleteEngine(c *context) *cobra.Command {
  
  
  var engineId int64
  cmd := newCmd(c, deleteEngineHelp, func(c *context, args []string) {
    
    
  err := c.remote.DeleteEngine(
  
    engineId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&engineId, "engine-id", engineId, "No description available")
  return cmd
}

var deleteModelHelp = `
model [?]
Delete Model
Examples:

    Delete a model
    $ steam delete model \
        --model-id=?

`
func deleteModel(c *context) *cobra.Command {
  
  
  var modelId int64
  cmd := newCmd(c, deleteModelHelp, func(c *context, args []string) {
    
    
  err := c.remote.DeleteModel(
  
    modelId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&modelId, "model-id", modelId, "No description available")
  return cmd
}

var deleteProjectHelp = `
project [?]
Delete Project
Examples:

    Delete a project
    $ steam delete project \
        --project-id=?

`
func deleteProject(c *context) *cobra.Command {
  
  
  var projectId int64
  cmd := newCmd(c, deleteProjectHelp, func(c *context, args []string) {
    
    
  err := c.remote.DeleteProject(
  
    projectId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&projectId, "project-id", projectId, "No description available")
  return cmd
}

var deleteRoleHelp = `
role [?]
Delete Role
Examples:

    Delete a role
    $ steam delete role \
        --role-id=?

`
func deleteRole(c *context) *cobra.Command {
  
  
  var roleId int64
  cmd := newCmd(c, deleteRoleHelp, func(c *context, args []string) {
    
    
  err := c.remote.DeleteRole(
  
    roleId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&roleId, "role-id", roleId, "No description available")
  return cmd
}

var deleteServiceHelp = `
service [?]
Delete Service
Examples:

    Delete a service
    $ steam delete service \
        --service-id=?

`
func deleteService(c *context) *cobra.Command {
  
  
  var serviceId int64
  cmd := newCmd(c, deleteServiceHelp, func(c *context, args []string) {
    
    
  err := c.remote.DeleteService(
  
    serviceId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&serviceId, "service-id", serviceId, "No description available")
  return cmd
}

var deleteWorkgroupHelp = `
workgroup [?]
Delete Workgroup
Examples:

    Delete a workgroup
    $ steam delete workgroup \
        --workgroup-id=?

`
func deleteWorkgroup(c *context) *cobra.Command {
  
  
  var workgroupId int64
  cmd := newCmd(c, deleteWorkgroupHelp, func(c *context, args []string) {
    
    
  err := c.remote.DeleteWorkgroup(
  
    workgroupId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&workgroupId, "workgroup-id", workgroupId, "No description available")
  return cmd
}




var getHelp = `
get [?]
Get entities
Commands:

    $ steam get all ...
    $ steam get cluster ...
    $ steam get clusters ...
    $ steam get dataset ...
    $ steam get datasets ...
    $ steam get datasource ...
    $ steam get datasources ...
    $ steam get engine ...
    $ steam get engines ...
    $ steam get history ...
    $ steam get identities ...
    $ steam get identity ...
    $ steam get job ...
    $ steam get jobs ...
    $ steam get model ...
    $ steam get models ...
    $ steam get permissions ...
    $ steam get privileges ...
    $ steam get project ...
    $ steam get projects ...
    $ steam get role ...
    $ steam get roles ...
    $ steam get service ...
    $ steam get services ...
    $ steam get workgroup ...
    $ steam get workgroups ...
`
func get(c *context) *cobra.Command {
  cmd := newCmd(c, getHelp, nil)

  cmd.AddCommand(getAll(c))
  cmd.AddCommand(getCluster(c))
  cmd.AddCommand(getClusters(c))
  cmd.AddCommand(getDataset(c))
  cmd.AddCommand(getDatasets(c))
  cmd.AddCommand(getDatasource(c))
  cmd.AddCommand(getDatasources(c))
  cmd.AddCommand(getEngine(c))
  cmd.AddCommand(getEngines(c))
  cmd.AddCommand(getHistory(c))
  cmd.AddCommand(getIdentities(c))
  cmd.AddCommand(getIdentity(c))
  cmd.AddCommand(getJob(c))
  cmd.AddCommand(getJobs(c))
  cmd.AddCommand(getModel(c))
  cmd.AddCommand(getModels(c))
  cmd.AddCommand(getPermissions(c))
  cmd.AddCommand(getPrivileges(c))
  cmd.AddCommand(getProject(c))
  cmd.AddCommand(getProjects(c))
  cmd.AddCommand(getRole(c))
  cmd.AddCommand(getRoles(c))
  cmd.AddCommand(getService(c))
  cmd.AddCommand(getServices(c))
  cmd.AddCommand(getWorkgroup(c))
  cmd.AddCommand(getWorkgroups(c))
  return cmd
}


var getAllHelp = `
all [?]
Get All
Examples:

    List all entity types
    $ steam get all --entity-types

    List all permissions
    $ steam get all --permissions

    List all cluster types
    $ steam get all --cluster-types

`
func getAll(c *context) *cobra.Command {
  var entityTypes bool
  var permissions bool
  var clusterTypes bool
  
  
  cmd := newCmd(c, getAllHelp, func(c *context, args []string) {
    
    
    
    if entityTypes {
      
  entityTypes, err := c.remote.GetAllEntityTypes(
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", entityTypes)
  
  return

    }
    
    
    
    if permissions {
      
  permissions, err := c.remote.GetAllPermissions(
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", permissions)
  
  return

    }
    
    
    
    if clusterTypes {
      
  clusterTypes, err := c.remote.GetAllClusterTypes(
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", clusterTypes)
  
  return

    }
    
    
    
    
    
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&entityTypes, "entity-types", entityTypes, "List all entity types")
  cmd.Flags().BoolVar(&permissions, "permissions", permissions, "List all permissions")
  cmd.Flags().BoolVar(&clusterTypes, "cluster-types", clusterTypes, "List all cluster types")
  
  
  return cmd
}

var getClusterHelp = `
cluster [?]
Get Cluster
Examples:

    Get cluster details
    $ steam get cluster \
        --cluster-id=?

    Get cluster details (Yarn only)
    $ steam get cluster --on-yarn \
        --cluster-id=?

    Get cluster status
    $ steam get cluster --status \
        --cluster-id=?

`
func getCluster(c *context) *cobra.Command {
  var onYarn bool
  var status bool
  
  
  var clusterId int64
  cmd := newCmd(c, getClusterHelp, func(c *context, args []string) {
    
    
    
    
    
    if onYarn {
      
  cluster, err := c.remote.GetClusterOnYarn(
  
    clusterId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", cluster)
  
  return

    }
    
    
    
    if status {
      
  clusterStatus, err := c.remote.GetClusterStatus(
  
    clusterId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", clusterStatus)
  
  return

    }
    
    
    
    
    if true {
      
  cluster, err := c.remote.GetCluster(
  
    clusterId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", cluster)
  
  return

    }
    
    
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&onYarn, "on-yarn", onYarn, "Get cluster details (Yarn only)")
  cmd.Flags().BoolVar(&status, "status", status, "Get cluster status")
  
  
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  return cmd
}

var getClustersHelp = `
clusters [?]
Get Clusters
Examples:

    List clusters
    $ steam get clusters \
        --offset=? \
        --limit=?

`
func getClusters(c *context) *cobra.Command {
  
  
  var offset int64
  var limit int64
  cmd := newCmd(c, getClustersHelp, func(c *context, args []string) {
    
    
  clusters, err := c.remote.GetClusters(
  
    offset,
  
    limit,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", clusters)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  return cmd
}

var getDatasetHelp = `
dataset [?]
Get Dataset
Examples:

    Get dataset details
    $ steam get dataset \
        --dataset-id=?

`
func getDataset(c *context) *cobra.Command {
  
  
  var datasetId int64
  cmd := newCmd(c, getDatasetHelp, func(c *context, args []string) {
    
    
  dataset, err := c.remote.GetDataset(
  
    datasetId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", dataset)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&datasetId, "dataset-id", datasetId, "No description available")
  return cmd
}

var getDatasetsHelp = `
datasets [?]
Get Datasets
Examples:

    List datasets
    $ steam get datasets \
        --datasource-id=? \
        --offset=? \
        --limit=?

`
func getDatasets(c *context) *cobra.Command {
  
  
  var datasourceId int64
  var offset int64
  var limit int64
  cmd := newCmd(c, getDatasetsHelp, func(c *context, args []string) {
    
    
  datasets, err := c.remote.GetDatasets(
  
    datasourceId,
  
    offset,
  
    limit,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", datasets)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&datasourceId, "datasource-id", datasourceId, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  return cmd
}

var getDatasourceHelp = `
datasource [?]
Get Datasource
Examples:

    Get datasource details
    $ steam get datasource \
        --datasource-id=?

`
func getDatasource(c *context) *cobra.Command {
  
  
  var datasourceId int64
  cmd := newCmd(c, getDatasourceHelp, func(c *context, args []string) {
    
    
  datasource, err := c.remote.GetDatasource(
  
    datasourceId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", datasource)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&datasourceId, "datasource-id", datasourceId, "No description available")
  return cmd
}

var getDatasourcesHelp = `
datasources [?]
Get Datasources
Examples:

    List datasources
    $ steam get datasources \
        --project-id=? \
        --offset=? \
        --limit=?

`
func getDatasources(c *context) *cobra.Command {
  
  
  var projectId int64
  var offset int64
  var limit int64
  cmd := newCmd(c, getDatasourcesHelp, func(c *context, args []string) {
    
    
  datasources, err := c.remote.GetDatasources(
  
    projectId,
  
    offset,
  
    limit,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", datasources)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&projectId, "project-id", projectId, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  return cmd
}

var getEngineHelp = `
engine [?]
Get Engine
Examples:

    Get engine details
    $ steam get engine \
        --engine-id=?

`
func getEngine(c *context) *cobra.Command {
  
  
  var engineId int64
  cmd := newCmd(c, getEngineHelp, func(c *context, args []string) {
    
    
  engine, err := c.remote.GetEngine(
  
    engineId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", engine)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&engineId, "engine-id", engineId, "No description available")
  return cmd
}

var getEnginesHelp = `
engines [?]
Get Engines
Examples:

    List engines
    $ steam get engines

`
func getEngines(c *context) *cobra.Command {
  
  
  cmd := newCmd(c, getEnginesHelp, func(c *context, args []string) {
    
    
  engines, err := c.remote.GetEngines(
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", engines)
  
  return

    
  })
  
  
  return cmd
}

var getHistoryHelp = `
history [?]
Get History
Examples:

    List audit trail records for an entity
    $ steam get history \
        --entity-type-id=? \
        --entity-id=? \
        --offset=? \
        --limit=?

`
func getHistory(c *context) *cobra.Command {
  
  
  var entityId int64
  var offset int64
  var limit int64
  var entityTypeId int64
  cmd := newCmd(c, getHistoryHelp, func(c *context, args []string) {
    
    
  history, err := c.remote.GetHistory(
  
    entityTypeId,
  
    entityId,
  
    offset,
  
    limit,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", history)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&entityId, "entity-id", entityId, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&entityTypeId, "entity-type-id", entityTypeId, "No description available")
  return cmd
}

var getIdentitiesHelp = `
identities [?]
Get Identities
Examples:

    List identities
    $ steam get identities \
        --offset=? \
        --limit=?

    List identities for a workgroup
    $ steam get identities --for-workgroup \
        --workgroup-id=?

    List identities for a role
    $ steam get identities --for-role \
        --role-id=?

`
func getIdentities(c *context) *cobra.Command {
  var forWorkgroup bool
  var forRole bool
  
  
  var roleId int64
  var offset int64
  var limit int64
  var workgroupId int64
  cmd := newCmd(c, getIdentitiesHelp, func(c *context, args []string) {
    
    
    
    
    
    if forWorkgroup {
      
  identities, err := c.remote.GetIdentitiesForWorkgroup(
  
    workgroupId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", identities)
  
  return

    }
    
    
    
    if forRole {
      
  identities, err := c.remote.GetIdentitiesForRole(
  
    roleId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", identities)
  
  return

    }
    
    
    
    
    if true {
      
  identities, err := c.remote.GetIdentities(
  
    offset,
  
    limit,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", identities)
  
  return

    }
    
    
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&forWorkgroup, "for-workgroup", forWorkgroup, "List identities for a workgroup")
  cmd.Flags().BoolVar(&forRole, "for-role", forRole, "List identities for a role")
  
  
  cmd.Flags().Int64Var(&roleId, "role-id", roleId, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&workgroupId, "workgroup-id", workgroupId, "No description available")
  return cmd
}

var getIdentityHelp = `
identity [?]
Get Identity
Examples:

    Get identity details
    $ steam get identity \
        --identity-id=?

    Get identity details by name
    $ steam get identity --by-name \
        --name=?

`
func getIdentity(c *context) *cobra.Command {
  var byName bool
  
  
  var identityId int64
  var name string
  cmd := newCmd(c, getIdentityHelp, func(c *context, args []string) {
    
    
    
    
    
    if byName {
      
  identity, err := c.remote.GetIdentityByName(
  
    name,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", identity)
  
  return

    }
    
    
    
    
    if true {
      
  identity, err := c.remote.GetIdentity(
  
    identityId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", identity)
  
  return

    }
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&byName, "by-name", byName, "Get identity details by name")
  
  
  cmd.Flags().Int64Var(&identityId, "identity-id", identityId, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  return cmd
}

var getJobHelp = `
job [?]
Get Job
Examples:

    Get job details
    $ steam get job \
        --cluster-id=? \
        --job-name=?

`
func getJob(c *context) *cobra.Command {
  
  
  var clusterId int64
  var jobName string
  cmd := newCmd(c, getJobHelp, func(c *context, args []string) {
    
    
  job, err := c.remote.GetJob(
  
    clusterId,
  
    jobName,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", job)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  cmd.Flags().StringVar(&jobName, "job-name", jobName, "No description available")
  return cmd
}

var getJobsHelp = `
jobs [?]
Get Jobs
Examples:

    List jobs
    $ steam get jobs \
        --cluster-id=?

`
func getJobs(c *context) *cobra.Command {
  
  
  var clusterId int64
  cmd := newCmd(c, getJobsHelp, func(c *context, args []string) {
    
    
  jobs, err := c.remote.GetJobs(
  
    clusterId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", jobs)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  return cmd
}

var getModelHelp = `
model [?]
Get Model
Examples:

    Get model details
    $ steam get model \
        --model-id=?

`
func getModel(c *context) *cobra.Command {
  
  
  var modelId int64
  cmd := newCmd(c, getModelHelp, func(c *context, args []string) {
    
    
  model, err := c.remote.GetModel(
  
    modelId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", model)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&modelId, "model-id", modelId, "No description available")
  return cmd
}

var getModelsHelp = `
models [?]
Get Models
Examples:

    List models
    $ steam get models \
        --project-id=? \
        --offset=? \
        --limit=?

    List models from a cluster
    $ steam get models --from-cluster \
        --cluster-id=?

`
func getModels(c *context) *cobra.Command {
  var fromCluster bool
  
  
  var projectId int64
  var offset int64
  var limit int64
  var clusterId int64
  cmd := newCmd(c, getModelsHelp, func(c *context, args []string) {
    
    
    
    
    
    if fromCluster {
      
  models, err := c.remote.GetModelsFromCluster(
  
    clusterId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", models)
  
  return

    }
    
    
    
    
    if true {
      
  models, err := c.remote.GetModels(
  
    projectId,
  
    offset,
  
    limit,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", models)
  
  return

    }
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&fromCluster, "from-cluster", fromCluster, "List models from a cluster")
  
  
  cmd.Flags().Int64Var(&projectId, "project-id", projectId, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  return cmd
}

var getPermissionsHelp = `
permissions [?]
Get Permissions
Examples:

    List permissions for a role
    $ steam get permissions --for-role \
        --role-id=?

    List permissions for an identity
    $ steam get permissions --for-identity \
        --identity-id=?

`
func getPermissions(c *context) *cobra.Command {
  var forRole bool
  var forIdentity bool
  
  
  var roleId int64
  var identityId int64
  cmd := newCmd(c, getPermissionsHelp, func(c *context, args []string) {
    
    
    
    if forRole {
      
  permissions, err := c.remote.GetPermissionsForRole(
  
    roleId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", permissions)
  
  return

    }
    
    
    
    if forIdentity {
      
  permissions, err := c.remote.GetPermissionsForIdentity(
  
    identityId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", permissions)
  
  return

    }
    
    
    
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&forRole, "for-role", forRole, "List permissions for a role")
  cmd.Flags().BoolVar(&forIdentity, "for-identity", forIdentity, "List permissions for an identity")
  
  
  cmd.Flags().Int64Var(&roleId, "role-id", roleId, "No description available")
  cmd.Flags().Int64Var(&identityId, "identity-id", identityId, "No description available")
  return cmd
}

var getPrivilegesHelp = `
privileges [?]
Get Privileges
Examples:

    List privileges for an entity
    $ steam get privileges \
        --entity-type-id=? \
        --entity-id=?

`
func getPrivileges(c *context) *cobra.Command {
  
  
  var entityTypeId int64
  var entityId int64
  cmd := newCmd(c, getPrivilegesHelp, func(c *context, args []string) {
    
    
  privileges, err := c.remote.GetPrivileges(
  
    entityTypeId,
  
    entityId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", privileges)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&entityTypeId, "entity-type-id", entityTypeId, "No description available")
  cmd.Flags().Int64Var(&entityId, "entity-id", entityId, "No description available")
  return cmd
}

var getProjectHelp = `
project [?]
Get Project
Examples:

    Get project details
    $ steam get project \
        --project-id=?

`
func getProject(c *context) *cobra.Command {
  
  
  var projectId int64
  cmd := newCmd(c, getProjectHelp, func(c *context, args []string) {
    
    
  project, err := c.remote.GetProject(
  
    projectId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", project)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&projectId, "project-id", projectId, "No description available")
  return cmd
}

var getProjectsHelp = `
projects [?]
Get Projects
Examples:

    List projects
    $ steam get projects \
        --offset=? \
        --limit=?

`
func getProjects(c *context) *cobra.Command {
  
  
  var offset int64
  var limit int64
  cmd := newCmd(c, getProjectsHelp, func(c *context, args []string) {
    
    
  projects, err := c.remote.GetProjects(
  
    offset,
  
    limit,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", projects)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  return cmd
}

var getRoleHelp = `
role [?]
Get Role
Examples:

    Get role details
    $ steam get role \
        --role-id=?

    Get role details by name
    $ steam get role --by-name \
        --name=?

`
func getRole(c *context) *cobra.Command {
  var byName bool
  
  
  var roleId int64
  var name string
  cmd := newCmd(c, getRoleHelp, func(c *context, args []string) {
    
    
    
    
    
    if byName {
      
  role, err := c.remote.GetRoleByName(
  
    name,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", role)
  
  return

    }
    
    
    
    
    if true {
      
  role, err := c.remote.GetRole(
  
    roleId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", role)
  
  return

    }
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&byName, "by-name", byName, "Get role details by name")
  
  
  cmd.Flags().Int64Var(&roleId, "role-id", roleId, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  return cmd
}

var getRolesHelp = `
roles [?]
Get Roles
Examples:

    List roles
    $ steam get roles \
        --offset=? \
        --limit=?

    List roles for an identity
    $ steam get roles --for-identity \
        --identity-id=?

`
func getRoles(c *context) *cobra.Command {
  var forIdentity bool
  
  
  var limit int64
  var identityId int64
  var offset int64
  cmd := newCmd(c, getRolesHelp, func(c *context, args []string) {
    
    
    
    
    
    if forIdentity {
      
  roles, err := c.remote.GetRolesForIdentity(
  
    identityId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", roles)
  
  return

    }
    
    
    
    
    if true {
      
  roles, err := c.remote.GetRoles(
  
    offset,
  
    limit,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", roles)
  
  return

    }
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&forIdentity, "for-identity", forIdentity, "List roles for an identity")
  
  
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&identityId, "identity-id", identityId, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
  return cmd
}

var getServiceHelp = `
service [?]
Get Service
Examples:

    Get service details
    $ steam get service \
        --service-id=?

`
func getService(c *context) *cobra.Command {
  
  
  var serviceId int64
  cmd := newCmd(c, getServiceHelp, func(c *context, args []string) {
    
    
  service, err := c.remote.GetService(
  
    serviceId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", service)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&serviceId, "service-id", serviceId, "No description available")
  return cmd
}

var getServicesHelp = `
services [?]
Get Services
Examples:

    List services
    $ steam get services \
        --offset=? \
        --limit=?

    List services for a model
    $ steam get services --for-model \
        --model-id=? \
        --offset=? \
        --limit=?

`
func getServices(c *context) *cobra.Command {
  var forModel bool
  
  
  var offset int64
  var limit int64
  var modelId int64
  cmd := newCmd(c, getServicesHelp, func(c *context, args []string) {
    
    
    
    
    
    if forModel {
      
  services, err := c.remote.GetServicesForModel(
  
    modelId,
  
    offset,
  
    limit,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", services)
  
  return

    }
    
    
    
    
    if true {
      
  services, err := c.remote.GetServices(
  
    offset,
  
    limit,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", services)
  
  return

    }
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&forModel, "for-model", forModel, "List services for a model")
  
  
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&modelId, "model-id", modelId, "No description available")
  return cmd
}

var getWorkgroupHelp = `
workgroup [?]
Get Workgroup
Examples:

    Get workgroup details
    $ steam get workgroup \
        --workgroup-id=?

    Get workgroup details by name
    $ steam get workgroup --by-name \
        --name=?

`
func getWorkgroup(c *context) *cobra.Command {
  var byName bool
  
  
  var name string
  var workgroupId int64
  cmd := newCmd(c, getWorkgroupHelp, func(c *context, args []string) {
    
    
    
    
    
    if byName {
      
  workgroup, err := c.remote.GetWorkgroupByName(
  
    name,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", workgroup)
  
  return

    }
    
    
    
    
    if true {
      
  workgroup, err := c.remote.GetWorkgroup(
  
    workgroupId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", workgroup)
  
  return

    }
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&byName, "by-name", byName, "Get workgroup details by name")
  
  
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().Int64Var(&workgroupId, "workgroup-id", workgroupId, "No description available")
  return cmd
}

var getWorkgroupsHelp = `
workgroups [?]
Get Workgroups
Examples:

    List workgroups
    $ steam get workgroups \
        --offset=? \
        --limit=?

    List workgroups for an identity
    $ steam get workgroups --for-identity \
        --identity-id=?

`
func getWorkgroups(c *context) *cobra.Command {
  var forIdentity bool
  
  
  var offset int64
  var limit int64
  var identityId int64
  cmd := newCmd(c, getWorkgroupsHelp, func(c *context, args []string) {
    
    
    
    
    
    if forIdentity {
      
  workgroups, err := c.remote.GetWorkgroupsForIdentity(
  
    identityId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", workgroups)
  
  return

    }
    
    
    
    
    if true {
      
  workgroups, err := c.remote.GetWorkgroups(
  
    offset,
  
    limit,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", workgroups)
  
  return

    }
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&forIdentity, "for-identity", forIdentity, "List workgroups for an identity")
  
  
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&identityId, "identity-id", identityId, "No description available")
  return cmd
}




var importHelp = `
import [?]
Import entities
Commands:

    $ steam import model ...
`
func import_(c *context) *cobra.Command {
  cmd := newCmd(c, importHelp, nil)

  cmd.AddCommand(importModel(c))
  return cmd
}


var importModelHelp = `
model [?]
Import Model
Examples:

    Import models from a cluster
    $ steam import model --from-cluster \
        --cluster-id=? \
        --project-id=? \
        --model-name=?

`
func importModel(c *context) *cobra.Command {
  var fromCluster bool
  
  
  var clusterId int64
  var projectId int64
  var modelName string
  cmd := newCmd(c, importModelHelp, func(c *context, args []string) {
    
    
    
    if fromCluster {
      
  model, err := c.remote.ImportModelFromCluster(
  
    clusterId,
  
    projectId,
  
    modelName,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", model)
  
  return

    }
    
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&fromCluster, "from-cluster", fromCluster, "Import models from a cluster")
  
  
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  cmd.Flags().Int64Var(&projectId, "project-id", projectId, "No description available")
  cmd.Flags().StringVar(&modelName, "model-name", modelName, "No description available")
  return cmd
}




var linkHelp = `
link [?]
Link entities
Commands:

    $ steam link identity ...
`
func link(c *context) *cobra.Command {
  cmd := newCmd(c, linkHelp, nil)

  cmd.AddCommand(linkIdentity(c))
  return cmd
}


var linkIdentityHelp = `
identity [?]
Link Identity
Examples:

    Link an identity with a workgroup
    $ steam link identity --with-workgroup \
        --identity-id=? \
        --workgroup-id=?

    Link an identity with a role
    $ steam link identity --with-role \
        --identity-id=? \
        --role-id=?

`
func linkIdentity(c *context) *cobra.Command {
  var withWorkgroup bool
  var withRole bool
  
  
  var identityId int64
  var workgroupId int64
  var roleId int64
  cmd := newCmd(c, linkIdentityHelp, func(c *context, args []string) {
    
    
    
    if withWorkgroup {
      
  err := c.remote.LinkIdentityWithWorkgroup(
  
    identityId,
  
    workgroupId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    }
    
    
    
    if withRole {
      
  err := c.remote.LinkIdentityWithRole(
  
    identityId,
  
    roleId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    }
    
    
    
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&withWorkgroup, "with-workgroup", withWorkgroup, "Link an identity with a workgroup")
  cmd.Flags().BoolVar(&withRole, "with-role", withRole, "Link an identity with a role")
  
  
  cmd.Flags().Int64Var(&identityId, "identity-id", identityId, "No description available")
  cmd.Flags().Int64Var(&workgroupId, "workgroup-id", workgroupId, "No description available")
  cmd.Flags().Int64Var(&roleId, "role-id", roleId, "No description available")
  return cmd
}




var pingHelp = `
ping [?]
Ping entities
Commands:

    $ steam ping server ...
`
func ping(c *context) *cobra.Command {
  cmd := newCmd(c, pingHelp, nil)

  cmd.AddCommand(pingServer(c))
  return cmd
}


var pingServerHelp = `
server [?]
Ping Server
Examples:

    Ping the Steam server
    $ steam ping server \
        --input=?

`
func pingServer(c *context) *cobra.Command {
  
  
  var input string
  cmd := newCmd(c, pingServerHelp, func(c *context, args []string) {
    
    
  output, err := c.remote.PingServer(
  
    input,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", output)
  
  return

    
  })
  
  
  cmd.Flags().StringVar(&input, "input", input, "Message to send")
  return cmd
}




var registerHelp = `
register [?]
Register entities
Commands:

    $ steam register cluster ...
`
func register(c *context) *cobra.Command {
  cmd := newCmd(c, registerHelp, nil)

  cmd.AddCommand(registerCluster(c))
  return cmd
}


var registerClusterHelp = `
cluster [?]
Register Cluster
Examples:

    Connect to a cluster
    $ steam register cluster \
        --address=?

`
func registerCluster(c *context) *cobra.Command {
  
  
  var address string
  cmd := newCmd(c, registerClusterHelp, func(c *context, args []string) {
    
    
  clusterId, err := c.remote.RegisterCluster(
  
    address,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", clusterId)
  
  return

    
  })
  
  
  cmd.Flags().StringVar(&address, "address", address, "No description available")
  return cmd
}




var shareHelp = `
share [?]
Share entities
Commands:

    $ steam share entity ...
`
func share(c *context) *cobra.Command {
  cmd := newCmd(c, shareHelp, nil)

  cmd.AddCommand(shareEntity(c))
  return cmd
}


var shareEntityHelp = `
entity [?]
Share Entity
Examples:

    Share an entity with a workgroup
    $ steam share entity \
        --kind=? \
        --workgroup-id=? \
        --entity-type-id=? \
        --entity-id=?

`
func shareEntity(c *context) *cobra.Command {
  
  
  var entityTypeId int64
  var entityId int64
  var kind string
  var workgroupId int64
  cmd := newCmd(c, shareEntityHelp, func(c *context, args []string) {
    
    
  err := c.remote.ShareEntity(
  
    kind,
  
    workgroupId,
  
    entityTypeId,
  
    entityId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&entityTypeId, "entity-type-id", entityTypeId, "No description available")
  cmd.Flags().Int64Var(&entityId, "entity-id", entityId, "No description available")
  cmd.Flags().StringVar(&kind, "kind", kind, "No description available")
  cmd.Flags().Int64Var(&workgroupId, "workgroup-id", workgroupId, "No description available")
  return cmd
}




var splitHelp = `
split [?]
Split entities
Commands:

    $ steam split dataset ...
`
func split(c *context) *cobra.Command {
  cmd := newCmd(c, splitHelp, nil)

  cmd.AddCommand(splitDataset(c))
  return cmd
}


var splitDatasetHelp = `
dataset [?]
Split Dataset
Examples:

    Split a dataset
    $ steam split dataset \
        --dataset-id=? \
        --ratio1=? \
        --ratio2=?

`
func splitDataset(c *context) *cobra.Command {
  
  
  var datasetId int64
  var ratio1 int
  var ratio2 int
  cmd := newCmd(c, splitDatasetHelp, func(c *context, args []string) {
    
    
  datasetIds, err := c.remote.SplitDataset(
  
    datasetId,
  
    ratio1,
  
    ratio2,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", datasetIds)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&datasetId, "dataset-id", datasetId, "No description available")
  cmd.Flags().IntVar(&ratio1, "ratio1", ratio1, "No description available")
  cmd.Flags().IntVar(&ratio2, "ratio2", ratio2, "No description available")
  return cmd
}




var startHelp = `
start [?]
Start entities
Commands:

    $ steam start cluster ...
    $ steam start service ...
`
func start(c *context) *cobra.Command {
  cmd := newCmd(c, startHelp, nil)

  cmd.AddCommand(startCluster(c))
  cmd.AddCommand(startService(c))
  return cmd
}


var startClusterHelp = `
cluster [?]
Start Cluster
Examples:

    Start a cluster using Yarn
    $ steam start cluster --on-yarn \
        --cluster-name=? \
        --engine-id=? \
        --size=? \
        --memory=? \
        --username=?

`
func startCluster(c *context) *cobra.Command {
  var onYarn bool
  
  
  var engineId int64
  var size int
  var memory string
  var username string
  var clusterName string
  cmd := newCmd(c, startClusterHelp, func(c *context, args []string) {
    
    
    
    if onYarn {
      
  clusterId, err := c.remote.StartClusterOnYarn(
  
    clusterName,
  
    engineId,
  
    size,
  
    memory,
  
    username,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", clusterId)
  
  return

    }
    
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&onYarn, "on-yarn", onYarn, "Start a cluster using Yarn")
  
  
  cmd.Flags().Int64Var(&engineId, "engine-id", engineId, "No description available")
  cmd.Flags().IntVar(&size, "size", size, "No description available")
  cmd.Flags().StringVar(&memory, "memory", memory, "No description available")
  cmd.Flags().StringVar(&username, "username", username, "No description available")
  cmd.Flags().StringVar(&clusterName, "cluster-name", clusterName, "No description available")
  return cmd
}

var startServiceHelp = `
service [?]
Start Service
Examples:

    Start a service
    $ steam start service \
        --model-id=? \
        --port=?

`
func startService(c *context) *cobra.Command {
  
  
  var modelId int64
  var port int
  cmd := newCmd(c, startServiceHelp, func(c *context, args []string) {
    
    
  service, err := c.remote.StartService(
  
    modelId,
  
    port,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  fmt.Printf("%+v\n", service)
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&modelId, "model-id", modelId, "No description available")
  cmd.Flags().IntVar(&port, "port", port, "No description available")
  return cmd
}




var stopHelp = `
stop [?]
Stop entities
Commands:

    $ steam stop cluster ...
    $ steam stop service ...
`
func stop(c *context) *cobra.Command {
  cmd := newCmd(c, stopHelp, nil)

  cmd.AddCommand(stopCluster(c))
  cmd.AddCommand(stopService(c))
  return cmd
}


var stopClusterHelp = `
cluster [?]
Stop Cluster
Examples:

    Stop a cluster using Yarn
    $ steam stop cluster --on-yarn \
        --cluster-id=?

`
func stopCluster(c *context) *cobra.Command {
  var onYarn bool
  
  
  var clusterId int64
  cmd := newCmd(c, stopClusterHelp, func(c *context, args []string) {
    
    
    
    if onYarn {
      
  err := c.remote.StopClusterOnYarn(
  
    clusterId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    }
    
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&onYarn, "on-yarn", onYarn, "Stop a cluster using Yarn")
  
  
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  return cmd
}

var stopServiceHelp = `
service [?]
Stop Service
Examples:

    Stop a service
    $ steam stop service \
        --service-id=?

`
func stopService(c *context) *cobra.Command {
  
  
  var serviceId int64
  cmd := newCmd(c, stopServiceHelp, func(c *context, args []string) {
    
    
  err := c.remote.StopService(
  
    serviceId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&serviceId, "service-id", serviceId, "No description available")
  return cmd
}




var unlinkHelp = `
unlink [?]
Unlink entities
Commands:

    $ steam unlink identity ...
`
func unlink(c *context) *cobra.Command {
  cmd := newCmd(c, unlinkHelp, nil)

  cmd.AddCommand(unlinkIdentity(c))
  return cmd
}


var unlinkIdentityHelp = `
identity [?]
Unlink Identity
Examples:

    Unlink an identity from a workgroup
    $ steam unlink identity --from-workgroup \
        --identity-id=? \
        --workgroup-id=?

    Unlink an identity from a role
    $ steam unlink identity --from-role \
        --identity-id=? \
        --role-id=?

`
func unlinkIdentity(c *context) *cobra.Command {
  var fromWorkgroup bool
  var fromRole bool
  
  
  var identityId int64
  var workgroupId int64
  var roleId int64
  cmd := newCmd(c, unlinkIdentityHelp, func(c *context, args []string) {
    
    
    
    if fromWorkgroup {
      
  err := c.remote.UnlinkIdentityFromWorkgroup(
  
    identityId,
  
    workgroupId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    }
    
    
    
    if fromRole {
      
  err := c.remote.UnlinkIdentityFromRole(
  
    identityId,
  
    roleId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    }
    
    
    
    
    
    
    
    
  })
  cmd.Flags().BoolVar(&fromWorkgroup, "from-workgroup", fromWorkgroup, "Unlink an identity from a workgroup")
  cmd.Flags().BoolVar(&fromRole, "from-role", fromRole, "Unlink an identity from a role")
  
  
  cmd.Flags().Int64Var(&identityId, "identity-id", identityId, "No description available")
  cmd.Flags().Int64Var(&workgroupId, "workgroup-id", workgroupId, "No description available")
  cmd.Flags().Int64Var(&roleId, "role-id", roleId, "No description available")
  return cmd
}




var unregisterHelp = `
unregister [?]
Unregister entities
Commands:

    $ steam unregister cluster ...
`
func unregister(c *context) *cobra.Command {
  cmd := newCmd(c, unregisterHelp, nil)

  cmd.AddCommand(unregisterCluster(c))
  return cmd
}


var unregisterClusterHelp = `
cluster [?]
Unregister Cluster
Examples:

    Disconnect from a cluster
    $ steam unregister cluster \
        --cluster-id=?

`
func unregisterCluster(c *context) *cobra.Command {
  
  
  var clusterId int64
  cmd := newCmd(c, unregisterClusterHelp, func(c *context, args []string) {
    
    
  err := c.remote.UnregisterCluster(
  
    clusterId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  return cmd
}




var unshareHelp = `
unshare [?]
Unshare entities
Commands:

    $ steam unshare entity ...
`
func unshare(c *context) *cobra.Command {
  cmd := newCmd(c, unshareHelp, nil)

  cmd.AddCommand(unshareEntity(c))
  return cmd
}


var unshareEntityHelp = `
entity [?]
Unshare Entity
Examples:

    Unshare an entity
    $ steam unshare entity \
        --kind=? \
        --workgroup-id=? \
        --entity-type-id=? \
        --entity-id=?

`
func unshareEntity(c *context) *cobra.Command {
  
  
  var kind string
  var workgroupId int64
  var entityTypeId int64
  var entityId int64
  cmd := newCmd(c, unshareEntityHelp, func(c *context, args []string) {
    
    
  err := c.remote.UnshareEntity(
  
    kind,
  
    workgroupId,
  
    entityTypeId,
  
    entityId,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().StringVar(&kind, "kind", kind, "No description available")
  cmd.Flags().Int64Var(&workgroupId, "workgroup-id", workgroupId, "No description available")
  cmd.Flags().Int64Var(&entityTypeId, "entity-type-id", entityTypeId, "No description available")
  cmd.Flags().Int64Var(&entityId, "entity-id", entityId, "No description available")
  return cmd
}




var updateHelp = `
update [?]
Update entities
Commands:

    $ steam update dataset ...
    $ steam update datasource ...
    $ steam update identity ...
    $ steam update role ...
    $ steam update workgroup ...
`
func update(c *context) *cobra.Command {
  cmd := newCmd(c, updateHelp, nil)

  cmd.AddCommand(updateDataset(c))
  cmd.AddCommand(updateDatasource(c))
  cmd.AddCommand(updateIdentity(c))
  cmd.AddCommand(updateRole(c))
  cmd.AddCommand(updateWorkgroup(c))
  return cmd
}


var updateDatasetHelp = `
dataset [?]
Update Dataset
Examples:

    Update a dataset
    $ steam update dataset \
        --dataset-id=? \
        --name=? \
        --description=? \
        --response-column-name=?

`
func updateDataset(c *context) *cobra.Command {
  
  
  var datasetId int64
  var name string
  var description string
  var responseColumnName string
  cmd := newCmd(c, updateDatasetHelp, func(c *context, args []string) {
    
    
  err := c.remote.UpdateDataset(
  
    datasetId,
  
    name,
  
    description,
  
    responseColumnName,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&datasetId, "dataset-id", datasetId, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().StringVar(&responseColumnName, "response-column-name", responseColumnName, "No description available")
  return cmd
}

var updateDatasourceHelp = `
datasource [?]
Update Datasource
Examples:

    Update a datasource
    $ steam update datasource \
        --datasource-id=? \
        --name=? \
        --description=? \
        --path=?

`
func updateDatasource(c *context) *cobra.Command {
  
  
  var datasourceId int64
  var name string
  var description string
  var path string
  cmd := newCmd(c, updateDatasourceHelp, func(c *context, args []string) {
    
    
  err := c.remote.UpdateDatasource(
  
    datasourceId,
  
    name,
  
    description,
  
    path,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&datasourceId, "datasource-id", datasourceId, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().StringVar(&path, "path", path, "No description available")
  return cmd
}

var updateIdentityHelp = `
identity [?]
Update Identity
Examples:

    Update an identity
    $ steam update identity \
        --identity-id=? \
        --password=?

`
func updateIdentity(c *context) *cobra.Command {
  
  
  var identityId int64
  var password string
  cmd := newCmd(c, updateIdentityHelp, func(c *context, args []string) {
    
    
  err := c.remote.UpdateIdentity(
  
    identityId,
  
    password,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&identityId, "identity-id", identityId, "No description available")
  cmd.Flags().StringVar(&password, "password", password, "No description available")
  return cmd
}

var updateRoleHelp = `
role [?]
Update Role
Examples:

    Update a role
    $ steam update role \
        --role-id=? \
        --name=? \
        --description=?

`
func updateRole(c *context) *cobra.Command {
  
  
  var roleId int64
  var name string
  var description string
  cmd := newCmd(c, updateRoleHelp, func(c *context, args []string) {
    
    
  err := c.remote.UpdateRole(
  
    roleId,
  
    name,
  
    description,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&roleId, "role-id", roleId, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  return cmd
}

var updateWorkgroupHelp = `
workgroup [?]
Update Workgroup
Examples:

    Update a workgroup
    $ steam update workgroup \
        --workgroup-id=? \
        --name=? \
        --description=?

`
func updateWorkgroup(c *context) *cobra.Command {
  
  
  var workgroupId int64
  var name string
  var description string
  cmd := newCmd(c, updateWorkgroupHelp, func(c *context, args []string) {
    
    
  err := c.remote.UpdateWorkgroup(
  
    workgroupId,
  
    name,
  
    description,
  )
  if err != nil {
    log.Fatalln(err)
  }
  
  return

    
  })
  
  
  cmd.Flags().Int64Var(&workgroupId, "workgroup-id", workgroupId, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  return cmd
}



