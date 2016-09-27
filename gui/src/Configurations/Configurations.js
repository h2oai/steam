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
 * Created by Jeff Fohl <jfohl@h2o.ai> on 7/29/16.
 */
var React = require('react');
var _ = require('lodash');
var PageHeader_1 = require('../Projects/components/PageHeader');
var TabNavigation_1 = require('../Projects/components/TabNavigation');
var Labels_1 = require('./components/Labels');
require('./styles/configurations.scss');
var Configurations = (function (_super) {
    __extends(Configurations, _super);
    /**
     * TODO(jefffohl): Make the tab container a generalized container, like <TabContainer>, to keep things DRY.
     */
    function Configurations() {
        _super.call(this);
        this.state = {
            tabs: {},
            isSelected: null
        };
    }
    Configurations.prototype.componentWillMount = function () {
        this.setState({
            tabs: {
                labels: {
                    label: 'Labels',
                    isSelected: true,
                    onClick: this.clickHandler.bind(this),
                    component: React.createElement(Labels_1.default, {projectid: parseInt(this.props.params.projectid, 10)})
                }
            },
            isSelected: 'labels'
        });
    };
    Configurations.prototype.clickHandler = function (tab) {
        var key = _.findKey(this.state.tabs, tab);
        var newState = _.cloneDeep(this.state);
        Object.keys(newState.tabs).map(function (tab) {
            newState.tabs[tab].isSelected = false;
        });
        newState.tabs[key].isSelected = true;
        newState.isSelected = key;
        this.setState(newState);
    };
    Configurations.prototype.render = function () {
        return (React.createElement("div", {className: "services"}, React.createElement(PageHeader_1.default, null, React.createElement("span", null, "Project Configurations")), React.createElement(TabNavigation_1.default, {tabs: this.state.tabs}), React.createElement("main", null, this.state.tabs[this.state.isSelected].component)));
    };
    return Configurations;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = Configurations;
//# sourceMappingURL=Configurations.js.map