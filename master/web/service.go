package web

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/h2oai/steamY/bindings"
	"github.com/h2oai/steamY/lib/fs"
	"github.com/h2oai/steamY/lib/proxy"
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
	clusterProxy              *proxy.RProxy
	cloudActivity             map[string]web.Timestamp // TODO: not threadsafe
	scoreActivity             map[string]web.Timestamp // TODO: not threadsafe
}

func toTimestamp(t time.Time) web.Timestamp {
	return web.Timestamp(t.UTC().Unix())
}

func now() web.Timestamp {
	return toTimestamp(time.Now())
}

func NewService(workingDir string, ds *db.DS, compilationServiceAddress, scoringServiceAddress string, clusterProxy *proxy.RProxy, kerberos bool, username, keytab string) *web.Impl {
	return &web.Impl{&Service{
		workingDir,
		ds,
		compilationServiceAddress,
		scoringServiceAddress,
		kerberos,
		username,
		keytab,
		clusterProxy,
		make(map[string]web.Timestamp),
		make(map[string]web.Timestamp),
	}}
}

func (s *Service) Ping(status bool) (bool, error) {
	return status, nil
}

func (s *Service) poll() {
	log.Println("Polling for cloud activity")
	cs, err := s.GetClouds()
	if err != nil {
		log.Println(err)
	}
	for _, c := range cs {
		if c.State != web.CloudStopped {
			js, err := s.GetJobs(c.Name)
			if err != nil {
				log.Println(err)
			}
			if len(js) > 0 {
				j := js[0]
				if j.Progress == "DONE" {
					s.cloudActivity[c.Name] = j.FinishedAt / 1000 // TODO: not threadsafe
				} else {
					s.cloudActivity[c.Name] = now() // TODO: not threadsafe
				}
			} else {
				s.cloudActivity[c.Name] = c.CreatedAt // TODO: not threadsafe
			}
		}
	}
	log.Println("Polling for scoring service activity")
	ss, err := s.GetScoringServices()
	if err != nil {
		log.Println(err)
	}
	for _, sc := range ss {
		if sc.State != web.ScoringServiceStopped {
			s.scoreActivity[sc.ModelName], err = svc.Poll(sc)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (s *Service) ActivityPoll(status bool) (bool, error) {
	log.Println("Polling for activity")
	s.poll() // Fill cloudActivity map with running clouds on startup

	go func() {
		tickInterval := time.Hour

		ticker := time.NewTicker(tickInterval)
		for {
			select {
			case <-ticker.C:
				s.poll()
			}
		}
	}()
	return status, nil
}

func (s *Service) InitClusterProxy() (bool, error) {
	cs, err := s.ds.ListCloud()
	if err != nil {
		return false, err //FIXME format error
	}

	for _, c := range cs {
		if c.State != "Stopped" {
			s.clusterProxy.NewProxy(c.ID, c.Address)
		}
	}
	return true, nil
}

func (s *Service) RegisterCloud(address string) (*web.Cloud, error) {

	h := h2ov3.NewClient(address)
	cloud, err := h.GetCloud()
	if err != nil {
		return nil, fmt.Errorf("Could not communicate with cloud %s.", address)
	}

	if _, err := s.getCloud(cloud.CloudName); err == nil {
		return nil, fmt.Errorf("Cloud registration failed. A cloud with the address %s is already registered.", address)
	}

	c := db.NewCloud(
		cloud.CloudName,
		"",
		int(cloud.CloudSize),
		"",
		address,
		"",
		"",
		string(web.CloudStarted),
		"",
	)
	if err := s.ds.CreateCloud(c); err != nil {
		return nil, err
	}
	s.clusterProxy.NewProxy(c.ID, c.Address)

	return toCloud(c), nil
}

func (s *Service) UnregisterCloud(cloudName string) error {

	// Make sure the cloud exists
	c, err := s.getCloud(cloudName)
	if err != nil {
		return err
	}

	// HACK: if the engine name is empty, this is not an external cloud. So bail out.
	if c.EngineName != "" {
		return fmt.Errorf("Cannot unregister internal clouds.")
	}

	// Permanently delete this cloud
	if err := s.ds.DeleteCloud(cloudName); err != nil {
		return err
	}

	return nil
}

func (s *Service) StartCloud(cloudName, engineName string, size int, memory, username string) (*web.Cloud, error) {
	// Make sure this cloud is unique
	if _, err := s.getCloud(cloudName); err == nil {
		return nil, fmt.Errorf("Cloud start failed. A cloud with the name %s already exists.", cloudName)
	}

	e, err := s.getEngine(engineName)
	if err != nil {
		return nil, fmt.Errorf("Cloud start failed. Cannot locate engine %s.", engineName)
	}

	// Make cloud with yarn
	appID, address, out, err := yarn.StartCloud(size, s.kerberosEnabled, memory, cloudName, e.Path, s.username, s.keytab) // FIXME: THIS IS USING ADMIN TO START ALL CLOUDS
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Add cloud to db
	c := db.NewCloud(
		cloudName,
		engineName,
		size,
		appID,
		address,
		memory,
		username,
		string(web.CloudStarted),
		out,
	)

	if err := s.ds.CreateCloud(c); err != nil {
		return nil, err
	}
	// Create an instance of this cloud in cloudActivity map
	s.cloudActivity[c.ID] = web.Timestamp(c.CreatedAt) // TODO: not threadsafe
	// Create a reverse proxy for cluster.
	s.clusterProxy.NewProxy(c.ID, c.Address)

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

	if err := yarn.StopCloud(s.kerberosEnabled, c.ID, c.ApplicationID, c.Out, s.username, s.keytab); err != nil { //FIXME: this is using adming kerberos credentials
		log.Println(err)
		return err
	}

	// Update the state and update DB
	c.State = string(web.CloudStopped)
	if err := s.ds.UpdateCloud(c); err != nil {
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
		clouds[i].Activity = web.Timestamp(s.cloudActivity[c.ID]) // update to last known activity
	}
	return clouds, nil
}

// Returns the Cloud status from H2O
// This method should only be called if the cluster reports a non-Stopped status
// If the cloud was shut down from the outside of steam, will report Unknown
// / status for cloud
//
// TODO: Maybe this should only report if non-Stopped,non-Unknown status
//       In the case of Unknown, should only check if forced?
func (s *Service) GetCloudStatus(cloudName string) (*web.Cloud, error) { // Only called if cloud status != found
	c, err := s.getCloud(cloudName)
	if err != nil {
		return nil, fmt.Errorf("Cannot find cluster %s in GetCloudStatus:\n%v", cloudName, err)
	}

	h := h2ov3.NewClient(c.Address)

	cloud, err := h.GetCloud()
	if err != nil { // Cloud just isn't found
		c.State = web.CloudUnknown
		log.Printf("Error from GetCloud in GetCloudStatus:\n%v", err)
		return nil, fmt.Errorf("Cannot find cluster %s, is it still running?", c.ID)
	}

	var (
		tot, all int32
		mem      int64
	)
	for _, n := range cloud.Nodes {
		mem += n.MaxMem
		tot += n.NumCpus
		all += n.CpusAllowed
	}
	var health web.CloudState
	if cloud.CloudHealthy {
		health = web.CloudHealthy
	} else {
		health = web.CloudUnknown
	}

	return &web.Cloud{
		web.Timestamp(c.CreatedAt),
		c.ID,
		c.EngineName,
		cloud.Version,
		c.Size,
		toSizeBytes(mem),
		int(tot),
		int(all),
		health,
		c.Address,
		c.Username,
		c.ApplicationID,
		web.Timestamp(s.cloudActivity[c.ID]),
	}, nil
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

//
// Functions for sorting jobs
//

type Jobs []*web.Job

func (k Jobs) Len() int {
	return len(k)
}

func (k Jobs) Less(i, j int) bool {
	switch {
	case k[i].Progress == "DONE" && k[j].Progress == "DONE":
		return k[i].FinishedAt < k[j].FinishedAt
	case k[i].Progress == "DONE":
		return true
	case k[j].Progress == "DONE":
		return false
	default:
		return k[i].FinishedAt < k[j].FinishedAt
	}
}

func (k Jobs) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

//
// End
//

func (s *Service) GetJob(cloudName, jobName string) (*web.Job, error) {
	c, err := s.getCloud(cloudName)
	if err != nil {
		return nil, err //FIXME format error
	}

	h := h2ov3.NewClient(c.Address)

	j, err := h.GetJob(jobName)
	if err != nil {
		return nil, err //FIXME format error
	}
	job := j.Jobs[0]

	return htoJob(job), nil
}

func (s *Service) GetJobs(cloudName string) ([]*web.Job, error) {
	c, err := s.getCloud(cloudName)
	if err != nil {
		return nil, err //FIXME format error
	}

	h := h2ov3.NewClient(c.Address)

	j, err := h.GetJobs()
	if err != nil {
		return nil, err //FIXME format error
	}

	jobs := make([]*web.Job, len(j.Jobs))
	for i, job := range j.Jobs {
		jobs[i] = htoJob(job)
	}

	sort.Sort(sort.Reverse(Jobs(jobs)))

	return jobs, nil
}

func (s *Service) BuildModel(cloudName string, dataset string, targetName string, maxRunTime int) (*web.Model, error) {
	c, err := s.GetCloud(cloudName)
	if err != nil {
		return nil, err
	}
	if c.State == web.CloudStopped {
		return nil, fmt.Errorf("%s is a stopped cloud. Cannot build a model.", cloudName)
	}
	h := h2ov3.NewClient(c.Address)

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
		"AutoML",
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

func (s *Service) GetCloudModels(cloudName string) ([]*web.Model, error) {
	c, err := s.getCloud(cloudName)
	if err != nil {
		return nil, fmt.Errorf("Cannot find cloud %s in GetCloudModels:\n%v", cloudName, err)
	}

	h := h2ov3.NewClient(c.Address)
	ms, err := h.GetModels()
	if err != nil {
		return nil, fmt.Errorf("Cannot reach models in cloud %s in GetCloudModels:\n%v", cloudName, err)
	}

	models := make([]*web.Model, len(ms.Models))
	for i, m := range ms.Models {
		models[i] = &web.Model{
			Name:       m.ModelId.Name,
			CloudName:  cloudName,
			Algo:       m.AlgoFullName,
			Dataset:    m.DataFrame.Name,
			TargetName: m.ResponseColumnName,
			CreatedAt:  web.Timestamp(m.Timestamp),
		}
	}

	return models, nil
}

//TODO this is messy, fix it
func (s *Service) GetModelFromCloud(cloudName string, modelName string) (*web.Model, error) {
	// get cloud, must exist in db
	c, err := s.getCloud(cloudName)
	if err != nil {
		return nil, err // FIXME fmt error
	}

	// get model from the cloud
	h := h2ov3.NewClient(c.Address)
	r, err := h.GetModel(modelName)
	if err != nil {
		return nil, err //FIXME fmt error
	}

	m := r.Models[0]

	// get pojo and genmodel
	modelDir := fs.GetModelPath(s.workingDir, modelName, "java")
	jm, err := h.ExportJavaModel(modelName, modelDir)
	if err != nil {
		return nil, err //FIXME fmt error
	}
	gm, err := h.ExportGenModel(modelDir)
	if err != nil {
		return nil, err //FIXME fmt error
	}

	model := db.NewModel(
		modelName,
		cloudName,
		m.AlgoFullName,
		m.DataFrame.Name,
		m.ResponseColumnName,
		-1,
		jm,
		gm,
	)
	if err := s.ds.CreateModel(model); err != nil {
		return nil, err //FIXME fmt error
	}

	mod := toModel(model)
	mod.Algo = m.AlgoFullName

	return mod, nil
}

func (s *Service) DeleteModel(modelName string) error {
	ss, err := s.getScoringService(modelName)
	if err == nil {

		if ss.State != web.ScoringServiceStopped {
			return fmt.Errorf("Cannot delete. A scoring service on model %s is"+
				" still running.", modelName)
		}
	}

	if _, err := s.getModel(modelName); err != nil {
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

	// e := fs.GetAssetsPath(s.workingDir, "makewar-extra.jar")

	w, err := c.CompilePojo(j, g, "makewar")
	if err != nil {
		return "", err
	}
	log.Printf("Warfile Dest: %s", w)

	return w, nil
}

func (s *Service) StartScoringService(modelName string, port int) (*web.ScoringService, error) {
	if _, ok := s.getScoringService(modelName); ok == nil {
		return nil, fmt.Errorf("A scoring service with the model %s already exists.", modelName)
	}

	w, err := s.compileModel(modelName)
	if err != nil {
		return nil, err
	}

	const jetty = "jetty-runner.jar"

	j := fs.GetAssetsPath(s.workingDir, jetty)

	pid, err := svc.Start(w, j, s.scoringServiceAddress, port)
	if err != nil {
		return nil, err
	}

	externalIP, err := fs.GetExternalHost()
	if err != nil {
		return nil, err
	}
	ss := db.NewScoringService(
		modelName,
		modelName,
		externalIP,
		port,
		string(web.ScoringServiceStarted),
		pid,
	)

	log.Printf("Scoring service started at %s:%d\n", externalIP, ss.Port)

	if err := s.ds.CreateScoringService(ss); err != nil {
		return nil, err
	}

	s.scoreActivity[modelName] = web.Timestamp(ss.CreatedAt)
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
	ss, err := s.ds.ReadScoringService(modelName)
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
	return toScoringService(ss), nil
}

func (s *Service) GetScoringServices() ([]*web.ScoringService, error) {
	scs, err := s.ds.ListScoringService()
	if err != nil {
		return nil, err
	}
	ss := make([]*web.ScoringService, len(scs))
	for i, sc := range scs {
		ss[i] = toScoringService(sc)
		ss[i].Activity = s.scoreActivity[sc.ModelName]
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

	return s.ds.DeleteScoringService(modelName)
}

func (s *Service) AddEngine(engineName, enginePath string) error {
	e := db.NewEngine(
		engineName,
		enginePath,
	)
	if err := s.ds.CreateEngine(e); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetEngine(engineName string) (*web.Engine, error) {
	e, err := s.getEngine(engineName)
	if err != nil {
		return nil, err
	}
	return toEngine(e), nil
}

func (s *Service) getEngine(engineName string) (*db.Engine, error) {
	e, err := s.ds.ReadEngine(engineName)
	if err != nil {
		return nil, err
	}
	if e == nil {
		return nil, fmt.Errorf("Engine %s does not exist.", engineName)
	}
	return e, err
}

func (s *Service) GetEngines() ([]*web.Engine, error) {
	es, err := s.ds.ListEngine()
	if err != nil {
		return nil, err
	}

	engines := make([]*web.Engine, len(es))
	for i, e := range es {
		engines[i] = toEngine(e)
	}

	return engines, nil
}

func (s *Service) DeleteEngine(engineName string) error {
	// TODO delete jarfile from disk?
	if _, err := s.getEngine(engineName); err != nil {
		return fmt.Errorf("Engine %s does not exits.", engineName)
	}
	return s.ds.DeleteEngine(engineName)
}

// Helper function to convert from int to bytes
func toSizeBytes(i int64) string {
	f := float64(i)

	s := 0
	for f > 1024 {
		f /= 1024
		s++
	}
	b := strconv.FormatFloat(f, 'f', 2, 64)

	switch s {
	case 0:
		return b + " B"
	case 1:
		return b + " KB"
	case 2:
		return b + " MB"
	case 3:
		return b + " GB"
	case 4:
		return b + " TB"
	case 5:
		return b + " PB"
	}

	return ""
}

//
// Routines to convert DB structs into API structs
//

func toCloud(c *db.Cloud) *web.Cloud {
	return &web.Cloud{
		web.Timestamp(c.CreatedAt), // CreatedAt
		c.ID,         // Name
		c.EngineName, // EngineName
		"",           // EngineVerion
		c.Size,       // Size
		"",           // Memory
		0,            // TotalCores
		0,            // AllowedCores
		web.CloudState(c.State), // State
		c.Address,               // Address
		c.Username,              // Username
		c.ApplicationID,         // ApplicationID
		0,                       //Activity
	}
}

func toModel(m *db.Model) *web.Model {
	return &web.Model{
		m.ID,                       // Name
		m.CloudName,                // CloudName
		m.Algo,                     // Algo
		m.Dataset,                  // Dataset
		m.TargetName,               // TargetName
		m.MaxRuntime,               // MaxRunTime
		m.JavaModelPath,            // JavaModelPath
		m.GenModelPath,             // GenModelPath
		web.Timestamp(m.CreatedAt), // CreatedAt
	}
}

func toScoringService(s *db.ScoringService) *web.ScoringService {
	return &web.ScoringService{
		ModelName: s.ModelName,
		Address:   s.Address,
		Port:      s.Port,
		State:     web.ScoringServiceState(s.State),
		Pid:       s.Pid,
		CreatedAt: web.Timestamp(s.CreatedAt),
	}
}

func toEngine(e *db.Engine) *web.Engine {
	return &web.Engine{
		e.ID,
		e.Path,
		web.Timestamp(e.CreatedAt),
	}
}

//
// Routines to convert H2O structs into API structs
//

func htoJob(j *bindings.JobV3) *web.Job {
	var end int64
	if j.Status == "DONE" {
		end = j.StartTime + j.Msec
	}

	return &web.Job{
		Name:        j.Key.Name,
		Description: j.Description,
		Progress:    j.Status,
		CreatedAt:   web.Timestamp(j.StartTime),
		FinishedAt:  web.Timestamp(end),
	}
}
