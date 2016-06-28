/**
 * Created by justin on 6/27/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import { Leaderboard }from './components/Leaderboard';
import { Pagination } from './components/Pagination';

export default class Projects extends React.Component<any, any> {
  render() {
    let items = [
      {
        id: 1,
        rank: 1,
        metadata: {
          name: 'DRF-1069085'
        }
      },
      {
        id: 2,
        rank: 2,
        metadata: {
          name: 'DRF-1069099'
        }
      }
    ];
    return (
      <div className="projects">
        <Leaderboard items={items}></Leaderboard>
        <Pagination></Pagination>
      </div>
    );
  }
}