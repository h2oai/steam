import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { groupedBarChart } from 'vis-components';

interface Props {
    data: any[]
}

export default class GroupedBarChart extends React.Component<Props, any> {

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

    renderGraph() {
        let options = {
            groupByVariable: 'value',
            barColors: ['#a6cee3', '#1f78b4', '#b2df8a', '#33a02c', '#fb9a99', '#e31a1c', '#fdbf6f']
        };

        groupedBarChart.plot(this._mountNode, this.props.data, options);
    }

    render() {
        return <div className="grouped-bar-container"></div>;
    }
}
