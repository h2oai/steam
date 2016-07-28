/**
 * Created by justin on 7/11/16.
 */
import * as React from 'react';
import * as _ from 'lodash';
import PageHeader from '../Projects/components/PageHeader';
import TabNavigation from '../Projects/components/TabNavigation';
import DeployedServices from '../Projects/components/DeployedServices';
import Packaging from './components/Packaging';
import './styles/deployment.scss';

export default class Services extends React.Component<any, any> {
  constructor() {
    super();
    this.state = {
      tabs: {
        deployedServices: {
          label: 'DEPLOYED SERVICES',
          isSelected: true,
          onClick: this.clickHandler.bind(this),
          component: <DeployedServices/>
        },
        packaging: {
          label: 'PACKAGING',
          isSelected: false,
          onClick: this.clickHandler.bind(this),
          component: <Packaging/>
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

  uploadNewPackage() {

  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="services">
        <PageHeader>
          <span>Deployment</span>
          <span><button className="default" onClick={this.uploadNewPackage.bind(this)}>Upload New Package</button></span>
        </PageHeader>
        <TabNavigation tabs={this.state.tabs}/>
        <main>
          {this.state.tabs[this.state.isSelected].component}
        </main>
      </div>
    );
  }
}
