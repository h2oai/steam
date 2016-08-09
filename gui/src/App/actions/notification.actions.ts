/**
 * Created by justin on 8/8/16.
 */

export const OPEN_NOTIFICATION = 'OPEN_NOTIFICATION';
export const CLOSE_NOTIFICATION = 'CLOSE_NOTIFICATION';

export function openNotification(notificationType, text, actions) {
  return {
    type: OPEN_NOTIFICATION,
    notificationType,
    text,
    actions,
    isOpen: true
  };
}

export function closeNotification() {
  return {
    type: CLOSE_NOTIFICATION
  };
}
