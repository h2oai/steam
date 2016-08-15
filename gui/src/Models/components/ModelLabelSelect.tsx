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
