import * as React from 'react';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/users.scss';
import {fetchUsersWithRolesAndProjects} from "../actions/users.actions";

interface Props {
  projects:Array<any>,
  roles:Array<any>
}

interface DispatchProps {
  fetchUsersWithRolesAndProjects:Function
}

export class UserAccess extends React.Component<Props & DispatchProps, any> {
  componentWillMount() {
    this.props.fetchUsersWithRolesAndProjects();
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="user-access">
        <div className="filter-column">
          FILTERS
          <Table className="full-size">
            <Row header={true}>
              <Cell>
                ROLES
              </Cell>
            </Row>
            {this.props.roles ?
              <div>{
                this.props.roles.map((role) => {
                  return <Row key={role.name}>
                    <Cell><input type="checkbox" value="on" checked=""></input> {role.name}</Cell>
                  </Row>
                })
              }</div>
              : null
            }
          </Table>
          <Table>
            <Row header={true}>
              <Cell>PROJECTS</Cell>
            </Row>
            {this.props.projects ? <div>
              {this.props.projects.map((project)=> {
                return <Row key={project.name}><Cell><input type="checkbox" value="on" checked=""></input> {project.name}</Cell></Row>
              })}
              </div>
            : null}
          </Table>
        </div>
        <div className="user-access-list">
          <Table>
            <Row header={true}>
              <Cell>User</Cell>
              <Cell>Role</Cell>
              <Cell>Project</Cell>
            </Row>

            <Row>
              <Cell>My User</Cell>
              <Cell>Admin</Cell>
              <Cell>Churn</Cell>
            </Row>

            <Row>
              <Cell>My User</Cell>
              <Cell>Project Lead</Cell>
              <Cell>Churn</Cell>
            </Row>

          </Table>
        </div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    projects:state.users.projects,
    roles:state.users.roles
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchUsersWithRolesAndProjects:bindActionCreators(fetchUsersWithRolesAndProjects,dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(UserAccess);
