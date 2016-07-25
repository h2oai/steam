/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import { Link, withRouter } from 'react-router';
import { Sidebar } from '../Sidebar/Sidebar';
import './navigation.scss';
import { routes } from '../../../routes';
import * as _ from 'lodash';
const logo = require('../../../../assets/h2o-home.png');

interface Props {
  router: any,
  routes: any
}

interface DispatchProps {
}

interface State {
  isOpen: boolean
}

export class Navigation extends React.Component<Props & DispatchProps, State> {

  constructor() {
    super();
  }

  sitemap = routes[0].childRoutes;

  render(): React.ReactElement<HTMLElement> {

    let submenu = null;

    _.forEach(this.props.routes, (route) => {
      if (!submenu && route.path && route.path !== "/" && this.props.router.isActive(route.path)) {
        submenu = (
          <Sidebar className='secondary-navigation'>
            <nav className="navigation--primary">
              <div className="navigation">
                <header>
                  <div className="header-navigation">
                    <i className="fa fa-angle-left"></i><span>{route.name}</span>
                  </div>
                </header>
                <div className="header-content">UNTITLED</div>
                <ul className="nav-list">
                  {_.map(route.childRoutes, (menuItem: any) => {
                    let path = route.path + '/' + menuItem.path;
                    return (!menuItem.showInNavigation) ? null : (
                      <li key={menuItem.path} className={classNames('nav-list--item', {active: this.props.router.isActive(path)})}>
                        <Link to={path}>{menuItem.name}</Link>
                      </li>
                    );
                  })}
                </ul>
              </div>
            </nav>
          </Sidebar>
        );
      }
    });
    return (
      <div className='nav-container'>
        <Sidebar className="primary-navigation">
          <nav className="navigation--primary">
            <div className="navigation">
              <header>
                <div className="logo-container">
                  <Link to="/"><div className="logo"><img src={logo}></img></div></Link>
                </div>
              </header>
              <div className="header-content">
              </div>
              <ul className='nav-list'>
              {_.map(this.sitemap, (route: any) => {
                  return (!route.showInNavigation) ? null : (
                    <li key={route.path} className={classNames('nav-list--item', {active: this.props.router.isActive(route.path)})}>
                      <Link to={route.path}><i className={route.icon}></i><div className="nav-list--label">{route.name}</div></Link>
                    </li>
                  );
                })
              }
              </ul>
            </div>
          </nav>
        </Sidebar>
        {submenu}
      </div>
    );
  }
}

export default withRouter(Navigation);
