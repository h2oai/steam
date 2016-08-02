/**
 * Created by justin on 7/29/16.
 */

import * as React from 'react';
import * as moment from 'moment';
import FilterDropdown from '../../Models/components/FilterDropdown';
import DefaultModal from '../../App/components/DefaultModal';
import PageHeader from '../../Projects/components/PageHeader';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { Model, RegressionModel, MultinomialModel, BinomialModel } from '../../Proxy/Proxy';
import { fetchModelsFromProject } from '../../Projects/actions/projects.actions';
import '../styles/modelselectionmodal.scss';
import { fetchLeaderboard } from '../../Models/actions/leaderboard.actions';

interface Props {
  open: boolean,
  models: any,
  projectId: string,
  onSelectModel: Function,
  onCancel: Function,
  project: any,
  onFilter: Function,
  sortCriteria: string[]
}

export default class ModelSelectionModal extends React.Component<Props, any> {
  refs: {
    [key: string]: Element
    filterModels: HTMLInputElement
  };
  onFilter(filters) {
    this.props.onFilter(filters, this.refs.filterModels.value);
  }

  render(): React.ReactElement<DefaultModal> {
    return (
      <DefaultModal className="model-comparison-modal" open={this.props.open}>
        <PageHeader>
          CHOOSE MODEL TO COMPARE
        </PageHeader>
        <div>
          <div>Filter models by name</div>
          <input ref="filterModels" type="text" placeholder="filter models" onChange={this.onFilter.bind(this)}/>
          <Table>
            <Row header={true}>
              <Cell>
                <FilterDropdown onFilter={this.onFilter.bind(this)} sortCriteria={this.props.sortCriteria}/>
              </Cell>
              <Cell>
                MODEL
              </Cell>
              <Cell>
                DATE
              </Cell>
              <Cell>
                MSE
              </Cell>
              <Cell>
                AUC
              </Cell>
              <Cell>
                F1
              </Cell>
              <Cell/>
            </Row>
            {this.props.models.map((model, i: number) => {
              return (
                <Row key={i}>
                  <Cell/>
                  <Cell>
                    {model.name}
                  </Cell>
                  <Cell>
                    {moment.unix(model.created_at).format('YYYY-MM-DD HH:mm')}
                  </Cell>
                  <Cell>
                    {model.mse.toFixed(6)}
                  </Cell>
                  <Cell>
                    {model.r_squared.toFixed(6)}
                  </Cell>
                  <Cell>
                    {model.mean_residual_deviance}
                  </Cell>
                  <Cell>
                    <button className="default" onClick={this.props.onSelectModel.bind(this, model)}>Select</button>
                  </Cell>
                </Row>
              );
            })}
          </Table>
        </div>
        <footer>
          <button className="default" onClick={this.props.onCancel.bind(this)}>Cancel</button>
        </footer>
      </DefaultModal>
    );
  }
}
