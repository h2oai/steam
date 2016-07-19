/**
 * Created by justin on 7/18/16.
 */

import * as Remote from '../../Proxy/proxy';
import { Model } from '../../Proxy/proxy';

export const REQUEST_CLUSTERS = 'REQUEST_CLUSTERS';
export const RECEIVE_CLUSTERS = 'RECEIVE_CLUSTERS';
export const REQUEST_MODELS = 'REQUEST_MODELS';
export const RECEIVE_MODELS = 'RECEIVE_MODELS';
export const CREATE_PROJECT_COMPLETED = 'CREATE_PROJECT_COMPLETED';
export const IMPORT_MODEL_FROM_CLUSTER_COMPLETED = 'IMPORT_MODEL_FROM_CLUSTER_COMPLETED';

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
  }
}

export function importModelFromClusterCompleted(model) {
  return {
    type: IMPORT_MODEL_FROM_CLUSTER_COMPLETED,
    model
  }
}

export function fetchModelsFromCluster(clusterId: number) {
  return (dispatch) => {
    dispatch(requestModels());
    Remote.getClusterModels(clusterId, (error, res) => {
      dispatch(receiveModelsFromCluster(res));
    });
  };
}

export function createProject(name: string) {
  return (dispatch) => {
    Remote.createProject(name, '', (error, res) => {
      dispatch(createProjectCompleted(res));
    });
  };
}

export function importModelFromCluster(clusterId: number, projectId: number, modelName: string) {
  return (dispatch) => {
    console.log(clusterId, projectId, modelName);
    Remote.importModelFromCluster(clusterId, projectId, modelName, (error, res) => {
      dispatch(importModelFromClusterCompleted(res))
    });
  }
}

export function createProjectAndImportModelsFromCluster(projectName: string, clusterId: number, models: string[]) {
  return (dispatch) => {
    Remote.createProject(name, '', (error, res) => {
      dispatch(createProjectCompleted(res));
      for (var i = 0; i < models.length; i++) {
        Remote.importModelFromCluster(clusterId, res, models[i], (error, res) => {
          dispatch(importModelFromClusterCompleted(res))
        });
      }
    });
  };
}
