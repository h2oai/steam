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
 * Created by Jeff Fohl <jfohl@h2o.ai> on 6/29/16.
 */
import * as React from 'react';
import * as _ from 'lodash';
import { Label } from '../../Proxy/Proxy';
interface Props {
  modelId: number,
  projectId: number,
  labels: {
    [projectId: number]: Label[]
  },
  onChangeHandler: Function
}

export default class ModelLabelSelect extends React.Component<Props, any> {
  onChangeHandler(event) {
    let value = parseInt(event.target.value, 10);
    if (value === -1) {
      this.props.onChangeHandler(_.find(this.props.labels[this.props.projectId], {model_id: this.props.modelId}).id, this.props.modelId, true);
    } else {
      this.props.onChangeHandler(parseInt(event.target.value, 10), this.props.modelId, false);
    }
  }
  render(): React.ReactElement<HTMLSelectElement> {
    if (_.isUndefined(this.props.labels[this.props.projectId])) {
      return <select name="labelSelect"><option value={-1}></option></select>;
    }
    return (
      <select name="labelSelect" onChange={this.onChangeHandler.bind(this)} value={_.find(this.props.labels[this.props.projectId], {model_id: this.props.modelId}) ? _.find(this.props.labels[this.props.projectId], {model_id: this.props.modelId}).id : -1}>
        <option value={-1}></option>
        {this.props.labels[this.props.projectId] ? this.props.labels[this.props.projectId].map((label: Label) => {
          return (
            <option key={label.id} value={label.id}>{label.name}</option>
          );
        }) : <option></option>}
      </select>
    );
  }
}
