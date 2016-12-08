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
	"fmt"

	"github.com/pkg/errors"
)

// --- Datastore-backed Principal Impl ---

type Principal struct {
	ds          *Datastore
	Identity    *Identity
	permissions map[int64]bool
	isSuperuser bool
}

func (pz *Principal) Id() int64 {
	return pz.Identity.Id
}

func (pz *Principal) WorkgroupId() int64 {
	return pz.Identity.WorkgroupId.Int64
}

func (pz *Principal) Name() string {
	return pz.Identity.Name
}

func (pz *Principal) Password() string {
	return pz.Identity.Password.String
}

func (pz *Principal) IsActive() bool {
	return pz.Identity.IsActive
}

func (pz *Principal) IsSuperuser() bool {
	return pz.isSuperuser
}

func (pz *Principal) HasPermission(code int64) bool {
	if pz.IsSuperuser() {
		return true
	}
	_, ok := pz.permissions[code]
	return ok
}

func (pz *Principal) CheckPermission(code int64) error {
	if pz.HasPermission(code) {
		return nil
	}
	return fmt.Errorf("identity %s does not have permission '%s' to perform this operation", pz.Name(), pz.ds.permissionMap[code])
}

// TODO use bitwise ops to simplify this
func (pz *Principal) hasPrivilege(entityTypeId, entityId int64, expectedPrivilege string) (bool, error) {
	if pz.IsSuperuser() {
		return true, nil
	}

	owns := false
	canEdit := false
	canView := false

	tx, err := pz.ds.db.Begin()
	if err != nil {
		return false, errors.Wrap(err, "beginning transaction")
	}

	var privileges []privilege
	if err := tx.Wrap(func() error {
		var err error
		privileges, err = readPrivileges(tx,
			ById(pz.Identity.Id),
			ByEntityTypeId(entityTypeId),
			ByEntityId(entityId),
		)
		return errors.Wrap(err, "reading privileges from database")
	}); err != nil {
		return false, errors.Wrap(err, "committing transaction")
	}

	if len(privileges) == 0 {
		return false, nil
	}

	for _, p := range privileges {
		switch p.Type {
		case Owns:
			owns = true
			canEdit = true
			canView = true
		case Edit:
			canEdit = true
			canView = true
		case View:
			canView = true
		}

		switch expectedPrivilege {
		case Owns:
			if owns {
				return true, nil
			}
		case Edit:
			if owns || canEdit {
				return true, nil
			}
		case View:
			if owns || canEdit || canView {
				return true, nil
			}
		}
	}
	return false, nil
}

func (pz *Principal) checkPrivilege(entityTypeId, entityId int64, expectedPrivilege string) error {
	ok, err := pz.hasPrivilege(entityTypeId, entityId, expectedPrivilege)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("identity %s does not have privilege '%s' on the entity %s:%d", pz.Name(), expectedPrivilege, pz.ds.entityTypeMap[entityTypeId], entityId)
	}
	return nil
}

func (pz *Principal) Owns(entityTypeId, entityId int64) (bool, error) {
	return pz.hasPrivilege(entityTypeId, entityId, Owns)
}

func (pz *Principal) CanEdit(entityTypeId, entityId int64) (bool, error) {
	return pz.hasPrivilege(entityTypeId, entityId, Edit)
}

func (pz *Principal) CanView(entityTypeId, entityId int64) (bool, error) {
	return pz.hasPrivilege(entityTypeId, entityId, View)
}

func (pz *Principal) CheckOwns(entityTypeId, entityId int64) error {
	return pz.checkPrivilege(entityTypeId, entityId, Owns)
}

func (pz *Principal) CheckEdit(entityTypeId, entityId int64) error {
	return pz.checkPrivilege(entityTypeId, entityId, Edit)
}

func (pz *Principal) CheckView(entityTypeId, entityId int64) error {
	return pz.checkPrivilege(entityTypeId, entityId, View)
}

func (pz *Principal) String() string {
	return pz.Name()
}
