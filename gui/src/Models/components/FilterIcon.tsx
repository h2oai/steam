/**
 * Created by Jeff Fohl <jfohl@h2o.ai> on 6/29/16.
 */
import * as React from 'react';
import '../styles/filterdropdown.scss';

export default class FilterIcon extends React.Component<any, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="filter-icon">
        <span></span>
        <span></span>
        <span></span>
      </div>
    );
  }
}
