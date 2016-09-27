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
 * Created by justin on 9/8/16.
 */
var React = require('react');
var _ = require('lodash');
var react_redux_1 = require('react-redux');
var clusters_actions_1 = require('../actions/clusters.actions');
var redux_1 = require('redux');
var Cell_1 = require('../../Projects/components/Cell');
var Row_1 = require('../../Projects/components/Row');
var PageHeader_1 = require('../../Projects/components/PageHeader');
var Table_1 = require('../../Projects/components/Table');
require('../styles/launchcluster.scss');
var h2oUIKit_1 = require('h2oUIKit');
var LaunchCluster = (function (_super) {
    __extends(LaunchCluster, _super);
    function LaunchCluster() {
        _super.call(this);
        this.state = {
            memorySizeUnit: 'm',
            engineId: null
        };
    }
    LaunchCluster.prototype.componentDidMount = function () {
        this.props.getEngines();
        this.props.getConfig();
    };
    LaunchCluster.prototype.startCluster = function (event) {
        event.preventDefault();
        var clusterName = this.refs.clusterForm.querySelector('input[name="name"]').value;
        var engineId = this.state.engineId;
        var size = this.refs.clusterForm.querySelector('input[name="size"]').value;
        var memory = this.refs.clusterForm.querySelector('input[name="memory"]').value;
        var keytab = _.get(this.refs.clusterForm.querySelector('input[name="keytab"]'), 'value', '');
        this.props.startYarnCluster(clusterName, parseInt(engineId, 10), parseInt(size, 10), memory + this.state.memorySizeUnit, keytab);
    };
    LaunchCluster.prototype.uploadEngine = function (event) {
        event.preventDefault();
        this.props.uploadEngine(this.refs.engine);
    };
    LaunchCluster.prototype.onChangeMemory = function (event) {
        this.setState({
            memorySizeUnit: event.target.value
        });
    };
    LaunchCluster.prototype.onChangeEngine = function (event) {
        this.setState({
            engineId: event.target.value
        });
    };
    LaunchCluster.prototype.render = function () {
        return (React.createElement("div", {className: "launch-cluster"}, React.createElement(PageHeader_1.default, null, "LAUNCH NEW CLUSTER"), React.createElement("form", {ref: "clusterForm", onSubmit: this.startCluster.bind(this)}, React.createElement(Table_1.default, null, React.createElement(Row_1.default, {header: true}), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "CLUSTER NAME"), React.createElement(Cell_1.default, null, React.createElement("input", {type: "text", name: "name"}))), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "NUMBER OF NODES"), React.createElement(Cell_1.default, null, React.createElement(h2oUIKit_1.NumericInput, {name: "size", min: "1"}))), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "MEMORY PER NODE"), React.createElement(Cell_1.default, null, React.createElement(h2oUIKit_1.NumericInput, {name: "memory", min: "1"}), React.createElement("select", {className: "memory-selection", onChange: this.onChangeMemory.bind(this)}, React.createElement("option", {value: "m"}, "MB"), React.createElement("option", {value: "g"}, "GB")))), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "H2O VERSION"), React.createElement(Cell_1.default, null, React.createElement("div", {className: "upload-engine"}, React.createElement("input", {ref: "engine", type: "file", name: "engine"}), React.createElement("div", {className: "button-primary", onClick: this.uploadEngine.bind(this)}, "Upload Engine")), React.createElement("select", {onChange: this.onChangeEngine.bind(this)}, React.createElement("option", null), this.props.engines.map(function (engine, i) {
            return React.createElement("option", {key: i, value: engine.id}, engine.name);
        })))), _.get(this.props.config, 'kerberos_enabled', false) === true ? React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "Kerberos Keytab"), React.createElement(Cell_1.default, null, React.createElement("input", {type: "text", name: "keytab"}))) : null), React.createElement("div", {type: "submit", className: "button-primary"}, "Launch New Clusters"))));
    };
    return LaunchCluster;
}(React.Component));
exports.LaunchCluster = LaunchCluster;
function mapStateToProps(state) {
    return {
        engines: state.clusters.engines,
        config: state.clusters.config
    };
}
function mapDispatchToProps(dispatch) {
    return {
        uploadEngine: redux_1.bindActionCreators(clusters_actions_1.uploadEngine, dispatch),
        startYarnCluster: redux_1.bindActionCreators(clusters_actions_1.startYarnCluster, dispatch),
        getEngines: redux_1.bindActionCreators(clusters_actions_1.getEngines, dispatch),
        getConfig: redux_1.bindActionCreators(clusters_actions_1.getConfig, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(LaunchCluster);
//# sourceMappingURL=LaunchCluster.js.map