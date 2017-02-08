# ------------------------------
# --- This is generated code ---
# ---      DO NOT EDIT       ---
# ------------------------------

#'
#' Steam backend API calls
#'
#' These are the methods to communicate to the backend Steam API through RPC.
#' @import jsonlite
#' @import RCurl
#'

# .steamRPC is a wrapper for the curl calls to Steam
.steamRPC <- function(conn, method, params) {
    # Validate fields
    if (!is.SteamConnection(conn)) { stop("conn must be a valid SteamConnection")}

    if (length(params) > 0) {
        params = list(params)
    }
    request <- list(
        method = paste("web", method, sep="."),
        params = params,
        id = conn$uid
    )

    # Args for valid curl
    payload <- toJSON(request, null="list", auto_unbox=TRUE)
    address <- paste(conn$ip, paste(conn$port, "web", sep="/"), sep=":")
    userpwd <- paste(conn$username, conn$password, sep=":")
    header <- c('Content-Type' = 'application/json')
    opts = curlOptions(userpwd = userpwd,httpauth = 1L, ssl.verifypeer = conn$verify_ssl)

    # Error handling values
    h <- basicHeaderGatherer()
    .__curlErr = FALSE
    .__curlErrMessage <- ""
    httpStatusCode = -1L
    httpStatusMessage = ""

    # Return values
    g <- basicTextGatherer(.mapUnicode = FALSE)

    tmp <- tryCatch(curlPerform(
            url = address,
            postfields = payload,
            writefunction = g$update,
            headerfunction = h$update,
            httpheader = header,
            .opts = opts
        ),
        error = function(x) { .__curlErr <<- TRUE; .__curlErrMessage <<- x$message }
    )
    if (.__curlErr) {
        stop(sprintf("%s: %s", .__curlErrMessage, g$value()[1]))
    } else if (h$value()["status"] != 200){
        stop(g$value()[1])
    }
    response <- fromJSON(g$value())
    if (!is.null(response$error)) {
        stop(response$error)
    }

    conn$uid <<- conn$uid + 1
    response$result
}


#' Ping the Steam server
#'
#' @param input \code{character} Message to send
#'
.pingServer <- function(conn, input) {
    request <- list(
        input = input
    )

    .steamRPC(conn, "PingServer", request)
}

#' Get Steam start up configurations
#'
#'
.getConfig <- function(conn) {
    request <- list(
    )

    .steamRPC(conn, "GetConfig", request)
}

#' Set this to enable kerberos usage when applicable
#'
#' @param enabled \code{logical} Whether kerberos should be enabled or disabled
#'
.setGlobalKerberos <- function(conn, enabled) {
    request <- list(
        enabled = enabled
    )

    .steamRPC(conn, "SetGlobalKerberos", request)
}

#' Check if an identity has admin privileges
#'
#'
.checkAdmin <- function(conn) {
    request <- list(
    )

    .steamRPC(conn, "CheckAdmin", request)
}

#' Set security configuration to local
#'
#'
.setLocalConfig <- function(conn) {
    request <- list(
    )

    .steamRPC(conn, "SetLocalConfig", request)
}

#' Set LDAP security configuration
#'
#' @param config No description available
#'
.setLdapConfig <- function(conn, config) {
    request <- list(
        config = config
    )

    .steamRPC(conn, "SetLdapConfig", request)
}

#' Get LDAP security configurations
#'
#'
.getLdapConfig <- function(conn) {
    request <- list(
    )

    .steamRPC(conn, "GetLdapConfig", request)
}

#' Test LDAP security configurations
#'
#' @param config No description available
#'
.testLdapConfig <- function(conn, config) {
    request <- list(
        config = config
    )

    .steamRPC(conn, "TestLdapConfig", request)
}

#' Get the keytab for the logged in user
#'
#'
.getUserKeytab <- function(conn) {
    request <- list(
    )

    .steamRPC(conn, "GetUserKeytab", request)
}

#' Get the keytab for Steam (used for polling)
#'
#'
.getSteamKeytab <- function(conn) {
    request <- list(
    )

    .steamRPC(conn, "GetSteamKeytab", request)
}

#' Test the keytab for the given user
#'
#' @param keytab_id \code{numeric} No description available
#'
.testKeytab <- function(conn, keytab_id) {
    request <- list(
        keytab_id = keytab_id
    )

    .steamRPC(conn, "TestKeytab", request)
}

#' Delete the keytab entry for the given user
#'
#' @param keytab_id \code{numeric} No description available
#'
.deleteKeytab <- function(conn, keytab_id) {
    request <- list(
        keytab_id = keytab_id
    )

    .steamRPC(conn, "DeleteKeytab", request)
}

#' Connect to a cluster
#'
#' @param address \code{character} No description available
#'
.registerCluster <- function(conn, address) {
    request <- list(
        address = address
    )

    .steamRPC(conn, "RegisterCluster", request)
}

#' Disconnect from a cluster
#'
#' @param cluster_id \code{numeric} No description available
#'
.unregisterCluster <- function(conn, cluster_id) {
    request <- list(
        cluster_id = cluster_id
    )

    .steamRPC(conn, "UnregisterCluster", request)
}

