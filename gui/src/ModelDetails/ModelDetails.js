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
var classNames = require('classnames');
var _ = require('lodash');
var Collapsible_1 = require('./components/Collapsible');
var ModelOverview_1 = require('./components/ModelOverview');
var GoodnessOfFit_1 = require('./components/GoodnessOfFit');
var PageHeader_1 = require('../Projects/components/PageHeader');
var ExportModal_1 = require('./components/ExportModal');
var Deploy_1 = require('../Models/components/Deploy');
var ModelSelectionModal_1 = require('./components/ModelSelectionModal');
var react_router_1 = require('react-router');
require('./styles/modeldetails.scss');
var model_overview_action_1 = require('./actions/model.overview.action');
var redux_1 = require('redux');
var react_redux_1 = require('react-redux');
var leaderboard_actions_1 = require('../Models/actions/leaderboard.actions');
var projects_actions_1 = require('../Projects/actions/projects.actions');
var deployment_actions_1 = require('../Deployment/actions/deployment.actions');
var ModelDetails = (function (_super) {
    __extends(ModelDetails, _super);
    function ModelDetails() {
        _super.call(this);
        this.state = {
            isModelOpen: true,
            isResidualOpen: true,
            isVariableOpen: true,
            isGoodnessOpen: true,
            isExportModalOpen: false,
            isModelSelectionModal: false,
            comparisonModel: null
        };
        this.exportModel = this.exportModel.bind(this);
    }
    ModelDetails.prototype.componentWillMount = function () {
        var _this = this;
        if (this.props.project) {
            this.props.fetchProject(parseInt(this.props.params.projectid, 10)).then(function (res) {
                _this.props.fetchLeaderboard(parseInt(_this.props.params.projectid, 10), res.model_category);
                _this.props.fetchSortCriteria(res.model_category.toLowerCase());
                _this.setState({
                    modelCategory: res.model_category.toLowerCase()
                });
            });
        }
        this.props.findModelsCount(parseInt(this.props.params.projectid, 10));
        this.props.fetchPackages(parseInt(this.props.params.projectid, 10));
        this.props.fetchModelOverview(parseInt(this.props.params.modelid, 10));
    };
    ModelDetails.prototype.toggleOpen = function (accordian) {
        /**
         * TODO(justinloyola): Fix the asynchronous state change issues
         */
        if (accordian === 'model') {
            this.setState({
                isModelOpen: !this.state.isModelOpen
            });
        }
        else if (accordian === 'residual') {
            this.setState({
                isResidualOpen: !this.state.isResidualOpen
            });
        }
        else if (accordian === 'variable') {
            this.setState({
                isVariableOpen: !this.state.isVariableOpen
            });
        }
        else if (accordian === 'goodness') {
            this.setState({
                isGoodnessOpen: !this.state.isGoodnessOpen
            });
        }
    };
    ModelDetails.prototype.forkModel = function () {
        react_router_1.hashHistory.push('/projects/forkmodel');
    };
    ModelDetails.prototype.exportModel = function () {
        this.setState({
            isExportModalOpen: !this.state.isExportModalOpen
        });
    };
    ModelDetails.prototype.cancel = function () {
        this.setState({
            isExportModalOpen: false
        });
    };
    ModelDetails.prototype.downloadModel = function (event) {
        event.preventDefault();
        this.props.downloadModel(event);
    };
    ModelDetails.prototype.deployModel = function () {
        this.setState({
            isDeployModalOpen: true
        });
    };
    ModelDetails.prototype.openComparisonModal = function () {
        this.setState({
            isModelSelectionModalOpen: true
        });
    };
    ModelDetails.prototype.closeComparisonModal = function () {
        this.setState({
            isModelSelectionModalOpen: false
        });
    };
    ModelDetails.prototype.onSelectModel = function (model) {
        this.closeComparisonModal();
        this.setState({
            comparisonModel: model
        });
    };
    ModelDetails.prototype.onCancel = function () {
        this.closeComparisonModal();
    };
    ModelDetails.prototype.onFilter = function (filters, name, offset) {
        this.props.fetchLeaderboard(parseInt(this.props.params.projectid, 10), this.state.modelCategory, name, filters.sortBy, filters.orderBy === 'asc', offset);
    };
    ModelDetails.prototype.closeDeployModal = function () {
        this.setState({
            isDeployModalOpen: false
        });
    };
    ModelDetails.prototype.onDeploy = function (model, serviceName, packageName) {
        this.setState({
            isDeployModalOpen: false
        });
        this.props.deployModel(model.id, serviceName, this.props.params.projectid, packageName);
    };
    ModelDetails.prototype.render = function () {
        if (_.isEmpty(this.props.model)) {
            return React.createElement("div", null);
        }
        return (React.createElement("div", {className: "model-details"}, React.createElement(ModelSelectionModal_1.default, {open: this.state.isModelSelectionModalOpen, onFilter: this.onFilter.bind(this), models: this.props.models, sortCriteria: this.props.sortCriteria, onSelectModel: this.onSelectModel.bind(this), onCancel: this.onCancel.bind(this), count: this.props.count}), React.createElement(ExportModal_1.default, {open: this.state.isExportModalOpen, name: this.props.model.name.toUpperCase(), onCancel: this.cancel.bind(this), modelId: parseInt(this.props.params.modelid, 10), projectId: parseInt(this.props.params.projectid, 10), onDownload: this.downloadModel.bind(this)}), React.createElement(Deploy_1.default, {open: this.state.isDeployModalOpen, onCancel: this.closeDeployModal.bind(this), model: this.props.model, onDeploy: this.onDeploy.bind(this), packages: this.props.packages}), React.createElement(PageHeader_1.default, null, React.createElement("span", null, this.props.model.name.toUpperCase()), React.createElement("div", {className: "buttons"}, React.createElement("button", {className: "default invert", onClick: this.exportModel.bind(this)}, "Export Model"), React.createElement("button", {className: "default", onClick: this.deployModel.bind(this)}, "Deploy Model")), React.createElement("div", {className: "comparison-selection"}, React.createElement("span", null, React.createElement("span", null, "compared to:"), React.createElement("button", {className: classNames('model-selection-button', { selected: this.state.comparisonModel }), onClick: this.openComparisonModal.bind(this)}, this.state.comparisonModel ? this.state.comparisonModel.name : 'SELECT MODEL FOR COMPARISON')))), React.createElement("header", {className: "overview-header"}, React.createElement("span", {onClick: this.toggleOpen.bind(this, 'model')}, React.createElement("i", {className: classNames('fa', { 'fa-minus-square-o': this.state.isModelOpen, 'fa-plus-square-o': !this.state.isModelOpen })}), "Model Overview")), React.createElement(Collapsible_1.default, {open: this.state.isModelOpen}, React.createElement(ModelOverview_1.default, {model: this.props.model})), React.createElement("header", {className: "overview-header"}, React.createElement("span", {onClick: this.toggleOpen.bind(this, 'goodness')}, React.createElement("i", {className: classNames('fa', { 'fa-minus-square-o': this.state.isGoodnessOpen, 'fa-plus-square-o': !this.state.isGoodnessOpen })}), "Goodness of Fit")), React.createElement(Collapsible_1.default, {open: this.state.isGoodnessOpen}, React.createElement(GoodnessOfFit_1.default, {model: this.props.model, comparisonModel: this.state.comparisonModel, modelCategory: this.state.modelCategory}))));
    };
    return ModelDetails;
}(React.Component));
exports.ModelDetails = ModelDetails;
function mapStateToProps(state) {
    return {
        model: state.model,
        count: state.leaderboard.count,
        packages: state.deployments.packages,
        project: state.projects.project,
        models: state.leaderboard.items,
        sortCriteria: state.leaderboard.criteria
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchLeaderboard: redux_1.bindActionCreators(leaderboard_actions_1.fetchLeaderboard, dispatch),
        fetchProject: redux_1.bindActionCreators(projects_actions_1.fetchProject, dispatch),
        fetchSortCriteria: redux_1.bindActionCreators(leaderboard_actions_1.fetchSortCriteria, dispatch),
        fetchModelOverview: redux_1.bindActionCreators(model_overview_action_1.fetchModelOverview, dispatch),
        downloadModel: redux_1.bindActionCreators(model_overview_action_1.downloadModel, dispatch),
        deployModel: redux_1.bindActionCreators(model_overview_action_1.deployModel, dispatch),
        fetchPackages: redux_1.bindActionCreators(deployment_actions_1.fetchPackages, dispatch),
        findModelsCount: redux_1.bindActionCreators(leaderboard_actions_1.findModelsCount, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(ModelDetails);
//# sourceMappingURL=ModelDetails.js.map