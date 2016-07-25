/**
 * Created by justin on 7/12/16.
 */
import * as React from 'react';
import * as _ from 'lodash';
import Panel from './Panel';
import { fetchServices, killService } from '../actions/services.actions';
import '../styles/deployedservices.scss';
import { ScoringService } from '../../Proxy/proxy';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

interface Props {
  services: {
    runningServices: ScoringService[]
  }
}

interface DispatchProps {
  fetchServices: Function,
  killService: Function
}


export class DeployedServices extends React.Component<Props & DispatchProps, any> {
  componentWillMount(): void {
    if (_.isEmpty(this.props.services.runningServices)) {
      this.props.fetchServices();
    }
  }

  killService(serviceId) {
    this.props.killService(serviceId);
  }

  render(): React.ReactElement<HTMLDivElement> {
    if (_.isEmpty(this.props.services.runningServices)) {
      return <div></div>;
    }
    return (
      <div className="deployed-services">
        <section>
          {this.props.services.runningServices.map((service, i) => {
            return (
              <Panel key={i} className="services-panel">
                <div className="panel-body">
                  <div className="panel-title">
                    <span>N/A @ <a href={'http://' + service.address + ':' + service.port} target="_blank" rel="noopener">{service.address + ':' + service.port}</a></span>
                    <div style={{color: service.state === 'stopped' ? 'red' : 'green'}}>{service.state}</div>
                  </div>
                  <div className="panel-info">
                    <div className="panel-info-row">
                      <span><i className="fa fa-cube"/></span><span>Model ID</span>
                      <span>{service.model_id}</span>
                    </div>
                    <div className="panel-info-row">
                      <span><i className="fa fa-folder-o"/></span><span>Project</span>
                      <span>N/A</span>
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
    fetchServices: bindActionCreators(fetchServices, dispatch),
    killService: bindActionCreators(killService, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(DeployedServices);
