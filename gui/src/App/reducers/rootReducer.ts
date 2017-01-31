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
 * Created by justin on 6/25/16.
 */

import { combineReducers } from 'redux';
import { routerReducer } from 'react-router-redux';
import { modelOverviewReducer } from '../../ModelDetails/reducers/model.overview.reducer';
import { leaderboardReducer } from '../../Models/reducers/leaderboard.reducer';
import { labelsReducer } from '../../Configurations/reducers/labels.reducer';
import { projectsReducer } from '../../Projects/reducers/projects.reducer';
import { servicesReducer } from '../../Projects/reducers/services.reducer';
import { profileReducer } from '../../Profile/reducers/profile.reducers';
import { deploymentReducer } from '../../Deployment/reducers/deployment.reducer';
import { usersReducer } from '../../Users/reducers/users.reducer';
import { notificationReducer } from './notification.reducer';
import { globalReducer } from './global.reducer';
import { collaboratorsReducer } from "../../Collaborators/reducers/collaborators.reducer";
import { clustersReducer } from '../../Clusters/reducers/clusters.reducer';

export const rootReducer = combineReducers({
  global: globalReducer,
  model: modelOverviewReducer,
  leaderboard: leaderboardReducer,
  labels: labelsReducer,
  projects: projectsReducer,
  services: servicesReducer,
  profile: profileReducer,
  deployments: deploymentReducer,
  notification: notificationReducer,
  users: usersReducer,
  routing: routerReducer,
  collaborators: collaboratorsReducer,
  clusters: clustersReducer
});
