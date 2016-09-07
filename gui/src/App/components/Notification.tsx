/**
 * Created by justin on 8/8/16.
 */
import * as React from 'react';
import '../styles/notification.scss';
import { TransitionMotion, spring } from 'react-motion';
import { NotificationData } from '../actions/notification.actions';

export enum NotificationType { Info, Confirm, Warning, Error }

interface Props {
  notificationData: NotificationData
}

interface DispatchProps {
  closeNotification: Function
}

export class Notification extends React.Component<Props & DispatchProps, any> {

  constructor() {
    super();
    this.state = {
      isActive: false,
      hasViewButton: false,
      styles: [{key: 'a', style : { top: -150, opacity: 1}}],
      defaultStyles: [{key: 'a', style : { top: -150, opacity: 1}}]
    };
  }

  componentWillMount(): void {
    if (this.props.notificationData) {
      if ( this.props.notificationData.detail.length > 100) {
        this.setState({
          hasViewButton: true
        });
      }
    }
  }

  componentWillReceiveProps(nextProps): void {
    if ( nextProps.notificationData ) {
      if ( nextProps.notificationData.isActive ) {
        this.setState( {
          styles: [{key: 'a', style : { top: spring(22), opacity: 1}}],
          defaultStyles: [{key: 'a', style : { top: -150, opacity: 1}}]
        });
      } else {
        this.setState( {
          styles: [{key: 'a', style: { top: 22, opacity: spring(0) }}],
          defaultStyles: [{key: 'a', style : { top: 22, opacity: 1}}]
        });
      }
    }
  }

  render(): React.ReactElement<HTMLElement> {
    if ( this.props.notificationData ) {
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
      console.log(guideClassName);

      return (
        <TransitionMotion styles={ this.state.styles } defaultStyles={ this.state.defaultStyles }>
          {animationValue =>
            <div className='notification' open={this.props.notificationData.isActive} style={animationValue[0].style}>
              { console.log(animationValue[0].style) }
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
                    <button onClick={this.props.closeNotification} class="view-button"> View</button> : null }
                  <button onClick={this.props.closeNotification} className="dismiss-button"> Dismiss</button>
                </div>
              </div>
            </div>
          }
        </TransitionMotion>
      );
    } else {
      console.log('returning null');
      return null;
    }
  }
}
