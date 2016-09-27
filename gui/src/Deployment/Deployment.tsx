/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

/**
 * Created by justin on 7/11/16.
 */
import * as React from 'react';
import * as _ from 'lodash';
import PageHeader from '../Projects/components/PageHeader';
import TabNavigation from '../Projects/components/TabNavigation';
import DeployedServices from '../Projects/components/DeployedServices';
import Packaging from './components/Packaging';
import UploadPreProcessingModal from './components/UploadPreProcessingModal';
import { connect } from 'react-redux';
import { uploadPackage } from './actions/deployment.actions';
import { bindActionCreators } from 'redux';
import './styles/deployment.scss';

interface Props {
  params: {
    projectid: string
  },
  packages: {
    packages: string[]
  }
}

interface DispatchProps {
  uploadPackage: Function
}

export class Deployment extends React.Component<Props & DispatchProps, any> {
  constructor() {
    super();
    this.state = {
      tabs: {
        deployedServices: {
          label: 'DEPLOYED SERVICES',
          isSelected: true,
          onClick: this.clickHandler.bind(this),
          component: <DeployedServices/>
        },
        packaging: {
          label: 'PACKAGING',
          isSelected: false,
          onClick: this.clickHandler.bind(this),
          component: <Packaging/>
        }
      },
      isSelected: 'deployedServices',
      uploadOpen: false,
      packages: [],
      projectId: null
    };
  }

  componentWillMount() {
    this.setState({
      tabs: {
        deployedServices: {
          label: 'DEPLOYED SERVICES',
          isSelected: true,
          onClick: this.clickHandler.bind(this),
          component: <DeployedServices/>
        },
        packaging: {
          label: 'PACKAGING',
          isSelected: false,
          onClick: this.clickHandler.bind(this),
          component: <Packaging projectId={this.props.params.projectid}/>
        }
      }
    });
  }

  clickHandler(tab) {
    let key = _.findKey(this.state.tabs, tab);
    let newState = _.cloneDeep(this.state);
    Object.keys(newState.tabs).map((tab) => {
      newState.tabs[tab].isSelected = false;
    });
    newState.tabs[key].isSelected = true;
    newState.isSelected = key;
    this.setState(newState);
  }

  openUpload() {
    this.setState({
      uploadOpen: true
    });
  }

  closeUpload() {
    this.setState({
      uploadOpen: false
    });
  }

  upload(event, uploadedPackage, formData) {
    event.preventDefault();
    this.props.uploadPackage(parseInt(this.props.params.projectid, 10), uploadedPackage.name, formData);
    this.closeUpload();
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="services">
        {!_.isUndefined(this.props.params.projectid) ?
          <UploadPreProcessingModal open={this.state.uploadOpen} cancel={this.closeUpload.bind(this)}
                                    upload={this.upload.bind(this)}/> : null}
        <PageHeader>
          <span>Deployment</span>
          {!_.isUndefined(this.props.params.projectid) ? <div className="button-primary header-buttons"
                                                                     onClick={this.openUpload.bind(this)}>Upload New Package</div> : null}
        </PageHeader>
        {!_.isUndefined(this.props.params.projectid) ? <TabNavigation tabs={this.state.tabs}/> : null}
        <main>
          {this.state.tabs.deployedServices.isSelected === true ?
            <DeployedServices projectId={this.props.params.projectid}/> : null}
          {this.state.tabs.packaging.isSelected === true && !_.isUndefined(this.props.params.projectid) ?
            <Packaging projectId={this.props.params.projectid}/> : null}
        </main>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    packages: state.packages
  };
}

function mapDispatchToProps(dispatch) {
  return {
    uploadPackage: bindActionCreators(uploadPackage, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Deployment);
