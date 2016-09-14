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
