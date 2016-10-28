package ldap

import (
	"log"
	"sync"
	"time"
)

type LdapUser struct {
	mu    sync.RWMutex
	users map[string]*Timers

	IdlTime time.Duration
	MaxTime time.Duration
}

type Timers struct {
	Idl *time.Timer
	Max *time.Timer
}

func (u *LdapUser) NewUser(auth, user string) {
	u.mu.Lock()
	t := &Timers{
		Idl: time.NewTimer(u.IdlTime),
		Max: time.NewTimer(u.MaxTime),
	}
	u.users[auth] = t
	u.mu.Unlock()

	// Launches delete cycle for itself
	go u.TimedDelete(auth, user)
}

func (u *LdapUser) TimedDelete(auth, user string) {
	u.mu.RLock()
	idl := u.users[auth].Idl
	max := u.users[auth].Max
	u.mu.RUnlock()

	select {
	case <-idl.C:
		u.mu.Lock()
		max.Stop()
		delete(u.users, auth)
		u.mu.Unlock()
		log.Println(user, "logged out from being idle")
	case <-max.C:
		u.mu.Lock()
		idl.Stop()
		delete(u.users, auth)
		u.mu.Unlock()
		log.Println(user, "logged out due to max time")
	}
}

func (u *LdapUser) Exists(auth string) bool {
	u.mu.Lock()
	user, ok := u.users[auth]
	if ok {
		if !user.Idl.Stop() {
			<-user.Idl.C
		}
		user.Idl.Reset(u.IdlTime)
	}
	u.mu.Unlock()

	return ok
}

func NewLdapUser(idleTime, maxTime time.Duration) *LdapUser {
	return &LdapUser{
		users:   make(map[string]*Timers),
		IdlTime: idleTime,
		MaxTime: maxTime,
	}
}
