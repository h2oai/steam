import { combineReducers } from 'redux';
import { statisticsReducer } from './statisticsReducer';
import { modelReducer } from './modelReducer';

const rootReducer = combineReducers({
  statistics: statisticsReducer,
  model: modelReducer
});

export default rootReducer;