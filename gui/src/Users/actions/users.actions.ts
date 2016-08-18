import * as Remote from '../../Proxy/Proxy';
import * as _ from 'lodash';
import { openNotification } from '../../App/actions/notification.actions';

export const REQUEST_PERMISSIONS_BY_ROLE = 'REQUEST_PERMISSIONS_BY_ROLE';
export const RECEIVE_PERMISSIONS_BY_ROLE = 'RECEIVE_PERMISSIONS_BY_ROLE';
export const REQUEST_USERS_WITH_ROLES_AND_PROJECTS = "REQUEST_USERS_WITH_ROLES_AND_PROJECTS";
export const RECEIVE_USERS_WITH_ROLES_AND_PROJECTS = "RECEIVE_USERS_WITH_ROLES_AND_PROJECTS"
export const REQUEST_ROLE_NAMES = 'REQUEST_ROLE_NAMES';
export const RECEIVE_ROLE_NAMES = 'RECEIVE_ROLE_NAMES';
export const REQUEST_PROJECTS = 'REQUEST_PROJECTS';
export const RECEIVE_PROJECTS = 'RECEIVE_PROJECTS';

export const requestPermissionsByRole = () => {
  return {
    type: REQUEST_PERMISSIONS_BY_ROLE
  };
};

export function receivePermissionsByRole(permissions) {
  return {
    type: RECEIVE_PERMISSIONS_BY_ROLE,
    permissions
  };
}
export const requestUsersWithRolesAndProjects = () => {
  return {
    type: REQUEST_PERMISSIONS_BY_ROLE
  };
};

export function receiveUsersWithRolesAndProjects(usersWithRolesAndProjects) {
  return {
    type: RECEIVE_USERS_WITH_ROLES_AND_PROJECTS,
    usersWithRolesAndProjects
  };
}


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

function getProjects(dispatch):Promise<Array<any>> {
  return new Promise((resolve, reject) => {
      dispatch(requestProjects());
      Remote.getProjects(0,1000, (error, res:any)=> {
        if(error) {
          openNotification('error', 'There was an error retrieving projects', null);
          reject();
        }
        dispatch(receiveProjects(res));
        resolve(res);
      })
    }
  );
}

/***
 * @returns {Promise<T>|Promise} [{ created:number, description:String, id:number, name:String }]
 */
function getRoles(dispatch):Promise<Array<any>> {
  return new Promise((resolve, reject) => {
      dispatch(requestRoleNames());
      Remote.getRoles(0,1000, (error, res:any)=> {
        if(error) {
          openNotification('error', 'There was an error retrieving roles', null);
          reject();
        }
        dispatch(receiveRoleNames(res));
        resolve(res);
      })
    }
  );
}

/***
 * @returns {Promise<T>|Promise}  { code:String, description:String, id:number }
 */
function getPermissionDescriptions():Promise<Array<any>> {
  return new Promise((resolve, reject) => {
    Remote.getAllPermissions((error, res) => {
      if (error) {
        openNotification('error', 'There was an error retrieving permissions list', null);
        reject(error);
      }
      resolve(res);
    })
  });
}

export function fetchUsersWithRolesAndProjects() {
  return (dispatch) => {
    //dispatch(requestUsersWithRolesAndProjects());
    const projectsPromise = getProjects(dispatch);

    dispatch(requestRoleNames());
    const rolesPromise = getRoles(dispatch);
  }
}

/***
 * @param roles the numeric roleID's to be returned in the permission set
 */
export function fetchPermissionsByRole() {
  return (dispatch) => {
    dispatch(requestPermissionsByRole());

    const descriptionsPromise = getPermissionDescriptions();

    getRoles(dispatch).then((roles)=> {
      let permissionRequests: Array<Promise<any>> = [];
      for (let role of roles) {
        permissionRequests.push(new Promise(
          (resolve, reject)=>Remote.getPermissionsForRole(role.id, (error, res) => {
            if (error) {
              openNotification('error', 'There was an error retrieving permissions list', null);
              reject();
            }
            resolve({
              roleId : role.id,
              permissions : res
            });
          })
        ));
      }
      Promise.all(permissionRequests).then((permissionsByRole)=> {
        let output = [];
        permissionsByRole.sort((a,b) => {
          if(a.roleId < b.roleId) return -1; else return 1;
        });

        descriptionsPromise.then((descriptions) => {
            let flags;
            let permissionSet;

            for(let i=0;i < descriptions.length; i++) {
              flags=[];
              for(let j=0; j < permissionsByRole.length; j++) {
                permissionSet = permissionsByRole[j];
                if(_.findIndex(permissionSet.permissions, (o:any)=> {
                  return o.id === descriptions[i].id;
                }) !== -1) {
                  flags.push(true);
                } else {
                  flags.push(false);
                }
              }
              output.push({
                description : descriptions[i].description,
                flags
              });
            }

            dispatch(receivePermissionsByRole(output));
        });
      });
    });
  };
}
