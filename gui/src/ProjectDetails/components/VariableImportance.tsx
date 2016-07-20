/**
 * Created by justin on 6/28/16.
 */
import * as React from 'react';
import DetailLine from './DetailLine';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import '../styles/variableimportance.scss';
import * as d3 from 'd3';

interface Props {
}

export default class VariableImportance extends React.Component<Props, any> {

  columns = [];

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
        type: 'numeric',
        importance: 0.21
      }            
    ];
    this.widthScale = d3.scaleLinear()
      .domain([0, 1])
      .range([0, this.rowWidth]);
  }
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="metrics">
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
              <Row>
              <Cell></Cell>
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
              <Cell>
              </Cell>
              <Cell className="graph">
              </Cell>
              <Cell className="graph">
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