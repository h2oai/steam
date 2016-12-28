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

// ------------------------------
// --- This is generated code ---
// ---      DO NOT EDIT       ---
// ------------------------------




// --- CLI Stub ---

import * as Proxy from './xhr';

function print(error: Error, data: any): void {
  if (error) {
    console.error(error);
  } else {
    console.log(data);
  }
}


export function pingServer(input: string): void {
  const req: any = { input: input };
  Proxy.Call("PingServer", req, print);
}

export function getConfig(): void {
  const req: any = {  };
  Proxy.Call("GetConfig", req, print);
}

export function registerCluster(address: string): void {
  const req: any = { address: address };
  Proxy.Call("RegisterCluster", req, print);
}

export function unregisterCluster(clusterId: number): void {
  const req: any = { cluster_id: clusterId };
  Proxy.Call("UnregisterCluster", req, print);
}

export function startClusterOnYarn(clusterName: string, engineId: number, size: number, memory: string, secure: boolean, keytab: string): void {
  const req: any = { cluster_name: clusterName, engine_id: engineId, size: size, memory: memory, secure: secure, keytab: keytab };
  Proxy.Call("StartClusterOnYarn", req, print);
}

export function stopClusterOnYarn(clusterId: number, keytab: string): void {
  const req: any = { cluster_id: clusterId, keytab: keytab };
  Proxy.Call("StopClusterOnYarn", req, print);
}

export function getCluster(clusterId: number): void {
  const req: any = { cluster_id: clusterId };
  Proxy.Call("GetCluster", req, print);
}

export function getClusterOnYarn(clusterId: number): void {
  const req: any = { cluster_id: clusterId };
  Proxy.Call("GetClusterOnYarn", req, print);
}

export function getClusters(offset: number, limit: number): void {
  const req: any = { offset: offset, limit: limit };
  Proxy.Call("GetClusters", req, print);
}

export function getClusterStatus(clusterId: number): void {
  const req: any = { cluster_id: clusterId };
  Proxy.Call("GetClusterStatus", req, print);
}

export function deleteCluster(clusterId: number): void {
  const req: any = { cluster_id: clusterId };
  Proxy.Call("DeleteCluster", req, print);
}

export function getJob(clusterId: number, jobName: string): void {
  const req: any = { cluster_id: clusterId, job_name: jobName };
  Proxy.Call("GetJob", req, print);
}

export function getJobs(clusterId: number): void {
  const req: any = { cluster_id: clusterId };
  Proxy.Call("GetJobs", req, print);
}

export function createProject(name: string, description: string, modelCategory: string): void {
  const req: any = { name: name, description: description, model_category: modelCategory };
  Proxy.Call("CreateProject", req, print);
}

export function getProjects(offset: number, limit: number): void {
  const req: any = { offset: offset, limit: limit };
  Proxy.Call("GetProjects", req, print);
}

export function getProject(projectId: number): void {
  const req: any = { project_id: projectId };
  Proxy.Call("GetProject", req, print);
}

export function deleteProject(projectId: number): void {
  const req: any = { project_id: projectId };
  Proxy.Call("DeleteProject", req, print);
}

export function createDatasource(projectId: number, name: string, description: string, path: string): void {
  const req: any = { project_id: projectId, name: name, description: description, path: path };
  Proxy.Call("CreateDatasource", req, print);
}

export function getDatasources(projectId: number, offset: number, limit: number): void {
  const req: any = { project_id: projectId, offset: offset, limit: limit };
  Proxy.Call("GetDatasources", req, print);
}

export function getDatasource(datasourceId: number): void {
  const req: any = { datasource_id: datasourceId };
  Proxy.Call("GetDatasource", req, print);
}

export function updateDatasource(datasourceId: number, name: string, description: string, path: string): void {
  const req: any = { datasource_id: datasourceId, name: name, description: description, path: path };
  Proxy.Call("UpdateDatasource", req, print);
}

export function deleteDatasource(datasourceId: number): void {
  const req: any = { datasource_id: datasourceId };
  Proxy.Call("DeleteDatasource", req, print);
}

export function createDataset(clusterId: number, datasourceId: number, name: string, description: string, responseColumnName: string): void {
  const req: any = { cluster_id: clusterId, datasource_id: datasourceId, name: name, description: description, response_column_name: responseColumnName };
  Proxy.Call("CreateDataset", req, print);
}

