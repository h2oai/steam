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
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/h2oai/steam/master/data"

	"github.com/go-ldap/ldap"
	"github.com/pkg/errors"
)

type Ldap struct {
	Address  string
	BindDN   string
	BindPass string

	UserBaseDn      string
	UserBaseFilter  string
	UserRNAttribute string

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
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0,
		false,
		fmt.Sprintf("(&%s(%s=%s))",
			l.UserBaseFilter, l.UserRNAttribute, user),
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
	userBaseDn, userRNAttribute, userBaseFilter string,
	forceBind bool,
	idleTime, maxTime time.Duration) *Ldap {
	return &Ldap{
		// Base LDAP settings
		Address: address, BindDN: bindDn, BindPass: bindPass,
		// User filter settings
		UserBaseDn: userBaseDn, UserRNAttribute: userRNAttribute, UserBaseFilter: userBaseFilter,
		// Additional Configs
		ForceBind: forceBind,
		Users:     NewLdapUser(idleTime, maxTime),
	}
}

func FromConfig(ds *data.Datastore) (*Ldap, error) {
	config, exists, err := ds.ReadSecurity(data.ByKey("ldap"))
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
		a.Address,
		a.BindDN, a.BindPass,
		a.UserBaseDn, a.UserRNAttribute, a.UserBaseFilter,

		a.ForceBind, time.Minute*1, 60*2*time.Minute), nil
}
