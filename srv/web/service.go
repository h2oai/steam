// ----------------------------------
// --- Generated with go:generate ---
// ---        DO NOT EDIT         ---
// ----------------------------------

package web

import (
	"net/http"
)

// --- Aliases ---

type CloudState string
type ScoringServiceState string
type Timestamp int64

// --- Consts ---

const (
	CloudStarted CloudState = "Started"
	CloudHealthy            = "Healthy"
	CloudStopped            = "Stopped"
	CloudUnknown            = "Unknown"
)

const (
	ScoringServiceStarted ScoringServiceState = "Started"
	ScoringServiceStopped                     = "Stopped"
)

// --- Types ---

type Cloud struct {
	CreatedAt     Timestamp  `json:"created_at"`
	Name          string     `json:"name"`
	EngineName    string     `json:"engine_name"`
	EngineVersion string     `json:"engine_version"`
	Size          int        `json:"size"`
	Memory        string     `json:"memory"`
	TotalCores    int        `json:"total_cores"`
	AllowedCores  int        `json:"allowed_cores"`
	State         CloudState `json:"state"`
	Address       string     `json:"address"`
	Username      string     `json:"username"`
	ApplicationID string     `json:"application_id"`
}

type Model struct {
	Name          string    `json:"name"`
	CloudName     string    `json:"cloud_name"`
	Dataset       string    `json:"dataset"`
	TargetName    string    `json:"target_name"`
	MaxRuntime    int       `json:"max_runtime"`
	JavaModelPath string    `json:"java_model_path"`
	GenModelPath  string    `json:"gen_model_path"`
	CreatedAt     Timestamp `json:"created_at"`
}

type ScoringService struct {
	ModelName string              `json:"model_name"`
	Address   string              `json:"address"`
	Port      int                 `json:"port"`
	State     ScoringServiceState `json:"state"`
	Pid       int                 `json:"pid"`
	CreatedAt Timestamp           `json:"created_at"`
}

type Engine struct {
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	CreatedAt Timestamp `json:"created_at"`
}

// --- Interfaces ---

type Service interface {
	Ping(status bool) (bool, error)
	StartCloud(cloudName string, engineName string, size int, memory string, username string) (*Cloud, error)
	StopCloud(cloudName string) error
	GetCloud(cloudName string) (*Cloud, error)
	GetClouds() ([]*Cloud, error)
	GetCloudStatus(cloudName string) (*Cloud, error)
	DeleteCloud(cloudName string) error
	BuildModel(cloudName string, dataset string, targetName string, maxRunTime int) (*Model, error)
	GetModel(modelName string) (*Model, error)
	GetModels() ([]*Model, error)
	GetCloudModels(cloudName string) ([]*Model, error)
	DeleteModel(modelName string) error
	StartScoringService(modelName string, port int) (*ScoringService, error)
	StopScoringService(modelName string, port int) error
	GetScoringService(modelName string) (*ScoringService, error)
	GetScoringServices() ([]*ScoringService, error)
	DeleteScoringService(modelName string, port int) error
	AddEngine(engineName string, enginePath string) error
	GetEngine(engineName string) (*Engine, error)
	GetEngines() ([]*Engine, error)
	DeleteEngine(engineName string) error
}

// --- Messages ---

type PingIn struct {
	Status bool `json:"status"`
}

type PingOut struct {
	Status bool `json:"status"`
}

type StartCloudIn struct {
	CloudName  string `json:"cloud_name"`
	EngineName string `json:"engine_name"`
	Size       int    `json:"size"`
	Memory     string `json:"memory"`
	Username   string `json:"username"`
}

type StartCloudOut struct {
	Cloud *Cloud `json:"cloud"`
}

type StopCloudIn struct {
	CloudName string `json:"cloud_name"`
}

type StopCloudOut struct {
}

type GetCloudIn struct {
	CloudName string `json:"cloud_name"`
}

type GetCloudOut struct {
	Cloud *Cloud `json:"cloud"`
}

type GetCloudsIn struct {
}

type GetCloudsOut struct {
	Clouds []*Cloud `json:"clouds"`
}

