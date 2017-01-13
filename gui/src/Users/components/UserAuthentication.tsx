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
import { bindActionCreators } from 'redux';
import { Collapse } from '@blueprintjs/core/dist/components/collapse/collapse';
import { Button } from '@blueprintjs/core/dist/components/button/buttons';
import { Tooltip } from '@blueprintjs/core/dist/components/tooltip/tooltip';
import { LdapConfig} from "../../Proxy/Proxy";
import { FocusStyleManager } from "@blueprintjs/core";
import { fetchLdapConfig, saveLdapConfig, testLdapConfig, setLocalConfig } from "../actions/users.actions";
import { getConfig } from "../../Clusters/actions/clusters.actions";

interface Props {
  doesLdapExist: boolean,
  ldapConfig: LdapConfig,
  authType: string
}
interface DispatchProps {
  fetchLdapConfig: Function,
  saveLdapConfig: Function,
  testLdapConfig: Function,
  setLocalConfig: Function,
  getConfig: Function
}

export class UserAuthentication extends React.Component<Props & DispatchProps, any> {

  dbSelectInput: HTMLSelectElement;
  hostInput: HTMLInputElement;
  portInput: HTMLInputElement;
  sslEnabledInput: HTMLInputElement;
  bindDnInput: HTMLInputElement;
  bindDnPasswordInput: HTMLInputElement;
  confirmPasswordInput: HTMLInputElement;
  userbaseDnInput: HTMLInputElement;
  userbaseFilterInput: HTMLInputElement;
  usernameAttribueInput: HTMLInputElement;
  //realnameAttributeInput: HTMLInputElement;
  groupDnInput: HTMLInputElement;
  staticMemberAttributeInput: HTMLInputElement;
  searchRequestSizeLimitInput: HTMLInputElement;
  searchRequestTimeLimitInput: HTMLInputElement;

