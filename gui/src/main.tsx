/**
 Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.
 */

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

ReactDOM.render(
  <Provider store={store}>
    <Router history={hashHistory} routes={routes}></Router>
  </Provider>,
  document.getElementById('app'));
