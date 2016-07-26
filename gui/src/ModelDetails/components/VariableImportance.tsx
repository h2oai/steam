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
import { getOrdinal } from '../../App/utils/getOrdinal';
import '../styles/variableimportance.scss';

// sample data
import { responseDistributionSubset } from '../tests/data/responseDistributionSubset';

interface Props {
  columns?: any[]
  data?: any
  rowHeight?: number
  rowWidth?: number
}

interface State {

}

export default class VariableImportance extends React.Component<Props, State> {

  widthScale: any = () => {};

  componentWillMount() {
    this.widthScale = d3.scaleLinear()
      .domain([0, 1])
      .range([0, this.props.rowWidth]);
  }

  static defaultProps: Props = {
    rowHeight: 70,
    rowWidth: 210,
    columns: [
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
    ],
    data : {
      responseDistributionSubset
    }
  };

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
              <div>
                RESPONSE DISTRIBUTION
              </div>
              <div className="legend">
                <svg width={this.props.rowWidth} height={0.37 * this.props.rowHeight}>
                  <g transform="translate(0, 0)">
                    <rect
                        x="0"
                        y="10"
                        width="20"
                        height="20"
                        rx="0"
                        ry="0"
                        className="symbol yes"
                    />
                    <text
                       x="30"
                       y="10"
                       dy="12"
                       className="legendText"
                    >
                       Yes
                     </text>
                  </g>
                  <g transform="translate(85, 0)">
                    <rect
                      x="0"
                      y="10"
                      width="20"
                      height="20"
                      rx="0"
                      ry="0"
                      className="symbol no"
                    />
                    <text
                       x="30"
                       y="10"
                       dy="12"
                       className="legendText"
                    >
                       No
                     </text>
                  </g>
                </svg>
              </div>
            </Cell>
            <Cell className="graph">
              NOTES
            </Cell>
          </Row>
          {this.props.columns.map((item, i) => {
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
                  <svg width={this.props.rowWidth} height={this.props.rowHeight}>
                    <rect
                      x="0"
                      y="0"
                      width={this.widthScale(item.importance)}
                      height={this.props.rowHeight}
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
                <GroupedBarChart data={this.props.data['responseDistributionSubset'][i]['responseCounts']}/>
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
