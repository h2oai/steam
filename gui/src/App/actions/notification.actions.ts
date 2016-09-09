/**
 * Created by justin on 8/8/16.
 */

import { NotificationType } from '../components/Notification';

export const OPEN_NOTIFICATION = 'OPEN_NOTIFICATION';
export const CLOSE_NOTIFICATION = 'CLOSE_NOTIFICATION';
export const KILL_ALL_INACTIVE_NOTIFICATIONS = 'KILL_ALL_INACTIVE_NOTIFICATIONS';
export const DISMISS_NOTIFICATION = 'DISMISS_NOTIFICATION';
export const CLOSE_NOTIFICATION_MANAGER = 'CLOSE_NOTIFICATION_MANAGER';

export interface NotificationData {
  isActive: Boolean;
  isAlive: Boolean;
  notificationType: NotificationType;
  header: string;
  detail;
  action;
  index: number;
}

export function openNotification(notificationType: NotificationType, header: string, detail, actions) {
  return (dispatch, getState) => {
    dispatch(_openNotification(notificationType, header, detail, actions, getState()));
  };
}

function _openNotification(notificationType: NotificationType, header: string, detail, actions, state) {
  let index = state.notification.allNotifications.length;
  return {
    type: OPEN_NOTIFICATION,
    notificationData: {
      isActive: true,
      isAlive: true,
      notificationType,
      header,
      detail,
      actions,
      index
    }
  };
}

export function dismissNotification(notification) {
  return (dispatch, getState) => {
    dispatch({
        type: DISMISS_NOTIFICATION,
        notification
    });
  };
}

export function closeNotificationManager() {
  return {
    type: CLOSE_NOTIFICATION_MANAGER
  };
}

export function killAllInactiveNotifications() {
  return {
    type: KILL_ALL_INACTIVE_NOTIFICATIONS
  };
}
