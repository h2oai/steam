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
