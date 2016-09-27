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
// --- CLI Stub ---
var Proxy = require('./xhr');
function print(error, data) {
    if (error) {
        console.error(error);
    }
    else {
        console.log(data);
    }
}
function pingServer(input) {
    var req = { input: input };
    Proxy.Call("PingServer", req, print);
}
exports.pingServer = pingServer;
function getConfig() {
    var req = {};
    Proxy.Call("GetConfig", req, print);
}
exports.getConfig = getConfig;
function registerCluster(address) {
    var req = { address: address };
    Proxy.Call("RegisterCluster", req, print);
}
exports.registerCluster = registerCluster;
function unregisterCluster(clusterId) {
    var req = { cluster_id: clusterId };
    Proxy.Call("UnregisterCluster", req, print);
}
exports.unregisterCluster = unregisterCluster;
function startClusterOnYarn(clusterName, engineId, size, memory, keytab) {
    var req = { cluster_name: clusterName, engine_id: engineId, size: size, memory: memory, keytab: keytab };
    Proxy.Call("StartClusterOnYarn", req, print);
}
exports.startClusterOnYarn = startClusterOnYarn;
function stopClusterOnYarn(clusterId, keytab) {
    var req = { cluster_id: clusterId, keytab: keytab };
    Proxy.Call("StopClusterOnYarn", req, print);
}
exports.stopClusterOnYarn = stopClusterOnYarn;
function getCluster(clusterId) {
    var req = { cluster_id: clusterId };
    Proxy.Call("GetCluster", req, print);
}
exports.getCluster = getCluster;
function getClusterOnYarn(clusterId) {
    var req = { cluster_id: clusterId };
    Proxy.Call("GetClusterOnYarn", req, print);
}
exports.getClusterOnYarn = getClusterOnYarn;
function getClusters(offset, limit) {
    var req = { offset: offset, limit: limit };
    Proxy.Call("GetClusters", req, print);
}
exports.getClusters = getClusters;
function getClusterStatus(clusterId) {
    var req = { cluster_id: clusterId };
    Proxy.Call("GetClusterStatus", req, print);
}
exports.getClusterStatus = getClusterStatus;
function deleteCluster(clusterId) {
    var req = { cluster_id: clusterId };
    Proxy.Call("DeleteCluster", req, print);
}
exports.deleteCluster = deleteCluster;
function getJob(clusterId, jobName) {
    var req = { cluster_id: clusterId, job_name: jobName };
    Proxy.Call("GetJob", req, print);
}
exports.getJob = getJob;
function getJobs(clusterId) {
    var req = { cluster_id: clusterId };
    Proxy.Call("GetJobs", req, print);
}
exports.getJobs = getJobs;
function createProject(name, description, modelCategory) {
    var req = { name: name, description: description, model_category: modelCategory };
    Proxy.Call("CreateProject", req, print);
}
exports.createProject = createProject;
function getProjects(offset, limit) {
    var req = { offset: offset, limit: limit };
    Proxy.Call("GetProjects", req, print);
}
exports.getProjects = getProjects;
function getProject(projectId) {
    var req = { project_id: projectId };
    Proxy.Call("GetProject", req, print);
}
exports.getProject = getProject;
function deleteProject(projectId) {
    var req = { project_id: projectId };
    Proxy.Call("DeleteProject", req, print);
}
exports.deleteProject = deleteProject;
function createDatasource(projectId, name, description, path) {
    var req = { project_id: projectId, name: name, description: description, path: path };
    Proxy.Call("CreateDatasource", req, print);
}
exports.createDatasource = createDatasource;
function getDatasources(projectId, offset, limit) {
    var req = { project_id: projectId, offset: offset, limit: limit };
    Proxy.Call("GetDatasources", req, print);
}
exports.getDatasources = getDatasources;
function getDatasource(datasourceId) {
    var req = { datasource_id: datasourceId };
    Proxy.Call("GetDatasource", req, print);
}
exports.getDatasource = getDatasource;
function updateDatasource(datasourceId, name, description, path) {
    var req = { datasource_id: datasourceId, name: name, description: description, path: path };
    Proxy.Call("UpdateDatasource", req, print);
}
exports.updateDatasource = updateDatasource;
function deleteDatasource(datasourceId) {
    var req = { datasource_id: datasourceId };
    Proxy.Call("DeleteDatasource", req, print);
}
exports.deleteDatasource = deleteDatasource;
function createDataset(clusterId, datasourceId, name, description, responseColumnName) {
    var req = { cluster_id: clusterId, datasource_id: datasourceId, name: name, description: description, response_column_name: responseColumnName };
    Proxy.Call("CreateDataset", req, print);
}
exports.createDataset = createDataset;
function getDatasets(datasourceId, offset, limit) {
    var req = { datasource_id: datasourceId, offset: offset, limit: limit };
    Proxy.Call("GetDatasets", req, print);
}
exports.getDatasets = getDatasets;
function getDataset(datasetId) {
    var req = { dataset_id: datasetId };
    Proxy.Call("GetDataset", req, print);
}
exports.getDataset = getDataset;
function getDatasetsFromCluster(clusterId) {
    var req = { cluster_id: clusterId };
    Proxy.Call("GetDatasetsFromCluster", req, print);
}
exports.getDatasetsFromCluster = getDatasetsFromCluster;
function updateDataset(datasetId, name, description, responseColumnName) {
    var req = { dataset_id: datasetId, name: name, description: description, response_column_name: responseColumnName };
    Proxy.Call("UpdateDataset", req, print);
}
exports.updateDataset = updateDataset;
function splitDataset(datasetId, ratio1, ratio2) {
    var req = { dataset_id: datasetId, ratio1: ratio1, ratio2: ratio2 };
    Proxy.Call("SplitDataset", req, print);
}
exports.splitDataset = splitDataset;
function deleteDataset(datasetId) {
    var req = { dataset_id: datasetId };
    Proxy.Call("DeleteDataset", req, print);
}
exports.deleteDataset = deleteDataset;
function buildModel(clusterId, datasetId, algorithm) {
    var req = { cluster_id: clusterId, dataset_id: datasetId, algorithm: algorithm };
    Proxy.Call("BuildModel", req, print);
}
exports.buildModel = buildModel;
function buildModelAuto(clusterId, dataset, targetName, maxRunTime) {
    var req = { cluster_id: clusterId, dataset: dataset, target_name: targetName, max_run_time: maxRunTime };
    Proxy.Call("BuildModelAuto", req, print);
}
exports.buildModelAuto = buildModelAuto;
function getModel(modelId) {
    var req = { model_id: modelId };
    Proxy.Call("GetModel", req, print);
}
exports.getModel = getModel;
function getModels(projectId, offset, limit) {
    var req = { project_id: projectId, offset: offset, limit: limit };
    Proxy.Call("GetModels", req, print);
}
exports.getModels = getModels;
function getModelsFromCluster(clusterId, frameKey) {
    var req = { cluster_id: clusterId, frame_key: frameKey };
    Proxy.Call("GetModelsFromCluster", req, print);
}
exports.getModelsFromCluster = getModelsFromCluster;
function findModelsCount(projectId) {
    var req = { project_id: projectId };
    Proxy.Call("FindModelsCount", req, print);
}
exports.findModelsCount = findModelsCount;
function getAllBinomialSortCriteria() {
    var req = {};
    Proxy.Call("GetAllBinomialSortCriteria", req, print);
}
exports.getAllBinomialSortCriteria = getAllBinomialSortCriteria;
function findModelsBinomial(projectId, namePart, sortBy, ascending, offset, limit) {
    var req = { project_id: projectId, name_part: namePart, sort_by: sortBy, ascending: ascending, offset: offset, limit: limit };
    Proxy.Call("FindModelsBinomial", req, print);
}
exports.findModelsBinomial = findModelsBinomial;
function getModelBinomial(modelId) {
    var req = { model_id: modelId };
    Proxy.Call("GetModelBinomial", req, print);
}
exports.getModelBinomial = getModelBinomial;
function getAllMultinomialSortCriteria() {
    var req = {};
    Proxy.Call("GetAllMultinomialSortCriteria", req, print);
}
exports.getAllMultinomialSortCriteria = getAllMultinomialSortCriteria;
function findModelsMultinomial(projectId, namePart, sortBy, ascending, offset, limit) {
    var req = { project_id: projectId, name_part: namePart, sort_by: sortBy, ascending: ascending, offset: offset, limit: limit };
    Proxy.Call("FindModelsMultinomial", req, print);
}
exports.findModelsMultinomial = findModelsMultinomial;
function getModelMultinomial(modelId) {
    var req = { model_id: modelId };
    Proxy.Call("GetModelMultinomial", req, print);
}
exports.getModelMultinomial = getModelMultinomial;
function getAllRegressionSortCriteria() {
    var req = {};
    Proxy.Call("GetAllRegressionSortCriteria", req, print);
}
exports.getAllRegressionSortCriteria = getAllRegressionSortCriteria;
function findModelsRegression(projectId, namePart, sortBy, ascending, offset, limit) {
    var req = { project_id: projectId, name_part: namePart, sort_by: sortBy, ascending: ascending, offset: offset, limit: limit };
    Proxy.Call("FindModelsRegression", req, print);
}
exports.findModelsRegression = findModelsRegression;
function getModelRegression(modelId) {
    var req = { model_id: modelId };
    Proxy.Call("GetModelRegression", req, print);
}
exports.getModelRegression = getModelRegression;
function importModelFromCluster(clusterId, projectId, modelKey, modelName) {
    var req = { cluster_id: clusterId, project_id: projectId, model_key: modelKey, model_name: modelName };
    Proxy.Call("ImportModelFromCluster", req, print);
}
exports.importModelFromCluster = importModelFromCluster;
function deleteModel(modelId) {
    var req = { model_id: modelId };
    Proxy.Call("DeleteModel", req, print);
}
exports.deleteModel = deleteModel;
function createLabel(projectId, name, description) {
    var req = { project_id: projectId, name: name, description: description };
    Proxy.Call("CreateLabel", req, print);
}
exports.createLabel = createLabel;
function updateLabel(labelId, name, description) {
    var req = { label_id: labelId, name: name, description: description };
    Proxy.Call("UpdateLabel", req, print);
}
exports.updateLabel = updateLabel;
function deleteLabel(labelId) {
    var req = { label_id: labelId };
    Proxy.Call("DeleteLabel", req, print);
}
exports.deleteLabel = deleteLabel;
function linkLabelWithModel(labelId, modelId) {
    var req = { label_id: labelId, model_id: modelId };
    Proxy.Call("LinkLabelWithModel", req, print);
}
exports.linkLabelWithModel = linkLabelWithModel;
function unlinkLabelFromModel(labelId, modelId) {
    var req = { label_id: labelId, model_id: modelId };
    Proxy.Call("UnlinkLabelFromModel", req, print);
}
exports.unlinkLabelFromModel = unlinkLabelFromModel;
function getLabelsForProject(projectId) {
    var req = { project_id: projectId };
    Proxy.Call("GetLabelsForProject", req, print);
}
exports.getLabelsForProject = getLabelsForProject;
function startService(modelId, name, packageName) {
    var req = { model_id: modelId, name: name, package_name: packageName };
    Proxy.Call("StartService", req, print);
}
exports.startService = startService;
function stopService(serviceId) {
    var req = { service_id: serviceId };
    Proxy.Call("StopService", req, print);
}
exports.stopService = stopService;
function getService(serviceId) {
    var req = { service_id: serviceId };
    Proxy.Call("GetService", req, print);
}
exports.getService = getService;
function getServices(offset, limit) {
    var req = { offset: offset, limit: limit };
    Proxy.Call("GetServices", req, print);
}
exports.getServices = getServices;
function getServicesForProject(projectId, offset, limit) {
    var req = { project_id: projectId, offset: offset, limit: limit };
    Proxy.Call("GetServicesForProject", req, print);
}
exports.getServicesForProject = getServicesForProject;
function getServicesForModel(modelId, offset, limit) {
    var req = { model_id: modelId, offset: offset, limit: limit };
    Proxy.Call("GetServicesForModel", req, print);
}
exports.getServicesForModel = getServicesForModel;
function deleteService(serviceId) {
    var req = { service_id: serviceId };
    Proxy.Call("DeleteService", req, print);
}
exports.deleteService = deleteService;
function addEngine(engineName, enginePath) {
    var req = { engine_name: engineName, engine_path: enginePath };
    Proxy.Call("AddEngine", req, print);
}
exports.addEngine = addEngine;
function getEngine(engineId) {
    var req = { engine_id: engineId };
    Proxy.Call("GetEngine", req, print);
}
exports.getEngine = getEngine;
function getEngines() {
    var req = {};
    Proxy.Call("GetEngines", req, print);
}
exports.getEngines = getEngines;
function deleteEngine(engineId) {
    var req = { engine_id: engineId };
    Proxy.Call("DeleteEngine", req, print);
}
exports.deleteEngine = deleteEngine;
function getAllEntityTypes() {
    var req = {};
    Proxy.Call("GetAllEntityTypes", req, print);
}
exports.getAllEntityTypes = getAllEntityTypes;
function getAllPermissions() {
    var req = {};
    Proxy.Call("GetAllPermissions", req, print);
}
exports.getAllPermissions = getAllPermissions;
function getAllClusterTypes() {
    var req = {};
    Proxy.Call("GetAllClusterTypes", req, print);
}
exports.getAllClusterTypes = getAllClusterTypes;
function getPermissionsForRole(roleId) {
    var req = { role_id: roleId };
    Proxy.Call("GetPermissionsForRole", req, print);
}
exports.getPermissionsForRole = getPermissionsForRole;
function getPermissionsForIdentity(identityId) {
    var req = { identity_id: identityId };
    Proxy.Call("GetPermissionsForIdentity", req, print);
}
exports.getPermissionsForIdentity = getPermissionsForIdentity;
function createRole(name, description) {
    var req = { name: name, description: description };
    Proxy.Call("CreateRole", req, print);
}
exports.createRole = createRole;
function getRoles(offset, limit) {
    var req = { offset: offset, limit: limit };
    Proxy.Call("GetRoles", req, print);
}
exports.getRoles = getRoles;
function getRolesForIdentity(identityId) {
    var req = { identity_id: identityId };
    Proxy.Call("GetRolesForIdentity", req, print);
}
exports.getRolesForIdentity = getRolesForIdentity;
function getRole(roleId) {
    var req = { role_id: roleId };
    Proxy.Call("GetRole", req, print);
}
exports.getRole = getRole;
function getRoleByName(name) {
    var req = { name: name };
    Proxy.Call("GetRoleByName", req, print);
}
exports.getRoleByName = getRoleByName;
function updateRole(roleId, name, description) {
    var req = { role_id: roleId, name: name, description: description };
    Proxy.Call("UpdateRole", req, print);
}
exports.updateRole = updateRole;
function linkRoleWithPermissions(roleId, permissionIds) {
    var req = { role_id: roleId, permission_ids: permissionIds };
    Proxy.Call("LinkRoleWithPermissions", req, print);
}
exports.linkRoleWithPermissions = linkRoleWithPermissions;
function linkRoleWithPermission(roleId, permissionId) {
    var req = { role_id: roleId, permission_id: permissionId };
    Proxy.Call("LinkRoleWithPermission", req, print);
}
exports.linkRoleWithPermission = linkRoleWithPermission;
function unlinkRoleFromPermission(roleId, permissionId) {
    var req = { role_id: roleId, permission_id: permissionId };
    Proxy.Call("UnlinkRoleFromPermission", req, print);
}
exports.unlinkRoleFromPermission = unlinkRoleFromPermission;
function deleteRole(roleId) {
    var req = { role_id: roleId };
    Proxy.Call("DeleteRole", req, print);
}
exports.deleteRole = deleteRole;
function createWorkgroup(name, description) {
    var req = { name: name, description: description };
    Proxy.Call("CreateWorkgroup", req, print);
}
exports.createWorkgroup = createWorkgroup;
function getWorkgroups(offset, limit) {
    var req = { offset: offset, limit: limit };
    Proxy.Call("GetWorkgroups", req, print);
}
exports.getWorkgroups = getWorkgroups;
function getWorkgroupsForIdentity(identityId) {
    var req = { identity_id: identityId };
    Proxy.Call("GetWorkgroupsForIdentity", req, print);
}
exports.getWorkgroupsForIdentity = getWorkgroupsForIdentity;
function getWorkgroup(workgroupId) {
    var req = { workgroup_id: workgroupId };
    Proxy.Call("GetWorkgroup", req, print);
}
exports.getWorkgroup = getWorkgroup;
function getWorkgroupByName(name) {
    var req = { name: name };
    Proxy.Call("GetWorkgroupByName", req, print);
}
exports.getWorkgroupByName = getWorkgroupByName;
function updateWorkgroup(workgroupId, name, description) {
    var req = { workgroup_id: workgroupId, name: name, description: description };
    Proxy.Call("UpdateWorkgroup", req, print);
}
exports.updateWorkgroup = updateWorkgroup;
function deleteWorkgroup(workgroupId) {
    var req = { workgroup_id: workgroupId };
    Proxy.Call("DeleteWorkgroup", req, print);
}
exports.deleteWorkgroup = deleteWorkgroup;
function createIdentity(name, password) {
    var req = { name: name, password: password };
    Proxy.Call("CreateIdentity", req, print);
}
exports.createIdentity = createIdentity;
function getIdentities(offset, limit) {
    var req = { offset: offset, limit: limit };
    Proxy.Call("GetIdentities", req, print);
}
exports.getIdentities = getIdentities;
function getIdentitiesForWorkgroup(workgroupId) {
    var req = { workgroup_id: workgroupId };
    Proxy.Call("GetIdentitiesForWorkgroup", req, print);
}
exports.getIdentitiesForWorkgroup = getIdentitiesForWorkgroup;
function getIdentitiesForRole(roleId) {
    var req = { role_id: roleId };
    Proxy.Call("GetIdentitiesForRole", req, print);
}
exports.getIdentitiesForRole = getIdentitiesForRole;
function getIdentitiesForEntity(entityType, entityId) {
    var req = { entity_type: entityType, entity_id: entityId };
    Proxy.Call("GetIdentitiesForEntity", req, print);
}
exports.getIdentitiesForEntity = getIdentitiesForEntity;
function getIdentity(identityId) {
    var req = { identity_id: identityId };
    Proxy.Call("GetIdentity", req, print);
}
exports.getIdentity = getIdentity;
function getIdentityByName(name) {
    var req = { name: name };
    Proxy.Call("GetIdentityByName", req, print);
}
exports.getIdentityByName = getIdentityByName;
function linkIdentityWithWorkgroup(identityId, workgroupId) {
    var req = { identity_id: identityId, workgroup_id: workgroupId };
    Proxy.Call("LinkIdentityWithWorkgroup", req, print);
}
exports.linkIdentityWithWorkgroup = linkIdentityWithWorkgroup;
function unlinkIdentityFromWorkgroup(identityId, workgroupId) {
    var req = { identity_id: identityId, workgroup_id: workgroupId };
    Proxy.Call("UnlinkIdentityFromWorkgroup", req, print);
}
exports.unlinkIdentityFromWorkgroup = unlinkIdentityFromWorkgroup;
function linkIdentityWithRole(identityId, roleId) {
    var req = { identity_id: identityId, role_id: roleId };
    Proxy.Call("LinkIdentityWithRole", req, print);
}
exports.linkIdentityWithRole = linkIdentityWithRole;
function unlinkIdentityFromRole(identityId, roleId) {
    var req = { identity_id: identityId, role_id: roleId };
    Proxy.Call("UnlinkIdentityFromRole", req, print);
}
exports.unlinkIdentityFromRole = unlinkIdentityFromRole;
function updateIdentity(identityId, password) {
    var req = { identity_id: identityId, password: password };
    Proxy.Call("UpdateIdentity", req, print);
}
exports.updateIdentity = updateIdentity;
function deactivateIdentity(identityId) {
    var req = { identity_id: identityId };
    Proxy.Call("DeactivateIdentity", req, print);
}
exports.deactivateIdentity = deactivateIdentity;
function shareEntity(kind, workgroupId, entityTypeId, entityId) {
    var req = { kind: kind, workgroup_id: workgroupId, entity_type_id: entityTypeId, entity_id: entityId };
    Proxy.Call("ShareEntity", req, print);
}
exports.shareEntity = shareEntity;
function getPrivileges(entityTypeId, entityId) {
    var req = { entity_type_id: entityTypeId, entity_id: entityId };
    Proxy.Call("GetPrivileges", req, print);
}
exports.getPrivileges = getPrivileges;
function unshareEntity(kind, workgroupId, entityTypeId, entityId) {
    var req = { kind: kind, workgroup_id: workgroupId, entity_type_id: entityTypeId, entity_id: entityId };
    Proxy.Call("UnshareEntity", req, print);
}
exports.unshareEntity = unshareEntity;
function getHistory(entityTypeId, entityId, offset, limit) {
    var req = { entity_type_id: entityTypeId, entity_id: entityId, offset: offset, limit: limit };
    Proxy.Call("GetHistory", req, print);
}
exports.getHistory = getHistory;
function createPackage(projectId, name) {
    var req = { project_id: projectId, name: name };
    Proxy.Call("CreatePackage", req, print);
}
exports.createPackage = createPackage;
function getPackages(projectId) {
    var req = { project_id: projectId };
    Proxy.Call("GetPackages", req, print);
}
exports.getPackages = getPackages;
function getPackageDirectories(projectId, packageName, relativePath) {
    var req = { project_id: projectId, package_name: packageName, relative_path: relativePath };
    Proxy.Call("GetPackageDirectories", req, print);
}
exports.getPackageDirectories = getPackageDirectories;
function getPackageFiles(projectId, packageName, relativePath) {
    var req = { project_id: projectId, package_name: packageName, relative_path: relativePath };
    Proxy.Call("GetPackageFiles", req, print);
}
exports.getPackageFiles = getPackageFiles;
function deletePackage(projectId, name) {
    var req = { project_id: projectId, name: name };
    Proxy.Call("DeletePackage", req, print);
}
exports.deletePackage = deletePackage;
function deletePackageDirectory(projectId, packageName, relativePath) {
    var req = { project_id: projectId, package_name: packageName, relative_path: relativePath };
    Proxy.Call("DeletePackageDirectory", req, print);
}
exports.deletePackageDirectory = deletePackageDirectory;
function deletePackageFile(projectId, packageName, relativePath) {
    var req = { project_id: projectId, package_name: packageName, relative_path: relativePath };
    Proxy.Call("DeletePackageFile", req, print);
}
exports.deletePackageFile = deletePackageFile;
function setAttributesForPackage(projectId, packageName, attributes) {
    var req = { project_id: projectId, package_name: packageName, attributes: attributes };
    Proxy.Call("SetAttributesForPackage", req, print);
}
exports.setAttributesForPackage = setAttributesForPackage;
function getAttributesForPackage(projectId, packageName) {
    var req = { project_id: projectId, package_name: packageName };
    Proxy.Call("GetAttributesForPackage", req, print);
}
exports.getAttributesForPackage = getAttributesForPackage;
//# sourceMappingURL=CLI.js.map