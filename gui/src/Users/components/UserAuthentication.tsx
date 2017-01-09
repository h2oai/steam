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
import {setLdap, LdapConfig} from "../../Proxy/Proxy";

interface Props {
}
interface DispatchProps {
}

export class UserAuthentication extends React.Component<Props & DispatchProps, any> {

  dbSelectInput: HTMLSelectElement;
  hostInput: HTMLInputElement;
  portInput: HTMLInputElement;
  sslEnabledInput: HTMLInputElement;
  connectionOrderInput: HTMLInputElement;
  bindDnInput: HTMLInputElement;
  bindDnPasswordInput: HTMLInputElement;
  confirmPasswordInput: HTMLInputElement;
  userbaseDnInput: HTMLInputElement;
  userbaseFilterInput: HTMLInputElement;
  usernameRNInput: HTMLInputElement;
  realNameAttributeInput: HTMLInputElement;

  constructor(params) {
    super(params);
    this.state = {
      isLDAPConnectionSettingsOpen: true,
      showDBOptions: true,
      hostInputValid: true,
      portInputValid: true,
      connectionOrderInputValid: true,
      passwordInputValid: true,
      userbaseFilterInputValid: true,
      usernameRNInputValid: true,
      realNameAttributeInputValid: true
    };
  }

  validateAll = () => {
    if (this.hostInput.value.length > 3) {
      this.setState({ hostInputValid: true });
    } else {
      this.setState({ hostInputValid: false });
    }

    if (this.portInput.value.length > 0) {
      this.setState({ portInputValid: true });
    } else {
      this.setState({ portInputValid: false });
    }

    if (this.connectionOrderInput.value.length > 0) {
      this.setState({ connectionOrderInputValid: true });
    } else {
      this.setState({ connectionOrderInputValid: false });
    }

    this.validatePasswords();

    if (this.userbaseFilterInput.value.length > 0) {
      this.setState({ userbaseFilterInput: true });
    } else {
      this.setState({ userbaseFilterInputValid: false});
    }

    if (this.usernameRNInput.value.length > 0) {
      this.setState({ usernameRNInputValid: true });
    } else {
      this.setState({ usernameRNInputValid: false });
    }

    if (this.realNameAttributeInput.value.length > 0) {
      this.setState({ realNameAttributeInputValid: true });
    } else {
      this.setState({ realNameAttributeInputValid: false });
    }

  };

  validatePasswords = () => {
    if (this.confirmPasswordInput.value.length > 0) {
      if (this.bindDnPasswordInput.value.length < 2) {
        this.state.passwordInputValid = false;
      } else {
        if (this.bindDnPasswordInput.value === this.confirmPasswordInput.value) {
          this.state.passwordInputValid = true;
        } else {
          this.state.passwordInputValid = false;
        }
      }
    } else {
      this.state.passwordInputValid = null;
    }
  };

  componentWillMount() {
  }

  onShowLDAPConnectionSettingsClicked = () => {
    this.setState({
      isLDAPConnectionSettingsOpen: !this.state.isLDAPConnectionSettingsOpen
    });
  };

  onSaveConfigClicked = (e) => {
    e.preventDefault();
    this.validateAll();

    console.log(this);
    let ldapConfig: LdapConfig = {
      host: this.hostInput.value,
      port: parseInt(this.portInput.value, 10),
      ldaps: true,
      bind_dn: this.bindDnInput.value,
      bind_password: this.bindDnPasswordInput.value,
      user_base_dn: this.userbaseDnInput.value,
      user_base_filter: this.userbaseFilterInput.value,
      user_rn_attribute: this.usernameRNInput.value,
      force_bind: true
    };
    let encrypt = true;

    setLdap(ldapConfig, encrypt, (error: Error) => {
      if (error) {
        console.log("ERROR", error);
      } else {
        console.log("success");
      }
    });
  };
  onDBChanged = (e) => {
    if (this.dbSelectInput.selectedIndex === 0) { //LDAP
      this.setState({
        showDBOptions: true
      });
    } else if (this.dbSelectInput.selectedIndex === 1) {
      this.setState({
        showDBOptions: false
      });
    }
  };

