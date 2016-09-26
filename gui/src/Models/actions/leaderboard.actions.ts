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
export const RECEIVE_MODEL_COUNT = 'RECEIVE_MODEL_COUNT';
export const REQUEST_DELETE_MODEL = 'REQUEST_DELETE_MODEL';
export const RECEIVE_DELETE_MODEL = 'RECEIVE_DELETE_MODEL';

interface Leaderboard {
  id: number,
  rank: number,
  metadata: any
}

export const MAX_ITEMS = 5;

export function requestDeleteModel(id) {
  return {
    type: REQUEST_DELETE_MODEL,
    id
  };
};
export function receiveDeleteModel(id, success) {
  return {
    type: RECEIVE_DELETE_MODEL,
    id,
    success
  };
};

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

export function receiveModelCount(count: number) {
  return {
    type: RECEIVE_MODEL_COUNT,
    count
  };
}

export function findModelsCount(projectId: number) {
  return (dispatch) => {
    return Remote.findModelsCount(projectId, (error, count) => {
      if (error) {
        dispatch(openNotification('error', error.toString(), null));
        return;
      }
      dispatch(receiveModelCount(count));
    });
  };
}

export function deleteModel(modelId: number, fetchLeaderboard: Function) {
  return(dispatch) => {
    dispatch(requestDeleteModel(modelId));
    Remote.deleteModel(modelId, (error) => {
      if (error) {
        dispatch(openNotification('error', error.toString(), null));
        dispatch(receiveDeleteModel(modelId, false));
        return;
      }
      dispatch(receiveDeleteModel(modelId, true));
      fetchLeaderboard();
    });
  };
}
