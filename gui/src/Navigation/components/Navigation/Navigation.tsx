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
import * as $ from 'jquery';
import * as classNames from 'classnames';
import { Link, withRouter } from 'react-router';
import { Sidebar } from '../Sidebar/Sidebar';
import { buildPath } from '../../../App/utils/buildPath';
import { getRoute } from '../../../App/utils/getRoute';
import { routes } from '../../../routes';
import * as _ from 'lodash';
import { connect } from 'react-redux';
import './navigation.scss';
import { Project, Config } from '../../../Proxy/Proxy';
import {Motion, spring} from 'react-motion';
import { getConfig } from '../../../Clusters/actions/clusters.actions';
import { bindActionCreators } from 'redux';

interface Props {
  routes: any
  params: any
  project: Project
  config: Config
}

interface DispatchProps {
  getConfig: Function
}


interface State {
  activeTopLevelPath: string
  isSubMenuActive: boolean
}

export class Navigation extends React.Component<Props & DispatchProps, any> {

  constructor() {
    super();
    this.state = {
      activeTopLevelPath: '',
      isSubMenuActive: false,
      config: {}
    };
  }

  componentWillMount(): void {
    this.setMenuState(this.props.routes);
    this.props.getConfig();
  }

  componentWillReceiveProps(nextProps: Props): void {
    this.setMenuState(nextProps.routes);
  }

  setMenuState(newRoutes: any[]): void {
    let currentRoutePath = newRoutes[newRoutes.length - 1].path;
    let topLevelPath = '';
    if (currentRoutePath) {
      topLevelPath = currentRoutePath.split('/')[0];
    }
    let submenuActive = false;
    _.forEach(routes[0].childRoutes, (route) => {
      if (this.isActive(route.path, newRoutes) && route.showChildrenAsSubmenu) {
        submenuActive = true;
      }
    });

    this.setState({
      activeTopLevelPath: topLevelPath,
      isSubMenuActive: submenuActive
    });
  }

  isActive(path: string, newRoutes: any[]): boolean {
    let currentRoutePath = newRoutes[newRoutes.length - 1].path;
    if (currentRoutePath && currentRoutePath.indexOf(path) !== -1) {
      return true;
    }
    return false;
  }

  getParentRouteName(currentPath: string): string {
    let newPathParts = currentPath.split('/');
    if (newPathParts.length < 2) {
      return currentPath;
    }
    newPathParts.pop();
    let newPath = newPathParts.join('/');
    let parentRoute = getRoute(newPath);
    return parentRoute.name;
  }

  logout() {
    $.ajax({
      url: window.location.protocol + '://' + window.location.host,
      beforeSend: function (xhr) {
        xhr.withCredentials = true;
        xhr.setRequestHeader('Authorization', 'Basic ' + btoa('fjkdshfhkjsdfjkhsdkfjhsdf:hfkjdshfdhff'));
      }
    });
  }

  renderSubmenu(activeRoute: any, shouldShow: boolean): JSX.Element {
    let childRoutes = routes[0].childRoutes.filter((route) => {
      return (route.path.indexOf(activeRoute.path) !== -1 && route.path !== activeRoute.path);
    });
    let styleTo;
    if (shouldShow) {
      styleTo = { left: spring(72) };
    } else {
      styleTo = { left: spring(300) };
    }
    return (
    <Motion defaultStyle={{left: 300}} style={styleTo}>
      {interpolatingStyle =>
      <div className="left-submenu" style={interpolatingStyle}>
        <Sidebar className='secondary-navigation'>
          <nav className="navigation--primary">
            <div className="navigation">
              <header>
                <div className="header-navigation">
                  <Link to={this.getParentRouteName(activeRoute.path)}><i
                    className="fa fa-angle-left"></i><span>{this.getParentRouteName(activeRoute.path)}</span></Link>
                </div>
              </header>
              <div className="header-content">{this.props.project.name}</div>
              <ul className="nav-list">
                {_.map(childRoutes, (menuItem: any) => {
                  let path = buildPath(menuItem.path, this.props.params);
                  return (!menuItem.showInNavigation) ? null : (
                    <li key={menuItem.path}
                        className={classNames('nav-list--item', {active: this.isActive(menuItem.path, this.props.routes)})}>
                      <Link to={path}>{menuItem.name}</Link>
                    </li>
                  );
                })}
              </ul>
            </div>
          </nav>
        </Sidebar>
      </div>
      }
    </Motion>
    );
  }

  render(): React.ReactElement<HTMLElement> {
    let submenu = <div></div>;
    return (
      <div className={classNames('nav-container')}>
        <Sidebar className="primary-navigation">
          <nav className="navigation--primary">
            <div className="navigation">
              <header>
                <div className="logo-container">
                  <Link to="/">
                    <div className="logo">STEAM</div>
                    <div>{this.props.config ? 'v' + this.props.config.version : null}</div>
                  </Link>
                </div>
              </header>
              <div className="header-content">
              </div>
              <ul className='nav-list'>
                {routes[0].childRoutes.map((route: any) => {
                  let isActive = false;
                  if (this.isActive(route.path, this.props.routes)) {
                    isActive = true;
                  }
                  if (route.showChildrenAsSubmenu) {
                    submenu = this.renderSubmenu(route, isActive);
                  }
                  if (route.path.split('/').length > 1 || !route.showInNavigation) {
                    return null;
                  }
                  let activeChildren = route.path === this.state.activeTopLevelPath && this.state.isSubMenuActive;
                  let path = '/' + route.path;
                  return (
                    <li key={path}
                        className={classNames('nav-list--item', { active: isActive}, {activeChildren: activeChildren}) }>
                      <Link to={path}><i className={route.icon}></i>
                        <div className="nav-list--label">{route.name}</div>
                      </Link>
                    </li>
                  );
                })
                }
                <li className="logout nav-list--item">
                  <a href="mailto:steam@h2o.ai?subject=STEAM: ">
                    <i className="fa fa-question-circle-o"/>
                    <div className="nav-list--label">
                      Support
                    </div>
                  </a>
                  <a onClick={this.logout.bind(this)}>
                    <i className="fa fa-sign-out"/>
                    <div className="nav-list--label">Logout</div>
                  </a>
                </li>
              </ul>
            </div>
          </nav>
        </Sidebar>
        {submenu}
      </div>
    );
  }
}


function mapStateToProps(state): any {
  return {
    project: state.projects.project,
    config: state.clusters.config
  };
}

function mapDispatchToProps(dispatch) {
  return {
    getConfig: bindActionCreators(getConfig, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(withRouter(Navigation));
