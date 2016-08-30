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
        <p>Cupcake ipsum dolor sit amet chocolate bar sesame snaps sugar plum dessert. Sugar plum sesame snaps oat cake jelly cake sugar plum cake danish pie. Jelly-o candy canes soufflé gummi bears jelly beans sweet roll bear claw.</p>
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
