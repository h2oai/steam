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
 * Created by justin on 9/15/16.
 */
import * as _ from 'lodash';
import { RECEIVE_ENGINES, FETCH_CONFIG_COMPLETED } from '../actions/clusters.actions';


let initialState = {
  engines: [],
  config: null
};

export const clustersReducer = (state = initialState, action) => {
  switch (action.type) {
    case RECEIVE_ENGINES:
      return _.assign({}, state, {
        engines: action.engines
      });
    case FETCH_CONFIG_COMPLETED:
      return _.assign({}, state, {
        config: action.config
      });
    default:
      return state;
  }
};
