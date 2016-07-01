/**
 * Created by justin on 6/27/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import { Link } from 'react-router';
import Deploy from '../components/Deploy';
import '../styles/leaderboard.scss';

interface Props {
  items: any[]
}

export class Leaderboard extends React.Component<Props, any> {
  constructor() {
    super();
    this.state = {
      isDeployOpen: false
    };
    this.openDeploy = this.openDeploy.bind(this);
    this.closeHandler = this.closeHandler.bind(this);
  }
  static getOrdinal(rank: number): string {
    let suffixes = ['th', 'st', 'nd', 'rd'];
    let remainder = rank % 100;
    return (suffixes[(remainder - 20) % 10] || suffixes[remainder] || suffixes[0]);
  }

  openDeploy() {
    this.setState({
      isDeployOpen: true
    });
  }

  closeHandler() {
    this.setState({
      isDeployOpen: false
    });
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="leaderboard">
        <header>
          MODEL LEADERBOARD
        </header>
        <Deploy open={this.state.isDeployOpen} closeHandler={this.closeHandler}></Deploy>
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
                  <Link className="action" to={'/projects/' + item.id}><span><i className="fa fa-eye"></i></span><span className="action-label">View Details</span></Link>
                  <div className="action">
                    <span><i className="fa fa-arrow-up"></i></span><span className="action-label">Promote</span>
                  </div>
                  <div className="action" onClick={this.openDeploy}>
                    <span><i className="fa fa-database"></i></span><span className="action-label">Deploy</span>
                  </div>
                  <div className="action">
                    <span><i className="fa fa-ellipsis-h"></i></span><span className="action-label">More Actions</span>
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