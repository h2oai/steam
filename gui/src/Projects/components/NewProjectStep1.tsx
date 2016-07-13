/**
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import { Link } from 'react-router';
import PageHeader from '../components/PageHeader';
import Table from '../components/Table';
import Row from '../components/Row';
import Cell from '../components/Cell';
import '../styles/newproject.scss';

interface Props {
  children?: React.ReactChildren
}

export default class NewProject extends React.Component<Props, any> {
  constructor() {
    super();
    this.state = {
      shapeIsDisabled: true
    }
  }
  render() {
    return (
      <div className="new-project">
        <PageHeader>New Project</PageHeader>
        <form>
          <label>Give your project a name</label>
          <input type="text" placeholder="Name"/>
        </form>
        <form className={classNames({disabled: this.state.shapeIsDisabled})}>
          <label>Upload a CSV file</label>
          <span>We will automatically infer a data dictionary based on the data</span>
          <input type="file"/>
          <button type="button" className="default">Upload</button>
        </form>
        <form>
          <label>Verify Data Shape</label>
          <span>Your models will be more accurate if H2O has an accurate understanding of the column types in your data.</span>
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
        </div>
        <div>
          <form>
            <label>Select Response Column</label>
            <span>Identify the column with the value you want to predict.</span>
            <select>
              <option>Test</option>
            </select>
          </form>
        </div>
        <Link to="/projects/new/2" className="default">Next: Train Initial Models</Link>
      </div>
    );
  }
}
