/**
 * Created by justin on 7/18/16.
 */
import * as React from 'react';
import * as $ from 'jquery';
import * as _ from 'lodash';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Table from './Table';
import Row from './Row';
import Cell from './Cell';
import {
  fetchClusters, fetchModelsFromCluster,
  importModelFromCluster, createProjectAndImportModelsFromCluster, registerCluster
} from '../actions/projects.actions';
import { Cluster, Model } from '../../Proxy/proxy';
import '../styles/importnewproject.scss';
import { hashHistory } from 'react-router';

interface DispatchProps {
  fetchClusters: Function,
  fetchModelsFromCluster: Function,
  importModelFromCluster: Function,
  createProjectAndImportModelsFromCluster: Function,
  registerCluster: Function
}

interface Props {
  clusters: Cluster[],
  models: Model[]
}

export class ImportNewProject extends React.Component<DispatchProps & Props, any> {
  refs: {
    [key: string]: (Element);
    projectName: (HTMLInputElement)
  };

  constructor() {
    super();
    this.state = {
      clusterId: null
    };
  }

  componentWillMount(): void {
    if (_.isEmpty(this.props.clusters)) {
      this.props.fetchClusters();
    }
  }

  selectCluster(id: number): void {
    this.setState({
      clusterId: id
    });
    this.props.fetchModelsFromCluster(id);
  }

  createProject(): void {
    let name = $(this.refs.projectName).val();
    let importModels = [];
    $('.import-models input:checked').map((i, input) => {
      importModels.push($(input).prop('name'));
    });
    this.props.createProjectAndImportModelsFromCluster(name, this.state.clusterId, importModels).then((res) => {
      hashHistory.push('/projects/' + res + '/models');
    });
  }

  registerCluster(event) {
    event.preventDefault();
    let ipAddress = $(event.target).find('input[name="ip-address"]').val();
    let port = $(event.target).find('input[name="port"]').val();
    this.props.registerCluster(ipAddress + ':' + port);
  }

  render(): React.ReactElement<HTMLDivElement> {
    if (!this.props.clusters) {
      return <div></div>;
    }
    return (
      <div className="import-new-project">
        <div className="step-1">
          <div className="select-cluster">
            <h1>1. Select H2O cluster</h1>
            <div>
              Select an H2O cluster to import models and datasets from.
              <Table>
                <Row header={true}>
                  <Cell>CLUSTER</Cell>
                  <Cell>DATASETS</Cell>
                  <Cell>MODELS</Cell>
                  <Cell></Cell>
                </Row>
                {this.props.clusters.map((cluster, i) => {
                  return (
                    <Row key={i}>
                      <Cell>{cluster.name}</Cell>
                      <Cell>N/A</Cell>
                      <Cell>N/A</Cell>
                      <Cell>
                        <button className="default" onClick={this.selectCluster.bind(this, cluster.id)}>Connect</button>
                      </Cell>
                    </Row>
                  );
                })}
              </Table>
            </div>
          </div>
          <div className="connect-cluster">
            <h1>&hellip; or connect to a new H2O cluster</h1>
            <div>
              Connect to a H2O cluster where your existing models and data sets are located.
            </div>
            <form onSubmit={this.registerCluster.bind(this)}>
              <input type="text" name="ip-address" placeholder="IP Address"/>
              <input type="text" name="port" placeholder="Port"/>
              <button type="submit" className="default">Connect</button>
            </form>
          </div>
        </div>
        {!_.isEmpty(this.props.models) ? <div>
          <h1>2. Pick Models to Import</h1>
          <div>
            Models in a project must share the same feature set and response column to enable comparison.
          </div>
          <Table className="import-models">
            <Row header={true}>
              <Cell>DATA SET</Cell>
              <Cell>RESPONSE COLUMN</Cell>
              <Cell>MODELS</Cell>
              <Cell></Cell>
            </Row>
            {this.props.models.map((model, i) => {
              console.log(model);
              return (
                <Row key={i}>
                  <Cell>{model.name}</Cell>
                  <Cell>{model.response_column_name}</Cell>
                  <Cell>N/A</Cell>
                  <Cell>
                    <input type="checkbox" name={model.name}/>&nbsp; Select for Import
                  </Cell>
                </Row>
              );
            })}
          </Table>
        </div> : null}
        {!_.isEmpty(this.props.models) ? <div className="name-project">
          <h1>3. Name Project</h1>
          <div>
            <input ref="projectName" type="text"/>
          </div>
        </div> : null}
        {!_.isEmpty(this.props.models) ? <div>
          <button className="default" onClick={this.createProject.bind(this)}>Create Project</button>
        </div> : null}
      </div>
    );
  }
}

function mapStateToProps(state): any {
  return {
    clusters: state.projects.clusters,
    models: state.projects.models,
    project: state.project
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    fetchClusters: bindActionCreators(fetchClusters, dispatch),
    fetchModelsFromCluster: bindActionCreators(fetchModelsFromCluster, dispatch),
    createProjectAndImportModelsFromCluster: bindActionCreators(createProjectAndImportModelsFromCluster, dispatch),
    importModelFromCluster: bindActionCreators(importModelFromCluster, dispatch),
    registerCluster: bindActionCreators(registerCluster, dispatch)
  };
}


export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ImportNewProject);
