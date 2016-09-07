/**
 * Created by justin on 7/22/16.
 */
import * as Remote from '../../Proxy/Proxy';

export const REQUEST_ALL_SERVICES = 'REQUEST_ALL_SERVICES';
export const RECEIVE_ALL_SERVICES = 'RECEIVE_ALL_SERVICES';
export const REQUEST_SERVICES_FOR_PROJECT = 'REQUEST_SERVICES_FOR_PROJECT';
export const RECEIVE_SERVICES_FOR_PROJECT = 'RECEIVE_SERVICES_FOR_PROJECT';
export const STOPPED_SERVICE = 'STOPPED_SERVICE';

export const requestAllServices = () => {
  return {
    type: REQUEST_ALL_SERVICES
  };
};
export function receiveAllServices(services) {
  return {
    type: RECEIVE_ALL_SERVICES,
    services
  };
}

export const requestServicesForProject = () => {
  return {
    type: REQUEST_SERVICES_FOR_PROJECT
  };
};
export function receiveServicesForProject(services) {
  return {
    type: RECEIVE_SERVICES_FOR_PROJECT,
    services
  };
}

export const stoppedService = () => {
  return {
    type: STOPPED_SERVICE
  };
};

export function fetchAllServices() {
  return (dispatch) => {
    dispatch(requestAllServices());
    Remote.getServices(0, 1000, (error, res) => {
      dispatch(receiveAllServices(res));
    });
  };
}

export function fetchServicesForProject(projectId: number) {
  return (dispatch) => {
    dispatch(requestServicesForProject());
    Remote.getServicesForProject(projectId, 0, 1000, (error, res) => {
      dispatch(receiveServicesForProject(res));
    });
  };
}

export function killService(serviceId: number) {
  return (dispatch) => {
    Remote.stopService(serviceId, (error) => {
      dispatch(fetchAllServices());
    });
  };
}
