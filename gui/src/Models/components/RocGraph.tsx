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
import * as ReactDOM from 'react-dom';
const rocChart: any = require('roc-chart');

import '../styles/rocgraph.scss';
import { BRAND_ORANGE, BRAND_BLUE } from '../../App/utils/colors';

interface Props {
  data: any[],
  config?: any
}

export default class RocGraph extends React.Component<Props, any> {

  _mountNode: Element;

  componentDidMount() {
    this._mountNode = ReactDOM.findDOMNode(this);
    this.renderGraph();
  }

  componentWillUnmount() {
    if (this._mountNode) {
      ReactDOM.unmountComponentAtNode(this._mountNode);
      this._mountNode.remove();
      this._mountNode = null;
    }
  }

  componentWillUpdate(nextProps) {
    let cfg = {
      margin: {top: 2, right: 2, bottom: 2, left: 2},
      width: '100%',
      height: '100%',
      interpolationMode: 'basis',
      smooth: true,
      animate: undefined,
      hideAxes: true,
      hideAUCText: true,
      curveColors: [BRAND_BLUE, BRAND_ORANGE]
    };
    this._mountNode.innerHTML = '';
    rocChart.plot(this._mountNode, nextProps.data, this.props.config || cfg);
  }

  renderGraph() {
    let cfg = {
      margin: {top: 2, right: 2, bottom: 2, left: 2},
      width: '100%',
      height: '100%',
      interpolationMode: 'basis',
      smooth: true,
      animate: false,
      hideAxes: true,
      hideAUCText: true,
      curveColors: [BRAND_BLUE]
    };
    rocChart.plot(this._mountNode, this.props.data, this.props.config || cfg);
  }

  render() {
    return <div className="roc-container"></div>;
  }
}
