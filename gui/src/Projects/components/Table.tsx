/**
 * Created by justin on 7/5/16.
 */


import * as React from 'react';
import * as classNames from 'classnames';
import '../styles/table.scss';

interface Props {
  className?: any
  children?: React.ReactChildren
}

export default class Table extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className={classNames('table', this.props.className)}>
        {this.props.children}
      </div>
    );
  }
}