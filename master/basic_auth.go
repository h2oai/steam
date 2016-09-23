package master

import (
	"github.com/abbot/go-http-auth"
	"github.com/h2oai/steam/master/az"
	"net/http"
)

type BasicAuthProvider struct {
	az    az.Az
	realm string
}

func (p *BasicAuthProvider) Secure(handler http.Handler) http.Handler {
	authenticator := auth.NewBasicAuthenticator(p.realm, func(user, realm string) string {
		return p.az.Authenticate(user)
	})
	return auth.JustCheck(authenticator, handler.ServeHTTP)
}

// Basic/Digest auth have no notion of logouts, so these handlers simply fail auth,
//  causing a 401 on the original realm, forcing the browser to re-auth.

func (p *BasicAuthProvider) Logout() http.Handler {
	authenticator := auth.NewBasicAuthenticator(p.realm, authNoop)
	return auth.JustCheck(authenticator, serveNoop)
}

func newBasicAuthProvider(az az.Az, realm string) AuthProvider {
	return &BasicAuthProvider{az, realm}
}
