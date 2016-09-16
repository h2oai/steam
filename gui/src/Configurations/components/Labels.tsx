/**
 * Created by justin on 7/12/16.
 */
import * as React from 'react';
import Table from '../../Projects/components/Table';
import Row from '../../Projects/components/Row';
import Cell from '../../Projects/components/Cell';
import CreateNewLabelModal from './CreateNewLabelModal';
import { fetchLabels, createLabel, deleteLabel, updateLabel } from '../actions/configuration.labels.action';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import '../styles/labels.scss';

import * as _ from 'lodash';

interface Props {
    labels?: any
    projectid: string
    fetchLabels: Function
    createLabel: Function
    deleteLabel: Function
    updateLabel: Function
}

interface State {
  modalOpen: boolean
  label: any
}

class Labels extends React.Component<Props, any> {

    constructor() {
        super();
        this.state = {
            modalOpen: false,
            label: {}
        };
    }

    componentWillMount() {
        if (!this.props.labels || !this.props.labels[this.props.projectid]) {
            this.props.fetchLabels(this.props.projectid);
        }
    }

    openModal() {
        this.setState({
            modalOpen: true
        });
    }

    closeModal() {
        this.setState({
            modalOpen: false,
            label: {}
        });
    }

    saveUpdateLabel(label) {
      if (!label.id) {
        let newLabel = {
          name: label.name,
          description: label.description
        };
        this.saveLabel(label);
      } else {
        this.updateLabel(label);
      }
    }

    updateLabel(label) {
      this.props.updateLabel(parseInt(label.id, 10), this.props.projectid, label.name, label.description).then((response) => {
          this.props.fetchLabels(this.props.projectid);
          this.closeModal();
      }, (error) => {
          alert(error);
      });
    }

    saveLabel(label) {
        this.props.createLabel(parseInt(this.props.projectid, 10), label.name, label.description).then((response) => {
            this.props.fetchLabels(this.props.projectid);
            this.closeModal();
        }, (error) => {
            alert(error);
        });
    }

    deleteLabel(labelId) {
      this.props.deleteLabel(labelId).then((response) => {
          this.props.fetchLabels(this.props.projectid);
      }, (error) => {
          alert(error);
      });
    }

    renderLabels() {
        if (!this.props.labels || !this.props.projectid || !this.props.labels[this.props.projectid]) {
            return null;
        }
        return this.props.labels[this.props.projectid].map((label) => {
            let deleteLabel = () => {
              this.deleteLabel(label.id);
            };
            let updateLabel = (event) => {
              this.setState({
                label: {
                  id: label.id,
                  name: label.name,
                  description: label.description
                },
                modalOpen: true
              });
            };
            return (
                <Row key={label.id}>
                    <Cell className="label-bullets">
                        <span className="label-bullet"></span>
                    </Cell>
                    <Cell className="label-names">
                        <div className="label-name">{label.name}</div>
                        <div className="label-description muted">{label.description}</div>
                    </Cell>
                    <Cell className="label-model">
                        <span className="model-icon"></span>
                        <span className="model-name">{(label.model_id >= 0) ? (<span className="fa fa-cube"></span>) : null} {(label.model_id >= 0) ? label.model_id : "Not currently applied to a model"}</span>
                    </Cell>
                    <Cell className="label-permissions">

                    </Cell>
                    <Cell>
                        <span className="fa fa-pencil" onClick={updateLabel}></span>
                        <span className="fa fa-trash" onClick={deleteLabel}></span>
                    </Cell>
                </Row>
            );
        });
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
                    <button className="button-primary" onClick={this.openModal.bind(this) }>
                        Create New Label
                    </button>
                </span>
                <CreateNewLabelModal label={this.state.label} open={this.state.modalOpen} cancel={this.closeModal.bind(this) } save={this.saveUpdateLabel.bind(this)} />
                <div className="label-table">
                    <Table>
                        <Row className="head">
                            <Cell/>
                            <Cell>Label</Cell>
                            <Cell>Model</Cell>
                            <Cell>Permissions</Cell>
                            <Cell/>
                        </Row>
                        {this.renderLabels() }
                    </Table>
                </div>
            </div>
        );
    }
}

function mapStateToProps(state: any): any {
    return {
        labels: state.labels
    };
}

function mapDispatchToProps(dispatch) {
    return {
        fetchLabels: bindActionCreators(fetchLabels, dispatch),
        createLabel: bindActionCreators(createLabel, dispatch),
        deleteLabel: bindActionCreators(deleteLabel, dispatch),
        updateLabel: bindActionCreators(updateLabel, dispatch)
    };
}

export default connect<Props, any, any>(mapStateToProps, mapDispatchToProps)(Labels);
