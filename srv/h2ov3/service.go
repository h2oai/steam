package h2ov3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strconv"

	"github.com/h2oai/steamY/bindings"
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

	nan := regexp.MustCompile(`"NaN"`)
	b = nan.ReplaceAllLiteral(b, []byte("0"))

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

type ModelKeyV3 struct {
	Name string `json:"name"`
}

type AutoMLBuilderV3 struct {
	Job bindings.JobV3 `json:"job"`
}

type AutoMLKeyV3 struct {
	Name string `json:"name"`
}

type AutoMLV3 struct {
	Leader bindings.ModelKeyV3 `json:"leader"`
}

func (h *H2O) jobPoll(jobID string) (*bindings.JobV3, error) {
	var (
		j bindings.JobsV3
		k bindings.JobV3
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

	k, err := h.jobPoll(a.Job.Key.Name)

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
	log.Printf("Completed AutoML with model: %s", m.Leader.Name)
	return m.Leader.Name, err
}

func (h *H2O) ExportJavaModel(modelID, p string) (string, error) {
	f, err := h.download("/3/Models.java/"+modelID, p, true)
	if err != nil {
		return "", fmt.Errorf("Java model export failed: %s", err)
	}
	return f, nil
}

func (h *H2O) ExportGenModel(p string) (string, error) {
	f, err := h.download("/3/h2o-genmodel.jar", path.Join(p, "h2o-genmodel.jar"), false)
	if err != nil {
		return "", fmt.Errorf("Java genmodel jar export failed: %s", err)
	}
	return f, nil
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
	_, err := h.post("/3/Shutdown", nil)
	if err != nil {
		return err
	}
	return nil
}

// TODO: THE FOLLOWING METHODS SHOULD BE IN GENERATED FILES
func (h *H2O) GetCloud() (*bindings.CloudV3, error) {
	b, err := h.get("/3/Cloud", nil)
	if err != nil {
		return nil, fmt.Errorf("Error from get in GetCloud:\n%v", err)
	}

	var c bindings.CloudV3
	if err := unmarshal(b, &c); err != nil {
		return nil, fmt.Errorf("Error from unmarshal in GetCloud:\n%v", err)
	}

	return &c, nil
}

func (h *H2O) GetModels() (*bindings.ModelsV3, error) {
	b, err := h.get("/3/Models", nil)
	if err != nil {
		return nil, fmt.Errorf("Error from get in GetModels:\n%v", err)
	}

	var m bindings.ModelsV3
	if err := unmarshal(b, &m); err != nil {
		return nil, fmt.Errorf("Error from unmarshal in GetModles\n%s", err)
	}

	return &m, nil
}

func (h *H2O) GetModel(modelName string) (*bindings.ModelsV3, error) {
	b, err := h.get("/3/Models", url.Values{"model_id": {modelName}})
	if err != nil {
		return nil, fmt.Errorf("Error from get in GetModels:\n%v", err)
	}

	var m bindings.ModelsV3
	if err := unmarshal(b, &m); err != nil {
		return nil, fmt.Errorf("Error from get in GetModels:\n%v", err)
	}

	return &m, nil
}

func (h *H2O) GetJobs() (*bindings.JobsV3, error) {
	b, err := h.get("/3/Jobs", nil)
	if err != nil {
		return nil, fmt.Errorf("Error getting jobs list: \n%v", err)
	}

	var j bindings.JobsV3
	if err := unmarshal(b, &j); err != nil {
		return nil, fmt.Errorf("Error unmarshalling jobs list: \n%v", err)
	}

	return &j, nil
}

func (h *H2O) GetJob(jobName string) (*bindings.JobsV3, error) {
	b, err := h.get("/3/Jobs", url.Values{"job_id": {jobName}})
	if err != nil {
		return nil, fmt.Errorf("Error getting job %s: \n%v", jobName, err)
	}

	var j bindings.JobsV3
	if err := unmarshal(b, &j); err != nil {
		return nil, fmt.Errorf("Error unmarshalling job %s: \n%v", jobName, err)
	}

	return &j, nil
}
