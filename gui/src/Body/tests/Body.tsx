/**
 * Created by justin on 8/16/16.
 */
jest.unmock('../Body');

import { shallow } from 'enzyme';
import * as React from 'react';
import Body from '../Body';

describe('Body', () => {
  beforeEach(() => {
    this.innerElements = <div>Test</div>;
    this.wrapper = shallow(<Body>{this.innerElements}</Body>);
  });

  it('exists', () => {
    expect(this.wrapper).toBeDefined();
  });
  
  it('passes children', () => {
    expect(this.wrapper.containsAllMatchingElements(this.innerElements));  
  });
});
