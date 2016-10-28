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
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/context"

	auth "github.com/abbot/go-http-auth"
)

const contentType = "Content-Type"

type BasicLdapAuth struct {
	Realm string

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

	if a.Conn.Users.Exists(s[1]) && !a.Conn.ForceBind {
		return user
	}

	log.Println("LDAP", user, "checking bind")
	if err := a.Conn.CheckBind(user, password); err != nil {
		log.Println(err)
		return ""
	}
	a.Conn.Users.NewUser(s[1], user)

	return user
}

func NewBasicLdapAuth(realm string, conn *Ldap) *BasicLdapAuth {
	return &BasicLdapAuth{Realm: realm, Conn: conn}
}
