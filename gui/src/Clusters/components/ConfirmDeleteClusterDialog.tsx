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
import '../styles/clusters.scss';
import DefaultModal from '../../App/components/DefaultModal';
import {Cluster} from "../../Proxy/Proxy";

interface Props {
  clusterToDelete: Cluster
  onDeleteConfirmed: Function
  onCancel: Function
  isOpen: boolean
}

export default class ConfirmDeleteClusterDialog extends React.Component<Props, any> {

  onCancelClicked = () => {
    this.props.onCancel();
  };

  onConfirmClicked = () => {
    this.props.onDeleteConfirmed();
  };

  render(): React.ReactElement<DefaultModal> {
    return (
      <DefaultModal open={this.props.isOpen}>
        <h1>CONFIRM DELETE CLUSTER</h1>&nbsp;
        {this.props.clusterToDelete ? <p>Are you sure you wish to delete cluster {this.props.clusterToDelete.name}?</p> : null}
        <br />&nbsp;
        <div className="button-primary" onClick={this.onConfirmClicked}>Confirm</div>
        <div className="button-secondary" onClick={this.onCancelClicked}>Cancel</div>
      </DefaultModal>
    );
  }
}
