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
var projectMembers_1 = require('./components/projectMembers');
var projectLabelsAccess_1 = require('./components/projectLabelsAccess');
require('./styles/collaborators.scss');
var Collaborators = (function (_super) {
    __extends(Collaborators, _super);
    function Collaborators() {
        _super.call(this);
        this.state = {
            tabs: {
                projectMembers: {
                    label: 'MEMBERS',
                    isSelected: true,
                    onClick: this.switchTab.bind(this),
                    component: React.createElement(projectMembers_1.default, null)
                },
                labelsAccess: {
                    label: 'LABELS ACCESS',
                    isSelected: false,
                    onClick: this.switchTab.bind(this),
                    component: React.createElement(projectLabelsAccess_1.default, null)
                }
            },
            isSelected: 'projectMembers'
        };
    }
    Collaborators.prototype.componentWillMount = function () {
        this.setState({
            tabs: {
                projectMembers: {
                    label: 'MEMBERS',
                    isSelected: true,
                    onClick: this.switchTab.bind(this),
                    component: React.createElement(projectMembers_1.default, null)
                },
                labelsAccess: {
                    label: 'LABELS ACCESS',
                    isSelected: false,
                    onClick: this.switchTab.bind(this),
                    component: React.createElement(projectLabelsAccess_1.default, null)
                }
            }
        });
    };
    Collaborators.prototype.switchTab = function (newTab) {
        var key = _.findKey(this.state.tabs, newTab);
        var newState = _.cloneDeep(this.state);
        Object.keys(newState.tabs).map(function (tab) {
            newState.tabs[tab].isSelected = false;
        });
        newState.tabs[key].isSelected = true;
        newState.isSelected = key;
        this.setState(newState);
    };
    Collaborators.prototype.render = function () {
        return (React.createElement("div", {className: "collaborators"}, React.createElement(PageHeader_1.default, null, React.createElement("span", null, "Collaborators")), React.createElement(TabNavigation_1.default, {tabs: this.state.tabs}), this.state.tabs.projectMembers.isSelected === true ?
            React.createElement(projectMembers_1.default, {projectid: this.props.params.projectid, loadLabelsTab: this.switchTab.bind(this, this.state.tabs.labelsAccess)}) : null, this.state.tabs.labelsAccess.isSelected === true ?
            React.createElement(projectLabelsAccess_1.default, {projectid: this.props.params.projectid}) : null));
    };
    return Collaborators;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = Collaborators;
//# sourceMappingURL=Collaborators.js.map