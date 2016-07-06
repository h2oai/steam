/**
 * Created by justin on 6/27/16.
 */
import * as React from 'react';  
import '../styles/pagination.scss';

interface Props {
  
}

export class Pagination extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="pagination-container">
        <div>
          <span><i className="fa fa-chevron-left"></i></span>
          Page 1 of 14
          <span><i className="fa fa-chevron-right"></i></span>
        </div>
      </div>
    );
  }
}