export function getDatasets(datasourceId: number, offset: number, limit: number): void {
  const req: any = { datasource_id: datasourceId, offset: offset, limit: limit };
  Proxy.Call("GetDatasets", req, print);
}

export function getDataset(datasetId: number): void {
  const req: any = { dataset_id: datasetId };
  Proxy.Call("GetDataset", req, print);
}

export function getDatasetsFromCluster(clusterId: number): void {
  const req: any = { cluster_id: clusterId };
  Proxy.Call("GetDatasetsFromCluster", req, print);
}

export function updateDataset(datasetId: number, name: string, description: string, responseColumnName: string): void {
  const req: any = { dataset_id: datasetId, name: name, description: description, response_column_name: responseColumnName };
  Proxy.Call("UpdateDataset", req, print);
}

export function splitDataset(datasetId: number, ratio1: number, ratio2: number): void {
  const req: any = { dataset_id: datasetId, ratio1: ratio1, ratio2: ratio2 };
  Proxy.Call("SplitDataset", req, print);
}

export function deleteDataset(datasetId: number): void {
  const req: any = { dataset_id: datasetId };
  Proxy.Call("DeleteDataset", req, print);
}

export function buildModel(clusterId: number, datasetId: number, algorithm: string): void {
  const req: any = { cluster_id: clusterId, dataset_id: datasetId, algorithm: algorithm };
  Proxy.Call("BuildModel", req, print);
}

export function buildModelAuto(clusterId: number, dataset: string, targetName: string, maxRunTime: number): void {
  const req: any = { cluster_id: clusterId, dataset: dataset, target_name: targetName, max_run_time: maxRunTime };
  Proxy.Call("BuildModelAuto", req, print);
}

export function getModel(modelId: number): void {
  const req: any = { model_id: modelId };
  Proxy.Call("GetModel", req, print);
}

export function getModels(projectId: number, offset: number, limit: number): void {
  const req: any = { project_id: projectId, offset: offset, limit: limit };
  Proxy.Call("GetModels", req, print);
}

export function getModelsFromCluster(clusterId: number, frameKey: string): void {
  const req: any = { cluster_id: clusterId, frame_key: frameKey };
  Proxy.Call("GetModelsFromCluster", req, print);
}

export function findModelsCount(projectId: number): void {
  const req: any = { project_id: projectId };
  Proxy.Call("FindModelsCount", req, print);
}

export function getAllBinomialSortCriteria(): void {
  const req: any = {  };
  Proxy.Call("GetAllBinomialSortCriteria", req, print);
}

export function findModelsBinomial(projectId: number, namePart: string, sortBy: string, ascending: boolean, offset: number, limit: number): void {
  const req: any = { project_id: projectId, name_part: namePart, sort_by: sortBy, ascending: ascending, offset: offset, limit: limit };
  Proxy.Call("FindModelsBinomial", req, print);
}

export function getModelBinomial(modelId: number): void {
  const req: any = { model_id: modelId };
  Proxy.Call("GetModelBinomial", req, print);
}

export function getAllMultinomialSortCriteria(): void {
  const req: any = {  };
  Proxy.Call("GetAllMultinomialSortCriteria", req, print);
}

export function findModelsMultinomial(projectId: number, namePart: string, sortBy: string, ascending: boolean, offset: number, limit: number): void {
  const req: any = { project_id: projectId, name_part: namePart, sort_by: sortBy, ascending: ascending, offset: offset, limit: limit };
  Proxy.Call("FindModelsMultinomial", req, print);
}

export function getModelMultinomial(modelId: number): void {
  const req: any = { model_id: modelId };
  Proxy.Call("GetModelMultinomial", req, print);
}

export function getAllRegressionSortCriteria(): void {
  const req: any = {  };
  Proxy.Call("GetAllRegressionSortCriteria", req, print);
}

export function findModelsRegression(projectId: number, namePart: string, sortBy: string, ascending: boolean, offset: number, limit: number): void {
  const req: any = { project_id: projectId, name_part: namePart, sort_by: sortBy, ascending: ascending, offset: offset, limit: limit };
  Proxy.Call("FindModelsRegression", req, print);
}