#' Start a cluster using Yarn
#'
#' @param cluster_name \code{character} No description available
#' @param engine_id \code{numeric} No description available
#' @param size \code{numeric} No description available
#' @param memory \code{character} No description available
#' @param secure \code{logical} No description available
#' @param keytab \code{character} No description available
#'
.startClusterOnYarn <- function(conn, cluster_name, engine_id, size, memory, secure, keytab) {
    request <- list(
        cluster_name = cluster_name,
        engine_id = engine_id,
        size = size,
        memory = memory,
        secure = secure,
        keytab = keytab
    )

    .steamRPC(conn, "StartClusterOnYarn", request)
}

#' Stop a cluster using Yarn
#'
#' @param cluster_id \code{numeric} No description available
#' @param keytab \code{character} No description available
#'
.stopClusterOnYarn <- function(conn, cluster_id, keytab) {
    request <- list(
        cluster_id = cluster_id,
        keytab = keytab
    )

    .steamRPC(conn, "StopClusterOnYarn", request)
}

#' Get cluster details
#'
#' @param cluster_id \code{numeric} No description available
#'
.getCluster <- function(conn, cluster_id) {
    request <- list(
        cluster_id = cluster_id
    )

    .steamRPC(conn, "GetCluster", request)
}

#' Get cluster details (Yarn only)
#'
#' @param cluster_id \code{numeric} No description available
#'
.getClusterOnYarn <- function(conn, cluster_id) {
    request <- list(
        cluster_id = cluster_id
    )

    .steamRPC(conn, "GetClusterOnYarn", request)
}

#' List clusters
#'
#' @param offset \code{numeric} No description available
#' @param limit \code{numeric} No description available
#'
.getClusters <- function(conn, offset, limit) {
    request <- list(
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "GetClusters", request)
}

#' Get cluster status
#'
#' @param cluster_id \code{numeric} No description available
#'
.getClusterStatus <- function(conn, cluster_id) {
    request <- list(
        cluster_id = cluster_id
    )

    .steamRPC(conn, "GetClusterStatus", request)
}

#' Delete a cluster
#'
#' @param cluster_id \code{numeric} No description available
#'
.deleteCluster <- function(conn, cluster_id) {
    request <- list(
        cluster_id = cluster_id
    )

    .steamRPC(conn, "DeleteCluster", request)
}

#' Get job details
#'
#' @param cluster_id \code{numeric} No description available
#' @param job_name \code{character} No description available
#'
.getJob <- function(conn, cluster_id, job_name) {
    request <- list(
        cluster_id = cluster_id,
        job_name = job_name
    )

    .steamRPC(conn, "GetJob", request)
}

#' List jobs
#'
#' @param cluster_id \code{numeric} No description available
#'
.getJobs <- function(conn, cluster_id) {
    request <- list(
        cluster_id = cluster_id
    )

    .steamRPC(conn, "GetJobs", request)
}

#' Create a project
#'
#' @param name \code{character} No description available
#' @param description \code{character} No description available
#' @param model_category \code{character} No description available
#'
.createProject <- function(conn, name, description, model_category) {
    request <- list(
        name = name,
        description = description,
        model_category = model_category
    )

    .steamRPC(conn, "CreateProject", request)
}

#' List projects
#'
#' @param offset \code{numeric} No description available
#' @param limit \code{numeric} No description available
#'
.getProjects <- function(conn, offset, limit) {
    request <- list(
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "GetProjects", request)
}

#' Get project details
#'
#' @param project_id \code{numeric} No description available
#'
.getProject <- function(conn, project_id) {
    request <- list(
        project_id = project_id
    )

    .steamRPC(conn, "GetProject", request)
}

#' Delete a project
#'
#' @param project_id \code{numeric} No description available
#'
.deleteProject <- function(conn, project_id) {
    request <- list(
        project_id = project_id
    )

    .steamRPC(conn, "DeleteProject", request)
}

#' Create a datasource
#'
#' @param project_id \code{numeric} No description available
#' @param name \code{character} No description available
#' @param description \code{character} No description available
#' @param path \code{character} No description available
#'
.createDatasource <- function(conn, project_id, name, description, path) {
    request <- list(
        project_id = project_id,
        name = name,
        description = description,
        path = path
    )

    .steamRPC(conn, "CreateDatasource", request)
}

#' List datasources
#'
#' @param project_id \code{numeric} No description available
#' @param offset \code{numeric} No description available
#' @param limit \code{numeric} No description available
#'
.getDatasources <- function(conn, project_id, offset, limit) {
    request <- list(
        project_id = project_id,
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "GetDatasources", request)
}

#' Get datasource details
#'
#' @param datasource_id \code{numeric} No description available
#'
.getDatasource <- function(conn, datasource_id) {
    request <- list(
        datasource_id = datasource_id
    )

    .steamRPC(conn, "GetDatasource", request)
}

#' Update a datasource
#'
#' @param datasource_id \code{numeric} No description available
#' @param name \code{character} No description available
#' @param description \code{character} No description available
#' @param path \code{character} No description available
#'
.updateDatasource <- function(conn, datasource_id, name, description, path) {
    request <- list(
        datasource_id = datasource_id,
        name = name,
        description = description,
        path = path
    )

    .steamRPC(conn, "UpdateDatasource", request)
}

