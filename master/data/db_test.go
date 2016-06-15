package data

import (
	"github.com/h2oai/steamY/master/az"
	"testing"
)

func connect(t *testing.T) (*Datastore, *az.Principal) {
	db, err := Connect("steam", "steam", "disable")
	if err != nil {
		t.Error(err)
	}
	if err := truncate(db); err != nil {
		t.Error(err)
	}
	if err := prime(db); err != nil {
		t.Error(err)
	}

	ds, err := NewDatastore(db)
	if err != nil {
		t.Error(err)
	}

	uid, wgid, err := ds.CreateSuperuser("Superuser", "")
	if err != nil {
		t.Fatal(err)
	}

	p := &az.Principal{uid, wgid}

	if err := ds.SetupSuperuser(p); err != nil {
		t.Fatal(err)
	}

	return ds, p
}

func TestPrivilegesForIdentity(t *testing.T) {
	ds, p := connect(t)

	// Create user
	uid, uwgid, err := ds.CreateIdentity(p, "user", "password1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uid)

	// Create entity
	eid, err := ds.CreateWorkgroup(p, "foo", "description")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(eid)

	// Make user own entity
	if err := ds.CreatePrivilege(p, Privilege{
		Owns,
		uwgid,
		ds.EntityTypes.Workgroup,
		eid,
	}); err != nil {
		t.Fatal(err)
	}

	// Get user's privilege on entity
	privileges, err := ds.ReadPrivileges(p, uid, ds.EntityTypes.Workgroup, eid)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(privileges)

	if privileges[0] != "own" {
		t.Fatal("wrong privilege")
	}

	// Drop privileges
	ds.DeletePrivilege(p, Privilege{
		Owns,
		uwgid,
		ds.EntityTypes.Workgroup,
		eid,
	})

	// Get user's privilege on entity
	privileges, err = ds.ReadPrivileges(p, uid, ds.EntityTypes.Workgroup, eid)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(privileges)

	if len(privileges) != 0 {
		t.Fatal("wrong privilege")
	}

}

func TestPrivilegesForWorkgroup(t *testing.T) {
	ds, p := connect(t)

	// Create user
	uid, _, err := ds.CreateIdentity(p, "user", "password1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uid)

	// Create group
	wgid, err := ds.CreateWorkgroup(p, "group", "description")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(wgid)

	// Add user to group
	if err := ds.LinkIdentityAndWorkgroup(p, uid, wgid); err != nil {
		t.Fatal(err)
	}

	// Create entity
	eid, err := ds.CreateWorkgroup(p, "othergroup", "description")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(eid)

	// Make group edit entity
	if err := ds.CreatePrivilege(p, Privilege{
		CanEdit,
		wgid,
		ds.EntityTypes.Workgroup,
		eid,
	}); err != nil {
		t.Fatal(err)
	}

	// Read user's privilege on entity
	privileges, err := ds.ReadPrivileges(p, uid, ds.EntityTypes.Workgroup, eid)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(privileges)

	if privileges[0] != "edit" {
		t.Fatal("wrong privilege")
	}

	// Drop group's privileges
	ds.DeletePrivilege(p, Privilege{
		CanEdit,
		wgid,
		ds.EntityTypes.Workgroup,
		eid,
	})

	// Read user's privilege on entity
	privileges, err = ds.ReadPrivileges(p, uid, ds.EntityTypes.Workgroup, eid)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(privileges)

	if len(privileges) != 0 {
		t.Fatal("wrong privilege")
	}
}

func TestPrivilegeCollationForIdentity(t *testing.T) {
	ds, p := connect(t)

	// Create user
	uid, uwgid, err := ds.CreateIdentity(p, "user", "password1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uid)

	// Create group
	wgid, err := ds.CreateWorkgroup(p, "group", "description")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(wgid)

	// Add user to group
	if err := ds.LinkIdentityAndWorkgroup(p, uid, wgid); err != nil {
		t.Fatal(err)
	}

	// Create entity
	eid, err := ds.CreateWorkgroup(p, "foo", "description")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(eid)

	// Make user own entity
	if err := ds.CreatePrivilege(p, Privilege{
		Owns,
		uwgid,
		ds.EntityTypes.Workgroup,
		eid,
	}); err != nil {
		t.Fatal(err)
	}

	// Make group edit entity
	if err := ds.CreatePrivilege(p, Privilege{
		CanEdit,
		wgid,
		ds.EntityTypes.Workgroup,
		eid,
	}); err != nil {
		t.Fatal(err)
	}

	// Get user's privilege on entity
	privileges, err := ds.ReadPrivileges(p, uid, ds.EntityTypes.Workgroup, eid)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(privileges)

	if len(privileges) != 2 {
		t.Fatal("wrong privileges")
	}

	// Drop group privileges
	ds.DeletePrivilege(p, Privilege{
		CanEdit,
		wgid,
		ds.EntityTypes.Workgroup,
		eid,
	})

	// Get user's privilege on entity
	privileges, err = ds.ReadPrivileges(p, uid, ds.EntityTypes.Workgroup, eid)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(privileges)

	if len(privileges) != 1 {
		t.Fatal("wrong privileges")
	}

	// Drop user's privileges
	ds.DeletePrivilege(p, Privilege{
		Owns,
		uwgid,
		ds.EntityTypes.Workgroup,
		eid,
	})

	// Get user's privilege on entity
	privileges, err = ds.ReadPrivileges(p, uid, ds.EntityTypes.Workgroup, eid)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(privileges)

	if len(privileges) != 0 {
		t.Fatal("wrong privileges")
	}
}

func TestSecurity(t *testing.T) {
	ds, p := connect(t)

	// read permissions
	permissions, err := ds.ReadAllPermissions(p)
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

	if len(roles) == 0 {
		t.Fatal("expected > 0 roles")
	}

	for _, role := range roles {
		if role.Name == "Superuser" {
			continue
		}
		r, err := ds.ReadRole(p, role.Id)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(r)
		if r.Id != role.Id || r.Name != role.Name || r.Description != role.Description {
			t.Fatal("role not saved correctly")
		}

		perms, err := ds.ReadPermissionsForRole(p, role.Id)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(perms)
		if len(perms) != 0 {
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
	pr, err := ds.ReadPermissionsForRole(p, role1Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pr)
	if len(pr) != permCount {
		t.Fatal("expected 5 permissions")
	}

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
	user1Id, _, err := ds.CreateIdentity(p, "user1", "password1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.CreatePrivilege(p, Privilege{
		Owns,
		p.WorkgroupId,
		ds.EntityTypes.Identity,
		user1Id,
	}); err != nil {
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

	if err := ds.LinkIdentityAndWorkgroup(p, user1Id, group1Id); err != nil {
		t.Fatal(err)
	}

	// set roles

	if err := ds.LinkIdentityAndRole(p, user1Id, role1Id); err != nil {
		t.Fatal(err)
	}

	// verify

	wgs, err := ds.ReadWorkgroupsForIdentity(p, user1Id)
	if err != nil {
		t.Fatal(err)
	}
	if len(wgs) != 1 {
		t.Fatal("expected 1 workgroup")
	}
	if wgs[0].Id != group1Id {
		t.Fatal("wrong workgroup id")
	}

	rs, err := ds.ReadRolesForIdentity(p, user1Id)
	if err != nil {
		t.Fatal(err)
	}
	if len(rs) != 1 {
		t.Fatal("expected 1 role")
	}
	if rs[0].Id != role1Id {
		t.Fatal("wrong role id")
	}

	pids, err := ds.ReadPermissionsForIdentity(p, user1Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pids)
	if len(pids) != 5 {
		t.Fatal("expected 5 permissions")
	}

	// change workgroup

	if err := ds.UnlinkIdentityAndWorkgroup(p, user1Id, group1Id); err != nil {
		t.Fatal(err)
	}
	if err := ds.LinkIdentityAndWorkgroup(p, user1Id, group2Id); err != nil {
		t.Fatal(err)
	}

	// change roles

	if err := ds.UnlinkIdentityAndRole(p, user1Id, role1Id); err != nil {
		t.Fatal(err)
	}
	if err := ds.LinkIdentityAndRole(p, user1Id, role2Id); err != nil {
		t.Fatal(err)
	}

	// verify

	wgs, err = ds.ReadWorkgroupsForIdentity(p, user1Id)
	if err != nil {
		t.Fatal(err)
	}
	if len(wgs) != 1 {
		t.Fatal("expected 1 workgroup")
	}
	if wgs[0].Id != group2Id {
		t.Fatal("wrong workgroup id")
	}

	rs, err = ds.ReadRolesForIdentity(p, user1Id)
	if err != nil {
		t.Fatal(err)
	}
	if len(rs) != 1 {
		t.Fatal("expected 1 role")
	}
	if rs[0].Id != role2Id {
		t.Fatal("wrong role id")
	}

	// read workgroup + identities

	wi, err := ds.ReadIdentitiesForWorkgroup(p, group2Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(wi)
	if len(wi) != 1 {
		t.Fatal("expected 1 identity")
	}

	// read roles + identities

	ri, err := ds.ReadIdentitiesForRole(p, role2Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ri)
	if len(ri) != 1 {
		t.Fatal("expected 1 identity")
	}

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

	id1, err := ds.ReadIdentity(p, user1Id)
	if err != nil {
		t.Fatal(err)
	}
	if id1.IsActive {
		t.Fatal("expected user to be inactive")
	}
}

func TestEngines(t *testing.T) {
	ds, p := connect(t)

	id1, err := ds.CreateEngine(p, "engine1", "location1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id1)

	id2, err := ds.CreateEngine(p, "engine2", "location2")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id2)

	engines, err := ds.ReadEngines(p)
	if err != nil {
		t.Fatal(err)
	}

	if len(engines) != 2 {
		t.Fatal("expected 2 engines")
	}

	e1, err := ds.ReadEngine(p, id1)
	if err != nil {
		t.Fatal(err)
	}

	if e1.Id != id1 || e1.Name != "engine1" || e1.Location != "location1" {
		t.Fatal("wrong engine")
	}

	e2, err := ds.ReadEngine(p, id2)
	if err != nil {
		t.Fatal(err)
	}

	if e2.Id != id2 || e2.Name != "engine2" || e2.Location != "location2" {
		t.Fatal("wrong engine")
	}

	if err := ds.DeleteEngine(p, id1); err != nil {
		t.Fatal(err)
	}

	engines, err = ds.ReadEngines(p)
	if err != nil {
		t.Fatal(err)
	}

	if len(engines) != 1 {
		t.Fatal("expected 1 engine")
	}

	if err := ds.DeleteEngine(p, id2); err != nil {
		t.Fatal(err)
	}

	engines, err = ds.ReadEngines(p)
	if err != nil {
		t.Fatal(err)
	}

	if len(engines) != 0 {
		t.Fatal("expected 0 engine")
	}
}
