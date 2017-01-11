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
	// Fetch identity
	identity, ok, err := ds.ReadIdentity(ByName(name))
	if err != nil {
		return nil, errors.Wrap(err, "reading identity")
	}
	if !ok {
		return nil, nil
	}
	// Fetch roles
	roles, err := ds.ReadRoles(ForIdentity(identity.Id))
	if err != nil {
		return nil, errors.Wrap(err, "reading roles")
	}

	isAdmin := false
	for _, role := range roles {
		if role.Name == AdminRN {
			isAdmin = true
			break
		}
	}

	perms, err := ds.ReadPermissions(
		ForIdentity(identity.Id),
	)
	if err != nil {
		return nil, err
	}

	permissions := make(map[int64]bool)
	for _, perm := range perms {
		permissions[perm.Id] = true
	}

	return &Principal{ds, &identity, permissions, isAdmin}, nil
}
