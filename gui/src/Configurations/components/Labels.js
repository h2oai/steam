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
 * Created by justin on 7/12/16.
 */
var React = require('react');
var Table_1 = require('../../Projects/components/Table');
var Row_1 = require('../../Projects/components/Row');
var Cell_1 = require('../../Projects/components/Cell');
var CreateNewLabelModal_1 = require('./CreateNewLabelModal');
var configuration_labels_action_1 = require('../actions/configuration.labels.action');
var redux_1 = require('redux');
var react_redux_1 = require('react-redux');
require('../styles/labels.scss');
var Labels = (function (_super) {
    __extends(Labels, _super);
    function Labels() {
        _super.call(this);
        this.state = {
            modalOpen: false,
            label: {}
        };
    }
    Labels.prototype.componentWillMount = function () {
        if (!this.props.labels || !this.props.labels[this.props.projectid]) {
            this.props.fetchLabels(this.props.projectid);
        }
    };
    Labels.prototype.openModal = function () {
        this.setState({
            modalOpen: true
        });
    };
    Labels.prototype.closeModal = function () {
        this.setState({
            modalOpen: false,
            label: {}
        });
    };
    Labels.prototype.saveUpdateLabel = function (label) {
        if (!label.id) {
            var newLabel = {
                name: label.name,
                description: label.description
            };
            this.saveLabel(label);
        }
        else {
            this.updateLabel(label);
        }
    };
    Labels.prototype.updateLabel = function (label) {
        var _this = this;
        this.props.updateLabel(parseInt(label.id, 10), this.props.projectid, label.name, label.description).then(function (response) {
            _this.props.fetchLabels(_this.props.projectid);
            _this.closeModal();
        }, function (error) {
            alert(error);
        });
    };
    Labels.prototype.saveLabel = function (label) {
        var _this = this;
        this.props.createLabel(parseInt(this.props.projectid, 10), label.name, label.description).then(function (response) {
            _this.props.fetchLabels(_this.props.projectid);
            _this.closeModal();
        }, function (error) {
            alert(error);
        });
    };
    Labels.prototype.deleteLabel = function (labelId) {
        var _this = this;
        this.props.deleteLabel(labelId).then(function (response) {
            _this.props.fetchLabels(_this.props.projectid);
        }, function (error) {
            alert(error);
        });
    };
    Labels.prototype.renderLabels = function () {
        var _this = this;
        if (!this.props.labels || !this.props.projectid || !this.props.labels[this.props.projectid]) {
            return null;
        }
        return this.props.labels[this.props.projectid].map(function (label) {
            var deleteLabel = function () {
                _this.deleteLabel(label.id);
            };
            var updateLabel = function (event) {
                _this.setState({
                    label: {
                        id: label.id,
                        name: label.name,
                        description: label.description
                    },
                    modalOpen: true
                });
            };
            return (React.createElement(Row_1.default, {key: label.id}, React.createElement(Cell_1.default, {className: "label-bullets"}, React.createElement("span", {className: "label-bullet"})), React.createElement(Cell_1.default, {className: "label-names"}, React.createElement("div", {className: "label-name"}, label.name), React.createElement("div", {className: "label-description muted"}, label.description)), React.createElement(Cell_1.default, {className: "label-model"}, React.createElement("span", {className: "model-icon"}), React.createElement("span", {className: "model-name"}, (label.model_id >= 0) ? (React.createElement("span", {className: "fa fa-cube"})) : null, " ", (label.model_id >= 0) ? label.model_id : "Not currently applied to a model")), React.createElement(Cell_1.default, {className: "label-permissions"}), React.createElement(Cell_1.default, null, React.createElement("span", {className: "fa fa-pencil", onClick: updateLabel}), React.createElement("span", {className: "fa fa-trash", onClick: deleteLabel}))));
        });
    };
    Labels.prototype.render = function () {
        return (React.createElement("div", {className: "labels"}, React.createElement("h1", null, "Model Labels"), React.createElement("p", {className: "lede"}, "You can create labels for the models. A label can only be associated" + ' ' + "with one model at a time. You can give different team members various" + ' ' + "permissions."), React.createElement("p", {className: "lede"}, "For example, you can create \"test\" and \"production\" labels. You could" + ' ' + "then allow the entire team to label a model \"test\", but only give admins" + ' ' + "the power to label a model \"production\"."), React.createElement("span", null, React.createElement("button", {className: "button-primary", onClick: this.openModal.bind(this)}, "Create New Label")), React.createElement(CreateNewLabelModal_1.default, {label: this.state.label, open: this.state.modalOpen, cancel: this.closeModal.bind(this), save: this.saveUpdateLabel.bind(this)}), React.createElement("div", {className: "label-table"}, React.createElement(Table_1.default, null, React.createElement(Row_1.default, {className: "head"}, React.createElement(Cell_1.default, null), React.createElement(Cell_1.default, null, "Label"), React.createElement(Cell_1.default, null, "Model"), React.createElement(Cell_1.default, null, "Permissions"), React.createElement(Cell_1.default, null)), this.renderLabels()))));
    };
    return Labels;
}(React.Component));
function mapStateToProps(state) {
    return {
        labels: state.labels
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchLabels: redux_1.bindActionCreators(configuration_labels_action_1.fetchLabels, dispatch),
        createLabel: redux_1.bindActionCreators(configuration_labels_action_1.createLabel, dispatch),
        deleteLabel: redux_1.bindActionCreators(configuration_labels_action_1.deleteLabel, dispatch),
        updateLabel: redux_1.bindActionCreators(configuration_labels_action_1.updateLabel, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(Labels);
//# sourceMappingURL=Labels.js.map