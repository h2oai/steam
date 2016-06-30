/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import { connect } from 'react-redux';
import Navigation from '../Navigation/components/Navigation/Navigation';
import Body from '../Body/Body';

import './styles/app.scss';

type NavigationState = {
  isOpen: boolean
}

interface Props {
  navigation: NavigationState
}

interface DispatchProps {
}

export default class App extends React.Component<Props & DispatchProps, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="app-container">
        <Navigation></Navigation>
        <Body>
          {this.props.children}
        </Body>
      </div>
    );
  }
}