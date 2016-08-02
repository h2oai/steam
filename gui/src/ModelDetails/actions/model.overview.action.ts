/**
 * Created by justin on 6/28/16.
 */
import * as Remote from '../../Proxy/Proxy';
export const FETCH_MODEL_OVERVIEW = 'FETCH_MODEL_OVERVIEW';
export const RECEIVE_MODEL_OVERVIEW = 'RECEIVE_MODEL_OVERVIEW';
export const FETCH_DOWNLOAD_MODEL = 'FETCH_DOWNLOAD_MODEL';
export const RECEIVE_DOWNLOAD_MODEL = 'RECEIVE_DOWNLOAD_MODEL';

export const requestModelOverview = () => {
  return {
    type: FETCH_MODEL_OVERVIEW
  };
};

export function receiveModelOverview(model) {
  return {
    type: RECEIVE_MODEL_OVERVIEW,
    model
  };
}

export const requestDownloadModel = () => {
  return {
    type: FETCH_DOWNLOAD_MODEL
  };
};

export function receiveDownloadModel(model) {
  return {
    type: RECEIVE_DOWNLOAD_MODEL,
    model
  };
}

export function fetchModelOverview(modelId: number): Function {
  return (dispatch) => {
    dispatch(requestModelOverview());
    Remote.getModel(modelId, (error, model) => {
      getModelStrategy(model.model_category.toLowerCase())(modelId, (error, res) => {
        dispatch(receiveModelOverview(res));
      });
    });
  };
}

function getModelStrategy(modelCategory): Function {
  if (modelCategory === 'binomial') {
    return Remote.getModelBinomial;
  } else if (modelCategory === 'multinomial') {
    return Remote.getModelMultinomial;
  } else if (modelCategory === 'regression') {
    return Remote.getModelRegression;
  }
}

export function downloadModel(): Function {
  /**
   * TODO(justinloyola): Waiting on endpoint
   */
  return (dispatch) => {
    dispatch(requestDownloadModel());
    dispatch(receiveDownloadModel({}));

  };
}

export function deployModel(modelId: number): Function {
  return (dispatch) => {
    /**
     * TODO(justinloyola):
     *   if arg2 is "", deploys a vanilla war file as a service
     *   if arg2 is a valid package name, the python scripts from the package are
     *     bundled into the war file. 
     */
    Remote.startService(modelId, "", (error, res) => {
      console.log(res);
    });
  };
}
