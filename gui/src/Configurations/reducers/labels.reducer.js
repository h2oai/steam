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
 * Created by Jeff Fohl <jfohl@h2o.ai> on 08/01/16.
 */
var _ = require('lodash');
var configuration_labels_action_1 = require('../actions/configuration.labels.action');
var configuration_labels_action_2 = require('../actions/configuration.labels.action');
var configuration_labels_action_3 = require('../actions/configuration.labels.action');
var initialState = {};
exports.labelsReducer = function (state, action) {
    if (state === void 0) { state = initialState; }
    switch (action.type) {
        case configuration_labels_action_1.RECEIVE_LABELS:
            return receiveLabels(state, action);
        case configuration_labels_action_2.RECEIVE_CREATE_LABEL:
            return receiveCreateLabel(state, action);
        case configuration_labels_action_3.RECEIVE_UPDATE_LABEL:
            return receiveUpdateLabel(state, action);
        default:
            return state;
    }
};
function receiveLabels(state, action) {
    var projectLabels = {};
    projectLabels[action.projectId] = action.labels;
    return _.assign({}, state, projectLabels);
}
function receiveCreateLabel(state, action) {
    var labels;
    if (state[action.projectId]) {
        labels = state[action.projectId].slice();
        labels.push(action.label);
    }
    else {
        labels = [action.label];
    }
    var projectLabels = {};
    projectLabels[action.projectId] = labels;
    return _.assign({}, state, projectLabels);
}
function receiveUpdateLabel(state, action) {
    return state;
}
//# sourceMappingURL=labels.reducer.js.map