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
"use strict";
/**
 * Created by justin on 6/28/16.
 */
var _ = require('lodash');
var leaderboard_actions_1 = require('../actions/leaderboard.actions');
var initialState = {
    items: [],
    modelCategory: null,
    criteria: null,
    count: 0
};
exports.leaderboardReducer = function (state, action) {
    if (state === void 0) { state = initialState; }
    switch (action.type) {
        case leaderboard_actions_1.RECEIVE_LEADERBOARD:
            return _.assign({}, state, {
                items: action.leaderboard
            });
        case leaderboard_actions_1.RECEIVE_SORT_CRITERIA:
            return _.assign({}, state, {
                criteria: action.criteria
            });
        case leaderboard_actions_1.RECEIVE_MODEL_COUNT:
            return _.assign({}, state, {
                count: action.count
            });
        default:
            return state;
    }
};
//# sourceMappingURL=leaderboard.reducer.js.map