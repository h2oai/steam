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
	"github.com/h2oai/steam/master/az"
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
