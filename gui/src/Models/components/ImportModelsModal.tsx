/**
 * Created by justin on 8/6/16.
 */
import * as React from 'react';
import DefaultModal from '../../App/components/DefaultModal';
import PageHeader from '../../Projects/components/PageHeader';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
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
  modelCategory: string
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
    this.props.fetchModelsFromCluster(parseInt(event.target.value, 10), 'Key_Frame__DGA3.hex');
  }

  importModelsFromCluster(event) {
    event.preventDefault();
    let inputs = event.target.querySelectorAll('input:checked');
    console.log(inputs);
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
    console.log(this.props.models);
    return (
      <DefaultModal className="import-modal" open={this.props.open}>
        <PageHeader>IMPORT MODELS</PageHeader>
        <form onSubmit={this.importModelsFromCluster.bind(this)}>
          <Table className="outer-table">
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
                <button className="default" type="submit">Import</button>
                <button onClick={this.props.onCancel.bind(this)} className="default invert">Cancel</button>
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
