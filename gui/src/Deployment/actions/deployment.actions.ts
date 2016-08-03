/**
 * Created by justin on 8/2/16.
 */
import * as Remote from '../../Proxy/Proxy';


export const UPLOADING_PACKAGE = 'UPLOADING_PACKAGE';
export const FINISH_UPLOADING_PACKAGE = 'FINISH_UPLOADING_PACKAGE';
export const RECEIVE_PACKAGES = 'RECEIVE_PACKAGES';

export function uploadingPackage() {
  return {
    type: UPLOADING_PACKAGE
  };
}

export function finishUploadingPackage() {
  return {
    type: FINISH_UPLOADING_PACKAGE
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
    let formFiles: NodeListOf<HTMLInputElement> = form.querySelectorAll('input[type="file"');
    console.log(formFiles);
    for (let i = 0; i < formFiles.length; i++) {
      let data = new FormData();
      for (let j = 0; j < formFiles[i].files.length; j++) {
        data.append('file', formFiles[i].files[j]);
        Remote.createPackage(projectId, packageName, (error) => {
          fetch(`/upload?type=file&project-id=${projectId}&package-name=${packageName}&relative-path=`, {
            credentials: 'include',
            method: 'post',
            body: data
          }).then(() => {
            dispatch(finishUploadingPackage());
          });
        });
      }
    }
  };
}

export function fetchPackages(projectId: number) {
  return (dispatch) => {
    Remote.getPackages(projectId, (error, res) => {
      dispatch(receivePackages(res));
    });
  };
}