#' Delete a datasource
#'
#' @param datasource_id \code{numeric} No description available
#'
.deleteDatasource <- function(conn, datasource_id) {
    request <- list(
        datasource_id = datasource_id
    )

    .steamRPC(conn, "DeleteDatasource", request)
}

#' Create a dataset
#'
#' @param cluster_id \code{numeric} No description available
#' @param datasource_id \code{numeric} No description available
#' @param name \code{character} No description available
#' @param description \code{character} No description available
#' @param response_column_name \code{character} No description available
#'
.createDataset <- function(conn, cluster_id, datasource_id, name, description, response_column_name) {
    request <- list(
        cluster_id = cluster_id,
        datasource_id = datasource_id,
        name = name,
        description = description,
        response_column_name = response_column_name
    )

    .steamRPC(conn, "CreateDataset", request)
}

#' List datasets
#'
#' @param datasource_id \code{numeric} No description available
#' @param offset \code{numeric} No description available
#' @param limit \code{numeric} No description available
#'
.getDatasets <- function(conn, datasource_id, offset, limit) {
    request <- list(
        datasource_id = datasource_id,
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "GetDatasets", request)
}

#' Get dataset details
#'
#' @param dataset_id \code{numeric} No description available
#'
.getDataset <- function(conn, dataset_id) {
    request <- list(
        dataset_id = dataset_id
    )

    .steamRPC(conn, "GetDataset", request)
}

#' Get a list of datasets on a cluster
#'
#' @param cluster_id \code{numeric} No description available
#'
.getDatasetsFromCluster <- function(conn, cluster_id) {
    request <- list(
        cluster_id = cluster_id
    )

    .steamRPC(conn, "GetDatasetsFromCluster", request)
}

#' Update a dataset
#'
#' @param dataset_id \code{numeric} No description available
#' @param name \code{character} No description available
#' @param description \code{character} No description available
#' @param response_column_name \code{character} No description available
#'
.updateDataset <- function(conn, dataset_id, name, description, response_column_name) {
    request <- list(
        dataset_id = dataset_id,
        name = name,
        description = description,
        response_column_name = response_column_name
    )

    .steamRPC(conn, "UpdateDataset", request)
}

#' Split a dataset
#'
#' @param dataset_id \code{numeric} No description available
#' @param ratio1 \code{numeric} No description available
#' @param ratio2 \code{numeric} No description available
#'
.splitDataset <- function(conn, dataset_id, ratio1, ratio2) {
    request <- list(
        dataset_id = dataset_id,
        ratio1 = ratio1,
        ratio2 = ratio2
    )

    .steamRPC(conn, "SplitDataset", request)
}

#' Delete a dataset
#'
#' @param dataset_id \code{numeric} No description available
#'
.deleteDataset <- function(conn, dataset_id) {
    request <- list(
        dataset_id = dataset_id
    )

    .steamRPC(conn, "DeleteDataset", request)
}

#' Build a model
#'
#' @param cluster_id \code{numeric} No description available
#' @param dataset_id \code{numeric} No description available
#' @param algorithm \code{character} No description available
#'
.buildModel <- function(conn, cluster_id, dataset_id, algorithm) {
    request <- list(
        cluster_id = cluster_id,
        dataset_id = dataset_id,
        algorithm = algorithm
    )

    .steamRPC(conn, "BuildModel", request)
}

#' Build an AutoML model
#'
#' @param cluster_id \code{numeric} No description available
#' @param dataset \code{character} No description available
#' @param target_name \code{character} No description available
#' @param max_run_time \code{numeric} No description available
#'
.buildModelAuto <- function(conn, cluster_id, dataset, target_name, max_run_time) {
    request <- list(
        cluster_id = cluster_id,
        dataset = dataset,
        target_name = target_name,
        max_run_time = max_run_time
    )

    .steamRPC(conn, "BuildModelAuto", request)
}

#' Get model details
#'
#' @param model_id \code{numeric} No description available
#'
.getModel <- function(conn, model_id) {
    request <- list(
        model_id = model_id
    )

    .steamRPC(conn, "GetModel", request)
}

#' List models
#'
#' @param project_id \code{numeric} No description available
#' @param offset \code{numeric} No description available
#' @param limit \code{numeric} No description available
#'
.getModels <- function(conn, project_id, offset, limit) {
    request <- list(
        project_id = project_id,
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "GetModels", request)
}

#' List models from a cluster
#'
#' @param cluster_id \code{numeric} No description available
#' @param frame_key \code{character} No description available
#'
.getModelsFromCluster <- function(conn, cluster_id, frame_key) {
    request <- list(
        cluster_id = cluster_id,
        frame_key = frame_key
    )

    .steamRPC(conn, "GetModelsFromCluster", request)
}

#' Get a count models in a project
#'
#' @param project_id \code{numeric} No description available
#'
.findModelsCount <- function(conn, project_id) {
    request <- list(
        project_id = project_id
    )

    .steamRPC(conn, "FindModelsCount", request)
}

#' List sort criteria for a binomial models
#'
#'
.getAllBinomialSortCriteria <- function(conn) {
    request <- list(
    )

    .steamRPC(conn, "GetAllBinomialSortCriteria", request)
}

