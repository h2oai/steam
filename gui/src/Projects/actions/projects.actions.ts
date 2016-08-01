/**
 * Created by justin on 7/18/16.
 */

import * as Remote from '../../Proxy/Proxy';
import { Project } from '../../Proxy/Proxy';

export const REQUEST_CLUSTERS = 'REQUEST_CLUSTERS';
export const RECEIVE_CLUSTERS = 'RECEIVE_CLUSTERS';
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

export function fetchClusters() {
  return (dispatch) => {
    dispatch(requestClusters());
    Remote.getClusters(0, 5, (error, res) => {
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
      dispatch(receiveModelsFromProject(res));
    });
  };
}

export function fetchProject(projectId: number) {
  return (dispatch) => {
    Remote.getProject(projectId, (error, res) => {
      dispatch(receiveProject(res));  
    });  
  };
}

export function fetchModelsFromCluster(clusterId: number, frameKey: string) {
  return (dispatch) => {
    dispatch(requestModels());
    Remote.getModelsFromCluster(clusterId, frameKey, (error, res) => {
      dispatch(receiveModelsFromCluster(res));
    });
  };
}

export function fetchDatasetsFromCluster(clusterId: number) {
  return (dispatch) => {
    dispatch(requestDatasetsFromCluster());
    Remote.getDatasetsFromCluster(clusterId, (error, res) => {
      dispatch(receiveDatasetsFromCluster(res));
    });
  };
}

export function createProject(name: string, modelCategory: string) {
  return (dispatch) => {
    return new Promise((resolve, reject) => {
      Remote.createProject(name, '', modelCategory, (error, res) => {
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
        dispatch(importModelFromClusterCompleted(res));
        resolve(res);
      });
    });
  };
}

export function createProjectAndImportModelsFromCluster(projectName: string, clusterId: number, modelCategory: string, models: string[]) {
  return (dispatch) => {
    return new Promise((resolve, reject) => {
      let promises = [];
      dispatch(createProject(projectName, modelCategory)).then((projectId) => {
        models.map((modelName) => {
          promises.push(dispatch(importModelFromCluster(clusterId, projectId, modelName)));
        });
        Promise.all(promises).then(() => {
          resolve(projectId);
        });
      });
    });
  };
}

export function registerCluster(address: string) {
  return (dispatch) => {
    Remote.registerCluster(address, (error, res) => {
      dispatch(fetchClusters());
    });
  };
}

export function fetchProjects() {
  return (dispatch) => {
    Remote.getProjects(0, 5, (error, res) => {
      dispatch(receiveProjects(<Project[]> res));
    });
  };
}
