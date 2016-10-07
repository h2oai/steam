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

// Testing PR Builder

import * as React from 'react';
import * as ReactDOM from 'react-dom';
import * as CLI from './Proxy/CLI';
import { Router, hashHistory } from 'react-router';
import { syncHistoryWithStore } from 'react-router-redux';
import { createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import { routes } from './routes';
import { rootReducer } from './App/reducers/rootReducer';
import './variables.scss';
import 'font-awesome/css/font-awesome.css';
import { compose } from 'redux';

// Export Steam API to browser console.
(window as any).steam = CLI;

const initialState = {};

const store: any = createStore(
  rootReducer,
  initialState,
  compose(
    applyMiddleware(thunk), window.devToolsExtension ? window.devToolsExtension() : f => f)
);

let history: any = syncHistoryWithStore(hashHistory, store);

hashHistory.listen((e) => {
  if (window.ga) {
    window.ga('send', {
      hitType: 'pageview',
      page: e.pathname
    });
  }
});

ReactDOM.render(
  <Provider store={store}>
    <Router history={hashHistory} routes={routes}></Router>
  </Provider>,
  document.getElementById('app'));
