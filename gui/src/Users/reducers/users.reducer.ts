/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

import * as _ from 'lodash';
import {
  RECEIVE_PERMISSIONS_WITH_ROLES, RECEIVE_ROLE_NAMES, RECEIVE_PROJECTS, RECEIVE_USERS, RECEIVE_SAVE_PERMISSIONS,
  RESET_UPDATES, RECEIVE_USERS_WITH_ROLES_AND_PROJECTS, FILTER_SELECTIONS_CHANGED, ENTER_NEW_ROLE, EXIT_NEW_ROLE,
  ENTER_NEW_USER,
  EXIT_NEW_USER, RECEIVE_CREATE_ROLE, RECEIVE_WORKGROUPS_FOR_IDENTITY, RECEIVE_LDAP_CONFIG, RECEIVE_ADMIN_CHECK
} from '../actions/users.actions';

let initialState = {
  permissionsWithRoles: [],
  roles: [],
  projects: [],
  updates: [],
  createNewUserIsEntered: false,
  createNewRoleIsEntered: false,
  ldapExists: false,
  ldapConfig: {
    host: "",
    port: 636,
    ldaps: true,
    bind_dn: "dc=xyz,dc=com",
    bind_password: "",
    user_base_dn: "",
    user_base_filter: "",
    group_dn: "dc=xyz,dc=com",
    static_member_attribute: "memberUid",
    search_request_size_limit: 0,
    search_request_time_limit: 0,
    force_bind: true
  },
  isAdmin: false
};

export const usersReducer = (state: any = initialState, action: any) => {
  switch (action.type) {
    case RECEIVE_ADMIN_CHECK : {
      return _.assign({}, state, {
        isAdmin: action.isAdmin
      });
    }
    case RECEIVE_LDAP_CONFIG : {
      return _.assign({}, state, {
        ldapExists: action.exists,
        ldapConfig: action.config
      });
    }
    case RECEIVE_WORKGROUPS_FOR_IDENTITY : {
      return _.assign({}, state, {
        userWithWorkgroups: {
          id: action.userId,
          workgroups: action.workgroups
        }
      });
    }
    case RECEIVE_CREATE_ROLE : {
      let newState: any = _.assign({}, state);
      newState.selectedRoles = state.selectedRoles.slice(0);
      newState.selectedRoles.push({
        id: action.roleId,
        selected: true
      });
      return newState;
    }
    case ENTER_NEW_USER : {
      return _.assign({}, state, {
        createNewUserIsEntered: true
      });
    }
    case EXIT_NEW_USER : {
      return _.assign({}, state, {
        createNewUserIsEntered: false
      });
    }
    case ENTER_NEW_ROLE : {
      return _.assign({}, state, {
        createNewRoleIsEntered: true
      });
    }
    case EXIT_NEW_ROLE : {
      return _.assign({}, state, {
        createNewRoleIsEntered: false
      });
    }
    case RESET_UPDATES : {
      return _.assign({}, state, {
        updates: []
      });
    }
    case FILTER_SELECTIONS_CHANGED : {
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
        selectedRoles: newSelectedRoles
      });
    }
    case RECEIVE_USERS_WITH_ROLES_AND_PROJECTS : {
      return _.assign({}, state, {
        usersWithRolesAndProjects: action.usersWithRolesAndProjects
      });
    }
    case RECEIVE_PERMISSIONS_WITH_ROLES : {
      return _.assign({}, state, {
        permissionsWithRoles: action.permissions
      });
    }
    case RECEIVE_ROLE_NAMES : {
      let roles = action.roleNames;
      roles.sort((a, b) => {
        if (a.id < b.id) return -1; else return 1;
      });

      let toAppend: any = {roles};
      if (!(state as any).selectedRoles) {
        let selectedRoles = new Array(roles.length + 1);

        selectedRoles[0] = {
          id: -1,
          selected: true
        };

        for (let i = 1; i < selectedRoles.length; i++) {
          selectedRoles[i] = {
            id: roles[i - 1].id,
            selected: true
          };
        }

        toAppend.selectedRoles = selectedRoles;
      }
      return _.assign({}, state, toAppend);
    }
    case RECEIVE_PROJECTS : {
      return _.assign({}, state, {
        projects: action.projects
      });
    }
    case RECEIVE_USERS : {
      return _.assign({}, state, {
        users: action.users
      });
    }
    case RECEIVE_SAVE_PERMISSIONS : {
      let newState: any = _.assign({}, state);
      if (action.hasOwnProperty("roleId")) {
        newState.updates = newState.updates.slice(0);
        newState.updates.push({
          roleId: action.roleId,
          permissionId: action.permissionId
        });
      } else if (action.hasOwnProperty("error")) {
        newState.updates = newState.updates.slice(0);
        newState.updates.push({
          error: action.error
        });
      } else {
        console.log("ERROR: invalid update state");
      }
      return newState;
    }
    default: {
      return state;
    }
  }
};
