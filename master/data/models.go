//go:generate scaneo $GOFILE

package data

import (
	"github.com/lib/pq"
	"time"
)

type Meta struct {
	Id    int64
	Key   string
	Value string
}

type Permission struct {
	Id          int64
	Name        string
	Description string
}

type EntityType struct {
	Id   int64
	Name string
}

type Role struct {
	Id          int64
	Name        string
	Description string
	Created     time.Time
}

type Workgroup struct {
	Id          int64
	Name        string
	Description string
	Created     time.Time
}

type Identity struct {
	Id        int64
	Name      string
	IsActive  bool
	LastLogin pq.NullTime
	Created   time.Time
}

type IdentityAndPassword struct {
	Id        int64
	Name      string
	Password  string
	IsActive  bool
	LastLogin pq.NullTime
	Created   time.Time
}
