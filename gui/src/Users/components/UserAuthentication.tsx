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
import * as _ from 'lodash';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/users.scss';
import { Collapse } from '@blueprintjs/core/dist/components/collapse/collapse';
import { Button } from '@blueprintjs/core/dist/components/button/buttons';

interface Props {
}
interface DispatchProps {
}

export class UserAuthentication extends React.Component<Props & DispatchProps, any> {

  constructor(params) {
    super(params);
    this.state = {
      isLDAPConnectionSettingsOpen: true
    };
  }

  componentWillMount() {
  }

  onShowLDAPConnectionSettingsClicked = () => {
    this.setState({
      isLDAPConnectionSettingsOpen: !this.state.isLDAPConnectionSettingsOpen
    });
  };

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="user-authentication">
        User Authentication
        <Button onClick={this.onShowLDAPConnectionSettingsClicked}>
          {this.state.isOpen ? "Hide" : "Show"}
        </Button>
        <Collapse isOpen={true}>
          LDAP Connection Settings
        </Collapse>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
  };
}

function mapDispatchToProps(dispatch) {
  return {
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(UserAuthentication);
