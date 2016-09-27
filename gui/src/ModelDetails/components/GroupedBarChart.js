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
var ReactDOM = require('react-dom');
var visComponents = require('vis-components');
var groupedBarChart = visComponents.groupedBarChart;
var GroupedBarChart = (function (_super) {
    __extends(GroupedBarChart, _super);
    function GroupedBarChart() {
        _super.apply(this, arguments);
    }
    GroupedBarChart.prototype.componentDidMount = function () {
        this._mountNode = ReactDOM.findDOMNode(this);
        this.renderGraph();
    };
    GroupedBarChart.prototype.componentWillUnmount = function () {
        if (this._mountNode) {
            ReactDOM.unmountComponentAtNode(this._mountNode);
            this._mountNode.remove();
            this._mountNode = null;
        }
    };
    GroupedBarChart.prototype.renderGraph = function () {
        var options = {
            groupByVariable: this.props.groupByVariable,
            barColors: this.props.barColors
        };
        groupedBarChart.plot(this._mountNode, this.props.data, options);
    };
    GroupedBarChart.prototype.render = function () {
        return React.createElement("div", {className: "grouped-bar-container"});
    };
    GroupedBarChart.defaultProps = {
        data: [],
        groupByVariable: 'value',
        barColors: ['#a6cee3', '#1f78b4', '#b2df8a', '#33a02c', '#fb9a99', '#e31a1c', '#fdbf6f']
    };
    return GroupedBarChart;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = GroupedBarChart;
//# sourceMappingURL=GroupedBarChart.js.map