/**
 * Created by justin on 8/17/16.
 */

import * as Remote from '../../Proxy/Proxy';
import { openNotification } from '../../App/actions/notification.actions';

export const RECEIVE_ENGINES = 'RECEIVE_ENGINES';
export const RECEIVE_CONFIG = 'RECEIVE_CONFIG';

export function receiveEngines(engines) {
  return {
    type: RECEIVE_ENGINES,
    engines
  };
}

export function receiveConfig(config) {
  return {
    type: RECEIVE_CONFIG,
    config
  };
}

export function uploadEngine(file) {
  if (!file) {
    openNotification('error', 'No engine file selected.', null);
  }
  return (dispatch) => {
    dispatch(openNotification('info', 'Uploading engine...', null));
    let data = new FormData();
    data.append('file', file.files[0]);
    fetch(`/upload?type=engine`, {
      credentials: 'include',
      method: 'post',
      body: data
    }).then((res) => {
      dispatch(openNotification('success', 'Engine uploaded', null));
      dispatch(getEngines());
    });
  };
}

export function startYarnCluster(clusterName, engineId, size, memory, keytab = '') {
  if (!clusterName || !engineId || !size || !memory) {
    openNotification('error', 'All fields are required', null);
  }
  return (dispatch) => {
    dispatch(openNotification('info', 'Connecting to YARN...', null));
    Remote.startClusterOnYarn(clusterName, engineId, size, memory, keytab, (error, res) => {
      if (error) {
        dispatch(openNotification('error', error.toString(), null));
        return;
      }
      dispatch(openNotification('success', 'Cluster Launched', null));
    });
  };
}

export function getEngines() {
  return (dispatch) => {
    Remote.getEngines((error, engines) => {
      if (error) {
        dispatch(openNotification('error', error.toString(), null));
        return;
      }
      dispatch(receiveEngines(engines));
    });
  };
}

export function getConfig() {
  return (dispatch) => {
    Remote.getConfig((error, config) => {
      if (error) {
        dispatch(openNotification('error', error.toString(), null));
        return;
      }
      dispatch(receiveConfig(config));
    });
  };
}
