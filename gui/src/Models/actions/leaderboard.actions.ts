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
          modelName: 'DRF-1069085',
          modelType: 'drf',
          createdBy: 'tonychu',
          creationDate: '2016-06-14',
          timing: '213 ± 12ms'
        }
      },
      {
        id: 1,
        rank: 2,
        metadata: {
          modelName: 'GBT-1071707',
          modelType: 'deeplearning',
          createdBy: 'ivywang',
          creationDate: '2016-06-14',
          timing: '213 ± 12ms'
        }
      },
      {
        id: 2,
        rank: 3,
        metadata: {
          modelName: 'GBT-1071707',
          modelType: 'gbm',
          createdBy: 'marklandry',
          creationDate: '2016-06-14',
          timing: '213 ± 12ms'
        }
      },
      {
        id: 3,
        rank: 4,
        metadata: {
          modelName: 'GBT-1071707',
          modelType: 'glm',
          createdBy: 'marklandry',
          creationDate: '2016-06-14',
          timing: '213 ± 12ms'
        }
      },
      {
        id: 4,
        rank: 5,
        metadata: {
          modelName: 'GBT-1071707',
          modelType: 'naivebayes',
          createdBy: 'marklandry',
          creationDate: '2016-06-14',
          timing: '213 ± 12ms'
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
