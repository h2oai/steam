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
"use strict";
var __extends = (this && this.__extends) || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
};
/**
 * Created by justin on 7/11/16.
 * TODO(justinloyola): WIP
 */
var React = require('react');
var ReactDOM = require('react-dom');
var classNames = require('classnames');
var $ = require('jquery');
require('../styles/dropdown.scss');
var Dropdown = (function (_super) {
    __extends(Dropdown, _super);
    function Dropdown() {
        _super.call(this);
        this.state = {
            open: false,
        };
    }
    Dropdown.prototype.updateDropdownStyle = function () {
        $(this._mountNode).css({
            top: this.refs.dropdownInvokerContainer.getBoundingClientRect().bottom,
            left: this.refs.dropdownInvokerContainer.getBoundingClientRect().left,
            width: $(this.refs.dropdownInvokerContainer).width()
        });
    };
    Dropdown.prototype.componentWillMount = function () {
        var _this = this;
        this.appendToElement();
        $(document).bind('click.dropdown', function (event) {
            if (ReactDOM.findDOMNode(_this.refs.dropdownInvokerContainer).contains(event.target)) {
                _this.setState({
                    open: !_this.state.open
                });
            }
            else {
                _this.setState({
                    open: false
                });
            }
            _this.updateDropdownStyle();
        });
        $(window).bind('resize.dropdown', function () {
            _this.updateDropdownStyle();
        });
    };
    Dropdown.prototype.componentDidUpdate = function () {
        if (this._mountNode) {
            this.renderDropdown();
            if (this.state.open === true) {
                $(this._mountNode).addClass('open');
            }
            else {
                $(this._mountNode).removeClass('open');
            }
        }
    };
    Dropdown.prototype.componentWillUnmount = function () {
        $(document).unbind('click.dropdown');
        $(document).unbind('resize.dropdown');
        if (this._mountNode) {
            ReactDOM.unmountComponentAtNode(this._mountNode);
            this._mountNode = null;
            this._mountNode.remove();
        }
    };
    Dropdown.prototype.renderDropdown = function () {
        ReactDOM.unstable_renderSubtreeIntoContainer(this, this.getDropdown(), this._mountNode);
    };
    Dropdown.prototype.appendToElement = function () {
        this._mountNode = document.createElement('div');
        $(this._mountNode).addClass('dropdown-container');
        $(this._mountNode).addClass(this.props.className);
        var target = $(document.body);
        target.append(this._mountNode);
    };
    Dropdown.prototype.getDropdown = function () {
        return (React.createElement("div", {className: "dropdown-menu"}, this.props.dropdownContent));
    };
    Dropdown.prototype.render = function () {
        return (React.createElement("div", {ref: "dropdownInvokerContainer", className: classNames('dropdown', this.props.invokerContainerClass, { open: this.state.open })}, this.props.children));
    };
    return Dropdown;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = Dropdown;
//# sourceMappingURL=Dropdown.js.map