/**
 * Created by Jeff Fohl <jfohl@h2o.ai> on 6/15/16.
 * This is where we configure the structure of the app, an add additional information, such as how to handle breadcrumbs
 */

 import App from '../App/App';
 import Clusters from '../Clusters/Clusters';
 import Models from '../Models/Models';
 import Projects from '../Projects/Projects';
 import WelcomeSplashScreen from '../Projects/components/WelcomeSplashScreen';
 import ProjectDetails from '../ProjectDetails/ProjectDetails';
 import NewProjectStep1 from '../Projects/components/NewProjectStep1';
 import NewProjectStep2 from '../Projects/components/NewProjectStep2';
 import NewProjectStep3 from '../Projects/components/NewProjectStep3';
 import Deployments from '../Projects/components/Deployments';
 import CreateNewModel from '../Projects/components/CreateNewModel';

export const routes = [{
  path: '/',
  component: App,
  name: 'Home',
  showInBreadcrumb: true,
  showInNavigation: false,
  indexRoute: {
    component: WelcomeSplashScreen,
    name: 'Welcome',
    showInBreadcrumb: false,
    showInNavigation: false
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
        component: WelcomeSplashScreen,
        name: 'Welcome',
        showInBreadcrumb: false,
        showInNavigation: false,
      },
      childRoutes: [
        // /projects/deployments
        {
          path: 'deployments',
          component: Deployments,
          name: 'Deployments',
          showInBreadcrumb: true,
          showInNavigation: true,
        },
        // /projects/new
        {
          path: 'new',
          name: 'New',
          showInBreadcrumb: false,
          showInNavigation: false,
          indexRoute: {
            component: NewProjectStep1,
            name: '1'
          },
          childRoutes: [
            // /projects/new/1
            {
              path: '1',
              component: NewProjectStep1,
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
        // /projects/clusters
        {
          path: 'clusters',
          component: Clusters,
          name: 'Clusters',
          showInBreadcrumb: true,
          showInNavigation: true
        },
        // /projects/models
        {
          path: 'models',
          component: Models,
          name: 'Models',
          showInBreadcrumb: true,
          showInNavigation: true
        },
        // /projects/forkmodel
        {
          path: 'forkmodel',
          component: CreateNewModel,
          name: 'Fork Model',
          showInBreadcrumb: true,
          showInNavigation: false
        },
        // /projects/models/:id
        {
          path: 'models/:id',
          component: ProjectDetails,
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
      showInBreadcrumb: true,
      showInNavigation: true,
      childRoutes: []
    },
    // /clusters
    {
      path: 'clusters',
      name: 'Clusters',
      icon: 'fa fa-cube',
      showInBreadcrumb: true,
      showInNavigation: true,
      childRoutes: [
        {
          path: 'models',
          name: 'Sub 1'
        }
      ]
    },
    // /team
    {
      path: 'team',
      name: 'Team',
      icon: 'fa fa-users',
      showInBreadcrumb: true,
      showInNavigation: true,
      childRoutes: []
    }
  ]
}];