type GetCloudStatusIn struct {
	CloudName string `json:"cloud_name"`
}

type GetCloudStatusOut struct {
	Cloud *Cloud `json:"cloud"`
}

type DeleteCloudIn struct {
	CloudName string `json:"cloud_name"`
}

type DeleteCloudOut struct {
}

type BuildModelIn struct {
	CloudName  string `json:"cloud_name"`
	Dataset    string `json:"dataset"`
	TargetName string `json:"target_name"`
	MaxRunTime int    `json:"max_run_time"`
}

type BuildModelOut struct {
	Model *Model `json:"model"`
}

type GetModelIn struct {
	ModelName string `json:"model_name"`
}

type GetModelOut struct {
	Model *Model `json:"model"`
}

type GetModelsIn struct {
}

type GetModelsOut struct {
	Models []*Model `json:"models"`
}

type GetCloudModelsIn struct {
	CloudName string `json:"cloud_name"`
}

type GetCloudModelsOut struct {
	Models []*Model `json:"models"`
}

type DeleteModelIn struct {
	ModelName string `json:"model_name"`
}

type DeleteModelOut struct {
}

type StartScoringServiceIn struct {
	ModelName string `json:"model_name"`
	Port      int    `json:"port"`
}

type StartScoringServiceOut struct {
	Service *ScoringService `json:"service"`
}

type StopScoringServiceIn struct {
	ModelName string `json:"model_name"`
	Port      int    `json:"port"`
}

type StopScoringServiceOut struct {
}

type GetScoringServiceIn struct {
	ModelName string `json:"model_name"`
}

type GetScoringServiceOut struct {
	Service *ScoringService `json:"service"`
}

type GetScoringServicesIn struct {
}

type GetScoringServicesOut struct {
	Services []*ScoringService `json:"services"`
}

type DeleteScoringServiceIn struct {
	ModelName string `json:"model_name"`
	Port      int    `json:"port"`
}

type DeleteScoringServiceOut struct {
}

type AddEngineIn struct {
	EngineName string `json:"engine_name"`
	EnginePath string `json:"engine_path"`
}

type AddEngineOut struct {
}

type GetEngineIn struct {
	EngineName string `json:"engine_name"`
}

type GetEngineOut struct {
	Engine *Engine `json:"engine"`
}

type GetEnginesIn struct {
}

type GetEnginesOut struct {
	Engines []*Engine `json:"engines"`
}

type DeleteEngineIn struct {
	EngineName string `json:"engine_name"`
}

type DeleteEngineOut struct {
}

// --- Client Stub ---

type Remote struct {
	Proc Proc
}

type Proc interface {
	Call(name string, in, out interface{}) error
}

func (this *Remote) Ping(status bool) (bool, error) {
	in := PingIn{status}
	var out PingOut
	err := this.Proc.Call("Ping", &in, &out)
	if err != nil {
		return false, err
	}
	return out.Status, nil
}

