/**
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import '../styles/breadcrumb.scss';

interface Props {
  crumbs: string[]
}

export default class Breadcrumb extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLElement> {
    return (
      <ol className="breadcrumb">
        {this.props.crumbs.map((crumb, i) => {
          return <li key={i}>{crumb}</li>;
        })}
      </ol>
    )
  }
}
