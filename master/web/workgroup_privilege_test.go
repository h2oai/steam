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

import (
	"github.com/h2oai/steam/master/data"
	"testing"
)

func TestOwnSharingOfWorkgroup(tt *testing.T) {
	t := newTest(tt)

	entityId, err := t.svc.CreateWorkgroup(t.su, "entity1", "description")
	t.nil(err)

	const user1Name = "user1"
	userId, err := t.svc.CreateIdentity(t.su, user1Name, "password1")
	t.nil(err)

	group1Id, err := t.svc.CreateWorkgroup(t.su, "group1", "group1 description")
	t.nil(err)

	err = t.svc.LinkIdentityWithWorkgroup(t.su, userId, group1Id)
	t.nil(err)

	user1, err := t.dir.Lookup(user1Name)
	t.nil(err)

	role1Id, err := t.svc.CreateRole(t.su, "role1", "role1 description")
	t.nil(err)

	permissionMap := buildPermissionMap(t)

	err = t.svc.LinkRoleWithPermissions(t.su, role1Id, []int64{
		permissionMap[data.ViewWorkgroup],
		permissionMap[data.ManageWorkgroup],
	})

	t.nil(err)

	err = t.svc.LinkIdentityWithRole(t.su, userId, role1Id)
	t.nil(err)

	// edit as user1 -- should fail
	err = t.svc.UpdateWorkgroup(user1, entityId, "entity2", "description")
	t.notnil(err)

	entityTypeMap := buildEntityTypeMap(t)

	// share as su
	err = t.svc.ShareEntity(t.su, data.CanEdit, group1Id, entityTypeMap[data.WorkgroupEntity], entityId)
	t.nil(err)

	user1, err = t.dir.Lookup(user1Name) // reload
	t.nil(err)

	// edit as user1 -- should pass
	err = t.svc.UpdateWorkgroup(user1, entityId, "entity2", "description")
	t.nil(err)

	// view as user1 -- should pass
	entity, err := t.svc.GetWorkgroup(user1, entityId)
	t.nil(err)

	t.ok(entity.Name == "entity2", "entity name")

	user2Id, err := t.svc.CreateIdentity(t.su, "user2", "password2")
	t.nil(err)

	group2Id, err := t.svc.CreateWorkgroup(t.su, "group2", "group2 description")
	t.nil(err)

	err = t.svc.LinkIdentityWithWorkgroup(t.su, user2Id, group2Id)
	t.nil(err)

	role2Id, err := t.svc.CreateRole(t.su, "role2", "role2 description")
	t.nil(err)

	err = t.svc.LinkRoleWithPermissions(t.su, role2Id, []int64{
		permissionMap[data.ViewWorkgroup],
	})
	t.nil(err)

	err = t.svc.LinkIdentityWithRole(t.su, user2Id, role2Id)
	t.nil(err)

	user2, err := t.dir.Lookup("user2")
	t.nil(err)

	// view as user2 -- should fail (group2 does not have view privilege)
	_, err = t.svc.GetWorkgroup(user2, role2Id)
	t.notnil(err)

	// share as user1 to group2 -- should fail (user1 cannot view group2)
	err = t.svc.ShareEntity(user1, data.CanView, group2Id, entityTypeMap[data.WorkgroupEntity], entityId)
	t.notnil(err)

	// share group2 with user1 as su
	err = t.svc.ShareEntity(t.su, data.CanView, group1Id, entityTypeMap[data.WorkgroupEntity], group2Id)
	t.nil(err)

	// share as user1 to group2 -- should fail (user1 does not have own privilege)
	err = t.svc.ShareEntity(user1, data.CanView, group2Id, entityTypeMap[data.WorkgroupEntity], entityId)
	t.notnil(err)

	// make user1 an owner
	err = t.svc.ShareEntity(t.su, data.Owns, group1Id, entityTypeMap[data.WorkgroupEntity], entityId)
	t.nil(err)

	// share as user1 to group2 -- should pass
	err = t.svc.ShareEntity(user1, data.CanView, group2Id, entityTypeMap[data.WorkgroupEntity], entityId)
	t.nil(err)

	// view as user2 -- should pass
	_, err = t.svc.GetWorkgroup(user2, entityId)
	t.nil(err)
}

