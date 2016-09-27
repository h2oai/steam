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
 * Created by justin on 7/5/16.
 */
var React = require('react');
var redux_1 = require('redux');
var react_redux_1 = require('react-redux');
var Leaderboard_1 = require('./components/Leaderboard');
var leaderboard_actions_1 = require('./actions/leaderboard.actions');
var projects_actions_1 = require('../Projects/actions/projects.actions');
var model_overview_action_1 = require('../ModelDetails/actions/model.overview.action');
var Models = (function (_super) {
    __extends(Models, _super);
    function Models() {
        _super.call(this);
        this.state = {
            modelCategory: null
        };
    }
    Models.prototype.componentWillMount = function () {
        var _this = this;
        if (this.props.project) {
            this.props.fetchProject(parseInt(this.props.params.projectid, 10)).then(function (res) {
                _this.props.fetchLeaderboard(parseInt(_this.props.params.projectid, 10), res.model_category);
                _this.props.fetchSortCriteria(res.model_category.toLowerCase());
                _this.setState({
                    modelCategory: res.model_category.toLowerCase()
                });
            });
        }
    };
    Models.prototype.onFilter = function (filters, name, offset) {
        this.props.fetchLeaderboard(parseInt(this.props.params.projectid, 10), this.state.modelCategory, name, filters.sortBy, filters.orderBy === 'asc', offset);
    };
    Models.prototype.render = function () {
        if (!this.props.leaderboard) {
            return React.createElement("div", null);
        }
        return (React.createElement("div", {className: "projects"}, React.createElement(Leaderboard_1.default, {items: this.props.leaderboard, projectId: parseInt(this.props.params.projectid, 10), modelCategory: this.state.modelCategory, sortCriteria: this.props.sortCriteria, onFilter: this.onFilter.bind(this), deployModel: this.props.deployModel, fetchLeaderboard: this.props.fetchLeaderboard})));
    };
    return Models;
}(React.Component));
exports.Models = Models;
function mapStateToProps(state) {
    return {
        leaderboard: state.leaderboard.items,
        sortCriteria: state.leaderboard.criteria,
        project: state.projects.project
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchLeaderboard: redux_1.bindActionCreators(leaderboard_actions_1.fetchLeaderboard, dispatch),
        deployModel: redux_1.bindActionCreators(model_overview_action_1.deployModel, dispatch),
        fetchSortCriteria: redux_1.bindActionCreators(leaderboard_actions_1.fetchSortCriteria, dispatch),
        fetchProject: redux_1.bindActionCreators(projects_actions_1.fetchProject, dispatch)
    };
}
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = react_redux_1.connect(mapStateToProps, mapDispatchToProps)(Models);
//# sourceMappingURL=Models.js.map