/**
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import * as _ from 'lodash';
import Leaderboard from './components/Leaderboard';
import { fetchModelsFromProject, fetchProject } from '../Projects/actions/projects.actions';
import { fetchLeaderboard } from './actions/leaderboard.actions';

interface Props {
  leaderboard: any,
  params: {
    projectid: string
  },
  project: any
}

interface DispatchProps {
  fetchLeaderboard: Function,
  fetchProject: Function
}

export class Models extends React.Component<Props & DispatchProps, any> {
  componentWillMount(): void {
    if (_.isEmpty(this.props.project)) {
      this.props.fetchProject(parseInt(this.props.params.projectid, 10));
    }
  }

  componentWillReceiveProps(nextProps) {
    if (_.isEmpty(this.props.leaderboard) && nextProps.project.model_category) {
      this.props.fetchLeaderboard(parseInt(this.props.params.projectid, 10), nextProps.project.model_category);
    }
  }

  onFilter(filters) {
    this.props.fetchLeaderboard(parseInt(this.props.params.projectid, 10), this.props.project.model_category, filters.sortBy, filters.orderBy === 'asc');
  }

  render(): React.ReactElement<HTMLDivElement> {
    if (!this.props.leaderboard) {
      return <div></div>;
    }
    return (
      <div className="projects">
        <Leaderboard items={this.props.leaderboard} projectId={parseInt(this.props.params.projectid, 10)} modelCategory={this.props.project.model_category} onFilter={this.onFilter.bind(this)}></Leaderboard>
      </div>
    );
  }
}

function mapStateToProps(state: any): any {
  return {
    leaderboard: state.leaderboard.items,
    project: state.projects.project
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    fetchLeaderboard: bindActionCreators(fetchLeaderboard, dispatch),
    fetchProject: bindActionCreators(fetchProject, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Models);
