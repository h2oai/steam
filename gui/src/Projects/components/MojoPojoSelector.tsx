/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

import * as React from 'react';
import '../styles/mojopojo.scss';

export default class MojoPojoSelector extends React.Component<any, any> {

  constructor(props) {
    super(props);
    this.state = {
      showMenu: false
    };
  }

  onAutoSelected() {
    localStorage.setItem("mojoPojoSelection", "auto");
  }
  onPojoSelected() {
    localStorage.setItem("mojoPojoSelection", "pojo");
  }
  onMojoSelected() {
    localStorage.setItem("mojoPojoSelection", "mojo");
  }

  onMouseOver() {
    this.setState({showMenu: true});
  }
  onMouseLeave() {
    this.setState({showMenu: false});
  }

  render(): React.ReactElement<HTMLSpanElement> {
    let mojoPojoSelection = localStorage.getItem("mojoPojoSelection");
    if (mojoPojoSelection !== "auto" && mojoPojoSelection !== "pojo" && mojoPojoSelection !== "mojo") {
      mojoPojoSelection = "auto";
    }

    return (
      <span className="mojo-select-launcher tooltip-launcher" onMouseEnter={this.onMouseOver.bind(this)} onMouseLeave={this.onMouseLeave.bind(this)}>
        { this.state.showMenu ?
        <div className="mojo-select tooltip">
          <h2>Choose Format</h2>
          <div className="format-select">{mojoPojoSelection === "auto" ? <input type="radio" name="format" value="auto" onClick={this.onAutoSelected} defaultChecked /> : <input type="radio" name="format" onClick={this.onAutoSelected} value="auto"/> }
            Auto Select
          </div>
          <div className="format-select">{mojoPojoSelection === "pojo" ? <input type="radio" name="format" value="pojo" onClick={this.onPojoSelected} defaultChecked/> : <input type="radio" name="format" onClick={this.onPojoSelected} value="pojo"/> }
            Prefer Pojo
          </div>
          <div className="format-select">{mojoPojoSelection === "mojo" ? <input type="radio" name="format" value="mojo" onClick={this.onMojoSelected} defaultChecked /> : <input type="radio" name="format" onClick={this.onMojoSelected} value="mojo"/> }
            Prefer Mojo
          </div>
        </div> : null}
        <span className="link">here</span>
      </span>
    );
  }
}
