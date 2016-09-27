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
 * Created by justin on 6/28/16.
 */
var Remote = require('../../Proxy/Proxy');
var notification_actions_1 = require('../../App/actions/notification.actions');
var Notification_1 = require('../../App/components/Notification');
var react_router_1 = require('react-router');
exports.FETCH_MODEL_OVERVIEW = 'FETCH_MODEL_OVERVIEW';
exports.RECEIVE_MODEL_OVERVIEW = 'RECEIVE_MODEL_OVERVIEW';
exports.FETCH_DOWNLOAD_MODEL = 'FETCH_DOWNLOAD_MODEL';
exports.RECEIVE_DOWNLOAD_MODEL = 'RECEIVE_DOWNLOAD_MODEL';
exports.requestModelOverview = function () {
    return {
        type: exports.FETCH_MODEL_OVERVIEW
    };
};
function receiveModelOverview(model) {
    return {
        type: exports.RECEIVE_MODEL_OVERVIEW,
        model: model
    };
}
exports.receiveModelOverview = receiveModelOverview;
exports.requestDownloadModel = function () {
    return {
        type: exports.FETCH_DOWNLOAD_MODEL
    };
};
function receiveDownloadModel(model) {
    return {
        type: exports.RECEIVE_DOWNLOAD_MODEL,
        model: model
    };
}
exports.receiveDownloadModel = receiveDownloadModel;
function fetchModelOverview(modelId) {
    return function (dispatch) {
        dispatch(exports.requestModelOverview());
        Remote.getModel(modelId, function (error, model) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                return;
            }
            getModelStrategy(model.model_category.toLowerCase())(modelId, function (error, res) {
                dispatch(receiveModelOverview(res));
            });
        });
    };
}
exports.fetchModelOverview = fetchModelOverview;
function getModelStrategy(modelCategory) {
    if (modelCategory === 'binomial') {
        return Remote.getModelBinomial;
    }
    else if (modelCategory === 'multinomial') {
        return Remote.getModelMultinomial;
    }
    else if (modelCategory === 'regression') {
        return Remote.getModelRegression;
    }
}
function downloadModel() {
    /**
     * TODO(justinloyola): Waiting on endpoint
     */
    return function (dispatch) {
        dispatch(exports.requestDownloadModel());
        dispatch(receiveDownloadModel({}));
    };
}
exports.downloadModel = downloadModel;
function deployModel(modelId, name, projectId, packageName) {
    return function (dispatch) {
        dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Info, 'Deploying model', null, null));
        Remote.startService(modelId, name, packageName, function (error, res) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, "Deployment Error", error.toString(), null));
                return;
            }
            dispatch(notification_actions_1.closeNotificationManager());
            react_router_1.hashHistory.push('/projects/' + projectId + '/deployment');
        });
    };
}
exports.deployModel = deployModel;
//# sourceMappingURL=model.overview.action.js.map