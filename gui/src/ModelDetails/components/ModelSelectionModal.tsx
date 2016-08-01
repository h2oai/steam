/**
 * Created by justin on 7/29/16.
 */

import * as React from 'react';
import * as moment from 'moment';
import FilterDropdown from '../../Models/components/FilterDropdown';
import DefaultModal from '../../App/components/DefaultModal';
import PageHeader from '../../Projects/components/PageHeader';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchLeaderboard } from '../../Models/actions/leaderboard.actions';
import { Model } from '../../Proxy/Proxy';
import '../styles/modelselectionmodal.scss';

interface Props {
  open: boolean,
  models: Model[]
  projectId: string,
  onSelectModel: Function,
  onCancel: Function

}

interface DispatchProps {
  fetchLeaderboard: Function
}

export class ModelSelectionModal extends React.Component<Props & DispatchProps, any> {
  componentWillMount() {
    this.props.fetchLeaderboard(parseInt(this.props.projectId, 10));
  }

  onFilter(filters) {
    /**
     * TODO(justinloyola): AJAX call to filter models
     */
  }

  render(): React.ReactElement<DefaultModal> {
    return (
      <DefaultModal className="model-comparison-modal" open={this.props.open}>
        <PageHeader>
          CHOOSE MODEL TO COMPARE
        </PageHeader>
        <div>
          <div>Filter models by name</div>
          <input type="text" placeholder="filter models"/>
          <Table>
            <Row header={true}>
              <Cell>
                <FilterDropdown onFilter={this.onFilter.bind(this)}/>
              </Cell>
              <Cell>
                MODEL
              </Cell>
              <Cell>
                DATE
              </Cell>
              <Cell>
                MSE
              </Cell>
              <Cell>
                AUC
              </Cell>
              <Cell>
                F1
              </Cell>
              <Cell/>
            </Row>
            {this.props.models.map((model, i) => {
              return (
                <Row key={i}>
                  <Cell/>
                  <Cell>
                    {model.name}
                  </Cell>
                  <Cell>
                    {moment.unix(model.created_at).format('YYYY-MM-DD HH:mm')}
                  </Cell>
                  <Cell>
                    {model.mse}
                  </Cell>
                  <Cell>
                    {model.r_squared}
                  </Cell>
                  <Cell>
                    {model.r_squared}
                  </Cell>
                  <Cell>
                    <button className="default" onClick={this.props.onSelectModel.bind(this, model)}>Select</button>
                  </Cell>
                </Row>
              );
            })}
          </Table>
        </div>
        <footer>
          <button className="default" onClick={this.props.onCancel.bind(this)}>Cancel</button>
        </footer>
      </DefaultModal>
    );
  }
}

function mapStateToProps(state: any): any {
  return {
    models: state.leaderboard.items
  };
}

function mapDispatchToProps(dispatch): DispatchProps {
  return {
    fetchLeaderboard: bindActionCreators(fetchLeaderboard, dispatch)
  };
}

export default connect<any, DispatchProps, any>(mapStateToProps, mapDispatchToProps)(ModelSelectionModal);
