/**
 * Created by justin on 6/28/16.
 */

import * as React from 'react';
import '../styles/detailline.scss';

interface Props {
  icon?: string,
  label: string,
  value: string | number
}

export default class DetailLine extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="details">
        <div className="details--label">
          {this.props.icon ? <i className={this.props.icon}></i> : null}{this.props.label}
        </div>
        <div className="details--line">
        </div>
        <div className="details--value">
          {this.props.value}
        </div>
      </div>  
    );
  }
}