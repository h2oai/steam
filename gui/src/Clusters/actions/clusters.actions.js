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
/**
 * Created by justin on 8/17/16.
 */
var Remote = require('../../Proxy/Proxy');
var notification_actions_1 = require('../../App/actions/notification.actions');
var Notification_1 = require('../../App/components/Notification');
exports.RECEIVE_ENGINES = 'RECEIVE_ENGINES';
exports.START_FETCH_CONFIG = 'START_FETCH_CONFIG';
exports.FETCH_CONFIG_COMPLETED = 'FETCH_CONFIG_COMPLETED';
exports.START_UPLOAD_ENGINE = 'START_UPLOAD_ENGINE';
exports.UPLOAD_ENGINE_COMPLETED = 'UPLOAD_ENGINE_COMPLETED';
exports.START_CLUSTER = 'START_CLUSTER';
exports.START_CLUSTER_COMPLETED = 'START_CLUSTER_COMPLETED';
exports.START_GET_ENGINES = 'START_GET_ENGINES';
function receiveEngines(engines) {
    return {
        type: exports.RECEIVE_ENGINES,
        engines: engines
    };
}
exports.receiveEngines = receiveEngines;
function startCluster() {
    return {
        type: exports.START_CLUSTER
    };
}
exports.startCluster = startCluster;
function startGetEngines() {
    return {
        type: exports.START_GET_ENGINES
    };
}
exports.startGetEngines = startGetEngines;
function startClusterCompleted(response) {
    return {
        type: exports.START_CLUSTER_COMPLETED,
        response: response
    };
}
exports.startClusterCompleted = startClusterCompleted;
function fetchConfig() {
    return {
        type: exports.START_FETCH_CONFIG
    };
}
exports.fetchConfig = fetchConfig;
function fetchConfigCompleted(config) {
    return {
        type: exports.FETCH_CONFIG_COMPLETED,
        config: config
    };
}
exports.fetchConfigCompleted = fetchConfigCompleted;
function startUploadEngine() {
    return {
        type: exports.START_UPLOAD_ENGINE
    };
}
exports.startUploadEngine = startUploadEngine;
function uploadEngineCompleted(response) {
    return {
        type: exports.UPLOAD_ENGINE_COMPLETED,
        response: response
    };
}
exports.uploadEngineCompleted = uploadEngineCompleted;
function uploadEngine(file) {
    if (!file) {
        notification_actions_1.openNotification(Notification_1.NotificationType.Error, "File Error", 'No engine file selected.', null);
    }
    return function (dispatch) {
        dispatch(startUploadEngine());
        dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Info, "Update", 'Uploading engine...', null));
        var data = new FormData();
        data.append('file', file.files[0]);
        fetch("/upload?type=engine", {
            credentials: 'include',
            method: 'post',
            body: data
        }).then(function () {
            dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Confirm, "Success", 'Engine uploaded', null));
            dispatch(uploadEngineCompleted(null));
            dispatch(getEngines());
        }).catch(function (error) {
            dispatch(uploadEngineCompleted(error));
            dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, "Error", error.toString(), null));
        });
    };
}
exports.uploadEngine = uploadEngine;
function startYarnCluster(clusterName, engineId, size, memory, keytab) {
    if (!clusterName || !engineId || !size || !memory) {
        notification_actions_1.openNotification(Notification_1.NotificationType.Error, "Error", 'All fields are required', null);
    }
    return function (dispatch) {
        dispatch(startCluster());
        dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Info, "Update", 'Connecting to YARN...', null));
        Remote.startClusterOnYarn(clusterName, engineId, size, memory, keytab, function (error, clusterId) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, "Error", error.toString(), null));
                dispatch(startClusterCompleted(error.toString()));
                return;
            }
            dispatch(startClusterCompleted(clusterId));
            dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Confirm, "Success", 'Cluster Launched', null));
        });
    };
}
exports.startYarnCluster = startYarnCluster;
function getEngines() {
    return function (dispatch) {
        dispatch(startGetEngines());
        Remote.getEngines(function (error, engines) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Error', error.toString(), null));
                dispatch(receiveEngines(null));
                return;
            }
            dispatch(receiveEngines(engines));
        });
    };
}
exports.getEngines = getEngines;
function getConfig() {
    return function (dispatch) {
        dispatch(fetchConfig());
        Remote.getConfig(function (error, config) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Error', error.toString(), null));
                dispatch(fetchConfigCompleted(null));
                return;
            }
            dispatch(fetchConfigCompleted(config));
        });
    };
}
exports.getConfig = getConfig;
//# sourceMappingURL=clusters.actions.js.map