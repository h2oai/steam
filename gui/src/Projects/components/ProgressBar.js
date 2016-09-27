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
var classNames = require('classnames');
var _ = require('lodash');
require('../styles/progressbar.scss');
var ProgressBar = (function (_super) {
    __extends(ProgressBar, _super);
    function ProgressBar() {
        _super.call(this);
        this.state = {
            progress: 0
        };
    }
    ProgressBar.prototype.componentDidMount = function () {
        this.start();
    };
    ProgressBar.prototype.componentWillUnmount = function () {
        clearInterval(this.interval);
    };
    ProgressBar.prototype.start = function () {
        var _this = this;
        var maxIncrements = Math.floor(Math.random() * (100 - 40 + 1)) + 40;
        var i = 0;
        this.interval = setInterval(function () {
            i++;
            var remaining = 100 - _this.state.progress;
            _this.setState({
                progress: _this.state.progress += (0.05 * Math.pow(1 - Math.sqrt(remaining), 2))
            });
            if (i >= maxIncrements) {
                _this.end();
            }
        }, 50);
    };
    ProgressBar.prototype.end = function () {
        this.setState({
            progress: 100
        });
        clearInterval(this.interval);
        if (this.props.onComplete) {
            this.props.onComplete(this);
        }
    };
    ProgressBar.prototype.render = function () {
        return (React.createElement("div", {ref: "progressBar", className: classNames('progress-bar-container', { complete: this.state.progress === 100 }, this.props.className), onClick: this.props.onClick}, React.createElement("div", {className: "progress-bar", style: { width: this.state.progress + '%' }}), React.createElement("div", {className: "progress-counter"}, this.props.showPercentage === true && _.isEmpty(this.props.children) ? Math.ceil(this.state.progress) + '%' : null)));
    };
    return ProgressBar;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = ProgressBar;
//# sourceMappingURL=ProgressBar.js.map