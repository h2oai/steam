package web

import (
	"github.com/h2oai/steam/lib/fs"
	"github.com/h2oai/steamY/lib/yarn"
	"github.com/h2oai/steamY/master/db"
	"github.com/h2oai/steamY/srv/h2ov3"
	"github.com/h2oai/steamY/srv/web"
	"time"
)

type Service struct {
	workingDir                string
	ds                        *db.DS
	compilationServiceAddress string
}

func Timestamp(t time.Time) web.Timestamp {
	return web.Timestamp(t.UTC().Unix())
}

func Now() web.Timestamp {
	return Timestamp(time.Now())
}

func NewService(workingDir string, ds *db.DS, compilationServiceAddress string) *web.Impl {
	return &web.Impl{&Service{workingDir, ds, compilationServiceAddress}}
}

func (s *Service) Ping(status bool) (bool, error) {
	return status, nil
}

func (s *Service) StartCloud(cloudName, engineName string, size int, memory string, useKerberos bool, username string) (*web.Cloud, error) {
	keytab := "" // FIXME
	appId, address, err := yarn.StartCloud(size, useKerberos, memory, cloudName, username, keytab)
	if err != nil {
		return nil, err
	}
	return &web.Cloud{
		cloudName,
		engineName,
		size,
		appId,
		address,
		memory,
		username,
		true,
		web.CloudStarted,
		Now(), // FIXME
	}, nil
}

func (s *Service) StopCloud(cloudName string) error {
	return nil
}

func (s *Service) StopCloud__FIXME(name string, useKerberos bool, applicationID, username, keytab string) error {
	// FIXME
	err := yarn.StopCloud(useKerberos, name, applicationID, username, keytab)
	if err != nil {
		return err
	}
	return nil
}

// func (s *Service) GetCloud(address string) (*web.Cloud, error) {
// 	h := h2ov3.NewClient(address)

// 	c, err := h.GetCloud()
// 	if err != nil {
// 		return nil, err
// 	}
// 	var health web.CloudState
// 	if c.IsHealthy {
// 		health = "Healthy"
// 	} else {
// 		health = "Unhealthy"
// 	}
// 	cc := &web.Cloud{
// 		c.Name,
// 		c.Version,
// 		health,
// 	}
// 	return cc, nil
// }

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
	// h := compileclient.newclient()

	// p, err := h.compilepojo(javamodel, jar)
	// if err != nil {
	// 	return err
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

func (s *Service) GetCloud(cloudName string) (*web.Cloud, error) {
	return nil, nil
}
func (s *Service) GetClouds() ([]*web.Cloud, error) {
	return nil, nil
}
func (s *Service) DeleteCloud(cloudName string) error {
	return nil
}
func (s *Service) BuildModel(cloudName string, dataset string, targetName string, maxRunTime int) (*web.Model, error) {
	return nil, nil
}
func (s *Service) GetModel(modelName string) (*web.Model, error) {
	return nil, nil
}
func (s *Service) GetModels() ([]*web.Model, error) {
	return nil, nil
}
func (s *Service) DeleteModel(modelName string) error {
	return nil
}
func (s *Service) StartScoringService(modelName string, port int) (*web.ScoringService, error) {
	return nil, nil
}
func (s *Service) StopScoringService(modelName string, port int) error {
	return nil
}
func (s *Service) GetScoringService(serviceName string) (*web.ScoringService, error) {
	return nil, nil
}
func (s *Service) GetScoringServices() ([]*web.ScoringService, error) {
	return nil, nil
}
func (s *Service) DeleteScoringService(modelName string, port int) error {
	return nil
}
func (s *Service) GetEngine(engineName string) (*web.Engine, error) {
	return nil, nil
}
func (s *Service) GetEngines() ([]*web.Engine, error) {
	return nil, nil
}
func (s *Service) DeleteEngine(engineName string) error {
	return nil
}
