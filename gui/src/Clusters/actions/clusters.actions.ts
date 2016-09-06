/**
 * Created by justin on 8/17/16.
 */

import * as Remote from '../../Proxy/Proxy';

export function uploadEngine(form) {
  return (dispatch) => {
    let file = form.querySelectorAll('input[type="file"]')[0];
    let data = new FormData();
    data.append('file', file.files[0]);
    let headers = new Headers();
    headers.append('Content-Type', 'text/plain');
    fetch(`/upload?type=engine`, {
      credentials: 'include',
      method: 'post',
      headers,
      body: data
    }).then((res) => {
      Remote.getEngines((error, engines) => {
        console.log(engines);
      });
    });
  };
}

export function startYarnCluster(clusterName, engineId, size, memory, username) {
  return (dispatch) => {
    Remote.startClusterOnYarn(clusterName, engineId, size, memory, username, (error, res) => {

    });
  };
}
