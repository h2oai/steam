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
 * Created by justin on 7/27/16.
 */
var React = require('react');
var $ = require('jquery');
var classNames = require('classnames');
var _ = require('lodash');
var DefaultModal_1 = require('../../App/components/DefaultModal');
var Table_1 = require('../../Projects/components/Table');
var Row_1 = require('../../Projects/components/Row');
var Cell_1 = require('../../Projects/components/Cell');
require('../styles/uploadpreprocessingmodal.scss');
var UploadPreProcessingModal = (function (_super) {
    __extends(UploadPreProcessingModal, _super);
    function UploadPreProcessingModal() {
        _super.call(this);
        this.state = {
            mainFiles: '',
            libraryFiles: [],
            missingPackageNameError: false
        };
    }
    UploadPreProcessingModal.prototype.selectMain = function () {
        $('input[name="selectMain"]').click();
    };
    UploadPreProcessingModal.prototype.selectLibraries = function () {
        $('input[name="selectLibraries"]').click();
    };
    UploadPreProcessingModal.prototype.selectMainHandler = function (event) {
        this.setState({
            mainFiles: event.target.files[0]
        });
    };
    UploadPreProcessingModal.prototype.selectLibrariesHandler = function (event) {
        this.setState({
            libraryFiles: Array.prototype.slice.call(event.target.files)
        });
    };
    UploadPreProcessingModal.prototype.uploadPackage = function (event) {
        if (_.isEmpty(this.refs.packageName.value)) {
            this.setState({
                missingPackageNameError: true
            });
            event.preventDefault();
            return false;
        }
        var uploadedPackage = {
            name: this.refs.packageName.value
        };
        this.props.upload(event, uploadedPackage, this.refs.uploadForm);
    };
    UploadPreProcessingModal.prototype.render = function () {
        return (React.createElement(DefaultModal_1.default, {className: "upload-preprocessing-modal", open: this.props.open}, React.createElement("header", {className: "page-header"}, "UPLOAD PRE-PROCESSING PACKAGE (PYTHON)"), React.createElement("section", null, React.createElement("form", {ref: "uploadForm", onSubmit: this.uploadPackage.bind(this)}, React.createElement(Table_1.default, null, React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "SELECT PYTHON MAIN"), React.createElement(Cell_1.default, null, React.createElement("div", null, "Select a main Python file for pre-processing."), React.createElement("span", {className: "muted"}, "The output from this Python file should be one row of an H2O data from that your model is expecting."), React.createElement("div", {className: "upload"}, React.createElement("div", {className: "upload-info", onClick: this.selectMain.bind(this)}, React.createElement("span", null, React.createElement("i", {className: "fa fa-folder-o"})), React.createElement("span", {className: "file-list"}, this.state.mainFiles ? this.state.mainFiles.name : 'N/A'), React.createElement("span", null, React.createElement("i", {className: "fa fa-close"})), React.createElement("input", {type: "file", name: "selectMain", onChange: this.selectMainHandler.bind(this)}))))), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "SELECT PYTHON LIBRARIES"), React.createElement(Cell_1.default, null, React.createElement("div", null, "Select a one or more Python files for your library."), React.createElement("span", {className: "muted"}, "Any non-standard libraries called here should be installed into your deployment environment prior to launching services."), React.createElement("div", {className: "upload"}, React.createElement("div", {className: "upload-info", onClick: this.selectLibraries.bind(this)}, React.createElement("span", null, React.createElement("i", {className: "fa fa-folder-o"})), React.createElement("span", {className: "file-list"}, this.state.libraryFiles.length > 0 ? this.state.libraryFiles.map(function (file, i) {
            return React.createElement("div", {key: i}, file.name);
        }) : 'N/A'), React.createElement("span", null, React.createElement("i", {className: "fa fa-close"})), React.createElement("input", {type: "file", name: "selectLibraries", onChange: this.selectLibrariesHandler.bind(this), multiple: true}))))), React.createElement(Row_1.default, null, React.createElement(Cell_1.default, null, "NAME THE PACKAGE"), React.createElement(Cell_1.default, null, React.createElement("div", null, "Pick a name for this pre-processing package. You will use it as a reference when deploying models."), React.createElement("div", {className: "package-name-label muted"}, "Package name"), React.createElement("input", {ref: "packageName", type: "text", className: classNames('package-name', { error: this.state.missingPackageNameError })}))), React.createElement(Row_1.default, {className: "button-row"}, React.createElement(Cell_1.default, null), React.createElement(Cell_1.default, null, React.createElement("button", {type: "submit", className: "button-primary", onClick: this.uploadPackage.bind(this)}, "Upload"), React.createElement("button", {className: "button-secondary", onClick: this.props.cancel.bind(this)}, "Cancel"))))))));
    };
    return UploadPreProcessingModal;
}(React.Component));
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = UploadPreProcessingModal;
//# sourceMappingURL=UploadPreProcessingModal.js.map