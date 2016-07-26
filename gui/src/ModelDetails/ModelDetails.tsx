/**
 * Created by justin on 6/27/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import * as $ from 'jquery';
import * as _ from 'lodash';
import Collapsible from './components/Collapsible';
import ModelOverview from './components/ModelOverview';
import GoodnessOfFit from './components/GoodnessOfFit';
import VariableImportance from './components/VariableImportance';
import PageHeader from '../Projects/components/PageHeader';
import ExportModal from './components/ExportModal';
import { hashHistory } from 'react-router';
import './styles/modeldetails.scss';
import { fetchModelOverview, downloadModel, deployModel } from './actions/model.overview.action';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

interface Props {
  params: {
    modelId: string
  },
  model: any
}

interface DispatchProps {
  fetchModelOverview: Function,
  downloadModel: Function,
  deployModel: Function
}

export class ModelDetails extends React.Component<Props & DispatchProps, any> {
  constructor() {
    super();
    this.state = {
      isModelOpen: true,
      isResidualOpen: true,
      isVariableOpen: true,
      isGoodnessOpen: true,
      isExportModalOpen: false
    };
    this.exportModel = this.exportModel.bind(this);
  }

  componentWillMount() {
    this.props.fetchModelOverview(parseInt(this.props.params.modelId, 10));
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
    /**
     * TODO(justinloyola): Backend not ready - pass projectId and port being eliminated as a parameter
     */
    this.props.deployModel(this.props.model.id, 12345);
  }

  render(): React.ReactElement<HTMLDivElement> {
    if (_.isEmpty(this.props.model)) {
      return <div></div>;
    }
    return (
      <div className="project-details">
        <ExportModal open={this.state.isExportModalOpen} name={this.props.model.name.toUpperCase()} onCancel={this.cancel.bind(this)} onDownload={this.downloadModel.bind(this)}/>
        <PageHeader>
          <span>{this.props.model.name.toUpperCase()}</span>
          <div className="buttons">
            <button className="default invert" onClick={this.exportModel.bind(this)}>Export Model</button>
            <button className="default" onClick={this.deployModel.bind(this)}>Deploy Model</button>
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
          <GoodnessOfFit></GoodnessOfFit>
        </Collapsible>
        <header className="overview-header">
          <span onClick={this.toggleOpen.bind(this, 'residual')}><i
            className={classNames('fa', {'fa-minus-square-o': this.state.isResidualOpen, 'fa-plus-square-o': !this.state.isResidualOpen})}></i
          >Residual Analysis</span>
        </header>
        <Collapsible open={this.state.isResidualOpen}>
          <div>
            Residual body
          </div>
        </Collapsible>
        <header className="overview-header">
          <span onClick={this.toggleOpen.bind(this, 'variable')}><i
            className={classNames('fa', {'fa-minus-square-o': this.state.isVariableOpen, 'fa-plus-square-o': !this.state.isVariableOpen})}></i
          >Variable Importance</span>
        </header>
        <Collapsible open={this.state.isVariableOpen}>
          <VariableImportance></VariableImportance>
        </Collapsible>
      </div>
    );
  }
}

function mapStateToProps(state: any): any {
  return {
    model: state.model
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchModelOverview: bindActionCreators(fetchModelOverview, dispatch),
    downloadModel: bindActionCreators(downloadModel, dispatch),
    deployModel: bindActionCreators(deployModel, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ModelDetails);

