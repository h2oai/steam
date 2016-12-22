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
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/users.scss';
import { Collapse } from '@blueprintjs/core/dist/components/collapse/collapse';
import { Button } from '@blueprintjs/core/dist/components/button/buttons';
import { Tooltip } from '@blueprintjs/core/dist/components/tooltip/tooltip';

interface Props {
}
interface DispatchProps {
}

export class UserAuthentication extends React.Component<Props & DispatchProps, any> {

  constructor(params) {
    super(params);
    this.state = {
      isLDAPConnectionSettingsOpen: true
    };
  }

  componentWillMount() {
  }

  onShowLDAPConnectionSettingsClicked = () => {
    console.log("flipping");
    this.setState({
      isLDAPConnectionSettingsOpen: !this.state.isLDAPConnectionSettingsOpen
    });
  };

  onTestConfigClicked = () => {

  };
  onResetClicked = () => {

  };
  onSaveConfigClicked = () => {

  };

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="user-authentication">
        <div className="space-20">User DB Connection Settings</div>
        <table className="space-20">
          <tbody>
            <tr className="auth-row">
              <td className="auth-left">User DB Type</td>
              <td className="auth-right">
                <select>
                  <option>LDAP</option>
                  <option>Steam DB</option>
                </select>
              </td>
            </tr>
          </tbody>
        </table>

        <div className="colapse-header">
          <Button onClick={this.onShowLDAPConnectionSettingsClicked}>
            {this.state.isLDAPConnectionSettingsOpen ?
              <i className="fa fa-minus" aria-hidden="true"></i> :
              <i className="fa fa-plus" aria-hidden="true"></i> }
          </Button> &nbsp;
          LDAP Connection Settings
        </div>

        <Collapse isOpen={this.state.isLDAPConnectionSettingsOpen} className="space-20">
          <table>
            <tbody>
              <tr className="auth-row">
                <td className="auth-left">HOST &nbsp; <Tooltip className="steam-tooltip-launcher" content="LDAP host server address">
                  <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                </Tooltip></td>
                <td className="auth-right">
                  <input type="text">

                  </input>
                </td>
              </tr>
              <tr className="auth-row">
                <td className="auth-left">PORT &nbsp; <Tooltip className="steam-tooltip-launcher" content="The LDAP server port">
                  <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                </Tooltip></td>
                <td className="auth-right">
                  <input type="text"></input>
                </td>
              </tr>
              <tr className="auth-row">
                <td className="auth-left">SSL-ENABLED</td>
                <td className="auth-right">
                  <input type="checkbox"></input>
                </td>
              </tr>
            </tbody>
          </table>

        </Collapse>

        <div id="actionButtonsContainer" className="space-20">
          <div className="button-secondary" onClick={this.onTestConfigClicked}>Test Config</div>
          <div className="button-secondary" onClick={this.onResetClicked}>Reset</div>
          <div className="button-primary" onClick={this.onSaveConfigClicked}>Save Config</div>
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
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(UserAuthentication);
