package web

import (
	"testing"

	"github.com/h2oai/steam/master/az"
)

var modelKey = "bin_gbm"

func testImportModelCRUD(pz az.Principal, svc *Service, clusterId int64) func(t *testing.T) {
	return func(t *testing.T) {
		projectId, err := svc.CreateProject(pz, "proj1", "", "binomial")
		if err != nil {
			t.Errorf("creating project: %v", err)
			t.FailNow()
		}

		modelId, err := svc.ImportModelFromCluster(pz, clusterId, projectId, modelKey, modelKey)
		if err != nil {
			t.Errorf("importing model from cluster: %v", err)
			t.FailNow()
		}
		t.Logf("Imported model %q with ID %d", modelKey, modelId)

		models, err := svc.GetModels(pz, modelId, 0, 1000)
		if err != nil {
			t.Errorf("reading models from cluster: %v", err)
			t.FailNow()
		}
		t.Logf("Returned model: %+v", models[0].Name)

		model, err := svc.GetModel(pz, modelId)
		if err != nil {
			t.Errorf("reading model from cluster: %v", err)
			t.FailNow()
		}
		t.Logf("Returned model: %+v", model.Name)

		count, err := svc.FindModelsCount(pz, projectId)
		if err != nil {
			t.Errorf("geting project model count: %v", err)
			t.FailNow()
		}
		t.Log("Returned model count:", count)

		models2, err := svc.FindModelsBinomial(pz, projectId, "bin", "mse", true, 0, 1000)
		if err != nil {
			t.Errorf("reading models from cluster: %v", err)
			t.FailNow()
		}
		t.Logf("Returned model: %+v", models2[0].Name)

		if err := svc.DeleteModel(pz, modelId); err != nil {
			t.Errorf("deleting model from clsuter %v", err)
			t.FailNow()
		}
		t.Log("Deleted model with ID:", modelId)
	}
}
