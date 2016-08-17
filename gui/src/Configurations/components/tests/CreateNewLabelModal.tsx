/**
 * Created by justin on 8/16/16.
 */
jest.unmock('../CreateNewLabelModal');

import { shallow } from 'enzyme';
import * as React from 'react';
import CreateNewLabelModal from '../CreateNewLabelModal';

describe('CreateNewLabelModal', () => {
  beforeEach(() => {
    this.mocks = {
      cancel: jest.fn(),
      open: jest.fn(),
      save: jest.fn()
    };
    this.wrapper = shallow(<CreateNewLabelModal cancel={this.mocks.cancel} open={this.mocks.open} save={this.mocks.save}></CreateNewLabelModal>);
  });

  it('exists', () => {
    expect(this.wrapper).toBeDefined();
  });
});
