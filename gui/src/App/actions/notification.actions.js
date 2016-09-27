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
exports.OPEN_NOTIFICATION = 'OPEN_NOTIFICATION';
exports.CLOSE_NOTIFICATION = 'CLOSE_NOTIFICATION';
exports.KILL_ALL_INACTIVE_NOTIFICATIONS = 'KILL_ALL_INACTIVE_NOTIFICATIONS';
exports.DISMISS_NOTIFICATION = 'DISMISS_NOTIFICATION';
exports.CLOSE_NOTIFICATION_MANAGER = 'CLOSE_NOTIFICATION_MANAGER';
function openNotification(notificationType, header, detail, actions) {
    return function (dispatch, getState) {
        dispatch(_openNotification(notificationType, header, detail, actions, getState()));
    };
}
exports.openNotification = openNotification;
function _openNotification(notificationType, header, detail, actions, state) {
    var index = state.notification.allNotifications.length;
    return {
        type: exports.OPEN_NOTIFICATION,
        notificationData: {
            isActive: true,
            isAlive: true,
            notificationType: notificationType,
            header: header,
            detail: detail,
            actions: actions,
            index: index
        }
    };
}
function dismissNotification(notification) {
    return function (dispatch, getState) {
        dispatch({
            type: exports.DISMISS_NOTIFICATION,
            notification: notification
        });
    };
}
exports.dismissNotification = dismissNotification;
function closeNotificationManager() {
    return {
        type: exports.CLOSE_NOTIFICATION_MANAGER
    };
}
exports.closeNotificationManager = closeNotificationManager;
function killAllInactiveNotifications() {
    return {
        type: exports.KILL_ALL_INACTIVE_NOTIFICATIONS
    };
}
exports.killAllInactiveNotifications = killAllInactiveNotifications;
//# sourceMappingURL=notification.actions.js.map