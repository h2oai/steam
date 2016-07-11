import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { rocChart } from 'vis-components';
import '../styles/rocgraph.scss';

interface Props {

}

export default class RocGraph extends React.Component<Props, any> {
  _mountNode: Element;

  componentWillMount() {
    let config = {
      'margin': 20,
      'width': 150,
      'height': 150,
      'interpolationMode': 'basis',
      'fpr': 'X',
      'tprVariables': [
        {
          'name': 'BPC',
          'label': 'Break Points'
        },
        {
          'name': 'WNR',
          'label': 'Winners'
        },
        {
          'name': 'FSP',
          'label': 'First Serve %',
        },
        {
          'name': 'NPW',
          'label': 'Net Points Won'
        }
      ],
      'animate': true,
      'smooth': true
    };
    rocChart.plot(this._mountNode, this.props.data, config);
  }

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
    ReactDOM.unstable_renderSubtreeIntoContainer(this, this.getGraph(), this._mountNode);
  }

  getGraph() {
    return (
      <div></div>
    );
  }

  render() {
    return null;
  }
}
