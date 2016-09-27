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
var $ = require('jquery');
var Panel_1 = require('../Projects/components/Panel');
var PageHeader_1 = require('../Projects/components/PageHeader');
var projects_actions_1 = require('../Projects/actions/projects.actions');
var redux_1 = require('redux');
var react_redux_1 = require('react-redux');
require('./styles/clusters.scss');
var react_router_1 = require('react-router');
var clusters_actions_1 = require('./actions/clusters.actions');
var Clusters = (function (_super) {
    __extends(Clusters, _super);
    function Clusters(props) {
        _super.call(this, props);
        this.state = {
            yarnClusterModalOpen: false,
            newClusterRequested: false
        };
    }
    Clusters.prototype.componentWillMount = function () {
        if (_.isEmpty(this.props.clusters)) {
            this.props.fetchClusters();
            this.props.getConfig();
        }
    };
    Clusters.prototype.removeCluster = function (cluster) {
        if (cluster.type_id === 2) {
            var keytabFilename = _.get(this.refs.keytabFilename, 'value', '');
            this.props.stopClusterOnYarn(cluster.id, keytabFilename);
        }
        else {
            this.props.unregisterCluster(cluster.id);
        }
    };
    Clusters.prototype.openYarnClusterModal = function () {
        this.setState({
            yarnClusterModalOpen: true
        });
    };
    Clusters.prototype.registerCluster = function (event) {
        event.preventDefault();
        var ipAddress = $(event.target).find('input[name="ip-address"]').val();
        var port = $(event.target).find('input[name="port"]').val();
        this.props.registerCluster(ipAddress + ':' + port);
        this.setState({ newClusterRequested: false });
    };
    Clusters.prototype.onCreateNewClusterClicked = function (e) {
        this.setState({ newClusterRequested: true });
    };
    Clusters.prototype.render = function () {
        var _this = this;
        if (!this.props.clusters) {
            return React.createElement("div", null);
        }
        return (React.createElement("div", {className: "clusters"}, React.createElement(PageHeader_1.default, null, "CLUSTERS", !this.state.newClusterRequested ?
            React.createElement("div", {className: "buttons header-buttons"}, React.createElement("div", {className: "button-secondary", onClick: this.onCreateNewClusterClicked.bind(this)}, "Connect to Cluster"), React.createElement(react_router_1.Link, {to: "clusters/new", className: "button-primary"}, "Launch New Cluster"))
            : null), React.createElement("div", {className: "panel-container"}, this.state.newClusterRequested ?
            React.createElement(Panel_1.default, null, React.createElement("header", null, React.createElement("h2", {className: "new-cluster-header"}, "New Cluster")), React.createElement("p", null, "Connect to a H2O cluster where your existing models and data sets are located."), React.createElement("div", {className: "new-cluster-form"}, React.createElement("form", {onSubmit: this.registerCluster.bind(this)}, React.createElement("input", {type: "text", name: "ip-address", placeholder: "IP Address"}), React.createElement("input", {type: "text", name: "port", placeholder: "Port"}), React.createElement("button", {type: "submit", className: "button-primary"}, "Connect"))))
            : null, this.props.clusters.map(function (cluster, i) {
            return (React.createElement(Panel_1.default, {key: i}, React.createElement("header", null, React.createElement("span", null, React.createElement("i", {className: "fa fa-cubes"}), " ", React.createElement("a", {href: window.location.protocol + '//' + window.location.hostname + _.get(_this.props.config, 'cluster_proxy_address', '') + '/flow/?cluster_id=' + cluster.id, target: "_blank", rel: "noopener"}, cluster.name)), React.createElement("span", {className: "remove-cluster"}, _.get(_this.props.config, 'kerberos_enabled', false) === true ? React.createElement("input", {ref: "keytabFilename", type: "text", placeholder: "Keytab filename"}) : null, React.createElement("button", {className: "remove-cluster-button", onClick: _this.removeCluster.bind(_this, cluster)}, React.createElement("i", {className: "fa fa-trash"})))), React.createElement("article", null, React.createElement("h2", null, "ID: ", cluster.id, "STATUS"), React.createElement("h2", {className: "cluster-status"}, cluster.state === 'started' ? 'OK' : cluster.state)), React.createElement("h2", null, "ACCESS"), cluster.identities ?
                cluster.identities.map(function (identity, index) {
                    return React.createElement("div", {key: index}, identity.identity_name, " ");
                })
                : null));
        }))));
    };
    return Clusters;
}(React.Component));
exports.Clusters = Clusters;
function mapStateToProps(state) {
    return {
        clusters: state.projects.clusters,
        config: state.clusters.config
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchClusters: redux_1.bindActionCreators(projects_actions_1.fetchClusters, dispatch),
        fetchModelsFromCluster: redux_1.bindActionCreators(projects_actions_1.fetchModelsFromCluster, dispatch),
        registerCluster: redux_1.bindActionCreators(projects_actions_1.registerCluster, dispatch),
        unregisterCluster: redux_1.bindActionCreators(projects_actions_1.unregisterCluster, dispatch),
        stopClusterOnYarn: redux_1.bindActionCreators(projects_actions_1.stopClusterOnYarn, dispatch),
        getConfig: redux_1.bindActionCreators(clusters_actions_1.getConfig, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(Clusters);
//# sourceMappingURL=Clusters.js.map