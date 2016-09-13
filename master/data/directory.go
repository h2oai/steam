package data

import (
	"github.com/h2oai/steamY/master/az"
	"github.com/pkg/errors"
)

// --- Datastore-backed Directory Impl ---

func (ds *Datastore) Lookup(name string) (az.Principal, error) {
	identity, err := ds.readIdentityAndPassword(name)
	if err != nil {
		return nil, errors.Wrap(err, "failed reading identity with password")
	}

	if identity == nil {
		return nil, nil
	}

	roleNames, err := ds.readRoleNamesForIdentity(identity.Id)
	if err != nil {
		return nil, err
	}

	isSuperuser := false
	for _, roleName := range roleNames {
		if roleName == SuperuserRoleName {
			isSuperuser = true
			break
		}
	}

	permissionIds, err := ds.readPermissionsForIdentity(identity.Id)
	if err != nil {
		return nil, err
	}

	permissions := make(map[int64]bool)
	for _, permissionId := range permissionIds {
		permissions[permissionId] = true
	}

	return &Principal{ds, identity, permissions, isSuperuser}, nil
}
