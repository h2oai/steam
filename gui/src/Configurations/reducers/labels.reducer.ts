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
 * Created by Jeff Fohl <jfohl@h2o.ai> on 08/01/16.
 */
import * as _ from 'lodash';
import { RECEIVE_LABELS } from '../actions/configuration.labels.action';
import { RECEIVE_CREATE_LABEL } from '../actions/configuration.labels.action';
import { RECEIVE_UPDATE_LABEL } from '../actions/configuration.labels.action';

let initialState = {};

export const labelsReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case RECEIVE_LABELS:
      return receiveLabels(state, action);
    case RECEIVE_CREATE_LABEL:
      return receiveCreateLabel(state, action);
    case RECEIVE_UPDATE_LABEL:
      return receiveUpdateLabel(state, action);
    default:
      return state;
  }
};

function receiveLabels(state, action) {
  let projectLabels = {};
  projectLabels[action.projectId] = action.labels;
  return _.assign({}, state, projectLabels);
}

function receiveCreateLabel(state, action) {
  let labels;
  if (state[action.projectId]) {
    labels = state[action.projectId].slice();
    labels.push(action.label);
  } else {
    labels = [action.label];
  }
  let projectLabels = {};
  projectLabels[action.projectId] = labels;
  return _.assign({}, state, projectLabels);
}

function receiveUpdateLabel(state, action) {
  return state;
}
