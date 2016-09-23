/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

/**
 * Created by justin on 7/10/16.
 */
import * as React from 'react';
import { Link } from 'react-router';
import Table from '../components/Table';
import Row from '../components/Row';
import Cell from '../components/Cell';
import '../styles/newproject.scss';

export default class NewProjectStep2 extends React.Component<any, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="new-project">
        <form>
          <label>
            Train Initial Model
          </label>
            <span>
              Create a training frame, test frame, and start building models.
            </span>
        </form>
        <Table className="build-model">
          <Row header={true}/>
          <Row>
            <Cell>SPLIT DATAFRAME</Cell>
            <Cell>
              <div className="dataframe-range">
                <div>
                  <input type="range"/>
                </div>
                <div className="dataframe-range-labels">
                  <div>
                    Train: 75%
                  </div>
                  <div>
                    Test: 25%
                  </div>
                </div>
              </div>
            </Cell>
          </Row>
          <Row>
            <Cell>DEFAULT MODELS</Cell>
            <Cell>
              <div className="mode-checkboxes">
                <span><input type="checkbox"/>&nbsp;Generalized Linear Model</span>
                <span><input type="checkbox"/>&nbsp;Gradient Boosting Machine</span>
                <span><input type="checkbox"/>&nbsp;Random Forest</span>
                <span><input type="checkbox"/>&nbsp;Deep Learning</span>
                <span><input type="checkbox"/>&nbsp;Naive Bayes</span>
              </div>
            </Cell>
          </Row>
          <Row>
            <Cell>PICK A TRAINING CLUSTER</Cell>
            <Cell>
              <select>
                <option>Test</option>
              </select>
            </Cell>
          </Row>
        </Table>
        <Link to="/projects/new/3" className="default">Train Models</Link>
      </div>
    );
  }
}
