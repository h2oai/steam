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
var React = require('react');
var ReactDOM = require('react-dom');
var rocChart = require('roc-chart');
require('../styles/rocgraph.scss');
var colors_1 = require('../../App/utils/colors');
var RocGraph = (function (_super) {
    __extends(RocGraph, _super);
    function RocGraph() {
        _super.apply(this, arguments);
    }
    RocGraph.prototype.componentDidMount = function () {
        this._mountNode = ReactDOM.findDOMNode(this);
        this.renderGraph();
    };
    RocGraph.prototype.componentWillUnmount = function () {
        if (this._mountNode) {
            ReactDOM.unmountComponentAtNode(this._mountNode);
            this._mountNode.remove();
            this._mountNode = null;
        }
    };
    RocGraph.prototype.componentWillUpdate = function (nextProps) {
        var cfg = {
            margin: { top: 2, right: 2, bottom: 2, left: 2 },
            width: '100%',
            height: '100%',
            interpolationMode: 'basis',
            smooth: true,
            animate: undefined,
            hideAxes: true,
            hideAUCText: true,
            curveColors: [colors_1.BRAND_BLUE, colors_1.BRAND_ORANGE]
        };
        this._mountNode.innerHTML = '';
        rocChart.plot(this._mountNode, nextProps.data, this.props.config || cfg);
    };
    RocGraph.prototype.renderGraph = function () {
        var cfg = {
            margin: { top: 2, right: 2, bottom: 2, left: 2 },
            width: '100%',
            height: '100%',
            interpolationMode: 'basis',
            smooth: true,
            animate: false,
            hideAxes: true,
            hideAUCText: true,
            curveColors: [colors_1.BRAND_BLUE]
        };
        rocChart.plot(this._mountNode, this.props.data, this.props.config || cfg);
    };
    RocGraph.prototype.render = function () {
        return React.createElement("div", {className: "roc-container"});
    };
    return RocGraph;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = RocGraph;
//# sourceMappingURL=RocGraph.js.map