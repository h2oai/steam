package web

import (
	"errors"
	"os"
	"testing"

	"github.com/h2oai/steam/master/az"
)

type projectIn struct {
	name   string
	desc   string
	modCat string
}

var projectTests = []struct {
	in   projectIn
	out  int64
	pass bool
	err  error
}{
	// Basic Implementation
	{in: projectIn{name: "project1", desc: "test project"}, out: 1, pass: true},
	// Duplicate description
	{in: projectIn{name: "project2", desc: "test project"}, out: 2, pass: true},
	// No description
	{in: projectIn{name: "project3"}, out: 3, pass: true},
	// With model category
	{in: projectIn{name: "project4", modCat: "Binomial"}, out: 4, pass: true},
	// Duplicate name
	{in: projectIn{name: "project1", desc: "test project"}, pass: false,
		err: errors.New("creating project in database: committing transaction: executing query: UNIQUE constraint failed: project.name"),
	},
}

var readProjectsTests = []struct {
	offset uint
	limit  uint
}{
	{0, 10},
	{2, 2},
	{1, 5},
}

func TestSQLiteProject(t *testing.T) {
	svc, pz, temp := testSetup("project", "sqlite3")
	defer os.RemoveAll(temp)

	t.Logf("Testing %d case(s)", len(projectTests))
	// -- C --
	if ok := t.Run("Create", testProjectCreate(pz, svc)); !ok {
		t.FailNow()
	}

	// -- R --
	if ok := t.Run("Read", testProjectRead(pz, svc)); !ok {
		t.FailNow()
	}

	// -- U --

	//FIXME: TODO

	// -- D --
	if ok := t.Run("Delete", testProjectDelete(pz, svc)); !ok {
		t.FailNow()
	}
}

func testProjectCreate(pz az.Principal, svc *Service) func(*testing.T) {
	return func(t *testing.T) {
		for _, test := range projectTests {
			in, out := test.in, test.out
			id, err := svc.CreateProject(pz, in.name, in.desc, in.modCat)
			if test.pass {
				if err != nil {
					t.Errorf("Create(%+v): unexpected error creating project: %+v", in, err)
				} else if id != out {
					t.Errorf("Create(%+v): incorrect project id: expected %d, got %d", in, out, id)
				}
			} else {
				if err == nil {
					t.Errorf("Create(%+v): expected error", in)
				} else if err.Error() != test.err.Error() {
					t.Errorf("Create(%+v): incorrect error: expected %q, got %q", in, test.err, err)
				}
			}
		}
	}
}

func testProjectRead(pz az.Principal, svc *Service) func(*testing.T) {
	return func(t *testing.T) {
		var totPass uint
		for _, test := range projectTests {
			in, out := test.in, test.out
			project, err := svc.GetProject(pz, out)
			if test.pass {
				if err != nil {
					t.Errorf("Read(%+v): unexpected error reading project: %+v", out, err)
				} else if project.Name != in.name {
					t.Errorf("Read(%+v): incorrect name: expected %s, got %s", out, in.name, project.Name)
				} else if project.Description != in.desc {
					t.Errorf("Read(%+v): incorrect description: expected %s, got %s", out, in.desc, project.Description)
				} else if project.ModelCategory != in.modCat {
					t.Errorf("Read(%+v): incorrect model category: expected %s, got %s", out, in.modCat, project.ModelCategory)
				}
				totPass++
			} else {
				if err == nil {
					t.Errorf("Read(%+v): expected error", out)
				}
			}
		}

		for _, get := range readProjectsTests {
			var count int
			if totPass-get.offset < get.limit {
				count = int(totPass - get.offset)
			} else {
				count = int(get.limit)
			}
			projects, err := svc.GetProjects(pz, get.offset, get.limit)
			if err != nil {
				t.Errorf("Reads(%+v): unexpected error reading projects: %+v", get, err)
			} else if len(projects) != count {
				t.Errorf("Reads(%+v): incorrect number of projects read: expected %d, got %d", get, count, len(projects))
			} else if len(projects) > 1 && projects[0].Id-1 != int64(get.offset) {
				t.Errorf("Reads(%+v): incorrect offset: expected %d, got %d", get, get.offset, projects[0].Id-1)
			}
		}
	}
}

func testProjectDelete(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		for _, test := range projectTests {
			out := test.out
			err := svc.DeleteProject(pz, out)
			if test.pass {
				if err != nil {
					t.Errorf("Delete(%+v): unexpected error deleting project: %+v", out, err)
				}
			} else {
				if err == nil {
					t.Errorf("Delete(%+v): expected error deleting project", out)
				}
			}
		}

		projects, _ := svc.GetProjects(pz, 0, 1)
		if len(projects) > 0 {
			t.Error("Delete: at least one project was not deleted")
		}
	}
}
