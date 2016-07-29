package web

import "testing"

// TODO: This function is dependent upon a more robust h2o script running, this
// will need to have multiple models of all types running in the cluster.
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
	// TODO Import at least one of each model

	// -- R --
	model, err := t.svc.GetModel(t.su, modelId)
	t.nil(err)

	t.ok(name == model.Name, "GetModel: Name: expected %s got %s", name, model.Name)
	t.ok(h2oModelKey == model.ModelKey, "GetModel: ModelKey: expected %s got %s", h2oModelKey, model.ModelKey)

	// TODO VERIFY MODEL INFORMATION FROM H2O

	// TODO: Deprecated?
	models, err := t.svc.GetModels(t.su, projectId, 0, 10000)
	t.nil(err)

	t.ok(name == models[0].Name, "GetModel: Name: expected %s got %s", name, models[0].Name)
	t.ok(h2oModelKey == models[0].ModelKey, "GetModel: ModelKey: expected %s got %s", h2oModelKey, models[0].ModelKey)

	// TODO MORE TESTING HERE DEPENDENT ON H2O SCRIPT
	binModels, err := t.svc.FindModelsBinomial(t.su, projectId, "", "", true, 0, 1000)
	t.nil(err)
	t.log(binModels)

	// binModel, err := t.svc.GetModelBinomial(t.su, binModelId)
	// t.nil(err)
	// t.log(binModel)

	mulModels, err := t.svc.FindModelsMultinomial(t.su, projectId, "", "", true, 0, 1000)
	t.nil(err)
	t.log(mulModels)

	// mulModel, err := t.svc.GetModelMultinomial(t.su, mulModelId)
	// t.nil(err)
	// t.log(mulModel)

	regModels, err := t.svc.FindModelsRegression(t.su, projectId, "", "", true, 0, 1000)
	t.nil(err)
	t.log(regModels)

	// regModel, err := t.svc.GetModelRegression(t.su, regModelId)
	// t.nil(err)
	// t.log(regModel)

	// -- U --

	// -- D --
	err = t.svc.DeleteModel(t.su, modelId)
	t.nil(err)
}
