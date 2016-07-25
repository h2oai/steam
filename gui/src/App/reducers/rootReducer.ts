/**
 * Created by justin on 6/25/16.
 */

import { combineReducers } from 'redux';
import { routerReducer } from 'react-router-redux';
import { modelOverviewReducer } from '../../ProjectDetails/reducers/model.overview.reducer';
import { leaderboardReducer } from '../../Models/reducers/leaderboard.reducer';
import { projectsReducer } from '../../Projects/reducers/projects.reducer';
import { servicesReducer } from '../../Projects/reducers/services.reducer';

export const rootReducer = combineReducers({
  model: modelOverviewReducer,
  leaderboard: leaderboardReducer,
  projects: projectsReducer,
  services: servicesReducer,
  routing: routerReducer
});
