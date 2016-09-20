/**
 * Created by justin on 7/12/16.
 */
import * as React from 'react';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import { fetchPackages } from '../actions/deployment.actions';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/packaging.scss';

interface Props {
  projectId: string,
  deployments: {
    packages: string[]
  }
}

interface DispatchProps {
  fetchPackages: Function
}

export class Packaging extends React.Component<Props & DispatchProps, any> {
  componentWillMount() {
    this.props.fetchPackages(parseInt(this.props.projectId, 10));
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="packaging">
        <h1>PREPROCESSING PACKAGES</h1>
        <div className="intro">
          Custom packaging methods for model deployment
        </div>
        <Table>
          <Row header={true}/>
          {this.props.deployments.packages.map((packageName, i) => {
            return (
              <Row key={i}>
                <Cell className="folder-icon">
                  <i className="fa fa-folder"/>
                </Cell>
                <Cell>{packageName}</Cell>
              </Row>
            );
          })}
        </Table>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    deployments: state.deployments,
    projects: state.projects.project
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchPackages: bindActionCreators(fetchPackages, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Packaging);
