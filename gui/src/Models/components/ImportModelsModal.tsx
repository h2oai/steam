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
 * Created by justin on 8/6/16.
 */
import * as React from 'react';
import DefaultModal from '../../App/components/DefaultModal';
import PageHeader from '../../Projects/components/PageHeader';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import MojoPojoSelector from '../../Projects/components/MojoPojoSelector';
import { bindActionCreators } from 'redux';
import {
  fetchClusters, fetchModelsFromCluster,
  importModelsFromCluster
} from '../../Projects/actions/projects.actions';
import { connect } from 'react-redux';
import '../styles/importmodelsmodal.scss';
import { Cluster, Model } from '../../Proxy/Proxy';

interface Props {
  open: boolean,
  onCancel: Function,
  projectId: string,
  clusters: Cluster[],
  models: Model[],
  fetchLeaderboard: Function,
  modelCategory: string,
  datasetName: string
}

interface DispatchProps {
  fetchClusters: Function,
  fetchModelsFromCluster: Function,
  importModelsFromCluster: Function
}

export class ImportModelsModal extends React.Component<Props & DispatchProps, any> {
  constructor() {
    super();
    this.state = {
      clusterId: null,

      models: []
    };
  }

  componentWillMount() {
    this.props.fetchClusters(parseInt(this.props.projectId, 10));
  }

  onChange(event) {
    this.setState({
      clusterId: event.target.value
    });
    this.props.fetchModelsFromCluster(parseInt(event.target.value, 10), this.props.datasetName);
  }

  importModelsFromCluster(event) {
    event.preventDefault();
    let inputs = event.target.querySelectorAll('input:checked');
    let models = [];
    for (var i = 0; i < inputs.length; i++) {
      models.push(inputs[i].value);
    }
    this.props.importModelsFromCluster(parseInt(this.state.clusterId, 10), parseInt(this.props.projectId, 10), models).then(() => {
      this.props.fetchLeaderboard(parseInt(this.props.projectId, 10), this.props.modelCategory);
      this.props.onCancel();
    });
  }

  render() {
    return (
      <DefaultModal className="import-modal" open={this.props.open}>
        <PageHeader>IMPORT MODELS</PageHeader>
        <form onSubmit={this.importModelsFromCluster.bind(this)}>
          <Table className="outer-table">
            <Row>
              <Cell>
                By default, Steam picks the most optimized model format for you to import. Advanced users can choose your own model type&nbsp;
                <MojoPojoSelector></MojoPojoSelector>.
              </Cell>
            </Row>
            <Row>
              <Cell>
                CLUSTER
              </Cell>
              <Cell>
                <div>Select the cluster to import models from</div>
                <select onChange={this.onChange.bind(this)}>
                  <option value=""></option>
                  {this.props.clusters.map((cluster, i) => {
                    return <option key={i} value={cluster.id}>{cluster.name} @ {cluster.address}</option>;
                  })}
                </select>
              </Cell>
            </Row>
            <Row>
              <Cell>
                SELECT MODELS
              </Cell>
              <Cell>

                <Table className="inner-table">
                  <Row>
                    <Cell>MODEL</Cell>
                    <Cell>DATAFRAME</Cell>
                    <Cell>RESPONSE COLUMN</Cell>
                    <Cell/>
                  </Row>
                  {this.props.models.map((model, i) => {
                    return (
                      <Row key={i}>
                        <Cell>{model.name}</Cell>
                        <Cell>{model.dataset_name}</Cell>
                        <Cell>{model.response_column_name}</Cell>
                        <Cell><input type="checkbox" value={model.model_key}/></Cell>
                      </Row>
                    );
                  })}
                </Table>
              </Cell>
            </Row>
            <Row>
              <Cell/>
              <Cell className="button-container">
                <button type="submit" className="button-primary">Import</button>
                <button type="button" onClick={this.props.onCancel.bind(this)} className="button-secondary">Cancel</button>
              </Cell>
            </Row>
          </Table>
        </form>
      </DefaultModal>
    );
  }
}

function mapStateToProps(state) {
  return {
    clusters: state.projects.clusters,
    models: state.projects.models
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchClusters: bindActionCreators(fetchClusters, dispatch),
    fetchModelsFromCluster: bindActionCreators(fetchModelsFromCluster, dispatch),
    importModelsFromCluster: bindActionCreators(importModelsFromCluster, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ImportModelsModal);
