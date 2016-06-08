package data

import (
	"github.com/h2oai/steamY/master/az"
	"testing"
)

func getConnection(t *testing.T) (ds *Datastore) {
	ds, err := NewDatastore("steam", "steam", "disable")
	if err != nil {
		t.Error(err)
	}
	return ds
}

func TestSecurity(t *testing.T) {

	ds := getConnection(t)

	p := &az.Principal{1} // FIXME

	// read permissions
	permissions, err := ds.ReadPermissions(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(permissions)

	// create roles

	role1Id, err := ds.CreateRole(p, "role", "a role")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(role1Id)

	role2Id, err := ds.CreateRole(p, "role2", "a role")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(role2Id)

	// read roles

	role1, err := ds.ReadRole(p, role1Id)
	if err != nil {
		t.Fatal(err)
	}
	if role1.Id != role1Id || role1.Name != "role" || role1.Description != "a role" {
		t.Fatal("role not saved correctly")
	}
	role2, err := ds.ReadRole(p, role2Id)
	if err != nil {
		t.Fatal(err)
	}
	if role2.Id != role2Id || role2.Name != "role2" || role2.Description != "a role" {
		t.Fatal("role not saved correctly")
	}

	roles, err := ds.ReadRoles(p, 0, 100000)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(roles)

	for _, role := range roles {
		r, err := ds.ReadRole(p, role.Id)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(r)
		if r.Id != role.Id || r.Name != role.Name || r.Description != role.Description {
			t.Fatal("role not saved correctly")
		}

		rp, err := ds.ReadRoleAndPermissions(p, role.Id)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(rp)
		if len(rp.Permissions) != 0 {
			t.Fatal("expected 0 permissions")
		}
	}

	// rename role

	if err := ds.UpdateRole(p, role1Id, "role1"); err != nil {
		t.Fatal(err)
	}

	role1, err = ds.ReadRole(p, role1Id)
	if err != nil {
		t.Fatal(err)
	}
	if role1.Name != "role1" {
		t.Fatal("role name mismatch")
	}

	// set permissions

	const permCount = 5
	perms := make([]int64, permCount)
	for i := 0; i < permCount; i++ {
		perms[i] = permissions[i].Id
	}
	if err := ds.SetRolePermissions(p, role1Id, perms); err != nil {
		t.Fatal(err)
	}
	rp, err := ds.ReadRoleAndPermissions(p, role1Id)
	if err != nil {
		t.Fatal(err)
	}
	if len(rp.Permissions) != permCount {
		t.Fatal("expected 5 permissions")
	}
	t.Log(rp)

	// create workgroups

	group1Id, err := ds.CreateWorkgroup(p, "group", "a group")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(group1Id)

	group2Id, err := ds.CreateWorkgroup(p, "group2", "a group")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(group2Id)

	// read workgroups
	group1, err := ds.ReadWorkgroup(p, group1Id)
	if err != nil {
		t.Fatal(err)
	}
	if group1.Id != group1Id || group1.Name != "group" || group1.Description != "a group" {
		t.Fatal("group not saved correctly")
	}
	group2, err := ds.ReadWorkgroup(p, group2Id)
	if err != nil {
		t.Fatal(err)
	}
	if group2.Id != group2Id || group2.Name != "group2" || group2.Description != "a group" {
		t.Fatal("group not saved correctly")
	}

	groups, err := ds.ReadWorkgroups(p, 0, 100000)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(groups)

	for _, group := range groups {
		r, err := ds.ReadWorkgroup(p, group.Id)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(r)
		if r.Id != group.Id || r.Name != group.Name || r.Description != group.Description {
			t.Fatal("group not saved correctly")
		}
	}

	// rename workgroup

	if err := ds.UpdateWorkgroup(p, group1Id, "group1"); err != nil {
		t.Fatal(err)
	}

	group1, err = ds.ReadWorkgroup(p, group1Id)
	if err != nil {
		t.Fatal(err)
	}
	if group1.Name != "group1" {
		t.Fatal("group name mismatch")
	}

	// create a user
	user1Id, err := ds.CreateIdentity(p, "user1", "password1")
	if err != nil {
		t.Fatal(err)
	}
	users, err := ds.ReadIdentities(p, 0, 100000)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(users)
	if len(users) == 0 {
		t.Fatal("expected > 0 users")
	}
	userpwd, err := ds.ReadIdentityAndPassword(p, user1Id)
	if err != nil {
		t.Fatal(err)
	}
	if userpwd.Id != user1Id || userpwd.Name != "user1" || userpwd.Password != "password1" {
		t.Fatal("user not saved correctly")
	}

	// set workgroup

	// XXX

	// set roles

	// XXX

	// change workgroup

	// XXX

	// change roles

	// XXX

	// delete roles

	if err := ds.DeleteRole(p, role1Id); err != nil {
		t.Fatal(err)
	}
	if _, err := ds.ReadRole(p, role1Id); err == nil {
		t.Fatal("expected role to be deleted")
	}

	if err := ds.DeleteRole(p, role2Id); err != nil {
		t.Fatal(err)
	}
	if _, err := ds.ReadRole(p, role2Id); err == nil {
		t.Fatal("expected role to be deleted")
	}

	// delete workgroups

	if err := ds.DeleteWorkgroup(p, group1Id); err != nil {
		t.Fatal(err)
	}
	if _, err := ds.ReadWorkgroup(p, group1Id); err == nil {
		t.Fatal("expected group to be deleted")
	}

	if err := ds.DeleteWorkgroup(p, group2Id); err != nil {
		t.Fatal(err)
	}
	if _, err := ds.ReadWorkgroup(p, group2Id); err == nil {
		t.Fatal("expected group to be deleted")
	}

	// delete user

	if err := ds.DeactivateIdentity(p, user1Id); err != nil {
		t.Fatal(err)
	}

	profile1, err := ds.ReadProfile(p, user1Id)
	if err != nil {
		t.Fatal(err)
	}
	if profile1.Identity.IsActive {
		t.Fatal("expected user to be inactive")
	}
}
