package master

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/h2oai/steam/master/az"

	auth "github.com/abbot/go-http-auth"
)

type MultiAuthProvider struct {
	az    az.Az
	realm string
}

func (p *MultiAuthProvider) Secure(handler http.Handler) http.Handler {
	authenticator := NewMultiAuthenticator(p.az, p.realm)

	return auth.JustCheck(authenticator, handler.ServeHTTP)
}

// Basic/Digest auth have no notion of logouts, so these handlers simply fail auth,
//  causing a 401 on the original realm, forcing the browser to re-auth.

func (p *MultiAuthProvider) Logout() http.Handler {
	authenticator := auth.NewBasicAuthenticator(p.realm, authNoop)
	return auth.JustCheck(authenticator, serveNoop)
}

func newMultiAuthProvider(az az.Az, realm string) AuthProvider {
	return &MultiAuthProvider{az, realm}
}

type MultiAuth struct {
	az az.Az
	auth.BasicAuth
}

func NewMultiAuthenticator(az az.Az, realm string) *MultiAuth {
	basic := auth.NewBasicAuthenticator(realm, func(string, string) string { return "" })

	return &MultiAuth{BasicAuth: *basic, az: az}
}

func (a *MultiAuth) Wrap(wrapped auth.AuthenticatedHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if username := a.CheckAuth(r); username == "" {
			a.RequireAuth(w, r)
		} else {
			ar := &auth.AuthenticatedRequest{Request: *r, Username: username}
			wrapped(w, ar)
		}
	}
}

func (a *MultiAuth) CheckAuth(r *http.Request) string {
	token, user, pass := a.parseBasic(r)
	if user == "" {
		return ""
	}

	local, dbUser, dbPassword := a.az.Authenticate(user, pass, token)
	if local { // In the case of a local user, run simple basic auth
		return auth.NewBasicAuthenticator(a.Realm, a.secret(dbPassword)).CheckAuth(r)
	}
	// If not a local user, dbUser must have been set
	return dbUser
}

func (a *MultiAuth) parseBasic(r *http.Request) (string, string, string) {
	s := strings.SplitN(r.Header.Get(a.Headers.V().Authorization), " ", 2)
	if len(s) != 2 || s[0] != "Basic" {
		return "", "", ""
	}
	basic := s[1]

	b, err := base64.StdEncoding.DecodeString(basic)
	if err != nil {
		return "", "", ""
	}
	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return "", "", ""
	}
	// BasicAuth, Username, Password
	return basic, pair[0], pair[1]
}

func (a *MultiAuth) secret(password string) auth.SecretProvider {
	return func(user, realm string) string {
		return password
	}
}
