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
      missingPackageNameError: false
    };
  }

  selectMain() {
    $('input[name="selectMain"]').click();
  }

  selectLibraries() {
    $('input[name="selectLibraries"]').click();
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

  render(): React.ReactElement<DefaultModal> {
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
                  <div>Select a main Python file for pre-processing.</div>
                  <span className="muted">The output from this Python file should be one row of an H2O data from that your model is expecting.</span>
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
                  <div>Select a one or more Python files for your library.</div>
                  <span className="muted">Any non-standard libraries called here should be installed into your deployment environment prior to launching services.</span>
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
                  NAME THE PACKAGE
                </Cell>
                <Cell>
                  <div>Pick a name for this pre-processing package. You will use it as a reference when deploying models.</div>
                  <div className="package-name-label muted">Package name</div>
                  <input ref="packageName" type="text" className={classNames('package-name', {error: this.state.missingPackageNameError})}/>
                </Cell>
              </Row>
              <Row className="button-row">
                <Cell/>
                <Cell>
                  <button type="submit" className="button-primary" onClick={this.uploadPackage.bind(this)}>
                    Upload
                  </button>
                  <button className="button-primary invert" onClick={this.props.cancel.bind(this)}>
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
