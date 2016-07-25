/**
 * Created by justin on 7/18/16.
 */
import * as _ from 'lodash';
import {
  RECEIVE_CLUSTERS, RECEIVE_MODELS, CREATE_PROJECT_COMPLETED,
  RECEIVE_PROJECTS
} from '../actions/projects.actions';

let initialState = {
  clusters: [],
  models: [],
  project: {},
  availableProjects: []
};

export const projectsReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case RECEIVE_CLUSTERS:
      return _.assign({}, state, {
        clusters: action.clusters
      });
    case RECEIVE_MODELS:
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
    default:
      return state;
  }
};
