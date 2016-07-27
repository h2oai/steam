package h2ov3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/h2oai/steamY/bindings"
)

///////////////////
///////////////////
////// Cloud //////
///////////////////
///////////////////

// GetCloudStatus Determine the status of the nodes in the H2O cloud.
func (h *H2O) GetCloudStatus() (*bindings.CloudV3, error) {
	//@GET
	u := h.url("/3/Cloud")

	res, err := http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("H2O get request failed: %s: %s", u, err)
	}

	data, err := h.handleResponse(res, u)
	if err != nil {
		return nil, err
	}

	var out bindings.CloudV3
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("H2O response unmarshal failed: %v", err)
	}
	return &out, nil
}

////////////////////
////////////////////
////// Frames //////
////////////////////
////////////////////

// GetFramesFetch Return the specified Frame. */
func (h *H2O) GetFramesFetch(frame_id string) ([]byte, *bindings.FramesV3, error) {
	//@GET
	u := h.url("/3/Frames/?{frame_id}", frame_id)

	res, err := http.Get(u)
	if err != nil {
		return nil, nil, fmt.Errorf("H2O get request failed: %s: %s", u, err)
	}

	data, err := h.handleResponse(res, u)
	if err != nil {
		return nil, nil, err
	}

	var out bindings.FramesV3
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, nil, fmt.Errorf("H2O response unmarshal failed: %v", err)
	}
	return data, &out, nil
}

// GetFramesList Return all Frames in the H2O distributed K/V store. */
func (h *H2O) GetFramesList() (*bindings.FramesV3, error) {
	//@GET
	u := h.url("/3/Frames")

	res, err := http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("H2O get request failed: %s: %s", u, err)
	}

	data, err := h.handleResponse(res, u)
	if err != nil {
		return nil, err
	}

	var out bindings.FramesV3
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("H2O response unmarshal failed: %v", err)
	}

	return &out, nil
}

////////////////////
////////////////////
////// InitID //////
////////////////////
////////////////////

// GetInitIDIssue Issue a new session ID. */
func (h *H2O) GetInitIDIssue() (*bindings.InitIDV3, error) {
	//@GET
	u := h.url("/3/InitID")

	res, err := http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("H2O get request failed: %s: %s", u, err)
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

//////////////////
//////////////////
////// Jobs //////
//////////////////
//////////////////

// GetJobsList Get a list of all the H2O Jobs (long-running actions). */
func (h *H2O) GetJobsList() (*bindings.JobsV3, error) {
	//@GET
	u := h.url("/3/Jobs")

	res, err := http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("H2O get request failed: %s: %s", u, err)
	}

	data, err := h.handleResponse(res, u)
	if err != nil {
		return nil, err
	}

	var out bindings.JobsV3
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("H2O response unmarshal failed: %v", err)
	}
	return &out, nil
}

// GetJobsFetch Get the status of the given H2O Job (long-running action). */
func (h *H2O) GetJobsFetch(job_id string) (*bindings.JobsV3, error) {
	//@GET
	u := h.url("/3/Jobs/?{job_id}", job_id)

	res, err := http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("H2O get request failed: %s: %s", u, err)
	}

	data, err := h.handleResponse(res, u)
	if err != nil {
		return nil, err
	}

	var out bindings.JobsV3
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("H2O response unmarshal failed: %v", err)
	}
	return &out, nil
}

////////////////////
////////////////////
////// Models //////
////////////////////
////////////////////

// GetModelsFetch Return the specified Model from the H2O distributed K/V store, optionally with the list of compatible Frames. */
func (h *H2O) GetModelsFetch(model_id string) ([]byte, *bindings.ModelsV3, error) {
	//@GET
	u := h.url("/3/Models/?{model_id}", model_id)

	res, err := http.Get(u)
	if err != nil {
		return nil, nil, fmt.Errorf("H2O get request failed: %s: %s", u, err)
	}

	data, err := h.handleResponse(res, u)
	if err != nil {
		return nil, nil, err
	}

	var out bindings.ModelsV3
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, nil, fmt.Errorf("H2O response unmarshal failed: %v", err)
	}
	return data, &out, nil
}

// GetModelsList Return all Models from the H2O distributed K/V store. */
func (h *H2O) GetModelsList() (*bindings.ModelsV3, error) {
	//@GET
	u := h.url("/3/Models")

	res, err := http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("H2O get request failed: %s: %s", u, err)
	}

	data, err := h.handleResponse(res, u)
	if err != nil {
		return nil, err
	}

	var out bindings.ModelsV3
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("H2O response unmarshal failed: %v", err)
	}
	return &out, nil
}

//////////////////////////
//////////////////////////
////// ModelMetrics //////
//////////////////////////
//////////////////////////

// GetModelMetricsListSchemaFetchByModel Return the saved scoring metrics for the specified Model. */
func (h *H2O) GetModelMetricsListSchemaFetchByModel(model string) ([]byte, *bindings.ModelMetricsListSchemaV3, error) {
	//@GET
	u := h.url("/3/ModelMetrics/models/?{model}", model)

	res, err := http.Get(u)
	if err != nil {
		return nil, nil, fmt.Errorf("H2O get request failed: %s: %s", u, err)
	}

	data, err := h.handleResponse(res, u)
	if err != nil {
		return nil, nil, err
	}

	var out bindings.ModelMetricsListSchemaV3
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, nil, fmt.Errorf("H2O response unmarshal failed: %v", err)
	}
	return data, &out, nil
}
