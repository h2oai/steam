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
import { NotificationType } from "../../App/components/Notification";
import {ToastDataFactory} from "../../App/actions/notification.actions";
import { toastManager } from '../../App/components/ToastManager';


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
  bindDnPasswordInput: HTMLInputElement;
  confirmPasswordInput: HTMLInputElement;

  constructor(params) {
    super(params);
    this.state = {
      hostValue: "",
      portValue: "636",
      sslEnabledValue: true,
      bindDnValue: "dc=xyz,dc=com",
      bindDnPasswordValue: "",
      confirmPasswordValue: "",
      userbaseDnValue: "",
      userbaseFilterValue: "",
      usernameAttributeValue: "",
      groupDnValue: "dc=xyz,dc=com",
      staticMemberAttributeValue: "memberUid",
      searchRequestSizeLimitValue: "",
      searchRequestTimeLimitValue: "",
      isLDAPConnectionSettingsOpen: true,
      selectedDB: "",
      hostInputValid: true,
      passwordInputValid: true,
      userbaseDnInputValid: true,
      usernameAttributeInputValid: true,
      groupDnInputValid: true,
    };
  }

  validateAll = () => {
    if (this.state.hostValue.length > 3) {
      this.setState({ hostInputValid: true });
    } else {
      this.setState({ hostInputValid: false });
    }

    this.validatePasswords();

    if (this.state.userbaseDnValue.length > 0) {
      this.setState({ userbaseDnInputValid: true });
    } else {
      this.setState({ userbaseDnInputValid: false});
    }

    if (this.state.usernameAttributeValue.length > 0) {
      this.setState({ usernameAttributeInputValid: true });
    } else {
      this.setState({ usernameAttributeInputValid: false });
    }

    if (this.state.groupDnValue.length > 0) {
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
    if (this.props.ldapConfig) {
      this.populateValuesFromConfig(this.props.ldapConfig);
    }
    this.props.getConfig();
    this.props.fetchLdapConfig();
    FocusStyleManager.onlyShowFocusOnTabs();
  }

  componentWillReceiveProps = (nextProps) => {
    if (this.props.ldapConfig && nextProps.ldapConfig && JSON.stringify(this.props.ldapConfig) !== JSON.stringify(nextProps.ldapConfig)) {
      this.populateValuesFromConfig(nextProps.ldapConfig);
    }
  };

  populateValuesFromConfig = (config: LdapConfig) => {
    this.setState({
      hostValue: config.host,
      portValue: config.port,
      sslEnabledValue: config.ldaps,
      bindDnValue: config.bind_dn,
      userbaseDnValue: config.user_base_dn,
      userbaseFilterValue: config.user_base_filter,
      usernameAttributeValue: config.user_name_attribute,
      groupDnValue: config.group_dn,
      staticMemberAttributeValue: config.static_member_attribute,
      searchRequestSizeLimitValue: config.search_request_size_limit,
      searchRequestTimeLimitValue: config.search_request_time_limit
    });
  };

  onShowLDAPConnectionSettingsClicked = () => {
    this.setState({
      isLDAPConnectionSettingsOpen: !this.state.isLDAPConnectionSettingsOpen
    });
  };

  buildLdapConfig = (): LdapConfig => {
    return {
      host: this.state.hostValue,
      port: parseInt(this.state.portValue, 10),
      ldaps: this.state.sslEnabledValue === 'true',
      bind_dn: this.state.bindDnValue,
      bind_password: this.bindDnPasswordInput.value,
      user_base_dn: this.state.userbaseDnValue,
      user_base_filter: this.state.userbaseFilterValue,
      user_name_attribute: this.state.usernameAttributeValue,
      //user_rn_attribute: this.realnameAttributeInput.value,
      group_dn: this.state.groupDnValue,
      static_member_attribute: this.state.staticMemberAttributeValue,
      search_request_size_limit: parseInt(this.state.searchRequestSizeLimitValue, 10),
      search_request_time_limit: parseInt(this.state.searchRequestTimeLimitValue, 10),
      force_bind: true
    };
  };

  onTestConfigClicked = () => {
    this.validateAll();
    if (this.bindDnPasswordInput.value === this.confirmPasswordInput.value) {
      this.props.testLdapConfig(this.buildLdapConfig());
    } else {
      toastManager.show(ToastDataFactory.create(NotificationType.Error, "Passwords do not match"));
    }
  };

  onSaveConfigClicked = (e) => {
    e.preventDefault();
    this.validateAll();
    this.props.saveLdapConfig(this.buildLdapConfig());
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
                      <input type="text"
                             className={"pt-input ldap-input " + (this.state.hostInputValid ? '' : 'pt-intent-danger')}
                             value={this.state.hostValue}
                             onChange={(e: any) => this.setState({"hostValue": e.target.value})}
                      ></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">PORT &nbsp; <Tooltip className="steam-tooltip-launcher" content="The LDAP server port">
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text"
                           className={"pt-input ldap-input "}
                           value={this.state.portValue}
                           onChange={(e: any) => this.setState({"portValue": e.target.value})}
                    ></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">SSL-ENABLED &nbsp; <Tooltip className="steam-tooltip-launcher" content="The LDAP server port">
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <label className="pt-control pt-switch .modifier">
                      <input type="checkbox"
                             checked={this.state.sslEnabledValue}
                             onChange={(e: any) => {
                                  this.setState({"sslEnabledValue": !this.state.sslEnabledValue});
                                }
                             }
                      />
                      <span className="pt-control-indicator"></span>
                    </label>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">BIND DN &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>'Distinguished name' used to bind to LDAP server if extended access is needed.<br /> Leave this blank if anonymous bind is sufficient.</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text"
                           className="pt-input ldap-input"
                           value={this.state.bindDnValue}
                           onChange={(e: any) => this.setState({"bindDnValue": e.target.value})}
                    ></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">BIND DN PASSWORD &nbsp; <Tooltip className="steam-tooltip-launcher" content="Password for the Bind DN user">
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="password" className={"pt-input ldap-input " + (this.state.passwordInputValid ? "" : 'pt-intent-danger')} ref={(ref) => this.bindDnPasswordInput = ref}></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">CONFIRM PASSWORD</td>
                  <td className="auth-right">
                    <input type="password" className={"pt-input ldap-input " + (this.state.passwordInputValid ? "" : 'pt-intent-danger')} ref={(ref) => this.confirmPasswordInput = ref}></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">USER BASE DN &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The location of your LDAP users, specified by the DN of your user subtree.<br/> If necessary, you can specify several DNs separated by semicolons.</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text"
                           className={"pt-input ldap-input " + (this.state.userbaseDnInputValid ? "" : "pt-intent-danger")}
                           value={this.state.userbaseDnValue}
                           onChange={(e: any) => this.setState({"userbaseDnValue": e.target.value})}
                    >
                    </input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">USER BASE FILTER &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The LDAP search filter used to filter users.<br/> Highly recommended if you have a large amount of user entries under your user base DN.<br/> For example, '(department=IT)'</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text"
                           className="pt-input ldap-input"
                           onChange={(e: any) => this.setState({"userbaseFilterValue": e.target.value})}
                           value={this.state.userbaseFilterValue}
                    ></input>
                  </td>
                </tr>
                <tr className="auth-row">
                  <td className="auth-left">USER NAME ATTRIBUTE &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The user attribute that contains the username.<br/> Note that this attribute's value should be case insensitive.<br/> Set to 'uid' for most configurations. In Active Directory (AD), this should be set to 'sAMAccountName'.</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text"
                           className={"pt-input ldap-input " + (this.state.usernameAttributeInputValid ? '' : 'pt-intent-danger')}
                           onChange={(e: any) => this.setState({"usernameAttributeValue": e.target.value})}
                           value={this.state.usernameAttributeValue}
                           ></input>
                  </td>
                </tr>
                <tr className="auth-row">
                    <td className="auth-left">GROUP DN &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The location of your LDAP groups, specified by the DN of your group subtree.<br/> If necessary, you can specify several DNs separated by semicolons.</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                    <td className="auth-right">
                      <input type="text"
                           className={"pt-input ldap-input " + (this.state.groupDnInputValid ? '' : 'pt-intent-danger')}
                           onChange={(e: any) => this.setState({groupDnValue: e.target.value})}
                           value={this.state.groupDnValue}
                           ></input>
                    </td>
                </tr>

                <tr className="auth-row">
                    <td className="auth-left">STATIC MEMBER ATTRIBUTE &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>The group attribute that contains the group name.<br/> A typical value for this is 'cn'.</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text"
                           className="pt-input ldap-input "
                           value={this.state.staticMemberAttributeValue}
                           onChange={(e: any) => this.setState({"staticMemberAttributeValue": e.target.value})}
                    ></input>
                  </td>
                </tr>


                <tr className="auth-row">
                  <td className="auth-left">SEARCH REQUEST SIZE LIMIT &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>H2O Steam can chase referrals with anonymous bind only.<br/> You must also have anonymous search enabled on your LDAP server. Turn this off if you have no need for referrals.</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text"
                           className="pt-input ldap-input "
                           value={this.state.searchRequestSizeLimitValue}
                           onChange={(e: any) => this.setState({"searchRequestSizeLimitValue": e.target.value})}
                    ></input>
                  </td>
                </tr>


                <tr className="auth-row">
                  <td className="auth-left">SEARCH REQUEST TIME LIMIT &nbsp; <Tooltip className="steam-tooltip-launcher" content={<div>Sets the maximum number of entries requested by LDAP searches.<br /> The number actually returned is subject to the limit imposed by the LDAP server.</div>}>
                    <i className="fa fa-question-circle-o" aria-hidden="true"></i>
                  </Tooltip></td>
                  <td className="auth-right">
                    <input type="text"
                           className="pt-input ldap-input "
                           value={this.state.searchRequestTimeLimitValue}
                           onChange={(e: any) => this.setState({ "searchRequestTimeLimitValue": e.target.value })}
                    ></input>
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
