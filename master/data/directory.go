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

package data

import (
	"crypto/tls"
	"log"

	"github.com/h2oai/steam/lib/ldap"
	"github.com/h2oai/steam/master/az"

	"fmt"

	"github.com/pkg/errors"
)

// --- Datastore-backed Directory Impl ---

func (ds *Datastore) Lookup(username, password, token string, tlsConfig *tls.Config) (az.Principal, error) {
	// Fetch identity
	identity, exists, err := ds.ReadIdentity(ByName(username))
	if err != nil {
		return nil, errors.Wrap(err, "reading identity")
	}

	auth, ok, err := ds.ReadAuthentication(ByEnabled)
	if err != nil {
		return nil, errors.Wrap(err, "reading authentication config")
	} else if !ok {
		return ds.localLookup(identity, exists)
	}

	switch auth.Key {
	case LDAPAuth:
		conn, err := ldap.FromDatabase(auth.Value, tlsConfig)
		if err != nil {
			log.Printf("LDAP ERROR creating ldap config from database: %v", err)
			log.Println("LDAP ERROR resorting to local auth")
			return ds.localLookup(identity, exists)
		}
		return ds.ldapLookup(identity, exists, username, password, token, conn)
	}
	return nil, nil
}

func (ds *Datastore) LookupUser(name string) (az.Principal, error) {
	// Fetch identity
	identity, ok, err := ds.ReadIdentity(ByName(name))
	if err != nil {
		return nil, errors.Wrap(err, "reading identity")
	}
	if !ok {
		return nil, nil
	}

	return ds.localLookup(identity, ok)
}

func (ds *Datastore) localLookup(identity Identity, exists bool) (az.Principal, error) {
	// Validate that this identity exists
	if !exists {
		return nil, errors.New("unable to locate user")
	}
	// Validate that this is a local user
	if identity.AuthType != LocalAuth {
		return nil, fmt.Errorf("%s is not a local user", identity.Name)
	}

	return ds.lookup(identity)
}

func (ds *Datastore) ldapLookup(identity Identity, exists bool, username, password, token string, conn *ldap.Ldap) (az.Principal, error) {
	// Validate if local
	if identity.AuthType == LocalAuth || ds.users.Exists(token) {
		return ds.lookup(identity)
	}

	identity, exists, err := ds.NewUser(identity, exists, username, password, token, conn)
	if err != nil {
		return nil, errors.Wrap(err, "creating new user")
	}

	return ds.lookup(identity)
}

func (ds *Datastore) lookup(identity Identity) (az.Principal, error) {
	// Fetch roles
	roles, err := ds.ReadRoles(ForIdentity(identity.Id))
	if err != nil {
		return nil, errors.Wrap(err, "reading roles")
	}

	isAdmin := false
	for _, role := range roles {
		if role.Name == AdminRN {
			isAdmin = true
			break
		}
	}

	perms, err := ds.ReadPermissions(
		ForIdentity(identity.Id),
	)
	if err != nil {
		return nil, err
	}

	permissions := make(map[int64]bool)
	for _, perm := range perms {
		permissions[perm.Id] = true
	}

	return &Principal{ds, &identity, permissions, isAdmin}, nil
}
