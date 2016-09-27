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
 * Created by Jeff Fohl <jfohl@h2o.ai> on 7/11/16.
 */
var React = require('react');
var PageHeader_1 = require('./PageHeader');
var Table_1 = require('./Table');
var Row_1 = require('./Row');
var Cell_1 = require('./Cell');
var ModelParameters_1 = require('./ModelParameters');
var react_router_1 = require('react-router');
require('../styles/createnew.scss');
var CreateNewModel = (function (_super) {
    __extends(CreateNewModel, _super);
    function CreateNewModel() {
        _super.call(this);
    }
    CreateNewModel.prototype.handleSubmit = function (e) {
        e.preventDefault();
    };
    CreateNewModel.prototype.render = function () {
        return (React.createElement("div", {className: "create-new-model"}, React.createElement(PageHeader_1.default, null, "Create New Model", React.createElement("div", {className: "subtitle"}, "Forked from ", React.createElement(react_router_1.Link, {to: "/projects/models/DRF-1069085"}, "DRF-1069085"))), React.createElement("form", {name: "new-model", onSubmit: this.handleSubmit}, React.createElement("section", {className: "data"}, React.createElement(Table_1.default, {className: "new-model"}, React.createElement(Row_1.default, {className: "sub-section"}, React.createElement(Cell_1.default, {className: "label"}, "Training Dataset", React.createElement("p", {className: "hint"}, "Dataset to use in training.")), React.createElement(Cell_1.default, {className: "value"}, React.createElement("select", {name: "training-dataset"}, React.createElement("option", {value: "telecom-churn-25-train"}, "telecom-churn-25-train"), React.createElement("option", {value: "telecom-churn-50-train"}, "telecom-churn-50-train"), React.createElement("option", {value: "telecom-churn-75-train"}, "telecom-churn-75-train"), React.createElement("option", {value: "telecom-churn-100-train"}, "telecom-churn-100-train")))), React.createElement(Row_1.default, {className: "sub-section"}, React.createElement(Cell_1.default, {className: "label"}, "Transformation Pipeline"), React.createElement(Cell_1.default, {className: "value"}, "transformers/transformation-pipeline.py ", React.createElement("button", {className: "link"}, "change"))), React.createElement(Row_1.default, {className: "sub-section"}, React.createElement(Cell_1.default, {className: "label"}, "Model Type", React.createElement("p", {className: "hint"}, "Select model type to train. Given this project's task, GBM's are recommended.")), React.createElement(Cell_1.default, {className: "value"}, React.createElement("select", {name: "model-type"}, React.createElement("option", {value: "gbm"}, "Gradient Boosting Machine"), React.createElement("option", {value: "rf"}, "Random Forest"), React.createElement("option", {value: "glm"}, "Generalized Linear Model"), React.createElement("option", {value: "svm"}, "Support Vector Machine"), React.createElement("option", {value: "dl"}, "Deep Learning"), React.createElement("option", {value: "nb"}, "Naive Bayes")))), React.createElement(Row_1.default, {className: "sub-section"}, React.createElement(Cell_1.default, {className: "label"}, "GBM Parameters", React.createElement("p", {className: "hint"}, "Set training parameters for model training."), React.createElement("p", {className: "hint"}, "Not sure what parameters to use? Try grid search to test a range of parameters at once."), React.createElement("p", null, React.createElement("button", {className: "link"}, "switch to grid search"))), React.createElement(Cell_1.default, {className: "value"}, React.createElement(ModelParameters_1.default, null))), React.createElement(Row_1.default, {className: "sub-section"}, React.createElement(Cell_1.default, {className: "label"}, "Cluster"), React.createElement(Cell_1.default, {className: "value"}, React.createElement("select", {name: "cluster"}, React.createElement("option", {value: "prithvi-2"}, "Prithvi - 2 nodes"), React.createElement("option", {value: "prithvi-4"}, "Prithvi - 4 nodes"), React.createElement("option", {value: "prithvi-8"}, "Prithvi - 8 nodes"), React.createElement("option", {value: "prithvi-16"}, "Prithvi - 16 nodes"), React.createElement("option", {value: "prithvi-32"}, "Prithvi - 32 nodes"), React.createElement("option", {value: "prithvi-64"}, "Prithvi - 64 nodes"), React.createElement("option", {value: "prithvi-128"}, "Prithvi - 128 nodes")))))), React.createElement("section", {className: "actions"}, React.createElement("button", {className: "default"}, "Beginning Model Training"), React.createElement("div", {className: "optional"}, React.createElement("div", {className: "checkbox"}, React.createElement("input", {type: "checkbox", name: "save-as-script"})), React.createElement("div", {className: "optional-label"}, React.createElement("p", null, "Save modeling procedures as a script."), React.createElement("p", null, "This is helpful for scheduled re-training of models.")))))));
    };
    return CreateNewModel;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = CreateNewModel;
//# sourceMappingURL=CreateNewModel.js.map