func (this *Remote) StartCloud(cloudName string, engineName string, size int, memory string, username string) (*Cloud, error) {
	in := StartCloudIn{cloudName, engineName, size, memory, username}
	var out StartCloudOut
	err := this.Proc.Call("StartCloud", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Cloud, nil
}

func (this *Remote) StopCloud(cloudName string) error {
	in := StopCloudIn{cloudName}
	var out StopCloudOut
	err := this.Proc.Call("StopCloud", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetCloud(cloudName string) (*Cloud, error) {
	in := GetCloudIn{cloudName}
	var out GetCloudOut
	err := this.Proc.Call("GetCloud", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Cloud, nil
}

func (this *Remote) GetClouds() ([]*Cloud, error) {
	in := GetCloudsIn{}
	var out GetCloudsOut
	err := this.Proc.Call("GetClouds", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Clouds, nil
}

func (this *Remote) GetCloudStatus(cloudName string) (*Cloud, error) {
	in := GetCloudStatusIn{cloudName}
	var out GetCloudStatusOut
	err := this.Proc.Call("GetCloudStatus", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Cloud, nil
}

func (this *Remote) DeleteCloud(cloudName string) error {
	in := DeleteCloudIn{cloudName}
	var out DeleteCloudOut
	err := this.Proc.Call("DeleteCloud", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) BuildModel(cloudName string, dataset string, targetName string, maxRunTime int) (*Model, error) {
	in := BuildModelIn{cloudName, dataset, targetName, maxRunTime}
	var out BuildModelOut
	err := this.Proc.Call("BuildModel", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Model, nil
}

func (this *Remote) GetModel(modelName string) (*Model, error) {
	in := GetModelIn{modelName}
	var out GetModelOut
	err := this.Proc.Call("GetModel", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Model, nil
}

func (this *Remote) GetModels() ([]*Model, error) {
	in := GetModelsIn{}
	var out GetModelsOut
	err := this.Proc.Call("GetModels", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Models, nil
}

func (this *Remote) GetCloudModels(cloudName string) ([]*Model, error) {
	in := GetCloudModelsIn{cloudName}
	var out GetCloudModelsOut
	err := this.Proc.Call("GetCloudModels", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Models, nil
}

func (this *Remote) DeleteModel(modelName string) error {
	in := DeleteModelIn{modelName}
	var out DeleteModelOut
	err := this.Proc.Call("DeleteModel", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) StartScoringService(modelName string, port int) (*ScoringService, error) {
	in := StartScoringServiceIn{modelName, port}
	var out StartScoringServiceOut
	err := this.Proc.Call("StartScoringService", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Service, nil
}

func (this *Remote) StopScoringService(modelName string, port int) error {
	in := StopScoringServiceIn{modelName, port}
	var out StopScoringServiceOut
	err := this.Proc.Call("StopScoringService", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetScoringService(modelName string) (*ScoringService, error) {
	in := GetScoringServiceIn{modelName}
	var out GetScoringServiceOut
	err := this.Proc.Call("GetScoringService", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Service, nil
}

func (this *Remote) GetScoringServices() ([]*ScoringService, error) {
	in := GetScoringServicesIn{}
	var out GetScoringServicesOut
	err := this.Proc.Call("GetScoringServices", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Services, nil
}

func (this *Remote) DeleteScoringService(modelName string, port int) error {
	in := DeleteScoringServiceIn{modelName, port}
	var out DeleteScoringServiceOut
	err := this.Proc.Call("DeleteScoringService", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) AddEngine(engineName string, enginePath string) error {
	in := AddEngineIn{engineName, enginePath}
	var out AddEngineOut
	err := this.Proc.Call("AddEngine", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetEngine(engineName string) (*Engine, error) {
	in := GetEngineIn{engineName}
	var out GetEngineOut
	err := this.Proc.Call("GetEngine", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Engine, nil
}

func (this *Remote) GetEngines() ([]*Engine, error) {
	in := GetEnginesIn{}
	var out GetEnginesOut
	err := this.Proc.Call("GetEngines", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Engines, nil
}

func (this *Remote) DeleteEngine(engineName string) error {
	in := DeleteEngineIn{engineName}
	var out DeleteEngineOut
	err := this.Proc.Call("DeleteEngine", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

// --- Server Stub ---

type Impl struct {
	Service Service
}

func (this *Impl) Ping(r *http.Request, in *PingIn, out *PingOut) error {
	it, err := this.Service.Ping(in.Status)
	if err != nil {
		return err
	}
	out.Status = it
	return nil
}

func (this *Impl) StartCloud(r *http.Request, in *StartCloudIn, out *StartCloudOut) error {
	it, err := this.Service.StartCloud(in.CloudName, in.EngineName, in.Size, in.Memory, in.Username)
	if err != nil {
		return err
	}
	out.Cloud = it
	return nil
}

func (this *Impl) StopCloud(r *http.Request, in *StopCloudIn, out *StopCloudOut) error {
	err := this.Service.StopCloud(in.CloudName)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) GetCloud(r *http.Request, in *GetCloudIn, out *GetCloudOut) error {
	it, err := this.Service.GetCloud(in.CloudName)
	if err != nil {
		return err
	}
	out.Cloud = it
	return nil
}

func (this *Impl) GetClouds(r *http.Request, in *GetCloudsIn, out *GetCloudsOut) error {
	it, err := this.Service.GetClouds()
	if err != nil {
		return err
	}
	out.Clouds = it
	return nil
}

func (this *Impl) GetCloudStatus(r *http.Request, in *GetCloudStatusIn, out *GetCloudStatusOut) error {
	it, err := this.Service.GetCloudStatus(in.CloudName)
	if err != nil {
		return err
	}
	out.Cloud = it
	return nil
}

func (this *Impl) DeleteCloud(r *http.Request, in *DeleteCloudIn, out *DeleteCloudOut) error {
	err := this.Service.DeleteCloud(in.CloudName)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) BuildModel(r *http.Request, in *BuildModelIn, out *BuildModelOut) error {
	it, err := this.Service.BuildModel(in.CloudName, in.Dataset, in.TargetName, in.MaxRunTime)
	if err != nil {
		return err
	}
	out.Model = it
	return nil
}

func (this *Impl) GetModel(r *http.Request, in *GetModelIn, out *GetModelOut) error {
	it, err := this.Service.GetModel(in.ModelName)
	if err != nil {
		return err
	}
	out.Model = it
	return nil
}

func (this *Impl) GetModels(r *http.Request, in *GetModelsIn, out *GetModelsOut) error {
	it, err := this.Service.GetModels()
	if err != nil {
		return err
	}
	out.Models = it
	return nil
}

func (this *Impl) GetCloudModels(r *http.Request, in *GetCloudModelsIn, out *GetCloudModelsOut) error {
	it, err := this.Service.GetCloudModels(in.CloudName)
	if err != nil {
		return err
	}
	out.Models = it
	return nil
}

func (this *Impl) DeleteModel(r *http.Request, in *DeleteModelIn, out *DeleteModelOut) error {
	err := this.Service.DeleteModel(in.ModelName)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) StartScoringService(r *http.Request, in *StartScoringServiceIn, out *StartScoringServiceOut) error {
	it, err := this.Service.StartScoringService(in.ModelName, in.Port)
	if err != nil {
		return err
	}
	out.Service = it
	return nil
}

func (this *Impl) StopScoringService(r *http.Request, in *StopScoringServiceIn, out *StopScoringServiceOut) error {
	err := this.Service.StopScoringService(in.ModelName, in.Port)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) GetScoringService(r *http.Request, in *GetScoringServiceIn, out *GetScoringServiceOut) error {
	it, err := this.Service.GetScoringService(in.ModelName)
	if err != nil {
		return err
	}
	out.Service = it
	return nil
}

func (this *Impl) GetScoringServices(r *http.Request, in *GetScoringServicesIn, out *GetScoringServicesOut) error {
	it, err := this.Service.GetScoringServices()
	if err != nil {
		return err
	}
	out.Services = it
	return nil
}

func (this *Impl) DeleteScoringService(r *http.Request, in *DeleteScoringServiceIn, out *DeleteScoringServiceOut) error {
	err := this.Service.DeleteScoringService(in.ModelName, in.Port)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) AddEngine(r *http.Request, in *AddEngineIn, out *AddEngineOut) error {
	err := this.Service.AddEngine(in.EngineName, in.EnginePath)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) GetEngine(r *http.Request, in *GetEngineIn, out *GetEngineOut) error {
	it, err := this.Service.GetEngine(in.EngineName)
	if err != nil {
		return err
	}
	out.Engine = it
	return nil
}

func (this *Impl) GetEngines(r *http.Request, in *GetEnginesIn, out *GetEnginesOut) error {
	it, err := this.Service.GetEngines()
	if err != nil {
		return err
	}
	out.Engines = it
	return nil
}

func (this *Impl) DeleteEngine(r *http.Request, in *DeleteEngineIn, out *DeleteEngineOut) error {
	err := this.Service.DeleteEngine(in.EngineName)
	if err != nil {
		return err
	}
	return nil
}
