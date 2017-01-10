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
import { Intent } from '@blueprintjs/core';
import { NotificationType } from '../components/Notification';
import { toastManager } from '../components/ToastManager';


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
  let message = React.createElement(
    'div',
    null,
    React.createElement('div', { className: 'notification-indicator' }),
    React.createElement(
      'div',
      { className: 'notification-content' },
      detail
    )
  );
  let intent;
  switch (notificationType) {
    case NotificationType.Confirm:
      intent = Intent.PRIMARY;
      break;
    case NotificationType.Error:
      intent = Intent.DANGER;
      break;
    case NotificationType.Info:
      intent = Intent.SUCCESS;
      break;
    case NotificationType.Warning:
      intent = Intent.WARNING;
      break;
    default :
      console.log("ERROR: Unexpected notification type");
  }
  let timeout = 5000;
  if (intent === Intent.DANGER) {
    timeout = 0;
  }

  toastManager.show({
    message,
    intent,
    className: "steam-notification",
    timeout
  });

  return {
    type: OPEN_NOTIFICATION,
    notificationData: {
      isActive: false,
      isAlive: false,
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
