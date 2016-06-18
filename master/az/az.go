package az

import (
	"net/http"
)

type Principal interface {
	Id() int64
	WorkgroupId() int64
	Name() string
	Password() string
	IsActive() bool
	HasPermission(code int64) bool
	CheckPermission(code int64) error
	Owns(entityTypeId, entityId int64) (bool, error)
	CanEdit(entityTypeId, entityId int64) (bool, error)
	CanView(entityTypeId, entityId int64) (bool, error)
	CheckOwns(entityTypeId, entityId int64) error
	CheckEdit(entityTypeId, entityId int64) error
	CheckView(entityTypeId, entityId int64) error
}

type Az interface {
	Authenticate(username string) string
	//Authenticate(r *http.Request) (Principal, error)
	Identify(r *http.Request) (Principal, error)
	//Authorize(name string) error
	//Deauthorize(name string) error
	//IsAuthorized(name string) (bool, error)
}
