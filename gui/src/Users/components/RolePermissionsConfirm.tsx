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
import InputFeedback from '../../App/components/InputFeedback';
import { FeedbackType } from '../../App/components/InputFeedback';

interface Props {
  requestedChanges: Array<any>,
  saveUpdatedPermissions: Function,
  open: boolean,
  closeHandler: Function,
  updates: Array<any>
}

export default class RolePermissionsConfirm extends React.Component<Props, any> {

  constructor(props) {
    super(props);
    this.state = {
      requestedChanges: this.props.requestedChanges,
      saveClicked: false
    };
  }

  componentWillReceiveProps(nextProps) {
    this.setState({
      requestedChanges: nextProps.requestedChanges,
      saveClicked: false
    });
  }

  cancelIndividualChange = (index) => {
    let clone = this.state.requestedChanges.slice(0);
    clone.splice(index, 1);

    this.setState({
      requestedChanges: clone,
      saveClicked: false
    });

    if (clone.length === 0) {
      this.props.closeHandler();
    }
  };

  confirmClicked = (index: number) => {
    let clone = this.state.requestedChanges.slice(0);
    clone[index].confirmed = true;

    this.setState({
      requestedChanges: clone,
      saveClicked: false
    });
  };

  onSaveChangesClicked = () => {
    this.props.saveUpdatedPermissions(this.state.requestedChanges);
    this.setState({
      requestedChanges: this.state.requestedChanges.slice(0),
      saveClicked: true
    });
  }

  render(): React.ReactElement<DefaultModal> {
    let saveChangesEnabled = true;
    for (let requestedChange of this.state.requestedChanges) {
      if (!requestedChange.confirmed) {
        saveChangesEnabled = false;
      }
    }

    var results = [];
    this.state.requestedChanges.map((requestedChange, index, array) => {
      let matchUpdate = null;

      for (let update of this.props.updates) {
        console.log("checking for update");
        if (update.roleId === requestedChange.newFlag.roleId && update.permissionId === requestedChange.permissionId) {
          matchUpdate = update;
        }
      }

      let message;
      let status;
      if (matchUpdate) {
        if (matchUpdate.hasOwnProperty("roleId")) {
          message = "";
          status = "saved";
        } else if (matchUpdate.hasOwnProperty("error")) {
          message = matchUpdate.error;
          status = "failed";
        }
      }

      results.push({
          requestedChange,
          status,
          message
        }
      );

      console.log(this.props.updates);
    });

    return (
      <DefaultModal open={this.props.open} closeHandler={this.props.closeHandler}>
        { this.state.saveClicked ?
        <div>
          <h1>SAVING CHANGES</h1>
          <Table>
            {results.map((result, index, array) => {
              return <Row key={index}>
                <Cell>
                  {result.status === null ? <InputFeedback message={ result.message } type={FeedbackType.Progress} /> : null }
                  {result.status === "failed" ? <InputFeedback message={ result.message } type={FeedbackType.Error} /> : null }
                  {result.status === "saved" ? <InputFeedback message={ result.message } type={FeedbackType.Confirm} /> : null }

                  {result.requestedChange.userDescription } &nbsp; {result.requestedChange.newFlag.value ? <span>gains</span> : <span>loses</span>} &nbsp; {result.requestedChange.description}

                  {result.status === "saved" ? <span className="green">Saved</span> : null}
                  {result.status === "failed" ? <span className="red">Failed</span> : null}
                </Cell>
              </Row>;
              })
            }
          </Table>
          <br />
          &nbsp;
          <div className="button-secondary" onClick={this.props.closeHandler}>Cancel</div>
        </div> :
        <div>
          <h1>CONFIRMING PERMISSION CHANGES</h1>
          <p>You are making the following changes</p>
          <Table>
            <Row header={true}>
              <Cell>CHANGE</Cell>
              <Cell>CONFIRM</Cell>
            </Row>
            { this.state.requestedChanges ? this.state.requestedChanges.map((requestedChange, index, array) => {
              return <Row key={index}>
                <Cell>{requestedChange.userDescription } &nbsp; {requestedChange.newFlag.value ? <span>gains</span> : <span>loses</span>} &nbsp; {requestedChange.description}</Cell>
                <Cell>{requestedChange.confirmed ? <span className="green"><i className="fa fa-check"></i> Confirmed</span> : <span className="button-primary" onClick={(e) => this.confirmClicked(index)}>Confirm</span>} &nbsp; <span onClick={() => this.cancelIndividualChange(index)} className="button">Cancel Change</span></Cell>
              </Row>;
            }) : null }
          </Table>
          <br />
          {saveChangesEnabled ? <div className="button-primary" onClick={this.onSaveChangesClicked}>Save Changes</div>
            : <div className="button-primary disabled">Save Changes</div> }
           &nbsp;
          <div className="button-secondary" onClick={this.props.closeHandler}>Cancel</div>
        </div>}


      </DefaultModal>
    );
  }
}
