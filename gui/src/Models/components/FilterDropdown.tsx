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
 * Created by justin on 7/28/16.
 */
import * as React from 'react';
import { connect } from 'react-redux';
import * as ReactDOM from 'react-dom';
import * as classNames from 'classnames';
import FilterIcon from './FilterIcon';
import { bindActionCreators } from 'redux';
import { fetchSortCriteria } from '../actions/leaderboard.actions';
import '../styles/filterdropdown.scss';

interface Props {
  target?: Element,
  onFilter: Function,
  sortCriteria: string[]
}

const FILTER_MAP = {
  mean_residual_deviance: 'MRD',
  r_squared: <span>R<sup>2</sup></span>,
  mse: 'MSE',
  logloss: 'LogLoss',
  auc: 'AUC',
  gini: 'Gini'
};

export default class FilterDropdown extends React.Component<Props, any> {
  refs: {
    [key: string]: Element
    filterDropdown: Element
    filterDropdownInvoker: Element
  };

  constructor() {
    super();
    this.state = {
      open: false,
      sortBy: null,
      orderBy: 'asc'
    };
    this.bodyClickHandler = this.bodyClickHandler.bind(this);
  }

  componentWillMount() {
    document.body.addEventListener('click', this.bodyClickHandler);
  }

  componentWillUnmount() {
    document.body.removeEventListener('click', this.bodyClickHandler);
  }

  bodyClickHandler(event) {
    if (!ReactDOM.findDOMNode(this.refs.filterDropdown).contains(event.target) && !ReactDOM.findDOMNode(this.refs.filterDropdownInvoker).contains(event.target)) {
      this.setState({
        open: false
      });
    }
  }

  openDropdown() {
    this.setState({
      open: !this.state.open
    });
  }

  selectSort(selection: string) {
    this.setState({
      sortBy: selection
    });
    this.props.onFilter({
      sortBy: selection,
      orderBy: this.state.orderBy
    });
  }

  selectOrder(selection: string) {
    this.setState({
      orderBy: selection
    });
    this.props.onFilter({
      sortBy: this.state.sortBy,
      orderBy: selection
    });
  }

  render(): React.ReactElement<HTMLDivElement> {
    if (this.props.sortCriteria === null) {
      return <div></div>;
    }
    return (
      <div className="filter-dropdown">
        <button ref="filterDropdownInvoker" className={classNames('filter-dropdown-invoker', {open: this.state.open})}
                onClick={this.openDropdown.bind(this)}><FilterIcon/></button>
        <div ref="filterDropdown" className={classNames('filter-dropdown-menu', {open: this.state.open})}>
          <div className="filter-option">
            <div className="filter-labels">SORT BY</div>
            <ul>
              {this.props.sortCriteria.map((criteria, i) => {
                return <li key={i} onClick={this.selectSort.bind(this, criteria)}
                    className={classNames({selected:this.state.sortBy === criteria})}>{FILTER_MAP[criteria]} {this.state.sortBy === criteria ?
                  <i className="fa fa-check"/> : null}</li>;
              })}
            </ul>
          </div>
          <div className="filter-option">
            <div className="filter-labels">ORDER</div>
            <ul>
              <li onClick={this.selectOrder.bind(this, 'asc')}
                  className={classNames({selected:this.state.orderBy === 'asc'})}>ASC {this.state.orderBy === 'asc' ?
                <i className="fa fa-check"/> : null}</li>
              <li onClick={this.selectOrder.bind(this, 'des')}
                  className={classNames({selected:this.state.orderBy === 'des'})}>DES {this.state.orderBy === 'des' ?
                <i className="fa fa-check"/> : null}</li>
            </ul>
          </div>
        </div>
      </div>
    );
  }
}
