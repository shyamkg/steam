/**
 * Created by justin on 7/8/16.
 */

import 'jest';

jest.unmock('../Models');

import * as React from 'react';

import Models from '../Models';

describe('Models', () => {
  it('exists', () => {
    expect(Models).toBeDefined();  
  });  
});