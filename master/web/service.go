package web

import (
	"fmt"
	"log"
	"time"

	"github.com/h2oai/steamY/lib/fs"
	"github.com/h2oai/steamY/lib/svc"
	"github.com/h2oai/steamY/lib/yarn"
	"github.com/h2oai/steamY/master/db"
	"github.com/h2oai/steamY/srv/comp"
	"github.com/h2oai/steamY/srv/h2ov3"
	"github.com/h2oai/steamY/srv/web"
)

type Service struct {
	workingDir                string
	ds                        *db.DS
	compilationServiceAddress string
	scoringServiceAddress     string
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

func NewService(workingDir string, ds *db.DS, compilationServiceAddress, scoringServiceAddress string, kerberos bool, username, keytab string) *web.Impl {
	return &web.Impl{&Service{
		workingDir,
		ds,
		compilationServiceAddress,
		scoringServiceAddress,
		kerberos,
		username,
		keytab,
	}}
}

func (s *Service) Ping(status bool) (bool, error) {
	return status, nil
}

func (s *Service) StartCloud(cloudName, engineName string, size int, memory, username string) (*web.Cloud, error) { // TODO: YARN DRIVER SHOULD BE THE ENGINE
	// Make sure this cloud is unique
	if _, ok := s.getCloud(cloudName); ok == nil {
		return nil, fmt.Errorf("Cloud start failed. A cloud with the name %s already exists.", cloudName)
	}
	// Make cloud with yarn
	appId, address, err := yarn.StartCloud(size, s.kerberosEnabled, memory, cloudName, s.username, s.keytab) // FIXME: THIS IS USING ADMIN TO START ALL CLOUDS
	if err != nil {
		return nil, err
	}
	// Add cloud to db
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
	// Make sure cloud is not running and exists
	if c, err := s.getCloud(cloudName); c != nil {
		if c.State != web.CloudStopped {
			return fmt.Errorf("Cannot delete. Cloud %s is still running.",
				cloudName)
		}
	} else if err != nil {
		return fmt.Errorf("Cloud %s does not exist.", cloudName)
	}
	return s.ds.DeleteCloud(cloudName)
}

func (s *Service) BuildModel(cloudName string, dataset string, targetName string, maxRunTime int) (*web.Model, error) {
	// c, err := s.GetCloud(cloudName)
	// if err != nil {
	// 	return nil, err
	// }
	// h := h2ov3.NewClient(c.Address)

	h := h2ov3.NewClient("172.16.2.108:54321") //FIXME: THIS SHOULD BE CLOUD A ADDRESS

	modelName, err := h.AutoML(dataset, targetName, maxRunTime) // TODO: can be a goroutine
	if err != nil {
		return nil, err
	}

	javaModelDir := fs.GetModelPath(s.workingDir, modelName, "java")
	jm, err := h.ExportJavaModel(modelName, javaModelDir)
	if err != nil {
		return nil, err
	}
	gm, err := h.ExportGenModel(javaModelDir)
	if err != nil {
		return nil, err
	}

	m := db.NewModel(
		modelName,
		cloudName,
		dataset,
		targetName,
		maxRunTime,
		jm,
		gm,
	)

	if err := s.ds.CreateModel(m); err != nil {
		return nil, err
	}

	return toModel(m), nil
}

func (s *Service) getModel(modelName string) (*db.Model, error) {
	m, err := s.ds.ReadModel(modelName)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, fmt.Errorf("Model %s does not exist.", modelName)
	}
	return m, err
}

func (s *Service) GetModel(modelName string) (*web.Model, error) {
	m, err := s.getModel(modelName)
	if err != nil {
		return nil, err
	}
	return toModel(m), nil
}

func (s *Service) GetModels() ([]*web.Model, error) {
	ms, err := s.ds.ListModel()
	if err != nil {
		return nil, err
	}

	models := make([]*web.Model, len(ms))
	for i, m := range ms {
		models[i] = toModel(m)
	}

	return models, nil
}

func (s *Service) DeleteModel(modelName string) error {
	if ss, _ := s.getScoringService(modelName); ss != nil {
		if ss.State != web.ScoringServiceStopped {
			return fmt.Errorf("Cannot delete. A scoring service on model %s is"+
				" still running.", modelName)
		}
	} else if _, err := s.getModel(modelName); err != nil {
		return fmt.Errorf("Model %s does not exits.", modelName)
	}
	return s.ds.DeleteModel(modelName)
}

