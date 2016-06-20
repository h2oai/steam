/**
 * Created by justin on 6/17/16.
 */

import * as React from 'react';
import Form from './Form';
import FormParameters from './FormParameters';
import QueryString from './QueryString';

interface Props {
  model:any
}

export default class ModelForm extends React.Component<Props, any> {
  render() {
    return (
      <Form>
        <FormParameters rows={this.props.model.m._names} model={this.props.model}/>
        <QueryString></QueryString>
        <div class="btnContainer">
          <button id="predict-btn" type="submit" className="btn btn-primary" name="okbutton">PREDICT</button>
          <button id="reset-btn" type="reset" className="btn btn-default">CLEAR</button>
        </div>
      </Form>
    );
  }
}