  constructor(params) {
    super(params);
    this.state = {
      isLDAPConnectionSettingsOpen: true,
      selectedDB: "",
      hostInputValid: true,
      portInputValid: true,
      passwordInputValid: true,
      userbaseDnInputValid: true,
      usernameAttributeInputValid: true,
      groupDnInputValid: true,
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

    this.validatePasswords();

    if (this.userbaseDnInput.value.length > 0) {
      this.setState({ userbaseDnInputValid: true });
    } else {
      this.setState({ userbaseDnInputValid: false});
    }

    if (this.usernameAttribueInput.value.length > 0) {
      this.setState({ usernameAttributeInputValid: true });
    } else {
      this.setState({ usernameAttributeInputValid: false });
    }

    if (this.groupDnInput.value.length > 0) {
      this.setState({ groupDnInputValid: true });
    } else {
      this.setState({ groupDnInputValid: false});
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
    this.props.getConfig();
    this.props.fetchLdapConfig();
    FocusStyleManager.onlyShowFocusOnTabs();
  }

  onShowLDAPConnectionSettingsClicked = () => {
    this.setState({
      isLDAPConnectionSettingsOpen: !this.state.isLDAPConnectionSettingsOpen
    });
  };

  buildLdapConfig = (): LdapConfig => {
    return {
      host: this.hostInput.value,
      port: parseInt(this.portInput.value, 10),
      ldaps: this.sslEnabledInput.value === 'true',
      bind_dn: this.bindDnInput.value,
      bind_password: this.bindDnPasswordInput.value,
      user_base_dn: this.userbaseDnInput.value,
      user_base_filter: this.userbaseFilterInput.value,
      user_name_attribute: this.usernameAttribueInput.value,
      //user_rn_attribute: this.realnameAttributeInput.value,
      group_dn: this.groupDnInput.value,
      static_member_attribute: this.staticMemberAttributeInput.value,
      search_request_size_limit: parseInt(this.searchRequestSizeLimitInput.value, 10),
      search_request_time_limit: parseInt(this.searchRequestTimeLimitInput.value, 10),
      force_bind: true
    };
  };

  onTestConfigClicked = () => {
    this.props.testLdapConfig(this.buildLdapConfig());
  };

  onSaveConfigClicked = (e) => {
    e.preventDefault();
    this.validateAll();

    if (
      this.state.hostInputValid &&
      this.state.portInputValid &&
      this.state.passwordInputValid &&
      this.state.userbaseDnInputValid &&
      this.state.usernameAttributeInputValid &&
      this.state.groupDnInputValid
    ) {
      this.props.saveLdapConfig(this.buildLdapConfig());
    }
  };
  onDBChanged = (e) => {
    if (e.target.value === "LDAP") {
      this.setState({
        selectedDB: e.target.value
      });
    } else if (e.target.value === "Steam Local DB") {
      this.setState({
        selectedDB: e.target.value
      });
    } else {
      console.log("ERROR", "Unknown Input Value"); //Should never be reached
    }
  };
  onRemoveLdapClicked = () => {
    this.props.setLocalConfig();
  };

  render(): React.ReactElement<HTMLDivElement> {
    let defaultSelectedDB;

    if (this.state.selectedDB && this.state.selectedDB !== "") {
      defaultSelectedDB = this.state.selectedDB;
    } else {
      if (this.props.authType) {
        if (this.props.authType.toLowerCase() === "ldap") {
          defaultSelectedDB = "LDAP";
        } else if (this.props.authType.toLowerCase() === "local") {
          defaultSelectedDB = "Steam Local DB";
        } else {
          console.log("ERROR", "Unrecognized auth type"); //should never be reached
        }
      } else {
        defaultSelectedDB = "Steam Local DB";
      }
    }

    return (
      <div className="user-authentication">

        <p>&nbsp;</p>
        <table className="space-20">
          <tbody>
            <tr className="auth-row">
              <td className="auth-left">User DB Type</td>
              <td className="auth-right">
                <select value={defaultSelectedDB} ref={(ref) => this.dbSelectInput = ref} onChange={this.onDBChanged}>
                  <option>Steam Local DB</option>
                  <option>LDAP</option>
                </select>
              </td>
            </tr>
          </tbody>
        </table>

        {defaultSelectedDB.toLowerCase() === 'ldap' ?
        <div>
          <div className="colapse-header">
            <Button onClick={this.onShowLDAPConnectionSettingsClicked}>
              {this.state.isLDAPConnectionSettingsOpen ?
                <i className="fa fa-minus" aria-hidden="true"></i> :
                <i className="fa fa-plus" aria-hidden="true"></i>
              }
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
                      <input type="text" className={"pt-input " + (this.state.hostInputValid ? '' : 'pt-intent-danger')} ref={(ref) => this.hostInput = ref} defaultValue={this.props.ldapConfig ? this.props.ldapConfig.host : ""}></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">PORT &nbsp; <Tooltip className="steam-tooltip-launcher" content="The LDAP server port">
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text" className={"pt-input " + (this.state.portInputValid ? '' : 'pt-intent-danger')} ref={(ref) => this.portInput = ref} defaultValue={this.props.ldapConfig ? this.props.ldapConfig.port.toString() : ""}></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">SSL-ENABLED</td>
                  <td className="auth-right">
                    <label className="pt-control pt-switch .modifier">
                      <input type="checkbox" defaultChecked={this.props.ldapConfig ? this.props.ldapConfig.ldaps : false} ref={(ref) => this.sslEnabledInput = ref} />
                      <span className="pt-control-indicator"></span>
                    </label>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">BIND DN &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>'Distinguished name' used to bind to LDAP server if extended access is needed.<br /> Leave this blank if anonymous bind is sufficient.</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text" defaultValue={this.props.ldapConfig ? this.props.ldapConfig.bind_dn : null} className="pt-input" ref={(ref) => this.bindDnInput = ref}></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">BIND DN PASSWORD &nbsp; <Tooltip className="steam-tooltip-launcher" content="Password for the Bind DN user">
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                      <input type="password" className={"pt-input " + (this.state.passwordInputValid ? "" : 'pt-intent-danger')} ref={(ref) => this.bindDnPasswordInput = ref}></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">CONFIRM PASSWORD</td>
                  <td className="auth-right">
                    <input type="password" className={"pt-input " + (this.state.passwordInputValid ? "" : 'pt-intent-danger')} ref={(ref) => this.confirmPasswordInput = ref}></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">USER BASE DN &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The location of your LDAP users, specified by the DN of your user subtree.<br/> If necessary, you can specify several DNs separated by semicolons.</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text" defaultValue={this.props.ldapConfig ? this.props.ldapConfig.user_base_dn : ""} className={"pt-input " + (this.state.userbaseDnInputValid ? "" : "pt-intent-danger")} ref={(ref) => this.userbaseDnInput = ref}></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">USER BASE FILTER &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The LDAP search filter used to filter users.<br/> Highly recommended if you have a large amount of user entries under your user base DN.<br/> For example, '(department=IT)'</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text" defaultValue={this.props.ldapConfig ? this.props.ldapConfig.user_base_filter : ""} className="pt-input" ref={(ref) => this.userbaseFilterInput = ref}></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">USER NAME ATTRIBUTE &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The user attribute that contains the username.<br/> Note that this attribute's value should be case insensitive.<br/> Set to 'uid' for most configurations. In Active Directory (AD), this should be set to 'sAMAccountName'.</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text" className={"pt-input " + (this.state.usernameAttributeInputValid ? '' : 'pt-intent-danger')} ref={(ref) => this.usernameAttribueInput = ref} defaultValue={this.props.ldapConfig ? this.props.ldapConfig.user_name_attribute : ""}></input>
                  </td>
                </tr>
                <tr className="auth-row">
                    <td className="auth-left">GROUP DN</td>
                    <td className="auth-right">
                      <input type="text" className={"pt-input " + (this.state.groupDnInputValid ? '' : 'pt-intent-danger')} ref={(ref) => this.groupDnInput = ref} defaultValue={this.props.ldapConfig ? this.props.ldapConfig.group_dn : ""}></input>
                    </td>
                </tr>

                <tr className="auth-row">
                  <td className="auth-left">STATIC MEMBER ATTRIBUTE</td>
                  <td className="auth-right">
                    <input type="text" className="pt-input" ref={(ref) => this.staticMemberAttributeInput = ref} defaultValue={this.props.ldapConfig ? this.props.ldapConfig.static_member_attribute : ""}></input>
                  </td>
                </tr>


                <tr className="auth-row">
                  <td className="auth-left">SEARCH REQUEST SIZE LIMIT</td>
                  <td className="auth-right">
                    <input type="text" className="pt-input" ref={(ref) => this.searchRequestSizeLimitInput = ref} defaultValue={this.props.ldapConfig ? this.props.ldapConfig.search_request_size_limit.toString() : ""}></input>
                  </td>
                </tr>


                <tr className="auth-row">
                  <td className="auth-left">SEARCH REQUEST TIME LIMIT</td>
                  <td className="auth-right">
                    <input type="text" className="pt-input" ref={(ref) => this.searchRequestTimeLimitInput = ref} defaultValue={this.props.ldapConfig ? this.props.ldapConfig.search_request_time_limit.toString() : ""}></input>
                  </td>
                </tr>

              </tbody>
            </table>

          </Collapse>
        </div> : null }

        <div id="actionButtonsContainer" className="space-20">
          {defaultSelectedDB.toLowerCase() === 'ldap' ?
            <span>
              <div className="button-secondary" onClick={this.onTestConfigClicked}>Test Config</div> &nbsp;
              <div className="button-primary" onClick={this.onSaveConfigClicked}>Save Config</div>
            </span>
            : null}
          {this.props.authType && this.props.authType.toLowerCase() === "ldap" && defaultSelectedDB.toLowerCase() !== 'ldap' ?
            <div className="button-primary" onClick={this.onRemoveLdapClicked}>Remove LDAP</div>
            : null}
        </div>

      </div>
    );
  }
}

function mapStateToProps(state): any {
  return {
    ldapConfig: state.users.ldapConfig,
    doesLdapExist: state.users.ldapExists,
    authType: state.clusters.config.authentication_type
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    fetchLdapConfig: bindActionCreators(fetchLdapConfig, dispatch),
    saveLdapConfig: bindActionCreators(saveLdapConfig, dispatch),
    testLdapConfig: bindActionCreators(testLdapConfig, dispatch),
    setLocalConfig: bindActionCreators(setLocalConfig, dispatch),
    getConfig: bindActionCreators(getConfig, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(UserAuthentication);
