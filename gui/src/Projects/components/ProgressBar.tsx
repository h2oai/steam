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
  start?: boolean,
  end?: boolean,
  onComplete?: Function,
  onClick?: Function,
  style?: any
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
    if (this.props.start === true || _.isUndefined(this.props.start)) {
      this.start();
    }
  }

  componentWillReceiveProps(nextProps) {
    if (_.isUndefined(nextProps.start) && this.state.progress === 0) {
      this.start();
    }
    if (nextProps.end === true && this.state.progress !== 100) {
      this.end();
    }
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  start() {
    this.interval = setInterval(() => {
      let remaining = 100 - this.state.progress;
      this.setState({
        progress: this.state.progress += (0.05 * Math.pow(1 - Math.sqrt(remaining), 2))
      });
    }, 50);
  }

  end() {
    this.setState({
      progress: 100
    });
    clearInterval(this.interval);
    let timeout = setTimeout(() => {
      if (this.props.onComplete) {
        this.props.onComplete(this);
      }
      clearTimeout(timeout);
    }, 100);
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