func TestEditSharingOfWorkgroup(tt *testing.T) {
	t := newTest(tt)

	entityId, err := t.svc.CreateWorkgroup(t.su, "entity1", "description")
	t.nil(err)

	const username = "user1"
	userId, err := t.svc.CreateIdentity(t.su, username, "password1")
	t.nil(err)

	groupId, err := t.svc.CreateWorkgroup(t.su, "group1", "group1 description")
	t.nil(err)

	err = t.svc.LinkIdentityWithWorkgroup(t.su, userId, groupId)
	t.nil(err)

	user, err := t.dir.Lookup(username)
	t.nil(err)

	roleId, err := t.svc.CreateRole(t.su, "role1", "role1 description")
	t.nil(err)

	permissionMap := buildPermissionMap(t)

	err = t.svc.LinkRoleWithPermissions(t.su, roleId, []int64{
		permissionMap[data.ViewWorkgroup],
		permissionMap[data.ManageWorkgroup],
	})
	t.nil(err)

	err = t.svc.LinkIdentityWithRole(t.su, userId, roleId)
	t.nil(err)

	// edit as user -- should fail
	err = t.svc.UpdateWorkgroup(user, entityId, "entity2", "description")
	t.notnil(err)

	entityTypeMap := buildEntityTypeMap(t)

	// share as su
	err = t.svc.ShareEntity(t.su, data.CanEdit, groupId, entityTypeMap[data.WorkgroupEntity], entityId)
	t.nil(err)

	user, err = t.dir.Lookup(username) // reload
	t.nil(err)

	// edit as user -- should pass
	err = t.svc.UpdateWorkgroup(user, entityId, "entity2", "description")
	t.nil(err)

	// view as user -- should pass
	entity, err := t.svc.GetWorkgroup(user, entityId)
	t.nil(err)

	t.ok(entity.Name == "entity2", "entity name")
}

func TestViewSharingOfWorkgroup(tt *testing.T) {
	t := newTest(tt)

	entityId, err := t.svc.CreateWorkgroup(t.su, "entity1", "description")
	t.nil(err)

	const username = "user1"
	userId, err := t.svc.CreateIdentity(t.su, username, "password1")
	t.nil(err)

	groupId, err := t.svc.CreateWorkgroup(t.su, "group1", "group1 description")
	t.nil(err)

	err = t.svc.LinkIdentityWithWorkgroup(t.su, userId, groupId)
	t.nil(err)

	user, err := t.dir.Lookup(username)
	t.nil(err)

	roleId, err := t.svc.CreateRole(t.su, "role1", "role1 description")
	t.nil(err)

	permissionMap := buildPermissionMap(t)

	err = t.svc.LinkRoleWithPermissions(t.su, roleId, []int64{
		permissionMap[data.ViewWorkgroup],
	})
	t.nil(err)

	err = t.svc.LinkIdentityWithRole(t.su, userId, roleId)
	t.nil(err)

	// view as user -- should fail
	entity, err := t.svc.GetWorkgroup(user, entityId)
	t.notnil(err)

	entityTypeMap := buildEntityTypeMap(t)

	// share as su
	err = t.svc.ShareEntity(t.su, data.CanView, groupId, entityTypeMap[data.WorkgroupEntity], entityId)
	t.nil(err)

	user, err = t.dir.Lookup(username) // reload
	t.nil(err)

	// view as user -- should pass
	entity, err = t.svc.GetWorkgroup(user, entityId)
	t.nil(err)

	t.ok(entity.Name == "entity1", "entity name")
}
