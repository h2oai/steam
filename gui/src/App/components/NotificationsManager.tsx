import * as React from 'react';
import * as _ from 'lodash';
import { connect } from 'react-redux';
import {
  closeNotificationManager, dismissNotification, NotificationData, killAllInactiveNotifications
} from '../actions/notification.actions';
import { bindActionCreators } from 'redux';
import { Notification } from "./Notification";
import { Motion, spring } from 'react-motion';


interface Props {
  notifications
}

interface DispatchProps {
  dismissNotification: Function,
  closeNotificationManager: Function,
  killAllInactiveNotifications: Function
}

export class NotificationsManager extends React.Component<Props & DispatchProps, any> {

  timeoutTimerID = -1;
  TIMEOUT_LENGTH = 5000;

  constructor() {
    super();
    this.state = {
      isActive: false,
      hasViewButton: false,
      style : { top: -150, opacity: 1},
      defaultStyle: { top: -150, opacity: 1}
    };
  }

  componentWillReceiveProps(nextProps): void {
    if ( nextProps.notifications ) {
      let isActive = false;
      for (let notificationData of nextProps.notifications.allNotifications) {
       if (notificationData.isActive) {
         isActive = true;
       }
      }

      if (isActive) {
        this.setState({
          isActive: true,
          style: {top: spring(22), opacity: 1},
          defaultStyle: { top: -150, opacity: 1}
        });
      } else {
        this.setFadeout();
      }
    }

    if (this.timeoutTimerID !== -1) {
      clearTimeout(this.timeoutTimerID);
      this.timeoutTimerID = setTimeout(this.onTimeout.bind(this), this.TIMEOUT_LENGTH);
    }
  }

  setFadeout() {
    this.setState({
      isActive: false,
      style: {top: -150, opacity: spring(0)},
      defaultStyle: {top: 22, opacity: 1}
    });
  }

  onAnimationComplete() {
    if (this.state.isActive) {
      this.startTimeoutTimer();
    } else {
      this.props.killAllInactiveNotifications();
    }
  }

  startTimeoutTimer() {
    this.timeoutTimerID = setTimeout(this.onTimeout.bind(this), this.TIMEOUT_LENGTH);
  }
  
  onTimeout() {
    for (let notificationData of this.props.notifications.allNotifications) {
      notificationData.isActive = false;
    }
    this.setFadeout();
  }

  dismissNotification(notificationData) {
    this.props.dismissNotification(notificationData);
  }

  render(): React.ReactElement<HTMLElement> {
    if ( this.props.notifications.allNotifications ) {
      return (
        <Motion defaultStyle={ this.state.defaultStyle } style={ this.state.style } onRest={this.onAnimationComplete.bind(this) }>
          {animationValue => {
            let modifiedAnimationValue: any = _.assign({},animationValue);
            if (modifiedAnimationValue.opacity < 1) {
              modifiedAnimationValue.top = 22;
            }
            return <div className="notification-manager"  style={modifiedAnimationValue}> {
              this.props.notifications.allNotifications.map((notificationData, index, array) => {
                if (notificationData.isAlive) {
                  return <Notification key={index} notificationData={ notificationData } dismissNotification={this.dismissNotification.bind(this)} />;
                }
              })
            }
            </div>;
            }
          }
        </Motion>
      );
    } else {
      return null;
    }
  }
}

function mapStateToProps(state) {
  return {
    notifications: state.notification
  };
}

function mapDispatchToProps(dispatch) {
  return {
    dismissNotification: bindActionCreators(dismissNotification, dispatch),
    closeNotificationManager: bindActionCreators(closeNotificationManager, dispatch),
    killAllInactiveNotifications: bindActionCreators(killAllInactiveNotifications, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(NotificationsManager);
