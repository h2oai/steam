/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import ContainerFluid from '../components/ContainerFluid';
import Section from '../components/Section';
import Title from '../components/Title';

export default class App extends React.Component<any, any> {
  render() {
    return (
      <ContainerFluid className="body-container">
        <Title></Title>
        <ContainerFluid>
          <Section className={classNames('col-md-12', 'col-sm-12')}>
            <p>
              Select input parameters, OR enter your own custom query string to predict2
            </p>
          </Section>
          <Section id="inputParams" className={classNames('col-md-12', 'col-sm-12')}>
            <header><span className={classNames('glyphicon', 'glyphicon-log-in')} aria-hidden="true"></span>Model input parameters</header>
          </Section>
        </ContainerFluid>
      </ContainerFluid>
    );
  }
}