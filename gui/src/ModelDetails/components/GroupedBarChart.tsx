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
const visComponents: any = require('vis-components');
const groupedBarChart = visComponents.groupedBarChart;

type IBarColors = [string, string, string, string, string, string, string];

interface Props {
    data: any[]
    groupByVariable?: string
    barColors?: IBarColors
}

export default class GroupedBarChart extends React.Component<Props, any> {

    _mountNode: Element;

    static defaultProps: Props = {
      data: [],
      groupByVariable: 'value',
      barColors: ['#a6cee3', '#1f78b4', '#b2df8a', '#33a02c', '#fb9a99', '#e31a1c', '#fdbf6f']
    };

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

    renderGraph() {
        let options = {
            groupByVariable: this.props.groupByVariable,
            barColors: this.props.barColors
        };

        groupedBarChart.plot(this._mountNode, this.props.data, options);
    }

    render() {
        return <div className="grouped-bar-container"></div>;
    }
}
