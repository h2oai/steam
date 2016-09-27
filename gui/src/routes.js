/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
"use strict";
/**
 * Created by Jeff Fohl <jfohl@h2o.ai> on 6/15/16.
 * This is where we configure the structure of the app,
 * and add additional information, such as how to handle breadcrumbs
 */
var App_1 = require('./App/App');
var Clusters_1 = require('./Clusters/Clusters');
var Users_1 = require('./Users/Users');
var Models_1 = require('./Models/Models');
var WelcomeSplashScreen_1 = require('./Projects/components/WelcomeSplashScreen');
var ModelDetails_1 = require('./ModelDetails/ModelDetails');
var ImportNewProject_1 = require('./Projects/components/ImportNewProject');
var NewProjectStep2_1 = require('./Projects/components/NewProjectStep2');
var NewProjectStep3_1 = require('./Projects/components/NewProjectStep3');
var CreateNewModel_1 = require('./Projects/components/CreateNewModel');
var Deployment_1 = require('./Deployment/Deployment');
var Configurations_1 = require('./Configurations/Configurations');
var ProjectScreenStrategy_1 = require('./Projects/components/ProjectScreenStrategy');
var Collaborators_1 = require('./Collaborators/Collaborators');
var LaunchCluster_1 = require('./Clusters/components/LaunchCluster');
exports.routes = [
    {
        path: '/',
        component: App_1.default,
        name: 'Home',
        showInBreadcrumb: true,
        showInNavigation: false,
        indexRoute: {
            component: WelcomeSplashScreen_1.default,
            name: 'Welcome'
        },
        childRoutes: [
            // /projects
            {
                path: 'projects',
                component: ProjectScreenStrategy_1.default,
                name: 'Projects',
                showInBreadcrumb: true,
                showInNavigation: true,
                icon: 'fa fa-folder'
            },
            {
                path: 'projects/:projectid',
                component: Models_1.default,
                name: 'Project',
                showInBreadcrumb: true,
                showInNavigation: false,
                showChildrenAsSubmenu: true
            },
            // /projects/:id/models
            {
                path: 'projects/:projectid/models',
                component: Models_1.default,
                name: "Models",
                showInBreadcrumb: true,
                showInNavigation: true
            },
            // /projects/:id/models/:id
            {
                path: 'projects/:projectid/models/:modelid',
                component: ModelDetails_1.default,
                name: "Model Detail",
                showInBreadcrumb: true,
                showInNavigation: false
            },
            // forkmodel
            {
                path: 'projects/:projectid/models/:modelid/forkmodel',
                component: CreateNewModel_1.default,
                name: 'Create New Model',
                showInBreadcrumb: true,
                showInNavigation: false
            },
            // /projects/:id/deployment
            {
                path: 'projects/:projectid/deployment',
                component: Deployment_1.default,
                name: 'Deployment',
                showInBreadcrumb: true,
                showInNavigation: true
            },
            {
                path: 'projects/:projectid/configurations',
                component: Configurations_1.default,
                name: 'Configurations',
                showInBreadcrumb: true,
                showInNavigation: true
            },
            {
                path: 'projects/:projectid/collaborators',
                component: Collaborators_1.default,
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
                    component: ImportNewProject_1.default,
                    name: '1'
                }
            },
            // /newproject/1
            {
                path: 'newproject/import',
                component: ImportNewProject_1.default,
                name: 'Step 1',
                showInBreadcrumb: true,
                showInNavigation: false
            },
            // /newproject/2
            {
                path: 'newproject/2',
                component: NewProjectStep2_1.default,
                name: 'Step 2',
                showInBreadcrumb: true,
                showInNavigation: false
            },
            // /newproject/3
            {
                path: 'newproject/3',
                component: NewProjectStep3_1.default,
                name: 'Step 3',
                showInBreadcrumb: true,
                showInNavigation: false
            },
            // /services
            {
                path: 'services',
                name: 'Services',
                icon: 'fa fa-cloud',
                component: Deployment_1.default,
                showInBreadcrumb: true,
                showInNavigation: true
            },
            // /clusters
            {
                path: 'clusters',
                component: Clusters_1.default,
                name: 'Clusters',
                icon: 'fa fa-cube',
                showInBreadcrumb: true,
                showInNavigation: true,
            },
            {
                path: 'clusters/new',
                component: LaunchCluster_1.default,
                name: 'Clusters',
                icon: 'fa fa-cube',
                showInBreadcrumb: false,
                showInNavigation: false
            },
            //users
            {
                path: 'users',
                component: Users_1.default,
                name: 'Users',
                icon: 'fa fa-user',
                showInBreadcrumb: true,
                showInNavigation: true
            }
        ]
    }
];
//# sourceMappingURL=routes.js.map