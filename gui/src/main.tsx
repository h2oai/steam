/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { Router, Route, IndexRoute, IndexRedirect, hashHistory, browserHistory } from 'react-router';
import { syncHistoryWithStore } from 'react-router-redux';
import { Store, createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import { routes } from './Navigation/routes';
import { rootReducer } from './App/reducers/rootReducer';

import './variables.scss';
import 'font-awesome/css/font-awesome.css';

const initialState = {};

const store: any = createStore(
  rootReducer,
  initialState,
  applyMiddleware(thunk)
);

let history: any = syncHistoryWithStore(hashHistory, store);

ReactDOM.render(
  <Provider store={store}>
    <Router history={hashHistory} routes={routes}></Router>
  </Provider>,
  document.getElementById('app'));
