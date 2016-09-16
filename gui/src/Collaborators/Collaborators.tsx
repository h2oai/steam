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
