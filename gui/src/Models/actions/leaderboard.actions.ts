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

export function fetchLeaderboard(projectId: number) {
  return (dispatch) => {
    dispatch(requestLeaderboard());
    Remote.getModels(projectId, 0, 5, (error, res) => {
      dispatch(receiveLeaderboard(res));
    });
  };
}

export function filterModels(projectId: number, namePart: string, sortBy: string, ascending: boolean) {
  return (dispatch) => {
    Remote.findModelsRegression(projectId, namePart, sortBy, ascending, 0, 5, (error, res) => {
      console.log(res);
    });
  };
}
