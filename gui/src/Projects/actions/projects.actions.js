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
/**
 * Created by justin on 7/18/16.
 */
var Remote = require('../../Proxy/Proxy');
var _ = require('lodash');
var global_actions_1 = require('../../App/actions/global.actions');
var notification_actions_1 = require('../../App/actions/notification.actions');
var Notification_1 = require('../../App/components/Notification');
exports.SET_CURRENT_PROJECT = 'SET_CURRENT_PROJECT';
exports.REQUEST_CLUSTERS = 'REQUEST_CLUSTERS';
exports.RECEIVE_CLUSTERS = 'RECEIVE_CLUSTERS';
exports.RESET_CLUSTER_SELECTION = 'RESET_CLUSTER_SELECTION';
exports.REQUEST_MODELS = 'REQUEST_MODELS';
exports.RECEIVE_MODELS = 'RECEIVE_MODELS';
exports.CREATE_PROJECT_COMPLETED = 'CREATE_PROJECT_COMPLETED';
exports.IMPORT_MODEL_FROM_CLUSTER_COMPLETED = 'IMPORT_MODEL_FROM_CLUSTER_COMPLETED';
exports.RECEIVE_PROJECTS = 'RECEIVE_PROJECTS';
exports.REQUEST_DATASETS_FROM_CLUSTER = 'REQUEST_DATASETS_FROM_CLUSTER';
exports.RECEIVE_DATASETS_FROM_CLUSTER = 'RECEIVE_DATASETS_FROM_CLUSTER';
exports.RECEIVE_MODELS_FROM_PROJECT = 'RECEIVE_MODELS_FROM_PROJECT';
exports.RECEIVE_PROJECT = 'RECEIVE_PROJECT';
exports.REGISTER_CLUSTER_ERROR = 'REGISTER_CLUSTER_ERROR';
function setCurrentProject(projectId) {
    return {
        type: exports.SET_CURRENT_PROJECT,
        projectId: projectId
    };
}
exports.setCurrentProject = setCurrentProject;
;
exports.requestClusters = function () {
    return {
        type: exports.REQUEST_CLUSTERS
    };
};
function receiveClusters(clusters) {
    return {
        type: exports.RECEIVE_CLUSTERS,
        clusters: clusters
    };
}
exports.receiveClusters = receiveClusters;
;
function resetClusterSelection() {
    return {
        type: exports.RESET_CLUSTER_SELECTION
    };
}
exports.resetClusterSelection = resetClusterSelection;
;
function registerClusterError(message) {
    return {
        type: exports.REGISTER_CLUSTER_ERROR,
        message: message
    };
}
exports.registerClusterError = registerClusterError;
;
function fetchClusters() {
    return function (dispatch, getState) {
        dispatch(exports.requestClusters());
        var state = getState();
        if (_.isEmpty(state.global.entityIds)) {
            dispatch(global_actions_1.fetchEntityIds()).then(function () {
                _fetchClusters(dispatch, getState);
            });
        }
        else {
            _fetchClusters(dispatch, getState);
        }
    };
}
exports.fetchClusters = fetchClusters;
function _fetchClusters(dispatch, getState) {
    Remote.getClusters(0, 1000, function (error, res) {
        if (error) {
            notification_actions_1.openNotification(Notification_1.NotificationType.Error, "Load Error", error.toString(), null);
            return;
        }
        var identityPromises = [];
        var toReturn = [];
        var state = getState();
        var _loop_1 = function(cluster) {
            identityPromises.push(new Promise(function (resolve, reject) {
                Remote.getIdentitiesForEntity(state.global.entityIds.cluster, cluster.id, function (identitiesError, users) {
                    if (identitiesError) {
                        notification_actions_1.openNotification(Notification_1.NotificationType.Error, "Load Error", identitiesError.toString(), null);
                        reject(identitiesError.toString());
                        return;
                    }
                    toReturn.push(_.assign({}, cluster, { identities: users }));
                    resolve();
                });
            }));
        };
        for (var _i = 0, res_1 = res; _i < res_1.length; _i++) {
            var cluster = res_1[_i];
            _loop_1(cluster);
        }
        Promise.all(identityPromises).then(function (results) {
            dispatch(receiveClusters(toReturn));
        });
    });
}
exports.requestModels = function () {
    return {
        type: exports.REQUEST_MODELS
    };
};
function receiveModelsFromCluster(models) {
    return {
        type: exports.RECEIVE_MODELS,
        models: models
    };
}
exports.receiveModelsFromCluster = receiveModelsFromCluster;
function createProjectCompleted(project) {
    return {
        type: exports.CREATE_PROJECT_COMPLETED,
        project: project
    };
}
exports.createProjectCompleted = createProjectCompleted;
function receiveProject(project) {
    return {
        type: exports.RECEIVE_PROJECT,
        project: project
    };
}
exports.receiveProject = receiveProject;
function importModelFromClusterCompleted(model) {
    return {
        type: exports.IMPORT_MODEL_FROM_CLUSTER_COMPLETED,
        model: model
    };
}
exports.importModelFromClusterCompleted = importModelFromClusterCompleted;
function receiveProjects(projects) {
    return {
        type: exports.RECEIVE_PROJECTS,
        projects: projects
    };
}
exports.receiveProjects = receiveProjects;
function requestDatasetsFromCluster() {
    return {
        type: exports.REQUEST_DATASETS_FROM_CLUSTER
    };
}
exports.requestDatasetsFromCluster = requestDatasetsFromCluster;
function receiveDatasetsFromCluster(datasets) {
    return {
        type: exports.RECEIVE_DATASETS_FROM_CLUSTER,
        datasets: datasets
    };
}
exports.receiveDatasetsFromCluster = receiveDatasetsFromCluster;
function receiveModelsFromProject(models) {
    return {
        type: exports.RECEIVE_MODELS_FROM_PROJECT,
        models: models
    };
}
exports.receiveModelsFromProject = receiveModelsFromProject;
function fetchModelsFromProject(projectId) {
    return function (dispatch) {
        Remote.getModels(projectId, 0, 5, function (error, res) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                return;
            }
            dispatch(receiveModelsFromProject(res));
        });
    };
}
exports.fetchModelsFromProject = fetchModelsFromProject;
function fetchProject(projectId) {
    return function (dispatch) {
        return new Promise(function (resolve, reject) {
            Remote.getProject(projectId, function (error, res) {
                if (error) {
                    dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                    reject(error);
                    return;
                }
                dispatch(receiveProject(res));
                resolve(res);
            });
        });
    };
}
exports.fetchProject = fetchProject;
function fetchModelsFromCluster(clusterId, frameKey) {
    return function (dispatch) {
        dispatch(exports.requestModels());
        Remote.getModelsFromCluster(clusterId, frameKey, function (error, res) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
            }
            dispatch(receiveModelsFromCluster(res));
        });
    };
}
exports.fetchModelsFromCluster = fetchModelsFromCluster;
function fetchDatasetsFromCluster(clusterId) {
    return function (dispatch) {
        dispatch(requestDatasetsFromCluster());
        Remote.getDatasetsFromCluster(clusterId, function (error, res) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                return;
            }
            dispatch(receiveDatasetsFromCluster(res));
        });
    };
}
exports.fetchDatasetsFromCluster = fetchDatasetsFromCluster;
function createProject(name, modelCategory) {
    return function (dispatch) {
        return new Promise(function (resolve, reject) {
            Remote.createProject(name, '', modelCategory, function (error, res) {
                if (error) {
                    dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                    reject(error);
                    return;
                }
                dispatch(createProjectCompleted(res));
                resolve(res);
            });
        });
    };
}
exports.createProject = createProject;
function importModelFromCluster(clusterId, projectId, modelName) {
    return function (dispatch) {
        return new Promise(function (resolve, reject) {
            Remote.importModelFromCluster(clusterId, projectId, modelName, modelName, function (error, res) {
                if (error) {
                    dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                    reject(error);
                    return;
                }
                dispatch(importModelFromClusterCompleted(res));
                resolve(res);
            });
        });
    };
}
exports.importModelFromCluster = importModelFromCluster;
function importModelsFromCluster(clusterId, projectId, models) {
    return function (dispatch) {
        var promises = [];
        return new Promise(function (resolve, reject) {
            models.map(function (modelName) {
                promises.push(dispatch(importModelFromCluster(clusterId, projectId, modelName)));
            });
            Promise.all(promises).then(function () {
                resolve(projectId);
            });
        });
    };
}
exports.importModelsFromCluster = importModelsFromCluster;
function createProjectAndImportModelsFromCluster(projectName, clusterId, modelCategory, models) {
    return function (dispatch) {
        return new Promise(function (resolve, reject) {
            dispatch(createProject(projectName, modelCategory)).then(function (projectId) {
                dispatch(importModelsFromCluster(clusterId, projectId, models)).then(function () {
                    resolve(projectId);
                });
            });
        });
    };
}
exports.createProjectAndImportModelsFromCluster = createProjectAndImportModelsFromCluster;
function registerCluster(address) {
    return function (dispatch) {
        Remote.registerCluster(address, function (error, res) {
            if (error) {
                dispatch(registerClusterError(error.message));
                return;
            }
            dispatch(fetchClusters());
        });
    };
}
exports.registerCluster = registerCluster;
function unregisterCluster(clusterId) {
    return function (dispatch) {
        Remote.unregisterCluster(clusterId, function (error) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                return;
            }
            dispatch(fetchClusters());
        });
    };
}
exports.unregisterCluster = unregisterCluster;
function stopClusterOnYarn(clusterId, keytabFilename) {
    return function (dispatch) {
        Remote.stopClusterOnYarn(clusterId, keytabFilename, function (error) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                return;
            }
            dispatch(fetchClusters());
        });
    };
}
exports.stopClusterOnYarn = stopClusterOnYarn;
function fetchProjects() {
    return function (dispatch) {
        Remote.getProjects(0, 1000, function (error, res) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                return;
            }
            dispatch(receiveProjects(res));
        });
    };
}
exports.fetchProjects = fetchProjects;
//# sourceMappingURL=projects.actions.js.map