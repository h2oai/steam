package web

import "testing"

func TestModelCRUD(tt *testing.T) {
	t := newTest(tt)

	const (
		name = "model1"
	)

	// -- Setup --

	projectId, err := t.svc.CreateProject(t.su, "project1", "test project", "")
	t.nil(err)
	clusterId, err := t.svc.RegisterCluster(t.su, ClusterAddress)
	t.nil(err)

	// -- C --

	modelId, err := t.svc.ImportModelFromCluster(t.su, clusterId, projectId, h2oModelKey, name)
	t.nil(err)

	// -- R --
	model, err := t.svc.GetModel(t.su, modelId)
	t.nil(err)

	t.ok(name == model.Name, "GetModel: Name: expected %s got %s", name, model.Name)
	t.ok(h2oModelKey == model.ModelKey, "GetModel: ModelKey: expected %s got %s", h2oModelKey, model.ModelKey)

	// VERIFY MODEL INFORMATION FROM H2O

	models, err := t.svc.GetModels(t.su, projectId, 0, 10000)
	t.nil(err)

	t.ok(name == models[0].Name, "GetModel: Name: expected %s got %s", name, models[0].Name)
	t.ok(h2oModelKey == models[0].ModelKey, "GetModel: ModelKey: expected %s got %s", h2oModelKey, models[0].ModelKey)

	// -- U --

	// -- D --
	err := t.svc.DeleteModel(t.su, modelId)
	t.nil(err)
}
