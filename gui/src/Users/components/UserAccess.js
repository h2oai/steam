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
var _ = require('lodash');
var Table_1 = require('../../Projects/components/Table');
var Row_1 = require('../../Projects/components/Row');
var Cell_1 = require('../../Projects/components/Cell');
var redux_1 = require('redux');
var react_redux_1 = require('react-redux');
require('../styles/users.scss');
var users_actions_1 = require("../actions/users.actions");
var UserAccess = (function (_super) {
    __extends(UserAccess, _super);
    function UserAccess() {
        _super.apply(this, arguments);
    }
    UserAccess.prototype.componentWillMount = function () {
        if (!this.props.usersWithRolesAndProjects) {
            this.props.fetchUsersWithRolesAndProjects();
        }
    };
    UserAccess.prototype.onRoleCheckboxClicked = function (e) {
        this.props.changeFilterSelections(parseInt(e.target.dataset.id, 10), e.target.checked);
    };
    UserAccess.prototype.checkIsRoleSelected = function (id) {
        if (_.isEmpty(this.props.selectedRoles)) {
            return false;
        }
        var index = _.findIndex(this.props.selectedRoles, function (o) {
            if (o.id === (id))
                return true;
            return false;
        });
        if (index === -1)
            console.log("ERROR: unable to find match");
        return this.props.selectedRoles[index].selected;
    };
    UserAccess.prototype.shouldRowBeShown = function (roles) {
        var _loop_1 = function(role) {
            var index = _.findIndex(this_1.props.selectedRoles, function (o) {
                if (o.id === role.id) {
                    return true;
                }
                else {
                    return false;
                }
            });
            if (index === -1)
                return { value: false };
            if (this_1.props.selectedRoles[index].selected) {
                return { value: true };
            }
        };
        var this_1 = this;
        for (var _i = 0, roles_1 = roles; _i < roles_1.length; _i++) {
            var role = roles_1[_i];
            var state_1 = _loop_1(role);
            if (typeof state_1 === "object") return state_1.value;
        }
        return false;
    };
    UserAccess.prototype.render = function () {
        var _this = this;
        return (React.createElement("div", {className: "user-access intro"}, React.createElement("div", {className: "filter-column"}, "FILTERS", React.createElement(Table_1.default, {className: "full-size"}, React.createElement(Row_1.default, {header: true}, React.createElement(Cell_1.default, null, "ROLES")), this.props.roles ?
            React.createElement("div", {ref: "roleBoxes"}, this.props.roles.map(function (role, index) {
                return React.createElement(Row_1.default, {key: role.name}, React.createElement(Cell_1.default, null, React.createElement("input", {type: "checkbox", name: "selectedRoles", "data-id": role.id, checked: _this.checkIsRoleSelected(role.id), onChange: _this.onRoleCheckboxClicked.bind(_this)}), " ", role.name));
            }))
            : null)), React.createElement("div", {className: "user-access-list"}, React.createElement(Table_1.default, null, React.createElement(Row_1.default, {header: true}, React.createElement(Cell_1.default, null, "User"), React.createElement(Cell_1.default, null, "Role")), this.props.usersWithRolesAndProjects ?
            this.props.usersWithRolesAndProjects.map(function (userWithRoleAndProject, index) {
                if (_this.shouldRowBeShown(userWithRoleAndProject.roles)) {
                    return React.createElement(Row_1.default, {key: index}, React.createElement(Cell_1.default, null, userWithRoleAndProject.user.name), React.createElement(Cell_1.default, null, " ", userWithRoleAndProject.roles.map(function (role, index) {
                        return React.createElement("span", {key: index}, role.name);
                    })));
                }
                else {
                    return null;
                }
            })
            : null))));
    };
    return UserAccess;
}(React.Component));
exports.UserAccess = UserAccess;
function mapStateToProps(state) {
    return {
        projects: state.users.projects,
        roles: state.users.roles,
        users: state.users.users,
        usersWithRolesAndProjects: state.users.usersWithRolesAndProjects,
        selectedRoles: state.users.selectedRoles
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchUsersWithRolesAndProjects: redux_1.bindActionCreators(users_actions_1.fetchUsersWithRolesAndProjects, dispatch),
        changeFilterSelections: redux_1.bindActionCreators(users_actions_1.changeFilterSelections, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(UserAccess);
//# sourceMappingURL=UserAccess.js.map