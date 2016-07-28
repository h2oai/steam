/**
 * Created by justin on 7/27/16.
 */
import * as _ from 'lodash';
import { RECEIVE_PROFILE } from '../actions/profile.actions';

let initialState = {
};

export const profileReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case RECEIVE_PROFILE:
      return _.assign({}, state, action.profile);
    default:
      return state;
  }
};
