import * as Remote from '../../Proxy/Proxy';
import * as _ from 'lodash';

export const REQUEST_MEMBERS = 'REQUEST_MEMBERS';
export const RECEIVE_MEMBERS = 'RECEIVE_MEMBERS';
export const REQUEST_LABELS = 'REQUEST_LABELS';
export const RECEIVE_LABELS = 'RECEIVE_LABELS';

export function requestMembers() {
  return {
    type: REQUEST_MEMBERS
  };
}
export function receiveMembers() {
  return {
    type: RECEIVE_MEMBERS
  };
}

export function requestLabels() {
  return {
    type: REQUEST_LABELS
  };
}
export function receiveLabels() {
  return {
    type: RECEIVE_LABELS
  };
}

export function fetchMembers() {
  return (dispatch) => {
    dispatch(requestMembers());
    /*Remote.getPackages(projectId, (error, res) => {
      if (error) {
        dispatch(openNotification('error', error.toString(), null));
        return;
      }
      dispatch(receivePackages(res));
    });*/
    dispatch(receiveMembers());
  };
}
export function fetchLabels() {
  return (dispatch) => {
    dispatch(requestLabels());
    /*Remote.getPackages(projectId, (error, res) => {
     if (error) {
     dispatch(openNotification('error', error.toString(), null));
     return;
     }
     dispatch(receivePackages(res));
     });*/
    dispatch(receiveLabels());
  };
}
