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
    rocChart.plot(this._mountNode, this.props.data);
  }

  render() {
    return <div className="roc-graph"></div>;
  }
}
