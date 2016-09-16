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
