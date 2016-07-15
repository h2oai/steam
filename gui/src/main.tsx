/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { Router, Route, IndexRoute, IndexRedirect, hashHistory } from 'react-router';
import { syncHistoryWithStore } from 'react-router-redux';
import { Store, createStore, applyMiddleware, compose } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import App from './App/App';
import Clusters from './Clusters/Clusters';
import Models from './Models/Models';
import Projects from './Projects/Projects';
import WelcomeSplashScreen from './Projects/components/WelcomeSplashScreen';
import ProjectDetails from './ProjectDetails/ProjectDetails';
import NewProjectStep1 from './Projects/components/NewProjectStep1';
import NewProjectStep2 from './Projects/components/NewProjectStep2';
import NewProjectStep3 from './Projects/components/NewProjectStep3';
import Deployments from './Projects/components/Deployments';
import CreateNewModel from './Projects/components/CreateNewModel';
import { rootReducer } from './App/reducers/rootReducer';

import './variables.scss';
import 'font-awesome/css/font-awesome.css';

const initialState = {};

const store: any = createStore(
  rootReducer,
  initialState,
  compose(
    applyMiddleware(thunk),
    window.devToolsExtension && window.devToolsExtension()
  )
);

let history: any = syncHistoryWithStore(hashHistory, store);

ReactDOM.render(
  <Provider store={store}>
    <Router history={history}>
      <Route path="/" component={App} isExcludedFromBreadcrumb={true}>
        <IndexRoute component={WelcomeSplashScreen}/>
        <Route path="projects" component={Projects} name="Projects" isExcludedFromBreadcrumb={true}>
          <IndexRoute component={WelcomeSplashScreen}/>
          <Route path="deployments" component={Deployments} name="Deployments"/>
          <Route path="new" isExcludedFromBreadcrumb={true}>
            <Route path="1" component={NewProjectStep1} name="Create New Project"/>
            <Route path="2" component={NewProjectStep2} isExcludedFromBreadcrumb={true}/>
            <Route path="3" component={NewProjectStep3} isExcludedFromBreadcrumb={true}/>
            <IndexRedirect to="1"/>
          </Route>
        </Route>
        <Route path="clusters" component={Clusters}/>
        <Route path="models" component={Models}/>
        <Route path="forkmodel" component={CreateNewModel} name="Create New Model"/>
        <Route path="models/:id" component={ProjectDetails}/>
      </Route>
    </Router>
  </Provider>,
  document.getElementById('app'));
