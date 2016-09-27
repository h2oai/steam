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
 * Created by justin on 6/28/16.
 */
var React = require('react');
var Table_1 = require('../../Projects/components/Table');
var Row_1 = require('../../Projects/components/Row');
var Cell_1 = require('../../Projects/components/Cell');
var GroupedBarChart_1 = require('./GroupedBarChart');
var d3 = require('d3');
var getOrdinal_1 = require('../../App/utils/getOrdinal');
require('../styles/variableimportance.scss');
// sample data
var responseDistributionSubset_1 = require('../data/responseDistributionSubset');
var VariableImportance = (function (_super) {
    __extends(VariableImportance, _super);
    function VariableImportance() {
        _super.apply(this, arguments);
        this.widthScale = function () { };
    }
    VariableImportance.prototype.componentWillMount = function () {
        this.widthScale = d3.scaleLinear()
            .domain([0, 1])
            .range([0, this.props.rowWidth]);
    };
    VariableImportance.prototype.render = function () {
        var _this = this;
        return (React.createElement("div", {className: "variable-importance metrics"}, React.createElement(Table_1.default, null, React.createElement(Row_1.default, {header: true}, React.createElement(Cell_1.default, null, React.createElement("i", {className: "fa fa-caret-down"})), React.createElement(Cell_1.default, null, "COLUMN"), React.createElement(Cell_1.default, {className: "graph"}, "IMPORTANCE"), React.createElement(Cell_1.default, {className: "graph"}, "PARTIAL DEPENDENCY"), React.createElement(Cell_1.default, {className: "graph"}, React.createElement("div", null, "RESPONSE DISTRIBUTION"), React.createElement("div", {className: "legend"}, React.createElement("svg", {width: this.props.rowWidth, height: 0.37 * this.props.rowHeight}, React.createElement("g", {transform: "translate(0, 0)"}, React.createElement("rect", {x: "0", y: "10", width: "20", height: "20", rx: "0", ry: "0", className: "symbol yes"}), React.createElement("text", {x: "30", y: "10", dy: "12", className: "legendText"}, "Yes")), React.createElement("g", {transform: "translate(85, 0)"}, React.createElement("rect", {x: "0", y: "10", width: "20", height: "20", rx: "0", ry: "0", className: "symbol no"}), React.createElement("text", {x: "30", y: "10", dy: "12", className: "legendText"}, "No"))))), React.createElement(Cell_1.default, {className: "graph"}, "NOTES")), this.props.columns.map(function (item, i) {
            return (React.createElement(Row_1.default, {key: i}, React.createElement(Cell_1.default, null, (i + 1) + getOrdinal_1.getOrdinal(i + 1)), React.createElement(Cell_1.default, null, React.createElement("div", {className: "variableImportance"}, React.createElement("div", {className: "columnName"}, item.name), React.createElement("div", {className: "detail"}, item.type))), React.createElement(Cell_1.default, null, React.createElement("div", null, React.createElement("svg", {width: _this.props.rowWidth, height: _this.props.rowHeight}, React.createElement("rect", {x: "0", y: "0", width: _this.widthScale(item.importance), height: _this.props.rowHeight, rx: "0", ry: "0", className: "bar"}))), React.createElement("div", {className: "detail"}, item.importance)), React.createElement(Cell_1.default, {className: "graph"}), React.createElement(Cell_1.default, null, React.createElement(GroupedBarChart_1.default, {data: _this.props.data['responseDistributionSubset'][i]['responseCounts']})), React.createElement(Cell_1.default, null)));
        }))));
    };
    VariableImportance.defaultProps = {
        rowHeight: 70,
        rowWidth: 210,
        columns: [
            {
                name: 'tenure',
                type: 'numeric',
                importance: 0.97
            },
            {
                name: 'gender',
                type: 'enum(2)',
                importance: 0.88
            },
            {
                name: 'PhoneService',
                type: 'enum(3)',
                importance: 0.72
            },
            {
                name: 'OnlineSecurity',
                type: 'enum(3)',
                importance: 0.32
            },
            {
                name: 'Dependents',
                type: 'categorical',
                importance: 0.21
            }
        ],
        data: {
            responseDistributionSubset: responseDistributionSubset_1.responseDistributionSubset
        }
    };
    return VariableImportance;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = VariableImportance;
//# sourceMappingURL=VariableImportance.js.map