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
import { Permission, Role, Identity, Workgroup, LdapConfig } from "../../Proxy/Proxy";
import { getConfig } from '../../Clusters/actions/clusters.actions';
import { LdapGroup } from "../../Proxy/Proxy";

export const REQUEST_SAVE_LOCAL_KERBEROS = 'REQUEST_SAVE_GLOBAL_KERBEROS';
export const RECEIVE_SAVE_LOCAL_KERBEROS = 'RECEIVE_SAVE_GLOBAL_KERBEROS';

export function requestSaveLocalKerberos() {
  return (dispatch) => {
    dispatch({
      type: REQUEST_SAVE_LOCAL_KERBEROS
    });
  };
};
export function receiveSaveLocalKerberos() {
  return (dispatch) => {
    dispatch({
      type: RECEIVE_SAVE_LOCAL_KERBEROS
    });
  };
};

export function saveLocalKerberos(file) {
  return (dispatch, getState) => {
    dispatch(requestSaveLocalKerberos());
    dispatch(openNotification(NotificationType.Info, "Update", 'Uploading keytab...', null));
    let data = new FormData();
    data.append('file', file.files[0]);
    fetch(`/upload?type=keytab`, {
      credentials: 'include',
      method: 'post',
      body: data
    }).then((response) => {
      if (response.status === 200) {
        dispatch(openNotification(NotificationType.Confirm, "Success", 'Keytab uploaded', null));
        dispatch(receiveSaveLocalKerberos());
      } else {
        dispatch(openNotification(NotificationType.Error, "Error", response.statusText, null));
      }
    });

  };
}
