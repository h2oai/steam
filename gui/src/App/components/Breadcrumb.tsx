/**
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import { Link } from 'react-router';
import * as _ from 'lodash';
import '../styles/breadcrumb.scss';
import { routes } from '../../routes';

interface Props {
  routes: any,
  params: any
}

export default class Breadcrumb extends React.Component<Props, any> {

  render(): React.ReactElement<HTMLElement> {

    // create a new array, including only routes that we want to show in the breadcrumb
    let crumbs = [];

    _.forEach(this.props.routes, (route) => {
      if (route.showInBreadcrumb) {
        if (route.path.indexOf('/:') > -1) {
          // split into two
          let pair = route.path.split('/:');
          crumbs.push({
            name: route.name,
            path: pair[0]
          });
          crumbs.push({
            name: this.props.params[pair[1]],
            path: this.props.params[pair[1]]
          });
        } else {
          crumbs.push(_.assign({}, route));
        }
      }
    });

    let path = '';

    return (
      <ol className="breadcrumb">
        {crumbs.map((route, i) => {
          if (crumbs.length > 1) {
            path = (i > 0) ? path += '/' + route.path : path;
            let crumb = (i === crumbs.length - 1) ? route.name : (<Link to={path}>{route.name}</Link>);
            return <li key={i}>{crumb}</li>;
          }
          return null;
        })}
      </ol>
    );
  }
}
