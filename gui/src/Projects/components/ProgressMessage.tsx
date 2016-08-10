import * as React from 'react';
import '../styles/progressmessage.scss';
import CircularProgress from 'material-ui/CircularProgress';
import {getMuiTheme, MuiThemeProvider} from 'material-ui/styles';

interface Props {
  showSpinner: boolean
  message: String
}

export default class ProgressMessage extends React.Component<Props, any> {
  constructor() {
    super();
  }

  render(): React.ReactElement<HTMLElement> {
    return (
      <div className="cluster-fetch-progress">
        { this.props.showSpinner ?
          <MuiThemeProvider muiTheme={getMuiTheme()}>
            <div className="spin-spacer">
              <CircularProgress style={{"position":"absolute","top":"-15px","left":"-15px"}} color="#f5a623" size={0.24} />
            </div>
          </MuiThemeProvider>
          : null }
        <span>{this.props.message}</span>
      </div>
    );
  }
}
