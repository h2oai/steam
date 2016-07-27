import * as React from 'react';
import * as ReactDOM from 'react-dom';
import '../styles/rocgraph.scss';
const visComponents: any = require('vis-components');
const rocChart = visComponents.rocChart;

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

  renderGraph() {
    let cfg = {
        margin: { top: 2, right: 2, bottom: 2, left: 2 },
        width: 60,
        height: 60,
        interpolationMode: 'basis',
        smooth: true,
        fpr: 'fpr',
        tprVariables: [{
          name: 'tpr',
          label: 'tpr'
        }],
        animate: false,
        hideAxes: true,
        hideAUCText: true
    };

    rocChart.plot(this._mountNode, this.props.data, this.props.config || cfg);
  }

  render() {
    return <div className="roc-container"></div>;
  }
}
