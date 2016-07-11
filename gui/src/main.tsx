/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { Router, Route, IndexRoute, hashHistory } from 'react-router';
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
import NewProject from './Projects/components/NewProject';
import { rootReducer } from './App/reducers/rootReducer';

import './variables.scss';
import 'font-awesome/css/font-awesome.css';

const initialState = {};

const store: Store = createStore(
  rootReducer,
  initialState,
  applyMiddleware(thunk)
);

let history: ReactRouterRedux.ReactRouterReduxHistory = syncHistoryWithStore(hashHistory, store);

ReactDOM.render(
  <Provider store={store}>
    <Router history={history}>
      <Route path="/" component={App}>
        <IndexRoute component={WelcomeSplashScreen}/>
        <Route path="projects" component={Projects}>
          <IndexRoute component={WelcomeSplashScreen}/>
          <Route path="new" component={NewProject}/>
        </Route>
        <Route path="clusters" component={Clusters}/>
        <Route path="models" component={Models}/>
        <Route path="models/:id" component={ProjectDetails}/>
      </Route>
    </Router>
  </Provider>,
  document.getElementById('app'));
