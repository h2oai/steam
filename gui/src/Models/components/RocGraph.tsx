import * as React from 'react';
import '../styles/rocgraph.scss';
import { rocChart } from 'vis-components';

interface Props {
  data: any[]
}

export default class RocGraph extends React.Component<Props, any> {

  _mountNode: Element;

  componentDidMount() {
    this._mountNode = ReactDOM.findDOMNode(this);
    this.renderGraph();
  }

  componentDidUpdate() {
    if (this._mountNode) {
      this.renderGraph();
    }
  }

  renderGraph() {
    let cfg = {
        margin: { top: 2, right: 2, bottom: 2, left: 2 },
        width: 60,
        height: 60,
        interpolationMode: 'basis',
        ticks: undefined,
        tickValues: [0, 0.1, 0.25, 0.5, 0.75, 0.9, 1],
        fpr: 'fpr',
        tprVariables: [{
          name: 'tpr',
        }],
        animate: false,
        hideTicks: true,
        hideAxes: true,
        hideBoundaries: false
    };

    rocChart.plot(this._mountNode, this.props.data, cfg);
  }

  render() {
    return <div className="roc-container"></div>;
  }
}
