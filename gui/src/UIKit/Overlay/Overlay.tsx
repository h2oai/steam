/**
 * Created by justin on 6/30/16.
 */
import * as React from 'react';
import * as ReactDOM from 'react-dom';
import './styles/overlay.scss';

interface Props {
  target?: string,
  placement?: string,
  appendToBody?: boolean,
  className?: any,
  open: boolean
}

export default class Overlay extends React.Component<Props, any> {
  _mountNode: HTMLElement;

  componentDidUpdate() {
    if (this._mountNode) {
      this.renderOverlay();
    }
  }

  componentWillMount() {
    this.appendToElement(this.getOverlay());
  }

  componentWillUnmount() {
    if (this._mountNode) {
      ReactDOM.unmountComponentAtNode(this._mountNode);
      this._mountNode = null;
    }
  }

  appendToElement(element) {
    let placement = this.props.placement || 'top';
    let target = this.props.target || 'body';
    if (this.props.appendToBody === true) {
      target = 'body';
    }
    let targetNode = document.querySelector(target);
    if (targetNode === null) {
      throw 'No such target element found';
    }
    this._mountNode = document.createElement('div');
    this._mountNode.className = 'overlay-container';
    if (placement === 'bottom') {
      this._mountNode.style.bottom = '0';
    } else {
      this._mountNode.style.top = '0';
    }
    targetNode.appendChild(this._mountNode);
  }

  renderOverlay() {
    ReactDOM.unstable_renderSubtreeIntoContainer(this, this.getOverlay(), this._mountNode);
  }

  getOverlay() {
    if (this._mountNode) {
      if (this.props.open === true) {
        this._mountNode.classList.add('open');
      } else {
        this._mountNode.classList.remove('open');
      }
    }
    return (
      <div className="overlay">
        {this.props.children}
      </div>
    );
  }

  render(): React.ReactElement<HTMLDivElement> {
    return null;
  }
}