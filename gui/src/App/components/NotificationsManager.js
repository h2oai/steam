"use strict";
var __extends = (this && this.__extends) || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
};
var React = require('react');
var _ = require('lodash');
var react_redux_1 = require('react-redux');
var notification_actions_1 = require('../actions/notification.actions');
var redux_1 = require('redux');
var Notification_1 = require("./Notification");
var react_motion_1 = require('react-motion');
var NotificationsManager = (function (_super) {
    __extends(NotificationsManager, _super);
    function NotificationsManager() {
        _super.call(this);
        this.timeoutTimerID = -1;
        this.TIMEOUT_LENGTH = 5000;
        this.state = {
            isActive: false,
            hasViewButton: false,
            style: { top: -150, opacity: 1 },
            defaultStyle: { top: -150, opacity: 1 }
        };
    }
    NotificationsManager.prototype.componentWillReceiveProps = function (nextProps) {
        if (nextProps.notifications) {
            var isActive = false;
            for (var _i = 0, _a = nextProps.notifications.allNotifications; _i < _a.length; _i++) {
                var notificationData = _a[_i];
                if (notificationData.isActive) {
                    isActive = true;
                }
            }
            if (isActive) {
                this.setState({
                    isActive: true,
                    style: { top: react_motion_1.spring(22), opacity: 1 },
                    defaultStyle: { top: -150, opacity: 1 }
                });
            }
            else {
                this.setFadeout();
            }
        }
        if (this.timeoutTimerID !== -1) {
            clearTimeout(this.timeoutTimerID);
            this.timeoutTimerID = setTimeout(this.onTimeout.bind(this), this.TIMEOUT_LENGTH);
        }
    };
    NotificationsManager.prototype.setFadeout = function () {
        this.setState({
            isActive: false,
            style: { top: -150, opacity: react_motion_1.spring(0) },
            defaultStyle: { top: 22, opacity: 1 }
        });
    };
    NotificationsManager.prototype.onAnimationComplete = function () {
        if (this.state.isActive) {
            this.startTimeoutTimer();
        }
        else {
            this.props.killAllInactiveNotifications();
        }
    };
    NotificationsManager.prototype.startTimeoutTimer = function () {
        this.timeoutTimerID = setTimeout(this.onTimeout.bind(this), this.TIMEOUT_LENGTH);
    };
    NotificationsManager.prototype.onTimeout = function () {
        for (var _i = 0, _a = this.props.notifications.allNotifications; _i < _a.length; _i++) {
            var notificationData = _a[_i];
            notificationData.isActive = false;
        }
        this.setFadeout();
    };
    NotificationsManager.prototype.dismissNotification = function (notificationData) {
        this.props.dismissNotification(notificationData);
    };
    NotificationsManager.prototype.render = function () {
        var _this = this;
        if (this.props.notifications.allNotifications) {
            return (React.createElement(react_motion_1.Motion, {defaultStyle: this.state.defaultStyle, style: this.state.style, onRest: this.onAnimationComplete.bind(this)}, function (animationValue) {
                var modifiedAnimationValue = _.assign({}, animationValue);
                if (modifiedAnimationValue.opacity < 1) {
                    modifiedAnimationValue.top = 22;
                }
                return React.createElement("div", {className: "notification-manager", style: modifiedAnimationValue}, " ", _this.props.notifications.allNotifications.map(function (notificationData, index, array) {
                    if (notificationData.isAlive) {
                        return React.createElement(Notification_1.Notification, {key: index, notificationData: notificationData, dismissNotification: _this.dismissNotification.bind(_this)});
                    }
                }));
            }));
        }
        else {
            return null;
        }
    };
    return NotificationsManager;
}(React.Component));
exports.NotificationsManager = NotificationsManager;
function mapStateToProps(state) {
    return {
        notifications: state.notification
    };
}
function mapDispatchToProps(dispatch) {
    return {
        dismissNotification: redux_1.bindActionCreators(notification_actions_1.dismissNotification, dispatch),
        closeNotificationManager: redux_1.bindActionCreators(notification_actions_1.closeNotificationManager, dispatch),
        killAllInactiveNotifications: redux_1.bindActionCreators(notification_actions_1.killAllInactiveNotifications, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(NotificationsManager);
//# sourceMappingURL=NotificationsManager.js.map