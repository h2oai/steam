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

/**
 * Created by justin on 6/27/16.
 */
import * as React from 'react';
import * as _ from 'lodash';
import * as $ from 'jquery';
import Panel from '../Projects/components/Panel';
import PageHeader from '../Projects/components/PageHeader';
import {
  fetchModelsFromCluster, fetchClusters, registerCluster,
  unregisterCluster, stopClusterOnYarn
} from '../Projects/actions/projects.actions';
import { bindActionCreators } from 'redux';
import { Cluster } from '../Proxy/Proxy';
import { connect } from 'react-redux';
import './styles/clusters.scss';
import { Link } from 'react-router';
import { getConfig } from './actions/clusters.actions';
import ConfirmDeleteClusterDialog from "./components/ConfirmDeleteClusterDialog";
import { hasPermissionToShow } from "../App/utils/permissions";

interface DispatchProps {
  fetchClusters: Function
  fetchModelsFromCluster: Function
  registerCluster: Function,
  unregisterCluster: Function,
  stopClusterOnYarn: Function,
  getConfig: Function
}

interface Props {
  clusters: Cluster[],
  config: any,
  isAdmin: boolean
}

export class Clusters extends React.Component<Props & DispatchProps, any> {
  refs: {
    [key: string]: Element
    keytabFilename: HTMLInputElement
  };

  constructor(props) {
    super(props);
    this.state = {
      yarnClusterModalOpen: false,
      newClusterRequested: false,
      selectedCluster: null,
      confirmDeleteClusterOpen: false,
      clustersDeletedIds: []
    };
  }

  componentWillMount(): void {
    this.props.fetchClusters();
    this.props.getConfig();
  }

  goProxy(cluster) {
    document.cookie = cluster.name + "=" + cluster.token;
    let url = "https://" + window.location.hostname + this.props.config.cluster_proxy_address + cluster.context_path + "flow/index.html";
    window.open(url, "_blank");
  }

  openYarnClusterModal() {
    this.setState({
      yarnClusterModalOpen: true
    });
  }

  registerCluster(event) {
    event.preventDefault();
    let ipAddress = $(event.target).find('input[name="ip-address"]').val();
    let port = $(event.target).find('input[name="port"]').val();
    this.props.registerCluster(ipAddress + ':' + port);
    this.setState({newClusterRequested: false});
  }

  onCreateNewClusterClicked(e) {
    this.setState({newClusterRequested: true});
  }

  removeCluster = (cluster) => {
    if (cluster.type_id === 2) {
      let keytabFilename = _.get((this.refs.keytabFilename as HTMLInputElement), 'value', '');
      this.props.stopClusterOnYarn(cluster.id, keytabFilename);
    } else {
      this.props.unregisterCluster(cluster.id);
    }
  };

  onDeleteClusterClicked = (cluster) => {
    this.setState({
      selectedCluster: cluster,
      confirmDeleteClusterOpen: true
    });
  };

  onDeleteClusterConfirmed = () => {
    let clusterToDeleteId = this.state.selectedCluster.id;
    let newClusterDeletedIdsArray = JSON.parse(JSON.stringify(this.state.clustersDeletedIds));
    newClusterDeletedIdsArray.push(clusterToDeleteId);
    this.removeCluster(this.state.selectedCluster);
    this.setState({
      clustersDeletedIds: newClusterDeletedIdsArray
    });

    this.setState({
      selectedCluster: null,
      confirmDeleteClusterOpen: false
    });
  };

  onDeleteClusterCanceled = () => {
    this.setState({
      selectedCluster: null,
      confirmDeleteClusterOpen: false
    });
  };

  isInDeletedClusters = (clusterId): boolean => {
    for (let deletedClusterId of this.state.clustersDeletedIds) {
      if (clusterId === deletedClusterId) {
        return true;
      }
    }
    return false;
  };

