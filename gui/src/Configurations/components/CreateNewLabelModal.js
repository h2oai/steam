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
 * Created by justin on 7/27/16.
 */
var React = require('react');
var DefaultModal_1 = require('../../App/components/DefaultModal');
var Table_1 = require('../../Projects/components/Table');
var Row_1 = require('../../Projects/components/Row');
var Cell_1 = require('../../Projects/components/Cell');
require('../styles/createnewlabelmodal.scss');
var initialState = {
    id: false,
    name: '',
    description: ''
};
var CreateNewLabelModal = (function (_super) {
    __extends(CreateNewLabelModal, _super);
    function CreateNewLabelModal() {
        _super.call(this);
        this.state = initialState;
    }
    CreateNewLabelModal.prototype.componentWillReceiveProps = function (nextProps) {
        if (nextProps.label.id) {
            this.setState({
                id: nextProps.label.id,
                name: nextProps.label.name,
                description: nextProps.label.description
            });
        }
    };
    CreateNewLabelModal.prototype.updateState = function (event) {
        var newState = {};
        newState[event.currentTarget.name] = event.currentTarget.value;
        this.setState(newState);
    };
    CreateNewLabelModal.prototype.cancel = function () {
        this.setState(initialState);
        this.props.cancel();
    };
    CreateNewLabelModal.prototype.save = function () {
        this.props.save(this.state);
        this.setState(initialState);
    };
    CreateNewLabelModal.prototype.render = function () {
        return (React.createElement(DefaultModal_1.default, {open: this.props.open, closeHandler: this.props.cancel}, React.createElement("div", {className: "create-edit-label-modal"}, React.createElement("header", {className: "page-header"}, "Create / Edit Label"), React.createElement("section", null, React.createElement(Table_1.default, null, React.createElement(Row_1.default, null, React.createElement(Cell_1.default, {className: "table-row-name"}, "Label Info"), React.createElement(Cell_1.default, {className: "table-row-item"}, React.createElement("p", null, "Enter a name and description of your label."), React.createElement("p", {className: "muted"}, "You can use this label in the project for exactly 1 model."), React.createElement("div", {className: "form-group"}, React.createElement("div", {className: "form-item"}, React.createElement("label", {className: "muted", htmlFor: "name"}, "Label name"), React.createElement("input", {type: "text", name: "name", value: this.state.name, onChange: this.updateState.bind(this)})), React.createElement("div", {className: "form-item"}, React.createElement("label", {className: "muted", htmlFor: "description"}, "Label description"), React.createElement("textarea", {name: "description", value: this.state.description, rows: "4", cols: "50", onChange: this.updateState.bind(this)}))))), React.createElement(Row_1.default, {className: "button-row"}, React.createElement(Cell_1.default, {className: "table-row-name"}), React.createElement(Cell_1.default, {className: "table-row-item"}, React.createElement("button", {className: "button-primary", onClick: this.save.bind(this)}, "Save"), React.createElement("button", {className: "button-secondary", onClick: this.cancel.bind(this)}, "Cancel"))))))));
    };
    return CreateNewLabelModal;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = CreateNewLabelModal;
//# sourceMappingURL=CreateNewLabelModal.js.map