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
 * Created by justin on 7/27/16.
 */
import * as React from 'react';
import * as $ from 'jquery';
import * as classNames from 'classnames';
import * as _ from 'lodash';
import DefaultModal from '../../App/components/DefaultModal';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import '../styles/uploadpreprocessingmodal.scss';


interface Props {
  open: boolean,
  upload: Function,
  cancel: Function
}

export default class UploadPreProcessingModal extends React.Component<Props, any> {
  refs: {
    [key: string]: Element
    packageName: HTMLInputElement,
    uploadForm: HTMLFormElement
  };

  constructor() {
    super();
    this.state = {
      mainFiles: '',
      libraryFiles: [],
      condaFiles: [],
      missingPackageNameError: false,
      packageNamed: false,
      showTooltipMain: false,
      showTooltipLibrary: false,
      showTooltipConfig: false,
      showTooltipName: false
    };
  }

  selectMain() {
    $('input[name="selectMain"]').click();
  }

  selectLibraries() {
    $('input[name="selectLibraries"]').click();
  }

  selectConda() {
    $('input[name="selectConda"]').click();
  }

  selectMainHandler(event) {
    this.setState({
      mainFiles: event.target.files[0]
    });
  }

  selectLibrariesHandler(event) {
    this.setState({
      libraryFiles: Array.prototype.slice.call(event.target.files)
    });
  }

  selectCondaHandler(event) {
    this.setState({
      condaFiles: event.target.files[0]
    });
  }

  uploadPackage(event) {
    if (_.isEmpty(this.refs.packageName.value)) {
      this.setState({
        missingPackageNameError: true
      });
      event.preventDefault();
      return false;
    }
    let uploadedPackage = {
      name: this.refs.packageName.value
    };
    this.props.upload(event, uploadedPackage, this.refs.uploadForm);
  }

  onPackageNameChanged = () => {
    if ((this.refs.packageName as any).value.length < 1) {
      this.setState({packageNamed: false});
    } else {
      this.setState({packageNamed: true});
    }
  };

  onMainTooltipOver = () => {
    this.setState({
      showTooltipMain: true
    });
  };

  onMainTooltipOut = () => {
    this.setState({
      showTooltipMain: false
    });
  };

  onLibraryTooltipOver = () => {
    this.setState({
      showTooltipLibrary: true
    });
  };

  onLibraryTooltipOut = () => {
    this.setState({
      showTooltipLibrary: false
    });
  };

  onCondaTooltipOver = () => {
    this.setState({
      showTooltipConda: true
    });
  };

  onCondaTooltipOut = () => {
    this.setState({
      showTooltipConda: false
    });
  };

  onNameTooltipOver = () => {
    this.setState({
      showTooltipName: true
    });
  };

  onNameTooltipOut = () => {
    this.setState({
      showTooltipName: false
    });
  };

