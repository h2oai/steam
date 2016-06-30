/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';

import './styles/sidebar.scss';

interface Props {
  className: any
}

interface DispatchProps {
}

interface State {
  isOpen: boolean
}

export class Sidebar extends React.Component<Props & DispatchProps, State> {
  render(): React.ReactElement<HTMLElement> {
    return (
      <aside className={classNames('sidebar', this.props.className)}>
        {this.props.children}
      </aside>
    );
  }
}