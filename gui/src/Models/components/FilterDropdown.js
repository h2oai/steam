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
 * Created by justin on 7/28/16.
 */
var React = require('react');
var ReactDOM = require('react-dom');
var classNames = require('classnames');
var FilterIcon_1 = require('./FilterIcon');
require('../styles/filterdropdown.scss');
var FILTER_MAP = {
    mean_residual_deviance: 'MRD',
    r_squared: React.createElement("span", null, "R", React.createElement("sup", null, "2")),
    mse: 'MSE',
    logloss: 'LogLoss',
    auc: 'AUC',
    gini: 'Gini'
};
var FilterDropdown = (function (_super) {
    __extends(FilterDropdown, _super);
    function FilterDropdown() {
        _super.call(this);
        this.state = {
            open: false,
            sortBy: null,
            orderBy: 'asc'
        };
        this.bodyClickHandler = this.bodyClickHandler.bind(this);
    }
    FilterDropdown.prototype.componentWillMount = function () {
        document.body.addEventListener('click', this.bodyClickHandler);
    };
    FilterDropdown.prototype.componentWillUnmount = function () {
        document.body.removeEventListener('click', this.bodyClickHandler);
    };
    FilterDropdown.prototype.bodyClickHandler = function (event) {
        if (!ReactDOM.findDOMNode(this.refs.filterDropdown).contains(event.target) && !ReactDOM.findDOMNode(this.refs.filterDropdownInvoker).contains(event.target)) {
            this.setState({
                open: false
            });
        }
    };
    FilterDropdown.prototype.openDropdown = function () {
        this.setState({
            open: !this.state.open
        });
    };
    FilterDropdown.prototype.selectSort = function (selection) {
        this.setState({
            sortBy: selection
        });
        this.props.onFilter({
            sortBy: selection,
            orderBy: this.state.orderBy
        });
    };
    FilterDropdown.prototype.selectOrder = function (selection) {
        this.setState({
            orderBy: selection
        });
        this.props.onFilter({
            sortBy: this.state.sortBy,
            orderBy: selection
        });
    };
    FilterDropdown.prototype.render = function () {
        var _this = this;
        if (this.props.sortCriteria === null) {
            return React.createElement("div", null);
        }
        return (React.createElement("div", {className: "filter-dropdown"}, React.createElement("button", {ref: "filterDropdownInvoker", className: classNames('filter-dropdown-invoker', { open: this.state.open }), onClick: this.openDropdown.bind(this)}, React.createElement(FilterIcon_1.default, null)), React.createElement("div", {ref: "filterDropdown", className: classNames('filter-dropdown-menu', { open: this.state.open })}, React.createElement("div", {className: "filter-option"}, React.createElement("div", {className: "filter-labels"}, "SORT BY"), React.createElement("ul", null, this.props.sortCriteria.map(function (criteria, i) {
            return React.createElement("li", {key: i, onClick: _this.selectSort.bind(_this, criteria), className: classNames({ selected: _this.state.sortBy === criteria })}, FILTER_MAP[criteria], " ", _this.state.sortBy === criteria ?
                React.createElement("i", {className: "fa fa-check"}) : null);
        }))), React.createElement("div", {className: "filter-option"}, React.createElement("div", {className: "filter-labels"}, "ORDER"), React.createElement("ul", null, React.createElement("li", {onClick: this.selectOrder.bind(this, 'asc'), className: classNames({ selected: this.state.orderBy === 'asc' })}, "ASC ", this.state.orderBy === 'asc' ?
            React.createElement("i", {className: "fa fa-check"}) : null), React.createElement("li", {onClick: this.selectOrder.bind(this, 'des'), className: classNames({ selected: this.state.orderBy === 'des' })}, "DES ", this.state.orderBy === 'des' ?
            React.createElement("i", {className: "fa fa-check"}) : null))))));
    };
    return FilterDropdown;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = FilterDropdown;
//# sourceMappingURL=FilterDropdown.js.map