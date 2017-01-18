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
	"log"

	"github.com/h2oai/steam/lib/ldap"
	"github.com/h2oai/steam/master/az"
	"github.com/pkg/errors"
)

// --- Datastore-backed Directory Impl ---

func (ds *Datastore) Lookup(username, password, token string) (az.Principal, error) {
	auth, ok, err := ds.ReadAuthentication(ByEnabled)
	if err != nil {
		return nil, errors.Wrap(err, "reading authentication config")
	} else if !ok {
		return ds.LookupUser(username)
	}

	exists := ds.users.Exists(token)
	if exists {
		return ds.LookupUser(username)
	}
	switch auth.Key {
	case LDAPAuth:
		conn, err := ldap.FromDatabase(auth.Value)
		if err != nil {
			return nil, errors.Wrap(err, "creating ldap config from database")
		}
		return ds.ldapLookup(username, password, token, conn)
	}
	return ds.LookupUser(username)
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

func (ds *Datastore) ldapLookup(username, password, token string, conn *ldap.Ldap) (az.Principal, error) {
	// Fetch identity
	identity, userExists, err := ds.ReadIdentity(ByName(username))
	if err != nil {
		return nil, errors.Wrap(err, "reading identity")
	}
	if identity.AuthType == LocalAuth {
		return ds.LookupUser(username)
	}

	log.Println("LDAP", username, "Checking bind")
	if err := conn.CheckBind(username, password); err != nil {
		return nil, errors.Wrap(err, "checking user bind")
	}
	log.Println("LDAP", username, "bind success")

	ds.users.NewUser(token)

	// Fetch valid roles
	// FIXME: THIS SHOULD INHERIT ROLES FROM LDAP
	role, _, err := ds.ReadRole(ByName(standard_user))
	if err != nil {
		return nil, errors.Wrap(err, "reading roles from database")
	}

	if !userExists { // at this point already validated so add instead of bail
		id, err := ds.CreateIdentity(
			username,
			WithDefaultIdentityWorkgroup, LinkRole(role.Id, false),
			WithAuthType(LDAPAuth),
		)
		if err != nil {
			return nil, errors.Wrap(err, "creating identity")
		}
		identity, _, err = ds.ReadIdentity(ById(id))
		if err != nil {
			return nil, errors.Wrap(err, "reading identity")
		}
	} else {
		if err := ds.UpdateIdentity(identity.Id, LinkRole(role.Id, true)); err != nil {
			return nil, errors.Wrap(err, "adding roles to identity")
		}
	}

	perms, err := ds.ReadPermissions(ForIdentity(identity.Id))
	if err != nil {
		return nil, errors.Wrap(err, "reading permissions")
	}

	permissions := make(map[int64]bool)
	for _, perm := range perms {
		permissions[perm.Id] = true
	}

	return &Principal{ds, &identity, permissions, false}, nil
}
