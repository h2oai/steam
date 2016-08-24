import * as React from 'react';
import PageHeader from '../Projects/components/PageHeader';
import TabNavigation from '../Projects/components/TabNavigation';
import ProjectMembers from './components/projectMembers';
import ProjectLabelsAccess from './components/projectLabelsAccess';
import './styles/collaborators.scss';

interface Props {
}

interface DispatchProps {
}

export class Collaborators extends React.Component<Props & DispatchProps, any> {

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="collaborators">
        <PageHeader>
          <span>Collaborators</span>
        </PageHeader>
        <TabNavigation tabs={this.state.tabs}/>

        {this.state.tabs.projectMembers.isSelected === true ?
          <ProjectMembers /> : null}
        {this.state.tabs.projectLabelsAccess.isSelected === true ?
          <ProjectLabelsAccess /> : null}
      </div>
    );
  }
}
