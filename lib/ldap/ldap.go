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
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/h2oai/steam/srv/web"

	"github.com/go-ldap/ldap"
	"github.com/pkg/errors"
)

type Ldap struct {
	Address  string
	BindDN   string
	BindPass string

	UserBaseDn        string // Location of LDAP users TODO: allow for multiple DNs
	UserBaseFilter    string // LDAP search filter e.g. (department=IT)
	UserNameAttribute string // Contains username (typically uid or sAMAccountName)

	GroupBaseDn             string   // Location of LDAP group TODO: allow for multiple DNs
	GroupNameAttribute      string   // The attribute that contains the name (typically cn)
	GroupNames              []string // The names of groups that have access to Steam
	StaticGroupSearchFilter string   // LDAP search filter e.g. (department=IT)
	StaticMemberAttribute   string   // The group member values (typically member or memberUid)

	SearchRequestSizeLimit int // The maximum number of entries request by LDAP searches
	SearchRequestTimeLimit int // The maximum number of seconds to wait for LDAP searches

	Ldaps     bool
	ForceBind bool

	tlsConfig *tls.Config
}

func (l *Ldap) TestConfig() (int, map[string]int, error) {
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
		return 0, nil, errors.Wrap(err, "dialing ldap")
	}
	defer conn.Close()
	if err := conn.Bind(l.BindDN, l.BindPass); err != nil {
		return 0, nil, errors.Wrap(err, "attempting bind")
	}

	userReq := ldap.NewSearchRequest(
		l.UserBaseDn, ldap.ScopeWholeSubtree, ldap.DerefAlways,
		l.SearchRequestSizeLimit, l.SearchRequestTimeLimit,
		false,
		"(objectClass=*)",
		[]string{l.UserNameAttribute}, nil,
	)

	q := joinOrQuery(l.GroupNameAttribute, l.GroupNames...)
	groupReq := ldap.NewSearchRequest(
		l.GroupBaseDn, ldap.ScopeWholeSubtree, ldap.DerefAlways,
		l.SearchRequestSizeLimit, l.SearchRequestTimeLimit,
		false,
		fmt.Sprintf("(%s)", q),
		[]string{l.StaticMemberAttribute, l.GroupNameAttribute}, nil,
	)

	userRes, err := conn.Search(userReq)
	if err != nil {
		return 0, nil, errors.Wrap(err, "searching for user base DN")
	}
	if len(userRes.Entries) < 1 {
		return 0, nil, errors.New(fmt.Sprint("unable to locate user base DN", l.UserBaseDn))
	}

	groupRes, err := conn.Search(groupReq)
	if err != nil {
		return 0, nil, errors.Wrap(err, "searching for groups")
	}
	if len(groupRes.Entries) < 1 {
		return 0, nil, errors.New(fmt.Sprint("unable to locate any matching groups", l.GroupBaseDn))
	}

	uCN := make(map[string]struct{})
	uDN := make(map[string]struct{})
	for _, u := range userRes.Entries {
		uCN[u.GetAttributeValue(l.UserNameAttribute)] = struct{}{}
		uDN[u.DN] = struct{}{}
	}

	var ct int
	groups := make(map[string]int)
	for _, e := range groupRes.Entries {
		// TODO: There may be a better implementation for this
		// Query is case insensitive... attribute keys are not. Additionally
		// the LDAP/AD key may not match the provided key.
		var aName string
		for _, a := range e.Attributes {
			if strings.ToLower(l.GroupNameAttribute) == strings.ToLower(a.Name) {
				aName = a.Name
			}
		}

		for _, member := range e.GetAttributeValues(l.StaticMemberAttribute) {
			if _, ok := uCN[member]; ok {
				ct++
				groups[e.GetAttributeValue(aName)]++
			} else if _, ok := uDN[member]; ok {
				ct++
				groups[e.GetAttributeValue(aName)]++
			}
		}
	}

	return ct, groups, nil
}

func joinOrQuery(nameAttribute string, names ...string) string {
	buf := new(bytes.Buffer)
	if len(names) > 1 {
		buf.WriteString("(|")
	}
	for _, name := range names {
		buf.WriteString(fmt.Sprintf("(%s=%s)", nameAttribute, name))
	}
	if len(names) > 1 {
		buf.WriteString(")")
	}

	return buf.String()
}

