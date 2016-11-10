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
import '../styles/users.scss';
import DefaultModal from '../../App/components/DefaultModal';
import {Identity} from "../../Proxy/Proxy";

interface Props {
  user: Identity
  deleteUserAction: Function
  closeHandler: Function
  open: boolean
}

export default class DeleteUserConfirm extends React.Component<Props, any> {

  onCancelClicked = () => {
    this.props.closeHandler();
  };

  onConfirmClicked = () => {
    this.props.deleteUserAction(this.props.user.id);
    this.props.closeHandler();
  };

  render(): React.ReactElement<DefaultModal> {
    return (
      <DefaultModal open={this.props.open}>
        <div>
          <h1>CONFIRM DELETE USER</h1>
          <br />&nbsp;
          <p>This will removed user: {this.props.user ? this.props.user.name : null}</p>
          <br />&nbsp;
          <div className="button-primary" onClick={this.onConfirmClicked}>Confirm</div>
          <div className="button-secondary" onClick={this.onCancelClicked}>Cancel</div>
        </div>
      </DefaultModal>
    );
  }
}
