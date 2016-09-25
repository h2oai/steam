/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
