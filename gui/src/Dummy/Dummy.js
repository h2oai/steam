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
var __extends = (this && this.__extends) || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
};
/**
 * Created by Jeff Fohl <jfohl@h2o.ai> on 7/17/16.
 * Dummy component for mocking
 */
var React = require('react');
;
var Dummy = (function (_super) {
    __extends(Dummy, _super);
    function Dummy() {
        _super.call(this);
    }
    Dummy.prototype.render = function () {
        return (React.createElement("div", null, React.createElement("h1", null, "Content to come.")));
    };
    return Dummy;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = Dummy;
;
//# sourceMappingURL=Dummy.js.map