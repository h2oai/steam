package ldap

import (
	"fmt"

	"github.com/go-ldap/ldap"
	"github.com/pkg/errors"
)

type Ldap struct {
	Conn *ldap.Conn

	userBaseDn      string
	userIdAttribute string
	userObjectClass string
}

func (l *Ldap) FindDN(user string) (string, error) {
	//
	req := ldap.NewSearchRequest(l.userBaseDn, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("&(&(objectClass=%s)({%s}={%s})",
			l.userObjectClass, l.userIdAttribute, user),
		nil, nil,
	)

	res, err := l.Conn.Search(req)
	if err != nil {
		return "", errors.Wrap(err, "searching ldap")
	}

	if len(res.Entries) < 1 {
		return "", fmt.Errorf("user %s does not exist", user)
	} else if len(res.Entries) > 1 {
		return "", fmt.Errorf("too many user entries")
	}

	return res.Entries[0].DN, nil
}

func (l *Ldap) Close() {
	l.Conn.Close()
}
