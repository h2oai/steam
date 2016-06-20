/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';

interface Props {
  className?: string,
  children?: any,
  id?: string,
}

export default class Row extends React.Component<Props, any> {
  render() {
    return (
      <div id={this.props.id} className={classNames('row', this.props.className)}>
        {this.props.children}
      </div>
    );
  }
}