#' List binomial models
#'
#' @param project_id \code{numeric} No description available
#' @param name_part \code{character} No description available
#' @param sort_by \code{character} No description available
#' @param ascending \code{logical} No description available
#' @param offset \code{numeric} No description available
#' @param limit \code{numeric} No description available
#'
.findModelsBinomial <- function(conn, project_id, name_part, sort_by, ascending, offset, limit) {
    request <- list(
        project_id = project_id,
        name_part = name_part,
        sort_by = sort_by,
        ascending = ascending,
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "FindModelsBinomial", request)
}

#' View a binomial model
#'
#' @param model_id \code{numeric} No description available
#'
.getModelBinomial <- function(conn, model_id) {
    request <- list(
        model_id = model_id
    )

    .steamRPC(conn, "GetModelBinomial", request)
}

#' List sort criteria for a multinomial models
#'
#'
.getAllMultinomialSortCriteria <- function(conn) {
    request <- list(
    )

    .steamRPC(conn, "GetAllMultinomialSortCriteria", request)
}

#' List multinomial models
#'
#' @param project_id \code{numeric} No description available
#' @param name_part \code{character} No description available
#' @param sort_by \code{character} No description available
#' @param ascending \code{logical} No description available
#' @param offset \code{numeric} No description available
#' @param limit \code{numeric} No description available
#'
.findModelsMultinomial <- function(conn, project_id, name_part, sort_by, ascending, offset, limit) {
    request <- list(
        project_id = project_id,
        name_part = name_part,
        sort_by = sort_by,
        ascending = ascending,
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "FindModelsMultinomial", request)
}

#' View a binomial model
#'
#' @param model_id \code{numeric} No description available
#'
.getModelMultinomial <- function(conn, model_id) {
    request <- list(
        model_id = model_id
    )

    .steamRPC(conn, "GetModelMultinomial", request)
}

#' List sort criteria for a regression models
#'
#'
.getAllRegressionSortCriteria <- function(conn) {
    request <- list(
    )

    .steamRPC(conn, "GetAllRegressionSortCriteria", request)
}

#' List regression models
#'
#' @param project_id \code{numeric} No description available
#' @param name_part \code{character} No description available
#' @param sort_by \code{character} No description available
#' @param ascending \code{logical} No description available
#' @param offset \code{numeric} No description available
#' @param limit \code{numeric} No description available
#'
.findModelsRegression <- function(conn, project_id, name_part, sort_by, ascending, offset, limit) {
    request <- list(
        project_id = project_id,
        name_part = name_part,
        sort_by = sort_by,
        ascending = ascending,
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "FindModelsRegression", request)
}

#' View a binomial model
#'
#' @param model_id \code{numeric} No description available
#'
.getModelRegression <- function(conn, model_id) {
    request <- list(
        model_id = model_id
    )

    .steamRPC(conn, "GetModelRegression", request)
}

#' Import models from a cluster
#'
#' @param cluster_id \code{numeric} No description available
#' @param project_id \code{numeric} No description available
#' @param model_key \code{character} No description available
#' @param model_name \code{character} No description available
#'
.importModelFromCluster <- function(conn, cluster_id, project_id, model_key, model_name) {
    request <- list(
        cluster_id = cluster_id,
        project_id = project_id,
        model_key = model_key,
        model_name = model_name
    )

    .steamRPC(conn, "ImportModelFromCluster", request)
}

#' Check if a model category can generate MOJOs
#'
#' @param algo \code{character} No description available
#'
.checkMojo <- function(conn, algo) {
    request <- list(
        algo = algo
    )

    .steamRPC(conn, "CheckMojo", request)
}

#' Import a model's POJO from a cluster
#'
#' @param model_id \code{numeric} No description available
#'
.importModelPojo <- function(conn, model_id) {
    request <- list(
        model_id = model_id
    )

    .steamRPC(conn, "ImportModelPojo", request)
}

#' Import a model's MOJO from a cluster
#'
#' @param model_id \code{numeric} No description available
#'
.importModelMojo <- function(conn, model_id) {
    request <- list(
        model_id = model_id
    )

    .steamRPC(conn, "ImportModelMojo", request)
}

#' Delete a model
#'
#' @param model_id \code{numeric} No description available
#'
.deleteModel <- function(conn, model_id) {
    request <- list(
        model_id = model_id
    )

    .steamRPC(conn, "DeleteModel", request)
}

#' Create a label
#'
#' @param project_id \code{numeric} No description available
#' @param name \code{character} No description available
#' @param description \code{character} No description available
#'
.createLabel <- function(conn, project_id, name, description) {
    request <- list(
        project_id = project_id,
        name = name,
        description = description
    )

    .steamRPC(conn, "CreateLabel", request)
}

#' Update a label
#'
#' @param label_id \code{numeric} No description available
#' @param name \code{character} No description available
#' @param description \code{character} No description available
#'
.updateLabel <- function(conn, label_id, name, description) {
    request <- list(
        label_id = label_id,
        name = name,
        description = description
    )

    .steamRPC(conn, "UpdateLabel", request)
}

