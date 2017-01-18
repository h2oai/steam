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
import { hasPermissionToShow } from "../../App/utils/permissions";
import { Popover, PopoverInteractionKind, Position } from "@blueprintjs/core";

interface Props {
  engines: any,
  config: any,
  isAdmin: boolean
  clusterLaunchIsInProgress: boolean
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
      selectedEngine: null,
      engineDropdownOpen: null
    };
  }

  componentDidMount() {
    this.props.getEngines();
    this.props.getConfig();
  }

  startCluster(event) {
    event.preventDefault();
    let clusterName = (this.refs.clusterForm.querySelector('input[name="name"]') as HTMLInputElement).value;
    let engineId = this.state.selectedEngine.id;
    let size = (this.refs.clusterForm.querySelector('input[name="size"]') as HTMLInputElement).value;
    let memory = (this.refs.clusterForm.querySelector('input[name="memory"]') as HTMLInputElement).value;
    let keytab = _.get((this.refs.clusterForm.querySelector('input[name="keytab"]') as HTMLInputElement), 'value', '');
    // let secure = (this.refs.clusterForm.querySelector('input[name="secure"]') as HTMLInputElement).checked;
    let secure = true;
    this.props.startYarnCluster(clusterName, parseInt(engineId, 10), parseInt(size, 10), memory + 'g', secure, keytab);
  }

  uploadEngine(event) {
    event.preventDefault();
    this.props.uploadEngine(this.refs.engine);
    this.setState({
      engineDropdownOpen: true
    });
  };

  onEngineUploadDialogOpened = () => {
    this.setState({
      engineDropdownOpen: true
    });
  };

  render() {

    let uploadEngine;
    if (hasPermissionToShow("ManageEngine", this.props.config, this.props.isAdmin)) {
      uploadEngine = (
        <label className="pt-file-upload" onClick={this.onEngineUploadDialogOpened} >
          <input ref="engine" type="file" onChange={this.uploadEngine.bind(this)} />
          <span className="pt-file-upload-input">Upload New Engine...</span>
        </label>
      );
    } else {
      if (this.props.engines && this.props.engines.length > 0) {
        uploadEngine = null;
      } else {
        uploadEngine = <div>Please request your Admin to upload more engines</div>;
      }
    }

    let currentEngineString;
    if (this.state.selectedEngine) {
      currentEngineString = this.state.selectedEngine.name;
    } else {
      currentEngineString = "Please select an engine";
    }

    let engineSelectContent = (
      <div>
        {this.props.engines.map((engine, i) => {
          return <div key={i} className="pt-menu-item pt-popover-dismiss" onClick={() => this.setState({ selectedEngine: engine })}>
            {engine.name}
          </div>;
        })}
        {uploadEngine}
      </div>
    );

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
                H2O VERSION
              </Cell>
              <Cell>

                <Popover content={engineSelectContent}
                         interactionKind={PopoverInteractionKind.CLICK}
                         popoverClassName="pt-popover-content-sizing"
                         position={Position.BOTTOM}
                         useSmartPositioning={true}>
                  <div className="pt-button">
                    { currentEngineString } &nbsp;
                    <span className="pt-icon-standa pt-icon-caret-down pt-align-right font-18"></span>
                  </div>
                </Popover>
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
          {this.props.clusterLaunchIsInProgress ? null : <button type="submit" className="button-primary">Launch New Clusters</button>}
          {this.props.clusterLaunchIsInProgress ? <div className="pt-spinner .modifier">
            <div className="pt-spinner-svg-container">
              <svg viewBox="0 0 100 100">
                <path className="pt-spinner-track" d="M 50,50 m 0,-44.5 a 44.5,44.5 0 1 1 0,89 a 44.5,44.5 0 1 1 0,-89"></path>
                <path className="pt-spinner-head" d="M 94.5 50 A 44.5 44.5 0 0 0 50 5.5"></path>
              </svg>
            </div>
          </div> : null}
        </form>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    engines: state.clusters.engines,
    config: state.clusters.config,
    isAdmin: state.global.isAdmin,
    clusterLaunchIsInProgress: state.clusters.clusterLaunchIsInProgress
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
