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
package ldap

import (
	"encoding/base64"
	"net/http"
	"strings"

	"golang.org/x/net/context"

	auth "github.com/abbot/go-http-auth"
	"github.com/h2oai/steam/master/az"
)

const contentType = "Content-Type"

type BasicLdapAuth struct {
	az        az.Az
	Directory az.Directory
	Realm     string

	// Headers used by authenticator. Set to ProxyHeaders to use with
	// proxy server. When nil, NormalHeaders are used.
	Headers *auth.Headers

	Conn *Ldap
}

// Ask for BasicAuth headers
func (a *BasicLdapAuth) RequireAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, a.Headers.V().UnauthContentType)
	w.Header().Set(a.Headers.V().Authenticate, `Basic realm="`+a.Realm+`"`)
	w.WriteHeader(a.Headers.V().UnauthCode)
	w.Write([]byte(a.Headers.V().UnauthResponse))
}

func (a *BasicLdapAuth) Wrap(wrapped auth.AuthenticatedHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if username := a.CheckAuth(r); username == "" {
			a.RequireAuth(w, r)
		} else {
			ar := &auth.AuthenticatedRequest{Request: *r, Username: username}
			wrapped(w, ar)
		}
	}
}

func (a *BasicLdapAuth) NewContext(ctx context.Context, r *http.Request) context.Context {
	info := &auth.Info{Username: a.CheckAuth(r), ResponseHeaders: make(http.Header)}
	info.Authenticated = (info.Username != "")
	if !info.Authenticated {
		info.ResponseHeaders.Set(a.Headers.V().Authenticate, `Basic realm="`+a.Realm+`"`)
	}
	return context.WithValue(ctx, 0, info)
}

// CheckAuth verifies the user is authenticated by either finding them in the
// LDAP config map or binding the LDAP
func (a *BasicLdapAuth) CheckAuth(r *http.Request) string {
	s := strings.SplitN(r.Header.Get(a.Headers.V().Authorization), " ", 2)
	if len(s) != 2 || s[0] != "Basic" {
		return ""
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return ""
	}
	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return ""
	}
	user, password := pair[0], pair[1]
	if strings.TrimSpace(password) == "" {
		return ""
	}

	// First lookup in local datastore
	pz, err := a.Directory.Lookup(user)
	if err != nil {
		return ""
	}
	// If found, validate user against local datastore
	if pz != nil { // TODO: This should check if user is an LDAP user or local
		return auth.NewBasicAuthenticator(a.Realm, func(user, realm string) string { return a.az.Authenticate(user) }).CheckAuth(r)
	}

	// If not local, continue to check in LDAP for user
	if a.Conn.Users.Exists(s[1]) {
		return user
	}

	return a.Conn.Users.NewUser(s[1], user, password, a.Conn)
}

func NewBasicLdapAuth(az az.Az, directory az.Directory, realm string, conn *Ldap) *BasicLdapAuth {
	return &BasicLdapAuth{az: az, Directory: directory, Realm: realm, Conn: conn}
}
