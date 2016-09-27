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
 * Created by justin on 7/5/16.
 */
var React = require('react');
var react_router_1 = require('react-router');
var PageHeader_1 = require('../components/PageHeader');
require('../styles/newproject.scss');
var NewProjectStep1 = (function (_super) {
    __extends(NewProjectStep1, _super);
    function NewProjectStep1() {
        _super.apply(this, arguments);
    }
    NewProjectStep1.prototype.render = function () {
        return (React.createElement("div", {className: "new-project"}, React.createElement(PageHeader_1.default, null, "New Project"), React.createElement("div", {className: "project-description"}, React.createElement("span", null, "Steam organizes your data sets, your models, and your deployment configurations into one cohesive project. This enables you to:"), React.createElement("ul", {className: "project-description-list"}, React.createElement("li", null, "Visually compare all models within a project"), React.createElement("li", null, "Manage how models from a project gets deployed"), React.createElement("li", null, "Track the history of model deployment"))), React.createElement("div", {className: "cards-container"}, React.createElement("div", {className: "card small"}, React.createElement("header", null, "Import Existing Models"), React.createElement("article", null, "Choose this option if you already have H2O models built and stored in a H2O cluster in your network."), React.createElement("footer", null, React.createElement(react_router_1.Link, {to: "newproject/import", className: "default"}, "Start Import"))), React.createElement("div", {className: "card small"}, React.createElement("header", null, "Start from Scratch"), React.createElement("article", null, "Choose this option if this is a completely new project."), React.createElement("footer", null, React.createElement(react_router_1.Link, {to: "newproject/import", className: "default"}, "Create New Project"))))));
    };
    return NewProjectStep1;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = NewProjectStep1;
//# sourceMappingURL=NewProjectStep1.js.map