/**
 * Created by justin on 7/22/16.
 */
import * as moment from 'moment';
import { Link } from 'react-router';
import PageHeader from './PageHeader';
import Panel from './Panel';
import * as React from 'react';
import { Project } from '../../Proxy/Proxy';
import '../styles/projectslist.scss';
import { hashHistory } from 'react-router';
import Progress from '../../App/components/Progress';

interface Props {
  projects: Project[],
  deleteProject: Function
}

export default class ProjectsList extends React.Component<Props, any> {
  openProject(projectId) {
    hashHistory.push('/projects/' + projectId + '/models');
  }

  deleteProject(projectId) {
    this.props.deleteProject(projectId);
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
            {this.props.projects.map((project: any, i) => {
              return (
                <Panel className="project-card" key={i}>
                  <article>
                    <div className="project-metadata" onClick={this.openProject.bind(this, project.id)}>
                      <header>{project.name}</header>
                      <div>{project.model_category}</div>
                      <div>{moment.unix(project.created_at).format('YYYY-MM-DD HH:mm')}</div>
                    </div>
                  </article>
                  { project.isDeleteInProgress ?
                    <Progress message="Deleting project" />
                    : <i className="fa fa-trash" aria-hidden="true" onClick={ () => this.deleteProject(project.id) }></i>
                  }
                </Panel>
              );
            })}
          </div>
        </div>
      </div>
    );
  }
}
