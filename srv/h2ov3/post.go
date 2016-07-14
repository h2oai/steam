package h2ov3

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/h2oai/steamY/bindings"
)

/////////////////
/////////////////
////// GBM //////
/////////////////
/////////////////

// PostGBMTrain Train a GBM model. */
func (h *H2O) PostGBMTrain(in *bindings.GBMParametersV3) (*bindings.GBMV3, error) {
	//@POST
	u := h.url("/3/ModelBuilders/gbm")

	v := url.Values{
		"training_frame":  {in.TrainingFrame.Name},
		"response_column": {in.ResponseColumn.ColumnName},
	}

	if in.ValidationFrame != nil {
		v["validation_frame"] = []string{in.ValidationFrame.Name}
	}

	res, err := http.PostForm(u, v)
	if err != nil {
		return nil, fmt.Errorf("H2O post request failed: %s: %s", u, err)
	}

	data, err := h.handleResponse(res, u)
	if err != nil {
		return nil, err
	}

	var out bindings.GBMV3
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("H2O response unmarshal failed: %v", err)
	}
	return &out, nil
}

/////////////////////////
/////////////////////////
////// ImportFiles //////
/////////////////////////
/////////////////////////

// PostImportFilesImportfiles Import raw data files into a single-column H2O Frame. */
func (h *H2O) PostImportFilesImportfiles(path string) (*bindings.ImportFilesV3, error) {
	//@POST
	u := h.url("/3/ImportFiles")

	// FIXME: remaining URL values req'd
	v := url.Values{
		"path": {path},
	}

	res, err := http.PostForm(u, v)
	if err != nil {
		return nil, fmt.Errorf("H2O post request failed: %s: %s", u, err)
	}

	data, err := h.handleResponse(res, u)
	if err != nil {
		return nil, err
	}

	var out bindings.ImportFilesV3
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("H2O response unmarshal failed: %v", err)
	}
	return &out, nil
}

///////////////////
///////////////////
////// Parse //////
///////////////////
///////////////////

// PostParseParse Parse a raw byte-oriented Frame into a useful columnar data Frame. */
func (h *H2O) PostParseParse(in *bindings.ParseV3) (*bindings.ParseV3, error) {
	//@POST
	u := h.url("/3/Parse")

	defaultParse := bindings.NewParseV3()

	// FIXME: remaining URL values req'd
	v := url.Values{
		"destination_frame": {in.DestinationFrame.Name},
		"source_frames":     keyArrayToStringArray(in.SourceFrames),
	}

	if in.ParseType != defaultParse.ParseType {
		v["parse_type"] = []string{string(in.ParseType)}
	}
	if in.Separator != defaultParse.Separator {
		v["separator"] = []string{strconv.FormatInt(int64(in.Separator), 10)}
	}
	if in.SingleQuotes != defaultParse.SingleQuotes {
		v["single_quotes"] = []string{strconv.FormatBool(in.SingleQuotes)}
	}
	if in.CheckHeader != defaultParse.CheckHeader {
		v["check_header"] = []string{strconv.FormatInt(int64(in.CheckHeader), 10)}
	}
	if in.NumberColumns != defaultParse.NumberColumns {
		v["number_columns"] = []string{strconv.FormatInt(int64(in.NumberColumns), 10)}
	}
	if in.ColumnNames != nil {
		v["column_names"] = in.ColumnNames
	}
	if in.ColumnTypes != nil {
		v["column_types"] = in.ColumnTypes
	}
	if in.Domains != nil {
		v["domains"] = DoubleStringArraysToSingle(in.Domains)
	}
	if in.NaStrings != nil {
		v["na_strings"] = DoubleStringArraysToSingle(in.NaStrings)
	}
	if in.ChunkSize != defaultParse.ChunkSize {
		v["chunk_size"] = []string{strconv.FormatInt(int64(in.ChunkSize), 10)}
	}
	if in.DeleteOnDone != defaultParse.DeleteOnDone {
		v["delete_on_done"] = []string{}
	}
	if in.Blocking != defaultParse.Blocking {
		v["blocking"] = []string{}
	}
	if in.ExcludeFields != defaultParse.ExcludeFields {
		v["_exclude_fields"] = []string{}
	}

	res, err := http.PostForm(u, v)
	if err != nil {
		return nil, err
	}
	data, err := h.handleResponse(res, u)
	if err != nil {
		return nil, err
	}

	var out bindings.ParseV3
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("H2O response unmarshal failed: %v", err)
	}

	return &out, nil
}

////////////////////////
////////////////////////
////// ParseSetup //////
////////////////////////
////////////////////////

// PostParseSetupGuesssetup Guess the parameters for parsing raw byte-oriented data into an H2O Frame. */
func (h *H2O) PostParseSetupGuesssetup(sourceFrames []string) (*bindings.ParseSetupV3, error) {
	//@POST
	u := h.url("/3/ParseSetup")

	// FIXME: remaining URL values req'd
	v := url.Values{
		"source_frames": sourceFrames,
	}

	res, err := http.PostForm(u, v)
	if err != nil {
		return nil, err
	}

	data, err := h.handleResponse(res, u)
	if err != nil {
		return nil, err
	}

	var out bindings.ParseSetupV3
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("H2O response unmarshal failed: %v", err)
	}
	return &out, nil
}
