/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import './styles/body.scss';

interface Props {
}

interface DispatchProps {
}

export default class Body extends React.Component<Props & DispatchProps, any> {
  render(): React.ReactElement<HTMLElement> {
    return (
      <section className="main-section">
        <article className="content">
          {this.props.children}
        </article>
      </section>
    );
  }
}
