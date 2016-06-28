/**
 * Created by justin on 6/27/16.
 */
import * as React from 'react';
import '../styles/modeloverview.scss';

export default class ModelOverview extends React.Component<any, any> {
  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="model-overview">
        <div>
          <div className="details">
            <div className="details--label">
              Author
            </div>
            <div className="details--line">
            </div>
            <div className="details--value">
              Mark Landry
            </div>
          </div>
          <div className="details">
            <div className="details--label">
              Date
            </div>
            <div className="details--line">
            </div>
            <div className="details--value">
              2016-06-06 17:17
            </div>
          </div>
        </div>
      </div>
    );
  }
}