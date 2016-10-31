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
import {Identity, Workgroup} from "../../Proxy/Proxy";
import {UserWithWorkgroups} from "../actions/users.actions";

interface Props {
  userToEdit: Identity
  closeHandler: Function
  open: boolean
  workgroups: Array<Workgroup>
  userWithWorkgroups: UserWithWorkgroups
}
interface DispatchProps {
  fetchWorkgroups: Function,
  updateUserWorkgroups: Function
}
interface InputWithWorkgroup {
  input: HTMLInputElement
  workgroup: Workgroup
}

export default class EditUserDialog extends React.Component<Props & DispatchProps, any> {

  inputsWithWorkgroups: Array<InputWithWorkgroup>;

  constructor(props) {
    super(props);
    this.inputsWithWorkgroups = [];
  }

  registerWorkgroupInput = (input: HTMLInputElement, workgroup: Workgroup): void => {
    if (!input) return;
    this.inputsWithWorkgroups.push({
      input,
      workgroup
    });
  };

  onCancelClicked = () => {
    this.props.closeHandler();
  };

  onConfirmClicked = () => {
    let requestedEnableWorkgroupIds: Array<number> = [];

    for (let inputWithWorkgroup of this.inputsWithWorkgroups) {
      if (inputWithWorkgroup.input.checked) {
        requestedEnableWorkgroupIds.push(inputWithWorkgroup.workgroup.id);
      }
    }

    this.props.updateUserWorkgroups(this.props.userToEdit.id, requestedEnableWorkgroupIds);
    this.props.closeHandler();
  };

  render(): React.ReactElement<DefaultModal> {
    this.inputsWithWorkgroups = [];

    let workgroupsList;
    let userHasAccess: boolean;
    if (this.props.userWithWorkgroups && this.props.userToEdit && this.props.userWithWorkgroups.id === this.props.userToEdit.id) {
      workgroupsList = this.props.workgroups.map((workgroup: Workgroup, index, array) => {
        userHasAccess = false;
        for (let workgroupWithAccess of this.props.userWithWorkgroups.workgroups) {
          if (workgroup.id === workgroupWithAccess.id) {
            userHasAccess = true;
          }
        }

        if (userHasAccess) {
          return <div key={index}>
            <input type="checkbox" defaultChecked={true} ref={(input) => this.registerWorkgroupInput(input, workgroup)} /> {workgroup.name}
          </div>;
        } else {
          return <div key={index}>
            <input type="checkbox" defaultChecked={false} ref={(input) => this.registerWorkgroupInput(input, workgroup)} /> {workgroup.name}
          </div>;
        }
      });
    }

    return (
      <DefaultModal open={this.props.open}>
        <h1>EDIT USER DETAILS</h1>
        <br />&nbsp;
        <p> Give {this.props.userToEdit ? this.props.userToEdit.name : null} access to these workgroups:</p>
        { workgroupsList }
        <br />&nbsp;
        <div className="button-primary" onClick={this.onConfirmClicked}>Confirm</div>
        <div className="button-secondary" onClick={this.onCancelClicked}>Cancel</div>
      </DefaultModal>
    );
  }
}
