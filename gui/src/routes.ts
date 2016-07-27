/**
 * Created by Jeff Fohl <jfohl@h2o.ai> on 6/15/16.
 * This is where we configure the structure of the app,
 * and add additional information, such as how to handle breadcrumbs
 */

import App from './App/App';
import Clusters from './Clusters/Clusters';
import Models from './Models/Models';
import Projects from './Projects/Projects';
import WelcomeSplashScreen from './Projects/components/WelcomeSplashScreen';
import ModelDetails from './ModelDetails/ModelDetails';
import NewProjectStep1 from './Projects/components/NewProjectStep1';
import ImportNewProject from './Projects/components/ImportNewProject';
import NewProjectStep2 from './Projects/components/NewProjectStep2';
import NewProjectStep3 from './Projects/components/NewProjectStep3';
import CreateNewModel from './Projects/components/CreateNewModel';
import Deployment from './Deployment/Deployment';
import Dummy from './Dummy/Dummy';
import ProjectScreenStrategy from './Projects/components/ProjectScreenStrategy';

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
  icon?: string
  indexRoute?: IIndexRoute
  childRoutes?: IRoute[]
}

export const routes: IRoute[] = [{
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
      component: Projects,
      name: 'Projects',
      showInBreadcrumb: true,
      showInNavigation: true,
      icon: 'fa fa-folder',
      indexRoute: {
        component: ProjectScreenStrategy,
        name: 'Welcome'
      },
      childRoutes: [
        // /projects/models
        {
          path: ':id/models',
          component: Models,
          name: 'Models',
          showInBreadcrumb: true,
          showInNavigation: true
        },
        // /projects/data
        {
          path: 'data',
          component: Dummy,
          name: 'Data',
          showInBreadcrumb: true,
          showInNavigation: true
        },
        // /models/:id
        {
          path: ':id/models/:modelId',
          component: ModelDetails,
          name: 'Models',
          showInBreadcrumb: true,
          showInNavigation: false
        },
        // /projects/data
        {
          path: 'assets',
          component: Dummy,
          name: 'Assets',
          showInBreadcrumb: true,
          showInNavigation: true
        },
        // /projects/deployments
        {
          path: 'deployment',
          component: Deployment,
          name: 'Deployment',
          showInBreadcrumb: true,
          showInNavigation: true
        },
        // /projects/collaborators
        {
          path: 'collaborators',
          component: Dummy,
          name: 'Collaborators',
          showInBreadcrumb: true,
          showInNavigation: true
        },
        // /projects/new
        {
          path: 'new',
          name: 'New',
          showInBreadcrumb: true,
          showInNavigation: false,
          indexRoute: {
            component: NewProjectStep1,
            name: '1'
          },
          childRoutes: [
            // /projects/new/1
            {
              path: 'import',
              component: ImportNewProject,
              name: 'Step 1',
              showInBreadcrumb: true,
              showInNavigation: false
            },
            // /projects/new/2
            {
              path: '2',
              component: NewProjectStep2,
              name: 'Step 2',
              showInBreadcrumb: true,
              showInNavigation: false
            },
            // /projects/new/3
            {
              path: '3',
              component: NewProjectStep3,
              name: 'Step 3',
              showInBreadcrumb: true,
              showInNavigation: false
            }
          ]
        },
        // /projects/forkmodel
        {
          path: 'forkmodel',
          component: CreateNewModel,
          name: 'Create New Model',
          showInBreadcrumb: true,
          showInNavigation: false
        }
      ]
    },
    // /services
    {
      path: 'services',
      name: 'Services',
      icon: 'fa fa-cloud',
      component: Deployment,
      showInBreadcrumb: true,
      showInNavigation: true,
      childRoutes: []
    },
    // /clusters
    {
      path: 'clusters',
      component: Clusters,
      name: 'Clusters',
      icon: 'fa fa-cube',
      showInBreadcrumb: true,
      showInNavigation: true,
      childRoutes: []
    },
    // /team
    {
      path: 'team',
      component: Dummy,
      name: 'Team',
      showInBreadcrumb: true,
      showInNavigation: true,
      icon: 'fa fa-users',
      childRoutes: [
        {
          path: 'submenu1',
          component: Dummy,
          name: 'Sub Menu 1',
          showInBreadcrumb: true,
          showInNavigation: true
        },
        {
          path: 'submenu2',
          component: Dummy,
          name: 'Sub Menu 2',
          showInBreadcrumb: true,
          showInNavigation: true
        },
        {
          path: 'submenu3',
          component: Dummy,
          name: 'Sub Menu 3',
          showInBreadcrumb: true,
          showInNavigation: true
        }
      ]
    }
  ]
}];
