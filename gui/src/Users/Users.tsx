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
import TabNavigation from '../Projects/components/TabNavigation';
import {Identity, Permission, Config} from '../Proxy/Proxy';
import UserAccess from './components/UserAccess';
import RolePermissions from './components/RolePermissions';
import './styles/users.scss';
import PageHeader from '../Projects/components/PageHeader';
import CreateUser from "./components/CreateUser";
import CreateRole from "./components/CreateRole";
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import {enterNewUser, enterNewRole, exitNewRole, exitNewUser, fetchLdapConfig} from "./actions/users.actions";
import UserAuthentication from "./components/UserAuthentication";
import {hasPermissionToShow} from "../App/utils/permissions";
import {fetchConfig} from "../Clusters/actions/clusters.actions";
import {fetchIsAdmin} from "../App/actions/global.actions";
import ClusterAuthentication from "./components/ClusterAuthentication";


interface Props {
  identity: Identity
  permission: Permission
  createNewUserIsEntered: boolean
  createNewRoleIsEntered: boolean
  isAdmin: boolean
  config: Config
}
interface DispatchProps {
  enterNewUser: Function
  enterNewRole: Function
  exitNewUser: Function
  exitNewRole: Function
  fetchLdapConfig: Function
  fetchConfig: Function
  fetchIsAdmin: Function
}
export class Users extends React.Component<Props & DispatchProps, any> {

  constructor() {
    super();
    this.state = {
      tabs: this.getTabs(this.props),
      isSelected: 'users'
    };
  }

  componentWillMount(): void {
    this.setState({
      tabs: this.getTabs(this.props)
    });
    this.props.fetchIsAdmin();
    this.props.fetchConfig();
  }

  componentWillReceiveProps(nextProps): void {
    if ((!this.props.config && nextProps.config) || (!this.props.isAdmin && nextProps.isAdmin)) {
      this.setState({
        tabs: this.getTabs(nextProps)
      });
    }
  }

  getTabs(props): Object {
    if (!props) return null;

    let tabs = {};
    tabs["users"] = {
      label: 'USERS',
      isSelected: true,
      onClick: this.clickHandler.bind(this),
      component: <UserAccess />
    };
    if (hasPermissionToShow("ViewRole", props.config, props.isAdmin)) {
      tabs["roles"] = {
        label: 'ROLES',
        isSelected: false,
        onClick: this.clickHandler.bind(this),
        component: <RolePermissions />
      };
    }
    if (props.isAdmin) {
      tabs["authentication"] = {
        label: 'USER AUTHENTICATION',
        isSelected: false,
        onClick: this.clickHandler.bind(this),
        component: <UserAuthentication />
      };
      tabs["cluster_authentication"] = {
        label: 'CLUSTER AUTHENTICATION',
        isSelected: false,
        onClick: this.clickHandler.bind(this),
        component: <ClusterAuthentication />
      };
    }
    return tabs;
  }

  clickHandler(tab) {
    let key = _.findKey(this.state.tabs, tab);
    let newState = _.cloneDeep(this.state);
    Object.keys(newState.tabs).map((tab) => {
      newState.tabs[tab].isSelected = false;
    });
    newState.tabs[key].isSelected = true;
    newState.isSelected = key;
    this.setState(newState);
  }

  onCreateUserClicked() {
    let newState = _.cloneDeep(this.state);
    newState.tabs.roles.isSelected = false;
    newState.tabs.users.isSelected = true;
    newState.tabs.authentication.isSelected = false;
    newState.isSelected = "users";
    this.setState(newState);
    this.props.enterNewUser();
  }

  onCreateRoleClicked() {
    let newState = _.cloneDeep(this.state);
    newState.tabs.roles.isSelected = true;
    newState.tabs.users.isSelected = false;
    newState.tabs.authentication.isSelected = false;
    newState.isSelected = "roles";
    this.setState(newState);
    this.props.enterNewRole();
  }

  onCancelCreateUserClicked() {
    this.props.exitNewUser();
  }

  onCancelCreateRoleClicked() {
    this.props.exitNewRole();
  }

  render(): React.ReactElement<HTMLElement> {
    if (this.props.createNewUserIsEntered && hasPermissionToShow("ManageIdentity", this.props.config, this.props.isAdmin)) {
      return (
        <CreateUser cancelHandler={this.onCancelCreateUserClicked.bind(this)} />
      );
    } else if (this.props.createNewRoleIsEntered && hasPermissionToShow("ManageRole", this.props.config, this.props.isAdmin)) {
      return (
        <CreateRole cancelHandler={this.onCancelCreateRoleClicked.bind(this)} />
      );
    } else {
      return (
        <div className="users">
          <PageHeader>Configurations (Steam Global)
            <div className="header-buttons">
              <div className="button-primary" onClick={this.onCreateUserClicked.bind(this)}>Create User</div>
              <div className="button-primary" onClick={this.onCreateRoleClicked.bind(this)}>Create Role</div>
            </div>
          </PageHeader>
          <p>All settings on this page affect the Steam installation globally for administrators. All users on Steam can potentially be affected. Please be careful and consult support if you are unsure about the implications of updating these settings.</p>

          <div className="panel-container">
            <TabNavigation tabs={this.state.tabs}/>
            {this.state.tabs.users && this.state.tabs.users.isSelected ?
              <UserAccess /> : null}
            {this.state.tabs.roles && this.state.tabs.roles.isSelected ?
              <RolePermissions /> : null}
            {this.state.tabs.authentication && this.state.tabs.authentication.isSelected ?
              <UserAuthentication onCreateRoleClicked={this.onCreateRoleClicked.bind(this)} onManageRoleClicked={() => this.clickHandler(this.state.tabs.roles)} /> : null}
            {this.state.tabs.cluster_authentication && this.state.tabs.cluster_authentication.isSelected ?
              <ClusterAuthentication /> : null}
          </div>
        </div>
      );
    }
  }
}

function mapStateToProps(state: any): any {
  return {
    createNewUserIsEntered: state.users.createNewUserIsEntered,
    createNewRoleIsEntered: state.users.createNewRoleIsEntered,
    isAdmin: state.global.isAdmin,
    config: state.clusters.config,
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    enterNewUser: bindActionCreators(enterNewUser, dispatch),
    enterNewRole: bindActionCreators(enterNewRole, dispatch),
    exitNewUser: bindActionCreators(exitNewUser, dispatch),
    exitNewRole: bindActionCreators(exitNewRole, dispatch),
    fetchLdapConfig: bindActionCreators(fetchLdapConfig, dispatch),
    fetchConfig: bindActionCreators(fetchConfig, dispatch),
    fetchIsAdmin: bindActionCreators(fetchIsAdmin, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Users);
