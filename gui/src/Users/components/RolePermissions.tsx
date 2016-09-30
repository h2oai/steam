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
import { fetchPermissionsWithRoles, PermissionsWithRoles, saveUpdatedPermissions } from "../actions/users.actions";
import {Role} from "../../Proxy/Proxy";

interface Props {
  permissionsWithRoles: Array<PermissionsWithRoles>,
  roles: Array<Role>
}

interface DispatchProps {
  fetchPermissionsWithRoles: Function,
  saveUpdatedPermissions: Function
}

export class RolePermissions extends React.Component<Props & DispatchProps, any> {
  permissionInputs = {};

  componentWillMount() {
    this.props.fetchPermissionsWithRoles();
  }

  registerInput(input, flag, flagIndex, permissionSet, permissionIndex) {
    if (!(this as any).permissionInputs.hasOwnProperty(permissionIndex)) {
      (this as any).permissionInputs[permissionIndex] = {
        permissionSet,
        flags: {}
      };
    }
    (this as any).permissionInputs[permissionIndex].flags[flagIndex] = {
      originalFlag: flag,
      input
    };
  }

  updatePermissions = () => {
    this.props.saveUpdatedPermissions((this as any).permissionInputs);
  };

  render(): React.ReactElement<HTMLDivElement> {
    let permissionRows;
    (this as any)._checkboxes = {};
    console.log(this);

    if (this.props.permissionsWithRoles) {
      permissionRows = this.props.permissionsWithRoles.map(function (permissionSet, permissionIndex) {
        return<Row key={permissionIndex}>
          <Cell className="right-table-bar" key={permissionSet.description}>{permissionSet.description}</Cell>
          {permissionSet.flags.map((flag,flagIndex) => {
            if (flagIndex === 0) {
              return <Cell className="center-text" key={flagIndex}><input ref={(input) => this.registerInput(input, true, flagIndex, permissionSet, permissionIndex)}  type="checkbox" value="on" defaultChecked={true} readOnly={false} disabled={false}></input></Cell>;
            } else {
              return <Cell className="center-text" key={flagIndex}><input ref={(input) => this.registerInput(input, flag, flagIndex, permissionSet, permissionIndex)}  type="checkbox" value="on" defaultChecked={flag} readOnly={false} disabled={false}></input></Cell>;
            }
          })}
        </Row>;
      }, this);
    }

    return (
      <div className="role-permissions intro">
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
        <div className="button-primary" onClick={this.updatePermissions}>Save</div>
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
    fetchPermissionsWithRoles: bindActionCreators(fetchPermissionsWithRoles, dispatch),
    saveUpdatedPermissions: bindActionCreators(saveUpdatedPermissions, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(RolePermissions);
