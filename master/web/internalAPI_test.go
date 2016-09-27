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

/*
Heiarchy of flow
================

Projects

Datasources, Clusters

Datasets

Model

Service

*/

const (
	ClusterAddress = "localhost:54321"
)

//
// -- Project --
//

func TestProjectCRUD(tt *testing.T) {
	t := newTest(tt)

	const (
		name1 = "project1"
		desc1 = "test project"
	)

	// -- C --

	id, err := t.svc.CreateProject(t.su, name1, desc1, "")
	t.nil(err)

	// -- R --

	project, err := t.svc.GetProject(t.su, id)
	t.nil(err)

	t.ok(name1 == project.Name, "name")
	t.ok(desc1 == project.Description, "description")

	projects, err := t.svc.GetProjects(t.su, 0, 1000)
	t.nil(err)

	t.ok(len(projects) == 1, "project count")
	t.ok(name1 == projects[0].Name, "multi name")
	t.ok(desc1 == projects[0].Description, "mult description")

	// -- U --

	// -- D --

	err = t.svc.DeleteProject(t.su, id)
	t.nil(err)
}

//
// -- Cluster --
//

func TestExternalClusterCRUD(tt *testing.T) {
	t := newTest(tt)

	// Setup

	// End setup

	// -- C --

	id, err := t.svc.RegisterCluster(t.su, ClusterAddress)
	t.nil(err)

	// -- R --

	cluster, err := t.svc.GetCluster(t.su, id)
	t.nil(err)
	t.log(cluster)

	clusters, err := t.svc.GetClusters(t.su, 0, 1000)
	t.nil(err)
	t.log(clusters)

	// -- U --

	// -- D --

	err = t.svc.UnregisterCluster(t.su, id)
	t.nil(err)
}

//
// -- Datasource --
//

func TestDatasourceCRUD(tt *testing.T) {
	t := newTest(tt)

	const (
		name1 = "datasource1"
		desc1 = "first description"
		path1 = "dummy/path1"
		name2 = "datasource2"
		desc2 = "second description"
		path2 = "dummy/path2"
	)

	// -- C --

	projectId, err := t.svc.CreateProject(t.su, "p1", "d1", "") // This is not being tested here

	id, err := t.svc.CreateDatasource(t.su, projectId, name1, desc1, path1)
	t.nil(err)

	// -- R --

	datasource, err := t.svc.GetDatasource(t.su, id)
	t.nil(err)

	t.ok(name1 == datasource.Name, "name")
	t.ok(desc1 == datasource.Description, "description")
	// UNMARSHALL JSON t.ok(path1 == datasource.Configuration, "configuration")

	datasources, err := t.svc.GetDatasources(t.su, projectId, 0, 1000)
	t.nil(err)

	t.ok(len(datasources) == 1, "datasource count")
	t.ok(name1 == datasources[0].Name, "multi name")
	t.ok(desc1 == datasources[0].Description, "multi description")
	// UNMARSHAL JSON t.ok(path1 == datasources[0].Configuration, "configuration")

	// -- U --

	err = t.svc.UpdateDatasource(t.su, id, name2, desc2, path2)
	t.nil(err)

	datasource, err = t.svc.GetDatasource(t.su, id)
	t.nil(err)

	t.ok(name2 == datasource.Name, "updated name")
	t.ok(desc2 == datasource.Description, "updated description")
	// UNMARSHAL JSON t.ok(path2 == datasource.Configuration, "updated configuration")

	// -- D --

	err = t.svc.DeleteDatasource(t.su, id)
	t.nil(err)
}

//
// -- Dataset --
//

//
// -- Model --
//

// FIXME: Sebastian: this wasn't compiling

// func TestExternalModelCRUD(tt *testing.T) {
// 	t := newTest(tt)

// 	// -- Setup --

// 	clusterId, _ := t.svc.RegisterCluster(t.su, ClusterAddress)
// 	projectId, _ := t.svc.CreateProject(t.su, "p1", "d1")

// 	// -- End Setup --

// 	// -- C --

// 	id, err := t.svc.ImportModelFromCluster(t.su, clusterId, projectId, "modelName", "modelName")
// 	t.nil(err)

// }

// //
// // -- Service --
// //

// func TestScoringServicesCRUD(tt *testing.T) {
// 	t := newTest(tt)

// 	// -- Setup --

// 	// -- End Setup --

// 	// -- C --

// 	// id, err := t.svc.StartService(t.su, modelId, port)
// }
