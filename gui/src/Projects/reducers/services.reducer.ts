/**
 * Created by justin on 7/22/16.
 */
import * as _ from 'lodash';
import { RECEIVE_SERVICES_FOR_PROJECT, RECEIVE_ALL_SERVICES } from '../actions/services.actions';

let initialState = {
  runningServices: []
};

export const servicesReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case RECEIVE_SERVICES_FOR_PROJECT:
      return _.assign({}, state, {
        runningServicesForProject: action.services
      });
    case RECEIVE_ALL_SERVICES:
      return _.assign({}, state, {
        allRunningServices: action.services
      });
    default:
      return state;
  }
};
