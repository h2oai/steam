/**
 * Created by justin on 6/27/16.
 */
import * as React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchModelOverview } from '../actions/model.overview.action';
import DetailLine from './DetailLine';
import '../styles/modeloverview.scss';

interface Props {
  modelOverview: any
}

interface DispatchProps {
  fetchModelOverview: Function
}

export class ModelOverview extends React.Component<Props & DispatchProps, any> {
  componentWillMount() {
    this.props.fetchModelOverview();
  }

  render(): React.ReactElement<HTMLDivElement> {
    console.log(this.props);
    if (!this.props.modelOverview) {
      return <div></div>;
    }
    let iconMap = {
      Author: 'fa fa-user',
      Date: 'fa fa-calendar-o',
      Size: 'fa fa-save',
      'Training Time': 'fa fa-hourglass-2',
      'Classification Speed': 'fa fa-clock-o',
      'Model Type': 'fa fa-cube'
    };

    return (
      <div className="metrics">
        <div className="metrics-summary">
          <span className="metrics-summary--title">Basics</span>
          {this.props.modelOverview.basics.map((item, i) => {
            return <DetailLine key={i} icon={iconMap[item.label]} label={item.label} value={item.value}/>;
          })}
        </div>
        <div className="metrics-summary">
          <span className="metrics-summary--title">Parameters</span>
          {this.props.modelOverview.parameters.map((item, i) => {
            return <DetailLine key={i} label={item.label} value={item.value}/>
          })}
        </div>
      </div>
    );
  }
}

function mapStateToProps(state: Props): Props {
  return {
    modelOverview: state.modelOverview
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchModelOverview: bindActionCreators(fetchModelOverview, dispatch)
  }
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ModelOverview);