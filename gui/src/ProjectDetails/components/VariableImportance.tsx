/**
 * Created by justin on 6/28/16.
 */
import * as React from 'react';
import DetailLine from './DetailLine';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import GroupedBarChart from './GroupedBarChart';
import * as d3 from 'd3';
import { getOrdinal } from '../../utils/utils';
import '../styles/variableimportance.scss';

// sample data
import { responseDistributionSubset } from '../tests/data/responseDistributionSubset';

interface Props {
}

export default class VariableImportance extends React.Component<Props, any> {

  columns = [];
  sampleData = {};

  constructor() {
    super();
    this.rowHeight = 70;
    this.rowWidth = 210;
    this.columns = [
      {
        name: 'tenure',
        type: 'numeric',
        importance: 0.97
      },
      {
        name: 'gender',
        type: 'enum(2)',
        importance: 0.88
      },
      {
        name: 'PhoneService',
        type: 'enum(3)',
        importance: 0.72
      },
      {
        name: 'OnlineSecurity',
        type: 'enum(3)',
        importance: 0.32
      },
      {
        name: 'Dependents',
        type: 'categorical',
        importance: 0.21
      }
    ];
    this.widthScale = d3.scaleLinear()
      .domain([0, 1])
      .range([0, this.rowWidth]);
    this.sampleData = {
      responseDistributionSubset
    }
  }
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="variable-importance metrics">
        <Table>
          <Row header={true}>
            <Cell>
              <i className="fa fa-caret-down"/>
            </Cell>
            <Cell>
              COLUMN
            </Cell>
            <Cell className="graph">
              IMPORTANCE
            </Cell>
            <Cell className="graph">
              PARTIAL DEPENDENCY
            </Cell>
            <Cell className="graph">
              RESPONSE DISTRIBUTION
            </Cell>
            <Cell className="graph">
              NOTES
            </Cell>
          </Row>
          {this.columns.map((item, i) => {
            return (
              <Row key={i}>
              <Cell>
                {(i + 1) + getOrdinal(i + 1)}
              </Cell>
              <Cell>
                <div className="variableImportance">
                  <div className="columnName">
                    {item.name}
                  </div>
                  <div className="detail">
                    {item.type}
                  </div>
                </div>
              </Cell>
              <Cell>
                <div>
                  <svg width={this.rowWidth} height={this.rowHeight}>
                    <rect
                      x="0"
                      y="0"
                      width={this.widthScale(item.importance)}
                      height={this.rowHeight}
                      rx="0"
                      ry="0"
                      className="bar"
                    />
                  </svg>
                </div>
                <div className="detail">
                  {item.importance}
                </div>
              </Cell>
              <Cell className="graph">
              </Cell>
              <Cell>
                <GroupedBarChart data={this.sampleData['responseDistributionSubset'][i]['responseCounts']}/>
              </Cell>
              <Cell></Cell>
            </Row>
            );
          })}
        </Table>
      </div>
    );
  }
}
