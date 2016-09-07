import * as React from 'react';
import '../styles/progress.scss';

interface Props {
  message: string
}

export default class Progress extends React.Component<Props, any> {
  constructor() {
    super();
  }

  render(): React.ReactElement<HTMLElement> {
    return (
      <div className="progress">
        <div className='uil-facebook-css'>
          <div></div>
          <div></div>
          <div></div>
        </div> {this.props.message}
      </div>
    );
  }
}
