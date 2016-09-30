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

import * as Remote from '../../Proxy/Proxy';
import * as _ from 'lodash';
import { openNotification } from '../../App/actions/notification.actions';
import { NotificationType } from '../../App/components/Notification';
import { Permission } from "../../Proxy/Proxy";
import { Role } from "../../Proxy/Proxy";

export const FILTER_SELECTIONS_CHANGED = 'FILTER_SELECTIONS_CHANGED';
export const REQUEST_PERMISSIONS_WITH_ROLES = 'REQUEST_PERMISSIONS_WITH_ROLES';
export const RECEIVE_PERMISSIONS_WITH_ROLES = 'RECEIVE_PERMISSIONS_WITH_ROLES';
export const REQUEST_USERS_WITH_ROLES_AND_PROJECTS = "REQUEST_USERS_WITH_ROLES_AND_PROJECTS";
export const RECEIVE_USERS_WITH_ROLES_AND_PROJECTS = "RECEIVE_USERS_WITH_ROLES_AND_PROJECTS";
export const REQUEST_USERS = 'REQUEST_USERS';
export const RECEIVE_USERS = 'RECEIVE_USERS';
export const REQUEST_ROLE_NAMES = 'REQUEST_ROLE_NAMES';
export const RECEIVE_ROLE_NAMES = 'RECEIVE_ROLE_NAMES';
export const REQUEST_PROJECTS = 'REQUEST_PROJECTS';
export const RECEIVE_PROJECTS = 'RECEIVE_PROJECTS';
export const REQUEST_SAVE_PERMISSIONS = 'REQUEST_SAVE_PERMISSIONS';
export const RECEIVE_SAVE_PERMISSIONS = 'RECEIVE_SAVE_PERMISSIONS';

export function requestSavePermissions() {
  return {
    type: REQUEST_SAVE_PERMISSIONS
  };
}
export function receiveSavePermissions() {
  return {
    type: RECEIVE_SAVE_PERMISSIONS
  };
}

export function filterSelectionsChanged(id, selected) {
  return {
    type : FILTER_SELECTIONS_CHANGED,
    id,
    selected
  };
}

export const requestPermissionsByRole = () => {
  return {
    type: REQUEST_PERMISSIONS_WITH_ROLES
  };
};

export function receivePermissionsByRole(permissions: Array<Permission>) {
  return {
    type: RECEIVE_PERMISSIONS_WITH_ROLES,
    permissions
  };
}
export const requestUsersWithRolesAndProjects = () => {
  return {
    type: REQUEST_USERS_WITH_ROLES_AND_PROJECTS
  };
};

export function receiveUsersWithRolesAndProjects(usersWithRolesAndProjects) {
  return {
    type: RECEIVE_USERS_WITH_ROLES_AND_PROJECTS,
    usersWithRolesAndProjects
  };
}

export const requestUsers = () => {
  return {
    type: REQUEST_USERS
  };
};
export function receiveUsers(users) {
  return {
    type: RECEIVE_USERS,
    users
  };
};
export const requestProjects = () => {
  return {
    type: REQUEST_PROJECTS
  };
};

export function receiveProjects(projects) {
  return {
    type: RECEIVE_PROJECTS,
    projects
  };
}

export const requestRoleNames = () => {
  return {
    type: REQUEST_ROLE_NAMES
  };
};

export function receiveRoleNames(roleNames) {
  return {
    type: RECEIVE_ROLE_NAMES,
    roleNames
  };
}

function getProjects(dispatch): Promise<Array<any>> {
  return new Promise((resolve, reject) => {
      dispatch(requestProjects());
      Remote.getProjects(0, 1000, (error, res: any) => {
        if (error) {
          openNotification(NotificationType.Error, 'Load Error', 'There was an error retrieving projects', null);
          reject();
        }
        dispatch(receiveProjects(res));
        resolve(res);
      });
    }
  );
}

function getUsers(dispatch): Promise<Array<Role>> {
  return new Promise((resolve, reject) => {
      dispatch(requestUsers());
      Remote.getIdentities(0, 1000, (error, res: any) => {
        if (error) {
          openNotification(NotificationType.Error, 'Load Error', 'There was an error retrieving users', null);
          reject();
        }
        dispatch(receiveUsers(res));
        resolve(res);
      });
    }
  );
}

/***
 * @returns {Promise<T>|Promise} [{ created:number, description:String, id:number, name:String }]
 */
