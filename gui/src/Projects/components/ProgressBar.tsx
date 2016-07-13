/**
 * Created by justin on 7/12/16.
 */
import * as React from 'react';
import * as classNames from 'classnames';
import '../styles/progressbar.scss';

interface Props {
  progress: number,
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
