package comp

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

type CompileServer struct {
	Address string
}

func NewServer(address string) *CompileServer {
	return &CompileServer{
		address,
	}
}

func (cs *CompileServer) url(p string) string {
	return (&url.URL{Scheme: "http", Host: cs.Address, Path: p}).String()
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

func uploadJavaFiles(url, pojoPath, jarPath, extraPath string) (*http.Response, error) {
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
	if err := uploadFile(extraPath, "extra", writer); err != nil {
		return nil, err
	}

	// pdst, err := writer.CreateFormFile("pojo", path.Base(pojoPath))
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed writing to buffer: %v", err)
	// }

	// psrc, err := os.Open(pojoPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed opening file: %v", err)
	// }

	// if _, err := io.Copy(pdst, psrc); err != nil {
	// 	return nil, fmt.Errorf("Failed copying file: %v", err)
	// }

	// jdst, err := writer.CreateFormFile("jar", path.Base(jarPath))
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed writing to buffer: %v", err)
	// }

	// jsrc, err := os.Open(jarPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed opening file: %v", err)
	// }

	// if _, err := io.Copy(jdst, jsrc); err != nil {
	// 	return nil, fmt.Errorf("Failed copying file: %v", err)
	// }

	// edst, err := writer.CreateFormFile("extra", path.Base(extraPath))
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed writing to buffer: %v", err)
	// }

	// esrc, err := os.Open(jarPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed opening file: %v", err)
	// }

	// if _, err := io.Copy(jdst, jsrc); err != nil {
	// 	return nil, fmt.Errorf("Failed copying file: %v", err)
	// }

	ct := writer.FormDataContentType()
	writer.Close()

	// log.Printf("Dumping ct: %v", ct)
	// log.Printf(" Dumping Buffer: %+v", b)

	res, err := http.Post(url, ct, b)
	if err != nil {
		return nil, fmt.Errorf("Failed uploading file: %v", err)
	}

	// req, err := http.NewRequest("POST", url, b)
	// if err != nil {
	// 	return fmt.Errorf("Error creating request: %v", err)
	// }
	// req.Header.Add("Content-Type", ct)

	// client := &http.Client{}
	// res, err := client.Do(req)
	// if err != nil {
	// 	return fmt.Errorf("Failed uploading file: %v", err)
	// }

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("Failed reading upload response: %v", err)
		}
		return nil, fmt.Errorf("Failed uploading file: %s / %s", res.Status, string(body))
	}

	return res, nil
}

func (cs *CompileServer) CompilePojo(modelPath, jarPath, extraPath, servlet string) (string, error) {
	res, err := uploadJavaFiles(cs.url(servlet), modelPath, jarPath, extraPath)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	outname := path.Base(modelPath)
	outname = outname[0 : len(outname)-5]

	var p string
	d := path.Dir(modelPath)
	if servlet == "compile" {
		p = path.Join(d, outname+".jar")
	} else if servlet == "makewar" {
		p = path.Join(d, outname+".war")
	}

	p, err = fs.ResolvePath(p)
	if err != nil {
		return "", err
	}

	dst, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", fmt.Errorf("Download file createion failed: %s: %v", p, err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, res.Body); err != nil {
		return "", fmt.Errorf("Download file copy failed: Service to %s: %v", p, err)
	}
	return p, nil
}
