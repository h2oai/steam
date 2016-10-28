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
