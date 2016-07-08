/**
 * Created by justin on 7/8/16.
 */
jest.unmock('../Models');

import * as React from 'react';
import * as ReactDOM from 'react-dom';
import * as TestUtils from 'react-addons-test-utils';

const Models = require('../Models');

describe('Models', () => {
  it('exists', () => {
    expect(Models).toBeDefined();  
  });  
});
