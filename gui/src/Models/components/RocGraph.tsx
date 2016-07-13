import * as React from 'react';
import '../styles/rocgraph.scss';

interface Props {

}

export default class RocGraph extends React.Component<Props, any> {
  render() {
    return <svg width="60px" height="60px" className="train-roc-graph" viewBox="0 0 100 100" preserveAspectRatio="none"></svg>;
  }
}
