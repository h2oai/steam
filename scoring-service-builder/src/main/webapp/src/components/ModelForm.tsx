/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import * as classNames from 'classnames';
import Form from './Form';
import FormRow from './FormRow';
import Label from './Label';
import Select from './Select';
import Input from './Input';

interface Props {
  className?:string,
  children?:any,
  id?:string,
  model:any
}

export default class ModelForm extends React.Component<Props, any> {
  rowMap:Array<any> = [];

  componentWillMount() {
    for (var i in this.props.model.modelColumnNameToIndexMap) {
      this.rowMap.push({
        label: i,
        data: this.props.model.m._domains[this.props.model.modelColumnNameToIndexMap[i]]
      });
    }

    this.setState({
      rows: this.props.model.m._names
    });
  }

  inputStrategy(data:Array<any>) {
    if (data === null) {
      return (
        <div className="col-sm-6">
          <Input type="text" className="form-control"/>
        </div>
      );
    } else {
      return (
        <div className="col-sm-6">
          <Select>
            <option></option>
          </Select>
        </div>
      );
    }
  }

  render() {
    if (!this.state && !this.state.rows) {
      return <div></div>;
    }

    // <Select>
    //   <option/>
    //   {row.data ? row.data.map((data, i) => {
    //     return <option key={i} value={data}>{data}</option>
    //   }) : null}
    // </Select>
    return (
      <Form>
        <fieldset id="fs-params" className={'form-group'}>
          <legend>Parameters</legend>
          {this.rowMap.map((row, i) => {
            return (
              <FormRow key={i}>
                <Label key={i} className={classNames('col-sm-6')}>{row.label}</Label>
                {this.inputStrategy(row.data)}
              </FormRow>
            );
          })}
        </fieldset>
        <fieldset>
          <legend>Query String</legend>
          <div className="form-group">
            The parameters above gets automatically built into a REST API query string. You can also input your own string if that's easier for you.
          </div>
          <div className="form-group">
            <label className="sr-only"></label>
            <div className="input-group">
              <div id="url-prefix" className="input-group-addon">http://localhost:55001/predict?</div>
                <input type="text" className="form-control" id="queryParams" name="p"/>
                <div className="input-group-addon">
                  <a id="query-link" href="" target="_blank" style={{cursor: 'not-allowed'}}><i
                    className="glyphicon glyphicon-new-window"></i></a>
                </div>
              </div>
            </div>
        </fieldset>
        <div class="btnContainer">
          <button id="predict-btn" type="submit" className="btn btn-primary" name="okbutton">PREDICT</button>
          <button id="reset-btn" type="reset" className="btn btn-default">CLEAR</button>
        </div>
      </Form>
    );
  }
}