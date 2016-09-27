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
 * Created by justin on 7/22/16.
 */
var moment = require('moment');
var react_router_1 = require('react-router');
var PageHeader_1 = require('./PageHeader');
var Panel_1 = require('./Panel');
var React = require('react');
require('../styles/projectslist.scss');
var react_router_2 = require('react-router');
var ProjectsList = (function (_super) {
    __extends(ProjectsList, _super);
    function ProjectsList() {
        _super.apply(this, arguments);
    }
    ProjectsList.prototype.openProject = function (projectId) {
        react_router_2.hashHistory.push('/projects/' + projectId + '/models');
    };
    ProjectsList.prototype.render = function () {
        var _this = this;
        return (React.createElement("div", {className: "project-details"}, React.createElement(PageHeader_1.default, null, React.createElement("span", null, "PROJECTS"), React.createElement(react_router_1.Link, {to: "/newproject", className: "button-primary header-buttons"}, "Create New Project")), React.createElement("div", null, React.createElement("h1", null, "All Projects"), React.createElement("div", {className: "panel-container"}, this.props.projects.map(function (project, i) {
            return (React.createElement(Panel_1.default, {key: i, onClick: _this.openProject.bind(_this, project.id)}, React.createElement("article", null, React.createElement("div", {className: "project-metadata"}, React.createElement("header", null, project.name), React.createElement("div", null, project.model_category), React.createElement("div", null, moment.unix(project.created_at).format('YYYY-MM-DD HH:mm'))))));
        })))));
    };
    return ProjectsList;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = ProjectsList;
//# sourceMappingURL=ProjectsList.js.map