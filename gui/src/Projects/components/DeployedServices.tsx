/**
 * Created by justin on 7/12/16.
 */
import * as React from 'react';
import Panel from './Panel';
import '../styles/deployedservices.scss';

export default class DeployedServices extends React.Component<any, any> {
  render(): React.ReactElement<HTMLDivElement> {
    let services = [
      {
        name: 'Classifier Service',
        model: 'DRF-1069085',
        project: 'Churn Prediction'
      },
      {
        name: 'Classifier Service',
        model: 'DRF-1069085',
        project: 'Churn Prediction'
      },
      {
        name: 'Classifier Service',
        model: 'DRF-1069085',
        project: 'Churn Prediction'
      }
    ];
    return (
      <div className="deployed-services">
        <header>Services deployed from this project.</header>
        <section>
          {services.map((service, i) => {
            return (
              <Panel key={i} className="services-panel">
                <div className="panel-body">
                  <div className="panel-title">
                    {service.name} @ <a href="http://localhost:54321">localhost:54321</a>
                  </div>
                  <div className="panel-info">
                    <div className="panel-info-row">
                      <span><i className="fa fa-cube"/></span><span>Model</span>
                      <span>{service.model}</span>
                    </div>
                    <div className="panel-info-row">
                      <span><i className="fa fa-folder-o"/></span><span>Project</span>
                      <span>{service.project}</span>
                    </div>
                  </div>
                </div>
                <div className="panel-actions">
                  <div className="panel-action">
                    <div><i className="fa fa-eye"/></div>
                    <div>View Details</div>
                  </div>
                  <div className="panel-action">
                    <div><i className="fa fa-close"/></div>
                    <div>Kill Service</div>
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
