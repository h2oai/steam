/**
 * Created by justin on 6/28/16.
 */
import * as _ from 'lodash';
import { RECEIVE_LEADERBOARD, RECEIVE_SORT_CRITERIA } from '../actions/leaderboard.actions';

let initialState = {
  items: [],
  modelCategory: null
};

export const leaderboardReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case RECEIVE_LEADERBOARD:
      return _.assign({}, state, {
        items: action.leaderboard
      });
    case RECEIVE_SORT_CRITERIA:
      return _.assign({}, state, {
        criteria: action.criteria
      });
    default:
      return state;
  }
};
