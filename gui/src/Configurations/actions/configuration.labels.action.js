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
 * Created by Jeff Fohl <jfohl@h2o.ai> on 7/30/16.
 */
var Remote = require('../../Proxy/Proxy');
exports.FETCH_LABELS = 'FETCH_LABELS';
exports.RECEIVE_LABELS = 'RECEIVE_LABELS';
exports.CREATE_LABEL = 'CREATE_LABEL';
exports.RECEIVE_CREATE_LABEL = 'RECEIVE_CREATE_LABEL';
exports.UPDATE_LABEL = 'UPDATE_LABEL';
exports.RECEIVE_UPDATE_LABEL = 'RECEIVE_UPDATE_LABEL';
exports.DELETE_LABEL = 'DELETE_LABEL';
exports.RECEIVE_DELETE_LABEL = 'RECEIVE_DELETE_LABEL';
exports.requestLabels = function () {
    return {
        type: exports.FETCH_LABELS
    };
};
function fetchLabels(projectId) {
    return function (dispatch) {
        dispatch(exports.requestLabels());
        Remote.getLabelsForProject(projectId, function (error, res) {
            dispatch(receiveLabels(res, projectId));
        });
    };
}
exports.fetchLabels = fetchLabels;
function receiveLabels(labels, projectId) {
    return {
        type: exports.RECEIVE_LABELS,
        projectId: projectId,
        labels: labels
    };
}
exports.receiveLabels = receiveLabels;
function createLabel(projectId, name, description) {
    return function (dispatch) {
        return new Promise(function (resolve, reject) {
            Remote.createLabel(projectId, name, description, function (error, res) {
                if (error) {
                    reject(error);
                    return;
                }
                dispatch(receiveCreateLabel(res, projectId, name, description));
                resolve(res);
            });
        });
    };
}
exports.createLabel = createLabel;
function receiveCreateLabel(id, projectId, name, description) {
    return {
        type: exports.RECEIVE_CREATE_LABEL,
        projectId: projectId,
        label: {
            id: id,
            name: name,
            description: description
        }
    };
}
exports.receiveCreateLabel = receiveCreateLabel;
function updateLabel(labelId, projectId, name, description) {
    return function (dispatch) {
        return new Promise(function (resolve, reject) {
            Remote.updateLabel(labelId, name, description, function (error) {
                if (error) {
                    reject(error);
                    return;
                }
                dispatch(receiveUpdateLabel(labelId, projectId, name, description));
                resolve();
            });
        });
    };
}
exports.updateLabel = updateLabel;
function receiveUpdateLabel(id, projectId, name, description) {
    return {
        type: exports.RECEIVE_UPDATE_LABEL,
        projectId: projectId,
        label: {
            id: id,
            name: name,
            description: description
        }
    };
}
exports.receiveUpdateLabel = receiveUpdateLabel;
function deleteLabel(labelId) {
    return function (dispatch) {
        return new Promise(function (resolve, reject) {
            Remote.deleteLabel(labelId, function (error) {
                if (error) {
                    reject(error);
                    return;
                }
                //dispatch(receiveDeleteLabel(labelId));
                resolve();
            });
        });
    };
}
exports.deleteLabel = deleteLabel;
function receiveDeleteLabel(labelId) {
    return {
        type: exports.RECEIVE_DELETE_LABEL
    };
}
exports.receiveDeleteLabel = receiveDeleteLabel;
//# sourceMappingURL=configuration.labels.action.js.map