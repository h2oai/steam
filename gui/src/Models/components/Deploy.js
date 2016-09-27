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
 * Created by justin on 6/30/16.
 */
var React = require('react');
var _ = require('lodash');
var PageHeader_1 = require('../../Projects/components/PageHeader');
var Table_1 = require('../../Projects/components/Table');
var Row_1 = require('../../Projects/components/Row');
var Cell_1 = require('../../Projects/components/Cell');
var DefaultModal_1 = require('../../App/components/DefaultModal');
require('../styles/deploy.scss');
var Deploy = (function (_super) {
    __extends(Deploy, _super);
    function Deploy() {
        _super.apply(this, arguments);
    }
    Deploy.prototype.deploy = function () {
        this.props.onDeploy(this.props.model, _.get(this.refs.serviceName, 'value', ''), _.get(this.refs.packageName, 'value', ''));
    };
    Deploy.prototype.render = function () {
        return (React.createElement(DefaultModal_1.default, {className: "deploy-modal", open: this.props.open}, React.createElement("div", null, React.createElement(PageHeader_1.default, null, "DEPLOY ", _.get(this.props.model, 'name', '')), React.createElement("section", null, React.createElement(Table_1.default, {className: "deployment-table"}, React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "CONFIGURE SERVICE"), React.createElement(Cell_1.default, null, React.createElement("span", null, "Steam automatically selects a port that's not in use based on the port range set by your admin."), React.createElement("label", {className: "muted"}, "Service name"), React.createElement("input", {ref: "serviceName", type: "text"}))), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null), React.createElement(Cell_1.default, null, React.createElement("label", {className: "muted"}, "Preprocessing Script"), React.createElement("select", {ref: "packageName"}, React.createElement("option", {value: ""}, "None (Default)"), this.props.packages.map(function (packageName, i) {
            return React.createElement("option", {key: i, value: packageName}, packageName);
        })))), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null), React.createElement(Cell_1.default, null, React.createElement("button", {type: "button", className: "default deploy-button", onClick: this.deploy.bind(this)}, "Deploy"), React.createElement("button", {type: "button", className: "default invert", onClick: this.props.onCancel.bind(this)}, "Cancel"))))))));
    };
    return Deploy;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = Deploy;
//# sourceMappingURL=Deploy.js.map