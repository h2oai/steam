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
 * Created by justin on 9/15/16.
 */
var _ = require('lodash');
var clusters_actions_1 = require('../actions/clusters.actions');
var initialState = {
    engines: [],
    config: null
};
exports.clustersReducer = function (state, action) {
    if (state === void 0) { state = initialState; }
    switch (action.type) {
        case clusters_actions_1.RECEIVE_ENGINES:
            return _.assign({}, state, {
                engines: action.engines
            });
        case clusters_actions_1.FETCH_CONFIG_COMPLETED:
            return _.assign({}, state, {
                config: action.config
            });
        default:
            return state;
    }
};
//# sourceMappingURL=clusters.reducer.js.map