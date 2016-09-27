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
var ProjectLabelsAccess = (function (_super) {
    __extends(ProjectLabelsAccess, _super);
    function ProjectLabelsAccess() {
        _super.apply(this, arguments);
    }
    ProjectLabelsAccess.prototype.componentWillMount = function () {
        this.props.setCurrentProject(parseInt(this.props.projectid, 10));
        this.props.fetchLabelsForProject();
    };
    ProjectLabelsAccess.prototype.render = function () {
        return (React.createElement("div", {className: "labelsAccess"}, React.createElement("p", null), React.createElement("h1", null, "Labels Access"), React.createElement("p", {className: "lede"}, "A label is applied to a particular model to designate it for particular use, e.g. a 'prod' label to make a model as fit for production. Labels have restricted access control to make sure only users with appropriate privileges can change which models have a particular label."), React.createElement("p", {className: "lede"}, "All labels for this project, along with users privileges for that label, are listed below."), React.createElement(Table_1.default, null, React.createElement(Row_1.default, {header: true}, React.createElement(Cell_1.default, null, "LABEL"), React.createElement(Cell_1.default, null, "USERS")), this.props.labels ?
            this.props.labels.map(function (label, labelIndex) {
                return React.createElement(Row_1.default, {key: labelIndex}, React.createElement(Cell_1.default, null, label.name), React.createElement(Cell_1.default, null, label.identities ? label.identities.map(function (identity, identityIndex) {
                    return React.createElement("div", {key: identityIndex}, React.createElement("span", {className: "access-name"}, identity.identity_name), " ", React.createElement("span", {className: "access-type"}, identity.kind));
                }) : null));
            }) : null)));
    };
    return ProjectLabelsAccess;
}(React.Component));
exports.ProjectLabelsAccess = ProjectLabelsAccess;
function mapStateToProps(state) {
    return {
        labels: state.collaborators.labels
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchLabelsForProject: redux_1.bindActionCreators(collaborators_actions_1.fetchLabelsForProject, dispatch),
        setCurrentProject: redux_1.bindActionCreators(projects_actions_1.setCurrentProject, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(ProjectLabelsAccess);
//# sourceMappingURL=projectLabelsAccess.js.map