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
