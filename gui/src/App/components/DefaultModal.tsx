/**
 * Created by justin on 7/21/16.
 */
import * as React from 'react';
import * as classNames from 'classnames';
import { Overlay } from 'h2oUIKit';
import '../styles/defaultmodal.scss';

interface Props {
  className?: any,
  open: boolean
  closeHandler?: Function
}

export default class Deploy extends React.Component<Props, any> {
  constructor() {
    super();
    this.onClick = this.onClick.bind(this);
  }

  onClick() {
    this.props.closeHandler();
  }

  render(): React.ReactElement<Overlay> {
    return (
      <Overlay className={classNames(this.props.className)} open={this.props.open}>
        <div className="default-modal-container">
          <div className="default-modal">
            <div className="content">
              {this.props.children}
            </div>
          </div>
        </div>
      </Overlay>
    );
  }
}
