import * as React from 'react';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/collaborators.scss';
import { fetchLabels } from '../actions/collaborators.actions';

interface Props {
  labels: Array<any>
}

interface DispatchProps {
  fetchLabels: Function
}

export class ProjectLabelsAccess extends React.Component<Props & DispatchProps, any> {

  componentWillMount(): void {
    this.props.fetchLabels();
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="labelsAccess">
        <p></p>
        <h1>Labels Access</h1>
        <p>Cupcake ipsum dolor sit amet chocolate bar sesame snaps sugar plum dessert. Sugar plum sesame snaps oat cake jelly cake sugar plum cake danish pie. Jelly-o candy canes soufflé gummi bears jelly beans sweet roll bear claw.</p>
        <Table>
          <Row header={true}>
            <Cell>LABEL</Cell>
            <Cell>USERS WITH ACCESS</Cell>
            <Cell>ACCESS</Cell>
          </Row>
          <Row>
            <Cell>PROD</Cell>
            <Cell>First Last</Cell>
            <Cell>OWNER</Cell>
          </Row>
          <Row>
            <Cell>LBL</Cell>
            <Cell>First Last</Cell>
            <Cell>COLLABORATOR</Cell>
          </Row>
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
    fetchLabels: bindActionCreators(fetchLabels, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ProjectLabelsAccess);
