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
 * Created by Jeff Fohl <jfohl@h2o.ai> on 6/29/16.
 */
var React = require('react');
var _ = require('lodash');
var ModelLabelSelect = (function (_super) {
    __extends(ModelLabelSelect, _super);
    function ModelLabelSelect() {
        _super.apply(this, arguments);
    }
    ModelLabelSelect.prototype.onChangeHandler = function (event) {
        var value = parseInt(event.target.value, 10);
        if (value === -1) {
            this.props.onChangeHandler(_.find(this.props.labels[this.props.projectId], { model_id: this.props.modelId }).id, this.props.modelId, true);
        }
        else {
            this.props.onChangeHandler(parseInt(event.target.value, 10), this.props.modelId, false);
        }
    };
    ModelLabelSelect.prototype.render = function () {
        if (_.isUndefined(this.props.labels[this.props.projectId])) {
            return React.createElement("select", {name: "labelSelect"}, React.createElement("option", {value: -1}));
        }
        return (React.createElement("select", {name: "labelSelect", onChange: this.onChangeHandler.bind(this), value: _.find(this.props.labels[this.props.projectId], { model_id: this.props.modelId }) ? _.find(this.props.labels[this.props.projectId], { model_id: this.props.modelId }).id : -1}, React.createElement("option", {value: -1}), this.props.labels[this.props.projectId] ? this.props.labels[this.props.projectId].map(function (label) {
            return (React.createElement("option", {key: label.id, value: label.id}, label.name));
        }) : React.createElement("option", null)));
    };
    return ModelLabelSelect;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = ModelLabelSelect;
//# sourceMappingURL=ModelLabelSelect.js.map