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
'use strict';
var $ = require('jquery');
var _rpcId = 0;
function nextId() {
    return ++_rpcId;
}
function _invoke(settings, go) {
    var p = $.ajax(settings);
    p.done(function (data, status, xhr) {
        if (data.error) {
            return go(new Error(data.error), null);
        }
        go(null, data.result);
    });
    p.fail(function (xhr, status, error) {
        // XXX handle 401
        var res = xhr.responseJSON;
        if (res && res.error) {
            return go(new Error(res.error), null);
        }
        // special-case net::ERR_CONNECTION_REFUSED
        if (status === 'error' && xhr.status === 0) {
            return go(new Error("Could not connect to Steam. The server is currently unresponsive."), null);
        }
        go(new Error("HTTP connection failure: status=" + status + ", code=" + xhr.status + ", error=" + (error ? error : '?')), null);
    });
}
function invoke(method, param, headers, go) {
    var req = {
        method: "web." + method,
        params: [param],
        id: nextId()
    };
    var settings = {
        url: "/web",
        type: "POST",
        data: JSON.stringify(req),
        contentType: "application/json; charset=utf-8",
        dataType: "json"
    };
    if (headers) {
        settings.headers = headers;
    }
    _invoke(settings, go);
}
function upload(formData, go) {
    var settings = {
        url: "/upload",
        type: "POST",
        data: formData,
        cache: false,
        contentType: false,
        processData: false
    };
    _invoke(settings, go);
}
exports.upload = upload;
function Call(method, param, go) {
    invoke(method, param, null, go);
}
exports.Call = Call;
function authenticate(username, password, go) {
    var headers = {
        "Authorization": "Custom " + username + " " + password
    };
    invoke("Ping", { status: true }, headers, go);
}
exports.authenticate = authenticate;
//# sourceMappingURL=xhr.js.map