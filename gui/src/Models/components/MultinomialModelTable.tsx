/**
 * Created by justin on 8/5/16.
 */
import * as React from 'react';
import * as moment from 'moment';
import * as _ from 'lodash';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import FilterDropdown from './FilterDropdown';
import RocGraph from './RocGraph';
import { Link } from 'react-router';

interface Props {
  onFilter: Function,
  sortCriteria: string[],
  items: any,
  projectId: number,
  openDeploy: Function
}

interface MultinomialMetrics {
  nobs: number,
  MSE: number,
  r2: number,
  logloss: number
}

export default class MultinomialModelTable extends React.Component<Props, any> {
  render() {
    return (
      <Table>
        <Row header={true}>
          <Cell>
            <FilterDropdown onFilter={this.props.onFilter.bind(this)} sortCriteria={this.props.sortCriteria}/>
          </Cell>
          <Cell>
            MODEL
          </Cell>
          <Cell>
            MSE
          </Cell>
          <Cell>
            Logloss
          </Cell>
          <Cell>
            R<sup>2</sup>
          </Cell>
          <Cell>
            <div className="actions">
              ACTIONS
            </div>
          </Cell>
        </Row>
        {this.props.items.map((model, i) => {
          let modelMetrics = JSON.parse(model.metrics);
          let trainingMetrics: MultinomialMetrics = _.get(modelMetrics, 'models[0].output.training_metrics', {}) as MultinomialMetrics;
          let fpr = _.get(modelMetrics, 'models[0].output.training_metrics.thresholds_and_metric_scores.data[17]', []);
          let tpr = _.get(modelMetrics, 'models[0].output.training_metrics.thresholds_and_metric_scores.data[18]', []);
          let data = [];
          tpr.map((val, i) => {
            data.push({
              tpr: val,
              fpr: fpr[i]
            });
          });
          return (
            <Row key={i}>
              <Cell></Cell>
              <Cell>
                <div className="metadata">
                  <div className="model-name">
                    {model.name}
                  </div>
                  <div>
                    <span>Created at:&nbsp;</span><span>{moment.unix(model.created_at).format('YYYY-MM-DD hh:mm:ss')}</span>
                  </div>
                  <div>
                    <span>Num of Observations:&nbsp;</span><span>{trainingMetrics.nobs}</span>
                  </div>
                  <div>
                    <span>Cluster:&nbsp;</span><span>{model.cluster_name}</span>
                  </div>
                </div>
              </Cell>
              <Cell>
                {trainingMetrics.MSE ? trainingMetrics.MSE.toFixed(6) : 'N/A'}
              </Cell>
              <Cell>
                {trainingMetrics.logloss ? trainingMetrics.logloss.toFixed(6) : 'N/A'}
              </Cell>
              <Cell>
                {trainingMetrics.r2.toFixed(6)}
              </Cell>
              <Cell>
                <ul className="actions">
                  <li><Link to={'/projects/' + this.props.projectId + '/models/' + model.id}><span><i
                    className="fa fa-eye"></i></span><span>view model details</span></Link></li>
                  <li className="labels"><span><i className="fa fa-tags"></i></span> label as
                        <span className="label-selector">
                          <select name="labelSelect">
                            <option value="prod">test</option>
                            <option value="test">stage</option>
                            <option value="prod">prod</option>
                          </select>
                        </span>
                  </li>
                  <li onClick={this.props.openDeploy.bind(this, model)}><span><i className="fa fa-arrow-up"></i></span>
                    <span>deploy model</span></li>
                </ul>
              </Cell>
            </Row>
          );
        })}
      </Table>
    );
  }
}
