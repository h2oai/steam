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
var notification_actions_1 = require('./notification.actions');
var Notification_1 = require('../components/Notification');
exports.REQUEST_ENTITY_IDS = 'REQUEST_ENTITY_IDS';
exports.RECEIVE_ENTITY_IDS = 'RECEIVE_ENTITY_IDS';
function requestEntityIds() {
    return {
        type: exports.REQUEST_ENTITY_IDS
    };
}
exports.requestEntityIds = requestEntityIds;
function receiveEntityIds(response) {
    return {
        type: exports.RECEIVE_ENTITY_IDS,
        response: response
    };
}
exports.receiveEntityIds = receiveEntityIds;
function fetchEntityIds() {
    return function (dispatch, getState) {
        dispatch(requestEntityIds());
        return new Promise(function (resolve, reject) {
            Remote.getAllEntityTypes(function (error, res) {
                if (error) {
                    notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', 'There was an error retrieving permissions list', null);
                    reject(error);
                    return;
                }
                dispatch(receiveEntityIds(res));
                resolve(res);
            });
        });
    };
}
exports.fetchEntityIds = fetchEntityIds;
//# sourceMappingURL=global.actions.js.map