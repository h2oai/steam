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
import '../styles/users.scss';
import DefaultModal from '../../App/components/DefaultModal';
import {Identity, Workgroup, Role} from "../../Proxy/Proxy";
import {UserWithWorkgroups, UserWithRolesAndProjects} from "../actions/users.actions";

interface Props {
  userToEdit: Identity
  closeHandler: Function
  open: boolean
  workgroups: Array<Workgroup>
  roles: Array<Role>
  userWithWorkgroups: UserWithWorkgroups
  usersWithRolesAndProjects: Array<UserWithRolesAndProjects>
}
interface DispatchProps {
  fetchWorkgroups: Function,
  fetchUsersWithRolesAndProjects: Function,
  updateUserWorkgroups: Function,
  updateUserRoles: Function
}
interface InputWithWorkgroup {
  input: HTMLInputElement
  workgroup: Workgroup
}
interface InputWithRole {
  input: HTMLInputElement
  role: Role
}

export default class EditUserDialog extends React.Component<Props & DispatchProps, any> {

  inputsWithWorkgroups: Array<InputWithWorkgroup>;
  inputsWithRoles: Array<InputWithRole>;

  constructor(props) {
    super(props);
    this.inputsWithWorkgroups = [];
    this.inputsWithRoles = [];
  }

  componentWillMount() {
    this.props.fetchWorkgroups();
    this.props.fetchUsersWithRolesAndProjects();
  }

  registerWorkgroupInput = (input: HTMLInputElement, workgroup: Workgroup): void => {
    if (!input) return;
    this.inputsWithWorkgroups.push({
      input,
      workgroup
    });
  };

  registerRoleInput = (input: HTMLInputElement, role: Role): void => {
    if (!input) return;
    this.inputsWithRoles.push({
      input,
      role
    });
  };

  onCancelClicked = () => {
    this.props.closeHandler();
  };

  onConfirmClicked = () => {
    let requestedEnableWorkgroupIds: Array<number> = [];
    let requestedEnableRoleIds: Array<number> = [];

    for (let inputWithWorkgroup of this.inputsWithWorkgroups) {
      if (inputWithWorkgroup.input.checked) {
        requestedEnableWorkgroupIds.push(inputWithWorkgroup.workgroup.id);
      }
    }
    for (let inputWithRole of this.inputsWithRoles) {
      if (inputWithRole.input.checked) {
        requestedEnableRoleIds.push(inputWithRole.role.id);
      }
    }

    this.props.updateUserWorkgroups(this.props.userToEdit.id, requestedEnableWorkgroupIds);
    this.props.updateUserRoles(this.props.userToEdit.id, requestedEnableRoleIds);
    this.props.closeHandler();
  };

  arePropsReady = (): boolean => {
    if (this.props.userWithWorkgroups && this.props.userToEdit && this.props.usersWithRolesAndProjects) { //state exists
      if (this.props.userWithWorkgroups.id === this.props.userToEdit.id) { //state is not stale
        return true;
      }
    }
    return false;
  };

  doesUserHaveAccessToWorkgroup = (workgroupToCheck, workgroupsUserHasAccessTo): boolean => {
    let toReturn = false;
    for (let workgroupWithAccess of workgroupsUserHasAccessTo) {
      if (workgroupToCheck.id === workgroupWithAccess.id) {
        toReturn = true;
      }
    }

    return toReturn;
  };

  renderWorkgroupsList = () => {
    return this.props.workgroups.map((workgroupToDisplay: Workgroup, index, array) => {
      if (this.doesUserHaveAccessToWorkgroup(workgroupToDisplay, this.props.userWithWorkgroups.workgroups)) {
        return <div key={index}>
          <input type="checkbox" defaultChecked={true} ref={(input) => this.registerWorkgroupInput(input, workgroupToDisplay)} /> {workgroupToDisplay.name}
        </div>;
      } else {
        return <div key={index}>
          <input type="checkbox" defaultChecked={false} ref={(input) => this.registerWorkgroupInput(input, workgroupToDisplay)} /> {workgroupToDisplay.name}
        </div>;
      }
    });
  };

  renderRolesList = () => {
    function doesUserCurrentlyHaveRole(roleToDisplay: Role, usersCurrentRoles: Array<Role>): boolean {
      for (let usersCurrentRole of usersCurrentRoles) {
        if (roleToDisplay.id === usersCurrentRole.id) {
          return true;
        }
      }
    }

    for (let userWithRoleAndProject of this.props.usersWithRolesAndProjects) {
      if (userWithRoleAndProject.user.id === this.props.userToEdit.id) {
        return this.props.roles.map((roleToDisplay: Role, index, array) => {
          if (doesUserCurrentlyHaveRole(roleToDisplay, userWithRoleAndProject.roles)) {
            return <div key={index}>
              <input type="checkbox" defaultChecked={true} ref={(input) => this.registerRoleInput(input, roleToDisplay)} /> {roleToDisplay.name}
            </div>;
          } else {
            return <div key={index}>
              <input type="checkbox" defaultChecked={false} ref={(input) => this.registerRoleInput(input, roleToDisplay)} /> {roleToDisplay.name}
            </div>;
          }
        });
      }
    }
  };

  render(): React.ReactElement<DefaultModal> {
    this.inputsWithWorkgroups = [];

    return (
      <DefaultModal open={this.props.open}>
        <h1>EDIT USER DETAILS</h1>&nbsp;
        <p> Give {this.props.userToEdit ? this.props.userToEdit.name : null} access to these workgroups:</p>
        { this.arePropsReady() ? this.renderWorkgroupsList() : null }
        <p> Give {this.props.userToEdit ? this.props.userToEdit.name : null} access to these roles:</p>
        { this.arePropsReady() ? this.renderRolesList() : null }
        <br />&nbsp;
        <div className="button-primary" onClick={this.onConfirmClicked}>Confirm</div>
        <div className="button-secondary" onClick={this.onCancelClicked}>Cancel</div>
      </DefaultModal>
    );
  }
}
