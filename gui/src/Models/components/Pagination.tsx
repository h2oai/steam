/**
 * Created by justin on 6/27/16.
 */
import * as React from 'react';
import '../styles/pagination.scss';
import { MAX_ITEMS } from '../actions/leaderboard.actions';

interface Props {
  items: any,
  onPageBack: Function,
  onPageForward: Function
}

export default class Pagination extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="pagination-container">
        <span onClick={this.props.onPageBack.bind(this)}><i className="fa fa-caret-left"></i></span>
        <span className="page-info">1 - {this.props.items.length < MAX_ITEMS ? this.props.items.length : MAX_ITEMS} of {this.props.items.length} models</span>
        <span onClick={this.props.onPageForward.bind(this)}><i className="fa fa-caret-right"></i></span>
      </div>
    );
  }
}
