/**
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import '../styles/cornerborders.scss';

export default class CornerBorders extends React.Component<any, any> {
  render() {
    return (
      <div className="corner-borders">
        {this.props.children}
      </div>
    );
  }
}