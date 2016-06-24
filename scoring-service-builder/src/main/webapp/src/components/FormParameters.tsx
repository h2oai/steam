import * as React from 'react';
import {bindActionCreators} from 'redux';
import {connect} from 'react-redux';
import * as _ from 'lodash';
import * as classNames from 'classnames';
import FormRow from './FormRow';
import Label from './Label';
import Select from './Select';
import {updateQueryString} from '../reducers/modelReducer';

interface DispatchProps {
  updateQueryString:Function,
}

interface Props {
  model: any,
  rows: any
}

export class FormParameters extends React.Component<Props & DispatchProps, any> {
  constructor() {
    super();
    this.updateQueryStringHandler = this.updateQueryStringHandler.bind(this);
  }

  getRow(row:string, i:number) {
    return (
      <div>
        <Label key={row} className={classNames('col-sm-6')}>{
          Number(i + 1) + '. ' + row}</Label>
        {this.inputCreationStrategy(row)}
      </div>
    );
  }

  inputCreationStrategy(rowName:string) {
    let row = this.props.model.domainMap[this.props.model.modelColumnNameToIndexMap[rowName]];
    if (_.isUndefined(row)) {
      return FormParameters.createInput(rowName);
    } else {
      return this.createSelect(rowName, row);
    }
  }

  private static createInput(rowName:string) {
    return (
      <div className="col-sm-6">
        <input type="text" className="form-control" name={rowName}/>
      </div>
    );
  }

  private updateQueryStringHandler(event: React.SyntheticEvent) {
    this.props.updateQueryString((event.target as HTMLSelectElement).value);
  }

  private createSelect(rowName:string, data:Array<any>) {
    return (
      <div className="col-sm-6">
        <select className="form-control" name={name} onChange={this.updateQueryStringHandler}>
          <option/>
          {(() => {
            return _.keys(data).sort((a, b) => {
              return Number(a.match(/\d+/)) - Number(b.match(/\d+/));
            }).map((key) => {
              return <option key={key} value={key}>{key}</option>
            })
          })()}
        </select>
      </div>
    );
  }

  render() {
    return (
      <fieldset id="fs-params" className={'form-group'}>
        <legend>Parameters</legend>
        {this.props.rows.map((row, i) => {
          return (
            <FormRow key={i}>
              {this.getRow(row, i)}
            </FormRow>
          );
        })}
      </fieldset>
    );
  }
}

function mapStateToProps(state) {
  return {};
}

function mapDispatchToProps(dispatch: Redux.Dispatch): DispatchProps {
  return {
    updateQueryString: bindActionCreators(updateQueryString, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(FormParameters)