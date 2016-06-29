package rpc

import (
	"bytes"
	"fmt"
	grpc "github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Service struct {
	namespace string
	handler   interface{}
}

func NewService(namespace string, handler interface{}) *Service {
	return &Service{namespace, handler}
}

type Proc struct {
	Address   string
	username  string
	password  string
	client    *http.Client
	url       string
	namespace string
}

func NewProc(scheme, path, namespace, address, username, password string) *Proc {
	u := url.URL{Scheme: scheme, Host: address, Path: path}
	return &Proc{
		address,
		username,
		password,
		&http.Client{},
		u.String(),
		namespace + ".",
	}
}

func (proc *Proc) Call(method string, in, out interface{}) error {
	buf, err := json.EncodeClientRequest(proc.namespace+method, in)
	if err != nil {
		return fmt.Errorf("Error encoding request: %v", err)
	}

	body := bytes.NewBuffer(buf)
	req, err := http.NewRequest("POST", proc.url, body)
	if err != nil {
		return fmt.Errorf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(proc.username, proc.password)

	res, err := proc.client.Do(req)
	if err != nil {
		return fmt.Errorf("Error making request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		b, err := ioutil.ReadAll(res.Body)
		if err == nil {
			fmt.Println(string(b))
		}
		return fmt.Errorf("Error making request: %s", res.Status)
	}

	if err = json.DecodeClientResponse(res.Body, &out); err != nil {
		return err
	}
	return nil
}

func NewServer(svc *Service) http.Handler {
	s := grpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(svc.handler, svc.namespace)
	return s
}
