/**
 * Created by justin on 6/25/16.
 */

import * as React from 'react';
import {bindActionCreators} from 'redux';
import { connect } from 'react-redux';
import * as classNames from 'classnames';
import HamburgerMenu from '../Navigation/components/HamburgerMenu/HamburgerMenu';
import { navigationReducer, toggleMenu } from '../Navigation/components/Navigation/reducers/navigation.reducer';
import './styles/body.scss';

interface Props {
}

interface DispatchProps {
  toggleMenu: Function
}

export class Body extends React.Component<Props & DispatchProps, any> {
  constructor() {
    super();
    this.state = {
      isActive: false
    };
    this.clickHandler = this.clickHandler.bind(this);
  }
  clickHandler(event: React.MouseEvent): void {
    this.props.toggleMenu(!this.state.isActive);
    this.setState({
      isActive: !this.state.isActive
    });
  }
  
  render(): React.ReactElement<HTMLElement> {
    return (
      <section className="body-container">
        <div className="menu-container">
          <HamburgerMenu onClick={this.clickHandler} className={{active: this.state.isActive}}></HamburgerMenu>
        </div>
        <div className="content">
          {this.props.children}
        </div>
      </section>
    );
  }
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    toggleMenu: bindActionCreators(toggleMenu, dispatch)
  };
}

function mapStateToProps(state: Props): Props {
  return {};
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(Body);