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
