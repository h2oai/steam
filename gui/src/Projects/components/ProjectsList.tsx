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

interface Props {
  projects: Project[]
}

export default class ProjectsList extends React.Component<Props, any> {
  openProject(projectId) {
    hashHistory.push('/projects/' + projectId + '/models');
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
                      <header>{project.name}</header>
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
