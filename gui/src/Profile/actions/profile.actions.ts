/**
 * Created by justin on 7/27/16.
 */
export const RECEIVE_PROFILE = 'RECEIVE_PROFILE';

function receiveProfile(profile) {
  return {
    type: RECEIVE_PROFILE,
    profile
  };
}

export function setProfile(profile) {
  return (dispatch) => {
    let profileString = JSON.stringify(profile);
    localStorage.setItem('steamProfile', profileString);
    dispatch(receiveProfile(profile));
  };
}

export function fetchProfile() {
  return (dispatch) => {
    let profile = localStorage.getItem('steamProfile');
    dispatch(receiveProfile(JSON.parse(profile)));
  };
}
