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
var react_router_1 = require('react-router');
var NotificationsManager_1 = require('./components/NotificationsManager');
var Navigation_1 = require('../Navigation/components/Navigation/Navigation');
var Breadcrumb_1 = require('./components/Breadcrumb');
var Body_1 = require('../Body/Body');
require('./styles/breadcrumb.scss');
require('./styles/app.scss');
var App = (function (_super) {
    __extends(App, _super);
    function App() {
        _super.apply(this, arguments);
    }
    App.prototype.render = function () {
        return (React.createElement("div", {className: "app-container"}, React.createElement(NotificationsManager_1.default, null), React.createElement(Navigation_1.default, {routes: this.props.routes, params: this.props.params}), React.createElement("div", {className: "body-container"}, React.createElement("header", null, React.createElement(Breadcrumb_1.default, {routes: this.props.routes, params: this.props.params})), React.createElement(Body_1.default, null, this.props.children))));
    };
    return App;
}(React.Component));
exports.App = App;
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_router_1.withRouter(App);
//# sourceMappingURL=App.js.map