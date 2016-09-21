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
  unregisterCluster
} from '../Projects/actions/projects.actions';
import { bindActionCreators } from 'redux';
import { Cluster } from '../Proxy/Proxy';
import { connect } from 'react-redux';
import './styles/clusters.scss';

interface DispatchProps {
  fetchClusters: Function
  fetchModelsFromCluster: Function
  registerCluster: Function,
  unregisterCluster: Function
}

interface Props {
  clusters: Cluster[]
}

export class Clusters extends React.Component<Props & DispatchProps, any> {

  constructor(props) {
    super(props);
    this.state = {
      newClusterRequested: false
    };
  }

  componentWillMount(): void {
    if (_.isEmpty(this.props.clusters)) {
      this.props.fetchClusters();
    }
  }

  removeCluster(clusterId) {
    this.props.unregisterCluster(clusterId);
  }

  registerCluster(event) {
    event.preventDefault();
    let ipAddress = $(event.target).find('input[name="ip-address"]').val();
    let port = $(event.target).find('input[name="port"]').val();
    this.props.registerCluster(ipAddress + ':' + port);
    this.setState({ newClusterRequested: false });
  }

  onCreateNewClusterClicked(e) {
    this.setState({ newClusterRequested: true });
  }

  render(): React.ReactElement<HTMLDivElement> {
    if (!this.props.clusters) {
      return <div></div>;
    }
    return (
      <div className="clusters">
        <PageHeader>CLUSTERS
          { !this.state.newClusterRequested ?
            <div className="button-primary header-buttons" onClick={this.onCreateNewClusterClicked.bind(this)}>Create Cluster</div>
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
                  <span><i className="fa fa-cubes"/> <a href={cluster.address} target="_blank"
                                                        rel="noopener">{cluster.name}
                    @ {cluster.address}</a></span><button className="remove-cluster" onClick={this.removeCluster.bind(this, cluster.id)}><i className="fa fa-trash"/></button>
                </header>
                <article>
                  <h2>
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
    clusters: state.projects.clusters
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    fetchClusters: bindActionCreators(fetchClusters, dispatch),
    fetchModelsFromCluster: bindActionCreators(fetchModelsFromCluster, dispatch),
    registerCluster: bindActionCreators(registerCluster, dispatch),
    unregisterCluster: bindActionCreators(unregisterCluster, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Clusters);
