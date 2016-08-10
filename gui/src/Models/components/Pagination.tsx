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
