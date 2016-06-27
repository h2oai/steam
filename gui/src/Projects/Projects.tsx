/**
 * Created by justin on 6/27/16.
 */

import * as React from 'react';
import { Leaderboard }from './components/Leaderboard';

export default class Projects extends React.Component<any, any> {
  render() {
    let items = [
      {
        rank: 1,
        metadata: {
          name: 'DRF-1069085'
        }
      },
      {
        rank: 2,
        metadata: {
          name: 'DRF-1069099'
        }
      }
    ];
    return (
      <div>
        <Leaderboard items={items}></Leaderboard>
      </div>
    );
  }
}