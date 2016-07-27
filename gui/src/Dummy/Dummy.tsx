/**
 * Created by Jeff Fohl <jfohl@h2o.ai> on 7/17/16.
 * Dummy component for mocking
 */

import * as React from 'react';

interface DummyProps {};

export default class Dummy extends React.Component<DummyProps, any> {

  constructor() {
    super();
  }

  render(): React.ReactElement<HTMLDivElement> {
    return (
      <div>
        <h1>Content to come.</h1>
      </div>
     );
   }

};
