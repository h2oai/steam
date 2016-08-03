/**
 * Created by justin on 8/2/16.
 */
import { RECEIVE_PACKAGES } from '../actions/deployment.actions';
let initialState = {
  packages: []
};

export const deploymentReducer = (state = initialState, action) => {
  switch(action.type) {
    case RECEIVE_PACKAGES:
      return _.assign({}, state, {
        packages: action.packages
      });
    default:
      return state;
  }
};
