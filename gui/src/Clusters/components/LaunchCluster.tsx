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
    let secure = (this.refs.clusterForm.querySelector('input[name="secure"]') as HTMLInputElement).checked;

    this.props.startYarnCluster(clusterName, parseInt(engineId, 10), parseInt(size, 10), memory + 'g', secure, keytab);
  }

  uploadEngine(event) {
    event.preventDefault();
    this.props.uploadEngine(this.refs.engine);
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
                <NumericInput name="memory" min="1"/>GB
              </Cell>
            </Row>
            <Row>
              <Cell>
                SECURE
              </Cell>
              <Cell>
                <div className="checkbox">
                  <input type="checkbox" name="secure" />
                </div>
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
            {_.get(this.props.config, 'kerberos_enabled', false) ? <Row>
              <Cell>
                Kerberos Keytab
              </Cell>
              <Cell>
                <input type="text" name="keytab"/>
              </Cell>
            </Row> : null}
          </Table>
          <button type="submit" className="button-primary">Launch New Clusters</button>
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
