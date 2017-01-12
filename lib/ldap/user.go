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
	"log"
	"sync"
	"time"

	"github.com/h2oai/steam/master/data"
)

type LdapUser struct {
	mu sync.RWMutex
	// users map[string]*Timers
	users map[string]*time.Timer

	IdlTime time.Duration
	MaxTime time.Duration
}

// NewUser creates a new LdapUser with it's own "self-destruct" timer
func (u *LdapUser) NewUser(auth, user, password string, conn *Ldap, ds *data.Datastore) string {
	u.mu.Lock()
	defer u.mu.Unlock()

	// Verify that user does not exist before continuing
	if _, ok := u.users[auth]; ok {
		return user
	}

	log.Println("LDAP", user, "checking bind")
	if err := conn.CheckBind(user, password); err != nil {
		log.Println(err)
		return ""
	}

	if ds == nil {
		log.Println("Failed to load datastore")
		return ""
	}

	// FIXME: THIS SHOULD INHERIT ROLES FROM LDAP
	role, _, err := ds.ReadRole(data.ByName(data.STANDARD_USER))
	if err != nil {
		log.Println("Mapping roles: reading role from database", err)
		return ""
	}

	if identity, exists, err := ds.ReadIdentity(data.ByName(user)); err != nil {
		log.Println("Mapping roles: reading identity from database", err)
		return ""
	} else if exists {
		log.Println("LDAP", user, "updating roles")
		if err := ds.UpdateIdentity(identity.Id, data.LinkRole(role.Id, true)); err != nil {
			log.Println("Mapping roles: adding roles to identity", err)
			return ""
		}
	} else {
		log.Println("LDAP", user, "initializing user")
		if _, err := ds.CreateIdentity(user,
			data.WithDefaultIdentityWorkgroup, data.LinkRole(role.Id, false),
			data.WithAuthType(data.LDAPAuth),
		); err != nil {
			log.Println("Mapping roles: initializing identity with roles", err)
			return ""
		}
	}
	t := time.AfterFunc(u.MaxTime, func() {
		u.mu.Lock()
		delete(u.users, auth)
		u.mu.Unlock()
	})
	u.users[auth] = t

	return user
}

// Exists verifies if a user is in the Users map
func (u *LdapUser) Exists(auth string) bool {
	u.mu.RLock()
	_, ok := u.users[auth]
	u.mu.RUnlock()

	return ok
}

// Delete removes a user from the LdapUsers map and stop/flushes the timer
func (u *LdapUser) Delete(auth string) {
	u.mu.Lock()
	if t, ok := u.users[auth]; ok {
		if !t.Stop() {
			<-t.C
		}
		delete(u.users, auth)
	}
	u.mu.Unlock()
}

func NewLdapUser(idleTime, maxTime time.Duration) *LdapUser {
	return &LdapUser{
		// users:   make(map[string]*Timers),
		users:   make(map[string]*time.Timer),
		IdlTime: idleTime,
		MaxTime: maxTime,
	}
}
