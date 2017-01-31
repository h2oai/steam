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

import * as React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/user.scss';
import { } from "../../Proxy/Proxy";
import {fetchUserKeytab, deleteKeytab, testKeytab, saveLocalKerberos} from "../actions/user.actions";
import {Keytab} from "../../Proxy/Proxy";
import { Dialog } from '@blueprintjs/core/dist/components/dialog/dialog';

interface Props {
  userKeytab: Keytab
}
interface DispatchProps {
  fetchUserKeytab: Function,
  deleteKeytab: Function,
  testKeytab: Function,
  saveLocalKerberos: Function
}

export class UserKerberos extends React.Component<Props & DispatchProps, any> {

  refs: {
    [key: string]: (Element);
    keytab: (HTMLInputElement)
  };

  constructor(params) {
    super(params);
    this.state = {
      uploadText: "Upload New Keytab",
      afterConfirmAction: null
    };

  }

  componentWillMount() {
    this.props.fetchUserKeytab();
  }

  onDeleteKeytab = (id) => {
    this.props.deleteKeytab(id, false);
    this.setState({afterConfirmAction: null});
  };
  onTestConfigClicked = () => {
    this.props.testKeytab(this.props.userKeytab.id);
  };
  onNewKeytabSelected = (e) => {
    this.setState({
      uploadText: e.target.value
    });
    this.props.saveLocalKerberos(this.refs.keytab);
  };

  render(): React.ReactElement<HTMLDivElement> {

    return (
      <div className="cluster-authentication intro">
        <table className="space-20">
          <tbody>
            <tr>
              <td className="auth-left">PRINCIPLE KEYTAB</td>
              <td>
                <p>Your principle keytab</p>
                {this.props.userKeytab ? <p>{this.props.userKeytab.name} &nbsp; <i className="fa fa-times" aria-hidden="true" onClick={() => this.setState({afterConfirmAction: () => this.onDeleteKeytab(this.props.userKeytab.id)})}></i></p>
                  : <p>
                    <label className="pt-file-upload">
                      <input ref="keytab" type="file" onChange={(e) => this.onNewKeytabSelected(e)} />
                      <span className="pt-file-upload-input">{this.state.uploadText}</span>
                    </label>
                  </p> }
              </td>
            </tr>
          </tbody>
        </table>

        {this.props.userKeytab ? <div id="actionButtonsContainer" className="space-20">
            <div>
              <div className="button-secondary" onClick={this.onTestConfigClicked}>Test Config</div>
            </div>
        </div> : null}

        <Dialog
          iconName="Confirm"
          isOpen={this.state.afterConfirmAction}
          onClose={() => this.setState({ afterConfirmAction: null })}
          title="Confirm exit"
        >
          <div className="pt-dialog-body">
            Are you sure you wish to delete this keytab?
          </div>
          <div className="pt-dialog-footer">
            <div className="pt-dialog-footer-actions">
              <div className="button-secondary" onClick={() => this.setState({ afterConfirmAction: null })}>Cancel</div> &nbsp;
              <div className="button-primary" onClick={this.state.afterConfirmAction}>Accept</div>
            </div>
          </div>
        </Dialog>

      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    userKeytab: state.user.userKeytab
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchUserKeytab: bindActionCreators(fetchUserKeytab, dispatch),
    deleteKeytab: bindActionCreators(deleteKeytab, dispatch),
    testKeytab: bindActionCreators(testKeytab, dispatch),
    saveLocalKerberos: bindActionCreators(saveLocalKerberos, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(UserKerberos);
