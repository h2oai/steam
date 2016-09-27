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
 * Created by justin on 7/11/16.
 */
var React = require('react');
var _ = require('lodash');
var PageHeader_1 = require('../Projects/components/PageHeader');
var TabNavigation_1 = require('../Projects/components/TabNavigation');
var DeployedServices_1 = require('../Projects/components/DeployedServices');
var Packaging_1 = require('./components/Packaging');
var UploadPreProcessingModal_1 = require('./components/UploadPreProcessingModal');
var react_redux_1 = require('react-redux');
var deployment_actions_1 = require('./actions/deployment.actions');
var redux_1 = require('redux');
require('./styles/deployment.scss');
var Deployment = (function (_super) {
    __extends(Deployment, _super);
    function Deployment() {
        _super.call(this);
        this.state = {
            tabs: {
                deployedServices: {
                    label: 'DEPLOYED SERVICES',
                    isSelected: true,
                    onClick: this.clickHandler.bind(this),
                    component: React.createElement(DeployedServices_1.default, null)
                },
                packaging: {
                    label: 'PACKAGING',
                    isSelected: false,
                    onClick: this.clickHandler.bind(this),
                    component: React.createElement(Packaging_1.default, null)
                }
            },
            isSelected: 'deployedServices',
            uploadOpen: false,
            packages: [],
            projectId: null
        };
    }
    Deployment.prototype.componentWillMount = function () {
        this.setState({
            tabs: {
                deployedServices: {
                    label: 'DEPLOYED SERVICES',
                    isSelected: true,
                    onClick: this.clickHandler.bind(this),
                    component: React.createElement(DeployedServices_1.default, null)
                },
                packaging: {
                    label: 'PACKAGING',
                    isSelected: false,
                    onClick: this.clickHandler.bind(this),
                    component: React.createElement(Packaging_1.default, {projectId: this.props.params.projectid})
                }
            }
        });
    };
    Deployment.prototype.clickHandler = function (tab) {
        var key = _.findKey(this.state.tabs, tab);
        var newState = _.cloneDeep(this.state);
        Object.keys(newState.tabs).map(function (tab) {
            newState.tabs[tab].isSelected = false;
        });
        newState.tabs[key].isSelected = true;
        newState.isSelected = key;
        this.setState(newState);
    };
    Deployment.prototype.openUpload = function () {
        this.setState({
            uploadOpen: true
        });
    };
    Deployment.prototype.closeUpload = function () {
        this.setState({
            uploadOpen: false
        });
    };
    Deployment.prototype.upload = function (event, uploadedPackage, formData) {
        event.preventDefault();
        this.props.uploadPackage(parseInt(this.props.params.projectid, 10), uploadedPackage.name, formData);
        this.closeUpload();
    };
    Deployment.prototype.render = function () {
        return (React.createElement("div", {className: "services"}, !_.isUndefined(this.props.params.projectid) ?
            React.createElement(UploadPreProcessingModal_1.default, {open: this.state.uploadOpen, cancel: this.closeUpload.bind(this), upload: this.upload.bind(this)}) : null, React.createElement(PageHeader_1.default, null, React.createElement("span", null, "Deployment"), !_.isUndefined(this.props.params.projectid) ? React.createElement("div", {className: "button-primary header-buttons", onClick: this.openUpload.bind(this)}, "Upload New Package") : null), !_.isUndefined(this.props.params.projectid) ? React.createElement(TabNavigation_1.default, {tabs: this.state.tabs}) : null, React.createElement("main", null, this.state.tabs.deployedServices.isSelected === true ?
            React.createElement(DeployedServices_1.default, {projectId: this.props.params.projectid}) : null, this.state.tabs.packaging.isSelected === true && !_.isUndefined(this.props.params.projectid) ?
            React.createElement(Packaging_1.default, {projectId: this.props.params.projectid}) : null)));
    };
    return Deployment;
}(React.Component));
exports.Deployment = Deployment;
function mapStateToProps(state) {
    return {
        packages: state.packages
    };
}
function mapDispatchToProps(dispatch) {
    return {
        uploadPackage: redux_1.bindActionCreators(deployment_actions_1.uploadPackage, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(Deployment);
//# sourceMappingURL=Deployment.js.map