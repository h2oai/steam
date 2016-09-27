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
var Remote = require('../../Proxy/Proxy');
var notification_actions_1 = require('../../App/actions/notification.actions');
var Notification_1 = require('../../App/components/Notification');
exports.FETCH_LEADERBOARD = 'FETCH_LEADERBOARD';
exports.RECEIVE_LEADERBOARD = 'RECEIVE_LEADERBOARD';
exports.RECEIVE_SORT_CRITERIA = 'RECEIVE_SORT_CRITERIA';
exports.RECEIVE_MODEL_COUNT = 'RECEIVE_MODEL_COUNT';
exports.MAX_ITEMS = 5;
exports.requestLeaderboard = function () {
    return {
        type: exports.FETCH_LEADERBOARD
    };
};
function receiveLeaderboard(leaderboard) {
    return {
        type: exports.RECEIVE_LEADERBOARD,
        leaderboard: leaderboard
    };
}
exports.receiveLeaderboard = receiveLeaderboard;
function fetchLeaderboard(projectId, modelCategory, name, sortBy, ascending, offset) {
    return function (dispatch) {
        dispatch(exports.requestLeaderboard());
        findModelStrategy(modelCategory.toLowerCase())(projectId, name, sortBy || '', ascending || false, offset, exports.MAX_ITEMS, function (error, models) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                return;
            }
            dispatch(receiveLeaderboard(models));
        });
    };
}
exports.fetchLeaderboard = fetchLeaderboard;
function findModelStrategy(modelCategory) {
    if (modelCategory === 'binomial') {
        return Remote.findModelsBinomial;
    }
    else if (modelCategory === 'multinomial') {
        return Remote.findModelsMultinomial;
    }
    else if (modelCategory === 'regression') {
        return Remote.findModelsRegression;
    }
}
exports.findModelStrategy = findModelStrategy;
function receiveSortCriteria(criteria) {
    return {
        type: exports.RECEIVE_SORT_CRITERIA,
        criteria: criteria
    };
}
exports.receiveSortCriteria = receiveSortCriteria;
function fetchSortCriteria(modelCategory) {
    return function (dispatch) {
        getSortStrategy(modelCategory)(function (error, criteria) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                return;
            }
            dispatch(receiveSortCriteria(criteria));
        });
    };
}
exports.fetchSortCriteria = fetchSortCriteria;
function getSortStrategy(modelCategory) {
    if (modelCategory === 'binomial') {
        return Remote.getAllBinomialSortCriteria;
    }
    else if (modelCategory === 'multinomial') {
        return Remote.getAllMultinomialSortCriteria;
    }
    else if (modelCategory === 'regression') {
        return Remote.getAllRegressionSortCriteria;
    }
}
function linkLabelWithModel(labelId, modelId) {
    return function (dispatch) {
        return new Promise(function (resolve, reject) {
            Remote.linkLabelWithModel(labelId, modelId, function (error) {
                if (error) {
                    dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                    reject(error);
                    return;
                }
                resolve();
            });
        });
    };
}
exports.linkLabelWithModel = linkLabelWithModel;
function unlinkLabelFromModel(labelId, modelId) {
    return function (dispatch) {
        return new Promise(function (resolve, reject) {
            Remote.unlinkLabelFromModel(labelId, modelId, function (error) {
                if (error) {
                    dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                    reject(error);
                    return;
                }
                resolve();
            });
        });
    };
}
exports.unlinkLabelFromModel = unlinkLabelFromModel;
function receiveModelCount(count) {
    return {
        type: exports.RECEIVE_MODEL_COUNT,
        count: count
    };
}
exports.receiveModelCount = receiveModelCount;
function findModelsCount(projectId) {
    return function (dispatch) {
        return Remote.findModelsCount(projectId, function (error, count) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                return;
            }
            dispatch(receiveModelCount(count));
        });
    };
}
exports.findModelsCount = findModelsCount;
//# sourceMappingURL=leaderboard.actions.js.map