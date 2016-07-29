/**
 * Created by justin on 7/12/16.
 */
import * as React from 'react';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import '../styles/packaging.scss';

export default class Packaging extends React.Component<any, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="packaging">
        <h1>PREPROCESSING PACKAGES</h1>
        <div>
          Custom packaging methods for model deployment
        </div>
        <Table>
          <Row header={true}/>
          <Row>
            <Cell className="folder-icon">
              <i className="fa fa-folder"/>
            </Cell>
            <Cell>Spam Detector</Cell>
            <Cell>Mark Landry</Cell>
            <Cell>6 files</Cell>
            <Cell>280KB</Cell>
            <Cell>3 months ago</Cell>
          </Row>
        </Table>
      </div>
    );
  }
}
