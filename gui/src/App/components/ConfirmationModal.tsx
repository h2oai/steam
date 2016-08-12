/**
 * Created by justin on 8/12/16.
 */
import * as React from 'react';
import DefaultModal from './DefaultModal';
import PageHeader from '../../Projects/components/PageHeader';
import { connect } from 'react-redux';
import '../styles/confirmation.scss';

interface Props {
  confirmation: {
    isOpen: boolean,
    onYes: Function,
    onNo: Function,
    text: string,
    title: string
  }
}

export class Notification extends React.Component<Props, any> {
  onYes() {
    this.props.confirmation.onYes();
  }

  onNo() {
    this.props.confirmation.onNo();
  }

  render(): React.ReactElement<DefaultModal> {
    return (
      <DefaultModal className="confirmation-modal"
               open={this.props.confirmation.isOpen}>
        <div className="content">
          <PageHeader>
            {this.props.confirmation.title}
          </PageHeader>
          <h3>{this.props.confirmation.text}</h3>
        </div>
        <div className="confirmation-buttons">
          <button className="default" onClick={this.onYes.bind(this)}>Yes</button>
          <button className="default invert" onClick={this.onNo.bind(this)}>No</button>
        </div>
      </DefaultModal>
    );
  }
}

function mapStateToProps(state) {
  return {
    confirmation: state.confirmation
  };
}

export default connect<any, any, any>(mapStateToProps, {})(Notification);
