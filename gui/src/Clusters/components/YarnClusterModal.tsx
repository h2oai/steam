/**
 * Created by justin on 9/6/16.
 */

import * as React from 'react';
import DefaultModal from '../../App/components/DefaultModal';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import PageHeader from '../../Projects/components/PageHeader';
import { connect } from 'react-redux';
import { uploadEngine, startYarnCluster } from '../actions/clusters.actions';
import { bindActionCreators } from 'redux';

interface Props {
  open: boolean
}

interface DispatchProps {
  uploadEngine: Function,
  startYarnCluster: Function
}

export class YarnClusterModal extends React.Component<Props & DispatchProps, any> {
  refs: {
    [key: string]: (Element);
    engineForm: (HTMLFormElement)
    clusterForm: (HTMLFormElement)
  };
  startCluster(event) {
    event.preventDefault();
    let clusterName = this.refs.clusterForm.querySelector('input[name="name"]').value;
    let engineId = this.refs.clusterForm.querySelector('input[name="engineId"]').value;
    let size = this.refs.clusterForm.querySelector('input[name="size"]').value;
    let memory = this.refs.clusterForm.querySelector('input[name="memory"]').value;
    let username = this.refs.clusterForm.querySelector('input[name="username"]').value;
    this.props.startYarnCluster(clusterName, parseInt(engineId, 10), parseInt(size, 10), memory, username);
  }

  uploadEngine(event) {
    event.preventDefault();
    console.log(this.refs.engineForm);
    this.props.uploadEngine(this.refs.engineForm);
  }

  render(): React.ReactElement<DefaultModal> {
    return (
      <DefaultModal className="yarn-cluster-modal" open={this.props.open}>
        <PageHeader>LAUNCH YARN CLUSTER</PageHeader>
        <form ref="engineForm">
          <input type="file" name="engine"/>
          <button onClick={this.uploadEngine.bind(this)}>Upload Engine</button>
        </form>
        <form ref="clusterForm" onSubmit={this.startCluster.bind(this)}>
          <Table>
            <Row>
              <Cell>
                Name
              </Cell>
              <Cell>
                <input type="text" name="name"/>
              </Cell>
            </Row>
            <Row>
              <Cell>
                Engine ID
              </Cell>
              <Cell>
                <input type="text" name="engineId"/>
              </Cell>
            </Row>
            <Row>
              <Cell>
                Virtual Cores
              </Cell>
              <Cell>
                <input type="text" name="size"/>
              </Cell>
            </Row>
            <Row>
              <Cell>
                Memory
              </Cell>
              <Cell>
                <input type="text" name="memory"/>
              </Cell>
            </Row>
            <Row>
              <Cell>
                Username
              </Cell>
              <Cell>
                <input type="text" name="username"/>
              </Cell>
            </Row>
          </Table>
          <button type="submit" className="default">Start</button>
        </form>
      </DefaultModal>
    );
  }
}

function mapStateToProps(state) {
  return {};
}

function mapDispatchToProps(dispatch) {
  return {
    uploadEngine: bindActionCreators(uploadEngine, dispatch),
    startYarnCluster: bindActionCreators(startYarnCluster, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(YarnClusterModal);
