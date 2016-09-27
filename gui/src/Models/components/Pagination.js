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
 * Created by justin on 6/27/16.
 */
var React = require('react');
var classNames = require('classnames');
require('../styles/pagination.scss');
var leaderboard_actions_1 = require('../actions/leaderboard.actions');
var Pagination = (function (_super) {
    __extends(Pagination, _super);
    function Pagination() {
        _super.apply(this, arguments);
    }
    Pagination.prototype.render = function () {
        return (React.createElement("div", {className: "pagination-container"}, React.createElement("span", {onClick: this.props.onPageBack.bind(this)}, React.createElement("i", {className: classNames('fa fa-caret-left', { disabled: this.props.currentPage === 0 })})), React.createElement("span", {className: "page-info"}, ((this.props.currentPage + 1) * leaderboard_actions_1.MAX_ITEMS) - (leaderboard_actions_1.MAX_ITEMS - 1), " - ", (this.props.currentPage + 1) * leaderboard_actions_1.MAX_ITEMS < this.props.count ? (this.props.currentPage + 1) * leaderboard_actions_1.MAX_ITEMS : this.props.count, " of ", this.props.count, " models"), React.createElement("span", {onClick: this.props.onPageForward.bind(this)}, React.createElement("i", {className: classNames('fa fa-caret-right', { disabled: (this.props.currentPage + 1) * leaderboard_actions_1.MAX_ITEMS >= this.props.count })}))));
    };
    return Pagination;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = Pagination;
//# sourceMappingURL=Pagination.js.map