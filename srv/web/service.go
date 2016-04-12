// ----------------------------------
// --- Generated with go:generate ---
// ---        DO NOT EDIT         ---
// ----------------------------------

package web

import(
	"net/http"
)

// --- Aliases ---

type CloudState string
type Timestamp int64

// --- Types ---

type Cloud struct {
	Name string `json:"name"`
	Pack string `json:"pack"`
	State CloudState `json:"state"`
}

type CloudModelSynopsis struct {
	Algorithm string `json:"algorithm"`
	AlgorithmFullName string `json:"algorithm_full_name"`
	FrameName string `json:"frame_name"`
	ModelName string `json:"model_name"`
	ResponseColumnName string `json:"response_column_name"`
	ModifiedAt Timestamp `json:"modified_at"`
}

// --- Interfaces ---

type Service interface {
	Ping(status bool) (bool, error)
	StartCloud(size int, kerberos bool, name string, username string, keytab string) (string, error)
	StopCloud(kerberos bool, name string, id string, username string, keytab string) error
	GetCloud(address string) (*Cloud, error)
	BuildAutoML(address string, dataset string, targetName string) error
	GetModels(address string) ([]*CloudModelSynopsis, error)
	GetModel(address string, modelID string) (*RawModel, error)
	CompilePojo(address string, javaModel string, jar string) error
	Shutdown(address string) error
}

// --- Messages ---

type PingIn struct {
	Status bool `json:"status"`
}

type PingOut struct {
	Status bool `json:"status"`
}

type StartCloudIn struct {
	Size int `json:"size"`
	Kerberos bool `json:"kerberos"`
	Name string `json:"name"`
	Username string `json:"username"`
	Keytab string `json:"keytab"`
}

type StartCloudOut struct {
	ApID string `json:"ap_id"`
}

type StopCloudIn struct {
	Kerberos bool `json:"kerberos"`
	Name string `json:"name"`
	Id string `json:"id"`
	Username string `json:"username"`
	Keytab string `json:"keytab"`
}

type StopCloudOut struct {
}

type GetCloudIn struct {
	Address string `json:"address"`
}

type GetCloudOut struct {
	Cloud *Cloud `json:"cloud"`
}

type BuildAutoMLIn struct {
	Address string `json:"address"`
	Dataset string `json:"dataset"`
	TargetName string `json:"target_name"`
}

type BuildAutoMLOut struct {
}

type GetModelsIn struct {
	Address string `json:"address"`
}

type GetModelsOut struct {
	Models []*CloudModelSynopsis `json:"models"`
}

type GetModelIn struct {
	Address string `json:"address"`
	ModelID string `json:"model_id"`
}

type GetModelOut struct {
	Model *RawModel `json:"model"`
}

type CompilePojoIn struct {
	Address string `json:"address"`
	JavaModel string `json:"java_model"`
	Jar string `json:"jar"`
}

type CompilePojoOut struct {
}

type ShutdownIn struct {
	Address string `json:"address"`
}

type ShutdownOut struct {
}

// --- Client Stub ---

type Remote struct {
	Proc Proc
}

type Proc interface {
	Call(name string, in, out interface{}) error
}

func (this *Remote) Ping(status bool) (bool,  error) {
	in := PingIn{status}
	var out PingOut
	err := this.Proc.Call("Ping", &in, &out)
	if err != nil {
		return false, err
	}
	return out.Status, nil
}

func (this *Remote) StartCloud(size int, kerberos bool, name string, username string, keytab string) (string,  error) {
	in := StartCloudIn{size, kerberos, name, username, keytab}
	var out StartCloudOut
	err := this.Proc.Call("StartCloud", &in, &out)
	if err != nil {
		return "", err
	}
	return out.ApID, nil
}

func (this *Remote) StopCloud(kerberos bool, name string, id string, username string, keytab string) ( error) {
	in := StopCloudIn{kerberos, name, id, username, keytab}
	var out StopCloudOut
	err := this.Proc.Call("StopCloud", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetCloud(address string) (*Cloud,  error) {
	in := GetCloudIn{address}
	var out GetCloudOut
	err := this.Proc.Call("GetCloud", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Cloud, nil
}

func (this *Remote) BuildAutoML(address string, dataset string, targetName string) ( error) {
	in := BuildAutoMLIn{address, dataset, targetName}
	var out BuildAutoMLOut
	err := this.Proc.Call("BuildAutoML", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetModels(address string) ([]*CloudModelSynopsis,  error) {
	in := GetModelsIn{address}
	var out GetModelsOut
	err := this.Proc.Call("GetModels", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Models, nil
}

func (this *Remote) GetModel(address string, modelID string) (*RawModel,  error) {
	in := GetModelIn{address, modelID}
	var out GetModelOut
	err := this.Proc.Call("GetModel", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Model, nil
}

func (this *Remote) CompilePojo(address string, javaModel string, jar string) ( error) {
	in := CompilePojoIn{address, javaModel, jar}
	var out CompilePojoOut
	err := this.Proc.Call("CompilePojo", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) Shutdown(address string) ( error) {
	in := ShutdownIn{address}
	var out ShutdownOut
	err := this.Proc.Call("Shutdown", &in, &out)
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
	it, err := this.Service.StartCloud(in.Size, in.Kerberos, in.Name, in.Username, in.Keytab)
	if err != nil {
		return err
	}
	out.ApID = it
	return nil
}

func (this *Impl) StopCloud(r *http.Request, in *StopCloudIn, out *StopCloudOut) error {
	err := this.Service.StopCloud(in.Kerberos, in.Name, in.Id, in.Username, in.Keytab)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) GetCloud(r *http.Request, in *GetCloudIn, out *GetCloudOut) error {
	it, err := this.Service.GetCloud(in.Address)
	if err != nil {
		return err
	}
	out.Cloud = it
	return nil
}

func (this *Impl) BuildAutoML(r *http.Request, in *BuildAutoMLIn, out *BuildAutoMLOut) error {
	err := this.Service.BuildAutoML(in.Address, in.Dataset, in.TargetName)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) GetModels(r *http.Request, in *GetModelsIn, out *GetModelsOut) error {
	it, err := this.Service.GetModels(in.Address)
	if err != nil {
		return err
	}
	out.Models = it
	return nil
}

func (this *Impl) GetModel(r *http.Request, in *GetModelIn, out *GetModelOut) error {
	it, err := this.Service.GetModel(in.Address, in.ModelID)
	if err != nil {
		return err
	}
	out.Model = it
	return nil
}

func (this *Impl) CompilePojo(r *http.Request, in *CompilePojoIn, out *CompilePojoOut) error {
	err := this.Service.CompilePojo(in.Address, in.JavaModel, in.Jar)
	if err != nil {
		return err
	}
	return nil
}

func (this *Impl) Shutdown(r *http.Request, in *ShutdownIn, out *ShutdownOut) error {
	err := this.Service.Shutdown(in.Address)
	if err != nil {
		return err
	}
	return nil
}

