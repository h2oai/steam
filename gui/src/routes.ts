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
import ProjectDetails from './ProjectDetails/ProjectDetails';
import NewProjectStep1 from './Projects/components/NewProjectStep1';
import ImportNewProject from './Projects/components/ImportNewProject';
import NewProjectStep2 from './Projects/components/NewProjectStep2';
import NewProjectStep3 from './Projects/components/NewProjectStep3';
import Deployments from './Projects/components/Deployments';
import CreateNewModel from './Projects/components/CreateNewModel';
import Dummy from './Dummy/Dummy';

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

// <<<<<<< HEAD
//   <Router history={history}>
// <Route path="/" component={App} isExcludedFromBreadcrumb={true}>
// <IndexRoute component={WelcomeSplashScreen}/>
// <Route path="projects" component={Projects} name="Projects" isExcludedFromBreadcrumb={true}>
// <IndexRoute component={WelcomeSplashScreen}/>
// <Route path=":id/models" component={Models}/>
// <Route path="deployments" component={Deployments} name="Deployments"/>
// <Route path="new" isExcludedFromBreadcrumb={true}>
// <IndexRoute component={NewProjectStep1} name="Create New Project"/>
// <Route path="import" component={ImportNewProject} name="Create New Project"/>
// <Route path="3" component={NewProjectStep3} isExcludedFromBreadcrumb={true}/>
// </Route>
// </Route>
// <Route path="clusters" component={Clusters}/>
// <Route path="forkmodel" component={CreateNewModel} name="Create New Model"/>
// <Route path="models/:id" component={ProjectDetails}/>
// </Route>
// </Router>
// =======

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
        component: WelcomeSplashScreen,
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
          component: ProjectDetails,
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
          component: Dummy,
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
      component: Deployments,
      showInBreadcrumb: true,
      showInNavigation: true,
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
    },
    // /clusters
    {
      path: 'clusters',
      component: Dummy,
      name: 'Clusters',
      icon: 'fa fa-cube',
      showInBreadcrumb: true,
      showInNavigation: true,
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
