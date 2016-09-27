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
 * Created by justin on 7/12/16.
 */
var React = require('react');
var ReactDOM = require('react-dom');
var react_router_1 = require('react-router');
var $ = require('jquery');
var react_router_2 = require('react-router');
var Panel_1 = require('./Panel');
var PageHeader_1 = require('./PageHeader');
var ProgressBar_1 = require('./ProgressBar');
require('../styles/newprojectstep3.scss');
var NewProjectStep3 = (function (_super) {
    __extends(NewProjectStep3, _super);
    function NewProjectStep3() {
        _super.call(this);
        var jobs = [
            {
                name: 'DRF-1070196',
                project: 'Churn Prediction',
                author: 'Mark Landry',
                startTime: new Date().getTime()
            },
            {
                name: 'DRF-1070196',
                project: 'Churn Prediction',
                author: 'Mark Landry',
                startTime: new Date().getTime()
            },
            {
                name: 'DRF-1070196',
                project: 'Churn Prediction',
                author: 'Mark Landry',
                startTime: new Date().getTime()
            },
            {
                name: 'DRF-1070196',
                project: 'Churn Prediction',
                author: 'Mark Landry',
                startTime: new Date().getTime()
            }
        ];
        this.state = {
            jobs: jobs
        };
    }
    NewProjectStep3.prototype.onComplete = function (progressBar) {
        var node = ReactDOM.findDOMNode(progressBar);
        $(node).addClass('progress-button');
        $(node).find('.progress-counter').text('Completed');
    };
    NewProjectStep3.prototype.onClick = function () {
        react_router_2.hashHistory.push('/models/0');
    };
    NewProjectStep3.prototype.render = function () {
        var _this = this;
        return (React.createElement("div", {className: "new-project-step-3"}, React.createElement(PageHeader_1.default, null, "GOOD WORK!"), React.createElement("div", {className: "sub-title"}, "5 training jobs have been added to the ", React.createElement("span", null, "Prithvi - 8 node"), " cluster."), React.createElement("section", null, this.state.jobs.map(function (job, i) {
            return (React.createElement(Panel_1.default, {key: i}, React.createElement("div", {className: "panel-body"}, React.createElement("div", {className: "panel-title"}, "Training Job: ", job.name, " from ", job.project, React.createElement("div", {className: "panel-sub-title"}, "Started ", job.startTime, " by ", job.author)), React.createElement("div", {className: "panel-info"}, React.createElement(ProgressBar_1.default, {showPercentage: true, onComplete: _this.onComplete.bind(_this), onClick: _this.onClick.bind(_this)}))), React.createElement("div", {className: "panel-actions"}, React.createElement("div", {className: "panel-action"}, React.createElement("div", null, React.createElement("i", {className: "fa fa-pause"})), React.createElement("div", null, "Pause")), React.createElement("div", {className: "panel-action"}, React.createElement("div", null, React.createElement("i", {className: "fa fa-stop"})), React.createElement("div", null, "Cancel")))));
        }), React.createElement(react_router_1.Link, {to: "/projects/models", className: "default link-leaderboard"}, "Return to Model Leaderboard"), React.createElement(react_router_1.Link, {to: "/projects/deployments"}, "See all jobs on Prithbi - 8 node"))));
    };
    return NewProjectStep3;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = NewProjectStep3;
//# sourceMappingURL=NewProjectStep3.js.map