function getRoles(dispatch): Promise<Array<Role>> {
  return new Promise((resolve, reject) => {
      dispatch(requestRoleNames());
      Remote.getRoles(0, 1000, (error, res: any) => {
        if (error) {
          openNotification(NotificationType.Error, 'Load Error', 'There was an error retrieving roles', null);
          reject();
        }
        dispatch(receiveRoleNames(res));
        resolve(res);
      });
    }
  );
}

/***
 * @returns {Promise<T>|Promise}  { code:String, description:String, id:number }
 */
function getPermissionDescriptions(): Promise<Array<Permission>> {
  return new Promise((resolve, reject) => {
    Remote.getAllPermissions((error, res) => {
      if (error) {
        openNotification(NotificationType.Error, 'Load Error', 'There was an error retrieving permissions list', null);
        reject(error);
      }
      resolve(res);
    });
  });
}

export function changeFilterSelections(id, selected) {
  return (dispatch, getState) => {
    dispatch(filterSelectionsChanged(id, selected));
  };
}

export function fetchUsersWithRolesAndProjects() {
  return (dispatch, getState) => {
    dispatch(requestUsersWithRolesAndProjects());
    const projectsPromise = getProjects(dispatch);
    const rolesPromise = getRoles(dispatch);
    const usersPromise = getUsers(dispatch);

    let roleForIdentityPromises = [];

    usersPromise.then((users) => {
      for (let user of users) {
        roleForIdentityPromises.push(new Promise((resolve, reject) => {
          Remote.getRolesForIdentity(user.id, (error, res) => {
            if (error) {
              reject(error);
            }
            resolve({
              user,
              roles: res
            });
          });
        }));
      }

      Promise.all(roleForIdentityPromises).then((values) => {
        dispatch(receiveUsersWithRolesAndProjects(values));
      });
    });
  };
}

export interface PermissionsWithRoles {
  description: string
  flags: Array<boolean>
}
/***
 * @param roles the numeric roleID's to be returned in the permission set
 */
export function fetchPermissionsWithRoles() {
  return (dispatch) => {
    dispatch(requestPermissionsByRole());

    const descriptionsPromise = getPermissionDescriptions();

    getRoles(dispatch).then((roles) => {
      let permissionRequests: Array<Promise<any>> = [];
      for (let role of roles) {
        permissionRequests.push(new Promise(
          (resolve, reject) => Remote.getPermissionsForRole(role.id, (error, res) => {
            if (error) {
              openNotification(NotificationType.Error, 'Load Error', 'There was an error retrieving permissions list', null);
              reject();
            }
            resolve({
              roleId : role.id,
              permissions : res
            });
          })
        ));
      }
      Promise.all(permissionRequests).then((permissionsByRole) => {
        let output = [];
        permissionsByRole.sort((a, b) => {
          if (a.roleId < b.roleId) return -1; else return 1;
        });

        descriptionsPromise.then((descriptions) => {
            let flags;
            let permissionSet;

            for (let i = 0; i < descriptions.length; i++) {
              flags = [];
              for (let j = 0; j < permissionsByRole.length; j++) {
                permissionSet = permissionsByRole[j];
                if (_.findIndex(permissionSet.permissions, (o: Permission) => {
                  return o.id === descriptions[i].id;
                }) !== -1) {
                  flags.push(true);
                } else {
                  flags.push(false);
                }
              }
              output.push({
                description : descriptions[i].description,
                id: descriptions[i].id,
                flags
              });
            }

            dispatch(receivePermissionsByRole(output));
        });
      });
    });
  };
}

export function saveUpdatedPermissions(permissionInputs) {
  return (dispatch) => {
    dispatch(requestSavePermissions());

    let updates = [];

    for (var permissionKey in permissionInputs) {
      for (var flagKey in permissionInputs[permissionKey].flags) {
        let flagset = permissionInputs[permissionKey].flags[flagKey];
        if (flagset.originalFlag !== flagset.input.checked) {
          updates.push({
            newFlag: flagset.input.checked,
            userIndex: flagKey,
            permissionIndex: permissionKey,
            description: permissionInputs[permissionKey].permissionSet.description,
            permissionId: permissionInputs[permissionKey].permissionSet.id
          });
        }
      }
    }

    for (let update of updates) {
      if(update.newFlag === true) {
        //Remote.linkRoleWithPermission(roleId, permissionId, (error) => console.log(error));
      } else {
        //Remote.unlinkRoleFromPermission(roleId, permissionId, (error) => console.log(error));
      }
    }
    console.log(updates);
  };
}
