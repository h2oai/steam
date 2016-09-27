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
var React = require('react');
var Table_1 = require('../../Projects/components/Table');
var Row_1 = require('../../Projects/components/Row');
var Cell_1 = require('../../Projects/components/Cell');
var redux_1 = require('redux');
var react_redux_1 = require('react-redux');
require('../styles/collaborators.scss');
var collaborators_actions_1 = require('../actions/collaborators.actions');
var projects_actions_1 = require('../../Projects/actions/projects.actions');
var ProjectMembers = (function (_super) {
    __extends(ProjectMembers, _super);
    function ProjectMembers() {
        _super.apply(this, arguments);
    }
    ProjectMembers.prototype.componentWillMount = function () {
        this.props.setCurrentProject(parseInt(this.props.projectid, 10));
        this.props.fetchMembersForProject();
    };
    ProjectMembers.prototype.render = function () {
        return (React.createElement("div", {className: "projectMembers"}, React.createElement("p", null), React.createElement("h1", null, "Members"), React.createElement("p", {className: "lede"}, "Theses are users who have access to this project, meaning they can see data, models and services associated with the project. Additionally, owners and collaborators can create new models, and new services based on those models."), React.createElement("p", {className: "lede"}, "Labels associated with projects have ", React.createElement("span", {className: "link", onClick: this.props.loadLabelsTab}, "their own access controls, shown here"), "."), React.createElement(Table_1.default, null, React.createElement(Row_1.default, {header: true}, React.createElement(Cell_1.default, null, "USER"), React.createElement(Cell_1.default, null, "ROLE"), React.createElement(Cell_1.default, null, "ACCESS")), this.props.members ?
            this.props.members.map(function (member, index) {
                return React.createElement(Row_1.default, {key: index}, React.createElement(Cell_1.default, null, member.identity_name), React.createElement(Cell_1.default, null, member.role_name), React.createElement(Cell_1.default, null, member.kind));
            })
            : null)));
    };
    return ProjectMembers;
}(React.Component));
exports.ProjectMembers = ProjectMembers;
function mapStateToProps(state) {
    return {
        members: state.collaborators.members
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchMembersForProject: redux_1.bindActionCreators(collaborators_actions_1.fetchMembersForProject, dispatch),
        setCurrentProject: redux_1.bindActionCreators(projects_actions_1.setCurrentProject, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(ProjectMembers);
//# sourceMappingURL=projectMembers.js.map