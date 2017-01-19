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
	"sync"
	"time"

	"github.com/h2oai/steam/lib/ldap"
	"github.com/pkg/errors"
)

type userTable struct {
	mu sync.RWMutex
	// Logged in users
	users map[string]*time.Timer

	MaxTime time.Duration
}

func (ds *Datastore) NewUser(identity Identity, exists bool, username, password, token string, conn *ldap.Ldap) (Identity, bool, error) {
	ds.users.mu.Lock()
	defer ds.users.mu.Unlock()

	// If user already exists, Noop
	if _, ok := ds.users.users[token]; ok {
		return identity, exists, nil
	}

	log.Println("LDAP", username, "Checking bind")
	if err := conn.CheckBind(username, password); err != nil {
		return Identity{}, false, errors.Wrap(err, "checking user bind")
	}
	log.Println("LDAP", username, "bind success")

	// Self deleting timer after max duration
	t := time.AfterFunc(ds.users.MaxTime, func() {
		ds.users.mu.Lock()
		delete(ds.users.users, token)
		ds.users.mu.Unlock()
	})
	ds.users.users[token] = t

	return ds.fetchLdapIdentity(identity, exists, username, password, token)
}

func (u *userTable) Exists(token string) bool {
	u.mu.RLock()
	_, ok := u.users[token]
	u.mu.RUnlock()

	return ok
}

func newUserTable(maxTime time.Duration) *userTable {
	return &userTable{
		users:   make(map[string]*time.Timer),
		MaxTime: maxTime,
	}
}

func (ds *Datastore) fetchLdapIdentity(identity Identity, exists bool, username, password, token string) (Identity, bool, error) {
	// Fetch valid roles
	// FIXME: THIS SHOULD INHERIT ROLES FROM LDAP
	role, _, err := ds.ReadRole(ByName(standard_user))
	if err != nil {
		return Identity{}, false, errors.Wrap(err, "reading roles from database")
	}

	// At this point the user is alread authenticated; so if they don't exists
	// add them to the DB instead of bailing
	if !exists {
		id, err := ds.CreateIdentity(
			username,
			WithDefaultIdentityWorkgroup, LinkRole(role.Id, false),
			WithAuthType(LDAPAuth),
		)
		if err != nil {
			return Identity{}, false, errors.Wrap(err, "creating identity")
		}
		identity, exists, err = ds.ReadIdentity(ById(id))
		if err != nil {
			return Identity{}, false, errors.Wrap(err, "reading identity")
		}
	} else {
		if err := ds.UpdateIdentity(identity.Id, LinkRole(role.Id, true)); err != nil {
			return Identity{}, false, errors.Wrap(err, "adding roles to identity")
		}
	}

	return identity, exists, nil
}
