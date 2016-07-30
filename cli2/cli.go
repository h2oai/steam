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

// registerGeneratedCommands appends generated commands to the
//   supplied Cobra context.
func registerGeneratedCommands(c *context, cmd *cobra.Command) {
  cmd.AddCommand(
    add(c),
    build(c),
    create(c),
    deactivate(c),
    delete_(c),
    find(c),
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
  var engineName string // No description available
  var enginePath string // No description available

  cmd := newCmd(c, addEngineHelp, func(c *context, args []string) {
    
      // Add an engine
      engineId, err := c.remote.AddEngine(
        engineName, // No description available
        enginePath, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("EngineId:\t%v\n", engineId)
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
  var auto bool // Switch for BuildModelAuto()
  var algorithm string // No description available
  var clusterId int64 // No description available
  var dataset string // No description available
  var datasetId int64 // No description available
  var maxRunTime int // No description available
  var targetName string // No description available

  cmd := newCmd(c, buildModelHelp, func(c *context, args []string) {
    if auto { // BuildModelAuto
      
      // Build an AutoML model
      model, err := c.remote.BuildModelAuto(
        clusterId, // No description available
        dataset, // No description available
        targetName, // No description available
        maxRunTime, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", model.Id), // No description available
        fmt.Sprintf("TrainingDatasetId:\t%v\t", model.TrainingDatasetId), // No description available
        fmt.Sprintf("ValidationDatasetId:\t%v\t", model.ValidationDatasetId), // No description available
        fmt.Sprintf("Name:\t%v\t", model.Name), // No description available
        fmt.Sprintf("ClusterName:\t%v\t", model.ClusterName), // No description available
        fmt.Sprintf("ModelKey:\t%v\t", model.ModelKey), // No description available
        fmt.Sprintf("Algorithm:\t%v\t", model.Algorithm), // No description available
        fmt.Sprintf("ModelCategory:\t%v\t", model.ModelCategory), // No description available
        fmt.Sprintf("DatasetName:\t%v\t", model.DatasetName), // No description available
        fmt.Sprintf("ResponseColumnName:\t%v\t", model.ResponseColumnName), // No description available
        fmt.Sprintf("LogicalName:\t%v\t", model.LogicalName), // No description available
        fmt.Sprintf("Location:\t%v\t", model.Location), // No description available
        fmt.Sprintf("MaxRuntime:\t%v\t", model.MaxRuntime), // No description available
        fmt.Sprintf("Metrics:\t%v\t", model.Metrics), // No description available
        fmt.Sprintf("CreatedAt:\t%v\t", model.CreatedAt), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
      return
    }
    if true { // default
      
      // Build a model
      modelId, err := c.remote.BuildModel(
        clusterId, // No description available
        datasetId, // No description available
        algorithm, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("ModelId:\t%v\n", modelId)
      return
    }
  })
  cmd.Flags().BoolVar(&auto, "auto", auto, "Build an AutoML model")
  
  
  cmd.Flags().StringVar(&algorithm, "algorithm", algorithm, "No description available")
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  cmd.Flags().StringVar(&dataset, "dataset", dataset, "No description available")
  cmd.Flags().Int64Var(&datasetId, "dataset-id", datasetId, "No description available")
  cmd.Flags().IntVar(&maxRunTime, "max-run-time", maxRunTime, "No description available")
  cmd.Flags().StringVar(&targetName, "target-name", targetName, "No description available")
  return cmd
}




var createHelp = `
create [?]
Create entities
Commands:

    $ steam create dataset ...
    $ steam create datasource ...
    $ steam create identity ...
    $ steam create label ...
    $ steam create project ...
    $ steam create role ...
    $ steam create workgroup ...
`
func create(c *context) *cobra.Command {
  cmd := newCmd(c, createHelp, nil)

  cmd.AddCommand(createDataset(c))
  cmd.AddCommand(createDatasource(c))
  cmd.AddCommand(createIdentity(c))
  cmd.AddCommand(createLabel(c))
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
  var clusterId int64 // No description available
  var datasourceId int64 // No description available
  var description string // No description available
  var name string // No description available
  var responseColumnName string // No description available

  cmd := newCmd(c, createDatasetHelp, func(c *context, args []string) {
    
      // Create a dataset
      datasetId, err := c.remote.CreateDataset(
        clusterId, // No description available
        datasourceId, // No description available
        name, // No description available
        description, // No description available
        responseColumnName, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("DatasetId:\t%v\n", datasetId)
      return
  })
  
  
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  cmd.Flags().Int64Var(&datasourceId, "datasource-id", datasourceId, "No description available")
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
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
  var description string // No description available
  var name string // No description available
  var path string // No description available
  var projectId int64 // No description available

  cmd := newCmd(c, createDatasourceHelp, func(c *context, args []string) {
    
      // Create a datasource
      datasourceId, err := c.remote.CreateDatasource(
        projectId, // No description available
        name, // No description available
        description, // No description available
        path, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("DatasourceId:\t%v\n", datasourceId)
      return
  })
  
  
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().StringVar(&path, "path", path, "No description available")
  cmd.Flags().Int64Var(&projectId, "project-id", projectId, "No description available")
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
  var name string // No description available
  var password string // No description available

  cmd := newCmd(c, createIdentityHelp, func(c *context, args []string) {
    
      // Create an identity
      identityId, err := c.remote.CreateIdentity(
        name, // No description available
        password, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("IdentityId:\t%v\n", identityId)
      return
  })
  
  
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().StringVar(&password, "password", password, "No description available")
  return cmd
}

var createLabelHelp = `
label [?]
Create Label
Examples:

    Create a label
    $ steam create label \
        --project-id=? \
        --name=? \
        --description=?

`

func createLabel(c *context) *cobra.Command {
  var description string // No description available
  var name string // No description available
  var projectId int64 // No description available

  cmd := newCmd(c, createLabelHelp, func(c *context, args []string) {
    
      // Create a label
      labelId, err := c.remote.CreateLabel(
        projectId, // No description available
        name, // No description available
        description, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("LabelId:\t%v\n", labelId)
      return
  })
  
  
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().Int64Var(&projectId, "project-id", projectId, "No description available")
  return cmd
}

var createProjectHelp = `
project [?]
Create Project
Examples:

    Create a project
    $ steam create project \
        --name=? \
        --description=? \
        --model-category=?

`

func createProject(c *context) *cobra.Command {
  var description string // No description available
  var modelCategory string // No description available
  var name string // No description available

  cmd := newCmd(c, createProjectHelp, func(c *context, args []string) {
    
      // Create a project
      projectId, err := c.remote.CreateProject(
        name, // No description available
        description, // No description available
        modelCategory, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("ProjectId:\t%v\n", projectId)
      return
  })
  
  
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().StringVar(&modelCategory, "model-category", modelCategory, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
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
  var description string // No description available
  var name string // No description available

  cmd := newCmd(c, createRoleHelp, func(c *context, args []string) {
    
      // Create a role
      roleId, err := c.remote.CreateRole(
        name, // No description available
        description, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("RoleId:\t%v\n", roleId)
      return
  })
  
  
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
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
  var description string // No description available
  var name string // No description available

  cmd := newCmd(c, createWorkgroupHelp, func(c *context, args []string) {
    
      // Create a workgroup
      workgroupId, err := c.remote.CreateWorkgroup(
        name, // No description available
        description, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("WorkgroupId:\t%v\n", workgroupId)
      return
  })
  
  
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
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
  var identityId int64 // No description available

  cmd := newCmd(c, deactivateIdentityHelp, func(c *context, args []string) {
    
      // Deactivate an identity
      err := c.remote.DeactivateIdentity(
        identityId, // No description available
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
    $ steam delete label ...
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
  cmd.AddCommand(deleteLabel(c))
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
  var clusterId int64 // No description available

  cmd := newCmd(c, deleteClusterHelp, func(c *context, args []string) {
    
      // Delete a cluster
      err := c.remote.DeleteCluster(
        clusterId, // No description available
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
  var datasetId int64 // No description available

  cmd := newCmd(c, deleteDatasetHelp, func(c *context, args []string) {
    
      // Delete a dataset
      err := c.remote.DeleteDataset(
        datasetId, // No description available
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
  var datasourceId int64 // No description available

  cmd := newCmd(c, deleteDatasourceHelp, func(c *context, args []string) {
    
      // Delete a datasource
      err := c.remote.DeleteDatasource(
        datasourceId, // No description available
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
  var engineId int64 // No description available

  cmd := newCmd(c, deleteEngineHelp, func(c *context, args []string) {
    
      // Delete an engine
      err := c.remote.DeleteEngine(
        engineId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
  })
  
  
  cmd.Flags().Int64Var(&engineId, "engine-id", engineId, "No description available")
  return cmd
}

var deleteLabelHelp = `
label [?]
Delete Label
Examples:

    Delete a label
    $ steam delete label \
        --label-id=?

`

func deleteLabel(c *context) *cobra.Command {
  var labelId int64 // No description available

  cmd := newCmd(c, deleteLabelHelp, func(c *context, args []string) {
    
      // Delete a label
      err := c.remote.DeleteLabel(
        labelId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
  })
  
  
  cmd.Flags().Int64Var(&labelId, "label-id", labelId, "No description available")
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
  var modelId int64 // No description available

  cmd := newCmd(c, deleteModelHelp, func(c *context, args []string) {
    
      // Delete a model
      err := c.remote.DeleteModel(
        modelId, // No description available
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
  var projectId int64 // No description available

  cmd := newCmd(c, deleteProjectHelp, func(c *context, args []string) {
    
      // Delete a project
      err := c.remote.DeleteProject(
        projectId, // No description available
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
  var roleId int64 // No description available

  cmd := newCmd(c, deleteRoleHelp, func(c *context, args []string) {
    
      // Delete a role
      err := c.remote.DeleteRole(
        roleId, // No description available
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
  var serviceId int64 // No description available

  cmd := newCmd(c, deleteServiceHelp, func(c *context, args []string) {
    
      // Delete a service
      err := c.remote.DeleteService(
        serviceId, // No description available
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
  var workgroupId int64 // No description available

  cmd := newCmd(c, deleteWorkgroupHelp, func(c *context, args []string) {
    
      // Delete a workgroup
      err := c.remote.DeleteWorkgroup(
        workgroupId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
  })
  
  
  cmd.Flags().Int64Var(&workgroupId, "workgroup-id", workgroupId, "No description available")
  return cmd
}




var findHelp = `
find [?]
Find entities
Commands:

    $ steam find models ...
`
func find(c *context) *cobra.Command {
  cmd := newCmd(c, findHelp, nil)

  cmd.AddCommand(findModels(c))
  return cmd
}


var findModelsHelp = `
models [?]
Find Models
Examples:

    List binomial models
    $ steam find models --binomial \
        --project-id=? \
        --name-part=? \
        --sort-by=? \
        --ascending=? \
        --offset=? \
        --limit=?

    List multinomial models
    $ steam find models --multinomial \
        --project-id=? \
        --name-part=? \
        --sort-by=? \
        --ascending=? \
        --offset=? \
        --limit=?

    List regression models
    $ steam find models --regression \
        --project-id=? \
        --name-part=? \
        --sort-by=? \
        --ascending=? \
        --offset=? \
        --limit=?

`

func findModels(c *context) *cobra.Command {
  var binomial bool // Switch for FindModelsBinomial()
  var multinomial bool // Switch for FindModelsMultinomial()
  var regression bool // Switch for FindModelsRegression()
  var ascending bool // No description available
  var limit int64 // No description available
  var namePart string // No description available
  var offset int64 // No description available
  var projectId int64 // No description available
  var sortBy string // No description available

  cmd := newCmd(c, findModelsHelp, func(c *context, args []string) {
    if binomial { // FindModelsBinomial
      
      // List binomial models
      models, err := c.remote.FindModelsBinomial(
        projectId, // No description available
        namePart, // No description available
        sortBy, // No description available
        ascending, // No description available
        offset, // No description available
        limit, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(models))
      for i, e := range models {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.TrainingDatasetId, // No description available
          e.ValidationDatasetId, // No description available
          e.Name, // No description available
          e.ClusterName, // No description available
          e.ModelKey, // No description available
          e.Algorithm, // No description available
          e.ModelCategory, // No description available
          e.DatasetName, // No description available
          e.ResponseColumnName, // No description available
          e.LogicalName, // No description available
          e.Location, // No description available
          e.MaxRuntime, // No description available
          e.Metrics, // No description available
          e.CreatedAt, // No description available
          e.Mse, // No description available
          e.RSquared, // No description available
          e.Logloss, // No description available
          e.Auc, // No description available
          e.Gini, // No description available
        )
      }
      c.printt("Id\tTrainingDatasetId\tValidationDatasetId\tName\tClusterName\tModelKey\tAlgorithm\tModelCategory\tDatasetName\tResponseColumnName\tLogicalName\tLocation\tMaxRuntime\tMetrics\tCreatedAt\tMse\tRSquared\tLogloss\tAuc\tGini\t", lines)
      return
    }
    if multinomial { // FindModelsMultinomial
      
      // List multinomial models
      models, err := c.remote.FindModelsMultinomial(
        projectId, // No description available
        namePart, // No description available
        sortBy, // No description available
        ascending, // No description available
        offset, // No description available
        limit, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(models))
      for i, e := range models {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.TrainingDatasetId, // No description available
          e.ValidationDatasetId, // No description available
          e.Name, // No description available
          e.ClusterName, // No description available
          e.ModelKey, // No description available
          e.Algorithm, // No description available
          e.ModelCategory, // No description available
          e.DatasetName, // No description available
          e.ResponseColumnName, // No description available
          e.LogicalName, // No description available
          e.Location, // No description available
          e.MaxRuntime, // No description available
          e.Metrics, // No description available
          e.CreatedAt, // No description available
          e.Mse, // No description available
          e.RSquared, // No description available
          e.Logloss, // No description available
        )
      }
      c.printt("Id\tTrainingDatasetId\tValidationDatasetId\tName\tClusterName\tModelKey\tAlgorithm\tModelCategory\tDatasetName\tResponseColumnName\tLogicalName\tLocation\tMaxRuntime\tMetrics\tCreatedAt\tMse\tRSquared\tLogloss\t", lines)
      return
    }
    if regression { // FindModelsRegression
      
      // List regression models
      models, err := c.remote.FindModelsRegression(
        projectId, // No description available
        namePart, // No description available
        sortBy, // No description available
        ascending, // No description available
        offset, // No description available
        limit, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(models))
      for i, e := range models {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.TrainingDatasetId, // No description available
          e.ValidationDatasetId, // No description available
          e.Name, // No description available
          e.ClusterName, // No description available
          e.ModelKey, // No description available
          e.Algorithm, // No description available
          e.ModelCategory, // No description available
          e.DatasetName, // No description available
          e.ResponseColumnName, // No description available
          e.LogicalName, // No description available
          e.Location, // No description available
          e.MaxRuntime, // No description available
          e.Metrics, // No description available
          e.CreatedAt, // No description available
          e.Mse, // No description available
          e.RSquared, // No description available
          e.MeanResidualDeviance, // No description available
        )
      }
      c.printt("Id\tTrainingDatasetId\tValidationDatasetId\tName\tClusterName\tModelKey\tAlgorithm\tModelCategory\tDatasetName\tResponseColumnName\tLogicalName\tLocation\tMaxRuntime\tMetrics\tCreatedAt\tMse\tRSquared\tMeanResidualDeviance\t", lines)
      return
    }
  })
  cmd.Flags().BoolVar(&binomial, "binomial", binomial, "List binomial models")
  cmd.Flags().BoolVar(&multinomial, "multinomial", multinomial, "List multinomial models")
  cmd.Flags().BoolVar(&regression, "regression", regression, "List regression models")
  
  
  cmd.Flags().BoolVar(&ascending, "ascending", ascending, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().StringVar(&namePart, "name-part", namePart, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
  cmd.Flags().Int64Var(&projectId, "project-id", projectId, "No description available")
  cmd.Flags().StringVar(&sortBy, "sort-by", sortBy, "No description available")
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
    $ steam get labels ...
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
  cmd.AddCommand(getLabels(c))
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

    List sort criteria for a binomial models
    $ steam get all --binomial-sort-criteria

    List sort criteria for a multinomial models
    $ steam get all --multinomial-sort-criteria

    List sort criteria for a regression models
    $ steam get all --regression-sort-criteria

    List all entity types
    $ steam get all --entity-types

    List all permissions
    $ steam get all --permissions

    List all cluster types
    $ steam get all --cluster-types

`

func getAll(c *context) *cobra.Command {
  var binomialSortCriteria bool // Switch for GetAllBinomialSortCriteria()
  var multinomialSortCriteria bool // Switch for GetAllMultinomialSortCriteria()
  var regressionSortCriteria bool // Switch for GetAllRegressionSortCriteria()
  var entityTypes bool // Switch for GetAllEntityTypes()
  var permissions bool // Switch for GetAllPermissions()
  var clusterTypes bool // Switch for GetAllClusterTypes()

  cmd := newCmd(c, getAllHelp, func(c *context, args []string) {
    if binomialSortCriteria { // GetAllBinomialSortCriteria
      
      // List sort criteria for a binomial models
      criteria, err := c.remote.GetAllBinomialSortCriteria(
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("Criteria:\t%v\n", criteria)
      return
    }
    if multinomialSortCriteria { // GetAllMultinomialSortCriteria
      
      // List sort criteria for a multinomial models
      criteria, err := c.remote.GetAllMultinomialSortCriteria(
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("Criteria:\t%v\n", criteria)
      return
    }
    if regressionSortCriteria { // GetAllRegressionSortCriteria
      
      // List sort criteria for a regression models
      criteria, err := c.remote.GetAllRegressionSortCriteria(
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("Criteria:\t%v\n", criteria)
      return
    }
    if entityTypes { // GetAllEntityTypes
      
      // List all entity types
      entityTypes, err := c.remote.GetAllEntityTypes(
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(entityTypes))
      for i, e := range entityTypes {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t",
          e.Id, // No description available
          e.Name, // No description available
        )
      }
      c.printt("Id\tName\t", lines)
      return
    }
    if permissions { // GetAllPermissions
      
      // List all permissions
      permissions, err := c.remote.GetAllPermissions(
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(permissions))
      for i, e := range permissions {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t",
          e.Id, // No description available
          e.Code, // No description available
          e.Description, // No description available
        )
      }
      c.printt("Id\tCode\tDescription\t", lines)
      return
    }
    if clusterTypes { // GetAllClusterTypes
      
      // List all cluster types
      clusterTypes, err := c.remote.GetAllClusterTypes(
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(clusterTypes))
      for i, e := range clusterTypes {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t",
          e.Id, // No description available
          e.Name, // No description available
        )
      }
      c.printt("Id\tName\t", lines)
      return
    }
  })
  cmd.Flags().BoolVar(&binomialSortCriteria, "binomial-sort-criteria", binomialSortCriteria, "List sort criteria for a binomial models")
  cmd.Flags().BoolVar(&multinomialSortCriteria, "multinomial-sort-criteria", multinomialSortCriteria, "List sort criteria for a multinomial models")
  cmd.Flags().BoolVar(&regressionSortCriteria, "regression-sort-criteria", regressionSortCriteria, "List sort criteria for a regression models")
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
  var onYarn bool // Switch for GetClusterOnYarn()
  var status bool // Switch for GetClusterStatus()
  var clusterId int64 // No description available

  cmd := newCmd(c, getClusterHelp, func(c *context, args []string) {
    if onYarn { // GetClusterOnYarn
      
      // Get cluster details (Yarn only)
      cluster, err := c.remote.GetClusterOnYarn(
        clusterId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", cluster.Id), // No description available
        fmt.Sprintf("EngineId:\t%v\t", cluster.EngineId), // No description available
        fmt.Sprintf("Size:\t%v\t", cluster.Size), // No description available
        fmt.Sprintf("ApplicationId:\t%v\t", cluster.ApplicationId), // No description available
        fmt.Sprintf("Memory:\t%v\t", cluster.Memory), // No description available
        fmt.Sprintf("Username:\t%v\t", cluster.Username), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
      return
    }
    if status { // GetClusterStatus
      
      // Get cluster status
      clusterStatus, err := c.remote.GetClusterStatus(
        clusterId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Version:\t%v\t", clusterStatus.Version), // No description available
        fmt.Sprintf("Status:\t%v\t", clusterStatus.Status), // No description available
        fmt.Sprintf("MaxMemory:\t%v\t", clusterStatus.MaxMemory), // No description available
        fmt.Sprintf("TotalCpuCount:\t%v\t", clusterStatus.TotalCpuCount), // No description available
        fmt.Sprintf("TotalAllowedCpuCount:\t%v\t", clusterStatus.TotalAllowedCpuCount), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
      return
    }
    if true { // default
      
      // Get cluster details
      cluster, err := c.remote.GetCluster(
        clusterId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", cluster.Id), // No description available
        fmt.Sprintf("Name:\t%v\t", cluster.Name), // No description available
        fmt.Sprintf("TypeId:\t%v\t", cluster.TypeId), // No description available
        fmt.Sprintf("DetailId:\t%v\t", cluster.DetailId), // No description available
        fmt.Sprintf("Address:\t%v\t", cluster.Address), // No description available
        fmt.Sprintf("State:\t%v\t", cluster.State), // No description available
        fmt.Sprintf("CreatedAt:\t%v\t", cluster.CreatedAt), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
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
  var limit int64 // No description available
  var offset int64 // No description available

  cmd := newCmd(c, getClustersHelp, func(c *context, args []string) {
    
      // List clusters
      clusters, err := c.remote.GetClusters(
        offset, // No description available
        limit, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(clusters))
      for i, e := range clusters {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.Name, // No description available
          e.TypeId, // No description available
          e.DetailId, // No description available
          e.Address, // No description available
          e.State, // No description available
          e.CreatedAt, // No description available
        )
      }
      c.printt("Id\tName\tTypeId\tDetailId\tAddress\tState\tCreatedAt\t", lines)
      return
  })
  
  
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
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
  var datasetId int64 // No description available

  cmd := newCmd(c, getDatasetHelp, func(c *context, args []string) {
    
      // Get dataset details
      dataset, err := c.remote.GetDataset(
        datasetId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", dataset.Id), // No description available
        fmt.Sprintf("DatasourceId:\t%v\t", dataset.DatasourceId), // No description available
        fmt.Sprintf("Name:\t%v\t", dataset.Name), // No description available
        fmt.Sprintf("Description:\t%v\t", dataset.Description), // No description available
        fmt.Sprintf("FrameName:\t%v\t", dataset.FrameName), // No description available
        fmt.Sprintf("ResponseColumnName:\t%v\t", dataset.ResponseColumnName), // No description available
        fmt.Sprintf("Properties:\t%v\t", dataset.Properties), // No description available
        fmt.Sprintf("CreatedAt:\t%v\t", dataset.CreatedAt), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
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

    Get a list of datasets on a cluster
    $ steam get datasets --from-cluster \
        --cluster-id=?

`

func getDatasets(c *context) *cobra.Command {
  var fromCluster bool // Switch for GetDatasetsFromCluster()
  var clusterId int64 // No description available
  var datasourceId int64 // No description available
  var limit int64 // No description available
  var offset int64 // No description available

  cmd := newCmd(c, getDatasetsHelp, func(c *context, args []string) {
    if fromCluster { // GetDatasetsFromCluster
      
      // Get a list of datasets on a cluster
      dataset, err := c.remote.GetDatasetsFromCluster(
        clusterId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(dataset))
      for i, e := range dataset {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.DatasourceId, // No description available
          e.Name, // No description available
          e.Description, // No description available
          e.FrameName, // No description available
          e.ResponseColumnName, // No description available
          e.Properties, // No description available
          e.CreatedAt, // No description available
        )
      }
      c.printt("Id\tDatasourceId\tName\tDescription\tFrameName\tResponseColumnName\tProperties\tCreatedAt\t", lines)
      return
    }
    if true { // default
      
      // List datasets
      datasets, err := c.remote.GetDatasets(
        datasourceId, // No description available
        offset, // No description available
        limit, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(datasets))
      for i, e := range datasets {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.DatasourceId, // No description available
          e.Name, // No description available
          e.Description, // No description available
          e.FrameName, // No description available
          e.ResponseColumnName, // No description available
          e.Properties, // No description available
          e.CreatedAt, // No description available
        )
      }
      c.printt("Id\tDatasourceId\tName\tDescription\tFrameName\tResponseColumnName\tProperties\tCreatedAt\t", lines)
      return
    }
  })
  cmd.Flags().BoolVar(&fromCluster, "from-cluster", fromCluster, "Get a list of datasets on a cluster")
  
  
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  cmd.Flags().Int64Var(&datasourceId, "datasource-id", datasourceId, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
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
  var datasourceId int64 // No description available

  cmd := newCmd(c, getDatasourceHelp, func(c *context, args []string) {
    
      // Get datasource details
      datasource, err := c.remote.GetDatasource(
        datasourceId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", datasource.Id), // No description available
        fmt.Sprintf("ProjectId:\t%v\t", datasource.ProjectId), // No description available
        fmt.Sprintf("Name:\t%v\t", datasource.Name), // No description available
        fmt.Sprintf("Description:\t%v\t", datasource.Description), // No description available
        fmt.Sprintf("Kind:\t%v\t", datasource.Kind), // No description available
        fmt.Sprintf("Configuration:\t%v\t", datasource.Configuration), // No description available
        fmt.Sprintf("CreatedAt:\t%v\t", datasource.CreatedAt), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
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
  var limit int64 // No description available
  var offset int64 // No description available
  var projectId int64 // No description available

  cmd := newCmd(c, getDatasourcesHelp, func(c *context, args []string) {
    
      // List datasources
      datasources, err := c.remote.GetDatasources(
        projectId, // No description available
        offset, // No description available
        limit, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(datasources))
      for i, e := range datasources {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.ProjectId, // No description available
          e.Name, // No description available
          e.Description, // No description available
          e.Kind, // No description available
          e.Configuration, // No description available
          e.CreatedAt, // No description available
        )
      }
      c.printt("Id\tProjectId\tName\tDescription\tKind\tConfiguration\tCreatedAt\t", lines)
      return
  })
  
  
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
  cmd.Flags().Int64Var(&projectId, "project-id", projectId, "No description available")
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
  var engineId int64 // No description available

  cmd := newCmd(c, getEngineHelp, func(c *context, args []string) {
    
      // Get engine details
      engine, err := c.remote.GetEngine(
        engineId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", engine.Id), // No description available
        fmt.Sprintf("Name:\t%v\t", engine.Name), // No description available
        fmt.Sprintf("Location:\t%v\t", engine.Location), // No description available
        fmt.Sprintf("CreatedAt:\t%v\t", engine.CreatedAt), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
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
    
      // List engines
      engines, err := c.remote.GetEngines(
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(engines))
      for i, e := range engines {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.Name, // No description available
          e.Location, // No description available
          e.CreatedAt, // No description available
        )
      }
      c.printt("Id\tName\tLocation\tCreatedAt\t", lines)
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
  var entityId int64 // No description available
  var entityTypeId int64 // No description available
  var limit int64 // No description available
  var offset int64 // No description available

  cmd := newCmd(c, getHistoryHelp, func(c *context, args []string) {
    
      // List audit trail records for an entity
      history, err := c.remote.GetHistory(
        entityTypeId, // No description available
        entityId, // No description available
        offset, // No description available
        limit, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(history))
      for i, e := range history {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t",
          e.IdentityId, // No description available
          e.Action, // No description available
          e.Description, // No description available
          e.CreatedAt, // No description available
        )
      }
      c.printt("IdentityId\tAction\tDescription\tCreatedAt\t", lines)
      return
  })
  
  
  cmd.Flags().Int64Var(&entityId, "entity-id", entityId, "No description available")
  cmd.Flags().Int64Var(&entityTypeId, "entity-type-id", entityTypeId, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
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
  var forWorkgroup bool // Switch for GetIdentitiesForWorkgroup()
  var forRole bool // Switch for GetIdentitiesForRole()
  var limit int64 // No description available
  var offset int64 // No description available
  var roleId int64 // No description available
  var workgroupId int64 // No description available

  cmd := newCmd(c, getIdentitiesHelp, func(c *context, args []string) {
    if forWorkgroup { // GetIdentitiesForWorkgroup
      
      // List identities for a workgroup
      identities, err := c.remote.GetIdentitiesForWorkgroup(
        workgroupId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(identities))
      for i, e := range identities {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.Name, // No description available
          e.IsActive, // No description available
          e.LastLogin, // No description available
          e.Created, // No description available
        )
      }
      c.printt("Id\tName\tIsActive\tLastLogin\tCreated\t", lines)
      return
    }
    if forRole { // GetIdentitiesForRole
      
      // List identities for a role
      identities, err := c.remote.GetIdentitiesForRole(
        roleId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(identities))
      for i, e := range identities {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.Name, // No description available
          e.IsActive, // No description available
          e.LastLogin, // No description available
          e.Created, // No description available
        )
      }
      c.printt("Id\tName\tIsActive\tLastLogin\tCreated\t", lines)
      return
    }
    if true { // default
      
      // List identities
      identities, err := c.remote.GetIdentities(
        offset, // No description available
        limit, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(identities))
      for i, e := range identities {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.Name, // No description available
          e.IsActive, // No description available
          e.LastLogin, // No description available
          e.Created, // No description available
        )
      }
      c.printt("Id\tName\tIsActive\tLastLogin\tCreated\t", lines)
      return
    }
  })
  cmd.Flags().BoolVar(&forWorkgroup, "for-workgroup", forWorkgroup, "List identities for a workgroup")
  cmd.Flags().BoolVar(&forRole, "for-role", forRole, "List identities for a role")
  
  
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
  cmd.Flags().Int64Var(&roleId, "role-id", roleId, "No description available")
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
  var byName bool // Switch for GetIdentityByName()
  var identityId int64 // No description available
  var name string // No description available

  cmd := newCmd(c, getIdentityHelp, func(c *context, args []string) {
    if byName { // GetIdentityByName
      
      // Get identity details by name
      identity, err := c.remote.GetIdentityByName(
        name, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", identity.Id), // No description available
        fmt.Sprintf("Name:\t%v\t", identity.Name), // No description available
        fmt.Sprintf("IsActive:\t%v\t", identity.IsActive), // No description available
        fmt.Sprintf("LastLogin:\t%v\t", identity.LastLogin), // No description available
        fmt.Sprintf("Created:\t%v\t", identity.Created), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
      return
    }
    if true { // default
      
      // Get identity details
      identity, err := c.remote.GetIdentity(
        identityId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", identity.Id), // No description available
        fmt.Sprintf("Name:\t%v\t", identity.Name), // No description available
        fmt.Sprintf("IsActive:\t%v\t", identity.IsActive), // No description available
        fmt.Sprintf("LastLogin:\t%v\t", identity.LastLogin), // No description available
        fmt.Sprintf("Created:\t%v\t", identity.Created), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
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
  var clusterId int64 // No description available
  var jobName string // No description available

  cmd := newCmd(c, getJobHelp, func(c *context, args []string) {
    
      // Get job details
      job, err := c.remote.GetJob(
        clusterId, // No description available
        jobName, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Name:\t%v\t", job.Name), // No description available
        fmt.Sprintf("ClusterName:\t%v\t", job.ClusterName), // No description available
        fmt.Sprintf("Description:\t%v\t", job.Description), // No description available
        fmt.Sprintf("Progress:\t%v\t", job.Progress), // No description available
        fmt.Sprintf("StartedAt:\t%v\t", job.StartedAt), // No description available
        fmt.Sprintf("CompletedAt:\t%v\t", job.CompletedAt), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
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
  var clusterId int64 // No description available

  cmd := newCmd(c, getJobsHelp, func(c *context, args []string) {
    
      // List jobs
      jobs, err := c.remote.GetJobs(
        clusterId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(jobs))
      for i, e := range jobs {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t%v\t",
          e.Name, // No description available
          e.ClusterName, // No description available
          e.Description, // No description available
          e.Progress, // No description available
          e.StartedAt, // No description available
          e.CompletedAt, // No description available
        )
      }
      c.printt("Name\tClusterName\tDescription\tProgress\tStartedAt\tCompletedAt\t", lines)
      return
  })
  
  
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  return cmd
}

var getLabelsHelp = `
labels [?]
Get Labels
Examples:

    List labels for a project, with corresponding models, if any
    $ steam get labels --for-project \
        --project-id=?

`

func getLabels(c *context) *cobra.Command {
  var forProject bool // Switch for GetLabelsForProject()
  var projectId int64 // No description available

  cmd := newCmd(c, getLabelsHelp, func(c *context, args []string) {
    if forProject { // GetLabelsForProject
      
      // List labels for a project, with corresponding models, if any
      labels, err := c.remote.GetLabelsForProject(
        projectId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(labels))
      for i, e := range labels {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.ProjectId, // No description available
          e.ModelId, // No description available
          e.Name, // No description available
          e.Description, // No description available
          e.CreatedAt, // No description available
        )
      }
      c.printt("Id\tProjectId\tModelId\tName\tDescription\tCreatedAt\t", lines)
      return
    }
  })
  cmd.Flags().BoolVar(&forProject, "for-project", forProject, "List labels for a project, with corresponding models, if any")
  
  
  cmd.Flags().Int64Var(&projectId, "project-id", projectId, "No description available")
  return cmd
}

var getModelHelp = `
model [?]
Get Model
Examples:

    Get model details
    $ steam get model \
        --model-id=?

    View a binomial model
    $ steam get model --binomial \
        --model-id=?

    View a binomial model
    $ steam get model --multinomial \
        --model-id=?

    View a binomial model
    $ steam get model --regression \
        --model-id=?

`

func getModel(c *context) *cobra.Command {
  var binomial bool // Switch for GetModelBinomial()
  var multinomial bool // Switch for GetModelMultinomial()
  var regression bool // Switch for GetModelRegression()
  var modelId int64 // No description available

  cmd := newCmd(c, getModelHelp, func(c *context, args []string) {
    if binomial { // GetModelBinomial
      
      // View a binomial model
      model, err := c.remote.GetModelBinomial(
        modelId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", model.Id), // No description available
        fmt.Sprintf("TrainingDatasetId:\t%v\t", model.TrainingDatasetId), // No description available
        fmt.Sprintf("ValidationDatasetId:\t%v\t", model.ValidationDatasetId), // No description available
        fmt.Sprintf("Name:\t%v\t", model.Name), // No description available
        fmt.Sprintf("ClusterName:\t%v\t", model.ClusterName), // No description available
        fmt.Sprintf("ModelKey:\t%v\t", model.ModelKey), // No description available
        fmt.Sprintf("Algorithm:\t%v\t", model.Algorithm), // No description available
        fmt.Sprintf("ModelCategory:\t%v\t", model.ModelCategory), // No description available
        fmt.Sprintf("DatasetName:\t%v\t", model.DatasetName), // No description available
        fmt.Sprintf("ResponseColumnName:\t%v\t", model.ResponseColumnName), // No description available
        fmt.Sprintf("LogicalName:\t%v\t", model.LogicalName), // No description available
        fmt.Sprintf("Location:\t%v\t", model.Location), // No description available
        fmt.Sprintf("MaxRuntime:\t%v\t", model.MaxRuntime), // No description available
        fmt.Sprintf("Metrics:\t%v\t", model.Metrics), // No description available
        fmt.Sprintf("CreatedAt:\t%v\t", model.CreatedAt), // No description available
        fmt.Sprintf("Mse:\t%v\t", model.Mse), // No description available
        fmt.Sprintf("RSquared:\t%v\t", model.RSquared), // No description available
        fmt.Sprintf("Logloss:\t%v\t", model.Logloss), // No description available
        fmt.Sprintf("Auc:\t%v\t", model.Auc), // No description available
        fmt.Sprintf("Gini:\t%v\t", model.Gini), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
      return
    }
    if multinomial { // GetModelMultinomial
      
      // View a binomial model
      model, err := c.remote.GetModelMultinomial(
        modelId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", model.Id), // No description available
        fmt.Sprintf("TrainingDatasetId:\t%v\t", model.TrainingDatasetId), // No description available
        fmt.Sprintf("ValidationDatasetId:\t%v\t", model.ValidationDatasetId), // No description available
        fmt.Sprintf("Name:\t%v\t", model.Name), // No description available
        fmt.Sprintf("ClusterName:\t%v\t", model.ClusterName), // No description available
        fmt.Sprintf("ModelKey:\t%v\t", model.ModelKey), // No description available
        fmt.Sprintf("Algorithm:\t%v\t", model.Algorithm), // No description available
        fmt.Sprintf("ModelCategory:\t%v\t", model.ModelCategory), // No description available
        fmt.Sprintf("DatasetName:\t%v\t", model.DatasetName), // No description available
        fmt.Sprintf("ResponseColumnName:\t%v\t", model.ResponseColumnName), // No description available
        fmt.Sprintf("LogicalName:\t%v\t", model.LogicalName), // No description available
        fmt.Sprintf("Location:\t%v\t", model.Location), // No description available
        fmt.Sprintf("MaxRuntime:\t%v\t", model.MaxRuntime), // No description available
        fmt.Sprintf("Metrics:\t%v\t", model.Metrics), // No description available
        fmt.Sprintf("CreatedAt:\t%v\t", model.CreatedAt), // No description available
        fmt.Sprintf("Mse:\t%v\t", model.Mse), // No description available
        fmt.Sprintf("RSquared:\t%v\t", model.RSquared), // No description available
        fmt.Sprintf("Logloss:\t%v\t", model.Logloss), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
      return
    }
    if regression { // GetModelRegression
      
      // View a binomial model
      model, err := c.remote.GetModelRegression(
        modelId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", model.Id), // No description available
        fmt.Sprintf("TrainingDatasetId:\t%v\t", model.TrainingDatasetId), // No description available
        fmt.Sprintf("ValidationDatasetId:\t%v\t", model.ValidationDatasetId), // No description available
        fmt.Sprintf("Name:\t%v\t", model.Name), // No description available
        fmt.Sprintf("ClusterName:\t%v\t", model.ClusterName), // No description available
        fmt.Sprintf("ModelKey:\t%v\t", model.ModelKey), // No description available
        fmt.Sprintf("Algorithm:\t%v\t", model.Algorithm), // No description available
        fmt.Sprintf("ModelCategory:\t%v\t", model.ModelCategory), // No description available
        fmt.Sprintf("DatasetName:\t%v\t", model.DatasetName), // No description available
        fmt.Sprintf("ResponseColumnName:\t%v\t", model.ResponseColumnName), // No description available
        fmt.Sprintf("LogicalName:\t%v\t", model.LogicalName), // No description available
        fmt.Sprintf("Location:\t%v\t", model.Location), // No description available
        fmt.Sprintf("MaxRuntime:\t%v\t", model.MaxRuntime), // No description available
        fmt.Sprintf("Metrics:\t%v\t", model.Metrics), // No description available
        fmt.Sprintf("CreatedAt:\t%v\t", model.CreatedAt), // No description available
        fmt.Sprintf("Mse:\t%v\t", model.Mse), // No description available
        fmt.Sprintf("RSquared:\t%v\t", model.RSquared), // No description available
        fmt.Sprintf("MeanResidualDeviance:\t%v\t", model.MeanResidualDeviance), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
      return
    }
    if true { // default
      
      // Get model details
      model, err := c.remote.GetModel(
        modelId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", model.Id), // No description available
        fmt.Sprintf("TrainingDatasetId:\t%v\t", model.TrainingDatasetId), // No description available
        fmt.Sprintf("ValidationDatasetId:\t%v\t", model.ValidationDatasetId), // No description available
        fmt.Sprintf("Name:\t%v\t", model.Name), // No description available
        fmt.Sprintf("ClusterName:\t%v\t", model.ClusterName), // No description available
        fmt.Sprintf("ModelKey:\t%v\t", model.ModelKey), // No description available
        fmt.Sprintf("Algorithm:\t%v\t", model.Algorithm), // No description available
        fmt.Sprintf("ModelCategory:\t%v\t", model.ModelCategory), // No description available
        fmt.Sprintf("DatasetName:\t%v\t", model.DatasetName), // No description available
        fmt.Sprintf("ResponseColumnName:\t%v\t", model.ResponseColumnName), // No description available
        fmt.Sprintf("LogicalName:\t%v\t", model.LogicalName), // No description available
        fmt.Sprintf("Location:\t%v\t", model.Location), // No description available
        fmt.Sprintf("MaxRuntime:\t%v\t", model.MaxRuntime), // No description available
        fmt.Sprintf("Metrics:\t%v\t", model.Metrics), // No description available
        fmt.Sprintf("CreatedAt:\t%v\t", model.CreatedAt), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
      return
    }
  })
  cmd.Flags().BoolVar(&binomial, "binomial", binomial, "View a binomial model")
  cmd.Flags().BoolVar(&multinomial, "multinomial", multinomial, "View a binomial model")
  cmd.Flags().BoolVar(&regression, "regression", regression, "View a binomial model")
  
  
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
        --cluster-id=? \
        --frame-key=?

`

func getModels(c *context) *cobra.Command {
  var fromCluster bool // Switch for GetModelsFromCluster()
  var clusterId int64 // No description available
  var frameKey string // No description available
  var limit int64 // No description available
  var offset int64 // No description available
  var projectId int64 // No description available

  cmd := newCmd(c, getModelsHelp, func(c *context, args []string) {
    if fromCluster { // GetModelsFromCluster
      
      // List models from a cluster
      models, err := c.remote.GetModelsFromCluster(
        clusterId, // No description available
        frameKey, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(models))
      for i, e := range models {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.TrainingDatasetId, // No description available
          e.ValidationDatasetId, // No description available
          e.Name, // No description available
          e.ClusterName, // No description available
          e.ModelKey, // No description available
          e.Algorithm, // No description available
          e.ModelCategory, // No description available
          e.DatasetName, // No description available
          e.ResponseColumnName, // No description available
          e.LogicalName, // No description available
          e.Location, // No description available
          e.MaxRuntime, // No description available
          e.Metrics, // No description available
          e.CreatedAt, // No description available
        )
      }
      c.printt("Id\tTrainingDatasetId\tValidationDatasetId\tName\tClusterName\tModelKey\tAlgorithm\tModelCategory\tDatasetName\tResponseColumnName\tLogicalName\tLocation\tMaxRuntime\tMetrics\tCreatedAt\t", lines)
      return
    }
    if true { // default
      
      // List models
      models, err := c.remote.GetModels(
        projectId, // No description available
        offset, // No description available
        limit, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(models))
      for i, e := range models {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.TrainingDatasetId, // No description available
          e.ValidationDatasetId, // No description available
          e.Name, // No description available
          e.ClusterName, // No description available
          e.ModelKey, // No description available
          e.Algorithm, // No description available
          e.ModelCategory, // No description available
          e.DatasetName, // No description available
          e.ResponseColumnName, // No description available
          e.LogicalName, // No description available
          e.Location, // No description available
          e.MaxRuntime, // No description available
          e.Metrics, // No description available
          e.CreatedAt, // No description available
        )
      }
      c.printt("Id\tTrainingDatasetId\tValidationDatasetId\tName\tClusterName\tModelKey\tAlgorithm\tModelCategory\tDatasetName\tResponseColumnName\tLogicalName\tLocation\tMaxRuntime\tMetrics\tCreatedAt\t", lines)
      return
    }
  })
  cmd.Flags().BoolVar(&fromCluster, "from-cluster", fromCluster, "List models from a cluster")
  
  
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  cmd.Flags().StringVar(&frameKey, "frame-key", frameKey, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
  cmd.Flags().Int64Var(&projectId, "project-id", projectId, "No description available")
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
  var forRole bool // Switch for GetPermissionsForRole()
  var forIdentity bool // Switch for GetPermissionsForIdentity()
  var identityId int64 // No description available
  var roleId int64 // No description available

  cmd := newCmd(c, getPermissionsHelp, func(c *context, args []string) {
    if forRole { // GetPermissionsForRole
      
      // List permissions for a role
      permissions, err := c.remote.GetPermissionsForRole(
        roleId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(permissions))
      for i, e := range permissions {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t",
          e.Id, // No description available
          e.Code, // No description available
          e.Description, // No description available
        )
      }
      c.printt("Id\tCode\tDescription\t", lines)
      return
    }
    if forIdentity { // GetPermissionsForIdentity
      
      // List permissions for an identity
      permissions, err := c.remote.GetPermissionsForIdentity(
        identityId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(permissions))
      for i, e := range permissions {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t",
          e.Id, // No description available
          e.Code, // No description available
          e.Description, // No description available
        )
      }
      c.printt("Id\tCode\tDescription\t", lines)
      return
    }
  })
  cmd.Flags().BoolVar(&forRole, "for-role", forRole, "List permissions for a role")
  cmd.Flags().BoolVar(&forIdentity, "for-identity", forIdentity, "List permissions for an identity")
  
  
  cmd.Flags().Int64Var(&identityId, "identity-id", identityId, "No description available")
  cmd.Flags().Int64Var(&roleId, "role-id", roleId, "No description available")
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
  var entityId int64 // No description available
  var entityTypeId int64 // No description available

  cmd := newCmd(c, getPrivilegesHelp, func(c *context, args []string) {
    
      // List privileges for an entity
      privileges, err := c.remote.GetPrivileges(
        entityTypeId, // No description available
        entityId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(privileges))
      for i, e := range privileges {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t",
          e.Kind, // No description available
          e.WorkgroupId, // No description available
          e.WorkgroupName, // No description available
          e.WorkgroupDescription, // No description available
        )
      }
      c.printt("Kind\tWorkgroupId\tWorkgroupName\tWorkgroupDescription\t", lines)
      return
  })
  
  
  cmd.Flags().Int64Var(&entityId, "entity-id", entityId, "No description available")
  cmd.Flags().Int64Var(&entityTypeId, "entity-type-id", entityTypeId, "No description available")
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
  var projectId int64 // No description available

  cmd := newCmd(c, getProjectHelp, func(c *context, args []string) {
    
      // Get project details
      project, err := c.remote.GetProject(
        projectId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", project.Id), // No description available
        fmt.Sprintf("Name:\t%v\t", project.Name), // No description available
        fmt.Sprintf("Description:\t%v\t", project.Description), // No description available
        fmt.Sprintf("ModelCategory:\t%v\t", project.ModelCategory), // No description available
        fmt.Sprintf("CreatedAt:\t%v\t", project.CreatedAt), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
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
  var limit int64 // No description available
  var offset int64 // No description available

  cmd := newCmd(c, getProjectsHelp, func(c *context, args []string) {
    
      // List projects
      projects, err := c.remote.GetProjects(
        offset, // No description available
        limit, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(projects))
      for i, e := range projects {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.Name, // No description available
          e.Description, // No description available
          e.ModelCategory, // No description available
          e.CreatedAt, // No description available
        )
      }
      c.printt("Id\tName\tDescription\tModelCategory\tCreatedAt\t", lines)
      return
  })
  
  
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
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
  var byName bool // Switch for GetRoleByName()
  var name string // No description available
  var roleId int64 // No description available

  cmd := newCmd(c, getRoleHelp, func(c *context, args []string) {
    if byName { // GetRoleByName
      
      // Get role details by name
      role, err := c.remote.GetRoleByName(
        name, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", role.Id), // No description available
        fmt.Sprintf("Name:\t%v\t", role.Name), // No description available
        fmt.Sprintf("Description:\t%v\t", role.Description), // No description available
        fmt.Sprintf("Created:\t%v\t", role.Created), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
      return
    }
    if true { // default
      
      // Get role details
      role, err := c.remote.GetRole(
        roleId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", role.Id), // No description available
        fmt.Sprintf("Name:\t%v\t", role.Name), // No description available
        fmt.Sprintf("Description:\t%v\t", role.Description), // No description available
        fmt.Sprintf("Created:\t%v\t", role.Created), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
      return
    }
  })
  cmd.Flags().BoolVar(&byName, "by-name", byName, "Get role details by name")
  
  
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().Int64Var(&roleId, "role-id", roleId, "No description available")
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
  var forIdentity bool // Switch for GetRolesForIdentity()
  var identityId int64 // No description available
  var limit int64 // No description available
  var offset int64 // No description available

  cmd := newCmd(c, getRolesHelp, func(c *context, args []string) {
    if forIdentity { // GetRolesForIdentity
      
      // List roles for an identity
      roles, err := c.remote.GetRolesForIdentity(
        identityId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(roles))
      for i, e := range roles {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.Name, // No description available
          e.Description, // No description available
          e.Created, // No description available
        )
      }
      c.printt("Id\tName\tDescription\tCreated\t", lines)
      return
    }
    if true { // default
      
      // List roles
      roles, err := c.remote.GetRoles(
        offset, // No description available
        limit, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(roles))
      for i, e := range roles {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.Name, // No description available
          e.Description, // No description available
          e.Created, // No description available
        )
      }
      c.printt("Id\tName\tDescription\tCreated\t", lines)
      return
    }
  })
  cmd.Flags().BoolVar(&forIdentity, "for-identity", forIdentity, "List roles for an identity")
  
  
  cmd.Flags().Int64Var(&identityId, "identity-id", identityId, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
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
  var serviceId int64 // No description available

  cmd := newCmd(c, getServiceHelp, func(c *context, args []string) {
    
      // Get service details
      service, err := c.remote.GetService(
        serviceId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", service.Id), // No description available
        fmt.Sprintf("ModelId:\t%v\t", service.ModelId), // No description available
        fmt.Sprintf("Address:\t%v\t", service.Address), // No description available
        fmt.Sprintf("Port:\t%v\t", service.Port), // No description available
        fmt.Sprintf("ProcessId:\t%v\t", service.ProcessId), // No description available
        fmt.Sprintf("State:\t%v\t", service.State), // No description available
        fmt.Sprintf("CreatedAt:\t%v\t", service.CreatedAt), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
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
  var forModel bool // Switch for GetServicesForModel()
  var limit int64 // No description available
  var modelId int64 // No description available
  var offset int64 // No description available

  cmd := newCmd(c, getServicesHelp, func(c *context, args []string) {
    if forModel { // GetServicesForModel
      
      // List services for a model
      services, err := c.remote.GetServicesForModel(
        modelId, // No description available
        offset, // No description available
        limit, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(services))
      for i, e := range services {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.ModelId, // No description available
          e.Address, // No description available
          e.Port, // No description available
          e.ProcessId, // No description available
          e.State, // No description available
          e.CreatedAt, // No description available
        )
      }
      c.printt("Id\tModelId\tAddress\tPort\tProcessId\tState\tCreatedAt\t", lines)
      return
    }
    if true { // default
      
      // List services
      services, err := c.remote.GetServices(
        offset, // No description available
        limit, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(services))
      for i, e := range services {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.ModelId, // No description available
          e.Address, // No description available
          e.Port, // No description available
          e.ProcessId, // No description available
          e.State, // No description available
          e.CreatedAt, // No description available
        )
      }
      c.printt("Id\tModelId\tAddress\tPort\tProcessId\tState\tCreatedAt\t", lines)
      return
    }
  })
  cmd.Flags().BoolVar(&forModel, "for-model", forModel, "List services for a model")
  
  
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&modelId, "model-id", modelId, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
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
  var byName bool // Switch for GetWorkgroupByName()
  var name string // No description available
  var workgroupId int64 // No description available

  cmd := newCmd(c, getWorkgroupHelp, func(c *context, args []string) {
    if byName { // GetWorkgroupByName
      
      // Get workgroup details by name
      workgroup, err := c.remote.GetWorkgroupByName(
        name, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", workgroup.Id), // No description available
        fmt.Sprintf("Name:\t%v\t", workgroup.Name), // No description available
        fmt.Sprintf("Description:\t%v\t", workgroup.Description), // No description available
        fmt.Sprintf("Created:\t%v\t", workgroup.Created), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
      return
    }
    if true { // default
      
      // Get workgroup details
      workgroup, err := c.remote.GetWorkgroup(
        workgroupId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", workgroup.Id), // No description available
        fmt.Sprintf("Name:\t%v\t", workgroup.Name), // No description available
        fmt.Sprintf("Description:\t%v\t", workgroup.Description), // No description available
        fmt.Sprintf("Created:\t%v\t", workgroup.Created), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
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
  var forIdentity bool // Switch for GetWorkgroupsForIdentity()
  var identityId int64 // No description available
  var limit int64 // No description available
  var offset int64 // No description available

  cmd := newCmd(c, getWorkgroupsHelp, func(c *context, args []string) {
    if forIdentity { // GetWorkgroupsForIdentity
      
      // List workgroups for an identity
      workgroups, err := c.remote.GetWorkgroupsForIdentity(
        identityId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(workgroups))
      for i, e := range workgroups {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.Name, // No description available
          e.Description, // No description available
          e.Created, // No description available
        )
      }
      c.printt("Id\tName\tDescription\tCreated\t", lines)
      return
    }
    if true { // default
      
      // List workgroups
      workgroups, err := c.remote.GetWorkgroups(
        offset, // No description available
        limit, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := make([]string, len(workgroups))
      for i, e := range workgroups {
        lines[i] = fmt.Sprintf(
          "%v\t%v\t%v\t%v\t",
          e.Id, // No description available
          e.Name, // No description available
          e.Description, // No description available
          e.Created, // No description available
        )
      }
      c.printt("Id\tName\tDescription\tCreated\t", lines)
      return
    }
  })
  cmd.Flags().BoolVar(&forIdentity, "for-identity", forIdentity, "List workgroups for an identity")
  
  
  cmd.Flags().Int64Var(&identityId, "identity-id", identityId, "No description available")
  cmd.Flags().Int64Var(&limit, "limit", limit, "No description available")
  cmd.Flags().Int64Var(&offset, "offset", offset, "No description available")
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
        --model-key=? \
        --model-name=?

`

func importModel(c *context) *cobra.Command {
  var fromCluster bool // Switch for ImportModelFromCluster()
  var clusterId int64 // No description available
  var modelKey string // No description available
  var modelName string // No description available
  var projectId int64 // No description available

  cmd := newCmd(c, importModelHelp, func(c *context, args []string) {
    if fromCluster { // ImportModelFromCluster
      
      // Import models from a cluster
      modelId, err := c.remote.ImportModelFromCluster(
        clusterId, // No description available
        projectId, // No description available
        modelKey, // No description available
        modelName, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("ModelId:\t%v\n", modelId)
      return
    }
  })
  cmd.Flags().BoolVar(&fromCluster, "from-cluster", fromCluster, "Import models from a cluster")
  
  
  cmd.Flags().Int64Var(&clusterId, "cluster-id", clusterId, "No description available")
  cmd.Flags().StringVar(&modelKey, "model-key", modelKey, "No description available")
  cmd.Flags().StringVar(&modelName, "model-name", modelName, "No description available")
  cmd.Flags().Int64Var(&projectId, "project-id", projectId, "No description available")
  return cmd
}




var linkHelp = `
link [?]
Link entities
Commands:

    $ steam link identity ...
    $ steam link label ...
    $ steam link role ...
`
func link(c *context) *cobra.Command {
  cmd := newCmd(c, linkHelp, nil)

  cmd.AddCommand(linkIdentity(c))
  cmd.AddCommand(linkLabel(c))
  cmd.AddCommand(linkRole(c))
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
  var withWorkgroup bool // Switch for LinkIdentityWithWorkgroup()
  var withRole bool // Switch for LinkIdentityWithRole()
  var identityId int64 // No description available
  var roleId int64 // No description available
  var workgroupId int64 // No description available

  cmd := newCmd(c, linkIdentityHelp, func(c *context, args []string) {
    if withWorkgroup { // LinkIdentityWithWorkgroup
      
      // Link an identity with a workgroup
      err := c.remote.LinkIdentityWithWorkgroup(
        identityId, // No description available
        workgroupId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
    }
    if withRole { // LinkIdentityWithRole
      
      // Link an identity with a role
      err := c.remote.LinkIdentityWithRole(
        identityId, // No description available
        roleId, // No description available
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
  cmd.Flags().Int64Var(&roleId, "role-id", roleId, "No description available")
  cmd.Flags().Int64Var(&workgroupId, "workgroup-id", workgroupId, "No description available")
  return cmd
}

var linkLabelHelp = `
label [?]
Link Label
Examples:

    Label a model
    $ steam link label --with-model \
        --label-id=? \
        --model-id=?

`

func linkLabel(c *context) *cobra.Command {
  var withModel bool // Switch for LinkLabelWithModel()
  var labelId int64 // No description available
  var modelId int64 // No description available

  cmd := newCmd(c, linkLabelHelp, func(c *context, args []string) {
    if withModel { // LinkLabelWithModel
      
      // Label a model
      err := c.remote.LinkLabelWithModel(
        labelId, // No description available
        modelId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
    }
  })
  cmd.Flags().BoolVar(&withModel, "with-model", withModel, "Label a model")
  
  
  cmd.Flags().Int64Var(&labelId, "label-id", labelId, "No description available")
  cmd.Flags().Int64Var(&modelId, "model-id", modelId, "No description available")
  return cmd
}

var linkRoleHelp = `
role [?]
Link Role
Examples:

    Link a role with a permission
    $ steam link role --with-permission \
        --role-id=? \
        --permission-id=?

`

func linkRole(c *context) *cobra.Command {
  var withPermission bool // Switch for LinkRoleWithPermission()
  var permissionId int64 // No description available
  var roleId int64 // No description available

  cmd := newCmd(c, linkRoleHelp, func(c *context, args []string) {
    if withPermission { // LinkRoleWithPermission
      
      // Link a role with a permission
      err := c.remote.LinkRoleWithPermission(
        roleId, // No description available
        permissionId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
    }
  })
  cmd.Flags().BoolVar(&withPermission, "with-permission", withPermission, "Link a role with a permission")
  
  
  cmd.Flags().Int64Var(&permissionId, "permission-id", permissionId, "No description available")
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
  var input string // Message to send

  cmd := newCmd(c, pingServerHelp, func(c *context, args []string) {
    
      // Ping the Steam server
      output, err := c.remote.PingServer(
        input, // Message to send
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("Output:\t%v\n", output)
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
  var address string // No description available

  cmd := newCmd(c, registerClusterHelp, func(c *context, args []string) {
    
      // Connect to a cluster
      clusterId, err := c.remote.RegisterCluster(
        address, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("ClusterId:\t%v\n", clusterId)
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
  var entityId int64 // No description available
  var entityTypeId int64 // No description available
  var kind string // No description available
  var workgroupId int64 // No description available

  cmd := newCmd(c, shareEntityHelp, func(c *context, args []string) {
    
      // Share an entity with a workgroup
      err := c.remote.ShareEntity(
        kind, // No description available
        workgroupId, // No description available
        entityTypeId, // No description available
        entityId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
  })
  
  
  cmd.Flags().Int64Var(&entityId, "entity-id", entityId, "No description available")
  cmd.Flags().Int64Var(&entityTypeId, "entity-type-id", entityTypeId, "No description available")
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
  var datasetId int64 // No description available
  var ratio1 int // No description available
  var ratio2 int // No description available

  cmd := newCmd(c, splitDatasetHelp, func(c *context, args []string) {
    
      // Split a dataset
      datasetIds, err := c.remote.SplitDataset(
        datasetId, // No description available
        ratio1, // No description available
        ratio2, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("DatasetIds:\t%v\n", datasetIds)
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
  var onYarn bool // Switch for StartClusterOnYarn()
  var clusterName string // No description available
  var engineId int64 // No description available
  var memory string // No description available
  var size int // No description available
  var username string // No description available

  cmd := newCmd(c, startClusterHelp, func(c *context, args []string) {
    if onYarn { // StartClusterOnYarn
      
      // Start a cluster using Yarn
      clusterId, err := c.remote.StartClusterOnYarn(
        clusterName, // No description available
        engineId, // No description available
        size, // No description available
        memory, // No description available
        username, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      fmt.Printf("ClusterId:\t%v\n", clusterId)
      return
    }
  })
  cmd.Flags().BoolVar(&onYarn, "on-yarn", onYarn, "Start a cluster using Yarn")
  
  
  cmd.Flags().StringVar(&clusterName, "cluster-name", clusterName, "No description available")
  cmd.Flags().Int64Var(&engineId, "engine-id", engineId, "No description available")
  cmd.Flags().StringVar(&memory, "memory", memory, "No description available")
  cmd.Flags().IntVar(&size, "size", size, "No description available")
  cmd.Flags().StringVar(&username, "username", username, "No description available")
  return cmd
}

var startServiceHelp = `
service [?]
Start Service
Examples:

    Start a service
    $ steam start service \
        --model-id=?

`

func startService(c *context) *cobra.Command {
  var modelId int64 // No description available

  cmd := newCmd(c, startServiceHelp, func(c *context, args []string) {
    
      // Start a service
      service, err := c.remote.StartService(
        modelId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      lines := []string{
        fmt.Sprintf("Id:\t%v\t", service.Id), // No description available
        fmt.Sprintf("ModelId:\t%v\t", service.ModelId), // No description available
        fmt.Sprintf("Address:\t%v\t", service.Address), // No description available
        fmt.Sprintf("Port:\t%v\t", service.Port), // No description available
        fmt.Sprintf("ProcessId:\t%v\t", service.ProcessId), // No description available
        fmt.Sprintf("State:\t%v\t", service.State), // No description available
        fmt.Sprintf("CreatedAt:\t%v\t", service.CreatedAt), // No description available
      }
      c.printt("Attribute\tValue\t", lines)
      return
  })
  
  
  cmd.Flags().Int64Var(&modelId, "model-id", modelId, "No description available")
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
  var onYarn bool // Switch for StopClusterOnYarn()
  var clusterId int64 // No description available

  cmd := newCmd(c, stopClusterHelp, func(c *context, args []string) {
    if onYarn { // StopClusterOnYarn
      
      // Stop a cluster using Yarn
      err := c.remote.StopClusterOnYarn(
        clusterId, // No description available
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
  var serviceId int64 // No description available

  cmd := newCmd(c, stopServiceHelp, func(c *context, args []string) {
    
      // Stop a service
      err := c.remote.StopService(
        serviceId, // No description available
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
    $ steam unlink label ...
    $ steam unlink role ...
`
func unlink(c *context) *cobra.Command {
  cmd := newCmd(c, unlinkHelp, nil)

  cmd.AddCommand(unlinkIdentity(c))
  cmd.AddCommand(unlinkLabel(c))
  cmd.AddCommand(unlinkRole(c))
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
  var fromWorkgroup bool // Switch for UnlinkIdentityFromWorkgroup()
  var fromRole bool // Switch for UnlinkIdentityFromRole()
  var identityId int64 // No description available
  var roleId int64 // No description available
  var workgroupId int64 // No description available

  cmd := newCmd(c, unlinkIdentityHelp, func(c *context, args []string) {
    if fromWorkgroup { // UnlinkIdentityFromWorkgroup
      
      // Unlink an identity from a workgroup
      err := c.remote.UnlinkIdentityFromWorkgroup(
        identityId, // No description available
        workgroupId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
    }
    if fromRole { // UnlinkIdentityFromRole
      
      // Unlink an identity from a role
      err := c.remote.UnlinkIdentityFromRole(
        identityId, // No description available
        roleId, // No description available
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
  cmd.Flags().Int64Var(&roleId, "role-id", roleId, "No description available")
  cmd.Flags().Int64Var(&workgroupId, "workgroup-id", workgroupId, "No description available")
  return cmd
}

var unlinkLabelHelp = `
label [?]
Unlink Label
Examples:

    Remove a label from a model
    $ steam unlink label --from-model \
        --label-id=? \
        --model-id=?

`

func unlinkLabel(c *context) *cobra.Command {
  var fromModel bool // Switch for UnlinkLabelFromModel()
  var labelId int64 // No description available
  var modelId int64 // No description available

  cmd := newCmd(c, unlinkLabelHelp, func(c *context, args []string) {
    if fromModel { // UnlinkLabelFromModel
      
      // Remove a label from a model
      err := c.remote.UnlinkLabelFromModel(
        labelId, // No description available
        modelId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
    }
  })
  cmd.Flags().BoolVar(&fromModel, "from-model", fromModel, "Remove a label from a model")
  
  
  cmd.Flags().Int64Var(&labelId, "label-id", labelId, "No description available")
  cmd.Flags().Int64Var(&modelId, "model-id", modelId, "No description available")
  return cmd
}

var unlinkRoleHelp = `
role [?]
Unlink Role
Examples:

    Unlink a role from a permission
    $ steam unlink role --from-permission \
        --role-id=? \
        --permission-id=?

`

func unlinkRole(c *context) *cobra.Command {
  var fromPermission bool // Switch for UnlinkRoleFromPermission()
  var permissionId int64 // No description available
  var roleId int64 // No description available

  cmd := newCmd(c, unlinkRoleHelp, func(c *context, args []string) {
    if fromPermission { // UnlinkRoleFromPermission
      
      // Unlink a role from a permission
      err := c.remote.UnlinkRoleFromPermission(
        roleId, // No description available
        permissionId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
    }
  })
  cmd.Flags().BoolVar(&fromPermission, "from-permission", fromPermission, "Unlink a role from a permission")
  
  
  cmd.Flags().Int64Var(&permissionId, "permission-id", permissionId, "No description available")
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
  var clusterId int64 // No description available

  cmd := newCmd(c, unregisterClusterHelp, func(c *context, args []string) {
    
      // Disconnect from a cluster
      err := c.remote.UnregisterCluster(
        clusterId, // No description available
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
  var entityId int64 // No description available
  var entityTypeId int64 // No description available
  var kind string // No description available
  var workgroupId int64 // No description available

  cmd := newCmd(c, unshareEntityHelp, func(c *context, args []string) {
    
      // Unshare an entity
      err := c.remote.UnshareEntity(
        kind, // No description available
        workgroupId, // No description available
        entityTypeId, // No description available
        entityId, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
  })
  
  
  cmd.Flags().Int64Var(&entityId, "entity-id", entityId, "No description available")
  cmd.Flags().Int64Var(&entityTypeId, "entity-type-id", entityTypeId, "No description available")
  cmd.Flags().StringVar(&kind, "kind", kind, "No description available")
  cmd.Flags().Int64Var(&workgroupId, "workgroup-id", workgroupId, "No description available")
  return cmd
}




var updateHelp = `
update [?]
Update entities
Commands:

    $ steam update dataset ...
    $ steam update datasource ...
    $ steam update identity ...
    $ steam update label ...
    $ steam update role ...
    $ steam update workgroup ...
`
func update(c *context) *cobra.Command {
  cmd := newCmd(c, updateHelp, nil)

  cmd.AddCommand(updateDataset(c))
  cmd.AddCommand(updateDatasource(c))
  cmd.AddCommand(updateIdentity(c))
  cmd.AddCommand(updateLabel(c))
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
  var datasetId int64 // No description available
  var description string // No description available
  var name string // No description available
  var responseColumnName string // No description available

  cmd := newCmd(c, updateDatasetHelp, func(c *context, args []string) {
    
      // Update a dataset
      err := c.remote.UpdateDataset(
        datasetId, // No description available
        name, // No description available
        description, // No description available
        responseColumnName, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
  })
  
  
  cmd.Flags().Int64Var(&datasetId, "dataset-id", datasetId, "No description available")
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
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
  var datasourceId int64 // No description available
  var description string // No description available
  var name string // No description available
  var path string // No description available

  cmd := newCmd(c, updateDatasourceHelp, func(c *context, args []string) {
    
      // Update a datasource
      err := c.remote.UpdateDatasource(
        datasourceId, // No description available
        name, // No description available
        description, // No description available
        path, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
  })
  
  
  cmd.Flags().Int64Var(&datasourceId, "datasource-id", datasourceId, "No description available")
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
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
  var identityId int64 // No description available
  var password string // No description available

  cmd := newCmd(c, updateIdentityHelp, func(c *context, args []string) {
    
      // Update an identity
      err := c.remote.UpdateIdentity(
        identityId, // No description available
        password, // No description available
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

var updateLabelHelp = `
label [?]
Update Label
Examples:

    Update a label
    $ steam update label \
        --label-id=? \
        --name=? \
        --description=?

`

func updateLabel(c *context) *cobra.Command {
  var description string // No description available
  var labelId int64 // No description available
  var name string // No description available

  cmd := newCmd(c, updateLabelHelp, func(c *context, args []string) {
    
      // Update a label
      err := c.remote.UpdateLabel(
        labelId, // No description available
        name, // No description available
        description, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
  })
  
  
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().Int64Var(&labelId, "label-id", labelId, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
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
  var description string // No description available
  var name string // No description available
  var roleId int64 // No description available

  cmd := newCmd(c, updateRoleHelp, func(c *context, args []string) {
    
      // Update a role
      err := c.remote.UpdateRole(
        roleId, // No description available
        name, // No description available
        description, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
  })
  
  
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().Int64Var(&roleId, "role-id", roleId, "No description available")
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
  var description string // No description available
  var name string // No description available
  var workgroupId int64 // No description available

  cmd := newCmd(c, updateWorkgroupHelp, func(c *context, args []string) {
    
      // Update a workgroup
      err := c.remote.UpdateWorkgroup(
        workgroupId, // No description available
        name, // No description available
        description, // No description available
      )
      if err != nil {
        log.Fatalln(err)
      }
      return
  })
  
  
  cmd.Flags().StringVar(&description, "description", description, "No description available")
  cmd.Flags().StringVar(&name, "name", name, "No description available")
  cmd.Flags().Int64Var(&workgroupId, "workgroup-id", workgroupId, "No description available")
  return cmd
}



