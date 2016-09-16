/**
 * Created by justin on 9/15/16.
 */
import * as _ from 'lodash';
import { RECEIVE_ENGINES } from '../actions/clusters.actions';


let initialState = {
  engines: []
};

export const clustersReducer = (state = initialState, action) => {
  switch (action.type) {
    case RECEIVE_ENGINES:
      console.log(action);
      return _.assign({}, state, {
        engines: action.engines
      });
    default:
      return state;
  }
};
