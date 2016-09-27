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
var _ = require('lodash');
var Panel_1 = require('./Panel');
var services_actions_1 = require('../actions/services.actions');
var redux_1 = require('redux');
var react_redux_1 = require('react-redux');
require('../styles/deployedservices.scss');
var DeployedServices = (function (_super) {
    __extends(DeployedServices, _super);
    function DeployedServices() {
        _super.apply(this, arguments);
    }
    DeployedServices.prototype.componentWillMount = function () {
        this.fetchServicesStrategy(this.props.projectId);
    };
    DeployedServices.prototype.fetchServicesStrategy = function (projectId) {
        if (projectId) {
            return this.props.fetchServicesForProject(parseInt(projectId, 10));
        }
        else {
            return this.props.fetchAllServices();
        }
    };
    DeployedServices.prototype.killService = function (serviceId) {
        this.props.killService(serviceId, parseInt(this.props.projectId, 10));
    };
    DeployedServices.prototype.render = function () {
        var _this = this;
        var runningServices;
        if (this.props.projectId) {
            runningServices = this.props.services.runningServicesForProject;
        }
        else {
            runningServices = this.props.services.allRunningServices;
        }
        if (_.isEmpty(runningServices)) {
            return (React.createElement("div", null, React.createElement("div", {className: "lede intro"}, "There are no services currently deployed.")));
        }
        return (React.createElement("div", {className: "deployed-services"}, React.createElement("section", null, runningServices.map(function (service, i) {
            return (React.createElement(Panel_1.default, {key: i, className: "services-panel"}, React.createElement("div", {className: "panel-body"}, React.createElement("div", {className: "panel-title"}, React.createElement("span", null, service.name, " @ ", React.createElement("a", {href: 'http://' + service.address + ':' + service.port, target: "_blank", rel: "noopener"}, service.address + ':' + service.port)), React.createElement("div", {style: { color: service.state === 'stopped' ? 'red' : 'green' }}, service.state)), React.createElement("div", {className: "panel-info"}, React.createElement("div", {className: "panel-info-row"}, React.createElement("span", null, React.createElement("i", {className: "fa fa-cube"})), React.createElement("span", null, "Model"), React.createElement("span", null, service.model_id)), React.createElement("div", {className: "panel-info-row"}, React.createElement("span", null, React.createElement("i", {className: "fa fa-folder-o"})), React.createElement("span", null, "Status"), React.createElement("span", null, service.state === 'started' ? 'OK' : 'Error')))), React.createElement("div", {className: "panel-actions"}, React.createElement("div", {className: "panel-action", onClick: _this.killService.bind(_this, service.id)}, React.createElement("div", null, React.createElement("i", {className: "fa fa-close"})), React.createElement("div", null, "Stop Service")))));
        }))));
    };
    return DeployedServices;
}(React.Component));
exports.DeployedServices = DeployedServices;
function mapStateToProps(state) {
    return {
        services: state.services
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchAllServices: redux_1.bindActionCreators(services_actions_1.fetchAllServices, dispatch),
        fetchServicesForProject: redux_1.bindActionCreators(services_actions_1.fetchServicesForProject, dispatch),
        killService: redux_1.bindActionCreators(services_actions_1.killService, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(DeployedServices);
//# sourceMappingURL=DeployedServices.js.map