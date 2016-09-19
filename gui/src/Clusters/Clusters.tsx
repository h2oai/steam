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

interface DispatchProps {
  fetchClusters: Function
  fetchModelsFromCluster: Function
  registerCluster: Function,
  unregisterCluster: Function,
  stopClusterOnYarn: Function
}

interface Props {
  clusters: Cluster[]
}

export class Clusters extends React.Component<Props & DispatchProps, any> {
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
    }
  }

  removeCluster(cluster) {
    if (cluster.type_id === 2) {
      this.props.stopClusterOnYarn(cluster.id);
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
            <div className="buttons default">
              <button className="default invert" onClick={this.onCreateNewClusterClicked.bind(this)}>
                Connect to Cluster
              </button>
              <Link to="clusters/new" className="default">Launch New Cluster</Link>
            </div>
            : null }
        </PageHeader>
        <div className="panel-container">
          { this.state.newClusterRequested ?
            <Panel>
              <header>
                <h3 className="new-cluster-header">New Cluster</h3>
              </header>
              <p>Connect to a H2O cluster where your existing models and data sets are located.</p>
              <div className="new-cluster-form">
                <form onSubmit={this.registerCluster.bind(this)}>
                  <input type="text" name="ip-address" placeholder="IP Address"/>&nbsp;
                  <input type="text" name="port" placeholder="Port"/>&nbsp;<br /><br />
                  <button type="submit" className="default">Connect</button>
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
                    @ {cluster.address}</a></span>
                  <button className="remove-cluster" onClick={this.removeCluster.bind(this, cluster)}><i
                    className="fa fa-trash"/></button>
                </header>
                <article>
                  <h3>
                    ID
                  </h3>
                  <h2>
                    {cluster.id}
                  </h2>
                  <h3>
                    STATUS
                  </h3>
                  <h2 className="cluster-status">
                    {cluster.state === 'started' ? 'OK' : cluster.state}
                  </h2>
                </article>
                <h3>ACCESS</h3>
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
    unregisterCluster: bindActionCreators(unregisterCluster, dispatch),
    stopClusterOnYarn: bindActionCreators(stopClusterOnYarn, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Clusters);
