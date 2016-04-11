// ----------------------------------
// --- Generated with go:generate ---
// ---        DO NOT EDIT         ---
// ----------------------------------

package web

import (
	"net/http"
)

// --- Interfaces ---

type Service interface {
	Ping(status bool) (bool, error)
}

// --- Messages ---

type PingIn struct {
	Status bool `json:"status"`
}

type PingOut struct {
	Status bool `json:"status"`
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
