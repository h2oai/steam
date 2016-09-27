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

/**
 * Created by justin on 8/2/16.
 */
import * as Remote from '../../Proxy/Proxy';
import * as _ from 'lodash';
import { openNotification } from '../../App/actions/notification.actions';
import { NotificationType } from '../../App/components/Notification';

export const UPLOADING_PACKAGE = 'UPLOADING_PACKAGE_COMPONENT';
export const FINISH_UPLOADING_PACKAGE_COMPONENT = 'FINISH_UPLOADING_PACKAGE_COMPONENT';
export const RECEIVE_PACKAGES = 'RECEIVE_PACKAGES';

export function uploadingPackage() {
  return {
    type: UPLOADING_PACKAGE
  };
}

export function finishUploadingPackageComponent() {
  return {
    type: FINISH_UPLOADING_PACKAGE_COMPONENT
  };
}

export function receivePackages(packages) {
  return {
    type: RECEIVE_PACKAGES,
    packages
  };
}

export function uploadPackage(projectId: number, packageName: string, form) {
  return (dispatch) => {
    dispatch(uploadingPackage());
    let formFiles: NodeListOf<HTMLInputElement> = form.querySelectorAll('input[type="file"]');
    Remote.createPackage(projectId, packageName, (error) => {
      let data;
      let requests = [];
      let main = null;
      for (let i = 0; i < formFiles.length; i++) {
        for (let j = 0; j < formFiles[i].files.length; j++) {
          data = new FormData();
          if (formFiles[i].name === 'selectMain') {
            main = formFiles[i].files[j].name;
          }
          data.append('file', formFiles[i].files[j]);
          if (error) {
            dispatch(openNotification(NotificationType.Error, "Load Error", error.toString(), null));
            return;
          }

          requests.push(fetch(`/upload?type=file&project-id=${projectId}&package-name=${packageName}&relative-path=`, {
            credentials: 'include',
            method: 'post',
            body: data
          }).then(() => {
            Remote.setAttributesForPackage(projectId, packageName, JSON.stringify({main: formFiles[i].files[j].name}), (error) => {
              if (error) {
                dispatch(openNotification(NotificationType.Error, "Load Error", error, null));
                return;
              }
            });
            dispatch(finishUploadingPackageComponent());
            dispatch(fetchPackages(projectId));
          }));
        }
      }
      Promise.all(requests).then(() => {
        Remote.setAttributesForPackage(projectId, packageName, JSON.stringify({main: main}), (error) => {
          if (error) {
            dispatch(openNotification(NotificationType.Error, 'Load Error', error, null));
            return;
          }
        });
        dispatch(finishUploadingPackageComponent());
        dispatch(fetchPackages(projectId));
      });
    });
  };
}

export function fetchPackages(projectId: number) {
  return (dispatch) => {
    Remote.getPackages(projectId, (error, res) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
        return;
      }
      dispatch(receivePackages(res));
    });
  };
}
