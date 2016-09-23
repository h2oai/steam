/**
 * Created by justin on 7/10/16.
 */
import * as React from 'react';
import * as classNames from 'classnames';
import { Link } from 'react-router';
import '../styles/welcomesplashscreen.scss';

interface Props {
}

interface DispatchProps {
}

export default class WelcomeSplashScreen extends React.Component<Props & DispatchProps, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="welcome-splash-screen">
        <div className="welcome-splash-screen--content">
          <div>WELCOME TO</div>
          <div className="welcome-splash-screen--content--product-name">H<sub>2</sub>O STEAM</div>
          <div>Fast, distributed data science for teams</div>
          <div><Link to="/newproject"
                     className={classNames('default', 'start-project')}>Start A New Project</Link>
          </div>
        </div>
      </div>
    );
  }
}
