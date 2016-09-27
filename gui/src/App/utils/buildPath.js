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
/*
* Takes a path and converts parameters to their appropriate
* @function
* @param {string} path
* @param {object} params
* @returns {string} Returns a string with the filled in params.
*/
exports.buildPath = function (path, params) {
    var parts = path.split('/');
    parts = parts.map(function (part) {
        if (part[0] === ':') {
            var newPart = part.slice(1, part.length);
            if (params[newPart]) {
                return params[newPart];
            }
        }
        return part;
    });
    return '/' + parts.join('/');
};
//# sourceMappingURL=buildPath.js.map