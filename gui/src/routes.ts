/**
 * Created by Jeff Fohl <jfohl@h2o.ai> on 6/15/16.
 * This is where we configure the structure of the app,
 * and add additional information, such as how to handle breadcrumbs
 */

import App from './App/App';
import Clusters from './Clusters/Clusters';
import Users from './Users/Users';
import Models from './Models/Models';
import WelcomeSplashScreen from './Projects/components/WelcomeSplashScreen';
import ModelDetails from './ModelDetails/ModelDetails';
import NewProjectStep1 from './Projects/components/NewProjectStep1';
import ImportNewProject from './Projects/components/ImportNewProject';
import NewProjectStep2 from './Projects/components/NewProjectStep2';
import NewProjectStep3 from './Projects/components/NewProjectStep3';
import CreateNewModel from './Projects/components/CreateNewModel';
import Deployment from './Deployment/Deployment';
import Configurations from './Configurations/Configurations';
import Dummy from './Dummy/Dummy';
import ProjectScreenStrategy from './Projects/components/ProjectScreenStrategy';
import {Collaborators} from "./Collaborators/Collaborators";

interface IIndexRoute {
  component: any
  name: string
}

interface IRoute {
  path: string
  component?: any
  name: string;
  showInBreadcrumb?: boolean
  showInNavigation?: boolean
  showChildrenAsSubmenu?: boolean
  icon?: string
  indexRoute?: IIndexRoute
  childRoutes?: IRoute[]
}

export const routes: IRoute[] = [
  {
    path: '/',
    component: App,
    name: 'Home',
    showInBreadcrumb: true,
    showInNavigation: false,
    indexRoute: {
      component: WelcomeSplashScreen,
      name: 'Welcome'
    },
    childRoutes: [
      // /projects
      {
        path: 'projects',
        component: ProjectScreenStrategy,
        name: 'Projects',
        showInBreadcrumb: true,
        showInNavigation: true,
        icon: 'fa fa-folder'
      },
      {
        path: 'projects/:projectid',
        component: Dummy,
        name: 'Project',
        showInBreadcrumb: true,
        showInNavigation: false,
        showChildrenAsSubmenu: true
      },
      // /projects/:id/models
      {
        path: 'projects/:projectid/models',
        component: Models,
        name: "Models",
        showInBreadcrumb: true,
        showInNavigation: true
      },
      // /projects/:id/models/:id
      {
        path: 'projects/:projectid/models/:modelid',
        component: ModelDetails,
        name: "Model Detail",
        showInBreadcrumb: true,
        showInNavigation: false
      },
      // forkmodel
      {
        path: 'projects/:projectid/models/:modelid/forkmodel',
        component: CreateNewModel,
        name: 'Create New Model',
        showInBreadcrumb: true,
        showInNavigation: false
      },
      // /projects/:id/deployment
      {
        path: 'projects/:projectid/deployment',
        component: Deployment,
        name: 'Deployment',
        showInBreadcrumb: true,
        showInNavigation: true
      },
      {
        path: 'projects/:projectid/configurations',
        component: Configurations,
        name: 'Configurations',
        showInBreadcrumb: true,
        showInNavigation: true
      },
      {
        path: 'projects/:projectid/collaborators',
        component: Collaborators,
        name: 'Collaborators',
        showInBreadcrumb: true,
        showInNavigation: true
      },
      {
        path: 'newproject',
        name: 'New',
        showInBreadcrumb: true,
        showInNavigation: false,
        indexRoute: {
          component: ImportNewProject,
          name: '1'
        }
      },
      // /newproject/1
      {
        path: 'newproject/import',
        component: ImportNewProject,
        name: 'Step 1',
        showInBreadcrumb: true,
        showInNavigation: false
      },
      // /newproject/2
      {
        path: 'newproject/2',
        component: NewProjectStep2,
        name: 'Step 2',
        showInBreadcrumb: true,
        showInNavigation: false
      },
      // /newproject/3
      {
        path: 'newproject/3',
        component: NewProjectStep3,
        name: 'Step 3',
        showInBreadcrumb: true,
        showInNavigation: false
      },
      // /services
      {
        path: 'services',
        name: 'Services',
        icon: 'fa fa-cloud',
        component: Deployment,
        showInBreadcrumb: true,
        showInNavigation: true
      },
      // /clusters
      {
        path: 'clusters',
        component: Clusters,
        name: 'Clusters',
        icon: 'fa fa-cube',
        showInBreadcrumb: true,
        showInNavigation: true
      },
      //users
      {
        path: 'users',
        component: Users,
        name: 'Users',
        icon: 'fa fa-user',
        showInBreadcrumb: true,
        showInNavigation: true
      }
    ]
  }
];
