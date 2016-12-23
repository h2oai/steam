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
import { Position } from '@blueprintjs/core';
import { Popover, PopoverInteractionKind } from '@blueprintjs/core/dist/components/popover/popover';

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

  render(): React.ReactElement<HTMLSpanElement> {
    let mojoPojoSelection = localStorage.getItem("mojoPojoSelection");
    if (mojoPojoSelection !== "auto" && mojoPojoSelection !== "pojo" && mojoPojoSelection !== "mojo") {
      mojoPojoSelection = "auto";
    }

    let autoInput;
    let preferPojoInput;
    let preferMojoInput;
    if (mojoPojoSelection === "auto") {
      autoInput = <input type="radio" name="format" value="auto" onClick={this.onAutoSelected} defaultChecked />;
    } else {
      autoInput = <input type="radio" name="format" value="auto" onClick={this.onAutoSelected} />;
    }
    if (mojoPojoSelection === "pojo") {
      preferPojoInput = <input type="radio" name="format" value="pojo" onClick={this.onPojoSelected} defaultChecked/>;
    } else {
      preferPojoInput = <input type="radio" name="format" value="pojo" onClick={this.onPojoSelected} />;
    }
    if (mojoPojoSelection === "mojo") {
      preferMojoInput = <input type="radio" name="format" value="mojo" onClick={this.onMojoSelected} defaultChecked />;
    } else {
      preferMojoInput = <input type="radio" name="format" value="mojo" onClick={this.onMojoSelected} />;
    }

    let popoverContent =
      <div>
        <h2>Choose Format</h2>
        <div>
          {autoInput} Auto Select
        </div>
        <div>
          { preferPojoInput} Prefer Pojo
        </div>
        <div>
          { preferMojoInput } Prefer Mojo
        </div>
      </div>;

    return (
    <Popover content={popoverContent}
             inline={true}
             interactionKind={PopoverInteractionKind.HOVER}
             popoverClassName="pt-popover-content-sizing"
             position={Position.BOTTOM}
             useSmartPositioning={false}>
      <span className="link">here</span>
    </Popover>
    );
  }
}
