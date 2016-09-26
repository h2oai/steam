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
 * Created by justin on 7/11/16.
 * TODO(justinloyola): WIP
 */
import * as React from 'react';
import * as ReactDOM from 'react-dom';
import * as classNames from 'classnames';
import * as $ from 'jquery';
import '../styles/dropdown.scss';

interface Props {
  dropdownContent: Element | React.ReactElement<any>,
  className?: any,
  invokerContainerClass: any,
  target?: Element
}

export default class Dropdown extends React.Component<Props, any> {
  _mountNode: HTMLElement;
  container: HTMLElement;
  refs: {
    [key: string]: Element;
    dropdownInvokerContainer: HTMLDivElement,
    dropdownInvoker: HTMLButtonElement
  };

  constructor() {
    super();
    this.state = {
      open: false,
    };
  }

  updateDropdownStyle() {
    $(this._mountNode).css({
      top: this.refs.dropdownInvokerContainer.getBoundingClientRect().bottom,
      left: this.refs.dropdownInvokerContainer.getBoundingClientRect().left,
      width: $(this.refs.dropdownInvokerContainer).width()
    });
  }

  componentWillMount() {
    this.appendToElement();
    $(document).bind('click.dropdown', (event) => {
      if (ReactDOM.findDOMNode(this.refs.dropdownInvokerContainer).contains(event.target)) {
        this.setState({
          open: !this.state.open
        });
      } else {
        this.setState({
          open: false
        });
      }
      this.updateDropdownStyle();
    });
    $(window).bind('resize.dropdown', () => {
      this.updateDropdownStyle();
    });
  }

  componentDidUpdate() {
    if (this._mountNode) {
      this.renderDropdown();
      if (this.state.open === true) {
        $(this._mountNode).addClass('open');
      } else {
        $(this._mountNode).removeClass('open');
      }
    }
  }

  componentWillUnmount() {
    $(document).unbind('click.dropdown');
    $(document).unbind('resize.dropdown');
    if (this._mountNode) {
      ReactDOM.unmountComponentAtNode(this._mountNode);
      this._mountNode = null;
      this._mountNode.remove();
    }
  }

  renderDropdown() {
    ReactDOM.unstable_renderSubtreeIntoContainer(this, this.getDropdown(), this._mountNode);
  }

  appendToElement() {
    this._mountNode = document.createElement('div');
    $(this._mountNode).addClass('dropdown-container');
    $(this._mountNode).addClass(this.props.className);
    let target = $(document.body);
    target.append(this._mountNode);
  }

  getDropdown() {
    return (
      <div className="dropdown-menu">
        {this.props.dropdownContent}
      </div>
    );
  }

  render() {
    return (
      <div ref="dropdownInvokerContainer" className={classNames('dropdown', this.props.invokerContainerClass, { open: this.state.open})}>
        {this.props.children}
      </div>
    );
  }
}
