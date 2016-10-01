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

/**
 * Created by justin on 7/18/16.
 */

import * as Remote from '../../Proxy/Proxy';
import * as _ from 'lodash';
import { Project, UserRole } from '../../Proxy/Proxy';
import { fetchEntityIds } from '../../App/actions/global.actions';
import { openNotification } from '../../App/actions/notification.actions';
import { NotificationType } from '../../App/components/Notification';
import {ClusterStatus, Cluster} from "../../Proxy/Proxy";

export const SET_CURRENT_PROJECT = 'SET_CURRENT_PROJECT';
export const REQUEST_CLUSTERS = 'REQUEST_CLUSTERS';
export const RECEIVE_CLUSTERS = 'RECEIVE_CLUSTERS';
export const RESET_CLUSTER_SELECTION = 'RESET_CLUSTER_SELECTION';
export const REQUEST_MODELS = 'REQUEST_MODELS';
export const RECEIVE_MODELS = 'RECEIVE_MODELS';
export const CREATE_PROJECT_COMPLETED = 'CREATE_PROJECT_COMPLETED';
export const IMPORT_MODEL_FROM_CLUSTER_COMPLETED = 'IMPORT_MODEL_FROM_CLUSTER_COMPLETED';
export const RECEIVE_PROJECTS = 'RECEIVE_PROJECTS';
export const REQUEST_DATASETS_FROM_CLUSTER = 'REQUEST_DATASETS_FROM_CLUSTER';
export const RECEIVE_DATASETS_FROM_CLUSTER = 'RECEIVE_DATASETS_FROM_CLUSTER';
export const RECEIVE_MODELS_FROM_PROJECT = 'RECEIVE_MODELS_FROM_PROJECT';
export const RECEIVE_PROJECT = 'RECEIVE_PROJECT';
export const REQUEST_DELETE_PROJECT = 'REQUEST_DELETE_PROJECT';
export const RECEIVE_DELETE_PROJECT = 'RECEIVE_DELETE_PROJECT';
export const REGISTER_CLUSTER_ERROR = 'REGISTER_CLUSTER_ERROR';

export function requestDeleteProject(projectId: number) {
  return {
    type: REQUEST_DELETE_PROJECT,
    projectId
  };
};
export function receiveDeleteProject(projectId: number, successful: boolean) {
  return {
    type: RECEIVE_DELETE_PROJECT,
    projectId,
    successful
  };
};

export function setCurrentProject(projectId) {
  return {
    type: SET_CURRENT_PROJECT,
    projectId
  };
};

export const requestClusters = () => {
  return {
    type: REQUEST_CLUSTERS
  };
};

export function receiveClusters(clusters) {
  return {
    type: RECEIVE_CLUSTERS,
    clusters
  };
};

export function resetClusterSelection()  {
  return {
    type: RESET_CLUSTER_SELECTION
  };
};

export function registerClusterError(message) {
  return {
    type: REGISTER_CLUSTER_ERROR,
    message
  };
};

export function fetchClusters() {
  return (dispatch, getState) => {
    dispatch(requestClusters());
    let state = getState();

    if (_.isEmpty(state.global.entityIds)) {
      dispatch(fetchEntityIds()).then(() => {
        _fetchClusters(dispatch, getState);
      });
    } else {
      _fetchClusters(dispatch, getState);
    }

  };
}
function _fetchClusters(dispatch, getState) {
    Remote.getClusters(0, 1000, (error, clusters: Cluster[]) => {
      if (error) {
        openNotification(NotificationType.Error, "Load Error", error.toString(), null);
        return;
      }

      let identityPromises = [];
      let detailsPromise = [];
      let toReturn = [];
      let state = getState();

      for (let cluster of clusters) {
        identityPromises.push(new Promise((resolve, reject) => {
          Remote.getIdentitiesForEntity(state.global.entityIds.cluster, cluster.id, (identitiesError: Error, users: UserRole[]) => {
            if (identitiesError) {
              openNotification(NotificationType.Error, "Load Error", identitiesError.toString(), null);
              reject(identitiesError.toString());
              return;
            }

            detailsPromise.push(new Promise((resolve, reject) => {
              Remote.getClusterStatus(cluster.id, (statusError: Error, clusterStatus: ClusterStatus) => {
                if (statusError) {
                  openNotification(NotificationType.Error, "Load Error", statusError.toString(), null);
                  reject(statusError.toString());
                  return;
                }

                toReturn.push(_.assign({}, cluster, { identities: users }, {status: clusterStatus}));
                resolve();
              });
            }).then(() => resolve()));
          });
        }));
      }
      Promise.all(identityPromises).then((results) => {
        dispatch(receiveClusters(toReturn));
      });
    });
}

export const requestModels = () => {
  return {
    type: REQUEST_MODELS
  };
};

export function receiveModelsFromCluster(models) {
  return {
    type: RECEIVE_MODELS,
    models
  };
}

export function createProjectCompleted(project) {
  return {
    type: CREATE_PROJECT_COMPLETED,
    project
  };
}

