/**
 * Created by justin on 6/28/16.
 */

import { RECEIVE_MODEL_OVERVIEW } from '../actions/model.overview.action';
import * as _ from 'lodash';

let initialState = {
};

export const modelOverviewReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case RECEIVE_MODEL_OVERVIEW:
      return _.assign({}, state, action.model);
    default:
      return state;
  }
};