  render(): React.ReactElement<HTMLDivElement> {
    console.log(this.state.hostInputValid);

    return (
      <div className="user-authentication">
        <div className="space-20">User DB Connection Settings</div>

        <table className="space-20">
          <tbody>
            <tr className="auth-row">
              <td className="auth-left">User DB Type</td>
              <td className="auth-right">
                <select ref={(ref) => this.dbSelectInput = ref} onChange={this.onDBChanged}>
                  <option>LDAP</option>
                  <option>Steam Local DB</option>
                </select>
              </td>
            </tr>
          </tbody>
        </table>

        {this.state.showDBOptions ?
        <div>
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
                    {this.state.hostInputValid ?
                      <input type="text" className="pt-input" ref={(ref) => this.hostInput = ref}></input> :
                      <input type="text" className="pt-input pt-intent-danger" ref={(ref) => this.hostInput = ref}></input>}
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">PORT &nbsp; <Tooltip className="steam-tooltip-launcher" content="The LDAP server port">
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    {this.state.portInputValid ?
                      <input type="text" className="pt-input" ref={(ref) => this.portInput = ref} defaultValue="689"></input> :
                      <input type="text" className="pt-input pt-intent-danger" ref={(ref) => this.portInput = ref} defaultValue="689"></input> }
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">SSL-ENABLED</td>
                  <td className="auth-right">
                    <label className="pt-control pt-switch .modifier">
                      <input type="checkbox" defaultChecked={true} ref={(ref) => this.sslEnabledInput = ref} />
                      <span className="pt-control-indicator"></span>
                    </label>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">CONNECTION ORDER &nbsp; <Tooltip className="steam-tooltip-launcher" content="The order in which Steam will query this LDAP server (among other enabled servers)">
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    {this.state.connectionOrderInputValid ?
                    <input type="text" className="pt-input" ref={(ref) => this.connectionOrderInput = ref} defaultValue="1"></input> :
                    <input type="text" className="pt-input pt-intent-danger" ref={(ref) => this.connectionOrderInput = ref} defaultValue="1"></input>}
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">BIND DN &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>'Distinguished name' used to bind to LDAP server if extended access is needed.<br /> Leave this blank if anonymous bind is sufficient.</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text" className="pt-input" ref={(ref) => this.bindDnInput = ref}></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">BIND DN PASSWORD &nbsp; <Tooltip className="steam-tooltip-launcher" content="Password for the Bind DN user">
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    {this.state.passwordInputValid ?
                      <input type="password" className="pt-input" ref={(ref) => this.bindDnPasswordInput = ref}></input> :
                      <input type="password" className="pt-input pt-intent-danger" ref={(ref) => this.bindDnPasswordInput = ref}></input>}
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">CONFIRM PASSWORD</td>
                  <td className="auth-right">
                    {this.state.passwordInputValid ?
                    <input type="password" className="pt-input" ref={(ref) => this.confirmPasswordInput = ref}></input> :
                    <input type="password" className="pt-input pt-intent-danger" ref={(ref) => this.confirmPasswordInput = ref}></input>}
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">USER BASE DN &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The location of your LDAP users, specified by the DN of your user subtree.<br/> If necessary, you can specify several DNs separated by semicolons.</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text" className="pt-input" ref={(ref) => this.userbaseDnInput = ref}></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">USER BASE FILTER &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The LDAP search filter used to filter users.<br/> Highly recommended if you have a large amount of user entries under your user base DN.<br/> For example, '(department=IT)'</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    {this.state.userbaseFilterInputValid ?
                    <input type="text" className="pt-input" ref={(ref) => this.userbaseFilterInput = ref}></input> :
                    <input type="text" className="pt-input pt-intent-danger" ref={(ref) => this.userbaseFilterInput = ref}></input> }
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">USER NAME ATTRIBUTE &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The user attribute that contains the username.<br/> Note that this attribute's value should be case insensitive.<br/> Set to 'uid' for most configurations. In Active Directory (AD), this should be set to 'sAMAccountName'.</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    {this.state.usernameRNInputValid ?
                    <input type="text" className="pt-input" ref={(ref) => this.usernameRNInput = ref}></input> :
                    <input type="text" className="pt-input pt-intent-danger" ref={(ref) => this.usernameRNInput = ref}></input> }
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">REAL NAME ATTRIBUTE &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The user attribute that contains a human readable name.<br/> This is typically 'cn' (common name) or 'displayName'.</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">

                    {this.state.realNameAttributeInputValid ?
                    <input type="text" className="pt-input" ref={(ref) => this.realNameAttributeInput = ref} defaultValue="cn"></input> :
                    <input type="text" className="pt-input pt-intent-danger" ref={(ref) => this.realNameAttributeInput = ref} defaultValue="cn"></input>}
                  </td>
                </tr>
              </tbody>
            </table>

          </Collapse>
        </div> : null }

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
