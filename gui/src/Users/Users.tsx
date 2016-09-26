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
import PageHeader from '../Projects/components/PageHeader';
import TabNavigation from '../Projects/components/TabNavigation';
import { Identity, Permission } from '../Proxy/Proxy';
import UserAccess from './components/UserAccess';
import RolePermissions from './components/RolePermissions';
import './styles/users.scss';

interface Props {
  identity: Identity
  permission: Permission
}

export default class Users extends React.Component<Props, any> {

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
        }
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

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="users">
        <PageHeader>USERS</PageHeader>
        <div className="panel-container">
          <TabNavigation tabs={this.state.tabs} />
          {this.state.tabs.users.isSelected === true ?
            <UserAccess /> : null}
          {this.state.tabs.roles.isSelected === true ?
            <RolePermissions /> : null}
        </div>
      </div>
    );
  }
}
