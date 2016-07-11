/**
 * Created by justin on 7/10/16.
 */
import * as React from 'react';
import '../styles/welcomesplashscreen.scss';

export default class WelcomeSplashScreen extends React.Component<any, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="welcome-splash-screen">
        <div className="welcome-splash-screen--content">
          <div>WELCOME TO</div>
          <div className="welcome-splash-screen--content--product-name">H<sub>2</sub>O STEAM</div>
          <div>Fast, distributed taglines for teams</div>
          <button type="button">Start A New Project</button>
        </div>
      </div>
    )
  }
}
