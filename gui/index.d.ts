/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

/**
 * Created by justin on 7/13/16.
 */

declare function require(string): string;

interface Window {
  devToolsExtension: any
  ga: any
  chrome?: {
    webstore: any
  }
  InstallTrigger?: any
}

declare module 'react-router-redux' {
  export var syncHistoryWithStore: any;
  export var routerReducer: any;
}

declare module 'react-router' {
  export var Link: any;
  export var hashHistory: any;
  export var withRouter: any;
  export var Router: any;
  export var PlainRoute: any;
}
