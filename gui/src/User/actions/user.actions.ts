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
import {Keytab} from "../../Proxy/Proxy";
import {fetchGlobalKeytab} from "../../Users/actions/users.actions";

export const REQUEST_USER_KEYTAB = 'REQUEST_USER_KEYTAB';
export const RECEIVE_USER_KEYTAB = 'RECEIVE_USER_KEYTAB';
export const REQUEST_TEST_KEYTAB = 'REQUEST_TEST_KEYTAB';
export const RECEIVE_TEST_KEYTAB = 'RECEIVE_TEST_KEYTAB';
export const REQUEST_DELETE_KEYTAB = 'REQUEST_DELETE_KEYTAB';
export const RECEIVE_DELETE_KEYTAB = 'RECEIVE_DELETE_KEYTAB';
export const REQUEST_SAVE_LOCAL_KERBEROS = 'REQUEST_SAVE_GLOBAL_KERBEROS';
export const RECEIVE_SAVE_LOCAL_KERBEROS = 'RECEIVE_SAVE_GLOBAL_KERBEROS';

export function requestUserKeytab() {
  return (dispatch) => {
    dispatch({
      type: REQUEST_USER_KEYTAB
    });
  };
};
export function receiveUserKeytab(keytab: Keytab, exists: boolean) {
  return (dispatch) => {
    dispatch({
      type: RECEIVE_USER_KEYTAB,
      keytab,
      exists
    });
  };
};
export function requestTestKeytab() {
  return (dispatch) => {
    dispatch({
      type: REQUEST_TEST_KEYTAB
    });
  };
};
export function receiveTestKeytab() {
  return (dispatch) => {
    dispatch({
      type: RECEIVE_TEST_KEYTAB
    });
  };
};
export function requestDeleteKeytab() {
  return (dispatch) => {
    dispatch({
      type: REQUEST_DELETE_KEYTAB
    });
  };
};
export function receiveDeleteKeytab() {
  return (dispatch) => {
    dispatch({
      type: RECEIVE_DELETE_KEYTAB
    });
  };
};
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

export function fetchUserKeytab() {
  return (dispatch) => {
    dispatch(requestUserKeytab());
    Remote.getUserKeytab((error: Error, keytab: Keytab, exists: boolean) => {
      dispatch(receiveUserKeytab(keytab, exists));
    });
  };
};

export function deleteKeytab(keytabId: number, global: boolean) {
  return (dispatch, getState) => {
    dispatch(requestDeleteKeytab());
    Remote.deleteKeytab(keytabId, (error: Error) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, "Error", error.toString(), null));
        return;
      }
      dispatch(openNotification(NotificationType.Info, "Update", "Keytab Deleted", null));
      dispatch(receiveDeleteKeytab());
      if (global) {
        fetchGlobalKeytab()(dispatch);
      } else {
        fetchUserKeytab()(dispatch);
      }
    });
  };
}

export function testKeytab(keytabId: number) {
  return (dispatch, getstate) => {
    dispatch(requestTestKeytab());
    Remote.testKeytab(keytabId, (error: Error) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, "Error", error.toString(), null));
        return;
      }
      dispatch(openNotification(NotificationType.Confirm, "Update", "Keytab Valid", null));
      dispatch(receiveTestKeytab());
    });
  };
}

export function saveLocalKerberos(file) {
  return (dispatch, getState) => {
    dispatch(requestSaveLocalKerberos());
    dispatch(openNotification(NotificationType.Info, "Update", 'Uploading keytab...', null));
    let data = new FormData();
    data.append('file', file.files[0]);
    fetch(`/upload?type=keytab&principal=user`, {
      credentials: 'include',
      method: 'post',
      body: data
    }).then((response) => {
      if (response.status === 200) {
        dispatch(openNotification(NotificationType.Confirm, "Success", 'User keytab uploaded', null));
        dispatch(receiveSaveLocalKerberos());
        fetchUserKeytab()(dispatch);
      } else {
        dispatch(openNotification(NotificationType.Error, "Error", response.statusText, null));
      }
    });

  };
}
