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

import * as Remote from '../../Proxy/Proxy';
import * as _ from 'lodash';
import { fetchEntityIds } from '../../App/actions/global.actions';
import { openNotification } from '../../App/actions/notification.actions';

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
      openNotification('error', error.toString(), null);
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
      openNotification('error', error.toString(), null);
      return;
    }

    let identityPromises = [];
    let toReturn = [];
    for (let label of labels) {
      identityPromises.push(new Promise((resolve, reject) => {
        Remote.getIdentitiesForEntity(state.global.entityIds.label, label.id, (identitiesError, identitiesRes) => {
          if (identitiesError) {
            openNotification('error', identitiesError.toString(), null);
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
