/**
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import '../styles/pageheader.scss';

interface Props {
  children?: React.ReactChildren
}

export default class PageHeader extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLElement> {
    return (
      <header className="page-header">{this.props.children}</header>  
    );
  }
}