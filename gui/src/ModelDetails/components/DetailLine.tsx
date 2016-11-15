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
    return (
      <div className={classNames('details', this.props.className, this.props.label.toString())}>
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
