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

package master

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/h2oai/steam/lib/fs"
	"github.com/h2oai/steam/master/az"
	"github.com/h2oai/steam/srv/compiler"
	srvweb "github.com/h2oai/steam/srv/web"
	"github.com/rs/xid"
)

const (
	paramType        = "type"
	paramTypeModel   = "model"
	paramArtifact    = "artifact"
	paramProjectId   = "project-id"
	paramLabelName   = "label-name"
	paramModelId     = "model-id"
	paramPackageName = "package-name"

	// model artifact types
	javaClass    = "java-class"     // foo.java
	javaClassDep = "java-class-dep" // gen-model.jar
	javaJar      = "java-jar"       // foo.jar
	javaWar      = "java-war"       // foo.war
	javaPyWar    = "java-py-war"    // foo_py.war
)

type DownloadHandler struct {
	az                     az.Az
	workingDirectory       string
	webService             srvweb.Service
	compilerServiceAddress string
}

func newDownloadHandler(az az.Az, workingDirectory string, webService srvweb.Service, compilerServiceAddress string) *DownloadHandler {
	return &DownloadHandler{
		az,
		workingDirectory,
		webService,
		compilerServiceAddress,
	}
}

func (s *DownloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	guid := xid.New().String()

	pz, azerr := s.az.Identify(r)
	if azerr != nil {
		log.Println(guid, "ERR", "?", "Download", azerr)
		http.Error(w, fmt.Sprintf("Authentication failed: %s", azerr), http.StatusUnauthorized)
	}

	values := r.URL.Query()

	typ := values.Get(paramType)

	log.Println(guid, "REQ", pz, "Download", values)

	if len(typ) == 0 {
		http.Error(w, fmt.Sprintf("Missing %s", paramType), http.StatusBadRequest)
		return
	}

	switch typ {
	case paramTypeModel:
		artifact := values.Get(paramArtifact)

		if len(artifact) == 0 {
			http.Error(w, fmt.Sprintf("Missing %s", paramArtifact), http.StatusBadRequest)
			return
		}

		packageName := values.Get(paramPackageName)

		projectIdValue := values.Get(paramProjectId)
		if len(projectIdValue) == 0 {
			http.Error(w, fmt.Sprintf("Missing %s", paramProjectId), http.StatusBadRequest)
			return
		}

		projectId, err := strconv.ParseInt(projectIdValue, 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("Not a serial number %s=%s: %s", paramProjectId, projectIdValue, err), http.StatusBadRequest)
			return
		}

		if projectId <= 0 {
			http.Error(w, fmt.Sprintf("Invalid %s: %s", paramProjectId, projectId), http.StatusBadRequest)
			return
		}

		modelIdValue := values.Get(paramModelId)
		if len(modelIdValue) != 0 {
			modelId, err := strconv.ParseInt(modelIdValue, 10, 64)
			if err != nil {
				http.Error(w, fmt.Sprintf("Not a serial number %s=%s: %s", paramModelId, modelIdValue, err), http.StatusBadRequest)
				return
			}

			if modelId <= 0 {
				http.Error(w, fmt.Sprintf("Invalid %s: %s", paramModelId, modelId), http.StatusBadRequest)
				return
			}

			s.serveModel(w, r, pz, projectId, modelId, artifact, packageName)
			return
		}

		labelName := values.Get(paramLabelName)
		if len(labelName) != 0 {
			labels, err := s.webService.GetLabelsForProject(pz, projectId)
			if err != nil {
				http.Error(w, fmt.Sprintf("Failed reading labels for project %d: %s", projectId, err), http.StatusInternalServerError)
				return
			}

			for _, label := range labels {
				if label.Name == labelName {
					if label.ModelId > 0 {
						s.serveModel(w, r, pz, label.ProjectId, label.ModelId, artifact, packageName)
						return
					}
					http.Error(w, fmt.Sprintf("No model associated with label: %s", labelName), http.StatusNotFound)
					return
				}
			}

			http.Error(w, fmt.Sprintf("No label found: %s", labelName), http.StatusNotFound)
			return

		}
	default:
		http.Error(w, fmt.Sprintf("Invalid %s: %s", paramType, typ), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	http.Error(w, "Not Found", http.StatusNotFound)
}

func (s *DownloadHandler) serveModel(w http.ResponseWriter, r *http.Request, pz az.Principal, projectId, modelId int64, artifact string, packageName string) {
	switch artifact {
	case javaClass, javaClassDep, javaJar, javaWar, javaPyWar:
		// Call the API to get the model details.
		// We assume that if the GetModel() call succeeds, the principal has
		//   permissions and privileges to read this model, and consequently
		//   allowed to download it.
		model, err := s.webService.GetModel(pz, modelId)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed reading model %d: %s", modelId, err), http.StatusForbidden)
			return
		}
		modelLocation := model.Location
		if len(modelLocation) == 0 {
			http.Error(w, fmt.Sprintf("Failed reading model %d: the model was not saved correctly", model.Id), http.StatusNotFound)
			return
		}

		var filePath string
		switch artifact {
		case javaClass:
			filePath = fs.GetJavaModelPath(s.workingDirectory, modelId, model.LogicalName)

		case javaClassDep:
			filePath = fs.GetGenModelPath(s.workingDirectory, modelId)

		case javaWar:
			warFilePath, err := compiler.CompileModel(
				s.compilerServiceAddress,
				s.workingDirectory,
				projectId,
				modelId,
				model.LogicalName,
				compiler.ArtifactWar,
				"",
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			filePath = warFilePath

		case javaPyWar:
			packageName = strings.TrimSpace(packageName)
			if len(packageName) == 0 {
				http.Error(w, "No package-name specified", http.StatusBadRequest)
				return
			}
			warFilePath, err := compiler.CompileModel(
				s.compilerServiceAddress,
				s.workingDirectory,
				projectId,
				modelId,
				model.LogicalName,
				compiler.ArtifactPythonWar,
				packageName,
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			filePath = warFilePath

		case javaJar:
			jarFilePath, err := compiler.CompileModel(
				s.compilerServiceAddress,
				s.workingDirectory,
				projectId,
				modelId,
				model.LogicalName,
				compiler.ArtifactJar,
				"",
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			filePath = jarFilePath
		}

		// Delegate to builtin.
		// Can result in 200, 404, 403 or 500 based on file availability and permissions.
		w.Header().Set("Content-Disposition", "attachment; filename=\""+path.Base(filePath)+"\"")
		http.ServeFile(w, r, filePath)
		return

	default:
		http.Error(w, fmt.Sprintf("Invalid %s: %s", paramArtifact, artifact), http.StatusBadRequest)
		return
	}
}
