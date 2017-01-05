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
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/h2oai/steam/lib/fs"

	"github.com/BurntSushi/toml"
	"github.com/go-ldap/ldap"
	"github.com/pkg/errors"
)

type Ldap struct {
	Address  string
	BindDN   string
	BindPass string `toml:"bindPassword"`

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
	forceBind bool,
	idleTime, maxTime time.Duration) *Ldap {
	return &Ldap{
		// Base LDAP settings
		Address: address, BindDN: bindDn, BindPass: bindPass,
		// User filter settings
		UserBaseDn: userBaseDn, UserIdAttribute: userIdAttribute, UserObjectClass: userObjectClass,
		// Additional Configs
		ForceBind: forceBind,
		Users:     NewLdapUser(idleTime, maxTime),
	}
}

func FromConfig(workingDir string) (*Ldap, error) {
	A := struct {
		Hostname        string
		Port            int
		BindDn          string
		BindPassword    string
		UserBaseDn      string
		UserIdAttribute string
		UserObjectClass string

		IsTLS     bool `toml:"useLdaps"`
		ForceBind bool
		IdleTime  time.Duration
		MaxTime   time.Duration
	}{}

	if _, err := os.Stat(fs.GetDBPath(workingDir, "ldap.conf")); os.IsNotExist(err) {
		return nil, errors.New("no LDAP config file provided. Please set up LDAP settings in local client login mode")
	}

	f, err := filepath.Abs(fs.GetDBPath(workingDir, "ldap.conf"))
	if err != nil {
		return nil, errors.Wrap(err, "getting absolute path")
	}
	if _, err := toml.DecodeFile(f, &A); err != nil {
		return nil, errors.Wrap(err, "decoding config file")
	}

	return NewLdap(
		fmt.Sprintf("%s:%d", A.Hostname, A.Port),
		A.BindDn, A.BindPassword,
		A.UserBaseDn, A.UserIdAttribute, A.UserObjectClass,

		A.ForceBind, time.Minute*A.IdleTime, time.Minute*A.MaxTime), nil
}
