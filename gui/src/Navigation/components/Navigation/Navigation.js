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
 * Created by justin on 6/25/16.
 */
var React = require('react');
var $ = require('jquery');
var classNames = require('classnames');
var react_router_1 = require('react-router');
var Sidebar_1 = require('../Sidebar/Sidebar');
var buildPath_1 = require('../../../App/utils/buildPath');
var getRoute_1 = require('../../../App/utils/getRoute');
var routes_1 = require('../../../routes');
var _ = require('lodash');
var react_redux_1 = require('react-redux');
require('./navigation.scss');
var react_motion_1 = require('react-motion');
var Navigation = (function (_super) {
    __extends(Navigation, _super);
    function Navigation() {
        _super.call(this);
        this.state = {
            activeTopLevelPath: '',
            isSubMenuActive: false,
        };
    }
    Navigation.prototype.componentWillMount = function () {
        this.setMenuState(this.props.routes);
    };
    Navigation.prototype.componentWillReceiveProps = function (nextProps) {
        this.setMenuState(nextProps.routes);
    };
    Navigation.prototype.setMenuState = function (newRoutes) {
        var _this = this;
        var currentRoutePath = newRoutes[newRoutes.length - 1].path;
        var topLevelPath = '';
        if (currentRoutePath) {
            topLevelPath = currentRoutePath.split('/')[0];
        }
        var submenuActive = false;
        _.forEach(routes_1.routes[0].childRoutes, function (route) {
            if (_this.isActive(route.path, newRoutes) && route.showChildrenAsSubmenu) {
                submenuActive = true;
            }
        });
        this.setState({
            activeTopLevelPath: topLevelPath,
            isSubMenuActive: submenuActive
        });
    };
    Navigation.prototype.isActive = function (path, newRoutes) {
        var currentRoutePath = newRoutes[newRoutes.length - 1].path;
        if (currentRoutePath && currentRoutePath.indexOf(path) !== -1) {
            return true;
        }
        return false;
    };
    Navigation.prototype.getParentRouteName = function (currentPath) {
        var newPathParts = currentPath.split('/');
        if (newPathParts.length < 2) {
            return currentPath;
        }
        newPathParts.pop();
        var newPath = newPathParts.join('/');
        var parentRoute = getRoute_1.getRoute(newPath);
        return parentRoute.name;
    };
    Navigation.prototype.logout = function () {
        $.ajax({
            url: window.location.protocol + '://' + window.location.host,
            beforeSend: function (xhr) {
                xhr.withCredentials = true;
                xhr.setRequestHeader('Authorization', 'Basic ' + btoa('fjkdshfhkjsdfjkhsdkfjhsdf:hfkjdshfdhff'));
            }
        });
    };
    Navigation.prototype.renderSubmenu = function (activeRoute, shouldShow) {
        var _this = this;
        var childRoutes = routes_1.routes[0].childRoutes.filter(function (route) {
            return (route.path.indexOf(activeRoute.path) !== -1 && route.path !== activeRoute.path);
        });
        var styleTo;
        if (shouldShow) {
            styleTo = { left: react_motion_1.spring(72) };
        }
        else {
            styleTo = { left: react_motion_1.spring(300) };
        }
        return (React.createElement(react_motion_1.Motion, {defaultStyle: { left: 300 }, style: styleTo}, function (interpolatingStyle) {
            return React.createElement("div", {className: "left-submenu", style: interpolatingStyle}, React.createElement(Sidebar_1.Sidebar, {className: 'secondary-navigation'}, React.createElement("nav", {className: "navigation--primary"}, React.createElement("div", {className: "navigation"}, React.createElement("header", null, React.createElement("div", {className: "header-navigation"}, React.createElement(react_router_1.Link, {to: _this.getParentRouteName(activeRoute.path)}, React.createElement("i", {className: "fa fa-angle-left"}), React.createElement("span", null, _this.getParentRouteName(activeRoute.path))))), React.createElement("div", {className: "header-content"}, _this.props.project.name), React.createElement("ul", {className: "nav-list"}, _.map(childRoutes, function (menuItem) {
                var path = buildPath_1.buildPath(menuItem.path, _this.props.params);
                return (!menuItem.showInNavigation) ? null : (React.createElement("li", {key: menuItem.path, className: classNames('nav-list--item', { active: _this.isActive(menuItem.path, _this.props.routes) })}, React.createElement(react_router_1.Link, {to: path}, menuItem.name)));
            }))))));
        }));
    };
    Navigation.prototype.render = function () {
        var _this = this;
        var submenu = React.createElement("div", null);
        return (React.createElement("div", {className: classNames('nav-container')}, React.createElement(Sidebar_1.Sidebar, {className: "primary-navigation"}, React.createElement("nav", {className: "navigation--primary"}, React.createElement("div", {className: "navigation"}, React.createElement("header", null, React.createElement("div", {className: "logo-container"}, React.createElement(react_router_1.Link, {to: "/"}, React.createElement("div", {className: "logo"}, "STEAM")))), React.createElement("div", {className: "header-content"}), React.createElement("ul", {className: 'nav-list'}, routes_1.routes[0].childRoutes.map(function (route) {
            var isActive = false;
            if (_this.isActive(route.path, _this.props.routes)) {
                isActive = true;
            }
            if (route.showChildrenAsSubmenu) {
                submenu = _this.renderSubmenu(route, isActive);
            }
            if (route.path.split('/').length > 1 || !route.showInNavigation) {
                return null;
            }
            var activeChildren = route.path === _this.state.activeTopLevelPath && _this.state.isSubMenuActive;
            var path = '/' + route.path;
            return (React.createElement("li", {key: path, className: classNames('nav-list--item', { active: isActive }, { activeChildren: activeChildren })}, React.createElement(react_router_1.Link, {to: path}, React.createElement("i", {className: route.icon}), React.createElement("div", {className: "nav-list--label"}, route.name))));
        }), React.createElement("li", {className: "logout nav-list--item"}, React.createElement("a", {href: "mailto:steam@h2o.ai?subject=STEAM: "}, React.createElement("i", {className: "fa fa-question-circle-o"}), React.createElement("div", {className: "nav-list--label"}, "Support")), React.createElement("a", {onClick: this.logout.bind(this)}, React.createElement("i", {className: "fa fa-sign-out"}), React.createElement("div", {className: "nav-list--label"}, "Logout"))))))), submenu));
    };
    return Navigation;
}(React.Component));
exports.Navigation = Navigation;
function mapStateToProps(state) {
    return {
        project: state.projects.project
    };
}
function mapDispatchToProps() {
    return {};
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(react_router_1.withRouter(Navigation));
//# sourceMappingURL=Navigation.js.map