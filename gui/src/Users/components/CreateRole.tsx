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
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/users.scss';
import PageHeader from '../../Projects/components/PageHeader';
import Cell from "../../Projects/components/Cell";
import Table from "../../Projects/components/Table";
import Row from "../../Projects/components/Row";
import {PermissionsWithRoles, fetchPermissionsWithRoles, createRole, NewRolePermission} from "../actions/users.actions";
import {Role} from "../../Proxy/Proxy";


interface Props {
  permissionsWithRoles: Array<PermissionsWithRoles>,
  projects: Array<any>,
  roles: Array<any>,
  cancelHandler: Function
}

interface DispatchProps {
  fetchPermissionsWithRoles: Function,
  createRole: Function
}

interface IPermissionInput {
  input: HTMLInputElement;
  permissionId: number;
}
class PermissionInput implements IPermissionInput {
  constructor(public input: HTMLInputElement, public permissionId: number) {}
}

export class CreateRole extends React.Component<Props & DispatchProps, any> {

  nameInput: HTMLInputElement;
  descriptionInput: HTMLInputElement;
  permissionInputs: Array<IPermissionInput>;

  constructor(props) {
    super(props);
    this.state = {
      validNameEntered: false,
      validDescriptionEntered: false
    };
  }

  componentWillMount() {
    this.permissionInputs = [];
    this.props.fetchPermissionsWithRoles();
  }

  registerInput(input: HTMLInputElement, permissionId: number) {
    if (input) {
      this.permissionInputs.push(
        new PermissionInput(input, permissionId)
      );
    }
  }

  onNameChanged() {
    if (this.nameInput && this.nameInput.value.length > 0) {
      if (!this.state.validNameEntered) {
        this.setState({
          validNameEntered: true
        });
      }
    } else {
      if (this.state.validNameEntered) {
        this.setState({
          validNameEntered: false
        });
      }
    }
  }

  onDescriptionChanged() {
    if (this.descriptionInput && this.descriptionInput.value.length > 0) {
      if (!this.state.validDescriptionEntered) {
        this.setState({
          validDescriptionEntered: true
        });
      }
    } else {
      if (this.state.validDescriptionEntered) {
        this.setState({
          validDescriptionEntered: false
        });
      }
    }
  }

  onCreateRoleClicked = () => {
    console.log("creating role");
    let newRolePermissions: Array<NewRolePermission> = [];
    let newRolePermission: NewRolePermission;

    for (let input of this.permissionInputs) {
      newRolePermissions.push(new NewRolePermission(input.permissionId, input.input.checked));
    }

    this.props.createRole(this.nameInput.value, this.descriptionInput.value, newRolePermissions);
  }

  render(): React.ReactElement<HTMLDivElement> {
    this.permissionInputs = [];
    return (
      <div className="users">
        <PageHeader>CREATE NEW ROLE</PageHeader>
        <div className="lede">To create a new type of role in Steam, provide a name for this role, and select the privileges it should have</div>
        <br />
        <div>Role Name</div>
        <input type="text" ref={(input) => this.nameInput = input} onChange={this.onNameChanged.bind(this) }/>
        <br />
        <div>Role Description</div>
        <input type="text" ref={(input) => this.descriptionInput = input} onChange={this.onDescriptionChanged.bind(this) }/>
        <br />
        <Table>
          <Row>
            <Cell>PERMISSION</Cell>
            <Cell>IS GRANTED</Cell>
          </Row>
          { this.props.permissionsWithRoles ? this.props.permissionsWithRoles.map((permission, index, array) => {
            return<Row key={index}>
              <Cell>{permission.description}</Cell>
              <Cell><input ref={(input) => this.registerInput(input, permission.id)} type="checkbox"/></Cell>
            </Row>;
            })
          : null}
        </Table>
        &nbsp;
        <br />
        &nbsp;
        {this.state.validNameEntered && this.state.validDescriptionEntered ?
         <div className="button-primary" onClick={this.onCreateRoleClicked}>Create Role</div>
         : <div className="button-primary disabled">Create Role</div>}
         <div className="button-secondary" onClick={this.props.cancelHandler}>Cancel</div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    permissionsWithRoles: state.users.permissionsWithRoles,
    roles: state.users.roles,
    users: state.users.users
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchPermissionsWithRoles: bindActionCreators(fetchPermissionsWithRoles, dispatch),
    createRole: bindActionCreators(createRole, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(CreateRole);
