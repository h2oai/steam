/**
 * Created by justin on 7/18/16.
 */
import * as React from 'react';
import * as $ from 'jquery';
import * as _ from 'lodash';
import * as classNames from 'classnames';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Table from './Table';
import Row from './Row';
import Cell from './Cell';
import {
  fetchClusters, fetchModelsFromCluster,
  importModelFromCluster, createProjectAndImportModelsFromCluster, registerCluster, fetchDatasetsFromCluster
} from '../actions/projects.actions';
import { Cluster, Model, Dataset } from '../../Proxy/Proxy';
import '../styles/importnewproject.scss';
import { hashHistory } from 'react-router';

interface DispatchProps {
  fetchClusters: Function,
  fetchModelsFromCluster: Function,
  importModelFromCluster: Function,
  createProjectAndImportModelsFromCluster: Function,
  registerCluster: Function,
  fetchDatasetsFromCluster: Function
}

interface Props {
  clusters: Cluster[],
  models: Model[],
  datasets: Dataset[]
}

export class ImportNewProject extends React.Component<DispatchProps & Props, any> {
  refs: {
    [key: string]: (Element);
    projectName: (HTMLInputElement)
  };

  constructor() {
    super();
    this.state = {
      clusterId: null,
      datasetId: null,
      modelCategory: null,
      isModelSelected: false
    };
  }

  componentWillMount(): void {
    if (_.isEmpty(this.props.clusters)) {
      this.props.fetchClusters();
    }
  }

  selectDataset(event): void {
    this.setState({
      datasetId: event.target.value
    });
    this.props.fetchModelsFromCluster(this.state.clusterId, event.target.value);
  }

  createProject(): void {
    let name = $(this.refs.projectName).val();
    let importModels = [];
    let checkedModels = $('.import-models input:checked');
    if (checkedModels.length > 0) {
      checkedModels.map((i, input) => {
        importModels.push($(input).prop('name'));
      });
      this.props.createProjectAndImportModelsFromCluster(name, this.state.clusterId, this.state.modelCategory, importModels).then((res) => {
        hashHistory.push('/projects/' + res + '/models');
      });
    }
  }

  registerCluster(event) {
    event.preventDefault();
    let ipAddress = $(event.target).find('input[name="ip-address"]').val();
    let port = $(event.target).find('input[name="port"]').val();
    this.props.registerCluster(ipAddress + ':' + port);
  }

  selectModel() {
    let checkedModels = $('.import-models input:checked');
    if (checkedModels) {
      this.setState({
        isModelSelected: true
      });
    } else {
      this.setState({
        isModelSelected: false
      });
    }
  }

  selectCategory(event) {
    this.setState({
      modelCategory: event.target.value
    });
  }

  retrieveClusterDataframes(clusterId: number) {
    this.setState({
      clusterId: clusterId
    });
    this.props.fetchDatasetsFromCluster(clusterId);
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
                        <button className="default" onClick={this.retrieveClusterDataframes.bind(this, cluster.id)}>
                          Connect
                        </button>
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
        {this.state.clusterId ? <div>
          <h1>2. Select Dataset</h1>
          <div className="select-dataset">
            <div>
              <div>Frame</div>
              <select onChange={this.selectDataset.bind(this)}>
                <option></option>
                {this.props.datasets ? this.props.datasets.map((dataset, i) => {
                  return <option key={i} value={dataset.frame_name}>{dataset.name}</option>;
                }) : null}
              </select>
            </div>
            <div>
              <div>Category</div>
              <select onChange={this.selectCategory.bind(this)}>
                <option></option>
                {this.props.models ? _.uniqBy(this.props.models, 'model_category').map((model, i) => {
                  return <option key={i} value={model.model_category}>{model.model_category}</option>;
                }) : null}
              </select>
            </div>
          </div>
          <div>
          </div>
        </div> : null}
        {!_.isEmpty(this.props.models) && this.state.modelCategory ? <div>
          <h1>3. Pick Models to Import</h1>
          <div>
            Models in a project must share the same feature set and response column to enable comparison.
          </div>
          <Table className="import-models">
            <Row header={true}>
              <Cell>MODEL</Cell>
              <Cell>RESPONSE COLUMN</Cell>
              <Cell>CATEGORICAL</Cell>
              <Cell></Cell>
            </Row>
            {_.filter(this.props.models, model => model.model_category === this.state.modelCategory).map((model, i) => {
              return (
                <Row key={i}>
                  <Cell>{model.name}</Cell>
                  <Cell>{model.response_column_name}</Cell>
                  <Cell>{model.model_category}</Cell>
                  <Cell>
                    <input type="checkbox" name={model.name} onChange={this.selectModel.bind(this, model)}/>&nbsp; Select for Import
                  </Cell>
                </Row>
              );
            })}
          </Table>
        </div> : null}
        {!_.isEmpty(this.props.models && this.state.modelCategory) ? <div className="name-project">
          <h1>4. Name Project</h1>
          <div>
            <input ref="projectName" type="text"/>
          </div>
        </div> : null}
        {!_.isEmpty(this.props.models) && this.state.modelCategory ? <div>
          <button className={classNames('default', {disabled: !this.state.isModelSelected})}
                  onClick={this.createProject.bind(this)}>Create Project
          </button>
        </div> : null}
      </div>
    );
  }
}

function mapStateToProps(state): any {
  return {
    clusters: state.projects.clusters,
    models: state.projects.models,
    datasets: state.projects.datasets,
    project: state.project
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    fetchClusters: bindActionCreators(fetchClusters, dispatch),
    fetchModelsFromCluster: bindActionCreators(fetchModelsFromCluster, dispatch),
    createProjectAndImportModelsFromCluster: bindActionCreators(createProjectAndImportModelsFromCluster, dispatch),
    importModelFromCluster: bindActionCreators(importModelFromCluster, dispatch),
    registerCluster: bindActionCreators(registerCluster, dispatch),
    fetchDatasetsFromCluster: bindActionCreators(fetchDatasetsFromCluster, dispatch)
  };
}


export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ImportNewProject);
