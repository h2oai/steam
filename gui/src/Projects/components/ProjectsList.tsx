/**
 * Created by justin on 7/22/16.
 */
import * as moment from 'moment';
import { Link } from 'react-router';
import PageHeader from './PageHeader';
import Panel from './Panel';
import * as React from 'react';
import { hashHistory } from 'react-router';
import { openConfirmation, closeConfirmation } from '../../App/actions/confirmation.actions';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/projectslist.scss';
import { deleteProject } from '../actions/projects.actions';
import { Project } from '../../Proxy/Proxy';

interface Props {
  projects: Project[]
}

interface DispatchProps {
  openConfirmation: Function,
  closeConfirmation: Function,
  deleteProject: Function
}

export class ProjectsList extends React.Component<Props & DispatchProps, any> {
  constructor() {
    super();
    this.state = {
      isOpenDeleteProjectModal: false
    };
  }

  openProject(projectId) {
    hashHistory.push('/projects/' + projectId + '/models');
  }

  openDeleteProjectModal(projectId, event) {
    event.stopPropagation();
    this.props.openConfirmation('Delete Project?', 'Would you like to delete this project?', () => {
      this.props.deleteProject(projectId);
      this.closeDeleteProjectModal();
    }, this.closeDeleteProjectModal.bind(this));
  }

  closeDeleteProjectModal() {
    this.props.closeConfirmation();
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="project-details">
        <PageHeader>
          <span>PROJECTS</span>
          <span className="new-project-button-container">
            <Link to="/newproject" className="default">Create New Project</Link>
          </span>
        </PageHeader>
        <div>
          <h1>All Projects</h1>
          <div className="panel-container">
            {this.props.projects.map((project, i) => {
              return (
                <Panel key={i} onClick={this.openProject.bind(this, project.id)}>
                  <article>
                    <div className="project-metadata">
                      <header>{project.name}<span className="delete-project"
                                                  onClick={this.openDeleteProjectModal.bind(this, project.id)}><i
                        className="fa fa-trash"/></span></header>
                      <div>{project.model_category}</div>
                      <div>{moment.unix(project.created_at).format('YYYY-MM-DD HH:mm')}</div>
                    </div>
                  </article>
                </Panel>
              );
            })}
          </div>
        </div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    confirmation: state.confirmation
  };
}

function mapDispatchToProps(dispatch) {
  return {
    openConfirmation: bindActionCreators(openConfirmation, dispatch),
    closeConfirmation: bindActionCreators(closeConfirmation, dispatch),
    deleteProject: bindActionCreators(deleteProject, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ProjectsList);
