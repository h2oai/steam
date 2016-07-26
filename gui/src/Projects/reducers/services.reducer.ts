/**
 * Created by justin on 7/22/16.
 */
import * as _ from 'lodash';
import { RECEIVE_SERVICES } from '../actions/services.actions';

let initialState = {
  runningServices: []
};

export const servicesReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case RECEIVE_SERVICES:
      return _.assign({}, state, {
        runningServices: action.services
      });
    default:
      return state;
  }
};