export function receiveProject(project) {
  return {
    type: RECEIVE_PROJECT,
    project
  };
}

export function importModelFromClusterCompleted(model) {
  return {
    type: IMPORT_MODEL_FROM_CLUSTER_COMPLETED,
    model
  };
}

export function receiveProjects(projects) {
  for (let project of projects) {
    project.isDeleteInProgress = false;
  }
  return {
    type: RECEIVE_PROJECTS,
    projects
  };
}

export function requestDatasetsFromCluster() {
  return {
    type: REQUEST_DATASETS_FROM_CLUSTER
  };
}

export function receiveDatasetsFromCluster(datasets) {
  return {
    type: RECEIVE_DATASETS_FROM_CLUSTER,
    datasets
  };
}

export function receiveModelsFromProject(models) {
  return {
    type: RECEIVE_MODELS_FROM_PROJECT,
    models
  };
}

export function fetchModelsFromProject(projectId: number) {
  return (dispatch) => {
    Remote.getModels(projectId, 0, 5, (error, res) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
        return;
      }
      dispatch(receiveModelsFromProject(res));
    });
  };
}

export function fetchProject(projectId: number) {
  return (dispatch) => {
    return new Promise((resolve, reject) => {
      Remote.getProject(projectId, (error, res) => {
        if (error) {
          dispatch(openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
          reject(error);
          return;
        }
        dispatch(receiveProject(res));
        resolve(res);
      });
    });
  };
}

export function fetchModelsFromCluster(clusterId: number, frameKey: string) {
  return (dispatch) => {
    dispatch(requestModels());
    Remote.getModelsFromCluster(clusterId, frameKey, (error, res) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
      }
      dispatch(receiveModelsFromCluster(res));
    });
  };
}

export function fetchDatasetsFromCluster(clusterId: number) {
  return (dispatch) => {
    dispatch(requestDatasetsFromCluster());
    Remote.getDatasetsFromCluster(clusterId, (error, res) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
        return;
      }
      dispatch(receiveDatasetsFromCluster(res));
    });
  };
}

export function createProject(name: string, modelCategory: string) {
  return (dispatch) => {
    return new Promise((resolve, reject) => {
      Remote.createProject(name, '', modelCategory, (error, res) => {
        if (error) {
          dispatch(openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
          reject(error);
          return;
        }
        dispatch(createProjectCompleted(res));
        resolve(res);
      });
    });
  };
}

export function importModelFromCluster(clusterId: number, projectId: number, modelName: string) {
  return (dispatch) => {
    return new Promise((resolve, reject) => {
      Remote.importModelFromCluster(clusterId, projectId, modelName, modelName, (error, res) => {
        if (error) {
          dispatch(openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
          reject(error);
          return;
        }
        dispatch(importModelFromClusterCompleted(res));
        resolve(res);
      });
    });
  };
}

export function importModelsFromCluster(clusterId: number, projectId: number, models: string[]) {
  return (dispatch) => {
    let promises = [];
    return new Promise((resolve, reject) => {
      models.map((modelName) => {
        promises.push(dispatch(importModelFromCluster(clusterId, projectId, modelName)));
      });
      Promise.all(promises).then(() => {
        resolve(projectId);
      });
    });
  };
}

export function createProjectAndImportModelsFromCluster(projectName: string, clusterId: number, modelCategory: string, models: string[]) {
  return (dispatch) => {
    return new Promise((resolve, reject) => {
      dispatch(createProject(projectName, modelCategory)).then((projectId) => {
        dispatch(importModelsFromCluster(clusterId, projectId, models)).then(() => {
          resolve(projectId);
        });
      });
    });
  };
}

export function registerCluster(address: string) {
  return (dispatch) => {
    Remote.registerCluster(address, (error, res) => {
      if (error) {
        dispatch(registerClusterError(error.message));
        return;
      }
      dispatch(fetchClusters());
    });
  };
}

export function unregisterCluster(clusterId: number) {
  return (dispatch) => {
    Remote.unregisterCluster(clusterId, (error) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
        return;
      }
      dispatch(fetchClusters());
    });
  };
}

export function stopClusterOnYarn(clusterId: number, keytabFilename: string) {
  return (dispatch) => {
    Remote.stopClusterOnYarn(clusterId, keytabFilename, (error) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
        return;
      }
      dispatch(fetchClusters());
    });
  };
}

export function fetchProjects() {
  return (dispatch) => {
    Remote.getProjects(0, 1000, (error, res) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
        return;
      }
      dispatch(receiveProjects(<Project[]> res));
    });
  };
}

export function deleteProject(projectId: number) {
  return (dispatch) => {
    dispatch(requestDeleteProject(projectId));
    Remote.deleteProject(projectId, (error) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, 'Delete error', error.toString(), null));
        dispatch(receiveDeleteProject(projectId, false));
        return;
      }
      dispatch(receiveDeleteProject(projectId, true));
      dispatch(fetchProjects());
    });
  };
}
