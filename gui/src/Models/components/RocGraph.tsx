import * as React from 'react';
import * as ReactDOM from 'react-dom';
import * as d3 from 'd3';
import '../styles/rocgraph.scss';

interface Props {

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
    let xScale = d3.scaleLinear().range([0, 100]);
    let yScale = d3.scaleLinear().range([100, 0]);
    xScale.domain([0, 1]);
    yScale.domain([0, 1]);

    let data = [
      {
        x: 0
      },
      {
        x: .3
      },
      {
        x: 0.5
      },
      {
        x: 1
      }
    ];
    let line = d3.line().x(d => xScale(d.x)).y(d => yScale(d.x));
    d3.select(this._mountNode).append('g').append('path').style('stroke', 'blue').attr('d', line(data));
  }

  render() {
    return <svg width="60px" height="60px" className="train-roc-graph" viewBox="0 0 100 100" preserveAspectRatio="none"></svg>;
  }
}
