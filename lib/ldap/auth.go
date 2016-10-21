package ldap

import "github.com/abbot/go-http-auth"

type BasicLdapAuth struct {
	Realm string

	Address  string
	BindDn   string
	BindPass string

	// Headers used by authenticator. Set to ProxyHeaders to use with
	// proxy server. When nil, NormalHeaders are used.
	Headers *auth.Headers
}
