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
import PageHeader from '../../Projects/components/PageHeader';
import Cell from "../../Projects/components/Cell";
import Table from "../../Projects/components/Table";
import Row from "../../Projects/components/Row";
import {Role, Workgroup} from "../../Proxy/Proxy";
import { fetchWorkgroups } from '../../Projects/actions/projects.actions';
import { bindActionCreators } from 'redux';
import {createUser, NewUserDetails} from "../actions/users.actions";
import InputFeedback from "../../App/components/InputFeedback";
import {FeedbackType} from "../../App/components/InputFeedback";
import MouseEvent = __React.MouseEvent;

interface Props {
  workgroups: Array<Workgroup>,
  roles: Array<Role>,
  cancelHandler: Function
}

interface DispatchProps {
  fetchWorkgroups: Function,
  createUser: Function
}

interface InputWithRole {
  input: HTMLInputElement;
  role: Role;
}
interface InputWithWorkgroup {
  input: HTMLInputElement,
  workgroup: Workgroup;
}

export class CreateUser extends React.Component<Props & DispatchProps, any> {

  inputsWithRoles: Array<InputWithRole>;
  inputsWithWorkgroups: Array<InputWithWorkgroup>;
  nameInput: HTMLInputElement;
  passwordInput: HTMLInputElement;
  passwordInputConfirm: HTMLInputElement;

  constructor(props) {
    super(props);
    this.inputsWithRoles = [];
    this.inputsWithWorkgroups = [];
    this.state = {
      validNameEntered: false,
      validPasswordEntered: false,
      validPasswordConfirmEntered: false,
      invalidPasswordConfirmEntered: false
    };
  }

  componentWillMount(): void {
    this.props.fetchWorkgroups();
  }

  registerRoleInput = (input: HTMLInputElement, role: Role) : void => {
    if (!input) return;
    this.inputsWithRoles.push({
      input,
      role
    });
  };

  registerWorkgroupInput = (input: HTMLInputElement, workgroup: Workgroup): void => {
    if (!input) return;
    this.inputsWithWorkgroups.push({
      input,
      workgroup
    });
  };

  onNameChanged = (): void => {
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
  };

  onPasswordChanged = (): void => {
    if (this.passwordInput && this.passwordInput.value.length > 1) {
      if (!this.state.validPasswordEntered) {
        this.setState({
          validPasswordEntered: true
        });
      }
    } else {
      if (this.state.validPasswordEntered) {
        this.setState({
          validPasswordEntered: false
        });
      }
    }
  };

  onPasswordConfirmChanged = (): void => {
    if (this.passwordInputConfirm.value.length < 1) {
      this.setState({
        validPasswordConfirmEntered: false,
        invalidPasswordConfirmEntered: false
      });
    } else {
      if (this.passwordInputConfirm.value === this.passwordInput.value) {
        this.setState({
          validPasswordConfirmEntered: true,
          invalidPasswordConfirmEntered: false
        });
      } else {
        this.setState({
          validPasswordConfirmEntered: false,
          invalidPasswordConfirmEntered: true
        });
      }
    }
  };

  onCreateUserClicked = (): void => {
    let workgroupIds = [];
    let roleIds = [];

    for (let inputWithWorkgroup of this.inputsWithWorkgroups) {
      if (inputWithWorkgroup.input.checked) {
        workgroupIds.push(inputWithWorkgroup.workgroup.id);
      }
    }
    for (let inputWithRole of this.inputsWithRoles) {
      if (inputWithRole.input.checked) {
        roleIds.push(inputWithRole.role.id);
      }
    }

    this.props.createUser(new NewUserDetails(
      this.nameInput.value,
      this.passwordInput.value,
      workgroupIds,
      roleIds
    ));
  };

  onCancelClicked = (e: MouseEvent) => {
    this.props.cancelHandler();
  };

  render(): React.ReactElement<HTMLDivElement> {
    this.inputsWithRoles = [];
    this.inputsWithWorkgroups = [];
    return (
      <div className="users">
        <PageHeader>CREATE NEW USER</PageHeader>
        <Table>
          <Row>
            <Cell>
              USERNAME
              This must correspond with a username in your YARN system
            </Cell>
            <Cell>
              <input type="text" ref={(input) => this.nameInput = input} onChange={this.onNameChanged.bind(this)} />
            </Cell>
          </Row>
          <Row>
            <Cell>
              PASSWORD
            </Cell>
            <Cell>
              <input type="password" ref={(input) => this.passwordInput = input} onChange={this.onPasswordChanged.bind(this)} />
            </Cell>
          </Row>
          <Row>
            <Cell>
              CONFIRM PASSWORD
            </Cell>
            <Cell>
              <input type="password" ref={(input) => this.passwordInputConfirm = input} onChange={this.onPasswordConfirmChanged.bind(this)} /><br />&nbsp;
              { this.state.invalidPasswordConfirmEntered ? <InputFeedback message="passwords do not match" type={FeedbackType.Error} /> : null }
              { this.state.validPasswordConfirmEntered ? <InputFeedback message="password confirmed" type={FeedbackType.Confirm} /> : null }
            </Cell>
          </Row>
          <Row>
            <Cell>
              ROLE
              The role(s) this user will have in Steam. At least one role is required.
            </Cell>
            <Cell>
              <div>
                {this.props.roles.map((role: Role, index, array) => {
                    return <div key={index}>
                        <input type="checkbox" ref={(input) => this.registerRoleInput(input, role)} /> {role.description}
                      </div>;
                  })
                }
              </div>
            </Cell>
          </Row>
          <Row>
            <Cell>
              WORKGROUPS
              The workgroup(s) this user will have access to
            </Cell>
            <Cell>
              <div>
                {this.props.workgroups ? this.props.workgroups.map((workgroup: Workgroup, index, array) => {
                  return <div key={index}>
                      <input type="checkbox" ref={(input) => this.registerWorkgroupInput(input, workgroup)} /> {workgroup.name}
                    </div>;
                }) : null
                }
              </div>
            </Cell>
          </Row>
        </Table>
        &nbsp;
        <br />
        &nbsp;
        {this.state.validNameEntered && this.state.validPasswordEntered && this.state.validPasswordConfirmEntered ?
        <div className="button-primary" onClick={this.onCreateUserClicked}>Create User</div>
        :<div className="button-primary disabled">Create User</div>}
        <div className="button-secondary" onClick={this.onCancelClicked}>Cancel</div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    roles: state.users.roles,
    workgroups: state.projects.workgroups
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchWorkgroups: bindActionCreators(fetchWorkgroups, dispatch),
    createUser: bindActionCreators(createUser, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(CreateUser);
