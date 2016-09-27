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
 * Created by justin on 6/27/16.
 */
var React = require('react');
var moment = require('moment');
var DetailLine_1 = require('./DetailLine');
require('../styles/modeloverview.scss');
var ModelOverview = (function (_super) {
    __extends(ModelOverview, _super);
    function ModelOverview() {
        _super.apply(this, arguments);
    }
    ModelOverview.prototype.render = function () {
        var iconMap = {
            Author: 'fa fa-user',
            Date: 'fa fa-calendar-o',
            Size: 'fa fa-save',
            'Training Time': 'fa fa-hourglass-2',
            'Classification Speed': 'fa fa-clock-o',
            'Model Type': 'fa fa-cube'
        };
        var basics = [
            {
                label: 'Date',
                value: moment.unix(this.props.model.created_at).format('YYYY-MM-DD HH:mm')
            },
            {
                label: 'Model Type',
                value: this.props.model.algorithm
            }
        ];
        var parameters = [
            {
                label: 'Dataset Name',
                value: this.props.model.dataset_name
            },
            {
                label: 'Response Column Name',
                value: this.props.model.response_column_name
            }
        ];
        return (React.createElement("div", {className: "metrics"}, React.createElement("div", {className: "metrics-summary"}, React.createElement("div", {className: "metrics-summary--title"}, "Basics"), basics.map(function (item, i) {
            return React.createElement(DetailLine_1.default, {key: i, icon: iconMap[item.label], label: item.label, value: item.value});
        })), React.createElement("div", {className: "metrics-summary"}, React.createElement("div", {className: "metrics-summary--title"}, "Model Parameters"), parameters.map(function (item, i) {
            return React.createElement(DetailLine_1.default, {key: i, label: item.label, value: item.value});
        }))));
    };
    return ModelOverview;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = ModelOverview;
//# sourceMappingURL=ModelOverview.js.map