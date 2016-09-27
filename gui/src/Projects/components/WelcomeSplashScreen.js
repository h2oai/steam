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
 * Created by justin on 7/10/16.
 */
var React = require('react');
var classNames = require('classnames');
var react_router_1 = require('react-router');
require('../styles/welcomesplashscreen.scss');
var WelcomeSplashScreen = (function (_super) {
    __extends(WelcomeSplashScreen, _super);
    function WelcomeSplashScreen() {
        _super.apply(this, arguments);
    }
    WelcomeSplashScreen.prototype.render = function () {
        return (React.createElement("div", {className: "welcome-splash-screen"}, React.createElement("div", {className: "welcome-splash-screen--content"}, React.createElement("div", null, "WELCOME TO"), React.createElement("div", {className: "welcome-splash-screen--content--product-name"}, "H", React.createElement("sub", null, "2"), "O STEAM"), React.createElement("div", null, "Fast, distributed data science for teams"), React.createElement("div", null, React.createElement(react_router_1.Link, {to: "/newproject", className: classNames('default', 'start-project')}, "Start A New Project")))));
    };
    return WelcomeSplashScreen;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = WelcomeSplashScreen;
//# sourceMappingURL=WelcomeSplashScreen.js.map