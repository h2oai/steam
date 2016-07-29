/**
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import * as _ from 'lodash';
import Leaderboard from './components/Leaderboard';
import { fetchLeaderboard, filterModels } from './actions/leaderboard.actions';
import { Model } from '../Proxy/Proxy';

interface Props {
  leaderboard: any,
  params: {
    projectid: string
  }
}

interface DispatchProps {
  fetchLeaderboard: Function,
  filterModels: Function
}

export class Models extends React.Component<Props & DispatchProps, any> {
  componentWillMount(): void {
    if (_.isEmpty(this.props.leaderboard)) {
      this.props.fetchLeaderboard(parseInt(this.props.params.projectid, 10));
    }
  }

  onFilter(filters) {
    console.log(filters);
    this.props.filterModels(parseInt(this.props.params.projectid, 10), '', 'mse', filters.orderBy === 'asc');
  }

  render(): React.ReactElement<HTMLDivElement> {
    if (!this.props.leaderboard) {
      return <div></div>;
    }
    return (
      <div className="projects">
        <Leaderboard items={this.props.leaderboard} projectId={parseInt(this.props.params.projectid, 10)} onFilter={this.onFilter.bind(this)}></Leaderboard>
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
    filterModels: bindActionCreators(filterModels, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Models);
