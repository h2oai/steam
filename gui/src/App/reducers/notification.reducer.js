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
 * Created by justin on 8/8/16.
 */
var _ = require('lodash');
var notification_actions_1 = require('../actions/notification.actions');
var initialState = {
    allNotifications: []
};
function notificationReducer(state, action) {
    if (state === void 0) { state = initialState; }
    var toReturn;
    switch (action.type) {
        case notification_actions_1.OPEN_NOTIFICATION:
            return {
                allNotifications: state.allNotifications.concat([action.notificationData])
            };
        case notification_actions_1.KILL_ALL_INACTIVE_NOTIFICATIONS:
            toReturn = _.assign({}, state);
            for (var _i = 0, _a = toReturn.allNotifications; _i < _a.length; _i++) {
                var notification = _a[_i];
                if (!notification.isActive) {
                    notification.isAlive = false;
                }
            }
            return toReturn;
        case notification_actions_1.DISMISS_NOTIFICATION:
            toReturn = _.assign({}, state);
            toReturn.allNotifications[action.notification.index].isActive = false;
            toReturn.allNotifications[action.notification.index].isAlive = false;
            return toReturn;
        default:
            return state;
    }
}
exports.notificationReducer = notificationReducer;
//# sourceMappingURL=notification.reducer.js.map