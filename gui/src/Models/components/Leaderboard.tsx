/**
 * Created by justin on 6/27/16.
 */

import * as React from 'react';
import * as _ from 'lodash';
import Deploy from '../components/Deploy';
import PageHeader from '../../Projects/components/PageHeader';
import Pagination from '../components/Pagination';
import BinomialModelTable from './BinomialModelTable';
import MultinomialModelTable from './MultinomialModelTable';
import RegressionModelTable from './RegressionModelTable';
import ImportModelsModal from './ImportModelsModal';
import { MAX_ITEMS, linkLabelWithModel, unlinkLabelFromModel, findModelsCount, deleteModel } from '../actions/leaderboard.actions';
import '../styles/leaderboard.scss';
import { fetchLabels } from '../../Configurations/actions/configuration.labels.action';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchPackages } from '../../Deployment/actions/deployment.actions';

interface Props {
  items: any[],
  projectId: number,
  deployModel: Function,
  modelCategory: string,
  onFilter: Function,
  sortCriteria: string[],
  labels: any[],
  packages: string[],
  fetchLeaderboard: Function,
  count: number
}

interface DispatchProps {
  fetchPackages: Function,
  fetchLabels: Function,
  linkLabelWithModel: Function,
  unlinkLabelFromModel: Function,
  findModelsCount: Function,
  deleteModel: Function
}

export class Leaderboard extends React.Component<Props & DispatchProps, any> {
  refs: {
    [key: string]: Element
    filterModels: HTMLInputElement
  };
  sampleData = {};

  constructor(props) {
    super(props);
    this.state = {
      isDeployOpen: false,
      isImportModelsOpen: false,
      openDeployModel: null,
      currentPage: 0,
      filters: {
        sortBy: '',
        orderBy: 'asc'
      }
    };
    this.openDeploy = this.openDeploy.bind(this);
    this.closeHandler = this.closeHandler.bind(this);
    this.onChangeHandler = this.onChangeHandler.bind(this);
  }

  componentWillMount() {
    if (!this.props.labels || !this.props.labels[this.props.projectId]) {
      this.props.fetchLabels(this.props.projectId);
    }
    this.props.fetchPackages(this.props.projectId);
    this.props.findModelsCount(this.props.projectId);
  }

  openDeploy(model): void {
    this.setState({
      isDeployOpen: true,
      openDeployModel: model
    });
  }

  openImportModels() {
    this.setState({
      isImportModelsOpen: true
    });
  }

  closeImportModels() {
    this.setState({
      isImportModelsOpen: false
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
    if (this.state.currentPage > 0) {
      this.setState({
        currentPage: --this.state.currentPage
      });
      this.props.onFilter(this.state.filters, this.refs.filterModels.value, this.state.currentPage * MAX_ITEMS);
    }
  }

  onDeploy(model, serviceName, packageName) {
    this.setState({
      isDeployOpen: false
    });
    this.props.deployModel(model.id, serviceName, this.props.projectId, packageName);
  }

  onChangeHandler(labelId, modelId, isUnlink) {
    if (isUnlink === true) {
      this.props.unlinkLabelFromModel(labelId, modelId).then(() => {
        this.props.fetchLabels(this.props.projectId);
      });
    } else {
      this.props.linkLabelWithModel(labelId, modelId).then(() => {
        this.props.fetchLabels(this.props.projectId);
      });
    }
  }

  getDataset() {
    return _.get(this.props, 'items[0].dataset_name');
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div ref="leaderboard" className="leaderboard">
        <ImportModelsModal projectId={this.props.projectId} open={this.state.isImportModelsOpen}
                           onCancel={this.closeImportModels.bind(this)} fetchLeaderboard={this.props.fetchLeaderboard}
                           modelCategory={this.props.modelCategory}
                           datasetName={this.getDataset()}/>
        <Deploy open={this.state.isDeployOpen} onCancel={this.closeHandler} model={this.state.openDeployModel}
                onDeploy={this.onDeploy.bind(this)} packages={this.props.packages || []}></Deploy>
        <PageHeader>
          <span>Models</span>
          <div className="buttons">
            <button className="default" onClick={this.openImportModels.bind(this)}>Import Models</button>
          </div>
        </PageHeader>
        <div className="filter">
          <input ref="filterModels" type="text" placeholder="filter models" onChange={this.onFilter.bind(this)}/>
        </div>
        {this.props.modelCategory === 'binomial' ?
          <BinomialModelTable onFilter={this.onFilter.bind(this)} sortCriteria={this.props.sortCriteria}
                              items={this.props.items} projectId={this.props.projectId}
                              openDeploy={this.openDeploy.bind(this)} labels={this.props.labels}
                              onChangeHandler={this.onChangeHandler} deleteModel={this.props.deleteModel}
                              fetchLeaderboard={() => { return this.props.fetchLeaderboard(this.props.projectId, this.props.modelCategory); }} /> : null}
        {this.props.modelCategory === 'multinomial' ?
          <MultinomialModelTable onFilter={this.onFilter.bind(this)} sortCriteria={this.props.sortCriteria}
                                 items={this.props.items} projectId={this.props.projectId}
                                 openDeploy={this.openDeploy.bind(this)} labels={this.props.labels}
                                 onChangeHandler={this.onChangeHandler} deleteModel={this.props.deleteModel}
                                 fetchLeaderboard={() => { return this.props.fetchLeaderboard(this.props.projectId, this.props.modelCategory); }} /> : null}
        {this.props.modelCategory === 'regression' ?
          <RegressionModelTable onFilter={this.onFilter.bind(this)} sortCriteria={this.props.sortCriteria}
                                items={this.props.items} projectId={this.props.projectId}
                                openDeploy={this.openDeploy.bind(this)} labels={this.props.labels}
                                onChangeHandler={this.onChangeHandler} deleteModel={this.props.deleteModel}
                                fetchLeaderboard={() => { return this.props.fetchLeaderboard(this.props.projectId, this.props.modelCategory); }} /> : null}
        <Pagination items={this.props.items} onPageBack={this.onPageBack.bind(this)}
                    onPageForward={this.onPageForward.bind(this)} currentPage={this.state.currentPage} count={this.props.count}></Pagination>
      </div>
    );
  }
}

function mapStateToProps(state: any): any {
  return {
    count: state.leaderboard.count,
    labels: state.labels,
    packages: state.deployments.packages
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchLabels: bindActionCreators(fetchLabels, dispatch),
    linkLabelWithModel: bindActionCreators(linkLabelWithModel, dispatch),
    unlinkLabelFromModel: bindActionCreators(unlinkLabelFromModel, dispatch),
    fetchPackages: bindActionCreators(fetchPackages, dispatch),
    findModelsCount: bindActionCreators(findModelsCount, dispatch),
    deleteModel: bindActionCreators(deleteModel, dispatch)
  };
}

export default connect<Props, any, any>(mapStateToProps, mapDispatchToProps)(Leaderboard);
