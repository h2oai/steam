/**
 * Created by justin on 7/12/16.
 */
import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { Link } from 'react-router';
import * as $ from 'jquery';
import { hashHistory } from 'react-router';
import Panel from './Panel';
import PageHeader from './PageHeader';
import ProgressBar from './ProgressBar';
import '../styles/newprojectstep3.scss';

export default class NewProjectStep3 extends React.Component<any, any> {
  constructor() {
    super();
    let jobs = [
      {
        name: 'DRF-1070196',
        project: 'Churn Prediction',
        author: 'Mark Landry',
        startTime: new Date().getTime()
      },
      {
        name: 'DRF-1070196',
        project: 'Churn Prediction',
        author: 'Mark Landry',
        startTime: new Date().getTime()
      },
      {
        name: 'DRF-1070196',
        project: 'Churn Prediction',
        author: 'Mark Landry',
        startTime: new Date().getTime()
      },
      {
        name: 'DRF-1070196',
        project: 'Churn Prediction',
        author: 'Mark Landry',
        startTime: new Date().getTime()
      }
    ];
    this.state = {
      jobs: jobs
    };
  }

  onComplete(progressBar) {
    let node = ReactDOM.findDOMNode(progressBar);
    $(node).addClass('progress-button');
    $(node).find('.progress-counter').text('Completed');
  }

  onClick() {
    hashHistory.push('/models/0');
  }

  render() {
    return (
      <div className="new-project-step-3">
        <PageHeader>GOOD WORK!</PageHeader>
        <div className="sub-title">
          5 training jobs have been added to the <span>Prithvi - 8 node</span> cluster.
        </div>
        <section>
          {this.state.jobs.map((job, i) => {
            return (
              <Panel key={i}>
                <div className="panel-body">
                  <div className="panel-title">
                    Training Job: {job.name} from {job.project}
                    <div className="panel-sub-title">
                      Started {job.startTime} by {job.author}
                    </div>
                  </div>
                  <div className="panel-info">
                    <ProgressBar showPercentage={true} onComplete={this.onComplete.bind(this)}
                                 onClick={this.onClick.bind(this)}>
                    </ProgressBar>
                  </div>
                </div>
                <div className="panel-actions">
                  <div className="panel-action">
                    <div><i className="fa fa-pause"/></div>
                    <div>Pause</div>
                  </div>
                  <div className="panel-action">
                    <div><i className="fa fa-stop"/></div>
                    <div>Cancel</div>
                  </div>
                </div>
              </Panel>
            );
          })}
          <Link to="/models" className="default link-leaderboard">Return to Model Leaderboard</Link><Link to="/projects/deployments">See all jobs on Prithbi - 8 node</Link>
        </section>
      </div>
    );
  }
}
