/**
 * Created by justin on 6/27/16.
 */
import * as React from 'react';
import * as _ from 'lodash';
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
  componentWillMount(): void {
    if (_.isEmpty(this.props.clusters)) {
      this.props.fetchClusters();
    }
  }

  removeCluster(clusterId) {
    this.props.unregisterCluster(clusterId);
  }

  render(): React.ReactElement<HTMLDivElement> {
    if (!this.props.clusters) {
      return <div></div>;
    }
    return (
      <div className="clusters">
        <PageHeader>CLUSTERS</PageHeader>
        <div className="panel-container">
          {this.props.clusters.map((cluster, i) => {
            return (
              <Panel key={i}>
                <header>
                  <span><i className="fa fa-cubes"/> <a href={cluster.address} target="_blank"
                                                        rel="noopener">{cluster.name}
                    @ {cluster.address}</a></span><button className="remove-cluster" onClick={this.removeCluster.bind(this, cluster.id)}><i className="fa fa-trash"/></button>
                </header>
                <article>
                  <h3>
                    STATUS
                  </h3>
                  <h2 className="cluster-status">
                    {cluster.state === 'started' ? 'OK' : cluster.state}
                  </h2>
                </article>
                <h3>ACCESS</h3>
                <div>First Last</div>
                <div><span className="fa fa-plus"></span><span> ADD USER</span></div>
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
