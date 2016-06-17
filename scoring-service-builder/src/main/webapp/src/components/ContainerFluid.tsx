/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';

interface Props {
  className?: string,
  children?: any,
  id?: string
}

export default class ContainerFluid extends React.Component<Props, any> {
  render() {
    return (
      <main id={this.props.id} className={classNames('container-fluid', this.props.className)}>
        {this.props.children}
      </main>
    );
  }
}