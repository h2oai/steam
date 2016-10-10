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
 * Created by justin on 7/22/16.
 */
import * as React from 'react';
import * as _ from 'lodash';
import ProjectsList from './ProjectsList';
import WelcomeSplashScreen from './WelcomeSplashScreen';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchProjects, deleteProject } from '../actions/projects.actions';
import { Project } from '../../Proxy/Proxy';

interface DispatchProps {
  fetchProjects: Function,
  deleteProject: Function
}

interface Props {
  projects: Project[]
}

export class ProjectScreenStrategy extends React.Component<Props & DispatchProps, any> {
  componentWillMount(): void {
    this.props.fetchProjects();
  }

  render(): any {
    if (this.props.projects === null) {
      return null;
    }
    if (_.isEmpty(this.props.projects)) {
      return <WelcomeSplashScreen/>;
    }
    return (
      <ProjectsList projects={this.props.projects} deleteProject={this.props.deleteProject}></ProjectsList>
    );
  }
}

function mapStateToProps(state) {
  return {
    projects: state.projects.availableProjects
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchProjects: bindActionCreators(fetchProjects, dispatch),
    deleteProject: bindActionCreators(deleteProject, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ProjectScreenStrategy);