#' Delete a label
#'
#' @param label_id \code{numeric} No description available
#'
.deleteLabel <- function(conn, label_id) {
    request <- list(
        label_id = label_id
    )

    .steamRPC(conn, "DeleteLabel", request)
}

#' Label a model
#'
#' @param label_id \code{numeric} No description available
#' @param model_id \code{numeric} No description available
#'
.linkLabelWithModel <- function(conn, label_id, model_id) {
    request <- list(
        label_id = label_id,
        model_id = model_id
    )

    .steamRPC(conn, "LinkLabelWithModel", request)
}

#' Remove a label from a model
#'
#' @param label_id \code{numeric} No description available
#' @param model_id \code{numeric} No description available
#'
.unlinkLabelFromModel <- function(conn, label_id, model_id) {
    request <- list(
        label_id = label_id,
        model_id = model_id
    )

    .steamRPC(conn, "UnlinkLabelFromModel", request)
}

#' List labels for a project, with corresponding models, if any
#'
#' @param project_id \code{numeric} No description available
#'
.getLabelsForProject <- function(conn, project_id) {
    request <- list(
        project_id = project_id
    )

    .steamRPC(conn, "GetLabelsForProject", request)
}

#' Start a service
#'
#' @param model_id \code{numeric} No description available
#' @param name \code{character} No description available
#' @param package_name \code{character} No description available
#'
.startService <- function(conn, model_id, name, package_name) {
    request <- list(
        model_id = model_id,
        name = name,
        package_name = package_name
    )

    .steamRPC(conn, "StartService", request)
}

#' Stop a service
#'
#' @param service_id \code{numeric} No description available
#'
.stopService <- function(conn, service_id) {
    request <- list(
        service_id = service_id
    )

    .steamRPC(conn, "StopService", request)
}

#' Get service details
#'
#' @param service_id \code{numeric} No description available
#'
.getService <- function(conn, service_id) {
    request <- list(
        service_id = service_id
    )

    .steamRPC(conn, "GetService", request)
}

#' List all services
#'
#' @param offset \code{numeric} No description available
#' @param limit \code{numeric} No description available
#'
.getServices <- function(conn, offset, limit) {
    request <- list(
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "GetServices", request)
}

#' List services for a project
#'
#' @param project_id \code{numeric} No description available
#' @param offset \code{numeric} No description available
#' @param limit \code{numeric} No description available
#'
.getServicesForProject <- function(conn, project_id, offset, limit) {
    request <- list(
        project_id = project_id,
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "GetServicesForProject", request)
}

#' List services for a model
#'
#' @param model_id \code{numeric} No description available
#' @param offset \code{numeric} No description available
#' @param limit \code{numeric} No description available
#'
.getServicesForModel <- function(conn, model_id, offset, limit) {
    request <- list(
        model_id = model_id,
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "GetServicesForModel", request)
}

#' Delete a service
#'
#' @param service_id \code{numeric} No description available
#'
.deleteService <- function(conn, service_id) {
    request <- list(
        service_id = service_id
    )

    .steamRPC(conn, "DeleteService", request)
}

#' Get engine details
#'
#' @param engine_id \code{numeric} No description available
#'
.getEngine <- function(conn, engine_id) {
    request <- list(
        engine_id = engine_id
    )

    .steamRPC(conn, "GetEngine", request)
}

#' Get an engine by a version substring
#'
#' @param version \code{character} No description available
#'
.getEngineByVersion <- function(conn, version) {
    request <- list(
        version = version
    )

    .steamRPC(conn, "GetEngineByVersion", request)
}

#' List engines
#'
#'
.getEngines <- function(conn) {
    request <- list(
    )

    .steamRPC(conn, "GetEngines", request)
}

#' Delete an engine
#'
#' @param engine_id \code{numeric} No description available
#'
.deleteEngine <- function(conn, engine_id) {
    request <- list(
        engine_id = engine_id
    )

    .steamRPC(conn, "DeleteEngine", request)
}

#' List all entity types
#'
#'
.getAllEntityTypes <- function(conn) {
    request <- list(
    )

    .steamRPC(conn, "GetAllEntityTypes", request)
}

#' List all permissions
#'
#'
.getAllPermissions <- function(conn) {
    request <- list(
    )

    .steamRPC(conn, "GetAllPermissions", request)
}

#' List all cluster types
#'
#'
.getAllClusterTypes <- function(conn) {
    request <- list(
    )

    .steamRPC(conn, "GetAllClusterTypes", request)
}

#' List permissions for a role
#'
#' @param role_id \code{numeric} Integer ID of a role in Steam.
#'
.getPermissionsForRole <- function(conn, role_id) {
    request <- list(
        role_id = role_id
    )

    .steamRPC(conn, "GetPermissionsForRole", request)
}

#' List permissions for an identity
#'
#' @param identity_id \code{numeric} Integer ID of an identity in Steam.
#'
.getPermissionsForIdentity <- function(conn, identity_id) {
    request <- list(
        identity_id = identity_id
    )

    .steamRPC(conn, "GetPermissionsForIdentity", request)
}

#' Create a role
#'
#' @param name \code{character} A string name.
#' @param description \code{character} A string description
#'
.createRole <- function(conn, name, description) {
    request <- list(
        name = name,
        description = description
    )

    .steamRPC(conn, "CreateRole", request)
}

