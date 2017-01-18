/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package master

import (
	"fmt"
	"log"
	"net/http"

	"github.com/h2oai/steam/master/az"

	auth "github.com/abbot/go-http-auth"
)

type DefaultAz struct {
	directory az.Directory
}

func NewDefaultAz(directory az.Directory) *DefaultAz {
	return &DefaultAz{directory}
}

// Authenticate either returns a local user with password, or else it validates
// a user against an external authentication provider
// Returns IsLocalUser, DBUsername, DBPassword
func (a *DefaultAz) Authenticate(username, password, token string) (bool, string, string) {
	// Looking and/or validate against an external authentication provider
	pz, err := a.directory.Lookup(username, password, token)
	if err != nil {
		log.Printf("User %s read failed: %s\n", username, err)
		return false, "", ""
	}
	if pz == nil {
		log.Printf("User %s does not exist\n", username)
		return false, "", ""
	}
	// If local, return true for local and the password to perform basic auth
	if pz.IsLocal() {
		return true, "", pz.Password()
	}
	// Else return false and the name (Name should be not "" if properly validated)
	return false, pz.Name(), ""
}

func (a *DefaultAz) Identify(r *http.Request) (az.Principal, error) {
	username := r.Header.Get(auth.AuthUsernameHeader)
	pz, err := a.directory.LookupUser(username)
	if err != nil {
		return nil, err
	}

	if pz == nil {
		return nil, fmt.Errorf("User %s does not exist\n", username)
	}

	return pz, nil
}

func serveNoop(w http.ResponseWriter, r *http.Request) {}
func authNoop(user, realm string) string               { return "" }
