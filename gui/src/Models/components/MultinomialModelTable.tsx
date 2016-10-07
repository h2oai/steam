/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
import ModelLabelSelect from './ModelLabelSelect';
import { Link } from 'react-router';
import { Label } from '../../Proxy/Proxy';

interface Props {
  onFilter: Function,
  sortCriteria: string[],
  items: any,
  projectId: number,
  openDeploy: Function,
  deleteModel: Function,
  fetchLeaderboard: Function,
  onChangeHandler: Function,
  labels: {
    [projectId: number]: Label[]
  }
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
          let modelMetrics = JSON.parse(model.json_metrics);
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
              <Cell name="mse">
                {trainingMetrics.MSE ? trainingMetrics.MSE.toFixed(6) : 'N/A'}
              </Cell>
              <Cell name="logloss">
                {trainingMetrics.logloss ? trainingMetrics.logloss.toFixed(6) : 'N/A'}
              </Cell>
              <Cell name="r2">
                {trainingMetrics.r2 ? trainingMetrics.r2.toFixed(6) : 'N/A'}
              </Cell>
              <Cell>
                <ul className="actions">
                  <li><Link to={'/projects/' + this.props.projectId + '/models/' + model.id}><span><i
                    className="fa fa-eye"></i></span><span>view model details</span></Link></li>
                  <li className="labels"><span><i className="fa fa-tags"></i></span> label as
                        <span className="label-selector">
                          <ModelLabelSelect projectId={this.props.projectId} modelId={model.id}
                                            labels={this.props.labels} onChangeHandler={this.props.onChangeHandler}/>
                        </span>
                  </li>
                  <li onClick={this.props.openDeploy.bind(this, model)}><span><i className="fa fa-arrow-up"></i></span>
                    <span>deploy model</span></li>
                  <span><i className="fa fa-trash" onClick={() => this.props.deleteModel(model.id, this.props.fetchLeaderboard)}></i> delete model</span>
                </ul>
              </Cell>
            </Row>
          );
        })}
      </Table>
    );
  }
}
