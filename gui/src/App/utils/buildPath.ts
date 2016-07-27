/*
* Takes a path and converts parameters to their appropriate
* @function
* @param {string} path
* @param {object} params
* @returns {string} Returns a string with the filled in params.
*/
export const buildPath = (path: string, params: any): string => {
  let parts = path.split('/');
  parts = parts.map((part) => {
    if (part[0] === ':') {
      let newPart = part.slice(1, part.length);
      if (params[newPart]) {
        return params[newPart];
      }
    }
    return part;
  });
  return '/' + parts.join('/');
};
