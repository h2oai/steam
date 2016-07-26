/**
 * Created by justin on 6/27/16.
 */
import * as React from 'react';
import * as classNames from 'classnames';
import '../styles/collapsible.scss';

interface Props {
  open: boolean
}
export default class Collapsible extends React.Component<Props, any> {
  constructor() {
    super();
    this.state = {
      open: false
    };
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className={classNames('collapsible', {closed: !this.props.open})}>
        {this.props.children}
      </div>  
    );
  }
}