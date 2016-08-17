/**
 * Created by justin on 8/16/16.
 */
jest.unmock('../Configurations');

import { shallow } from 'enzyme';
import * as React from 'react';
import Configurations from '../Configurations';

describe('Configurations', () => {
  beforeEach(() => {
    this.wrapper = shallow(<Configurations params={{projectid: "0"}}></Configurations>);
  });

  it('exists', () => {
    expect(this.wrapper).toBeDefined();
  });
});
