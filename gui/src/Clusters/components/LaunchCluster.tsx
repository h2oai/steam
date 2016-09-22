/**
 * Created by justin on 9/8/16.
 */

import * as React from 'react';
import * as _ from 'lodash';
import { connect } from 'react-redux';
import { startYarnCluster, uploadEngine, getEngines, getConfig } from '../actions/clusters.actions';
import { bindActionCreators } from 'redux';
import Cell from '../../Projects/components/Cell';
import Row from '../../Projects/components/Row';
import PageHeader from '../../Projects/components/PageHeader';
import Table from '../../Projects/components/Table';
import '../styles/launchcluster.scss';
import { NumericInput } from 'h2oUIKit';

interface Props {
  engines: any,
  config: any
}

interface DispatchProps {
  startYarnCluster: Function,
  uploadEngine: Function,
  getEngines: Function,
  getConfig: Function
}

export class LaunchCluster extends React.Component<Props & DispatchProps, any> {
  refs: {
    [key: string]: (Element);
    engine: (HTMLInputElement)
    clusterForm: (HTMLFormElement)
    engineList: (HTMLSelectElement)
  };

  constructor() {
    super();
    this.state = {
      memorySizeUnit: 'm',
      engineId: null
    };
  }

  componentDidMount() {
    this.props.getEngines();
    this.props.getConfig();
  }

  startCluster(event) {
    event.preventDefault();
    let clusterName = (this.refs.clusterForm.querySelector('input[name="name"]') as HTMLInputElement).value;
    let engineId = this.state.engineId;
    let size = (this.refs.clusterForm.querySelector('input[name="size"]') as HTMLInputElement).value;
    let memory = (this.refs.clusterForm.querySelector('input[name="memory"]') as HTMLInputElement).value;
    let keytab = _.get((this.refs.clusterForm.querySelector('input[name="keytab"]') as HTMLInputElement), 'value', '');
    this.props.startYarnCluster(clusterName, parseInt(engineId, 10), parseInt(size, 10), memory + this.state.memorySizeUnit, keytab);
  }

  uploadEngine(event) {
    event.preventDefault();
    this.props.uploadEngine(this.refs.engine);
  }

  onChangeMemory(event) {
    this.setState({
      memorySizeUnit: event.target.value
    });
  }

  onChangeEngine(event) {
    this.setState({
      engineId: event.target.value
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
                <select className="memory-selection" onChange={this.onChangeMemory.bind(this)}>
                  <option value="m">MB</option>
                  <option value="g">GB</option>
                </select>
              </Cell>
            </Row>
            <Row>
              <Cell>
                H2O VERSION
              </Cell>
              <Cell>
                <div className="upload-engine">
                  <input ref="engine" type="file" name="engine"/>
                  <div className="button-primary" onClick={this.uploadEngine.bind(this)}>Upload Engine</div>
                </div>
                <select onChange={this.onChangeEngine.bind(this)}>
                  <option></option>
                  {this.props.engines.map((engine, i) => {
                    return <option key={i} value={engine.id}>{engine.name}</option>;
                  })}
                </select>
              </Cell>
            </Row>
            {_.get(this.props.config, 'kerberos_enabled', false) === true ? <Row>
              <Cell>
                Kerberos Keytab
              </Cell>
              <Cell>
                <input type="text" name="keytab"/>
              </Cell>
            </Row> : null}
          </Table>
          <div type="submit" className="button-primary">Launch New Clusters</div>
        </form>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    engines: state.clusters.engines,
    config: state.clusters.config
  };
}

function mapDispatchToProps(dispatch) {
  return {
    uploadEngine: bindActionCreators(uploadEngine, dispatch),
    startYarnCluster: bindActionCreators(startYarnCluster, dispatch),
    getEngines: bindActionCreators(getEngines, dispatch),
    getConfig: bindActionCreators(getConfig, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(LaunchCluster);
