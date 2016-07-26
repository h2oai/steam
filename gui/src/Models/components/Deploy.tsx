/**
 * Created by justin on 6/30/16.
 */
import * as React from 'react';
import PageHeader from '../../Projects/components/PageHeader';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import DefaultModal from '../../App/components/DefaultModal';
import '../styles/deploy.scss';

interface Props {
  open: boolean
  closeHandler: Function
}

export default class Deploy extends React.Component<Props, any> {
  constructor() {
    super();
    this.onClick = this.onClick.bind(this);
  }

  onClick() {
    this.props.closeHandler();
  }

  render(): React.ReactElement<DefaultModal> {
    return (
      <DefaultModal className="deploy-modal" open={this.props.open}>
        <div>
          <PageHeader>DEPLOY DRF-1069085</PageHeader>
          <section>
            <Table className="deployment-table">
              <Row>
                <Cell>
                  SELECT PACKAGER
                </Cell>
                <Cell>
                  <select>
                    <option>
                      Storm Bolt Packager /deployment/stormbolt-packager.py
                    </option>
                  </select>
                </Cell>
              </Row>
              <Row>
                <Cell>
                  DEPLOYMENT MODE
                </Cell>
                <Cell>
                  <div className="radio-container">
                    <div className="radio-group">
                      <input type="radio" name="deploy-mode"/>
                      <span>As a Steam Service</span>
                    </div>
                    <div className="radio-group">
                      <input type="radio" name="deploy-mode"/>
                      <span>Download for Manual Deployment</span>
                    </div>
                  </div>
                </Cell>
              </Row>
              <Row>
                <Cell>
                  CONFIGURE SERVICE
                </Cell>
                <Cell>
                  <div>
                    <div className="caption">
                      Select existing service to deploy model into:
                    </div>
                    <select className="">
                      <option>
                        localhost:54321
                      </option>
                    </select>
                  </div>
                </Cell>
              </Row>
              <Row>
                <Cell/>
                <Cell>
                  <button type="button" className="default deploy-button">Deploy DRF-1069085 into localhost:54321</button>
                  <span><a href="javascript:void(0)">or, create a new service with this model.</a></span>
                </Cell>
              </Row>
            </Table>
          </section>
        </div>
      </DefaultModal>
    );
  }
}
