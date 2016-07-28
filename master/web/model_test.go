package web

import "testing"

func TestImportModelFromCluster(tt *testing.T) {
	t := newTest(tt)

	projectId, err := t.svc.CreateProject(t.su, "project1", "test project", "")
	t.nil(err)
	clusterId, err := t.svc.RegisterCluster(t.su, ClusterAddress)
	t.nil(err)

	modelId, err := t.svc.ImportModelFromCluster(t.su, clusterId, projectId, h2oModelKey, "model1")

	// VERIFY INFORMATION HERE
	t.log(modelId)
}
