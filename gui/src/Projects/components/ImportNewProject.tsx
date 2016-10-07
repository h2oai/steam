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
  fetchClusters, fetchModelsFromCluster, resetClusterSelection,
  importModelFromCluster, createProjectAndImportModelsFromCluster, registerCluster, fetchDatasetsFromCluster
} from '../actions/projects.actions';
import { Cluster, Model, Dataset } from '../../Proxy/Proxy';
import '../styles/importnewproject.scss';
import { hashHistory } from 'react-router';
import InputFeedback from '../../App/components/InputFeedback';
import { FeedbackType } from '../../App/components/InputFeedback';

interface DispatchProps {
  fetchClusters: Function,
  fetchModelsFromCluster: Function,
  importModelFromCluster: Function,
  createProjectAndImportModelsFromCluster: Function,
  registerCluster: Function,
  fetchDatasetsFromCluster: Function,
  resetClusterSelection: Function
}

interface Props {
  clusters: Cluster[],
  models: Model[],
  datasets: Dataset[],
  isClusterFetchInProcess: boolean,
  isModelFetchInProcess: boolean
  registerClusterError: string
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
    if (event.target.value) {
      this.props.fetchModelsFromCluster(this.state.clusterId, event.target.value);
    } else {
      this.setState({
        modelCategory: null
      });
    }
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

  resetClusterSelection(event) {
    event.preventDefault();
    this.setState({
      clusterId: null
    });
    this.props.resetClusterSelection();
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
    var selectedClusterName;
    var selectedClusterAddress;

    for (let cluster of this.props.clusters) {
      if (cluster.id === this.state.clusterId) {
        selectedClusterName = cluster.name;
        selectedClusterAddress = cluster.address;
      }
    }

    if (!this.props.clusters) {
      return <div></div>;
    }
    return (
      <div className="import-new-project">
        <div className="step-1">
          <div className="select-cluster">
            <h2>1. Select H2O cluster</h2>
            { this.state.clusterId ?
              <div className="cluster-info intro">
                <span><i className="fa fa-cubes cluster-image"/></span>
                <div className="cluster-details">
                  <div>{selectedClusterName}</div>
                  <div>{selectedClusterAddress}</div>
                  <div onClick={this.resetClusterSelection.bind(this)} className="select-new-cluster"><i className="fa fa-close"/> use a different cluster</div>
                </div>
              </div> :
              <div className="intro">
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
                        <Cell><span className="name-cell">{cluster.name}</span></Cell>
                        <Cell>N/A</Cell>
                        <Cell>N/A</Cell>
                        <Cell>
                          <button className="button-primary" onClick={this.retrieveClusterDataframes.bind(this, cluster.id)}>
                            Connect
                          </button>
                        </Cell>
                      </Row>
                    );
                  })}
                </Table>
              </div>
            }
          </div>
          { !this.state.clusterId ?
            <div className="connect-cluster">
              <h2>&hellip; or connect to a new H2O cluster</h2>
              <div className="intro">
                Connect to a H2O cluster where your existing models and data sets are located.
              </div>
              <form onSubmit={this.registerCluster.bind(this)}>
                <input type="text" name="ip-address" placeholder="IP Address"/>
                { this.props.registerClusterError ?
                  <InputFeedback message={ this.props.registerClusterError } type={FeedbackType.Error} />
                  : null }
                <input type="text" name="port" placeholder="Port"/>
                <button type="submit" className="button-primary">Connect</button>
              </form>
              { this.props.isClusterFetchInProcess ?
                  <InputFeedback message="Connecting..." type={FeedbackType.Progress} />
                 : null }
            </div>
          : null }
        </div>
        {this.state.clusterId ? <div>
          <h2>2. Select Dataframe</h2>
          <div className="intro">
            <select name="selectDataframe" onChange={this.selectDataset.bind(this)}>
              <option></option>
              {this.props.datasets ? this.props.datasets.map((dataset, i) => {
                return <option key={i} value={dataset.frame_name}>{dataset.name}</option>;
              }) : null}
            </select>
            { this.props.isModelFetchInProcess ?
                <InputFeedback message="Connecting..." type={FeedbackType.Progress} />
              : null }
          </div>
        </div> : null}
        {(this.state.datasetId && !this.props.isModelFetchInProcess) ?
          <div>
            <h2>3. Select Model Category</h2>
            <div className="intro">
              <select name="selectModelCategory" onChange={this.selectCategory.bind(this)}>
                <option></option>
                {this.props.models ? _.uniqBy(this.props.models, 'model_category').map((model, i) => {
                  return <option key={i} value={model.model_category}>{model.model_category}</option>;
                }) : null}
              </select>
            </div>
          </div> : null}
        {this.state.datasetId && !_.isEmpty(this.props.models) && this.state.modelCategory ? <div>
          <h2>4. Pick Models to Import</h2>
          <div className="intro">
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
        {this.state.datasetId && !_.isEmpty(this.props.models && this.state.modelCategory) ? <div className="name-project">
          <h2>5. Name Project</h2>
          <div className="intro">
            <input ref="projectName" type="text"/>
          </div>
        </div> : null}
        {this.state.datasetId && !_.isEmpty(this.props.models) && this.state.modelCategory ? <div>
          <button className={classNames('button-primary', {disabled: !this.state.isModelSelected})}
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
    project: state.project,
    isClusterFetchInProcess: state.projects.isClusterFetchInProcess,
    isModelFetchInProcess: state.projects.isModelFetchInProcess,
    registerClusterError: state.projects.registerClusterError
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    fetchClusters: bindActionCreators(fetchClusters, dispatch),
    fetchModelsFromCluster: bindActionCreators(fetchModelsFromCluster, dispatch),
    createProjectAndImportModelsFromCluster: bindActionCreators(createProjectAndImportModelsFromCluster, dispatch),
    importModelFromCluster: bindActionCreators(importModelFromCluster, dispatch),
    registerCluster: bindActionCreators(registerCluster, dispatch),
    fetchDatasetsFromCluster: bindActionCreators(fetchDatasetsFromCluster, dispatch),
    resetClusterSelection: bindActionCreators(resetClusterSelection, dispatch)
  };
}


export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ImportNewProject);
