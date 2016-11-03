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
import * as React from 'react';
import '../styles/input_feedback.scss';

export enum FeedbackType { Progress, Info, Confirm, Warning, Error }

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
        <div className="feedback-inner">
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

        { this.props.type === FeedbackType.Confirm ?
        <div className="input-feedback feedback-confirm">
          <div className="feedback-confirm-icon">
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