export function getModelRegression(modelId: number): void {
  const req: any = { model_id: modelId };
  Proxy.Call("GetModelRegression", req, print);
}

export function importModelFromCluster(clusterId: number, projectId: number, modelKey: string, modelName: string): void {
  const req: any = { cluster_id: clusterId, project_id: projectId, model_key: modelKey, model_name: modelName };
  Proxy.Call("ImportModelFromCluster", req, print);
}

export function checkMojo(algo: string): void {
  const req: any = { algo: algo };
  Proxy.Call("CheckMojo", req, print);
}

export function importModelPojo(modelId: number): void {
  const req: any = { model_id: modelId };
  Proxy.Call("ImportModelPojo", req, print);
}

export function importModelMojo(modelId: number): void {
  const req: any = { model_id: modelId };
  Proxy.Call("ImportModelMojo", req, print);
}

export function deleteModel(modelId: number): void {
  const req: any = { model_id: modelId };
  Proxy.Call("DeleteModel", req, print);
}

export function createLabel(projectId: number, name: string, description: string): void {
  const req: any = { project_id: projectId, name: name, description: description };
  Proxy.Call("CreateLabel", req, print);
}

export function updateLabel(labelId: number, name: string, description: string): void {
  const req: any = { label_id: labelId, name: name, description: description };
  Proxy.Call("UpdateLabel", req, print);
}

export function deleteLabel(labelId: number): void {
  const req: any = { label_id: labelId };
  Proxy.Call("DeleteLabel", req, print);
}

export function linkLabelWithModel(labelId: number, modelId: number): void {
  const req: any = { label_id: labelId, model_id: modelId };
  Proxy.Call("LinkLabelWithModel", req, print);
}

export function unlinkLabelFromModel(labelId: number, modelId: number): void {
  const req: any = { label_id: labelId, model_id: modelId };
  Proxy.Call("UnlinkLabelFromModel", req, print);
}

export function getLabelsForProject(projectId: number): void {
  const req: any = { project_id: projectId };
  Proxy.Call("GetLabelsForProject", req, print);
}

export function startService(modelId: number, name: string, packageName: string): void {
  const req: any = { model_id: modelId, name: name, package_name: packageName };
  Proxy.Call("StartService", req, print);
}

export function stopService(serviceId: number): void {
  const req: any = { service_id: serviceId };
  Proxy.Call("StopService", req, print);
}

export function getService(serviceId: number): void {
  const req: any = { service_id: serviceId };
  Proxy.Call("GetService", req, print);
}

export function getServices(offset: number, limit: number): void {
  const req: any = { offset: offset, limit: limit };
  Proxy.Call("GetServices", req, print);
}

export function getServicesForProject(projectId: number, offset: number, limit: number): void {
  const req: any = { project_id: projectId, offset: offset, limit: limit };
  Proxy.Call("GetServicesForProject", req, print);
}

export function getServicesForModel(modelId: number, offset: number, limit: number): void {
  const req: any = { model_id: modelId, offset: offset, limit: limit };
  Proxy.Call("GetServicesForModel", req, print);
}

export function deleteService(serviceId: number): void {
  const req: any = { service_id: serviceId };
  Proxy.Call("DeleteService", req, print);
}

export function getEngine(engineId: number): void {
  const req: any = { engine_id: engineId };
  Proxy.Call("GetEngine", req, print);
}

export function getEngines(): void {
  const req: any = {  };
  Proxy.Call("GetEngines", req, print);
}

export function deleteEngine(engineId: number): void {
  const req: any = { engine_id: engineId };
  Proxy.Call("DeleteEngine", req, print);
}

export function getAllEntityTypes(): void {
  const req: any = {  };
  Proxy.Call("GetAllEntityTypes", req, print);
}

export function getAllPermissions(): void {
  const req: any = {  };
  Proxy.Call("GetAllPermissions", req, print);
}

export function getAllClusterTypes(): void {
  const req: any = {  };
  Proxy.Call("GetAllClusterTypes", req, print);
}

export function getPermissionsForRole(roleId: number): void {
  const req: any = { role_id: roleId };
  Proxy.Call("GetPermissionsForRole", req, print);
}

