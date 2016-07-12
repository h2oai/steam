/**
 * Created by justin on 7/11/16.
 */
import * as React from 'react';
import PageHeader from './PageHeader';
import '../styles/deployments.scss';

export default class Deployments extends React.Component<any, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div>
        <PageHeader>DEPLOYMENT</PageHeader>
        <nav className="tabs">
          <a className="tab selected">
            Deployed Services
          </a>
          <a className="tab">
            PACKAGING
          </a>
          <a className="tab">
            MODEL API
          </a>
        </nav>
        <main>
          CONTENT
        </main>
      </div>
    );
  }
}
