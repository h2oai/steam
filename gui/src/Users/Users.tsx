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
import {Identity, Permission} from '../Proxy/Proxy';
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


interface Props {
  identity: Identity
  permission: Permission
  createNewUserIsEntered: boolean
  createNewRoleIsEntered: boolean
}
interface DispatchProps {
  enterNewUser: Function
  enterNewRole: Function
  exitNewUser: Function
  exitNewRole: Function
  fetchLdapConfig: Function
}
export class Users extends React.Component<Props & DispatchProps, any> {

  constructor() {
    super();
    this.state = {
      tabs: {
        users: {
          label: 'USERS',
          isSelected: true,
          onClick: this.clickHandler.bind(this),
          component: <UserAccess />
        },
        packaging: {
          label: 'ROLES',
          isSelected: false,
          onClick: this.clickHandler.bind(this),
          component: <RolePermissions />
        },
        authentication: {
          label: 'USER AUTHENTICATION',
          isSelected: false,
          onClick: this.clickHandler.bind(this),
          component: <UserAuthentication />
        }
      },
      isSelected: 'users'
    };
  }

  componentWillMount(): void {
    this.setState({
      tabs: {
        users: {
          label: 'USERS',
          isSelected: true,
          onClick: this.clickHandler.bind(this),
          component: <UserAccess />
        },
        roles: {
          label: 'ROLES',
          isSelected: false,
          onClick: this.clickHandler.bind(this),
          component: <RolePermissions />
        },
        authentication: {
          label: 'AUTHENTICATION',
          isSelected: false,
          onClick: this.clickHandler.bind(this),
          component: <UserAuthentication />
        },
      }
    });
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
    if (this.props.createNewUserIsEntered) {
      return (
        <CreateUser cancelHandler={this.onCancelCreateUserClicked.bind(this)} />
      );
    } else if (this.props.createNewRoleIsEntered) {
      return (
        <CreateRole cancelHandler={this.onCancelCreateRoleClicked.bind(this)} />
      );
    } else {
      return (
        <div className="users">
          <PageHeader>USERS & ROLES
            <div className="header-buttons">
              <div className="button-primary" onClick={this.onCreateUserClicked.bind(this)}>Create User</div>
              <div className="button-primary" onClick={this.onCreateRoleClicked.bind(this)}>Create Role</div>
            </div>
          </PageHeader>

          <div className="panel-container">
            <TabNavigation tabs={this.state.tabs}/>
            {this.state.tabs.users.isSelected === true ?
              <UserAccess /> : null}
            {this.state.tabs.roles.isSelected === true ?
              <RolePermissions /> : null}
            {this.state.tabs.authentication.isSelected === true ?
              <UserAuthentication /> : null}
          </div>
        </div>
      );
    }
  }
}

function mapStateToProps(state: any): any {
  return {
    createNewUserIsEntered: state.users.createNewUserIsEntered,
    createNewRoleIsEntered: state.users.createNewRoleIsEntered
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    enterNewUser: bindActionCreators(enterNewUser, dispatch),
    enterNewRole: bindActionCreators(enterNewRole, dispatch),
    exitNewUser: bindActionCreators(exitNewUser, dispatch),
    exitNewRole: bindActionCreators(exitNewRole, dispatch),
    fetchLdapConfig: bindActionCreators(fetchLdapConfig, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Users);
