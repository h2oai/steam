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
 * Created by justin on 7/10/16.
 */
var React = require('react');
var react_router_1 = require('react-router');
var Table_1 = require('../components/Table');
var Row_1 = require('../components/Row');
var Cell_1 = require('../components/Cell');
require('../styles/newproject.scss');
var NewProjectStep2 = (function (_super) {
    __extends(NewProjectStep2, _super);
    function NewProjectStep2() {
        _super.apply(this, arguments);
    }
    NewProjectStep2.prototype.render = function () {
        return (React.createElement("div", {className: "new-project"}, React.createElement("form", null, React.createElement("label", null, "Train Initial Model"), React.createElement("span", null, "Create a training frame, test frame, and start building models.")), React.createElement(Table_1.default, {className: "build-model"}, React.createElement(Row_1.default, {header: true}), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "SPLIT DATAFRAME"), React.createElement(Cell_1.default, null, React.createElement("div", {className: "dataframe-range"}, React.createElement("div", null, React.createElement("input", {type: "range"})), React.createElement("div", {className: "dataframe-range-labels"}, React.createElement("div", null, "Train: 75%"), React.createElement("div", null, "Test: 25%"))))), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "DEFAULT MODELS"), React.createElement(Cell_1.default, null, React.createElement("div", {className: "mode-checkboxes"}, React.createElement("span", null, React.createElement("input", {type: "checkbox"}), " Generalized Linear Model"), React.createElement("span", null, React.createElement("input", {type: "checkbox"}), " Gradient Boosting Machine"), React.createElement("span", null, React.createElement("input", {type: "checkbox"}), " Random Forest"), React.createElement("span", null, React.createElement("input", {type: "checkbox"}), " Deep Learning"), React.createElement("span", null, React.createElement("input", {type: "checkbox"}), " Naive Bayes")))), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "PICK A TRAINING CLUSTER"), React.createElement(Cell_1.default, null, React.createElement("select", null, React.createElement("option", null, "Test"))))), React.createElement(react_router_1.Link, {to: "/projects/new/3", className: "default"}, "Train Models")));
    };
    return NewProjectStep2;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = NewProjectStep2;
//# sourceMappingURL=NewProjectStep2.js.map