/**
 * Created by justin on 7/28/16.
 */
import * as React from 'react';
import { connect } from 'react-redux';
import * as ReactDOM from 'react-dom';
import * as classNames from 'classnames';
import FilterIcon from './FilterIcon';
import { bindActionCreators } from 'redux';
import { fetchSortCriteria } from '../actions/leaderboard.actions';
import '../styles/filterdropdown.scss';

interface Props {
  target?: Element,
  onFilter: Function,
  criteria: string[]
}

interface DispatchProps {
  fetchSortCriteria: Function
}

const FILTER_MAP = {
  mean_residual_deviance: 'MRD',
  r_squared: <span>R<sup>2</sup></span>,
  mse: 'MSE'
};

export class FilterDropdown extends React.Component<Props & DispatchProps, any> {
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

  componentWillReceiveProps(nextProps) {
    if (nextProps.category && !this.props.criteria) {
      this.props.fetchSortCriteria(nextProps.category.toLowerCase());
    }
  }

  bodyClickHandler(event) {
    if (!ReactDOM.findDOMNode(this.refs.filterDropdown).contains(event.target) && !ReactDOM.findDOMNode(this.refs.filterDropdownInvoker).contains(event.target)) {
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
    if (!this.props.criteria) {
      return <div></div>;
    }
    return (
      <div className="filter-dropdown">
        <button ref="filterDropdownInvoker" className={classNames('filter-dropdown-invoker', {open: this.state.open})}
                onClick={this.openDropdown.bind(this)}><FilterIcon/></button>
        <div ref="filterDropdown" className={classNames('filter-dropdown-menu', {open: this.state.open})}>
          <div className="filter-option">
            <div className="filter-labels">SORT BY</div>
            <ul>
              {this.props.criteria.map((criteria, i) => {
                return <li key={i} onClick={this.selectSort.bind(this, criteria)}
                    className={classNames({selected:this.state.sortBy === criteria})}>{FILTER_MAP[criteria]} {this.state.sortBy === criteria ?
                  <i className="fa fa-check"/> : null}</li>;
              })}
            </ul>
          </div>
          <div className="filter-option">
            <div className="filter-labels">ORDER</div>
            <ul>
              <li onClick={this.selectOrder.bind(this, 'asc')}
                  className={classNames({selected:this.state.orderBy === 'asc'})}>ASC {this.state.orderBy === 'asc' ?
                <i className="fa fa-check"/> : null}</li>
              <li onClick={this.selectOrder.bind(this, 'des')}
                  className={classNames({selected:this.state.orderBy === 'des'})}>DES {this.state.orderBy === 'des' ?
                <i className="fa fa-check"/> : null}</li>
            </ul>
          </div>
        </div>
      </div>
    );
  }

}

function mapStateToProps(state) {
  return {
    criteria: state.leaderboard.criteria
  };
}

function mapDispatchToProps(dispatch) {
  return {
    fetchSortCriteria: bindActionCreators(fetchSortCriteria, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(FilterDropdown);
