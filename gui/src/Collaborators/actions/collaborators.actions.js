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
var Remote = require('../../Proxy/Proxy');
var _ = require('lodash');
var global_actions_1 = require('../../App/actions/global.actions');
var notification_actions_1 = require('../../App/actions/notification.actions');
var Notification_1 = require('../../App/components/Notification');
exports.REQUEST_MEMBERS_FOR_PROJECT = 'REQUEST_MEMBERS_FOR_PROJECT';
exports.RECEIVE_MEMBERS_FOR_PROJECT = 'RECEIVE_MEMBERS_FOR_PROJECT';
exports.REQUEST_LABELS_FOR_PROJECT = 'REQUEST_LABELS_FOR_PROJECT';
exports.RECEIVE_LABELS_FOR_PROJECT = 'RECEIVE_LABELS_FOR_PROJECT';
function requestMembersForProject() {
    return {
        type: exports.REQUEST_MEMBERS_FOR_PROJECT
    };
}
exports.requestMembersForProject = requestMembersForProject;
function receiveMembersForProject(response) {
    return {
        type: exports.RECEIVE_MEMBERS_FOR_PROJECT,
        members: response
    };
}
exports.receiveMembersForProject = receiveMembersForProject;
function requestLabelsForProject() {
    return {
        type: exports.REQUEST_LABELS_FOR_PROJECT
    };
}
exports.requestLabelsForProject = requestLabelsForProject;
function receiveLabelsForProject(response) {
    return {
        type: exports.RECEIVE_LABELS_FOR_PROJECT,
        labels: response
    };
}
exports.receiveLabelsForProject = receiveLabelsForProject;
function _fetchMembersForProject(dispatch, state) {
    Remote.getIdentitiesForEntity(state.global.entityIds.project, state.projects.project.id, function (error, res) {
        if (error) {
            notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null);
            return;
        }
        dispatch(receiveMembersForProject(res));
    });
}
function fetchMembersForProject() {
    return function (dispatch, getState) {
        dispatch(requestMembersForProject());
        var state = getState();
        if (_.isEmpty(state.global.entityIds)) {
            dispatch(global_actions_1.fetchEntityIds()).then(function () {
                state = getState();
                _fetchMembersForProject(dispatch, state);
            });
        }
        else {
            _fetchMembersForProject(dispatch, state);
        }
    };
}
exports.fetchMembersForProject = fetchMembersForProject;
function _fetchLabelsForProject(dispatch, state) {
    Remote.getLabelsForProject(state.projects.project.id, function (error, labels) {
        if (error) {
            notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null);
            return;
        }
        var identityPromises = [];
        var toReturn = [];
        var _loop_1 = function(label) {
            identityPromises.push(new Promise(function (resolve, reject) {
                Remote.getIdentitiesForEntity(state.global.entityIds.label, label.id, function (identitiesError, identitiesRes) {
                    if (identitiesError) {
                        notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', identitiesError.toString(), null);
                        reject(identitiesError.toString());
                        return;
                    }
                    toReturn.push(_.assign({}, label, { identities: identitiesRes }));
                    resolve();
                });
            }));
        };
        for (var _i = 0, labels_1 = labels; _i < labels_1.length; _i++) {
            var label = labels_1[_i];
            _loop_1(label);
        }
        Promise.all(identityPromises).then(function (results) {
            dispatch(receiveLabelsForProject(toReturn));
        });
    });
}
function fetchLabelsForProject() {
    return function (dispatch, getState) {
        dispatch(requestLabelsForProject());
        var state = getState();
        if (_.isEmpty(state.global.entityIds)) {
            dispatch(global_actions_1.fetchEntityIds()).then(function () {
                state = getState();
                _fetchLabelsForProject(dispatch, state);
            });
        }
        else {
            _fetchLabelsForProject(dispatch, state);
        }
    };
}
exports.fetchLabelsForProject = fetchLabelsForProject;
//# sourceMappingURL=collaborators.actions.js.map