package cli2

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/h2oai/steam/lib/fs"
)

func transmitFile(url, username, password, filename string, attrs map[string]string) error {
	filename, err := fs.ResolvePath(filename)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	for key, value := range attrs {
		formField, err := writer.CreateFormField(key)
		if err != nil {
			return fmt.Errorf("Failed creating form field %s: %v", key, err)
		}

		if _, err := formField.Write([]byte(value)); err != nil {
			return fmt.Errorf("Failed writing form field %s: %v", key, err)
		}
	}

	dst, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return fmt.Errorf("Failed writing to buffer: %v", err)
	}

	src, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Failed opening file: %v", err)
	}

	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("Failed copying file: %v", err)
	}

	ct := writer.FormDataContentType()
	writer.Close()

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return fmt.Errorf("Error creating request: %v", err)
	}
	req.Header.Set("Content-type", ct)
	req.SetBasicAuth(username, password)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Failed uploading file: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Failed reading upload response: %v", err)
		}
		return fmt.Errorf("Failed uploading file: %s / %s", res.Status, string(body))
	}

	return nil
}
