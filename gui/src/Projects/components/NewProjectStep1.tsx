/**
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import { Link } from 'react-router';
import PageHeader from '../components/PageHeader';
import Table from '../components/Table';
import Row from '../components/Row';
import Cell from '../components/Cell';
import Dropdown from '../components/Dropdown';
import '../styles/newproject.scss';

interface Props {
  children?: React.ReactChildren
}

export default class NewProjectStep1 extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="new-project">
        <PageHeader>New Project</PageHeader>
        <div className="project-description">
          <span>Steam organizes your data sets, your models, and your deployment configurations into one cohesive project. This enables you to:</span>
          <ul className="project-description-list">
            <li>Visually compare all models within a project</li>
            <li>Manage how models from a project gets deployed</li>
            <li>Track the history of model deployment</li>
          </ul>
        </div>
        <div className="cards-container">
          <div className="card small">
            <header>Import Existing Models</header>
            <article>
              Choose this option if you already have H2O models built and stored in a H2O cluster in your network.
            </article>
            <footer>
              <Link to="newproject/import" className="default">Start Import</Link>
            </footer>
          </div>
          <div className="card small">
            <header>Start from Scratch</header>
            <article>
              Choose this option if this is a completely new project.
            </article>
            <footer>
              <Link to="newproject/import" className="default">Create New Project</Link>
            </footer>
          </div>
        </div>
      </div>
    );
  }
}
