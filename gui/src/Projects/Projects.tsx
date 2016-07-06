/**
 * Created by justin on 6/27/16.
 */

import * as React from 'react';
import * as _ from 'lodash';
import NewProject from './components/NewProject';

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
        <NewProject></NewProject>
      </div>
    );
  }
}