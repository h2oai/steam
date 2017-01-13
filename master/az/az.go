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
	IsAdmin() bool
	AuthType() string
	HasPermission(code int64) bool
	CheckPermission(code int64) error
	Owns(entityTypeId, entityId int64) (bool, error)
	CanEdit(entityTypeId, entityId int64) (bool, error)
	CanView(entityTypeId, entityId int64) (bool, error)
	CheckOwns(entityTypeId, entityId int64) error
	CheckEdit(entityTypeId, entityId int64) error
	CheckView(entityTypeId, entityId int64) error
}

type Directory interface {
	Lookup(username string) (Principal, error)
}

type Az interface {
	Authenticate(username string) string
	Identify(r *http.Request) (Principal, error)
}
