/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import * as classNames from 'classnames';
import ContainerFluid from '../components/ContainerFluid';
import Section from '../components/Section';
import Title from '../components/Title';
import ModelForm from '../components/ModelForm';
import { fetchStatistics } from '../reducers/statisticsReducer';
import { fetchModel } from "../reducers/modelReducer";


interface Props {
  statistics?: any,
  model?: any
}

interface DispatchProps {
  fetchStatistics: Function,
  fetchModel: Function
}

export class App extends React.Component<Props & DispatchProps, any> {
  componentWillMount() {
    this.props.fetchStatistics();
    this.props.fetchModel();
  }

  render() {
    if (!this.props.model) {
      return <div></div>;
    }
    return (
      <ContainerFluid className="body-container">
        <Title></Title>
        <ContainerFluid>
          <Section className={classNames('col-md-12', 'col-sm-12')}>
            <p>
              Select input parameters, OR enter your own custom query string to predict
            </p>
          </Section>
          <Section id="inputParams" className={classNames('col-md-6', 'col-sm-12')}>
            <header><span className={classNames('glyphicon', 'glyphicon-log-in')} aria-hidden="true"></span>Model input parameters</header>
            <ModelForm model={this.props.model}/>
          </Section>
        </ContainerFluid>
      </ContainerFluid>
    );
  }
}

function mapStateToProps(state) {
  return {
    statistics: state.statistics.data,
    model: state.model.data
  }
}

function mapDispatchToProps(dispatch) {
  return {
    fetchStatistics: bindActionCreators(fetchStatistics, dispatch),
    fetchModel: bindActionCreators(fetchModel, dispatch)
  }
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(App)