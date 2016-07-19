/**
 * Created by justin on 6/25/16.
 */

import { combineReducers } from 'redux';
import { routerReducer } from 'react-router-redux';
import { navigationReducer } from '../../Navigation/components/Navigation/reducers/navigation.reducer';
import { modelOverviewReducer } from '../../ProjectDetails/reducers/model.overview.reducer';
import { leaderboardReducer } from '../../Models/reducers/leaderboard.reducer';
import { projectsReducer } from '../../Projects/reducers/projects.reducer';

export const rootReducer = combineReducers({
  navigation: navigationReducer,
  modelOverview: modelOverviewReducer,
  leaderboard: leaderboardReducer,
  projects: projectsReducer,
  routing: routerReducer
});
