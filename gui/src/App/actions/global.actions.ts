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

export const REQUEST_ENTITY_IDS = 'REQUEST_ENTITY_IDS';
export const RECEIVE_ENTITY_IDS = 'RECEIVE_ENTITY_IDS';

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
          openNotification('error', 'There was an error retrieving permissions list', null);
          reject(error);
          return;
        }
        dispatch(receiveEntityIds(res));
        resolve(res);
      });
    });

  };
}
