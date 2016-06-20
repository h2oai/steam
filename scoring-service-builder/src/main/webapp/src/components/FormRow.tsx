/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import Row from './Row';
import * as classNames from 'classnames';

interface Props {
  className?: string,
  children?: any,
  id?: string,
  onSubmit?: Function
}

export default class FormRow extends React.Component<Props, any> {
  render() {
    return (
      <Row className={'form-group'}>
        {this.props.children}
      </Row>
    );
  }
}