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
var __extends = (this && this.__extends) || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
};
/**
 * Created by justin on 8/8/16.
 */
var React = require('react');
require('../styles/notification.scss');
(function (NotificationType) {
    NotificationType[NotificationType["Info"] = 0] = "Info";
    NotificationType[NotificationType["Confirm"] = 1] = "Confirm";
    NotificationType[NotificationType["Warning"] = 2] = "Warning";
    NotificationType[NotificationType["Error"] = 3] = "Error";
})(exports.NotificationType || (exports.NotificationType = {}));
var NotificationType = exports.NotificationType;
var Notification = (function (_super) {
    __extends(Notification, _super);
    function Notification() {
        _super.call(this);
        this.state = {
            hasViewButton: false,
        };
    }
    Notification.prototype.onDismissClicked = function () {
        this.props.dismissNotification(this.props.notificationData);
    };
    Notification.prototype.render = function () {
        if (this.props.notificationData) {
            var guideClassName = "left-type-guide ";
            switch (this.props.notificationData.notificationType) {
                case NotificationType.Info:
                    guideClassName += "left-type-guide-info ";
                    break;
                case NotificationType.Warning:
                    guideClassName += "left-type-guide-warning ";
                    break;
                case NotificationType.Error:
                    guideClassName += "left-type-guide-error ";
                    break;
                case NotificationType.Confirm:
                    guideClassName += "left-type-guide-confirm ";
                    break;
            }
            return React.createElement("div", {className: 'notification'}, React.createElement("div", {className: "inner-notification"}, React.createElement("div", {className: guideClassName}), React.createElement("div", {className: "text-container"}, React.createElement("div", {className: "header"}, this.props.notificationData.header), React.createElement("div", {className: "detail"}, this.props.notificationData.detail)), React.createElement("div", {className: "actions-container"}, this.state.hasViewButton ?
                React.createElement("div", {onClick: this.onDismissClicked.bind(this), className: "view-button"}, " View") : null, React.createElement("div", {onClick: this.onDismissClicked.bind(this), className: "dismiss-button"}, " Dismiss"))));
        }
        else {
            return null;
        }
    };
    return Notification;
}(React.Component));
exports.Notification = Notification;
//# sourceMappingURL=Notification.js.map