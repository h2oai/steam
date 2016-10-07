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
 * Created by justin on 7/10/16.
 */
import * as React from 'react';
import * as classNames from 'classnames';
import { Link } from 'react-router';
import '../styles/welcomesplashscreen.scss';
import PageHeader from './PageHeader';

interface Props {
}

interface DispatchProps {
}

export default class WelcomeSplashScreen extends React.Component<Props & DispatchProps, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="welcome-splash-screen">
        <div className="welcome-splash-screen--content">
          <h1>WELCOME TO</h1>
          <div className="page-header product-title">H<span className="half-size"><sub>2</sub></span>O STEAM</div>
          <h1>Fast, distributed data science for teams</h1>
          <div><Link to="/newproject"
                     className={classNames('default', 'start-project')}>Start A New Project</Link>
          </div>
        </div>
      </div>
    );
  }
}
