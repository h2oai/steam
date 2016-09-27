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
 * Created by justin on 6/27/16.
 */
import * as React from 'react';
import * as classNames from 'classnames';
import '../styles/pagination.scss';
import { MAX_ITEMS } from '../actions/leaderboard.actions';

interface Props {
  items: any,
  onPageBack: Function,
  onPageForward: Function,
  currentPage: number,
  count: number
}

export default class Pagination extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="pagination-container">
        <span onClick={this.props.onPageBack.bind(this)}><i className={classNames('fa fa-caret-left', {disabled: this.props.currentPage === 0})}></i></span>
        <span className="page-info">{((this.props.currentPage + 1) * MAX_ITEMS) - (MAX_ITEMS - 1)} - {(this.props.currentPage + 1) * MAX_ITEMS < this.props.count ? (this.props.currentPage + 1) * MAX_ITEMS : this.props.count} of {this.props.count} models</span>
        <span onClick={this.props.onPageForward.bind(this)}><i className={classNames('fa fa-caret-right', {disabled: (this.props.currentPage + 1) * MAX_ITEMS >= this.props.count})}></i></span>
      </div>
    );
  }
}
