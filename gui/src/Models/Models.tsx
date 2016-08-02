/**
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import * as _ from 'lodash';
import Leaderboard from './components/Leaderboard';
import { fetchLeaderboard } from './actions/leaderboard.actions';
import { Model } from '../Proxy/Proxy';
import { deployModel } from '../ModelDetails/actions/model.overview.action';

interface Props {
  leaderboard: any,
  params: {
    projectid: string
  }
}

interface DispatchProps {
  fetchLeaderboard: Function,
  deployModel: Function
}

export class Projects extends React.Component<Props & DispatchProps, any> {
  componentWillMount(): void {
    if (_.isEmpty(this.props.leaderboard)) {
      this.props.fetchLeaderboard(parseInt(this.props.params.projectid, 10));
    }
  }

  render(): React.ReactElement<HTMLDivElement> {
    if (!this.props.leaderboard) {
      return <div></div>;
    }
    return (
      <div className="projects">
        <Leaderboard items={this.props.leaderboard} projectId={parseInt(this.props.params.projectid, 10)} deployModel={this.props.deployModel}></Leaderboard>
      </div>
    );
  }
}

function mapStateToProps(state: any): any {
  return {
    leaderboard: state.leaderboard.items
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    fetchLeaderboard: bindActionCreators(fetchLeaderboard, dispatch),
    deployModel: bindActionCreators(deployModel, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Projects);
