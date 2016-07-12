/**
 * Created by justin on 7/11/16.
 */
import * as React from 'react';
import * as classNames from 'classnames';
import '../styles/tabnavigation.scss';

interface Props {
  tabs: {
    label: string
    isSelected: boolean,
    onClick?: Function
  }[],
}

export default class TabNavigation extends React.Component<Props, any> {
  render(): React.ReactElement<HTMLElement> {
    return (
      <nav className="tabs">
        {Object.keys(this.props.tabs).map((tab, i) => {
          return (
            <a key={i} className={classNames('tab', {selected: this.props.tabs[tab].isSelected === true})} onClick={this.props.tabs[tab].onClick ? this.props.tabs[tab].onClick.bind(this, this.props.tabs[tab]) : null}>
              {this.props.tabs[tab].label}
            </a>
          );
        })}
      </nav>
    );
  }
}
