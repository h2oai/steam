/**
 * Created by justin on 6/30/16.
 */
import * as React from 'react';
import * as _ from 'lodash';
import PageHeader from '../../Projects/components/PageHeader';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import DefaultModal from '../../App/components/DefaultModal';
import '../styles/deploy.scss';
import { Model } from '../../Proxy/Proxy';

interface Props {
  open: boolean,
  model: Model,
  onCancel: Function,
  onDeploy: Function
}

export default class Deploy extends React.Component<Props, any> {
  refs: {
    [key: string]: Element
    serviceName: HTMLInputElement
  };
  render(): React.ReactElement<DefaultModal> {
    return (
      <DefaultModal className="deploy-modal" open={this.props.open}>
        <div>
          <PageHeader>DEPLOY {_.get(this.props.model, 'name', '')}</PageHeader>
          <section>
            <Table className="deployment-table">
              <Row>
                <Cell>
                  CONFIGURE SERVICE
                </Cell>
                <Cell>
                  <span>Steam automatically selects a port that's not in use based on the port range set by your admin.</span>
                  <label className="muted">Service name</label>
                  <input ref="serviceName" type="text"/>
                </Cell>
              </Row>
              <Row>
                <Cell/>
                <Cell>
                  <button type="button" className="default deploy-button" onClick={this.props.onDeploy.bind(this, this.props.model, _.get(this.refs.serviceName, 'value', ''))}>
                    Deploy
                  </button>
                  <button type="button" className="default invert" onClick={this.props.onCancel.bind(this)}>Cancel
                  </button>
                </Cell>
              </Row>
            </Table>
          </section>
        </div>
      </DefaultModal>
    );
  }
}
