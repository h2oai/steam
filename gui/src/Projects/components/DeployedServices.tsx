/**
 * Created by justin on 7/12/16.
 */
import * as React from 'react';
import * as _ from 'lodash';
import Panel from './Panel';
import { fetchAllServices, killService, fetchServicesForProject } from '../actions/services.actions';
import { ScoringService } from '../../Proxy/Proxy';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/deployedservices.scss';

interface Props {
  services: {
    runningServicesForProject: ScoringService[],
    allRunningServices: ScoringService[]
  },
  projectId: string
}

interface DispatchProps {
  fetchAllServices: Function,
  killService: Function,
  fetchServicesForProject: Function
}


export class DeployedServices extends React.Component<Props & DispatchProps, any> {
  componentWillMount(): void {
    this.fetchServicesStrategy(this.props.projectId);
  }

  fetchServicesStrategy(projectId: string) {
    if (projectId) {
      return this.props.fetchServicesForProject(parseInt(projectId, 10));
    } else {
      return this.props.fetchAllServices();
    }
  }

  killService(serviceId) {
    this.props.killService(serviceId, parseInt(this.props.projectId, 10));
  }

  render(): React.ReactElement<HTMLDivElement> {
    let runningServices;
    if (this.props.projectId) {
      runningServices = this.props.services.runningServicesForProject;
    } else {
      runningServices = this.props.services.allRunningServices;
    }

    if (_.isEmpty(runningServices)) {
      return (
        <div>
          <h3>There are no services currently deployed.</h3>
        </div>
      );
    }

    return (
      <div className="deployed-services">
        <section>
          {runningServices.map((service, i) => {
            return (
              <Panel key={i} className="services-panel">
                <div className="panel-body">
                  <div className="panel-title">
                    <span>{service.name} @ <a href={'http://' + service.address + ':' + service.port} target="_blank" rel="noopener">{service.address + ':' + service.port}</a></span>
                    <div style={{color: service.state === 'stopped' ? 'red' : 'green'}}>{service.state}</div>
                  </div>
                  <div className="panel-info">
                    <div className="panel-info-row">
                      <span><i className="fa fa-cube"/></span><span>Model</span>
                      <span>{service.model_id}</span>
                    </div>
                    <div className="panel-info-row">
                      <span><i className="fa fa-folder-o"/></span><span>Status</span>
                      <span>{service.state === 'started' ? 'OK' : 'Error'}</span>
                    </div>
                  </div>
                </div>
                <div className="panel-actions">
                  <div className="panel-action" onClick={this.killService.bind(this, service.id)}>
                    <div><i className="fa fa-close"/></div>
                    <div>Stop Service</div>
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

function mapStateToProps(state): any {
  return {
    services: state.services
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    fetchAllServices: bindActionCreators(fetchAllServices, dispatch),
    fetchServicesForProject: bindActionCreators(fetchServicesForProject, dispatch),
    killService: bindActionCreators(killService, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(DeployedServices);
