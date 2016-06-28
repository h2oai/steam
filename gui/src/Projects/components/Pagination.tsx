/**
 * Created by justin on 6/27/16.
 */
import * as React from 'react';  
import '../styles/pagination.scss';
import 'bootstrap/dist/css/bootstrap.css';

interface Props {
  
}

export class Pagination extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="pagination-container">
        <div>
          <span><i className="glyphicon glyphicon-chevron-left"></i></span>
          Page 1 of 14
          <span><i className="glyphicon glyphicon-chevron-right"></i></span>
        </div>
      </div>
    );
  }
}