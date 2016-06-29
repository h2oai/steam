/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import { Link } from 'react-router';
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
      <nav className="navigation--primary">
        <div className="navigation">
          {this.props.children}    
        </div>
      </nav>
    );
  }
}