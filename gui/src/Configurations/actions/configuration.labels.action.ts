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

export function fetchLabels(projectId: number) {
  return (dispatch) => {
    dispatch(requestLabels());
    Remote.getLabelsForProject(projectId, (error, res) => {
      dispatch(receiveLeaderboard(res));
    });
  };
}

export function requestLabels() {
  return {
    type: FETCH_LABELS
  };
};

export function receiveLabels(labels) {
  return {
    type: RECEIVE_LABELS,
    labels
  };
};

export function createLabel(label) {
  return {
    type: CREATE_LABEL,
    label
  };
};

export function receiveCreateLabel(label) {
  return {
    type: RECEIVE_UPDATE_LABEL
  };
};

export function updateLabel(label) {
  return {
    type: UPDATE_LABEL
  };
};

export function receiveUpdateLabel(label) {
  return {
    type: RECEIVE_UPDATE_LABEL
  };
};

export function deleteLabel(labelId) {
  return {
    type: DELETE_LABEL
  };
};

export function receiveDeleteLabel(response) {
  return {
    type: RECEIVE_DELETE_LABEL
  };
};
