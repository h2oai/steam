/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import { withRouter } from 'react-router';
import Navigation from '../Navigation/components/Navigation/Navigation';
import Breadcrumb from './components/Breadcrumb';
import Body from '../Body/Body';

import './styles/breadcrumb.scss';
import './styles/app.scss';

type NavigationState = {
  isOpen: boolean
}

interface Props {
  navigation: NavigationState,
  routes: ReactRouter.PlainRoute & {
    isHiddenBreadcrumb: boolean,
    isExcludedFromBreadcrumb: boolean,
    name: string
  }[]
}

interface DispatchProps {
}

export class App extends React.Component<Props & DispatchProps, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="app-container">
        <Navigation routes={this.props.routes}></Navigation>
        <div className="body-container">
          <header>
            <Breadcrumb routes={this.props.routes}></Breadcrumb>
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

