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
	"net/http"

	"github.com/abbot/go-http-auth"
	"github.com/h2oai/steam/lib/ldap"
	"github.com/h2oai/steam/master/data"
)

type BasicLdapAuthProvider struct {
	ds    *data.Datastore
	realm string

	conn *ldap.Ldap
}

func (p *BasicLdapAuthProvider) Secure(handler http.Handler) http.Handler {
	az := NewDefaultAz(p.ds)
	authenticator := ldap.NewBasicLdapAuth(az, az.directory, p.realm, p.conn, p.ds)

	return auth.JustCheck(authenticator, handler.ServeHTTP)
}

func NewBasicLdapAuthProvider(ds *data.Datastore, realm string, conn *ldap.Ldap) *BasicLdapAuthProvider {
	return &BasicLdapAuthProvider{ds: ds, realm: realm, conn: conn}
}

// Basic/Digest auth have no notion of logouts, so these handlers simply fail auth,
//  causing a 401 on the original realm, forcing the browser to re-auth.
func (p *BasicLdapAuthProvider) Logout() http.Handler {
	authenticator := auth.NewBasicAuthenticator(p.realm, authNoop)
	return auth.JustCheck(authenticator, serveNoop)
}
