/**
 * Created by justin on 6/28/16.
 */
import * as React from 'react';
import * as _ from 'lodash';
import * as classNames from 'classnames';
import DetailLine from './DetailLine';
import RocGraph from '../../Models/components/RocGraph';
import '../styles/goodnessoffit.scss';
import { BRAND_ORANGE, BRAND_BLUE } from '../../App/utils/colors';

interface Props {
  model: any,
  comparisonModel: any,
  modelCategory: string
}

export default class GoodnessOfFit extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLDivElement> {
    let modelMetrics = JSON.parse(this.props.model.metrics);
    let comparisonModelMetrics = this.props.comparisonModel ? JSON.parse(this.props.comparisonModel.metrics) : null;
    let trainingMetrics = _.get(modelMetrics, 'models[0].output.training_metrics', {});
    let comparisonTrainingMetrics = _.get(comparisonModelMetrics, 'models[0].output.training_metrics', {});
    let metrics = {
      mse: {
        label: 'Mean Squared Error',
        value: _.get(this.props.model, 'mse', null) ? _.get(this.props.model, 'mse', null).toFixed(6) : null,
        comparisonValue: _.get(this.props.comparisonModel, 'mse', null) ? _.get(this.props.comparisonModel, 'mse', null).toFixed(6) : null
      },
      logloss: {
        label: 'LogLoss',
        value: _.get(this.props.model, 'logloss', null) ? _.get(this.props.model, 'logloss', null).toFixed(6) : null,
        comparisonValue: _.get(this.props.comparisonModel, 'logloss', null) ? _.get(this.props.comparisonModel, 'logloss', null).toFixed(6) : null
      },
      r_squared: {
        label: <span>R<sup>2</sup></span>,
        value: _.get(this.props.model, 'r_squared', null) ? _.get(this.props.model, 'r_squared', null).toFixed(6) : null,
        comparisonValue: _.get(this.props.comparisonModel, 'r_squared', null) ? _.get(this.props.comparisonModel, 'r_squared', null).toFixed(6) : null
      },
      auc: {
        label: 'AUC',
        value: _.get(this.props.model, 'auc', null) ? _.get(this.props.model, 'auc', null).toFixed(6) : null,
        comparisonValue: _.get(this.props.comparisonModel, 'auc', null) ? _.get(this.props.comparisonModel, 'auc', null).toFixed(6) : null
      },
      gini: {
        label: 'Gini',
        value: _.get(trainingMetrics, 'Gini', null) ? _.get(trainingMetrics, 'Gini', null).toFixed(6) : null,
        comparisonValue: _.get(comparisonTrainingMetrics, 'Gini', null) ? _.get(comparisonTrainingMetrics, 'Gini', null).toFixed(6) : null
      }
    };
    let fpr = _.get(modelMetrics, 'models[0].output.training_metrics.thresholds_and_metric_scores.data[17]', []);
    let tpr = _.get(modelMetrics, 'models[0].output.training_metrics.thresholds_and_metric_scores.data[18]', []);
    let comparisonTpr = _.get(comparisonModelMetrics, 'models[0].output.training_metrics.thresholds_and_metric_scores.data[18]', []);
    let data = [];
    let modelCurveData = {
      name: this.props.model.name,
      values: []
    };
    tpr.map((val, i) => {
      let newEntry: {tpr?: number, fpr?: number, comparisonTpr?: number} = {};
      newEntry.tpr = val;
      newEntry.fpr = fpr[i];
      modelCurveData.values.push(newEntry);
    });
    data.push(modelCurveData);
    if (this.props.comparisonModel) {
      let comparisonCurveData = {
        name: this.props.comparisonModel.name,
        values: []
      };
      comparisonTpr.map((val, i) => {
        let newEntry: {tpr?: number, fpr?: number, comparisonTpr?: number} = {};
        newEntry.tpr = val;
        newEntry.fpr = fpr[i];
        comparisonCurveData.values.push(newEntry);
      });
      data.push(comparisonCurveData);
    }

    let config = {
      margin: {top: 2, right: 2, bottom: 2, left: 2},
      interpolationMode: 'basis',
      height: '100%',
      width: '100%',
      smooth: true,
      fpr: 'fpr',
      tprVariables: [
        {
          name: 'tpr',
          label: 'tpr'
        },
        {
          name: 'comparisonTpr',
          label: 'comparisonTpr'
        }
      ],
      animate: undefined,
      hideAxes: true,
      hideAUCText: true,
      curveColors: [BRAND_BLUE, BRAND_ORANGE]
    };
    return (
      <div className="metrics goodness-of-fit">
        <div className="metrics-summary">
          <div className="metrics-summary--title">Metrics</div>
          {Object.keys(metrics).map((key) => {
            if (metrics[key].value) {
              return <DetailLine className={classNames({compare: metrics[key].comparisonValue})} key={key}
                                 label={metrics[key].label} value={metrics[key].value}
                                 comparisonValue={metrics[key].comparisonValue}/>;
            }
            return null;
          })}
        </div>
        {this.props.modelCategory === 'binomial' ? <div className="roc-chart"><RocGraph data={data} config={config}/></div> : null}
      </div>
    );
  }
}
