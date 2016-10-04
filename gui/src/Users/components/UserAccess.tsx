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
import {fetchUsersWithRolesAndProjects, changeFilterSelections} from "../actions/users.actions";
import {Role} from "../../Proxy/Proxy";

interface Props {
  projects: Array<any>,
  roles: Array<any>,
  users: Array<any>,
  usersWithRolesAndProjects: Array<any>,
  selectedRoles: any
}

interface DispatchProps {
  fetchUsersWithRolesAndProjects: Function,
  changeFilterSelections: Function
}

export class UserAccess extends React.Component<Props & DispatchProps, any> {

  componentWillMount() {
    if (!this.props.usersWithRolesAndProjects) {
      this.props.fetchUsersWithRolesAndProjects();
    }
  }

  onRoleCheckboxClicked(e) {
    this.props.changeFilterSelections(parseInt((e.target.dataset as any).id, 10), e.target.checked);
  }

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

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="user-access intro">
        <div className="filter-column">
          FILTERS
          <Table className="full-size">
            <Row header={true}>
              <Cell>
                ROLES
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
          </Table>
        </div>
        <div className="user-access-list">
          <Table>
            <Row header={true}>
              <Cell>User</Cell>
              <Cell>Role</Cell>
            </Row>

            { this.props.usersWithRolesAndProjects ?
                this.props.usersWithRolesAndProjects.map((userWithRoleAndProject,index) => {
                  if(this.shouldRowBeShown(userWithRoleAndProject.roles)) {
                    return <Row key={index}>
                      <Cell>{userWithRoleAndProject.user.name}</Cell>
                      <Cell> {
                        userWithRoleAndProject.roles.map((role, index) => {
                          return <span key={index}>
                            {role.name}
                          </span>;
                        })
                      }</Cell>
                    </Row>;
                  } else {
                    return null;
                  }
                })
              : null
            }

          </Table>
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
    selectedRoles: state.users.selectedRoles
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchUsersWithRolesAndProjects: bindActionCreators(fetchUsersWithRolesAndProjects, dispatch),
    changeFilterSelections: bindActionCreators(changeFilterSelections, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(UserAccess);
