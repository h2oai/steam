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
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/h2oai/steam/master/data"
	"github.com/h2oai/steam/srv/web"

	"github.com/pkg/errors"
	ldap "gopkg.in/ldap.v2"
)

type Ldap struct {
	Address  string
	BindDN   string
	BindPass string

	UserBaseDn        string // Location of LDAP users TODO: allow for multiple DNs
	UserBaseFilter    string // LDAP search filter e.g. (department=IT)
	UserNameAttribute string // Contains username (typically uid or sAMAccountName)

	GroupDn                 string // The group to allow users from //XXX TEMPORARY
	GroupBaseDn             string // Location of LDAP group TODO: allow for multiple DNs
	GroupNameAttribute      string // The attribute that contains the name (typically cn)
	StaticGroupSearchFilter string // LDAP search filter e.g. (department=IT)
	StaticMemberAttribute   string // The group member values (typically member or memberUid)

	SearchRequestSizeLimit int // The maximum number of entries request by LDAP searches
	SearchRequestTimeLimit int // The maximum number of seconds to wait for LDAP searches

	Ldaps     bool
	ForceBind bool

	tlsConfig *tls.Config

	// Users who are logged in
	Users *LdapUser
}

func (l *Ldap) Test() error {
	var (
		conn *ldap.Conn
		err  error
	)

	if l.Ldaps {
		conn, err = ldap.DialTLS("tcp", l.Address, l.tlsConfig)
	} else {
		conn, err = ldap.Dial("tcp", l.Address)
	}
	if err != nil {
		return errors.Wrap(err, "dialing ldap")
	}
	defer conn.Close()
	if err := conn.Bind(l.BindDN, l.BindPass); err != nil {
		return errors.Wrap(err, "attempting bind")
	}

	req := ldap.NewSearchRequest(
		l.GroupDn, ldap.ScopeBaseObject, ldap.DerefAlways,
		l.SearchRequestSizeLimit, l.SearchRequestTimeLimit,
		false,
		"(objectClass=*)",
		[]string{l.StaticMemberAttribute}, nil,
	)

	res, err := conn.Search(req)
	if err != nil {
		return errors.Wrap(err, "searching for group")
	}
	if len(res.Entries) < 1 {
		return errors.New(fmt.Sprint("unable to locate group", l.GroupDn))
	} else if len(res.Entries) > 2 {
		return errors.New("too many group entries")
	}

	return nil
}

func (l *Ldap) checkGroup(conn *ldap.Conn, user string) (bool, error) {
	req := ldap.NewSearchRequest(
		l.GroupDn, ldap.ScopeBaseObject, ldap.DerefAlways,
		l.SearchRequestSizeLimit, l.SearchRequestTimeLimit,
		false,
		"(objectClass=*)",
		[]string{l.StaticMemberAttribute}, nil,
	)

	res, err := conn.Search(req)
	if err != nil {
		return false, errors.Wrap(err, "searching ldap")
	}

	users := res.Entries[0].GetAttributeValues(l.StaticMemberAttribute)
	sort.Strings(users)
	if i := sort.SearchStrings(users, user); i != len(users) {
		for _, u := range users[i:] {
			if u == user {
				return true, nil
			}
			if !strings.HasPrefix(u, user) {
				break
			}
		}
	}

	return false, nil
}

func (l *Ldap) CheckBind(user, password string) error {
	// Make connection to LDAP with read-only user
	var (
		conn *ldap.Conn
		err  error
	)

	if l.Ldaps {
		conn, err = ldap.DialTLS("tcp", l.Address, l.tlsConfig)
	} else {
		conn, err = ldap.Dial("tcp", l.Address)
	}
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
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0,
		false,
		fmt.Sprintf("(&%s(%s=%s))",
			l.UserBaseFilter, l.UserNameAttribute, user),
		nil, nil,
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

	if ok, err := l.checkGroup(conn, user); err != nil {
		return errors.Wrap(err, "checking valid groups")
	} else if !ok {
		return errors.New("LDAP user has no valid Steam permissions")
	}

	userDn := res.Entries[0].DN

	// Verify user Bind
	return errors.Wrapf(conn.Bind(userDn, password), "user %s binding to ldap", user)
}

func NewLdap(
	// Connection settings
	address, bindDn, bindPass string, ldaps, forceBind bool,
	// User settings
	userBaseDn, userBaseFilter, userNameAttribute string,
	// Group settings
	groupDN, staticMemberAttribute string,
	// Advanced Settings
	searchRequestSizeLimit, searchRequestTimeLimit int,
) *Ldap {
	return &Ldap{
		// Connection settings
		Address: address, BindDN: bindDn, BindPass: bindPass, Ldaps: ldaps, ForceBind: forceBind,
		// User settings
		UserBaseDn: userBaseDn, UserNameAttribute: userNameAttribute, UserBaseFilter: userBaseFilter,
		// Group Settings
		GroupDn: groupDN, StaticMemberAttribute: staticMemberAttribute,
		// Additional Configs
		Users: NewLdapUser(time.Minute*1, 2*time.Hour),
	}
}

func FromConfig(config *web.LdapConfig) *Ldap {
	return NewLdap(
		fmt.Sprintf("%s:%d", config.Host, config.Port), config.BindDn, config.BindPassword, config.Ldaps, config.ForceBind,
		config.UserBaseDn, config.UserBaseFilter, config.UserNameAttribute,
		config.GroupDn, config.StaticMemberAttribute,
		0, 0,
	)
}

func FromDatabase(ds *data.Datastore) (*Ldap, error) {
	config, exists, err := ds.ReadAuthentication(data.ByKey("ldap"))
	if err != nil {
		return nil, errors.Wrap(err, "reading security config from database")
	} else if !exists {
		return nil, errors.New("no valid LDAP configurations. Please set LDAP configurations prior to starting steam with LDAP")
	}

	aux := struct {
		Bind string
		Ldap
	}{}
	if err := json.Unmarshal([]byte(config.Value), &aux); err != nil {
		return nil, errors.Wrap(err, "deserializing config")
	}

	b, err := base64.StdEncoding.DecodeString(aux.Bind)
	if err != nil {
		return nil, errors.Wrap(err, "decoding bind")
	}

	decrypt := strings.Split(string(b), ":")
	aux.Ldap.BindDN, aux.Ldap.BindPass = decrypt[0], decrypt[1]

	a := aux.Ldap

	return NewLdap(
		a.Address, a.BindDN, a.BindPass, a.Ldaps, a.ForceBind,
		a.UserBaseDn, a.UserBaseFilter, a.UserNameAttribute,
		a.GroupDn, a.StaticMemberAttribute,
		0, 0,
	), nil
}
