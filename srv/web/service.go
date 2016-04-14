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
type Timestamp int64

// --- Types ---

type CloudOpts struct {
	Name          string `json:"name"`
	ApplicationID string `json:"application_id"`
	Address       string `json:"address"`
}

type Cloud struct {
	Name  string     `json:"name"`
	Pack  string     `json:"pack"`
	State CloudState `json:"state"`
}

type CloudModelSynopsis struct {
	Algorithm          string    `json:"algorithm"`
	AlgorithmFullName  string    `json:"algorithm_full_name"`
	FrameName          string    `json:"frame_name"`
	ModelName          string    `json:"model_name"`
	ResponseColumnName string    `json:"response_column_name"`
	ModifiedAt         Timestamp `json:"modified_at"`
}

// --- Interfaces ---

type Service interface {
	Ping(status bool) (bool, error)
	StartCloud(name string, size int, mem string, useKerberos bool, username string, keytab string) (*CloudOpts, error)
	StopCloud(name string, useKerberos bool, applicationID string, username string, keytab string) error
	GetCloud(address string) (*Cloud, error)
	BuildAutoML(address string, dataset string, targetName string, maxRunTime int) (string, error)
	DeployPojo(address string, javaModelPath string, genModelPath string) error
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
	Name        string `json:"name"`
	Size        int    `json:"size"`
	Mem         string `json:"mem"`
	UseKerberos bool   `json:"use_kerberos"`
	Username    string `json:"username"`
	Keytab      string `json:"keytab"`
}

type StartCloudOut struct {
	Cloud *CloudOpts `json:"cloud"`
}

type StopCloudIn struct {
	Name          string `json:"name"`
	UseKerberos   bool   `json:"use_kerberos"`
	ApplicationID string `json:"application_id"`
	Username      string `json:"username"`
	Keytab        string `json:"keytab"`
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
	Address    string `json:"address"`
	Dataset    string `json:"dataset"`
	TargetName string `json:"target_name"`
	MaxRunTime int    `json:"max_run_time"`
}

type BuildAutoMLOut struct {
	ModelID string `json:"model_id"`
}

type DeployPojoIn struct {
	Address       string `json:"address"`
	JavaModelPath string `json:"java_model_path"`
	GenModelPath  string `json:"gen_model_path"`
}

type DeployPojoOut struct {
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

func (this *Remote) Ping(status bool) (bool, error) {
	in := PingIn{status}
	var out PingOut
	err := this.Proc.Call("Ping", &in, &out)
	if err != nil {
		return false, err
	}
	return out.Status, nil
}

func (this *Remote) StartCloud(name string, size int, mem string, useKerberos bool, username string, keytab string) (*CloudOpts, error) {
	in := StartCloudIn{name, size, mem, useKerberos, username, keytab}
	var out StartCloudOut
	err := this.Proc.Call("StartCloud", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Cloud, nil
}

func (this *Remote) StopCloud(name string, useKerberos bool, applicationID string, username string, keytab string) error {
	in := StopCloudIn{name, useKerberos, applicationID, username, keytab}
	var out StopCloudOut
	err := this.Proc.Call("StopCloud", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) GetCloud(address string) (*Cloud, error) {
	in := GetCloudIn{address}
	var out GetCloudOut
	err := this.Proc.Call("GetCloud", &in, &out)
	if err != nil {
		return nil, err
	}
	return out.Cloud, nil
}

func (this *Remote) BuildAutoML(address string, dataset string, targetName string, maxRunTime int) (string, error) {
	in := BuildAutoMLIn{address, dataset, targetName, maxRunTime}
	var out BuildAutoMLOut
	err := this.Proc.Call("BuildAutoML", &in, &out)
	if err != nil {
		return "", err
	}
	return out.ModelID, nil
}

func (this *Remote) DeployPojo(address string, javaModelPath string, genModelPath string) error {
	in := DeployPojoIn{address, javaModelPath, genModelPath}
	var out DeployPojoOut
	err := this.Proc.Call("DeployPojo", &in, &out)
	if err != nil {
		return err
	}
	return nil
}

func (this *Remote) Shutdown(address string) error {
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
	it, err := this.Service.StartCloud(in.Name, in.Size, in.Mem, in.UseKerberos, in.Username, in.Keytab)
	if err != nil {
		return err
	}
	out.Cloud = it
	return nil
}

func (this *Impl) StopCloud(r *http.Request, in *StopCloudIn, out *StopCloudOut) error {
	err := this.Service.StopCloud(in.Name, in.UseKerberos, in.ApplicationID, in.Username, in.Keytab)
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
	it, err := this.Service.BuildAutoML(in.Address, in.Dataset, in.TargetName, in.MaxRunTime)
	if err != nil {
		return err
	}
	out.ModelID = it
	return nil
}

func (this *Impl) DeployPojo(r *http.Request, in *DeployPojoIn, out *DeployPojoOut) error {
	err := this.Service.DeployPojo(in.Address, in.JavaModelPath, in.GenModelPath)
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
