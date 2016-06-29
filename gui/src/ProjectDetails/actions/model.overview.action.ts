import { MockFetchStrategy } from '../../App/utils/FetchStrategy/MockFetchStrategy';
/**
 * Created by justin on 6/28/16.
 */
export const FETCH_MODEL_OVERVIEW = 'FETCH_MODEL_OVERVIEW';
export const RECEIVE_MODEL_OVERVIEW = 'RECEIVE_MODEL_OVERVIEW';

export const requestModelOverview = () => {
  return {
    type: FETCH_MODEL_OVERVIEW
  };
};

export function receiveModelOverview(modelOverview) {
  return {
    type: RECEIVE_MODEL_OVERVIEW,
    modelOverview
  }
}

export function fetchModelOverview() {
  return (dispatch) => {
    dispatch(requestModelOverview());
    let basics = [
      {
        label: 'Author',
        value: 'Mark Landry'
      },
      {
        label: 'Date',
        value: '2016-06-06 17:17'
      },
      {
        label: 'Size',
        value: '286.3MB'
      },
      {
        label: 'Training Time',
        value: 47000
      },
      {
        label: 'Classification Speed',
        value: 131
      },
      {
        label: 'Model Type',
        value: 'GBM'
      }
    ];
    let parameters = [
      {
        label: 'ntrees',
        value: 50,
      },
      {
        label: 'max-depth',
        value: 5,
      },
      {
        label: 'min-rows',
        value: 10,
      },
      {
        label: 'learn rate',
        value: 0.1,
      },
      {
        label: 'nfold',
        value: 6.0,
      }
    ];
    new MockFetchStrategy().request(dispatch, {
      callback: receiveModelOverview,
      data: {
        basics,
        parameters
      }
    });
  };
}
