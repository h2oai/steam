/**
 * Created by justin on 6/27/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import * as _ from 'lodash';
import Collapsible from './components/Collapsible';
import ModelOverview from './components/ModelOverview';
import GoodnessOfFit from './components/GoodnessOfFit';
import PageHeader from '../Projects/components/PageHeader';
import ExportModal from './components/ExportModal';
import Deploy from '../Models/components/Deploy';
import ModelSelectionModal from './components/ModelSelectionModal';
import { hashHistory } from 'react-router';
import './styles/modeldetails.scss';
import { fetchModelOverview, downloadModel, deployModel } from './actions/model.overview.action';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchLeaderboard, fetchSortCriteria } from '../Models/actions/leaderboard.actions';
import { fetchProject } from '../Projects/actions/projects.actions';
import { fetchPackages } from '../Deployment/actions/deployment.actions';

interface Props {
  params: {
    modelid: string,
    projectid: string
  },
  model: any,
  models: any,
  project: any,
  sortCriteria: string[],
  packages: string[]
}

interface DispatchProps {
  fetchModelOverview: Function,
  fetchSortCriteria: Function,
  fetchLeaderboard: Function,
  fetchProject: Function,
  downloadModel: Function,
  deployModel: Function,
  fetchPackages: Function
}

export class ModelDetails extends React.Component<Props & DispatchProps, any> {
  constructor() {
    super();
    this.state = {
      isModelOpen: true,
      isResidualOpen: true,
      isVariableOpen: true,
      isGoodnessOpen: true,
      isExportModalOpen: false,
      isModelSelectionModal: false,
      comparisonModel: null
    };
    this.exportModel = this.exportModel.bind(this);
  }

  componentWillMount() {
    if (this.props.project) {
      this.props.fetchProject(parseInt(this.props.params.projectid, 10)).then((res) => {
        this.props.fetchLeaderboard(parseInt(this.props.params.projectid, 10), res.model_category);
        this.props.fetchSortCriteria(res.model_category.toLowerCase());
        this.setState({
          modelCategory: res.model_category.toLowerCase()
        });
      });
    }
    this.props.fetchPackages(parseInt(this.props.params.projectid, 10));
    this.props.fetchModelOverview(parseInt(this.props.params.modelid, 10));
  }

  toggleOpen(accordian: string) {
    /**
     * TODO(justinloyola): Fix the asynchronous state change issues
     */
    if (accordian === 'model') {
      this.setState({
        isModelOpen: !this.state.isModelOpen
      });
    } else if (accordian === 'residual') {
      this.setState({
        isResidualOpen: !this.state.isResidualOpen
      });
    } else if (accordian === 'variable') {
      this.setState({
        isVariableOpen: !this.state.isVariableOpen
      });
    } else if (accordian === 'goodness') {
      this.setState({
        isGoodnessOpen: !this.state.isGoodnessOpen
      });
    }
  }

  forkModel() {
    hashHistory.push('/projects/forkmodel');
  }

  exportModel() {
    this.setState({
      isExportModalOpen: !this.state.isExportModalOpen
    });
  }

  cancel() {
    this.setState({
      isExportModalOpen: false
    });
  }

  downloadModel(event) {
    event.preventDefault();
    this.props.downloadModel(event);
  }

  deployModel() {
    this.setState({
      isDeployModalOpen: true
    });
  }

  openComparisonModal() {
    this.setState({
      isModelSelectionModalOpen: true
    });
  }

  closeComparisonModal() {
    this.setState({
      isModelSelectionModalOpen: false
    });
  }

  onSelectModel(model) {
    this.closeComparisonModal();
    this.setState({
      comparisonModel: model
    });
  }

  onCancel() {
    this.closeComparisonModal();
  }

  onFilter(filters, name, offset) {
    this.props.fetchLeaderboard(parseInt(this.props.params.projectid, 10), this.state.modelCategory, name, filters.sortBy, filters.orderBy === 'asc', offset);
  }

  closeDeployModal() {
    this.setState({
      isDeployModalOpen: false
    });
  }

  onDeploy(model, serviceName, packageName) {
    console.log(packageName);
    this.setState({
      isDeployModalOpen: false
    });
    this.props.deployModel(model.id, serviceName, this.props.params.projectid, packageName);
  }

  render(): React.ReactElement<HTMLDivElement> {
    if (_.isEmpty(this.props.model)) {
      return <div></div>;
    }
    return (
      <div className="model-details">
        <ModelSelectionModal open={this.state.isModelSelectionModalOpen}
                             onFilter={this.onFilter.bind(this)}
                             models={this.props.models}
                             sortCriteria={this.props.sortCriteria}
                             onSelectModel={this.onSelectModel.bind(this)}
                             onCancel={this.onCancel.bind(this)}/>
        <ExportModal open={this.state.isExportModalOpen} name={this.props.model.name.toUpperCase()}
                     onCancel={this.cancel.bind(this)} modelId={parseInt(this.props.params.modelid, 10)} projectId={parseInt(this.props.params.projectid, 10)} onDownload={this.downloadModel.bind(this)}/>
        <Deploy open={this.state.isDeployModalOpen} onCancel={this.closeDeployModal.bind(this)} model={this.props.model} onDeploy={this.onDeploy.bind(this)} packages={this.props.packages}></Deploy>
        <PageHeader>
          <span>{this.props.model.name.toUpperCase()}</span>
          <div className="buttons">
            <button className="default invert" onClick={this.exportModel.bind(this)}>Export Model</button>
            <button className="default" onClick={this.deployModel.bind(this)}>Deploy Model</button>
          </div>
          <div className="comparison-selection">
            <span><span>compared to:</span><button className={classNames('model-selection-button', {selected: this.state.comparisonModel})}
                                                   onClick={this.openComparisonModal.bind(this)}>{this.state.comparisonModel ? this.state.comparisonModel.name : 'SELECT MODEL FOR COMPARISON'}</button></span>
          </div>
        </PageHeader>
        <header className="overview-header">
          <span onClick={this.toggleOpen.bind(this, 'model')}><i
            className={classNames('fa', {'fa-minus-square-o': this.state.isModelOpen, 'fa-plus-square-o': !this.state.isModelOpen})}></i
          >Model Overview</span>
        </header>
        <Collapsible open={this.state.isModelOpen}>
          <ModelOverview model={this.props.model}></ModelOverview>
        </Collapsible>
        <header className="overview-header">
          <span onClick={this.toggleOpen.bind(this, 'goodness')}><i
            className={classNames('fa', {'fa-minus-square-o': this.state.isGoodnessOpen, 'fa-plus-square-o': !this.state.isGoodnessOpen})}></i
          >Goodness of Fit</span>
        </header>
        <Collapsible open={this.state.isGoodnessOpen}>
          <GoodnessOfFit model={this.props.model} comparisonModel={this.state.comparisonModel} modelCategory={this.state.modelCategory}></GoodnessOfFit>
        </Collapsible>
      </div>
    );
  }
}

function mapStateToProps(state: any): any {
  return {
    model: state.model,
    packages: state.deployments.packages,
    project: state.projects.project,
    models: state.leaderboard.items,
    sortCriteria: state.leaderboard.criteria
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchLeaderboard: bindActionCreators(fetchLeaderboard, dispatch),
    fetchProject: bindActionCreators(fetchProject, dispatch),
    fetchSortCriteria: bindActionCreators(fetchSortCriteria, dispatch),
    fetchModelOverview: bindActionCreators(fetchModelOverview, dispatch),
    downloadModel: bindActionCreators(downloadModel, dispatch),
    deployModel: bindActionCreators(deployModel, dispatch),
    fetchPackages: bindActionCreators(fetchPackages, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ModelDetails);
