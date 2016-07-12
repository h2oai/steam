/**
 * Created by justin on 7/12/16.
 */
import * as React from 'react';
import '../styles/deployedservices.scss';

export default class DeployedServices extends React.Component<any, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="deployed-services">
        <header>Services deployed from this project.</header>
        <section>
          <div className="panel">
            <div className="panel-body">
              <div className="panel-title">
                Classifier Service @ <a href="http://localhost:54321">localhost:54321</a>
              </div>
              <div className="panel-info">
                <div>
                  <span><i className="fa fa-cube"/></span><span>Model</span>
                  <span>DRF-1069085</span>
                </div>
                <div>
                  <span><i className="fa fa-folder-o"/></span><span>Project</span>
                  <span>Churn Prediction</span>
                </div>
              </div>
            </div>
            <div className="panel-actions">
              <div className="panel-action">
                <div><i className="fa fa-eye"/></div><div>View Details</div>
              </div>
              <div className="panel-action">
                <div><i className="fa fa-close"/></div><div>Kill Service</div>
              </div>
            </div>
          </div>
        </section>
      </div>
    );
  }
}
