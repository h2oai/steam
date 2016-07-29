/**
 * Created by justin on 7/28/16.
 */
import * as React from 'react';
import * as ReactDOM from 'react-dom';
import * as classNames from 'classnames';
import '../styles/filterdropdown.scss';

interface Props {
  target?: Element,
  onFilter: Function
}

export default class FilterDropdown extends React.Component<Props, any> {
  refs: {
    [key: string]: Element
    filterDropdown: Element
    filterDropdownInvoker: Element
  };

  constructor() {
    super();
    this.state = {
      open: false,
      sortBy: null,
      orderBy: null
    };
    this.bodyClickHandler = this.bodyClickHandler.bind(this);
  }

  componentWillMount() {
    document.body.addEventListener('click', this.bodyClickHandler);
  }

  componentWillUnmount() {
    document.body.removeEventListener('click', this.bodyClickHandler);
  }

  bodyClickHandler(event) {
    if (!ReactDOM.findDOMNode(this.refs.filterDropdown).contains(event.target) &&
      !ReactDOM.findDOMNode(this.refs.filterDropdownInvoker).contains(event.target)) {
      this.setState({
        open: false
      });
    }
  }

  openDropdown() {
    this.setState({
      open: !this.state.open
    });
  }

  selectSort(selection: string) {
    this.setState({
      sortBy: selection
    });
    this.props.onFilter({
      sortBy: selection,
      orderBy: this.state.orderBy
    });
  }

  selectOrder(selection: string) {
    this.setState({
      orderBy: selection
    });
    this.props.onFilter({
      sortBy: this.state.sortBy,
      orderBy: selection
    });
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div className="filter-dropdown">
        <button ref="filterDropdownInvoker" className={classNames('filter-dropdown-invoker', {open: this.state.open})} onClick={this.openDropdown.bind(this)}><i className="fa fa-sort-amount-desc"/></button>
        <div ref="filterDropdown" className={classNames('filter-dropdown-menu', {open: this.state.open})}>
          <div className="filter-option">
            <div className="filter-labels">SORT BY</div>
            <ul>
              <li onClick={this.selectSort.bind(this, 'date')} className={classNames({selected:this.state.sortBy === 'date'})}>Date {this.state.sortBy === 'date' ? <i className="fa fa-check"/> : null}</li>
              <li onClick={this.selectSort.bind(this, 'mse')} className={classNames({selected:this.state.sortBy === 'mse'})}>MSE {this.state.sortBy === 'mse' ? <i className="fa fa-check"/> : null}</li>
              <li onClick={this.selectSort.bind(this, 'auc')} className={classNames({selected:this.state.sortBy === 'auc'})}>AUC {this.state.sortBy === 'auc' ? <i className="fa fa-check"/> : null}</li>
              <li onClick={this.selectSort.bind(this, 'f1')} className={classNames({selected:this.state.sortBy === 'f1'})}>F1 {this.state.sortBy === 'f1' ? <i className="fa fa-check"/> : null}</li>
              <li onClick={this.selectSort.bind(this, 'logloss')} className={classNames({selected:this.state.sortBy === 'logloss'})}>Logloss {this.state.sortBy === 'logloss' ? <i className="fa fa-check"/> : null}</li>
            </ul>
          </div>
          <div className="filter-option">
            <div className="filter-labels">ORDER</div>
            <ul>
              <li onClick={this.selectOrder.bind(this, 'asc')} className={classNames({selected:this.state.orderBy === 'asc'})}>ASC {this.state.orderBy === 'asc' ? <i className="fa fa-check"/> : null}</li>
              <li onClick={this.selectOrder.bind(this, 'des')} className={classNames({selected:this.state.orderBy === 'des'})}>DES {this.state.orderBy === 'des' ? <i className="fa fa-check"/> : null}</li>
            </ul>
          </div>
        </div>
      </div>
    );
  }
}
