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
import {
  RECEIVE_ENGINES, FETCH_CONFIG_COMPLETED, START_CLUSTER_COMPLETED,
  START_CLUSTER, UPLOAD_ENGINE_COMPLETED, START_UPLOAD_ENGINE
} from '../actions/clusters.actions';


let initialState = {
  engines: [],
  config: null,
  clusterLaunchIsInProgress: false,
  engineUploading: false
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
    case START_CLUSTER:
      return _.assign({}, state, {
        clusterLaunchIsInProgress: true
      });
    case START_CLUSTER_COMPLETED:
      return _.assign({}, state, {
        clusterLaunchIsInProgress: false
      });
    case START_UPLOAD_ENGINE:
      return _.assign({}, state, {
        engineUploading: true
      });
    case UPLOAD_ENGINE_COMPLETED:
      return _.assign({}, state, {
        engineUploading: false
      });
    default:
      return state;
  }
};
