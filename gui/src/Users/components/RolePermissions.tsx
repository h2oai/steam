import * as React from 'react';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/users.scss';
import {fetchPermissionsByRole} from "../actions/users.actions";

interface Props {
  permissionsByRole:Array<any>,
  roles:Array<any>
}

interface DispatchProps {
  fetchPermissionsByRole: Function
}

export class RolePermissions extends React.Component<Props & DispatchProps, any> {

  componentWillMount() {
    this.props.fetchPermissionsByRole();
  }

  render(): React.ReactElement<HTMLDivElement> {
    let permissionRows = this.props.permissionsByRole.map(function(permissionSet) {
      return<Row>
          <Cell className="right-table-bar" key={permissionSet.id}>{permissionSet.description}</Cell>
          {permissionSet.flags.map((flag)=> {
            return <Cell className="center-text"><input type="checkbox" value="on" checked={flag} readOnly={true}></input></Cell>
          })}
        </Row>
    });

    /*let permissionRows = "";
    let flagSet;
    for(let permissionSet of this.props.permissionsByRole) {
      flagSet = "";
      for(let flag of permissionSet.flags) {
        flagSet += <Cell className="center-text"><input type="checkbox" value="on" checked=""></input></Cell>
      }
      permissionRows +=
        <Row>
          <Cell className="right-table-bar">{permissionSet.description}</Cell>
          {flagSet}
        </Row>
    }*/

    return (
      <div className="role-permissions">
        {this.props.roles ? <Table>
          <Row header={true}>
            <Cell className="right-table-bar">Permission Name</Cell>
            {this.props.roles.map((role)=> {
              return <Cell className="center-text">{role.description}</Cell>
            })}
          </Row>

          {permissionRows}

        </Table>
        : null}
      </div>
    );
  }
}

function mapStateToProps(state):any {
  return {
    permissionsByRole: state.users.permissionsByRole,
    roles: state.users.roles
  };
}

function mapDispatchToProps(dispatch):DispatchProps {
  return {
    fetchPermissionsByRole: bindActionCreators(fetchPermissionsByRole, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(RolePermissions);
