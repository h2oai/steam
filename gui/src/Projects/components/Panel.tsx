/**
 * Created by justin on 7/12/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import '../styles/panel.scss';

interface Props {
  className?: any,
  onClick?: Function
}

export default class Panel extends React.Component<Props, any> {
  render() {
    return (
      <div className={classNames('panel', this.props.className)} onClick={this.props.onClick}>
        {this.props.children}
      </div>
    );
  }
}
