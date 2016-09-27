"use strict";
var __extends = (this && this.__extends) || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
};
var React = require('react');
require('../styles/input_feedback.scss');
(function (FeedbackType) {
    FeedbackType[FeedbackType["Progress"] = 0] = "Progress";
    FeedbackType[FeedbackType["Info"] = 1] = "Info";
    FeedbackType[FeedbackType["Confirm"] = 2] = "Confirm";
    FeedbackType[FeedbackType["Warning"] = 3] = "Warning";
    FeedbackType[FeedbackType["Error"] = 4] = "Error";
})(exports.FeedbackType || (exports.FeedbackType = {}));
var FeedbackType = exports.FeedbackType;
var InputFeedback = (function (_super) {
    __extends(InputFeedback, _super);
    function InputFeedback() {
        _super.call(this);
    }
    InputFeedback.prototype.render = function () {
        return (React.createElement("div", {className: "input-feedback feedback-progress"}, this.props.type === FeedbackType.Progress ?
            React.createElement("div", {className: "feedback-inner"}, React.createElement("div", {className: 'uil-facebook-css'}, React.createElement("div", null), React.createElement("div", null), React.createElement("div", null)), React.createElement("span", null, this.props.message))
            : null, this.props.type === FeedbackType.Info ?
            React.createElement("div", {className: "input-feedback feedback-info"}, React.createElement("span", null, this.props.message))
            : null, this.props.type === FeedbackType.Confirm ?
            React.createElement("div", {className: "input-feedback feedback-confirm"}, React.createElement("div", {className: "feedback-confirm-icon"}, React.createElement("i", {className: "fa fa-check-circle", "aria-hidden": "true"})), React.createElement("span", null, this.props.message))
            : null, this.props.type === FeedbackType.Warning ?
            React.createElement("div", {className: "input-feedback feedback-warning"}, React.createElement("div", {className: "feedback-warning-icon"}, React.createElement("i", {className: "fa fa-exclamation-triangle", "aria-hidden": "true"})), React.createElement("span", null, this.props.message))
            : null, this.props.type === FeedbackType.Error ?
            React.createElement("div", {className: "input-feedback feedback-error"}, React.createElement("div", {className: "feedback-error-icon"}, React.createElement("i", {className: "fa fa-exclamation-triangle", "aria-hidden": "true"})), React.createElement("span", null, this.props.message))
            : null));
    };
    return InputFeedback;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = InputFeedback;
//# sourceMappingURL=InputFeedback.js.map