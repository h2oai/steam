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
 * Created by justin on 7/12/16.
 */
var React = require('react');
var Table_1 = require('../../Projects/components/Table');
var Row_1 = require('../../Projects/components/Row');
var Cell_1 = require('../../Projects/components/Cell');
var deployment_actions_1 = require('../actions/deployment.actions');
var redux_1 = require('redux');
var react_redux_1 = require('react-redux');
require('../styles/packaging.scss');
var Packaging = (function (_super) {
    __extends(Packaging, _super);
    function Packaging() {
        _super.apply(this, arguments);
    }
    Packaging.prototype.componentWillMount = function () {
        this.props.fetchPackages(parseInt(this.props.projectId, 10));
    };
    Packaging.prototype.render = function () {
        return (React.createElement("div", {className: "packaging"}, React.createElement("h1", null, "PREPROCESSING PACKAGES"), React.createElement("div", {className: "intro"}, "Custom packaging methods for model deployment"), React.createElement(Table_1.default, null, React.createElement(Row_1.default, {header: true}), this.props.deployments.packages.map(function (packageName, i) {
            return (React.createElement(Row_1.default, {key: i}, React.createElement(Cell_1.default, {className: "folder-icon"}, React.createElement("i", {className: "fa fa-folder"})), React.createElement(Cell_1.default, null, packageName)));
        }))));
    };
    return Packaging;
}(React.Component));
exports.Packaging = Packaging;
function mapStateToProps(state) {
    return {
        deployments: state.deployments,
        projects: state.projects.project
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchPackages: redux_1.bindActionCreators(deployment_actions_1.fetchPackages, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(Packaging);
//# sourceMappingURL=Packaging.js.map