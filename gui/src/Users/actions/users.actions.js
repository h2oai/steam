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
var Remote = require('../../Proxy/Proxy');
var _ = require('lodash');
var notification_actions_1 = require('../../App/actions/notification.actions');
var Notification_1 = require('../../App/components/Notification');
exports.FILTER_SELECTIONS_CHANGED = 'FILTER_SELECTIONS_CHANGED';
exports.REQUEST_PERMISSIONS_WITH_ROLES = 'REQUEST_PERMISSIONS_WITH_ROLES';
exports.RECEIVE_PERMISSIONS_WITH_ROLES = 'RECEIVE_PERMISSIONS_WITH_ROLES';
exports.REQUEST_USERS_WITH_ROLES_AND_PROJECTS = "REQUEST_USERS_WITH_ROLES_AND_PROJECTS";
exports.RECEIVE_USERS_WITH_ROLES_AND_PROJECTS = "RECEIVE_USERS_WITH_ROLES_AND_PROJECTS";
exports.REQUEST_USERS = 'REQUEST_USERS';
exports.RECEIVE_USERS = 'RECEIVE_USERS';
exports.REQUEST_ROLE_NAMES = 'REQUEST_ROLE_NAMES';
exports.RECEIVE_ROLE_NAMES = 'RECEIVE_ROLE_NAMES';
exports.REQUEST_PROJECTS = 'REQUEST_PROJECTS';
exports.RECEIVE_PROJECTS = 'RECEIVE_PROJECTS';
function filterSelectionsChanged(id, selected) {
    return {
        type: exports.FILTER_SELECTIONS_CHANGED,
        id: id,
        selected: selected
    };
}
exports.filterSelectionsChanged = filterSelectionsChanged;
exports.requestPermissionsByRole = function () {
    return {
        type: exports.REQUEST_PERMISSIONS_WITH_ROLES
    };
};
function receivePermissionsByRole(permissions) {
    return {
        type: exports.RECEIVE_PERMISSIONS_WITH_ROLES,
        permissions: permissions
    };
}
exports.receivePermissionsByRole = receivePermissionsByRole;
exports.requestUsersWithRolesAndProjects = function () {
    return {
        type: exports.REQUEST_USERS_WITH_ROLES_AND_PROJECTS
    };
};
function receiveUsersWithRolesAndProjects(usersWithRolesAndProjects) {
    return {
        type: exports.RECEIVE_USERS_WITH_ROLES_AND_PROJECTS,
        usersWithRolesAndProjects: usersWithRolesAndProjects
    };
}
exports.receiveUsersWithRolesAndProjects = receiveUsersWithRolesAndProjects;
exports.requestUsers = function () {
    return {
        type: exports.REQUEST_USERS
    };
};
function receiveUsers(users) {
    return {
        type: exports.RECEIVE_USERS,
        users: users
    };
}
exports.receiveUsers = receiveUsers;
;
exports.requestProjects = function () {
    return {
        type: exports.REQUEST_PROJECTS
    };
};
function receiveProjects(projects) {
    return {
        type: exports.RECEIVE_PROJECTS,
        projects: projects
    };
}
exports.receiveProjects = receiveProjects;
exports.requestRoleNames = function () {
    return {
        type: exports.REQUEST_ROLE_NAMES
    };
};
function receiveRoleNames(roleNames) {
    return {
        type: exports.RECEIVE_ROLE_NAMES,
        roleNames: roleNames
    };
}
exports.receiveRoleNames = receiveRoleNames;
function getProjects(dispatch) {
    return new Promise(function (resolve, reject) {
        dispatch(exports.requestProjects());
        Remote.getProjects(0, 1000, function (error, res) {
            if (error) {
                notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', 'There was an error retrieving projects', null);
                reject();
            }
            dispatch(receiveProjects(res));
            resolve(res);
        });
    });
}
function getUsers(dispatch) {
    return new Promise(function (resolve, reject) {
        dispatch(exports.requestUsers());
        Remote.getIdentities(0, 1000, function (error, res) {
            if (error) {
                notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', 'There was an error retrieving users', null);
                reject();
            }
            dispatch(receiveUsers(res));
            resolve(res);
        });
    });
}
/***
 * @returns {Promise<T>|Promise} [{ created:number, description:String, id:number, name:String }]
 */
