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

import * as _ from 'lodash';
import { RECEIVE_ENTITY_IDS } from '../actions/global.actions';

const initialState = {
  entityIds: {}
};

export function globalReducer(state = initialState, action) {
  switch (action.type) {
    case RECEIVE_ENTITY_IDS:
      let entityIds = {};
      for (let entityType of action.response) {
        entityIds[entityType.name] = entityType.id;
      }
      return _.assign({}, state, { entityIds });
    default:
      return state;
  }
}
