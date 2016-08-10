package web

import "testing"

// TODO: This function is dependent upon a more robust h2o script running, this
// will need to have multiple models of all types running in the cluster.
func TestModelCRUD(tt *testing.T) {
	t := newTest(tt)

	// -- Setup --

	projectId, err := t.svc.CreateProject(t.su, "project1", "test project", "")
	t.nil(err)
	clusterId, err := t.svc.RegisterCluster(t.su, ClusterAddress)
	t.nil(err)

	//
	// -- C --
	//

	modelId := make([]int64, len(h2oModels))
	for i, model := range h2oModels {
		var err error
		// If no name is supplied, model should inherit name from H2O key
		modelId[i], err = t.svc.ImportModelFromCluster(t.su, clusterId, projectId, model.name, "")
		t.nil(err)
	}

	//
	// -- R --
	//

	// model, err := t.svc.GetModel(t.su, modelId)
	// t.nil(err)

	// t.ok(name == model.Name, "GetModel: Name: expected %s got %s", name, model.Name)
	// t.ok(h2oModelKey == model.ModelKey, "GetModel: ModelKey: expected %s got %s", h2oModelKey, model.ModelKey)

	// TODO VERIFY MODEL INFORMATION FROM H2O

	// TODO: Deprecated?
	models, err := t.svc.GetModels(t.su, projectId, 0, 10000)
	t.nil(err)
	t.log(models)

	// t.ok(name == models[0].Name, "GetModel: Name: expected %s got %s", name, models[0].Name)
	// t.ok(h2oModelKey == models[0].ModelKey, "GetModel: ModelKey: expected %s got %s", h2oModelKey, models[0].ModelKey)

	// TODO MORE TESTING HERE DEPENDENT ON H2O SCRIPT
	binModels, err := t.svc.FindModelsBinomial(t.su, projectId, "", "", true, 0, 1000)
	t.nil(err)
	t.log(binModels)

	mulModels, err := t.svc.FindModelsMultinomial(t.su, projectId, "", "", true, 0, 1000)
	t.nil(err)
	t.log(mulModels)

	regModels, err := t.svc.FindModelsRegression(t.su, projectId, "", "", true, 0, 1000)
	t.nil(err)
	t.log(regModels)

	//
	// -- U --
	//

	updatedName := "NewName"

	err = t.svc.UpdateModel(t.su, models[0].Id, updatedName)
	t.nil(err)
	model, err := t.svc.GetModel(t.su, models[0].Id)
	t.nil(err)

	t.ok(model.Name == updatedName, "Updated model name")

	//
	// -- D --
	//

	for _, id := range modelId {
		err := t.svc.DeleteModel(t.su, id)
		t.nil(err)
	}
}
