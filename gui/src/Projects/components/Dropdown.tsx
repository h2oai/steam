/**
 * Created by justin on 7/11/16.
 */
import * as React from 'react';
import * as classNames from 'classnames';
import * as $ from 'jquery';
import '../styles/dropdown.scss';

export default class Dropdown extends React.Component<any, any> {
  constructor() {
    super();
    this.state = {
      open: false
    };
  }

  componentWillMount() {
    $(document).click((event) => {
      console.log(event.target === this.refs.dropdownInvoker);
      if (event.target === this.refs.dropdownInvoker) {
        this.setState({
          open: !this.state.open
        });
      } else {
        this.setState({
          open: false
        });
      }
    });
  }

  render() {
    return (
      <div className={classNames('dropdown', { open: this.state.open})}>
        <button ref="dropdownInvoker" type="button">Open Dropdown</button>
        <div className="dropdown-menu">
          <ul>
            <li>Option 1</li>
            <li>Option 2</li>
          </ul>
        </div>
      </div>
    );
  }
}