export function getPermissionsForIdentity(identityId: number): void {
  const req: any = { identity_id: identityId };
  Proxy.Call("GetPermissionsForIdentity", req, print);
}

export function createRole(name: string, description: string): void {
  const req: any = { name: name, description: description };
  Proxy.Call("CreateRole", req, print);
}

export function getRoles(offset: number, limit: number): void {
  const req: any = { offset: offset, limit: limit };
  Proxy.Call("GetRoles", req, print);
}

export function getRolesForIdentity(identityId: number): void {
  const req: any = { identity_id: identityId };
  Proxy.Call("GetRolesForIdentity", req, print);
}

export function getRole(roleId: number): void {
  const req: any = { role_id: roleId };
  Proxy.Call("GetRole", req, print);
}

export function getRoleByName(name: string): void {
  const req: any = { name: name };
  Proxy.Call("GetRoleByName", req, print);
}

export function updateRole(roleId: number, name: string, description: string): void {
  const req: any = { role_id: roleId, name: name, description: description };
  Proxy.Call("UpdateRole", req, print);
}

export function linkRoleWithPermissions(roleId: number, permissionIds: number[]): void {
  const req: any = { role_id: roleId, permission_ids: permissionIds };
  Proxy.Call("LinkRoleWithPermissions", req, print);
}

export function linkRoleWithPermission(roleId: number, permissionId: number): void {
  const req: any = { role_id: roleId, permission_id: permissionId };
  Proxy.Call("LinkRoleWithPermission", req, print);
}

export function unlinkRoleFromPermission(roleId: number, permissionId: number): void {
  const req: any = { role_id: roleId, permission_id: permissionId };
  Proxy.Call("UnlinkRoleFromPermission", req, print);
}

export function deleteRole(roleId: number): void {
  const req: any = { role_id: roleId };
  Proxy.Call("DeleteRole", req, print);
}

export function createWorkgroup(name: string, description: string): void {
  const req: any = { name: name, description: description };
  Proxy.Call("CreateWorkgroup", req, print);
}

export function getWorkgroups(offset: number, limit: number): void {
  const req: any = { offset: offset, limit: limit };
  Proxy.Call("GetWorkgroups", req, print);
}

export function getWorkgroupsForIdentity(identityId: number): void {
  const req: any = { identity_id: identityId };
  Proxy.Call("GetWorkgroupsForIdentity", req, print);
}

export function getWorkgroup(workgroupId: number): void {
  const req: any = { workgroup_id: workgroupId };
  Proxy.Call("GetWorkgroup", req, print);
}

export function getWorkgroupByName(name: string): void {
  const req: any = { name: name };
  Proxy.Call("GetWorkgroupByName", req, print);
}

export function updateWorkgroup(workgroupId: number, name: string, description: string): void {
  const req: any = { workgroup_id: workgroupId, name: name, description: description };
  Proxy.Call("UpdateWorkgroup", req, print);
}

export function deleteWorkgroup(workgroupId: number): void {
  const req: any = { workgroup_id: workgroupId };
  Proxy.Call("DeleteWorkgroup", req, print);
}

export function createIdentity(name: string, password: string): void {
  const req: any = { name: name, password: password };
  Proxy.Call("CreateIdentity", req, print);
}

export function getIdentities(offset: number, limit: number): void {
  const req: any = { offset: offset, limit: limit };
  Proxy.Call("GetIdentities", req, print);
}

export function getIdentitiesForWorkgroup(workgroupId: number): void {
  const req: any = { workgroup_id: workgroupId };
  Proxy.Call("GetIdentitiesForWorkgroup", req, print);
}

export function getIdentitiesForRole(roleId: number): void {
  const req: any = { role_id: roleId };
  Proxy.Call("GetIdentitiesForRole", req, print);
}

export function getIdentitiesForEntity(entityType: number, entityId: number): void {
  const req: any = { entity_type: entityType, entity_id: entityId };
  Proxy.Call("GetIdentitiesForEntity", req, print);
}

export function getIdentity(identityId: number): void {
  const req: any = { identity_id: identityId };
  Proxy.Call("GetIdentity", req, print);
}

export function getIdentityByName(name: string): void {
  const req: any = { name: name };
  Proxy.Call("GetIdentityByName", req, print);
}

