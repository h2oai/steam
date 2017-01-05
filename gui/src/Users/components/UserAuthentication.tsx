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
import { connect } from 'react-redux';
import '../styles/users.scss';
import { Collapse } from '@blueprintjs/core/dist/components/collapse/collapse';
import { Button } from '@blueprintjs/core/dist/components/button/buttons';
import { Tooltip } from '@blueprintjs/core/dist/components/tooltip/tooltip';
import findDOMNode = ReactDOM.findDOMNode;

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
    this.setState({
      isLDAPConnectionSettingsOpen: !this.state.isLDAPConnectionSettingsOpen
    });
  };

  onResetClicked = () => {

  };
  onSaveConfigClicked = (e) => {
    e.preventDefault();
    console.log(findDOMNode(this.refs["hostInput"])["value"]);
  };
  onDBChanged = (e) => {
    console.log(e);
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
                <select ref="dbSelect" onChange={this.onDBChanged}>
                  <option>LDAP</option>
                  <option>Steam Local DB</option>
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
                  <input type="text" ref="hostInput"></input>
                </td>
              </tr>
              <tr className="auth-row">
                <td className="auth-left">PORT &nbsp; <Tooltip className="steam-tooltip-launcher" content="The LDAP server port">
                  <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                </Tooltip></td>
                <td className="auth-right">
                  <input type="text" value="689"></input>
                </td>
              </tr>
              <tr className="auth-row">
                <td className="auth-left">SSL-ENABLED</td>
                <td className="auth-right">
                  <input type="checkbox" defaultChecked={true}></input>
                </td>
              </tr>
              <tr className="auth-row">
                <td className="auth-left">CONNECTION ORDER &nbsp; <Tooltip className="steam-tooltip-launcher" content="The order in which Steam will query this LDAP server (among other enabled servers)">
                  <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                </Tooltip></td>
                <td className="auth-right">
                  <input type="text" value="1"></input>
                </td>
              </tr>
              <tr className="auth-row">
                <td className="auth-left">BIND DN &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>'Distinguished name' used to bind to LDAP server if extended access is needed.<br /> Leave this blank if anonymous bind is sufficient.</div>}>
                  <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                </Tooltip></td>
                <td className="auth-right">
                  <input type="text"></input>
                </td>
              </tr>
              <tr className="auth-row">
                <td className="auth-left">BIND DN PASSWORD &nbsp; <Tooltip className="steam-tooltip-launcher" content="Password for the Bind DN user">
                  <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                </Tooltip></td>
                <td className="auth-right">
                  <input type="password"></input>
                </td>
              </tr>
              <tr className="auth-row">
                <td className="auth-left">CONFIRM PASSWORD</td>
                <td className="auth-right">
                  <input type="password"></input>
                </td>
              </tr>
              <tr className="auth-row">
                <td className="auth-left">USER BASE DN &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The location of your LDAP users, specified by the DN of your user subtree.<br/> If necessary, you can specify several DNs separated by semicolons.</div>}>
                  <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                </Tooltip></td>
                <td className="auth-right">
                  <input type="text"></input>
                </td>
              </tr>
              <tr className="auth-row">
                <td className="auth-left">USER BASE FILTER &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The LDAP search filter used to filter users.<br/> Highly recommended if you have a large amount of user entries under your user base DN.<br/> For example, '(department=IT)'</div>}>
                  <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                </Tooltip></td>
                <td className="auth-right">
                  <input type="text"></input>
                </td>
              </tr>
              <tr className="auth-row">
                <td className="auth-left">USER NAME ATTRIBUTE &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The user attribute that contains the username.<br/> Note that this attribute's value should be case insensitive.<br/> Set to 'uid' for most configurations. In Active Directory (AD), this should be set to 'sAMAccountName'.</div>}>
                  <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                </Tooltip></td>
                <td className="auth-right">
                  <input type="text"></input>
                </td>
              </tr>
              <tr className="auth-row">
                <td className="auth-left">REAL NAME ATTRIBUTE &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The user attribute that contains a human readable name.<br/> This is typically 'cn' (common name) or 'displayName'.</div>}>
                  <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                </Tooltip></td>
                <td className="auth-right">
                  <input type="text" value="cn"></input>
                </td>
              </tr>
            </tbody>
          </table>

        </Collapse>

        <div id="actionButtonsContainer" className="space-20">
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