func (l *Ldap) checkGroup(conn *ldap.Conn, userCN, userDN string) (bool, error) {
	query := joinOrQuery(l.GroupNameAttribute, l.GroupNames...)
	req := ldap.NewSearchRequest(
		l.GroupBaseDn, ldap.ScopeWholeSubtree, ldap.DerefAlways,
		l.SearchRequestSizeLimit, l.SearchRequestTimeLimit,
		false,
		fmt.Sprintf("(%s)", query),
		[]string{l.StaticMemberAttribute}, nil,
	)

	res, err := conn.Search(req)
	if err != nil {
		return false, errors.Wrap(err, "searching ldap")
	}

	if len(res.Entries) < 1 {
		log.Println("no groups found matching query")
		return false, nil
	}

	// Search each group for at least one match with this user
	for _, entry := range res.Entries {
		users := entry.GetAttributeValues(l.StaticMemberAttribute)

		// builtin sort.Search requires that the array be sorted first

		// Sort.Search returns the smallest index at which userName is less the next value
		// This means you need to verify that users[i] == userName because some values may incorrectly
		// return true (i.e. "Jetty1", may return true for "Jetty")
		sort.Strings(users)
		if i := sort.SearchStrings(users, userDN); i != len(users) {
			if users[i] == userDN {
				return true, nil
			}
		}
		if i := sort.SearchStrings(users, userCN); i != len(users) {
			if users[i] == userCN {
				return true, nil
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
	// If a username is entered this is not an Anonymous bind >> Check password
	if strings.TrimSpace(password) == "" {
		return errors.New("no password provided. Please enter a valid password")
	}

	userDn := res.Entries[0].DN
	if ok, err := l.checkGroup(conn, user, userDn); err != nil {
		return errors.Wrap(err, "checking valid groups")
	} else if !ok {
		return errors.New("LDAP user has no valid Steam permissions")
	}

	// Verify user Bind
	return errors.Wrapf(conn.Bind(userDn, password), "user %s binding to ldap", user)
}

func NewLdap(
	// Connection settings
	address, bindDn, bindPass string, ldaps, forceBind bool,
	// User settings
	userBaseDn, userBaseFilter, userNameAttribute string,
	// Group settings
	groupBaseDn, groupNameAttribute, staticMemberAttribute string, groupNames []string,
	// Advanced Settings
	searchRequestSizeLimit, searchRequestTimeLimit int,
	// TLS Settings
	tlsConfig *tls.Config,
) *Ldap {
	return &Ldap{
		// Connection settings
		Address: address, BindDN: bindDn, BindPass: bindPass, Ldaps: ldaps, ForceBind: forceBind,
		// User settings
		UserBaseDn: userBaseDn, UserNameAttribute: userNameAttribute, UserBaseFilter: userBaseFilter,
		// Group settings
		GroupBaseDn: groupBaseDn, GroupNameAttribute: groupNameAttribute, StaticMemberAttribute: staticMemberAttribute, GroupNames: groupNames,
		// Advanced settings
		SearchRequestSizeLimit: searchRequestSizeLimit, SearchRequestTimeLimit: searchRequestTimeLimit,
		// TLS Settings
		tlsConfig: tlsConfig,
	}
}

func FromConfig(config *web.LdapConfig, tlsConfig *tls.Config) *Ldap {
	groupNames := strings.Split(config.GroupNames, ",")

	return NewLdap(
		fmt.Sprintf("%s:%d", config.Host, config.Port), config.BindDn, config.BindPassword, config.Ldaps, config.ForceBind,
		config.UserBaseDn, config.UserBaseFilter, config.UserNameAttribute,
		config.GroupBaseDn, config.GroupNameAttribute, config.StaticMemberAttribute, groupNames,
		config.SearchRequestSizeLimit, config.SearchRequestTimeLimit,
		tlsConfig,
	)
}

func FromDatabase(config string, tlsConfig *tls.Config) (*Ldap, error) {
	aux := struct {
		Bind string
		Ldap
	}{}
	if err := json.Unmarshal([]byte(config), &aux); err != nil {
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
		a.GroupBaseDn, a.GroupNameAttribute, a.StaticMemberAttribute, a.GroupNames,
		a.SearchRequestSizeLimit, a.SearchRequestTimeLimit,
		tlsConfig,
	), nil
}

func CreateTLSConfig(certFilePath, keyFilePath string) (*tls.Config, error) {
	// FIXME: should NOT continue to remove verify
	tlsConfig := &tls.Config{InsecureSkipVerify: true}

	// cert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	// if err != nil {
	// 	return nil, err
	// }

	// tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
	// tlsConfig.BuildNameToCertificate()
	return tlsConfig, nil
}
