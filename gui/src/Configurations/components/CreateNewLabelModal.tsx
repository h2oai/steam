/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

/**
 * Created by justin on 7/27/16.
 */
import * as React from 'react';
import * as $ from 'jquery';
import DefaultModal from '../../App/components/DefaultModal';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import '../styles/createnewlabelmodal.scss';


interface Props {
    open: boolean,
    save: Function,
    cancel: Function,
    label?: any
}

interface State {
  id: number|boolean
  name: string
  description: string
}

const initialState: State = {
    id: false,
    name: '',
    description: ''
};


export default class CreateNewLabelModal extends React.Component<Props, any> {

    constructor() {
        super();
        this.state = initialState;
    }

    componentWillReceiveProps(nextProps) {
      if (nextProps.label.id) {
        this.setState({
          id: nextProps.label.id,
          name: nextProps.label.name,
          description: nextProps.label.description
        });
      }
    }

    updateState(event) {
      let newState = {};
      newState[event.currentTarget.name] = event.currentTarget.value;
      this.setState(newState);
    }

    cancel() {
      this.setState(initialState);
      this.props.cancel();
    }

    save() {
      this.props.save(this.state);
      this.setState(initialState);
    }

    render(): React.ReactElement<DefaultModal> {
        return (
            <DefaultModal open={this.props.open} closeHandler={this.props.cancel}>
                <div className="create-edit-label-modal">
                    <header className="page-header">
                        Create / Edit Label
                    </header>
                    <section>
                        <Table>
                            <Row>
                                <Cell className="table-row-name">
                                    Label Info
                                </Cell>
                                <Cell className="table-row-item">
                                    <p>Enter a name and description of your label.</p>
                                    <p className="muted">You can use this label in the project for exactly 1 model.</p>
                                    <div className="form-group">
                                      <div className="form-item">
                                          <label className="muted" htmlFor="name">Label name</label>
                                          <input type="text" name="name" value={this.state.name} onChange={this.updateState.bind(this)} />
                                      </div>
                                      <div className="form-item">
                                          <label className="muted" htmlFor="description">Label description</label>
                                          <textarea name="description" value={this.state.description} rows="4" cols="50" onChange={this.updateState.bind(this)}></textarea>
                                      </div>
                                    </div>
                                </Cell>
                            </Row>
                            <Row className="button-row">
                                <Cell className="table-row-name"></Cell>
                                <Cell className="table-row-item">
                                    <button className="button-primary" onClick={this.save.bind(this) }>
                                        Save
                                    </button>
                                    <button className="button-secondary" onClick={this.cancel.bind(this) }>
                                        Cancel
                                    </button>
                                </Cell>
                            </Row>
                        </Table>
                    </section>
                </div>
            </DefaultModal>
        );
    }
}
