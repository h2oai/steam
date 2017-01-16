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
 * Created by justin on 7/22/16.
 */
import * as Remote from '../../Proxy/Proxy';
import { openNotification } from '../../App/actions/notification.actions';
import { NotificationType } from '../../App/components/Notification';

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

export function killService(serviceId: number, projectId: number) {
  return (dispatch) => {
    Remote.stopService(serviceId, (error) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, 'Create Error', error.toString(), null));
      } else {
        if (projectId) {
          dispatch(fetchServicesForProject(projectId));
        } else {
          dispatch(fetchAllServices());
        }
      }
    });
  };
}
