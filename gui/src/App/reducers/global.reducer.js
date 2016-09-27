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
var _ = require('lodash');
var global_actions_1 = require('../actions/global.actions');
var initialState = {
    entityIds: {}
};
function globalReducer(state, action) {
    if (state === void 0) { state = initialState; }
    switch (action.type) {
        case global_actions_1.RECEIVE_ENTITY_IDS:
            var entityIds = {};
            for (var _i = 0, _a = action.response; _i < _a.length; _i++) {
                var entityType = _a[_i];
                entityIds[entityType.name] = entityType.id;
            }
            return _.assign({}, state, { entityIds: entityIds });
        default:
            return state;
    }
}
exports.globalReducer = globalReducer;
//# sourceMappingURL=global.reducer.js.map