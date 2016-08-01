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

	"github.com/h2oai/steamY/lib/fs"
)

type Service struct {
	Address string
}

func NewService(address string) *Service {
	return &Service{
		address,
	}
}

func (s *Service) urlFor(slug string) string {
	return (&url.URL{Scheme: "http", Host: s.Address, Path: slug}).String()
}

func (s *Service) Ping() error {
	if _, err := http.Get(s.urlFor("Ping")); err != nil {
		return err
	}
	return nil
}

func uploadFile(filePath, kind string, w *multipart.Writer) error {
	dst, err := w.CreateFormFile(kind, path.Base(filePath))
	if err != nil {
		return fmt.Errorf("Failed writing to buffer: %v", err)
	}
	src, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Failed opening file: %v", err)
	}
	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("Failed copying file: %v", err)
	}

	return nil
}

func compile(url, pojoPath, jarPath string) (*http.Response, error) {
	pojoPath, err := fs.ResolvePath(pojoPath)
	if err != nil {
		return nil, err
	}
	jarPath, err = fs.ResolvePath(jarPath)
	if err != nil {
		return nil, err
	}

	b := &bytes.Buffer{}
	writer := multipart.NewWriter(b)

	if err := uploadFile(pojoPath, "pojo", writer); err != nil {
		return nil, err
	}
	if err := uploadFile(jarPath, "jar", writer); err != nil {
		return nil, err
	}

	ct := writer.FormDataContentType()
	writer.Close()

	res, err := http.Post(url, ct, b)
	if err != nil {
		return nil, fmt.Errorf("Failed uploading file: %v", err)
	}

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("Failed reading upload response: %v", err)
		}
		return nil, fmt.Errorf("Failed uploading file: %s / %s", res.Status, string(body))
	}

	return res, nil
}

func (s *Service) CompileModel(wd, modelLocation, modelLogicalName, artifact string) (string, error) {

	genModelPath := fs.GetGenModelPath(wd, modelLocation)
	javaModelPath := fs.GetJavaModelPath(wd, modelLocation, modelLogicalName)

	var targetFile string

	switch artifact {
	case "war":
		targetFile = fs.GetWarFilePath(wd, modelLocation, modelLogicalName)
	case "jar":
		targetFile = fs.GetModelJarFilePath(wd, modelLocation, modelLogicalName)
	}

	if _, err := os.Stat(targetFile); os.IsNotExist(err) {
	} else {
		return targetFile, nil
	}

	if err := s.Ping(); err != nil {
		return "", err
	}

	slug := "makewar"
	if artifact == "jar" {
		slug = "compile"
	}

	res, err := compile(s.urlFor(slug), javaModelPath, genModelPath)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	dst, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", fmt.Errorf("Download file createion failed: %s: %v", targetFile, err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, res.Body); err != nil {
		return "", fmt.Errorf("Download file copy failed: Service to %s: %v", targetFile, err)
	}
	return targetFile, nil
}
