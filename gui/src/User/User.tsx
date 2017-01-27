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
import TabNavigation from '../Projects/components/TabNavigation';
import './styles/user.scss';
import PageHeader from '../Projects/components/PageHeader';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import ClusterAuthentication from "./components/ClusterAuthentication";


interface Props {
}
interface DispatchProps {
}
export class User extends React.Component<Props & DispatchProps, any> {

  constructor() {
    super();
    this.state = {
      tabs: this.getTabs(this.props),
      isSelected: 'cluster_authentication'
    };
  }

  componentWillMount(): void {
    this.setState({
      tabs: this.getTabs(this.props)
    });
  }

  getTabs(props): Object {
    if (!props) return null;

    let tabs = {};
    tabs["user"] = {
      label: 'CLUSTER AUTHENTICATION',
      isSelected: true,
      onClick: this.clickHandler.bind(this),
      component: <ClusterAuthentication />
    };
    return tabs;
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

  render(): React.ReactElement<HTMLElement> {
    return (
      <div className="users">
        <PageHeader>USER INFO
        </PageHeader>

        <div className="panel-container">
          <TabNavigation tabs={this.state.tabs}/>
          {this.state.tabs.user && this.state.tabs.user.isSelected ?
            <ClusterAuthentication /> : null}
        </div>
      </div>
    );
  }
}

function mapStateToProps(state: any): any {
  return {
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(User);
