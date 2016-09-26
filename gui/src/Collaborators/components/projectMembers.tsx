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
import '../styles/collaborators.scss';
import { fetchMembersForProject } from '../actions/collaborators.actions';
import { setCurrentProject } from '../../Projects/actions/projects.actions';

interface Props {
  projectid: string,
  loadLabelsTab: Function,
  members: Array<any>
}

interface DispatchProps {
  fetchMembersForProject: Function,
  setCurrentProject: Function
}

export class ProjectMembers extends React.Component<Props & DispatchProps, any> {
  componentWillMount(): void {
    this.props.setCurrentProject(parseInt(this.props.projectid, 10));
    this.props.fetchMembersForProject();
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="projectMembers">
        <p></p>
        <h1>Members</h1>
        <p>Theses are users who have access to this project, meaning they can see data, models and services associated with the project. Additionally, owners and collaborators can create new models, and new services based on those models.</p>
        <p>Labels associated with projects have <span className="link" onClick={ this.props.loadLabelsTab }>their own access controls, shown here</span>.</p>
        <Table>
          <Row header={true}>
            <Cell>USER</Cell>
            <Cell>ROLE</Cell>
            <Cell>ACCESS</Cell>
          </Row>
          {this.props.members ?
          this.props.members.map((member, index) => {
            return <Row key={index}>
              <Cell>{member.identity_name}</Cell>
              <Cell>{member.role_name}</Cell>
              <Cell>{member.kind}</Cell>
            </Row>;
          })
            : null }
        </Table>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    members: state.collaborators.members
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchMembersForProject: bindActionCreators(fetchMembersForProject, dispatch),
    setCurrentProject: bindActionCreators(setCurrentProject, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ProjectMembers);
