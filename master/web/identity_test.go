package web

import (
	"testing"
)

func TestIdentityCRUD(tt *testing.T) {
	t := newTest(tt)

	const name1 = "user1"
	const password1 = "password1"

	id, err := t.svc.CreateIdentity(t.su, name1, password1)
	t.nil(err)

	user, err := t.svc.GetIdentity(t.su, id)
	t.nil(err)

	t.ok(name1 == user.Name, "name")
	t.ok(user.IsActive, "active")

	user, err = t.svc.GetIdentityByName(t.su, name1)
	t.nil(err)
	t.ok(name1 == user.Name, "name")

	users, err := t.svc.GetIdentities(t.su, 0, 1000)
	t.nil(err)

	t.ok(len(users) == 2, "user count")

	t.ok(t.su.Name() == users[0].Name, "name")
	t.ok(name1 == users[1].Name, "name")

	err = t.svc.DeactivateIdentity(t.su, id)
	t.nil(err)

	user, err = t.svc.GetIdentity(t.su, id)
	t.nil(err)

	t.ok(name1 == user.Name, "name")
	t.ok(!user.IsActive, "active")
}

func TestIdentityWorkgroupLinking(tt *testing.T) {
	t := newTest(tt)

	userId, err := t.svc.CreateIdentity(t.su, "user1", "password1")
	t.nil(err)

	groupId, err := t.svc.CreateWorkgroup(t.su, "group1", "group description1")
	t.nil(err)

	users, err := t.svc.GetIdentitiesForWorkgroup(t.su, groupId)
	t.nil(err)

	t.ok(len(users) == 0, "users for group")

	err = t.svc.LinkIdentityWithWorkgroup(t.su, userId, groupId)
	t.nil(err)

	users, err = t.svc.GetIdentitiesForWorkgroup(t.su, groupId)
	t.nil(err)

	t.ok(len(users) == 1, "users for group")
	t.ok(users[0].Name == "user1", "user name")

	groups, err := t.svc.GetWorkgroupsForIdentity(t.su, userId)
	t.nil(err)

	t.ok(len(groups) == 1, "groups for user")
	t.ok(groups[0].Name == "group1", "group name")

	err = t.svc.UnlinkIdentityFromWorkgroup(t.su, userId, groupId)
	t.nil(err)

	users, err = t.svc.GetIdentitiesForWorkgroup(t.su, groupId)
	t.nil(err)

	t.ok(len(users) == 0, "users for group")
}

func TestIdentityAndRoleLinking(tt *testing.T) {
	t := newTest(tt)

	userId, err := t.svc.CreateIdentity(t.su, "user1", "password1")
	t.nil(err)

	roleId, err := t.svc.CreateRole(t.su, "role1", "role description1")
	t.nil(err)

	users, err := t.svc.GetIdentitiesForRole(t.su, roleId)
	t.nil(err)

	t.ok(len(users) == 0, "users for role")

	err = t.svc.LinkIdentityWithRole(t.su, userId, roleId)
	t.nil(err)

	users, err = t.svc.GetIdentitiesForRole(t.su, roleId)
	t.nil(err)

	t.ok(len(users) == 1, "users for role")
	t.ok(users[0].Name == "user1", "user name")

	roles, err := t.svc.GetRolesForIdentity(t.su, userId)
	t.nil(err)

	t.ok(len(roles) == 1, "roles for user")
	t.ok(roles[0].Name == "role1", "role name")

	perms, err := t.svc.GetPermissionsForIdentity(t.su, userId)
	t.ok(len(perms) == 0, "permissions for user")

	allPerms, err := t.svc.GetAllPermissions(t.su)
	t.nil(err)

	permIds := make([]int64, len(allPerms))
	for i, p := range allPerms {
		permIds[i] = p.Id
	}

	err = t.svc.LinkRoleWithPermissions(t.su, roleId, permIds)
	t.nil(err)

	perms, err = t.svc.GetPermissionsForIdentity(t.su, userId)
	t.ok(len(perms) == len(allPerms), "permissions for user")

	err = t.svc.UnlinkIdentityFromRole(t.su, userId, roleId)
	t.nil(err)

	users, err = t.svc.GetIdentitiesForRole(t.su, roleId)
	t.nil(err)

	t.ok(len(users) == 0, "users for role")

	perms, err = t.svc.GetPermissionsForIdentity(t.su, userId)
	t.ok(len(perms) == 0, "permissions for user")
}
