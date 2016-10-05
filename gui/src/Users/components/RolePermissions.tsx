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

  constructor(params) {
    super(params);
    this.state = {
      requestedChanges: null
    };
  }

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

  cancelIndividualChange = (index) => {
    let clone = this.state.requestedChanges.slice(0);
    clone.splice(index, 1);

    this.setState({
      requestedChanges: clone
    });
  };

  requestChanges = () => {
    let updates = [];

    for (var permissionKey in this.permissionInputs) {
      for (var flagKey in this.permissionInputs[permissionKey].flags) {
        let flagset = this.permissionInputs[permissionKey].flags[flagKey];
        if (flagset.originalFlag !== flagset.input.checked) {
          updates.push({
            newFlag: flagset.input.checked,
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
      requestedChanges: updates
    });
    //this.props.saveUpdatedPermissions((this as any).permissionInputs);
  };

  render(): React.ReactElement<HTMLDivElement> {
    console.log(this.state);
    if (this.state.requestedChanges === null) { //list options
      let permissionRows;
      (this as any)._checkboxes = {};

      if (this.props.permissionsWithRoles) {
        permissionRows = this.props.permissionsWithRoles.map(function (permissionSet, permissionIndex) {
          return<Row key={permissionIndex}>
            <Cell className="right-table-bar" key={permissionSet.description}>{permissionSet.description}</Cell>
            {permissionSet.flags.map((flag, flagIndex) => {
              if (flagIndex === 0) {
                return <Cell className="center-text" key={flagIndex}><input
                  ref={(input) => this.registerInput(input, true, flagIndex, permissionSet, permissionIndex)}
                  type="checkbox" value="on" defaultChecked={true} readOnly={false} disabled={false}></input></Cell>;
              } else {
                return <Cell className="center-text" key={flagIndex}><input
                  ref={(input) => this.registerInput(input, flag, flagIndex, permissionSet, permissionIndex)}
                  type="checkbox" value="on" defaultChecked={flag} readOnly={false} disabled={false}></input></Cell>;
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
          <div className="button-primary" onClick={this.requestChanges}>Save</div>
        </div>
      );
    } else { //confirm screen
      return (
        <div>
          <h1>CONFIRMING PERMISSION CHANGES</h1>
          <p>You are making the following changes</p>
          <Table>
            <Row header={true}>
              <Cell>CHANGE</Cell>
              <Cell>CONFIRM</Cell>
            </Row>
            { this.state.requestedChanges.map((requestedChange, index, array) => {
              return <Row key={index}>
                <Cell>{requestedChange.userDescription} &nbsp; {requestedChange.newFlag ? <span>gains</span> : <span>loses</span>} &nbsp; {requestedChange.description}</Cell>
                <Cell>{requestedChange.confirmed ? <span>Confirmed</span> : <span className="button-primary">Confirm</span>} <span onClick={() => this.cancelIndividualChange(index)} className="button">Cancel Change</span></Cell>
              </Row>;
            })}
          </Table>
        </div>
      );
    }
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
