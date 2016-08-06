/**
 * Created by justin on 6/27/16.
 */

import * as React from 'react';
import Deploy from '../components/Deploy';
import PageHeader from '../../Projects/components/PageHeader';
import Pagination from '../components/Pagination';
import BinomialModelTable from './BinomialModelTable';
import MultinomialModelTable from './MultinomialModelTable';
import RegressionModelTable from './RegressionModelTable';
import { MAX_ITEMS } from '../actions/leaderboard.actions';
import '../styles/leaderboard.scss';

interface Props {
  items: any[],
  projectId: number,
  deployModel: Function,
  modelCategory: string,
  onFilter: Function,
  sortCriteria: string[]
}

interface DispatchProps {
}

export default class Leaderboard extends React.Component<Props & DispatchProps, any> {
  refs: {
    [key: string]: Element
    filterModels: HTMLInputElement
  };
  sampleData = {};

  constructor() {
    super();
    this.state = {
      isDeployOpen: false,
      openDeployModel: null,
      currentPage: 0,
      filters: {
        sortBy: '',
        orderBy: 'asc'
      }
    };
    this.openDeploy = this.openDeploy.bind(this);
    this.closeHandler = this.closeHandler.bind(this);
  }

  openDeploy(model): void {
    this.setState({
      isDeployOpen: true,
      openDeployModel: model
    });
  }

  closeHandler(): void {
    this.setState({
      isDeployOpen: false
    });
  }

  onFilter(filters) {
    this.setState({
      filters: filters
    });
    this.props.onFilter(filters, this.refs.filterModels.value);
  }

  onPageForward() {
    this.setState({
      currentPage: ++this.state.currentPage
    });
    this.props.onFilter(this.state.filters, this.refs.filterModels.value, this.state.currentPage * MAX_ITEMS);
  }

  onPageBack() {
    if (this.state.currentPage >= 0) {
      this.setState({
        currentPage: --this.state.currentPage
      });
      this.props.onFilter(this.state.filters, this.refs.filterModels.value, this.state.currentPage * MAX_ITEMS);
    }
  }

  onDeploy(model, name) {
    this.setState({
      isDeployOpen: false
    });
    this.props.deployModel(model.id, name);
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div ref="leaderboard" className="leaderboard">
        <Deploy open={this.state.isDeployOpen} onCancel={this.closeHandler} model={this.state.openDeployModel} onDeploy={this.onDeploy.bind(this)}></Deploy>
        <PageHeader>
          <span>Models</span>
          <div className="buttons">
            <button className="default">Import Model</button>
          </div>
        </PageHeader>
        <div className="filter">
          <input ref="filterModels" type="text" placeholder="filter models" onChange={this.onFilter.bind(this)}/>
        </div>
        {this.props.modelCategory === 'binomial' ? <BinomialModelTable onFilter={this.onFilter.bind(this)} sortCriteria={this.props.sortCriteria} items={this.props.items} projectId={this.props.projectId} openDeploy={this.openDeploy.bind(this)}/> : null}
        {this.props.modelCategory === 'multinomial' ? <MultinomialModelTable onFilter={this.onFilter.bind(this)} sortCriteria={this.props.sortCriteria} items={this.props.items} projectId={this.props.projectId} openDeploy={this.openDeploy.bind(this)}/> : null}
        {this.props.modelCategory === 'regression' ? <RegressionModelTable onFilter={this.onFilter.bind(this)} sortCriteria={this.props.sortCriteria} items={this.props.items} projectId={this.props.projectId} openDeploy={this.openDeploy.bind(this)}/> : null}
        <Pagination items={this.props.items} onPageBack={this.onPageBack.bind(this)} onPageForward={this.onPageForward.bind(this)}></Pagination>
      </div>
    );
  }
}
