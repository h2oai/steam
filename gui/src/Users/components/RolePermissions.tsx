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
import { fetchPermissionsWithRoles, PermissionsWithRoles, saveUpdatedPermissions, resetUpdates } from "../actions/users.actions";
import {Role} from "../../Proxy/Proxy";
import RolePermissionsConfirm from "./RolePermissionsConfirm";

interface Props {
  permissionsWithRoles: Array<PermissionsWithRoles>,
  roles: Array<Role>,
  updates: Array<any>
}

interface DispatchProps {
  fetchPermissionsWithRoles: Function,
  saveUpdatedPermissions: Function,
  resetUpdates: Function
}

export class RolePermissions extends React.Component<Props & DispatchProps, any> {

  constructor(params) {
    super(params);
    this.state = {
      requestedChanges: [],
      confirmOpen: false
    };
  }

  permissionInputs = {};

  componentWillMount() {
    this.props.fetchPermissionsWithRoles();
  }

  componentWillReceiveProps(nextProps) {
    if (this.props.permissionsWithRoles !== nextProps.permissionsWithRoles) {
      this.permissionInputs = {};
    }
  }

  componentDidUpdate(prevProps, prevState) {
    if (prevProps.permissionsWithRoles !== this.props.permissionsWithRoles) {
      this.requestChanges();
    }
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

  modalCloseHandler = () => {
    this.props.fetchPermissionsWithRoles();
    this.props.resetUpdates();
  };

  requestChanges = () => {
    let updates = [];

    for (var permissionKey in this.permissionInputs) {
      for (var flagKey in this.permissionInputs[permissionKey].flags) {
        let flagset = this.permissionInputs[permissionKey].flags[flagKey];
        if (flagset.originalFlag.value !== flagset.input.checked) {
          updates.push({
            newFlag: { value: flagset.input.checked, roleId: parseInt(flagset.input.dataset["roleid"], 10) },
            userIndex: flagKey,
            userDescription: this.props.roles[flagKey].description,
            permissionIndex: permissionKey,
            description: this.permissionInputs[permissionKey].permissionSet.description,
            permissionId: this.permissionInputs[permissionKey].permissionSet.id,
            confirmed: false
          });
        }
      }
    }

    updates.sort((a, b) => {
      if (a.newFlag) {
        return -1;
      } else {
        return 1;
      }
    });

    this.setState({
      requestedChanges: updates,
      confirmOpen: false
    });
  };

  requestConfirm = () => {
    this.setState({
      requestedChanges: this.state.requestedChanges,
      confirmOpen: true
    });
  };

  render(): React.ReactElement<HTMLDivElement> {
    let permissionRows;
    (this as any)._checkboxes = {};

    if (this.props.permissionsWithRoles) {
      permissionRows = this.props.permissionsWithRoles.map(function (permissionSet, permissionIndex) {
        return<Row key={permissionIndex}>
          <Cell className="right-table-bar" key={permissionSet.description}>{permissionSet.description}</Cell>
          {permissionSet.flags.map((flag: any, flagIndex) => {
            if (flagIndex === 0) {
              return <Cell className="center-text" key={flagIndex}><input data-roleid={flag.roleId}
                ref={(input) => this.registerInput(input, {value: true, roleId: flag.roleId}, flagIndex, permissionSet, permissionIndex)}
                type="checkbox" value="on" defaultChecked={true} readOnly={true} disabled={true}></input></Cell>;
            } else {
              return <Cell className="center-text" key={flagIndex}><input data-roleid={flag.roleId}
                ref={(input) => this.registerInput(input, flag, flagIndex, permissionSet, permissionIndex)}
                type="checkbox" value="on" defaultChecked={flag.value} readOnly={false} disabled={false} onClick={this.requestChanges}></input></Cell>;
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
        <br />
        { this.state.requestedChanges.length === 0 ?
          <div className="button-primary disabled">Review Changes</div> : null }
        { this.state.requestedChanges.length === 1 ?
          <div className="button-primary" onClick={this.requestConfirm}>Review 1 Change</div> : null }
        { this.state.requestedChanges.length > 1 ?
          <div className="button-primary" onClick={this.requestConfirm}>Review {this.state.requestedChanges.length} Changes</div> : null }
      <RolePermissionsConfirm open={this.state.confirmOpen} closeHandler={this.modalCloseHandler.bind(this)} requestedChanges={this.state.requestedChanges} saveUpdatedPermissions={this.props.saveUpdatedPermissions} updates={this.props.updates} />
    </div>
    );
  }
}

function mapStateToProps(state): any {
  return {
    permissionsWithRoles: state.users.permissionsWithRoles,
    roles: state.users.roles,
    updates: state.users.updates
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    fetchPermissionsWithRoles: bindActionCreators(fetchPermissionsWithRoles, dispatch),
    saveUpdatedPermissions: bindActionCreators(saveUpdatedPermissions, dispatch),
    resetUpdates: bindActionCreators(resetUpdates, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(RolePermissions);
