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
