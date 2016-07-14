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

func (h *H2O) SplitFrame(frameId string, seed int, ratios ...float32) ([]string, string) {

	var (
		buf   bytes.Buffer
		i     int
		ratio float32
		split string
	)
	tmpVec := frameId + "_TMP"

	buf.WriteString("(,")                                                                              // 1 open
	buf.WriteString("  (tmp= " + tmpVec + " (h2o.runif " + frameId + " " + strconv.Itoa(seed) + ")) ") // 3 open 2 close

	splitIds := make([]string, len(ratios)+1)
	for i, ratio = range ratios {
		split = strconv.FormatFloat(float64(ratio), 'f', 2, 64)
		splitIds[i] = frameId + "_split_" + strconv.Itoa(i)
		buf.WriteString("(assign " + splitIds[i] + " ") // 2 open
		buf.WriteString("  (rows " + frameId + " ")     // 3 open
		if i > 0 {
			prev := strconv.FormatFloat(float64(ratios[i-1]), 'f', 2, 64)
			buf.WriteString("(& (> " + frameId + " " + prev + ")")
		}
		buf.WriteString("(<= " + tmpVec + " " + split + "))") // 4 open 2 close
		buf.WriteString(")")                                  // 1 close
		if i > 0 {
			buf.WriteString(")")
		}
	}
	splitIds[i+1] = frameId + "_split_" + strconv.Itoa(i+1)
	buf.WriteString("(assign " + splitIds[i+1] + " ")                            // 2 open
	buf.WriteString("  (rows " + frameId + " (> " + tmpVec + " " + split + "))") // 4 open 2 close
	buf.WriteString(")")                                                         // 1 close
	buf.WriteString("(rm " + tmpVec + "))")                                      // 2 open 2 close

	return splitIds, buf.String()
}
