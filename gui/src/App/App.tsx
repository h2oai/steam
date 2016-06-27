/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router';
import * as classNames from 'classnames';
import { Navigation } from '../Navigation/components/Navigation/Navigation';
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

export class App extends React.Component<Props & DispatchProps, any> {
  constructor() {
    super();
    this.state = {
      navigation: {
        isOpen: false
      }
    };
  }

  render() {
    return (
      <div className="app-container">
        <div className="stage">
          <div className={classNames('navigation-container', {open: this.props.navigation.isOpen})}>
            <Navigation>
              <li>
                <Link to="clusters">Clusters</Link>
              </li>
              <li>
                Models
              </li>
              <li>
                Services
              </li>
              <li>
                Assets
              </li>
              <li>
                <Link to="projects">Projects</Link>
              </li>
            </Navigation>
          </div>
          <div className={classNames('pusher', {open: this.props.navigation.isOpen})}>
            <Body>
              {this.props.children}
            </Body>
          </div>
        </div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    navigation: state.navigation
  };
}

function mapDispatchToProps(dispatch) {
  return {}
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(App)