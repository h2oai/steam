/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';

interface Props {
  type?: string
}

export default class Input extends React.Component<Props, any> {
  render() {
    return (
      <input {...this.props}>
        {this.props.children}
      </input>
    );
  }
}