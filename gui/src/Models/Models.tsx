/**
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import * as _ from 'lodash';
import Leaderboard from './components/Leaderboard';
import { fetchLeaderboard } from './actions/leaderboard.actions';

interface Props {
  leaderboard
}

interface DispatchProps {
  fetchLeaderboard: Function
}

export class Projects extends React.Component<Props & DispatchProps, any> {
  componentWillMount(): void {
    if (_.isEmpty(this.props.leaderboard)) {
      this.props.fetchLeaderboard();
    }
  }

  render(): React.ReactElement<HTMLDivElement> {
    if (!this.props.leaderboard) {
      return <div></div>;
    }

    return (
      <div className="projects">
        <Leaderboard items={this.props.leaderboard}></Leaderboard>
      </div>
    );
  }
}

function mapStateToProps(state: Props): Props {
  console.log(state);
  return {
    leaderboard: state.leaderboard.items
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    fetchLeaderboard: bindActionCreators(fetchLeaderboard, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Projects);
