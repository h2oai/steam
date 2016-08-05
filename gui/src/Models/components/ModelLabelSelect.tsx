/**
 * Created by Jeff Fohl <jfohl@h2o.ai> on 6/29/16.
 */
import * as React from 'react';

interface Props {
  modelId: number,
  projectId: number,
  labels: any[]
}

export default class ModelLabelSelect extends React.Component<Props, any> {

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <select name="labelSelect">
        {this.props.labels[this.props.projectId].map((label) => {
          return (
            <option key={label.id} value={label.id}>{label.name}</option>
          );
        })}
      </select>
    );
  }
}
