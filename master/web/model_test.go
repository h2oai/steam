package web

import (
	"testing"

	"github.com/h2oai/steam/master/az"
)

var modelKey = "glm-97176b71-c4dc-4d3a-bf36-53b5a4e4ab03"

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
	}
}
