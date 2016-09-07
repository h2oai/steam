/**
 * Created by justin on 8/8/16.
 */
import * as _ from 'lodash';
import {OPEN_NOTIFICATION, CLOSE_NOTIFICATION, NotificationData} from '../actions/notification.actions';

const initialState = {
  allNotifications: []
};

export function notificationReducer(state = initialState, action) {
  switch (action.type) {
    case OPEN_NOTIFICATION:
      return {
        allNotifications: state.allNotifications.concat([action.notificationData])
      };
    case CLOSE_NOTIFICATION:
      let toReturn: any = _.assign({}, state);
      toReturn.allNotifications[toReturn.allNotifications.length - 1].isActive = false;
      return toReturn;
    default:
      return state;
  }
}
