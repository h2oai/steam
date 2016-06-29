/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router';
import * as classNames from 'classnames';
import { Navigation } from '../Navigation/components/Navigation/Navigation';
import { Sidebar } from '../Navigation/components/Sidebar/Sidebar';
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

  render(): React.ReactElement<HTMLDivElement> {
    let sidebar1 = (
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
    );

    let sidebar2 = (
      <div className={classNames('project-sidebar')}>
        <div className="project-sidebar-header">
          <h5>Project</h5>
          <h3>Airlines</h3>
          <div className="avatar-container">
            <div className="avatar"></div>
            <div className="avatar"></div>
            <div className="avatar"></div>
          </div>
        </div>
        <div className="models item">
          <span className="item--label">Models</span>
          <span className="badge">64</span>
        </div>
        <div className="item">
          <span className="item--label">Dataframe</span>
        </div>
        <div className="item transformations">
          <span className="item--label">Transformations</span>
        </div>
      </div>
    );

    let sidebar3 = (
      <Navigation>
        <header>
          <div className="logo"></div>
            Header goes here
        </header>
        <ul className="nav-list">

        </ul>
      </Navigation>
    );
    return (
      <div className="app-container">
        <div>
          <Sidebar>
            <Navigation>
              <header>
                <div className="logo-container">
                  <div className="logo"></div>
                </div>
              </header>
              <ul className="nav-list">
                <li className="nav-list--item">
                  <i className="fa fa-cloud"></i><Link to="clusters">Clusters</Link>
                </li>
                <li className="nav-list--item">
                  <i className="fa fa-cube"></i><Link to="clusters">Models</Link>
                </li>
                <li className="nav-list--item">
                  <i className="fa fa-cubes"></i><Link to="">Services</Link>
                </li>
                <li className="nav-list--item">
                  <i className="fa fa-clone"></i><Link to="clusters">Assets</Link>
                </li>
                <li className="nav-list--item">
                  <i className="fa fa-folder-open"></i><Link to="projects">Projects</Link>
                </li>
              </ul>
            </Navigation>
          </Sidebar>
          <Sidebar className="sub-navigation">
            <Navigation>
              <header>
              </header>
              <ul className="nav-list">
                <li className="nav-list--item">
                  <Link to="clusters">Sub 1</Link>
                </li>
                <li className="nav-list--item">
                  <Link to="clusters">Sub 2</Link>
                </li>
                <li className="nav-list--item">
                  <Link to="">Sub 3</Link>
                </li>
                <li className="nav-list--item">
                  <Link to="clusters">Sub 4</Link>
                </li>
                <li className="nav-list--item">
                  <Link to="projects">Sub 5</Link>
                </li>
              </ul>
            </Navigation>
          </Sidebar>
        </div>
        <Body>
          {this.props.children}
        </Body>
      </div>
    );
  }
}

function mapStateToProps(state: Props): Props {
  return {
    navigation: state.navigation
  };
}

function mapDispatchToProps(dispatch) {
  return {}
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(App)