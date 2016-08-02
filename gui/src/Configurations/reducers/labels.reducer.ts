/**
 * Created by Jeff Fohl <jfohl@h2o.ai> on 08/01/16.
 */
import * as _ from 'lodash';
import { RECEIVE_LABELS } from '../actions/configuration.labels.action';

let initialState = {
  labels: {}
};

export const labelsReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case RECEIVE_LABELS:
      let projectLabels = {};
      projectLabels[action.projectId] = action.labels;
      let newLabels = _.assign(state.labels, projectLabels);
      return _.assign({}, state, {
        labels: newLabels
      });
    default:
      return state;
  }
};
