/**
 * Created by justin on 6/27/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import { Link } from 'react-router';
import '../styles/leaderboard.scss';

interface Props {
  items: any[]
}

export class Leaderboard extends React.Component<Props, any> {
  static getOrdinal(rank: number): string {
    let suffixes = ['th', 'st', 'nd', 'rd'];
    let remainder = rank % 100;
    return (suffixes[(remainder - 20) % 10] || suffixes[remainder] || suffixes[0]);
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="leaderboard">
        <header>
          MODEL LEADERBOARD
        </header>
        <ul>
          {this.props.items.map((item, i) => {
            return (
              <li key={i} className={classNames('col-xs-12', {leader: item.rank === 1})}>
                <div className="col-sm-1 col-xs-12 rank">
                  {item.rank + Leaderboard.getOrdinal(item.rank)}
                </div>
                <div className="col-sm-2 col-xs-12 metadata">
                  {item.metadata.name}
                </div>
                <div className="col-sm-7 col-xs-12 graphs">
                  GRAPH
                </div>
                <div className="col-sm-2 col-xs-12 actions">
                  <div>
                    <Link to={'/projects/' + item.id}>View Details</Link>
                  </div>
                  <div>
                    Promote
                  </div>
                  <div>
                    Deploy
                  </div>
                  <div>
                    &hellip; More Actions
                  </div>
                </div>
              </li>
            );
          })}
        </ul>
      </div>
    );
  }
}