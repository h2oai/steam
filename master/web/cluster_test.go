package web

import "testing"

func TestGetClusterDatasets(tt *testing.T) {
	t := newTest(tt)

	id, err := t.svc.RegisterCluster(t.su, ClusterAddress)
	t.nil(err)

	datasets, err := t.svc.GetDatasetsFromCluster(t.su, id)
	t.nil(err)

	// VERIFY INFORMATION IN DATASETS HERE (INTEGRATION)
	t.log("Datasets", datasets)
}

func TestGetClusterModels(tt *testing.T) {
	t := newTest(tt)

	id, err := t.svc.RegisterCluster(t.su, ClusterAddress)
	t.nil(err)

	models, err := t.svc.GetModelsFromCluster(t.su, id, h2oFrameKey)
	t.nil(err)

	// VERIFY INFORMATION HERE
	t.log("Models", models)
}
