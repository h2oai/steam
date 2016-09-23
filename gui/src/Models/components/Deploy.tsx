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
  onDeploy: Function,
  packages: string[]
}

export default class Deploy extends React.Component<Props, any> {
  refs: {
    [key: string]: Element
    serviceName: HTMLInputElement,
    packageName: HTMLSelectElement
  };
  deploy() {
    this.props.onDeploy(this.props.model, _.get(this.refs.serviceName, 'value', ''), _.get(this.refs.packageName, 'value', ''));
  }
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
                  <label className="muted">Preprocessing Script</label>
                  <select ref="packageName">
                    <option value="">None (Default)</option>
                    {this.props.packages.map((packageName, i) => {
                      return <option key={i} value={packageName}>{packageName}</option>;
                    })}
                  </select>
                </Cell>
              </Row>
              <Row>
                <Cell/>
                <Cell>
                  <button type="button" className="default deploy-button" onClick={this.deploy.bind(this)}>
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
