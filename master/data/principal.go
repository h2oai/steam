package data

import (
	"fmt"
)

// --- Datastore-backed Principal Impl ---

type Principal struct {
	ds          *Datastore
	identity    *IdentityAndPassword
	permissions map[int64]bool
	isSuperuser bool
}

func (pz *Principal) Id() int64 {
	return pz.identity.Id
}

func (pz *Principal) WorkgroupId() int64 {
	return pz.identity.WorkgroupId
}

func (pz *Principal) Name() string {
	return pz.identity.Name
}

func (pz *Principal) Password() string {
	return pz.identity.Password
}

func (pz *Principal) IsActive() bool {
	return pz.identity.IsActive
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
	return fmt.Errorf("Identity %s does not have permission '%s' to perform this operation", pz.Name(), pz.ds.permissionMap[code].Description)
}

// TODO use bitwise ops to simplify this
func (pz *Principal) hasPrivilege(entityTypeId, entityId int64, expectedPrivilege string) (bool, error) {
	if pz.IsSuperuser() {
		return true, nil
	}

	owns := false
	canEdit := false
	canView := false

	privileges, err := pz.ds.readPrivileges(pz.identity.Id, entityTypeId, entityId)
	if err != nil {
		return false, err
	}

	if len(privileges) == 0 {
		return false, nil
	}

	for _, p := range privileges {
		switch p {
		case Owns:
			owns = true
			canEdit = true
			canView = true
		case CanEdit:
			canEdit = true
			canView = true
		case CanView:
			canView = true
		}

		switch expectedPrivilege {
		case Owns:
			if owns {
				return true, nil
			}
		case CanEdit:
			if owns || canEdit {
				return true, nil
			}
		case CanView:
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
		return fmt.Errorf("Identity %s does not have privilege '%s' on the entity %s:%d", pz.Name(), expectedPrivilege, pz.ds.entityTypeMap[entityTypeId].Name, entityId)
	}
	return nil
}

func (pz *Principal) Owns(entityTypeId, entityId int64) (bool, error) {
	return pz.hasPrivilege(entityTypeId, entityId, Owns)
}

func (pz *Principal) CanEdit(entityTypeId, entityId int64) (bool, error) {
	return pz.hasPrivilege(entityTypeId, entityId, CanEdit)
}

func (pz *Principal) CanView(entityTypeId, entityId int64) (bool, error) {
	return pz.hasPrivilege(entityTypeId, entityId, CanView)
}

func (pz *Principal) CheckOwns(entityTypeId, entityId int64) error {
	return pz.checkPrivilege(entityTypeId, entityId, Owns)
}

func (pz *Principal) CheckEdit(entityTypeId, entityId int64) error {
	return pz.checkPrivilege(entityTypeId, entityId, CanEdit)
}

func (pz *Principal) CheckView(entityTypeId, entityId int64) error {
	return pz.checkPrivilege(entityTypeId, entityId, CanView)
}
