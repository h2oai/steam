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
import '../styles/user.scss';
import { saveLocalKerberos } from "../actions/user.actions";
import { } from "../../Proxy/Proxy";
import {Tooltip} from "@blueprintjs/core";

interface Props {
}
interface DispatchProps {
  saveLocalKerberos: Function
}

export class ClusterAuthentication extends React.Component<Props & DispatchProps, any> {

  refs: {
    [key: string]: (Element);
    keytab: (HTMLInputElement)
  };

  constructor(params) {
    super(params);
    this.state = {
      uploadText: "Upload New Keytab"
    };

  }

  componentWillMount() {
  }

  onTestConfigClicked = () => {

  };
  onSaveConfigClicked = () => {
    this.props.saveLocalKerberos(this.refs.keytab);
  };
  onNewKeytabSelected = (e) => {
    this.setState({
      uploadText: e.target.value
    });
  };

  render(): React.ReactElement<HTMLDivElement> {

    return (
      <div className="cluster-authentication intro">
        <table className="space-20">
          <tbody>
            <tr>
              <td className="auth-left">PRINCIPLE KEYTAB</td>
              <td>
                <p>Your principle keytab</p>
                <p>
                  <label className="pt-file-upload">
                    <input ref="keytab" type="file" onChange={(e) => this.onNewKeytabSelected(e)} />
                    <span className="pt-file-upload-input">{this.state.uploadText}</span>
                  </label>
                </p>
              </td>
            </tr>
          </tbody>
        </table>

        <div id="actionButtonsContainer" className="space-20">
            <div>
              <div className="button-secondary" onClick={this.onTestConfigClicked}>Test Config</div> &nbsp;
              <div className="button-primary" onClick={this.onSaveConfigClicked}>Save Config</div>
            </div>
        </div>

      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
  };
}

function mapDispatchToProps(dispatch) {
  return {
    saveLocalKerberos: bindActionCreators(saveLocalKerberos, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ClusterAuthentication);
