/**
 * Created by justin on 7/12/16.
 */
import * as React from 'react';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import CreateNewLabelModal from './CreateNewLabelModal';
import '../styles/labels.scss';

interface Props {
  labels: any
}

export default class Labels extends React.Component<Props, any> {

    constructor() {
        super();
        this.state = {
            modalOpen: false
        };
    }

    openModal() {
        this.setState({
            modalOpen: true
        });
    }

    closeModal() {
        this.setState({
            modalOpen: false
        });
    }

    newLabel(label) {
        this.closeModal();
    }

    render(): React.ReactElement<HTMLDivElement> {
        return (
            <div className="labels">
                <h1>Model Labels</h1>
                <p>
                    You can create labels for the models. A label can only be associated
                    with one model at a time. You can give different team members various
                    permissions.
                </p>
                <p>
                    For example, you can create "test" and "production" labels. You could
                    then allow the entire team to label a model "test", but only give admins
                    the power to label a model "production".
                </p>
                <span>
                    <button className="default" onClick={this.openModal.bind(this) }>
                        Create New Label
                    </button>
                </span>
                <CreateNewLabelModal open={this.state.modalOpen} cancel={this.closeModal.bind(this) } save={this.newLabel.bind(this) }/>
                <div className="label-table">
                  <Table>
                    <Row className="head">
                      <Cell/>
                      <Cell>Label</Cell>
                      <Cell>Model</Cell>
                      <Cell>Permissions</Cell>
                      <Cell/>
                    </Row>
                    <Row>
                      <Cell className="label-bullets">
                        <span className="label-bullet"></span>
                      </Cell>
                      <Cell className="label-names">
                        <div className="label-name">test</div>
                        <div className="label-description muted">Label description</div>
                      </Cell>
                      <Cell className="label-model">
                        <span className="model-icon"></span>
                        <span className="model-name"><span className="fa fa-cube"></span> DRF-1568963</span>
                      </Cell>
                      <Cell className="label-permissions">
                        <ul>
                          <li>Anyone on ISRM team can edit.</li>
                          <li>Anyone on the project can view.</li>
                        </ul>
                      </Cell>
                      <Cell>
                        <span className="fa fa-pencil"></span>
                        <span className="fa fa-trash"></span>
                      </Cell>
                    </Row>
                  </Table>
                </div>
            </div>
        );
    }
}
