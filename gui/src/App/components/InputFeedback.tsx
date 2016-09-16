import * as React from 'react';
import '../styles/input_feedback.scss';

export enum FeedbackType { Progress, Info, Success, Warning, Error }

interface Props {
  message: string,
  type: FeedbackType
}

export default class InputFeedback extends React.Component<Props, any> {
  constructor() {
    super();
  }

  render(): React.ReactElement<HTMLElement> {
    return (
      <div className="input-feedback feedback-progress">
        { this.props.type === FeedbackType.Progress ?
        <div>
          <div className='uil-facebook-css'>
            <div></div>
            <div></div>
            <div></div>
          </div>
          <span>{this.props.message}</span>
        </div>
        : null }

        { this.props.type === FeedbackType.Info ?
          <div className="input-feedback feedback-info">
            <span>{this.props.message}</span>
          </div>
        : null }

        { this.props.type === FeedbackType.Success ?
        <div className="input-feedback feedback-success">
          <div className="feedback-success-icon">
            <i className="fa fa-check-circle" aria-hidden="true"></i>
          </div>
          <span>{this.props.message}</span>
        </div>
        : null }

        { this.props.type === FeedbackType.Warning ?
        <div className="input-feedback feedback-warning">
          <div className="feedback-warning-icon">
            <i className="fa fa-exclamation-triangle" aria-hidden="true"></i>
          </div>
          <span>{this.props.message}</span>
        </div>
        : null }

        { this.props.type === FeedbackType.Error ?
        <div className="input-feedback feedback-error">
          <div className="feedback-error-icon">
            <i className="fa fa-exclamation-triangle" aria-hidden="true"></i>
          </div>
          <span>{this.props.message}</span>
        </div>
        : null }
      </div>
    );
  }
}
