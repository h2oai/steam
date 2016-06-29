/**
 * Created by justin on 6/28/16.
 */

import { RECEIVE_MODEL_OVERVIEW } from '../actions/model.overview.action';
import * as _ from 'lodash';

let initialState = {
  basics: [],
  parameters: []
};

export const modelOverviewReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case RECEIVE_MODEL_OVERVIEW:
      return _.assign({}, state, {
        basics: action.modelOverview.basics,
        parameters: action.modelOverview.parameters
      });
    default:
      return state;
  }
};