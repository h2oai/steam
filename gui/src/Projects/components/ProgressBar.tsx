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
import * as _ from 'lodash';
import '../styles/progressbar.scss';

interface Props {
  showPercentage: boolean,
  className?: any,
  onComplete?: Function,
  onClick?: Function
}
export default class ProgressBar extends React.Component<Props, any> {
  interval: number;

  constructor() {
    super();
    this.state = {
      progress: 0
    };
  }

  componentDidMount() {
    this.start();
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  start() {
    let maxIncrements = Math.floor(Math.random() * (100 - 40 + 1)) + 40;
    let i = 0;
    this.interval = setInterval(() => {
      i++;
      let remaining = 100 - this.state.progress;
      this.setState({
        progress: this.state.progress += (0.05 * Math.pow(1 - Math.sqrt(remaining), 2))
      });
      if (i >= maxIncrements) {
        this.end();
      }
    }, 50);
  }

  end() {
    this.setState({
      progress: 100
    });
    clearInterval(this.interval);
    if (this.props.onComplete) {
      this.props.onComplete(this);
    }
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div ref="progressBar"
           className={classNames('progress-bar-container', {complete: this.state.progress === 100}, this.props.className)} onClick={this.props.onClick}>
        <div className="progress-bar" style={{width: this.state.progress + '%'}}>
        </div>
        <div
          className="progress-counter">{this.props.showPercentage === true && _.isEmpty(this.props.children) ? Math.ceil(this.state.progress) + '%' : null}</div>
      </div>
    );
  }
}