#' List roles
#'
#' @param offset \code{numeric} An offset uint start the search on.
#' @param limit \code{numeric} The maximum uint objects.
#'
.getRoles <- function(conn, offset, limit) {
    request <- list(
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "GetRoles", request)
}

#' List roles for an identity
#'
#' @param identity_id \code{numeric} Integer ID of an identity in Steam.
#'
.getRolesForIdentity <- function(conn, identity_id) {
    request <- list(
        identity_id = identity_id
    )

    .steamRPC(conn, "GetRolesForIdentity", request)
}

#' Get role details
#'
#' @param role_id \code{numeric} Integer ID of a role in Steam.
#'
.getRole <- function(conn, role_id) {
    request <- list(
        role_id = role_id
    )

    .steamRPC(conn, "GetRole", request)
}

#' Get role details by name
#'
#' @param name \code{character} A role name.
#'
.getRoleByName <- function(conn, name) {
    request <- list(
        name = name
    )

    .steamRPC(conn, "GetRoleByName", request)
}

#' Update a role
#'
#' @param role_id \code{numeric} Integer ID of a role in Steam.
#' @param name \code{character} A string name.
#' @param description \code{character} A string description
#'
.updateRole <- function(conn, role_id, name, description) {
    request <- list(
        role_id = role_id,
        name = name,
        description = description
    )

    .steamRPC(conn, "UpdateRole", request)
}

#' Link a role with permissions
#'
#' @param role_id \code{numeric} Integer ID of a role in Steam.
#' @param permission_ids \code{numeric} A list of Integer IDs for permissions in Steam.
#'
.linkRoleWithPermissions <- function(conn, role_id, permission_ids) {
    request <- list(
        role_id = role_id,
        permission_ids = permission_ids
    )

    .steamRPC(conn, "LinkRoleWithPermissions", request)
}

#' Link a role with a permission
#'
#' @param role_id \code{numeric} Integer ID of a role in Steam.
#' @param permission_id \code{numeric} Integer ID of a permission in Steam.
#'
.linkRoleWithPermission <- function(conn, role_id, permission_id) {
    request <- list(
        role_id = role_id,
        permission_id = permission_id
    )

    .steamRPC(conn, "LinkRoleWithPermission", request)
}

#' Unlink a role from a permission
#'
#' @param role_id \code{numeric} Integer ID of a role in Steam.
#' @param permission_id \code{numeric} Integer ID of a permission in Steam.
#'
.unlinkRoleFromPermission <- function(conn, role_id, permission_id) {
    request <- list(
        role_id = role_id,
        permission_id = permission_id
    )

    .steamRPC(conn, "UnlinkRoleFromPermission", request)
}

#' Delete a role
#'
#' @param role_id \code{numeric} Integer ID of a role in Steam.
#'
.deleteRole <- function(conn, role_id) {
    request <- list(
        role_id = role_id
    )

    .steamRPC(conn, "DeleteRole", request)
}

#' Create a workgroup
#'
#' @param name \code{character} A string name.
#' @param description \code{character} A string description
#'
.createWorkgroup <- function(conn, name, description) {
    request <- list(
        name = name,
        description = description
    )

    .steamRPC(conn, "CreateWorkgroup", request)
}

#' List workgroups
#'
#' @param offset \code{numeric} An offset uint start the search on.
#' @param limit \code{numeric} The maximum uint objects.
#'
.getWorkgroups <- function(conn, offset, limit) {
    request <- list(
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "GetWorkgroups", request)
}

#' List workgroups for an identity
#'
#' @param identity_id \code{numeric} Integer ID of an identity in Steam.
#'
.getWorkgroupsForIdentity <- function(conn, identity_id) {
    request <- list(
        identity_id = identity_id
    )

    .steamRPC(conn, "GetWorkgroupsForIdentity", request)
}

#' Get workgroup details
#'
#' @param workgroup_id \code{numeric} Integer ID of a workgroup in Steam.
#'
.getWorkgroup <- function(conn, workgroup_id) {
    request <- list(
        workgroup_id = workgroup_id
    )

    .steamRPC(conn, "GetWorkgroup", request)
}

#' Get workgroup details by name
#'
#' @param name \code{character} A string name.
#'
.getWorkgroupByName <- function(conn, name) {
    request <- list(
        name = name
    )

    .steamRPC(conn, "GetWorkgroupByName", request)
}

#' Update a workgroup
#'
#' @param workgroup_id \code{numeric} Integer ID of a workgrou in Steam.
#' @param name \code{character} A string name.
#' @param description \code{character} A string description
#'
.updateWorkgroup <- function(conn, workgroup_id, name, description) {
    request <- list(
        workgroup_id = workgroup_id,
        name = name,
        description = description
    )

    .steamRPC(conn, "UpdateWorkgroup", request)
}

#' Delete a workgroup
#'
#' @param workgroup_id \code{numeric} Integer ID of a workgroup in Steam.
#'
.deleteWorkgroup <- function(conn, workgroup_id) {
    request <- list(
        workgroup_id = workgroup_id
    )

    .steamRPC(conn, "DeleteWorkgroup", request)
}

