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
import * as moment from 'moment';
import DetailLine from './DetailLine';
import '../styles/modeloverview.scss';

interface Props {
  model: any
}

export default class ModelOverview extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLDivElement> {
    let iconMap = {
      Author: 'fa fa-user',
      Date: 'fa fa-calendar-o',
      Size: 'fa fa-save',
      'Training Time': 'fa fa-hourglass-2',
      'Classification Speed': 'fa fa-clock-o',
      'Model Type': 'fa fa-cube'
    };
    let basics = [
      {
        label: 'Date',
        value: moment.unix(this.props.model.created_at).format('YYYY-MM-DD HH:mm')
      },
      {
        label: 'Model Type',
        value: this.props.model.algorithm
      }
    ];

    let parameters = [
      {
        label: 'Dataset Name',
        value: this.props.model.dataset_name
      },
      {
        label: 'Response Column Name',
        value: this.props.model.response_column_name
      }
    ];
    return (
      <div className="metrics">
        <div className="metrics-summary">
          <div className="metrics-summary--title">Basics</div>
          {basics.map((item, i) => {
            return <DetailLine key={i} icon={iconMap[item.label]} label={item.label} value={item.value}/>;
          })}
        </div>
        <div className="metrics-summary">
          <div className="metrics-summary--title">Model Parameters</div>
          {parameters.map((item, i) => {
            return <DetailLine key={i} label={item.label} value={item.value}/>;
          })}
        </div>
      </div>
    );
  }
}
