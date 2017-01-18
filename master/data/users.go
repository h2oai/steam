package data

import (
	"sync"
	"time"
)

type userTable struct {
	mu sync.RWMutex
	// Logged in users
	users map[string]*time.Timer

	MaxTime time.Duration
}

func (u *userTable) NewUser(token string) {
	u.mu.Lock()
	defer u.mu.Unlock()

	// If user already exists, Noop
	if _, ok := u.users[token]; ok {
		return
	}
	// Self deleting timer after max duration
	t := time.AfterFunc(u.MaxTime, func() {
		u.mu.Lock()
		delete(u.users, token)
		u.mu.Unlock()
	})
	u.users[token] = t
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
