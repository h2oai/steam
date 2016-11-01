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
import {Model} from "../../Proxy/Proxy";
import {Workgroup} from "../../Proxy/Proxy";
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
export const RECEIVE_WORKGROUPS = 'RECEIVE_WORKGROUPS';
export const REQUEST_DELETE_WORKGROUP = 'REQUEST_DELETE_WORKGROUP';
export const RECEIVE_DELETE_WORKGROUP = 'RECEIVE_DELETE_WORKGROUP';
export const REQUEST_CREATE_WORKGROUP = 'REQUEST_CREATE_WORKGROUP';
export const RECEIVE_CREATE_WORKGROUP = 'RECEIVE_CREATE_WORKGROUP';

export function requestCreateWorkgroup() {
  return {
    type: REQUEST_CREATE_WORKGROUP
  };
};
export function receiveCreateWorkgroup() {
  return {
    type: RECEIVE_CREATE_WORKGROUP
  };
};
export function requestDeleteWorkgroup(workgroupId) {
  return {
    type: REQUEST_DELETE_WORKGROUP,
    workgroupId
  };
};
export function receiveDeleteWorkgroup(workgroupId) {
  return {
    type: RECEIVE_DELETE_WORKGROUP,
    workgroupId
  };
};
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

export function receiveWorkgroups(workgroups) {
  return {
    type: RECEIVE_WORKGROUPS,
    workgroups
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

export function createProjectAsync(dispatch, name: string, modelCategory: string) {
  return new Promise((resolve, reject) => {
    Remote.createProject(name, '', modelCategory, (error, res) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
        reject(error);
        return;
      }
      createWorkgroupAsync(dispatch, name).then(() => {
        dispatch(createProjectCompleted(res));
        resolve(res);
      });
    });
  });
}
export function createWorkgroupAsync(dispatch, name: string) {
  return new Promise((resolve, reject) => {
    dispatch(requestCreateWorkgroup());
    Remote.createWorkgroup(name, name, (error: Error, workgroupId: number) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, 'Create Error', error.toString(), null));
        reject(error);
        return;
      }
      dispatch(receiveCreateWorkgroup());
      resolve();
    });
  });
}
export function createProject(name: string, modelCategory: string) {
  return (dispatch) => {
    createProjectAsync(dispatch, name, modelCategory);
  };
}

export function importModelFromCluster(clusterId: number, projectId: number, modelName: string) {
  return (dispatch) => {

    function importModelFromClusterAsync(clusterId: number, projectId: number, modelName: string) {
      return new Promise(function(resolve, reject) {
        Remote.importModelFromCluster(clusterId, projectId, modelName, modelName, (error, modelId: number) => {
          if (error) {
            reject(error);
          }
          resolve(modelId);
        });
      });
    }

    function getAlgoAsync(modelId: number) {
      return new Promise((resolve, reject) => {
        Remote.getModel(modelId, (error: Error, model: Model) => {
          if (error) {
            reject(error);
          }
          resolve({modelId, algo: model.algorithm});
        });
      });
    }

    function checkCanMojoAsync(algo: string, modelId: number) {
      return new Promise((resolve, reject) => {
        Remote.checkMojo(algo, (error: Error, canMojo: boolean) => {
          if (error) {
            reject(error);
          }
          resolve({canMojo, modelId});
        });
      });
    }

    function importMojoAsync(modelId: number) {
      return new Promise((resolve, reject) => {
        Remote.importModelMojo(modelId, (error) => {
          if (error) {
            reject(error);
          }
          resolve();
        });
      });
    }

    function importPojoAsync(modelId: number) {
      return new Promise((resolve, reject) => {
        Remote.importModelPojo(modelId, (error) => {
          if (error) {
            reject(error);
          }
          resolve();
        });
      });
    }

    importModelFromClusterAsync(clusterId, projectId, modelName)
      .then((modelId: number) => getAlgoAsync(modelId))
      .catch((error) => openNotification(NotificationType.Error, 'Load Error', error.toString(), null))
      .then((algoRes: any) => checkCanMojoAsync(algoRes.algo, algoRes.modelId))
      .catch((error) => openNotification(NotificationType.Error, 'Load Error', error.toString(), null))
      .then((canMojoRes: any) => {
        if (canMojoRes.canMojo && localStorage.getItem("mojoPojoSelection") !== "pojo") {
          importMojoAsync(canMojoRes.modelId)
            .then(() => {
              dispatch(importModelFromClusterCompleted(canMojoRes.modelId));
            })
            .catch((error) => openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
        } else {
          importPojoAsync(canMojoRes.modelId)
            .then(() => {
              dispatch(importModelFromClusterCompleted(canMojoRes.modelId));
            })
            .catch((error) => openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
        }
      })
      .catch((error) => openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
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
      createProjectAsync(dispatch, projectName, modelCategory).then((projectId) => {
        dispatch(importModelsFromCluster(clusterId, projectId as number, models)).then(() => {
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

function deleteWorkgroupAsync(dispatch, workgroupId) {
  return new Promise((deleteWorkgroupResolve, deleteWorkgroupReject) => {
    dispatch(requestDeleteWorkgroup(workgroupId));
    Remote.deleteWorkgroup(workgroupId, (error: Error) => {
      if (error) {
        openNotification(NotificationType.Error, 'Load Error', error.toString(), null);
        deleteWorkgroupReject(error);
        return;
      }
      dispatch(receiveDeleteWorkgroup(workgroupId));
      deleteWorkgroupResolve(workgroupId);
      return;
    });
  });
}
function fetchWorkgroupsAsync(dispatch) {
  return new Promise((fetchWorkgroupResolve, fetchWorkgroupReject) => {
    Remote.getWorkgroups(0, 1000, (error, res) => {
      if (error) {
        openNotification(NotificationType.Error, 'Load Error', error.toString(), null);
        fetchWorkgroupReject();
        return;
      }
      dispatch(receiveWorkgroups(<Workgroup[]> res));
      fetchWorkgroupResolve(res);
    });
  });
}
export function fetchWorkgroups() {
  return (dispatch) => {
    fetchWorkgroupsAsync(dispatch);
  };
}

export function deleteProject(projectId: number) {
  return (dispatch, getState) => {
    fetchWorkgroupsAsync(dispatch).then((workgroups) => {
      dispatch(requestDeleteProject(projectId));

      let deleteWorkgroupPromises = [];
      let state = getState();
      for (let project of state.projects.availableProjects) {
        if (project.id === projectId) {
          for (let workgroup of workgroups as Array<Workgroup>) {
            if (workgroup.name === project.name) {
              deleteWorkgroupPromises.push(deleteWorkgroupAsync(dispatch, workgroup.id));
            }
          }
        }
      }

      Promise.all(deleteWorkgroupPromises).then((dwpResult) => {
        Remote.deleteProject(projectId, (error) => {
          if (error) {
            dispatch(openNotification(NotificationType.Error, 'Delete error', error.toString(), null));
            dispatch(receiveDeleteProject(projectId, false));
            return;
          }
          dispatch(receiveDeleteProject(projectId, true));
          dispatch(fetchProjects());
        });
      });
    });
  };
}
