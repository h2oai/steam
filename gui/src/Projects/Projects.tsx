/**
 * Created by justin on 6/27/16.
 */

import * as React from 'react';
import ProjectScreenStrategy from './components/ProjectScreenStrategy';

interface Props {
  leaderboard
}

interface DispatchProps {
  fetchLeaderboard: Function
}

export default class Projects extends React.Component<Props & DispatchProps, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="projects">
        {this.props.children}
      </div>
    );
  }
}
