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
"use strict";
var _ = require('lodash');
var users_actions_1 = require('../actions/users.actions');
var initialState = {
    permissionWithRoles: [],
    roles: [],
    projects: []
};
exports.usersReducer = function (state, action) {
    if (state === void 0) { state = initialState; }
    switch (action.type) {
        case users_actions_1.FILTER_SELECTIONS_CHANGED:
            var index = _.findIndex(state.selectedRoles, function (o) {
                if (o.id === action.id) {
                    return true;
                }
                else {
                    return false;
                }
            });
            if (index === -1)
                console.log("ERROR: unable to find role");
            var newSelectedRoles = _.cloneDeep(state.selectedRoles);
            newSelectedRoles[index].selected = action.selected;
            return _.assign({}, state, {
                selectedRoles: newSelectedRoles
            });
        case users_actions_1.RECEIVE_USERS_WITH_ROLES_AND_PROJECTS:
            return _.assign({}, state, {
                usersWithRolesAndProjects: action.usersWithRolesAndProjects
            });
        case users_actions_1.RECEIVE_PERMISSIONS_WITH_ROLES:
            return _.assign({}, state, {
                permissionsWithRoles: action.permissions
            });
        case users_actions_1.RECEIVE_ROLE_NAMES:
            var roles = action.roleNames;
            roles.sort(function (a, b) {
                if (a.id < b.id)
                    return -1;
                else
                    return 1;
            });
            var toAppend = { roles: roles };
            if (!state.selectedRoles) {
                var selectedRoles = new Array(roles.length);
                for (var i = 0; i < selectedRoles.length; i++) {
                    selectedRoles[i] = {
                        id: roles[i].id,
                        selected: true
                    };
                }
                toAppend.selectedRoles = selectedRoles;
            }
            return _.assign({}, state, toAppend);
        case users_actions_1.RECEIVE_PROJECTS:
            return _.assign({}, state, {
                projects: action.projects
            });
        case users_actions_1.RECEIVE_USERS:
            return _.assign({}, state, {
                users: action.users
            });
        default:
            return state;
    }
};
//# sourceMappingURL=users.reducer.js.map