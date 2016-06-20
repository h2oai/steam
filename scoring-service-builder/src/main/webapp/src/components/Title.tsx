/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import ContainerFluid from './ContainerFluid';

interface Props {
  className?: string,
  children?: any
}

export default class Title extends React.Component<Props, any> {
  render() {
    return (
      <ContainerFluid className={classNames('title', 'grad', this.props.className)}>
        <h1>Steam</h1>
        <h2>Prediction Service</h2>
      </ContainerFluid>
    );
  }
}