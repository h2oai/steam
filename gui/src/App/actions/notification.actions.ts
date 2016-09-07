/**
 * Created by justin on 8/8/16.
 */

import { NotificationType } from '../components/Notification';

export const OPEN_NOTIFICATION = 'OPEN_NOTIFICATION';
export const CLOSE_NOTIFICATION = 'CLOSE_NOTIFICATION';

export interface NotificationData {
  isActive: Boolean;
  notificationType: NotificationType;
  header: string;
  detail;
  action;
}

export function openNotification(notificationType: NotificationType, header: string, detail, actions) {
  return {
    type: OPEN_NOTIFICATION,
    notificationData: {
      isActive: true,
      notificationType,
      header,
      detail,
      actions
    }
  };
}

export function closeNotification() {
  return {
    type: CLOSE_NOTIFICATION
  };
}
