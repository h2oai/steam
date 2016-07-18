/**
 * Created by justin on 7/15/16.
 */
import * as Remote from '../../Proxy/Proxy';
export const FETCH_PROJECTS = 'FETCH_LEADERBOARD';
export const RECEIVE_PROJECTS = 'RECEIVE_LEADERBOARD';
export const CREATING_PROJECT = 'CREATING_PROJECT';
export const CREATED_PROJECT = 'CREATED_PROJECT';
export const SET_CURRENT_PROJECT = 'SET_CURRENT_PROJECT';
export const GET_CURRENT_PROJECT = 'GET_CURRENT_PROJECT';

export const requestProjects = () => {
  return {
    type: FETCH_PROJECTS
  };
};

export function receiveProjects(projects) {
  return {
    type: RECEIVE_PROJECTS,
    projects
  };
}

export function creatingProject(project) {
  return {
    type: CREATING_PROJECT,
    project
  };
}

export function createdProject(project) {
  return {
    type: CREATED_PROJECT,
    project
  }
}

export function setCurrentProject(project) {
  return {
    type: SET_CURRENT_PROJECT,
    project
  };
}

export function getCurrentProject() {
  return {
    type: GET_CURRENT_PROJECT
  };
}

export function fetchProjects() {
  return (dispatch) => {
    dispatch(requestProjects());
    Remote.getProjects(0, 5, (error, res) => {
      dispatch(receiveProjects(res));
    });
  };
}

export function createProject(project: {name: string, description: string}) {
  return (dispatch) => {
    dispatch(creatingProject(project));
    Remote
    Remote.createProject(project.name, project.description, (error, res) => {
      dispatch(createdProject(res));
    })
  }
}
