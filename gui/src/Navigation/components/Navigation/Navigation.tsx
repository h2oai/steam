/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import { Link, withRouter } from 'react-router';
import { Sidebar } from '../Sidebar/Sidebar';
import './navigation.scss';
import * as logo from '../../../../assets/h2o-home.png';

interface Props {
  router: any
}

interface DispatchProps {
}

interface State {
  isOpen: boolean
}

export class Navigation extends React.Component<Props & DispatchProps, State> {
  sitemap = {
    projects: {
      path: 'projects',
      label: 'Projects',
      icon: 'fa fa-folder',
      childRoutes: [
        {
          path: 'models',
          label: 'Models'
        },
        {
          path: 'projects/dataframes',
          label: 'Dataframes'
        },
        {
          path: 'projects/repository',
          label: 'Repository'
        },
        {
          path: 'projects/configurations',
          label: 'Configurations'
        },
        {
          path: 'projects/deployments',
          label: 'Deployments'
        },
        {
          path: 'projects/collaborators',
          label: 'Collaborators'
        }
      ]
    },
    services: {
      path: 'services',
      label: 'Services',
      icon: 'fa fa-cloud',
      childRoutes: []
    },
    clusters: {
      path: 'clusters',
      label: 'Clusters',
      icon: 'fa fa-cube',
      childRoutes: [
        {
          path: 'models',
          label: 'Sub 1'
        }
      ]
    },
    team: {
      path: 'team',
      label: 'Team',
      icon: 'fa fa-users',
      childRoute: []
    }
  };

  constructor() {
    super();
    this.openSubmenu = this.openSubmenu.bind(this);
    this.closeSubmenu = this.closeSubmenu.bind(this);
    this.state = {
      isOpen: false
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
    return Object.keys(this.sitemap).filter((key) => {
      return this.props.router.isActive(key, true);
    })[0];
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
                {Object.keys(this.sitemap).map((route: string, i: number) => {
                    return (
                      <li key={i} className={classNames('nav-list--item', {active: this.getPath() === route && !this.sitemap[route].childRoutes})} onMouseOver={this.openSubmenu}>
                        <Link to={this.sitemap[route].path}><i className={this.sitemap[route].icon}></i><div className="nav-list--label">{this.sitemap[route].label}</div></Link>
                      </li>
                    );
                  })}
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
                {this.sitemap[this.getPath()] ? this.sitemap[this.getPath()].childRoutes.map((route, i: number) => {
                  return (
                    <li key={i} className={classNames('nav-list--item', {active: this.getPath() === route.path})}>
                      <Link to={route.path}>{route.label}</Link>
                    </li>
                  );
                }) : null}
              </ul>
            </div>
          </nav>
        </Sidebar>
      </div>
    );
  }
}

export default withRouter(Navigation);
