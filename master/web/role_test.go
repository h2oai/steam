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

package web

import "testing"

func TestRoleCRUD(tt *testing.T) {
	t := newTest(tt)

	const name1 = "role1"
	const description1 = "description1"

	id, err := t.svc.CreateRole(t.su, name1, description1)
	t.nil(err)

	role, err := t.svc.GetRole(t.su, id)
	t.nil(err)

	t.ok(name1 == role.Name, "name")
	t.ok(description1 == role.Description, "description")

	role, err = t.svc.GetRoleByName(t.su, name1)
	t.nil(err)
	t.ok(name1 == role.Name, "name")
	t.ok(description1 == role.Description, "description")

	roles, err := t.svc.GetRoles(t.su, 0, 1000)
	t.nil(err)

	t.ok(len(roles) == 1, "role count: expected 1, got %d", len(roles))

	role = roles[0]
	t.ok(name1 == role.Name, "name")
	t.ok(description1 == role.Description, "description")

	const name2 = "role2"
	const description2 = "description2"

	err = t.svc.UpdateRole(t.su, id, name2, description2)
	t.nil(err)

	role, err = t.svc.GetRole(t.su, id)
	t.nil(err)

	t.ok(name2 == role.Name, "name")
	t.ok(description2 == role.Description, "description")

	role, err = t.svc.GetRoleByName(t.su, name2)
	t.nil(err)
	t.ok(name2 == role.Name, "name")
	t.ok(description2 == role.Description, "description")

	roles, err = t.svc.GetRoles(t.su, 0, 2000)
	t.nil(err)

	t.ok(len(roles) == 1, "role count")

	role = roles[0]
	t.ok(name2 == role.Name, "name")
	t.ok(description2 == role.Description, "description")

	err = t.svc.DeleteRole(t.su, id)
	t.nil(err)
}

func TestRoleDeletion(tt *testing.T) {
	t := newTest(tt)

	const name1 = "role1"
	const description1 = "description1"

	id, err := t.svc.CreateRole(t.su, name1, description1)
	t.nil(err)

	err = t.svc.DeleteRole(t.su, id)
	t.nil(err)

	_, err = t.svc.GetRole(t.su, id)
	t.notnil(err)

	_, err = t.svc.GetRoleByName(t.su, name1)
	t.notnil(err)

	roles, err := t.svc.GetRoles(t.su, 0, 1000)
	t.nil(err)
	t.ok(len(roles) == 0, "role count")

	// err = t.svc.DeleteRole(t.su, id) // should fail on a duplicate attempt
	// t.notnil(err)
}

func TestRolePermissionLinking(tt *testing.T) {
	t := newTest(tt)

	roleId, err := t.svc.CreateRole(t.su, "name1", "description1")
	t.nil(err)

	expected, err := t.svc.GetAllPermissions(t.su)
	t.nil(err)

	pids1 := make([]int64, len(expected))
	for i, p := range expected {
		pids1[i] = p.Id
	}

	err = t.svc.LinkRoleWithPermissions(t.su, roleId, pids1)
	t.nil(err)

	actual, err := t.svc.GetPermissionsForRole(t.su, roleId)
	t.nil(err)

	t.ok(len(expected) == len(actual), "permissions")

	// change permissions

	pids2 := pids1[0:5]

	err = t.svc.LinkRoleWithPermissions(t.su, roleId, pids2)
	t.nil(err)

	actual, err = t.svc.GetPermissionsForRole(t.su, roleId)
	t.nil(err)

	t.ok(len(pids2) == len(actual), "permissions")

	// remove all permissions

	err = t.svc.LinkRoleWithPermissions(t.su, roleId, []int64{})
	t.nil(err)

	actual, err = t.svc.GetPermissionsForRole(t.su, roleId)
	t.nil(err)

	t.ok(len(actual) == 0, "permissions")
}
