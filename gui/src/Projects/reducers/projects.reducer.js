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
 * Created by justin on 7/18/16.
 */
var _ = require('lodash');
var projects_actions_1 = require('../actions/projects.actions');
var initialState = {
    clusters: [],
    models: [],
    project: {},
    availableProjects: null,
    isClusterFetchInProcess: false
};
exports.projectsReducer = function (state, action) {
    if (state === void 0) { state = initialState; }
    switch (action.type) {
        case projects_actions_1.SET_CURRENT_PROJECT:
            if (state.project.hasOwnProperty("id")) {
                if (state.project.id === action.projectId) {
                    return state;
                }
            }
            var toReturn = _.assign({}, state);
            toReturn.project = { id: action.projectId };
            return toReturn;
        case projects_actions_1.REQUEST_CLUSTERS:
            return _.assign({}, state, {
                isClusterFetchInProcess: true
            });
        case projects_actions_1.RECEIVE_CLUSTERS:
            return _.assign({}, state, {
                clusters: action.clusters,
                isClusterFetchInProcess: false,
                registerClusterError: null
            });
        case projects_actions_1.REQUEST_MODELS:
            return _.assign({}, state, {
                isModelFetchInProcess: true
            });
        case projects_actions_1.RECEIVE_MODELS:
            return _.assign({}, state, {
                models: action.models,
                isModelFetchInProcess: false
            });
        case projects_actions_1.RECEIVE_PROJECT:
            return _.assign({}, state, {
                project: action.project
            });
        case projects_actions_1.RECEIVE_MODELS_FROM_PROJECT:
            return _.assign({}, state, {
                models: action.models
            });
        case projects_actions_1.CREATE_PROJECT_COMPLETED:
            return _.assign({}, state, {
                project: action.project
            });
        case projects_actions_1.RECEIVE_PROJECTS:
            return _.assign({}, state, {
                availableProjects: action.projects
            });
        case projects_actions_1.RECEIVE_DATASETS_FROM_CLUSTER:
            return _.assign({}, state, {
                datasets: action.datasets
            });
        case projects_actions_1.REGISTER_CLUSTER_ERROR:
            return _.assign({}, state, {
                registerClusterError: action.message
            });
        default:
            return state;
    }
};
//# sourceMappingURL=projects.reducer.js.map