import * as _ from 'lodash';
import {
  RECEIVE_MEMBERS_FOR_PROJECT, RECEIVE_LABELS_FOR_PROJECT
} from '../actions/collaborators.actions';

let initialState = {
  members: [],
  labels: []
};

export const collaboratorsReducer = (state = initialState, action) => {
  switch (action.type) {
    case RECEIVE_MEMBERS_FOR_PROJECT :
      return _.assign({}, state, { members: action.members });
    case RECEIVE_LABELS_FOR_PROJECT :
      return _.assign({}, state, { labels: action.labels });
    default:
      return state;
  }
};