// Returns a Warfile for use in deployment
func (s *Service) compileModel(modelName string) (string, error) {
	c := comp.NewServer(s.compilationServiceAddress)

	m, err := s.getModel(modelName)
	if err != nil {
		return "", err
	}
	j := m.JavaModelPath
	g := m.GenModelPath

	e := fs.GetAssetsPath(s.workingDir, "makewar-extra.jar")

	w, err := c.CompilePojo(j, g, e, "makewar")
	if err != nil {
		return "", err
	}
	log.Printf("Warfile Dest: %s", w)

	return w, nil
}

func genScoringServiceName(modelName string) string { return modelName + " Scoring Service" }

func (s *Service) StartScoringService(modelName string, port int) (*web.ScoringService, error) {

	if _, ok := s.getScoringService(modelName); ok == nil {
		return nil, fmt.Errorf("A scoring service with the model %s already exists.", modelName)
	}

	w, err := s.compileModel(modelName)
	if err != nil {
		return nil, err
	}

	const jetty = "jetty-runner-9.3.9.M1.jar"

	j := fs.GetAssetsPath(s.workingDir, jetty)

	pid, err := svc.Start(w, j, s.scoringServiceAddress, port)
	if err != nil {
		return nil, err
	}

	ss := db.NewScoringService(
		genScoringServiceName(modelName),
		modelName,
		s.scoringServiceAddress,
		port,
		string(web.ScoringServiceStarted),
		pid,
	)

	if err := s.ds.CreateScoringService(ss); err != nil {
		return nil, err
	}

	return toScoringService(ss), nil
}

func (s *Service) StopScoringService(modelName string, port int) error {
	// Find the cloud in db
	ss, err := s.getScoringService(modelName)
	if err != nil {
		return err
	}
	// Verify Scoring Service is still running
	if ss.State == string(web.ScoringServiceStopped) {
		return fmt.Errorf("Scoring Service on %s is already stopped", modelName)
	}
	// Stop Scoring Service
	if err := svc.Stop(ss.Pid); err != nil {
		return err
	}
	// Update the state and update DB
	ss.State = string(web.ScoringServiceStopped)
	if err := s.ds.UpdateScoringService(ss); err != nil {
		return err
	}
	return nil
}

func (s *Service) getScoringService(modelName string) (*db.ScoringService, error) {
	ss, err := s.ds.ReadScoringService(genScoringServiceName(modelName))
	if err != nil {
		return nil, err
	}
	if ss == nil {
		return nil, fmt.Errorf("Scoring service for model %s does not exits.", modelName)
	}

	return ss, nil
}

func (s *Service) GetScoringService(modelName string) (*web.ScoringService, error) {
	ss, err := s.getScoringService(modelName)
	if err != nil {
		return nil, err
	}
	return toScoringService(ss), nil //FIXME
}

func (s *Service) GetScoringServices() ([]*web.ScoringService, error) {
	scs, err := s.ds.ListScoringService()
	if err != nil {
		return nil, err
	}
	ss := make([]*web.ScoringService, len(scs))
	for i, sc := range scs {
		ss[i] = toScoringService(sc)
	}

	return ss, nil
}

func (s *Service) DeleteScoringService(modelName string, port int) error {
	if ss, err := s.getScoringService(modelName); ss != nil {
		if ss.State != web.ScoringServiceStopped {
			return fmt.Errorf("Cannot delete. Scoring service on model %s is "+
				"still running.", modelName)
		}
	} else if err != nil {
		return fmt.Errorf("Scoring service for model %s does not exits.",
			modelName)
	}

	return s.ds.DeleteScoringService(genScoringServiceName(modelName))
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

func toModel(m *db.Model) *web.Model {
	return &web.Model{
		m.ID,
		m.CloudName,
		m.Dataset,
		m.TargetName,
		m.MaxRuntime,
		m.JavaModelPath,
		m.GenModelPath,
		web.Timestamp(m.CreatedAt),
	}
}

func toScoringService(s *db.ScoringService) *web.ScoringService {
	return &web.ScoringService{
		s.ModelName,
		s.Address,
		s.Port,
		web.ScoringServiceState(s.State),
		s.Pid,
		web.Timestamp(s.CreatedAt),
	}
}
