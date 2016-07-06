/**
 * Created by justin on 6/27/16.
 */
import * as React from 'react';  
import '../styles/pagination.scss';

interface Props {
  items: any
}

export default class Pagination extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="pagination-container">
        <span><i className="fa fa-caret-left"></i></span>
        <span className="page-info">1 - 5 of {this.props.items.length} models</span>
        <span><i className="fa fa-caret-right"></i></span>
      </div>
    );
  }
}