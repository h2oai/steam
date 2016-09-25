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
 * Created by justin on 7/29/16.
 */

import * as React from 'react';
import * as moment from 'moment';
import FilterDropdown from '../../Models/components/FilterDropdown';
import DefaultModal from '../../App/components/DefaultModal';
import PageHeader from '../../Projects/components/PageHeader';
import Pagination from '../../Models/components/Pagination';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import '../styles/modelselectionmodal.scss';
import { MAX_ITEMS } from '../../Models/actions/leaderboard.actions';

interface Props {
  open: boolean,
  models: any,
  onSelectModel: Function,
  onCancel: Function,
  onFilter: Function,
  sortCriteria: string[],
  count: number
}

export default class ModelSelectionModal extends React.Component<Props, any> {
  refs: {
    [key: string]: Element
    filterModels: HTMLInputElement
  };

  constructor() {
    super();
    this.state = {
      currentPage: 0,
      filters: {
        sortBy: '',
        orderBy: 'asc'
      }
    };
  }

  onFilter(filters) {
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
                    {model.mse ? model.mse.toFixed(6) : 'N/A'}
                  </Cell>
                  <Cell>
                    {model.auc ? model.auc.toFixed(6) : 'N/A'}
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
          <Pagination items={this.props.models} onPageForward={this.onPageForward.bind(this)}
                      onPageBack={this.onPageBack.bind(this)} currentPage={this.state.currentPage} count={this.props.count}/>
          <button className="default" onClick={this.props.onCancel.bind(this)}>Cancel</button>
        </footer>
      </DefaultModal>
    );
  }
}