  render(): React.ReactElement<HTMLDivElement> {
    if (!this.props.clusters) {
      return <div></div>;
    }

    return (
      <div className="clusters">
        <ConfirmDeleteClusterDialog isOpen={this.state.confirmDeleteClusterOpen} clusterToDelete={this.state.selectedCluster} onDeleteConfirmed={this.onDeleteClusterConfirmed} onCancel={this.onDeleteClusterCanceled} />

        <PageHeader>CLUSTERS
          { !this.state.newClusterRequested ?
            <div className="buttons header-buttons">
              <div className="button-secondary" onClick={this.onCreateNewClusterClicked.bind(this)}>
                Connect to Cluster
              </div>
              {hasPermissionToShow("ViewEngine", this.props.config, this.props.isAdmin) ?
              <Link to="clusters/new" className="button-primary">Launch New Cluster</Link> : null}
            </div>
            : null }
        </PageHeader>
        <div className="panel-container">
          { this.state.newClusterRequested ?
            <Panel>
              <header>
                <h2 className="new-cluster-header">New Cluster</h2>
              </header>
              <p>Connect to a H2O cluster where your existing models and data sets are located.</p>
              <div className="new-cluster-form">
                <form onSubmit={this.registerCluster.bind(this)}>
                  <input type="text" name="ip-address" placeholder="IP Address"/>
                  <input type="text" name="port" placeholder="Port"/>
                  <button type="submit" className="button-primary">Connect</button>
                </form>
              </div>
            </Panel>
            : null
          }
          {this.props.clusters.map((cluster :any, i) => {
            return (
              <Panel key={i}>
                <header>
                  <span><i className="fa fa-cubes mar-bot-20"/> <a onClick={cluster.context_path !== "" ? this.goProxy.bind(this, cluster) : null} href={cluster.context_path !== "" ? null : 'http://' + cluster.address + cluster.context_path} target="_blank"
                                                        rel="noopener" className="charcoal-grey semibold link">{cluster.name}</a> {cluster.status.total_cpu_count ? <span> -- {cluster.status.total_cpu_count} &nbsp;cores</span> : null}</span>
                  <span className="remove-cluster">
                    {hasPermissionToShow("ManageCluster", this.props.config, this.props.isAdmin) ?
                      this.isInDeletedClusters(cluster.id) ?
                        <div className="pt-spinner modifier pt-small">
                          <div className="pt-spinner-svg-container">
                            <svg viewBox="0 0 100 100">
                              <path className="pt-spinner-track"
                                    d="M 50,50 m 0,-44.5 a 44.5,44.5 0 1 1 0,89 a 44.5,44.5 0 1 1 0,-89"></path>
                              <path className="pt-spinner-head" d="M 94.5 50 A 44.5 44.5 0 0 0 50 5.5"></path>
                            </svg>
                          </div>
                        </div>
                        :
                        <button className="remove-cluster-button" onClick={(e) => this.onDeleteClusterClicked(cluster)}>
                          <i className="fa fa-trash no-margin"/>
                        </button>
                      : null }
                  </span>
                </header>
                <div className="flexrow">
                  <div className="flexcolumn">
                    <div className="info-header">STATUS</div>
                    <div className="flexrow mar-right-71">
                      {cluster.state === "stopped" ?
                        <div className="flexrow">
                          <div className="infodot-container"><i className="fa fa-circle orange mar-right-3"/> Stopped</div>
                        </div> :
                        <div className="flexrow">
                          { cluster.status.status === "Healthy" ?
                            <div className="infodot-container"><i className="fa fa-circle green mar-right-3"/> Healthy</div>
                            : <div className="infodot-container"><i className="fa fa-circle orange mar-right-3"/> {cluster.status.status}</div>
                          }
                          { cluster.state === "started" ?
                            <div className="infodot-container"><i className="fa fa-circle green mar-right-3"/> Started</div>
                            : <div className="infodot-container"><i className="fa fa-circle orange mar-right-3"/> {cluster.state}</div>
                          }
                        </div>
                      }
                    </div>
                </div>
                {cluster.status.total_cpu_count ? <div className="flexcolumn">
                  <div className="info-header">VERSION</div>
                  <div className="charcoal-grey">{cluster.status.version}</div>
                </div> : null }
              </div>
              </Panel>
            );
          })}
        </div>
      </div>
    );
  }
}

function mapStateToProps(state): any {
  return {
    clusters: state.projects.clusters,
    config: state.clusters.config,
    isAdmin: state.global.isAdmin
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    fetchClusters: bindActionCreators(fetchClusters, dispatch),
    fetchModelsFromCluster: bindActionCreators(fetchModelsFromCluster, dispatch),
    registerCluster: bindActionCreators(registerCluster, dispatch),
    unregisterCluster: bindActionCreators(unregisterCluster, dispatch),
    stopClusterOnYarn: bindActionCreators(stopClusterOnYarn, dispatch),
    getConfig: bindActionCreators(getConfig, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Clusters);
