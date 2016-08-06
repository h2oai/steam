/**
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Leaderboard from './components/Leaderboard';
import { fetchLeaderboard, fetchSortCriteria } from './actions/leaderboard.actions';
import { fetchProject } from '../Projects/actions/projects.actions';
import { deployModel } from '../ModelDetails/actions/model.overview.action';

interface Props {
  leaderboard: any,
  params: {
    projectid: string
  },
  project: any,
  sortCriteria: string[]
}

interface DispatchProps {
  fetchLeaderboard: Function,
  deployModel: Function,
  fetchProject: Function,
  fetchSortCriteria: Function
}

export class Models extends React.Component<Props & DispatchProps, any> {
  constructor() {
    super();
    this.state = {
      modelCategory: null
    };
  }

  componentWillMount(): void {
    if (this.props.project) {
      this.props.fetchProject(parseInt(this.props.params.projectid, 10)).then((res) => {
        this.props.fetchLeaderboard(parseInt(this.props.params.projectid, 10), res.model_category);
        this.props.fetchSortCriteria(res.model_category.toLowerCase());
        this.setState({
          modelCategory: res.model_category.toLowerCase()
        });
      });
    }
  }

  onFilter(filters, name, offset) {
    this.props.fetchLeaderboard(parseInt(this.props.params.projectid, 10), this.state.modelCategory, name, filters.sortBy, filters.orderBy === 'asc', offset);
  }

  render(): React.ReactElement<HTMLDivElement> {
    if (!this.props.leaderboard) {
      return <div></div>;
    }
    return (
      <div className="projects">
        <Leaderboard items={this.props.leaderboard} projectId={parseInt(this.props.params.projectid, 10)}
                     modelCategory={this.state.modelCategory} sortCriteria={this.props.sortCriteria}
                     onFilter={this.onFilter.bind(this)} deployModel={this.props.deployModel}
                     fetchLeaderboard={this.props.fetchLeaderboard}></Leaderboard>
      </div>
    );
  }
}

function mapStateToProps(state: any): any {
  return {
    leaderboard: state.leaderboard.items,
    sortCriteria: state.leaderboard.criteria,
    project: state.projects.project
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    fetchLeaderboard: bindActionCreators(fetchLeaderboard, dispatch),
    deployModel: bindActionCreators(deployModel, dispatch),
    fetchSortCriteria: bindActionCreators(fetchSortCriteria, dispatch),
    fetchProject: bindActionCreators(fetchProject, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Models);
