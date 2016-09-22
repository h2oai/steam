/**
 * Created by justin on 8/2/16.
 */
import * as Remote from '../../Proxy/Proxy';
import * as _ from 'lodash';
import { openNotification } from '../../App/actions/notification.actions';

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
            dispatch(openNotification('error', error.toString(), null));
            return;
          }

          requests.push(fetch(`/upload?type=file&project-id=${projectId}&package-name=${packageName}&relative-path=`, {
            credentials: 'include',
            method: 'post',
            body: data
          }));

        }
      }
      Promise.all(requests).then(() => {
        Remote.setAttributesForPackage(projectId, packageName, JSON.stringify({main: main}), (error) => {
          if (error) {
            dispatch(openNotification('error', error, null));
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
        dispatch(openNotification('error', error.toString(), null));
        return;
      }
      dispatch(receivePackages(res));
    });
  };
}
