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
"use strict";
/**
 * Created by justin on 7/22/16.
 */
var Remote = require('../../Proxy/Proxy');
exports.REQUEST_ALL_SERVICES = 'REQUEST_ALL_SERVICES';
exports.RECEIVE_ALL_SERVICES = 'RECEIVE_ALL_SERVICES';
exports.REQUEST_SERVICES_FOR_PROJECT = 'REQUEST_SERVICES_FOR_PROJECT';
exports.RECEIVE_SERVICES_FOR_PROJECT = 'RECEIVE_SERVICES_FOR_PROJECT';
exports.STOPPED_SERVICE = 'STOPPED_SERVICE';
exports.requestAllServices = function () {
    return {
        type: exports.REQUEST_ALL_SERVICES
    };
};
function receiveAllServices(services) {
    return {
        type: exports.RECEIVE_ALL_SERVICES,
        services: services
    };
}
exports.receiveAllServices = receiveAllServices;
exports.requestServicesForProject = function () {
    return {
        type: exports.REQUEST_SERVICES_FOR_PROJECT
    };
};
function receiveServicesForProject(services) {
    return {
        type: exports.RECEIVE_SERVICES_FOR_PROJECT,
        services: services
    };
}
exports.receiveServicesForProject = receiveServicesForProject;
exports.stoppedService = function () {
    return {
        type: exports.STOPPED_SERVICE
    };
};
function fetchAllServices() {
    return function (dispatch) {
        dispatch(exports.requestAllServices());
        Remote.getServices(0, 1000, function (error, res) {
            dispatch(receiveAllServices(res));
        });
    };
}
exports.fetchAllServices = fetchAllServices;
function fetchServicesForProject(projectId) {
    return function (dispatch) {
        dispatch(exports.requestServicesForProject());
        Remote.getServicesForProject(projectId, 0, 1000, function (error, res) {
            dispatch(receiveServicesForProject(res));
        });
    };
}
exports.fetchServicesForProject = fetchServicesForProject;
function killService(serviceId, projectId) {
    return function (dispatch) {
        Remote.stopService(serviceId, function (error) {
            if (projectId) {
                dispatch(fetchServicesForProject(projectId));
            }
            else {
                dispatch(fetchAllServices());
            }
        });
    };
}
exports.killService = killService;
//# sourceMappingURL=services.actions.js.map