/**
 * Created by justin on 7/15/16.
 */
import * as _ from 'lodash';
import { RECEIVE_PROJECTS, CREATED_PROJECT, SET_CURRENT_PROJECT, GET_CURRENT_PROJECT } from '../actions/projects.actions';

let initialState = {
  projects: [],
  project: {
    name: ''
  }
};

export const projectsReducer = (state = initialState, action: any) => {
  switch (action.type) {
    case RECEIVE_PROJECTS:
      return _.assign({}, state, {
        projects: action.projects
      });
    case CREATED_PROJECT:
      console.log(action.project);
      return _.assign({}, state, {
        project: action.project.result
      });
    case SET_CURRENT_PROJECT:
      return _.assign({}, state, {
        project: action.project
      });
    case GET_CURRENT_PROJECT:
      return _.assign({}, state, {
        project: state.project
      });
    default:
      return state;
  }
};
