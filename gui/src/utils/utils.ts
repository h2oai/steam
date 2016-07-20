/*
* Utility functions for Steam UI
*/

/*
* Takes an integer and returns a human readable ordinal.
* @function
* @example
* // returns 'st'
* getOrdinal(1);
* @param {number} rank
* @returns {string} Returns the ordinal of the supplied integer.
*/
export const getOrdinal = (rank: number): string => {
  let suffixes = ['th', 'st', 'nd', 'rd'];
  let remainder = rank % 100;
  return (suffixes[(remainder - 20) % 10] || suffixes[remainder] || suffixes[0]);
};
