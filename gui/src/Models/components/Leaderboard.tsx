/**
 * Created by justin on 6/27/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import { Link } from 'react-router';
import Deploy from '../components/Deploy';
import RocGraph from '../components/RocGraph';
import PageHeader from '../../Projects/components/PageHeader';
import Pagination from '../components/Pagination';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import { getOrdinal } from '../../App/utils/getOrdinal';
import '../styles/leaderboard.scss';

// sample data
import { deeplearningTrain } from '../tests/data/deeplearningTrain';
import { deeplearningValidation } from '../tests/data/deeplearningValidation';
import { drfTrain } from '../tests/data/drfTrain';
import { drfValidation } from '../tests/data/drfValidation';
import { gbmTrain } from '../tests/data/gbmTrain';
import { gbmValidation } from '../tests/data/gbmValidation';
import { glmTrain } from '../tests/data/glmTrain';
import { glmValidation } from '../tests/data/glmValidation';
import { naivebayesTrain } from '../tests/data/naivebayesTrain';
import { naivebayesValidation } from '../tests/data/naivebayesValidation';

interface Props {
  items: any[],
  projectId: number
}

interface DispatchProps {
}

export default class Leaderboard extends React.Component<Props & DispatchProps, any> {

  sampleData = {};

  constructor() {
    super();
    this.state = {
      isDeployOpen: false
    };
    this.openDeploy = this.openDeploy.bind(this);
    this.closeHandler = this.closeHandler.bind(this);
    this.sampleData = {
      deeplearningTrain,
      deeplearningValidation,
      drfTrain,
      drfValidation,
      gbmTrain,
      gbmValidation,
      glmTrain,
      glmValidation,
      naivebayesTrain,
      naivebayesValidation
    };
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
        <Deploy open={this.state.isDeployOpen} closeHandler={this.closeHandler}></Deploy>
        <PageHeader>
          <span>Models</span>
          <div className="buttons">
            <button className="default">Import Model</button>
          </div>
        </PageHeader>
        <div className="filter">
          <input type="text" placeholder="filter models"/>
        </div>
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
            <Cell>
              <div className="actions">
                ACTIONS
              </div>
            </Cell>
          </Row>
          {this.props.items.map((item, i) => {
            return (
              <Row key={i}>
                <Cell>{item.id + getOrdinal(item.id)}</Cell>
                <Cell>
                  <div className="metadata">
                    <div className="model-name">
                      {item.name}
                    </div>
                    <div>
                      {item.cluster_name}
                    </div>
                    <div>
                      {item.createdAt}
                    </div>
                    <div>
                      {item.max_runtime}
                    </div>
                  </div>
                </Cell>
                <Cell className="graph">
                  <RocGraph data={this.sampleData['gbmTrain']}/>
                </Cell>
                <Cell className="graph">
                  <RocGraph data={this.sampleData['gbmValidation']}/>
                </Cell>
                <Cell>
                  <ul className="actions">
                    <li><Link to={'/projects/' + this.props.projectId + '/models/' + item.id}><span><i className="fa fa-eye"></i></span><span>view model details</span></Link></li>
                    <li onClick={this.openDeploy}><span><i className="fa fa-arrow-up"></i></span><span>deploy model</span></li>
                    <li><span><i className="fa fa-ellipsis-h"></i></span><span>more actions</span></li>
                  </ul>
                </Cell>
              </Row>
            );
          })}
        </Table>
        <Pagination items={this.props.items}></Pagination>
      </div>
    );
  }
}
