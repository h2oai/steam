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
 * Created by justin on 7/12/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import '../styles/panel.scss';

interface Props {
  className?: any,
  onClick?: Function
}

export default class Panel extends React.Component<Props, any> {

  onClick = (e: React.MouseEvent<any>) => {
    if (this.props.onClick) {
      this.props.onClick();
    }
  };

  render() {
    return (
      <div className={classNames('panel', this.props.className)} onClick={this.onClick}>
        {this.props.children}
      </div>
    );
  }
}
