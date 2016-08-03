/**
 * Created by justin on 7/12/16.
 */
import * as React from 'react';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import '../styles/packaging.scss';
import { fetchPackages } from '../actions/deployment.actions';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

interface Props {
  projectId: string,
  packages: string[]
}

interface DispatchProps {
  fetchPackages: Function
}

export default class Packaging extends React.Component<Props, any> {
  componentWillMount() {
    console.log(this.props);
    this.props.fetchPackages();
  }

  render(): React.ReactElement<HTMLDivElement> {
    if (_.isEmpty(this.props.packages)) {
      return <div></div>;
    }
    return (
      <div className="packaging">
        <h1>PREPROCESSING PACKAGES</h1>
        <div>
          Custom packaging methods for model deployment
        </div>
        <Table>
          <Row header={true}/>
          {this.props.packages.map((packageName) => {
            return (
              <Row>
                <Cell className="folder-icon">
                  <i className="fa fa-folder"/>
                </Cell>
                <Cell>{packageName}</Cell>
                <Cell>Mark Landry</Cell>
                <Cell>6 files</Cell>
                <Cell>280KB</Cell>
                <Cell>3 months ago</Cell>
              </Row>
            );
          })}
        </Table>
      </div>
    );
  }
}

function mapStateToProps(state) {
  console.log(state);
  return {
    packages: state.packages,
    projects: state.projects.project
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchPackages: bindActionCreators(fetchPackages, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Packaging);
