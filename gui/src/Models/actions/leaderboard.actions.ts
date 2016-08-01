/**
 * Created by justin on 6/28/16.
 */
import { MockFetchStrategy } from '../../App/utils/FetchStrategy/MockFetchStrategy';
import { AjaxStrategy } from '../../App/utils/FetchStrategy/AjaxStrategy';
import * as Remote from '../../Proxy/Proxy';
import { BinomialModel } from '../../Proxy/Proxy';
import { RegressionModel } from '../../Proxy/Proxy';
import { MultinomialModel } from '../../Proxy/Proxy';
export const FETCH_LEADERBOARD = 'FETCH_LEADERBOARD';
export const RECEIVE_LEADERBOARD = 'RECEIVE_LEADERBOARD';
export const RECEIVE_SORT_CRITERIA = 'RECEIVE_SORT_CRITERIA';

interface Leaderboard {
  id: number,
  rank: number,
  metadata: any
}

export const requestLeaderboard = () => {
  return {
    type: FETCH_LEADERBOARD
  };
};

export function receiveLeaderboard(leaderboard) {
  return {
    type: RECEIVE_LEADERBOARD,
    leaderboard
  };
}

export function fetchLeaderboard(projectId: number) {
  return (dispatch) => {
    dispatch(requestLeaderboard());
    Remote.getModels(projectId, 0, 5, (error, res) => {
      findModelStrategy(res[0].model_category.toLowerCase())(projectId, '', '', true, 0, 5, (error, models) => {
        dispatch(receiveLeaderboard(models as BinomialModel[] | MultinomialModel[] | RegressionModel[]));
      });
    });
  };
}

export function findModelStrategy(modelCategory: string): Function {
  if (modelCategory === 'binomial') {
    return Remote.findModelsBinomial;
  } else if (modelCategory === 'multinomial') {
    return Remote.findModelsMultinomial;
  } else if (modelCategory === 'regression') {
    return Remote.findModelsRegression;
  }
}


export function receiveSortCriteria(criteria) {
  return {
    type: RECEIVE_SORT_CRITERIA,
    criteria
  };
}

export function fetchSortCriteria(modelCategory: string) {
  return (dispatch) => {
    getSortStrategy(modelCategory)((error, criteria: string[]) => {
      dispatch(receiveSortCriteria(criteria));
    });
  };
}

function getSortStrategy(modelCategory): Function {
  if (modelCategory === 'binomial') {
    return Remote.getAllMultinomialSortCriteria;
  } else if (modelCategory === 'multinomial') {
    return Remote.getAllMultinomialSortCriteria;
  } else if (modelCategory === 'regression') {
    return Remote.getAllRegressionSortCriteria;
  }
}
