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
	"sync"
	"time"
)

type LdapUser struct {
	mu sync.RWMutex
	// users map[string]*Timers
	users map[string]*time.Timer

	IdlTime time.Duration
	MaxTime time.Duration
}

// type Timers struct {
// 	Idl *time.Timer
// 	Max *time.Timer
// }

// NewUser creates a new LdapUser with it's own "self-destruct" timer
func (u *LdapUser) NewUser(auth, user string) {
	u.mu.Lock()
	t := time.AfterFunc(u.MaxTime, func() { u.Delete(auth) })
	u.users[auth] = t
	u.mu.Unlock()
}

// Exists verifies if a user is in the Users map
func (u *LdapUser) Exists(auth string) bool {
	u.mu.RLock()
	_, ok := u.users[auth]
	// if ok {
	// 	if !user.Idl.Stop() {
	// 		<-user.Idl.C
	// 	}
	// 	user.Idl.Reset(u.IdlTime)
	// }
	u.mu.RUnlock()

	return ok
}

// Delete removes a user from the LdapUsers map and stop/flushes the timer
func (u *LdapUser) Delete(auth string) {
	u.mu.Lock()
	t, ok := u.users[auth]
	if ok {
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
