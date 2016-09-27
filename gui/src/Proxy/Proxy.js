/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
"use strict";
// ------------------------------
// --- This is generated code ---
// ---      DO NOT EDIT       ---
// ------------------------------
// --- Types ---
var Proxy = require('./xhr');
// --- Client Stub ---
function pingServer(input, go) {
    var req = { input: input };
    Proxy.Call("PingServer", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.output);
        }
    });
}
exports.pingServer = pingServer;
function getConfig(go) {
    var req = {};
    Proxy.Call("GetConfig", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.config);
        }
    });
}
exports.getConfig = getConfig;
function registerCluster(address, go) {
    var req = { address: address };
    Proxy.Call("RegisterCluster", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.cluster_id);
        }
    });
}
exports.registerCluster = registerCluster;
function unregisterCluster(clusterId, go) {
    var req = { cluster_id: clusterId };
    Proxy.Call("UnregisterCluster", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.unregisterCluster = unregisterCluster;
function startClusterOnYarn(clusterName, engineId, size, memory, keytab, go) {
    var req = { cluster_name: clusterName, engine_id: engineId, size: size, memory: memory, keytab: keytab };
    Proxy.Call("StartClusterOnYarn", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.cluster_id);
        }
    });
}
exports.startClusterOnYarn = startClusterOnYarn;
function stopClusterOnYarn(clusterId, keytab, go) {
    var req = { cluster_id: clusterId, keytab: keytab };
    Proxy.Call("StopClusterOnYarn", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.stopClusterOnYarn = stopClusterOnYarn;
function getCluster(clusterId, go) {
    var req = { cluster_id: clusterId };
    Proxy.Call("GetCluster", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.cluster);
        }
    });
}
exports.getCluster = getCluster;
function getClusterOnYarn(clusterId, go) {
    var req = { cluster_id: clusterId };
    Proxy.Call("GetClusterOnYarn", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.cluster);
        }
    });
}
exports.getClusterOnYarn = getClusterOnYarn;
function getClusters(offset, limit, go) {
    var req = { offset: offset, limit: limit };
    Proxy.Call("GetClusters", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.clusters);
        }
    });
}
exports.getClusters = getClusters;
function getClusterStatus(clusterId, go) {
    var req = { cluster_id: clusterId };
    Proxy.Call("GetClusterStatus", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.cluster_status);
        }
    });
}
exports.getClusterStatus = getClusterStatus;
function deleteCluster(clusterId, go) {
    var req = { cluster_id: clusterId };
    Proxy.Call("DeleteCluster", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.deleteCluster = deleteCluster;
function getJob(clusterId, jobName, go) {
    var req = { cluster_id: clusterId, job_name: jobName };
    Proxy.Call("GetJob", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.job);
        }
    });
}
exports.getJob = getJob;
function getJobs(clusterId, go) {
    var req = { cluster_id: clusterId };
    Proxy.Call("GetJobs", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.jobs);
        }
    });
}
exports.getJobs = getJobs;
function createProject(name, description, modelCategory, go) {
    var req = { name: name, description: description, model_category: modelCategory };
    Proxy.Call("CreateProject", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.project_id);
        }
    });
}
exports.createProject = createProject;
function getProjects(offset, limit, go) {
    var req = { offset: offset, limit: limit };
    Proxy.Call("GetProjects", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.projects);
        }
    });
}
exports.getProjects = getProjects;
function getProject(projectId, go) {
    var req = { project_id: projectId };
    Proxy.Call("GetProject", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.project);
        }
    });
}
exports.getProject = getProject;
function deleteProject(projectId, go) {
    var req = { project_id: projectId };
    Proxy.Call("DeleteProject", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.deleteProject = deleteProject;
function createDatasource(projectId, name, description, path, go) {
    var req = { project_id: projectId, name: name, description: description, path: path };
    Proxy.Call("CreateDatasource", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.datasource_id);
        }
    });
}
exports.createDatasource = createDatasource;
function getDatasources(projectId, offset, limit, go) {
    var req = { project_id: projectId, offset: offset, limit: limit };
    Proxy.Call("GetDatasources", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.datasources);
        }
    });
}
exports.getDatasources = getDatasources;
function getDatasource(datasourceId, go) {
    var req = { datasource_id: datasourceId };
    Proxy.Call("GetDatasource", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.datasource);
        }
    });
}
exports.getDatasource = getDatasource;
function updateDatasource(datasourceId, name, description, path, go) {
    var req = { datasource_id: datasourceId, name: name, description: description, path: path };
    Proxy.Call("UpdateDatasource", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.updateDatasource = updateDatasource;
function deleteDatasource(datasourceId, go) {
    var req = { datasource_id: datasourceId };
    Proxy.Call("DeleteDatasource", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.deleteDatasource = deleteDatasource;
function createDataset(clusterId, datasourceId, name, description, responseColumnName, go) {
    var req = { cluster_id: clusterId, datasource_id: datasourceId, name: name, description: description, response_column_name: responseColumnName };
    Proxy.Call("CreateDataset", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.dataset_id);
        }
    });
}
exports.createDataset = createDataset;
function getDatasets(datasourceId, offset, limit, go) {
    var req = { datasource_id: datasourceId, offset: offset, limit: limit };
    Proxy.Call("GetDatasets", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.datasets);
        }
    });
}
exports.getDatasets = getDatasets;
function getDataset(datasetId, go) {
    var req = { dataset_id: datasetId };
    Proxy.Call("GetDataset", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.dataset);
        }
    });
}
exports.getDataset = getDataset;
function getDatasetsFromCluster(clusterId, go) {
    var req = { cluster_id: clusterId };
    Proxy.Call("GetDatasetsFromCluster", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.dataset);
        }
    });
}
exports.getDatasetsFromCluster = getDatasetsFromCluster;
function updateDataset(datasetId, name, description, responseColumnName, go) {
    var req = { dataset_id: datasetId, name: name, description: description, response_column_name: responseColumnName };
    Proxy.Call("UpdateDataset", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.updateDataset = updateDataset;
function splitDataset(datasetId, ratio1, ratio2, go) {
    var req = { dataset_id: datasetId, ratio1: ratio1, ratio2: ratio2 };
    Proxy.Call("SplitDataset", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.dataset_ids);
        }
    });
}
exports.splitDataset = splitDataset;
function deleteDataset(datasetId, go) {
    var req = { dataset_id: datasetId };
    Proxy.Call("DeleteDataset", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.deleteDataset = deleteDataset;
function buildModel(clusterId, datasetId, algorithm, go) {
    var req = { cluster_id: clusterId, dataset_id: datasetId, algorithm: algorithm };
    Proxy.Call("BuildModel", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.model_id);
        }
    });
}
exports.buildModel = buildModel;
function buildModelAuto(clusterId, dataset, targetName, maxRunTime, go) {
    var req = { cluster_id: clusterId, dataset: dataset, target_name: targetName, max_run_time: maxRunTime };
    Proxy.Call("BuildModelAuto", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.model);
        }
    });
}
exports.buildModelAuto = buildModelAuto;
function getModel(modelId, go) {
    var req = { model_id: modelId };
    Proxy.Call("GetModel", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.model);
        }
    });
}
exports.getModel = getModel;
function getModels(projectId, offset, limit, go) {
    var req = { project_id: projectId, offset: offset, limit: limit };
    Proxy.Call("GetModels", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.models);
        }
    });
}
exports.getModels = getModels;
function getModelsFromCluster(clusterId, frameKey, go) {
    var req = { cluster_id: clusterId, frame_key: frameKey };
    Proxy.Call("GetModelsFromCluster", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.models);
        }
    });
}
exports.getModelsFromCluster = getModelsFromCluster;
function findModelsCount(projectId, go) {
    var req = { project_id: projectId };
    Proxy.Call("FindModelsCount", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.count);
        }
    });
}
exports.findModelsCount = findModelsCount;
function getAllBinomialSortCriteria(go) {
    var req = {};
    Proxy.Call("GetAllBinomialSortCriteria", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.criteria);
        }
    });
}
exports.getAllBinomialSortCriteria = getAllBinomialSortCriteria;
function findModelsBinomial(projectId, namePart, sortBy, ascending, offset, limit, go) {
    var req = { project_id: projectId, name_part: namePart, sort_by: sortBy, ascending: ascending, offset: offset, limit: limit };
    Proxy.Call("FindModelsBinomial", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.models);
        }
    });
}
exports.findModelsBinomial = findModelsBinomial;
function getModelBinomial(modelId, go) {
    var req = { model_id: modelId };
    Proxy.Call("GetModelBinomial", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.model);
        }
    });
}
exports.getModelBinomial = getModelBinomial;
function getAllMultinomialSortCriteria(go) {
    var req = {};
    Proxy.Call("GetAllMultinomialSortCriteria", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.criteria);
        }
    });
}
exports.getAllMultinomialSortCriteria = getAllMultinomialSortCriteria;
function findModelsMultinomial(projectId, namePart, sortBy, ascending, offset, limit, go) {
    var req = { project_id: projectId, name_part: namePart, sort_by: sortBy, ascending: ascending, offset: offset, limit: limit };
    Proxy.Call("FindModelsMultinomial", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.models);
        }
    });
}
exports.findModelsMultinomial = findModelsMultinomial;
function getModelMultinomial(modelId, go) {
    var req = { model_id: modelId };
    Proxy.Call("GetModelMultinomial", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.model);
        }
    });
}
exports.getModelMultinomial = getModelMultinomial;
function getAllRegressionSortCriteria(go) {
    var req = {};
    Proxy.Call("GetAllRegressionSortCriteria", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.criteria);
        }
    });
}
exports.getAllRegressionSortCriteria = getAllRegressionSortCriteria;
function findModelsRegression(projectId, namePart, sortBy, ascending, offset, limit, go) {
    var req = { project_id: projectId, name_part: namePart, sort_by: sortBy, ascending: ascending, offset: offset, limit: limit };
    Proxy.Call("FindModelsRegression", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.models);
        }
    });
}
exports.findModelsRegression = findModelsRegression;
function getModelRegression(modelId, go) {
    var req = { model_id: modelId };
    Proxy.Call("GetModelRegression", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.model);
        }
    });
}
exports.getModelRegression = getModelRegression;
function importModelFromCluster(clusterId, projectId, modelKey, modelName, go) {
    var req = { cluster_id: clusterId, project_id: projectId, model_key: modelKey, model_name: modelName };
    Proxy.Call("ImportModelFromCluster", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.model_id);
        }
    });
}
exports.importModelFromCluster = importModelFromCluster;
function deleteModel(modelId, go) {
    var req = { model_id: modelId };
    Proxy.Call("DeleteModel", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.deleteModel = deleteModel;
function createLabel(projectId, name, description, go) {
    var req = { project_id: projectId, name: name, description: description };
    Proxy.Call("CreateLabel", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.label_id);
        }
    });
}
exports.createLabel = createLabel;
function updateLabel(labelId, name, description, go) {
    var req = { label_id: labelId, name: name, description: description };
    Proxy.Call("UpdateLabel", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.updateLabel = updateLabel;
function deleteLabel(labelId, go) {
    var req = { label_id: labelId };
    Proxy.Call("DeleteLabel", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.deleteLabel = deleteLabel;
function linkLabelWithModel(labelId, modelId, go) {
    var req = { label_id: labelId, model_id: modelId };
    Proxy.Call("LinkLabelWithModel", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.linkLabelWithModel = linkLabelWithModel;
function unlinkLabelFromModel(labelId, modelId, go) {
    var req = { label_id: labelId, model_id: modelId };
    Proxy.Call("UnlinkLabelFromModel", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.unlinkLabelFromModel = unlinkLabelFromModel;
function getLabelsForProject(projectId, go) {
    var req = { project_id: projectId };
    Proxy.Call("GetLabelsForProject", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.labels);
        }
    });
}
exports.getLabelsForProject = getLabelsForProject;
function startService(modelId, name, packageName, go) {
    var req = { model_id: modelId, name: name, package_name: packageName };
    Proxy.Call("StartService", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.service_id);
        }
    });
}
exports.startService = startService;
function stopService(serviceId, go) {
    var req = { service_id: serviceId };
    Proxy.Call("StopService", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.stopService = stopService;
function getService(serviceId, go) {
    var req = { service_id: serviceId };
    Proxy.Call("GetService", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.service);
        }
    });
}
exports.getService = getService;
function getServices(offset, limit, go) {
    var req = { offset: offset, limit: limit };
    Proxy.Call("GetServices", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.services);
        }
    });
}
exports.getServices = getServices;
function getServicesForProject(projectId, offset, limit, go) {
    var req = { project_id: projectId, offset: offset, limit: limit };
    Proxy.Call("GetServicesForProject", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.services);
        }
    });
}
exports.getServicesForProject = getServicesForProject;
function getServicesForModel(modelId, offset, limit, go) {
    var req = { model_id: modelId, offset: offset, limit: limit };
    Proxy.Call("GetServicesForModel", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.services);
        }
    });
}
exports.getServicesForModel = getServicesForModel;
function deleteService(serviceId, go) {
    var req = { service_id: serviceId };
    Proxy.Call("DeleteService", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.deleteService = deleteService;
function addEngine(engineName, enginePath, go) {
    var req = { engine_name: engineName, engine_path: enginePath };
    Proxy.Call("AddEngine", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.engine_id);
        }
    });
}
exports.addEngine = addEngine;
function getEngine(engineId, go) {
    var req = { engine_id: engineId };
    Proxy.Call("GetEngine", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.engine);
        }
    });
}
exports.getEngine = getEngine;
function getEngines(go) {
    var req = {};
    Proxy.Call("GetEngines", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.engines);
        }
    });
}
exports.getEngines = getEngines;
function deleteEngine(engineId, go) {
    var req = { engine_id: engineId };
    Proxy.Call("DeleteEngine", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.deleteEngine = deleteEngine;
function getAllEntityTypes(go) {
    var req = {};
    Proxy.Call("GetAllEntityTypes", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.entity_types);
        }
    });
}
exports.getAllEntityTypes = getAllEntityTypes;
function getAllPermissions(go) {
    var req = {};
    Proxy.Call("GetAllPermissions", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.permissions);
        }
    });
}
exports.getAllPermissions = getAllPermissions;
function getAllClusterTypes(go) {
    var req = {};
    Proxy.Call("GetAllClusterTypes", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.cluster_types);
        }
    });
}
exports.getAllClusterTypes = getAllClusterTypes;
function getPermissionsForRole(roleId, go) {
    var req = { role_id: roleId };
    Proxy.Call("GetPermissionsForRole", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.permissions);
        }
    });
}
exports.getPermissionsForRole = getPermissionsForRole;
function getPermissionsForIdentity(identityId, go) {
    var req = { identity_id: identityId };
    Proxy.Call("GetPermissionsForIdentity", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.permissions);
        }
    });
}
exports.getPermissionsForIdentity = getPermissionsForIdentity;
function createRole(name, description, go) {
    var req = { name: name, description: description };
    Proxy.Call("CreateRole", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.role_id);
        }
    });
}
exports.createRole = createRole;
function getRoles(offset, limit, go) {
    var req = { offset: offset, limit: limit };
    Proxy.Call("GetRoles", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.roles);
        }
    });
}
exports.getRoles = getRoles;
function getRolesForIdentity(identityId, go) {
    var req = { identity_id: identityId };
    Proxy.Call("GetRolesForIdentity", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.roles);
        }
    });
}
exports.getRolesForIdentity = getRolesForIdentity;
function getRole(roleId, go) {
    var req = { role_id: roleId };
    Proxy.Call("GetRole", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.role);
        }
    });
}
exports.getRole = getRole;
function getRoleByName(name, go) {
    var req = { name: name };
    Proxy.Call("GetRoleByName", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.role);
        }
    });
}
exports.getRoleByName = getRoleByName;
function updateRole(roleId, name, description, go) {
    var req = { role_id: roleId, name: name, description: description };
    Proxy.Call("UpdateRole", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.updateRole = updateRole;
function linkRoleWithPermissions(roleId, permissionIds, go) {
    var req = { role_id: roleId, permission_ids: permissionIds };
    Proxy.Call("LinkRoleWithPermissions", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.linkRoleWithPermissions = linkRoleWithPermissions;
function linkRoleWithPermission(roleId, permissionId, go) {
    var req = { role_id: roleId, permission_id: permissionId };
    Proxy.Call("LinkRoleWithPermission", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.linkRoleWithPermission = linkRoleWithPermission;
function unlinkRoleFromPermission(roleId, permissionId, go) {
    var req = { role_id: roleId, permission_id: permissionId };
    Proxy.Call("UnlinkRoleFromPermission", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.unlinkRoleFromPermission = unlinkRoleFromPermission;
function deleteRole(roleId, go) {
    var req = { role_id: roleId };
    Proxy.Call("DeleteRole", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.deleteRole = deleteRole;
function createWorkgroup(name, description, go) {
    var req = { name: name, description: description };
    Proxy.Call("CreateWorkgroup", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.workgroup_id);
        }
    });
}
exports.createWorkgroup = createWorkgroup;
function getWorkgroups(offset, limit, go) {
    var req = { offset: offset, limit: limit };
    Proxy.Call("GetWorkgroups", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.workgroups);
        }
    });
}
exports.getWorkgroups = getWorkgroups;
function getWorkgroupsForIdentity(identityId, go) {
    var req = { identity_id: identityId };
    Proxy.Call("GetWorkgroupsForIdentity", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.workgroups);
        }
    });
}
exports.getWorkgroupsForIdentity = getWorkgroupsForIdentity;
function getWorkgroup(workgroupId, go) {
    var req = { workgroup_id: workgroupId };
    Proxy.Call("GetWorkgroup", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.workgroup);
        }
    });
}
exports.getWorkgroup = getWorkgroup;
function getWorkgroupByName(name, go) {
    var req = { name: name };
    Proxy.Call("GetWorkgroupByName", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.workgroup);
        }
    });
}
exports.getWorkgroupByName = getWorkgroupByName;
function updateWorkgroup(workgroupId, name, description, go) {
    var req = { workgroup_id: workgroupId, name: name, description: description };
    Proxy.Call("UpdateWorkgroup", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.updateWorkgroup = updateWorkgroup;
function deleteWorkgroup(workgroupId, go) {
    var req = { workgroup_id: workgroupId };
    Proxy.Call("DeleteWorkgroup", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.deleteWorkgroup = deleteWorkgroup;
function createIdentity(name, password, go) {
    var req = { name: name, password: password };
    Proxy.Call("CreateIdentity", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.identity_id);
        }
    });
}
exports.createIdentity = createIdentity;
function getIdentities(offset, limit, go) {
    var req = { offset: offset, limit: limit };
    Proxy.Call("GetIdentities", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.identities);
        }
    });
}
exports.getIdentities = getIdentities;
function getIdentitiesForWorkgroup(workgroupId, go) {
    var req = { workgroup_id: workgroupId };
    Proxy.Call("GetIdentitiesForWorkgroup", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.identities);
        }
    });
}
exports.getIdentitiesForWorkgroup = getIdentitiesForWorkgroup;
function getIdentitiesForRole(roleId, go) {
    var req = { role_id: roleId };
    Proxy.Call("GetIdentitiesForRole", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.identities);
        }
    });
}
exports.getIdentitiesForRole = getIdentitiesForRole;
function getIdentitiesForEntity(entityType, entityId, go) {
    var req = { entity_type: entityType, entity_id: entityId };
    Proxy.Call("GetIdentitiesForEntity", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.users);
        }
    });
}
exports.getIdentitiesForEntity = getIdentitiesForEntity;
function getIdentity(identityId, go) {
    var req = { identity_id: identityId };
    Proxy.Call("GetIdentity", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.identity);
        }
    });
}
exports.getIdentity = getIdentity;
function getIdentityByName(name, go) {
    var req = { name: name };
    Proxy.Call("GetIdentityByName", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.identity);
        }
    });
}
exports.getIdentityByName = getIdentityByName;
function linkIdentityWithWorkgroup(identityId, workgroupId, go) {
    var req = { identity_id: identityId, workgroup_id: workgroupId };
    Proxy.Call("LinkIdentityWithWorkgroup", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.linkIdentityWithWorkgroup = linkIdentityWithWorkgroup;
function unlinkIdentityFromWorkgroup(identityId, workgroupId, go) {
    var req = { identity_id: identityId, workgroup_id: workgroupId };
    Proxy.Call("UnlinkIdentityFromWorkgroup", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.unlinkIdentityFromWorkgroup = unlinkIdentityFromWorkgroup;
function linkIdentityWithRole(identityId, roleId, go) {
    var req = { identity_id: identityId, role_id: roleId };
    Proxy.Call("LinkIdentityWithRole", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.linkIdentityWithRole = linkIdentityWithRole;
function unlinkIdentityFromRole(identityId, roleId, go) {
    var req = { identity_id: identityId, role_id: roleId };
    Proxy.Call("UnlinkIdentityFromRole", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.unlinkIdentityFromRole = unlinkIdentityFromRole;
function updateIdentity(identityId, password, go) {
    var req = { identity_id: identityId, password: password };
    Proxy.Call("UpdateIdentity", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.updateIdentity = updateIdentity;
function deactivateIdentity(identityId, go) {
    var req = { identity_id: identityId };
    Proxy.Call("DeactivateIdentity", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.deactivateIdentity = deactivateIdentity;
function shareEntity(kind, workgroupId, entityTypeId, entityId, go) {
    var req = { kind: kind, workgroup_id: workgroupId, entity_type_id: entityTypeId, entity_id: entityId };
    Proxy.Call("ShareEntity", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.shareEntity = shareEntity;
function getPrivileges(entityTypeId, entityId, go) {
    var req = { entity_type_id: entityTypeId, entity_id: entityId };
    Proxy.Call("GetPrivileges", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.privileges);
        }
    });
}
exports.getPrivileges = getPrivileges;
function unshareEntity(kind, workgroupId, entityTypeId, entityId, go) {
    var req = { kind: kind, workgroup_id: workgroupId, entity_type_id: entityTypeId, entity_id: entityId };
    Proxy.Call("UnshareEntity", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.unshareEntity = unshareEntity;
function getHistory(entityTypeId, entityId, offset, limit, go) {
    var req = { entity_type_id: entityTypeId, entity_id: entityId, offset: offset, limit: limit };
    Proxy.Call("GetHistory", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.history);
        }
    });
}
exports.getHistory = getHistory;
function createPackage(projectId, name, go) {
    var req = { project_id: projectId, name: name };
    Proxy.Call("CreatePackage", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.createPackage = createPackage;
function getPackages(projectId, go) {
    var req = { project_id: projectId };
    Proxy.Call("GetPackages", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.packages);
        }
    });
}
exports.getPackages = getPackages;
function getPackageDirectories(projectId, packageName, relativePath, go) {
    var req = { project_id: projectId, package_name: packageName, relative_path: relativePath };
    Proxy.Call("GetPackageDirectories", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.directories);
        }
    });
}
exports.getPackageDirectories = getPackageDirectories;
function getPackageFiles(projectId, packageName, relativePath, go) {
    var req = { project_id: projectId, package_name: packageName, relative_path: relativePath };
    Proxy.Call("GetPackageFiles", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.files);
        }
    });
}
exports.getPackageFiles = getPackageFiles;
function deletePackage(projectId, name, go) {
    var req = { project_id: projectId, name: name };
    Proxy.Call("DeletePackage", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.deletePackage = deletePackage;
function deletePackageDirectory(projectId, packageName, relativePath, go) {
    var req = { project_id: projectId, package_name: packageName, relative_path: relativePath };
    Proxy.Call("DeletePackageDirectory", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.deletePackageDirectory = deletePackageDirectory;
function deletePackageFile(projectId, packageName, relativePath, go) {
    var req = { project_id: projectId, package_name: packageName, relative_path: relativePath };
    Proxy.Call("DeletePackageFile", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.deletePackageFile = deletePackageFile;
function setAttributesForPackage(projectId, packageName, attributes, go) {
    var req = { project_id: projectId, package_name: packageName, attributes: attributes };
    Proxy.Call("SetAttributesForPackage", req, function (error, data) {
        if (error) {
            return go(error);
        }
        else {
            var d = data;
            return go(null);
        }
    });
}
exports.setAttributesForPackage = setAttributesForPackage;
function getAttributesForPackage(projectId, packageName, go) {
    var req = { project_id: projectId, package_name: packageName };
    Proxy.Call("GetAttributesForPackage", req, function (error, data) {
        if (error) {
            return go(error, null);
        }
        else {
            var d = data;
            return go(null, d.attributes);
        }
    });
}
exports.getAttributesForPackage = getAttributesForPackage;
//# sourceMappingURL=Proxy.js.map