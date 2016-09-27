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
var React = require('react');
var _ = require('lodash');
var ProjectsList_1 = require('./ProjectsList');
var WelcomeSplashScreen_1 = require('./WelcomeSplashScreen');
var react_redux_1 = require('react-redux');
var redux_1 = require('redux');
var projects_actions_1 = require('../actions/projects.actions');
var ProjectScreenStrategy = (function (_super) {
    __extends(ProjectScreenStrategy, _super);
    function ProjectScreenStrategy() {
        _super.apply(this, arguments);
    }
    ProjectScreenStrategy.prototype.componentWillMount = function () {
        this.props.fetchProjects();
    };
    ProjectScreenStrategy.prototype.render = function () {
        if (this.props.projects === null) {
            return null;
        }
        if (_.isEmpty(this.props.projects)) {
            return React.createElement(WelcomeSplashScreen_1.default, null);
        }
        return (React.createElement(ProjectsList_1.default, {projects: this.props.projects}));
    };
    return ProjectScreenStrategy;
}(React.Component));
exports.ProjectScreenStrategy = ProjectScreenStrategy;
function mapStateToProps(state) {
    return {
        projects: state.projects.availableProjects
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchProjects: redux_1.bindActionCreators(projects_actions_1.fetchProjects, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(ProjectScreenStrategy);
//# sourceMappingURL=ProjectScreenStrategy.js.map