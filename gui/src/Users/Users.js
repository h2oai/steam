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
var PageHeader_1 = require('../Projects/components/PageHeader');
var TabNavigation_1 = require('../Projects/components/TabNavigation');
var UserAccess_1 = require('./components/UserAccess');
var RolePermissions_1 = require('./components/RolePermissions');
require('./styles/users.scss');
var Users = (function (_super) {
    __extends(Users, _super);
    function Users() {
        _super.call(this);
        this.state = {
            tabs: {
                users: {
                    label: 'USERS',
                    isSelected: true,
                    onClick: this.clickHandler.bind(this),
                    component: React.createElement(UserAccess_1.default, null)
                },
                packaging: {
                    label: 'ROLES',
                    isSelected: false,
                    onClick: this.clickHandler.bind(this),
                    component: React.createElement(RolePermissions_1.default, null)
                }
            },
            isSelected: 'users'
        };
    }
    Users.prototype.componentWillMount = function () {
        this.setState({
            tabs: {
                users: {
                    label: 'USERS',
                    isSelected: true,
                    onClick: this.clickHandler.bind(this),
                    component: React.createElement(UserAccess_1.default, null)
                },
                roles: {
                    label: 'ROLES',
                    isSelected: false,
                    onClick: this.clickHandler.bind(this),
                    component: React.createElement(RolePermissions_1.default, null)
                }
            }
        });
    };
    Users.prototype.clickHandler = function (tab) {
        var key = _.findKey(this.state.tabs, tab);
        var newState = _.cloneDeep(this.state);
        Object.keys(newState.tabs).map(function (tab) {
            newState.tabs[tab].isSelected = false;
        });
        newState.tabs[key].isSelected = true;
        newState.isSelected = key;
        this.setState(newState);
    };
    Users.prototype.render = function () {
        return (React.createElement("div", {className: "users"}, React.createElement(PageHeader_1.default, null, "USERS"), React.createElement("div", {className: "panel-container"}, React.createElement(TabNavigation_1.default, {tabs: this.state.tabs}), this.state.tabs.users.isSelected === true ?
            React.createElement(UserAccess_1.default, null) : null, this.state.tabs.roles.isSelected === true ?
            React.createElement(RolePermissions_1.default, null) : null)));
    };
    return Users;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = Users;
//# sourceMappingURL=Users.js.map