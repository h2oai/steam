package web

import "testing"

func TestServiceCRUD(tt *testing.T) {
	t := newTest(tt)

	// var names [...]string{"service1","service2","service3","service4", "service5"}

	// -- Setup --

	projectId, err := t.svc.CreateProject(t.su, "project1", "test project", "")
	t.nil(err)
	clusterId, err := t.svc.RegisterCluster(t.su, ClusterAddress)
	t.nil(err)
	modelId, err := t.svc.ImportModelFromCluster(t.su, clusterId, projectId, h2oModels[0].name, "")
	t.nil(err)

	// -- C --

	serviceId := make([]int64, 5)
	for i := 0; i < 5; i++ {
		var err error
		serviceId[i], err = t.svc.StartService(t.su, modelId, "", "")
		if err != nil {
			t.nil(err)
		}
	}

	// -- R --

	services, err := t.svc.GetServices(t.su, 0, 1000)
	t.nil(err)

	t.log(services)

	service, err := t.svc.GetService(t.su, serviceId[0])

	t.log(service)

	// -- U --

	// -- D --
	for i := 0; i < 5; i++ {
		err := t.svc.StopService(t.su, serviceId[i])
		t.nil(err)
	}
}
