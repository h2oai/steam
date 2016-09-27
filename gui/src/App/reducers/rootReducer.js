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
 * Created by justin on 6/25/16.
 */
var redux_1 = require('redux');
var react_router_redux_1 = require('react-router-redux');
var model_overview_reducer_1 = require('../../ModelDetails/reducers/model.overview.reducer');
var leaderboard_reducer_1 = require('../../Models/reducers/leaderboard.reducer');
var labels_reducer_1 = require('../../Configurations/reducers/labels.reducer');
var projects_reducer_1 = require('../../Projects/reducers/projects.reducer');
var services_reducer_1 = require('../../Projects/reducers/services.reducer');
var profile_reducers_1 = require('../../Profile/reducers/profile.reducers');
var deployment_reducer_1 = require('../../Deployment/reducers/deployment.reducer');
var users_reducer_1 = require('../../Users/reducers/users.reducer');
var notification_reducer_1 = require('./notification.reducer');
var global_reducer_1 = require('./global.reducer');
var collaborators_reducer_1 = require("../../Collaborators/reducers/collaborators.reducer");
var clusters_reducer_1 = require('../../Clusters/reducers/clusters.reducer');
exports.rootReducer = redux_1.combineReducers({
    global: global_reducer_1.globalReducer,
    model: model_overview_reducer_1.modelOverviewReducer,
    leaderboard: leaderboard_reducer_1.leaderboardReducer,
    labels: labels_reducer_1.labelsReducer,
    projects: projects_reducer_1.projectsReducer,
    services: services_reducer_1.servicesReducer,
    profile: profile_reducers_1.profileReducer,
    deployments: deployment_reducer_1.deploymentReducer,
    notification: notification_reducer_1.notificationReducer,
    users: users_reducer_1.usersReducer,
    routing: react_router_redux_1.routerReducer,
    collaborators: collaborators_reducer_1.collaboratorsReducer,
    clusters: clusters_reducer_1.clustersReducer
});
//# sourceMappingURL=rootReducer.js.map