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

	"github.com/h2oai/steamY/lib/fs"
)

const (
	ArtifactWar       = "war"
	ArtifactJar       = "jar"
	ArtifactPythonWar = "pywar"
)

const (
	fileTypeJava        = "pojo"
	fileTypeJavaDep     = "jar"
	fileTypePythonMain  = "python"
	fileTypePythonOther = "pythonextra"
)

func CompileModel(address, wd string, projectId, modelId int64, modelLogicalName, artifact, packageName string) (string, error) {

	genModelPath := fs.GetGenModelPath(wd, modelId)
	javaModelPath := fs.GetJavaModelPath(wd, modelId, modelLogicalName)

	var targetFile, slug string

	switch artifact {
	case ArtifactWar:
		targetFile = fs.GetWarFilePath(wd, modelId, modelLogicalName)
		slug = "makewar"
	case ArtifactPythonWar:
		targetFile = fs.GetPythonWarFilePath(wd, modelId, modelLogicalName)
		slug = "makepythonwar"
	case ArtifactJar:
		targetFile = fs.GetModelJarFilePath(wd, modelId, modelLogicalName)
		slug = "callCompiler"
	}

	if _, err := os.Stat(targetFile); os.IsNotExist(err) {
	} else {
		return targetFile, nil
	}

	// ping to check if service is up
	if _, err := http.Get(toUrl(address, "ping")); err != nil {
		return "", fmt.Errorf("Failed connecting to scoring service builder: %s", err)
	}

	packageName = strings.TrimSpace(packageName)

	var pythonMainFilePath string
	var pythonOtherFilePaths []string

	if artifact == ArtifactPythonWar && len(packageName) > 0 {
		var err error
		pythonMainFilePath, pythonOtherFilePaths, err = getPythonFilePaths(wd, projectId, packageName)
		if err != nil {
			return "", err
		}
	}

	if err := callCompiler(toUrl(address, slug), targetFile, javaModelPath, genModelPath, pythonMainFilePath, pythonOtherFilePaths); err != nil {
		return "", err
	}

	return targetFile, nil
}

func callCompiler(url, targetFile, javaFilePath, javaDepPath, pythonMainFilePath string, pythonOtherFilePaths []string) error {
	b := &bytes.Buffer{}
	writer := multipart.NewWriter(b)

	if err := attachFile(writer, javaFilePath, fileTypeJava); err != nil {
		return fmt.Errorf("Failed attaching Java file to compilation request: %s", err)
	}

	if err := attachFile(writer, javaDepPath, fileTypeJavaDep); err != nil {
		return fmt.Errorf("Failed attaching Java dependency to compilation request: %s", err)
	}

	if len(pythonMainFilePath) > 0 {
		if err := attachFile(writer, pythonMainFilePath, fileTypePythonMain); err != nil {
			return fmt.Errorf("Failed attaching Python main file to compilation request: %s", err)
		}
		if pythonOtherFilePaths != nil && len(pythonOtherFilePaths) > 0 {
			for _, p := range pythonOtherFilePaths {
				if err := attachFile(writer, p, fileTypePythonOther); err != nil {
					return fmt.Errorf("Failed attaching Python file to compilation request: %s", err)
				}
			}
		}
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

	dst, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Failed creating compiled artifact %s: %v", targetFile, err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, res.Body); err != nil {
		return fmt.Errorf("Failed writing compiled artifact %s: %v", targetFile, err)
	}

	return nil
}

func getPythonFilePaths(wd string, projectId int64, packageName string) (string, []string, error) {
	var pythonMainFilePath string
	var pythonOtherFilePaths []string

	packagePath := fs.GetPackagePath(wd, projectId, packageName)

	if !fs.DirExists(packagePath) {
		return "", nil, fmt.Errorf("Package %s does not exist")
	}

	packageAttrsBytes, err := fs.GetPackageAttributes(wd, projectId, packageName)
	if err != nil {
		return "", nil, fmt.Errorf("Failed reading package attributes: %s", err)
	}

	packageAttrs, err := fs.JsonToMap(packageAttrsBytes)
	if err != nil {
		return "", nil, fmt.Errorf("Failed parsing package attributes: %s", err)
	}

	pythonMain, ok := packageAttrs["main"]
	if !ok {
		return "", nil, fmt.Errorf("Failed determining Python main file from package attributes")
	}

	packageFileList, err := fs.ListFiles(packagePath)
	if err != nil {
		return "", nil, fmt.Errorf("Failed reading package file list: %s", err)
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
		return "", nil, fmt.Errorf("Failed locating Python main file in package file listing")
	}

	return pythonMainFilePath, pythonOtherFilePaths, nil
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
