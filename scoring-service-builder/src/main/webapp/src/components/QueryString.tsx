import * as React from 'react';

interface Props {

}

export default class QueryString extends React.Component<Props, any> {
  buildQueryString() {
    
  }
  
  render() {
    return (
      <fieldset>
        <legend>Query String</legend>
        <div className="form-group">
          The parameters above gets automatically built into a REST API query string. You can also input your own string if that's easier for you.
        </div>
        <div className="form-group">
          <label className="sr-only"></label>
          <div className="input-group">
            <div id="url-prefix" className="input-group-addon">
              http://localhost:55001/predict?
            </div>
            <input type="text" className="form-control" id="queryParams" name="p"></input>
            <div className="input-group-addon">
              <a id="query-link" href="" target="_blank" style={{cursor: 'not-allowed'}}>
                <i className="glyphicon glyphicon-new-window"></i>
              </a>
            </div>
          </div>
        </div>
      </fieldset>
    );
  }
}