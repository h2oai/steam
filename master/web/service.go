package web

import (
	// "fmt"
	// "github.com/h2oai/steam/lib/az"
	// "github.com/h2oai/steam/lib/fs"
	// "github.com/h2oai/steam/lib/grid"
	// "github.com/h2oai/steam/lib/sec"
	// "github.com/h2oai/steam/master/cli"
	// "github.com/h2oai/steam/master/ctl"

	"github.com/h2oai/steamY/lib/yarn"
	"github.com/h2oai/steamY/master/db"
	"github.com/h2oai/steamY/srv/h2ov3"
	// "github.com/h2oai/steam/master/job"
	// "github.com/h2oai/steam/master/proc"
	// "github.com/h2oai/steam/master/usr/auth"
	// "github.com/h2oai/steam/srv/h2ov3"
	"github.com/h2oai/steamY/srv/web"
	// "os"
	// "regexp"
	// "strconv"
	// "strings"
	// "time"
)

type Service struct {
	wd string
	ds *db.DS
}

func NewService(wd string, ds *db.DS) *web.Impl {
	return &web.Impl{&Service{wd, ds}}
}

func (s *Service) Ping(status bool) (bool, error) {
	return status, nil
}

func (s *Service) StartCloud(size int, kerberos bool, name, username, keytab string) (string, error) {
	id, err := yarn.StartCloud(size, kerberos, name, username, keytab)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *Service) StopCloud(kerberos bool, name, id, username, keytab string) error {
	err := yarn.StopCloud(kerberos, name, id, username, keytab)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetCloud(address string) (*web.Cloud, error) {
	h := h2ov3.NewClient(address)

	c, err := h.GetCloud()
	if err != nil {
		return nil, err
	}
	var health web.CloudState
	if c.IsHealthy {
		health = "Healthy"
	} else {
		health = "Unhealthy"
	}
	cc := &web.Cloud{
		c.Name,
		c.Version,
		health,
	}
	return cc, nil
}

func (s *Service) BuildAutoML(address, dataset, targetName string) error {
	h := h2ov3.NewClient(address)

	err := h.BuildAutoML(dataset, targetName)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetModels(address string) ([]*web.CloudModelSynopsis, error) {
	h := h2ov3.NewClient(address)

	ms, err := h.GetModels()
	if err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *Service) GetModel(address string) (*RawModel, error) {
	h := h2ov3.NewClient(address)

	m, err := h.GetModels()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Service) CompilePojo(address, javaModel, jar string) error {
	// h := compileClient.NewClient()
	//
	// p, err := h.CompilePojo(javaModel, jar)
	// if err != nil {
	// 	  return err
	// }

	return nil
}

func (s *Service) Shutdown(address string) error {
	h := h2ov3.NewClient(address)

	if err := h.Shutdown(); err != nil {
		return err
	}

	return nil
}
