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
import { openNotification } from './notification.actions';
import { NotificationType } from '../components/Notification';

export const REQUEST_ENTITY_IDS = 'REQUEST_ENTITY_IDS';
export const RECEIVE_ENTITY_IDS = 'RECEIVE_ENTITY_IDS';
export const REQUEST_IS_ADMIN = 'REQUEST_IS_ADMIN';
export const RECEIVE_IS_ADMIN = 'RECEIVE_IS_ADMIN';

export function requestIsAdmin() {
  return {
    type: REQUEST_IS_ADMIN
  };
}
export function receiveIsAdmin(isAdmin) {
  return {
    type: RECEIVE_IS_ADMIN,
    isAdmin
  };
}
export function requestEntityIds() {
  return {
    type: REQUEST_ENTITY_IDS
  };
}

export function receiveEntityIds(response) {
  return {
    type: RECEIVE_ENTITY_IDS,
    response: response
  };
}

export function fetchEntityIds() {
  return (dispatch, getState) => {
    dispatch(requestEntityIds());

    return new Promise((resolve, reject) => {
      Remote.getAllEntityTypes((error, res) => {
        if (error) {
          dispatch(openNotification(NotificationType.Error, 'Load Error', 'There was an error retrieving permissions list', null));
          reject(error);
          return;
        }
        dispatch(receiveEntityIds(res));
        resolve(res);
      });
    });
  };
}
export function fetchIsAdmin() {
  return (dispatch, getState) => {
    dispatch(requestIsAdmin());
    return new Promise((resolve, reject) => {
      Remote.checkAdmin((error: Error, isAdmin: boolean) => {
        if (error) {
          dispatch(openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
          reject(error);
          return;
        }
        dispatch(receiveIsAdmin(isAdmin));
        resolve(isAdmin);
      });
    });
  };
};
