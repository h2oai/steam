/**
 * Created by Jeff Fohl <jfohl@h2o.ai> on 7/12/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import Table from './Table';
import Row from './Row';
import Cell from './Cell';
import { NumericInput } from 'h2oUIKit';
import '../styles/modelparameters.scss';

interface Props {

}

export default class ModelParameters extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLElement> {
    return (
      <div className="model-parameters">
        <Table>
          <Row>
            <Cell><span>Number of Trees</span></Cell>
            <Cell><NumericInput name="numberOfTrees"/></Cell>
            <Cell>Number of trees to train</Cell>
          </Row>
          <Row>
            <Cell><span>Max Depth</span></Cell>
            <Cell><NumericInput name="maxDepth"/></Cell>
            <Cell>Maximum depth for any number in a tree</Cell>
          </Row>
          <Row>
            <Cell><span>Minimum Number of Rows</span></Cell>
            <Cell><NumericInput name="minimumNumberOfRows"/></Cell>
            <Cell>Minimum number of rows in each leaf node</Cell>
          </Row>
          <Row>
            <Cell><span>Learning Rate</span></Cell>
            <Cell><NumericInput name="learningRate"/></Cell>
            <Cell>Learning rate</Cell>
          </Row>
        </Table>
        <button className="link">see full parameters list</button>
      </div>
    );
  }
}
