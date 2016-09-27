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

import { routes } from '../../routes';

/*
* Takes a path and converts parameters to their appropriate
* @function
* @param {string} path
* @returns {object} Returns a route object
*/

export const getRoute = (path: string): any => {
    let route = null;
    for (let i = 0; i < routes[0].childRoutes.length; i++) {
        if (routes[0].childRoutes[i].path === path) {
            route = routes[0].childRoutes[i];
            break;
        }
    }
    return route;
};
