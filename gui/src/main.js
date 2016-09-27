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
/**
 *    Copyright (C) 2016 H2Oai Inc.
 *
 *    This program is free software: you can redistribute it and/or  modify
 *    it under the terms of the GNU Affero General Public License, version 3,
 *    as published by the Free Software Foundation.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    As a special exception, the copyright holders give permission to link the
 *    code of portions of this program with the OpenSSL library under certain
 *    conditions as described in each individual source file and distribute
 *    linked combinations including the program with the OpenSSL library. You
 *    must comply with the GNU Affero General Public License in all respects for
 *    all of the code used other than as permitted herein. If you modify file(s)
 *    with this exception, you may extend this exception to your version of the
 *    file(s), but you are not obligated to do so. If you do not wish to do so,
 *    delete this exception statement from your version. If you delete this
 *    exception statement from all source files in the program, then also delete
 *    it in the license file.
 */
var React = require('react');
var ReactDOM = require('react-dom');
var CLI = require('./Proxy/CLI');
var react_router_1 = require('react-router');
var react_router_redux_1 = require('react-router-redux');
var redux_1 = require('redux');
var react_redux_1 = require('react-redux');
var redux_thunk_1 = require('redux-thunk');
var routes_1 = require('./routes');
var rootReducer_1 = require('./App/reducers/rootReducer');
require('./variables.scss');
require('font-awesome/css/font-awesome.css');
var redux_2 = require('redux');
// Export Steam API to browser console.
window.steam = CLI;
var initialState = {};
var store = redux_1.createStore(rootReducer_1.rootReducer, initialState, redux_2.compose(redux_1.applyMiddleware(redux_thunk_1.default), window.devToolsExtension ? window.devToolsExtension() : function (f) { return f; }));
var history = react_router_redux_1.syncHistoryWithStore(react_router_1.hashHistory, store);
ReactDOM.render(React.createElement(react_redux_1.Provider, {store: store}, React.createElement(react_router_1.Router, {history: react_router_1.hashHistory, routes: routes_1.routes})), document.getElementById('app'));
//# sourceMappingURL=main.js.map