/**
 * Created by justin on 6/28/16.
 */

import { IFetchStrategy } from './IFetchStrategy';
import { IFetchStrategyConfig } from './IFetchStrategyConfig';

export class MockFetchStrategy implements IFetchStrategy {
  request(dispatch: Redux.Dispatch, config: IFetchStrategyConfig): void {
    setTimeout(() => {
      dispatch(config.callback(config.data));
    }, 100);
  }
}