package web

import (
	"os"
	"testing"

	"github.com/h2oai/steam/master/az"
)

type roleIn struct {
	name string
	desc string
}

var roleTests = []struct {
	in   roleIn
	out  int64
	pass bool
	err  error
}{
	{in: roleIn{name: "role1", desc: "test role"}, out: 2, pass: true},
	{in: roleIn{name: "role2", desc: "test role"}, out: 3, pass: true},
}

var readRoleTests = []struct {
	offset uint
	limit  uint
}{
	{0, 10},
}

var updateRoleTests = map[int64][]struct {
	in   roleIn
	pass bool
	err  error
}{
	2: {
		{in: roleIn{name: "role1_rename", desc: "test role"}, pass: true},
	},
}

func TestSQLiteRole(t *testing.T) {
	svc, pz, temp := testSetup("role", "sqlite3")
	defer os.RemoveAll(temp)

	t.Logf("Testing %d case(s)", len(roleTests))
	// -- C --
	if ok := t.Run("Create", testRoleCreate(pz, svc)); !ok {
		t.FailNow()
	}

	// -- R --
	if ok := t.Run("Read", testRoleRead(pz, svc)); !ok {
		t.FailNow()
	}

	// -- U --
	if ok := t.Run("Update", testRoleUpdate(pz, svc)); !ok {
		t.FailNow()
	}

	// -- D --
	if ok := t.Run("Delete", testRoleDelete(pz, svc)); !ok {
		t.FailNow()
	}
}

func testRoleCreate(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		for _, test := range roleTests {
			in, out := test.in, test.out
			id, err := svc.CreateRole(pz, in.name, in.desc)
			if test.pass {
				if err != nil {
					t.Errorf("Create(%+v): unexpected error creating role: %+v", in, err)
				} else if id != out {
					t.Errorf("Create(%+v): incorrect cluster id: expected %d, got %d", out, out, id)
				}
			} else {
				if err == nil {
					t.Errorf("Create(%+v): expected error creating role", in)
				} else if err.Error() != test.err.Error() {
					t.Errorf("Create(%+v): incorrect error: expected %q, got %q", in, test.err, err)
				}
			}
		}
	}
}

func testRoleRead(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		var totPass uint = 1
		for _, test := range roleTests {
			in, out := test.in, test.out
			role, err := svc.GetRole(pz, out)
			if test.pass {
				if err != nil {
					t.Errorf("Read(%+v): unexpected error reading role: %+v", out, err)
				} else if in.name != role.Name {
					t.Errorf("Read(%+v): incorrect role name: expected %s, got %s", out, in.name, role.Name)
				} else if in.desc != role.Description {
					t.Errorf("Read(%+v): incorrect role description: expected %s, got %s", out, in.desc, role.Description)
				}
				totPass++
			} else {
				if err == nil {
					t.Errorf("Read(%+v): expected error reading role", out)
				}
			}

			role, err = svc.GetRoleByName(pz, in.name)
			if test.pass {
				if err != nil {
					t.Errorf("Read(%+v): unexpected error reading role: %+v", in.name, err)
				} else if out != role.Id {
					t.Errorf("Read(%+v): incorrect role id: expected %s, got %s", in.name, out, role.Id)
				} else if in.desc != role.Description {
					t.Errorf("Read(%+v): incorrect role description: expected %s, got %s", in.name, in.desc, role.Description)
				}
			} else {
				if err == nil {
					t.Errorf("Read(%+v): expected error reading role", in.name)
				}
			}

		}

		for _, get := range readRoleTests {
			var count int
			if totPass-get.offset < get.limit {
				count = int(totPass - get.offset)
			} else {
				count = int(get.limit)
			}
			roles, err := svc.GetRoles(pz, get.offset, get.limit)
			if err != nil {
				t.Errorf("Read(%+v): unexpected error reading roles: %+v", get, err)
			} else if len(roles) != count {
				t.Errorf("Read(%+v): incorrect number of roles read: expected %d, got %d", get, count, len(roles))
			} else if len(roles) > 0 && roles[0].Id-1 != int64(get.offset) {
				t.Errorf("Read(%+v): incorrect offset: expected %d, got %d)", get, get.offset, roles[0].Id-1)
			}
		}
	}
}

func testRoleUpdate(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		for id, upds := range updateRoleTests {
			for _, test := range upds {
				in := test.in
				err := svc.UpdateRole(pz, id, in.name, in.desc)
				if test.pass {
					if err != nil {
						t.Errorf("Update(%d:%+v): unexpected error updating role: %+v", id, in, err)
					}
					role, err := svc.GetRole(pz, id)
					if err != nil {
						t.Errorf("Update(%d:%+v): unexpected error reading role: %+v", id, in, err)
					} else if in.name != role.Name {
						t.Errorf("Update(%d:%+v): incorrect role name, expected %s, got %s", in.name, role.Name)
					} else if in.desc != role.Description {
						t.Errorf("Update(%d:%+v): incorrect role description, expected %s, got %s", in.desc, role.Description)
					}
				} else {
					if err == nil {
						t.Errorf("Update(%d:%+v): expected error updating role: %+v", id, in, err)
					} else if err.Error() != test.err.Error() {
						t.Errorf("Update(%d:%+v): incorrect error: expected %q, got %q", id, in, test.err, err)
					}
				}
			}
		}
	}
}

func testRoleDelete(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		for _, test := range roleTests {
			out := test.out
			err := svc.DeleteRole(pz, out)
			if test.pass {
				if err != nil {
					t.Errorf("Delete(%+v): unexpected error deleting role: %+v", out, err)
				}
			} else {
				if err == nil {
					t.Errorf("Delete(%+v): expected error deleting role", out)
				}
			}
		}

		roles, _ := svc.GetRoles(pz, 0, 2)
		if len(roles) > 1 {
			t.Errorf("Delete: at least one role was not deleted")
		}
	}
}
