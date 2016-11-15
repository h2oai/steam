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
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/users.scss';
import {
  fetchUsersWithRolesAndProjects, changeFilterSelections,
  UserWithRolesAndProjects, deleteUser, undeleteUser, UserWithWorkgroups, fetchWorkgroupsForUserId, updateUserWorkgroups, updateUserRoles
} from "../actions/users.actions";
import {Role, Identity, Project, Workgroup} from "../../Proxy/Proxy";
import {fetchWorkgroups} from "../../Projects/actions/projects.actions";
import EditUserDialog from "./EditUserDialog";
import InputFeedback from "../../App/components/InputFeedback";
import {FeedbackType} from "../../App/components/InputFeedback";

interface Props {
  projects: Array<Project>,
  roles: Array<Role>,
  users: Array<Identity>,
  usersWithRolesAndProjects: Array<UserWithRolesAndProjects>,
  selectedRoles: any,
  workgroups: Array<Workgroup>,
  userWithWorkgroups: UserWithWorkgroups
}
interface DispatchProps {
  fetchUsersWithRolesAndProjects: Function,
  changeFilterSelections: Function,
  deleteUser: Function,
  undeleteUser: Function,
  fetchWorkgroups: Function,
  fetchWorkgroupsForUserId: Function
  updateUserWorkgroups: Function
  updateUserRoles: Function
}

export class UserAccess extends React.Component<Props & DispatchProps, any> {

  constructor(params) {
    super(params);
    this.state = {
      editUserOpen: false,
      userToEdit: null
    };
  }

  componentWillMount() {
    if (!this.props.usersWithRolesAndProjects) {
      this.props.fetchUsersWithRolesAndProjects();
    }
  }

  onRoleCheckboxClicked = (e) => {
    this.props.changeFilterSelections(parseInt((e.target.dataset as any).id, 10), e.target.checked);
  };

  checkIsRoleSelected(id): boolean {
    if (_.isEmpty(this.props.selectedRoles)) {
      return false;
    }

    let index = _.findIndex(this.props.selectedRoles, (o) => {
      if ((o as any).id === (id)) return true;
      return false;
    });
    if (index === -1) console.log("ERROR: unable to find match");

    return this.props.selectedRoles[index].selected;
  }

  shouldRowBeShown(roles: Array<Role>): boolean {
    if (roles.length === 0) {
      for (let selectedRole of this.props.selectedRoles) {
        if (selectedRole.id === -1) {
          return selectedRole.selected;
        }
      }
    }

    for (let role of roles) {
      let index = _.findIndex(this.props.selectedRoles, (o) => {
        if ((o as any).id === role.id) {
          return true;
        } else {
          return false;
        }
      });

      if (index === -1) return false;

      if (this.props.selectedRoles[index].selected) {
        return true;
      }
    }
    return false;
  }

  onEditUserClicked = (user: Identity) => {
    this.props.fetchWorkgroups();
    this.props.fetchWorkgroupsForUserId(user.id);
    this.setState({
      editUserOpen: true,
      userToEdit: user
    });
  };

  editUserCloseHandler = () => {
    this.setState({
      editUserOpen: false
    });
  };

  renderTableRows = () => {
    let isSuperuser = (userWithRoleAndProject): boolean => {
      if (userWithRoleAndProject.user.id !== 1) {
        return true;
      }
      return false;
    };


    let renderLastCell = (userWithRoleAndProject) => {
      if (isSuperuser(userWithRoleAndProject)) {
        if (userWithRoleAndProject.user.is_active) {
          return(<Cell className="link"><span onClick={() => this.props.deleteUser(userWithRoleAndProject.user.id)}><i className="fa fa-times" aria-hidden="true" alt="Deactivate user"></i>&nbsp;Deactivate User</span><br/><br/><span onClick={() => this.onEditUserClicked(userWithRoleAndProject.user)}><i className="fa fa-edit" aria-hidden="true" alt="edit"></i>&nbsp;Edit</span></Cell>);
        } else {
          return(<Cell className="link"><span onClick={() => this.props.undeleteUser(userWithRoleAndProject.user.id)}><i className="fa fa-undo" aria-hidden="true" alt="Reactivate user"></i>&nbsp;Reactivate User</span></Cell>);
        }
      } else {
        return(<Cell></Cell>);
      }
    };

    return(this.props.usersWithRolesAndProjects.map((userWithRoleAndProject, index) => {
      if (this.shouldRowBeShown(userWithRoleAndProject.roles)) {
        return <Row key={index} className={userWithRoleAndProject.user.is_active ? "" : "inactive"}>
          <Cell>{userWithRoleAndProject.user.name}</Cell>
          <Cell> {
            userWithRoleAndProject.roles.length === 0 ? <span><InputFeedback type={FeedbackType.Error} message="No roles assigned." /><br /><span className="link" onClick={() => this.onEditUserClicked(userWithRoleAndProject.user)}>Assign one or more roles</span></span> :
            userWithRoleAndProject.roles.map((role, index) => {
              return  <span key={index}>
                              {role.name}<br/>
                        </span>;
            })
          }</Cell>
          { renderLastCell(userWithRoleAndProject) }
        </Row>;
      }
    }));
  };

