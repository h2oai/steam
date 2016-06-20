import * as _ from 'lodash';

import {
  FETCH_STATISTICS,
  REQUEST_STATISTICS,
  RECEIVE_STATISTICS
} from '../actions/statisticsActions';

let initialState = {};

export const statisticsReducer = (state = initialState, action) => {
  switch(action.type) {
    case FETCH_STATISTICS:
      return _.assign({}, state, {
        isFetching: true
      });
    case RECEIVE_STATISTICS:
      return _.assign({}, state, {
        data: action.data,
        isFetching: false
      });
    default:
      return state;
  }
};

function requestStatistics() {
  return {
    type: REQUEST_STATISTICS
  }
}

function receiveStatistics(data) {
  return {
    type: RECEIVE_STATISTICS,
    data
  }
}

const exampleStatisticsResponse = {"lastTime":0,"lastTimeAgoMs":1466191381324,"pythonget":{"averageAfterWarmupTime":0.0,"totalTimeSquaredMs":0.0,"lastMs":0.0,"averageTime":0.0,"totalTimeMs":0.0,"count":0.0,"warmupTimeSquaredMs":0.0,"warmupTimeMs":0.0},"post":{"averageAfterWarmupTime":0.0,"totalTimeSquaredMs":0.0,"lastMs":0.0,"averageTime":0.0,"totalTimeMs":0.0,"count":0.0,"warmupTimeSquaredMs":0.0,"warmupTimeMs":0.0},"get":{"averageAfterWarmupTime":0.0,"totalTimeSquaredMs":0.0,"lastMs":0.0,"averageTime":0.0,"totalTimeMs":0.0,"count":0.0,"warmupTimeSquaredMs":0.0,"warmupTimeMs":0.0},"prediction":{"averageAfterWarmupTime":0.0,"totalTimeSquaredMs":0.0,"lastMs":0.0,"averageTime":0.0,"totalTimeMs":0.0,"count":0.0,"warmupTimeSquaredMs":0.0,"warmupTimeMs":0.0},"startTime":1466191337349,"pythonpost":{"averageAfterWarmupTime":0.0,"totalTimeSquaredMs":0.0,"lastMs":0.0,"averageTime":0.0,"totalTimeMs":0.0,"count":0.0,"warmupTimeSquaredMs":0.0,"warmupTimeMs":0.0},"upTimeMs":43975,"lastTimeUTC":"","startTimeUTC":"2016-06-17 19:22:17 UTC","warmUpCount":5};

export const fetchStatistics = (): Function => {
  return function(dispatch) {
    dispatch(requestStatistics());
    return dispatch(receiveStatistics(exampleStatisticsResponse));
    // return fetch('/stats')
    //   .then(response => response.json())
    //   .then(json => dispatch(receiveStatistics(exampleStatisticsResponse)));
  };
};