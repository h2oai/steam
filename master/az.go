package master

import (
	"github.com/abbot/go-http-auth"
	"github.com/h2oai/steamY/master/az"
	"github.com/h2oai/steamY/master/data"
	"log"
	"net/http"
)

type DefaultAz struct {
	ds *data.Datastore
	pz az.Principal
}

func newDefaultAz(ds *data.Datastore, pz az.Principal) *DefaultAz {
	return &DefaultAz{ds, pz}
}

func (a *DefaultAz) Authenticate(username string) string {
	pz, err := a.ds.NewPrincipal(username)
	if err != nil {
		log.Printf("User %s read failed: %s\n", username, err)
		return ""
	}
	log.Println("User logged in:", username)
	return pz.Password()
}

func (a *DefaultAz) Identify(r *http.Request) (az.Principal, error) {
	username := r.Header.Get(auth.AuthUsernameHeader)
	log.Println("User identified:", username)
	return a.ds.NewPrincipal(username)
}
