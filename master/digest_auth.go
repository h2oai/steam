package master

import (
	"github.com/abbot/go-http-auth"
	"github.com/h2oai/steamY/master/az"
	"net/http"
)

type DigestAuthProvider struct {
	az    az.Az
	realm string
}

func (p *DigestAuthProvider) Secure(handler http.Handler) http.Handler {
	authenticator := auth.NewDigestAuthenticator(p.realm, func(user, realm string) string {
		return p.az.Authenticate(user)
	})
	return auth.JustCheck(authenticator, handler.ServeHTTP)
}

// Basic/Digest auth have no notion of logouts, so these handlers simply fail auth,
//  causing a 401 on the original realm, forcing the browser to re-auth.

func (p *DigestAuthProvider) Logout() http.Handler {
	authenticator := auth.NewDigestAuthenticator(p.realm, authNoop)
	return auth.JustCheck(authenticator, serveNoop)
}

func newDigestAuthProvider(az az.Az, realm string) AuthProvider {
	return &DigestAuthProvider{az, realm}
}
