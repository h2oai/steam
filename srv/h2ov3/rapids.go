package h2ov3

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"

	"github.com/h2oai/steamY/bindings"
)

// PostRapidsExec Execute an Rapids AST.
func (h *H2O) PostRapidsExec(in *bindings.RapidsSchema) ([]byte, error) {
	//@POST

	u := h.url("/99/Rapids")

	defaultIn := bindings.NewRapidsSchema()

	v := url.Values{
		"ast": {in.Ast},
	}

	if in.Id != defaultIn.Id {
		v["id"] = []string{in.Id}
	}
	if in.SessionId != defaultIn.SessionId {
		v["session_id"] = []string{in.SessionId}
	}

	res, err := http.PostForm(u, v)
	if err != nil {
		return nil, err
	}
	data, err := h.handleResponse(res, u)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *H2O) SplitFrame(frameId string, seed int, ratios ...float32) string {

	var buf bytes.Buffer

	buf.WriteString("(, (tmp= " + tmpVec + "(h2o.runif " + frameId + " " + strconv.Itoa(seed) + "))")
	for i, ratio := range ratios {
		val := strconv.FormatFloat(ratio, 'f', 2, 32)
		buf.WriteString("assign split" + strconv.Itoa(i) + "(rows " + frameId + " (")
		if i < len(ratios) {
			buf.WriteString("<= " + tmpVec + " " + val)
		}
	}

	return
}
