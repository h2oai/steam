/**
 * Created by justin on 6/28/16.
 */
import { MockFetchStrategy } from '../../App/utils/FetchStrategy/MockFetchStrategy';
import { AjaxStrategy } from '../../App/utils/FetchStrategy/AjaxStrategy';
import * as Remote from '../../Proxy/Proxy';
export const FETCH_LEADERBOARD = 'FETCH_LEADERBOARD';
export const RECEIVE_LEADERBOARD = 'RECEIVE_LEADERBOARD';

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

export function fetchLeaderboard() {
  return (dispatch) => {
    dispatch(requestLeaderboard());
    let leaderboard: Leaderboard[] = [
      {
        id: 0,
        rank: 1,
        metadata: {
          modelName: 'glm-aaed8e67-2a86-476f-b1ab-b8232c0c553e',
          modelType: 'glm',
          createdBy: 'laurendiperna',
          creationDate: '2016-07-18',
          timing: '213 ± 12ms',
          auc: 0.847345
        }
      },
      {
        id: 1,
        rank: 2,
        metadata: {
          modelName: 'gbm-007328bb-1d09-4dcd-b659-7f58acfa4da4',
          modelType: 'gbm',
          createdBy: 'navdeep',
          creationDate: '2016-07-18',
          timing: '213 ± 12ms',
          auc: 0.846126
        }
      },
      {
        id: 2,
        rank: 3,
        metadata: {
          modelName: 'deeplearning-ef0ea82a-acf4-4392-9b91-44d583fb826f',
          modelType: 'deeplearning',
          createdBy: 'marklandry',
          creationDate: '2016-07-18',
          timing: '213 ± 12ms',
          auc: 0.837695
        }
      },
      {
        id: 3,
        rank: 4,
        metadata: {
          modelName: 'drf-5c77a28d-04e5-46f6-964a-a3b7f80ba1a6',
          modelType: 'drf',
          createdBy: 'laurendiperna',
          creationDate: '2016-07-18',
          timing: '213 ± 12ms',
          auc: 0.835059
        }
      },
      {
        id: 4,
        rank: 5,
        metadata: {
          modelName: 'naivebayes-940f6d45-54e4-41f9-aa34-8562e7f9390f',
          modelType: 'naivebayes',
          createdBy: 'marklandry',
          creationDate: '2016-07-18',
          timing: '213 ± 12ms',
          auc: 0.819140
        }
      }
    ];
    // Remote.Proxy.getModels(0, 5, (error, res) => {
    //   console.log(error, res);
    // });
    new MockFetchStrategy().request(dispatch, {
      callback: receiveLeaderboard,
      data: leaderboard
    });
  };
}
