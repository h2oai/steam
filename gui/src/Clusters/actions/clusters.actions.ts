/**
 * Created by justin on 8/17/16.
 */

import * as Remote from '../../Proxy/Proxy';
import { openNotification } from '../../App/actions/notification.actions';
import { NotificationType } from '../../App/components/Notification';

export const RECEIVE_ENGINES = 'RECEIVE_ENGINES';
export const START_FETCH_CONFIG = 'START_FETCH_CONFIG';
export const FETCH_CONFIG_COMPLETED = 'FETCH_CONFIG_COMPLETED';
export const START_UPLOAD_ENGINE = 'START_UPLOAD_ENGINE';
export const UPLOAD_ENGINE_COMPLETED = 'UPLOAD_ENGINE_COMPLETED';
export const START_CLUSTER = 'START_CLUSTER';
export const START_CLUSTER_COMPLETED = 'START_CLUSTER_COMPLETED';
export const START_GET_ENGINES = 'START_GET_ENGINES';

export function receiveEngines(engines) {
  return {
    type: RECEIVE_ENGINES,
    engines
  };
}

export function startCluster() {
  return {
    type: START_CLUSTER
  };
}

export function startGetEngines() {
  return {
    type: START_GET_ENGINES
  };
}

export function startClusterCompleted(response: number | string) {
  return {
    type: START_CLUSTER_COMPLETED,
    response
  };
}

export function fetchConfig() {
  return {
    type: START_FETCH_CONFIG
  };
}

export function fetchConfigCompleted(config) {
  return {
    type: FETCH_CONFIG_COMPLETED,
    config
  };
}

export function startUploadEngine() {
  return {
    type: START_UPLOAD_ENGINE
  };
}

export function uploadEngineCompleted(response) {
  return {
    type: UPLOAD_ENGINE_COMPLETED,
    response
  };
}

export function uploadEngine(file) {
  if (!file) {
    openNotification(NotificationType.Error, "File Error", 'No engine file selected.', null);
  }
  return (dispatch) => {
    dispatch(startUploadEngine());
    dispatch(openNotification(NotificationType.Info, "Update", 'Uploading engine...', null));
    let data = new FormData();
    data.append('file', file.files[0]);
    fetch(`/upload?type=engine`, {
      credentials: 'include',
      method: 'post',
      body: data
    }).then(() => {
      dispatch(openNotification(NotificationType.Confirm, "Success", 'Engine uploaded', null));
      dispatch(uploadEngineCompleted(null));
      dispatch(getEngines());
    }).catch((error) => {
      dispatch(uploadEngineCompleted(error));
      dispatch(openNotification(NotificationType.Error, "Error", error.toString(), null));
    });
  };
}

export function startYarnCluster(clusterName, engineId, size, memory, keytab) {
  if (!clusterName || !engineId || !size || !memory) {
    openNotification(NotificationType.Error, "Error", 'All fields are required', null);
  }
  return (dispatch) => {
    dispatch(startCluster());
    dispatch(openNotification(NotificationType.Info, "Update", 'Connecting to YARN...', null));
    Remote.startClusterOnYarn(clusterName, engineId, size, memory, keytab, (error, clusterId) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, "Error", error.toString(), null));
        dispatch(startClusterCompleted(error.toString()));
        return;
      }
      dispatch(startClusterCompleted(clusterId));
      dispatch(openNotification(NotificationType.Confirm, "Success", 'Cluster Launched', null));
    });
  };
}

export function getEngines() {
  return (dispatch) => {
    dispatch(startGetEngines());
    Remote.getEngines((error, engines) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, 'Error', error.toString(), null));
        dispatch(receiveEngines(null));
        return;
      }
      dispatch(receiveEngines(engines));
    });
  };
}

export function getConfig() {
  return (dispatch) => {
    dispatch(fetchConfig());
    Remote.getConfig((error, config) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, 'Error', error.toString(), null));
        dispatch(fetchConfigCompleted(null));
        return;
      }
      dispatch(fetchConfigCompleted(config));
    });
  };
}
