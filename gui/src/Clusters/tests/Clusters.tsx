/**
 * Created by justin on 8/16/16.
 */

jest.unmock('../Clusters');

import { shallow } from 'enzyme';
import * as React from 'react';
import { Clusters } from '../Clusters';

describe('Clusters', () => {
  beforeEach(() => {
    let mockClusters = [
      {
        id: 0,
        name: 'TestCluster',
        type_id: 0,
        detail_id: 0,
        address: 'http://localhost:54321',
        state: 'started',
        created_at: 1471370542359
      }
    ];
    this.mocks = {
      fetchClusters: jest.fn().mockReturnValue(mockClusters),
      unregisterCluster: jest.fn(),
      mockClusters: mockClusters
    };
    this.wrapper = shallow(<Clusters fetchClusters={this.mocks.fetchClusters} clusters={this.mocks.mockClusters} unregisterCluster={this.mocks.unregisterCluster}></Clusters>);
  });

  it('exists', () => {
    expect(this.wrapper).toBeDefined();
  });

  it('should not fetch clusters if already exists', () => {
    expect(this.mocks.fetchClusters).not.toBeCalled();
  });
  
  it('should fetch clusters if already doesn\'t already exist', () => {
    this.wrapper = shallow(<Clusters fetchClusters={this.mocks.fetchClusters} unregisterCluster={this.mocks.unregisterCluster} clusters={undefined}></Clusters>);
    expect(this.mocks.fetchClusters).toBeCalled();
  });


  it('should call removeCluster', () => {
    this.wrapper = shallow(<Clusters fetchClusters={this.mocks.fetchClusters} clusters={this.mocks.mockClusters} unregisterCluster={this.mocks.unregisterCluster}></Clusters>);
    this.wrapper.find('button').simulate('click');
    expect(this.mocks.unregisterCluster).toBeCalled();
  });
});
