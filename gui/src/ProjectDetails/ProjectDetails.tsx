/**
 * Created by justin on 6/27/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import Collapsible from './components/Collapsible';
import ModelOverview from './components/ModelOverview';
import GoodnessOfFit from './components/GoodnessOfFit';
import PageHeader from '../Projects/components/PageHeader';
import { hashHistory } from 'react-router';
import './styles/projectdetails.scss';

interface Props {
  params: any
}


export default class ProjectDetails extends React.Component<Props, any> {
  constructor() {
    super();
    this.state = {
      isModelOpen: true,
      isResidualOpen: true,
      isVariableOpen: true,
      isGoodnessOpen: true
    };
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
    hashHistory.push('/forkmodel');
  }

  deployModel() {
    return null;
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="project-details">
        <PageHeader>
          <span>DRF-1069085</span>
          <div className="buttons">
            <button className="default invert" onClick={this.forkModel}>Fork Model</button>
            <button className="default" onClick={this.deployModel}>Deploy Model</button>
          </div>
        </PageHeader>
        <header className="overview-header">
          <span onClick={this.toggleOpen.bind(this, 'model')}><i
            className={classNames('fa', {'fa-minus-square-o': this.state.isModelOpen, 'fa-plus-square-o': !this.state.isModelOpen})}></i
          >Model Overview</span>
        </header>
        <Collapsible open={this.state.isModelOpen}>
          <ModelOverview></ModelOverview>
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
          >Variable Analysis</span>
        </header>
        <Collapsible open={this.state.isVariableOpen}>
          <div>
            Variable body
          </div>
        </Collapsible>
      </div>
    );
  }
}
