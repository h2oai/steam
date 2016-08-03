/**
 * Created by Jeff Fohl <jfohl@h2o.ai> on 7/30/16.
 */
import * as Remote from '../../Proxy/Proxy';
export const FETCH_LABELS = 'FETCH_LABELS';
export const RECEIVE_LABELS = 'RECEIVE_LABELS';
export const CREATE_LABEL = 'CREATE_LABEL';
export const RECEIVE_CREATE_LABEL = 'RECEIVE_CREATE_LABEL';
export const UPDATE_LABEL = 'UPDATE_LABEL';
export const RECEIVE_UPDATE_LABEL = 'RECEIVE_UPDATE_LABEL';
export const DELETE_LABEL = 'DELETE_LABEL';
export const RECEIVE_DELETE_LABEL = 'RECEIVE_DELETE_LABEL';

export const requestLabels = () => {
  return {
    type: FETCH_LABELS
  };
};

export function fetchLabels(projectId: number) {
  return (dispatch) => {
    dispatch(requestLabels());
    Remote.getLabelsForProject(projectId, (error, res) => {
      dispatch(receiveLabels(res, projectId));
    });
  };
}

export function receiveLabels(labels: any[], projectId: number) {
  return {
    type: RECEIVE_LABELS,
    projectId,
    labels
  };
};

export function createLabel(projectId: number, name: string, description: string) {
  return (dispatch) => {
    return new Promise((resolve, reject) => {
      Remote.createLabel(projectId, name, description, (error, res) => {
        if (error) {
          reject(error);
          return;
        }
        dispatch(receiveCreateLabel(res, projectId, name, description));
        resolve(res);
      });
    });
  };
}

export function receiveCreateLabel(id: number, projectId: number, name: string, description: string) {
  return {
    type: RECEIVE_CREATE_LABEL,
    projectId,
    label: {
      id,
      name,
      description
    }
  };
};

export function updateLabel(labelId: number, projectId: number, name: string, description: string) {
  return (dispatch) => {
    return new Promise((resolve, reject) => {
      Remote.updateLabel(labelId, name, description, (error) => {
        if (error) {
          reject(error);
          return;
        }
        dispatch(receiveUpdateLabel(labelId, projectId, name, description));
        resolve();
      });
    });
  };
}

export function receiveUpdateLabel(id: number, projectId: number, name: string, description: string) {
  return {
    type: RECEIVE_UPDATE_LABEL,
    projectId,
    label: {
      id,
      name,
      description
    }
  };
};

export function deleteLabel(labelId: number) {
  return (dispatch) => {
    return new Promise((resolve, reject) => {
      Remote.deleteLabel(labelId, (error) => {
        if (error) {
          reject(error);
          return;
        }
        //dispatch(receiveDeleteLabel(labelId));
        resolve();
      });
    });
  };
}

export function receiveDeleteLabel(labelId) {
  return {
    type: RECEIVE_DELETE_LABEL,

  };
};
