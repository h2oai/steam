package data

import (
	"database/sql"
	"fmt"
	"github.com/h2oai/steamY/master/az"
)

// --- Datastore-backed Principal Impl ---

func (ds *Datastore) NewPrincipal(name string) (az.Principal, error) {
	identity, err := ds.readIdentityAndPassword(name)
	if err != nil {
		return nil, err
	}

	if identity == nil {
		return nil, nil
	}

	permissionIds, err := ds.readPermissionsForIdentity(identity.Id)
	if err != nil {
		return nil, err
	}

	permissionMap := ds.permissionMap
	permissions := make(map[int64]bool)
	for _, permissionId := range permissionIds {
		permissions[permissionMap[permissionId].Code] = true
	}

	return &Principal{ds, identity, permissions}, nil
}

type Principal struct {
	ds          *Datastore
	identity    *IdentityAndPassword
	permissions map[int64]bool
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

func (pz *Principal) HasPermission(code int64) bool {
	_, ok := pz.permissions[code]
	return ok
}

func (pz *Principal) CheckPermission(code int64) error {
	if pz.HasPermission(code) {
		return nil
	}
	// FIXME return string representation of permission code
	return fmt.Errorf("Identity %s does not have permission %d to perform this operation", pz.Name(), code)
}

// TODO use bitwise ops to simplify this
func (pz *Principal) hasPrivilege(entityTypeId, entityId int64, expectedPrivilege string) (bool, error) {

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
		return fmt.Errorf("Identity %s does not have privilege '%s' on the entity", pz.Name(), expectedPrivilege)
	}
	return nil
}

func (pz *Principal) Owns(entityTypeId, entityId int64) (bool, error) {
	return pz.hasPrivilege(entityId, entityId, Owns)
}

func (pz *Principal) CanEdit(entityTypeId, entityId int64) (bool, error) {
	return pz.hasPrivilege(entityId, entityId, CanEdit)
}

func (pz *Principal) CanView(entityTypeId, entityId int64) (bool, error) {
	return pz.hasPrivilege(entityId, entityId, CanView)
}

func (pz *Principal) CheckOwns(entityTypeId, entityId int64) error {
	return pz.checkPrivilege(entityId, entityId, Owns)
}

func (pz *Principal) CheckEdit(entityTypeId, entityId int64) error {
	return pz.checkPrivilege(entityId, entityId, CanEdit)
}

func (pz *Principal) CheckView(entityTypeId, entityId int64) error {
	return pz.checkPrivilege(entityId, entityId, CanView)
}

const SystemIdentityName = "system"

func CreateSystemIdentity(db *sql.DB) (int64, int64, error) {
	var id, workgroupId int64
	err := executeTransaction(db, func(tx *sql.Tx) error {
		var err error

		workgroupId, err = createDefaultWorkgroup(tx, SystemIdentityName)
		if err != nil {
			return err
		}

		id, err = createIdentity(tx, SystemIdentityName, "", workgroupId)
		if err != nil {
			return err
		}

		return linkIdentityAndWorkgroup(tx, id, workgroupId)
	})
	return id, workgroupId, err
}

type SystemPrincipal struct {
	identity *IdentityAndPassword
}

func (ds *Datastore) NewSystemPrincipal() (az.Principal, error) {
	identity, err := ds.readIdentityAndPassword(SystemIdentityName)
	if err != nil {
		return nil, err
	}

	if identity == nil {
		return nil, nil
	}

	return &SystemPrincipal{identity}, nil
}

func (pz *SystemPrincipal) Id() int64 {
	return pz.identity.Id
}

func (pz *SystemPrincipal) WorkgroupId() int64 {
	return pz.identity.WorkgroupId
}

func (pz *SystemPrincipal) Name() string {
	return pz.identity.Name
}

func (pz *SystemPrincipal) Password() string {
	return pz.identity.Password
}

func (pz *SystemPrincipal) IsActive() bool {
	return true
}

func (pz *SystemPrincipal) HasPermission(code int64) bool {
	return true
}

func (pz *SystemPrincipal) CheckPermission(code int64) error {
	return nil
}

func (pz *SystemPrincipal) Owns(entityTypeId, entityId int64) (bool, error) {
	return true, nil
}

func (pz *SystemPrincipal) CanEdit(entityTypeId, entityId int64) (bool, error) {
	return true, nil
}

func (pz *SystemPrincipal) CanView(entityTypeId, entityId int64) (bool, error) {
	return true, nil
}

func (pz *SystemPrincipal) CheckOwns(entityTypeId, entityId int64) error {
	return nil
}

func (pz *SystemPrincipal) CheckEdit(entityTypeId, entityId int64) error {
	return nil
}

func (pz *SystemPrincipal) CheckView(entityTypeId, entityId int64) error {
	return nil
}
