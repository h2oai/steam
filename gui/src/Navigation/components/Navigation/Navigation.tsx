/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import { Link, withRouter } from 'react-router';
import { Sidebar } from '../Sidebar/Sidebar';
import './navigation.scss';
import { routes } from '../../routes';
import * as _ from 'lodash';
const logo = require('../../../../assets/h2o-home.png');

interface Props {
  router: any
}

interface DispatchProps {
}

interface State {
  isOpen: boolean
}

export class Navigation extends React.Component<Props & DispatchProps, State> {

  sitemap = routes[0].childRoutes;

  constructor() {
    super();
    this.openSubmenu = this.openSubmenu.bind(this);
    this.closeSubmenu = this.closeSubmenu.bind(this);
    this.state = {
      isOpen: true
    };
  }

  openSubmenu() {
    if (this.state.isOpen === false) {
      this.setState({
        isOpen: true
      });
    }
  }

  closeSubmenu() {
    if (this.state.isOpen === true) {
      this.setState({
        isOpen: false
      });
    }

  }

  getPath(): string {
    let foo = _.filter(this.sitemap, (route) => {
      console.log('route: ', route);
      return this.props.router.isActive(route.path, true);
    });
    console.log("foo", foo);
    return (foo[0]) ? foo[0].path : '';
  }

  render(): React.ReactElement<HTMLElement> {
    return (
      <div className="nav-container" onMouseLeave={this.closeSubmenu}>
        <Sidebar className="primary-navigation">
          <nav className="navigation--primary">
            <div className="navigation">
              <header>
                <div className="logo-container">
                  <div className="logo"><img src={logo}></img></div>
                </div>
              </header>
              <div className="header-content">
              </div>
              <ul className={classNames('nav-list', {open: this.state.isOpen})}>
              {_.map(this.sitemap, (route: any) => {
                  return (
                    <li key={route.path} className={classNames('nav-list--item', {active: this.props.router.isActive(route.path, true)})} onMouseOver={this.openSubmenu}>
                      <Link to={route.path}><i className={route.icon}></i><div className="nav-list--label">{route.name}</div></Link>
                    </li>
                  );
                })
              }
              </ul>
            </div>
          </nav>
        </Sidebar>
        <Sidebar className={classNames('secondary-navigation', {open: this.state.isOpen})}>
          <nav className="navigation--primary">
            <div className="navigation">
              <header>
                <div className="header-navigation">
                  <i className="fa fa-angle-left"></i><span>Projects</span>
                </div>
              </header>
              <div className="header-content">UNTITLED</div>
              <ul className="nav-list">
                {_.map(this.sitemap[0].childRoutes, (route: any) => {
                  let path = this.sitemap[0].path + '/' + route.path;
                  return (!route.showInNavigation) ? null : (
                    <li key={route.path} className={classNames('nav-list--item', {active: this.props.router.isActive(path, true)})}>
                      <Link to={path}>{route.name}</Link>
                    </li>
                  );
                })}
              </ul>
            </div>
          </nav>
        </Sidebar>
      </div>
    );
  }
}

export default withRouter(Navigation);
