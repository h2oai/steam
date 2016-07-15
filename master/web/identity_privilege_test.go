package web

import (
	"github.com/h2oai/steamY/master/data"
	"testing"
)

func TestOwnSharingOfIdentity(tt *testing.T) {
	t := newTest(tt)

	entityId, err := t.svc.CreateIdentity(t.su, "entity1", "description")
	t.nil(err)

	const user1Name = "user1"
	userId, err := t.svc.CreateIdentity(t.su, user1Name, "password1")
	t.nil(err)

	group1Id, err := t.svc.CreateWorkgroup(t.su, "group1", "group1 description")
	t.nil(err)

	err = t.svc.LinkIdentityAndWorkgroup(t.su, userId, group1Id)
	t.nil(err)

	user1, err := t.dir.Lookup(user1Name)
	t.nil(err)

	role1Id, err := t.svc.CreateRole(t.su, "role1", "role1 description")
	t.nil(err)

	permissionMap := buildPermissionMap(t)

	err = t.svc.LinkRoleAndPermissions(t.su, role1Id, []int64{
		permissionMap[data.ViewWorkgroup],
		permissionMap[data.ManageIdentity],
		permissionMap[data.ViewIdentity],
	})

	t.nil(err)

	err = t.svc.LinkIdentityAndRole(t.su, userId, role1Id)
	t.nil(err)

	// edit as user1 -- should fail
	err = t.svc.UpdateIdentity(user1, entityId, "password2")
	t.notnil(err)

	entityTypeMap := buildEntityTypeMap(t)

	// share as su
	err = t.svc.ShareEntity(t.su, data.CanEdit, group1Id, entityTypeMap[data.IdentityEntity], entityId)
	t.nil(err)

	user1, err = t.dir.Lookup(user1Name) // reload
	t.nil(err)

	// edit as user1 -- should pass
	err = t.svc.UpdateIdentity(user1, entityId, "password2")
	t.nil(err)

	// view as user1 -- should pass
	_, err = t.svc.GetIdentity(user1, entityId)
	t.nil(err)

	user2Id, err := t.svc.CreateIdentity(t.su, "user2", "password2")
	t.nil(err)

	group2Id, err := t.svc.CreateWorkgroup(t.su, "group2", "group2 description")
	t.nil(err)

	err = t.svc.LinkIdentityAndWorkgroup(t.su, user2Id, group2Id)
	t.nil(err)

	role2Id, err := t.svc.CreateRole(t.su, "role2", "role2 description")
	t.nil(err)

	err = t.svc.LinkRoleAndPermissions(t.su, role2Id, []int64{
		permissionMap[data.ViewIdentity],
	})
	t.nil(err)

	err = t.svc.LinkIdentityAndRole(t.su, user2Id, role2Id)
	t.nil(err)

	user2, err := t.dir.Lookup("user2")
	t.nil(err)

	// view as user2 -- should fail (group2 does not have view privilege)
	_, err = t.svc.GetIdentity(user2, role2Id)
	t.notnil(err)

	// share as user1 to group2 -- should fail (user1 cannot view group2)
	err = t.svc.ShareEntity(user1, data.CanView, group2Id, entityTypeMap[data.IdentityEntity], entityId)
	t.notnil(err)

	// share group2 with user1 as su
	err = t.svc.ShareEntity(t.su, data.CanView, group1Id, entityTypeMap[data.WorkgroupEntity], group2Id)
	t.nil(err)

	// share as user1 to group2 -- should fail (user1 does not have own privilege)
	err = t.svc.ShareEntity(user1, data.CanView, group2Id, entityTypeMap[data.IdentityEntity], entityId)
	t.notnil(err)

	// make user1 an owner
	err = t.svc.ShareEntity(t.su, data.Owns, group1Id, entityTypeMap[data.IdentityEntity], entityId)
	t.nil(err)

	// share as user1 to group2 -- should pass
	err = t.svc.ShareEntity(user1, data.CanView, group2Id, entityTypeMap[data.IdentityEntity], entityId)
	t.nil(err)

	// view as user2 -- should pass
	_, err = t.svc.GetIdentity(user2, entityId)
	t.nil(err)
}

func TestEditSharingOfIdentity(tt *testing.T) {
	t := newTest(tt)

	entityId, err := t.svc.CreateIdentity(t.su, "entity1", "description")
	t.nil(err)

	const username = "user1"
	userId, err := t.svc.CreateIdentity(t.su, username, "password1")
	t.nil(err)

	groupId, err := t.svc.CreateWorkgroup(t.su, "group1", "group1 description")
	t.nil(err)

	err = t.svc.LinkIdentityAndWorkgroup(t.su, userId, groupId)
	t.nil(err)

	user, err := t.dir.Lookup(username)
	t.nil(err)

	roleId, err := t.svc.CreateRole(t.su, "role1", "role1 description")
	t.nil(err)

	permissionMap := buildPermissionMap(t)

	err = t.svc.LinkRoleAndPermissions(t.su, roleId, []int64{
		permissionMap[data.ViewIdentity],
		permissionMap[data.ManageIdentity],
	})
	t.nil(err)

	err = t.svc.LinkIdentityAndRole(t.su, userId, roleId)
	t.nil(err)

	// edit as user -- should fail
	err = t.svc.UpdateIdentity(user, entityId, "password2")
	t.notnil(err)

	entityTypeMap := buildEntityTypeMap(t)

	// share as su
	err = t.svc.ShareEntity(t.su, data.CanEdit, groupId, entityTypeMap[data.IdentityEntity], entityId)
	t.nil(err)

	user, err = t.dir.Lookup(username) // reload
	t.nil(err)

	// edit as user -- should pass
	err = t.svc.UpdateIdentity(user, entityId, "password2")
	t.nil(err)

	// view as user -- should pass
	_, err = t.svc.GetIdentity(user, entityId)
	t.nil(err)
}

func TestViewSharingOfIdentity(tt *testing.T) {
	t := newTest(tt)

	entityId, err := t.svc.CreateIdentity(t.su, "entity1", "description")
	t.nil(err)

	const username = "user1"
	userId, err := t.svc.CreateIdentity(t.su, username, "password1")
	t.nil(err)

	groupId, err := t.svc.CreateWorkgroup(t.su, "group1", "group1 description")
	t.nil(err)

	err = t.svc.LinkIdentityAndWorkgroup(t.su, userId, groupId)
	t.nil(err)

	user, err := t.dir.Lookup(username)
	t.nil(err)

	roleId, err := t.svc.CreateRole(t.su, "role1", "role1 description")
	t.nil(err)

	permissionMap := buildPermissionMap(t)

	err = t.svc.LinkRoleAndPermissions(t.su, roleId, []int64{
		permissionMap[data.ViewIdentity],
	})
	t.nil(err)

	err = t.svc.LinkIdentityAndRole(t.su, userId, roleId)
	t.nil(err)

	// view as user -- should fail
	_, err = t.svc.GetIdentity(user, entityId)
	t.notnil(err)

	entityTypeMap := buildEntityTypeMap(t)

	// share as su
	err = t.svc.ShareEntity(t.su, data.CanView, groupId, entityTypeMap[data.IdentityEntity], entityId)
	t.nil(err)

	user, err = t.dir.Lookup(username) // reload
	t.nil(err)

	// view as user -- should pass
	_, err = t.svc.GetIdentity(user, entityId)
	t.nil(err)
}
