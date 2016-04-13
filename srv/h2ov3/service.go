package h2ov3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/h2oai/steamY/lib/fs"
)

type H2O struct {
	Address string
}

func NewClient(address string) *H2O {
	return &H2O{
		address,
	}
}

func (h *H2O) url(p string) string {
	return (&url.URL{Scheme: "http", Host: h.Address, Path: p}).String()
}

type H2OException struct {
	Message string `json:"msg"`
	// Stacktrace []string `json:"stacktrace"`
}

func readException(b []byte) (*H2OException, error) {
	var exception H2OException
	if err := unmarshal(b, &exception); err != nil {
		return nil, err
	}
	return &exception, nil
}

func (h *H2O) handleResponse(u string, res *http.Response, httperr error) ([]byte, error) {
	if httperr != nil {
		return nil, fmt.Errorf("H2O request failed: %s: %s", u, httperr)
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("H2O response read failed: %s: %v:", u, err)
	}

	if res.StatusCode != 200 {
		if ex, err := readException(b); err == nil {
			return nil, fmt.Errorf("H2O failed: %s", ex.Message)
		}
		return nil, fmt.Errorf("H2O response status failed: %s: %s / %s", u, res.Status, string(b))
	}

	return b, nil
}

func (h *H2O) get(s string, data url.Values) ([]byte, error) {
	u := h.url(s)
	if data != nil {
		u += "?" + data.Encode()
	}
	res, err := http.Get(u)
	return h.handleResponse(u, res, err)
}

func (h *H2O) post(s string, data url.Values) ([]byte, error) {
	u := h.url(s)
	res, err := http.PostForm(u, data)
	return h.handleResponse(u, res, err)
}

func (h *H2O) download(s, p string, preserveFilename bool) (string, error) {
	u := h.url(s)

	_, filepath, err := fs.Download(p, u, preserveFilename)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

func unmarshal(b []byte, v interface{}) error {
	err := json.Unmarshal(b, v)
	if err != nil {
		return fmt.Errorf("H2O response unmarshal failed: %v", err)
	}
	return nil
}

type Cloud struct {
	Version      string `json:"version"` //TODO use this to determine json decoder
	Name         string `json:"cloud_name"`
	Size         int    `json:"cloud_size"`
	IsHealthy    bool   `json:"cloud_healthy"`
	BadNodes     int    `json:"bad_nodes"`
	HasConsensus bool   `json:"consensus"`
	IsLocked     bool   `json:"locked"`
}

func (h *H2O) GetCloud() (*Cloud, error) {
	b, err := h.get("/3/Cloud", nil)
	if err != nil {
		return nil, err
	}

	var cloud Cloud
	if err := unmarshal(b, &cloud); err != nil {
		return nil, err
	}

	return &cloud, nil
}

type RawModel struct {
	ID   string
	Data []byte
}

type ModelSynopsisV3 struct {
	Algo               string     `json:"algo"`
	AlgoFullName       string     `json:"algo_full_name"`
	DataFrame          FrameKeyV3 `json:"data_frame"`
	ModelId            ModelKeyV3 `json:"model_id"`
	ResponseColumnName string     `json:"response_column_name"`
	Timestamp          int64      `json:"timestamp"`
}

type GetModelsResponseV3 struct {
	Models []*ModelSynopsisV3 `json:"models"`
}

type FrameKeyV3 struct {
	Name string `json:"name"`
}

type ModelKeyV3 struct {
	Name string `json:"name"`
}

type AutoMLBuilderV3 struct {
	Job JobV3 `json:"job"`
}

type AutoMLKeyV3 struct {
	Name string `json:"name"`
}

type AutoMLV3 struct {
	Leader ModelKeyV3 `json:"leader"`
}

type JobsV3 struct {
	Jobs []*JobV3 `json:"jobs"`
}

type JobV3 struct {
	JobKey    JobKeyV3    `json:"key"`
	Status    string      `json:"status"`
	Dest      AutoMLKeyV3 `json:"dest"` //FIXME: This should be a general key
	Exception string      `json:"exception"`
}

type JobKeyV3 struct {
	Name string `json:"name"`
}

func (h *H2O) jobPoll(jobID string) (*JobV3, error) {
	var (
		j JobsV3
		k JobV3
	)

	for { // Polling for job
		b, err := h.get("/3/Jobs/"+jobID, nil)
		if err != nil {
			return nil, err
		}
		if err := unmarshal(b, &j); err != nil {
			return nil, err
		}
		k = *j.Jobs[0]
		if k.Status == "CREATED" || k.Status == "RUNNING" {
			continue
		} else {
			if k.Exception != "" {
				return nil, fmt.Errorf("Error running AutoML: %s", k.Exception)
			}
		}
		break
	}

	return &k, nil
}

// AutoML will use the built in AutoML tool to create a model with minimal
// user interaction.
func (h *H2O) AutoML(dataset, targetName string, maxTime int) (string, error) {
	b, err := h.post("/3/AutoMLBuilder", url.Values{
		"dataset":     {dataset},
		"target_name": {targetName},
		"max_time":    {strconv.Itoa(maxTime)},
	})
	if err != nil {
		return "", err
	}
	var a AutoMLBuilderV3
	if err := unmarshal(b, &a); err != nil {
		return "", err
	}

	k, err := h.jobPoll(a.Job.JobKey.Name)

	b, err = h.get("/3/AutoML/"+k.Dest.Name, nil)
	if err != nil {
		return "", err
	}
	var m AutoMLV3
	if err = unmarshal(b, &m); err != nil {
		return "", err
	}

	if m.Leader.Name == "" {
		return "", fmt.Errorf("Unable to complete model in %d seconds", maxTime)
	}

	return m.Leader.Name, err
}

func (h *H2O) ExportJavaModel(modelID, p string) error {
	_, err := h.download("/3/Models.java/"+modelID, p, true)
	if err != nil {
		return fmt.Errorf("Java model export failed: %s", err)
	}
	return nil
}

func (h *H2O) ExportGenModel(p string) error {
	if _, err := h.download("/3/h2o-genmodel.jar", path.Join(p, "h2o-genmodel.jar"), false); err != nil {
		return fmt.Errorf("Java genmodel jar export failed: %s", err)
	}
	return nil
}

func (h *H2O) CompilePojo(javaModel, jar string) error {
	if _, err := h.post("/compile", url.Values{
		"pojo": {javaModel},
		"jar":  {jar},
	}); err != nil {
		return err
	}
	return nil
}

func (h *H2O) Shutdown() error {
	_, err := h.post("/3/Shutdown", url.Values{})
	if err != nil {
		return err
	}
	return nil
}
