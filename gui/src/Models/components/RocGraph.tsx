import * as React from 'react';
import * as ReactDOM from 'react-dom';
import * as visComponents from 'vis-components';
import '../styles/rocgraph.scss';

interface Props {

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
    visComponents.rocChart.plot(this._mountNode, this.props.data);
  }

  render() {
    return <div></div>;
  }
}
