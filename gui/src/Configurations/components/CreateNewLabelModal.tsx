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
                    <header>
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
                                    <button className="default" onClick={this.save.bind(this) }>
                                        Save
                                    </button>
                                    <button className="default invert" onClick={this.cancel.bind(this) }>
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
