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

package h2ov3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/h2oai/steam/bindings"
)

////////////////////
////////////////////
////// InitID //////
////////////////////
////////////////////

// DeleteInitIDEndsession End a session. */
func (h *H2O) DeleteInitIDEndsession() (*bindings.InitIDV3, error) {
	//@DELETE
	u := h.url("3/InitID")

	req, err := http.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, fmt.Errorf("error making delete request: %s: %s", u, err)
	}

	if h.Token != "" {
		req.Header.Set("Authorization", "Basic "+h.Token)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("H2O delete request failed: %s: %s", u, err)
	}

	data, err := h.handleResponse(res, u)
	if err != nil {
		return nil, err
	}

	var out bindings.InitIDV3
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("H2O response unmarshal failed: %v", err)
	}
	return &out, nil
}
