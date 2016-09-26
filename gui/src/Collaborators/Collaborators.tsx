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

import * as React from 'react';
import * as _ from 'lodash';
import PageHeader from '../Projects/components/PageHeader';
import TabNavigation from '../Projects/components/TabNavigation';
import ProjectMembers from './components/projectMembers';
import ProjectLabelsAccess from './components/projectLabelsAccess';

import './styles/collaborators.scss';

interface Props {
  params: {
    projectid: string,
  }
}

interface DispatchProps {
}

export default class Collaborators extends React.Component<Props & DispatchProps, any> {

  constructor() {
    super();
    this.state = {
      tabs: {
        projectMembers: {
          label: 'MEMBERS',
          isSelected: true,
          onClick: this.switchTab.bind(this),
          component: <ProjectMembers />
        },
        labelsAccess: {
          label: 'LABELS ACCESS',
          isSelected: false,
          onClick: this.switchTab.bind(this),
          component: <ProjectLabelsAccess />
        }
      },
      isSelected: 'projectMembers'
    };
  }

  componentWillMount(): void {
    this.setState({
      tabs: {
        projectMembers: {
          label: 'MEMBERS',
          isSelected: true,
          onClick: this.switchTab.bind(this),
          component: <ProjectMembers />
        },
        labelsAccess: {
          label: 'LABELS ACCESS',
          isSelected: false,
          onClick: this.switchTab.bind(this),
          component: <ProjectLabelsAccess />
        }
      }

    });
  }

  switchTab(newTab) {
    let key = _.findKey(this.state.tabs, newTab);
    let newState = _.cloneDeep(this.state);
    Object.keys(newState.tabs).map((tab) => {
      newState.tabs[tab].isSelected = false;
    });
    newState.tabs[key].isSelected = true;
    newState.isSelected = key;
    this.setState(newState);
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="collaborators">
        <PageHeader>
          <span>Collaborators</span>
        </PageHeader>
        <TabNavigation tabs={this.state.tabs}/>

        {this.state.tabs.projectMembers.isSelected === true ?
          <ProjectMembers projectid={this.props.params.projectid} loadLabelsTab={this.switchTab.bind(this, this.state.tabs.labelsAccess)} /> : null}
        {this.state.tabs.labelsAccess.isSelected === true ?
          <ProjectLabelsAccess projectid={this.props.params.projectid} /> : null}
      </div>
    );
  }
}
