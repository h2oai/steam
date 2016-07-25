/**
 * Created by justin on 7/22/16.
 */
import * as React from 'react';
import * as _ from 'lodash';
import ProjectDetails from './ProjectDetails';
import WelcomeSplashScreen from './WelcomeSplashScreen';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchProjects } from '../actions/projects.actions';
import { Project } from '../../Proxy/proxy';

interface DispatchProps {
  fetchProjects: Function
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
      <ProjectDetails projects={this.props.projects}></ProjectDetails>
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
    fetchProjects: bindActionCreators(fetchProjects, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ProjectScreenStrategy);
