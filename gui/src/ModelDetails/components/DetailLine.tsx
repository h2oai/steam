/**
 * Created by justin on 6/28/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import '../styles/detailline.scss';

interface Props {
  icon?: string,
  label: string | React.ReactElement<Element>,
  value: any,
  comparisonValue?: any,
  className?: any
}

export default class DetailLine extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLDivElement> {
    console.log(this.props);
    return (
      <div className={classNames('details', this.props.className)}>
        <div className="details--label">
          {this.props.icon ? <i className={this.props.icon}></i> : null}{this.props.label}
        </div>
        <div className="details--line">
        </div>
        <div className="details--value">
          <span>{this.props.value}</span><span>{this.props.comparisonValue}</span>
        </div>
      </div>
    );
  }
}
