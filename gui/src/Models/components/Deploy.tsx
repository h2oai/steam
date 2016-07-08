/**
 * Created by justin on 6/30/16.
 */
import * as React from 'react';
import '../styles/deploy.scss';
import { Overlay } from 'h2oUIKit';

interface Props {
  open: boolean
  closeHandler: Function
}

export default class Deploy extends React.Component<Props, any> {
  constructor() {
    super();
    this.onClick = this.onClick.bind(this);
    this.state = {
      isOpen: false
    };
    console.log(Overlay);
  }

  onClick() {
    this.props.closeHandler();
  }

  render(): React.ReactElement<Overlay> {
    return (
      <Overlay open={this.props.open}>
        <div className="deploy-modal">
          <div className="content">
            <button className="close-button" onClick={this.onClick}><i className="fa fa-close"></i></button>
            <header>Model</header>
            <section>
              <div>Select method of deployment</div>
              <div>
                <ul>
                  <li>Jar download</li>
                  <li>POJO download</li>
                  <li>Remote WAR</li>
                  <li>Local WAR</li>
                </ul>
              </div>
            </section>
          </div>
        </div>
      </Overlay>
    );
  }
}