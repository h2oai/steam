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
var _ = require('lodash');
var classNames = require('classnames');
var DetailLine_1 = require('./DetailLine');
var RocGraph_1 = require('../../Models/components/RocGraph');
require('../styles/goodnessoffit.scss');
var colors_1 = require('../../App/utils/colors');
var GoodnessOfFit = (function (_super) {
    __extends(GoodnessOfFit, _super);
    function GoodnessOfFit() {
        _super.apply(this, arguments);
    }
    GoodnessOfFit.prototype.render = function () {
        var modelMetrics = JSON.parse(this.props.model.json_metrics);
        var comparisonModelMetrics = this.props.comparisonModel ? JSON.parse(this.props.comparisonModel.json_metrics) : null;
        var trainingMetrics = _.get(modelMetrics, 'models[0].output.training_metrics', {});
        var comparisonTrainingMetrics = _.get(comparisonModelMetrics, 'models[0].output.training_metrics', {});
        var metrics = {
            mse: {
                label: 'Mean Squared Error',
                value: _.get(this.props.model, 'mse', null) ? _.get(this.props.model, 'mse', null).toFixed(6) : null,
                comparisonValue: _.get(this.props.comparisonModel, 'mse', null) ? _.get(this.props.comparisonModel, 'mse', null).toFixed(6) : null
            },
            logloss: {
                label: 'LogLoss',
                value: _.get(this.props.model, 'logloss', null) ? _.get(this.props.model, 'logloss', null).toFixed(6) : null,
                comparisonValue: _.get(this.props.comparisonModel, 'logloss', null) ? _.get(this.props.comparisonModel, 'logloss', null).toFixed(6) : null
            },
            r_squared: {
                label: React.createElement("span", null, "R", React.createElement("sup", null, "2")),
                value: _.get(this.props.model, 'r_squared', null) ? _.get(this.props.model, 'r_squared', null).toFixed(6) : null,
                comparisonValue: _.get(this.props.comparisonModel, 'r_squared', null) ? _.get(this.props.comparisonModel, 'r_squared', null).toFixed(6) : null
            },
            auc: {
                label: 'AUC',
                value: _.get(this.props.model, 'auc', null) ? _.get(this.props.model, 'auc', null).toFixed(6) : null,
                comparisonValue: _.get(this.props.comparisonModel, 'auc', null) ? _.get(this.props.comparisonModel, 'auc', null).toFixed(6) : null
            },
            gini: {
                label: 'Gini',
                value: _.get(trainingMetrics, 'Gini', null) ? _.get(trainingMetrics, 'Gini', null).toFixed(6) : null,
                comparisonValue: _.get(comparisonTrainingMetrics, 'Gini', null) ? _.get(comparisonTrainingMetrics, 'Gini', null).toFixed(6) : null
            }
        };
        var fpr = _.get(modelMetrics, 'models[0].output.training_metrics.thresholds_and_metric_scores.data[17]', []);
        var tpr = _.get(modelMetrics, 'models[0].output.training_metrics.thresholds_and_metric_scores.data[18]', []);
        var comparisonTpr = _.get(comparisonModelMetrics, 'models[0].output.training_metrics.thresholds_and_metric_scores.data[18]', []);
        var data = [];
        var modelCurveData = {
            name: this.props.model.name,
            values: []
        };
        tpr.map(function (val, i) {
            var newEntry = {};
            newEntry.tpr = val;
            newEntry.fpr = fpr[i];
            modelCurveData.values.push(newEntry);
        });
        data.push(modelCurveData);
        if (this.props.comparisonModel) {
            var comparisonCurveData_1 = {
                name: this.props.comparisonModel.name,
                values: []
            };
            comparisonTpr.map(function (val, i) {
                var newEntry = {};
                newEntry.tpr = val;
                newEntry.fpr = fpr[i];
                comparisonCurveData_1.values.push(newEntry);
            });
            data.push(comparisonCurveData_1);
        }
        var config = {
            margin: { top: 2, right: 2, bottom: 2, left: 2 },
            interpolationMode: 'basis',
            height: '100%',
            width: '100%',
            smooth: true,
            fpr: 'fpr',
            tprVariables: [
                {
                    name: 'tpr',
                    label: 'tpr'
                },
                {
                    name: 'comparisonTpr',
                    label: 'comparisonTpr'
                }
            ],
            animate: undefined,
            hideAxes: true,
            hideAUCText: true,
            curveColors: [colors_1.BRAND_BLUE, colors_1.BRAND_ORANGE]
        };
        return (React.createElement("div", {className: "metrics goodness-of-fit"}, React.createElement("div", {className: "metrics-summary"}, React.createElement("div", {className: "metrics-summary--title"}, "Metrics"), Object.keys(metrics).map(function (key) {
            if (metrics[key].value) {
                return React.createElement(DetailLine_1.default, {className: classNames({ compare: metrics[key].comparisonValue }), key: key, label: metrics[key].label, value: metrics[key].value, comparisonValue: metrics[key].comparisonValue});
            }
            return null;
        })), this.props.modelCategory === 'binomial' ? React.createElement("div", {className: "roc-chart"}, React.createElement(RocGraph_1.default, {data: data, config: config})) : null));
    };
    return GoodnessOfFit;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = GoodnessOfFit;
//# sourceMappingURL=GoodnessOfFit.js.map