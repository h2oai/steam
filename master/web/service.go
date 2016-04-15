package web

import (
	"fmt"
	"log"
	"time"

	"github.com/h2oai/steam/lib/fs"
	"github.com/h2oai/steamY/lib/yarn"
	"github.com/h2oai/steamY/master/db"
	"github.com/h2oai/steamY/srv/h2ov3"
	"github.com/h2oai/steamY/srv/web"
)

type Service struct {
	workingDir                string
	ds                        *db.DS
	compilationServiceAddress string
	kerberosEnabled           bool
	username                  string
	keytab                    string
}

func toTimestamp(t time.Time) web.Timestamp {
	return web.Timestamp(t.UTC().Unix())
}

func now() web.Timestamp {
	return toTimestamp(time.Now())
}

func NewService(workingDir string, ds *db.DS, compilationServiceAddress string, kerberos bool, username, keytab string) *web.Impl {
	return &web.Impl{&Service{workingDir, ds, compilationServiceAddress, kerberos, username, keytab}}
}

func (s *Service) Ping(status bool) (bool, error) {
	return status, nil
}

func (s *Service) StartCloud(cloudName, engineName string, size int, memory, username string) (*web.Cloud, error) { // TODO: YARN DRIVER SHOULD BE THE ENGINE
	appId, address, err := yarn.StartCloud(size, s.kerberosEnabled, memory, cloudName, s.username, s.keytab) // FIXME: THIS IS USING ADMIN TO START ALL CLOUDS
	if err != nil {
		return nil, err
	}

	c := db.NewCloud(
		cloudName,
		engineName,
		size,
		appId,
		address,
		memory,
		username,
		string(web.CloudStarted),
	)

	log.Printf("Created Cloud:\n%+v", c)

	if err := s.ds.CreateCloud(c); err != nil {
		return nil, err
	}

	return toCloud(c), nil
}

func (s *Service) StopCloud(cloudName string) error {

	// Make sure the cloud exists
	c, err := s.getCloud(cloudName)
	if err != nil {
		return err
	}

	// Bail out if already stopped
	if c.State == string(web.CloudStopped) {
		return fmt.Errorf("Cloud %s is already stopped", cloudName)
	}

	if err := yarn.StopCloud(s.kerberosEnabled, c.ID, c.ApplicationID, s.username, s.keytab); err != nil { //FIXME: this is using adming kerberos credentials
		return err
	}

	// Update the state and update DB
	c.State = string(web.CloudStopped)
	if err := s.ds.UpdateCloud(c); err != nil {
		return err
	}

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

func (s *Service) GetCloud(cloudName string) (*web.Cloud, error) {
	c, err := s.getCloud(cloudName)
	if err != nil {
		return nil, err
	}
	return toCloud(c), nil
}

func (s *Service) getCloud(cloudName string) (*db.Cloud, error) {
	c, err := s.ds.ReadCloud(cloudName)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, fmt.Errorf("Cloud %s does not exist.", cloudName)
	}
	return c, nil

}

func (s *Service) GetClouds() ([]*web.Cloud, error) {
	cs, err := s.ds.ListCloud()
	if err != nil {
		return nil, err
	}

	clouds := make([]*web.Cloud, len(cs))
	for i, c := range cs {
		clouds[i] = toCloud(c)
	}
	return clouds, nil
}

func (s *Service) DeleteCloud(cloudName string) error {
	return s.ds.DeleteCloud(cloudName)
}

func (s *Service) BuildModel(cloudName string, dataset string, targetName string, maxRunTime int) (*web.Model, error) {
	return nil, nil //FIXME
}

func (s *Service) GetModel(modelName string) (*web.Model, error) {
	return nil, nil //FIXME
}

func (s *Service) GetModels() ([]*web.Model, error) {
	return nil, nil //FIXME
}

func (s *Service) DeleteModel(modelName string) error {
	return nil
}

func (s *Service) StartScoringService(modelName string, port int) (*web.ScoringService, error) {
	return nil, nil //FIXME
}

func (s *Service) StopScoringService(modelName string, port int) error {
	return nil //FIXME
}

func (s *Service) GetScoringService(serviceName string) (*web.ScoringService, error) {
	return nil, nil //FIXME
}

func (s *Service) GetScoringServices() ([]*web.ScoringService, error) {
	return nil, nil //FIXME
}

func (s *Service) DeleteScoringService(modelName string, port int) error {
	return nil //FIXME
}

func (s *Service) GetEngine(engineName string) (*web.Engine, error) {
	return nil, nil //FIXME
}

func (s *Service) GetEngines() ([]*web.Engine, error) {
	return nil, nil //FIXME
}

func (s *Service) DeleteEngine(engineName string) error {
	return nil //FIXME
}

//
// Routines to convert DB structs into API structs
//

func toCloud(c *db.Cloud) *web.Cloud {
	return &web.Cloud{
		c.ID,
		c.EngineName,
		c.Size,
		c.ApplicationID,
		c.Address,
		c.Memory,
		c.Username,
		web.CloudState(c.State),
		web.Timestamp(c.CreatedAt),
	}
}
