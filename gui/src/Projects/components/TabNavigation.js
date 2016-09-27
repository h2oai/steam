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
 * Created by justin on 7/11/16.
 */
var React = require('react');
var classNames = require('classnames');
require('../styles/tabnavigation.scss');
var TabNavigation = (function (_super) {
    __extends(TabNavigation, _super);
    function TabNavigation() {
        _super.apply(this, arguments);
    }
    TabNavigation.prototype.render = function () {
        var _this = this;
        return (React.createElement("nav", {className: "tabs"}, Object.keys(this.props.tabs).map(function (tab, i) {
            return (React.createElement("a", {key: i, className: classNames('tab', { selected: _this.props.tabs[tab].isSelected === true }), onClick: _this.props.tabs[tab].onClick ? _this.props.tabs[tab].onClick.bind(_this, _this.props.tabs[tab]) : null}, _this.props.tabs[tab].label));
        })));
    };
    return TabNavigation;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = TabNavigation;
//# sourceMappingURL=TabNavigation.js.map