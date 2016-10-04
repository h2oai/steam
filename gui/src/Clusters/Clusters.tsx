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
      newClusterRequested: false
    };
  }

  componentWillMount(): void {
    if (_.isEmpty(this.props.clusters)) {
      this.props.fetchClusters();
      this.props.getConfig();
    }
  }

  removeCluster(cluster) {
    if (cluster.type_id === 2) {
      let keytabFilename = _.get((this.refs.keytabFilename as HTMLInputElement), 'value', '');
      this.props.stopClusterOnYarn(cluster.id, keytabFilename);
    } else {
      this.props.unregisterCluster(cluster.id);
    }
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

  render(): React.ReactElement<HTMLDivElement> {
    if (!this.props.clusters) {
      return <div></div>;
    }
    return (
      <div className="clusters">
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
          {this.props.clusters.map((cluster, i) => {
            return (
              <Panel key={i}>
                <header>
                  <span><i className="fa fa-cubes"/> <a href={window.location.protocol + '//' + window.location.hostname + _.get(this.props.config, 'cluster_proxy_address', '') + '/flow/?cluster_id=' + cluster.id} target="_blank"
                                                        rel="noopener">{cluster.name}</a></span>
                  <span className="remove-cluster">
                    {_.get(this.props.config, 'kerberos_enabled', false) === true ? <input ref="keytabFilename" type="text" placeholder="Keytab filename"/> : null}
                    <button className="remove-cluster-button" onClick={this.removeCluster.bind(this, cluster)}><i
                      className="fa fa-trash"/></button>
                  </span>
                </header>
                <article>
                  <h2>
                    ID: {cluster.id}
                    STATUS
                  </h2>
                  <h2 className="cluster-status">
                    {cluster.state === 'started' ? 'OK' : cluster.state}
                  </h2>
                </article>
                <h2>ACCESS</h2>
                { (cluster as any).identities ?
                  (cluster as any).identities.map((identity, index) => {
                    return <div key={index}>{ identity.identity_name }&nbsp;</div>;
                  })
                  : null }
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
