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
import * as moment from 'moment';
import { Link } from 'react-router';
import PageHeader from './PageHeader';
import Panel from './Panel';
import * as React from 'react';
import { Project } from '../../Proxy/Proxy';
import '../styles/projectslist.scss';
import { hashHistory } from 'react-router';
import InputFeedback from '../../App/components/InputFeedback';
import { FeedbackType } from '../../App/components/InputFeedback';

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
            <Link to="/newproject" className="button-primary header-buttons">Create New Project</Link>
        </PageHeader>
        <div>
          <h1>All Projects</h1>
          <div className="panel-container">
            {this.props.projects.map((project: any, i) => {
              return (
                <Panel className="project-card" key={i}>
                  <article>
                    <div className="project-metadata" onClick={this.openProject.bind(this, project.id)}>
                      <header className="link">{project.name}</header>
                      <div>{project.model_category}</div>
                      <div>{moment.unix(project.created_at).format('YYYY-MM-DD HH:mm')}</div>
                    </div>
                  </article>
                  { project.isDeleteInProgress ?
                    <div className="deleting-progress">
                      <InputFeedback type={FeedbackType.Info} message="Deleting project" />
                    </div>
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
