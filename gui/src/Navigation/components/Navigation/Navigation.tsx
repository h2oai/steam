/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import { Sidebar } from '../Sidebar/Sidebar';
import * as classNames from 'classnames';

import './navigation.scss';

interface Props {
}

interface DispatchProps {
}

interface State {
  isOpen: boolean
}

export class Navigation extends React.Component<Props & DispatchProps, State> {
  render(): React.ReactElement<HTMLElement> {
    return (
      <Sidebar>
        <nav className="navigation__div">
          <ul className="navigation__div__ul">
            {this.props.children}
          </ul>
        </nav>
      </Sidebar>
    );
  }
}