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
 * Created by justin on 7/18/16.
 */
var React = require('react');
var $ = require('jquery');
var _ = require('lodash');
var classNames = require('classnames');
var redux_1 = require('redux');
var react_redux_1 = require('react-redux');
var Table_1 = require('./Table');
var Row_1 = require('./Row');
var Cell_1 = require('./Cell');
var projects_actions_1 = require('../actions/projects.actions');
require('../styles/importnewproject.scss');
var react_router_1 = require('react-router');
var InputFeedback_1 = require('../../App/components/InputFeedback');
var InputFeedback_2 = require('../../App/components/InputFeedback');
var ImportNewProject = (function (_super) {
    __extends(ImportNewProject, _super);
    function ImportNewProject() {
        _super.call(this);
        this.state = {
            clusterId: null,
            datasetId: null,
            modelCategory: null,
            isModelSelected: false
        };
    }
    ImportNewProject.prototype.componentWillMount = function () {
        if (_.isEmpty(this.props.clusters)) {
            this.props.fetchClusters();
        }
    };
    ImportNewProject.prototype.selectDataset = function (event) {
        this.setState({
            datasetId: event.target.value
        });
        if (event.target.value) {
            this.props.fetchModelsFromCluster(this.state.clusterId, event.target.value);
        }
        else {
            this.setState({
                modelCategory: null
            });
        }
    };
    ImportNewProject.prototype.createProject = function () {
        var name = $(this.refs.projectName).val();
        var importModels = [];
        var checkedModels = $('.import-models input:checked');
        if (checkedModels.length > 0) {
            checkedModels.map(function (i, input) {
                importModels.push($(input).prop('name'));
            });
            this.props.createProjectAndImportModelsFromCluster(name, this.state.clusterId, this.state.modelCategory, importModels).then(function (res) {
                react_router_1.hashHistory.push('/projects/' + res + '/models');
            });
        }
    };
    ImportNewProject.prototype.registerCluster = function (event) {
        event.preventDefault();
        var ipAddress = $(event.target).find('input[name="ip-address"]').val();
        var port = $(event.target).find('input[name="port"]').val();
        this.props.registerCluster(ipAddress + ':' + port);
    };
    ImportNewProject.prototype.resetClusterSelection = function (event) {
        event.preventDefault();
        this.setState({
            clusterId: null
        });
        this.props.resetClusterSelection();
    };
    ImportNewProject.prototype.selectModel = function () {
        var checkedModels = $('.import-models input:checked');
        if (checkedModels) {
            this.setState({
                isModelSelected: true
            });
        }
        else {
            this.setState({
                isModelSelected: false
            });
        }
    };
    ImportNewProject.prototype.selectCategory = function (event) {
        this.setState({
            modelCategory: event.target.value
        });
    };
    ImportNewProject.prototype.retrieveClusterDataframes = function (clusterId) {
        this.setState({
            clusterId: clusterId
        });
        this.props.fetchDatasetsFromCluster(clusterId);
    };
    ImportNewProject.prototype.render = function () {
        var _this = this;
        var selectedClusterName;
        var selectedClusterAddress;
        for (var _i = 0, _a = this.props.clusters; _i < _a.length; _i++) {
            var cluster = _a[_i];
            if (cluster.id === this.state.clusterId) {
                selectedClusterName = cluster.name;
                selectedClusterAddress = cluster.address;
            }
        }
        if (!this.props.clusters) {
            return React.createElement("div", null);
        }
        return (React.createElement("div", {className: "import-new-project"}, React.createElement("div", {className: "step-1"}, React.createElement("div", {className: "select-cluster"}, React.createElement("h2", null, "1. Select H2O cluster"), this.state.clusterId ?
            React.createElement("div", {className: "cluster-info intro"}, React.createElement("span", null, React.createElement("i", {className: "fa fa-cubes cluster-image"})), React.createElement("div", {className: "cluster-details"}, React.createElement("div", null, selectedClusterName), React.createElement("div", null, selectedClusterAddress), React.createElement("div", {onClick: this.resetClusterSelection.bind(this), className: "select-new-cluster"}, React.createElement("i", {className: "fa fa-close"}), " use a different cluster"))) :
            React.createElement("div", {className: "intro"}, "Select an H2O cluster to import models and datasets from.", React.createElement(Table_1.default, null, React.createElement(Row_1.default, {header: true}, React.createElement(Cell_1.default, null, "CLUSTER"), React.createElement(Cell_1.default, null, "DATASETS"), React.createElement(Cell_1.default, null, "MODELS"), React.createElement(Cell_1.default, null)), this.props.clusters.map(function (cluster, i) {
                return (React.createElement(Row_1.default, {key: i}, React.createElement(Cell_1.default, null, React.createElement("span", {className: "name-cell"}, cluster.name)), React.createElement(Cell_1.default, null, "N/A"), React.createElement(Cell_1.default, null, "N/A"), React.createElement(Cell_1.default, null, React.createElement("button", {className: "button-primary", onClick: _this.retrieveClusterDataframes.bind(_this, cluster.id)}, "Connect"))));
            })))), !this.state.clusterId ?
            React.createElement("div", {className: "connect-cluster"}, React.createElement("h2", null, "… or connect to a new H2O cluster"), React.createElement("div", {className: "intro"}, "Connect to a H2O cluster where your existing models and data sets are located."), React.createElement("form", {onSubmit: this.registerCluster.bind(this)}, React.createElement("input", {type: "text", name: "ip-address", placeholder: "IP Address"}), this.props.registerClusterError ?
                React.createElement(InputFeedback_1.default, {message: this.props.registerClusterError, type: InputFeedback_2.FeedbackType.Error})
                : null, React.createElement("input", {type: "text", name: "port", placeholder: "Port"}), React.createElement("button", {type: "submit", className: "button-primary"}, "Connect")), this.props.isClusterFetchInProcess ?
                React.createElement(InputFeedback_1.default, {message: "Connecting...", type: InputFeedback_2.FeedbackType.Progress})
                : null)
            : null), this.state.clusterId ? React.createElement("div", null, React.createElement("h2", null, "2. Select Dataframe"), React.createElement("div", {className: "intro"}, React.createElement("select", {name: "selectDataframe", onChange: this.selectDataset.bind(this)}, React.createElement("option", null), this.props.datasets ? this.props.datasets.map(function (dataset, i) {
            return React.createElement("option", {key: i, value: dataset.frame_name}, dataset.name);
        }) : null), this.props.isModelFetchInProcess ?
            React.createElement(InputFeedback_1.default, {message: "Connecting...", type: InputFeedback_2.FeedbackType.Progress})
            : null)) : null, (this.state.datasetId && !this.props.isModelFetchInProcess) ?
            React.createElement("div", null, React.createElement("h2", null, "3. Select Model Category"), React.createElement("div", {className: "intro"}, React.createElement("select", {name: "selectModelCategory", onChange: this.selectCategory.bind(this)}, React.createElement("option", null), this.props.models ? _.uniqBy(this.props.models, 'model_category').map(function (model, i) {
                return React.createElement("option", {key: i, value: model.model_category}, model.model_category);
            }) : null))) : null, this.state.datasetId && !_.isEmpty(this.props.models) && this.state.modelCategory ? React.createElement("div", null, React.createElement("h2", null, "4. Pick Models to Import"), React.createElement("div", {className: "intro"}, "Models in a project must share the same feature set and response column to enable comparison."), React.createElement(Table_1.default, {className: "import-models"}, React.createElement(Row_1.default, {header: true}, React.createElement(Cell_1.default, null, "MODEL"), React.createElement(Cell_1.default, null, "RESPONSE COLUMN"), React.createElement(Cell_1.default, null, "CATEGORICAL"), React.createElement(Cell_1.default, null)), _.filter(this.props.models, function (model) { return model.model_category === _this.state.modelCategory; }).map(function (model, i) {
            return (React.createElement(Row_1.default, {key: i}, React.createElement(Cell_1.default, null, model.name), React.createElement(Cell_1.default, null, model.response_column_name), React.createElement(Cell_1.default, null, model.model_category), React.createElement(Cell_1.default, null, React.createElement("input", {type: "checkbox", name: model.name, onChange: _this.selectModel.bind(_this, model)}), "  Select for Import")));
        }))) : null, this.state.datasetId && !_.isEmpty(this.props.models && this.state.modelCategory) ? React.createElement("div", {className: "name-project"}, React.createElement("h2", null, "5. Name Project"), React.createElement("div", {className: "intro"}, React.createElement("input", {ref: "projectName", type: "text"}))) : null, this.state.datasetId && !_.isEmpty(this.props.models) && this.state.modelCategory ? React.createElement("div", null, React.createElement("button", {className: classNames('button-primary', { disabled: !this.state.isModelSelected }), onClick: this.createProject.bind(this)}, "Create Project")) : null));
    };
    return ImportNewProject;
}(React.Component));
exports.ImportNewProject = ImportNewProject;
function mapStateToProps(state) {
    return {
        clusters: state.projects.clusters,
        models: state.projects.models,
        datasets: state.projects.datasets,
        project: state.project,
        isClusterFetchInProcess: state.projects.isClusterFetchInProcess,
        isModelFetchInProcess: state.projects.isModelFetchInProcess,
        registerClusterError: state.projects.registerClusterError
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchClusters: redux_1.bindActionCreators(projects_actions_1.fetchClusters, dispatch),
        fetchModelsFromCluster: redux_1.bindActionCreators(projects_actions_1.fetchModelsFromCluster, dispatch),
        createProjectAndImportModelsFromCluster: redux_1.bindActionCreators(projects_actions_1.createProjectAndImportModelsFromCluster, dispatch),
        importModelFromCluster: redux_1.bindActionCreators(projects_actions_1.importModelFromCluster, dispatch),
        registerCluster: redux_1.bindActionCreators(projects_actions_1.registerCluster, dispatch),
        fetchDatasetsFromCluster: redux_1.bindActionCreators(projects_actions_1.fetchDatasetsFromCluster, dispatch),
        resetClusterSelection: redux_1.bindActionCreators(projects_actions_1.resetClusterSelection, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(ImportNewProject);
//# sourceMappingURL=ImportNewProject.js.map