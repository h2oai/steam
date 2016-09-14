/**
 * Created by justin on 6/28/16.
 */
import * as _ from 'lodash';
import { RECEIVE_LEADERBOARD, RECEIVE_SORT_CRITERIA, RECEIVE_MODEL_COUNT, RECEIVE_DELETE_MODEL, REQUEST_DELETE_MODEL } from '../actions/leaderboard.actions';

let initialState = {
  items: [],
  modelCategory: null,
  criteria: null,
  count: 0
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
    case RECEIVE_MODEL_COUNT:
      return _.assign({}, state, {
        count: action.count
      });
    case REQUEST_DELETE_MODEL:
      var toReturn: any = _.assign({}, state);
      toReturn.items = toReturn.items.slice();
      for (let model of toReturn.items) {
        if (model.id === action.id) {
          model.isDeleteInProgress = true;
        }
      }
      return toReturn;
    case RECEIVE_DELETE_MODEL:
      if (!action.successful) {
        var toReturn: any = _.assign({}, state);
        toReturn.items = toReturn.items.slice();
        for (let model of toReturn.items) {
          if (model.id === action.id) {
            model.isDeleteInProgress = false;
          }
        }
        return toReturn;
      } else {
        return state;
      }
    default:
      return state;
  }
};
