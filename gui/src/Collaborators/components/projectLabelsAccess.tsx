import * as React from 'react';
import * as _ from 'lodash';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/collaborators.scss';
import { fetchLabelsForProject } from '../actions/collaborators.actions';
import { setCurrentProject } from '../../Projects/actions/projects.actions';

interface Props {
  projectid: string,
  labels: Array<any>
}

interface DispatchProps {
  fetchLabelsForProject: Function,
  setCurrentProject: Function
}

export class ProjectLabelsAccess extends React.Component<Props & DispatchProps, any> {

  componentWillMount(): void {
    this.props.setCurrentProject(parseInt(this.props.projectid, 10));
    this.props.fetchLabelsForProject();
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="labelsAccess">
        <p></p>
        <h1>Labels Access</h1>
        <p>Cupcakey ipsum dolor sit amet chocolate bar sesame snaps sugar plum dessert. Sugar plum sesame snaps oat cake jelly cake sugar plum cake danish pie. Jelly-o candy canes soufflé gummi bears jelly beans sweet roll bear claw.</p>
        <Table>
          <Row header={true}>
            <Cell>LABEL</Cell>
            <Cell>USERS</Cell>
          </Row>
            {this.props.labels ?
              this.props.labels.map((label, labelIndex) => {
                return <Row key={labelIndex}>
                  <Cell>{label.name}</Cell>
                  <Cell>
                  {label.identities ? label.identities.map((identity, identityIndex) => {
                    return <div key={identityIndex}>
                      <span className="access-name">{identity.identity_name}</span>&nbsp;
                      <span className="access-type">{identity.kind}</span>
                    </div>;
                  }) : null}
                  </Cell>
                </Row>;
              }) : null}
        </Table>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    labels: state.collaborators.labels
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchLabelsForProject: bindActionCreators(fetchLabelsForProject, dispatch),
    setCurrentProject: bindActionCreators(setCurrentProject, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ProjectLabelsAccess);
