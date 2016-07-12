/**
 * Created by justin on 7/12/16.
 */
import * as React from 'react';
import * as $ from 'jquery';
import * as classNames from 'classnames';
import Panel from './Panel';
import PageHeader from './PageHeader';
import ProgressBar from './ProgressBar';
import '../styles/newprojectstep3.scss';

export default class NewProjectStep3 extends React.Component<any, any> {
  progressBars: Array = [];
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
    }
  }
  onComplete() {
    console.log(this.progressBars);
    $(this).addClass('complete');
  }

  render() {
    return (
      <div className="new-project-step-3">
        <PageHeader>GOOD WORK!</PageHeader>
        <div>
          5 training jobs have been added to the Prithvi - 8 node cluster.
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
                    <ProgressBar ref={() => this.progressBars.push(this)} className={classNames({complete: this.state.jobs[i].isComplete})} showPercentage={true} onComplete={this.onComplete.bind(this)}/>
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
        </section>
      </div>
    );
  }
}
