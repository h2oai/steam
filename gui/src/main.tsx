/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { Router, Route, IndexRoute, IndexRedirect, hashHistory } from 'react-router';
import { syncHistoryWithStore } from 'react-router-redux';
import { Store, createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import App from './App/App';
import Clusters from './Clusters/Clusters';
import Models from './Models/Models';
import Projects from './Projects/Projects';
import WelcomeSplashScreen from './Projects/components/WelcomeSplashScreen';
import ProjectDetails from './ProjectDetails/ProjectDetails';
import NewProjectStep1 from './Projects/components/NewProjectStep1';
import NewProjectStep3 from './Projects/components/NewProjectStep3';
import Deployments from './Projects/components/Deployments';
import CreateNewModel from './Projects/components/CreateNewModel';
import { routes } from './routes';
import { rootReducer } from './App/reducers/rootReducer';

import './variables.scss';
import 'font-awesome/css/font-awesome.css';
import ImportNewProject from './Projects/components/ImportNewProject';

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