function getRoles(dispatch) {
    return new Promise(function (resolve, reject) {
        dispatch(exports.requestRoleNames());
        Remote.getRoles(0, 1000, function (error, res) {
            if (error) {
                notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', 'There was an error retrieving roles', null);
                reject();
            }
            dispatch(receiveRoleNames(res));
            resolve(res);
        });
    });
}
/***
 * @returns {Promise<T>|Promise}  { code:String, description:String, id:number }
 */
function getPermissionDescriptions() {
    return new Promise(function (resolve, reject) {
        Remote.getAllPermissions(function (error, res) {
            if (error) {
                notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', 'There was an error retrieving permissions list', null);
                reject(error);
            }
            resolve(res);
        });
    });
}
function changeFilterSelections(id, selected) {
    return function (dispatch, getState) {
        dispatch(filterSelectionsChanged(id, selected));
    };
}
exports.changeFilterSelections = changeFilterSelections;
function fetchUsersWithRolesAndProjects() {
    return function (dispatch, getState) {
        dispatch(exports.requestUsersWithRolesAndProjects());
        var projectsPromise = getProjects(dispatch);
        var rolesPromise = getRoles(dispatch);
        var usersPromise = getUsers(dispatch);
        var roleForIdentityPromises = [];
        usersPromise.then(function (users) {
            var _loop_1 = function(user) {
                roleForIdentityPromises.push(new Promise(function (resolve, reject) {
                    Remote.getRolesForIdentity(user.id, function (error, res) {
                        if (error) {
                            reject(error);
                        }
                        resolve({
                            user: user,
                            roles: res
                        });
                    });
                }));
            };
            for (var _i = 0, users_1 = users; _i < users_1.length; _i++) {
                var user = users_1[_i];
                _loop_1(user);
            }
            Promise.all(roleForIdentityPromises).then(function (values) {
                dispatch(receiveUsersWithRolesAndProjects(values));
            });
        });
    };
}
exports.fetchUsersWithRolesAndProjects = fetchUsersWithRolesAndProjects;
/***
 * @param roles the numeric roleID's to be returned in the permission set
 */
function fetchPermissionsWithRoles() {
    return function (dispatch) {
        dispatch(exports.requestPermissionsByRole());
        var descriptionsPromise = getPermissionDescriptions();
        getRoles(dispatch).then(function (roles) {
            var permissionRequests = [];
            var _loop_2 = function(role) {
                permissionRequests.push(new Promise(function (resolve, reject) { return Remote.getPermissionsForRole(role.id, function (error, res) {
                    if (error) {
                        notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', 'There was an error retrieving permissions list', null);
                        reject();
                    }
                    resolve({
                        roleId: role.id,
                        permissions: res
                    });
                }); }));
            };
            for (var _i = 0, roles_1 = roles; _i < roles_1.length; _i++) {
                var role = roles_1[_i];
                _loop_2(role);
            }
            Promise.all(permissionRequests).then(function (permissionsByRole) {
                var output = [];
                permissionsByRole.sort(function (a, b) {
                    if (a.roleId < b.roleId)
                        return -1;
                    else
                        return 1;
                });
                descriptionsPromise.then(function (descriptions) {
                    var flags;
                    var permissionSet;
                    var _loop_3 = function(i) {
                        flags = [];
                        for (var j = 0; j < permissionsByRole.length; j++) {
                            permissionSet = permissionsByRole[j];
                            if (_.findIndex(permissionSet.permissions, function (o) {
                                return o.id === descriptions[i].id;
                            }) !== -1) {
                                flags.push(true);
                            }
                            else {
                                flags.push(false);
                            }
                        }
                        output.push({
                            description: descriptions[i].description,
                            flags: flags
                        });
                    };
                    for (var i = 0; i < descriptions.length; i++) {
                        _loop_3(i);
                    }
                    dispatch(receivePermissionsByRole(output));
                });
            });
        });
    };
}
exports.fetchPermissionsWithRoles = fetchPermissionsWithRoles;
//# sourceMappingURL=users.actions.js.map