export function linkIdentityWithWorkgroup(identityId: number, workgroupId: number): void {
  const req: any = { identity_id: identityId, workgroup_id: workgroupId };
  Proxy.Call("LinkIdentityWithWorkgroup", req, print);
}

export function unlinkIdentityFromWorkgroup(identityId: number, workgroupId: number): void {
  const req: any = { identity_id: identityId, workgroup_id: workgroupId };
  Proxy.Call("UnlinkIdentityFromWorkgroup", req, print);
}

export function linkIdentityWithRole(identityId: number, roleId: number): void {
  const req: any = { identity_id: identityId, role_id: roleId };
  Proxy.Call("LinkIdentityWithRole", req, print);
}

export function unlinkIdentityFromRole(identityId: number, roleId: number): void {
  const req: any = { identity_id: identityId, role_id: roleId };
  Proxy.Call("UnlinkIdentityFromRole", req, print);
}

export function updateIdentity(identityId: number, password: string): void {
  const req: any = { identity_id: identityId, password: password };
  Proxy.Call("UpdateIdentity", req, print);
}

export function activateIdentity(identityId: number): void {
  const req: any = { identity_id: identityId };
  Proxy.Call("ActivateIdentity", req, print);
}

export function deactivateIdentity(identityId: number): void {
  const req: any = { identity_id: identityId };
  Proxy.Call("DeactivateIdentity", req, print);
}

export function shareEntity(kind: string, workgroupId: number, entityTypeId: number, entityId: number): void {
  const req: any = { kind: kind, workgroup_id: workgroupId, entity_type_id: entityTypeId, entity_id: entityId };
  Proxy.Call("ShareEntity", req, print);
}

export function getPrivileges(entityTypeId: number, entityId: number): void {
  const req: any = { entity_type_id: entityTypeId, entity_id: entityId };
  Proxy.Call("GetPrivileges", req, print);
}

export function unshareEntity(kind: string, workgroupId: number, entityTypeId: number, entityId: number): void {
  const req: any = { kind: kind, workgroup_id: workgroupId, entity_type_id: entityTypeId, entity_id: entityId };
  Proxy.Call("UnshareEntity", req, print);
}

export function getHistory(entityTypeId: number, entityId: number, offset: number, limit: number): void {
  const req: any = { entity_type_id: entityTypeId, entity_id: entityId, offset: offset, limit: limit };
  Proxy.Call("GetHistory", req, print);
}

export function createPackage(projectId: number, name: string): void {
  const req: any = { project_id: projectId, name: name };
  Proxy.Call("CreatePackage", req, print);
}

export function getPackages(projectId: number): void {
  const req: any = { project_id: projectId };
  Proxy.Call("GetPackages", req, print);
}

export function getPackageDirectories(projectId: number, packageName: string, relativePath: string): void {
  const req: any = { project_id: projectId, package_name: packageName, relative_path: relativePath };
  Proxy.Call("GetPackageDirectories", req, print);
}

export function getPackageFiles(projectId: number, packageName: string, relativePath: string): void {
  const req: any = { project_id: projectId, package_name: packageName, relative_path: relativePath };
  Proxy.Call("GetPackageFiles", req, print);
}

export function deletePackage(projectId: number, name: string): void {
  const req: any = { project_id: projectId, name: name };
  Proxy.Call("DeletePackage", req, print);
}

export function deletePackageDirectory(projectId: number, packageName: string, relativePath: string): void {
  const req: any = { project_id: projectId, package_name: packageName, relative_path: relativePath };
  Proxy.Call("DeletePackageDirectory", req, print);
}

export function deletePackageFile(projectId: number, packageName: string, relativePath: string): void {
  const req: any = { project_id: projectId, package_name: packageName, relative_path: relativePath };
  Proxy.Call("DeletePackageFile", req, print);
}

export function setAttributesForPackage(projectId: number, packageName: string, attributes: string): void {
  const req: any = { project_id: projectId, package_name: packageName, attributes: attributes };
  Proxy.Call("SetAttributesForPackage", req, print);
}

export function getAttributesForPackage(projectId: number, packageName: string): void {
  const req: any = { project_id: projectId, package_name: packageName };
  Proxy.Call("GetAttributesForPackage", req, print);
}




