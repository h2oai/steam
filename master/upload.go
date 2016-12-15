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
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/h2oai/steam/lib/fs"
	"github.com/h2oai/steam/master/az"
	"github.com/h2oai/steam/master/data"
	srvweb "github.com/h2oai/steam/srv/web"
	"github.com/pkg/errors"
)

type UploadHandler struct {
	az               az.Az
	workingDirectory string
	webService       srvweb.Service
	ds               *data.Datastore
}

func newUploadHandler(az az.Az, wd string, webService srvweb.Service, ds *data.Datastore) *UploadHandler {
	return &UploadHandler{az, wd, webService, ds}
}

func (s *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("File upload request received.")

	pz, azerr := s.az.Identify(r)
	if azerr != nil {
		log.Println(azerr)
		http.Error(w, fmt.Sprintf("Authentication failed: %s", azerr), http.StatusUnauthorized)
		return
	}

	r.ParseMultipartForm(0)

	src, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("Upload form parse failed:", err)
		http.Error(w, fmt.Sprintf("Malformed request: %v", err), http.StatusBadRequest)
		return
	}
	defer src.Close()

	typ := r.FormValue("type")
	var dstDir string

	switch typ {
	case fs.KindEngine:
		if err := pz.CheckPermission(s.ds.Permission.ManageEngine); err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		dstDir = path.Join(s.workingDirectory, fs.LibDir, typ)

		if path.Ext(handler.Filename) != ".zip" {
			http.Error(w, fmt.Sprintf("Invalid file type. Accepted filetype(s): zip"), http.StatusUnsupportedMediaType)
		}

	case fs.KindFile:
		if err := pz.CheckPermission(s.ds.Permission.ManageProject); err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		projectIdValue := r.FormValue("project-id")
		projectId, err := strconv.ParseInt(projectIdValue, 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid project id: %s", projectIdValue), http.StatusBadRequest)
			return
		}

		if err := pz.CheckEdit(s.ds.EntityType.Project, projectId); err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		packageName := r.FormValue("package-name")
		if err := fs.ValidateName(packageName); err != nil {
			http.Error(w, fmt.Sprintf("Invalid package name: %s", err), http.StatusBadRequest)
			return
		}

		packagePath := fs.GetPackagePath(s.workingDirectory, projectId, packageName)
		if !fs.DirExists(packagePath) {
			http.Error(w, fmt.Sprintf("Package %s does not exist", packageName), http.StatusBadRequest)
			return
		}

		relativePath := r.FormValue("relative-path")
		dstDir, err = fs.GetPackageRelativePath(s.workingDirectory, projectId, packageName, relativePath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid relative path: %s", err), http.StatusBadRequest)
		}

	default:
		http.Error(w, fmt.Sprintf("Invalid upload type: %s", typ), http.StatusBadRequest)
		return
	}

	log.Println("Remote file: ", handler.Filename)

	fileBaseName := path.Base(handler.Filename)
	dstPath := path.Join(dstDir, fileBaseName)

	if err := os.MkdirAll(path.Dir(dstPath), fs.DirPerm); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	dst, err := os.OpenFile(dstPath, os.O_CREATE|os.O_TRUNC, fs.FilePerm)
	if err != nil {
		log.Println("Upload file open operation failed:", err)
		http.Error(w, fmt.Sprintf("Error writing uploaded file to disk: %s", err), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	io.Copy(dst, src)

	switch typ {
	case fs.KindEngine:
		defer os.Remove(dstPath)

		if err := s.handleEngine(w, pz, fileBaseName, dstDir, dstPath); err != nil {
			log.Println("Failed saving engine to disk:", err)
			return
		}
	}
}

func (s *UploadHandler) handleEngine(w http.ResponseWriter, pz az.Principal, fileName, fileDir, filePath string) error {
	// Open zip file and defer close
	r, err := zip.OpenReader(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error opening zip file: %v", err), http.StatusInternalServerError)
		return errors.Wrap(err, "failed opening zip file")
	}
	defer r.Close()

	// Search zipFile for h2o.jar or h2odriver.jar
	var rc io.ReadCloser
	for i, f := range r.File {
		if path.Base(f.Name) == "h2o.jar" || path.Base(f.Name) == "h2odriver.jar" {
			var err error
			rc, err = f.Open()
			if err != nil {
				http.Error(w, fmt.Sprintf("Upload file open operation failed: %v", err), http.StatusInternalServerError)
				return errors.Wrap(err, "failed opening engine file")
			}
			defer rc.Close()
			break
		}

		if i == len(r.File)-1 {
			http.Error(w, "Unable to locate valid engine", http.StatusBadRequest)
			return fmt.Errorf("failed to locate engine")
		}
	}

	// Create a .jar file with engine version name
	dstBase := strings.Replace(fileName, path.Ext(fileName), ".jar", 1)
	dstPath := path.Join(fileDir, dstBase)
	dst, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, fs.FilePerm)
	if err != nil {
		http.Error(w, fmt.Sprintf("Upload file create operation failed: %v", err), http.StatusInternalServerError)
		return errors.Wrap(err, "failed creating engine file")
	}

	// Copy file from rc to dst
	if _, err := io.Copy(dst, rc); err != nil {
		http.Error(w, fmt.Sprintf("File copy operation failed:", err), http.StatusInternalServerError)
		return errors.Wrap(err, "failed copying uploaded file to disk")
	}

	// Add Engine to datastore
	if _, err := s.ds.CreateEngine(dstBase, dstPath,
		data.WithPrivilege(pz, data.Owns), data.WithAudit(pz),
	); err != nil {
		http.Error(w, fmt.Sprintf("Error saving engine to datastore: %v", err), http.StatusInternalServerError)
		return errors.Wrap(err, "failed saving engine to datastore")
	}

	return nil
}
