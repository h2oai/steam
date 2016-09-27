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
var react_router_1 = require('react-router');
var buildPath_1 = require('../utils/buildPath');
var getRoute_1 = require('../utils/getRoute');
require('../styles/breadcrumb.scss');
var Breadcrumb = (function (_super) {
    __extends(Breadcrumb, _super);
    function Breadcrumb() {
        _super.apply(this, arguments);
    }
    Breadcrumb.prototype.render = function () {
        var _this = this;
        var crumbs = [];
        if (this.props.routes && this.props.routes[this.props.routes.length - 1] && this.props.routes[this.props.routes.length - 1].path) {
            var pathParts = this.props.routes[this.props.routes.length - 1].path.split('/');
            var path_1 = '';
            pathParts.forEach(function (part) {
                path_1 += (path_1 === '') ? part : '/' + part;
                var route = getRoute_1.getRoute(path_1);
                if (!route.showInBreadcrumb) {
                    return;
                }
                var crumb = {
                    path: buildPath_1.buildPath(route.path, _this.props.params),
                    name: _this.props.params[part.slice(1, part.length)] || route.name
                };
                crumbs.push(crumb);
            });
        }
        crumbs.unshift({
            path: '/',
            name: 'Home'
        });
        var breadClass = (crumbs.length <= 1) ? "breadcrumb" : "filled breadcrumb";
        return (React.createElement("ol", {className: breadClass}, crumbs.map(function (route, i) {
            if (crumbs.length > 1) {
                var crumb = (i === crumbs.length - 1) ? route.name : (React.createElement(react_router_1.Link, {to: route.path}, route.name));
                return React.createElement("li", {key: i}, crumb);
            }
            return null;
        })));
    };
    return Breadcrumb;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = Breadcrumb;
//# sourceMappingURL=Breadcrumb.js.map