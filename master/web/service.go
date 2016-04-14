package web

import (
	// "fmt"
	// "github.com/h2oai/steam/lib/az"
	// "github.com/h2oai/steam/lib/fs"
	// "github.com/h2oai/steam/lib/grid"
	// "github.com/h2oai/steam/lib/sec"
	// "github.com/h2oai/steam/master/cli"
	// "github.com/h2oai/steam/master/ctl"

	"github.com/h2oai/steam/lib/fs"
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
	workingDir                string
	ds                        *db.DS
	compilationServiceAddress string
}

func NewService(workingDir string, ds *db.DS, compilationServiceAddress string) *web.Impl {
	return &web.Impl{&Service{workingDir, ds, compilationServiceAddress}}
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

func (s *Service) BuildAutoML(address, dataset, targetName string, maxTime int) (string, error) {
	h := h2ov3.NewClient(address)

	modelName, err := h.AutoML(dataset, targetName, maxTime) //TODO: j is a job that can be started and waited for
	if err != nil {
		return "", err
	}

	javaModelDir := fs.GetModelPath(s.workingDir, modelName, "java")
	if err := h.ExportJavaModel(modelName, javaModelDir); err != nil {
		return "", err
	}
	if err := h.ExportGenModel(javaModelDir); err != nil {
		return "", err
	}

	return modelName, nil
}

// func (s *Service) GetModels(address string) ([]*web.CloudModelSynopsis, error) {
// 	h := h2ov3.NewClient(address)

// 	ms, err := h.GetModels()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ms, nil
// }

// func (s *Service) GetModel(address string, modelID string) (*web.ModelInfo, error) {

// 	h := h2ov3.NewClient(address)

// 	m, err := h.GetModel(modelID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	log.Println()

// 	return nil, nil
// }

func (s *Service) DeployPojo(address, javaModel, jar string) error {
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

// func toWebModel(m *db.Model) *web.ModelInfo {
// 	return nil
// }
