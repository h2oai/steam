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
 * Created by justin on 8/8/16.
 */
import * as React from 'react';
import '../styles/notification.scss';
import { NotificationData } from '../actions/notification.actions';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { dismissNotification } from '../actions/notification.actions';

export enum NotificationType { Info, Confirm, Warning, Error }

interface Props {
  notificationData: NotificationData,
  dismissNotification: Function
}

export class Notification extends React.Component<Props, any> {

  constructor() {
    super();
    this.state = {
      hasViewButton: false,
    };
  }

  onDismissClicked() {
    this.props.dismissNotification(this.props.notificationData);
  }

  render(): React.ReactElement<HTMLElement> {
    if (this.props.notificationData) {
      let guideClassName = "left-type-guide ";
      switch (this.props.notificationData.notificationType) {
        case NotificationType.Info :
          guideClassName += "left-type-guide-info ";
          break;
        case NotificationType.Warning :
          guideClassName += "left-type-guide-warning ";
          break;
        case NotificationType.Error :
          guideClassName += "left-type-guide-error ";
          break;
        case NotificationType.Confirm :
          guideClassName += "left-type-guide-confirm ";
          break;
      }

      return <div className='notification'>
        <div className="inner-notification">
          <div className={guideClassName}></div>
          <div className="text-container">
            <div className="header">
              {this.props.notificationData.header}
            </div>
            <div className="detail">
              {this.props.notificationData.detail}
            </div>
          </div>
          <div className="actions-container">
            { this.state.hasViewButton ?
              <div onClick={ this.onDismissClicked.bind(this) } className="view-button"> View</div> : null }
            <div onClick={ this.onDismissClicked.bind(this) } className="dismiss-button"> Dismiss</div>
          </div>
        </div>
      </div>;
    } else {
      return null;
    }

  }
}
