/**
 * Created by justin on 6/27/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import { Link } from 'react-router';
import Deploy from '../components/Deploy';
import RocGraph from '../components/RocGraph';
import CornerBorders from '../components/CornerBorders';
import PageHeader from '../../Projects/components/PageHeader';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import '../styles/leaderboard.scss';

interface Props {
  items: any[]
}

interface DispatchProps {
}

export default class Leaderboard extends React.Component<Props & DispatchProps, any> {
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

  openDeploy(): void {
    this.setState({
      isDeployOpen: true
    });
  }

  closeHandler(): void {
    this.setState({
      isDeployOpen: false
    });
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="leaderboard">
        <PageHeader>Models</PageHeader>
        <Deploy open={this.state.isDeployOpen} closeHandler={this.closeHandler}></Deploy>
        <Table>
          <Row header={true}>
            <Cell>
              <i className="fa fa-caret-down"/>
            </Cell>
            <Cell>
              MODEL
            </Cell>
            <Cell className="graph">
              TRAIN ROC
            </Cell>
            <Cell className="graph">
              TEST ROC
            </Cell>
            <Cell className="graph">
              RESIDUALS
            </Cell>
            <Cell>
              <div className="actions">
                ACTIONS
              </div>
            </Cell>
          </Row>
          {this.props.items.map((item, i) => {
            return (
              <Row key={i}>
                <Cell>{item.rank + Leaderboard.getOrdinal(item.rank)}</Cell>
                <Cell>
                  <div className="metadata">
                    <div className="model-name">
                      {item.metadata.modelName}
                    </div>
                    <div>
                      {item.metadata.createdBy} 
                    </div>
                    <div>
                      {item.metadata.creationDate}
                    </div>
                    <div>
                      {item.metadata.timing}
                    </div>
                  </div>
                </Cell>
                <Cell className="graph">
                  <RocGraph/>
                </Cell>
                <Cell className="graph">
                  <RocGraph/>
                </Cell>
                <Cell className="graph">
                  <RocGraph/>
                </Cell>
                <Cell>
                  <ul className="actions">
                    <li><Link to={"models/" + item.id}><span><i className="fa fa-eye"></i></span><span>view model details</span></Link></li>
                    <li><span><i className="fa fa-database"></i></span><span>designate as baseline</span></li>
                    <li onClick={this.openDeploy}><span><i className="fa fa-arrow-up"></i></span><span>deploy model</span></li>
                    <li><span><i className="fa fa-ellipsis-h"></i></span><span>more actions</span></li>
                  </ul>
                </Cell>
              </Row>
            );
          })}
        </Table>
      </div>
    );
  }
}