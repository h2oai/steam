package ldap

import (
	"fmt"
	"time"

	"github.com/go-ldap/ldap"
	"github.com/pkg/errors"
)

type Ldap struct {
	Address  string
	BindDN   string
	BindPass string

	UserBaseDn      string
	UserIdAttribute string
	UserObjectClass string

	// TODO implement TLS case
	isTLS     bool
	ForceBind bool

	// Users who are logged in
	Users *LdapUser
}

func (l *Ldap) CheckBind(user, password string) error {
	// Make connection to LDAP with read-only user
	conn, err := ldap.Dial("tcp", l.Address)
	if err != nil {
		return errors.Wrap(err, "dialing ldap")
	}
	defer conn.Close()
	if err := conn.Bind(l.BindDN, l.BindPass); err != nil {
		return errors.Wrap(err, "read user binding to ldap")
	}

	// Search request for userDN
	req := ldap.NewSearchRequest(
		l.UserBaseDn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=%s)(%s=%s))",
			l.UserObjectClass, l.UserIdAttribute, user),
		nil,
		nil,
	)
	res, err := conn.Search(req)
	if err != nil {
		return errors.Wrap(err, "searching ldap")
	}
	if len(res.Entries) < 1 {
		return fmt.Errorf("user %s does not exist", user)
	} else if len(res.Entries) > 1 {
		return fmt.Errorf("too many user entries")
	}
	userDn := res.Entries[0].DN

	// Verify user Bind
	return errors.Wrapf(conn.Bind(userDn, password), "user %s binding to ldap", user)
}

func NewLdap(
	address, bindDn, bindPass string,
	userBaseDn, userIdAttribute, userObjectClass string,
	idleTime, maxTime time.Duration) *Ldap {
	return &Ldap{
		Address: address, BindDN: bindDn, BindPass: bindPass,

		UserBaseDn: userBaseDn, UserIdAttribute: userIdAttribute, UserObjectClass: userObjectClass,

		Users: NewLdapUser(idleTime, maxTime),
	}
}
