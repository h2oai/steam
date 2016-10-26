package ldap

import (
	"net/http"

	"golang.org/x/net/context"

	auth "github.com/abbot/go-http-auth"
)

type BasicLdapAuth struct {
	Realm string

	Address  string
	BindDn   string
	BindPass string

	// Headers used by authenticator. Set to ProxyHeaders to use with
	// proxy server. When nil, NormalHeaders are used.
	Headers *auth.Headers
}

func (a *BasicLdapAuth) Wrap(auth.AuthenticatedHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if username := "/* VERIFY VALID USER */"; username == "" {
			/* RETURN BASIC AUTH REQUIRED HEADERS */
		} else {
			ar := *auth.AuthenticatedRequest{Request: *r, Username: username}
			wrapped(a, ar)
		}
	}
}

func (a *BasicLdapAuth) NewContext(ctx context.Context, r *http.Request) context.Context {
}
