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
 * Created by justin on 8/6/16.
 */
var React = require('react');
var DefaultModal_1 = require('../../App/components/DefaultModal');
var PageHeader_1 = require('../../Projects/components/PageHeader');
var Table_1 = require('../../Projects/components/Table');
var Row_1 = require('../../Projects/components/Row');
var Cell_1 = require('../../Projects/components/Cell');
var redux_1 = require('redux');
var projects_actions_1 = require('../../Projects/actions/projects.actions');
var react_redux_1 = require('react-redux');
require('../styles/importmodelsmodal.scss');
var ImportModelsModal = (function (_super) {
    __extends(ImportModelsModal, _super);
    function ImportModelsModal() {
        _super.call(this);
        this.state = {
            clusterId: null,
            models: []
        };
    }
    ImportModelsModal.prototype.componentWillMount = function () {
        this.props.fetchClusters(parseInt(this.props.projectId, 10));
    };
    ImportModelsModal.prototype.onChange = function (event) {
        this.setState({
            clusterId: event.target.value
        });
        this.props.fetchModelsFromCluster(parseInt(event.target.value, 10), this.props.datasetName);
    };
    ImportModelsModal.prototype.importModelsFromCluster = function (event) {
        var _this = this;
        event.preventDefault();
        var inputs = event.target.querySelectorAll('input:checked');
        var models = [];
        for (var i = 0; i < inputs.length; i++) {
            models.push(inputs[i].value);
        }
        this.props.importModelsFromCluster(parseInt(this.state.clusterId, 10), parseInt(this.props.projectId, 10), models).then(function () {
            _this.props.fetchLeaderboard(parseInt(_this.props.projectId, 10), _this.props.modelCategory);
            _this.props.onCancel();
        });
    };
    ImportModelsModal.prototype.render = function () {
        return (React.createElement(DefaultModal_1.default, {className: "import-modal", open: this.props.open}, React.createElement(PageHeader_1.default, null, "IMPORT MODELS"), React.createElement("form", {onSubmit: this.importModelsFromCluster.bind(this)}, React.createElement(Table_1.default, {className: "outer-table"}, React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "CLUSTER"), React.createElement(Cell_1.default, null, React.createElement("div", null, "Select the cluster to import models from"), React.createElement("select", {onChange: this.onChange.bind(this)}, React.createElement("option", {value: ""}), this.props.clusters.map(function (cluster, i) {
            return React.createElement("option", {key: i, value: cluster.id}, cluster.name, " @ ", cluster.address);
        })))), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "SELECT MODELS"), React.createElement(Cell_1.default, null, React.createElement(Table_1.default, {className: "inner-table"}, React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "MODEL"), React.createElement(Cell_1.default, null, "DATAFRAME"), React.createElement(Cell_1.default, null, "RESPONSE COLUMN"), React.createElement(Cell_1.default, null)), this.props.models.map(function (model, i) {
            return (React.createElement(Row_1.default, {key: i}, React.createElement(Cell_1.default, null, model.name), React.createElement(Cell_1.default, null, model.dataset_name), React.createElement(Cell_1.default, null, model.response_column_name), React.createElement(Cell_1.default, null, React.createElement("input", {type: "checkbox", value: model.model_key}))));
        })))), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null), React.createElement(Cell_1.default, {className: "button-container"}, React.createElement("button", {type: "submit", className: "button-primary"}, "Import"), React.createElement("button", {type: "button", onClick: this.props.onCancel.bind(this), className: "button-secondary"}, "Cancel")))))));
    };
    return ImportModelsModal;
}(React.Component));
exports.ImportModelsModal = ImportModelsModal;
function mapStateToProps(state) {
    return {
        clusters: state.projects.clusters,
        models: state.projects.models
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchClusters: redux_1.bindActionCreators(projects_actions_1.fetchClusters, dispatch),
        fetchModelsFromCluster: redux_1.bindActionCreators(projects_actions_1.fetchModelsFromCluster, dispatch),
        importModelsFromCluster: redux_1.bindActionCreators(projects_actions_1.importModelsFromCluster, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(ImportModelsModal);
//# sourceMappingURL=ImportModelsModal.js.map