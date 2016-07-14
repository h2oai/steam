/**
 * Created by justin on 7/5/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import * as $ from 'jquery';
import { Link } from 'react-router';
import PageHeader from '../components/PageHeader';
import Table from '../components/Table';
import Row from '../components/Row';
import Cell from '../components/Cell';
import ProgressBar from '../components/ProgressBar';
import '../styles/newproject.scss';

interface Props {
  children?: React.ReactChildren
}

export default class NewProject extends React.Component<Props, any> {
  refs: {
    [key: string]: (Element);
    uploadFile: HTMLInputElement;
  };

  constructor() {
    super();
    this.state = {
      shapeIsDisabled: true,
      uploadFilename: '',
      isUploading: false
    };
  }

  componentDidMount(): void {
    let uploadFile = ($(this.refs.uploadFile) as any);
    uploadFile.bind('change.file', () => {
      this.setState({
        uploadFilename: uploadFile[0].files[0].name
      });
    });
  }

  componentWillUnmount(): void {
    $(this.refs.uploadFile).unbind('change.file');
  }

  onClickFileText(): void {
    $(this.refs.uploadFile).click();
  }

  upload(): void {
    /**
     * TODO(justinloyola): Add upload functionality to backend
     */
    this.setState({
      isUploading: true,
      stopUploading: false
    });
    setTimeout(() => {
      this.setState({
        stopUploading: true
      });
    }, 2000)
  }

  getAddDataSource() {
    return (
      <form>
        <label>Add a Data Source</label>
        <span>We will automatically infer a data dictionary based on the data</span>
        <input ref="uploadText" type="text" className="upload-file" onClick={this.onClickFileText.bind(this)}
               value={this.state.uploadFilename}/>
        <input ref="uploadFile" type="file"/>
        <button type="button" className="default" onClick={this.upload.bind(this)}>Upload</button>
      </form>
    );
  }

  onCompleteUploading() {
    console.log('show');
    this.setState({
      isUploading: false
    });
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="new-project">
        <PageHeader>New Project</PageHeader>
        <form>
          <label>Give your project a name</label>
          <input type="text" placeholder="Name"/>
        </form>
        {this.state.isUploading === false ? this.getAddDataSource() : null}
        {this.state.isUploading === true ? <ProgressBar showPercentage={true} onComplete={this.onCompleteUploading.bind(this)} end={this.state.stopUploading} style={{display: this.state.stopUploading ? 'block' : 'none'}}/>: null}
        <form className={classNames({disabled: !this.state.uploadFilename})}>
          <label>Verify Data Shape</label>
          <span>Your models will be more accurate if H2O has an accurate understanding of the column types in your data.</span>
        </form>
        <div>
          <Table className={classNames({disabled: !this.state.uploadFilename})}>
            <Row header={true}>
              <Cell>COLUMN NAME</Cell>
              <Cell>DATA TYPE</Cell>
              <Cell>DISTRIBUTION</Cell>
              <Cell>STATISTICS</Cell>
            </Row>
            <Row>
              <Cell><span>setosa_length</span></Cell>
              <Cell><span>numeric</span></Cell>
              <Cell/>
              <Cell>
                <div className="statistics">
                  <div className="headings">
                    <div><span>min</span></div>
                    <div><span>max</span></div>
                    <div><span>mean</span></div>
                  </div>
                  <div className="values">
                    <div><span>0.4</span></div>
                    <div><span>5.1</span></div>
                    <div><span>3.4</span></div>
                  </div>
                </div>
              </Cell>
            </Row>
            <Row>
              <Cell>
                <span>class</span>
              </Cell>
              <Cell>
                <span>categorical</span>
              </Cell>
              <Cell/>
              <Cell>
                <div className="statistics">
                  <div className="headings">
                    <div><span>classes</span></div>
                  </div>
                  <div className="values">
                    <div><span>3</span></div>
                  </div>
                </div>
              </Cell>
            </Row>
          </Table>
        </div>
        <div>
          <form className={classNames({disabled: !this.state.uploadFilename})}>
            <label>Select Response Column</label>
            <span>Identify the column with the value you want to predict.</span>
            <select>
              <option>Test</option>
            </select>
          </form>
        </div>
        <Link to="/projects/new/2" className={classNames('default next', {disabled: !this.state.uploadFilename})}>Next: Train Initial Models</Link>
      </div>
    );
  }
}
