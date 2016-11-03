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
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import { Link } from 'react-router';
import * as _ from 'lodash';
import { buildPath } from '../utils/buildPath';
import { getRoute } from '../utils/getRoute';
import '../styles/breadcrumb.scss';
import { routes } from '../../routes';

interface Props {
    routes: any
    params: any
}

interface Crumb {
    name: string
    path: string
}

export default class Breadcrumb extends React.Component<Props, any> {

    render(): React.ReactElement<HTMLElement> {
        let crumbs: Crumb[] = [];
        if (this.props.routes && this.props.routes[this.props.routes.length - 1] && this.props.routes[this.props.routes.length - 1].path) {
            let pathParts = this.props.routes[this.props.routes.length - 1].path.split('/');
            let path = '';
            pathParts.forEach((part) => {
                path += (path === '') ? part : '/' + part;
                let route = getRoute(path);
                if (!route.showInBreadcrumb) {
                  return;
                }
                let crumb: Crumb = {
                    path: buildPath(route.path, this.props.params),
                    name: this.props.params[part.slice(1, part.length)] || route.name
                };
                crumbs.push(crumb);
            });
        }
        crumbs.unshift({
            path: '/',
            name: 'Home'
        });

        let breadClass = (crumbs.length <= 1) ? "breadcrumb" : "filled breadcrumb";

        return (
            <ol className={breadClass}>
                {crumbs.map((route, i) => {
                    if (crumbs.length > 1) {
                        let crumb = (i === crumbs.length - 1) ? route.name : (<Link to={route.path}>{route.name}</Link>);
                        if (i === 2) {
                          if (crumbs[1].path === "/projects" ) {
                            return <li id="projectIdCrumb" key={i}>{crumb}</li>;
                          }
                        }
                        if (i === 4) {
                          if (crumbs[3].path.indexOf("/models") === crumbs[3].path.length - "/models".length) {
                            return <li id="modelIdCrumb" key={i}>{crumb}</li>;
                          }
                        }
                        return <li key={i}>{crumb}</li>;
                    }
                    return null;
                }) }
            </ol>
        );
    }
}
