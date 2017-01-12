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

/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import { withRouter, PlainRoute } from 'react-router';
import NotificationsManager from './components/NotificationsManager';
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
    let isChrome = !!window.chrome && !!window.chrome.webstore;
    let isFirefox = typeof window.InstallTrigger !== 'undefined';
    let app = (
      <div className="app-container">
        <NotificationsManager />
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
    let unsupportedBrowser = (
      <div className="unsupported-browser">
        <h1>Browser not supported</h1>
        <h3><a href="http://docs.h2o.ai/steam/latest-stable/steam-docs/Installation.html">Please review the list of supported browsers.</a></h3>
      </div>
    );
    if (isChrome || isFirefox) {
      return app;
    } else {
      return unsupportedBrowser;
    }
  }
}
export default withRouter(App);
