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
require('../styles/users.scss');
var users_actions_1 = require("../actions/users.actions");
var RolePermissions = (function (_super) {
    __extends(RolePermissions, _super);
    function RolePermissions() {
        _super.apply(this, arguments);
    }
    RolePermissions.prototype.componentWillMount = function () {
        this.props.fetchPermissionsWithRoles();
    };
    RolePermissions.prototype.render = function () {
        var permissionRows;
        if (this.props.permissionsWithRoles) {
            permissionRows = this.props.permissionsWithRoles.map(function (permissionSet, index) {
                return React.createElement(Row_1.default, {key: index}, React.createElement(Cell_1.default, {className: "right-table-bar", key: permissionSet.description}, permissionSet.description), permissionSet.flags.map(function (flag, flagIndex) {
                    if (flagIndex === 0) {
                        return React.createElement(Cell_1.default, {className: "center-text", key: flagIndex}, React.createElement("input", {type: "checkbox", value: "on", checked: true, readOnly: true, disabled: true}));
                    }
                    else {
                        return React.createElement(Cell_1.default, {className: "center-text", key: flagIndex}, React.createElement("input", {type: "checkbox", value: "on", checked: flag, readOnly: true, disabled: true}));
                    }
                }));
            });
        }
        return (React.createElement("div", {className: "role-permissions intro"}, this.props.permissionsWithRoles && this.props.roles ? React.createElement(Table_1.default, null, React.createElement(Row_1.default, {header: true}, React.createElement(Cell_1.default, {className: "right-table-bar"}, "Permission Name"), this.props.roles.map(function (role, rolesIndex) {
            return React.createElement(Cell_1.default, {className: "center-text", key: rolesIndex}, role.description);
        })), permissionRows)
            : null));
    };
    return RolePermissions;
}(React.Component));
exports.RolePermissions = RolePermissions;
function mapStateToProps(state) {
    return {
        permissionsWithRoles: state.users.permissionsWithRoles,
        roles: state.users.roles
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchPermissionsWithRoles: redux_1.bindActionCreators(users_actions_1.fetchPermissionsWithRoles, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(RolePermissions);
//# sourceMappingURL=RolePermissions.js.map