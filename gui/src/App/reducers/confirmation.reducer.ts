/**
 * Created by justin on 8/12/16.
 */
import * as _ from 'lodash';
import { OPEN_CONFIRMATION, CLOSE_CONFIRMATION } from '../actions/confirmation.actions';

const initialState = {
  isOpen: false,
  notificationType: null,
  text: '',
  title: ''
};

export function confirmationReducer(state = initialState, action) {
  switch (action.type) {
    case OPEN_CONFIRMATION:
      return _.assign({}, state, action);
    case CLOSE_CONFIRMATION:
      return _.assign({}, state, initialState);
    default:
      return state;
  }
}
