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
  text: string,
  options: any
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
      if (event.target === this.refs.dropdownInvoker) {
        this.setState({
          open: !this.state.open
        });
      } else {
        this.setState({
          open: false
        });
        this.updateDropdownStyle();
      }
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
    }
  }

  renderDropdown() {
    ReactDOM.unstable_renderSubtreeIntoContainer(this, this.getDropdown(), this._mountNode);
  }

  appendToElement() {
    this._mountNode = document.createElement('div');
    $(this._mountNode).addClass('dropdown-container');
    document.body.appendChild(this._mountNode);
  }

  getDropdown() {
    return (
      <div className="dropdown-menu">
        <ul>
          {this.props.options.map((option, i) => {
            return <li key={i} value={option.value}>{option.text}</li>;
          })}
        </ul>
      </div>
    );
  }

  render() {
    return (
      <div ref="dropdownInvokerContainer" className={classNames('dropdown', { open: this.state.open})}>
        <button ref="dropdownInvoker" type="button">{this.props.text}</button>
      </div>
    );
  }
}