  onSelectNoneClicked = () => {
    for (let selectedRole of this.props.selectedRoles) {
      this.props.changeFilterSelections(selectedRole.id, false);
    }
  };
  onSelectAllClicked = () => {
    for (let selectedRole of this.props.selectedRoles) {
      this.props.changeFilterSelections(selectedRole.id, true);
    }
  };

  render(): React.ReactElement<HTMLDivElement> {
    let numRolesSelected = 0;
    if (this.props.selectedRoles) {
      for (let selectedRole of this.props.selectedRoles) {
        if (selectedRole.selected) {
          numRolesSelected++;
        }
      }
    }

    return (
      <div className="user-access intro">
        <EditUserDialog open={this.state.editUserOpen} userToEdit={this.state.userToEdit} closeHandler={this.editUserCloseHandler} fetchWorkgroups={this.props.fetchWorkgroups} fetchUsersWithRolesAndProjects={this.props.fetchUsersWithRolesAndProjects} userWithWorkgroups={this.props.userWithWorkgroups } workgroups={this.props.workgroups} updateUserWorkgroups={this.props.updateUserWorkgroups} updateUserRoles={this.props.updateUserRoles} usersWithRolesAndProjects={this.props.usersWithRolesAndProjects} roles={this.props.roles} />
        <div className="filter-and-list">
          <div className="filter-column">
            FILTERS
            <Table className="full-size">
              <Row header={true}>
                <Cell>
                  ROLES<br/>
                  <div className="bulk-select">
                    {numRolesSelected === 0 ?
                    <div>
                      <input type="radio" name="roleBulkSelect" onChange={this.onSelectNoneClicked} checked={true}></input>
                      Select None<br/>
                    </div>
                    : <div>
                        <input type="radio" name="roleBulkSelect" onChange={this.onSelectNoneClicked} checked={false}></input>
                        Select None<br/>
                    </div> }

                    {this.props.selectedRoles && numRolesSelected > 0 && numRolesSelected < this.props.selectedRoles.length ?
                      <div>
                        <input type="radio" name="roleBulkSelect" checked={true}></input>Select Some<br/>
                      </div>
                      : <div>
                        <input type="radio" name="roleBulkSelect" checked={false}></input>Select Some<br/>
                      </div>
                    }

                    {this.props.selectedRoles && numRolesSelected === this.props.selectedRoles.length ?
                      <div>
                        <input type="radio" name="roleBulkSelect" onChange={this.onSelectAllClicked} checked={true}></input>Select All
                      </div>
                      :<div>
                        <input type="radio" name="roleBulkSelect" onChange={this.onSelectAllClicked} checked={false}></input>Select All
                      </div>
                    }

                  </div>
                </Cell>
              </Row>
              {this.props.roles ?
                <div ref="roleBoxes">{
                  this.props.roles.map((role, index) => {
                    return <Row key={role.name}>
                      <Cell><input type="checkbox" name="selectedRoles" data-id={role.id} checked={this.checkIsRoleSelected(role.id)} onChange={this.onRoleCheckboxClicked.bind(this)}></input> {role.name}</Cell>
                    </Row>;
                  })
                }</div>
                : null
              }
              <Row>
                <Cell className="light-grey">
                  <input type="checkbox" name="selectedRoles" data-id={-1} checked={this.checkIsRoleSelected(-1)} onChange={this.onRoleCheckboxClicked}></input> [No roles]
                </Cell>
              </Row>
            </Table>
          </div>
          <div className="user-access-list">
            <Table>
              <Row header={true}>
                <Cell>User</Cell>
                <Cell>Role(s)</Cell>
                <Cell>Actions</Cell>
              </Row>

              { this.props.usersWithRolesAndProjects ? this.renderTableRows() : null }

            </Table>
          </div>
        </div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    projects: state.users.projects,
    roles: state.users.roles,
    users: state.users.users,
    usersWithRolesAndProjects: state.users.usersWithRolesAndProjects,
    selectedRoles: state.users.selectedRoles,
    workgroups: state.projects.workgroups,
    userWithWorkgroups: state.users.userWithWorkgroups
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchUsersWithRolesAndProjects: bindActionCreators(fetchUsersWithRolesAndProjects, dispatch),
    changeFilterSelections: bindActionCreators(changeFilterSelections, dispatch),
    deleteUser: bindActionCreators(deleteUser, dispatch),
    undeleteUser: bindActionCreators(undeleteUser, dispatch),
    fetchWorkgroups: bindActionCreators(fetchWorkgroups, dispatch),
    fetchWorkgroupsForUserId: bindActionCreators(fetchWorkgroupsForUserId, dispatch),
    updateUserWorkgroups: bindActionCreators(updateUserWorkgroups, dispatch),
    updateUserRoles: bindActionCreators(updateUserRoles, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(UserAccess);
