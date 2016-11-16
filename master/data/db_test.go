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
	"testing"
	"time"
)

func setup(t *testing.T) (*Datastore, az.Principal) {
	db, err := connect("steam", "steam", "disable")
	if err != nil {
		t.Error(err)
	}
	if err := truncate(db); err != nil {
		t.Error(err)
	}
	if err := prime(db); err != nil {
		t.Error(err)
	}

	ds, err := newDatastore(db)
	if err != nil {
		t.Error(err)
	}

	const suName = "Superuser"
	const suPassword = "Password"
	_, _, err = ds.CreateSuperuser(suName, suPassword)
	if err != nil {
		t.Fatal(err)
	}

	p, err := ds.Lookup(suName)

	return ds, p
}

func TestInvalidIdentity(t *testing.T) {
	ds, _ := setup(t)

	userpwd, err := ds.readIdentityAndPassword("user1")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(userpwd)

}

func TestPrivilegesForIdentity(t *testing.T) {
	ds, p := setup(t)

	t.Log("superuser=", p.IsSuperuser())

	// Create identity
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
	privileges, err := ds.readPrivileges(uid, ds.EntityTypes.Workgroup, eid)
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
	privileges, err = ds.readPrivileges(uid, ds.EntityTypes.Workgroup, eid)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(privileges)

	if len(privileges) != 0 {
		t.Fatal("wrong privilege")
	}

}

func TestPrivilegesForWorkgroup(t *testing.T) {
	ds, p := setup(t)

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
	privileges, err := ds.readPrivileges(uid, ds.EntityTypes.Workgroup, eid)
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
	privileges, err = ds.readPrivileges(uid, ds.EntityTypes.Workgroup, eid)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(privileges)

	if len(privileges) != 0 {
		t.Fatal("wrong privilege")
	}
}

func TestPrivilegeCollationForIdentity(t *testing.T) {
	ds, p := setup(t)

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
	privileges, err := ds.readPrivileges(uid, ds.EntityTypes.Workgroup, eid)
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
	privileges, err = ds.readPrivileges(uid, ds.EntityTypes.Workgroup, eid)
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
	privileges, err = ds.readPrivileges(uid, ds.EntityTypes.Workgroup, eid)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(privileges)

	if len(privileges) != 0 {
		t.Fatal("wrong privileges")
	}
}

