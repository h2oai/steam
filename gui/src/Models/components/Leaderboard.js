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
var _ = require('lodash');
var Deploy_1 = require('../components/Deploy');
var PageHeader_1 = require('../../Projects/components/PageHeader');
var Pagination_1 = require('../components/Pagination');
var BinomialModelTable_1 = require('./BinomialModelTable');
var MultinomialModelTable_1 = require('./MultinomialModelTable');
var RegressionModelTable_1 = require('./RegressionModelTable');
var ImportModelsModal_1 = require('./ImportModelsModal');
var leaderboard_actions_1 = require('../actions/leaderboard.actions');
require('../styles/leaderboard.scss');
var configuration_labels_action_1 = require('../../Configurations/actions/configuration.labels.action');
var redux_1 = require('redux');
var react_redux_1 = require('react-redux');
var deployment_actions_1 = require('../../Deployment/actions/deployment.actions');
var Leaderboard = (function (_super) {
    __extends(Leaderboard, _super);
    function Leaderboard(props) {
        _super.call(this, props);
        this.sampleData = {};
        this.state = {
            isDeployOpen: false,
            isImportModelsOpen: false,
            openDeployModel: null,
            currentPage: 0,
            filters: {
                sortBy: '',
                orderBy: 'asc'
            }
        };
        this.openDeploy = this.openDeploy.bind(this);
        this.closeHandler = this.closeHandler.bind(this);
        this.onChangeHandler = this.onChangeHandler.bind(this);
    }
    Leaderboard.prototype.componentWillMount = function () {
        if (!this.props.labels || !this.props.labels[this.props.projectId]) {
            this.props.fetchLabels(this.props.projectId);
        }
        this.props.fetchPackages(this.props.projectId);
        this.props.findModelsCount(this.props.projectId);
    };
    Leaderboard.prototype.openDeploy = function (model) {
        this.setState({
            isDeployOpen: true,
            openDeployModel: model
        });
    };
    Leaderboard.prototype.openImportModels = function () {
        this.setState({
            isImportModelsOpen: true
        });
    };
    Leaderboard.prototype.closeImportModels = function () {
        this.setState({
            isImportModelsOpen: false
        });
    };
    Leaderboard.prototype.closeHandler = function () {
        this.setState({
            isDeployOpen: false
        });
    };
    Leaderboard.prototype.onFilter = function (filters) {
        this.setState({
            filters: filters
        });
        this.props.onFilter(filters, this.refs.filterModels.value);
    };
    Leaderboard.prototype.onPageForward = function () {
        this.setState({
            currentPage: ++this.state.currentPage
        });
        this.props.onFilter(this.state.filters, this.refs.filterModels.value, this.state.currentPage * leaderboard_actions_1.MAX_ITEMS);
    };
    Leaderboard.prototype.onPageBack = function () {
        if (this.state.currentPage > 0) {
            this.setState({
                currentPage: --this.state.currentPage
            });
            this.props.onFilter(this.state.filters, this.refs.filterModels.value, this.state.currentPage * leaderboard_actions_1.MAX_ITEMS);
        }
    };
    Leaderboard.prototype.onDeploy = function (model, serviceName, packageName) {
        this.setState({
            isDeployOpen: false
        });
        this.props.deployModel(model.id, serviceName, this.props.projectId, packageName);
    };
    Leaderboard.prototype.onChangeHandler = function (labelId, modelId, isUnlink) {
        var _this = this;
        if (isUnlink === true) {
            this.props.unlinkLabelFromModel(labelId, modelId).then(function () {
                _this.props.fetchLabels(_this.props.projectId);
            });
        }
        else {
            this.props.linkLabelWithModel(labelId, modelId).then(function () {
                _this.props.fetchLabels(_this.props.projectId);
            });
        }
    };
    Leaderboard.prototype.getDataset = function () {
        return _.get(this.props, 'items[0].dataset_name');
    };
    Leaderboard.prototype.render = function () {
        return (React.createElement("div", {ref: "leaderboard", className: "leaderboard"}, React.createElement(ImportModelsModal_1.default, {projectId: this.props.projectId, open: this.state.isImportModelsOpen, onCancel: this.closeImportModels.bind(this), fetchLeaderboard: this.props.fetchLeaderboard, modelCategory: this.props.modelCategory, datasetName: this.getDataset()}), React.createElement(Deploy_1.default, {open: this.state.isDeployOpen, onCancel: this.closeHandler, model: this.state.openDeployModel, onDeploy: this.onDeploy.bind(this), packages: this.props.packages || []}), React.createElement(PageHeader_1.default, null, React.createElement("span", null, "Models"), React.createElement("div", {className: "button-primary header-buttons", onClick: this.openImportModels.bind(this)}, "Import Models")), React.createElement("div", {className: "filter"}, React.createElement("input", {ref: "filterModels", type: "text", placeholder: "filter models", onChange: this.onFilter.bind(this)})), this.props.modelCategory === 'binomial' ?
            React.createElement(BinomialModelTable_1.default, {onFilter: this.onFilter.bind(this), sortCriteria: this.props.sortCriteria, items: this.props.items, projectId: this.props.projectId, openDeploy: this.openDeploy.bind(this), labels: this.props.labels, onChangeHandler: this.onChangeHandler}) : null, this.props.modelCategory === 'multinomial' ?
            React.createElement(MultinomialModelTable_1.default, {onFilter: this.onFilter.bind(this), sortCriteria: this.props.sortCriteria, items: this.props.items, projectId: this.props.projectId, openDeploy: this.openDeploy.bind(this), labels: this.props.labels, onChangeHandler: this.onChangeHandler}) : null, this.props.modelCategory === 'regression' ?
            React.createElement(RegressionModelTable_1.default, {onFilter: this.onFilter.bind(this), sortCriteria: this.props.sortCriteria, items: this.props.items, projectId: this.props.projectId, openDeploy: this.openDeploy.bind(this), labels: this.props.labels, onChangeHandler: this.onChangeHandler}) : null, React.createElement(Pagination_1.default, {items: this.props.items, onPageBack: this.onPageBack.bind(this), onPageForward: this.onPageForward.bind(this), currentPage: this.state.currentPage, count: this.props.count})));
    };
    return Leaderboard;
}(React.Component));
exports.Leaderboard = Leaderboard;
function mapStateToProps(state) {
    return {
        count: state.leaderboard.count,
        labels: state.labels,
        packages: state.deployments.packages
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchLabels: redux_1.bindActionCreators(configuration_labels_action_1.fetchLabels, dispatch),
        linkLabelWithModel: redux_1.bindActionCreators(leaderboard_actions_1.linkLabelWithModel, dispatch),
        unlinkLabelFromModel: redux_1.bindActionCreators(leaderboard_actions_1.unlinkLabelFromModel, dispatch),
        fetchPackages: redux_1.bindActionCreators(deployment_actions_1.fetchPackages, dispatch),
        findModelsCount: redux_1.bindActionCreators(leaderboard_actions_1.findModelsCount, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(Leaderboard);
//# sourceMappingURL=Leaderboard.js.map