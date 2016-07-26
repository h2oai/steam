/**
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import '../styles/cell.scss';

interface Props {
  children?: React.ReactChildren,
  className?: any
}

export default class Cell extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className={classNames('cell', this.props.className)}>
        {this.props.children}
      </div>
    );
  }
}