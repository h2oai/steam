package web

import (
	"os"
	"testing"

	"github.com/h2oai/steam/master/az"
)

type workgroupIn struct {
	name string
	desc string
}

var workgroupTests = []struct {
	in   workgroupIn
	out  int64
	pass bool
	err  error
}{
	{in: workgroupIn{name: "workgroup1", desc: "test workgroup"}, out: 2, pass: true},
}

var readWorkgroupTests = []struct {
	offset uint
	limit  uint
}{
	{0, 10},
}

var updateWorkgroupTests = map[int64][]struct {
	in   workgroupIn
	pass bool
	err  error
}{
	2: {
		{in: workgroupIn{name: "workgroup1_rename", desc: "test workgroup"}, pass: true},
	},
}

func TestSQLiteWorkgroup(t *testing.T) {
	svc, pz, temp := testSetup("workgroup", "sqlite3")
	defer os.RemoveAll(temp)

	t.Logf("Testing %d case(s)", len(workgroupTests))
	// -- C --
	if ok := t.Run("Create", testWorkgroupCreate(pz, svc)); !ok {
		t.FailNow()
	}

	// -- R --
	if ok := t.Run("Read", testWorkgroupRead(pz, svc)); !ok {
		t.FailNow()
	}

	// -- U --
	if ok := t.Run("Update", testWorkgroupUpdate(pz, svc)); !ok {
		t.FailNow()
	}

	// -- D --
	if ok := t.Run("Delete", testWorkgroupDelete(pz, svc)); !ok {
		t.FailNow()
	}
}

func testWorkgroupCreate(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		for _, test := range workgroupTests {
			in, out := test.in, test.out
			id, err := svc.CreateWorkgroup(pz, in.name, in.desc)
			if test.pass {
				if err != nil {
					t.Errorf("Create(%+v): unexpected error creating workgroup: %+v", in, err)
				} else if id != out {
					t.Errorf("Create(%+v): incorrect cluster id: expected %d, got %d", out, out, id)
				}
			} else {
				if err == nil {
					t.Errorf("Create(%+v): expected error creating workgroup", in)
				} else if err.Error() != test.err.Error() {
					t.Errorf("Create(%+v): incorrect error: expected %q, got %q", in, test.err, err)
				}
			}
		}
	}
}

func testWorkgroupRead(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		var totPass uint = 1
		for _, test := range workgroupTests {
			in, out := test.in, test.out
			workgroup, err := svc.GetWorkgroup(pz, out)
			if test.pass {
				if err != nil {
					t.Errorf("Read(%+v): unexpected error reading workgroup: %+v", out, err)
				} else if in.name != workgroup.Name {
					t.Errorf("Read(%+v): incorrect workgroup name: expected %s, got %s", out, in.name, workgroup.Name)
				} else if in.desc != workgroup.Description {
					t.Errorf("Read(%+v): incorrect workgroup description: expected %s, got %s", out, in.desc, workgroup.Description)
				}
				totPass++
			} else {
				if err == nil {
					t.Errorf("Read(%+v): expected error reading workgroup", out)
				}
			}

			workgroup, err = svc.GetWorkgroupByName(pz, in.name)
			if test.pass {
				if err != nil {
					t.Errorf("Read(%+v): unexpected error reading workgroup: %+v", in.name, err)
				} else if out != workgroup.Id {
					t.Errorf("Read(%+v): incorrect workgroup id: expected %s, got %s", in.name, out, workgroup.Id)
				} else if in.desc != workgroup.Description {
					t.Errorf("Read(%+v): incorrect workgroup description: expected %s, got %s", in.name, in.desc, workgroup.Description)
				}
			} else {
				if err == nil {
					t.Errorf("Read(%+v): expected error reading workgroup", in.name)
				}
			}

		}

		for _, get := range readWorkgroupTests {
			var count int
			if totPass-get.offset < get.limit {
				count = int(totPass - get.offset)
			} else {
				count = int(get.limit)
			}
			workgroups, err := svc.GetWorkgroups(pz, get.offset, get.limit)
			if err != nil {
				t.Errorf("Read(%+v): unexpected error reading workgroups: %+v", get, err)
			} else if len(workgroups) != count {
				t.Errorf("Read(%+v): incorrect number of workgroups read: expected %d, got %d", get, count, len(workgroups))
			} else if len(workgroups) > 0 && workgroups[0].Id-1 != int64(get.offset) {
				t.Errorf("Read(%+v): incorrect offset: expected %d, got %d)", get, get.offset, workgroups[0].Id-1)
			}
		}
	}
}

func testWorkgroupUpdate(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		for id, upds := range updateWorkgroupTests {
			for _, test := range upds {
				in := test.in
				err := svc.UpdateWorkgroup(pz, id, in.name, in.desc)
				if test.pass {
					if err != nil {
						t.Errorf("Update(%d:%+v): unexpected error updating workgroup: %+v", id, in, err)
					}
					workgroup, err := svc.GetWorkgroup(pz, id)
					if err != nil {
						t.Errorf("Update(%d:%+v): unexpected error reading workgroup: %+v", id, in, err)
					} else if in.name != workgroup.Name {
						t.Errorf("Update(%d:%+v): incorrect workgroup name, expected %s, got %s", id, in, in.name, workgroup.Name)
					} else if in.desc != workgroup.Description {
						t.Errorf("Update(%d:%+v): incorrect workgroup description, expected %s, got %s", in.desc, workgroup.Description)
					}
				} else {
					if err == nil {
						t.Errorf("Update(%d:%+v): expected error updating workgroup: %+v", id, in, err)
					} else if err.Error() != test.err.Error() {
						t.Errorf("Update(%d:%+v): incorrect error: expected %q, got %q", id, in, test.err, err)
					}
				}
			}
		}
	}
}

func testWorkgroupDelete(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		for _, test := range workgroupTests {
			out := test.out
			err := svc.DeleteWorkgroup(pz, out)
			if test.pass {
				if err != nil {
					t.Errorf("Delete(%+v): unexpected error deleting workgroup: %+v", out, err)
				}
			} else {
				if err == nil {
					t.Errorf("Delete(%+v): expected error deleting workgroup", out)
				}
			}
		}

		workgroups, _ := svc.GetWorkgroups(pz, 0, 2)
		if len(workgroups) > 1 {
			t.Errorf("Delete: at least one workgroup was not deleted")
		}
	}
}