#' Create an identity
#'
#' @param name \code{character} A string name.
#' @param password \code{character} A string password
#'
.createIdentity <- function(conn, name, password) {
    request <- list(
        name = name,
        password = password
    )

    .steamRPC(conn, "CreateIdentity", request)
}

#' List identities
#'
#' @param offset \code{numeric} An offset uint start the search on.
#' @param limit \code{numeric} The maximum uint objects.
#'
.getIdentities <- function(conn, offset, limit) {
    request <- list(
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "GetIdentities", request)
}

#' List identities for a workgroup
#'
#' @param workgroup_id \code{numeric} Integer ID of a workgroup in Steam.
#'
.getIdentitiesForWorkgroup <- function(conn, workgroup_id) {
    request <- list(
        workgroup_id = workgroup_id
    )

    .steamRPC(conn, "GetIdentitiesForWorkgroup", request)
}

#' List identities for a role
#'
#' @param role_id \code{numeric} Integer ID of a role in Steam.
#'
.getIdentitiesForRole <- function(conn, role_id) {
    request <- list(
        role_id = role_id
    )

    .steamRPC(conn, "GetIdentitiesForRole", request)
}

#' Get a list of identities and roles with access to an entity
#'
#' @param entity_type \code{numeric} An entity type ID.
#' @param entity_id \code{numeric} An entity ID.
#'
.getIdentitiesForEntity <- function(conn, entity_type, entity_id) {
    request <- list(
        entity_type = entity_type,
        entity_id = entity_id
    )

    .steamRPC(conn, "GetIdentitiesForEntity", request)
}

#' Get identity details
#'
#' @param identity_id \code{numeric} Integer ID of an identity in Steam.
#'
.getIdentity <- function(conn, identity_id) {
    request <- list(
        identity_id = identity_id
    )

    .steamRPC(conn, "GetIdentity", request)
}

#' Get identity details by name
#'
#' @param name \code{character} An identity name.
#'
.getIdentityByName <- function(conn, name) {
    request <- list(
        name = name
    )

    .steamRPC(conn, "GetIdentityByName", request)
}

#' Link an identity with a workgroup
#'
#' @param identity_id \code{numeric} Integer ID of an identity in Steam.
#' @param workgroup_id \code{numeric} Integer ID of a workgroup in Steam.
#'
.linkIdentityWithWorkgroup <- function(conn, identity_id, workgroup_id) {
    request <- list(
        identity_id = identity_id,
        workgroup_id = workgroup_id
    )

    .steamRPC(conn, "LinkIdentityWithWorkgroup", request)
}

#' Unlink an identity from a workgroup
#'
#' @param identity_id \code{numeric} Integer ID of an identity in Steam.
#' @param workgroup_id \code{numeric} Integer ID of a workgroup in Steam.
#'
.unlinkIdentityFromWorkgroup <- function(conn, identity_id, workgroup_id) {
    request <- list(
        identity_id = identity_id,
        workgroup_id = workgroup_id
    )

    .steamRPC(conn, "UnlinkIdentityFromWorkgroup", request)
}

#' Link an identity with a role
#'
#' @param identity_id \code{numeric} Integer ID of an identity in Steam.
#' @param role_id \code{numeric} Integer ID of a role in Steam.
#'
.linkIdentityWithRole <- function(conn, identity_id, role_id) {
    request <- list(
        identity_id = identity_id,
        role_id = role_id
    )

    .steamRPC(conn, "LinkIdentityWithRole", request)
}

#' Unlink an identity from a role
#'
#' @param identity_id \code{numeric} Integer ID of an identity in Steam.
#' @param role_id \code{numeric} Integer ID of a role in Steam.
#'
.unlinkIdentityFromRole <- function(conn, identity_id, role_id) {
    request <- list(
        identity_id = identity_id,
        role_id = role_id
    )

    .steamRPC(conn, "UnlinkIdentityFromRole", request)
}

#' Update an identity
#'
#' @param identity_id \code{numeric} Integer ID of an identity in Steam.
#' @param password \code{character} Password for identity
#'
.updateIdentity <- function(conn, identity_id, password) {
    request <- list(
        identity_id = identity_id,
        password = password
    )

    .steamRPC(conn, "UpdateIdentity", request)
}

#' Activate an identity
#'
#' @param identity_id \code{numeric} Integer ID of an identity in Steam.
#'
.activateIdentity <- function(conn, identity_id) {
    request <- list(
        identity_id = identity_id
    )

    .steamRPC(conn, "ActivateIdentity", request)
}

#' Deactivate an identity
#'
#' @param identity_id \code{numeric} Integer ID of an identity in Steam.
#'
.deactivateIdentity <- function(conn, identity_id) {
    request <- list(
        identity_id = identity_id
    )

    .steamRPC(conn, "DeactivateIdentity", request)
}

#' Share an entity with a workgroup
#'
#' @param kind \code{character} Type of permission. Can be view, edit, or own.
#' @param workgroup_id \code{numeric} Integer ID of a workgroup in Steam.
#' @param entity_type_id \code{numeric} Integer ID for the type of entity.
#' @param entity_id \code{numeric} Integer ID for an entity in Steam.
#'
.shareEntity <- function(conn, kind, workgroup_id, entity_type_id, entity_id) {
    request <- list(
        kind = kind,
        workgroup_id = workgroup_id,
        entity_type_id = entity_type_id,
        entity_id = entity_id
    )

    .steamRPC(conn, "ShareEntity", request)
}

