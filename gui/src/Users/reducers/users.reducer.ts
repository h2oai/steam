import * as _ from 'lodash';
import {
  RECEIVE_PERMISSIONS_BY_ROLE, RECEIVE_ROLE_NAMES, RECEIVE_PROJECTS
} from '../actions/users.actions';

let initialState = {
  permissionsByRole: [],
};

export const usersReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case RECEIVE_PERMISSIONS_BY_ROLE:
      return _.assign({}, state, {
        permissionsByRole: action.permissions
      });
    case RECEIVE_ROLE_NAMES:
      let roles = action.roleNames;
      roles.sort((a,b) => {
        if(a.id < b.id) return -1; else return 1;
      });
      return _.assign({}, state, {
        roles: action.roleNames
      });
    case RECEIVE_PROJECTS :
      console.log(action.projects);
      return _.assign({}, state, {
        projects: action.projects
      });
    default:
      return state;
  }
};
