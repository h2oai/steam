/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import { Link } from 'react-router';
import { Sidebar } from '../Sidebar/Sidebar';
import { buildPath } from '../../../App/utils/buildPath';
import { getRoute } from '../../../App/utils/getRoute';
import './navigation.scss';
import { routes } from '../../../routes';
import * as _ from 'lodash';
const logo = require('../../../../assets/h2o-home.png');

interface Props {
    routes: any
    params: any
}

interface DispatchProps {
}

interface State {
  activeTopLevelPath: string
  isSubMenuActive: boolean
}

export class Navigation extends React.Component<Props & DispatchProps, State> {

    constructor() {
        super();
        this.state = {
          activeTopLevelPath: '',
          isSubMenuActive: false
        };
    }

    componentWillMount(): void {
      this.setMenuState(this.props.routes);
    }

    componentWillReceiveProps(nextProps): void {
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

    renderSubmenu(activeRoute: any): JSX.Element {
      let childRoutes = routes[0].childRoutes.filter((route) => {
        return (route.path.indexOf(activeRoute.path) !== -1 && route.path !== activeRoute.path);
      });
      return (
        <Sidebar className='secondary-navigation'>
          <nav className="navigation--primary">
            <div className="navigation">
              <header>
                <div className="header-navigation">
                  <i className="fa fa-angle-left"></i><span>{this.getParentRouteName(activeRoute.path)}</span>
                </div>
              </header>
              <div className="header-content">{activeRoute.name}</div>
              <ul className="nav-list">
                {_.map(childRoutes, (menuItem: any) => {
                  let path = buildPath(menuItem.path, this.props.params);
                  return (!menuItem.showInNavigation) ? null : (
                    <li key={menuItem.path} className={classNames('nav-list--item', {active: this.isActive(menuItem.path, this.props.routes)})}>
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

    render(): React.ReactElement<HTMLElement> {
        let submenu = <div></div>;
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
                                {routes[0].childRoutes.map((route: any) => {
                                    let isActive = false;
                                    if (this.isActive(route.path, this.props.routes)) {
                                      isActive = true;
                                      if (route.showChildrenAsSubmenu) {
                                        submenu = this.renderSubmenu(route);
                                      }
                                    }
                                    if (route.path.split('/').length > 1 || !route.showInNavigation) {
                                      return null;
                                    }
                                    let activeChildren = route.path === this.state.activeTopLevelPath && this.state.isSubMenuActive;
                                    let path = '/' + route.path;
                                    return (
                                        <li key={path} className={classNames('nav-list--item', { active: isActive}, {activeChildren: activeChildren}) }>
                                            <Link to={path}><i className={route.icon}></i><div className="nav-list--label">{route.name}</div></Link>
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

export default Navigation;
