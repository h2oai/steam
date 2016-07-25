/**
 * Created by justin on 7/22/16.
 */
import * as React from 'react';
import PageHeader from './PageHeader';
import Panel from './Panel';
import RocGraph from '../../Models/components/RocGraph';
import { Project } from '../../Proxy/proxy';
import { glmTrain } from '../../Models/tests/data/glmTrain';
import '../styles/projectdetails.scss';

interface Props {
  projects: Project[]
}

export default class ProjectDetails extends React.Component<Props, any> {
  render() {
    let config = {
      margin: { top: 2, right: 2, bottom: 2, left: 2 },
      width: 231,
      height: 231,
      interpolationMode: 'basis',
      smooth: true,
      fpr: 'fpr',
      tprVariables: [{
        name: 'tpr',
        label: 'tpr'
      }],
      animate: false,
      hideAxes: true,
      hideAUCText: true,
      hideBoundaries: true
    };
    return (
      <div className="project-details">
        <PageHeader>PROJECTS</PageHeader>
        <div>
          <h1>Your Recent Projects</h1>
          <div className="panel-container">
            {this.props.projects.slice(0, 2).map((project, i) => {
              return (
                <Panel key={'recent' + i}>
                  <article>
                    <div>
                      <RocGraph config={config} data={glmTrain}/>
                    </div>
                    <div className="project-metadata">
                      <header>{project.name}</header>
                    </div>
                  </article>
                </Panel>
              );
            })}
          </div>
          <div>
            <h1>All Projects</h1>
            <div className="panel-container">
              {this.props.projects.map((project, i) => {
                return (
                  <Panel key={'all' + i}>
                    <article>
                      <div>
                        <RocGraph config={config} data={glmTrain}/>
                      </div>
                      <div className="project-metadata">
                        <header>{project.name}</header>
                      </div>
                    </article>
                  </Panel>
                );
              })}
            </div>
          </div>
        </div>
      </div>
    );
  }
}
