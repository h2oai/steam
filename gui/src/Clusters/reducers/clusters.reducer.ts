/**
 * Created by justin on 9/15/16.
 */
import * as _ from 'lodash';
import { RECEIVE_ENGINES, RECEIVE_CONFIG } from '../actions/clusters.actions';


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
    case RECEIVE_CONFIG:
      return _.assign({}, state, {
        config: action.config
      });
    default:
      return state;
  }
};
