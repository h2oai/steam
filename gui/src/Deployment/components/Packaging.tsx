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
