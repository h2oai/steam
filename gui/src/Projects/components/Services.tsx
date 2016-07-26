/**
 * Created by justin on 7/11/16.
 */
import * as React from 'react';
import * as _ from 'lodash';
import PageHeader from './PageHeader';
import TabNavigation from './TabNavigation';
import DeployedServices from './DeployedServices';
import Packaging from './Packaging';
import ModelApi from './ModelApi';
import '../styles/services.scss';

export default class Services extends React.Component<any, any> {
  constructor() {
    super();
    this.state = {
      tabs: {
        deployedServices: {
          label: 'Services',
          isSelected: true,
          onClick: this.clickHandler.bind(this),
          component: <DeployedServices/>
        },
        packaging: {
          label: 'Packaging',
          isSelected: false,
          onClick: this.clickHandler.bind(this),
          component: <Packaging/>
        },
        modelApi: {
          label: 'Model API',
          isSelected: false,
          onClick: this.clickHandler.bind(this),
          component: <ModelApi/>
        }
      },
      isSelected: 'deployedServices'
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
        <PageHeader>Services</PageHeader>
        <main>
          {this.state.tabs[this.state.isSelected].component}
        </main>
      </div>
    );
  }
}