#' List privileges for an entity
#'
#' @param entity_type_id \code{numeric} Integer ID for the type of entity.
#' @param entity_id \code{numeric} Integer ID for an entity in Steam.
#'
.getPrivileges <- function(conn, entity_type_id, entity_id) {
    request <- list(
        entity_type_id = entity_type_id,
        entity_id = entity_id
    )

    .steamRPC(conn, "GetPrivileges", request)
}

#' Unshare an entity
#'
#' @param kind \code{character} Type of permission. Can be view, edit, or own.
#' @param workgroup_id \code{numeric} Integer ID of a workgroup in Steam.
#' @param entity_type_id \code{numeric} Integer ID for the type of entity.
#' @param entity_id \code{numeric} Integer ID for an entity in Steam.
#'
.unshareEntity <- function(conn, kind, workgroup_id, entity_type_id, entity_id) {
    request <- list(
        kind = kind,
        workgroup_id = workgroup_id,
        entity_type_id = entity_type_id,
        entity_id = entity_id
    )

    .steamRPC(conn, "UnshareEntity", request)
}

#' List audit trail records for an entity
#'
#' @param entity_type_id \code{numeric} Integer ID for the type of entity.
#' @param entity_id \code{numeric} Integer ID for an entity in Steam.
#' @param offset \code{numeric} An offset uint start the search on.
#' @param limit \code{numeric} The maximum uint objects.
#'
.getHistory <- function(conn, entity_type_id, entity_id, offset, limit) {
    request <- list(
        entity_type_id = entity_type_id,
        entity_id = entity_id,
        offset = offset,
        limit = limit
    )

    .steamRPC(conn, "GetHistory", request)
}

#' Create a package for a project
#'
#' @param project_id \code{numeric} No description available
#' @param name \code{character} No description available
#'
.createPackage <- function(conn, project_id, name) {
    request <- list(
        project_id = project_id,
        name = name
    )

    .steamRPC(conn, "CreatePackage", request)
}

#' List packages for a project 
#'
#' @param project_id \code{numeric} No description available
#'
.getPackages <- function(conn, project_id) {
    request <- list(
        project_id = project_id
    )

    .steamRPC(conn, "GetPackages", request)
}

#' List directories in a project package
#'
#' @param project_id \code{numeric} No description available
#' @param package_name \code{character} No description available
#' @param relative_path \code{character} No description available
#'
.getPackageDirectories <- function(conn, project_id, package_name, relative_path) {
    request <- list(
        project_id = project_id,
        package_name = package_name,
        relative_path = relative_path
    )

    .steamRPC(conn, "GetPackageDirectories", request)
}

#' List files in a project package
#'
#' @param project_id \code{numeric} No description available
#' @param package_name \code{character} No description available
#' @param relative_path \code{character} No description available
#'
.getPackageFiles <- function(conn, project_id, package_name, relative_path) {
    request <- list(
        project_id = project_id,
        package_name = package_name,
        relative_path = relative_path
    )

    .steamRPC(conn, "GetPackageFiles", request)
}

#' Delete a project package
#'
#' @param project_id \code{numeric} No description available
#' @param name \code{character} No description available
#'
.deletePackage <- function(conn, project_id, name) {
    request <- list(
        project_id = project_id,
        name = name
    )

    .steamRPC(conn, "DeletePackage", request)
}

#' Delete a directory in a project package
#'
#' @param project_id \code{numeric} No description available
#' @param package_name \code{character} No description available
#' @param relative_path \code{character} No description available
#'
.deletePackageDirectory <- function(conn, project_id, package_name, relative_path) {
    request <- list(
        project_id = project_id,
        package_name = package_name,
        relative_path = relative_path
    )

    .steamRPC(conn, "DeletePackageDirectory", request)
}

#' Delete a file in a project package
#'
#' @param project_id \code{numeric} No description available
#' @param package_name \code{character} No description available
#' @param relative_path \code{character} No description available
#'
.deletePackageFile <- function(conn, project_id, package_name, relative_path) {
    request <- list(
        project_id = project_id,
        package_name = package_name,
        relative_path = relative_path
    )

    .steamRPC(conn, "DeletePackageFile", request)
}

#' Set attributes on a project package
#'
#' @param project_id \code{numeric} No description available
#' @param package_name \code{character} No description available
#' @param attributes \code{character} No description available
#'
.setAttributesForPackage <- function(conn, project_id, package_name, attributes) {
    request <- list(
        project_id = project_id,
        package_name = package_name,
        attributes = attributes
    )

    .steamRPC(conn, "SetAttributesForPackage", request)
}

#' List attributes for a project package
#'
#' @param project_id \code{numeric} No description available
#' @param package_name \code{character} No description available
#'
.getAttributesForPackage <- function(conn, project_id, package_name) {
    request <- list(
        project_id = project_id,
        package_name = package_name
    )

    .steamRPC(conn, "GetAttributesForPackage", request)
}

