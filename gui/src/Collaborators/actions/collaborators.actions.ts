import * as Remote from '../../Proxy/Proxy';
import * as _ from 'lodash';
import { fetchEntityIds } from '../../App/actions/global.actions';
import { openNotification } from '../../App/actions/notification.actions';
import { NotificationType } from '../../App/components/Notification';

export const REQUEST_MEMBERS_FOR_PROJECT = 'REQUEST_MEMBERS_FOR_PROJECT';
export const RECEIVE_MEMBERS_FOR_PROJECT = 'RECEIVE_MEMBERS_FOR_PROJECT';
export const REQUEST_LABELS_FOR_PROJECT = 'REQUEST_LABELS_FOR_PROJECT';
export const RECEIVE_LABELS_FOR_PROJECT = 'RECEIVE_LABELS_FOR_PROJECT';

export function requestMembersForProject() {
  return {
    type: REQUEST_MEMBERS_FOR_PROJECT
  };
}
export function receiveMembersForProject(response) {
  return {
    type: RECEIVE_MEMBERS_FOR_PROJECT,
    members: response
  };
}

export function requestLabelsForProject() {
  return {
    type: REQUEST_LABELS_FOR_PROJECT
  };
}
export function receiveLabelsForProject(response) {
  return {
    type: RECEIVE_LABELS_FOR_PROJECT,
    labels: response
  };
}

function _fetchMembersForProject(dispatch, state) {
  Remote.getIdentitiesForEntity(state.global.entityIds.project, state.projects.project.id, (error, res) => {
    if (error) {
      openNotification(NotificationType.Error, 'Load Error', error.toString(), null);
      return;
    }
    dispatch(receiveMembersForProject(res));
   });
}
export function fetchMembersForProject() {
  return (dispatch, getState) => {
    dispatch(requestMembersForProject());
    let state = getState();

    if (_.isEmpty(state.global.entityIds)) {
      dispatch(fetchEntityIds()).then(() => {
        state = getState();
        _fetchMembersForProject(dispatch, state);
      });
    } else {
      _fetchMembersForProject(dispatch, state);
    }
  };
}

function _fetchLabelsForProject(dispatch, state) {
  Remote.getLabelsForProject(state.projects.project.id, (error, labels) => {
    if (error) {
      openNotification(NotificationType.Error, 'Load Error', error.toString(), null);
      return;
    }

    let identityPromises = [];
    let toReturn = [];
    for (let label of labels) {
      identityPromises.push(new Promise((resolve, reject) => {
        Remote.getIdentitiesForEntity(state.global.entityIds.label, label.id, (identitiesError, identitiesRes) => {
          if (identitiesError) {
            openNotification(NotificationType.Error, 'Load Error', identitiesError.toString(), null);
            reject(identitiesError.toString());
            return;
          }
          toReturn.push(_.assign({}, label, { identities: identitiesRes }));
          resolve();
        });
      }));
    }
    Promise.all(identityPromises).then((results) => {
      dispatch(receiveLabelsForProject(toReturn));
    });
  });
}
export function fetchLabelsForProject() {
  return (dispatch, getState) => {
    dispatch(requestLabelsForProject());
    let state = getState();

    if (_.isEmpty(state.global.entityIds)) {
        dispatch(fetchEntityIds()).then(() => {
          state = getState();
          _fetchLabelsForProject(dispatch, state);
        });
    } else {
      _fetchLabelsForProject(dispatch, state);
    }
  };
}
