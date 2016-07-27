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
                component: Projects,
                name: 'Projects',
                showInBreadcrumb: true,
                showInNavigation: true,
                icon: 'fa fa-folder'
            },
            {
                path: 'projects/:projectid',
                component: ProjectDetails,
                name: 'Project Details',
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
                component: Dummy,
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
            // /projects/:id/data
            {
                path: 'projects/:projectid/data',
                component: Dummy,
                name: 'Data',
                showInBreadcrumb: true,
                showInNavigation: true
            },
            // /projects/:id/deployments
            {
                path: 'projects/:projectid/deployment',
                component: Dummy,
                name: 'Deployment',
                showInBreadcrumb: true,
                showInNavigation: true
            },
            // /projects/:id/collaborators
            {
                path: 'projects/:projectid/collaborators',
                component: Dummy,
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
                    component: NewProjectStep1,
                    name: '1'
                }
            },
            // /newproject/1
            {
                path: 'newproject/import',
                component: Dummy,
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
                component: Deployments,
                showInBreadcrumb: true,
                showInNavigation: true
            },
            // /clusters
            {
                path: 'clusters',
                component: Dummy,
                name: 'Clusters',
                icon: 'fa fa-cube',
                showInBreadcrumb: true,
                showInNavigation: true
            },
            // /team
            {
                path: 'team',
                component: Dummy,
                name: 'Team',
                icon: 'fa fa-cube',
                showInBreadcrumb: true,
                showInNavigation: true
            }

        ]
    }
];
