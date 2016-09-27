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
 * Created by justin on 8/4/16.
 */
var React = require('react');
var moment = require('moment');
var _ = require('lodash');
var Table_1 = require('../../Projects/components/Table');
var Row_1 = require('../../Projects/components/Row');
var Cell_1 = require('../../Projects/components/Cell');
var FilterDropdown_1 = require('./FilterDropdown');
var RocGraph_1 = require('./RocGraph');
var ModelLabelSelect_1 = require('./ModelLabelSelect');
var react_router_1 = require('react-router');
var BinomialModelTable = (function (_super) {
    __extends(BinomialModelTable, _super);
    function BinomialModelTable() {
        _super.apply(this, arguments);
    }
    BinomialModelTable.prototype.render = function () {
        var _this = this;
        return (React.createElement(Table_1.default, null, React.createElement(Row_1.default, {header: true}, React.createElement(Cell_1.default, null, React.createElement(FilterDropdown_1.default, {onFilter: this.props.onFilter.bind(this), sortCriteria: this.props.sortCriteria})), React.createElement(Cell_1.default, null, "MODEL"), React.createElement(Cell_1.default, null, "AUC"), React.createElement(Cell_1.default, null, "Gini"), React.createElement(Cell_1.default, null, "MSE"), React.createElement(Cell_1.default, null, "Logloss"), React.createElement(Cell_1.default, {className: "graph"}, "ROC"), React.createElement(Cell_1.default, null, React.createElement("div", {className: "actions"}, "ACTIONS"))), this.props.items.map(function (model, i) {
            var modelMetrics = JSON.parse(model.json_metrics);
            var trainingMetrics = _.get(modelMetrics, 'models[0].output.training_metrics', {});
            var fpr = _.get(modelMetrics, 'models[0].output.training_metrics.thresholds_and_metric_scores.data[17]', []);
            var tpr = _.get(modelMetrics, 'models[0].output.training_metrics.thresholds_and_metric_scores.data[18]', []);
            var data = [
                {
                    name: model.name,
                    values: []
                }
            ];
            tpr.map(function (val, i) {
                data[0].values.push({
                    tpr: val,
                    fpr: fpr[i]
                });
            });
            return (React.createElement(Row_1.default, {key: i}, React.createElement(Cell_1.default, null), React.createElement(Cell_1.default, null, React.createElement("div", {className: "metadata"}, React.createElement("div", {className: "model-name"}, model.name), React.createElement("div", null, React.createElement("span", null, "Created at: "), React.createElement("span", null, moment.unix(model.created_at).format('YYYY-MM-DD hh:mm:ss'))), React.createElement("div", null, React.createElement("span", null, "Num of Observations: "), React.createElement("span", null, trainingMetrics.nobs)), React.createElement("div", null, React.createElement("span", null, "Cluster: "), React.createElement("span", null, model.cluster_name)))), React.createElement(Cell_1.default, {name: "auc"}, trainingMetrics.AUC ? trainingMetrics.AUC.toFixed(6) : 'N/A'), React.createElement(Cell_1.default, {name: "gini"}, trainingMetrics.Gini ? trainingMetrics.Gini.toFixed(6) : 'N/A'), React.createElement(Cell_1.default, {name: "mse"}, trainingMetrics.MSE ? trainingMetrics.MSE.toFixed(6) : 'N/A'), React.createElement(Cell_1.default, {name: "logloss"}, trainingMetrics.logloss ? trainingMetrics.logloss.toFixed(6) : 'N/A'), React.createElement(Cell_1.default, {className: "graph", name: "roc"}, React.createElement(RocGraph_1.default, {data: data})), React.createElement(Cell_1.default, null, React.createElement("ul", {className: "actions"}, React.createElement("li", null, React.createElement(react_router_1.Link, {to: '/projects/' + _this.props.projectId + '/models/' + model.id}, React.createElement("span", null, React.createElement("i", {className: "fa fa-eye"})), React.createElement("span", null, "view model details"))), React.createElement("li", {className: "labels"}, React.createElement("span", null, React.createElement("i", {className: "fa fa-tags"})), " label as", React.createElement("span", {className: "label-selector"}, React.createElement(ModelLabelSelect_1.default, {projectId: _this.props.projectId, modelId: model.id, labels: _this.props.labels, onChangeHandler: _this.props.onChangeHandler}))), React.createElement("li", {onClick: _this.props.openDeploy.bind(_this, model)}, React.createElement("span", null, React.createElement("i", {className: "fa fa-arrow-up"})), React.createElement("span", null, "deploy model"))))));
        }, this)));
    };
    return BinomialModelTable;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = BinomialModelTable;
//# sourceMappingURL=BinomialModelTable.js.map