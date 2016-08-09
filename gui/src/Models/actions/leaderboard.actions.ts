/**
 * Created by justin on 6/28/16.
 */
import { MockFetchStrategy } from '../../App/utils/FetchStrategy/MockFetchStrategy';
import { AjaxStrategy } from '../../App/utils/FetchStrategy/AjaxStrategy';
import * as Remote from '../../Proxy/Proxy';
import { BinomialModel } from '../../Proxy/Proxy';
import { RegressionModel } from '../../Proxy/Proxy';
import { MultinomialModel } from '../../Proxy/Proxy';
import { openNotification } from '../../App/actions/notification.actions';
export const FETCH_LEADERBOARD = 'FETCH_LEADERBOARD';
export const RECEIVE_LEADERBOARD = 'RECEIVE_LEADERBOARD';
export const RECEIVE_SORT_CRITERIA = 'RECEIVE_SORT_CRITERIA';

interface Leaderboard {
  id: number,
  rank: number,
  metadata: any
}

export const MAX_ITEMS = 5;

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

export function fetchLeaderboard(projectId: number, modelCategory: string, name: string, sortBy: string, ascending: boolean, offset: number) {
  return (dispatch) => {
    dispatch(requestLeaderboard());
    findModelStrategy(modelCategory.toLowerCase())(projectId, name, sortBy || '', ascending || false, offset, MAX_ITEMS, (error, models) => {
      if (error) {
        dispatch(openNotification('error', error.toString(), null));
        return;
      }
      dispatch(receiveLeaderboard(models as BinomialModel[] | MultinomialModel[] | RegressionModel[]));
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
      if (error) {
        dispatch(openNotification('error', error.toString(), null));
        return;
      }
      dispatch(receiveSortCriteria(criteria));
    });
  };
}

function getSortStrategy(modelCategory): Function {
  if (modelCategory === 'binomial') {
    return Remote.getAllBinomialSortCriteria;
  } else if (modelCategory === 'multinomial') {
    return Remote.getAllMultinomialSortCriteria;
  } else if (modelCategory === 'regression') {
    return Remote.getAllRegressionSortCriteria;
  }
}

export function linkLabelWithModel(labelId: number, modelId: number) {
  return (dispatch) => {
    return new Promise((resolve, reject) => {
      Remote.linkLabelWithModel(labelId, modelId, (error) => {
        if (error) {
          dispatch(openNotification('error', error.toString(), null));
          reject(error);
          return;
        }
        resolve();
      });
    });
  };
}

export function unlinkLabelFromModel(labelId: number, modelId: number) {
  return (dispatch) => {
    return new Promise((resolve, reject) => {
      Remote.unlinkLabelFromModel(labelId, modelId, (error) => {
        if (error) {
          dispatch(openNotification('error', error.toString(), null));
          reject(error);
          return;
        }
        resolve();
      });
    });
  };
}
