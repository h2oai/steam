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
 * Created by justin on 8/2/16.
 */
var Remote = require('../../Proxy/Proxy');
var notification_actions_1 = require('../../App/actions/notification.actions');
var Notification_1 = require('../../App/components/Notification');
exports.UPLOADING_PACKAGE = 'UPLOADING_PACKAGE_COMPONENT';
exports.FINISH_UPLOADING_PACKAGE_COMPONENT = 'FINISH_UPLOADING_PACKAGE_COMPONENT';
exports.RECEIVE_PACKAGES = 'RECEIVE_PACKAGES';
function uploadingPackage() {
    return {
        type: exports.UPLOADING_PACKAGE
    };
}
exports.uploadingPackage = uploadingPackage;
function finishUploadingPackageComponent() {
    return {
        type: exports.FINISH_UPLOADING_PACKAGE_COMPONENT
    };
}
exports.finishUploadingPackageComponent = finishUploadingPackageComponent;
function receivePackages(packages) {
    return {
        type: exports.RECEIVE_PACKAGES,
        packages: packages
    };
}
exports.receivePackages = receivePackages;
function uploadPackage(projectId, packageName, form) {
    return function (dispatch) {
        dispatch(uploadingPackage());
        var formFiles = form.querySelectorAll('input[type="file"]');
        Remote.createPackage(projectId, packageName, function (error) {
            var data;
            var requests = [];
            var main = null;
            var _loop_1 = function(i) {
                var _loop_2 = function(j) {
                    data = new FormData();
                    if (formFiles[i].name === 'selectMain') {
                        main = formFiles[i].files[j].name;
                    }
                    data.append('file', formFiles[i].files[j]);
                    if (error) {
                        dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, "Load Error", error.toString(), null));
                        return { value: void 0 };
                    }
                    requests.push(fetch("/upload?type=file&project-id=" + projectId + "&package-name=" + packageName + "&relative-path=", {
                        credentials: 'include',
                        method: 'post',
                        body: data
                    }).then(function () {
                        Remote.setAttributesForPackage(projectId, packageName, JSON.stringify({ main: formFiles[i].files[j].name }), function (error) {
                            if (error) {
                                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, "Load Error", error, null));
                                return;
                            }
                        });
                        dispatch(finishUploadingPackageComponent());
                        dispatch(fetchPackages(projectId));
                    }));
                };
                for (var j = 0; j < formFiles[i].files.length; j++) {
                    var state_1 = _loop_2(j);
                    if (typeof state_1 === "object") return state_1;
                }
            };
            for (var i = 0; i < formFiles.length; i++) {
                var state_2 = _loop_1(i);
                if (typeof state_2 === "object") return state_2.value;
            }
            Promise.all(requests).then(function () {
                Remote.setAttributesForPackage(projectId, packageName, JSON.stringify({ main: main }), function (error) {
                    if (error) {
                        dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error, null));
                        return;
                    }
                });
                dispatch(finishUploadingPackageComponent());
                dispatch(fetchPackages(projectId));
            });
        });
    };
}
exports.uploadPackage = uploadPackage;
function fetchPackages(projectId) {
    return function (dispatch) {
        Remote.getPackages(projectId, function (error, res) {
            if (error) {
                dispatch(notification_actions_1.openNotification(Notification_1.NotificationType.Error, 'Load Error', error.toString(), null));
                return;
            }
            dispatch(receivePackages(res));
        });
    };
}
exports.fetchPackages = fetchPackages;
//# sourceMappingURL=deployment.actions.js.map