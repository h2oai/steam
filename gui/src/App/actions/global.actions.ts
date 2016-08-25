import * as Remote from '../../Proxy/Proxy';
import { openNotification } from './notification.actions';

export const REQUEST_ENTITY_IDS = 'REQUEST_ENTITY_IDS';
export const RECEIVE_ENTITY_IDS = 'RECEIVE_ENTITY_IDS';

export function requestEntityIds() {
  return {
    type: REQUEST_ENTITY_IDS
  };
}

export function receiveEntityIds(response) {
  return {
    type: RECEIVE_ENTITY_IDS,
    response: response
  };
}

export function fetchEntityIds() {
  return (dispatch, getState) => {
    dispatch(requestEntityIds());

    return new Promise((resolve, reject) => {
      Remote.getAllEntityTypes((error, res) => {
        if (error) {
          openNotification('error', 'There was an error retrieving permissions list', null);
          reject(error);
        }
        dispatch(receiveEntityIds(res));
        resolve(res);
      });
    });

  };
}
