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
    cancel: Function
}

export default class CreateNewLabelModal extends React.Component<Props, any> {
    refs: {
        [key: string]: Element
    };

    constructor() {
        super();
        this.state = {

        };
    }

    render(): React.ReactElement<DefaultModal> {
        return (
            <DefaultModal open={this.props.open}>
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
                                          <label className="muted" htmlFor="labelName">Label name</label>
                                          <input name="labelName" type="text" />
                                      </div>
                                      <div className="form-item">
                                          <label className="muted" htmlFor="labelDescription">Label description</label>
                                          <textarea name="labelDescription" rows="4" cols="50"></textarea>
                                      </div>
                                    </div>
                                </Cell>
                            </Row>
                            <Row className="button-row">
                                <Cell className="table-row-name"></Cell>
                                <Cell className="table-row-item">
                                    <button className="default" onClick={this.props.save.bind(this) }>
                                        Save
                                    </button>
                                    <button className="default invert" onClick={this.props.cancel.bind(this) }>
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
