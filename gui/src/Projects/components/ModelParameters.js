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
 * Created by Jeff Fohl <jfohl@h2o.ai> on 7/12/16.
 */
var React = require('react');
var Table_1 = require('./Table');
var Row_1 = require('./Row');
var Cell_1 = require('./Cell');
var h2oUIKit_1 = require('h2oUIKit');
require('../styles/modelparameters.scss');
var ModelParameters = (function (_super) {
    __extends(ModelParameters, _super);
    function ModelParameters() {
        _super.apply(this, arguments);
    }
    ModelParameters.prototype.render = function () {
        return (React.createElement("div", {className: "model-parameters"}, React.createElement(Table_1.default, null, React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, React.createElement("span", null, "Number of Trees")), React.createElement(Cell_1.default, null, React.createElement(h2oUIKit_1.NumericInput, {name: "numberOfTrees"})), React.createElement(Cell_1.default, null, "Number of trees to train")), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, React.createElement("span", null, "Max Depth")), React.createElement(Cell_1.default, null, React.createElement(h2oUIKit_1.NumericInput, {name: "maxDepth"})), React.createElement(Cell_1.default, null, "Maximum depth for any number in a tree")), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, React.createElement("span", null, "Minimum Number of Rows")), React.createElement(Cell_1.default, null, React.createElement(h2oUIKit_1.NumericInput, {name: "minimumNumberOfRows"})), React.createElement(Cell_1.default, null, "Minimum number of rows in each leaf node")), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, React.createElement("span", null, "Learning Rate")), React.createElement(Cell_1.default, null, React.createElement(h2oUIKit_1.NumericInput, {name: "learningRate"})), React.createElement(Cell_1.default, null, "Learning rate"))), React.createElement("button", {className: "link"}, "see full parameters list")));
    };
    return ModelParameters;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = ModelParameters;
//# sourceMappingURL=ModelParameters.js.map