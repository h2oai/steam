/**
 * Created by justin on 9/8/16.
 */

import * as React from 'react';
import { connect } from 'react-redux';
import { startYarnCluster, uploadEngine, getEngines } from '../actions/clusters.actions';
import { bindActionCreators } from 'redux';
import Cell from '../../Projects/components/Cell';
import Row from '../../Projects/components/Row';
import PageHeader from '../../Projects/components/PageHeader';
import Table from '../../Projects/components/Table';
import '../styles/launchcluster.scss';
import { NumericInput } from 'h2oUIKit';

interface Props {

}

interface DispatchProps {
  startYarnCluster: Function,
  uploadEngine: Function
}

export class LaunchCluster extends React.Component<Props & DispatchProps, any> {
  refs: {
    [key: string]: (Element);
    engineForm: (HTMLFormElement)
    clusterForm: (HTMLFormElement)
  };

  constructor() {
    super();
    this.state = {
      memorySizeUnit: 'm'
    };
  }

  startCluster(event) {
    event.preventDefault();
    let clusterName = (this.refs.clusterForm.querySelector('input[name="name"]') as HTMLInputElement).value;
    let engineId = (this.refs.clusterForm.querySelector('input[name="engineId"]') as HTMLInputElement).value;
    let size = (this.refs.clusterForm.querySelector('input[name="size"]') as HTMLInputElement).value;
    let memory = (this.refs.clusterForm.querySelector('input[name="memory"]') as HTMLInputElement).value;
    this.props.startYarnCluster(clusterName, parseInt(engineId, 10), parseInt(size, 10), memory + this.state.memorySizeUnit);
  }

  uploadEngine(event) {
    event.preventDefault();
    console.log(this.refs.engineForm);
    this.props.uploadEngine(this.refs.engineForm);
  }

  onChangeMemory(event) {
    this.setState({
      memorySizeUnit: event.target.value
    });
  }

  render() {
    return (
      <div className="launch-cluster">
        <PageHeader>LAUNCH NEW CLUSTER</PageHeader>
        <form ref="clusterForm" onSubmit={this.startCluster.bind(this)}>
          <Table>
            <Row header={true}/>
            <Row>
              <Cell>
                CLUSTER NAME
              </Cell>
              <Cell>
                <input type="text" name="name"/>
              </Cell>
            </Row>
            <Row>
              <Cell>
                ENGINE ID
              </Cell>
              <Cell>
                <form ref="engineForm">
                  <input type="file" name="engine"/>
                  <button className="default" onClick={this.uploadEngine.bind(this)}>Upload Engine</button>
                </form>
                <NumericInput name="engineId" min="1"/>
              </Cell>
            </Row>
            <Row>
              <Cell>
                NUMBER OF NODES
              </Cell>
              <Cell>
                <NumericInput name="size" min="1"/>
              </Cell>
            </Row>
            <Row>
              <Cell>
                MEMORY PER NODE
              </Cell>
              <Cell>
                <NumericInput name="memory" min="1"/>
                <select onChange={this.onChangeMemory.bind(this)}>
                  <option value="m">MB</option>
                  <option value="g">GB</option>
                </select>
              </Cell>
            </Row>
          </Table>
          <button type="submit" className="default">Launch New Clusters</button>
        </form>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {};
}

function mapDispatchToProps(dispatch) {
  return {
    uploadEngine: bindActionCreators(uploadEngine, dispatch),
    startYarnCluster: bindActionCreators(startYarnCluster, dispatch),
    getEngines: bindActionCreators(getEngines, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(LaunchCluster);
