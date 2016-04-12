package h2ov3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/h2oai/steamY/lib/fs"
	"github.com/h2oai/steamY/srv/web"
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

func (h *H2O) BuildAutoML(dataset, targetName string) error {
	_, err := h.get("/3/AutoMLBuilder", url.Values{ // FIXME: This needs first value needs to be stored
		"dataset":     {dataset},
		"target_name": {targetName},
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *H2O) GetModels() ([]*web.CloudModelSynopsis, error) {
	b, err := h.get("/3/Models", nil)
	if err != nil {
		return nil, err
	}
	var r GetModelsResponseV3
	if err := unmarshal(b, &r); err != nil {
		return nil, err
	}

	models := make([]*web.CloudModelSynopsis, len(r.Models))
	for i, s := range r.Models {
		models[i] = &web.CloudModelSynopsis{
			s.Algo,
			s.AlgoFullName,
			s.DataFrame.Name,
			s.ModelId.Name,
			s.ResponseColumnName,
			web.Timestamp(s.Timestamp / 1000),
		}
	}

	return models, nil
}

func (h *H2O) GetModel(modelID string) (*RawModel, error) {
	b, err := h.get("/3/Models/"+url.QueryEscape(modelID), nil)
	if err != nil {
		return nil, err
	}
	return &RawModel{modelID, b}, nil
}

func (h *H2O) ExportBinaryModel(modelID, p string) error {
	_, err := h.get("/99/Models.bin/"+url.QueryEscape(modelID), url.Values{
		"dir":   {p},
		"force": {"true"},
	})
	if err != nil {
		return fmt.Errorf("Binary model export failed: %s", err)
	}
	return nil
}

func (h *H2O) ImportBinaryModel(modelName, p string) error {
	_, err := h.post("/99/Models.bin/"+url.QueryEscape(modelName), url.Values{
		"dir":   {p},
		"force": {"true"},
	})

	if err != nil {
		return fmt.Errorf("Binary model import failed: %s", err)
	}
	return nil
}

func (h *H2O) ExportJavaModel(modelID, p string) error {
	_, err := h.download("/3/Models.java/"+url.QueryEscape(modelID), p, true)
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
