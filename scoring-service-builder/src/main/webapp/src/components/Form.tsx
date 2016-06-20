/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';

interface Props {
  className?: string,
  children?: any,
  id?: string,
  onSubmit?: Function
}

export default class Form extends React.Component<Props, any> {
  render() {
    return (
      <form
        id={this.props.id} autocomplete="off"
        className={this.props.className}
        onSubmit={this.props.onSubmit}>
        {this.props.children}
      </form>
    );
  }
}