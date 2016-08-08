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
