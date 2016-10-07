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

import * as React from 'react';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import '../styles/users.scss';
import { saveUpdatedPermissions } from "../actions/users.actions";
import DefaultModal from '../../App/components/DefaultModal';

interface Props {
  requestedChanges: Array<any>,
  saveUpdatedPermissions: Function,
  open: boolean,
  closeHandler: Function
}

export default class RolePermissionsConfirm extends React.Component<Props, any> {

  constructor(props) {
    super(props);
    this.state = {
      requestedChanges: this.props.requestedChanges
    };
  }

  componentWillReceiveProps(nextProps) {
    this.setState({
      requestedChanges: nextProps.requestedChanges
    });
  }

  cancelIndividualChange = (index) => {
    let clone = this.state.requestedChanges.slice(0);
    clone.splice(index, 1);

    this.setState({
      requestedChanges: clone
    });
  };

  confirmClicked = (index: number) => {
    let clone = this.state.requestedChanges.slice(0);
    clone[index].confirmed = true;

    this.setState({
      requestedChanges: clone
    });
  };

  render(): React.ReactElement<DefaultModal> {
    let saveChangesEnabled = true;
    for (let requestedChange of this.state.requestedChanges) {
      if (!requestedChange.confirmed) {
        saveChangesEnabled = false;
      }
    }

    return (
      <DefaultModal open={this.props.open} closeHandler={this.props.closeHandler}>
        <h1>CONFIRMING PERMISSION CHANGES</h1>
        <p>You are making the following changes</p>
        <Table>
          <Row header={true}>
            <Cell>CHANGE</Cell>
            <Cell>CONFIRM</Cell>
          </Row>
          { this.state.requestedChanges ? this.state.requestedChanges.map((requestedChange, index, array) => {
            return <Row key={index}>
              <Cell>{requestedChange.userDescription } &nbsp; {requestedChange.newFlag ? <span>gains</span> : <span>loses</span>} &nbsp; {requestedChange.description}</Cell>
              <Cell>{requestedChange.confirmed ? <span className="green"><i className="fa fa-check"></i> Confirmed</span> : <span className="button-primary" onClick={(e) => this.confirmClicked(index)}>Confirm</span>} &nbsp; <span onClick={() => this.cancelIndividualChange(index)} className="button">Cancel Change</span></Cell>
            </Row>;
          }) : null }
        </Table>
        <br />
        {saveChangesEnabled ? <div className="button-primary" onClick={() => this.props.saveUpdatedPermissions(this.state.requestedChanges)}>Save Changes</div>
          : <div className="button-primary disabled">Save Changes</div> }
         &nbsp;
        <div className="button-secondary" onClick={this.props.closeHandler}>Cancel</div>
      </DefaultModal>
    );
  }
}
