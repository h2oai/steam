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
  ENTER_NEW_USER, EXIT_NEW_USER, RECEIVE_CREATE_ROLE, RECEIVE_WORKGROUPS_FOR_IDENTITY, RECEIVE_LDAP_CONFIG,
  RECEIVE_ADMIN_CHECK, RECEIVE_TEST_LDAP, REQUEST_CLEAR_TEST_LDAP
} from '../actions/users.actions';

export const DEFAULT_HOST = "";
export const DEFAULT_PORT = "636";
export const DEFAULT_SSL_ENABLED = true;
export const DEFAULT_BIND_DN = "dc=xyz,dc=com";
export const DEFAULT_BIND_DN_PASSWORD = "";
export const DEFAULT_CONFIRM_PASSWORD = "";
export const DEFAULT_USERBASE_DN = "";
export const DEFAULT_USERBASE_FILTER = "";
export const DEFAULT_USERNAME_ATTRIBUTE = "";
export const DEFAULT_GROUP_DN = "dc=xyz,dc=com";
export const DEFAULT_STATIC_MEMBER_ATTRIBUTE = "memberUid";
export const DEFAULT_SEARCH_REQUEST_SIZE_LIMIT = "";
export const DEFAULT_SEARCH_REQUEST_TIME_LIMIT = "";
export const DEFAULT_GROUP_NAME_ATTRIBUTE = "";
export const DEFAULT_GROUP_NAMES = "";

let initialState = {
  permissionsWithRoles: [],
  roles: [],
  projects: [],
  updates: [],
  createNewUserIsEntered: false,
  createNewRoleIsEntered: false,
  ldapExists: false,
  ldapConfig: {
    host: DEFAULT_HOST,
    port: DEFAULT_PORT,
    ldaps: DEFAULT_SSL_ENABLED,
    bind_dn: DEFAULT_BIND_DN,
    bind_password: DEFAULT_BIND_DN_PASSWORD,
    user_base_dn: DEFAULT_USERBASE_DN,
    user_base_filter: DEFAULT_USERBASE_FILTER,
    usernameAttribute: DEFAULT_USERNAME_ATTRIBUTE,
    group_base_dn: DEFAULT_GROUP_DN,
    group_name_attribute: DEFAULT_GROUP_NAME_ATTRIBUTE,
    static_member_attribute: DEFAULT_STATIC_MEMBER_ATTRIBUTE,
    group_names: DEFAULT_GROUP_NAMES,
    search_request_size_limit: DEFAULT_SEARCH_REQUEST_SIZE_LIMIT,
    search_request_time_limit: DEFAULT_SEARCH_REQUEST_TIME_LIMIT,
    force_bind: true
  },
  isAdmin: false,
  testResult: null
};

export const usersReducer = (state: any = initialState, action: any) => {
  switch (action.type) {
    case REQUEST_CLEAR_TEST_LDAP : {
      return _.assign({}, state, {
          testResult: null
        }
      );
    }
    case RECEIVE_TEST_LDAP : {
      return _.assign({}, state, {
        testResult: action.ldapTestResult
      });
    }
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