func TestSecurity(t *testing.T) {
	ds, p := setup(t)

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

	roles, err := ds.ReadRoles(p, 0, 100)
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

	if err := ds.UpdateRole(p, role1Id, "role1", "description1"); err != nil {
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
	if err := ds.LinkRoleAndPermissions(p, role1Id, perms); err != nil {
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

	groups, err := ds.ReadWorkgroups(p, 0, 100)
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

	if err := ds.UpdateWorkgroup(p, group1Id, "group1", "description"); err != nil {
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
		p.WorkgroupId(),
		ds.EntityTypes.Identity,
		user1Id,
	}); err != nil {
		t.Fatal(err)
	}

	users, err := ds.ReadIdentities(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(users)
	if len(users) == 0 {
		t.Fatal("expected > 0 users")
	}
	userpwd, err := ds.readIdentityAndPassword("user1")
	if err != nil {
		t.Fatal(err)
	}
	if userpwd == nil {
		t.Fatal("user not saved correctly")
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
	ds, p := setup(t)

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

func TestExternalClusters(t *testing.T) {
	ds, p := setup(t)

	id1, err := ds.CreateExternalCluster(p, "cluster1", "address1", "started")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id1)

	id2, err := ds.CreateExternalCluster(p, "cluster2", "address2", "started")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id2)

	clusters, err := ds.ReadClusters(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(clusters) != 2 {
		t.Fatal("expected 2 clusters")
	}

	c1, err := ds.ReadCluster(p, id1)
	if err != nil {
		t.Fatal(err)
	}

	if c1.Id != id1 || c1.Name != "cluster1" || c1.Address != "address1" || c1.State != "started" {
		t.Fatal("wrong cluster")
	}

	c2, err := ds.ReadCluster(p, id2)
	if err != nil {
		t.Fatal(err)
	}

	if c2.Id != id2 || c2.Name != "cluster2" || c2.Address != "address2" || c2.State != "started" {
		t.Fatal("wrong cluster")
	}

	if err := ds.DeleteCluster(p, id1); err != nil {
		t.Fatal(err)
	}

	clusters, err = ds.ReadClusters(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(clusters) != 1 {
		t.Fatal("expected 1 cluster")
	}

	if err := ds.DeleteCluster(p, id2); err != nil {
		t.Fatal(err)
	}

	clusters, err = ds.ReadClusters(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(clusters) != 0 {
		t.Fatal("expected 0 cluster")
	}
}

func TestYarnClusters(t *testing.T) {
	ds, p := setup(t)

	eid, err := ds.CreateEngine(p, "engine", "location")
	if err != nil {
		t.Fatal(err)
	}

	id1, err := ds.CreateYarnCluster(p, "cluster1", "address1", "", "started", YarnCluster{
		0,
		eid,
		4,
		"applicationId1",
		"memory1",
		"username1",
		"outputDir1",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id1)

	id2, err := ds.CreateYarnCluster(p, "cluster2", "address2", "", "started", YarnCluster{
		0,
		eid,
		4,
		"applicationId2",
		"memory2",
		"username2",
		"outputDir2",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id2)

	clusters, err := ds.ReadClusters(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(clusters) != 2 {
		t.Fatal("expected 2 clusters")
	}

	c1, err := ds.ReadCluster(p, id1)
	if err != nil {
		t.Fatal(err)
	}

	if c1.Id != id1 || c1.Name != "cluster1" || c1.Address != "address1" || c1.State != "started" {
		t.Fatal("wrong cluster")
	}

	y1, err := ds.ReadYarnCluster(p, c1.Id)
	if err != nil {
		t.Fatal(err)
	}

	if y1.Memory != "memory1" {
		t.Fatal("wrong yarn cluster")
	}

	c2, err := ds.ReadCluster(p, id2)
	if err != nil {
		t.Fatal(err)
	}

	if c2.Id != id2 || c2.Name != "cluster2" || c2.Address != "address2" || c2.State != "started" {
		t.Fatal("wrong cluster")
	}

	y2, err := ds.ReadYarnCluster(p, c2.Id)
	if err != nil {
		t.Fatal(err)
	}

	if y2.Memory != "memory2" {
		t.Fatal("wrong yarn cluster")
	}

	if err := ds.DeleteCluster(p, id1); err != nil {
		t.Fatal(err)
	}

	clusters, err = ds.ReadClusters(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(clusters) != 1 {
		t.Fatal("expected 1 cluster")
	}

	if err := ds.DeleteCluster(p, id2); err != nil {
		t.Fatal(err)
	}

	clusters, err = ds.ReadClusters(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(clusters) != 0 {
		t.Fatal("expected 0 cluster")
	}
}

func TestProjects(t *testing.T) {
	ds, p := setup(t)

	id1, err := ds.CreateProject(p, "project1", "description1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id1)

	id2, err := ds.CreateProject(p, "project2", "description2")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id2)

	projects, err := ds.ReadProjects(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(projects) != 2 {
		t.Fatal("expected 2 projects")
	}

	p1, err := ds.ReadProject(p, id1)
	if err != nil {
		t.Fatal(err)
	}

	if p1.Id != id1 || p1.Name != "project1" || p1.Description != "description1" {
		t.Fatal("wrong project")
	}

	p2, err := ds.ReadProject(p, id2)
	if err != nil {
		t.Fatal(err)
	}

	if p2.Id != id2 || p2.Name != "project2" || p2.Description != "description2" {
		t.Fatal("wrong project")
	}

	if err := ds.DeleteProject(p, id1); err != nil {
		t.Fatal(err)
	}

	projects, err = ds.ReadProjects(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(projects) != 1 {
		t.Fatal("expected 1 project")
	}

	if err := ds.DeleteProject(p, id2); err != nil {
		t.Fatal(err)
	}

	projects, err = ds.ReadProjects(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(projects) != 0 {
		t.Fatal("expected 0 project")
	}
}

func TestModels(t *testing.T) {
	ds, p := setup(t)

	id1, err := ds.CreateModel(p, Model{
		0,
		"model1",
		"cluster1",
		"algo1",
		"dataset1",
		"column1",
		"name1",
		"location1",
		0,
		time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id1)

	id2, err := ds.CreateModel(p, Model{
		0,
		"model2",
		"cluster2",
		"algo2",
		"dataset2",
		"column2",
		"name2",
		"location2",
		0,
		time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id2)

	models, err := ds.ReadModels(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(models) != 2 {
		t.Fatal("expected 2 models")
	}

	m1, err := ds.ReadModel(p, id1)
	if err != nil {
		t.Fatal(err)
	}

	if m1.Id != id1 || m1.Name != "model1" {
		t.Fatal("wrong model")
	}

	m2, err := ds.ReadModel(p, id2)
	if err != nil {
		t.Fatal(err)
	}

	if m2.Id != id2 || m2.Name != "model2" {
		t.Fatal("wrong model")
	}

	if err := ds.DeleteModel(p, id1); err != nil {
		t.Fatal(err)
	}

	models, err = ds.ReadModels(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(models) != 1 {
		t.Fatal("expected 1 model")
	}

	if err := ds.DeleteModel(p, id2); err != nil {
		t.Fatal(err)
	}

	models, err = ds.ReadModels(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(models) != 0 {
		t.Fatal("expected 0 model")
	}
}

func TestProjectModels(t *testing.T) {
	ds, p := setup(t)

	pid, err := ds.CreateProject(p, "project1", "description1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pid)

	mid1, err := ds.CreateModel(p, Model{
		0,
		"model1",
		"cluster1",
		"algo1",
		"dataset1",
		"column1",
		"name1",
		"location1",
		0,
		time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(mid1)

	mid2, err := ds.CreateModel(p, Model{
		0,
		"model2",
		"cluster2",
		"algo2",
		"dataset2",
		"column2",
		"name2",
		"location2",
		0,
		time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(mid2)

	if err := ds.LinkProjectAndModel(p, pid, mid1); err != nil {
		t.Fatal(err)
	}

	if err := ds.LinkProjectAndModel(p, pid, mid2); err != nil {
		t.Fatal(err)
	}

	models, err := ds.ReadModelsForProject(p, pid, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(models) != 2 {
		t.Fatal("expected 2 models")
	}

	if err := ds.UnlinkProjectAndModel(p, pid, mid1); err != nil {
		t.Fatal(err)
	}

	models, err = ds.ReadModelsForProject(p, pid, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(models) != 1 {
		t.Fatal("expected 1 model")
	}

	if err := ds.UnlinkProjectAndModel(p, pid, mid2); err != nil {
		t.Fatal(err)
	}

	models, err = ds.ReadModelsForProject(p, pid, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(models) != 0 {
		t.Fatal("expected 0 models")
	}

}

func TestServices(t *testing.T) {
	ds, p := setup(t)

	mid1, err := ds.CreateModel(p, Model{
		0,
		"model1",
		"cluster1",
		"algo1",
		"dataset1",
		"column1",
		"name1",
		"location1",
		0,
		time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(mid1)

	mid2, err := ds.CreateModel(p, Model{
		0,
		"model2",
		"cluster2",
		"algo2",
		"dataset2",
		"column2",
		"name2",
		"location2",
		0,
		time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(mid2)

	id1, err := ds.CreateService(p, Service{
		0,
		mid1,
		"address1",
		9001,
		1111,
		"started",
		time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id1)

	id2, err := ds.CreateService(p, Service{
		0,
		mid2,
		"address2",
		9002,
		2222,
		"started",
		time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id2)

	services, err := ds.ReadServices(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(services) != 2 {
		t.Fatal("expected 2 services")
	}

	s1, err := ds.ReadService(p, id1)
	if err != nil {
		t.Fatal(err)
	}

	if s1.Id != id1 || s1.ModelId != mid1 {
		t.Fatal("wrong service")
	}

	s2, err := ds.ReadService(p, id2)
	if err != nil {
		t.Fatal(err)
	}

	if s2.Id != id2 || s2.ModelId != mid2 {
		t.Fatal("wrong service")
	}

	if err := ds.DeleteService(p, id1); err != nil {
		t.Fatal(err)
	}

	services, err = ds.ReadServices(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(services) != 1 {
		t.Fatal("expected 1 service")
	}

	if err := ds.DeleteService(p, id2); err != nil {
		t.Fatal(err)
	}

	services, err = ds.ReadServices(p, 0, 100)
	if err != nil {
		t.Fatal(err)
	}

	if len(services) != 0 {
		t.Fatal("expected 0 service")
	}
}
