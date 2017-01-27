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

import * as React from 'react';
import * as _ from 'lodash';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/users.scss';
import { } from "../../Proxy/Proxy";
import {Tooltip} from "@blueprintjs/core";
import { getConfig } from "../../Clusters/actions/clusters.actions";
import {Config} from "../../Proxy/Proxy";
import {saveGlobalKerberos, fetchGlobalKeytab, setGlobalKerberosEnabled} from "../actions/users.actions";
import {deleteKeytab, testKeytab} from "../../user/actions/user.actions";
import {Keytab} from "../../Proxy/Proxy";

interface Props {
  config: Config,
  globalKeytab: Keytab
}
interface DispatchProps {
  saveGlobalKerberos: Function,
  getConfig: Function
  fetchGlobalKeytab: Function,
  deleteKeytab: Function,
  testKeytab: Function,
  setGlobalKerberosEnabled: Function
}

export class GlobalKerberos extends React.Component<Props & DispatchProps, any> {

  refs: {
    [key: string]: (Element);
    file: HTMLInputElement
  };

  constructor(params) {
    super(params);
    this.state = {
      uploadText: "Upload New Keytab",
      steamPrincipleValue: ""
    };

  }

  componentWillMount() {
    this.props.getConfig();
  }

  componentWillReceiveProps = (nextProps) => {
    if (this.props.config && nextProps.config && JSON.stringify(this.props.config.kerberos_enabled) !== JSON.stringify(nextProps.config.kerberosEnabledValue)) {
      this.setState({
        kerberosEnabledValue : nextProps.config.kerberos_enabled
      });
    };

    if (this.props.config && nextProps.config && JSON.stringify(this.props.config) !== JSON.stringify(nextProps.config)) {
      this.populateValuesFromConfig(nextProps.config);
    };
  };

  populateValuesFromConfig(config: Config) {
    this.setState({
      kerberosEnabledValue: config.kerberos_enabled
    });
  };

  onDeleteKeytab = (id) => {
    this.props.deleteKeytab(id);
  };
  onTestConfigClicked = (id) => {
    this.props.testKeytab(id);
  };
  onNewKeytabSelected = (e) => {
    this.setState({
      uploadText: e.target.value
    });
    this.props.saveGlobalKerberos(this.refs.file, this.state.steamPrincipleValue);
  };

  render(): React.ReactElement<HTMLDivElement> {

    return (
      <div className="cluster-authentication intro">
        <table className="space-20">
          <tbody>
            <tr>
              <td className="auth-left">KERBEROS ENABLED</td>
              <td>
                <label className="pt-control pt-switch .modifier">
                  <input type="checkbox"
                         checked={this.props.config ? this.props.config.kerberos_enabled : false}
                         onChange={
                           (e: any) => {
                              this.props.setGlobalKerberosEnabled(this.props.config ? !this.props.config.kerberos_enabled : true);
                            }
                          }
                  />
                  <span className="pt-control-indicator"></span>
                </label>

                </td>
            </tr>
            {this.state.kerberosEnabledValue ? <tr>
              <td className="auth-left">STEAM PRINCIPLE &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>This is the kerberos principle used for steam monitoring and management</div>}>
                <i className="fa fa-question-circle-o" aria-hidden="true"></i>
              </Tooltip></td>
              <td>
                <input type="text"
                       className="pt-input ldap-input "
                       value={this.state.steamPrincipleValue}
                       onChange={(e: any) => this.setState({ "steamPrincipleValue": e.target.value })}
                ></input>
              </td>
            </tr> : null }
            {this.state.steamPrincipleValue !== "" && this.state.kerberosEnabledValue ? <tr>
              <td className="auth-left">PRINCIPLE KEYTAB</td>
              <td>
                <p>This keytab is used for the steam installation in the background. Personal principle keytabs for each Steam users are configured by themselves in Steam "User Preferences"</p>
                {this.props.globalKeytab ? <p>{this.props.globalKeytab.name} &nbsp; <i className="fa fa-trash" aria-hidden="true" onClick={() => this.onDeleteKeytab(this.props.globalKeytab.id)}></i></p> :
                <p>
                  <label className="pt-file-upload">
                    <input ref="file" type="file" onChange={(e) => this.onNewKeytabSelected(e)} />
                    <span className="pt-file-upload-input">{this.state.uploadText}</span>
                  </label>
                </p>}
              </td>
            </tr> : null }
          </tbody>
        </table>

        {this.props.globalKeytab ? <div id="actionButtonsContainer" className="space-20">
            <div>
              <div className="button-secondary" onClick={this.onTestConfigClicked}>Test Config</div>
            </div>
        </div> : null }

      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    config: state.clusters.config
  };
}

function mapDispatchToProps(dispatch) {
  return {
    saveGlobalKerberos: bindActionCreators(saveGlobalKerberos, dispatch),
    getConfig: bindActionCreators(getConfig, dispatch),
    fetchGlobalKeytab: bindActionCreators(fetchGlobalKeytab(), dispatch),
    deleteKeytab: bindActionCreators(deleteKeytab, dispatch),
    testKeytab: bindActionCreators(testKeytab, dispatch),
    setGlobalKerberosEnabled: bindActionCreators(setGlobalKerberosEnabled, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(GlobalKerberos);
