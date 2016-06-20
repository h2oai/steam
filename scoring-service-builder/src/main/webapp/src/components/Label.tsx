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

export default class Label extends React.Component<Props, any> {
  render() {
    return (
      <label id={this.props.id} className={classNames('form-control-label', this.props.className)}>
        {this.props.children}
      </label>
    );
  }
}