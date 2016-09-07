import * as React from 'react';
import * as _ from 'lodash';
import { connect } from 'react-redux';
import { closeNotification, NotificationData } from '../actions/notification.actions';
import { bindActionCreators } from 'redux';
import { Notification } from "./Notification";


interface Props {
  notifications
}

interface DispatchProps {
  closeNotification: Function
}

export class NotificationsManager extends React.Component<Props & DispatchProps, any> {

  constructor() {
    super();
  }

  componentWillMount(): void {

  }

  render(): React.ReactElement<HTMLElement> {
    //todo: switch to animating the whole NotificationsManager, then just add new static notification elements within a FlexBox and reset the fadeout timer
    if (!_.isEmpty(this.props.notifications)) {
      return <Notification notificationData={ this.props.notifications.allNotifications[] }
                           closeNotification={ this.props.closeNotification }/>;
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
    closeNotification: bindActionCreators(closeNotification, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(NotificationsManager);