  render(): React.ReactElement<DefaultModal> {
    let disableSubmit = false;
    if (this.state.libraryFiles.length < 1 || !this.state.mainFiles || !this.state.packageNamed) {
      disableSubmit = true;
    }

    return (
      <DefaultModal className="upload-preprocessing-modal" open={this.props.open}>
        <header className="page-header">
          UPLOAD PRE-PROCESSING PACKAGE (PYTHON)
        </header>
        <section>
          <form ref="uploadForm" onSubmit={this.uploadPackage.bind(this)}>
            <Table>
              <Row>
                <Cell>
                  SELECT PYTHON MAIN
                </Cell>
                <Cell>
                  <div>
                    Select a main Python file for pre-processing.&nbsp;
                    <i className="fa fa-question-circle-o orange tooltip-launcher" aria-hidden="true" onMouseEnter={this.onMainTooltipOver} onMouseLeave={this.onMainTooltipOut}>
                      {this.state.showTooltipMain ?
                        <div className="tooltip tooltip-question">
                          <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" viewBox="0 0 200.6 200.1" className="caret">
                            <polygon className="caret-triangle" points="100.3,3.4 200.6,200.1 0,200.1 "/>
                            <rect x="0" y="196" className="caret-cover" width="200" height="20"/>
                          </svg>
                          The output from this Python file should be one of row of an H2O data form that your model is expecting.
                        </div>
                        : null }
                    </i>
                  </div>
                  <div className="upload">
                    <div className="upload-info" onClick={this.selectMain.bind(this)}>
                      <span>
                        <i className="fa fa-folder-o"/>
                      </span>
                      <span className="file-list">{this.state.mainFiles ? this.state.mainFiles.name : 'N/A'}</span>
                      <span>
                        <i className="fa fa-close"/>
                      </span>
                      <input type="file" name="selectMain" onChange={this.selectMainHandler.bind(this)}/>
                    </div>
                  </div>
                </Cell>
              </Row>
              <Row>
                <Cell>
                  SELECT PYTHON LIBRARIES
                </Cell>
                <Cell>
                  <div>
                    Select a one or more Python files for your library.&nbsp;
                    <i className="fa fa-question-circle-o orange tooltip-launcher tooltip-launcher-libraries" aria-hidden="true" onMouseEnter={this.onLibraryTooltipOver} onMouseLeave={this.onLibraryTooltipOut}>
                      {this.state.showTooltipLibrary ?
                        <div className="tooltip tooltip-question tooltip-libraries">
                        <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" viewBox="0 0 200.6 200.1" className="caret">
                          <polygon className="caret-triangle" points="100.3,3.4 200.6,200.1 0,200.1 "/>
                          <rect x="0" y="196" className="caret-cover" width="200" height="20"/>
                        </svg>

                          Any non-standard libraries called here should be installed into your deployment environment prior to launching services
                      </div>
                        : null }
                    </i>
                  </div>
                  <div className="upload">
                    <div className="upload-info" onClick={this.selectLibraries.bind(this)}>
                      <span>
                        <i className="fa fa-folder-o"/>
                      </span>
                      <span className="file-list">{this.state.libraryFiles.length > 0 ? this.state.libraryFiles.map((file, i) => {
                        return <div key={i}>{file.name}</div>;
                      }) : 'N/A'}</span>
                      <span>
                        <i className="fa fa-close"/>
                      </span>
                      <input type="file" name="selectLibraries" onChange={this.selectLibrariesHandler.bind(this)} multiple/>
                    </div>
                  </div>
                </Cell>
              </Row>

              <Row>
                <Cell>
                  SELECT CONDA CONFIG
                </Cell>
                <Cell>
                  <div>
                    Pick a .yaml file that defines your conda environment.&nbsp;
                    <i className="fa fa-question-circle-o orange tooltip-launcher tooltip-launcher-conda" aria-hidden="true" onMouseEnter={this.onCondaTooltipOver} onMouseLeave={this.onCondaTooltipOut}>
                      {this.state.showTooltipConda ?
                      <div className="tooltip tooltip-question tooltip-conda">
                        <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" viewBox="0 0 200.6 200.1" className="caret">
                          <polygon className="caret-triangle" points="100.3,3.4 200.6,200.1 0,200.1 "/>
                          <rect x="0" y="196" className="caret-cover" width="200" height="20"/>
                        </svg>
                        you can get this file by doing this in your commandline of conda environment
                        <p>$ conda env export > mypackage.yaml</p>
                      </div>
                        : null }
                    </i>
                  </div>
                  <div className="upload">
                    <div className="upload-info" onClick={this.selectConda.bind(this)}>
                      <span>
                        <i className="fa fa-folder-o"/>
                      </span>
                      <span className="file-list">{this.state.condaFiles ? this.state.condaFiles.name : 'N/A'}</span>
                      <span>
                        <i className="fa fa-close"/>
                      </span>
                      <input type="file" name="selectConda" onChange={this.selectCondaHandler.bind(this)} />
                    </div>
                  </div>
                </Cell>
              </Row>

              <Row>
                <Cell>
                  NAME THE PACKAGE
                </Cell>
                <Cell>
                  <div>Pick a name for this pre-processing package.&nbsp;
                    <i className="fa fa-question-circle-o orange tooltip-launcher tooltip-launcher-name" aria-hidden="true" onMouseEnter={this.onNameTooltipOver} onMouseLeave={this.onNameTooltipOut}>
                      { this.state.showTooltipName ? <div className="tooltip tooltip-name tooltip-question">
                        <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" viewBox="0 0 200.6 200.1" className="caret">
                          <polygon className="caret-triangle" points="100.3,3.4 200.6,200.1 0,200.1 "/>
                          <rect x="0" y="196" className="caret-cover" width="200" height="20"/>
                        </svg>

                        You will use it as a reference when deploying models.
                      </div>
                        : null }
                    </i>
                  </div>
                  <div className="package-name-label muted">Package name</div>
                  <input ref="packageName" type="text" className={classNames('package-name', {error: this.state.missingPackageNameError})} onChange={this.onPackageNameChanged} />
                </Cell>
              </Row>
              <Row className="button-row">
                <Cell/>
                <Cell>
                  {disableSubmit ?
                      <button type="submit" className="button-primary disabled">
                        Upload
                      </button> :
                      <button type="submit" className="button-primary" onClick={this.uploadPackage.bind(this)}>
                        Upload
                      </button>}
                  <button className="button-secondary" onClick={this.props.cancel.bind(this)}>
                    Cancel
                  </button>
                </Cell>
              </Row>
            </Table>
          </form>
        </section>
      </DefaultModal>
    );
  }
}
