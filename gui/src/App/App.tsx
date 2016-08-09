/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import { withRouter, PlainRoute } from 'react-router';
import Notification from './components/Notification';
import Navigation from '../Navigation/components/Navigation/Navigation';
import Breadcrumb from './components/Breadcrumb';
import Body from '../Body/Body';

import './styles/breadcrumb.scss';
import './styles/app.scss';

interface Props {
  routes: PlainRoute[],
  params: any
}

interface DispatchProps {
}

export class App extends React.Component<Props & DispatchProps, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="app-container">
        <Notification/>
        <Navigation routes={this.props.routes} params={this.props.params}></Navigation>
        <div className="body-container">
          <header>
            <Breadcrumb routes={this.props.routes} params={this.props.params}></Breadcrumb>
          </header>
          <Body>
          {this.props.children}
          </Body>
        </div>
      </div>
    );
  }
}
export default withRouter(App);
