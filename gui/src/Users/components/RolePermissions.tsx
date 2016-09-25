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
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/users.scss';
import {fetchPermissionsWithRoles, PermissionsWithRoles} from "../actions/users.actions";
import {Role} from "../../Proxy/Proxy";

interface Props {
  permissionsWithRoles: Array<PermissionsWithRoles>,
  roles: Array<Role>
}

interface DispatchProps {
  fetchPermissionsWithRoles: Function
}

export class RolePermissions extends React.Component<Props & DispatchProps, any> {

  componentWillMount() {
    this.props.fetchPermissionsWithRoles();
  }

  render(): React.ReactElement<HTMLDivElement> {
    let permissionRows;
    if (this.props.permissionsWithRoles) {
      permissionRows = this.props.permissionsWithRoles.map(function (permissionSet, index) {
        return<Row key={index}>
          <Cell className="right-table-bar" key={permissionSet.description}>{permissionSet.description}</Cell>
          {permissionSet.flags.map((flag,flagIndex) => {
            if (flagIndex === 0) {
              return <Cell className="center-text" key={flagIndex}><input type="checkbox" value="on" checked={true} readOnly={true} disabled={true}></input></Cell>;
            } else {
              return <Cell className="center-text" key={flagIndex}><input type="checkbox" value="on" checked={flag} readOnly={true} disabled={true}></input></Cell>;
            }
          })}
        </Row>;
      });
    }

    return (
      <div className="role-permissions">
        {this.props.permissionsWithRoles && this.props.roles ? <Table>
          <Row header={true}>
            <Cell className="right-table-bar">Permission Name</Cell>
            {this.props.roles.map((role, rolesIndex) => {
              return <Cell className="center-text" key={rolesIndex}>{role.description}</Cell>;
            })}
          </Row>

          {permissionRows}

        </Table>
        : null}
      </div>
    );
  }
}

function mapStateToProps(state): any {
  return {
    permissionsWithRoles: state.users.permissionsWithRoles,
    roles: state.users.roles
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    fetchPermissionsWithRoles: bindActionCreators(fetchPermissionsWithRoles, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(RolePermissions);
