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
	u := h.url("/3/InitID")

	req, err := http.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, fmt.Errorf("error making delete request: %s: %s", u, err)
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
