import * as React from 'react';
import * as _ from 'lodash';
import PageHeader from '../Projects/components/PageHeader';
import TabNavigation from '../Projects/components/TabNavigation';
import ProjectMembers from './components/projectMembers';
import ProjectLabelsAccess from './components/projectLabelsAccess';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import './styles/collaborators.scss';

interface Props {
  params: {
    projectid: string
  }
}

interface DispatchProps {
}

export class Collaborators extends React.Component<Props & DispatchProps, any> {
  constructor() {
    super();
    this.state = {
      tabs: {
        projectMembers: {
          label: 'MEMBERS',
          isSelected: true,
          onClick: this.clickHandler.bind(this),
          component: <ProjectMembers/>
        },
        projectLabelsAccess: {
          label: 'LABEL ACCESS',
          isSelected: false,
          onClick: this.clickHandler.bind(this),
          component: <ProjectLabelsAccess/>
        }
      },
      isSelected: 'projectMembers',
      projectId: null
    };
  }

  componentWillMount() {
    this.setState({
      tabs: {
        projectMembers: {
          label: 'MEMBERS',
          isSelected: true,
          onClick: this.clickHandler.bind(this),
          component: <ProjectMembers/>
        },
        projectLabelsAccess: {
          label: 'LABEL ACCESS',
          isSelected: false,
          onClick: this.clickHandler.bind(this),
          component: <ProjectLabelsAccess/>
        }
      }
    });
  }

  clickHandler(tab) {
    let key = _.findKey(this.state.tabs, tab);
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
          <ProjectMembers /> : null}
        {this.state.tabs.projectLabelsAccess.isSelected === true ?
          <ProjectLabelsAccess /> : null}
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
  };
}

function mapDispatchToProps(dispatch) {
  return {
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Collaborators);
