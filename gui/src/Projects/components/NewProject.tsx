/**
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import PageHeader from '../components/PageHeader';
import Table from '../components/Table';
import Row from '../components/Row';
import Cell from '../components/Cell';
import '../styles/newproject.scss';

interface Props {
  children?: React.ReactChildren
}

export default class NewProject extends React.Component<Props, any> {
  render() {
    return (
      <div className="new-project">
        <PageHeader>New Project</PageHeader>
        <form>
          <label>Give your project a name</label>
          <input type="text" placeholder="Name"/>
        </form>
        <form>
          <label>Upload a CSV file</label>
          <span>We will automatically infer a data dictionary based on the data</span>
          <input type="file"/>
          <button type="button">Upload</button>
        </form>
        <form>
          <label>Verify Data Dictionary</label>
          <span>Start building models by verifying the import and selecting a response column</span>
        </form>
        <div>
          <Table>
            <Row header={true}>
              <Cell>COLUMN NAME</Cell>
              <Cell>DATA TYPE</Cell>
              <Cell>DISTRIBUTION</Cell>
              <Cell>STATISTICS</Cell>
              <Cell>ISSUES</Cell>
            </Row>
            <Row>
              <Cell>setosa_length</Cell>
              <Cell>numeric</Cell>
              <Cell/>
              <Cell>
                <div className="statistics">
                  <div className="headings">
                    <div>min</div>
                    <div>max</div>
                    <div>mean</div>
                  </div>
                  <div className="values">
                    <div>0.4</div>
                    <div>5.1</div>
                    <div>3.4</div>
                  </div>
                </div>
              </Cell>
              <Cell/>
            </Row>
            <Row>
              <Cell>
                class
              </Cell>
              <Cell>
                categorical
              </Cell>
              <Cell/>
              <Cell>
                <div className="statistics">
                  <div className="headings">
                    <div>classes</div>
                  </div>
                  <div className="values">
                    <div>3</div>
                  </div>
                </div>
              </Cell>
              <Cell/>
            </Row>
          </Table>
          <button type="button" className="save-dictionary">
            Save Dictionary
          </button>
        </div>
        <div>
          <form>
            <label>
              Build Initial Model
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
                </div>
              </Cell>
            </Row>
          </Table>
          <button type="button" className="train-models">Train Models</button>
        </div>
      </div>
    );
  }
}