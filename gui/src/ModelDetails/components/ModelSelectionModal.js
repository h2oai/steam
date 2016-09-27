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
 * Created by justin on 7/29/16.
 */
var React = require('react');
var moment = require('moment');
var FilterDropdown_1 = require('../../Models/components/FilterDropdown');
var DefaultModal_1 = require('../../App/components/DefaultModal');
var PageHeader_1 = require('../../Projects/components/PageHeader');
var Pagination_1 = require('../../Models/components/Pagination');
var Table_1 = require('../../Projects/components/Table');
var Row_1 = require('../../Projects/components/Row');
var Cell_1 = require('../../Projects/components/Cell');
require('../styles/modelselectionmodal.scss');
var leaderboard_actions_1 = require('../../Models/actions/leaderboard.actions');
var ModelSelectionModal = (function (_super) {
    __extends(ModelSelectionModal, _super);
    function ModelSelectionModal() {
        _super.call(this);
        this.state = {
            currentPage: 0,
            filters: {
                sortBy: '',
                orderBy: 'asc'
            }
        };
    }
    ModelSelectionModal.prototype.onFilter = function (filters) {
        this.props.onFilter(filters, this.refs.filterModels.value);
    };
    ModelSelectionModal.prototype.onPageForward = function () {
        this.setState({
            currentPage: ++this.state.currentPage
        });
        this.props.onFilter(this.state.filters, this.refs.filterModels.value, this.state.currentPage * leaderboard_actions_1.MAX_ITEMS);
    };
    ModelSelectionModal.prototype.onPageBack = function () {
        if (this.state.currentPage >= 0) {
            this.setState({
                currentPage: --this.state.currentPage
            });
            this.props.onFilter(this.state.filters, this.refs.filterModels.value, this.state.currentPage * leaderboard_actions_1.MAX_ITEMS);
        }
    };
    ModelSelectionModal.prototype.render = function () {
        var _this = this;
        return (React.createElement(DefaultModal_1.default, {className: "model-comparison-modal", open: this.props.open}, React.createElement(PageHeader_1.default, null, "CHOOSE MODEL TO COMPARE"), React.createElement("div", null, React.createElement("div", null, "Filter models by name"), React.createElement("input", {ref: "filterModels", type: "text", placeholder: "filter models", onChange: this.onFilter.bind(this)}), React.createElement(Table_1.default, null, React.createElement(Row_1.default, {header: true}, React.createElement(Cell_1.default, null, React.createElement(FilterDropdown_1.default, {onFilter: this.onFilter.bind(this), sortCriteria: this.props.sortCriteria})), React.createElement(Cell_1.default, null, "MODEL"), React.createElement(Cell_1.default, null, "DATE"), React.createElement(Cell_1.default, null, "MSE"), React.createElement(Cell_1.default, null, "AUC"), React.createElement(Cell_1.default, null)), this.props.models.map(function (model, i) {
            return (React.createElement(Row_1.default, {key: i}, React.createElement(Cell_1.default, null), React.createElement(Cell_1.default, null, model.name), React.createElement(Cell_1.default, null, moment.unix(model.created_at).format('YYYY-MM-DD HH:mm')), React.createElement(Cell_1.default, null, model.mse ? model.mse.toFixed(6) : 'N/A'), React.createElement(Cell_1.default, null, model.auc ? model.auc.toFixed(6) : 'N/A'), React.createElement(Cell_1.default, null, React.createElement("button", {className: "default", onClick: _this.props.onSelectModel.bind(_this, model)}, "Select"))));
        }))), React.createElement("footer", null, React.createElement(Pagination_1.default, {items: this.props.models, onPageForward: this.onPageForward.bind(this), onPageBack: this.onPageBack.bind(this), currentPage: this.state.currentPage, count: this.props.count}), React.createElement("button", {className: "default", onClick: this.props.onCancel.bind(this)}, "Cancel"))));
    };
    return ModelSelectionModal;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = ModelSelectionModal;
//# sourceMappingURL=ModelSelectionModal.js.map