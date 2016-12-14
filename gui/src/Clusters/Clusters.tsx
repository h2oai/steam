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
  config: any
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
      confirmDeleteClusterOpen: false
    };
  }

  componentWillMount(): void {
    this.props.fetchClusters();
    this.props.getConfig();
  }

  goProxy(cluster) {
    document.cookie = cluster.name + "=" + cluster.token;
    let url = "http://" + window.location.hostname + ":9999" + cluster.context_path;
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
    this.removeCluster(this.state.selectedCluster);
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
              <Link to="clusters/new" className="button-primary">Launch New Cluster</Link>
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
                  <span><i className="fa fa-cubes mar-bot-20"/> <a href={'http://' + cluster.address + cluster.context_path} target="_blank"
                                                        rel="noopener" className="charcoal-grey semibold">{cluster.name}</a> -- {cluster.status.total_cpu_count}&nbsp;cores</span>
                  <span className="remove-cluster">
                    <button className="remove-cluster-button test" onClick={this.goProxy.bind(this, cluster)}>
                      <i className="fa fa-arrow-circle-o-right no-margin"/>
                    </button>
                  </span>
                  <span className="remove-cluster">
                    {_.get(this.props.config, 'kerberos_enabled', false) ? <input ref="keytabFilename" type="text" placeholder="Keytab filename"/> : null}
                    <button className="remove-cluster-button" onClick={(e) => this.onDeleteClusterClicked(cluster)}><i
                      className="fa fa-trash no-margin"/></button>
                  </span>
                </header>
                <div className="flexrow">
                  <div className="flexcolumn">
                    <div className="info-header">STATUS</div>
                    <div className="flexrow mar-right-71">
                      { cluster.status.status === "healthy" ?
                        <div className="infodot-container"><i className="fa fa-circle green mar-right-3"/> Healthy</div>
                        : <div className="infodot-container"><i className="fa fa-circle orange mar-right-3"/> {cluster.status.status}</div>
                      }
                      { cluster.state === "started" ?
                        <div className="infodot-container"><i className="fa fa-circle green mar-right-3"/> Started</div>
                        : <div className="infodot-container"><i className="fa fa-circle orange mar-right-3"/> {cluster.state}</div>
                      }
                    </div>
                </div>
                <div className="flexcolumn">
                  <div className="info-header">VERSION</div>
                  <div className="charcoal-grey">{cluster.status.version}</div>
                </div>
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
    config: state.clusters.config
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
