/**
 * Created by justin on 8/8/16.
 */
import * as _ from 'lodash';
import {OPEN_NOTIFICATION, DISMISS_NOTIFICATION, KILL_ALL_INACTIVE_NOTIFICATIONS, NotificationData} from '../actions/notification.actions';

const initialState = {
  allNotifications: []
};

export function notificationReducer(state = initialState, action) {
  let toReturn: any;
  switch (action.type) {
    case OPEN_NOTIFICATION:
      return {
        allNotifications: state.allNotifications.concat([action.notificationData])
      };
    case KILL_ALL_INACTIVE_NOTIFICATIONS:
      toReturn = _.assign({}, state);
      for (let notification of toReturn.allNotifications) {
        if (!notification.isActive) {
          notification.isAlive = false;
        }
      }
      return toReturn;
    case DISMISS_NOTIFICATION:
      toReturn = _.assign({}, state);
      toReturn.allNotifications[action.notification.index].isActive = false;
      toReturn.allNotifications[action.notification.index].isAlive = false;
      return toReturn;
    default:
      return state;
  }
}
