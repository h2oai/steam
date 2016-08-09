/**
 * Created by justin on 8/8/16.
 */
import * as _ from 'lodash';
import { OPEN_NOTIFICATION, CLOSE_NOTIFICATION } from '../actions/notification.actions';

const initialState = {
  isOpen: false,
  notificationType: null,
  text: ''
};

export function notificationReducer(state = initialState, action) {
  switch (action.type) {
    case OPEN_NOTIFICATION:
      return _.assign({}, state, action);
    case CLOSE_NOTIFICATION:
      return _.assign({}, state, initialState);
    default:
      return state;
  }
}
