import * as _ from 'lodash';
import {
  RECEIVE_PERMISSIONS_WITH_ROLES, RECEIVE_ROLE_NAMES, RECEIVE_PROJECTS, RECEIVE_USERS,
  RECEIVE_USERS_WITH_ROLES_AND_PROJECTS, FILTER_SELECTIONS_CHANGED
} from '../actions/users.actions';

let initialState = {
  permissionWithRoles: [],
  roles: [],
  projects: []
};

export const usersReducer = (state: any = initialState, action: any) => {
  switch (action.type) {
    case FILTER_SELECTIONS_CHANGED :
      let index = _.findIndex(state.selectedRoles, (o) => {
        if ((o as any).id === action.id) {
          return true;
        } else {
          return false;
        }
      });
      if (index === -1) console.log("ERROR: unable to find role");

      let newSelectedRoles = _.cloneDeep(state.selectedRoles);
      newSelectedRoles[index].selected = action.selected;

      return _.assign({}, state, {
        selectedRoles : newSelectedRoles
      });
    case RECEIVE_USERS_WITH_ROLES_AND_PROJECTS:
      return _.assign({}, state, {
        usersWithRolesAndProjects: action.usersWithRolesAndProjects
      });
    case RECEIVE_PERMISSIONS_WITH_ROLES:
      return _.assign({}, state, {
        permissionsWithRoles: action.permissions
      });
    case RECEIVE_ROLE_NAMES:
      let roles = action.roleNames;
      roles.sort((a, b) => {
        if (a.id < b.id) return -1; else return 1;
      });

      let toAppend: any = { roles };
      if (!(state as any).selectedRoles) {
        let selectedRoles = new Array(roles.length);

        for (let i = 0; i < selectedRoles.length; i++) {
          selectedRoles[i] = {
            id: roles[i].id,
            selected: true
          };
        }

        toAppend.selectedRoles = selectedRoles;
      }
      return _.assign({}, state, toAppend);
    case RECEIVE_PROJECTS :
      return _.assign({}, state, {
        projects: action.projects
      });
    case RECEIVE_USERS :
      return _.assign({}, state, {
        users: action.users
      });
    default:
      return state;
  }
};
