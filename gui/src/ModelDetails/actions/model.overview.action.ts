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
 * Created by justin on 6/28/16.
 */
import * as Remote from '../../Proxy/Proxy';
import { openNotification, closeNotificationManager } from '../../App/actions/notification.actions';
import { NotificationType } from '../../App/components/Notification';
import { hashHistory } from 'react-router';
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
      if (error) {
        dispatch(openNotification(NotificationType.Error, 'Load Error', error.toString(), null));
        return;
      }
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

export function deployModel(modelId: number, name: string, projectId: string, packageName: string): Function {
  return (dispatch) => {
    dispatch(openNotification(NotificationType.Info, 'Deploying model', "Deploying Model", null));
    Remote.startService(modelId, name, packageName, (error, res) => {
      if (error) {
        dispatch(openNotification(NotificationType.Error, "Deployment Error", error.toString(), null));
        return;
      }
      dispatch(closeNotificationManager());
      hashHistory.push('/projects/' + projectId + '/deployment');
    });
  };
}
