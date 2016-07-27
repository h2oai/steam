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
}
