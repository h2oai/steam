/**
 * Created by justin on 7/22/16.
 */
import * as Remote from '../../Proxy/proxy';

export const REQUEST_SERVICES = 'REQUEST_SERVICES';
export const RECEIVE_SERVICES = 'RECEIVE_SERVICES';
export const STOPPED_SERVICE = 'STOPPED_SERVICE';
export const requestServices = () => {
  return {
    type: REQUEST_SERVICES
  };
};

export function receiveServices(services) {
  return {
    type: RECEIVE_SERVICES,
    services
  };
}

export const stoppedService = () => {
  return {
    type: STOPPED_SERVICE  
  };
};

export function fetchServices() {
  return (dispatch) => {
    dispatch(requestServices());
    Remote.getScoringServices(0, 5, (error, res) => {
      dispatch(receiveServices(res));
    });
  };
}

export function killService(serviceId: number) {
  return (dispatch) => {
    Remote.stopScoringService(serviceId, (error) => {
      dispatch(fetchServices());
    });
  };
}
