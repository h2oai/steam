/**
 * Created by justin on 7/18/16.
 */

import * as Remote from '../../Proxy/Proxy';
import { Project } from '../../Proxy/Proxy';
import { openNotification } from '../../App/actions/notification.actions';

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
}

export function resetClusterSelection()  {
  return {
    type: RESET_CLUSTER_SELECTION
  };
};

export function fetchClusters() {
  return (dispatch) => {
    dispatch(requestClusters());
    Remote.getClusters(0, 1000, (error, res) => {
      if (error) {
        openNotification('error', 'There was an error retrieving your list of clusters', null);
      }
      dispatch(receiveClusters(res));
    });
  };
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
        dispatch(openNotification('error', error.toString(), null));
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
          dispatch(openNotification('error', error.toString(), null));
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
        dispatch(openNotification('error', error.toString(), null));
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
        dispatch(openNotification('error', error.toString(), null));
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
          dispatch(openNotification('error', error.toString(), null));
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
          dispatch(openNotification('error', error.toString(), null));
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
        dispatch(openNotification('error', error.toString(), null));
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
        dispatch(openNotification('error', error.toString(), null));
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
        dispatch(openNotification('error', error.toString(), null));
        return;
      }
      dispatch(receiveProjects(<Project[]> res));
    });
  };
}
