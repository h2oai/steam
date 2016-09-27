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
"use strict";
/**
 * Created by justin on 7/27/16.
 */
exports.RECEIVE_PROFILE = 'RECEIVE_PROFILE';
function receiveProfile(profile) {
    return {
        type: exports.RECEIVE_PROFILE,
        profile: profile
    };
}
function setProfile(profile) {
    return function (dispatch) {
        var profileString = JSON.stringify(profile);
        localStorage.setItem('steamProfile', profileString);
        dispatch(receiveProfile(profile));
    };
}
exports.setProfile = setProfile;
function fetchProfile() {
    return function (dispatch) {
        var profile = localStorage.getItem('steamProfile');
        dispatch(receiveProfile(JSON.parse(profile)));
    };
}
exports.fetchProfile = fetchProfile;
//# sourceMappingURL=profile.actions.js.map