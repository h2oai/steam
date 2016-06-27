/**
 * Created by justin on 6/27/16.
 */

import { TOGGLE_MENU } from '../actions/navigation.action';
import * as _ from 'lodash';

let initialState = {
  isOpen: false
};

export const toggleMenu = (isOpen: boolean) => {
  return {
    type: TOGGLE_MENU,
    isOpen
  };
};

export const navigationReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case TOGGLE_MENU:
      return _.assign({}, state, {
        isOpen: action.isOpen
      });
    default:
      return state;
  }
};