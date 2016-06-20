/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';

interface Props {
  onChange?: Function,
  className?: string,
  children?: any,
  id?: string,
  name?: string
}

export default class Select extends React.Component<Props, any> {
  render() {
    return (
      <select name={this.props.name} id={this.props.id} className={classNames('form-control', this.props.className)} onChange={this.props.onChange}>
        {this.props.children}
      </select>
    );
  }
}