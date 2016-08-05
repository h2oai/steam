/**
 * Created by justin on 6/27/16.
 */

import * as React from 'react';
import * as $ from 'jquery';
import { Link } from 'react-router';
import Deploy from '../components/Deploy';
import RocGraph from '../components/RocGraph';
import PageHeader from '../../Projects/components/PageHeader';
import Pagination from '../components/Pagination';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import FilterDropdown from './FilterDropdown';
import { getOrdinal } from '../../App/utils/getOrdinal';
import ModelLabelSelect from './ModelLabelSelect';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchLabels } from '../../Configurations/actions/configuration.labels.action';
import '../styles/leaderboard.scss';

// sample data
import { deeplearningTrain } from '../data/deeplearningTrain';
import { deeplearningValidation } from '../data/deeplearningValidation';
import { drfTrain } from '../data/drfTrain';
import { drfValidation } from '../data/drfValidation';
import { gbmTrain } from '../data/gbmTrain';
import { gbmValidation } from '../data/gbmValidation';
import { glmTrain } from '../data/glmTrain';
import { glmValidation } from '../data/glmValidation';
import { naivebayesTrain } from '../data/naivebayesTrain';
import { naivebayesValidation } from '../data/naivebayesValidation';
import { MAX_ITEMS } from '../actions/leaderboard.actions';

interface Props {
  items: any[],
  projectId: number,
  deployModel: Function,
  modelCategory: string,
  onFilter: Function,
  sortCriteria: string[],
  labels: any[],
  fetchLabels: Function
}

interface DispatchProps {
}

class Leaderboard extends React.Component<Props & DispatchProps, any> {
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
    this.sampleData = {
      deeplearningTrain,
      deeplearningValidation,
      drfTrain,
      drfValidation,
      gbmTrain,
      gbmValidation,
      glmTrain,
      glmValidation,
      naivebayesTrain,
      naivebayesValidation
    };
  }

  componentWillMount() {
      if (!this.props.labels || !this.props.labels[this.props.projectId]) {
          this.props.fetchLabels(this.props.projectId);
      }
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
        <Table>
          <Row header={true}>
            <Cell>
              <FilterDropdown onFilter={this.onFilter.bind(this)} sortCriteria={this.props.sortCriteria}/>
            </Cell>
            <Cell>
              MODEL
            </Cell>
            <Cell className="graph">
              TRAIN ROC
            </Cell>
            <Cell className="graph">
              TEST ROC
            </Cell>
            <Cell>
              <div className="actions">
                ACTIONS
              </div>
            </Cell>
          </Row>
          {this.props.items.map((item, i) => {
            return (
              <Row key={i}>
                <Cell>{item.id + getOrdinal(item.id)}</Cell>
                <Cell>
                  <div className="metadata">
                    <div className="model-name">
                      {item.name}
                    </div>
                    <div>
                      {item.cluster_name}
                    </div>
                    <div>
                      {item.createdAt}
                    </div>
                    <div>
                      {item.max_runtime}
                    </div>
                  </div>
                </Cell>
                <Cell className="graph">
                  <RocGraph data={this.sampleData['gbmTrain']}/>
                </Cell>
                <Cell className="graph">
                  <RocGraph data={this.sampleData['gbmValidation']}/>
                </Cell>
                <Cell>
                  <ul className="actions">
                    <li><Link to={'/projects/' + this.props.projectId + '/models/' + item.id}><span><i className="fa fa-eye"></i></span><span>view model details</span></Link></li>
                    <li className="labels"><span><i className="fa fa-tags"></i></span> label as
                        <span className="label-selector">
                          <ModelLabelSelect projectId={this.props.projectId} modelId={item.id} labels={this.props.labels}/>
                        </span>
                    </li>
                    <li onClick={this.openDeploy.bind(this, item)}><span><i className="fa fa-arrow-up"></i></span> <span>deploy model</span></li>
                  </ul>
                </Cell>
              </Row>
            );
          })}
        </Table>
        <Pagination items={this.props.items} onPageBack={this.onPageBack.bind(this)} onPageForward={this.onPageForward.bind(this)}></Pagination>
      </div>
    );
  }
}

function mapStateToProps(state: any): any {
    return {
        labels: state.labels
    };
}

function mapDispatchToProps(dispatch) {
    return {
        fetchLabels: bindActionCreators(fetchLabels, dispatch),
    };
}

export default connect<Props, any, any>(mapStateToProps, mapDispatchToProps)(Leaderboard);
