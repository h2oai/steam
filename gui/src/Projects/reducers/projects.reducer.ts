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

/**
 * Created by justin on 7/18/16.
 */
import * as _ from 'lodash';
import {
  RECEIVE_CLUSTERS, RECEIVE_MODELS, CREATE_PROJECT_COMPLETED, SET_CURRENT_PROJECT,
  RECEIVE_PROJECTS, RECEIVE_DATASETS_FROM_CLUSTER, RECEIVE_MODELS_FROM_PROJECT, RECEIVE_PROJECT, REQUEST_CLUSTERS, REQUEST_MODELS, REQUEST_DELETE_PROJECT, RECEIVE_DELETE_PROJECT
} from '../actions/projects.actions';

let initialState = {
  clusters: [],
  models: [],
  project: {},
  availableProjects: null,
  isClusterFetchInProcess: false
};

export const projectsReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case SET_CURRENT_PROJECT :
      if (state.project.hasOwnProperty("id")) {
        if ((state.project as any).id === action.projectId) {
          return state;
        }
      }
      var toReturn: any =  _.assign({}, state);
      toReturn.project = { id: action.projectId };
      return toReturn;
    case REQUEST_CLUSTERS:
      return _.assign({}, state, {
        isClusterFetchInProcess: true
      });
    case RECEIVE_CLUSTERS:
      return _.assign({}, state, {
        clusters: action.clusters,
        isClusterFetchInProcess: false
      });
    case REQUEST_MODELS:
      return _.assign({}, state, {
        isModelFetchInProcess: true
      });
    case RECEIVE_MODELS:
      return _.assign({}, state, {
        models: action.models,
        isModelFetchInProcess: false
      });
    case RECEIVE_PROJECT:
      return _.assign({}, state, {
        project: action.project
      });
    case RECEIVE_MODELS_FROM_PROJECT:
      return _.assign({}, state, {
        models: action.models
      });
    case CREATE_PROJECT_COMPLETED:
      return _.assign({}, state, {
        project: action.project
      });
    case RECEIVE_PROJECTS:
      return _.assign({}, state, {
        availableProjects: action.projects
      });
    case RECEIVE_DATASETS_FROM_CLUSTER:
      return _.assign({}, state, {
        datasets: action.datasets
      });
    case REQUEST_DELETE_PROJECT:
      var toReturn: any = _.assign({}, state);
      toReturn.availableProjects = toReturn.availableProjects.slice();
      for (let project of toReturn.availableProjects) {
        if (project.id === action.projectId) {
          project.isDeleteInProgress = true;
        }
      }
      return toReturn;
    case RECEIVE_DELETE_PROJECT:
      if (!action.successful) {
        var toReturn: any = _.assign({}, state);
        toReturn.availableProjects = toReturn.availableProjects.slice();
        for (let project of toReturn.availableProjects) {
          if (project.id === action.projectId) {
            project.isDeleteInProgress = false;
          }
        }
        return toReturn;
      } else {
        return state;
      }
    default:
      return state;
  }
};
