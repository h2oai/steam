/**
 * Created by justin on 6/28/16.
 */
import * as React from 'react';
import * as _ from 'lodash';
import * as classNames from 'classnames';
import DetailLine from './DetailLine';
import RocGraph from '../../Models/components/RocGraph';
import '../styles/goodnessoffit.scss';

interface Props {
  model: any,
  comparisonModel: any
}

export default class GoodnessOfFit extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLDivElement> {
    let metrics = {
      mse: {
        label: 'Mean Squared Error',
        value: _.get(this.props.model, 'mse', null),
        comparisonValue: _.get(this.props.comparisonModel, 'mse', null)
      },
      logloss: {
        label: 'LogLoss',
        value: _.get(this.props.model, 'logloss', null),
        comparisonValue: _.get(this.props.comparisonModel, 'logloss', null)
      },
      r_squared: {
        label: <span>R<sup>2</sup></span>,
        value: _.get(this.props.model, 'r_squared', null),
        comparisonValue: _.get(this.props.comparisonModel, 'r_squared', null)
      },

    };
    let modelMetrics = JSON.parse(this.props.model.metrics);
    let fpr = _.get(modelMetrics, 'models[0].output.training_metrics.thresholds_and_metric_scores.data[17]', []);
    let tpr = _.get(modelMetrics, 'models[0].output.training_metrics.thresholds_and_metric_scores.data[18]', []);
    let data = [];
    tpr.map((val, i) => {
      data.push({
        tpr: val,
        fpr: fpr[i]
      });
    });
    return (
      <div className="metrics">
        <div className="metrics-summary">
          <div className="metrics-summary--title">Metrics</div>
          {Object.keys(metrics).map((key) => {
            if (metrics[key].value) {
              return <DetailLine className={classNames({compare: metrics[key].comparisonValue})} key={key} label={metrics[key].label} value={metrics[key].value}
                                 comparisonValue={metrics[key].comparisonValue}/>;
            }
            return null;
          })}
        </div>
        <div className="roc-chart">
          <RocGraph data={data}/>
        </div>
      </div>
    );
  }
}
