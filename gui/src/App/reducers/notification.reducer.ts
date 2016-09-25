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

/**
 * Created by justin on 8/8/16.
 */
import * as _ from 'lodash';
import { OPEN_NOTIFICATION, CLOSE_NOTIFICATION } from '../actions/notification.actions';

const initialState = {
  isOpen: false,
  notificationType: null,
  text: ''
};

export function notificationReducer(state = initialState, action) {
  switch (action.type) {
    case OPEN_NOTIFICATION:
      return _.assign({}, state, action);
    case CLOSE_NOTIFICATION:
      return _.assign({}, state, initialState);
    default:
      return state;
  }
}
