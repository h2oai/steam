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
                        return <li key={i}>{crumb}</li>;
                    }
                    return null;
                }) }
            </ol>
        );
    }
}
