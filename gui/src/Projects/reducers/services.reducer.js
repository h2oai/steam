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
 * Created by justin on 7/22/16.
 */
var _ = require('lodash');
var services_actions_1 = require('../actions/services.actions');
var initialState = {
    runningServices: []
};
exports.servicesReducer = function (state, action) {
    if (state === void 0) { state = initialState; }
    switch (action.type) {
        case services_actions_1.RECEIVE_SERVICES_FOR_PROJECT:
            return _.assign({}, state, {
                runningServicesForProject: action.services
            });
        case services_actions_1.RECEIVE_ALL_SERVICES:
            return _.assign({}, state, {
                allRunningServices: action.services
            });
        default:
            return state;
    }
};
//# sourceMappingURL=services.reducer.js.map