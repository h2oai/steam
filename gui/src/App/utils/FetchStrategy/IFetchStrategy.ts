/**
 * Created by justin on 6/28/16.
 */

import { IFetchStrategyConfig } from './IFetchStrategyConfig';

export interface IFetchStrategy {
  request(dispatch: Redux.Dispatch, config: IFetchStrategyConfig): void;
}