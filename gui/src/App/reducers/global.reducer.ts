import * as _ from 'lodash';
import { RECEIVE_ENTITY_IDS } from '../actions/global.actions';

const initialState = {
  entityIds: {}
};

export function globalReducer(state = initialState, action) {
  switch (action.type) {
    case RECEIVE_ENTITY_IDS:
      let entityIds = {};
      for (let entityType of action.response) {
        entityIds[entityType.name] = entityType.id;
      }
      return _.assign({}, state, { entityIds });
    default:
      return state;
  }
}
