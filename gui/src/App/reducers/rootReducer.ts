/**
 * Created by justin on 6/25/16.
 */

import { combineReducers } from 'redux';
import { routerReducer } from 'react-router-redux';
import { modelOverviewReducer } from '../../ProjectDetails/reducers/model.overview.reducer';
import { leaderboardReducer } from '../../Models/reducers/leaderboard.reducer';

export const rootReducer = combineReducers({
  modelOverview: modelOverviewReducer,
  leaderboard: leaderboardReducer,
  routing: routerReducer
});
