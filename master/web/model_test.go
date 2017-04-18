package web

import (
	"os"
	"testing"

	"github.com/h2oai/steam/master/az"
)

type modelIn struct {
	key  string
	name string
}

var modelTests = []struct {
	in   modelIn
	out  int64
	pass bool
	err  error
}{
	{in: modelIn{key: "bin_gbm", name: "binary gbm"}, out: 1, pass: true},
	{in: modelIn{key: "bin_glm", name: "binary glm"}, out: 2, pass: true},
	{in: modelIn{key: "bin_dpl", name: "binary deeplearning"}, out: 3, pass: true},
	{in: modelIn{key: "mul_gbm", name: "multinomial gbm"}, out: 4, pass: true},
	{in: modelIn{key: "mul_dpl", name: "multinomial deeplearning"}, out: 5, pass: true},
	{in: modelIn{key: "reg_gbm", name: "regression gbm"}, out: 6, pass: true},
	{in: modelIn{key: "reg_glm", name: "regression glm"}, out: 7, pass: true},
	{in: modelIn{key: "reg_dpl", name: "regression deeplearning"}, out: 8, pass: true},
}

var readModelTests = []struct {
	offset uint
	limit  uint
}{
	{0, 10},
	{1, 5},
	{3, 9},
}

func TestSQLiteModel(t *testing.T) {
	svc, pz, temp := testSetup("model", "sqlite3")
	defer os.RemoveAll(temp)

	if !test_h2o {
		t.Skip("skipping model tests: requires h2o")
	}
	projectId, err := svc.CreateProject(pz, "proj1", "desc", "")
	if err != nil {
		t.Fatalf("Setup: creating project: %+v", err)
	}
	clusterId, err := svc.RegisterCluster(pz, test_h2oAddress)
	if err != nil {
		t.Fatalf("Setup: creating cluster: %+v", err)
	}

	t.Logf("Testing %d case(s)", len(modelTests))
	// -- C --
	if ok := t.Run("Create", testModelCreate(pz, svc, clusterId, projectId)); !ok {
		t.FailNow()
	}

	// -- R --
	if ok := t.Run("Read", testModelRead(pz, svc, projectId)); !ok {
		t.FailNow()
	}

	// -- U --

	// -- D --
	if ok := t.Run("Delete", testModelDelete(pz, svc, projectId)); !ok {
		t.FailNow()
	}
}

func testModelCreate(pz az.Principal, svc *Service, clusterId, projectId int64) func(t *testing.T) {
	return func(t *testing.T) {
		for _, test := range modelTests {
			in, out := test.in, test.out
			id, err := svc.ImportModelFromCluster(pz, clusterId, projectId, in.key, in.name)
			if test.pass {
				if err != nil {
					t.Errorf("Create(%+v): unexpected error creating model: %+v", in, err)
				} else if id != out {
					t.Errorf("Create(%+v): incorrect cluster id: expected %d, got %d", out, out, id)
				}
			} else {
				if err == nil {
					t.Errorf("Create(%+v): expected error creating model", in)
				} else if err.Error() != test.err.Error() {
					t.Errorf("Create(%+v): incorrect error: expected %q, got %q", in, test.err, err)
				}
			}
		}
	}
}

func testModelRead(pz az.Principal, svc *Service, projectId int64) func(t *testing.T) {
	return func(t *testing.T) {
		var totPass uint
		for _, test := range modelTests {
			in, out := test.in, test.out
			model, err := svc.GetModel(pz, out)
			if test.pass {
				if err != nil {
					t.Errorf("Read(%+v): unexpected error reading model: %+v", out, err)
				} else if in.key != model.ModelKey {
					t.Errorf("Read(%+v): incorrect model key, expected %s, got %s", out, in.key, model.ModelKey)
				} else if in.name != model.Name {
					t.Errorf("Read(%+v): incorrect model name, expected %s, got %s", out, in.name, model.Name)
				}
				totPass++
			} else {
				if err == nil {
					t.Errorf("Read(%+v): expected error reading model", out)
				}
			}
		}

		for _, get := range readModelTests {
			var count int
			if totPass-get.offset < get.limit {
				count = int(totPass - get.offset)
			} else {
				count = int(get.limit)
			}
			models, err := svc.GetModels(pz, projectId, get.offset, get.limit)
			if err != nil {
				t.Errorf("Read(%+v): unexpected error reading models: %+v", get, err)
			} else if len(models) != count {
				t.Errorf("Read(%+v): incorrect number of models read: expected %d, got %d", get, count, len(models))
			} else if len(models) > 0 && models[0].Id-1 != int64(get.offset) {
				t.Errorf("Read(%+v): incorrect offset: expected %d, got %d)", get, get.offset, models[0].Id-1)
			}
		}
	}
}

func testModelDelete(pz az.Principal, svc *Service, projectId int64) func(t *testing.T) {
	return func(t *testing.T) {
		for _, test := range modelTests {
			out := test.out
			err := svc.DeleteModel(pz, out)
			if test.pass {
				if err != nil {
					t.Errorf("Delete(%+v): unexpected error deleting model: %+v", out, err)
				}
			} else {
				if err == nil {
					t.Errorf("Delete(%+v): expected error deleting model", out)
				}
			}
		}

		models, _ := svc.GetModels(pz, projectId, 0, 1)
		if len(models) > 0 {
			t.Errorf("Delete: at least one model was not deleted")

		}
	}
}
