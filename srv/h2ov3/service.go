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
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/h2oai/steam/bindings"
	"github.com/h2oai/steam/lib/fs"
)

type H2O struct {
	Address string
}

func NewClient(address string) *H2O {
	return &H2O{
		address,
	}
}

func (h *H2O) url(path string, parms ...interface{}) string {
	parm := regexp.MustCompile(`/(\?{.+?})`)

	mats := parm.FindAllStringSubmatch(path, -1)
	if len(mats) != len(parms) {
		panic(fmt.Sprintf("length of parms doesn't match: %s: %v", path, parms))
	}

	for i, parm := range parms {
		mat := mats[i][1]

		switch parm.(type) {
		case string:
			path = strings.Replace(path, mat, parm.(string), 1)
		case int:
			path = strings.Replace(path, mat, strconv.Itoa(parm.(int)), 1)
		case bool:
			path = strings.Replace(path, mat, strconv.FormatBool(parm.(bool)), 1)
		default:
			panic(fmt.Sprintf("unexepcted type %T: %v", parm, parm))

		}
	}

	return (&url.URL{Scheme: "http", Host: h.Address, Path: path}).String()
}

type H2OException struct {
	Message string `json:"msg"`
	// Stacktrace []string `json:"stacktrace"`
}

func readException(b []byte) (*H2OException, error) {
	var exception H2OException
	if err := json.Unmarshal(b, &exception); err != nil {
		return nil, fmt.Errorf("H2O response unmarshal failed: %v", err)
	}
	return &exception, nil
}

func (h *H2O) handleResponse(res *http.Response, u string) ([]byte, error) {
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

func (h *H2O) download(s, p string, preserveFilename bool) (string, error) {
	u := h.url(s)

	_, filepath, err := fs.Download(p, u, preserveFilename)
	if err != nil {
		return "", err
	}

	return filepath, nil
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

func (h *H2O) JobPoll(jobId string) (*bindings.JobV3, error) {
	for {
		jobBase, err := h.GetJobsFetch(jobId)
		if err != nil {
			return nil, err
		}
		fmt.Println(jobBase)

		job := jobBase.Jobs[0]
		if job.Status == "CREATED" || job.Status == "RUNNING" {
			continue
		} else {
			if job.Exception != "" {
				return nil, fmt.Errorf("H2O Job Error: %v", job.Exception)
			}
		}
		return job, nil
	}
}

// AutoML will use the built in AutoML tool to create a model with minimal
// user interaction.
func (h *H2O) AutoML(dataset, targetName string, maxTime int) (string, error) {
	// b, err := h.post("/3/AutoMLBuilder", url.Values{
	// 	"dataset":     {dataset},
	// 	"target_name": {targetName},
	// 	"max_time":    {strconv.Itoa(maxTime)},
	// })
	// if err != nil {
	// 	return "", err
	// }
	// var a AutoMLBuilderV3
	// if err := unmarshal(b, &a); err != nil {
	// 	return "", err
	// }

	// k, err := h.jobPoll(a.Job.Key.Name)

	// b, err = h.get("/3/AutoML/"+k.Dest.Name, nil)
	// if err != nil {
	// 	return "", err
	// }
	// var m AutoMLV3
	// if err = unmarshal(b, &m); err != nil {
	// 	return "", err
	// }

	// if m.Leader.Name == "" {
	// 	return "", fmt.Errorf("Unable to complete model in %d seconds", maxTime)
	// }
	// log.Printf("Completed AutoML with model: %s", m.Leader.Name)
	// return m.Leader.Name, err
	return "", nil
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

// func (h *H2O) CompilePojo(javaModel, jar string) error {
// 	if _, err := h.post("/compile", url.Values{
// 		"pojo": {javaModel},
// 		"jar":  {jar},
// 	}); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (h *H2O) Shutdown() error {
// 	_, err := h.post("/3/Shutdown", nil)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
