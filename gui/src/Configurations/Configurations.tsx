/**
 * Created by Jeff Fohl <jfohl@h2o.ai> on 7/29/16.
 */
import * as React from 'react';
import * as _ from 'lodash';
import PageHeader from '../Projects/components/PageHeader';
import TabNavigation from '../Projects/components/TabNavigation';
import Labels from './components/Labels';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchLabels } from './actions/configuration.labels.action';
import './styles/configurations.scss';

interface Props {
  labels?: any
}

interface DispatchProps {
  fetchLabels: Function
}

class Configurations extends React.Component<Props & DispatchProps, any> {

  /**
   * TODO(jefffohl): Make the tab container a generalized container, like <TabContainer>, to keep things DRY.
   */

  constructor() {
    super();
    this.state = {
      tabs: {
        labels: {
          label: 'Labels',
          isSelected: true,
          onClick: this.clickHandler.bind(this),
          component: <Labels labels={this.props.labels}/>
        }
      },
      isSelected: 'labels'
    };
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
      <div className="services">
        <PageHeader>
          <span>Project Configurations</span>
        </PageHeader>
        <TabNavigation tabs={this.state.tabs}/>
        <main>
          {this.state.tabs[this.state.isSelected].component}
        </main>
      </div>
    );
  }
}

function mapStateToProps(state: any): any {
  return {
    labels: state.projects[this.props.params.projectId].labels
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchLabels: bindActionCreators(fetchLabels, dispatch)
  };
}

export default connect<Props, any, any>(mapStateToProps, mapDispatchToProps)(Configurations);
