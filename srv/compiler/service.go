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

package compiler

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/h2oai/steam/lib/fs"
	"github.com/pkg/errors"
)

const (
	ArtifactWar       = "war"
	ArtifactJar       = "jar"
	ArtifactPythonWar = "pywar"
)

type fileType string

const (
	fileTypeJava        = "pojo"
	fileTypeJavaDep     = "jar"
	fileTypeMOJO        = "mojo"
	fileTypePythonMain  = "python"
	fileTypePythonOther = "pythonextra"
)

type ModelAsset interface {
	AttachFiles(w *multipart.Writer) error
}

func CompileModel(address, wd string, projectId, modelId int64, logicalName, modelType, algorithm, artifact, packageName string) (string, error) {
	// Verify that model has assets set
	switch modelType {
	case "mojo", "pojo":
	case "":
		return "", errors.New("model type unset")
	default:
		return "", errors.New(fmt.Sprintf("invalid model type %q", modelType))
	}

	// Check if target file exist
	targetFile, slug, ok := getTargetFile(artifact, wd, modelId, logicalName)
	if ok {
		return targetFile, nil
	}

	var pythonFilePaths []string
	if artifact == ArtifactPythonWar {
		var err error
		pythonFilePaths, err = getPythonFilePaths(wd, packageName, projectId)
		if err != nil {
			return "", errors.Wrap(err, "getting Python file paths")
		}
	}

	var assets ModelAsset
	if algorithm == "Deep Water" {
		assets = NewDeepwater(wd, modelId, logicalName, modelType, pythonFilePaths...)
	} else {
		assets = NewModel(wd, modelId, logicalName, modelType, pythonFilePaths...)
	}

	// ping to check if service is up
	if _, err := http.Get(toUrl(address, "ping")); err != nil {
		return "", errors.Wrap(err, "could not connect to prediction service builder")
	}

	if err := callCompiler(toUrl(address, slug), targetFile, assets); err != nil {
		return "", errors.Wrap(err, "failed compiler request")
	}

	return targetFile, nil
}

func getTargetFile(artifact, workingDirectory string, modelId int64, logicalName string) (string, string, bool) {
	var targetFile, slug string

	switch artifact {
	case ArtifactWar:
		targetFile = fs.GetWarFilePath(workingDirectory, modelId, logicalName)
		slug = "makewar"
	case ArtifactPythonWar:
		targetFile = fs.GetPythonWarFilePath(workingDirectory, modelId, logicalName)
		slug = "makepythonwar"
	case ArtifactJar:
		targetFile = fs.GetModelJarFilePath(workingDirectory, modelId, logicalName)
		slug = "compile"
	}

	_, err := os.Stat(targetFile)
	return targetFile, slug, os.IsExist(err)
}

func callCompiler(url, targetFile string, model ModelAsset) error {
	b := &bytes.Buffer{}
	writer := multipart.NewWriter(b)

	if err := model.AttachFiles(writer); err != nil {
		return err
	}

	ct := writer.FormDataContentType()
	writer.Close()

	res, err := http.Post(url, ct, b)
	if err != nil {
		return fmt.Errorf("Failed making compilation request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Failed reading compilation response: %v", err)
		}
		return fmt.Errorf("Failed compiling scoring service: %s / %s", res.Status, string(body))
	}

	dst, err := os.Create(targetFile)
	if err != nil {
		return fmt.Errorf("Failed creating compiled artifact %s: %v", targetFile, err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, res.Body); err != nil {
		return fmt.Errorf("Failed writing compiled artifact %s: %v", targetFile, err)
	}

	return nil
}

func getPythonFilePaths(workingDirectory, packageName string, projectId int64) ([]string, error) {
	packageName = strings.TrimSpace(packageName)
	if len(packageName) < 1 {
		return nil, errors.New("package not set for PythonWar")
	}

	var pythonMainFilePath string
	var pythonOtherFilePaths []string

	packagePath := fs.GetPackagePath(workingDirectory, projectId, packageName)

	if !fs.DirExists(packagePath) {
		return nil, fmt.Errorf("Package %s does not exist")
	}

	packageAttrsBytes, err := fs.GetPackageAttributes(workingDirectory, projectId, packageName)
	if err != nil {
		return nil, fmt.Errorf("Failed reading package attributes: %s", err)
	}

	packageAttrs, err := fs.JsonToMap(packageAttrsBytes)
	if err != nil {
		return nil, fmt.Errorf("Failed parsing package attributes: %s", err)
	}

	pythonMain, ok := packageAttrs["main"]
	if !ok {
		return nil, fmt.Errorf("Failed determining Python main file from package attributes")
	}

	packageFileList, err := fs.ListFiles(packagePath)
	if err != nil {
		return nil, fmt.Errorf("Failed reading package file list: %s", err)
	}

	// Filter .py files; separate ancillary files from the main one.
	pythonOtherFilePaths = make([]string, 0)
	for _, f := range packageFileList {
		if strings.ToLower(path.Ext(f)) == ".py" {
			p := path.Join(packagePath, f)
			if f == pythonMain {
				pythonMainFilePath = p
			} else {
				pythonOtherFilePaths = append(pythonOtherFilePaths, p)
			}
		}
	}

	if len(pythonMainFilePath) == 0 {
		return nil, fmt.Errorf("Failed locating Python main file in package file listing")
	}

	return append([]string{pythonMainFilePath}, pythonOtherFilePaths...), nil
}

func toUrl(address, slug string) string {
	return (&url.URL{Scheme: "http", Host: address, Path: slug}).String()
}

func attachFile(w *multipart.Writer, filePath, fileType string) error {
	dst, err := w.CreateFormFile(fileType, path.Base(filePath))
	if err != nil {
		return fmt.Errorf("Failed creating form attachment: %s", err)
	}
	src, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Failed opening file for attachment: %s", err)
	}
	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("Failed attaching file: %s", err)
	}

	return nil
}
