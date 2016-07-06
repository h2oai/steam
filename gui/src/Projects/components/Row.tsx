/**
 * Created by justin on 7/5/16.
 */
import * as React from 'react';
import * as classNames from 'classnames';
import '../styles/row.scss';

interface Props {
  header?: boolean,
  className?: any
  children?: React.ReactChildren
}

export default class Row extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className={classNames('row', {header: this.props.header}, this.props.className)}>
        {this.props.children}
      </div>
    );
  }
}